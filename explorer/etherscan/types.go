package etherscan

import "github.com/dapplink-labs/chain-explorer-api/common"

type AccountBalance struct {
	Account string         `json:"account"`
	Balance *common.BigInt `json:"balance"`
}

type ApiResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  T      `json:"result"`
}

type GasTrackerGasOracleResp struct {
	LastBlock       string `json:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}
