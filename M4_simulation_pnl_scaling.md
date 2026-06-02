# M4 — Product Simulation, P&L, Scaling Plan & Cross-Sell Strategy
**Wake Forest FTA-799 Capstone | Insurance Claim Assistant (ICA)**
**Author:** Capstone Researcher | **Date:** 2025 | **Status:** Milestone 4

---

## 1. Product Simulation: Synthetic 1,000-Claim Cohort

### 1.1 Simulation Design

To quantify ICA's impact before production data is available, a synthetic cohort of 1,000 claims was designed to mirror the distribution of a mid-size regional carrier writing Auto, Homeowners, and General Liability across 10 states. Inputs are anchored to published industry benchmarks; all derivations are labeled.

**Cohort Composition:**

| Line of Business | Claim Count | % of Cohort | Basis |
|-----------------|------------|-------------|-------|
| Private Passenger Auto (PD + BI) | 500 | 50% | Auto is the highest-volume P&C line; reflects carrier book typical of regional carriers ([III/NAIC, 2024](https://www.iii.org/fact-statistic/facts-statistics-auto-insurance)) |
| Homeowners (wind/water/fire) | 350 | 35% | Homeowners second-largest by claim count at ~5.3% annual claim frequency ([III, 2023](https://www.iii.org/fact-statistic/facts-statistics-homeowners-and-renters-insurance)) |
| General Liability | 150 | 15% | GL claims are lower-volume but higher-complexity; INTERNAL ASSUMPTION based on typical carrier book mix |

**Baseline Inputs (Pre-ICA):**

| Metric | Baseline | Source |
|--------|---------|--------|
| Cost per standard claim (processing only) | $50 (midpoint) | $40–$60 range ([Research & Markets / Decerto, 2026](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers)) |
| Average cycle time (days, all lines) | 25 days | INTERNAL ASSUMPTION — consistent with industry average of 7–10 days routine, 30 days average ([Research & Markets, 2026](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers)) |
| FNOL intake time | 25 min | Manual process average ([Kagen AI / Deloitte, 2026](https://www.kagen.ai/blog/ai-voice-agents-reduce-claims-processing-time-insurance)) |
| Adjuster hours per claim (non-complex) | 3.5 hours | INTERNAL ASSUMPTION — derived from $50/claim cost at $14/hr equivalent labor cost |
| Claimant NPS (baseline) | 22 | INTERNAL ASSUMPTION — consistent with typical P&C claims NPS range of 15–35 |

### 1.2 Simulation Outputs (ICA Deployed)

**Methodology:** Applied McKinsey's 20–40% cost reduction ([McKinsey, 2021](https://www.mckinsey.com/industries/financial-services/our-insights/insurance-2030-the-impact-of-ai-on-the-future-of-insurance)), Govindaswamy Subbian's 50% cycle-time reduction ([AJRCOS, 2025](https://journalajrcos.com/index.php/AJRCOS/article/view/604)), and Accenture's 73% process efficiency increase ([Accenture](https://www.accenture.com/content/dam/accenture/final/accenture-com/document/Accenture-Why-AI-In-Insurance-Claims-And-Underwriting.pdf)). Used the conservative end of each range to produce defensible projections.

| Output Metric | Pre-ICA | Post-ICA | Delta | Assumption |
|--------------|---------|---------|-------|-----------|
| **Processing cost per claim** | $50.00 | $32.50 | -$17.50 (-35%) | 35% reduction (mid-range of 20–40%, [McKinsey, 2021](https://www.mckinsey.com/industries/financial-services/our-insights/insurance-2030-the-impact-of-ai-on-the-future-of-insurance)) |
| **Total LAE for 1,000 claims** | $50,000 | $32,500 | **-$17,500** | Based on per-claim cost above |
| **Average cycle time** | 25 days | 12.5 days | -12.5 days (-50%) | 50% reduction ([AJRCOS, 2025](https://journalajrcos.com/index.php/AJRCOS/article/view/604)); conservative vs. 75% at advanced AI deployments |
| **FNOL intake time** | 25 min | 7 min | -18 min (-72%) | AI-assisted intake; INTERNAL ASSUMPTION anchored to 65–70% reduction in FNOL cost ([Research & Markets, 2026](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers)) |
| **Adjuster hours freed (per claim)** | 3.5 hrs | 1.75 hrs | -1.75 hrs | 50% human intervention reduction ([McKinsey, 2021](https://www.mckinsey.com/industries/financial-services/our-insights/insurance-2030-the-complete-guide-for-us-carriers)) |
| **Total adjuster hours freed (1,000 claims)** | 3,500 hrs | 1,750 hrs | **-1,750 hrs** | |
| **Claimant NPS lift** | 22 | 38 | +16 pts | INTERNAL ASSUMPTION — proactive communication + faster resolution; directionally consistent with Accenture customer experience findings |
| **SIU flag accuracy** | Rules-based, ~65% | ML model, ~82% | +17% AUROC | INTERNAL ASSUMPTION — consistent with [Saikia et al., IEEE 2024](https://ieeexplore.ieee.org/document/10724028/) XGBoost at 84% accuracy |

**At scale (annual 50,000 claims / mid-size carrier):** ICA generates approximately **$875,000 in LAE savings** and frees **87,500 adjuster hours** annually — time redirectable to complex, high-value claims requiring human judgment.

---

## 2. Market Sizing: TAM / SAM / SOM

### 2.1 TAM — Total Addressable Market

**Definition:** All U.S. P&C claims technology and LAE spend addressable by an AI claims assistant platform.

- US P&C industry LAE incurred (2024): **$85.4 billion** ([NAIC 2024 Annual Report](https://content.naic.org/sites/default/files/2024-annual-property-casualty-and-title-insurance-industries-analysis-report.pdf))
- Global insurtech market (2023): **$7.87 billion** ([Grand View Research, 2023](https://www.grandviewresearch.com/industry-analysis/insurtech-market))
- AI in insurance claims processing sub-market (2025): **$0.46 billion**, growing to **$0.97 billion by 2030** at 16.2% CAGR ([Research and Markets, 2026](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers))

**TAM = $0.97 billion** (AI claims processing market, U.S. addressable segment, 2030 projection) — conservative framing that excludes the broader LAE pool ICA can displace over a longer horizon.

### 2.2 SAM — Serviceable Addressable Market

**Definition:** Mid-size U.S. P&C carriers (Tier 2/3: $100M–$2B DWP), TPAs, and MGAs writing Auto, Homeowners, and GL — the segments with sufficient claims volume to justify ICA's platform fee but without in-house AI engineering capacity.

- Estimated mid-size carrier count in target profile: ~250 carriers ([NAIC Market Share Data, 2024](https://content.naic.org/sites/default/files/publication-msr-pb-property-casualty.pdf))
- Estimated TPA count (claims-bearing): ~100 firms (INTERNAL ASSUMPTION)
- Average claims processed per carrier: 50,000–500,000/yr
- Per-claim revenue at $4.00/claim × 100,000 claims/carrier: $400K ARR per carrier
- Platform fee per carrier: $60K/yr (INTERNAL ASSUMPTION)

**SAM = ~$115 million ARR** (250 carriers + 100 TPAs × blended $320K ARR potential)

### 2.3 SOM — Serviceable Obtainable Market (Year 3)

**Definition:** Reachable portion of SAM given ICA's go-to-market capacity, current product maturity, and competitive intensity.

- Year 3 target: 15 carriers / TPAs at blended $350K ARR
- Plus MGA channel: 5 MGAs at $150K ARR

**SOM Year 3 = $6.0M ARR** (INTERNAL ASSUMPTION — based on scaled deployment rate from Year 1 pilot → Year 3 expansion per scaling plan below)

---

## 3. Five-Year P&L Summary

*Full model in [ICA_PnL_Model.xlsx](./ICA_PnL_Model.xlsx). This section summarizes structure, assumptions, and trajectory.*

### 3.1 Revenue Model

| Revenue Stream | Mechanism | Year 1 | Year 3 | Year 5 |
|----------------|-----------|--------|--------|--------|
| **Per-claim SaaS fee** | $3.50–$5.00/claim processed through ICA | $120K | $2.1M | $6.0M |
| **Platform fee** | $50K–$100K/yr per carrier (onboarding, support, analytics) | $100K | $750K | $2.0M |
| **Professional services** | Implementation, carrier integration, training | $80K | $300K | $400K |
| **Total Revenue** | | **$300K** | **$3.15M** | **$8.4M** |

*All Year 1–5 figures are INTERNAL ASSUMPTIONS anchored to comparable insurtech per-claim pricing (comps: Snapsheet, Five Sigma, Hi Marley — INTERNAL ASSUMPTION as specific pricing not publicly disclosed).*

### 3.2 COGS Structure

| COGS Item | Driver | Cost Assumption | Source |
|-----------|--------|----------------|--------|
| LLM API tokens | ~$0.03/claim (GPT-4o input + output) | 30–40% of revenue Year 1; declining to ~12% by Year 5 | INTERNAL ASSUMPTION — GPT-4o pricing $0.005/1K input + $0.015/1K output tokens |
| AWS infrastructure (ECS, RDS, S3, ElastiCache) | Fixed base + per-claim variable | $8K/month base (Year 1) → $25K/month (Year 3) | INTERNAL ASSUMPTION |
| OCR (AWS Textract) | $0.015/page; ~5 pages/claim | ~$0.075/claim | [AWS Textract pricing](https://aws.amazon.com/textract/pricing/) |
| Customer success (CSM) | 1 CSM per 8 carrier accounts | $95K fully loaded per CSM | INTERNAL ASSUMPTION |

### 3.3 Gross Margin Trajectory

| Year | Revenue | COGS | Gross Profit | Gross Margin % |
|------|---------|------|-------------|---------------|
| 1 | $300K | $280K | $20K | 6.7% |
| 2 | $1.2M | $720K | $480K | 40% |
| 3 | $3.15M | $1.45M | $1.70M | 54% |
| 4 | $5.8M | $2.3M | $3.5M | 60% |
| 5 | $8.4M | $3.0M | $5.4M | 64% |

*Target: 65%+ gross margin by Year 5 — consistent with SaaS benchmarks; insurance AI platforms have elevated COGS vs. pure SaaS due to LLM token costs (INTERNAL ASSUMPTION).*

### 3.4 EBITDA and Breakeven

| Year | Gross Profit | OpEx (Sales + R&D + G&A) | EBITDA | EBITDA Margin |
|------|-------------|--------------------------|--------|--------------|
| 1 | $20K | $1.8M | **-$1.78M** | NM |
| 2 | $480K | $2.4M | **-$1.92M** | NM |
| 3 | $1.70M | $2.8M | **-$1.1M** | NM |
| 4 | $3.5M | $3.2M | **+$300K** | 5.2% |
| 5 | $5.4M | $3.8M | **+$1.6M** | 19% |

**EBITDA breakeven: Year 4** — consistent with comparable insurtech SaaS companies ([INTERNAL ASSUMPTION based on comparable SaaS CAC/LTV dynamics])

---

## 4. Scaling Plan: Years 1–5

```mermaid
gantt
    title ICA Go-to-Market Milestones (5-Year)
    dateFormat YYYY-Q[Q]
    axisFormat %Y-Q%q

    section Year 1 — Foundation
    Seed funding close         :done, 2025-Q1, 2025-Q1
    v1 product launch          :done, 2025-Q1, 2025-Q2
    Pilot Carrier 1 (50 claims):active, 2025-Q2, 2025-Q3
    Pilot Carrier 2            :2025-Q3, 2025-Q4
    Series A fundraise         :2025-Q4, 2026-Q1

    section Year 2 — Expansion
    Carrier 3-5 signed         :2026-Q1, 2026-Q2
    TPA channel launch         :2026-Q2, 2026-Q3
    Multi-tenant architecture  :2026-Q1, 2026-Q3
    SOC 2 Type II issued       :2026-Q2, 2026-Q2

    section Year 3 — Platform
    MGA platform launch        :2027-Q1, 2027-Q2
    15 carriers/TPAs live      :2027-Q1, 2027-Q4
    Guidewire integration      :2027-Q2, 2027-Q4
    $3M ARR milestone          :2027-Q4, 2027-Q4

    section Years 4-5 — Scale
    Series B                   :2028-Q1, 2028-Q2
    Enterprise carrier tier    :2028-Q1, 2029-Q4
    EBITDA positive            :2028-Q2, 2028-Q2
    $8M ARR                    :2029-Q4, 2029-Q4
```

### 4.1 Team Scaling by Function

| Year | Total HC | Engineering | Product | Sales (AE/SE/SDR) | Customer Success | G&A |
|------|---------|-------------|---------|-------------------|-----------------|-----|
| 1 | 6 | 3 | 1 | 1 | 1 | 0 (founder) |
| 2 | 14 | 5 | 2 | 3 | 2 | 2 |
| 3 | 26 | 8 | 3 | 7 | 4 | 4 |
| 4 | 42 | 12 | 4 | 12 | 8 | 6 |
| 5 | 60 | 16 | 5 | 18 | 12 | 9 |

*All headcount figures INTERNAL ASSUMPTIONS*

### 4.2 ARR Milestones

| Milestone | Target Date | ARR |
|-----------|-------------|-----|
| First paid carrier | Q3 Year 1 | $120K |
| 5 carriers live | Q2 Year 2 | $1.2M |
| TPA channel first revenue | Q3 Year 2 | $1.5M |
| Series A close | Q1 Year 2 | — |
| 15 carriers/TPAs | Q4 Year 3 | $3.15M |
| MGA platform revenue | Q2 Year 3 | — |
| EBITDA breakeven | Q2 Year 4 | ~$4.5M ARR |
| Series B close | Q1 Year 4 | — |
| 40+ accounts | Q4 Year 5 | $8.4M ARR |

---

## 5. Cross-Sell Opportunities

Once ICA has established a carrier/TPA integration, the marginal cost of delivering adjacent products over the same data and workflow infrastructure is low. The following cross-sells are sequenced by technical dependency and market readiness:

| Product | Description | Buyer | Revenue Model | Tech Dependency | Priority |
|---------|-------------|-------|---------------|----------------|----------|
| **Claims Analytics Dashboard** | Carrier-facing BI: cycle time by adjuster, LAE by claim type, SIU hit rate, leakage analysis | Carrier VP of Claims | Platform add-on: $25K–$50K/yr | Existing data warehouse | Year 2 |
| **Fraud / SIU Intelligence Module** | ML-powered fraud score + network graph analysis (multi-claimant fraud rings); integrates with ISO ClaimSearch | Carrier SIU dept. | Per-claim score: $0.50/claim | ML model on claim features | Year 2 |
| **Subrogation Acceleration** | AI identifies subrogation opportunities at FNOL; flags liens; drafts recovery letters | Carrier subrogation team | % of recovered subrogation OR flat per-flag | Evidence summarization + legal entity detection | Year 3 |
| **Claimant Self-Service Portal (White-Label)** | Branded claimant portal deployable by carrier without ICA branding | Carrier IT / Marketing | $50K implementation + $30K/yr | Multi-tenant architecture (Year 2 required) | Year 3 |
| **Predictive Reserve Adequacy** | ML model predicts ultimate claim cost at FNOL + reserve adequacy flag at 60/90/120 days | Carrier actuarial / finance | Platform add-on: $75K/yr | Historical claim outcome data (requires 12+ months of live data) | Year 4 |

---

## 6. UX Advantages vs. Competitive Landscape

| Feature | ICA | Snapsheet | Tractable | Traditional Carrier Portal |
|---------|-----|-----------|-----------|---------------------------|
| AI FNOL narrative drafting | ✓ | Limited | ✗ | ✗ |
| Plain-English policy explainer | ✓ | ✗ | ✗ | ✗ |
| Claimant-facing status board with AI rationale | ✓ | Partial | ✗ | Basic status only |
| Next-best-action for claimants AND adjusters | ✓ | Adjuster-only | ✗ | ✗ |
| Multi-LOB (Auto + HO + GL) in one wizard | ✓ | Auto-focused | Auto/vehicle damage | Carrier-specific |
| API-first, white-label ready | ✓ (Year 2) | ✓ | ✓ | ✗ |
| LLM provider agnostic (swap OpenAI ↔ Anthropic) | ✓ | ✗ | ✗ | N/A |
| Timeline estimation with SOL surfacing | ✓ | ✗ | ✗ | ✗ |
| Built-in PII redaction + GLBA/HIPAA controls | ✓ | Partial | Partial | Varies |
| Open-source friendly / audit trail | ✓ | Proprietary | Proprietary | Proprietary |

**Key differentiator:** ICA is the only platform that serves both the **claimant** (consumer-grade UX, plain English) and the **adjuster** (AI-drafted artifacts, next actions) from a single API-first codebase, across three lines of business, with built-in regulatory controls. Tractable excels at vehicle damage AI but does not touch the claims workflow; Snapsheet offers digital claims management but relies on carrier-specific LLM implementations. ICA's LLM-native architecture and insurance-specific prompt library (see `AI_PROMPTS.md`) represent a compound moat that deepens with every claim processed.

---

*Page 1 of approximately 6 | Next: ICA_PnL_Model.xlsx (see separate file)*
