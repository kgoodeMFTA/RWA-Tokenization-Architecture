# Incident Response Runbook (§12.4)

Five time-boxed phases. SLAs measured from confirmed-breach timestamp.

| Phase | Window | Trigger | Owner | Required actions |
|---|---|---|---|---|
| 1. Detect | 0–15 min | SIEM alert OR contract `EmergencyPaused` event OR KYT high-risk hit | On-call SRE | Acknowledge in PagerDuty; open incident channel; freeze the affected contract via `Pausable`; snapshot logs |
| 2. Contain | 15–60 min | Phase 1 done | Security Lead | Rotate compromised keys via CloudHSM; revoke OAuth tokens; pause Galileo webhook ingestion; engage Chainalysis for chain tracing |
| 3. Eradicate | 1–4 hr | Root cause confirmed | Eng Lead | Patch + redeploy contracts (audit waiver requires CTO + Compliance sign-off); rotate all webhook secrets; force-renew JWTs |
| 4. Recover | 4–24 hr | Patch deployed | Eng Lead + SRE | Unpause contracts after canary verification; replay missed Galileo events from durable queue; reconcile on-chain vs ledger state |
| 5. Lessons-learned | ≤ 5 business days | Post-recover | Compliance Lead | Post-mortem document; threat register update; ML detector retrained with new signature; regulator notice (SAR / 8-K / consumer notification as applicable) |

## Required notifications

| Trigger | Notify within | Recipient |
|---|---|---|
| Suspected unauthorized PII access | 72 hr | State AG / CCPA bodies |
| Suspicious activity > $10k or structuring pattern | 30 days | FinCEN (SAR) |
| Material cyber incident at public co. parent | 4 business days | SEC (8-K Item 1.05) |
| Consumer financial harm | Per state law | Affected consumers |
