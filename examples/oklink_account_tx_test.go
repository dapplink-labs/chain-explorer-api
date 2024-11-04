package tests

import (
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"github.com/dapplink-labs/chain-explorer-api/explorer/oklink"
	"testing"
)

func Test_oklink_ETH_GetTxByAddress_normal(t *testing.T) {
	oklinkCli, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName: "ETH",
		ExplorerName:   oklink.ChainExplorerName,
		Action:         account.OkLinkActionNormal,
		Address:        "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		IsFromOrTo:     string(account.From),
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := oklinkCli.GetTxByAddress(request)
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

func Test_oklink_GetTxByAddress_internal(t *testing.T) {
	oklinkCli, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName: "ETH",
		ExplorerName:   oklink.ChainExplorerName,
		Action:         account.OkLinkActionInternal,
		Address:        "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		IsFromOrTo:     string(account.From),
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := oklinkCli.GetTxByAddress(request)
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

func Test_oklink_ETH__GetTxByAddress_token_20(t *testing.T) {
	oklinkCli, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName: "ETH",
		ExplorerName:   oklink.ChainExplorerName,
		Action:         account.OkLinkActionToken,
		ProtocolType:   account.ProtocolTypeToken20,
		Address:        "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		IsFromOrTo:     string(account.From),
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := oklinkCli.GetTxByAddress(request)
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

func Test_oklink_ETH__GetTxByAddress_token_721(t *testing.T) {
	oklinkCli, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName: "ETH",
		ExplorerName:   oklink.ChainExplorerName,
		Action:         account.OkLinkActionToken,
		ProtocolType:   account.ProtocolTypeToken721,
		Address:        "0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97",
		IsFromOrTo:     string(account.From),
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := oklinkCli.GetTxByAddress(request)
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

func Test_oklink_ETH__GetTxByAddress_token_1155(t *testing.T) {
	oklinkCli, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName: "ETH",
		ExplorerName:   oklink.ChainExplorerName,
		Action:         account.OkLinkActionToken,
		ProtocolType:   account.ProtocolTypeToken1155,
		Address:        "0xe29Fa394Eef7d92e92A419c351B58Be24e907c7A",
		IsFromOrTo:     string(account.From),
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := oklinkCli.GetTxByAddress(request)
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

func Test_oklink_ETH_GetTxByAddress_tron(t *testing.T) {
	oklinkCli, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName: "Tron",
		ExplorerName:   oklink.ChainExplorerName,
		Action:         account.OkLinkActionNormal,
		Address:        "THfTJMWpcM6tmQsevN8zeCys5iohHWSQtF",
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 10,
		},
	}
	resp, err := oklinkCli.GetTxByAddress(request)
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

func Test_oklink_BTC_GetTxByAddress_normal(t *testing.T) {
	oklinkCli, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}

	request := &account.AccountTxRequest{
		ChainShortName: "BTC",
		ExplorerName:   oklink.ChainExplorerName,
		Action:         account.OkLinkActionUtxo,
		Address:        "bc1qcx6evk3my3cmepwpyv022fwg4gu6nwte3fcfak",
		//IsFromOrTo:     string(account.From),
		PageRequest: chain.PageRequest{
			Page:  1,
			Limit: 20,
		},
	}
	resp, err := oklinkCli.GetTxByAddress(request)
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
