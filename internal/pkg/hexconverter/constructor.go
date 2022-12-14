package hexconverter

import (
	"github.com/ethereum/go-ethereum/common/hexutil" // I used lib, but it is easily switchable
	"math/big"
)

type HexConverter struct {
}

func New() *HexConverter {
	return &HexConverter{}
}

func (h HexConverter) DecodeBig(input string) (*big.Int, error) {
	output, err := hexutil.DecodeBig(input)
	if err != nil {
		return &big.Int{}, err
	}

	return output, nil
}

func (h HexConverter) EncodeUint64(input uint64) string {
	output := hexutil.EncodeUint64(input)

	return output
}
