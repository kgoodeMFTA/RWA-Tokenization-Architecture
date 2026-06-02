# Architecture Overview

This document is the engineering counterpart to **Sections 8–12** of the capstone paper. Each subsection maps a paper section to the concrete code in this repo.

## Layered model (six logical layers)

| Layer | Responsibility | Implementation |
|---|---|---|
| L1 — Consumer engagement | iOS/Android-style React UI; no crypto jargon | `frontend/` |
| L2 — API gateway | OAuth 2.0, FDX 6.0 consent, rate limiting | `orchestration/src/middleware/` |
| L3 — Orchestration | Galileo→EVM bridge, programmable finance engines | `orchestration/src/services/` |
| L4 — Intelligence | XGBoost thin-file credit scoring, RL yield optimization, SHAP explainability | `analytics/` |
| L5 — Blockchain execution | Solidity contracts on Besu PoA + Fabric chaincode | `contracts/`, `fabric/` |
| L6 — Trust anchor | Chainlink-style oracle adapter, Chainalysis KYT mock, AWS HSM stubs | `orchestration/src/blockchain/oracle.js`, `orchestration/src/middleware/aml.js` |

## Multi-chain strategy

The paper anchors on a permissioned EVM (Hyperledger Besu PoA). We also include a Hyperledger Fabric implementation because it shows the same business logic running on a non-EVM permissioned chain, which is a useful comparison for traditional banks evaluating the architecture.

| Concern | EVM (Besu / Ethereum) | Fabric |
|---|---|---|
| Contract language | Solidity 0.8.x | Go chaincode |
| Consensus | IBFT 2.0 PoA | Raft / PBFT |
| Privacy | Account-level + ZKP layer | Channels + private data collections |
| Throughput | ~300 TPS | ~1,000+ TPS |
| Best fit | Neobanks, BaaS, DeFi composability | Traditional banks, consortium settlement |

## Three-tier smart contract architecture (§8.1.b)

1. **Asset origination layer** — `DepositReceiptToken.sol` mints ERC-20 yield-bearing receipts on payroll deposit events.
2. **Programmable logic layer** — `RWACreditLine.sol` issues collateralized credit; `RewardAssetAMM.sol` provides the redemption secondary market.
3. **Consumer interface layer** — `HITLGovernor.sol` mediates any automated action above the consumer-set USD threshold (default $50).

## Programmable finance engines (§9.5)

- **Overdraft Prevention Engine** — Galileo `transaction_posted` webhook → balance check via FDX 6.0 → waterfall (Savings → SpotMe → HITL alert).
- **Late-Fee Prevention Engine** — daily cron checks upcoming due dates, stages payment 3 days out, reschedules to next predicted paycheck if short.
- **Debt Payoff Acceleration Engine** — avalanche/snowball ranking + post-direct-deposit surplus sweep with 24-hour HITL window.

All three live in `orchestration/src/services/programmable-finance.js` and are unit-tested in `orchestration/tests/`.

## Compliance-as-architecture

Compliance is enforced by *the absence of alternative code paths*, not by policy documents. See [`compliance.md`](compliance.md).

## See also

- [`security.md`](security.md) — 7-layer attack surface, controls, incident response.
- [`buy-vs-build.md`](buy-vs-build.md) — DivByZero six-step matrix applied to every component.
- [`../diagrams/`](../diagrams/) — Mermaid system, sequence, attack-surface diagrams.
