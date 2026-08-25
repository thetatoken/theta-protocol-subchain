package core

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	ec "github.com/thetatoken/thetasubchain/eth/ethclient"
)

func chainIDNode(t *testing.T, result string, delay time.Duration) *ec.Client {
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
		fmt.Fprintf(w, `{"jsonrpc":"2.0","id":%s,"result":%q}`, req.ID, result)
	}))
	t.Cleanup(srv.Close)
	c, err := ec.Dial(srv.URL)
	if err != nil {
		t.Fatalf("dial: %v", err)
	}
	t.Cleanup(c.Close)
	return c
}

// The misconfiguration this guards against: an endpoint pointing at a different
// subchain. Predeploy addresses are identical across subchains, so the wrong endpoint
// answers plausibly at the expected TokenBank addresses and its events would be relayed
// under this chain's identity.
func TestVerifyChainIDRejectsTheWrongChain(t *testing.T) {
	assert := assert.New(t)

	err := VerifyChainID(chainIDNode(t, "0x581bb", 0), big.NewInt(360890), "subchain") // 360891
	assert.NotNil(err)
	assert.Contains(err.Error(), "360891")
	assert.Contains(err.Error(), "360890")
}

func TestVerifyChainIDAcceptsTheRightChain(t *testing.T) {
	assert.Nil(t, VerifyChainID(chainIDNode(t, "0x581ba", 0), big.NewInt(360890), "subchain"))
}

// A misconfigured or absent expectation must fail closed, not silently accept anything.
func TestVerifyChainIDRejectsAMissingExpectation(t *testing.T) {
	assert := assert.New(t)
	assert.NotNil(VerifyChainID(chainIDNode(t, "0x581ba", 0), nil, "subchain"))
	assert.NotNil(VerifyChainID(chainIDNode(t, "0x581ba", 0), big.NewInt(0), "subchain"))
	assert.NotNil(VerifyChainID(nil, big.NewInt(360890), "subchain"))
}

// The check runs at startup; an unresponsive endpoint must fail startup rather than
// hang it.
func TestVerifyChainIDIsBounded(t *testing.T) {
	start := time.Now()
	err := VerifyChainID(chainIDNode(t, "0x581ba", 60*time.Second), big.NewInt(360890), "subchain")
	assert.NotNil(t, err)
	assert.Less(t, time.Since(start), chainIDQueryTimeout+10*time.Second)
}
