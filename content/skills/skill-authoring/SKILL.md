---
name: skill-authoring
description: "Create or revise coding-assistant skills that match this repository's structure and conventions. Use when: writing, renaming, or reviewing a skill. Do not use: for general writing or non-skill documentation."
---


# Skill Authoring

Create portable skills that are small, discoverable, and usable by the
assistants Geremmyas supports.

## Process

1. Define the skill's job and trigger phrases.
2. Choose an action-oriented name in clear English.
3. Write `SKILL.md` with:
   - frontmatter `name`
   - frontmatter `description`
   - short process
   - rules
   - output expectations when useful
4. Move long examples, templates, tables, or rarely needed details to `assets/`
   or `references/`.
5. Update README, prompts, agents, instructions, and installer references when
   a skill is added or renamed.

## Checklist

- [ ] The directory name and `name:` match.
- [ ] The description says what the skill does and when to use it.
- [ ] The skill is a workflow or capability, not a passive document title.
- [ ] `SKILL.md` is short enough to load into context comfortably.
- [ ] Examples are concrete and reusable.
- [ ] The skill is assistant-neutral unless a target-specific constraint is explicit.
