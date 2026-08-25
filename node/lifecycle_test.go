package node

import (
	"os"
	"regexp"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Node.Wait() must wait for every component Node.Start() starts.
//
// This is asserted against the source rather than against a live Node, because the
// struct holds concrete types that cannot be stubbed out. The invariant is worth
// encoding somewhere regardless: Orchestrator.Start() was called by Start() while
// Wait() never waited for it, so the process could report a graceful shutdown with a
// relay call still in flight. Component-level lifecycle tests could not see that --
// they exercised Stop()/Wait() on the component directly, which is not how production
// shuts down.
func TestNodeWaitsForEverythingItStarts(t *testing.T) {
	raw, err := os.ReadFile("node.go")
	if err != nil {
		t.Fatalf("failed to read node.go: %v", err)
	}
	// Strip line comments first: node.go carries commented-out component wiring (the
	// reporter), and matching it would report a component that is not actually started.
	src := regexp.MustCompile(`(?m)^\s*//.*$`).ReplaceAll(raw, []byte(""))

	body := func(fn string) string {
		re := regexp.MustCompile(`(?s)func \(n \*Node\) ` + fn + `\(\) \{(.*?)\n\}`)
		m := re.FindSubmatch(src)
		if m == nil {
			t.Fatalf("could not locate func (n *Node) %v()", fn)
		}
		return string(m[1])
	}

	// Start() takes a ctx, so match it separately.
	startRe := regexp.MustCompile(`(?s)func \(n \*Node\) Start\(ctx context\.Context\) \{(.*?)\n\}`)
	sm := startRe.FindSubmatch(src)
	if sm == nil {
		t.Fatal("could not locate func (n *Node) Start(ctx context.Context)")
	}

	comps := func(s, method string) []string {
		re := regexp.MustCompile(`n\.([A-Za-z0-9_]+)\.` + method + `\(`)
		seen := map[string]bool{}
		for _, m := range re.FindAllStringSubmatch(s, -1) {
			seen[m[1]] = true
		}
		out := []string{}
		for k := range seen {
			out = append(out, k)
		}
		sort.Strings(out)
		return out
	}

	started := comps(string(sm[1]), "Start")
	waited := comps(body("Wait"), "Wait")

	assert.NotEmpty(t, started, "sanity: Start() should start something")

	// Pre-existing gaps, out of scope for this change and recorded rather than hidden.
	// Both predate PR #30. Neither issues cross-chain calls, which is why the
	// Orchestrator was the one fixed here.
	knownNotWaited := map[string]string{
		"Dispatcher": "pre-existing: started but never waited on",
		"Mempool":    "pre-existing: has Wait(), started but never waited on",
	}

	waitedSet := map[string]bool{}
	for _, w := range waited {
		waitedSet[w] = true
	}
	for _, c := range started {
		if reason, known := knownNotWaited[c]; known {
			if waitedSet[c] {
				t.Logf("n.%v is now waited on; remove it from knownNotWaited", c)
			} else {
				t.Logf("KNOWN GAP: n.%v -- %v", c, reason)
			}
			continue
		}
		assert.True(t, waitedSet[c],
			"Node.Start() starts n.%v but Node.Wait() never waits for it: "+
				"shutdown can be reported while it is still running", c)
	}
}
