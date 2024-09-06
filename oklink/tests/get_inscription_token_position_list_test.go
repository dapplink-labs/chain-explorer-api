//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionTokenPositionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionTokenPositionList("btc", "brc20", "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0", "", "", "", "", 3)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionTokenPositionListResp{
		Data: make([]types.InscriptionTokenPositionListData, 0),
	}

	expectData := types.InscriptionTokenPositionListData{
		Page:         "1",
		Limit:        "3",
		TotalPage:    "3334",
		PositionList: make([]types.InscriptionTokenPositionListPosition, 3),
	}

	expectData.PositionList[0] = types.InscriptionTokenPositionListPosition{
		HolderAddress: "bc1qhuv3dhpnm0wktasd3v0kt6e4aqfqsd0uhfdu7d",
		Amount:        "8704276.68038247",
		Rank:          "1",
	}

	expectData.PositionList[1] = types.InscriptionTokenPositionListPosition{
		HolderAddress: "bc1qggf48ykykz996uv5vsp5p9m9zwetzq9run6s64hm6uqfn33nhq0ql9t85q",
		Amount:        "1675781.3938851",
		Rank:          "2",
	}

	expectData.PositionList[2] = types.InscriptionTokenPositionListPosition{
		HolderAddress: "bc1qm64dsdz853ntzwleqsrdt5p53w75zfrtnmyzcx",
		Amount:        "1121981.97971559",
		Rank:          "3",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	for i := 0; i < len(data.PositionList); i++ {
		// TODO: The data obtained for HolderAddress and Amount does not match the data provided in the documentation
		//assert.Equal(t, data.PositionList[i].HolderAddress, expectData.PositionList[i].HolderAddress)
		//assert.Equal(t, data.PositionList[i].Amount, expectData.PositionList[i].Amount)
		assert.Equal(t, data.PositionList[i].Rank, expectData.PositionList[i].Rank)
	}
}
