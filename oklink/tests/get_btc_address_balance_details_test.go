//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcAddressBalanceDetails(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcAddressBalanceDetails("bc1ph0057nc25ka94z8ydg43j8tnnp38u3hxpadutnt4n3jyfrmjzmcqw99mk2", "meme", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcAddressBalanceDetailsResp{
		Data: make([]types.BtcAddressBalanceDetailsData, 0),
	}

	expectData := types.BtcAddressBalanceDetailsData{
		Page:                "1",
		Limit:               "20",
		TotalPage:           "1",
		Token:               "meme",
		TokenType:           "BRC20",
		Balance:             "18",
		AvailableBalance:    "",
		TransferBalance:     "18",
		TransferBalanceList: make([]types.BtcAddressBalanceDetailsTransferBalance, 0),
	}

	expectBalance := types.BtcAddressBalanceDetailsTransferBalance{
		InscriptionId:     "a1002519472f9a1d45db5a3df30ea521ecd5425e546a63a79f3a4a9ff4e6e582i0",
		InscriptionNumber: "4615101",
		Amount:            "3",
	}

	expectData.TransferBalanceList = append(expectData.TransferBalanceList, expectBalance)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
}
