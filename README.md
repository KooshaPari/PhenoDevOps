# PhenoDevOps

**Status:** maintenance

Unified CI/CD pipeline orchestration, infrastructure automation, and DevOps tooling for the Phenotype ecosystem. Streamlines the path from code commit to production with built-in quality gates, security scanning, and intelligent deployment strategies.

## Overview

**PhenoDevOps** is the backbone of the Phenotype platform's deployment infrastructure. It eliminates manual deployment steps and reduces risk through standardized, automated CI/CD pipelines, infrastructure-as-code (IaC) management, and operational tooling that enables fast, safe deployments at scale.

**Core Mission**: Provide a seamless, automated path from code to production with zero-manual-intervention deployment pipelines, infrastructure consistency checks, and comprehensive observability.

## Technology Stack

- **Pipeline Orchestration**: Custom Go-based engine with YAML pipeline definitions
- **Infrastructure as Code**: Terraform + Pulumi integration for cloud resource management
- **Container Orchestration**: Kubernetes manifest management and GitOps workflows
- **Deployment Strategies**: Rolling, blue-green, and canary deployment patterns with automatic rollback
- **Security**: SAST scanning, secrets detection, container image scanning, supply chain security
- **Observability**: Log aggregation, metrics dashboards, alert management, incident automation
- **Git Integration**: GitHub Actions compatible, Forgejo/Woodpecker support

## Quick Start

```bash
# Clone and explore
git clone <repo-url>
cd PhenoDevOps

# Review governance and architecture
cat CLAUDE.md          # Project governance
cat PRD.md             # Full product requirements
cat AGENTS.md          # Agent operating contract

# Explore pipeline examples
ls -la examples/pipelines/

# Run quality checks
go build ./cmd/phenodevops
go test ./...
```

## Project Structure

```
PhenoDevOps/
├── cmd/                   # CLI commands (pipeline, deploy, infra)
├── pkg/
│   ├── pipeline/          # Pipeline engine & YAML parsing
│   ├── deployer/          # Deployment orchestration (rolling, canary, blue-green)
│   ├── infrastructure/    # IaC integration (Terraform, Pulumi, K8s)
│   ├── observability/     # Logging, metrics, alerting
│   └── security/          # SAST, secret scanning, compliance
├── examples/              # Pipeline and deployment examples
├── docs/
│   ├── FUNCTIONAL_REQUIREMENTS.md  # Feature specifications
│   └── guides/            # Deployment guides, runbooks
└── PRD.md, AGENTS.md      # Product & governance docs
```

## Key Features

- **Pipeline Engine**: YAML-based pipeline definitions with parallel job execution, dependency management, artifact caching, and matrix builds
- **Stages**: Standard build, test (with parallelization), security scan, deploy (with approval), and smoke test stages
- **Deployment Strategies**: Rolling updates, blue-green deployments, canary releases with automatic rollback on failure
- **Infrastructure Management**: Declarative IaC with Terraform/Pulumi, drift detection, Kubernetes manifest handling
- **Day-2 Operations**: Log aggregation, metrics dashboards, alert management, runbook automation, incident response

## Related Phenotype Projects

- **[AgilePlus](../AgilePlus)** — Specification and work tracking for DevOps features
- **[Tracera](../Tracera)** — Distributed tracing backend for pipeline observability
- **[PhenoObservability](../PhenoObservability)** — Comprehensive observability platform (metrics, logs, APM)

## License

MIT — see [LICENSE](./LICENSE).
