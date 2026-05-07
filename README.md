# PhenoDevOps — Phenotype Infrastructure Kit

> **Pinned references (Phenotype-org)**
> - MSRV: see rust-toolchain.toml
> - cargo-deny config: see deny.toml
> - cargo-audit: rustsec/audit-check@v2 weekly
> - Branch protection: 1 reviewer required, no force-push
> - Authority: phenotype-org-governance/SUPERSEDED.md

## What is the Phenotype Infrastructure Kit

The **Phenotype Infrastructure Kit** is a Rust workspace providing shared, reusable
infrastructure crates consumed across the Phenotype polyrepo. It centralizes
logging, error handling, caching, git operations, GitHub API clients, gRPC
infrastructure, NATS event bus adapters, and more to avoid duplication.

## Workspace Structure

```
crates/
  agileplus-*/        # AgilePlus-specific crates (API, dashboard, domain, etc.)
  phenotype-*/        # Shared Phenotype-org crates (async-traits, cache-adapter, etc.)
  pheno-*/            # General Phenotype crates (core, db, crypto, cli)
```

## Quick Start

```bash
# Build the workspace
cargo build --workspace

# Run tests
cargo test --workspace

# Lint
cargo clippy --workspace
```

## Key Crates

| Crate | Purpose |
|-------|---------|
| `phenotype-logging` | Structured logging via `tracing` |
| `phenotype-error-core` | Canonical error types and `thiserror` wrappers |
| `phenotype-cache-adapter` | Redis and in-memory caching |
| `phenotype-git-core` | Git operations (blobs, trees, commits, diffs) |
| `phenotype-github` | GitHub API client (REST + GraphQL) |
| `phenotype-nats` | NATS event bus adapter |
| `agileplus-api` | AgilePlus gRPC + REST API surface |
