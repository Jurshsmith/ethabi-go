package human_readable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParsingEventAbi(t *testing.T) {
	humanReadableAbi := "event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)"
	abi, _ := parseABI(&humanReadableAbi)
	assert.Equal(t, abi.Events["Transfer"].Sig, "Transfer(address,address,uint256)")
}
