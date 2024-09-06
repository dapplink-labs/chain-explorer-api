//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTransactionList(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressTransactionList("eth", "0x85c6627c4ed773cb7c32644b041f58a058b00d30", "", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := types.AddressTransactionListResp{
		Data: make([]types.AddressTransactionListData, 0),
	}

	expectData := types.AddressTransactionListData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "3",
		ChainFullName:   "Ethereum",
		ChainShortName:  "ETH",
		TransactionList: make([]types.AddressTransaction, 1),
	}

	expectTransaction := types.AddressTransaction{
		TxId:                 "0x18714d659c9022eecd29bff3cd05cb78adc6c0b9522b04d713bfb2cc7a2f62f0",
		MethodId:             "",
		BlockHash:            "0xea0ee963034d87aeaccd6a0513725bec2a604b6a890e85d96289bfcd547154db",
		Height:               "16361564",
		TransactionTime:      "1673174795000",
		From:                 "0x85c6627c4ed773cb7c32644b041f58a058b00d30",
		To:                   "0xcffad3200574698b78f32232aa9d63eabd290703",
		IsFromContract:       false,
		IsToContract:         false,
		Amount:               "0.000567475798167289",
		TransactionSymbol:    "ETH",
		TxFee:                "0.000430211335176",
		State:                "success",
		TokenId:              "",
		TokenContractAddress: "",
		ChallengeStatus:      "",
		L1OriginHash:         "",
	}

	expectData.TransactionList[0] = expectTransaction
	expectResp.Data = append(expectResp.Data, expectData)

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.ChainFullName, expectData.ChainFullName)
	assert.Equal(t, data.ChainShortName, expectData.ChainShortName)

	tx := data.TransactionList[0]
	expectTx := expectData.TransactionList[0]
	assert.Equal(t, tx, expectTx)
}
