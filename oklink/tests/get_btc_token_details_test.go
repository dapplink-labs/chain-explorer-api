//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcTokenDetails(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcTokenDetails("sats")
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcTokenDetailsResp{
		Data: make([]types.BtcTokenDetailsTokenData, 0),
	}

	expectData := types.BtcTokenDetailsTokenData{
		Token:             "sats",
		Precision:         "18",
		TotalSupply:       "2100000000000000",
		MintAmount:        "2100000000000000",
		LimitPerMint:      "100000000",
		Holder:            "35825",
		DeployAddress:     "bc1prtawdt82wfgrujx6d0heu0smxt4yykq440t447wan88csf3mc7csm3ulcn",
		LogoUrl:           "https://static.oklink.com/cdn/web3/currency/token/btc-sats-9b664bdd6f5ed80d8d88957b63364c41f3ad4efb8eee11366aa16435974d9333i0.png/type=png_350_0",
		TxId:              "9b664bdd6f5ed80d8d88957b63364c41f3ad4efb8eee11366aa16435974d9333",
		InscriptionId:     "9b664bdd6f5ed80d8d88957b63364c41f3ad4efb8eee11366aa16435974d9333i0",
		DeployHeight:      "779971",
		DeployTime:        "1678339934000",
		InscriptionNumber: "357097",
		State:             "success",
		TokenType:         "BRC20",
		Msg:               "",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Token, expectData.Token)
}
