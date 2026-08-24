package core

import (
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/theta/common"
)

// The caller advances its scan checkpoint on the strength of this result, so every
// failure mode has to be reported as an error. Returning an empty slice makes a
// failed query indistinguishable from an empty block range, and the range is then
// skipped permanently -- nothing rescans it.

func queryAgainst(t *testing.T, handler http.HandlerFunc) ([]interface{}, error) {
	t.Helper()
	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)
	events, err := QueryInterChainEventLog(
		big.NewInt(360890), big.NewInt(1), big.NewInt(100),
		common.Address{}, common.Address{}, common.Address{}, common.Address{},
		"", srv.URL)
	out := make([]interface{}, len(events))
	return out, err
}

func TestQueryReportsHttpErrorStatus(t *testing.T) {
	_, err := queryAgainst(t, func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "bad gateway", http.StatusBadGateway)
	})
	assert.NotNil(t, err, "an HTTP 502 must not be reported as an empty block range")
	assert.Contains(t, err.Error(), "502")
}

func TestQueryReportsJsonRpcError(t *testing.T) {
	_, err := queryAgainst(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"jsonrpc":"2.0","id":74,"error":{"code":-32000,"message":"block range too large"}}`)
	})
	assert.NotNil(t, err, "a JSON-RPC error must not be reported as an empty block range")
	assert.Contains(t, err.Error(), "block range too large")
}

func TestQueryReportsMalformedResponse(t *testing.T) {
	_, err := queryAgainst(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"jsonrpc":"2.0","id":74,"result":`) // truncated
	})
	assert.NotNil(t, err, "an undecodable response must not be reported as an empty block range")
}

func TestQueryReportsUnreachableNode(t *testing.T) {
	events, err := QueryInterChainEventLog(
		big.NewInt(360890), big.NewInt(1), big.NewInt(100),
		common.Address{}, common.Address{}, common.Address{}, common.Address{},
		"", "http://127.0.0.1:1") // nothing listening
	assert.NotNil(t, err)
	assert.Nil(t, events)
}

// The success path must still report no error for a genuinely empty range, or the
// witness would never make progress.
func TestQueryAcceptsAGenuinelyEmptyRange(t *testing.T) {
	events, err := queryAgainst(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"jsonrpc":"2.0","id":74,"result":[]}`)
	})
	assert.Nil(t, err)
	assert.Empty(t, events)
}
