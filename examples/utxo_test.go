package tests

import (
	"fmt"
	"testing"

	"github.com/dapplink-labs/chain-explorer-api/common/account"
)

func TestGetAccountUtxo(t *testing.T) {
	oklinkClient, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	accountRequest := &account.AccountUtxoRequest{
		ExplorerName:   "Bitcoin",
		ChainShortName: "BTC",
		Address:        "bc1qffuugwxpyd72j5kat9q89hje8atmfetcmznfq0",
		Page:           "",
		Limit:          "",
	}
	accountRequestUtxo, err := oklinkClient.GetAccountUtxo(accountRequest)
	if err != nil {
		fmt.Println("get transaction by hash fail", "err", err)
	}
	fmt.Println("============Account Request Utxo===========")
	for _, value := range accountRequestUtxo {
		fmt.Println(value.TxId)
		fmt.Println(value.BlockTime)
		fmt.Println(value.Address)
		fmt.Println(value.UnspentAmount)
		fmt.Println(value.Index)
	}
	fmt.Println("============Account Request Utxo===========")
}
