//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTokenTransactionListMulti(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetAddressTokenTransactionListMulti("eth", "0xd501520326d41aead2a70d4b5bf0c4646c0c9bd8,0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701", "18370000", "18374470", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressTokenTransactionListMultiResp{
		Data: make([]types.AddressTokenTransactionListMultiData, 0),
	}

	expectData := types.AddressTokenTransactionListMultiData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "88",
		TransactionList: make([]types.AddressTokenTransactionListMultiTransaction, 0),
	}

	expectTransaction := types.AddressTokenTransactionListMultiTransaction{
		TxId:                 "0x644fe2fbc53316c3146760ecdb1405fc2dcb4893ac19552ad0898ea669176300",
		BlockHash:            "0xd027f203664d2911cb2bf2f73134539e6f7b5ac20be6ca998b7ea3691acfcd3d",
		Height:               "18373112",
		TransactionTime:      "1697582183000",
		From:                 "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
		To:                   "0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701",
		IsFromContract:       true,
		IsToContract:         false,
		Amount:               "1",
		TokenId:              "1",
		Symbol:               "Airdrop",
		TokenContractAddress: "0xf7b24c63fe941f2caadaee49f10e565d850be067",
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
