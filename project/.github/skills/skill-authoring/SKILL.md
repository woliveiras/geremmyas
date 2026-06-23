---
name: skill-authoring
description: "Create or revise Copilot skills that match this repository's structure and naming conventions. Use when: writing a new skill, renaming a skill, reviewing skill quality. Do not use: for general writing, non-skill documentation."
---


# Skill Authoring

Create skills that are small, discoverable, and specific to GitHub Copilot.

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
- [ ] The skill is adapted for GitHub Copilot, not copied from another agent.
