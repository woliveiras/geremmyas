# Venue Submission Checklists

Formatting requirements, expectations, and submission checklists for target venues.

---

## Quick Reference Table

| Venue | Type | Review | Page limit | Format | Artifact policy |
|---|---|---|---|---|---|
| arXiv | Preprint | None | Unlimited | Any | N/A |
| EMSE | Journal | Single-blind (revision) | Unlimited (guideline: 40pp) | Springer LNCS | Encouraged |
| IEEE TSE | Journal | Double-blind | ~30 pp | IEEE two-column | Encouraged |
| ICSE | Conference | Double-blind | 10pp (+ 2 ref) | IEEE two-column | Required (AE) |
| FSE/ESEC | Conference | Double-blind | 10pp (+ 2 ref) | ACM two-column | Required (AE) |
| MSR | Conference | Double-blind | 10pp (mining) / 4pp (short) | IEEE two-column | Required (AE) |
| EASE | Conference | Double-blind | 10pp | ACM | Encouraged |
| IST | Journal | Single-blind | Unlimited (~35pp) | Elsevier | Encouraged |
| JSS | Journal | Single-blind | ~25pp | Elsevier | Encouraged |

AE = Artifact Evaluation track.

---

## arXiv - Preprint

**Purpose**: Establish priority, share openly, get feedback before formal submission.

**Checklist:**
- [ ] `\anonymizedfalse` - arXiv version is deanonymized
- [ ] Author list complete with ORCID
- [ ] Data availability statement included
- [ ] Repository URL(s) included
- [ ] License stated (CC BY 4.0 for paper)
- [ ] Subject classification: cs.SE and any domain-specific (cs.AI, cs.AR, etc.)
- [ ] Abstract: clear enough for readers outside your sub-field
- [ ] Build: `pdflatex` twice; no `??` references in compiled PDF
- [ ] PDF metadata: author name and title present (NOT suppressed)

**Submission URL**: https://arxiv.org/submit

---

## EMSE - Empirical Software Engineering (Springer)

**Scope**: Empirical studies of SE processes, products, people. Strong preference for rigorous methodology.

**Checklist:**
- [ ] Study type clearly labeled in abstract/title (case study, experiment, SLR, survey)
- [ ] Methodology section follows established guidelines (cite Runeson & Höst, Wohlin, or Kitchenham as appropriate)
- [ ] Threats to Validity section: all four categories (Wohlin framework)
- [ ] Replication package: link to data + scripts (or data availability statement)
- [ ] Inter-rater reliability reported for qualitative analysis
- [ ] Related work positions paper explicitly (gap statement)
- [ ] Abstract: structured (Background / Objective / Method / Results / Conclusion)
- [ ] Springer LNCS format not required for submission (but check current author instructions)
- [ ] Cover letter: explain fit to EMSE scope, novelty claim, target editor if known

**Key EMSE expectations:**
- Empirical rigor over novelty of technique
- Transparency about methodology limitations
- Replication is valued (but new contribution expected)

---

## ICSE - International Conference on Software Engineering

**Scope**: Flagship SE conference. Highly competitive (~20–25% acceptance rate for research track).

**Checklist (Research Track):**
- [ ] `\anonymizedtrue` - double-blind; NO author names, affiliations, or identifying URLs
- [ ] Self-citations: anonymized as "[Anonymous, 2024]" or "[Removed for review]"
- [ ] PDF metadata: author field empty (`pdfauthor={}`)
- [ ] Page limit: 10 pages + 2 pages for references only
- [ ] IEEE format: two-column, 10pt
- [ ] Figures: readable at full-page print scale (fonts ≥ 8pt in figures)
- [ ] Artifact evaluation: prepare replication package; register for AE track separately
- [ ] Line numbers: required for review submissions (check call for papers)
- [ ] CCS concepts and keywords: required
- [ ] Data availability statement: required

**Key ICSE expectations:**
- Strong motivation: why does this problem matter to SE?
- Solid evaluation: clear metrics, baselines, statistical analysis
- Related work: must engage with top-cited SE papers on the topic

---

## FSE/ESEC - ACM/SIGSOFT Symposium on Foundations of Software Engineering

**Scope**: SE methods, tools, empirical studies. Co-located with ESEC (European SE Conference).

**Checklist:**
- [ ] Double-blind: same rules as ICSE
- [ ] ACM format: two-column (`acmart` class, `sigconf` option)
- [ ] Page limit: 10pp + references (no page limit on references)
- [ ] ACM CCS concepts required
- [ ] ACM rights management forms required (camera-ready)
- [ ] Artifact evaluation track: submit separately if applicable
- [ ] Review: typically 3 reviewers + meta-review; rebuttal phase exists

---

## MSR - Mining Software Repositories

**Scope**: Mining Git repos, issue trackers, code review, Stack Overflow, etc. Data-driven SE.

**Checklist (Technical Papers):**
- [ ] Double-blind
- [ ] IEEE format: two-column
- [ ] 10 pages + 2 reference pages (technical papers)
- [ ] Data availability: MSR strongly expects public dataset or replication package
- [ ] Mining methodology described: what data, how collected, what filters applied
- [ ] Ethical considerations: if human subjects data (GitHub users, SO posts), discuss ethics
- [ ] Short papers: 4 pages (emerging results, tools)
- [ ] Mining challenge: separate track for challenge tasks

**Key MSR expectations:**
- Dataset is first-class contribution
- Replication and transparency highly valued
- Negative results and replications explicitly welcome

---

## IST - Information and Software Technology (Elsevier)

**Scope**: Broad SE; accepts empirical studies, reviews, technical papers.

**Checklist:**
- [ ] Single-blind (author names visible to reviewers; reviewers anonymous)
- [ ] Elsevier template (elsarticle)
- [ ] Highlights: 3–5 bullet points required by Elsevier (what is new, what was found)
- [ ] Graphical abstract: optional but recommended
- [ ] Structured abstract
- [ ] Data in Brief: optionally submit supplementary data as a separate Data in Brief article

---

## Pre-Submission Final Checklist (All Venues)

### Content
- [ ] All RQs answered in Results and Conclusion
- [ ] Contributions listed in Introduction match what is delivered
- [ ] Abstract claims have evidence in the paper body
- [ ] Threats to Validity section present and complete

### Formatting
- [ ] Build produces PDF with no `??` for undefined references
- [ ] All figures have captions and are referenced in text
- [ ] All tables have captions above the table
- [ ] Listings have captions/labels
- [ ] No overfull `\hbox` warnings in critical sections (check build log)

### Double-blind (if applicable)
- [ ] `\anonymizedtrue` set
- [ ] `pdfinfo paper.pdf | grep Author` returns empty
- [ ] No first-person possessives near citations ("our [X]", "we developed [Y]~\cite{}")
- [ ] No identifying URLs (personal website, specific GitHub username in repo URL)
- [ ] Acknowledgments section hidden

### Data and Ethics
- [ ] Data availability statement present
- [ ] Replication package URL or DOI provided (or explained why not available)
- [ ] Ethical approval mentioned if study involves human subjects
- [ ] Licenses for datasets and code stated

### Submission
- [ ] Cover letter addresses: scope fit, novelty, and (if journal) conflicts of interest
- [ ] Suggested reviewers listed where requested (choose outside your institution)
- [ ] Check submission deadline + time zone (AoE = Anywhere on Earth)
