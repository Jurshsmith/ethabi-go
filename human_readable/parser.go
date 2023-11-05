package human_readable

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/jurshsmith/ethabi-go"
)

func parseABI(humanReadableAbi *string) (abi.ABI, error) {
	if strings.HasPrefix(*humanReadableAbi, string(ethabi.Event)) == true {
		eventAbi := ethabi.New(humanReadableAbi, ethabi.Event)
		eventAbiJson, _ := json.Marshal(&eventAbi)
		eventAbiJsonString := fmt.Sprintf("[%s]", string(eventAbiJson))
		abiReader := strings.NewReader(eventAbiJsonString)

		return abi.JSON(abiReader)
	}

	return abi.ABI{}, errors.New("HumanReadable ABI is either unsupported or unvalid")
}
