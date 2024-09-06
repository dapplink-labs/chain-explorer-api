//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionTokenTransactionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionTokenTransactionList("btc", "brc20", "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0", "1", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionTokenTransactionListResp{
		Data: make([]types.InscriptionTokenTransactionListData, 0),
	}

	expectData := types.InscriptionTokenTransactionListData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "10000",
		ChainFullName:   "Bitcoin",
		ChainShortName:  "BTC",
		TotalTransfer:   "337393",
		TransactionList: make([]types.InscriptionTokenTransactionListTransaction, 0),
	}

	expectTransaction := types.InscriptionTokenTransactionListTransaction{
		TxId:               "0cf990907ef51eb0607f7fc6bb81809137d5750f4b64de9a8fc7917770bd1ad5",
		BlockHash:          "00000000000000000000256813d252f532a57f5473f2e723d8c7483a7df4d42f",
		Height:             "832486",
		TransactionTime:    "1709177741000",
		From:               "",
		To:                 "bc1pqjwg8673seyk0t8jtz9j9g78uddps3cppd6nmnvjpp42sn22fqkqy8h700",
		Amount:             "141",
		Symbol:             "ordi",
		Action:             "inscribeTransfer",
		TokenInscriptionId: "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0",
		ProtocolType:       "BRC20",
		State:              "success",
		InscriptionId:      "0cf990907ef51eb0607f7fc6bb81809137d5750f4b64de9a8fc7917770bd1ad5i0",
		InscriptionNumber:  "62753536",
		OutputIndex:        "",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTransaction)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.ChainFullName, expectData.ChainFullName)
	assert.Equal(t, data.ChainShortName, expectData.ChainShortName)
	assert.Equal(t, data.TotalTransfer, expectData.TotalTransfer)

	transaction := data.TransactionList[0]
	assert.Equal(t, transaction.State, expectTransaction.State)
}
