# M1 — Problem Statement, Research Objectives, Literature Review & Methodology
**Wake Forest FTA-799 Capstone | Insurance Claim Assistant (ICA)**
**Author:** Capstone Researcher | **Date:** 2025 | **Status:** Milestone 1

---

## 1. Problem Statement

The U.S. property and casualty insurance industry paid **$558.8 billion in net losses and $85.4 billion in loss adjustment expenses (LAE) in 2024** — yet the operational infrastructure processing those payments remains largely manual, opaque, and adversarial to claimants ([NAIC 2024 P&C Annual Report](https://content.naic.org/sites/default/files/2024-annual-property-casualty-and-title-insurance-industries-analysis-report.pdf)). The First Notice of Loss (FNOL) process averages 20–30 minutes per call, standard claims cost $40–$60 to process before AI deployment, and more than 60% of U.S. insurers report zero straight-through processing capability in claims operations ([Aite-Novarica, 2023, cited in Decerto](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers)). Claimants — who are experiencing financial and emotional stress — face confusing policy language, opaque status updates, and call-center dependency at the precise moment they need clarity. Meanwhile, carriers face LAE ratios that compound combined ratios above 100, making claims operations a material drag on profitability. The Insurance Claim Assistant (ICA) addresses both sides: an AI-powered platform that reduces carrier LAE by automating intake, summarization, and next-best-action delivery, while simultaneously delivering a consumer-grade experience to claimants across Auto, Homeowners, and General Liability lines.

---

## 2. Research Objectives

1. **Quantify the LAE reduction achievable** through AI-assisted FNOL intake, evidence summarization, and adjuster workflow automation in a representative cohort of U.S. P&C claims, establishing a measurable baseline for carrier ROI.
2. **Design and validate a responsible AI governance framework** for claims processing that satisfies NAIC Model Bulletin on AI Systems (2023/2024), Colorado SB21-169/Reg 10-1-1, and NY DFS Circular Letter 7 (July 2024) — including bias testing, hallucination monitoring, and human-in-loop controls.
3. **Develop and test a B2B2C go-to-market model** targeting mid-size carriers, TPAs, and MGAs as platform buyers, quantifying TAM/SAM/SOM and identifying the pricing structure (per-claim SaaS + platform fee) that achieves EBITDA breakeven by Year 3.
4. **Assess the equity impact** of AI-assisted claims on underserved claimant populations, including time-to-settlement variance by zip code and demographic segment, establishing KPIs aligned to SASB Insurance Sustainability Accounting Standards and UN Principles for Sustainable Insurance.
5. **Evaluate a technology build-vs.-buy-vs.-partner decision matrix** across core components (LLM gateway, OCR, identity, e-signature, payments), deriving a cost- and risk-optimal architecture for a seed-stage insurtech serving the carrier market.

---

## 3. Industry Overview: U.S. P&C Claims Landscape

### 3.1 Market Scale

The U.S. P&C industry crossed **$1.06 trillion in direct written premiums for the first time in 2024** — an 8.0% year-over-year increase driven by rate hardening in Auto and Homeowners ([NAIC / S&P Global, 2024](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers)). Total net losses and LAE incurred reached **$644.3 billion** ($558.8B losses + $85.4B LAE), compared to $640.7 billion in 2023 ([NAIC 2024 P&C Annual Report](https://content.naic.org/sites/default/files/2024-annual-property-casualty-and-title-insurance-industries-analysis-report.pdf)).

| Line | 2024 Incurred Losses ($M) | Key Driver |
|------|--------------------------|------------|
| Private passenger auto liability | $132,012 | Social inflation, bodily injury severity |
| Private passenger auto physical damage | $89,500 | Repair cost inflation, parts shortages |
| Commercial auto liability | $33,573 | Nuclear verdicts, litigation financing |
| Commercial auto physical damage | $7,969 | Fleet CAT losses |
| **Total auto** | **$263,054** | |
| Homeowners / Farmowners | ~$120,000 est. | CAT events, reconstruction cost inflation |
| General liability | ~$85,000 est. | Social inflation, PFAS/CAT liability |

*Sources: [III / NAIC data via S&P Global](https://www.iii.org/fact-statistic/facts-statistics-auto-insurance); homeowners and GL figures are INTERNAL ASSUMPTIONS based on NAIC aggregate minus auto lines.*

### 3.2 Claims Frequency and Severity

For personal auto — the largest claims volume driver — 2024 data from ISO/Verisk shows ([III, 2024](https://www.iii.org/fact-statistic/facts-statistics-auto-insurance)):
- **Bodily injury claim frequency:** 0.80 claims per 100 earned car years; **average severity: $28,278**
- **Property damage claim frequency:** 2.50 per 100 earned car years; **average severity: $6,770**
- **Collision claim frequency:** 4.16 per 100 insured vehicle years; **average severity: $5,489**

For homeowners, **5.3% of insured homes had a claim in 2023** — approximately 1-in-18 insured homes annually ([III, 2023](https://www.iii.org/fact-statistic/facts-statistics-homeowners-and-renters-insurance)). U.S. property claims volume rose **36% in 2024**, propelled by a 113% increase in catastrophe claims ([III / Verisk, 2025](https://insuranceindustryblog.iii.org/claims-volume-up-36in-2024-climate-costs-litigation-drive-trend/)).

### 3.3 LAE Structure and Reduction Opportunity

The industry's **2024 LAE-to-net-losses ratio is approximately 15.3%** ($85.4B / $558.8B). LAE is a composite of:
- **ALAE (Allocated Loss Adjustment Expense):** Costs directly attributable to a specific claim — independent adjuster fees, expert witnesses, defense counsel, appraisal costs.
- **ULAE (Unallocated Loss Adjustment Expense):** Overhead of the claims department — staff salaries, claims management systems, facilities.

AI impact on LAE is substantial: McKinsey projects AI systems can **reduce claims processing costs 20–40%** and **cut cycle time up to 30%**, while reducing human intervention by 50% ([McKinsey, "Insurance 2030", 2021](https://www.mckinsey.com/industries/financial-services/our-insights/insurance-2030-the-impact-of-ai-on-the-future-of-insurance)). Accenture's research found AI-based claims systems achieve up to **73% increase in claims process cost efficiency** in production deployments ([Accenture](https://www.accenture.com/content/dam/accenture/final/accenture-com/document/Accenture-Why-AI-In-Insurance-Claims-And-Underwriting.pdf)).

### 3.4 Operational Benchmarks

| Metric | Pre-AI Benchmark | With AI | Source |
|--------|-----------------|---------|--------|
| Cost per standard claim | $40–$60 | $25–$36 | [Research & Markets / Decerto, 2026](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers) |
| Cost per FNOL intake | $15–$25 | $5–$8 | [Research & Markets, 2026](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers) |
| Routine resolution time | 7–10 days | 24–48 hours | [Research & Markets, 2026](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers) |
| STP rate (industry average) | <10% | 35%+ (top carriers) | [Aite-Novarica 2023, via Decerto](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers) |
| Human intervention reduction | Baseline | ~50% | [McKinsey, 2021](https://www.mckinsey.com/industries/financial-services/our-insights/insurance-2030-the-impact-of-ai-on-the-future-of-insurance) |

### 3.5 Competitive and Technology Landscape

57% of P&C carriers are actively replacing, enhancing, or recently replaced their core claims system ([Celent CIO Priorities Survey, 2024](https://www.celent.com/en/insights/916794809)), indicating a window of platform displacement. The insurtech claims technology sub-market reached **$0.46 billion in 2025**, projected to reach **$0.97 billion by 2030** at a 16.2% CAGR ([Research and Markets, 2026, via Decerto](https://www.decerto.com/eu/post/ai-claims-processing-the-complete-2026-guide-for-us-carriers)).

---

## 4. Literature Review

### Theme 1: AI and Machine Learning in Claims Processing

**Sheehan, B., Castignani, G., Masello, L., & Bhattacharya, S. (2025).** "AI revolution in insurance: bridging research and reality." *Frontiers in Artificial Intelligence.* DOI: [10.3389/frai.2025.1568266](https://pmc.ncbi.nlm.nih.gov/articles/PMC12014612/). Comprehensive review finding key gaps between academic AI claims research and real-world carrier implementation; emphasizes need for explainable AI and regulatory alignment.

**Govindaswamy Subbian, R. (2025).** "Enhancing Operational Efficiency in Claims Processing Through Technology." *AJRCOS.* DOI: [10.9734/ajrcos/2025/v18i3604](https://journalajrcos.com/index.php/AJRCOS/article/view/604). Quasi-experimental study at a production insurer documenting **50% reduction in claim cycle time, 40% reduction in operational costs, 25% improvement in fraud detection** after AI/automation deployment.

**Sudabathula, B. (2025).** "Revolutionizing Insurance: The Impact of AI on Claims Processing." *CSEIT.* DOI: [10.32628/cseit251112302](https://ijsrcseit.com/index.php/home/article/view/CSEIT251112302). Reviews ML, computer vision, and NLP integration for real-time fraud detection, automated document verification, and intelligent data extraction across the claims lifecycle.

**Saikia, D., Barua, R., Gourisaria, M., et al. (2024).** "Machine Learning Enhancements for Car Insurance Claim Prediction." *IEEE ICCCNT.* DOI: [10.1109/ICCCNT61001.2024.10724028](https://ieeexplore.ieee.org/document/10724028/). Demonstrates XGBoost achieving 0.84 accuracy in auto claims prediction; provides methodology for training ML models on carrier loss data.

**Velmurugan, K., et al. (2023).** "Data Driven Analysis of Insurance Claims Using Machine Learning Algorithm." *IJARSCT.* DOI: [10.48175/ijarsct-9689](http://ijarsct.co.in/Paper9689.pdf). Validates agile development of AI claims systems with iterative improvement cycles; documents personalized service, one-click FNOL, and intelligent claims processing outcomes.

**Hanafy, M. & Ming, R. (2021).** "Machine Learning Approaches for Auto Insurance Big Data." *Risks.* DOI: [10.3390/RISKS9020042](https://www.mdpi.com/2227-9091/9/2/42/pdf?version=1614263593). Establishes ML baseline for auto claims frequency/severity prediction; foundational for ICA's timeline estimation and reserve adequacy modules.

### Theme 2: Customer Experience in Claims

**McKinsey & Company (2021).** "Insurance 2030 — The impact of AI on the future of insurance." [McKinsey.com](https://www.mckinsey.com/industries/financial-services/our-insights/insurance-2030-the-impact-of-ai-on-the-future-of-insurance). Landmark industry analysis projecting 20–40% cost reduction and 30% cycle time improvement through AI in claims. Projects AI will handle 70% of routine claims by 2030 with minimal human intervention.

**Accenture (2022).** "Poor Claims Experiences Could Put Up to $170B of Global Insurance Premiums at Risk." [Accenture Newsroom](https://newsroom.accenture.com/news/2022/poor-claims-experiences-could-put-up-to-170b-of-global-insurance-premiums-at-risk-by-2027-according-to-new-accenture-research). Quantifies the revenue-at-risk from poor claims UX — $170B at risk globally by 2027 — establishing the business case for claimant-facing AI tools as retention instruments.

**Adhwaryu, H. (2025).** "Real-Time Data Ecosystems in Insurance." *IJAEM.* DOI: [10.35629/5252-0704286293](https://ijaem.net/issue_dcp/Real%20Time%20Data%20Ecosystems%20in%20Insurance%20A%20Comprehensive%20Analysis%20of%20Claims%20Processing%20and%20Policy%20Management%20Transformation.pdf). Examines Apache Kafka/Flink architectures enabling real-time claims data processing; directly informs ICA's SSE push notification and event-driven architecture choices.

### Theme 3: Regulatory Landscape for AI in Insurance

**NAIC (2023/2024).** "Model Bulletin: Use of Algorithms, Predictive Models, and Artificial Intelligence Systems by Insurers." [NAIC.org](https://content.naic.org/sites/default/files/cmte-h-big-data-artificial-intelligence-wg-ai-model-bulletin.pdf.pdf). The foundational U.S. regulatory framework for AI in insurance; establishes AIS Program requirements, governance, risk management, third-party oversight, and consumer protection expectations.

**Colorado Division of Insurance (2023).** "Regulation 10-1-1: Governance and Risk Management Framework Requirements for Life Insurers' Use of ECDIS, Algorithms, and Predictive Models." Effective November 14, 2023. [Colorado SB21-169 Summary](https://www.credo.ai/blog/colorado-sb21-169-8-things-you-need-to-know-about-colorados-new-ai-insurance-regulation). Requires annual bias testing attestation, inventory of all algorithmic models, and disparate impact testing across protected classes.

**New York DFS (2024).** "Insurance Circular Letter No. 7 (2024): Use of Artificial Intelligence Systems and External Consumer Data." July 11, 2024. [NYDFS.gov](https://www.dfs.ny.gov/industry-guidance/circular-letters/cl2024-07). Extends AI governance requirements to all insurance lines in New York; requires quantitative testing, governance framework, and consumer disclosure obligations.

### Theme 4: Fraud Detection, Privacy, and Responsible AI

**Luciano, E., Cattaneo, M., & Kenett, R. (2023).** "Adversarial AI in Insurance: Pervasiveness and Resilience." *arXiv.* DOI: [10.48550/arXiv.2301.07520](https://arxiv.org/ftp/arxiv/papers/2301/2301.07520.pdf). Analyzes adversarial attack vectors on insurance AI — directly applicable to ICA's threat model for prompt injection and model manipulation risks.

**Dong, P., Quan, Z., Edwards, B., et al. (2024).** "Privacy-Enhancing Collaborative Information Sharing through Federated Learning — A Case of the Insurance Industry." *arXiv.* DOI: [10.48550/arXiv.2402.14983](https://arxiv.org/abs/2402.14983). Proposes federated learning for privacy-preserving multi-carrier data sharing; informs ICA's long-term data network strategy without requiring raw claim data portability.

**Alam, A. & Prybutok, V.R. (2024).** "Use of responsible artificial intelligence to predict health insurance claims in the USA using machine learning algorithms." *Exploration of Digital Health Technologies.* DOI: [10.37349/edht.2024.00009](https://www.explorationpub.com/uploads/Article/A10119/10119.pdf). Evaluates XGBoost for claims prediction with explicit fairness constraints; methodology applicable to ICA's bias monitoring framework for BI claims.

**Dorle, A. (2025).** "An AI-Powered and Blockchain-Integrated Model for Automated Claims Processing and Fraud Detection." *IJITM.* DOI: [10.29070/sxye6f09](https://ignited.in/index.php/ijitm/article/view/15926). Demonstrates blockchain-integrated audit trails for immutable claims records; informs ICA's future roadmap for audit log tamper-evidence.

---

## 5. Methodology

ICA employs a **mixed-methods research design** combining three complementary approaches:

### 5.1 Design Thinking (Discovery Phase)

User research rooted in design thinking frames the claimant and adjuster experience problem before technology solutions are proposed. Persona development, journey mapping, and prototype testing with actual insurance claimants (recruited from adjuster networks and prior State Farm claims operations experience) produce empirically grounded feature requirements rather than assumed needs. Design thinking is appropriate here because claims processing is deeply contextual — the emotional and procedural experience of filing an auto claim differs fundamentally from a GL claim, and AI-generated outputs (FNOL drafts, policy explanations) must meet professional credibility standards that only practitioner review can validate.

### 5.2 Agile Build (Construction Phase)

A Scrum-based sprint cadence with dual-track discovery executes product development iteratively, exposing working software to simulated claim scenarios at the end of each 2-week sprint. This methodology is appropriate because AI product development requires rapid experimentation with prompt engineering, model selection, and output evaluation — requirements that cannot be fully specified in advance. The agile approach also enables continuous integration of regulatory feedback (NAIC bulletin interpretations, carrier legal review) without derailing delivery.

### 5.3 Quantitative Impact Modeling (Validation Phase)

A synthetic cohort simulation (N=1,000 claims, designed to mirror the distribution of Auto, HO, and GL claims in a mid-size carrier) computes LAE reduction, cycle-time delta, and claimant NPS lift using cited industry benchmarks as baseline inputs. Monte Carlo sensitivity analysis tests outcomes across assumption ranges, producing a defensible range of ROI projections rather than point estimates. Financial models (5-year P&L, Unit Economics) apply standard SaaS metrics (ARR, gross margin, CAC/LTV) alongside insurance-industry-specific metrics (ALAE saved per claim, loss-cost-to-cycle-time elasticity) to present a multidimensional picture of value creation.

### 5.4 Justification

The mixed-methods design mirrors best practice in applied FinTech research (combining practitioner observation with quantitative modeling) and satisfies the Wake Forest FTA-799 rubric requirement for both qualitative rigor and quantitative defensibility. The researcher's 5+ years managing auto injury claims at State Farm — including direct supervisory responsibility for FNOL intake quality, adjuster productivity, and reserve adequacy — provides domain validity that academic surveys alone cannot replicate.

---

*Page 1 of approximately 5 | Next: M2 — Technology & Architecture*
