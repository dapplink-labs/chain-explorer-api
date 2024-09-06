//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcInscriptionsList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcInscriptionsList("", "", "", "", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcInscriptionsListResp{
		Data: make([]types.BtcInscriptionsListData, 0),
	}

	expectData := types.BtcInscriptionsListData{
		Page:             "1",
		Limit:            "20",
		TotalPage:        "500",
		TotalInscription: "60109663",
		InscriptionsList: make([]types.BtcInscriptionsListInscription, 0),
	}

	expectInscription := types.BtcInscriptionsListInscription{
		InscriptionId:     "92780ef845a5190a1027724c24b5adbe6713abdaa43c5f273ff8a87d41f6cc8ci0",
		InscriptionNumber: "9999999",
		Location:          "92780ef845a5190a1027724c24b5adbe6713abdaa43c5f273ff8a87d41f6cc8c",
		Token:             "10MM",
		State:             "success",
		Msg:               "",
		TokenType:         "BRC20",
		ActionType:        "mint",
		LogoUrl:           "",
		OwnerAddress:      "bc1p53rrfs7l3fdyd60wzsk9gphnjm6y9hr4vhfrp207tsltyxjatfqqj9ly8k",
		TxId:              "92780ef845a5190a1027724c24b5adbe6713abdaa43c5f273ff8a87d41f6cc8c",
		BlockHeight:       "792013",
		ContentSize:       "",
		Time:              "1685401405000",
	}
	expectData.InscriptionsList = append(expectData.InscriptionsList, expectInscription)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.TotalInscription, expectData.TotalInscription)

	inscription := data.InscriptionsList[0]
	assert.Equal(t, inscription.TokenType, expectInscription.TokenType)
}
