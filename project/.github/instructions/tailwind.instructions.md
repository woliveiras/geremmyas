---
description: "Use when writing, reviewing, or refactoring Tailwind CSS v4 utility classes in components. Covers breaking changes from v3, renamed utilities, and patterns to avoid."
applyTo: "**/components/**/*.tsx, **/features/**/*.tsx, **/pages/**/*.tsx"
---

# Tailwind CSS v4 Conventions

- Use complete class strings. Tailwind scans source as text, so dynamic
  fragments like ``text-${color}-600`` will not be generated reliably.
- Use lookup maps for variants that depend on props or state.
- Prefer slash opacity syntax such as `bg-black/50` and `text-black/50`; do not
  use removed opacity utilities.
- Use modern flex utilities: `shrink-0` and `grow`.
- In v4, small shadow, blur, and radius utilities shifted down a step. Verify
  visual intent when migrating `shadow*`, `blur*`, and `rounded*` classes.
- Be explicit with borders because the default border color is `currentColor`.
- Be explicit with focus ring width when visual weight matters.
- Use CSS variable arbitrary values with v4 syntax like `bg-(--brand-color)`.
- Put `!important` at the end of the utility, for example `bg-red-500!`.
- Remember that `hover:` applies only on devices that support hover.
