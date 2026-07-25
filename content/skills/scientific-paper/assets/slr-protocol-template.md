# SLR Protocol Template

**Instructions**: Fill in all `[TODO]` placeholders before starting data collection. This document is the pre-registered protocol. Do not modify inclusion/exclusion criteria or search strings after data collection begins without documenting the change and reason.

---

## Protocol Metadata

| Field | Value |
|---|---|
| Protocol title | [TODO: Full title, e.g., "Protocol for a Systematic Literature Review on X"] |
| Version | 1.0 |
| Date | [TODO: YYYY-MM-DD] |
| Authors | William Oliveira |
| Status | Draft / Under review / Final |

---

## 1. Background and Motivation

[TODO: 2–4 sentences. Why is this topic worth an SLR? What problem exists that this SLR will help solve? What is the expected output - framework, taxonomy, evidence synthesis, gap map?]

---

## 2. Research Questions

### Primary RQ

**RQ1**: [TODO: Full research question formulated using PICO: Population, Intervention, Comparison, Outcome]

### Secondary RQs (optional)

**RQ2**: [TODO or DELETE]

**RQ3**: [TODO or DELETE]

### PICO decomposition

| PICO element | Value |
|---|---|
| Population | [TODO: e.g., "ARM-based single-board computers"] |
| Intervention | [TODO: e.g., "small language model inference"] |
| Comparison | [TODO: e.g., "different quantization levels" / "no comparison"] |
| Outcome | [TODO: e.g., "inference throughput, energy per inference, task accuracy"] |

---

## 3. Search Strategy

### 3.1 Search String

**Concept 1** (Population):
```
[TODO: e.g., "single-board computer" OR "SBC" OR "Raspberry Pi" OR "embedded hardware" OR "ARM"]
```

**Concept 2** (Intervention):
```
[TODO: e.g., "small language model" OR "SLM" OR "large language model" OR "LLM" OR "on-device inference"]
```

**Concept 3** (Outcome, if applicable):
```
[TODO: e.g., "performance" OR "benchmark" OR "evaluation" OR "energy" OR "inference"]
```

**Full combined string**:
```
([Concept 1]) AND ([Concept 2]) AND ([Concept 3])
```

**Validation**: The following known-relevant seed papers must be retrieved by this search string. If any are missed, revise the string before proceeding:
- [TODO: Paper 1 - Author, Title, Year]
- [TODO: Paper 2 - ...]
- [TODO: Paper 3 - ...]

### 3.2 Databases

| Database | URL | Date searched | Results (N) |
|---|---|---|---|
| ACM Digital Library | dl.acm.org | [TODO] | |
| IEEE Xplore | ieeexplore.ieee.org | [TODO] | |
| Scopus | scopus.com | [TODO] | |
| Web of Science | webofscience.com | [TODO] | |
| arXiv (cs.SE + cs.AI) | arxiv.org | [TODO] | |
| Google Scholar | scholar.google.com | [TODO] | |
| **Total** | | | |

### 3.3 Snowballing

Backward snowballing: review reference lists of all included papers (round 1).
Forward snowballing: check "cited by" for all included papers via Google Scholar (round 1).
Repeat until no new papers found.

---

## 4. Inclusion Criteria

A study is **included** if ALL of the following are true:

| # | Criterion |
|---|---|
| IC1 | [TODO: e.g., "The study reports empirical results on inference performance of LLMs/SLMs on hardware"] |
| IC2 | [TODO: e.g., "The study uses ARM-based hardware or explicitly comparable embedded hardware"] |
| IC3 | [TODO: e.g., "The study reports at least one of: throughput (tokens/s), latency, energy consumption, or accuracy"] |
| IC4 | Published in English |
| IC5 | Published between [TODO: YYYY] and [TODO: YYYY] (or: no year restriction) |
| IC6 | Peer-reviewed (conference, journal, workshop) OR preprint with ≥ [TODO: N] citations |

---

## 5. Exclusion Criteria

A study is **excluded** if ANY of the following is true:

| # | Criterion |
|---|---|
| EC1 | [TODO: e.g., "Study only evaluates cloud-based or server-grade hardware (no embedded/SBC)"] |
| EC2 | [TODO: e.g., "Study evaluates LLMs with >7B parameters only (beyond SLM scope)"] |
| EC3 | Secondary study (another SLR, mapping study, or survey) |
| EC4 | Short paper (<4 pages) without sufficient methodological detail |
| EC5 | Duplicate publication - keep most complete version |
| EC6 | Not accessible in full text |

---

## 6. Selection Process

### Stage 1: Title and Abstract Screening

- Reviewer 1: [TODO: Name]
- Reviewer 2: [TODO: Name or "TBD"]
- Disagreements resolved by: discussion, or third reviewer if unresolved
- Inter-rater agreement target: Cohen's κ ≥ 0.6

### Stage 2: Full-Text Screening

- Same reviewers
- Record exclusion reason for every excluded paper
- Document disagreements

### PRISMA Flow (fill after selection)

```
Records from databases (N=__)
  + Records from snowballing (N=__)
= Total records (N=__)
    ↓ Duplicates removed (N=__)
Records after deduplication (N=__)
    ↓ Excluded at title/abstract (N=__)
      Reasons: [TODO list top reasons]
Full-text assessed (N=__)
    ↓ Excluded at full-text (N=__)
      Reasons: [TODO list top reasons]
Studies included (N=__)
```

---

## 7. Quality Assessment

### Instrument

Score each criterion: 1 (yes), 0.5 (partial), 0 (no or unclear).

| # | Criterion |
|---|---|
| Q1 | Is the study design clearly described? |
| Q2 | Is the data collection method appropriate for the RQ? |
| Q3 | Is the sample/subject selection described and justified? |
| Q4 | Are validity threats discussed? |
| Q5 | Are results described with appropriate statistical analysis? |
| Q6 | Is the study reproducible (sufficient detail to replicate)? |
| Q7 | Is the study context described? |

**Minimum threshold for inclusion**: Total score ≥ [TODO: e.g., 3.0] / 7.

### Reviewers

QA performed by: [TODO: Reviewer 1] with validation by [TODO: Reviewer 2 or "spot-check on 20% of papers"].

---

## 8. Data Extraction

### Extraction Form

Extract the following fields for each included study:

| Field | Description |
|---|---|
| Paper ID | Sequential ID (P01, P02, ...) |
| Authors, Year, Title, Venue | Full reference |
| Study type | benchmarking / experiment / case study / survey / mixed |
| Hardware | SBC model(s), SoC, RAM configuration |
| Models evaluated | Model name, parameter count, quantization |
| Inference engine | llama.cpp / Ollama / TFLite / other |
| Task types | structured output / classification / summarization / QA / other |
| Primary metrics reported | tokens/s / TTFT / latency / energy / accuracy / other |
| N (runs per condition) | Number of measurement runs |
| Statistical analysis | Described? Which tests? |
| Dataset used | Name, size, availability |
| Results summary | Key quantitative finding (brief) |
| Data availability | public / on-request / none / partial |
| Relevant to RQ1 | yes / no + short note |
| Relevant to RQ2 | yes / no + short note |
| Gaps noted | What the paper did NOT study that is relevant |

### Extraction process

- Extracted by: [TODO]
- Validated by: [TODO] (10% random sample minimum)

---

## 9. Synthesis Plan

| RQ | Synthesis method | Rationale |
|---|---|---|
| RQ1 | [TODO: e.g., "Descriptive narrative + frequency table by metric type"] | [TODO] |
| RQ2 | [TODO: e.g., "Gap analysis matrix (hardware × model × metric)"] | [TODO] |

### Gap Analysis Matrix Structure

Rows: [TODO: e.g., hardware platform tier]
Columns: [TODO: e.g., model size range × quantization × task type]
Cells: ✓ well-studied / △ partially studied / ✗ not studied

---

## 10. Protocol Deviations Log

| Date | Change | Reason |
|---|---|---|
| (none at start) | | |

*Document any deviations from this protocol after data collection begins.*

---

## 11. References

[TODO: Key papers that motivated this SLR and foundational methodology references]

- Kitchenham, B., & Charters, S. (2007). Guidelines for performing Systematic Literature Reviews in Software Engineering. Technical Report EBSE 2007-001.
- Dybå, T., & Dingsøyr, T. (2008). Empirical studies of agile software development: A systematic review. Information and Software Technology, 50(9–10), 833–859.
- Moher, D., et al. (2009). Preferred Reporting Items for Systematic Reviews and Meta-Analyses: The PRISMA Statement. PLOS Medicine.
