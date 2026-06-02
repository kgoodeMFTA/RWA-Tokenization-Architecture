# Cybersecurity Architecture

Implementation of capstone §12 (Cybersecurity Risk Mitigation Plan).

## 7-Layer attack surface

See [`../diagrams/attack-surface.mmd`](../diagrams/attack-surface.mmd) for the visual map.

| # | Layer | Top threats | Primary controls (where implemented) |
|---|---|---|---|
| 1 | Galileo webhook receiver | Forgery, replay | HMAC-SHA256 + nonce + 5-min window (`orchestration/src/middleware/hmac.js`) |
| 2 | API gateway | JWT theft, scope escalation | OAuth 2.0 + JWT + tight scopes, rate limit (`orchestration/src/middleware/auth.js`) |
| 3 | Orchestration service | Logic bypass, race conditions | Idempotency keys, ML anomaly detector, mandatory KYT screen |
| 4 | Smart contracts | Reentrancy, overflow, access control | Checks-Effects-Interactions, OpenZeppelin `ReentrancyGuard`, `AccessControl`, Slither in CI |
| 5 | Oracle feed | Single-source manipulation, stale data | Median-of-N aggregator + z-score circuit breaker + staleness check (`contracts/src/OracleAdapter.sol`) |
| 6 | Besu validators | 51% / colluding validators, node compromise | Vetted validator set, 2-of-3 multisig on state-changing ops, CloudHSM (FIPS 140-2 L3) |
| 7 | Off-chain identity registry | PII exfiltration, key compromise | AES-256-GCM encryption, AWS Secrets Manager rotation, no PII on-chain ever |

## Control catalog (paper §12.2)

| Control | Where it lives |
|---|---|
| HMAC-SHA256 webhook signature + replay prevention | `orchestration/src/middleware/hmac.js` |
| Checks-Effects-Interactions + ReentrancyGuard | All Solidity contracts |
| Circuit breaker + z-score anomaly + staleness check | `contracts/src/OracleAdapter.sol` |
| AWS CloudHSM + 2-of-3 multisig (stub) | `orchestration/src/blockchain/signer.js` |
| AWS Secrets Manager + rotation (stub) | `orchestration/src/services/secrets.js` |
| AES-256-GCM off-chain encryption + data minimization | `orchestration/src/services/identity.js` |
| SHAP explainability + ECOA adverse-action codes | `analytics/src/credit_model.py` |
| Chainalysis KYT (mock) gate on every transfer | `orchestration/src/middleware/aml.js` |

## Attack scenario walkthrough — reentrancy on `RewardAssetAMM.sol`

Reproduces capstone §12.3 end-to-end. See `contracts/test/reentrancy.test.js` for the executable version.

```
Recon  → Slither catches CEI violation in Sprint 3 CI       → patched
Attempt → ReentrancyGuard.nonReentrant blocks 2nd call      → reverted
Detect → Isolation-Forest spikes >0.95 on attacker wallet   → wallet frozen
Respond → KYT flag + SOC alert <60s + SAR filed if needed   → contained
Learn  → Threat register + model retrained with signature   → improved
```

## Incident response (§12.4)

Five time-boxed phases, defined in [`incident-response.md`](incident-response.md) and codified as a runbook in `infrastructure/runbooks/`.
