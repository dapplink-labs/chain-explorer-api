package tests

import (
	"fmt"
	"testing"

	"github.com/dapplink-labs/chain-explorer-api/common/token"
)

func TestOklinkGetETHToken(t *testing.T) {
	oklinkClient, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	trps := &token.TokenRequest{
		ChainShortName:  "ETH",
		ExplorerName:    "oklink",
		ContractAddress: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
		ProtocolType:    "token_20",
		Page:            "1",
		Limit:           "10",
	}
	ethscanRespList, err := etherscanClient.GetTokenList(trps)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========etherscanClient============")
	for _, v := range ethscanRespList {
		fmt.Println(v.TokenId)
		fmt.Println(v.TokenContractAddress)
		fmt.Println(v.TotalSupply)
	}
	fmt.Println("===========etherscanClient===========")

	okRespList, err := oklinkClient.GetTokenList(trps)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("=========okRespList=============")
	for _, v := range okRespList {
		fmt.Println(v.TokenId)
		fmt.Println(v.TokenContractAddress)
		fmt.Println(v.TotalSupply)
	}
	fmt.Println("===========okRespList===========")
}

func TestOklinkGetNFTDetails(t *testing.T) {
	oklinkClient, _, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	trps := &token.GetNFTDetailsRequest{
		ExplorerName:    "oklink",
		ChainShortName:  "ETH",
		ContractAddress: "0xed5af388653567af2f388e6224dc7c4b3241c544",
		TokenId:         "5401",
	}
	okResp, err := oklinkClient.GetNFTDetails(trps)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("=========okRespList=============")
	fmt.Println(okResp)
	fmt.Println("===========okRespList===========")
}
