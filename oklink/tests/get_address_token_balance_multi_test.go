//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTokenBalanceMulti(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressTokenBalanceMulti("eth", "0x85c6627c4ed773cb7c32644b041f58a058b00d30,0xb13a8883d5116b418066c379bc3b3f40d087b8d8", "", "", 2)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressTokenBalanceMultiResp{
		Data: make([]types.AddressTokenBalanceMultiData, 0),
	}

	expectData := types.AddressTokenBalanceMultiData{
		Page:        "1",
		Limit:       "2",
		TotalPage:   "0",
		BalanceList: make([]types.AddressTokenBalanceMultiBalance, 2),
	}

	expectData.BalanceList[0] = types.AddressTokenBalanceMultiBalance{
		Address:              "0xf977814e90da44bfa03b6295a0616a897441acec",
		HoldingAmount:        "400",
		TokenContractAddress: "0x7379cbce70bba5a9871f97d33b391afba377e885",
	}

	expectData.BalanceList[1] = types.AddressTokenBalanceMultiBalance{
		Address:              "0xf977814e90da44bfa03b6295a0616a897441acec",
		HoldingAmount:        "123101078.45198849",
		TokenContractAddress: "0x5c885be435a9b5b55bcfc992d8c085e4e549661e",
	}

	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	// TODO: There is no data returned, so it's impossible to compare the balance-related data.
	//for i := 0; i < len(data.BalanceList); i++ {
	//	balance := data.BalanceList[i]
	//	expectBalance := expectData.BalanceList[i]
	//	assert.Equal(t, balance.Address, expectBalance.Address)
	//	assert.Equal(t, balance.HoldingAmount, expectBalance.HoldingAmount)
	//	assert.Equal(t, balance.TokenContractAddress, expectBalance.TokenContractAddress)
	//}
}
