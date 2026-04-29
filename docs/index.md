# PhenoDevOps

> DevOps automation, deployment orchestration, and operational tooling for the Phenotype ecosystem.

PhenoDevOps is the Phenotype DevOps workspace. It documents the shared CI/CD,
deployment, infrastructure, and day-2 operations patterns that keep project
delivery reproducible across the shelf.

## Surfaces

| Surface | Purpose |
| --- | --- |
| Pipeline orchestration | Standard build, test, security, deploy, and smoke-test stages |
| Deployment automation | Rolling, blue-green, and canary rollout patterns |
| Infrastructure workflows | Terraform, Pulumi, Kubernetes, and GitOps coordination |
| Operations docs | Validation, release, and maintenance procedures |

## Maintenance Posture

This repo is in maintenance mode. Favor documentation clarity, runnable
examples, and governance alignment over expanding the runtime surface.

## Quick Start

```bash
bun install
bun run docs:build
```

The docs site is static and can be built locally without deployment secrets.

## Public Routes

- [Guide](/guide) for the workspace model and repo contract
- [Pipelines](/pipelines) for the delivery-stage taxonomy
- [Operations](/operations) for local checks and Pages publishing

## Public Contract

- Keep deployment behavior documented before automation is added.
- Keep examples secret-free and safe to copy.
- Route reusable DevOps patterns back to Phenotype shared tooling when they
  become cross-project policy.

