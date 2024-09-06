//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressUtxo(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetAddressUtxo("btc", "bc1ql49ydapnjafl5t2cp9zqpjwe6pdgmxy98859v2", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressUtxoResp{
		Data: make([]types.AddressUtxoData, 0),
	}

	expectData := types.AddressUtxoData{
		Page:      "1",
		Limit:     "1",
		TotalPage: "160",
		UtxoList:  make([]types.AddressUtxoUtxo, 0),
	}

	expectUtxo := types.AddressUtxoUtxo{
		TxId:          "d11638ea2cf68c4b49c1d97ef681a9e7e4658ba6cb7290dd73d476db371b9037",
		Height:        "796599",
		BlockTime:     "1688150365",
		Address:       "bc1ql49ydapnjafl5t2cp9zqpjwe6pdgmxy98859v2",
		UnspentAmount: "0.0003",
		Index:         "0",
	}

	expectData.UtxoList = append(expectData.UtxoList, expectUtxo)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	utxo := data.UtxoList[0]
	assert.Equal(t, utxo.Address, expectUtxo.Address)
}
