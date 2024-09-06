//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressInternalTransactionListMulti(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetAddressInternalTransactionListMulti("eth", "0xd501520326d41aead2a70d4b5bf0c4646c0c9bd8,0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701", "18370000", "18374470", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressInternalTransactionListMultiResp{
		Data: make([]types.AddressInternalTransactionListMultiData, 0),
	}

	expectData := types.AddressInternalTransactionListMultiData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "7",
		TransactionList: make([]types.AddressInternalTransactionListMultiTransaction, 0),
	}

	expectTransaction := types.AddressInternalTransactionListMultiTransaction{
		TxId:            "0xfcc0b4d18791f5932ba7e3563012a176ef0d9f959e0beefc34f6956a0d0043b6",
		BlockHash:       "0x6eab9220138600d0cd66b73737aec77c016c18c91e13e93a938b7477e1e18865",
		Height:          "18373050",
		TransactionTime: "1697581427000",
		Operation:       "call",
		From:            "0x09fa0d3154363036ea406f254808c53f5f975518",
		To:              "0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701",
		IsFromContract:  true,
		IsToContract:    false,
		Amount:          "2450",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTransaction)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	expectTx := expectData.TransactionList[0]
	assert.Equal(t, tx.TxId, expectTx.TxId)
	assert.Equal(t, tx.From, expectTx.From)
}
