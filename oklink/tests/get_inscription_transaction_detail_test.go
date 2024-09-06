//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionTransactionDetail(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionTransactionDetail("btc", "c29fc5f33756c572fc55152435d9314059f8639797708b39471330536b94ed0c", "brc20", "", 0)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionTransactionDetailResp{
		Data: make([]types.InscriptionTransactionDetailData, 0),
	}

	expectData := types.InscriptionTransactionDetailData{
		Page:            "1",
		Limit:           "20",
		TotalPage:       "1",
		TransactionList: make([]types.InscriptionTransactionDetailTransaction, 0),
	}

	expectTx := types.InscriptionTransactionDetailTransaction{
		TxId:               "c29fc5f33756c572fc55152435d9314059f8639797708b39471330536b94ed0c",
		BlockHash:          "0000000000000000000159cea4e78229ccffab8ecfa94354729ee2b0c52b7a3f",
		Height:             "831823",
		TransactionTime:    "1708779611000",
		From:               "",
		To:                 "bc1qhuv3dhpnm0wktasd3v0kt6e4aqfqsd0uhfdu7d",
		Amount:             "2198220440",
		Action:             "inscribeTransfer",
		TokenInscriptionId: "9b664bdd6f5ed80d8d88957b63364c41f3ad4efb8eee11366aa16435974d9333i0",
		ProtocolType:       "BRC20",
		State:              "success",
		InscriptionId:      "c29fc5f33756c572fc55152435d9314059f8639797708b39471330536b94ed0ci0",
		InscriptionNumber:  "62184839",
		Symbol:             "sats",
		OutputIndex:        "",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.State, expectTx.State)
	assert.Equal(t, tx.Symbol, expectTx.Symbol)
}
