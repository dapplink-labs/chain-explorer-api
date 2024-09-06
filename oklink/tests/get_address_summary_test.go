//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressSummary(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressSummary("eth", "0x85c6627c4ed773cb7c32644b041f58a058b00d30")
	if err != nil {
		t.Error(err)
	}
	data := resp.Data[0]

	expectResp := &types.AddressSummaryResp{
		Data: make([]types.AddressSummaryData, 0),
	}

	expectData := types.AddressSummaryData{
		ChainFullName:  "Ethereum",
		ChainShortName: "ETH",
		Address:        "0x85c6627c4ed773cb7c32644b041f58a058b00d30",
		IsAaAddress:    false,
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	assert.Equal(t, data.ChainFullName, expectData.ChainFullName)
	assert.Equal(t, data.ChainShortName, expectData.ChainShortName)
	assert.Equal(t, data.Address, expectData.Address)
	assert.Equal(t, data.IsAaAddress, expectData.IsAaAddress)
}
