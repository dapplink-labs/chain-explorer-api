//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcTransactionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcTransactionList("", "", "", "", "", "", "", "", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcTransactionListResp{
		Data: make([]types.BtcTransactionListData, 0),
	}

	expectData := types.BtcTransactionListData{
		Page:             "1",
		Limit:            "20",
		TotalPage:        "500",
		TotalTransaction: "31124061",
		InscriptionsList: make([]types.GetBtcTransactionInscription, 0),
	}

	expectInscription := types.GetBtcTransactionInscription{
		TxId:              "af6bb18c64c57296ae07b8f4b1857c655160402a8d332fdb523915d7887476e2",
		BlockHeight:       "812873",
		State:             "success",
		TokenType:         "BRC20",
		ActionType:        "mint",
		FromAddress:       "",
		ToAddress:         "bc1qx2l6pzt7aknazsgdnvf5vnhp8ezx38ykg2wvfz",
		Amount:            "",
		Token:             "SoIa",
		InscriptionId:     "af6bb18c64c57296ae07b8f4b1857c655160402a8d332fdb523915d7887476e2i0",
		InscriptionNumber: "35346117",
		Index:             "0",
		Location:          "af6bb18c64c57296ae07b8f4b1857c655160402a8d332fdb523915d7887476e2",
		Msg:               "tick: $ORC has been minted",
		Time:              "1697695573000",
	}
	expectData.InscriptionsList = append(expectData.InscriptionsList, expectInscription)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	inscriptionsList := data.InscriptionsList[0]
	assert.Equal(t, inscriptionsList.Token, expectInscription.Token)
	assert.Equal(t, inscriptionsList.State, expectInscription.State)
}
