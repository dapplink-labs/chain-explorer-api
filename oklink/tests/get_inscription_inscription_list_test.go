//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionInscriptionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionInscriptionList("btc", "brc20", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionInscriptionListResp{
		Data: make([]types.InscriptionInscriptionListData, 0),
	}

	expectData := types.InscriptionInscriptionListData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "10000",
		InscriptionList: make([]types.InscriptionInscriptionListInscription, 0),
	}

	expectInscription := types.InscriptionInscriptionListInscription{
		InscriptionId:      "03fbeb834260fed03f87f0a09e06339379efc5fd3d6d08cc0a87451e509c32f1i0",
		InscriptionNumber:  "62752328",
		TokenInscriptionId: "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0",
		Symbol:             "ordi",
		State:              "success",
		ProtocolType:       "BRC20",
		Action:             "inscribeTransfer",
		OwnerAddress:       "bc1q6fh6ll49efsjrgcdwh7hp3cswt8faw85agghkk",
	}

	expectData.InscriptionList = append(expectData.InscriptionList, expectInscription)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	inscription := data.InscriptionList[0]
	assert.Equal(t, inscription.ProtocolType, expectInscription.ProtocolType)
}
