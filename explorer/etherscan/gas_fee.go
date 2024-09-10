package etherscan

import (
	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
)

func (cea *ChainExplorerAdaptor) GetEstimateGasFee(request *gas_fee.GasEstimateFeeRequest) (*gas_fee.GasEstimateFeeResponse, error) {
	resp := &ApiResponse[GasTrackerGasOracleResp]{}
	err := cea.baseClient.Call(ChainExplorerName, "gastracker", "gasoracle", "", nil, resp)

	if err != nil {
		return nil, err
	}
	return &gas_fee.GasEstimateFeeResponse{
		ChainFullName:         "Ethereum",
		ChainShortName:        "ETH",
		Symbol:                "ETH",
		BestTransactionFee:    "",
		BestTransactionFeeSat: "",
		RecommendedGasPrice:   resp.Result.ProposeGasPrice,
		RapidGasPrice:         resp.Result.FastGasPrice,
		StandardGasPrice:      resp.Result.ProposeGasPrice,
		SlowGasPrice:          resp.Result.SafeGasPrice,
		BaseFee:               resp.Result.SuggestBaseFee,
		GasUsedRatio:          resp.Result.GasUsedRatio,
	}, nil
}
