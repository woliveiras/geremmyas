# Error Catalog - Scientific Paper Errors

Structured catalog of methodological, validity, and writing errors. Used during `review`, `critique`, and `peer-review` modes. Errors are ranked by severity: **Critical** (likely rejection), **Major** (major revision), **Minor** (polish).

---

## Category 1: Conclusion Validity Errors

Errors that undermine whether conclusions follow from the data.

### 1.1 Correlation Claimed as Causation [Critical]
- **Pattern**: "X causes Y" when study only shows correlation between X and Y.
- **Detection**: Any causal language ("causes", "leads to", "results in", "is responsible for") without a controlled experiment or structural causal model.
- **Fix**: Replace causal language with correlational ("is associated with", "co-occurs with", "predicts in this context"); or redesign study as controlled experiment.

### 1.2 Underpowered Claims [Critical]
- **Pattern**: Claiming statistical significance with N < 30 without power analysis, or claiming a trend is meaningful from 3–5 data points.
- **Detection**: N stated in methodology; claims proportional to N.
- **Fix**: Add power analysis justifying N; soften claims to "preliminary evidence" or "exploratory finding."

### 1.3 Post-Hoc Rationalization [Major]
- **Pattern**: Hypothesis is stated as if it preceded data collection, but the hypothesis was clearly formed after seeing the data (HARKing: Hypothesizing After Results are Known).
- **Detection**: Hypotheses that match results exactly, without prior literature support; no pre-registration.
- **Fix**: Clearly label exploratory vs. confirmatory analysis. Pre-register if possible. Be transparent about hypothesis origin.

### 1.4 Selective Reporting [Critical]
- **Pattern**: Only favorable conditions or metrics are reported; unfavorable results are omitted or buried.
- **Detection**: Results section covers only conditions that support the main claim; negative or null results absent.
- **Fix**: Report all conditions. Discuss negative results explicitly in Discussion or Threats.

### 1.5 Multiple Comparisons Without Correction [Major]
- **Pattern**: Many statistical tests performed; p-values reported without Bonferroni, FDR, or similar correction.
- **Detection**: Count of statistical tests vs. mention of correction.
- **Fix**: Apply Bonferroni correction (conservative) or Benjamini-Hochberg FDR (less conservative). Report corrected p-values.

---

## Category 2: Internal Validity Errors

Errors in the relationship between study variables.

### 2.1 Uncontrolled Confounders [Critical]
- **Pattern**: Study compares A vs. B but multiple factors vary simultaneously (e.g., comparing tools by developer AND task type simultaneously).
- **Detection**: Experimental design section; whether conditions differ in more than the treatment variable.
- **Fix**: Isolate the variable; use blocking design; or acknowledge confounder explicitly as a threat.

### 2.2 Selection Bias [Major]
- **Pattern**: Cases/participants selected in a way that systematically favors one outcome (e.g., studying only successful projects to understand success factors).
- **Detection**: Case selection rationale; representativeness of sample.
- **Fix**: Describe selection procedure; use random or stratified sampling; acknowledge survivor bias where applicable.

### 2.3 Instrumentation Bias [Major]
- **Pattern**: Measurement instrument changes during the study, or different instruments used for different conditions.
- **Detection**: Methodology section; consistency of measurement across conditions.
- **Fix**: Use identical instruments for all conditions; pilot-test instruments before main data collection.

### 2.4 Maturation / Learning Effect [Major]
- **Pattern**: In within-subjects studies, participants improve across tasks due to practice, not treatment.
- **Detection**: Order of tasks; counterbalancing described.
- **Fix**: Counterbalance task order; include practice/warmup tasks that are discarded.

### 2.5 Researcher Bias (Subjectivity) [Major]
- **Pattern**: Researcher is also the subject (practitioner case study), introducing potential favorable interpretation.
- **Detection**: Single-author case studies of own work.
- **Fix**: Acknowledge explicitly in Threats; use automated/objective metrics where possible; seek external review of interpretations.

---

## Category 3: Construct Validity Errors

Errors in how concepts are operationalized.

### 3.1 Proxy Metric Without Validation [Major]
- **Pattern**: A metric is used as a proxy for a construct, but the proxy relationship is not validated.
- **Example**: Using "tokens/s" as proxy for "user experience quality" without validating that faster inference actually improves UX.
- **Detection**: Metric definitions; gap between what is measured and what is claimed.
- **Fix**: Either validate the proxy (cite prior work), directly measure the construct, or bound claims to the measurable proxy only.

### 3.2 Definition Drift [Major]
- **Pattern**: A key term is defined one way in Section 2 but used differently in Section 4.
- **Detection**: Track key term definitions across sections.
- **Fix**: Define key terms in a single place; use exactly that definition throughout.

### 3.3 Operationalization Without Justification [Minor]
- **Pattern**: A category, threshold, or classification is used without explaining why that threshold was chosen.
- **Example**: "Commits with ≥1 LLM-related file change are classified as AI commits" - why ≥1?
- **Fix**: Cite precedent for threshold, or explain rationale explicitly.

---

## Category 4: External Validity Errors

Errors in generalizability.

### 4.1 Overgeneralization From Single Case [Critical]
- **Pattern**: Finding from one system/project/team is stated as a general principle without qualification.
- **Detection**: Language like "this shows that all X will Y" in a single-case study.
- **Fix**: Add explicit scope qualifier: "in the context of this case study, our findings suggest..."; add a section on context factors that bound generalizability.

### 4.2 Sample Unrepresentative of Target Population [Major]
- **Pattern**: Study uses convenience sample (e.g., students) but claims apply to practitioners.
- **Detection**: Demographic profile vs. claim scope.
- **Fix**: Acknowledge population gap in Threats; cite validation studies or limit claims to the sampled population.

### 4.3 Technology/Version Lock-In [Minor]
- **Pattern**: Findings are strongly tied to specific library/framework versions that may change rapidly.
- **Detection**: Dependency versions in methodology; how quickly the ecosystem moves.
- **Fix**: State exact versions studied; acknowledge temporal scope limitation.

---

## Category 5: Writing and Presentation Errors

### 5.1 Abstract–Conclusion Mismatch [Major]
- **Pattern**: Abstract claims a finding that is stated differently (or more weakly) in Conclusion, or vice versa.
- **Detection**: Compare abstract's claims sentence-by-sentence with conclusion.
- **Fix**: Write abstract last; ensure every claim in abstract has a direct counterpart in conclusion.

### 5.2 Unanswered Research Questions [Critical]
- **Pattern**: A RQ is stated in Introduction but not explicitly answered in Results or Discussion.
- **Detection**: Map each RQ → section that answers it. If no section, it is unanswered.
- **Fix**: Add a "summary of RQ answers" table, or ensure each RQ has a dedicated subsection.

### 5.3 Contribution List Inflation [Major]
- **Pattern**: Introduction lists 5 contributions but the paper only fully delivers 2–3 of them.
- **Detection**: Trace each listed contribution to the specific section where it is delivered.
- **Fix**: Remove or qualify contributions that are not fully substantiated; or add the missing content.

### 5.4 Related Work Without Positioning [Major]
- **Pattern**: Related work describes what papers do but never explains what THIS paper does that related work does NOT.
- **Detection**: Absence of phrases like "unlike X, we...", "X studied Y but not Z; we address Z", "no prior work has...".
- **Fix**: Add an explicit gap statement at the end of each related work subsection; add a positioning table or matrix.

### 5.5 Hedging Language in Claims [Minor]
- **Pattern**: Excessive hedging that makes contributions unclear: "results may suggest that it is possible that X could potentially...".
- **Fix**: Make the strength of evidence explicit: "in the context of this study, X was observed in 9/10 cases."

### 5.6 Missing Replication Package / Data Availability [Major]
- **Pattern**: Paper describes experiments but provides no link to data, scripts, or replication package.
- **Detection**: Check for data availability statement and URLs.
- **Fix**: Publish dataset and scripts; add a Data Availability section.

---

## Category 6: LaTeX and Double-Blind Errors

### 6.1 Author-Identifying Self-Citation [Critical for double-blind]
- **Pattern**: Paper cites "our previous work [X]" or "in [our GitHub repo]" in blind submission.
- **Detection**: Search for "our", first-person possessives near `\cite{}`, and URLs to personal repos.
- **Fix**: Use `\ifanonymized` toggle to hide self-references; replace with "[anonymous]" in blind version.

### 6.2 Author Metadata in PDF [Critical for double-blind]
- **Pattern**: PDF metadata (Author field in Document Properties) contains author names.
- **Detection**: `pdfinfo paper.pdf | grep Author`.
- **Fix**: Add `\hypersetup{pdfauthor={}}` in preamble when `\ifanonymized` is true.

### 6.3 Broken Cross-References [Minor]
- **Pattern**: `\ref{}` or `\cref{}` keys resolve to "??" in compiled PDF (undefined label).
- **Detection**: LaTeX build log warnings; grep for `??` in PDF.
- **Fix**: Ensure `\label{}` exists for every `\ref{}` target; build twice.

### 6.4 Undefined `\cite{}` Keys [Minor]
- **Pattern**: `\cite{somekey}` appears in .tex but `somekey` is absent from .bib.
- **Detection**: Run `validate-bib.py`; LaTeX build log warns "Citation `X' undefined".
- **Fix**: Add missing BibTeX entry or correct the key.

### 6.5 Missing Required BibTeX Fields [Minor]
- **Pattern**: BibTeX entry lacks required fields for its type (e.g., `@article` without `journal`).
- **Detection**: `validate-bib.py` field check; BibTeX warnings in build log.
- **Fix**: Complete the entry; use DOI lookup or Semantic Scholar for authoritative metadata.
