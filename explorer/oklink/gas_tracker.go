package oklink

import "github.com/dapplink-labs/chain-explorer-api/common/gas_tracker"

func (cea *ChainExplorerAdaptor) GasOracle(req *gas_tracker.GasOracleRequest) (*gas_tracker.GasOracleResponse, error) {
	reqMap := map[string]interface{}{
		"chainShortName": req.ExplorerName,
	}

	resp := &gas_tracker.GasOracleResponse{}

	err := cea.baseClient.Call(ChainExplorerName, "gastracker",
		"/api/v5/explorer/gastracker/gasoracle", reqMap, resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GasEstimate(req *gas_tracker.GasEstimateRequest) (*gas_tracker.GasEstimateResponse, error) {
	reqMap := map[string]interface{}{
		"chainShortName": req.ExplorerName,
		"gasLimit":       req.GasLimit,
	}

	resp := &gas_tracker.GasEstimateResponse{}

	err := cea.baseClient.Call(ChainExplorerName, "gastracker",
		"/api/v5/explorer/gastracker/gasestimate", reqMap, resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
