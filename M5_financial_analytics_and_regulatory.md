# M5 — Financial Analytics Strategy & Regulatory/Risk Environment
**Wake Forest FTA-799 Capstone | Insurance Claim Assistant (ICA)**
**Author:** Capstone Researcher | **Date:** 2025 | **Status:** Milestone 5

---

## 1. Financial Analytics Strategy

### 1.1 ICA Internal Unit Economics

ICA tracks the following unit economics dashboard to manage its own SaaS business health:

| Metric | Definition | Formula | Year 1 Target | Year 3 Target | Year 5 Target |
|--------|-----------|---------|--------------|--------------|--------------|
| **ARR** | Annual Recurring Revenue | Sum of annualized contract values | $300K | $3.15M | $8.4M |
| **MRR** | Monthly Recurring Revenue | ARR / 12 | $25K | $262K | $700K |
| **Gross Margin %** | Revenue less COGS / Revenue | (Rev - COGS) / Rev | 7% | 54% | 64% |
| **CAC** | Customer Acquisition Cost | Total sales + marketing spend / new logos | $95K | $65K | $45K |
| **LTV** | Customer Lifetime Value | Gross Margin $ × (1 / logo churn rate) | $190K | $680K | $1.1M |
| **LTV/CAC Ratio** | SaaS health indicator | LTV / CAC | 2.0x | 10.5x | 24.4x |
| **Logo Churn Rate** | Annual customer loss rate | Lost logos / total logos | 20% (pilot exit risk) | 10% | 7% |
| **Net Revenue Retention** | Expansion + retention | (Beg ARR + expansion - contraction - churn) / Beg ARR | 100% | 115% | 125% |
| **Months to Recover CAC** | Payback period | CAC / (MRR × Gross Margin %) | 37 months | 10 months | 5 months |
| **Claims Processed (annual)** | Platform volume | Sum of claims passing through ICA | 30K | 650K | 2.1M |

*All Year 1–5 targets are INTERNAL ASSUMPTIONS unless cited. LTV/CAC and NRR targets modeled on comparable vertical SaaS benchmarks.*

### 1.2 Cohort Retention Analytics

Carrier cohorts (grouped by contract start quarter) are tracked on:
- **Gross Revenue Retention (GRR):** The floor — percent of prior period revenue retained, ignoring expansions. Target: ≥90% by Year 3.
- **Net Revenue Retention (NRR):** Expansion above GRR. Carriers add claims volume as they onboard more LOBs and branch offices. Target: ≥115% by Year 3.
- **Claims volume ramp:** Carrier cohorts typically start with 1 LOB (Auto) and expand to HO and GL within 12 months — tracked as LOB expansion rate per cohort.

### 1.3 Gross Margin per Carrier Segment

| Carrier Tier | Avg ARR | COGS (LLM + infra + CS) | Gross Margin | Notes |
|-------------|---------|------------------------|-------------|-------|
| Pilot (Tier 3, <100K claims/yr) | $120K | $90K | 25% | High CS cost relative to ARR; educational phase |
| Growth (Tier 2, 100K–500K claims/yr) | $400K | $175K | 56% | Automation offsets CS burden; LLM cost per claim falls with batching |
| Enterprise (Tier 1, >500K claims/yr) | $1.2M | $450K | 63% | Near target SaaS margin; multi-LOB expansion drives NRR |

*INTERNAL ASSUMPTIONS — all figures*

---

### 1.4 Value Delivered to Carriers: Financial Models

ICA's carrier ROI narrative is built on three quantifiable value levers:

#### A. ALAE Reduction Model

**Formula:** ALAE Savings = (Claims Processed × Pre-AI ALAE per claim) − (Claims Processed × Post-AI ALAE per claim)

| Input | Value | Source |
|-------|-------|--------|
| Pre-AI ALAE per claim (ALAE portion of $40–$60 total LAE) | ~$28 | INTERNAL ASSUMPTION — ALAE estimated as ~56% of total LAE processing cost, consistent with industry ALAE/ULAE splits |
| AI cost reduction on ALAE-intensive steps | 35% | Mid-range of 20–40% ([McKinsey, 2021](https://www.mckinsey.com/industries/financial-services/our-insights/insurance-2030-the-impact-of-ai-on-the-future-of-insurance)) |
| Post-AI ALAE per claim | ~$18 | $28 × (1 − 0.35) |
| ALAE saving per claim | $10 | |
| At 100,000 claims/yr (Tier 2 carrier) | **$1.0M/yr** | |

**Carrier ROI:** At $400K ARR to ICA, carrier nets **$600K+ ALAE saving annually** — a 2.5x ROI before accounting for cycle-time and NPS benefits.

#### B. Cycle-Time-to-Loss-Cost Elasticity

Extended claim cycles correlate with increased indemnity through two mechanisms:
1. **Attorney involvement:** Claims open >30 days are significantly more likely to involve legal representation; litigated BI claims settle for 3–5× unrepresented claims (INTERNAL ASSUMPTION — consistent with State Farm loss development data from researcher experience).
2. **Medical cost escalation:** Open BI claims accumulate treatment costs that would not occur if claims resolved promptly.

**Elasticity model:** Each 10-day reduction in cycle time reduces expected indemnity on BI claims by ~3–5% through reduced attorney involvement rate (INTERNAL ASSUMPTION — based on researcher's 5 years State Farm auto injury claims management experience). At an average BI severity of $28,278 ([III, 2024](https://www.iii.org/fact-statistic/facts-statistics-auto-insurance)) and 500 BI claims/carrier cohort, a 10-day cycle time reduction (ICA's conservative target per simulation) saves:
- $28,278 × 4% × 500 = **$565,560 in indemnity per carrier per year**

#### C. Leakage Reduction Estimation

Claims leakage — payments beyond what policy and facts warrant — is estimated at 5–10% of incurred losses industry-wide (INTERNAL ASSUMPTION — consistent with practitioner estimates; precise figure is proprietary at individual carriers). ICA's AI-driven next-best-action module and SIU scoring reduce leakage through:
- Timely SIU referral (AI flags suspicious patterns at FNOL; currently flagged days/weeks later in manual process)
- Evidence completeness checks (missing police reports, medical records caught at intake rather than at settlement)
- Subrogation identification at FNOL vs. at close

At a conservative 1% leakage reduction on $10M annual losses at a Tier 2 carrier: **$100K additional carrier value** beyond ALAE savings.

### 1.5 Future Product Revenue: Predictive Analytics Suite

| Product | Potential Revenue Model | Year Available |
|---------|------------------------|---------------|
| Predictive severity scoring | $1.50/claim scored | Year 3 |
| Reserve adequacy AI (with actuarial sign-off) | $75K/yr platform add-on | Year 4 |
| Reinsurance optimization dashboard | $150K/yr enterprise | Year 5 |
| Carrier benchmarking network (anonymized data) | $25K/yr per participant | Year 4 |

---

## 2. Regulatory Environment — Deep-Dive Matrix

| Regulation | Jurisdiction | Scope | Key ICA Requirements | Compliance Status |
|-----------|-------------|-------|---------------------|------------------|
| **GLBA Safeguards Rule** (16 CFR Part 314, amended 2023) | Federal | All financial institutions incl. insurers; protects NPI | AES-256 at rest, TLS 1.3 in transit, RBAC, audit logging, annual pen test, vendor risk management (LLM BAA), designated security program officer | Implemented in v1 architecture; BAAs required before production |
| **HIPAA Privacy + Security Rule** (45 CFR §§ 160, 164) | Federal | BI claims involving PHI; carrier as Covered Entity, ICA as Business Associate | BAA with carrier + LLM providers; HIPAA-eligible S3 bucket for medical evidence; minimum necessary standard; 60-day breach notification | BAA framework designed; PHI storage architecture in ARCHITECTURE.md |
| **NAIC Model Act #900 — UCSPA** | All 50 states | Unfair claims settlement practices | AI outputs must not misrepresent coverage, delay investigation, or coerce settlement; all AI outputs marked advisory; adjuster approval required | Implemented via Tier 4 blocks (coverage opinions blocked); disclaimers on all AI outputs |
| **NAIC AI Model Bulletin** (2023/2024) | States adopting (growing list) | AI systems used in insurance decisions | Written AIS Program; governance; model inventory; testing for unfair bias; third-party oversight; consumer notice | AIS Program framework in this document; model registry in data model |
| **Colorado SB21-169 + Reg 10-1-1** (effective Nov 14, 2023) | Colorado | Life (now); Auto/health in rulemaking 2025–2026 | Governance framework; ECDIS inventory; ongoing disparate impact testing; annual attestation to DOI; vendor due diligence | Applicable when ICA carriers write Colorado life/auto; bias testing framework in M3 |
| **NY DFS Circular Letter No. 7** (July 11, 2024) | New York | All insurance lines; AI in underwriting/pricing | Quantitative bias testing (Step 1–3 framework); governance; board oversight; consumer disclosure when AI used; adverse action notice | Note: CL7 focuses on underwriting/pricing, not claims; ICA's AI is claims-only — monitor for expansion |
| **EU AI Act** (effective 2025–2026) | EU (extraterritorial for EU claimants) | High-risk AI systems including insurance | If any EU claimants served: conformity assessment, technical documentation, human oversight requirements | Out of scope v1 (US-only); design AI oversight architecture to be EU AI Act-compatible |
| **CCPA / CPRA** (Cal. Civ. Code §1798) | California | California residents' personal data | Right to know, delete, opt-out of sale/sharing; annual privacy report; data minimization | Required for California claimants; implement privacy portal and data deletion workflow |
| **NAIC Model Unfair Trade Practices Act #880** | All 50 states | Unfair/deceptive insurance practices | AI outputs cannot constitute unfair trade practices; no false/misleading policy interpretation | Enforced via disclaimer requirements and T4 blocks on coverage opinions |
| **State Rate Filing Implications** | Varies | When AI used in pricing/adjudication decisions | If ICA outputs influence rate or adjudication, may require regulatory filing in states with prior approval | ICA advisory-only in v1; monitor regulatory interpretation as AI adjudication scope expands |

---

## 3. Enterprise Risk Register

| # | Risk | Category | Likelihood (1–5) | Impact (1–5) | Risk Score | Mitigation Strategy | Owner | Monitoring KPI |
|---|------|----------|-----------------|-------------|------------|--------------------|----|---------------|
| R1 | LLM hallucination causes incorrect claim guidance relied upon by claimant | AI/Ops | 3 | 5 | **15** | T4 blocks on coverage opinions; mandatory disclaimers; adjuster review for T2 outputs; faithfulness monitoring | AI/ML Lead | Weekly hallucination rate |
| R2 | PHI/PII data breach via LLM API or S3 misconfiguration | Security | 2 | 5 | **10** | PII redaction + BAA + S3 block public; annual pen test; incident response plan | Security Lead | Zero PII-in-LLM-call incidents |
| R3 | Carrier rejects AI outputs due to regulatory non-compliance | Regulatory | 3 | 4 | **12** | AIS Program documentation; pre-sale legal review; NAIC bulletin alignment | CEO + Counsel | Carrier compliance objections per quarter |
| R4 | AI bias causes disparate claim outcomes by protected class | Regulatory/Ethical | 2 | 5 | **10** | Quarterly disparate impact testing; bias monitoring dashboard; claims SME review | AI/ML Lead + Legal | Disparate Impact Ratio by protected class |
| R5 | Prompt injection attack via uploaded evidence file | Security | 3 | 4 | **12** | Input validation; injection pattern detection; LLM output schema validation | Engineering Lead | Prompt injection detection rate |
| R6 | Key man dependency (researcher/founder leaves) | Operational | 2 | 4 | **8** | Document all insurance domain knowledge; train claims SME team; board advisory from industry | CEO | Knowledge transfer % complete |
| R7 | LLM provider API outage (OpenAI/Anthropic) | Operational | 3 | 3 | **9** | Provider failover in `ai_service.py` (env-var swap); Redis cache for recent outputs; SLA from provider | Engineering Lead | LLM failover time; SLA uptime |
| R8 | Carrier IT security requirements block cloud deployment | Commercial | 3 | 3 | **9** | SOC 2 Type II; carrier-managed key option; on-prem Kubernetes deployment option in roadmap | Engineering + Sales | # carrier deals lost to security objections |
| R9 | Loss cost inflation outpaces LAE reduction savings (carrier ROI narrative breaks) | Market | 2 | 3 | **6** | ROI model includes sensitivity analysis; update benchmarks quarterly; position ICA as inflation hedge (cycle time reduces attorney involvement) | Product + Sales | Carrier reported LAE delta |
| R10 | Competitor (Snapsheet, Five Sigma, Hi Marley) adds LLM features matching ICA capability | Competitive | 4 | 3 | **12** | Ship faster; deepen insurance domain specificity; build claims SME network as moat; patent prompt architecture | Product | Feature gap vs. competitors quarterly |

### Risk Heatmap

```
IMPACT
  5 | R1          | R2,R4        |
  4 | R6          | R3,R5,R10   |
  3 |             | R7,R8        | R9
  2 |             |              |
  1 |             |              |
    | Low (1-2)   | Med (3)      | High (4-5)
                    LIKELIHOOD
```

*R1 (hallucination → harm) and R3 (regulatory rejection) are the highest-priority risks requiring active management protocols.*

---

*Page 1 of approximately 5.5 | Next: M6 — ESG, Ethics & KPIs*
