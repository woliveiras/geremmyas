---
name: text-review
description: >
  Rewrites blog post drafts to eliminate AI-generated writing patterns while
  preserving the author's voice, technical accuracy, and all factual content.
  Use when a post draft feels generic, padded, or clearly machine-generated.
  Trigger: /text-review, "rewrite this post", "remove AI patterns", "humanize this text".
---

Rewrite blog drafts so they sound like a real technical author with lived context, not like a polished generic assistant.

Prioritize three checks:

1. **Voice**: does this sound like the author has a point of view?
2. **Evidence**: does the text contain concrete context, constraints, failures, and decisions?
3. **Shape**: does the post use the right structure for what it is trying to do?

## What to eliminate

### Lexical tells

Remove or replace every instance of:

- The em dash character: use a comma, parentheses, colon, semicolon, or a new sentence instead
- "delve into", "dive deep", "dive into", "unpack", "tackle", "explore" (as a section opener), "demystify"
- "robust", "seamless", "powerful", "cutting-edge", "game-changing", "leverage" (verb), "harness" (verb), "utilize" (use "use")
- "crucial", "essential", "vital", "pivotal", "key" (when used as a filler adjective)
- "not just X, but Y" constructions: flatten into a direct claim
- "It's worth noting that", "It's important to remember that", "It goes without saying that", "Needless to say"
- "In the realm of", "In the world of", "In the landscape of", "When it comes to", "In today's X world", "In an era of"
- "Let's", "we'll" used to create false intimacy when the author is writing alone
- Forced tricolons: "fast, scalable, and reliable" - keep if genuinely meaningful, cut if decorative
- "Happy coding!", "Let's get started!", "Stay tuned!" - remove closing platitudes

### Structural tells

Fix every instance of:

- **Generic opening paragraph**: the first paragraph must contain a thesis, a case, or a concrete fact - not context-setting prose about "why X matters nowadays"
- **Recapping conclusion**: if the conclusion only restates what was already said, cut the restatement and replace with an implication, a decision, or a next step the reader can take
- **Mechanical transitions**: "Now that we've covered X, let's move on to Y" → cut entirely or merge into the next paragraph's opening sentence
- **Bullets where prose works**: if a list has fewer than 4 items and each item is a sentence-length thought, rewrite as prose
- **Uniform paragraph length**: vary sentence and paragraph length. A short punch sentence after a dense technical block is a deliberate choice - use it.
- **Hedging that avoids commitment**: "may", "might", "could potentially", "it depends on many factors" where a concrete recommendation is possible → make the call

### Tonal tells

Fix every instance of:

- **No opinion**: if the text presents trade-offs without recommending, add the recommendation ("I use X here because Y")
- **Generic examples**: replace foo/bar/MyClass/arbitrary round numbers with realistic, specific examples drawn from the post's context
- **Missing lived experience**: if the author is describing a problem or solution, add the cost, the failure mode, or the specific moment where this mattered - even one sentence is enough
- **False modesty**: "This is just my experience", "Your mileage may vary" as boilerplate disclaimers - cut unless genuinely necessary

### Functional tells

Remove sentences whose only job is:

- Announcing the topic instead of saying something about it
- Praising a technology before explaining its behavior
- Creating suspense without a concrete catch
- Summarizing the previous paragraph
- Explaining that the next section will explain something
- Saying a topic is complex, important, evolving, or challenging without evidence
- Framing obvious advice as a discovery

## Authorial fingerprints

Preserve or add signals that a real person wrote the post:

- Specific stakes: what broke, what was annoying, what cost time, what changed after the decision
- Decision scars: the failed option, trade-off accepted, misleading assumption, or thing the author would not repeat
- Temporal anchors: when this happened, what changed after a release, migration, deploy, review, or debugging session
- Concrete constraints: team size, repo shape, dependency version, CI time, memory limit, deployment target, editor, OS, or exact error message
- Opinion with reason: "I use X here because Y" or "I would not use X in this case because Y"
- One uncomfortable detail when useful: the workaround, wrong assumption, misleading error, restart, flaky check, or manual step

Do not invent facts. If the draft lacks concrete context, either preserve the gap or add a concise placeholder such as `[add exact error message here]` when editing a draft.

## Anti-symmetry rules

Avoid artificial balance.

- Do not give equal weight to options when the author clearly prefers one.
- Replace "it depends" with the actual condition that changes the decision.
- Cut generic trade-off paragraphs unless they end with a recommendation.
- Avoid "pros and cons" framing for posts that are really about a hard-won lesson.
- Prefer "I choose X when Y is true" over neutral comparison.
- If one option is mostly bad, say so and explain the failure mode.

## Structural diagnosis

Before rewriting, identify the post's real shape and make the structure match it:

- Incident: something broke, here is the fix.
- Decision: I chose X over Y, here is why.
- Tutorial: do these steps, avoid these traps.
- Opinion: I believe X because of Y.
- Postmortem: this failed, here is what changed.
- Research note: here is what I observed and what it may mean.
- Release note: this changed, here is who should care.

Avoid defaulting to: introduction, what is X, why it matters, conclusion. Use that shape only when it genuinely serves the post.

## Technical voice

For technical posts:

- Prefer exact command outputs, config snippets, file paths, package names, versions, and failure messages.
- Keep technical nouns precise. Do not replace exact terms with softer synonyms.
- Do not over-explain beginner concepts unless the post is explicitly introductory.
- Do not soften criticism of tools when the behavior is objectively bad.
- When a claim is based on personal experience, state the boundary: project size, stack, environment, workflow, or time period.
- Preserve code blocks, command snippets, links, tables, and frontmatter exactly unless the user asks to change them.
- Never change a technical claim to make prose smoother.

## Rewrite rules

1. **Start with the point.** First sentence of any section: thesis or concrete case, not background.
2. **Cut anything that does not add a fact, opinion, or step.** If a sentence could be deleted without losing meaning, delete it.
3. **Replace hedge with choice.** "It's often better to use X" → "Use X. It avoids Y."
4. **Vary rhythm.** After three long sentences, write one short one. Intentional fragments are acceptable.
5. **Prefer specific.** Real dates, real numbers, real names, real error messages over placeholders.
6. **One idea per paragraph.** If a paragraph ends and it covered two ideas, split it.
7. **Preserve every technical fact.** Do not simplify, omit, or soften technical claims. Accuracy is non-negotiable.
8. **Preserve the author's voice.** The goal is to make the text sound like a senior engineer who has opinions, not like a style guide.
9. **Keep useful roughness.** Do not polish away all personality, irritation, uncertainty, or specificity.
10. **Prefer earned confidence.** Strong claims need concrete support; weak evidence should produce scoped claims.

## Output format

If the user provides a file path or asks to rewrite in place, edit the file directly with file editing tools. Overwrite the body content while preserving frontmatter exactly as-is.

If the user asks for feedback, return:

1. LLM-like patterns found
2. Structural issues
3. Suggested rewrite strategy
4. Optional rewritten excerpt

If the user pastes text in chat and asks for a rewrite, return the rewritten text directly unless they ask for diagnostics.

Same structure (headings, code blocks, links) as the original. If a section is clean and needs no changes, leave it unchanged.
