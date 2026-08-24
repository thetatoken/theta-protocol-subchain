package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/theta/common"
)

// The denom travels with the event and is therefore attacker-influenced, so the
// extractor has to reject anything that is not exactly the expected shape rather
// than returning a zero address that a caller might treat as valid.
func TestExtractContractAddressFromDenom(t *testing.T) {
	assert := assert.New(t)

	addr, err := ExtractContractAddressFromDenom("361/2/0xaf537fb7e4c77c97403de94ce141b7edb9f7fcf0")
	assert.Nil(err)
	assert.Equal(common.HexToAddress("0xaf537fb7e4c77c97403de94ce141b7edb9f7fcf0"), addr)

	for _, bad := range []string{
		"",
		"361/2",                  // too few parts
		"361/2/0xaf537fb7/extra", // too many parts
		"361/2/not-an-address",   // not hex
		"361/2/0xaf537fb7e4c77c97403de94ce141b7edb9f7fc",      // too short
		"361/0/0x0000000000000000000000000000000000000000000", // too long
	} {
		_, err := ExtractContractAddressFromDenom(bad)
		assert.NotNil(err, "expected %q to be rejected", bad)
	}
}
