//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressInternalTransactionList(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressInternalTransactionList("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressInternalTransactionListResp{
		Data: make([]types.AddressInternalTransactionListData, 0),
	}

	expectData := types.AddressInternalTransactionListData{
		Limit:           "1",
		Page:            "1",
		TotalPage:       "10000",
		TransactionList: make([]types.AddressInternalTransactionListTransaction, 0),
	}

	expectTx := types.AddressInternalTransactionListTransaction{
		TxId:            "0x34d5bd0c44da0864cfb8b9d7f3311d5eb598a4093b26dd5df5d25ec0e9df4942",
		Operation:       "call",
		BlockHash:       "0xee4e80ebc8a4b8071b07abd63906a4201bcf76d66100369a39148a0f529d098c",
		Height:          "18376673",
		TransactionTime: "1697625227000",
		From:            "0x3a5cc8689d1b0cef2c317bc5c0ad6ce88b27d597",
		To:              "0xdac17f958d2ee523a2206206994597c13d831ec7",
		IsFromContract:  true,
		IsToContract:    true,
		Amount:          "0",
		Symbol:          "ETH",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.Operation, expectTx.Operation)
}
