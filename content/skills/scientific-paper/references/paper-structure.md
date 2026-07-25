# Paper Structure Templates by Type

Templates for section structure, with required vs. optional markings and quality checklist per section.

---

## Repo Conventions (all paper types)

```
papers/YYYY-short-slug/
├── README.md           # Abstract + metadata table + build instructions
├── paper.tex           # Main manuscript
├── title-page.tex      # Separate title page for double-blind submission packet
├── references.bib      # Bibliography
└── scripts/            # Data collection / analysis scripts (MIT licensed)
```

README metadata table template:
```markdown
| Field    | Value |
|----------|-------|
| Status   | Draft / Under Review / Published |
| arXiv    | [XXXX.XXXXX](https://arxiv.org/abs/XXXX.XXXXX) |
| DOI      | (if published) |
| Date     | Month YYYY |
| Venue    | Journal/Conference name |
```

---

## 1. Practitioner Case Study

**Target venues**: EMSE, JSS, ICSE (SEIP track), FSE (industry track), MSR

```
Abstract (structured: Background / Objective / Method / Results / Conclusion)
1. Introduction
   1.1 Motivation and Problem Statement
   1.2 Research Questions (numbered list)
   1.3 Contributions (numbered list)
   1.4 Paper Organization ("The remainder is organized as follows...")
2. Background and Related Work
   2.1 [Core technology/concept 1]
   2.2 [Core technology/concept 2]
   2.3 Related Work + gap statement
3. Study Design
   3.1 Subject System (what was studied, scale, context)
   3.2 Data Collection (sources: commits, logs, interviews, etc.)
   3.3 Data Analysis (classification procedure, coding scheme)
   3.4 Validity Considerations (brief forward reference to Section N)
4. [Study-specific section, e.g., Timeline / Development Arc]
5. Findings
   5.1 RQ1: [question]
   5.2 RQ2: [question]
   5.3 RQ3: [question]
6. Discussion
   6.1 Implications for Practitioners
   6.2 Design Heuristics / Recommendations
   6.3 Limitations
7. Threats to Validity
   7.1 Internal Validity
   7.2 External Validity
   7.3 Construct Validity
   7.4 Conclusion Validity
8. Conclusion
   8.1 Summary of Findings (answer each RQ)
   8.2 Future Work (specific, not "more research is needed")
References
Appendix [optional: full prompt text, raw data tables]
```

**Section quality checklist:**
- [ ] Introduction: problem is concrete, gap is supported, contributions are numbered
- [ ] Study Design: sufficient detail for replication judgment
- [ ] Findings: each subsection title contains the RQ answer, not just the RQ
- [ ] Discussion: separate implications for practitioners vs. researchers
- [ ] Threats: uses Wohlin framework categories (see threats-to-validity.md)
- [ ] Conclusion: answers each RQ explicitly, future work is specific

---

## 2. Benchmarking / Measurement Study

**Target venues**: EMSE, MSR, ICSE, ISSTA, IEEE TSE

```
Abstract (structured)
1. Introduction
   1.1 Motivation
   1.2 Research Questions
   1.3 Contributions
   1.4 Paper Organization
2. Background
   2.1 [Domain background]
   2.2 Related Benchmarking Studies (gap: what existing studies do NOT measure)
3. Study Design
   3.1 Research Questions
   3.2 Subjects (hardware, software, models - exact versions)
   3.3 Metrics and Measurements
   3.4 Experimental Protocol (warm-up, N, randomization, cool-down, system isolation)
   3.5 Datasets and Inputs (fixed, versioned, publicly available)
   3.6 Statistical Analysis Plan (tests chosen, correction for multiple comparisons)
4. Implementation
   4.1 Measurement Infrastructure
   4.2 Validation of Measurement (calibration, pilot runs)
5. Results
   5.1 RQ1: [Descriptive statistics + visualization]
   5.2 RQ2: [Statistical analysis + effect sizes]
   5.3 RQ3: [...]
   5.4 Supplementary Analysis [optional]
6. Discussion
   6.1 Interpretation of Findings
   6.2 Practical Implications
   6.3 Comparison to Related Work
7. Threats to Validity
8. Conclusion
References
Appendix: Raw data tables, full configuration details, replication package URL
```

**Key sections for benchmarking:**
- [ ] Section 3.4 must specify: N per condition, warm-up runs (count), run order (randomized/blocked), thermal policy, system isolation steps
- [ ] Section 3.6 must name the statistical test and when it is used
- [ ] Section 5: every result includes mean AND SD (or median AND IQR); no orphaned metrics
- [ ] Appendix: replication package URL with scripts, datasets, and environment specs

---

## 3. Systematic Literature Review (SLR)

**Target venues**: IST, EMSE, JSS, CSUR, IEEE TSE

```
Abstract (structured)
1. Introduction
   1.1 Motivation
   1.2 Research Questions
   1.3 Contributions (what this review provides that primary studies don't)
   1.4 Organization
2. Review Protocol
   2.1 Research Questions
   2.2 Search Strategy (string + databases)
   2.3 Inclusion and Exclusion Criteria
   2.4 Quality Assessment Criteria
   2.5 Data Extraction Form
   2.6 Synthesis Method
3. Results of Search
   3.1 PRISMA Flow (figure)
   3.2 Characteristics of Included Studies (year, venue, method, context)
4. Results
   4.1 RQ1: [synthesis]
   4.2 RQ2: [synthesis]
   4.3 [...]
5. Discussion
   5.1 Cross-RQ Synthesis
   5.2 Research Gaps and Future Directions
   5.3 Implications for Practice
   5.4 Implications for Research
6. Threats to Validity
7. Conclusion
References
Appendix: Full list of included studies, quality assessment scores, search strings per database
```

**Key requirements:**
- [ ] PRISMA flow diagram with exact counts at each stage
- [ ] Search string reproducible (show exact string used per database)
- [ ] Inter-rater reliability for selection and quality assessment
- [ ] Each RQ has a defined synthesis method (narrative, thematic, vote-counting, etc.)

---

## 4. Controlled Experiment

**Target venues**: EMSE, TSE, ICSE, FSE, EASE

```
Abstract (structured)
1. Introduction
   1.1 Motivation
   1.2 Hypotheses (H0 and H1 stated formally)
   1.3 Contributions
2. Background and Related Work
3. Experiment Design
   3.1 Goal (GQM template)
   3.2 Research Questions and Hypotheses
   3.3 Variables (independent, dependent, controlled)
   3.4 Subjects (population, sample, recruitment)
   3.5 Objects (tasks, artifacts)
   3.6 Instrumentation (forms, environments, measurements)
   3.7 Procedure (training, task sequence, debriefing)
   3.8 Analysis Procedure (statistical tests, α level)
4. Execution
   4.1 Pilot Study
   4.2 Main Study (deviations from protocol, if any)
5. Results
   5.1 Descriptive Statistics
   5.2 Hypothesis Testing
   5.3 Effect Size
6. Discussion
   6.1 Interpretation
   6.2 Relation to Related Work
7. Threats to Validity
8. Conclusion
References
Appendix: Materials (tasks, questionnaires), raw data
```

---

## 5. Survey Study

```
Abstract
1. Introduction (motivation, RQs, contributions)
2. Related Work
3. Survey Design
   3.1 Research Questions
   3.2 Instrument Design (questionnaire, pilot)
   3.3 Target Population and Sampling
   3.4 Data Collection Procedure
   3.5 Analysis Method
4. Results
   4.1 Response Rate and Demographic Profile
   4.2 RQ1: [results]
   4.3 RQ2: [results]
5. Discussion
6. Threats to Validity
7. Conclusion
Appendix: Questionnaire instrument
```

---

## 6. Technical / Experience Report

```
Abstract
1. Introduction (context, scope, what makes this report valuable)
2. System / Project Description
3. Challenges Encountered
   3.1 [Challenge 1] - symptoms, root cause, solution
   3.2 [Challenge 2] - ...
4. Lessons Learned (actionable, specific, scoped)
5. Comparison to Related Approaches (brief)
6. Limitations and Scope
7. Conclusion
```

---

## 7. Position / Vision Paper

```
Abstract
1. Introduction (the position in one sentence; why it matters now)
2. Background and Motivation (evidence the current approach has a problem)
3. The Position (what should change / what should be built)
4. Feasibility Evidence (preliminary results, existence proofs, analogies)
5. Challenges and Risks (honest acknowledgment)
6. Research Agenda (concrete research questions this position opens)
7. Conclusion
```
