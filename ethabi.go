package ethabi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/jurshsmith/ethabi-go/utils"
)

func ParseABI(humanReadableAbi *string) (abi.ABI, error) {
	if strings.HasPrefix(*humanReadableAbi, string(Event)) == true {
		eventAbi := New(humanReadableAbi, Event)
		eventAbiJson, _ := json.Marshal(&eventAbi)
		eventAbiJsonString := fmt.Sprintf("[%s]", string(eventAbiJson))
		abiReader := strings.NewReader(eventAbiJsonString)

		return abi.JSON(abiReader)
	}

	return abi.ABI{}, errors.New("HumanReadableABI is either invalid or unsupported")
}

func GetABIName(humanReadableAbi *string) string {
	abiTokens := strings.Split(*humanReadableAbi, "(")
	abiTokens = strings.Split(abiTokens[0], " ")

	return abiTokens[1]
}

type Abi struct {
	Type      string     `json:"type"`
	Name      string     `json:"name"`
	Inputs    []AbiInput `json:"inputs"`
	Anonymous bool       `json:"anonymous"`
}

type AbiType string

const (
	Event = "event"
)

func New(humanReadableAbi *string, Type AbiType) Abi {
	name := GetABIName(humanReadableAbi)
	inputs := NewAbiInputs(humanReadableAbi)

	return Abi{Type: string(Type), Name: name, Inputs: inputs, Anonymous: false}
}

type AbiInput struct {
	Type         string `json:"type"`
	InternalType string `json:"internalType"`
	Name         string `json:"name"`
	Indexed      bool   `json:"indexed"`
}

func NewAbiInputs(humanReadableAbi *string) []AbiInput {
	inputTokens := getInputTokens(humanReadableAbi)

	return utils.MapOverSlice(inputTokens, NewAbiInput)
}

func getInputTokens(humanReadableAbi *string) []string {
	abi := removeInputTokensBeforePart(humanReadableAbi)
	abi = removeInputTokensAfterPart(&abi)

	return strings.Split(abi, ",")
}

func removeInputTokensBeforePart(humanReadableAbi *string) string {
	abiTokens := strings.Split(*humanReadableAbi, "(")
	return abiTokens[1]
}

func removeInputTokensAfterPart(humanReadableAbi *string) string {
	abiTokens := strings.Split(*humanReadableAbi, ")")
	return abiTokens[0]
}

func NewAbiInput(inputToken *string) AbiInput {
	sanitizedInputToken := strings.TrimSpace(*inputToken)
	inputTokenParts := strings.Split(sanitizedInputToken, " ")

	theType := inputTokenParts[0]
	theType = maybeNormalizeType(&theType)

	name, indexed := getNameAndIndexStatus(&inputTokenParts)

	return AbiInput{Type: theType, InternalType: theType, Name: name, Indexed: indexed}
}

func maybeNormalizeType(theType *string) string {
	if *theType == "uint" {
		return "uint256"
	}

	return *theType
}

func getNameAndIndexStatus(inputTokenParts *[]string) (string, bool) {
	if len(*inputTokenParts) == 3 {
		return (*inputTokenParts)[2], true
	} else {
		return (*inputTokenParts)[1], false
	}
}
