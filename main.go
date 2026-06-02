// rwa-cc: Hyperledger Fabric chaincode mirroring the EVM contracts from the
// capstone (DepositReceiptToken, RewardAssetAMM, RWACreditLine) for traditional
// banks evaluating a non-EVM permissioned deployment path.
//
// World-state keys:
//   POS_<addr>       Position    deposit-receipt position
//   POOL             Pool        single AMM pool (rewardReserve, usdcReserve)
//   LINE_<addr>      Line        credit line state
//   DRAW_<drawId>    PendingDraw HITL-gated pending draw
//   EVT_<eventId>    string("1") idempotency marker for Galileo events
//
// Identity / authorization: the chaincode trusts the orchestration MSP via
// the standard Fabric x509 client identity. ORCHESTRATOR_OU is the gating OU.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const ORCHESTRATOR_OU = "orchestration"
const HITL_THRESHOLD_USDC = 50_000_000 // 6 decimals
const HITL_WINDOW_SECONDS = 24 * 3600

// ---------------------------- state types ----------------------------

type Position struct {
	Owner         string `json:"owner"`
	Principal     uint64 `json:"principal"`
	AccruedYield  uint64 `json:"accruedYield"`
	LastAccrualTs int64  `json:"lastAccrualTs"`
	APYBps        uint32 `json:"apyBps"`
}

type Pool struct {
	ReserveReward uint64 `json:"reserveReward"`
	ReserveUSDC   uint64 `json:"reserveUsdc"`
}

type Line struct {
	Owner            string `json:"owner"`
	Limit            uint64 `json:"limit"`
	Outstanding      uint64 `json:"outstanding"`
	APRBps           uint32 `json:"aprBps"`
	LastAccrualTs    int64  `json:"lastAccrualTs"`
	Active           bool   `json:"active"`
	ModelVersionHash string `json:"modelVersionHash"`
}

type PendingDraw struct {
	Owner     string `json:"owner"`
	Amount    uint64 `json:"amount"`
	UnlockTs  int64  `json:"unlockTs"`
	Cancelled bool   `json:"cancelled"`
	Executed  bool   `json:"executed"`
}

// ---------------------------- contract ----------------------------

type RWAContract struct {
	contractapi.Contract
}

func (c *RWAContract) Init(ctx contractapi.TransactionContextInterface) error {
	exists, err := ctx.GetStub().GetState("POOL")
	if err != nil {
		return err
	}
	if exists == nil {
		empty, _ := json.Marshal(Pool{})
		return ctx.GetStub().PutState("POOL", empty)
	}
	return nil
}

// MintDepositReceipt issues yield-bearing receipt tokens 1:1 on a Galileo
// direct-deposit webhook. ref: DepositReceiptToken.sol::mintFor
func (c *RWAContract) MintDepositReceipt(
	ctx contractapi.TransactionContextInterface,
	owner string, amount uint64, eventId string, apyBps uint32,
) error {
	if err := requireOrchestrator(ctx); err != nil {
		return err
	}
	if amount == 0 {
		return errors.New("zero amount")
	}
	if seen, _ := ctx.GetStub().GetState("EVT_" + eventId); seen != nil {
		return fmt.Errorf("event %s already processed", eventId)
	}
	if err := ctx.GetStub().PutState("EVT_"+eventId, []byte("1")); err != nil {
		return err
	}
	pos, err := c.getPosition(ctx, owner)
	if err != nil {
		return err
	}
	c.accrue(&pos)
	pos.Principal += amount
	pos.APYBps = apyBps
	pos.LastAccrualTs = nowSeconds()
	return c.savePosition(ctx, pos)
}

// RedeemDepositReceipt burns receipt tokens (principal + yield) and emits an
// event the orchestration layer translates into a USDC transfer.
func (c *RWAContract) RedeemDepositReceipt(
	ctx contractapi.TransactionContextInterface,
	owner string, amount uint64,
) error {
	if err := requireOrchestrator(ctx); err != nil {
		return err
	}
	pos, err := c.getPosition(ctx, owner)
	if err != nil {
		return err
	}
	c.accrue(&pos)
	available := pos.Principal + pos.AccruedYield
	if amount > available {
		return fmt.Errorf("insufficient balance: %d > %d", amount, available)
	}
	if amount <= pos.AccruedYield {
		pos.AccruedYield -= amount
	} else {
		fromPrincipal := amount - pos.AccruedYield
		pos.AccruedYield = 0
		pos.Principal -= fromPrincipal
	}
	if err := c.savePosition(ctx, pos); err != nil {
		return err
	}
	return ctx.GetStub().SetEvent("YieldRedeemed", []byte(fmt.Sprintf(`{"owner":"%s","amount":%d}`, owner, amount)))
}

// Underwrite is called by the underwriter MSP (compliance officer) after the
// XGBoost credit model has produced SHAP-explained, ECOA-compliant decisions.
func (c *RWAContract) Underwrite(
	ctx contractapi.TransactionContextInterface,
	owner string, limit uint64, aprBps uint32, modelVersionHash string,
) error {
	if err := requireOU(ctx, "underwriter"); err != nil {
		return err
	}
	line, _ := c.getLine(ctx, owner)
	line.Owner = owner
	line.Limit = limit
	line.APRBps = aprBps
	line.Active = limit > 0
	line.ModelVersionHash = modelVersionHash
	if line.LastAccrualTs == 0 {
		line.LastAccrualTs = nowSeconds()
	}
	return c.saveLine(ctx, line)
}

// RequestDraw mirrors RWACreditLine.requestDraw — small amounts execute
// immediately, large amounts enter the HITL queue.
func (c *RWAContract) RequestDraw(
	ctx contractapi.TransactionContextInterface,
	owner string, amount uint64, drawId string,
) (bool, error) {
	if err := requireOrchestrator(ctx); err != nil {
		return false, err
	}
	line, err := c.getLine(ctx, owner)
	if err != nil {
		return false, err
	}
	if !line.Active {
		return false, errors.New("line inactive")
	}
	c.accrueLine(&line)
	if amount > (line.Limit - line.Outstanding) {
		return false, errors.New("limit exceeded")
	}
	if amount < HITL_THRESHOLD_USDC {
		line.Outstanding += amount
		if err := c.saveLine(ctx, line); err != nil {
			return false, err
		}
		_ = ctx.GetStub().SetEvent("DrawExecuted",
			[]byte(fmt.Sprintf(`{"owner":"%s","drawId":"%s","amount":%d}`, owner, drawId, amount)))
		return true, nil
	}
	d := PendingDraw{
		Owner:    owner,
		Amount:   amount,
		UnlockTs: nowSeconds() + HITL_WINDOW_SECONDS,
	}
	if err := c.savePendingDraw(ctx, drawId, d); err != nil {
		return false, err
	}
	_ = ctx.GetStub().SetEvent("DrawRequested",
		[]byte(fmt.Sprintf(`{"owner":"%s","drawId":"%s","amount":%d,"unlockTs":%d}`, owner, drawId, amount, d.UnlockTs)))
	return false, nil
}

func (c *RWAContract) CancelDraw(ctx contractapi.TransactionContextInterface, drawId string) error {
	d, err := c.getPendingDraw(ctx, drawId)
	if err != nil {
		return err
	}
	if d.Cancelled || d.Executed {
		return errors.New("draw already resolved")
	}
	d.Cancelled = true
	if err := c.savePendingDraw(ctx, drawId, d); err != nil {
		return err
	}
	return ctx.GetStub().SetEvent("DrawCancelled", []byte(fmt.Sprintf(`{"drawId":"%s"}`, drawId)))
}

func (c *RWAContract) FinalizeDraw(ctx contractapi.TransactionContextInterface, drawId string) error {
	if err := requireOrchestrator(ctx); err != nil {
		return err
	}
	d, err := c.getPendingDraw(ctx, drawId)
	if err != nil {
		return err
	}
	if d.Cancelled || d.Executed {
		return errors.New("draw already resolved")
	}
	if nowSeconds() < d.UnlockTs {
		return errors.New("draw not ready")
	}
	d.Executed = true
	if err := c.savePendingDraw(ctx, drawId, d); err != nil {
		return err
	}
	line, _ := c.getLine(ctx, d.Owner)
	line.Outstanding += d.Amount
	if err := c.saveLine(ctx, line); err != nil {
		return err
	}
	return ctx.GetStub().SetEvent("DrawExecuted",
		[]byte(fmt.Sprintf(`{"owner":"%s","drawId":"%s","amount":%d}`, d.Owner, drawId, d.Amount)))
}

// Repay decreases outstanding balance — typically driven by a direct-deposit
// webhook in the orchestration layer.
func (c *RWAContract) Repay(ctx contractapi.TransactionContextInterface, owner string, amount uint64) error {
	if err := requireOrchestrator(ctx); err != nil {
		return err
	}
	line, err := c.getLine(ctx, owner)
	if err != nil {
		return err
	}
	if line.Outstanding == 0 {
		return errors.New("nothing to repay")
	}
	c.accrueLine(&line)
	paid := amount
	if paid > line.Outstanding {
		paid = line.Outstanding
	}
	line.Outstanding -= paid
	if err := c.saveLine(ctx, line); err != nil {
		return err
	}
	return ctx.GetStub().SetEvent("Repaid",
		[]byte(fmt.Sprintf(`{"owner":"%s","amount":%d}`, owner, paid)))
}

// ---------- AMM (simple x*y=k) ----------

func (c *RWAContract) AddLiquidity(ctx contractapi.TransactionContextInterface, rewardIn, usdcIn uint64) error {
	if err := requireOrchestrator(ctx); err != nil {
		return err
	}
	pool, err := c.getPool(ctx)
	if err != nil {
		return err
	}
	pool.ReserveReward += rewardIn
	pool.ReserveUSDC += usdcIn
	return c.savePool(ctx, pool)
}

func (c *RWAContract) SwapRewardForUSDC(
	ctx contractapi.TransactionContextInterface,
	rewardIn uint64, minOut uint64,
) (uint64, error) {
	pool, err := c.getPool(ctx)
	if err != nil {
		return 0, err
	}
	if pool.ReserveReward == 0 || pool.ReserveUSDC == 0 {
		return 0, errors.New("insufficient liquidity")
	}
	grossOut := (rewardIn * pool.ReserveUSDC) / (pool.ReserveReward + rewardIn)
	fee := (grossOut * 30) / 10_000
	out := grossOut - fee
	if out < minOut {
		return 0, fmt.Errorf("slippage: %d < %d", out, minOut)
	}
	pool.ReserveReward += rewardIn
	pool.ReserveUSDC -= (out + fee)
	if err := c.savePool(ctx, pool); err != nil {
		return 0, err
	}
	_ = ctx.GetStub().SetEvent("Swapped",
		[]byte(fmt.Sprintf(`{"rewardIn":%d,"usdcOut":%d,"fee":%d}`, rewardIn, out, fee)))
	return out, nil
}

// ---------- read helpers ----------

func (c *RWAContract) GetPosition(ctx contractapi.TransactionContextInterface, owner string) (Position, error) {
	return c.getPosition(ctx, owner)
}
func (c *RWAContract) GetLine(ctx contractapi.TransactionContextInterface, owner string) (Line, error) {
	return c.getLine(ctx, owner)
}
func (c *RWAContract) GetPool(ctx contractapi.TransactionContextInterface) (Pool, error) {
	return c.getPool(ctx)
}

// ---------- internals ----------

func (c *RWAContract) accrue(p *Position) {
	if p.Principal == 0 || p.LastAccrualTs == 0 {
		p.LastAccrualTs = nowSeconds()
		return
	}
	elapsed := uint64(nowSeconds() - p.LastAccrualTs)
	if elapsed == 0 {
		return
	}
	earned := (p.Principal * uint64(p.APYBps) * elapsed) / (10_000 * 365 * 24 * 3600)
	p.AccruedYield += earned
	p.LastAccrualTs = nowSeconds()
}
func (c *RWAContract) accrueLine(l *Line) {
	if l.Outstanding == 0 || l.LastAccrualTs == 0 {
		l.LastAccrualTs = nowSeconds()
		return
	}
	elapsed := uint64(nowSeconds() - l.LastAccrualTs)
	if elapsed == 0 {
		return
	}
	interest := (l.Outstanding * uint64(l.APRBps) * elapsed) / (10_000 * 365 * 24 * 3600)
	l.Outstanding += interest
	l.LastAccrualTs = nowSeconds()
}

func (c *RWAContract) getPosition(ctx contractapi.TransactionContextInterface, owner string) (Position, error) {
	raw, err := ctx.GetStub().GetState("POS_" + owner)
	if err != nil {
		return Position{}, err
	}
	if raw == nil {
		return Position{Owner: owner}, nil
	}
	var p Position
	if err := json.Unmarshal(raw, &p); err != nil {
		return Position{}, err
	}
	return p, nil
}
func (c *RWAContract) savePosition(ctx contractapi.TransactionContextInterface, p Position) error {
	b, _ := json.Marshal(p)
	return ctx.GetStub().PutState("POS_"+p.Owner, b)
}
func (c *RWAContract) getLine(ctx contractapi.TransactionContextInterface, owner string) (Line, error) {
	raw, err := ctx.GetStub().GetState("LINE_" + owner)
	if err != nil {
		return Line{}, err
	}
	if raw == nil {
		return Line{Owner: owner}, nil
	}
	var l Line
	if err := json.Unmarshal(raw, &l); err != nil {
		return Line{}, err
	}
	return l, nil
}
func (c *RWAContract) saveLine(ctx contractapi.TransactionContextInterface, l Line) error {
	b, _ := json.Marshal(l)
	return ctx.GetStub().PutState("LINE_"+l.Owner, b)
}
func (c *RWAContract) getPendingDraw(ctx contractapi.TransactionContextInterface, id string) (PendingDraw, error) {
	raw, err := ctx.GetStub().GetState("DRAW_" + id)
	if err != nil {
		return PendingDraw{}, err
	}
	if raw == nil {
		return PendingDraw{}, fmt.Errorf("draw %s not found", id)
	}
	var d PendingDraw
	if err := json.Unmarshal(raw, &d); err != nil {
		return PendingDraw{}, err
	}
	return d, nil
}
func (c *RWAContract) savePendingDraw(ctx contractapi.TransactionContextInterface, id string, d PendingDraw) error {
	b, _ := json.Marshal(d)
	return ctx.GetStub().PutState("DRAW_"+id, b)
}
func (c *RWAContract) getPool(ctx contractapi.TransactionContextInterface) (Pool, error) {
	raw, err := ctx.GetStub().GetState("POOL")
	if err != nil {
		return Pool{}, err
	}
	if raw == nil {
		return Pool{}, nil
	}
	var p Pool
	if err := json.Unmarshal(raw, &p); err != nil {
		return Pool{}, err
	}
	return p, nil
}
func (c *RWAContract) savePool(ctx contractapi.TransactionContextInterface, p Pool) error {
	b, _ := json.Marshal(p)
	return ctx.GetStub().PutState("POOL", b)
}

func nowSeconds() int64 { return time.Now().Unix() }

func requireOrchestrator(ctx contractapi.TransactionContextInterface) error {
	return requireOU(ctx, ORCHESTRATOR_OU)
}
func requireOU(ctx contractapi.TransactionContextInterface, ou string) error {
	cid := ctx.GetClientIdentity()
	cert, err := cid.GetX509Certificate()
	if err != nil {
		return err
	}
	for _, o := range cert.Subject.OrganizationalUnit {
		if strings.EqualFold(o, ou) {
			return nil
		}
	}
	return fmt.Errorf("caller OU %v missing required %s", cert.Subject.OrganizationalUnit, ou)
}

// ---------------------------- entrypoint ----------------------------

func main() {
	cc, err := contractapi.NewChaincode(&RWAContract{})
	if err != nil {
		panic(err)
	}
	if err := cc.Start(); err != nil {
		panic(err)
	}
}
