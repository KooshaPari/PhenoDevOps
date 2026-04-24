# AGENTS.md — PhenoDevOps

## Project Overview

- **Name**: PhenoDevOps
- **Description**: DevOps automation, CI/CD, deployments, and infrastructure as code for Phenotype
- **Location**: `/Users/kooshapari/CodeProjects/Phenotype/repos/PhenoDevOps`
- **Language Stack**: Go (primary), Rust, TypeScript
- **Published**: Internal (Phenotype ecosystem)

## Quick Start Commands

```bash
# Navigate to PhenoDevOps
cd /Users/kooshapari/CodeProjects/Phenotype/repos/PhenoDevOps

# Go components
go build ./...
go test ./...

# Rust components
cd rust && cargo build

# TypeScript components
cd ts && npm install && npm run build
```

## Architecture

```
PhenoDevOps/
├── agent-devops-setups/      # Agent-specific DevOps configs
├── chaos/                    # Chaos engineering
├── ci/                       # CI configurations
├── deploy/                   # Deployment scripts
├── docker/                   # Docker configurations
├── examples/                 # Example setups
├── go.mod                    # Go module
├── jobs/                     # Job definitions
├── migrations/               # Database migrations
├── README.md                 # Project overview
├── rust/                     # Rust components
└── ts/                       # TypeScript components
```

## Quality Standards

### Go Standards
- **Line length**: 100 characters
- **Formatter**: `gofmt`, `goimports`
- **Linter**: `golangci-lint`
- **Tests**: `go test ./...`

### IaC Standards
- Docker files must pass `docker build --no-cache`
- Kubernetes manifests must pass `kubectl apply --dry-run=client`

## Git Workflow

### Branch Naming
Format: `phenodevops/<type>/<description>`

Examples:
- `phenodevops/feat/argo-cd-integration`
- `phenodevops/fix/docker-build`

### Commit Format
```
<type>(<area>): <description>

Examples:
- feat(ci): add GitHub Actions workflows
- fix(deploy): resolve k8s namespace issue
```

## File Structure

```
PhenoDevOps/
├── ci/
│   └── github-actions/       # GitHub workflows
├── deploy/
│   ├── kubernetes/           # K8s manifests
│   └── terraform/            # TF modules
├── docker/
│   └── Dockerfile.*          # Container definitions
└── [language-specific]
```

## CLI Commands

```bash
# Go
make build
make test

# Docker
docker build -f docker/Dockerfile.<name> .

# K8s (dry run)
kubectl apply -f deploy/k8s/ --dry-run=client
```

## Troubleshooting

| Issue | Solution |
|-------|----------|
| Go module errors | `go mod tidy` |
| Docker build fails | Check base image tags |
| K8s validation fails | Use `kubectl validate` |

## Dependencies

- **pheno/**: Core configurations
- **phenodocs/**: Documentation deployment
- **AgilePlus**: CI integration

## Agent Notes

When working in PhenoDevOps:
1. Changes affect production infrastructure
2. Test IaC changes with dry-run first
3. Coordinate with cloud/ and nanovms/ for deployment targets
