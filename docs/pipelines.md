# Pipelines

## Standard Stages

| Stage | Expected Evidence |
| --- | --- |
| Build | Reproducible package, binary, image, or static artifact |
| Test | Unit, integration, and smoke checks appropriate to the project |
| Security | Secret scanning, dependency checks, and container/image policy |
| Deploy | Target environment, strategy, approval gate, and rollback plan |
| Verify | Health checks, logs, metrics, and post-deploy smoke proof |

## Deployment Strategies

### Rolling

Use rolling updates for low-risk services with compatible versions and strong
health checks.

### Blue-Green

Use blue-green deployments when fast rollback matters and the platform can keep
two complete environments live.

### Canary

Use canaries when production traffic can be gradually shifted and measured with
clear error-budget guardrails.

## Traceability

Every deployable automation path should identify:

- source repository and revision
- target environment
- validation command or workflow
- rollback path
- owning project or team

