package tests

import (
	"fmt"
	"testing"

	"github.com/dapplink-labs/chain-explorer-api/common/transaction"
)

func TestOklinkGetTxByHash(t *testing.T) {
	oklinkClient, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	transactionRequest := &transaction.TxRequest{
		ExplorerName:   "Bitcoin",
		ChainShortName: "BTC",
		Txid:           "f80b960fc894dc0cc311bee3a34e529d966a65fb3691efadd160716252bd26f1",
	}

	txRes, err := oklinkClient.GetTxByHash(transactionRequest)
	if err != nil {
		fmt.Println("get tx by hash fail", "err", err)
	}

	fmt.Println("============txRes===========")
	fmt.Println(txRes.Txid)
	fmt.Println(txRes.Txfee)
	fmt.Println(txRes.InputDetails)
	fmt.Println(txRes.OutputDetails)
	fmt.Println("============txRes===========")

}
