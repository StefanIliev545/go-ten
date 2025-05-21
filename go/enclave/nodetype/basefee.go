package nodetype

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

// calcNextBaseFee calculates the next base fee following the
// EIP-1559 algorithm. It takes the previous base fee, gas used
// in the previous batch and the gas limit of that batch.
func calcNextBaseFee(parentBaseFee *big.Int, gasUsed, gasLimit uint64) *big.Int {
	if parentBaseFee == nil {
		return big.NewInt(params.InitialBaseFee)
	}
	// If gas limit is zero just return previous base fee.
	if gasLimit == 0 {
		return new(big.Int).Set(parentBaseFee)
	}

	target := gasLimit / params.ElasticityMultiplier
	if target == 0 {
		return new(big.Int).Set(parentBaseFee)
	}

	baseFee := new(big.Int).Set(parentBaseFee)
	changeDenom := new(big.Int).SetUint64(params.BaseFeeChangeDenominator)

	if gasUsed == target {
		return baseFee
	} else if gasUsed > target {
		delta := new(big.Int).Mul(baseFee, big.NewInt(int64(gasUsed-target)))
		delta.Div(delta, new(big.Int).SetUint64(target))
		delta.Div(delta, changeDenom)
		if delta.Sign() == 0 {
			delta.SetInt64(1)
		}
		return baseFee.Add(baseFee, delta)
	}

	delta := new(big.Int).Mul(baseFee, big.NewInt(int64(target-gasUsed)))
	delta.Div(delta, new(big.Int).SetUint64(target))
	delta.Div(delta, changeDenom)
	if delta.Sign() == 0 {
		delta.SetInt64(1)
	}
	if baseFee.Cmp(delta) <= 0 {
		return big.NewInt(0)
	}
	return baseFee.Sub(baseFee, delta)
}
