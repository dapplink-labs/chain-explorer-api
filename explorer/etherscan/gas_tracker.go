package etherscan

import (
	"time"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/gas_tracker"
)

func (cea *ChainExplorerAdaptor) GasEstimate(req *gas_tracker.GasEstimateRequest) (*gas_tracker.GasEstimateResponse, error) {
	param := common.M{
		"gasPrice": req.GasLimit,
	}

	var confTime string
	err := cea.baseClient.Call("gastracker", "gasestimate", "GET", param, &confTime)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(confTime + "s")
	if err != nil {
		return nil, err
	}

	return &gas_tracker.GasEstimateResponse{
		ConfirmationTimes: gas_tracker.GasConfirmationTimes{
			Medium: duration,
		},
	}, nil
}

func (cea *ChainExplorerAdaptor) GasOracle(req *gas_tracker.GasOracleRequest) (*gas_tracker.GasOracleResponse, error) {
	param := common.M{}

	var response gas_tracker.GasOracleResponse
	err := cea.baseClient.Call("gastracker", "gasoracle", "GET", param, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
