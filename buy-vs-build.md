# Buy vs Build Matrix (capstone §10)

Following Chieruzzi's (2021) DivByZero six-step framework. Cost is step six, not step one.

| Component | Decision | Why |
|---|---|---|
| Galileo Core Banking API | **Buy** | Already powers Chime's ledger; webhooks are the real-time event source for tokenization. No new infra required. |
| RWA Orchestration Service (middleware) | **Build** | The Galileo→EVM bridge is the core competitive moat. No vendor sells this. Owned IP. → `orchestration/` |
| Smart Contracts (Deposit / AMM / Credit) | **Build + Audit** | Encode unique consumer-banking RWA business rules. Third-party audit (Certik / Trail of Bits / OpenZeppelin) is non-negotiable before prod. → `contracts/` |
| Permissioned Blockchain Node (Besu) | **Buy (managed)** | Use a managed Hyperledger Besu service. Enterprise SLA, FIPS 140-2 keys, monitoring out of the box. → infra config in `infrastructure/docker/` |
| AI Credit Scoring Model | **Build** | On-chain transaction history is proprietary data. ECOA Reg B requires full explainability — black-box vendor APIs = regulatory liability. → `analytics/src/credit_model.py` |
| Chainlink Oracle Feeds | **Buy** | Decentralized network already secures >$100B in DeFi. Building in-house re-introduces single-point manipulation risk. → adapter stub in `contracts/src/OracleAdapter.sol` |
| Circle USDC Mint API | **Buy** | Issuing proprietary stablecoin = OCC charter / 50-state MTL. USDC is already reserve-audited and Visa-integrated. → `orchestration/src/services/usdc-mint.js` |
| Chainalysis KYT (AML) | **Buy** | Replicating BSA-grade forensic analytics in-house would take years of data. → middleware in `orchestration/src/middleware/aml.js` |
| Consumer UX / App | **Build** | Trust is the product. Tokenization must feel like a normal Chime feature — no crypto jargon, no wallet addresses. → `frontend/` |
