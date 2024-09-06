//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressNormalTransactionListMulti(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetAddressNormalTransactionListMulti("eth", "00x533a7ae90fee4cafbc00e6a551cfb39a954cbf48,0xc0a3465b50a47848b7d04e145df61565d3e10566", "18374341", "18374343", "", "", 0)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressNormalTransactionListMultiResp{
		Data: make([]types.AddressNormalTransactionListMultiData, 0),
	}

	expectData := types.AddressNormalTransactionListMultiData{
		Page:            "1",
		Limit:           "20",
		TotalPage:       "1",
		TransactionList: make([]types.AddressNormalTransactionListMultiTransaction, 0),
	}

	expectTransaction := types.AddressNormalTransactionListMultiTransaction{
		TxId:            "0x76e650150abadac6c781c9c90a0fcda69fce6e69f0fbbfb08d8cadc26076802a",
		MethodId:        "",
		BlockHash:       "0x58c6a91b0b5ff04bb7ea9b9f264c547472a96dafbdb069acc1e2e8d63792db16",
		Height:          "18374343",
		TransactionTime: "1697597087000",
		From:            "0x533a7ae90fee4cafbc00e6a551cfb39a954cbf48",
		To:              "0xc0a3465b50a47848b7d04e145df61565d3e10566",
		IsFromContract:  false,
		IsToContract:    false,
		Amount:          "15296",
		Symbol:          "ETH",
		TxFee:           "0.000139810264818",
		GasLimit:        "21000",
		GasUsed:         "21000",
		GasPrice:        "",
		Nonce:           "6657631658",
		TransactionType: "2",
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

	assert.Equal(t, tx.From, expectTx.From)
}
