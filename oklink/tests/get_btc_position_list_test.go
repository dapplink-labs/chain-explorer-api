//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcPositionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcPositionList("sats", "", 2)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcPositionListResp{
		Data: make([]types.BtcPositionListData, 0),
	}

	expectData := types.BtcPositionListData{
		Page:         "1",
		Limit:        "2",
		TotalPage:    "5000",
		PositionList: make([]types.BtcPositionListPosition, 2),
	}

	expectData.PositionList[0] = types.BtcPositionListPosition{
		HolderAddress: "bc1plff0sqm6ym55eak9vjljghd55h7hkheg22c84w55w646l857elqsfzmfdv",
		Amount:        "31740686608926",
		Rank:          "1",
	}

	expectData.PositionList[1] = types.BtcPositionListPosition{
		HolderAddress: "bc1pun3whtlzac75f2vcznxmpfc09dnzyp0luw8tpfwc7wrruwav30pqsu6l9u",
		Amount:        "22651199774504",
		Rank:          "1",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)

	for i := 0; i < len(data.PositionList); i++ {
		position := data.PositionList[i]
		expectPosition := expectData.PositionList[i]

		assert.Equal(t, position.HolderAddress, expectPosition.HolderAddress)
		assert.Equal(t, position.Amount, expectPosition.Amount)
		assert.Equal(t, position.Rank, expectPosition.Rank)
	}
}
