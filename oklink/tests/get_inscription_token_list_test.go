//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionTokenList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionTokenList("btc", "runes", "", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionTokenListResp{
		Data: make([]types.InscriptionTokenListData, 0),
	}

	expectData := types.InscriptionTokenListData{
		Page:      "1",
		Limit:     "1",
		TotalPage: "10000",
		TokenList: make([]types.InscriptionTokenListToken, 0),
	}

	expectToken := types.InscriptionTokenListToken{
		Symbol:             "UNCOMMON•GOODS",
		TokenInscriptionId: "1:0",
		ProtocolType:       "RUNES",
		TotalSupply:        "3362172",
		MintAmount:         "3362172",
		DeployTime:         "0",
		Holder:             "40851",
		TransactionCount:   "3374587",
		CirculatingSupply:  "3362067",
		MintBitwork:        "",
		LimitPerMint:       "1",
		RunesSymbol:        "⧉",
		Divisibility:       "0",
		MintStatus:         "mintable",
		Premint:            "0",
		Burn:               "135",
		MintStartBlock:     "840000",
		MintEndBlock:       "1050000",
		MintCap:            "340282366920938463463374607431768211455",
	}
	expectData.TokenList = append(expectData.TokenList, expectToken)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	token := data.TokenList[0]
	assert.Equal(t, token.Symbol, expectToken.Symbol)
}
