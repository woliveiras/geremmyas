---
name: scientific-paper
description: "End-to-end workflow for creating, reviewing, and critiquing scientific papers with rigorous methodology. Use when: writing papers, reviewing papers, literature reviews. Do not use: for blog posts, technical documentation."
---


# Scientific Paper Skill

Complete workflow for empirical SE research: paper creation, deep review, SLR, methodology critique, peer-review simulation, and LaTeX conventions.

## Modes of Operation

| Mode | Trigger phrases | What it does |
|---|---|---|
| `create` | "create paper", "new paper", "paper template", "criar artigo" | Scaffolds paper structure + LaTeX template |
| `review` | "review paper", "revisar paper", "find errors" | Full structured review: methodology, writing, validity |
| `critique` | "critique", "metodologia", "methodological errors", "erros metodológicos" | Deep methodology critique using error catalog |
| `peer-review` | "peer review", "referee report", "simulate reviewer", "parecerista" | Generates structured referee report (Summary / Strengths / Weaknesses / Questions / Recommendation) |
| `slr` | "systematic review", "SLR", "revisão sistemática", "literature review" | Guides SLR protocol from scratch |
| `methodology-check` | "check methodology", "rigor", "validity", "research design" | Runs methodology checklist for study type |
| `data-presentation` | "figures", "tables", "charts", "data viz", "plots", "benchmark results" | Reviews data presentation choices |
| `bib-validate` | "validate bibliography", "check citations", "missing citations", "bib" | Runs [validate-bib.py](./scripts/validate-bib.py) to catch orphaned \cite{} and missing .bib entries |

## Step-by-Step Procedures

### Mode: `review` - Full Paper Review

1. Ask for the paper file or content (paper.tex or text paste).
2. Load [error catalog](./references/error-catalog.md) to identify issues by category.
3. Load [methodology checklist](./references/methodology-checklist.md) and run the appropriate checklist for the paper's study type.
4. Load [threats-to-validity](./references/threats-to-validity.md) to audit the validity section.
5. Load [paper-structure](./references/paper-structure.md) to verify completeness of all sections.
6. Produce structured output:
   - **Critical errors** (must fix before submission)
   - **Major concerns** (likely desk-reject or major revision)
   - **Minor issues** (polish, clarity, consistency)
   - **Positive observations** (what is done well - important for balance)
7. For each issue: location (section/paragraph) → problem → recommended fix.

### Mode: `peer-review` - Referee Report Simulation

1. Ask for paper content + target venue (EMSE, ICSE, FSE, MSR, arXiv, etc.).
2. Load [venues.md](./references/venues.md) for venue-specific criteria and expectations.
3. Load [error-catalog.md](./references/error-catalog.md) and [threats-to-validity.md](./references/threats-to-validity.md).
4. Write a complete referee report in the style of the target venue:

```
SUMMARY
[2–3 sentences: what the paper does and its main claim]

STRENGTHS
1. ...
2. ...
3. ...

WEAKNESSES / MAJOR CONCERNS
1. ...
2. ...

MINOR COMMENTS
1. ...

SPECIFIC QUESTIONS FOR AUTHORS
1. ...

RECOMMENDATION
[ ] Accept  [ ] Minor Revision  [X] Major Revision  [ ] Reject

CONFIDENCE: [Reviewer expertise 1–5]
```

5. Provide a separate "author's perspective" note: what changes would move the paper from current recommendation to Accept.

### Mode: `critique` - Methodology Critique

1. Identify the study type from the paper (case study / benchmarking / experiment / SLR / survey).
2. If the paper is a case study, use the `scientific-case-study-research` skill for case definition, units of analysis, protocol, triangulation, chain of evidence, ethics, and threats to validity.
3. Load [methodology-checklist.md](./references/methodology-checklist.md) for that type.
4. Run each checklist item and flag gaps.
5. Load [error-catalog.md](./references/error-catalog.md) and map identified issues to error categories.
6. Prioritize by severity: conclusion validity > internal validity > construct validity > external validity.

### Mode: `slr` - Systematic Literature Review

1. Load [systematic-review-protocol.md](./references/systematic-review-protocol.md).
2. Walk through each phase: motivation → RQs → search string → databases → inclusion/exclusion → quality assessment → data extraction → synthesis.
3. Use [assets/slr-protocol-template.md](./assets/slr-protocol-template.md) to scaffold the protocol document.
4. Validate search string completeness (synonyms, acronyms, Boolean operators).
5. Check that each RQ maps to a data extraction field and a synthesis method.

### Mode: `create` - New Paper Scaffold

1. Determine study type and target venue.
2. Load [paper-structure.md](./references/paper-structure.md) for the appropriate template.
3. Use [assets/paper-template.tex](./assets/paper-template.tex) as the LaTeX starting point.
4. Load [latex-conventions.md](./references/latex-conventions.md) for repo-specific setup.
5. Load [venues.md](./references/venues.md) for venue formatting requirements.
6. Scaffold folder: `papers/YYYY-short-slug/` with `paper.tex`, `title-page.tex`, `README.md`.

### Mode: `data-presentation` - Figures and Tables Review

1. Load [data-presentation.md](./references/data-presentation.md).
2. Audit each figure/table: chart type appropriateness, axis labels, units, error bars, statistical significance markers, color accessibility, caption completeness.
3. Flag misleading visualizations (e.g., truncated y-axis, missing baseline, no variance shown).

### Mode: `bib-validate` - Bibliography Validation

1. Locate `paper.tex` and `*.bib` file in the workspace.
2. Run [validate-bib.py](./scripts/validate-bib.py):
   ```bash
   python scripts/validate-bib.py --tex paper.tex --bib references.bib
   ```
3. Report:
   - **Orphaned `\cite{}`**: keys used in .tex but missing from .bib
   - **Unused entries**: keys in .bib never cited in .tex
   - **Malformed entries**: missing required BibTeX fields (author, title, year)
4. Run after any round of work where new citations were added; skip if no new `\cite{}` keys were introduced.

## Key Resources

| Resource | Purpose |
|---|---|
| [methodology-checklist.md](./references/methodology-checklist.md) | Rigor checklists by study type |
| [error-catalog.md](./references/error-catalog.md) | Catalog of methodological and writing errors |
| [paper-structure.md](./references/paper-structure.md) | Section templates by paper type |
| [systematic-review-protocol.md](./references/systematic-review-protocol.md) | Full Kitchenham SLR protocol |
| [threats-to-validity.md](./references/threats-to-validity.md) | Wohlin validity framework for SE research |
| [latex-conventions.md](./references/latex-conventions.md) | Repo LaTeX conventions and anonymization |
| [venues.md](./references/venues.md) | EMSE / ICSE / FSE / MSR / arXiv submission checklists |
| [data-presentation.md](./references/data-presentation.md) | Figures, tables, and benchmark data visualization |
| [paper-template.tex](./assets/paper-template.tex) | Ready-to-use LaTeX template |
| [slr-protocol-template.md](./assets/slr-protocol-template.md) | Fillable SLR protocol document |
| [validate-bib.py](./scripts/validate-bib.py) | Bibliography validation script |

## Write like a human

Even in formal academic prose, keep out the AI writing tells reviewers now flag:

- No em dashes and no curly quotes; use straight quotes.
- Cut AI vocabulary: "delve", "leverage", "robust", "crucial", "pivotal", "testament", "underscore", "showcase", "intricate", "meticulous", "landscape" (as an abstract noun).
- Remove significance padding ("stands as a testament to", "plays a pivotal role"), trailing "-ing" analysis clauses ("...highlighting its significance"), and copula avoidance ("serves as" for "is").
- Do not use "novel" or "state-of-the-art" without a named comparator. Attribute claims to a citation, not to "experts" or "studies".
- Open sections with the finding or claim, not "In this paper" or "It is important to note". Do not close sections with "In summary"/"In conclusion" restatements.
