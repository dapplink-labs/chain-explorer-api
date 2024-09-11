package tests

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
	"testing"
)

func TestGasFee(t *testing.T) {
	oklinkClient, etherscanClient, err := NewMockClient()
	if err != nil {
		fmt.Println("new mock client fail", "err", err)
	}
	gasLimit := 2000000000
	acbr := &gas_fee.GasEstimateFeeRequest{
		ChainShortName: "ETH",
		ExplorerName:   "etherescan",
		GasLimit:       gasLimit,
	}
	etherscanResp, err := etherscanClient.GetEstimateGasFee(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========etherscanResp============")
	fmt.Println(etherscanResp)

	fmt.Println(etherscanResp.BaseFee)
	fmt.Println(etherscanResp.StandardGasPrice)
	fmt.Println(etherscanResp.SlowGasPrice)
	fmt.Println("===========etherscanResp===========")

	oklinkResp, err := oklinkClient.GetEstimateGasFee(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========oklinkResp============")
	fmt.Println(oklinkResp.BaseFee)
	fmt.Println(oklinkResp.StandardGasPrice)
	fmt.Println(oklinkResp.SlowGasPrice)
	fmt.Println("==========oklinkResp============")
}
