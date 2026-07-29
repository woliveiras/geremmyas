# Style Pack Contract

Use this reference when adding an artistic style to the skill.

Create a sibling directory under `references/`:

```text
references/
  <style-slug>/
    style-guide.md
```

Use a lowercase kebab-case slug. Link the new guide directly from `SKILL.md`.
Keep all essential guidance in `style-guide.md`; add supporting files only when
they provide real reusable value.

## Required sections

Define:

1. **Visual signature**: observable traits, not mood-only labels.
2. **Resolution and scale**: logical dimensions, density, and display rules.
3. **Shape and line**: silhouette, contour, edge, and internal detail rules.
4. **Palette and value**: color limits, ramps, contrast, and semantic roles.
5. **Light and material**: shading, highlights, texture, and surface language.
6. **Characters and objects**: proportion and readability rules.
7. **Environments**: projection, depth, tiling, and detail hierarchy.
8. **Animation**: cadence, deformation, interpolation, and effects.
9. **Prompt block**: concise reusable prompt language.
10. **Quality checklist**: observable pass/fail signals.

## Rules

- Describe a repeatable system rather than a list of style adjectives.
- Keep runtime readability ahead of decorative fidelity.
- Use original art-language descriptions. Do not instruct imitation of a living
  artist or protected franchise.
- State which rules are strict and which can vary by project.
- Preserve the project's established style when it conflicts with a bundled
  pack.
- Include negative constraints that target common generation drift.
- Test the style on at least one character, object, environment slice, and
  effect before treating it as a complete project direction.
