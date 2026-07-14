---
name: paper-review
description: "Rewrite academic drafts to remove AI writing patterns while preserving claims and citations. Use when editing abstracts or paper sections. Do not use for blog posts or generating new content."
---

Rewrite the given text directly. No preamble, no explanation, no list of changes — deliver the rewritten text and nothing else.

## What to eliminate

### General AI tells (applies to all sections)

Remove or replace every instance of:

- Em dashes (—): use a comma, parentheses, or a new sentence instead
- "delve into", "dive deep", "unpack", "explore" (as a section opener), "tackle"
- "robust", "seamless", "powerful", "cutting-edge", "leverage" (verb), "harness" (verb), "utilize" (use "use")
- "crucial", "essential", "vital", "pivotal", "key" (when used as a filler adjective)
- "It's worth noting that", "It goes without saying", "Needless to say", "It is important to note that"
- "In the realm of", "In the landscape of", "When it comes to", "In today's X world"
- Forced tricolons: "accurate, efficient, and scalable" — keep if the three properties are independently measured, cut if decorative
- Hedging strings: "may potentially", "could possibly suggest", "might arguably be considered" → pick one hedge or eliminate

### Academic-specific AI tells

Fix every instance of:

- **"This paper aims to / presents / proposes / contributes"** as the sole opening of an abstract: the reader already knows it is a paper; open with the problem or the result instead
- **"novel", "innovative", "state-of-the-art"** without an empirical comparator or an explicit citation: if you cannot name what it outperforms and by how much, remove the adjective
- **"clearly", "obviously"**: if it were clear, you would not need to write the paper. Remove both.
- **"as can be seen from Figure X", "we can see that"**: replace with a direct factual claim ("Figure X shows a 12% reduction in latency")
- **Passive voice as a universal rule**: passive is appropriate in Methodology ("Participants were recruited…"); it is a tell in Introduction, Discussion, and Conclusion where active voice makes authorship and causality explicit
- **Abstract bloat**: background paragraphs longer than one sentence, motivation padding, context-setting that repeats the Introduction
- **"Further research is needed"** as a standalone conclusion: if you cannot specify what research, where, and why, cut the sentence
- **Contribution lists framed as intentions**: "We aim to show / We try to demonstrate / We seek to explore" → state what you did, not what you attempted
- **"In this paper, we"** repeated more than once in the abstract
- **Limitations section hedging**: "This may limit generalizability" without specifying to whom, under what conditions, and why

### Structural tells

Fix every instance of:

- **Abstract that does not follow the 5-beat structure** (see Rewrite Rules below)
- **Introduction that restates the abstract**: the Introduction adds context and motivation; if it copies abstract sentences, cut the copies
- **Discussion that merely narrates results**: Discussion interprets, not narrates — "X was higher" belongs in Results; "X was higher because Y, which suggests Z" belongs in Discussion
- **Conclusion that summarizes without implication**: end with what the findings mean for practice or future work, not a recap of your own paper

## Rewrite rules

### Shared rules (apply to all sections)

1. **Start with the point.** First sentence of any section: the finding, the problem, or the claim — not context about why the field matters.
2. **Cut anything that does not add a fact, result, or argument.** If a sentence could be deleted without losing scientific content, delete it.
3. **Replace double hedges with single or none.** "may potentially" → "may"; if the evidence supports a stronger claim, make it.
4. **Prefer specific.** Real numbers, real baselines, real effect sizes over qualitative claims.
5. **One idea per paragraph.** If a paragraph ends and it covered two arguments, split it.
6. **Preserve every technical fact and every citation.** Do not alter, omit, or strengthen claims beyond what the original supports.

### Abstract rewrite rule

Enforce this five-beat structure. Each beat is one sentence, two at most:

1. **Problem / Gap** — what is missing or broken in the field (one sentence, active voice preferred)
2. **What we did** — method or approach, stated as past tense action ("We analyzed…", "We introduce…")
3. **Key result** — the single most important finding, with numbers if available
4. **Secondary result or scope** — supporting finding or boundary conditions
5. **Implication** — what this means for practitioners, researchers, or the field

Do not add a sixth beat. Do not expand any beat beyond two sentences.

### Contribution framing rule

Contributions must be stated with result verbs, not intention verbs:

- Not: "We aim to show that X improves Y"
- Not: "We seek to explore the relationship between X and Y"
- Yes: "We show that X reduces Y by 18% across three datasets"
- Yes: "We identify three conditions under which X fails"
- Yes: "We provide a replication package with all scripts and data"

### Passive voice rule

| Section | Passive | Active |
|---|---|---|
| Abstract | Avoid | Prefer |
| Introduction | Avoid | Prefer |
| Methodology | OK for procedures | OK for decisions |
| Results | OK | OK |
| Discussion | Avoid | Prefer |
| Conclusion | Avoid | Prefer |
| Threats to Validity | OK | OK |

### Limitations / Threats rule

Each limitation must specify:
- **What** is limited (the construct, the sample, the method)
- **To whom** it does not generalize
- **Why** the limitation exists
- **What future work** could address it (one sentence, specific)

"This may limit generalizability" alone is not a valid limitation statement.

## Output format

Write the rewritten text directly back to the source file using file editing tools. Do not output the rewritten text in chat. Edit the file in place — overwrite the body content while preserving the frontmatter exactly as-is.

Same structure (section headings, figures references, citation keys) as the original. If a section is clean and needs no changes, leave it unchanged.
