//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTokenBalance(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressTokenBalance("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", "token_20", 2)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressTokenBalanceResp{
		Data: make([]types.AddressTokenBalanceData, 0),
	}

	expectData := types.AddressTokenBalanceData{
		Limit:     "2",
		Page:      "1",
		TotalPage: "195",
		TokenList: make([]types.AddressTokenBalanceToken, 2, 20),
	}

	expectData.TokenList[0] = types.AddressTokenBalanceToken{
		Symbol:               "MCO",
		TokenContractAddress: "0x135f915f081e750baa63e9ce97421cc2d9dbfac0",
		HoldingAmount:        "115792089237316200000000000000000000000000000000000000000000",
		PriceUsd:             "0",
		ValueUsd:             "",
		TokenId:              "",
	}

	expectData.TokenList[1] = types.AddressTokenBalanceToken{
		Symbol:               "ShibDoge",
		TokenContractAddress: "0x6adb2e268de2aa1abf6578e4a8119b960e02928f",
		HoldingAmount:        "14720806439512404",
		PriceUsd:             "0.00000000000000003",
		ValueUsd:             "0.44162419318537212",
		TokenId:              "",
	}

	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]

	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, len(data.TokenList), len(expectData.TokenList))

	for i := 0; i < 2; i++ {
		assert.Equal(t, data.TokenList[i].Symbol, expectData.TokenList[i].Symbol)
		assert.Equal(t, data.TokenList[i].TokenContractAddress, expectData.TokenList[i].TokenContractAddress)
		assert.Equal(t, data.TokenList[i].HoldingAmount, expectData.TokenList[i].HoldingAmount)
		assert.Equal(t, data.TokenList[i].PriceUsd, expectData.TokenList[i].PriceUsd)
		assert.Equal(t, data.TokenList[i].ValueUsd, expectData.TokenList[i].ValueUsd)
		assert.Equal(t, data.TokenList[i].TokenId, expectData.TokenList[i].TokenId)
	}
}
