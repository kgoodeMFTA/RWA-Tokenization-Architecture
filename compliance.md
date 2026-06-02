# Compliance by Architecture (§12.5)

Every regulatory obligation is enforced by the **absence of alternative code paths** — not by a policy document.

## CFPB Rule 1033 / FDX 6.0

- All consumer data access is funneled through `orchestration/src/middleware/consent.js`.
- No endpoint reads from `WALLET_MAPPING` or `ACCOUNT` without a matching valid `CONSENT_GRANT`.
- Enforced in middleware, asserted in `orchestration/tests/consent.test.js`.

## FinCEN BSA / AML

- Every `mint`, `transfer`, `redeem`, `creditDraw` route invokes `kytScreen(address)` before submitting the transaction.
- A failed KYT screen returns `403 Forbidden` and writes a `kyt_blocked` audit record.
- Mock implementation in `orchestration/src/middleware/aml.js`; production swaps in Chainalysis KYT.

## GDPR / CCPA

- The PII model lives entirely off-chain in the `CONSUMER` and `ACCOUNT` tables under AES-256-GCM at rest.
- The on-chain record carries only a `pseudonymous_id` (one-way hash) and pure protocol state.
- Right-to-erasure is satisfied by deleting the off-chain row; the on-chain pseudonym becomes meaningless.

## ECOA / Regulation B (12 CFR Part 1002)

- Every credit decision in `CREDIT_DECISION` is paired with N `SHAP_EXPLANATION` rows.
- For declines, the top-3 negative-SHAP features map to standardized adverse-action reason codes (`analytics/src/ecoa_codes.py`).
- Compliance officer review queue lives in `frontend/src/pages/admin/adverse-action-review.jsx`.

## NIST CSF 2.0 — Govern function

- Risk decisions are codified in `docs/architecture/security.md`, owned by the Compliance Lead, reviewed quarterly.
- All controls mapped to NIST CSF 2.0 sub-categories in `docs/architecture/nist-csf-mapping.md`.

## OWASP API Security Top 10 (2023)

- API1: Broken object-level authorization → resource-level scope checks on every route.
- API4: Unrestricted resource consumption → rate limiting + JWT-bound quotas.
- API7: SSRF → outbound HTTP allowlist for oracle and KYT only.
- API8: Security misconfiguration → CI fails the build if any `process.env.SECRET` ends up in `git diff`.
