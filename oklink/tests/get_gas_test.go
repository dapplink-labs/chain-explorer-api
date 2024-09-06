//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

// TestGetGas Test client.GetGas() function
func TestGetGas(t *testing.T) {
	client := oklink.New()
	// This API does not support querying gas for multiple currencies in a single request,
	// for example, client.GetGas("BTC, ETH") will trigger a 403 error.
	resp, err := client.GetGas("BTC")
	if err != nil {
		t.Error(err)
	}
	data := resp.Data[0]

	expectResp := &types.GetGasResp{
		Data: make([]types.GetGasData, 0),
	}

	expectData := types.GetGasData{
		ChainFullName:  "Bitcoin",
		ChainShortName: "BTC",
		Symbol:         "BTC",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	// Compare the length of response
	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	// Compare some invariable fields of response
	assert.Equal(t, data.ChainFullName, expectData.ChainFullName)
	assert.Equal(t, data.ChainShortName, expectData.ChainShortName)
	assert.Equal(t, data.Symbol, expectData.Symbol)
}
