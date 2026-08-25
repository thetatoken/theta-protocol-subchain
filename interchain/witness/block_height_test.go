package witness

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	scom "github.com/thetatoken/thetasubchain/common"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

func heightNode(t *testing.T, result string, rpcErr string, delay time.Duration) *ec.Client {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			ID json.RawMessage `json:"id"`
		}
		json.NewDecoder(r.Body).Decode(&req)
		if delay > 0 {
			select {
			case <-time.After(delay):
			case <-r.Context().Done():
				return
			}
		}
		if rpcErr != "" {
			fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"error":{"code":-32000,"message":%q}}`, req.ID, rpcErr)
			return
		}
		fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"result":%q}`, req.ID, result)
	}))
	t.Cleanup(srv.Close)

	client, err := ec.Dial(srv.URL)
	if err != nil {
		t.Fatalf("failed to dial: %v", err)
	}
	t.Cleanup(client.Close)
	return client
}

// update() feeds the mainchain height straight into CalculateDynasty(), which divides
// by it -- and big.Int.Div panics on a nil operand. On a freshly started node the
// height is still nil, so updateMainchainBlockHeight() MUST report failure rather
// than leaving the field nil and returning silently, or one failed RPC call takes the
// validator down during exactly the degraded conditions that caused it.
func TestUpdateMainchainBlockHeightReportsFailure(t *testing.T) {
	assert := assert.New(t)

	mw := &MetachainWitness{mainchainEthRpcClient: heightNode(t, "", "connection refused", 0)}

	err := mw.updateMainchainBlockHeight()
	assert.NotNil(err, "a failed height read must be reported, not swallowed")
	assert.Nil(mw.mainchainBlockHeight, "the height must remain unset after a failure")

	// The contract update() relies on: a nil height would panic here.
	assert.Panics(func() { scom.CalculateDynasty(mw.mainchainBlockHeight) },
		"this is why the error must be propagated")
}

func TestUpdateMainchainBlockHeightSucceeds(t *testing.T) {
	assert := assert.New(t)

	mw := &MetachainWitness{mainchainEthRpcClient: heightNode(t, "0x2209a59", "", 0)}

	assert.Nil(mw.updateMainchainBlockHeight())
	assert.Equal(big.NewInt(0x2209a59), mw.mainchainBlockHeight)
	assert.NotPanics(func() { scom.CalculateDynasty(mw.mainchainBlockHeight) })
}

// The height reads run on every witness tick, so an unresponsive node must not hold
// the loop open indefinitely.
func TestBlockHeightReadIsBounded(t *testing.T) {
	assert := assert.New(t)

	mw := &MetachainWitness{
		mainchainEthRpcClient: heightNode(t, "0x1", "", 60*time.Second),
		subchainEthRpcClient:  heightNode(t, "0x1", "", 60*time.Second),
	}

	start := time.Now()
	err := mw.updateMainchainBlockHeight()
	elapsed := time.Since(start)

	assert.NotNil(err)
	assert.Less(elapsed, blockHeightQueryTimeout+10*time.Second,
		"the head-height read must carry a deadline")
}

// The test above proves the helper reports failure. It does NOT prove update() acts
// on that report -- removing the early return would leave it green, which is the same
// unit-versus-wiring gap that hid S5 and T4. This drives update() itself.
func TestUpdateSkipsTheTickWhenTheHeightIsUnavailable(t *testing.T) {
	mw := &MetachainWitness{
		mainchainEthRpcClient: heightNode(t, "", "connection refused", 0),
		subchainEthRpcClient:  heightNode(t, "", "connection refused", 0),
	}

	// With the early return in place, update() bails before CalculateDynasty() and
	// before any of the collection work, none of which is wired up here. Without it,
	// CalculateDynasty(nil) panics.
	assert.NotPanics(t, func() { mw.update() },
		"update() must skip the tick when the mainchain height could not be read")
	assert.Nil(t, mw.mainchainBlockHeight)
}

// Stop() must actually stop the loop. All three components stored a derived context
// but handed mainloop the parent, so cancel() cancelled a context nothing was
// selecting on and Wait() would hang. Production masked it: Node.Stop() cancels the
// shared parent.
func TestStopCancelsTheDerivedContext(t *testing.T) {
	mw := &MetachainWitness{
		wg:                    &sync.WaitGroup{},
		updateInterval:        1000,
		mainchainEthRpcClient: heightNode(t, "", "connection refused", 0),
		subchainEthRpcClient:  heightNode(t, "", "connection refused", 0),
	}

	// A parent that is never cancelled: only the derived context can end the loop.
	mw.Start(context.Background())

	done := make(chan struct{})
	go func() {
		mw.Stop()
		mw.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(15 * time.Second):
		t.Fatal("Stop(); Wait() hung: mainloop is not watching the context Stop() cancels")
	}
}

// The simulated witness repeats the same lifecycle shape, so it gets the same test.
// The previous round covered only MetachainWitness, which is how the ticker race
// survived in all three components.
func TestSimulatedWitnessStopIsClean(t *testing.T) {
	mw := &SimulatedMetachainWitness{wg: &sync.WaitGroup{}}
	mw.Start(context.Background())

	done := make(chan struct{})
	go func() { mw.Stop(); mw.Wait(); close(done) }()

	select {
	case <-done:
	case <-time.After(15 * time.Second):
		t.Fatal("SimulatedMetachainWitness Stop(); Wait() hung")
	}
}
