//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressBalanceFills(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressBalanceFills("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", 2)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressBalanceFillsResp{
		Data: make([]types.AddressBalanceFillsData, 0),
	}

	expectData := types.AddressBalanceFillsData{
		Limit:     "2",
		Page:      "1",
		TotalPage: "195",
		TokenList: make([]types.AddressBalanceFillsToken, 2),
	}

	expectData.TokenList[0] = types.AddressBalanceFillsToken{
		Token:                "USDT",
		TokenId:              "",
		HoldingAmount:        "402413.20994",
		TotalTokenValue:      "168.38785369",
		Change24h:            "0.00013",
		PriceUsd:             "1.00054",
		ValueUsd:             "402630.5130733676",
		TokenContractAddress: "0xdac17f958d2ee523a2206206994597c13d831ec7",
	}

	expectData.TokenList[1] = types.AddressBalanceFillsToken{
		Token:                "CRO",
		TokenId:              "",
		HoldingAmount:        "4800000",
		TotalTokenValue:      "156.86203294",
		Change24h:            "-0.032126",
		PriceUsd:             "0.07813984132122903",
		ValueUsd:             "375071.238341899344",
		TokenContractAddress: "0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b",
	}

	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]

	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, len(data.TokenList), len(expectData.TokenList))

	for i := 0; i < 2; i++ {
		// Only some fields can be compared here,
		// because the values of the returned data are dynamic and cannot be predicted
		assert.Equal(t, data.TokenList[i].Token, expectData.TokenList[i].Token)
		assert.Equal(t, data.TokenList[i].TokenId, expectData.TokenList[i].TokenId)
		assert.Equal(t, data.TokenList[i].HoldingAmount, expectData.TokenList[i].HoldingAmount)
	}
}
