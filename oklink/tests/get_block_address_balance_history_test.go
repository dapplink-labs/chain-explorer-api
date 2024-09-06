//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBlockAddressBalanceHistory(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetBlockAddressBalanceHistory("eth", "18376634", "0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701", "")
	if err != nil {
		t.Error(err)
	}

	data := resp.Data[0]

	expectResp := &types.BlockAddressBalanceHistoryResp{
		Data: make([]types.BlockAddressBalanceHistoryData, 0),
	}

	expectData := types.BlockAddressBalanceHistoryData{
		Address:              "0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701",
		Height:               "18376634",
		Balance:              "5.895934930980364414",
		BalanceSymbol:        "ETH",
		TokenContractAddress: "",
		BlockTime:            "1697624735000",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	assert.Equal(t, data.Address, expectData.Address)
	assert.Equal(t, data.Height, expectData.Height)
	assert.Equal(t, data.Balance, expectData.Balance)
	assert.Equal(t, data.BalanceSymbol, expectData.BalanceSymbol)
	assert.Equal(t, data.TokenContractAddress, expectData.TokenContractAddress)
	assert.Equal(t, data.BlockTime, expectData.BlockTime)
}
