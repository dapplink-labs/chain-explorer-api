package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/explorer/etherscan"
	"github.com/dapplink-labs/chain-explorer-api/explorer/oklink"
)

var etherscanCli *etherscan.ChainExplorerAdaptor
var oklinkCli *oklink.ChainExplorerAdaptor

var (
	EtherscanBaseUrl = "https://api.etherscan.io/api?"
	EtherscanApiKey  = "HZEZGEPJJDA633N421AYW9NE8JFNZZC7JT"
	EtherscanTimeout = time.Second * 20

	OklinkBaseUrl = "https://www.oklink.com/"
	OklinkApiKey  = "5181d535-b68f-41cf-bbc6-25905e46b6a6"
	OkTimeout     = time.Second * 20
)

func TestOklinkGetAddressSummary(t *testing.T) {
	var err error
	oklinkCli, err = oklink.NewChainExplorerAdaptor(OklinkApiKey, OklinkBaseUrl, false, time.Duration(OkTimeout))
	if err != nil {
		fmt.Println("mock etherscan cli fail", "err", err)
	}
	accountItem := []string{"0x4838B106FCe9647Bdf1E7877BF73cE8B0BAD5f97"}
	symbol := []string{"ETH"}
	contractAddress := []string{"0x00"}
	protocolType := []string{""}
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
	resp, err := oklinkCli.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("======================")
	fmt.Println(resp.BalanceStr)
	fmt.Println(resp.Account)
	fmt.Println(resp.Symbol)
	fmt.Println("======================")
}

func TestEtherscanGetAddressSummary(t *testing.T) {
	var err error
	etherscanCli, err = etherscan.NewChainExplorerAdaptor(EtherscanApiKey, EtherscanBaseUrl, false, time.Duration(EtherscanTimeout))
	if err != nil {
		fmt.Println("mock etherscan cli fail", "err", err)
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
	resp, err := etherscanCli.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("======================")
	fmt.Println(resp.Balance.Int())
	fmt.Println(resp.Account)
	fmt.Println(resp.Symbol)
	fmt.Println("======================")
}

func TestOklinkTokenGetAddressSummary(t *testing.T) {
	var err error
	oklinkCli, err = oklink.NewChainExplorerAdaptor(OklinkApiKey, OklinkBaseUrl, false, time.Duration(OkTimeout))
	if err != nil {
		fmt.Println("mock etherscan cli fail", "err", err)
	}
	accountItem := []string{"0x55FE002aefF02F77364de339a1292923A15844B8"}
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
	resp, err := oklinkCli.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("======================")
	fmt.Println(resp.BalanceStr)
	fmt.Println(resp.Account)
	fmt.Println(resp.Symbol)
	fmt.Println("======================")
}

func TestEtherscanTokenGetAddressSummary(t *testing.T) {
	var err error
	etherscanCli, err = etherscan.NewChainExplorerAdaptor(EtherscanApiKey, EtherscanBaseUrl, false, time.Duration(EtherscanTimeout))
	if err != nil {
		fmt.Println("mock etherscan cli fail", "err", err)
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
	resp, err := etherscanCli.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("======================")
	fmt.Println(resp.Balance.Int())
	fmt.Println(resp.Account)
	fmt.Println(resp.Symbol)
	fmt.Println("======================")
}
