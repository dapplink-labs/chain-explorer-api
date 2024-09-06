//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionAddressInscriptionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionAddressInscriptionList("btc", "bc1qhuv3dhpnm0wktasd3v0kt6e4aqfqsd0uhfdu7d", "brc20", "", "", "", "", 2)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionAddressInscriptionListResp{
		Data: make([]types.InscriptionAddressInscriptionListData, 0),
	}

	expectData := types.InscriptionAddressInscriptionListData{
		Page:             "1",
		Limit:            "2",
		TotalPage:        "3154",
		TotalInscription: "6307",
		InscriptionList:  make([]types.InscriptionAddressInscriptionListInscription, 2),
	}

	expectData.InscriptionList[0] = types.InscriptionAddressInscriptionListInscription{
		InscriptionId:      "bca21f193e5f16a3fa1207df8021a5923539175e7bab92235c8a7e6ef9cf8db7i0",
		TokenInscriptionId: "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0",
		InscriptionNumber:  "62748365",
		Symbol:             "ordi",
		State:              "success",
		ProtocolType:       "BRC20",
		Action:             "inscribeTransfer",
	}

	expectData.InscriptionList[1] = types.InscriptionAddressInscriptionListInscription{
		InscriptionId:      "e5bc024b23d6e3a2cb499300b34234a625ff2c1d70fa43a28d50250efba5c7d1i0",
		TokenInscriptionId: "9b664bdd6f5ed80d8d88957b63364c41f3ad4efb8eee11366aa16435974d9333i0",
		InscriptionNumber:  "62748359",
		Symbol:             "sats",
		State:              "success",
		ProtocolType:       "BRC20",
		Action:             "inscribeTransfer",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.TotalInscription, expectData.TotalInscription)

	for i := 0; i < len(data.InscriptionList); i++ {
		inscription := data.InscriptionList[i]
		expectInscription := expectData.InscriptionList[i]
		assert.Equal(t, inscription.State, expectInscription.State)
	}
}
