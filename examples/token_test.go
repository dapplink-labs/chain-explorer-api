package tests

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/token"
	"testing"
)

// go test -v examples/mock.go examples/token_test.go -run TestOklinkGetETHToken  token_test.go
func TestOklinkGetETHToken(t *testing.T) {
	oklinkClient, etherscanClient, err := NewMockClient()
	//_, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	// ETH
	contractAddress := "0x0e3a2a1f2146d86a604adc220b4967a898d7fe07"
	protocolType := ""
	page := "1"
	limit := "10"
	trps := &token.TokenRequest{
		ChainShortName:  "ETH",
		ExplorerName:    "oklink",
		ContractAddress: contractAddress,
		ProtocolType:    protocolType,
		Page:            page,
		Limit:           limit,
	}
	fmt.Println(trps, "trps trps trps")

	ethscanResp, err := etherscanClient.GetTokenList(trps)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ethscanResp, "ethscanResp")
	fmt.Println(err, "err")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("======================")
	//fmt.Println(okResp.BalanceStr)
	//fmt.Println(okResp.Account)
	//fmt.Println(okResp.Symbol)
	fmt.Println("======================")

	okResp, err := oklinkClient.GetTokenList(trps)
	fmt.Println(okResp, "okResp")
	fmt.Println(err, "err")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("======================")
	//fmt.Println(okResp.BalanceStr)
	//fmt.Println(okResp.Account)
	//fmt.Println(okResp.Symbol)
	fmt.Println("======================")
}
