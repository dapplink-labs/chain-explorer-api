//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetUtxoTransactionStats(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetUtxoTransactionStats("btc", "", "", "", 2)
	if err != nil {
		t.Errorf("GetUtxoTransactionStats error: %v", err)
	}

	expectResp := &types.UtxoTransactionStatsResp{
		Data: make([]types.UtxoTransactionStatsData, 0),
	}

	expectData := types.UtxoTransactionStatsData{
		Page:                   "1",
		Limit:                  "2",
		TotalPage:              "2805",
		TransactionHistoryList: make([]types.UtxoTransactionStatsTransactionHistory, 2),
	}

	expectData.TransactionHistoryList[0] = types.UtxoTransactionStatsTransactionHistory{
		Time:                      "1715616000000",
		TotalTransactionCount:     "480918",
		NormalTransactionCount:    "454299",
		AtomicalsTransactionCount: "1888",
		StampTransactionCount:     "859",
		OrdinalsTransactionCount:  "32594",
	}

	expectData.TransactionHistoryList[1] = types.UtxoTransactionStatsTransactionHistory{
		Time:                      "1715529600000",
		TotalTransactionCount:     "596134",
		NormalTransactionCount:    "573757",
		AtomicalsTransactionCount: "2478",
		StampTransactionCount:     "739",
		OrdinalsTransactionCount:  "42489",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)

	for i := 0; i < len(data.TransactionHistoryList); i++ {
		txHistory := data.TransactionHistoryList[i]
		expectTxHistory := expectData.TransactionHistoryList[i]
		assert.Equal(t, txHistory.OrdinalsTransactionCount, expectTxHistory.OrdinalsTransactionCount)
	}
}
