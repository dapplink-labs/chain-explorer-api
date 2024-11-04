package tests

import (
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/explorer/solscan"
	"testing"
	"time"

	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"github.com/dapplink-labs/chain-explorer-api/explorer/etherscan"
)

func Test_etherscan_GetTxByAddress_txList(t *testing.T) {
	_, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	request := &account.AccountTxRequest{
		ChainShortName:   "ETH",
		ExplorerName:     etherscan.ChainExplorerName,
		Action:           account.EtherscanActionTxList,
		Address:          "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		StartBlockHeight: 0,
		EndBlockHeight:   99999999,
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if len(resp.TransactionList) > 0 {
		tx := resp.TransactionList[0]
		fmt.Println("txId", tx.TxId)
		fmt.Println("BlockHash", tx.BlockHash)
		fmt.Println("Height", tx.Height)
		fmt.Println("From", tx.From)
		fmt.Println("To", tx.To)
		fmt.Println("Amount", tx.Amount)
		fmt.Println("Nonce", tx.Nonce)
		fmt.Println("GasPrice", tx.GasPrice)
		fmt.Println("GasLimit", tx.GasLimit)
		fmt.Println("GasUsed", tx.GasUsed)
		fmt.Println("State", tx.State)
		fmt.Println("IsFromContract", tx.IsFromContract)
		fmt.Println("IsToContract", tx.IsToContract)
	}

	fmt.Println("======================")
	fmt.Println("Test_GetTxByAddress")
	jsonByteArr, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}

func Test_etherscan_GetTxByAddress_txlistinternal(t *testing.T) {
	_, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName:   "ETH",
		ExplorerName:     etherscan.ChainExplorerName,
		Action:           account.EtherscanActionTxListInternal,
		Address:          "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		StartBlockHeight: 0,
		EndBlockHeight:   99999999,
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if len(resp.TransactionList) > 0 {
		tx := resp.TransactionList[0]
		fmt.Println("txId", tx.TxId)
		fmt.Println("BlockHash", tx.BlockHash)
		fmt.Println("Height", tx.Height)
		fmt.Println("From", tx.From)
		fmt.Println("To", tx.To)
		fmt.Println("Amount", tx.Amount)
		fmt.Println("Nonce", tx.Nonce)
		fmt.Println("GasPrice", tx.GasPrice)
		fmt.Println("GasLimit", tx.GasLimit)
		fmt.Println("GasUsed", tx.GasUsed)
		fmt.Println("State", tx.State)
		fmt.Println("IsFromContract", tx.IsFromContract)
		fmt.Println("IsToContract", tx.IsToContract)
	}
	fmt.Println("======================")
	fmt.Println("Test_GetTxByAddress")
	jsonByteArr, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}

func Test_etherscan_GetTxByAddress_tokentx(t *testing.T) {

	_, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName:   "ETH",
		ExplorerName:     etherscan.ChainExplorerName,
		Action:           account.EtherscanActionTokenTx,
		Address:          "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		StartBlockHeight: 0,
		EndBlockHeight:   99999999,
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if len(resp.TransactionList) > 0 {
		tx := resp.TransactionList[0]
		fmt.Println("txId", tx.TxId)
		fmt.Println("BlockHash", tx.BlockHash)
		fmt.Println("Height", tx.Height)
		fmt.Println("From", tx.From)
		fmt.Println("To", tx.To)
		fmt.Println("TokenContractAddress", tx.TokenContractAddress)
		fmt.Println("Amount", tx.Amount)
		fmt.Println("Symbol", tx.Symbol)
		fmt.Println("Nonce", tx.Nonce)
		fmt.Println("GasPrice", tx.GasPrice)
		fmt.Println("GasLimit", tx.GasLimit)
		fmt.Println("GasUsed", tx.GasUsed)
		fmt.Println("State", tx.State)
		fmt.Println("IsFromContract", tx.IsFromContract)
		fmt.Println("IsToContract", tx.IsToContract)
	}
	fmt.Println("======================")
	fmt.Println("Test_GetTxByAddress")
	jsonByteArr, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}

func Test_etherscan_GetTxByAddress_tokennfttx(t *testing.T) {

	_, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName:   "ETH",
		ExplorerName:     etherscan.ChainExplorerName,
		Action:           account.EtherscanActionTokenNftTx,
		Address:          "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		StartBlockHeight: 0,
		EndBlockHeight:   99999999,
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if len(resp.TransactionList) > 0 {
		tx := resp.TransactionList[0]
		fmt.Println("txId", tx.TxId)
		fmt.Println("BlockHash", tx.BlockHash)
		fmt.Println("Height", tx.Height)
		fmt.Println("From", tx.From)
		fmt.Println("To", tx.To)
		fmt.Println("TokenContractAddress", tx.TokenContractAddress)
		fmt.Println("Amount", tx.Amount)
		fmt.Println("Symbol", tx.Symbol)
		fmt.Println("Nonce", tx.Nonce)
		fmt.Println("GasPrice", tx.GasPrice)
		fmt.Println("GasLimit", tx.GasLimit)
		fmt.Println("GasUsed", tx.GasUsed)
		fmt.Println("State", tx.State)
		fmt.Println("IsFromContract", tx.IsFromContract)
		fmt.Println("IsToContract", tx.IsToContract)
	}
	fmt.Println("======================")
	fmt.Println("Test_GetTxByAddress")
	jsonByteArr, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}

func Test_etherscan_GetTxByAddress_token1155tx(t *testing.T) {

	_, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName:   "ETH",
		ExplorerName:     etherscan.ChainExplorerName,
		Action:           account.EtherscanActionToken1155Tx,
		Address:          "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		StartBlockHeight: 0,
		EndBlockHeight:   99999999,
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if len(resp.TransactionList) > 0 {
		tx := resp.TransactionList[0]
		fmt.Println("txId", tx.TxId)
		fmt.Println("BlockHash", tx.BlockHash)
		fmt.Println("Height", tx.Height)
		fmt.Println("From", tx.From)
		fmt.Println("To", tx.To)
		fmt.Println("TokenContractAddress", tx.TokenContractAddress)
		fmt.Println("Amount", tx.Amount)
		fmt.Println("Symbol", tx.Symbol)
		fmt.Println("Nonce", tx.Nonce)
		fmt.Println("GasPrice", tx.GasPrice)
		fmt.Println("GasLimit", tx.GasLimit)
		fmt.Println("GasUsed", tx.GasUsed)
		fmt.Println("State", tx.State)
		fmt.Println("IsFromContract", tx.IsFromContract)
		fmt.Println("IsToContract", tx.IsToContract)
	}
	fmt.Println("======================")
	fmt.Println("Test_GetTxByAddress")
	jsonByteArr, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}

func Test_solscan_GetTxByAddress(t *testing.T) {
	sol, err := solscan.NewChainExplorerAdaptor("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjcmVhdGVkQXQiOjE3MjQwNjIxMzk5MDYsImVtYWlsIjoiemFja2d1by5ndW9AZ21haWwuY29tIiwiYWN0aW9uIjoidG9rZW4tYXBpIiwiYXBpVmVyc2lvbiI6InYxIiwiaWF0IjoxNzI0MDYyMTM5fQ.EaWDC25lyGNx_LqRL5sAYYKLMbq10brnexKnAz9C3UY", "https://pro-api.solscan.io", true, time.Second*20)
	if err == nil {
		request := account.AccountTxRequest{
			PageRequest: chain.PageRequest{
				Page:  1,
				Limit: 50,
			},
			Address: "4Y19AR3cQ76UmLPeEYsvwkUXaiS8GEfivyggcYSmL4M6",
		}
		resp, _ := sol.GetTxByAddress(&request)
		fmt.Println(resp)
	} else {
		fmt.Println(err)
	}
}
