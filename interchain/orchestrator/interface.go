package orchestrator

import (
	"context"
)

type ChainOrchestrator interface {
	Start(ctx context.Context)
	Stop()
	Wait()
}
