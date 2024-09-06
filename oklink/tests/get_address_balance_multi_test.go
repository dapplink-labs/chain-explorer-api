//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressBalanceMulti(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressBalanceMulti("eth", "0x85c6627c4ed773cb7c32644b041f58a058b00d30,0xb13a8883d5116b418066c379bc3b3f40d087b8d8")
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressBalanceMultiResp{
		Data: make([]types.AddressBalanceMultiData, 0),
	}

	expectData := types.AddressBalanceMultiData{
		Symbol:      "ETH",
		BalanceList: make([]types.AddressBalanceMultiBalance, 2),
	}

	expectData.BalanceList[0] = types.AddressBalanceMultiBalance{
		Address: "0x85c6627c4ed773cb7c32644b041f58a058b00d30",
		Balance: "0",
	}

	expectData.BalanceList[1] = types.AddressBalanceMultiBalance{
		Address: "0xb13a8883d5116b418066c379bc3b3f40d087b8d8",
		Balance: "0.00019330554147975",
	}

	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Symbol, expectData.Symbol)

	for i := 0; i < len(data.BalanceList); i++ {
		balance := data.BalanceList[i]
		expectBalance := expectData.BalanceList[i]
		assert.Equal(t, balance.Address, expectBalance.Address)
		assert.Equal(t, balance.Balance, expectBalance.Balance)
	}
}
