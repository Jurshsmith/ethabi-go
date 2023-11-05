package ethabi

import (
	"testing"

	goEthereumAbi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/stretchr/testify/assert"
)

func TestParsingEventAbis(t *testing.T) {
	humanReadableAbis := []string{
		"event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)",
		"event Transfer(address indexed from, address indexed to, uint256 tokenId)",
		"event Transfer(address indexed from, address indexed to, uint256     tokenId)",
		"event Transfer(address   indexed    from, address indexed   to, uint256 tokenId)",
	}

	for _, abi := range humanReadableAbis {
		abi, err := ParseABI(&abi)
		assert.Nil(t, err)
		assert.Equal(t, abi.Events["Transfer"].Sig, "Transfer(address,address,uint256)")
	}
}

func TestParsingNonEventAbis(t *testing.T) {
	invalidHumanReadableAbi := "contract SampleContract{}"
	abi, err := ParseABI(&invalidHumanReadableAbi)

	assert.NotNil(t, err)
	assert.Equal(t, abi, goEthereumAbi.ABI{})
}
