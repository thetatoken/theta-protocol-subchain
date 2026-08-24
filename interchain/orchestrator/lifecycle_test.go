package orchestrator

import (
	"context"
	"sync"
	"testing"
	"time"
)

// Same lifecycle contract as the witnesses: Stop() must end mainloop, and Wait() must
// return. Both halves are load-bearing -- passing the parent context to mainloop, or
// omitting wg.Done(), each hangs on its own.
//
// Run this under -race as well: mainloop and Stop() previously touched a shared ticker.
func TestOrchestratorStopIsClean(t *testing.T) {
	oc := &Orchestrator{
		wg:                    &sync.WaitGroup{},
		updateInterval:        1000,
		firstContradictedTime: make(map[string]time.Time),
		eventProcessedTime:    make(map[string]time.Time),
	}

	// A parent that is never cancelled: only the derived context can end the loop.
	oc.Start(context.Background())

	done := make(chan struct{})
	go func() { oc.Stop(); oc.Wait(); close(done) }()

	select {
	case <-done:
	case <-time.After(15 * time.Second):
		t.Fatal("Orchestrator Stop(); Wait() hung")
	}
}
