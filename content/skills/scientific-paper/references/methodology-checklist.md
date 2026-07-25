# Methodology Checklist by Study Type

Reference for auditing research rigor. Run the appropriate section for the paper under review.

---

## 1. Case Study (Runeson & Höst 2009, Yin 2018)

### Research Questions
- [ ] RQs are clearly stated and bounded (not "how does X work" but "what types of failures emerge when X under constraint Y")
- [ ] Each RQ is **specific** (narrow scope), **measurable** (observable evidence), and **relevant** (addresses a gap)
- [ ] The number of RQs is appropriate (2–5 for most case studies)
- [ ] All RQs are answered in the paper (cross-check abstract contributions vs. RQ answers)

### Subject System / Case Selection
- [ ] The case (system, project, organization) is described in sufficient detail for replication judgment
- [ ] Rationale for case selection is stated (purposive, theoretical, convenience - and why appropriate)
- [ ] Generalizability claims are proportional to case breadth (1 case ≠ universal claim)

### Data Collection
- [ ] Data sources are identified (interviews, documents, logs, code history, observations)
- [ ] Triangulation is used (≥2 independent data sources for key claims)
- [ ] Raw data or derived evidence is available or described precisely enough to audit
- [ ] Git history, commit logs, or issue trackers: if used as evidence, commit count / time span / methodology for classifying commits is described

### Data Analysis
- [ ] Analysis procedure is described (how themes/categories were derived)
- [ ] Coding scheme is defined if qualitative (open coding, axial coding, or deductive categories)
- [ ] Inter-rater reliability is reported if ≥2 coders (Cohen's κ or similar)
- [ ] Saturation point is discussed if applicable

### Validity
- [ ] Internal validity: confounders acknowledged (researcher involvement, selection of evidence)
- [ ] External validity: generalizability scope is explicitly bounded
- [ ] Construct validity: operationalizations (e.g., "AI-related commit" definition) are defined
- [ ] Conclusion validity: claims do not exceed what evidence supports

---

## 2. Benchmarking / Measurement Study

### Research Questions
- [ ] RQs are **comparative** or **relational** (not just descriptive) where appropriate
- [ ] Each metric maps to at least one RQ (no orphaned metrics)
- [ ] Primary metric is defined and justified (e.g., energy per inference, not just tokens/s)

### Hardware / Software Configuration
- [ ] All hardware is fully described: SoC, RAM, memory bandwidth, OS version
- [ ] All software is versioned: engine commit hash, model version, compiler flags
- [ ] System isolation is described: CPU governor, swap state, background services

### Experimental Design
- [ ] **N ≥ 30** runs per condition is the default; deviations require justification
- [ ] **Warm-up runs** (discarded) are used and their count is specified (minimum 3)
- [ ] **Run order is randomized** to avoid temporal/thermal order bias
- [ ] **Thermal cool-down** between runs is defined (e.g., wait until temp ≤ idle + 5°C)
- [ ] Cold start vs. warm cache distinction is documented
- [ ] Outlier handling criteria are defined a priori (not post-hoc)

### Measurement
- [ ] Each measurement instrument is described (power meter model, polling interval, etc.)
- [ ] Measurement error / precision is reported where relevant
- [ ] `drop_caches` or equivalent used before cold-start measurements
- [ ] Energy per inference = Watts × time (not just average power)

### Statistical Analysis
- [ ] Mean AND standard deviation (or IQR) reported for each metric
- [ ] Appropriate test for normality before choosing parametric vs. non-parametric
- [ ] Correlation: Pearson (linear + normal) or Spearman (non-linear / non-normal) - justify choice
- [ ] Effect size reported alongside p-values (Cohen's d, rank-biserial r)
- [ ] Bonferroni correction or FDR if multiple comparisons

### Reproducibility
- [ ] Fixed random seeds documented
- [ ] Prompts / inputs are fixed and publicly available
- [ ] Datasets are versioned and publicly accessible
- [ ] Scripts are published (or described precisely enough to re-implement)

---

## 3. Controlled Experiment (Wohlin et al. 2012)

### GQM Validation
- [ ] **Goal**: defined using GQM template - "Analyze X for the purpose of Y with respect to Z from the viewpoint of W in the context of V"
- [ ] **Questions**: each question operationalizes the goal
- [ ] **Metrics**: each metric answers exactly one question

### Subjects & Objects
- [ ] Population is defined; sample is described (convenience, stratified, random)
- [ ] Object (artifact, task, system) is described in detail
- [ ] Subject characteristics reported (experience level, background)

### Design
- [ ] Experimental design named: between-subjects, within-subjects, factorial, Latin square
- [ ] Treatment assignment: random assignment documented
- [ ] Blocking variables are identified and controlled
- [ ] Pilot study conducted and reported

### Threats
- [ ] Instrumentation: measurement instruments validated before main study
- [ ] Maturation: effect of learning/fatigue addressed (e.g., counterbalancing)
- [ ] Selection: group equivalence tested (pre-test if applicable)

---

## 4. Systematic Literature Review (Kitchenham & Charters 2007)

### Protocol
- [ ] A written protocol exists before data collection begins
- [ ] Protocol includes: RQs, search string, databases, inclusion/exclusion criteria, quality assessment form, data extraction form, synthesis method

### Search
- [ ] Search string covers synonyms, acronyms, and spelling variants
- [ ] Boolean structure: (concept1 OR synonym1) AND (concept2 OR synonym2)
- [ ] Databases searched: ACM DL, IEEE Xplore, Scopus, Web of Science, arXiv (at minimum)
- [ ] Search date is recorded; search is reproducible
- [ ] Forward and backward snowballing from seed set

### Selection
- [ ] Inclusion/exclusion criteria applied in stages (title/abstract → full text)
- [ ] ≥2 reviewers for full-text selection; disagreements resolved by discussion or third reviewer
- [ ] PRISMA flow diagram included

### Quality Assessment
- [ ] Quality assessment instrument is defined (e.g., Dybå & Dingsøyr checklist)
- [ ] Minimum quality threshold for inclusion is stated
- [ ] Quality scores reported per paper

### Synthesis
- [ ] Narrative synthesis or meta-analysis: choice is justified
- [ ] Heterogeneity discussed if quantitative synthesis attempted
- [ ] Gaps identified in a structured way

---

## 5. Survey / Questionnaire Study

- [ ] Survey instrument (questionnaire) is available (supplementary material or appendix)
- [ ] Pilot survey conducted and used to refine questions
- [ ] Response rate reported; non-response bias discussed
- [ ] Likert scale analysis: use median/mode not mean; use Mann-Whitney not t-test
- [ ] Demographic profile of respondents reported
- [ ] Closed vs. open questions: open questions coded and coding scheme described

---

## 6. Technical Report / Experience Report

- [ ] Scope is clearly bounded: what project, team size, time period
- [ ] Claims are scoped to the described context (not generalized without qualification)
- [ ] Lessons learned are actionable and specific (not "communication is important")
- [ ] Comparison to existing approaches or literature (even brief)

---

## 7. Position Paper / Vision Paper

- [ ] Core position/vision is stated in one sentence in the abstract
- [ ] Evidence for feasibility is provided (preliminary results, analogies, existence proofs)
- [ ] Distinguishes "what is" from "what could be" consistently
- [ ] Challenges and risks are acknowledged (not only the positive vision)
- [ ] Related positions in the field are engaged with (not strawmanned)

---

## Cross-Cutting: All Paper Types

### Abstract Quality
- [ ] Structured abstract where venue requires it: Background / Objective / Method / Results / Conclusion
- [ ] Abstract contains a specific result (number, finding, recommendation) - not just "we study X"
- [ ] Abstract does not contain claims absent from the paper body

### Contribution Alignment
- [ ] Contributions listed in introduction match contributions delivered in paper body
- [ ] Contributions are numbered and cross-referenceable to specific sections

### Related Work
- [ ] Related work **positions the paper** - explains what existing work does NOT cover
- [ ] Related work is recent (check for papers from last 2–3 years)
- [ ] Related work table or taxonomy used where ≥5 related papers

### Introduction
- [ ] Problem statement is concrete (not "research is important")
- [ ] Gap is clearly identified with evidence
- [ ] Contributions are listed explicitly (numbered list)
- [ ] Paper roadmap included ("The remainder is organized as follows...")

### Conclusion
- [ ] Conclusion does not introduce new claims absent from the paper
- [ ] Conclusion answers each RQ explicitly
- [ ] Future work is specific (not "more research is needed")
