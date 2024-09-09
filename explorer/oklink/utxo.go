package oklink

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/utxo"
)

const ChainExplorerModelUtxo = "utxo"

func (cea *ChainExplorerAdaptor) GetUtxoTransactionStats(req *utxo.TransactionStatsRequest) (*utxo.TransactionStatsResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName": req.ChainShortName,
		"startTime":      req.StartTime,
		"endTime":        req.EndTime,
		"page":           req.Page,
		"limit":          req.Limit,
	}
	resp := &utxo.TransactionStatsResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelUtxo,
		"/api/v5/explorer/utxo/transaction-stats", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetUtxoRevenueStats(req *utxo.RevenueStatsRequest) (*utxo.RevenueStatsResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName": req.ChainShortName,
		"startTime":      req.StartTime,
		"endTime":        req.EndTime,
		"page":           req.Page,
		"limit":          req.Limit,
	}
	resp := &utxo.RevenueStatsResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelUtxo,
		"/api/v5/explorer/utxo/revenue-stats", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}
