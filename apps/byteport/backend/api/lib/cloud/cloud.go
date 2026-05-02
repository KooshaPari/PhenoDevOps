// Package cloud provides a unified abstraction layer for multi-cloud deployments
// supporting AWS, GCP, Azure, Vercel, Render, Supabase, Fly.io, Neon, and PlanetScale
package cloud

import (
	"context"
	"fmt"
	"sync"
)

// CloudProvider is the interface all cloud providers must implement.
type CloudProvider interface {
	// GetMetadata returns the provider's metadata.
	GetMetadata() ProviderMetadata

	// GetCapabilities returns the list of capabilities supported by this provider.
	GetCapabilities() []Capability

	// SupportsResource reports whether this provider can manage the given resource type.
	SupportsResource(resourceType ResourceType) bool

	// Initialize configures the provider with credentials and validates them.
	Initialize(ctx context.Context, credentials Credentials) error

	// ValidateCredentials verifies that the supplied credentials are valid.
	ValidateCredentials(ctx context.Context) error

	// CreateResource creates a new cloud resource.
	CreateResource(ctx context.Context, config ResourceConfig) (*Resource, error)

	// GetResource retrieves an existing resource by ID.
	GetResource(ctx context.Context, id string) (*Resource, error)

	// UpdateResource updates an existing resource's configuration.
	UpdateResource(ctx context.Context, id string, config ResourceConfig) (*Resource, error)

	// DeleteResource deletes a resource by ID.
	DeleteResource(ctx context.Context, id string) error

	// ListResources returns all resources matching the given filter.
	ListResources(ctx context.Context, filter ResourceFilter) ([]*Resource, error)

	// Deploy triggers a new deployment for an existing resource.
	Deploy(ctx context.Context, config DeploymentConfig) (*Deployment, error)

	// GetDeploymentStatus returns the current status of a deployment.
	GetDeploymentStatus(ctx context.Context, id string) (*DeploymentStatus, error)

	// RollbackDeployment rolls back a deployment to the previous version.
	RollbackDeployment(ctx context.Context, id string) error

	// GetLogs returns a log stream for a resource.
	GetLogs(ctx context.Context, resource *Resource, opts LogOptions) (LogStream, error)

	// GetMetrics returns metrics for a resource over the given time range.
	GetMetrics(ctx context.Context, resource *Resource, opts MetricOptions) ([]Metric, error)

	// EstimateCost returns a cost estimate for a resource configuration.
	EstimateCost(ctx context.Context, config ResourceConfig) (*CostEstimate, error)

	// GetActualCost returns the actual incurred cost for a resource.
	GetActualCost(ctx context.Context, resource *Resource, timeRange TimeRange) (*Cost, error)
}

// ProviderFactory creates a provider instance from credentials.
type ProviderFactory func(credentials Credentials) (CloudProvider, error)

// providerRegistry is the global registry of cloud providers.
var (
	registry     = make(map[string]ProviderFactory)
	registryMu   sync.RWMutex
)

// MustRegister registers a provider factory under the given name.
// It panics if a provider with that name is already registered.
func MustRegister(metadata ProviderMetadata, factory ProviderFactory) {
	registryMu.Lock()
	defer registryMu.Unlock()
	if _, exists := registry[metadata.Name]; exists {
		panic("cloud: provider " + metadata.Name + " already registered")
	}
	registry[metadata.Name] = factory
}

// GetProvider returns a provider by name, creating it with the supplied credentials.
func GetProvider(name string, credentials Credentials) (CloudProvider, error) {
	registryMu.RLock()
	factory, ok := registry[name]
	registryMu.RUnlock()
	if !ok {
		return nil, fmt.Errorf("cloud: no provider registered for %q", name)
	}
	return factory(credentials)
}

// ListProviders returns the names of all registered providers.
func ListProviders() []string {
	registryMu.RLock()
	defer registryMu.RUnlock()
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	return names
}

// notSupportedError is returned when a provider does not support a requested operation.
type notSupportedError struct {
	provider  string
	operation string
}

func (e *notSupportedError) Error() string {
	return fmt.Sprintf("cloud: provider %q does not support %s", e.provider, e.operation)
}

// NewNotSupportedError returns an error indicating the provider does not support an operation.
func NewNotSupportedError(provider, operation string) error {
	return &notSupportedError{provider: provider, operation: operation}
}

// IsNotSupported reports whether err is a not-supported error.
func IsNotSupported(err error) bool {
	_, ok := err.(*notSupportedError)
	return ok
}

