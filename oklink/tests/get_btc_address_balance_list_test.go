//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcAddressBalanceList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcAddressBalanceList("bc1ph0057nc25ka94z8ydg43j8tnnp38u3hxpadutnt4n3jyfrmjzmcqw99mk2", "", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcAddressBalanceListResp{
		Data: make([]types.BtcAddressBalanceListData, 0),
	}

	expectData := types.BtcAddressBalanceListData{
		Page:        "1",
		Limit:       "20",
		TotalPage:   "1",
		BalanceList: make([]types.BtcAddressBalanceListBalance, 0),
	}

	expectBalance := types.BtcAddressBalanceListBalance{
		Token:            "X@AI",
		TokenType:        "BRC20",
		Balance:          "1350000000000",
		AvailableBalance: "1350000000000",
		TransferBalance:  "0",
	}
	expectData.BalanceList = append(expectData.BalanceList, expectBalance)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	balance := data.BalanceList[0]
	assert.Equal(t, balance.Token, expectBalance.Token)
	assert.Equal(t, balance.TokenType, expectBalance.TokenType)
}
