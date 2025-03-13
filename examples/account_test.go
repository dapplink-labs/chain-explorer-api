package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dapplink-labs/chain-explorer-api/common/account"
)

func TestGetAccountBalance(t *testing.T) {
	oklinkClient, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	accountItem := []string{"0xD79053a14BC465d9C1434d4A4fAbdeA7b6a2A94b"}
	symbol := []string{"ETH"}
	contractAddress := []string{"0x00"}
	protocolType := []string{""}
	page := []string{"1"}
	limit := []string{"10"}
	acbr := &account.AccountBalanceRequest{
		ChainShortName:  "ETH",
		ExplorerName:    "etherescan",
		Account:         accountItem,
		Symbol:          symbol,
		ContractAddress: contractAddress,
		ProtocolType:    protocolType,
		Page:            page,
		Limit:           limit,
	}
	etherscanResp, err := etherscanClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========etherscanResp============")
	fmt.Println(etherscanResp.BalanceStr)
	fmt.Println(etherscanResp.Balance.Int())
	fmt.Println(etherscanResp.Account)
	fmt.Println(etherscanResp.Symbol)
	fmt.Println("===========etherscanResp===========")

	oklinkResp, err := oklinkClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========oklinkResp============")
	fmt.Println(oklinkResp.BalanceStr)
	fmt.Println(oklinkResp.Account)
	fmt.Println(oklinkResp.Symbol)
	fmt.Println("==========oklinkResp============")
}

func TestTokenGetAccountBalance(t *testing.T) {
	oklinkClient, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	accountItem := []string{"0x55FE002aefF02F77364de339a1292923A15844B8"}
	symbol := []string{"USDT"}
	contractAddress := []string{"0xdAC17F958D2ee523a2206206994597C13D831ec7"}
	protocolType := []string{"token_20"}
	page := []string{"1"}
	limit := []string{"10"}
	acbr := &account.AccountBalanceRequest{
		ChainShortName:  "ETH",
		ExplorerName:    "etherescan",
		Account:         accountItem,
		Symbol:          symbol,
		ContractAddress: contractAddress,
		ProtocolType:    protocolType,
		Page:            page,
		Limit:           limit,
	}

	etherscanResp, err := etherscanClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========etherscanResp============")
	fmt.Println(etherscanResp.Balance.Int())
	fmt.Println(etherscanResp.Account)
	fmt.Println(etherscanResp.Symbol)
	fmt.Println("===========etherscanResp==========")

	oklinkResp, err := oklinkClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========oklinkResp============")
	fmt.Println(oklinkResp.BalanceStr)
	fmt.Println(oklinkResp.Account)
	fmt.Println(oklinkResp.Symbol)
	fmt.Println("==========oklinkResp============")

}

func TestGetMultiAccountBalance(t *testing.T) {
	oklinkClient, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	accountItem := []string{"0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97", "0x4E5B2e1dc63F6b91cb6Cd759936495434C7e972F"}
	symbol := []string{"ETH", "ETH"}
	contractAddress := []string{"0x00", "0x00"}
	protocolType := []string{"", ""}
	page := []string{"1", "1"}
	limit := []string{"10", "10"}
	acbr := &account.AccountBalanceRequest{
		ChainShortName:  "ETH",
		ExplorerName:    "oklink",
		Account:         accountItem,
		Symbol:          symbol,
		ContractAddress: contractAddress,
		ProtocolType:    protocolType,
		Page:            page,
		Limit:           limit,
	}
	oklinkResp, err := oklinkClient.GetMultiAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========Multi oklinkResp============")
	for _, oklinkRespItem := range oklinkResp {
		fmt.Println(oklinkRespItem.BalanceStr)
		fmt.Println(oklinkRespItem.Symbol)
		fmt.Println(oklinkRespItem.Account)
	}
	fmt.Println("==========Multi oklinkResp============")

	etherscanResp, err := etherscanClient.GetMultiAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========Multi etherscanResp============")
	for _, etherscanRespItem := range etherscanResp {
		fmt.Println(etherscanRespItem.Balance.Int())
		fmt.Println(etherscanRespItem.Symbol)
		fmt.Println(etherscanRespItem.Account)
	}
	fmt.Println("==========Multi etherscanResp============")
}

func TestTokenGetMultiAccountBalance(t *testing.T) {
	oklinkClient, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	accountItem := []string{"0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97", "0x4E5B2e1dc63F6b91cb6Cd759936495434C7e972F"}
	symbol := []string{"USDT"}
	contractAddress := []string{"0xdAC17F958D2ee523a2206206994597C13D831ec7"}
	protocolType := []string{"token_20"}
	page := []string{"1"}
	limit := []string{"10"}
	acbr := &account.AccountBalanceRequest{
		ChainShortName:  "ETH",
		ExplorerName:    "oklink",
		Account:         accountItem,
		Symbol:          symbol,
		ContractAddress: contractAddress,
		ProtocolType:    protocolType,
		Page:            page,
		Limit:           limit,
	}
	oklinkResp, err := oklinkClient.GetMultiAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========Multi Token OklinkResp============")
	for _, oklinkRespItem := range oklinkResp {
		fmt.Println(oklinkRespItem.BalanceStr)
		fmt.Println(oklinkRespItem.Symbol)
		fmt.Println(oklinkRespItem.Account)
	}
	fmt.Println("==========Multi Token OklinkResp============")
}

func TestGetAccountBalanceV2(t *testing.T) {
	oklinkClient, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	// accountItem := []string{"0x28C6c06298d514Db089934071355E5743bf21d60", "0xF17ACEd3c7A8DAA29ebb90Db8D1b6efD8C364a18"}
	accountItem := []string{"0x4D392a08f8198a8E56c58BDaB6D01E99Ab33c086"}
	acbr := &account.GetAccountBalanceRequest{
		ChainShortName:  "ETH",
		ExplorerName:    "etherescan",
		Account:         accountItem,
		ContractAddress: "0xF39c410Dac956BA98004f411E182FB4EEd595270",
		TokenType:       "contract",
		ProtocolType:    "token_721",
		Page:            "1",
		Limit:           "10",
	}
	// accountItem := []string{"0xE0554a476A092703abdB3Ef35c80e0D76d32939F"}
	// acbr := &account.GetAccountBalanceRequest{
	// 	ChainShortName:  "ETH",
	// 	ExplorerName:    "etherescan",
	// 	Account:         accountItem,
	// 	ContractAddress: "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
	// 	TokenType:       "contract",
	// 	ProtocolType:    "token_20",
	// 	Page:            "1",
	// 	Limit:           "10",
	// }
	// accountItem := []string{"0x5eE84D30c7EE57F63f71c92247Ff31f95E26916B"}
	// acbr := &account.GetAccountBalanceRequest{
	// 	ChainShortName:  "ETH",
	// 	ExplorerName:    "etherescan",
	// 	Account:         accountItem,
	// 	ContractAddress: "",
	// 	TokenType:       "native",
	// 	ProtocolType:    "",
	// 	Page:            "1",
	// 	Limit:           "10",
	// }
	etherscanResp, err := etherscanClient.GetAccountBalanceV2(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========etherscanResp============")
	jsonBytes, _ := json.Marshal(etherscanResp)
	fmt.Printf("%s\n", string(jsonBytes))
	fmt.Println("===========etherscanResp===========")

	oklinkResp, err := oklinkClient.GetAccountBalanceV2(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========oklinkResp============")
	jsonBytes, _ = json.Marshal(oklinkResp)
	fmt.Printf("%s\n", string(jsonBytes))
	fmt.Println("==========oklinkResp============")
}
