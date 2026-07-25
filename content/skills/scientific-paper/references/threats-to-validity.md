# Threats to Validity - Wohlin Framework for SE Research

Based on Wohlin et al. "Experimentation in Software Engineering" (2012) and Runeson & Höst "Guidelines for conducting and reporting case study research in software engineering" (2009).

---

## Framework Overview

| Category | Core question | Most relevant for |
|---|---|---|
| **Construct validity** | Are we measuring what we claim to measure? | All study types |
| **Internal validity** | Is the observed effect caused by the treatment? | Experiments, case studies |
| **External validity** | Can findings be generalized beyond this study? | All study types |
| **Conclusion validity** | Are our statistical conclusions correct? | Quantitative studies |

---

## Construct Validity

**Threat**: The operationalization of the dependent or independent variable does not capture the intended theoretical construct.

### Common threats and mitigations in SE research

| Threat | Example | Mitigation |
|---|---|---|
| Mono-operation bias | Using tokens/s as the only measure of "inference efficiency" | Use multiple metrics (tokens/s, TTFT, energy/inference) |
| Mono-method bias | Using only commit logs to understand developer behavior | Triangulate with interviews or surveys |
| Hypothesis guessing | Subjects behave differently because they know what is being measured | Blind conditions; use automated measurements |
| Evaluation apprehension | Survey respondents give socially desirable answers | Anonymous surveys; emphasize no right/wrong answers |
| Inadequate preoperational explication | "AI-related commit" defined vaguely | Define the classification criterion explicitly before data collection |

### Writing template

```
Construct Validity. We operationalized [concept] as [metric/measurement]. This risks
mono-operation bias, as a single metric may not fully capture [concept]. We mitigated
this by [e.g., using N additional metrics, triangulating with...]. The definition of
[key term] was established a priori as [definition] to avoid post-hoc redefinition.
```

---

## Internal Validity

**Threat**: Observed effects are caused by confounding variables rather than the treatment.

### Common threats and mitigations

| Threat | Example | Mitigation |
|---|---|---|
| Selection bias | Comparing two teams where Team A is more experienced | Random assignment; match on confounders; block design |
| Maturation | Participants improve over time in a within-subjects study | Counterbalance order; include practice trials (discarded) |
| History | External event occurs between pre- and post-test | Control group; short study duration |
| Instrumentation | Measurement tool changes between conditions | Use identical instruments for all conditions |
| Diffusion of treatment | Control group learns from treatment group | Physical separation; no communication between groups |
| Compensatory rivalry | Control group tries harder to compensate | Blind participants to group assignment where possible |
| Researcher bias | Researcher analyzes data with knowledge of hypothesis | Blind analysis; automated measurement where possible |
| Thermal throttling (benchmarking) | CPU/SoC throttles under sustained load, inflating latency | Thermal cool-down protocol; log temperature curves; flag throttled runs |
| Order effects (benchmarking) | First model run is slower due to cold cache | Warm-up runs (minimum 3, discarded); randomize run order |

### For practitioner case studies (self-study)

The primary internal validity threat in practitioner case studies is **researcher as instrument**: the researcher who built the system is also analyzing it. Mitigations:
1. Use objective, automated evidence (Git history, build logs, CI results) rather than recollection.
2. Document decision timeline contemporaneously (commits, notes) rather than retrospectively.
3. Acknowledge interpretive role explicitly; invite independent review of key classifications.

### Writing template

```
Internal Validity. The primary threat is [specific threat]. We mitigated this by
[mitigation]. [If unmitigated]: We acknowledge that [threat] may have influenced
[specific aspect]; future studies should [control for X / use Y design].
```

---

## External Validity

**Threat**: Findings cannot be generalized beyond the specific study context.

### Common threats and mitigations

| Threat | Example | Mitigation |
|---|---|---|
| Population validity | Study uses CS students; claims apply to industry practitioners | Be explicit about population; replicate with different populations |
| Ecological validity | Lab task does not reflect real-world complexity | Use realistic tasks; study in situ |
| Technology/version lock-in | Findings tied to specific library version that changes rapidly | State exact versions; acknowledge temporal scope |
| Single-case generalizability | 1 project/company; broad claims | Explicitly bound claims to context; compare to similar contexts |
| Platform specificity (benchmarking) | Results measured on Raspberry Pi 5 only | State explicitly; list context factors that may affect transferability |

### For benchmarking studies

External validity is bounded by:
- Hardware platform (SoC family, memory bandwidth tier, thermal envelope)
- Inference engine version
- Model family and quantization level
- Task type (structured output ≠ free-form generation ≠ classification)

Claims must be scoped: "for ARM Cortex-A76 class processors with LPDDR4X memory, in the context of ..."

### Writing template

```
External Validity. Our findings are based on [N cases / specific hardware / specific
context]. Generalizability is limited by [key bounding factors]. Practitioners in contexts
with [similar characteristics] may find the findings applicable, but replication is
recommended for [different SoC class / different task type / different scale].
```

---

## Conclusion Validity

**Threat**: Statistical conclusions are incorrect due to low power, violated assumptions, or inappropriate tests.

### Common threats and mitigations

| Threat | Example | Mitigation |
|---|---|---|
| Low statistical power | N=10 runs per condition; claims null result | Power analysis a priori; N≥30 as default |
| Violated test assumptions | t-test on non-normal data | Test for normality (Shapiro-Wilk); use Mann-Whitney for non-normal |
| Multiple comparisons | 20 tests at α=0.05 → 1 false positive expected | Bonferroni or Benjamini-Hochberg FDR correction |
| Reliability of measures | High variance in repeated measurements | Report SD/IQR; identify and address variance sources (thermal, cache) |
| Fishing / p-hacking | Testing many hypotheses and reporting only significant ones | Pre-register; report all tests; use correction |
| Random heterogeneity | Results vary across runs due to uncontrolled factors | Identify and control sources of variance; report variance |

### Writing template

```
Conclusion Validity. We used [test name] to test [hypothesis] because [normality
assumption check / data type]. With N=[N] runs per condition, the study has [power
analysis result or acknowledgment of limitation]. [If multiple comparisons]: We applied
[Bonferroni / BH-FDR] correction for [N] simultaneous tests, adjusting α to [value].
```

---

## Quick Reference: Threat Map by Study Type

| Study Type | Most Critical Threats |
|---|---|
| Practitioner case study (self) | Construct (operationalization), Internal (researcher bias), External (single case) |
| Benchmarking | Internal (thermal throttling, order effects), Conclusion (multiple comparisons, low N), External (platform specificity) |
| Controlled experiment | Internal (selection, maturation, instrumentation), Conclusion (power, test assumptions) |
| SLR | Construct (search completeness), Internal (selection bias), External (publication bias) |
| Survey | Construct (question ambiguity, response bias), External (non-response, sample representativeness) |

---

## Threats Section: Common Mistakes

1. **Generic boilerplate**: "We acknowledge that results may not generalize" - not acceptable. Be specific: what factors limit generalizability and why.
2. **Only listing threats without mitigations**: For each threat, state whether it was mitigated and how.
3. **Threats contradict contributions**: If a contribution claims broad applicability, external validity threats must be addressed.
4. **Missing conclusion validity entirely**: Often omitted in qualitative studies, but even qualitative work has conclusion validity concerns (coding reliability, saturation).
