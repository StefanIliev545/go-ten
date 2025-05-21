package nodetype

import (
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ten-protocol/go-ten/go/common"
)

// calcNextBaseFee determines the following base fee using go-ethereum's
// EIP-1559 implementation. The calculation is performed on the provided
// parent batch header using the supplied chain configuration.
func calcNextBaseFee(config *params.ChainConfig, parentHeader *common.BatchHeader) *big.Int {
	if parentHeader == nil {
		return big.NewInt(params.InitialBaseFee)
	}

	ethHeader := common.ConvertBatchHeaderToHeader(parentHeader)
	return misc.CalcBaseFee(config, ethHeader)
}
