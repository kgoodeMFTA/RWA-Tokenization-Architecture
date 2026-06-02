# M6 — ESG/CSR Framework, Ethical Decision-Making & KPI Scorecard
**Wake Forest FTA-799 Capstone | Insurance Claim Assistant (ICA)**
**Author:** Capstone Researcher | **Date:** 2025 | **Status:** Milestone 6

---

## 1. Sustainable Ethical Decision-Making Framework

### 1.1 Multi-Lens Ethical Analysis

Ethical reasoning in AI claims systems requires simultaneous consideration of competing stakeholder interests. ICA applies three complementary philosophical frameworks to every material product decision:

| Framework | Core Question | Application to ICA |
|-----------|-------------|-------------------|
| **Utilitarian** | What produces the greatest good for the greatest number? | Optimize for aggregate claimant welfare + carrier efficiency; accept trade-off if AI speeds 95% of claims at some cost to 5% of edge cases — provided no systematic harm |
| **Deontological (Kantian)** | Does this treat people as ends, not means? Do the rules hold universally? | No AI decision should reduce a claimant to a data point; every adverse action must be explainable and contestable; rules must apply equally regardless of claimant demographics |
| **Virtue Ethics** | Would a person of good character make this decision? | Adjuster integrity as the standard; AI acts as a virtuous assistant, not a replacement; system should make adjusters better at their job, not eliminate judgment |

### 1.2 Applied Ethical Case Study: Auto-Deny Without Human Review

**Scenario:** ICA's SIU fraud score identifies a claim as high-probability fraud (score ≥ 0.85). Should the system auto-deny the claim without human adjuster review?

**Utilitarian Analysis:**
- *For auto-deny:* Faster resolution, lower LAE, reduced fraud loss — benefits the majority of policyholders through lower premiums.
- *Against:* False positive rate of even 5% on a 0.85 threshold means 1 in 20 "fraud flags" is a legitimate claimant wrongly denied. At scale, this is thousands of claimants harmed annually.
- **Utilitarian verdict:** Auto-deny fails the utilitarian test when false positive rate × scale of harm exceeds fraud prevention benefit.

**Deontological Analysis:**
- Auto-deny without human review treats the claimant as a data pattern, not a person with rights.
- The NAIC Model Unfair Claims Settlement Practices Act explicitly prohibits denying claims without reasonable investigation ([NAIC Model Act #900](https://content.naic.org/)).
- **Deontological verdict:** Auto-deny categorically fails — it violates the claimant's right to fair investigation and violates existing law.

**Virtue Ethics Analysis:**
- An experienced adjuster of good character would never deny a claim solely on a score without reviewing evidence and providing the claimant an opportunity to respond.
- **Virtue ethics verdict:** Auto-deny fails.

**ICA Resolution:** The fraud score triggers a **T3 advisory** (flag shown to adjuster; adjuster makes the decision). The system provides: (a) the score, (b) the features driving the score, (c) a suggested investigation checklist (e.g., "Request recorded statement; verify repair shop affiliation"). Auto-denial is never permitted (T4 block). This aligns to NIST AI RMF Govern function and OECD AI Principle of accountability.

### 1.3 AI Ethics Frameworks Referenced

| Framework | Relevant ICA Principles |
|-----------|------------------------|
| **NIST AI Risk Management Framework (AI RMF 1.0)** | Govern: AIS Program, accountability ownership; Map: claim type risk tiering; Measure: hallucination/bias KPIs; Manage: T1–T4 human oversight tiers |
| **OECD AI Principles (2019, updated 2023)** | Transparency (explainable AI outputs with citations); Robustness (model monitoring + drift detection); Accountability (audit log for all AI actions); Fairness (disparate impact testing) |
| **IEEE Ethically Aligned Design** | Human agency preserved (no T4 AI-only decisions); privacy by design (PII redaction first); well-being (faster fair settlements benefit claimants) |

---

## 2. ESG/CSR Alignment

### 2.1 Environmental (E)

| ICA Practice | Environmental Benefit | Measurement |
|-------------|----------------------|-------------|
| **Paperless FNOL and evidence management** | Eliminates physical claim file: ~50–200 pages/claim × millions of claims industry-wide | Pages eliminated per claim; carrier paper/printing cost reduction |
| **Digital evidence over physical inspection** | AI evidence summarization reduces need for in-person adjuster inspections for routine claims, cutting vehicle miles driven | Estimated inspector miles avoided per AI-summarized claim |
| **Cloud-native infrastructure (ECS Fargate)** | Serverless containers use compute only when needed; AWS data centers are 4× more energy-efficient than typical enterprise data center ([AWS Sustainability](https://aws.amazon.com/sustainability/)) | GHG scope 3 emissions per claim processed |
| **Faster settlements → reduced rental car duration** | Faster cycle time for auto claims reduces rental vehicle days during repair | Rental days avoided per BI claim resolved faster |

### 2.2 Social (S)

| ICA Practice | Social Impact | Target Population |
|-------------|-------------|------------------|
| **Plain-English policy explainer** | Eliminates information asymmetry between carrier and claimant; democratizes access to policy comprehension previously requiring an attorney | All claimants; highest impact for first-time claimants, lower-income policyholders |
| **Multilingual support (roadmap v2)** | NLP translation layer enables non-English speakers to navigate claims in their primary language | Spanish, Mandarin, Vietnamese-speaking claimants — segments historically underserved by carrier portals |
| **24/7 AI availability** | Claimants can file FNOL and get status at 2 AM during a crisis (car accident, house fire) without waiting for business hours | All claimants; highest impact for hourly workers who cannot make calls during business hours |
| **Faster settlement for economic recovery** | 12.5-day resolution (vs. 25-day baseline) means faster replacement of totaled vehicles and home repairs for claimants living paycheck-to-paycheck | Lower-income policyholders for whom delayed settlement has cascading financial impact |
| **Equity in time-to-settlement** | KPI monitoring for settlement time variance by zip code and demographic segment; flags systematic disparities for carrier review | Minority and lower-income policyholders historically experiencing longer claim cycles |
| **Accessible UX design** | WCAG 2.1 AA compliance; screen-reader compatible components; clear visual hierarchy in shadcn/ui | Claimants with visual or motor disabilities |

### 2.3 Governance (G)

| ICA Practice | Governance Impact |
|-------------|-----------------|
| **Auditable AI decisions** | Every AI output stored with: model version, prompt hash, token count, timestamp, actor. Adjuster edit history retained. Full provenance for regulatory examination. |
| **Bias monitoring** | Quarterly disparate impact analysis by protected class; public-facing commitment to bias monitoring as part of carrier contracts |
| **Consumer redress mechanism** | Claimants can flag AI outputs they believe are inaccurate; flags route to adjuster review queue; response within 2 business days |
| **Transparent AI disclosure** | Every AI-generated artifact labeled "AI Draft — Review Required"; claimant notice when AI system used in their claim processing |
| **Third-party audit** | SOC 2 Type II (Year 1); annual penetration test; model risk review aligned to SR 11-7; carrier audit rights in contract |

---

## 3. TCFD / SASB Insurance Standard / UN PSI Alignment

### 3.1 TCFD — How ICA Moves Carrier Climate KPIs

The [Task Force on Climate-related Financial Disclosures (TCFD)](https://www.fsb-tcfd.org/) framework identifies climate risks as physical (extreme weather claims) and transition (policy change). ICA directly moves carrier TCFD-relevant KPIs:

| TCFD Pillar | Carrier TCFD Metric | ICA Contribution |
|-------------|--------------------|-----------------| 
| **Governance** | Board oversight of climate risk in claims | ICA's claims analytics dashboard surfaces CAT claim volume by geography, enabling board-level climate risk reporting |
| **Risk Management** | CAT exposure concentration | ICA's FNOL intake standardizes peril classification (wind, flood, fire) enabling carrier CAT model inputs |
| **Metrics & Targets** | Claims-cycle time during CAT events | ICA's AI-assisted FNOL reduces CAT surge bottleneck; carrier can report faster claimant service during declared disasters |
| **Strategy** | Physical risk impact on loss ratios | ICA's claim volume analytics by peril and geography feed carrier climate scenario analysis |

### 3.2 SASB Insurance Sustainability Standard

The [SASB Insurance Standard (FN-IN)](https://www.ifrs.org/groups/international-sustainability-standards-board/about-the-issb/) covers insurers' exposure to environmental and social risks. ICA helps carriers improve SASB-reportable metrics:

| SASB Topic | Metric | ICA Impact |
|-----------|--------|-----------|
| Transparent Information and Fair Advice | Customer complaint rate | ICA's plain-English policy explainer reduces complaints from policy misunderstanding |
| Transparent Information and Fair Advice | % claims settled without dispute | Faster, clearer settlement communication reduces disputed claims |
| Policies Designed to Incentivize Responsible Behavior | N/A (underwriting focus) | ICA supports carriers in demonstrating fair claims treatment |
| Environmental Risk Exposure | CAT net incurred losses | ICA's FNOL categorization improves accuracy of CAT loss reporting |

### 3.3 UN Principles for Sustainable Insurance (PSI)

The [UNEP FI Principles for Sustainable Insurance](https://www.unepfi.org/industries/insurance/the-principles-for-sustainable-insurance/) provide four principles for incorporating ESG into insurance:

| PSI Principle | ICA Alignment |
|--------------|--------------|
| **Principle 1:** Embed ESG into decision-making | ICA embeds equity KPIs (time-to-settlement by demographic) and fairness metrics (bias testing) into every model deployment decision |
| **Principle 2:** Work with clients/brokers on ESG awareness | ICA's plain-English policy explainer educates claimants on coverage; digital FNOL guides responsible documentation practices |
| **Principle 3:** Work with governments/regulators on ESG | ICA's regulatory alignment (NAIC, CO, NY DFS) and proactive bias testing support regulators' ESG objectives for fair claims |
| **Principle 4:** Demonstrate accountability + transparency | Annual carrier impact reports: claims processed, average cycle time, equity KPIs, AI fairness audit results |

---

## 4. KPI Scorecard

| # | KPI | Definition | Target | Frequency | Owner | Measurement Source |
|---|-----|-----------|--------|-----------|-------|-------------------|
| **Performance KPIs** | | | | | | |
| K1 | FNOL Completion Rate | % of started FNOL wizards resulting in submitted claim | ≥ 85% | Weekly | Product | `wizard_started` vs `fnol_submitted` events |
| K2 | Time to FNOL Submission | Median minutes from wizard start to submission | ≤ 8 min | Weekly | Product | Event timestamps |
| K3 | Claims Cycle Time (platform) | Median days from FNOL to claim close for ICA-processed claims | ≤ 15 days (Year 1 pilot) | Monthly | Customer Success | `claim_events` table |
| K4 | AI FNOL Draft Acceptance Rate | % of AI FNOL drafts accepted by adjuster without major edit | ≥ 40% | Sprint | AI/ML Lead | `draft_action = accepted` |
| K5 | Evidence OCR Success Rate | % of evidence uploads completing OCR successfully | ≥ 95% | Weekly | Engineering | `ocr_status = complete` / total |
| K6 | Policy Explainer Satisfaction | Claimant/adjuster rating | ≥ 4.2 / 5.0 | Continuous | Product | In-app rating |
| K7 | Status Board DAU | % of active claims with daily status board page view | ≥ 60% | Daily | Product | Page view analytics |
| **AI Quality KPIs** | | | | | | |
| K8 | LLM Hallucination Rate | % of sampled LLM outputs containing unsupported factual claims | ≤ 2% | Weekly (5% sample) | AI/ML Lead | Manual review + automated fact-check |
| K9 | Summary Faithfulness Score | % of summary claims verifiable in source document | ≥ 95% | Weekly | AI/ML Lead | QA pipeline |
| K10 | PII Redaction Failure Rate | % of LLM API calls where PII detected post-redaction in audit logs | 0.0% | Continuous | Security Lead | Log scanner alert |
| K11 | Prompt Injection Block Rate | % of injection attempts detected and blocked | ≥ 99.5% | Continuous | Engineering | WAF + application logs |
| **Equity & Fairness KPIs** | | | | | | |
| K12 | Disparate Impact Ratio — SIU Score | P(SIU flag | minority group) / P(SIU flag | majority group) | 0.80–1.25 | Quarterly | AI/ML Lead | Model output × demographic inference |
| K13 | Equal Opportunity Difference — SIU Score | TPR difference across race/ethnicity groups | ≤ 0.05 | Quarterly | AI/ML Lead | Model evaluation pipeline |
| K14 | Time-to-Settlement Equity | Difference in median cycle time between zip code income quartiles (Q1 vs Q4) | ≤ 3 days difference | Quarterly | Customer Success | `claim_events` + zip code income data (FFIEC) |
| K15 | Customer Complaint Rate | Claims resulting in formal complaint to carrier/DOI | ≤ 0.5% | Monthly | Customer Success | Carrier complaint log |
| **Financial KPIs** | | | | | | |
| K16 | ARR | Annual Recurring Revenue | See M4 targets | Monthly | CEO/Finance | CRM + billing system |
| K17 | Gross Margin % | (Revenue − COGS) / Revenue | 54% by Year 3 | Monthly | Finance | P&L |
| K18 | CAC Payback Period | Months to recover Customer Acquisition Cost | ≤ 18 months by Year 3 | Quarterly | Finance | CRM + P&L |
| K19 | Carrier LAE Delta | Reported reduction in LAE per claim by carrier vs. pre-ICA baseline | ≥ $12/claim | Quarterly | Customer Success | Carrier data + ICA claims volume |
| **Security & Compliance KPIs** | | | | | | |
| K20 | System Uptime | Platform availability | ≥ 99.5% | Continuous | Engineering | Uptime monitoring (PagerDuty / BetterUptime) |
| K21 | SOC 2 Findings Open | Open material findings from SOC 2 audit | 0 critical, ≤ 2 moderate | Quarterly | Security Lead | Audit report |
| K22 | Pen Test Critical/High Findings Remediated | % of P0/P1 pen test findings resolved within 30 days | 100% | Per test cycle | Engineering Lead | Pen test tracking |

---

*Page 1 of approximately 4 | Next: M7 — Final Synthesis Outline & Executive Summary*
