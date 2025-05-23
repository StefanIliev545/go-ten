package common

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
)

type PrivateTransactionsQueryResponse struct {
	Receipts types.Receipts
	Total    uint64
}

type TransactionListingResponse struct {
	TransactionsData []PublicTransaction
	Total            uint64
}

type BatchListingResponse struct {
	BatchesData []PublicBatch
	Total       uint64
}

type BatchListingResponseDeprecated struct {
	BatchesData []PublicBatchDeprecated
	Total       uint64
}

type BlockListingResponse struct {
	BlocksData []PublicBlock
	Total      uint64
}

type RollupListingResponse struct {
	RollupsData []PublicRollup
	Total       uint64
}

type PublicTransaction struct {
	TransactionHash TxHash
	BatchHeight     *big.Int
	BatchTimestamp  uint64
	Finality        FinalityType
}

type PublicBatch struct {
	SequencerOrderNo *big.Int              `json:"sequence"`
	FullHash         common.Hash           `json:"fullHash"`
	Height           *big.Int              `json:"height"`
	TxCount          *big.Int              `json:"txCount"`
	Header           *BatchHeader          `json:"header"`
	EncryptedTxBlob  EncryptedTransactions `json:"encryptedTxBlob"`
}

// TODO (@will) remove when tenscan UI has been updated
type PublicBatchDeprecated struct {
	BatchHeader
	TxHashes []TxHash `json:"txHashes"`
}

type PublicRollup struct {
	ID        *big.Int
	Hash      string
	FirstSeq  *big.Int
	LastSeq   *big.Int
	Timestamp uint64
	Header    *RollupHeader
	L1Hash    string
}

type PublicBlock struct {
	BlockHeader types.Header `json:"blockHeader"`
	RollupHash  common.Hash  `json:"rollupHash"`
}

type FinalityType string

const (
	MempoolPending FinalityType = "Pending"
	BatchFinal     FinalityType = "Final"
)

type QueryPagination struct {
	Offset uint64
	Size   uint
}

func (p *QueryPagination) UnmarshalJSON(data []byte) error {
	// Use a temporary struct to avoid infinite unmarshalling loop
	type Temp struct {
		Size   uint `json:"size"`
		Offset uint64
	}

	var temp Temp
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if temp.Size < 1 || temp.Size > 100 {
		return fmt.Errorf("size must be between 1 and 100")
	}

	p.Size = temp.Size
	p.Offset = temp.Offset
	return nil
}

type TenNetworkInfo struct {
	NetworkConfig             NetworkConfigAddress
	EnclaveRegistry           EnclaveRegistryAddress
	CrossChain                CrossChainAddress
	DataAvailabilityRegistry  DARegistryAddress
	L1MessageBus              L1MessageBusAddress
	L2MessageBus              L2MessageBusAddress
	L1Bridge                  L1BridgeAddress
	L2Bridge                  L2BridgeAddress
	L1CrossChainMessenger     L1CrossChainMessengerAddress
	L2CrossChainMessenger     L2CrossChainMessengerAddress
	TransactionsPostProcessor TransactionPostProcessorAddress
	SystemContractsUpgrader   SystemContractsUpgraderAddress
	L1StartHash               common.Hash
	PublicSystemContracts     map[string]common.Address
	AdditionalContracts       []*NamedAddress
}

// NetworkConfigAddresses return type of the addresses function on the NetworkConfig contract
type NetworkConfigAddresses struct {
	EnclaveRegistry          EnclaveRegistryAddress
	CrossChain               CrossChainAddress
	DataAvailabilityRegistry DARegistryAddress
	L1MessageBus             L1MessageBusAddress
	L1Bridge                 L1BridgeAddress
	L2Bridge                 L2BridgeAddress
	L1CrossChainMessenger    L1CrossChainMessengerAddress
	L2CrossChainMessenger    L2CrossChainMessengerAddress
	AdditionalContracts      []*NamedAddress // Dynamically named additional contracts
}

// NamedAddress matches the Solidity struct
type NamedAddress struct {
	Name string
	Addr common.Address
}
