package core

import (
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/theta/common"
	score "github.com/thetatoken/thetasubchain/core"
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

// A 200 response with no "result" key at all, or an explicit null, is not an empty
// block range. Accepting it lets the caller advance its checkpoint past a range it
// never actually read.
func TestQueryRejectsMissingOrNullResult(t *testing.T) {
	for _, body := range []string{
		`{"jsonrpc":"2.0","id":74}`,
		`{"jsonrpc":"2.0","id":74,"result":null}`,
	} {
		body := body
		_, err := queryAgainst(t, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, body)
		})
		assert.NotNil(t, err, "response %q must not be accepted as an empty range", body)
	}
}

// A log with no topics would previously panic on Topics[0]. The query filters by
// topic so it should not happen, which is exactly why it must be reported.
func TestQueryRejectsLogWithNoTopics(t *testing.T) {
	_, err := queryAgainst(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"jsonrpc":"2.0","id":74,"result":[{"topics":[],"data":"0x","blockNumber":"0x1"}]}`)
	})
	assert.NotNil(t, err, "a topic-less log must be reported, not indexed blindly")
}

// A log that matches one of our event selectors but whose payload will not decode
// must fail the whole range. Skipping it silently would drop a real transfer; keeping
// it would put an event with a nil nonce or block height into the relay pipeline.
func TestQueryRejectsMalformedMatchingLog(t *testing.T) {
	assert := assert.New(t)
	burnTopic := EventSelectors[score.IMCEventTypeCrossChainVoucherBurnTFuel]

	// Non-hex payload.
	_, err := queryAgainst(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"jsonrpc":"2.0","id":74,"result":[{"topics":[%q],"data":"0xZZZZ","blockNumber":"0x1"}]}`, burnTopic)
	})
	assert.NotNil(err, "a log whose data is not hex must fail the range")

	// Unparseable block number.
	_, err = queryAgainst(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"jsonrpc":"2.0","id":74,"result":[{"topics":[%q],"data":"0x","blockNumber":"latest"}]}`, burnTopic)
	})
	assert.NotNil(err, "a log whose block number will not parse must fail the range")
}
