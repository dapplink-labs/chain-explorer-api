//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionAddressTokenList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionAddressTokenList("btc", "bc1qhuv3dhpnm0wktasd3v0kt6e4aqfqsd0uhfdu7d", "brc20", "", "", "", "", 2)

	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionAddressTokenListResp{
		Data: make([]types.InscriptionAddressTokenListData, 0),
	}

	expectData := types.InscriptionAddressTokenListData{
		Page:      "1",
		Limit:     "2",
		TotalPage: "68",
		TokenList: make([]types.InscriptionAddressTokenListToken, 2),
	}

	expectData.TokenList[0] = types.InscriptionAddressTokenListToken{
		Symbol:             "maibi",
		TokenInscriptionId: "b7456b7e688edea8fb814df146a83a062260b596616bfceff3ae2b9ceb8dbab2i0",
		HoldingAmount:      "8188888888888889300",
		InscriptionAmount:  "2",
		AvailableAmount:    "8188888888888889300",
		TransferableAmount: "0",
		InscriptionNumber:  "",
	}

	expectData.TokenList[1] = types.InscriptionAddressTokenListToken{
		Symbol:             "GPTu",
		TokenInscriptionId: "25e7f4ca11734aa1d1d8c9d9be262b8ea8b09a660a93a826084f7b21b3b41518i0",
		HoldingAmount:      "8000000000000000000",
		InscriptionAmount:  "2",
		AvailableAmount:    "8000000000000000000",
		TransferableAmount: "0",
		InscriptionNumber:  "",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	for i := 0; i < len(data.TokenList); i++ {
		token := data.TokenList[i]
		expectToken := expectData.TokenList[i]

		assert.Equal(t, token.Symbol, expectToken.Symbol)
	}
}
