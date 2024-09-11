package etherscan

import (
	"errors"
	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
)

func (cea *ChainExplorerAdaptor) GetEstimateGasFee(request *gas_fee.GasEstimateFeeRequest) (*gas_fee.GasEstimateFeeResponse, error) {
	resp := &GasTrackerGasOracleResp{}
	err := cea.baseClient.Call(ChainExplorerName, "gastracker", "gasoracle", "", nil, &resp)

	if err != nil {
		return nil, err
	}
	if resp == nil || *resp == (GasTrackerGasOracleResp{}) {
		return nil, errors.New("result is empty")
	}
	return &gas_fee.GasEstimateFeeResponse{
		ChainFullName:         "Ethereum",
		ChainShortName:        "ETH",
		Symbol:                "ETH",
		BestTransactionFee:    "",
		BestTransactionFeeSat: "",
		RecommendedGasPrice:   resp.ProposeGasPrice,
		RapidGasPrice:         resp.FastGasPrice,
		StandardGasPrice:      resp.ProposeGasPrice,
		SlowGasPrice:          resp.SafeGasPrice,
		BaseFee:               resp.SuggestBaseFee,
		GasUsedRatio:          resp.GasUsedRatio,
	}, nil
}
