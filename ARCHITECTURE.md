# Architecture

## Overview

DevOps tooling workspace combining Rust and Go for build, deployment, and infrastructure automation across the Phenotype ecosystem.

## Components

### Rust Crates (`crates/`)
DevOps-specific Rust crates for the workspace. Specific crate purposes determined by `Cargo.toml` workspace members.

### Go CLI/Utilities
Go binaries and utilities for DevOps operations. `go.mod` present at workspace root.

## Data Flow

Rust crates provide core DevOps primitives (build, config, artifact management). Go layer provides CLI and orchestration on top.

## Key Files

- `Cargo.toml` — Rust workspace manifest
- `go.mod` — Go module
- `crates/` — Rust DevOps crates
