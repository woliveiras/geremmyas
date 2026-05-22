# Systematic Literature Review Protocol Guide

Based on Kitchenham & Charters (2007) "Guidelines for performing Systematic Literature Reviews in Software Engineering."

---

## Phase 0: Motivation and Scope

Before starting, answer:
1. **Why an SLR?** - What primary question requires aggregating evidence across studies? (not just "I want to know what people have done on X")
2. **Is an SLR appropriate?** - Are there enough primary studies? (rule of thumb: if you can already enumerate all relevant papers, a mapping study may suffice)
3. **What will the SLR produce that primary studies don't?** - Synthesis, meta-analysis, gap map, or framework?

---

## Phase 1: Research Questions

### RQ Formulation (PICO framework for SE)
- **P** - Population: what system/developer/project type?
- **I** - Intervention: what technology, technique, process?
- **C** - Comparison: compared to what? (or: in what context?)
- **O** - Outcome: what effect, quality attribute, or finding?

**Good SLR RQ examples:**
- "What inference performance metrics are used to evaluate on-device LLMs on embedded hardware?" (descriptive)
- "What is the relationship between model quantization level and task accuracy in structured output generation?" (relational)
- "What methodological gaps exist in existing benchmarks of SLMs on ARM-based single-board computers?" (gap analysis)

**Checklist:**
- [ ] Each RQ is answerable by reading and synthesizing primary studies
- [ ] RQs are specific enough that two independent researchers would extract the same data
- [ ] RQs collectively cover the SLR scope without overlap

---

## Phase 2: Search Strategy

### 2.1 Search String Construction

Template structure:
```
(concept1 OR synonym1 OR acronym1) AND (concept2 OR synonym2 OR acronym2)
```

Steps:
1. Extract main concepts from your RQs (PICO: Population + Intervention)
2. For each concept, list synonyms, acronyms, and spelling variants
3. Combine concepts with AND; variants with OR
4. Test the string on a small set of known-relevant papers - if any are missed, add the missing terms

**Example** (SLMs on embedded hardware):
```
("small language model" OR "SLM" OR "large language model" OR "LLM" OR "on-device LLM")
AND
("single-board computer" OR "SBC" OR "embedded hardware" OR "edge computing" OR "Raspberry Pi" OR "ARM")
AND
("inference" OR "performance" OR "benchmark" OR "evaluation")
```

### 2.2 Databases

Mandatory (for SE SLRs):
| Database | URL | Notes |
|---|---|---|
| ACM Digital Library | dl.acm.org | Strong for SE venues (ICSE, FSE, ISSTA, MSR) |
| IEEE Xplore | ieeexplore.ieee.org | Strong for TSE, EMSE, ICSME |
| Scopus | scopus.com | Broad; may need institution access |
| Web of Science | webofscience.com | Strong for impact factor journals |
| arXiv (cs.SE, cs.AI) | arxiv.org | Preprints; latest work |
| Google Scholar | scholar.google.com | Catch-all; use for snowballing |

Optional (domain-specific):
| Database | Use for |
|---|---|
| Semantic Scholar | ML/AI papers; good API |
| Papers With Code | Benchmarking papers with results |
| DBLP | Computer science venue index |

### 2.3 Snowballing

**Backward snowballing**: For each included paper, check its reference list for additional relevant papers.

**Forward snowballing**: For each included paper, check papers that cite it (Google Scholar "Cited by").

Run at least one round of snowballing after initial search. Stop when no new papers are found.

**Record**: For each snowballed paper, note which included paper it was found through.

---

## Phase 3: Study Selection

### 3.1 Inclusion Criteria

Define a priori (before starting selection). Examples:
- Written in English
- Published in peer-reviewed venue (conference, journal, workshop) OR preprint with ≥ N citations
- Published between YYYY and YYYY
- Studies [population] in the context of [domain]
- Reports [specific metric or outcome relevant to RQs]

### 3.2 Exclusion Criteria

Define a priori. Examples:
- Secondary studies (other SLRs, mapping studies, surveys) - unless you are meta-analyzing them
- Short papers (< N pages) without sufficient methodological detail
- Gray literature without peer review (blog posts, vendor white papers)
- Papers that do not report quantitative results on the target metric
- Duplicate publications (same study published in multiple venues - keep most complete version)

### 3.3 Selection Process

**Stage 1: Title and Abstract Screening**
- Two reviewers independently screen all results
- Apply inclusion/exclusion criteria at title/abstract level
- Record decision + reason for exclusion
- Compute inter-rater agreement (Cohen's κ ≥ 0.6 = substantial agreement)
- Resolve disagreements by discussion; third reviewer for unresolved cases

**Stage 2: Full-Text Screening**
- Retrieve full text of papers passing Stage 1
- Apply criteria at full-text level
- Same two-reviewer process; document exclusion reasons
- Record final included set

**PRISMA Flow:**
```
Records identified via databases (N=X)
    + Records from snowballing (N=Y)
    = Total records (N=X+Y)
        ↓ Duplicates removed (N=D)
    Records screened (N=X+Y-D)
        ↓ Excluded at title/abstract (N=E1, reasons)
    Full-text assessed (N=X+Y-D-E1)
        ↓ Excluded at full-text (N=E2, reasons)
    Studies included (N=final)
```

---

## Phase 4: Quality Assessment

### Quality Assessment Instrument (Dybå & Dingsøyr adapted for SE)

Score each item 0 (no), 0.5 (partial), 1 (yes):

| # | Item |
|---|---|
| Q1 | Is the study design clearly described? |
| Q2 | Is the data collection method appropriate for the RQ? |
| Q3 | Is the sample/subject selection described and justified? |
| Q4 | Are validity threats discussed? |
| Q5 | Are results described with appropriate statistical analysis? |
| Q6 | Is the study reproducible (enough detail to replicate)? |
| Q7 | Is the study context described (generalizability assessment possible)? |

**Minimum threshold**: Total score ≥ 3.0 (out of 7) for inclusion. Adjust based on your domain.

---

## Phase 5: Data Extraction

### Data Extraction Form Template

Define before extraction. Each field should answer exactly one RQ or sub-question:

```
Paper ID:
Title:
Authors:
Venue:
Year:
Study Type: [benchmarking / experiment / case study / survey]
Hardware/Platform: [SBC model, SoC, RAM]
Models Evaluated: [model names, parameter counts, quantization]
Inference Engine: [llama.cpp / Ollama / TensorFlow Lite / etc.]
Primary Metrics: [tokens/s / TTFT / energy / accuracy / etc.]
Task Types: [structured output / free-form / classification / etc.]
N (runs per condition):
Statistical Analysis Used: [yes/no + method]
Results Summary: [key quantitative findings]
Dataset / Benchmark Used:
Data Availability: [public / on-request / none]
Relevant to RQ1: [yes/no + notes]
Relevant to RQ2: [yes/no + notes]
Notes / Limitations:
```

---

## Phase 6: Synthesis

### Synthesis Methods by RQ type

| RQ type | Appropriate synthesis |
|---|---|
| Descriptive (what has been studied?) | Narrative synthesis + frequency tables |
| Comparative (which approach is better?) | Vote-counting or meta-analysis (if data allows) |
| Relational (what factors correlate with outcome?) | Narrative + correlation table |
| Gap analysis (what is missing?) | Taxonomy / matrix of coverage |

### Gap Analysis Matrix

Useful for identifying research gaps:

```
Rows: hardware platforms / contexts
Cols: models, metrics, tasks studied
Cells: ✓ (well-studied), △ (partially studied), ✗ (gap)
```

### Statistical Meta-Analysis (when applicable)

Only if: ≥3 studies report the same outcome metric on comparable conditions.
- Check for heterogeneity (I² statistic)
- If I² > 75%, use narrative synthesis instead of pooled estimate
- Report both fixed-effects and random-effects estimates

---

## Phase 7: Write-Up Requirements

- [ ] Protocol section describes all phases (can reference published protocol document)
- [ ] PRISMA flow diagram with exact counts
- [ ] Appendix: full search strings per database with date
- [ ] Appendix: full list of included studies (ID + reference)
- [ ] Appendix: quality assessment scores per study
- [ ] Appendix: data extraction table (or link to supplementary material)
- [ ] Inter-rater reliability reported for selection and quality assessment
- [ ] Threats to validity section addresses: search completeness, selection bias, data extraction reliability

---

## Common SLR Mistakes

1. **Search string too narrow**: Missing synonyms or domain-specific terms; test against known seed set.
2. **Single reviewer selection**: Inter-rater agreement is mandatory; otherwise selection is subjective.
3. **No snowballing**: Database search alone misses conference papers and recent preprints.
4. **Synthesis is just a list**: "Paper A did X, paper B did Y" is not synthesis - add: what does this mean together?
5. **No gap statement**: The value of an SLR is explicitly identifying what is NOT known. A map without gaps is incomplete.
6. **PRISMA numbers don't add up**: Every paper must be accounted for at every stage.
