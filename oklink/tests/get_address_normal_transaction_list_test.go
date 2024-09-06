//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressNormalTransactionList(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressNormalTransactionList("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressNormalTransactionListResp{
		Data: make([]types.AddressNormalTransactionListData, 0),
	}

	expectData := types.AddressNormalTransactionListData{
		Limit:           "1",
		Page:            "1",
		TotalPage:       "10000",
		TransactionList: make([]types.AddressNormalTransactionListTransaction, 0),
	}

	expectTx := types.AddressNormalTransactionListTransaction{
		TxId:            "0xac39ce204486c83fa1aef285456a7e0d76f4a76976ab5ab65bcea98d97ee8508",
		MethodId:        "0xa9059cbb",
		Nonce:           "0",
		GasPrice:        "8438956495",
		GasLimit:        "94813",
		GasUsed:         "63209",
		BlockHash:       "0x62a73bc006e481f6f6da53c3d71ea7a8f20c78de4b12a8eaa89d59d68501eefc",
		Height:          "18383240",
		TransactionTime: "1697704715000",
		From:            "0xf284512f225b350bf6e71d5a327891fcd26f640c",
		To:              "0xdac17f958d2ee523a2206206994597c13d831ec7",
		IsFromContract:  false,
		IsToContract:    true,
		Amount:          "0",
		Symbol:          "ETH",
		TxFee:           "0.000533418001092455",
		State:           "success",
		TransactionType: "2",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.State, expectTx.State)
}
