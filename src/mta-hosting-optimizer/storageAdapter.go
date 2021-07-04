package hostingoptimizer

import (
	"context"
)

type storageAdapter interface {
	getInactiveIPHosts(ctx context.Context,
		activeIPsCount int) ([]string, error)
}
