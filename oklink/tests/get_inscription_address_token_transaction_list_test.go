//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionAddressTokenTransactionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionAddressTokenTransactionList("btc", "bc1qvwqt8vtn2k7vrjqrsct63pkfw9ufqjldmjm439", "brc20", "", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionAddressTokenTransactionListResp{
		Data: make([]types.InscriptionAddressTokenTransactionListData, 0),
	}

	expectData := types.InscriptionAddressTokenTransactionListData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "81",
		ChainFullName:   "Bitcoin",
		ChainShortName:  "BTC",
		TotalTransfer:   "81",
		TransactionList: make([]types.InscriptionAddressTokenTransactionListTransaction, 0),
	}

	expectTransaction := types.InscriptionAddressTokenTransactionListTransaction{
		TxId:               "4cf9aefbc9febf80b68376fa773849aabfdd8e3f7a5254ad11fd7ec6c32d3e89",
		BlockHash:          "000000000000000000001765a54bc80e84b856d70a77884544839256b42e9a4e",
		Height:             "832015",
		TransactionTime:    "1708885500000",
		From:               "",
		To:                 "bc1qvwqt8vtn2k7vrjqrsct63pkfw9ufqjldmjm439",
		Amount:             "5000",
		Symbol:             "ANSM",
		Action:             "inscribeTransfer",
		TokenInscriptionId: "4f54d82160bf08bab83bbe89276b2fd9bed514ce843c91a796daa07bafb85239i0",
		ProtocolType:       "BRC20",
		State:              "success",
		InscriptionId:      "4cf9aefbc9febf80b68376fa773849aabfdd8e3f7a5254ad11fd7ec6c32d3e89i0",
		InscriptionNumber:  "62377691",
		OutputIndex:        "",
	}

	expectData.TransactionList = append(expectData.TransactionList, expectTransaction)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.ChainFullName, expectData.ChainFullName)
	assert.Equal(t, data.ChainShortName, expectData.ChainShortName)

	transaction := data.TransactionList[0]
	assert.Equal(t, transaction.State, expectTransaction.State)
	assert.Equal(t, transaction.Symbol, expectTransaction.Symbol)
}
