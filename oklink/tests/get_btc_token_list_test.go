//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcTokenList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcTokenList("", "", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcTokenListResp{
		Data: make([]types.BtcTokenListData, 0),
	}

	expectData := types.BtcTokenListData{
		Page:      "1",
		Limit:     "20",
		TotalPage: "500",
		TokenList: make([]types.BtcTokenListToken, 0),
	}

	expectToken := types.BtcTokenListToken{
		Token:             "ordi",
		DeployTime:        "1678248991000",
		InscriptionId:     "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0",
		InscriptionNumber: "348020",
		TotalSupply:       "21000000",
		MintAmount:        "21000000",
		TransactionCount:  "484169",
		Holder:            "13080",
		MintRate:          "1",
	}
	expectData.TokenList = append(expectData.TokenList, expectToken)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	token := data.TokenList[0]
	assert.Equal(t, token.Token, expectToken.Token)
}
