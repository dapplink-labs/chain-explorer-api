package tests

import (
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/explorer/etherscan"
	"os"
	"testing"

	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
)

func Test_GetTxByAddress_txList(t *testing.T) {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7890")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")

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
	ethscanResp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if ethscanResp.Limit > 0 {
		tx := ethscanResp.TransactionList[0]
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
	fmt.Println("Test_GetTxByAddress_txList")
	jsonByteArr, err := json.Marshal(ethscanResp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}

func Test_GetTxByAddress_txlistinternal(t *testing.T) {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7890")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")

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
	ethscanResp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if ethscanResp.Limit > 0 {
		tx := ethscanResp.TransactionList[0]
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
	fmt.Println("Test_GetTxByAddress_txList")
	jsonByteArr, err := json.Marshal(ethscanResp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}

func Test_GetTxByAddress_tokentx(t *testing.T) {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7890")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")

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
	ethscanResp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	if ethscanResp.Limit > 0 {
		tx := ethscanResp.TransactionList[0]
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
	fmt.Println("Test_GetTxByAddress_txList")
	jsonByteArr, err := json.Marshal(ethscanResp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}

func Test_GetTxByAddress_tokennfttx(t *testing.T) {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7890")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")

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
	ethscanResp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("======================")
	fmt.Println("Test_GetTxByAddress_txList")
	jsonByteArr, err := json.Marshal(ethscanResp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}

func Test_GetTxByAddress_token1155tx(t *testing.T) {
	os.Setenv("HTTP_PROXY", "http://127.0.0.1:7890")
	os.Setenv("HTTPS_PROXY", "http://127.0.0.1:7890")

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
	ethscanResp, err := etherscanClient.GetTxByAddress(request)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("======================")
	fmt.Println("Test_GetTxByAddress_txList")
	jsonByteArr, err := json.Marshal(ethscanResp)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(jsonByteArr))
	fmt.Println("======================")
}
