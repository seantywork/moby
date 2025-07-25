package system

import (
	"context"
	"time"

	"github.com/moby/moby/api/types"
	"github.com/moby/moby/api/types/build"
	"github.com/moby/moby/api/types/events"
	"github.com/moby/moby/api/types/filters"
	"github.com/moby/moby/api/types/registry"
	"github.com/moby/moby/api/types/swarm"
	"github.com/moby/moby/api/types/system"
)

// DiskUsageOptions holds parameters for system disk usage query.
type DiskUsageOptions struct {
	// Containers controls whether container disk usage should be computed.
	Containers bool

	// Images controls whether image disk usage should be computed.
	Images bool

	// Volumes controls whether volume disk usage should be computed.
	Volumes bool
}

// Backend is the methods that need to be implemented to provide
// system specific functionality.
type Backend interface {
	SystemInfo(context.Context) (*system.Info, error)
	SystemVersion(context.Context) (types.Version, error)
	SystemDiskUsage(ctx context.Context, opts DiskUsageOptions) (*system.DiskUsage, error)
	SubscribeToEvents(since, until time.Time, ef filters.Args) ([]events.Message, chan interface{})
	UnsubscribeFromEvents(chan interface{})
	AuthenticateToRegistry(ctx context.Context, authConfig *registry.AuthConfig) (string, error)
}

// ClusterBackend is all the methods that need to be implemented
// to provide cluster system specific functionality.
type ClusterBackend interface {
	Info(context.Context) swarm.Info
}

// BuildBackend provides build specific system information.
type BuildBackend interface {
	DiskUsage(context.Context) ([]*build.CacheRecord, error)
}

// StatusProvider provides methods to get the swarm status of the current node.
type StatusProvider interface {
	Status() string
}
