# Guide

## Workspace Model

PhenoDevOps owns the operational language around how Phenotype projects move
from source changes to validated deploys. Treat it as the docs-first home for
automation patterns before those patterns become shared tools.

## Core Domains

- **Pipeline orchestration:** build, test, scan, package, deploy, and verify
  stages with explicit dependencies.
- **Deployment strategies:** rolling, blue-green, and canary rollouts with
  rollback criteria.
- **Infrastructure coordination:** Terraform, Pulumi, Kubernetes manifests, and
  GitOps workflows.
- **Day-2 operations:** observability, alert response, runbooks, and incident
  automation.

## Change Policy

- Keep examples generic and secret-free.
- Prefer documented contracts over one-off scripts.
- Connect new project-wide policy back to AgilePlus and the Phenotype tooling
  repos when it becomes reusable.

## Local Requirements

- Bun for docs site validation.
- Go 1.21+ for future runtime work.
- Access to deployment providers only when testing provider-specific automation.

