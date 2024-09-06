//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionBlockTokenTransaction(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionBlockTokenTransaction("btc", "831823", "brc20", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionBlockTokenTransactionResp{
		Data: make([]types.InscriptionBlockTokenTransactionData, 0),
	}

	expectData := types.InscriptionBlockTokenTransactionData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "559",
		TotalTransfer:   "559",
		TransactionList: make([]types.InscriptionBlockTokenTransactionTransaction, 0),
	}

	expectTx := types.InscriptionBlockTokenTransactionTransaction{
		TxId:               "5ac740cad1c29266bf3615fc4f108c082431c7c0be74944e352edd75eed471ff",
		BlockHash:          "0000000000000000000159cea4e78229ccffab8ecfa94354729ee2b0c52b7a3f",
		Height:             "831823",
		From:               "",
		To:                 "bc1pnad2fk3fw6q3d20mhyacdekl7wf96rpg5yqxhtchzvpwet989lpsmvuc9n",
		Amount:             "1000",
		Action:             "mint",
		TokenInscriptionId: "4865f05b9132f12bb09d6215f13da5a304a502a95315d0a49463d6f8c0bb7740i0",
		ProtocolType:       "BRC20",
		State:              "success",
		InscriptionId:      "5ac740cad1c29266bf3615fc4f108c082431c7c0be74944e352edd75eed471ffi0",
		InscriptionNumber:  "62196529",
		Symbol:             "LABS",
		TransactionTime:    "1708779611000",
		OutputIndex:        "",
	}

	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.TotalTransfer, expectData.TotalTransfer)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.State, expectTx.State)
	assert.Equal(t, tx.Symbol, expectTx.Symbol)
}
