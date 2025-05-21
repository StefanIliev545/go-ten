package nodetype

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/misc"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// calcNextBaseFee calculates the next base fee using go-ethereum's
// implementation of the EIP-1559 algorithm. The parent header must
// contain the previous base fee, gas limit and gas used fields.
func calcNextBaseFee(cfg *params.ChainConfig, parent *gethtypes.Header) *big.Int {
	if parent == nil {
		return big.NewInt(params.InitialBaseFee)
	}

	next, err := misc.CalcBaseFee(cfg, parent)
	if err != nil {
		return big.NewInt(params.InitialBaseFee)
	}
	return next
}
