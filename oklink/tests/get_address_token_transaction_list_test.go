//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTokenTransactionList(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressTokenTransactionList("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", "token_20", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := types.AddressTokenTransactionListResp{
		Data: make([]types.AddressTokenTransactionListData, 0),
	}

	expectData := types.AddressTokenTransactionListData{
		Limit:           "1",
		Page:            "1",
		TotalPage:       "10000",
		TransactionList: make([]types.AddressTokenTransactionListTransaction, 0),
	}

	expectTx := types.AddressTokenTransactionListTransaction{
		TxId:                 "0x66e4b648d6b82c5e2cfdd2121af896a11618c69a356c307e2403a885a8503c88",
		BlockHash:            "0x6199c61f711a797e7bc88b213a5ae71374898a413e5e20f9f8ad420189088e82",
		Height:               "18376245",
		TransactionTime:      "1697620043000",
		From:                 "0xdac17f958d2ee523a2206206994597c13d831ec7",
		To:                   "0xce7a837e1717301cb30d668ec6fcc9f4d312f30f",
		TokenContractAddress: "0xd8daa146dc3d7f2e5a7df7074164b6ad99bed260",
		TokenId:              "",
		Amount:               "1450000000",
		Symbol:               "",
		IsFromContract:       true,
		IsToContract:         false,
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	data := resp.Data[0]
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.From, expectTx.From)
}
