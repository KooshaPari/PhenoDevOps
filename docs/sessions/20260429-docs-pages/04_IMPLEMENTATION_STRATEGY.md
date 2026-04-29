# Implementation Strategy

## Approach

- Keep the docs shell dependency-free outside VitePress.
- Use static Markdown pages that can be extended as tracked runtime code lands.
- Ignore generated VitePress output and let CI recreate it.

