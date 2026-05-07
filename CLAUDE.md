# PhenoDevOps — CLAUDE.md

Extends thegent governance base. See `platforms/thegent/dotfiles/governance/CLAUDE.base.md` for canonical definitions.

## Project Overview

- **Name**: PhenoDevOps
- **Description**: Phenotype DevOps infrastructure: CI/CD workflows, deployment tooling, Docker/Rust build chains, and developer experience automation
- **Location**: `/Users/kooshapari/CodeProjects/Phenotype/repos/PhenoDevOps/`
- **Language Stack**: Rust (Cargo workspace), TypeScript/Node, Shell, Docker
- **Published**: Internal (Phenotype org)

## AgilePlus Mandate

All work MUST be tracked in AgilePlus:
- Reference: `/Users/kooshapari/CodeProjects/Phenotype/repos/AgilePlus`
- CLI: `cd /Users/kooshapari/CodeProjects/Phenotype/repos/AgilePlus && agileplus <command>`

## Work Requirements

1. **Check for AgilePlus spec before implementing**
2. **Create spec for new work**: `agileplus specify --title "<feature>" --description "<desc>"`
3. **Update work package status**: `agileplus status <feature-id> --wp <wp-id> --state <state>`
4. **No code without corresponding AgilePlus spec**

## Branch Discipline

- Feature branches in `.worktrees/<topic>/`
- Legacy `PROJECT-wtrees/` and `repo-wtrees/` roots are for migration only and must not receive new work
- Canonical repository remains on `main` for final integration and verification
- Return to `main` for merge/integration checkpoints

## UTF-8 Encoding

All files must use UTF-8. Validate with:
```bash
cd /Users/kooshapari/CodeProjects/Phenotype/repos/PhenoDevOps
agileplus validate-encoding --all --fix
```

## Local Quality Checks

From this repository root:

```bash
cargo test --workspace
cargo clippy --workspace -- -D warnings
cargo fmt --check
```

## Testing & Specification Traceability

All tests MUST reference a Functional Requirement (FR):

```rust
// Traces to: FR-XXX-NNN
#[test]
fn test_feature_name() {
    // Test body
}
```

**Verification**:
- Every FR in FUNCTIONAL_REQUIREMENTS.md MUST have >=1 test
- Every test MUST reference >=1 FR
- Run: `cargo test` to verify

---

## Project Structure

```
PhenoDevOps/
├── Cargo.toml              # Rust workspace root
├── crates/                 # Rust workspace members
│   ├── phenotype-cache-adapter/
│   ├── phenotype-event-sourcing/
│   ├── phenotype-health/
│   └── ...
├── apps/                   # Application projects
├── agent-devops-setups/    # DevOps agent configurations
├── deploy/                 # Deployment scripts and configs
├── docker/                 # Dockerfiles and compose files
├── ci/                     # CI/CD workflow configs
├── docs/                   # Documentation
└── dotfiles/              # Developer dotfiles
```

## Architecture

### Rust Workspace

- Workspace members defined in root `Cargo.toml`
- Each crate is independently consumable
- Shared dependencies managed at workspace level
- Error types use `thiserror` with `#[from]` for conversions
- Serialization via `serde` with `Serialize`/`Deserialize` derives

### CI/CD

- GitHub Actions workflows in `.github/workflows/`
- Secrets scanning: trufflehog + gitleaks
- SBOM generation: cargo-bom
- License compliance: cargo-deny
- Security scanning: CodeQL + Semgrep

## Governance Reference

See thegent governance base for:
- Complete CI completeness policy
- Phenotype Git and Delivery Workflow Protocol
- Phenotype Org Cross-Project Reuse Protocol
- Phenotype Long-Term Stability and Non-Destructive Change Protocol
- Worktree Discipline guidelines

Location: `platforms/thegent/dotfiles/governance/CLAUDE.base.md`
