// Package cloud — Local Docker/Podman provider implementation.
// NOTE: This provider requires the docker/docker Go module which has a known
// incompatibility with Go 1.24's replace directive restrictions.
// Until resolved, this provider is gated behind the "docker_provider" build tag.
// Build with: go build -tags=docker_provider
package cloud

import (
	"context"
	"fmt"
)

// LocalProvider implements CloudProvider for local Docker/Podman deployments.
// NOTE: This is a stub. Full implementation requires the docker/docker module
// which is currently excluded due to Go 1.24 replace directive incompatibilities.
// To enable: go build -tags=docker_provider
type LocalProvider struct {
	credentials Credentials
	socketPath string
	host       string
	metadata   ProviderMetadata
}

// compile-time interface assertion
var _ CloudProvider = (*LocalProvider)(nil)

// NewLocalProvider creates a new LocalProvider (stub).
// Full implementation requires docker/docker module — build with -tags=docker_provider.
func NewLocalProvider(credentials Credentials) (*LocalProvider, error) {
	return &LocalProvider{
		credentials: credentials,
		socketPath: credentials.Data["socket_path"],
		host:       credentials.Data["host"],
		metadata: ProviderMetadata{
			Name:    "local",
			Version: "0.1.0",
			Description: "Local Docker/Podman provider (stub — requires -tags=docker_provider for full impl)",
		},
	}, nil
}

func (p *LocalProvider) GetMetadata() ProviderMetadata { return p.metadata }
func (p *LocalProvider) GetCapabilities() []Capability { return nil }

func (p *LocalProvider) SupportsResource(rt ResourceType) bool {
	return rt == ResourceTypeComputeContainer
}

func (p *LocalProvider) Initialize(ctx context.Context, creds Credentials) error {
	p.credentials = creds
	return nil
}

func (p *LocalProvider) ValidateCredentials(ctx context.Context) error {
	return NewNotSupportedError("local", "ValidateCredentials — requires docker_provider build tag")
}

func (p *LocalProvider) CreateResource(ctx context.Context, config ResourceConfig) (*Resource, error) {
	return nil, NewNotSupportedError("local", "CreateResource — requires docker_provider build tag")
}

func (p *LocalProvider) GetResource(ctx context.Context, id string) (*Resource, error) {
	return nil, NewNotSupportedError("local", "GetResource — requires docker_provider build tag")
}

func (p *LocalProvider) UpdateResource(ctx context.Context, id string, config ResourceConfig) (*Resource, error) {
	return nil, NewNotSupportedError("local", "UpdateResource — requires docker_provider build tag")
}

func (p *LocalProvider) DeleteResource(ctx context.Context, id string) error {
	return NewNotSupportedError("local", "DeleteResource — requires docker_provider build tag")
}

func (p *LocalProvider) ListResources(ctx context.Context, filter ResourceFilter) ([]*Resource, error) {
	return nil, NewNotSupportedError("local", "ListResources — requires docker_provider build tag")
}

func (p *LocalProvider) Deploy(ctx context.Context, config DeploymentConfig) (*Deployment, error) {
	return nil, NewNotSupportedError("local", "Deploy — requires docker_provider build tag")
}

func (p *LocalProvider) GetDeploymentStatus(ctx context.Context, id string) (*DeploymentStatus, error) {
	return nil, NewNotSupportedError("local", "GetDeploymentStatus — requires docker_provider build tag")
}

func (p *LocalProvider) RollbackDeployment(ctx context.Context, id string) error {
	return NewNotSupportedError("local", "RollbackDeployment — requires docker_provider build tag")
}

func (p *LocalProvider) GetLogs(ctx context.Context, resource *Resource, opts LogOptions) (LogStream, error) {
	return nil, NewNotSupportedError("local", "GetLogs — requires docker_provider build tag")
}

func (p *LocalProvider) GetMetrics(ctx context.Context, resource *Resource, opts MetricOptions) ([]Metric, error) {
	return nil, NewNotSupportedError("local", "GetMetrics — requires docker_provider build tag")
}

func (p *LocalProvider) EstimateCost(ctx context.Context, config ResourceConfig) (*CostEstimate, error) {
	return &CostEstimate{HourlyUSD: 0, DailyUSD: 0, MonthlyUSD: 0, Confidence: "high", Currency: "USD"}, nil
}

func (p *LocalProvider) GetActualCost(ctx context.Context, resource *Resource, timeRange TimeRange) (*Cost, error) {
	return nil, NewNotSupportedError("local", "GetActualCost — requires docker_provider build tag")
}

func autoDetectSocket() (host string, socketPath string, err error) {
	return "", "", fmt.Errorf("local: auto-detect not available — set socket_path or host explicitly")
}
