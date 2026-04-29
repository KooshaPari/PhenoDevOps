# Known Issues

## Current

- The canonical checkout still contains broad untracked scaffold churn.
- Runtime Go checks are not meaningful in this isolated tracked tree because
  the runtime code is not present on `main`.

## Resolved

- Docs promotion is no longer blocked on the dirty canonical checkout.

