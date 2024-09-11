package tests

import (
	"fmt"
	"testing"

	"github.com/dapplink-labs/chain-explorer-api/common/account"
)

func TestGetAccountBalance(t *testing.T) {
	oklinkClient, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	accountItem := []string{"0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"}
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
	ethscanResp, err := etherscanClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("======================")
	fmt.Println(ethscanResp.Balance.Int())
	fmt.Println(ethscanResp.Account)
	fmt.Println(ethscanResp.Symbol)
	fmt.Println("======================")

	okResp, err := oklinkClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("======================")
	fmt.Println(okResp.BalanceStr)
	fmt.Println(okResp.Account)
	fmt.Println(okResp.Symbol)
	fmt.Println("======================")
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
	fmt.Println("======================")
	fmt.Println(etherscanResp.Balance.Int())
	fmt.Println(etherscanResp.Account)
	fmt.Println(etherscanResp.Symbol)
	fmt.Println("======================")

	oklinkResp, err := oklinkClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("======================")
	fmt.Println(oklinkResp.BalanceStr)
	fmt.Println(oklinkResp.Account)
	fmt.Println(oklinkResp.Symbol)
	fmt.Println("======================")

}

//func TestGetMultiAccountBalance(t *testing.T) {
//	oklinkClient, etherscanClient, err := NewMockClient()
//	if err != nil {
//		fmt.Println("new mock client fail", "err", err)
//	}
//	accountItem := []string{"0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97", "0x4E5B2e1dc63F6b91cb6Cd759936495434C7e972F"}
//	symbol := []string{"ETH", "ETH"}
//	contractAddress := []string{"0x00", "0x00"}
//	protocolType := []string{"", ""}
//	page := []string{"1", "1"}
//	limit := []string{"10", "10"}
//	acbr := &account.AccountBalanceRequest{
//		ChainShortName:  "ETH",
//		ExplorerName:    "oklink",
//		Account:         accountItem,
//		Symbol:          symbol,
//		ContractAddress: contractAddress,
//		ProtocolType:    protocolType,
//		Page:            page,
//		Limit:           limit,
//	}
//}

//func TestTokenGetMultiAccountBalance(t *testing.T) {
//	oklinkClient, etherscanClient, err := NewMockClient()
//	if err != nil {
//		fmt.Println("new mock client fail", "err", err)
//	}
//	accountItem := []string{"0x55FE002aefF02F77364de339a1292923A15844B8"}
//	symbol := []string{"USDT"}
//	contractAddress := []string{"0xdAC17F958D2ee523a2206206994597C13D831ec7"}
//	protocolType := []string{"token_20"}
//	page := []string{"1"}
//	limit := []string{"10"}
//	acbr := &account.AccountBalanceRequest{
//		ChainShortName:  "ETH",
//		ExplorerName:    "oklink",
//		Account:         accountItem,
//		Symbol:          symbol,
//		ContractAddress: contractAddress,
//		ProtocolType:    protocolType,
//		Page:            page,
//		Limit:           limit,
//	}
//	oklinkResp, err := oklinkClient.GetMultiAccountBalance(acbr)
//	if err != nil {
//		t.Error(err)
//	}
//
//	etherscanResp, err := etherscanClient.GetMultiAccountBalance(acbr)
//	if err != nil {
//		t.Error(err)
//	}
//
//}
