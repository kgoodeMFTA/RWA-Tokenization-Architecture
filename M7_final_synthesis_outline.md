# M7 — Final Synthesis: Outline & Executive Summary
**Wake Forest FTA-799 Capstone | Insurance Claim Assistant (ICA)**
**Author:** Capstone Researcher | **Date:** 2025 | **Status:** Milestone 7

---

## SECTION A: EXECUTIVE SUMMARY (Full Draft — 1 Page)

The U.S. property and casualty insurance industry paid **$558.8 billion in net losses and $85.4 billion in Loss Adjustment Expenses in 2024** ([NAIC, 2024](https://content.naic.org/sites/default/files/2024-annual-property-casualty-and-title-insurance-industries-analysis-report.pdf)) — yet the operational systems processing those payments were largely designed in the 1990s. More than 60% of U.S. insurers report zero straight-through processing capability in claims ([Aite-Novarica, 2023](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers)). The average claim still costs $40–$60 to process manually and takes 25 days to resolve. Claimants — filing at the worst moments of their lives — navigate confusing portals, opaque policy language, and call-center dependency. Poor claims experiences put **$170 billion in global insurance premiums at risk by 2027** ([Accenture, 2022](https://newsroom.accenture.com/news/2022/poor-claims-experiences-could-put-up-to-170b-of-global-insurance-premiums-at-risk-by-2027-according-to-new-accenture-research)).

**The Insurance Claim Assistant (ICA)** is an AI-powered claims management platform purpose-built for mid-size U.S. carriers, third-party administrators (TPAs), and managing general agents (MGAs) writing Auto, Homeowners, and General Liability. ICA embeds large language models at every stage of the claims lifecycle: AI-drafted First Notice of Loss narratives, evidence summarization, plain-English policy explanation, next-best-action recommendations, and timeline estimation — all with carrier-grade security, GLBA/HIPAA compliance, and PII redaction before any data reaches an external LLM provider.

The business model is B2B2C SaaS: carriers and TPAs pay a per-claim fee ($3.50–$5.00) plus an annual platform fee ($50K–$100K), while their claimants receive a consumer-grade experience at no direct cost. The unit economics are compelling: McKinsey projects AI reduces claims processing costs 20–40% and cycle time up to 30% ([McKinsey, 2021](https://www.mckinsey.com/industries/financial-services/our-insights/insurance-2030-the-impact-of-ai-on-the-future-of-insurance)); at a 100,000-claim/yr carrier, ICA's $400K ARR displaces $1.0M+ in ALAE savings — a 2.5x ROI before counting cycle-time and indemnity reduction benefits. The AI claims processing market is growing at 16.2% CAGR toward $0.97 billion by 2030 ([Research and Markets, 2026](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers)), and 57% of carriers are actively replacing or enhancing their claims systems ([Celent, 2024](https://www.celent.com/en/insights/916794809)) — an unprecedented window of displacement for an AI-native challenger.

ICA's regulatory framework is proactive: the system's AIS Program aligns to the NAIC Model Bulletin on AI (2023/2024), Colorado SB21-169/Reg 10-1-1, and NY DFS Circular Letter 7 (July 2024). Bias monitoring (quarterly disparate impact testing), immutable audit trails, and explicit human-in-the-loop tiers for every AI output category ensure the platform strengthens — rather than undermines — carrier compliance posture.

The founding team brings an unusual combination: 5+ years managing State Farm Auto Injury Claims (direct expertise in FNOL quality, adjuster productivity, reserve adequacy, and SIU referral) combined with academic rigor from Wake Forest's M.S. in FinTech & Analytics. This practitioner-researcher convergence produces a product that speaks the language of carrier claims leadership and a financial model that a board of directors can defend to investors and regulators alike.

**Five-year outlook:** $8.4M ARR by Year 5, EBITDA positive by Year 4, with clear cross-sell expansion into fraud/SIU intelligence, subrogation acceleration, and predictive reserve adequacy — building a compounding data network effect as every processed claim deepens the models' insurance domain knowledge.

---

## SECTION B: FULL PAPER OUTLINE

### Final Paper Structure (15+ Pages)

| Section | Title | Est. Pages | Source Material |
|---------|-------|-----------|----------------|
| Executive Summary | — | 1 | This document, Section A |
| 1 | Problem Statement & Research Objectives | 1.5 | M1, §1–2 |
| 2 | Industry Overview & Market Landscape | 2 | M1, §3; M4, §2 |
| 3 | Literature Review | 2.5 | M1, §4 |
| 4 | Methodology | 1 | M1, §5 |
| 5 | Technology & Architecture | 2 | M2, all sections |
| 6 | Data Science & AI Strategy | 1.5 | M3, §1 |
| 7 | Agile PM Framework | 1 | M3, §2 |
| 8 | Cybersecurity & Compliance | 1.5 | M3, §3 |
| 9 | Product Simulation | 1.5 | M4, §1 |
| 10 | Financial Model & Scaling Plan | 2 | M4, §2–5; ICA_PnL_Model.xlsx |
| 11 | Financial Analytics Strategy | 1.5 | M5, §1 |
| 12 | Regulatory Environment | 1.5 | M5, §2 |
| 13 | Risk Management | 1 | M5, §3 |
| 14 | ESG, Ethics & KPI Scorecard | 2 | M6, all sections |
| 15 | Synthesis & Practical Application | 1.5 | M7, Section D |
| Appendices | See list below | 3–5 | All milestone documents |
| **Total** | | **~24–27 pages** | |

---

### Section-by-Section Outline

#### Section 1: Problem Statement & Research Objectives (1.5 pp)

- **1.1** The claims operations crisis: $85.4B LAE, 60%+ carriers with zero STP, $40–$60/claim manual cost
- **1.2** The claimant experience gap: 25-min FNOL calls, opaque status, $170B premium attrition risk
- **1.3** The AI opportunity: 20–40% cost reduction, 30%+ cycle time compression
- **1.4** Five research objectives (from M1, §2) — state as numbered, measurable outcomes
- **[Faculty feedback placeholder: Are objectives scoped appropriately for a 1-year capstone? Add 4th-order measurable target?]**

#### Section 2: Industry Overview & Market Landscape (2 pp)

- **2.1** U.S. P&C market scale: $935.7B NPW, $558.8B net losses, $85.4B LAE
- **2.2** Claim frequency/severity by LOB (auto, homeowners, GL) — table format
- **2.3** LAE structure (ALAE/ULAE) and reduction opportunity
- **2.4** Claims technology vendor landscape: Guidewire, Duck Creek, Snapsheet, Tractable, Five Sigma, Hi Marley — competitive positioning
- **2.5** Market sizing: TAM ($0.97B AI claims), SAM ($115M mid-tier), SOM ($6M Year 3)

#### Section 3: Literature Review (2.5 pp)

- **3.1** AI/ML in claims: 15 cited sources, grouped by theme (from M1, §4)
- **3.2** Synthesis: What academic literature establishes vs. gaps ICA addresses
- **[Faculty feedback placeholder: Review thematic groupings; ensure diversity of publication types — suggest adding 1–2 practitioner white papers from AM Best or Swiss Re sigma]**

#### Section 4: Methodology (1 pp)

- Design thinking → Agile build → Quantitative impact modeling
- Justification for mixed-methods approach
- Role of practitioner domain expertise as validity mechanism

#### Section 5: Technology & Architecture (2 pp)

- Full technology table (from M2, §1)
- Condensed Mermaid architecture diagram
- Buy vs. build matrix — key decisions highlighted
- Payment integration: Modern Treasury / FedNow rationale

#### Section 6: Data Science & AI Strategy (1.5 pp)

- Decision taxonomy: rules → ML → LLM
- Human-in-the-loop tiers (T1–T4)
- Model evaluation framework: faithfulness, hallucination, disparate impact metrics
- SR 11-7 / NAIC AIS Program alignment

#### Section 7: Agile PM Framework (1 pp)

- Dual-track discovery diagram
- Sprint cadence + RACI
- Quarterly OKRs (Q1–Q4)
- Definition of Done

#### Section 8: Cybersecurity & Compliance (1.5 pp)

- STRIDE threat model table
- Control matrix: NIST CSF × GLBA × HIPAA × NY DFS Part 500
- LLM-specific risk mitigations (prompt injection, PII leakage, data exfiltration)
- SOC 2 + pen test roadmap

#### Section 9: Product Simulation (1.5 pp)

- Synthetic 1,000-claim cohort design and inputs
- Output table: LAE delta, cycle time, NPS lift, adjuster hours freed
- Extrapolation to 50,000 claims/carrier/yr
- Sensitivity analysis: results at 20% vs. 40% cost reduction scenario

#### Section 10: Financial Model & Scaling Plan (2 pp)

- Revenue model: per-claim + platform fee
- COGS breakdown: LLM tokens, OCR, infra, CS
- 5-year P&L summary (reference ICA_PnL_Model.xlsx)
- Gross margin trajectory: 7% → 64%
- EBITDA breakeven: Year 4
- Scaling milestones: carrier count, ARR, headcount by year
- Cross-sell roadmap

#### Section 11: Financial Analytics Strategy (1.5 pp)

- ICA unit economics: ARR, CAC, LTV, CAC payback
- Carrier ROI narrative: ALAE reduction, cycle-time elasticity, leakage reduction
- Future product analytics: predictive severity, reserve adequacy, reinsurance optimization

#### Section 12: Regulatory Environment (1.5 pp)

- Regulatory matrix (from M5, §2): 9 key regulations, ICA requirements, compliance status
- State-specific nuances: Colorado SB21-169 (life → auto expansion), NY DFS CL7
- Rate filing implications of AI in adjudication decisions
- ECOA/Reg B not applicable (P&C claims ≠ credit decision) — explicit clarification

#### Section 13: Risk Management (1 pp)

- Enterprise risk register: top 10 risks
- Risk heatmap (narrative description)
- Priority risks: R1 (hallucination → harm) and R3 (regulatory rejection)
- Mitigation ownership and monitoring KPIs

#### Section 14: ESG, Ethics & KPI Scorecard (2 pp)

- Ethical framework: utilitarian / deontological / virtue — applied to auto-deny case study
- NIST AI RMF + OECD AI Principles alignment
- ESG matrix: E (paperless, cloud efficiency), S (equity in settlement, multilingual), G (auditable AI, bias monitoring)
- TCFD / SASB Insurance / UN PSI alignment
- Full 22-KPI scorecard table

#### Section 15: Synthesis & Practical Application (1.5 pp)

- See Section D below

---

## SECTION C: FACULTY FEEDBACK PLACEHOLDERS

| Location | Placeholder | Action Required |
|----------|-------------|----------------|
| Section 1 (Objectives) | Are 5 objectives appropriate scope for 1-year capstone? | Revise scope if feedback suggests narrowing to 3 most measurable |
| Section 3 (Lit Review) | Diversity of publication types? | Add 1 AM Best or Swiss Re sigma paper if flagged |
| Section 9 (Simulation) | Is 1,000-claim synthetic cohort defensible without primary data? | If challenged, add interview data from 3 active adjusters as qualitative validation |
| Section 10 (P&L) | Are Year 1 revenue assumptions ($300K) defensible? | If challenged, cite comparable seed-stage insurtech ARR (Hi Marley raised $25M at early ARR); adjust to $150K Year 1 if conservative required |
| Section 12 (Regulatory) | Is ECOA/Reg B exclusion clearly justified? | Add 2-sentence explicit explanation: P&C claims process doesn't involve credit determination; cite 15 U.S.C. § 1691 definition of "credit transaction" |
| All sections | Citation format consistent? | Standardize to Chicago Author-Date throughout |

---

## SECTION D: PRACTICAL APPLICATION SYNTHESIS

### How State Farm + Wake Forest Converge in ICA

This capstone is unusual in that the researcher is not hypothesizing about a problem domain — they lived inside it for five years. The convergence of practitioner experience and academic rigor produces a different quality of insight than either produces alone:

**From State Farm Auto Injury Claims (5+ years, Team Manager):**
- Direct experience with FNOL intake quality — the researcher knows firsthand how a poorly structured FNOL creates downstream reserve inadequacy and litigation risk. ICA's FNOL wizard design is not academic; it mirrors the exact fields that experienced adjusters know they need at first contact.
- Adjuster productivity bottlenecks: the 3+ hours an adjuster spends on documentation per claim is not theoretical — the researcher managed teams doing exactly this work. The 50% human intervention reduction target in McKinsey's projection is believable precisely because the researcher has watched adjusters spend that time on tasks AI can now handle.
- SIU referral processes: the researcher understands from the inside what patterns trigger SIU referral and why rule-based systems miss nuanced fraud signals. ICA's T3 advisory (not T4 auto-action) design reflects practitioner judgment about when AI should support, not replace, the SIU process.
- Carrier relationships: the researcher understands that the decision-maker for a claims technology purchase is the VP of Claims or Chief Claims Officer — not the CIO. ICA's pitch materials and ROI narrative are designed for that audience.

**From Wake Forest FTA-799 (M.S. FinTech & Analytics):**
- Quantitative rigor transforms practitioner intuition into defensible financial models. The ALAE reduction calculation, cycle-time elasticity model, and unit economics framework are academic contributions that make the practitioner knowledge investable.
- Regulatory landscape: the coursework provided structured frameworks (GLBA, HIPAA, NAIC Model Bulletin analysis) that practitioners absorb organically but rarely synthesize into a coherent compliance architecture. The regulatory matrix in M5 would not exist without the academic grounding.
- Data science methodology: SR 11-7 alignment, disparate impact testing, and the model evaluation framework (faithfulness, hallucination rate) reflect academic exposure to model risk governance that claims operations practitioners typically lack.

**The synthesis:** ICA is what happens when an experienced claims professional stops accepting manual processes as inevitable and applies FinTech rigor to redesign them. The product knows what adjusters actually need (not what a product manager hypothesizes they need), and the financial model reflects how carriers actually evaluate technology purchases (LAE reduction + NPS, not abstract AI capability metrics).

---

## SECTION E: APPENDICES LIST

| Appendix | Description | Source Document |
|----------|-------------|----------------|
| A | Full System Architecture Diagram (Mermaid + narrative) | ARCHITECTURE.md |
| B | ICA 5-Year P&L Model | ICA_PnL_Model.xlsx |
| C | Regulatory Compliance Matrix | M5, §2 |
| D | Full 22-KPI Scorecard | M6, §4 |
| E | AI Prompt Library (selected prompts) | AI_PROMPTS.md |
| F | Data Model (entity-relationship) | DATA_MODEL.md |
| G | UX Flow Diagrams | UX_FLOWS.md |
| H | API Contract Reference | API_CONTRACTS.md |
| I | Risk Register (full with mitigations) | M5, §3 |
| J | Literature Citations (Chicago Author-Date, full bibliography) | M1, §4 |

---

*M7 complete — final synthesis document. Proceed to board presentation preparation.*
