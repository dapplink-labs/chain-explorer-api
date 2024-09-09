package oklink

import "github.com/dapplink-labs/chain-explorer-api/common/block"

func (cea *ChainExplorerAdaptor) GetBlockAddressBalanceHistory(req *block.BlockAddressBalanceHistoryRequest) (*block.BlockAddressBalanceHistoryResponse, error) {
	reqMap := map[string]interface{}{
		"chainShortName": req.ChainShortName,
		"height":         req.Height,
		"address":        req.Address,
	}

	if req.TokenContractAddress != "" {
		reqMap["tokenContractAddress"] = req.TokenContractAddress
	}

	resp := &block.BlockAddressBalanceHistoryResponse{}

	err := cea.baseClient.Call(ChainExplorerName, "block",
		"/api/v5/explorer/block/address-balance-history", reqMap, resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetBlockByNumber(req *block.BlockByNumberRequest) (*block.BlockByNumberResponse, error) {
	reqMap := map[string]interface{}{
		"chainShortName": req.ChainShortName,
		"height":         req.Height,
	}

	resp := &block.BlockByNumberResponse{}

	err := cea.baseClient.Call(ChainExplorerName, "block",
		"/api/v5/explorer/block/address-balance-history", reqMap, resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ?????
func (cea *ChainExplorerAdaptor) GetBlockNumber(req *block.BlockNumberRequest) (*block.BlockNumberResponse, error) {
	reqMap := map[string]interface{}{
		"chainShortName": req.ChainShortName,
		"timestamp":      req.Timestamp,
		"closest":        req.Closest,
	}

	resp := &block.BlockNumberResponse{}

	err := cea.baseClient.Call(ChainExplorerName, "block",
		"/api/v5/explorer/block/address-balance-history", reqMap, resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
