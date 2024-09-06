//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetUtxoRevenueStats(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetUtxoRevenueStats("btc", "", "", "", 2)
	if err != nil {
		t.Errorf("GetUtxoRevenueStats error: %s", err)
		return
	}

	expectResp := &types.UtxoRevenueStatsResp{
		Data: make([]types.UtxoRevenueStatsData, 0),
	}

	expectData := types.UtxoRevenueStatsData{
		Page:               "1",
		Limit:              "2",
		TotalPage:          "2806",
		RevenueHistoryList: make([]types.UtxoRevenueStatsRevenueHistory, 2),
	}

	expectData.RevenueHistoryList[0] = types.UtxoRevenueStatsRevenueHistory{
		Time:           "1715702400000",
		BlockReward:    "181.25",
		TransactionFee: "9.6476646",
	}

	expectData.RevenueHistoryList[1] = types.UtxoRevenueStatsRevenueHistory{
		Time:           "1715616000000",
		BlockReward:    "440.625",
		TransactionFee: "24.94226618",
	}

	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, len(data.RevenueHistoryList), len(expectData.RevenueHistoryList))
}
