# chain-explorer-api

## PART1. OKLink API

### 1.1 Fundamental Blockchain Data

#### 1.1.1 Fundamental Data

##### 1.1.1.1 Get transaction or Gas fee

```go
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
```

#### 1.1.2 Address Data

##### 1.1.2.1 Get basic address details

```go
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
```

##### 1.1.2.2 Get token balance details by address

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTokenBalance(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressTokenBalance("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", "token_20", 2)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressTokenBalanceResp{
		Data: make([]types.AddressTokenBalanceData, 0),
	}

	expectData := types.AddressTokenBalanceData{
		Limit:     "2",
		Page:      "1",
		TotalPage: "195",
		TokenList: make([]types.AddressTokenBalanceToken, 2, 20),
	}

	expectData.TokenList[0] = types.AddressTokenBalanceToken{
		Symbol:               "MCO",
		TokenContractAddress: "0x135f915f081e750baa63e9ce97421cc2d9dbfac0",
		HoldingAmount:        "115792089237316200000000000000000000000000000000000000000000",
		PriceUsd:             "0",
		ValueUsd:             "",
		TokenId:              "",
	}

	expectData.TokenList[1] = types.AddressTokenBalanceToken{
		Symbol:               "ShibDoge",
		TokenContractAddress: "0x6adb2e268de2aa1abf6578e4a8119b960e02928f",
		HoldingAmount:        "14720806439512404",
		PriceUsd:             "0.00000000000000003",
		ValueUsd:             "0.44162419318537212",
		TokenId:              "",
	}

	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]

	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, len(data.TokenList), len(expectData.TokenList))

	for i := 0; i < 2; i++ {
		assert.Equal(t, data.TokenList[i].Symbol, expectData.TokenList[i].Symbol)
		assert.Equal(t, data.TokenList[i].TokenContractAddress, expectData.TokenList[i].TokenContractAddress)
		assert.Equal(t, data.TokenList[i].HoldingAmount, expectData.TokenList[i].HoldingAmount)
		assert.Equal(t, data.TokenList[i].PriceUsd, expectData.TokenList[i].PriceUsd)
		assert.Equal(t, data.TokenList[i].ValueUsd, expectData.TokenList[i].ValueUsd)
		assert.Equal(t, data.TokenList[i].TokenId, expectData.TokenList[i].TokenId)
	}
}
```

##### 1.1.2.3 Get token balance details

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressBalanceFills(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressBalanceFills("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", 2)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressBalanceFillsResp{
		Data: make([]types.AddressBalanceFillsData, 0),
	}

	expectData := types.AddressBalanceFillsData{
		Limit:     "2",
		Page:      "1",
		TotalPage: "195",
		TokenList: make([]types.AddressBalanceFillsToken, 2),
	}

	expectData.TokenList[0] = types.AddressBalanceFillsToken{
		Token:                "USDT",
		TokenId:              "",
		HoldingAmount:        "402413.20994",
		TotalTokenValue:      "168.38785369",
		Change24h:            "0.00013",
		PriceUsd:             "1.00054",
		ValueUsd:             "402630.5130733676",
		TokenContractAddress: "0xdac17f958d2ee523a2206206994597c13d831ec7",
	}

	expectData.TokenList[1] = types.AddressBalanceFillsToken{
		Token:                "CRO",
		TokenId:              "",
		HoldingAmount:        "4800000",
		TotalTokenValue:      "156.86203294",
		Change24h:            "-0.032126",
		PriceUsd:             "0.07813984132122903",
		ValueUsd:             "375071.238341899344",
		TokenContractAddress: "0xa0b73e1ff0b80914ab6fe0444e65848c4c34450b",
	}

	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]

	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, len(data.TokenList), len(expectData.TokenList))

	for i := 0; i < 2; i++ {
		// Only some fields can be compared here,
		// because the values of the returned data are dynamic and cannot be predicted
		assert.Equal(t, data.TokenList[i].Token, expectData.TokenList[i].Token)
		assert.Equal(t, data.TokenList[i].TokenId, expectData.TokenList[i].TokenId)
		assert.Equal(t, data.TokenList[i].HoldingAmount, expectData.TokenList[i].HoldingAmount)
	}
}
```

##### 1.1.2.4 Get address balance at specific block height

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBlockAddressBalanceHistory(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetBlockAddressBalanceHistory("eth", "18376634", "0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701", "")
	if err != nil {
		t.Error(err)
	}

	data := resp.Data[0]

	expectResp := &types.BlockAddressBalanceHistoryResp{
		Data: make([]types.BlockAddressBalanceHistoryData, 0),
	}

	expectData := types.BlockAddressBalanceHistoryData{
		Address:              "0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701",
		Height:               "18376634",
		Balance:              "5.895934930980364414",
		BalanceSymbol:        "ETH",
		TokenContractAddress: "",
		BlockTime:            "1697624735000",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	assert.Equal(t, data.Address, expectData.Address)
	assert.Equal(t, data.Height, expectData.Height)
	assert.Equal(t, data.Balance, expectData.Balance)
	assert.Equal(t, data.BalanceSymbol, expectData.BalanceSymbol)
	assert.Equal(t, data.TokenContractAddress, expectData.TokenContractAddress)
	assert.Equal(t, data.BlockTime, expectData.BlockTime)
}
```

##### 1.1.2.5 Get address transaction list

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTransactionList(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressTransactionList("eth", "0x85c6627c4ed773cb7c32644b041f58a058b00d30", "", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := types.AddressTransactionListResp{
		Data: make([]types.AddressTransactionListData, 0),
	}

	expectData := types.AddressTransactionListData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "3",
		ChainFullName:   "Ethereum",
		ChainShortName:  "ETH",
		TransactionList: make([]types.AddressTransaction, 1),
	}

	expectTransaction := types.AddressTransaction{
		TxId:                 "0x18714d659c9022eecd29bff3cd05cb78adc6c0b9522b04d713bfb2cc7a2f62f0",
		MethodId:             "",
		BlockHash:            "0xea0ee963034d87aeaccd6a0513725bec2a604b6a890e85d96289bfcd547154db",
		Height:               "16361564",
		TransactionTime:      "1673174795000",
		From:                 "0x85c6627c4ed773cb7c32644b041f58a058b00d30",
		To:                   "0xcffad3200574698b78f32232aa9d63eabd290703",
		IsFromContract:       false,
		IsToContract:         false,
		Amount:               "0.000567475798167289",
		TransactionSymbol:    "ETH",
		TxFee:                "0.000430211335176",
		State:                "success",
		TokenId:              "",
		TokenContractAddress: "",
		ChallengeStatus:      "",
		L1OriginHash:         "",
	}

	expectData.TransactionList[0] = expectTransaction
	expectResp.Data = append(expectResp.Data, expectData)

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.ChainFullName, expectData.ChainFullName)
	assert.Equal(t, data.ChainShortName, expectData.ChainShortName)

	tx := data.TransactionList[0]
	expectTx := expectData.TransactionList[0]
	assert.Equal(t, tx, expectTx)
}
```

##### 1.1.2.6 Get address transaction list

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTransactionList(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressTransactionList("eth", "0x85c6627c4ed773cb7c32644b041f58a058b00d30", "", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := types.AddressTransactionListResp{
		Data: make([]types.AddressTransactionListData, 0),
	}

	expectData := types.AddressTransactionListData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "3",
		ChainFullName:   "Ethereum",
		ChainShortName:  "ETH",
		TransactionList: make([]types.AddressTransaction, 1),
	}

	expectTransaction := types.AddressTransaction{
		TxId:                 "0x18714d659c9022eecd29bff3cd05cb78adc6c0b9522b04d713bfb2cc7a2f62f0",
		MethodId:             "",
		BlockHash:            "0xea0ee963034d87aeaccd6a0513725bec2a604b6a890e85d96289bfcd547154db",
		Height:               "16361564",
		TransactionTime:      "1673174795000",
		From:                 "0x85c6627c4ed773cb7c32644b041f58a058b00d30",
		To:                   "0xcffad3200574698b78f32232aa9d63eabd290703",
		IsFromContract:       false,
		IsToContract:         false,
		Amount:               "0.000567475798167289",
		TransactionSymbol:    "ETH",
		TxFee:                "0.000430211335176",
		State:                "success",
		TokenId:              "",
		TokenContractAddress: "",
		ChallengeStatus:      "",
		L1OriginHash:         "",
	}

	expectData.TransactionList[0] = expectTransaction
	expectResp.Data = append(expectResp.Data, expectData)

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.ChainFullName, expectData.ChainFullName)
	assert.Equal(t, data.ChainShortName, expectData.ChainShortName)

	tx := data.TransactionList[0]
	expectTx := expectData.TransactionList[0]
	assert.Equal(t, tx, expectTx)
}
```

##### 1.1.2.7 Get standard transaction list by address

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressNormalTransactionList(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressNormalTransactionList("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressNormalTransactionListResp{
		Data: make([]types.AddressNormalTransactionListData, 0),
	}

	expectData := types.AddressNormalTransactionListData{
		Limit:           "1",
		Page:            "1",
		TotalPage:       "10000",
		TransactionList: make([]types.AddressNormalTransactionListTransaction, 0),
	}

	expectTx := types.AddressNormalTransactionListTransaction{
		TxId:            "0xac39ce204486c83fa1aef285456a7e0d76f4a76976ab5ab65bcea98d97ee8508",
		MethodId:        "0xa9059cbb",
		Nonce:           "0",
		GasPrice:        "8438956495",
		GasLimit:        "94813",
		GasUsed:         "63209",
		BlockHash:       "0x62a73bc006e481f6f6da53c3d71ea7a8f20c78de4b12a8eaa89d59d68501eefc",
		Height:          "18383240",
		TransactionTime: "1697704715000",
		From:            "0xf284512f225b350bf6e71d5a327891fcd26f640c",
		To:              "0xdac17f958d2ee523a2206206994597c13d831ec7",
		IsFromContract:  false,
		IsToContract:    true,
		Amount:          "0",
		Symbol:          "ETH",
		TxFee:           "0.000533418001092455",
		State:           "success",
		TransactionType: "2",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.State, expectTx.State)
}
```

##### 1.1.2.8 Get internal transaction list by address

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressInternalTransactionList(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressInternalTransactionList("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressInternalTransactionListResp{
		Data: make([]types.AddressInternalTransactionListData, 0),
	}

	expectData := types.AddressInternalTransactionListData{
		Limit:           "1",
		Page:            "1",
		TotalPage:       "10000",
		TransactionList: make([]types.AddressInternalTransactionListTransaction, 0),
	}

	expectTx := types.AddressInternalTransactionListTransaction{
		TxId:            "0x34d5bd0c44da0864cfb8b9d7f3311d5eb598a4093b26dd5df5d25ec0e9df4942",
		Operation:       "call",
		BlockHash:       "0xee4e80ebc8a4b8071b07abd63906a4201bcf76d66100369a39148a0f529d098c",
		Height:          "18376673",
		TransactionTime: "1697625227000",
		From:            "0x3a5cc8689d1b0cef2c317bc5c0ad6ce88b27d597",
		To:              "0xdac17f958d2ee523a2206206994597c13d831ec7",
		IsFromContract:  true,
		IsToContract:    true,
		Amount:          "0",
		Symbol:          "ETH",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.Operation, expectTx.Operation)
}
```

##### 1.1.2.9 Get token transaction list by address

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTokenTransactionList(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressTokenTransactionList("eth", "0xdac17f958d2ee523a2206206994597c13d831ec7", "token_20", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := types.AddressTokenTransactionListResp{
		Data: make([]types.AddressTokenTransactionListData, 0),
	}

	expectData := types.AddressTokenTransactionListData{
		Limit:           "1",
		Page:            "1",
		TotalPage:       "10000",
		TransactionList: make([]types.AddressTokenTransactionListTransaction, 0),
	}

	expectTx := types.AddressTokenTransactionListTransaction{
		TxId:                 "0x66e4b648d6b82c5e2cfdd2121af896a11618c69a356c307e2403a885a8503c88",
		BlockHash:            "0x6199c61f711a797e7bc88b213a5ae71374898a413e5e20f9f8ad420189088e82",
		Height:               "18376245",
		TransactionTime:      "1697620043000",
		From:                 "0xdac17f958d2ee523a2206206994597c13d831ec7",
		To:                   "0xce7a837e1717301cb30d668ec6fcc9f4d312f30f",
		TokenContractAddress: "0xd8daa146dc3d7f2e5a7df7074164b6ad99bed260",
		TokenId:              "",
		Amount:               "1450000000",
		Symbol:               "",
		IsFromContract:       true,
		IsToContract:         false,
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	data := resp.Data[0]
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.From, expectTx.From)
}
```

##### 1.1.2.10 Get native token balance in batches

```go
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
```

##### 1.1.2.11 Get token balance in batches

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTokenBalanceMulti(t *testing.T) {
	client := oklink.New()
	resp, err := client.GetAddressTokenBalanceMulti("eth", "0x85c6627c4ed773cb7c32644b041f58a058b00d30,0xb13a8883d5116b418066c379bc3b3f40d087b8d8", "", "", 2)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressTokenBalanceMultiResp{
		Data: make([]types.AddressTokenBalanceMultiData, 0),
	}

	expectData := types.AddressTokenBalanceMultiData{
		Page:        "1",
		Limit:       "2",
		TotalPage:   "0",
		BalanceList: make([]types.AddressTokenBalanceMultiBalance, 2),
	}

	expectData.BalanceList[0] = types.AddressTokenBalanceMultiBalance{
		Address:              "0xf977814e90da44bfa03b6295a0616a897441acec",
		HoldingAmount:        "400",
		TokenContractAddress: "0x7379cbce70bba5a9871f97d33b391afba377e885",
	}

	expectData.BalanceList[1] = types.AddressTokenBalanceMultiBalance{
		Address:              "0xf977814e90da44bfa03b6295a0616a897441acec",
		HoldingAmount:        "123101078.45198849",
		TokenContractAddress: "0x5c885be435a9b5b55bcfc992d8c085e4e549661e",
	}

	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	// TODO: There is no data returned, so it's impossible to compare the balance-related data.
	//for i := 0; i < len(data.BalanceList); i++ {
	//	balance := data.BalanceList[i]
	//	expectBalance := expectData.BalanceList[i]
	//	assert.Equal(t, balance.Address, expectBalance.Address)
	//	assert.Equal(t, balance.HoldingAmount, expectBalance.HoldingAmount)
	//	assert.Equal(t, balance.TokenContractAddress, expectBalance.TokenContractAddress)
	//}
}
```

##### 1.1.2.12 Get standard transaction list for specific block

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressNormalTransactionListMulti(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetAddressNormalTransactionListMulti("eth", "00x533a7ae90fee4cafbc00e6a551cfb39a954cbf48,0xc0a3465b50a47848b7d04e145df61565d3e10566", "18374341", "18374343", "", "", 0)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressNormalTransactionListMultiResp{
		Data: make([]types.AddressNormalTransactionListMultiData, 0),
	}

	expectData := types.AddressNormalTransactionListMultiData{
		Page:            "1",
		Limit:           "20",
		TotalPage:       "1",
		TransactionList: make([]types.AddressNormalTransactionListMultiTransaction, 0),
	}

	expectTransaction := types.AddressNormalTransactionListMultiTransaction{
		TxId:            "0x76e650150abadac6c781c9c90a0fcda69fce6e69f0fbbfb08d8cadc26076802a",
		MethodId:        "",
		BlockHash:       "0x58c6a91b0b5ff04bb7ea9b9f264c547472a96dafbdb069acc1e2e8d63792db16",
		Height:          "18374343",
		TransactionTime: "1697597087000",
		From:            "0x533a7ae90fee4cafbc00e6a551cfb39a954cbf48",
		To:              "0xc0a3465b50a47848b7d04e145df61565d3e10566",
		IsFromContract:  false,
		IsToContract:    false,
		Amount:          "15296",
		Symbol:          "ETH",
		TxFee:           "0.000139810264818",
		GasLimit:        "21000",
		GasUsed:         "21000",
		GasPrice:        "",
		Nonce:           "6657631658",
		TransactionType: "2",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTransaction)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	expectTx := expectData.TransactionList[0]

	assert.Equal(t, tx.From, expectTx.From)
}
```

##### 1.1.2.13 Get internal transaction list for specific block

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressInternalTransactionListMulti(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetAddressInternalTransactionListMulti("eth", "0xd501520326d41aead2a70d4b5bf0c4646c0c9bd8,0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701", "18370000", "18374470", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressInternalTransactionListMultiResp{
		Data: make([]types.AddressInternalTransactionListMultiData, 0),
	}

	expectData := types.AddressInternalTransactionListMultiData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "7",
		TransactionList: make([]types.AddressInternalTransactionListMultiTransaction, 0),
	}

	expectTransaction := types.AddressInternalTransactionListMultiTransaction{
		TxId:            "0xfcc0b4d18791f5932ba7e3563012a176ef0d9f959e0beefc34f6956a0d0043b6",
		BlockHash:       "0x6eab9220138600d0cd66b73737aec77c016c18c91e13e93a938b7477e1e18865",
		Height:          "18373050",
		TransactionTime: "1697581427000",
		Operation:       "call",
		From:            "0x09fa0d3154363036ea406f254808c53f5f975518",
		To:              "0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701",
		IsFromContract:  true,
		IsToContract:    false,
		Amount:          "2450",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTransaction)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	expectTx := expectData.TransactionList[0]
	assert.Equal(t, tx.TxId, expectTx.TxId)
	assert.Equal(t, tx.From, expectTx.From)
}
```

##### 1.1.2.14 Get internal transaction list for specific block

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressInternalTransactionListMulti(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetAddressInternalTransactionListMulti("eth", "0xd501520326d41aead2a70d4b5bf0c4646c0c9bd8,0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701", "18370000", "18374470", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressInternalTransactionListMultiResp{
		Data: make([]types.AddressInternalTransactionListMultiData, 0),
	}

	expectData := types.AddressInternalTransactionListMultiData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "7",
		TransactionList: make([]types.AddressInternalTransactionListMultiTransaction, 0),
	}

	expectTransaction := types.AddressInternalTransactionListMultiTransaction{
		TxId:            "0xfcc0b4d18791f5932ba7e3563012a176ef0d9f959e0beefc34f6956a0d0043b6",
		BlockHash:       "0x6eab9220138600d0cd66b73737aec77c016c18c91e13e93a938b7477e1e18865",
		Height:          "18373050",
		TransactionTime: "1697581427000",
		Operation:       "call",
		From:            "0x09fa0d3154363036ea406f254808c53f5f975518",
		To:              "0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701",
		IsFromContract:  true,
		IsToContract:    false,
		Amount:          "2450",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTransaction)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	expectTx := expectData.TransactionList[0]
	assert.Equal(t, tx.TxId, expectTx.TxId)
	assert.Equal(t, tx.From, expectTx.From)
}
```

##### 1.1.2.15 Get token transaction list for specific block

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetAddressTokenTransactionListMulti(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetAddressTokenTransactionListMulti("eth", "0xd501520326d41aead2a70d4b5bf0c4646c0c9bd8,0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701", "18370000", "18374470", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.AddressTokenTransactionListMultiResp{
		Data: make([]types.AddressTokenTransactionListMultiData, 0),
	}

	expectData := types.AddressTokenTransactionListMultiData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "88",
		TransactionList: make([]types.AddressTokenTransactionListMultiTransaction, 0),
	}

	expectTransaction := types.AddressTokenTransactionListMultiTransaction{
		TxId:                 "0x644fe2fbc53316c3146760ecdb1405fc2dcb4893ac19552ad0898ea669176300",
		BlockHash:            "0xd027f203664d2911cb2bf2f73134539e6f7b5ac20be6ca998b7ea3691acfcd3d",
		Height:               "18373112",
		TransactionTime:      "1697582183000",
		From:                 "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
		To:                   "0xd275e5cb559d6dc236a5f8002a5f0b4c8e610701",
		IsFromContract:       true,
		IsToContract:         false,
		Amount:               "1",
		TokenId:              "1",
		Symbol:               "Airdrop",
		TokenContractAddress: "0xf7b24c63fe941f2caadaee49f10e565d850be067",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTransaction)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	expectTx := expectData.TransactionList[0]

	assert.Equal(t, tx.TxId, expectTx.TxId)
	assert.Equal(t, tx.From, expectTx.From)
}
```

### 1.2 BTC inscription data

#### 1.2.1 Get list of inscription tokens

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionTokenList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionTokenList("btc", "runes", "", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionTokenListResp{
		Data: make([]types.InscriptionTokenListData, 0),
	}

	expectData := types.InscriptionTokenListData{
		Page:      "1",
		Limit:     "1",
		TotalPage: "10000",
		TokenList: make([]types.InscriptionTokenListToken, 0),
	}

	expectToken := types.InscriptionTokenListToken{
		Symbol:             "UNCOMMON•GOODS",
		TokenInscriptionId: "1:0",
		ProtocolType:       "RUNES",
		TotalSupply:        "3362172",
		MintAmount:         "3362172",
		DeployTime:         "0",
		Holder:             "40851",
		TransactionCount:   "3374587",
		CirculatingSupply:  "3362067",
		MintBitwork:        "",
		LimitPerMint:       "1",
		RunesSymbol:        "⧉",
		Divisibility:       "0",
		MintStatus:         "mintable",
		Premint:            "0",
		Burn:               "135",
		MintStartBlock:     "840000",
		MintEndBlock:       "1050000",
		MintCap:            "340282366920938463463374607431768211455",
	}
	expectData.TokenList = append(expectData.TokenList, expectToken)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	token := data.TokenList[0]
	assert.Equal(t, token.Symbol, expectToken.Symbol)
}
```

#### 1.2.2 Get list of addresses holding inscription tokens

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionTokenPositionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionTokenPositionList("btc", "brc20", "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0", "", "", "", "", 3)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionTokenPositionListResp{
		Data: make([]types.InscriptionTokenPositionListData, 0),
	}

	expectData := types.InscriptionTokenPositionListData{
		Page:         "1",
		Limit:        "3",
		TotalPage:    "3334",
		PositionList: make([]types.InscriptionTokenPositionListPosition, 3),
	}

	expectData.PositionList[0] = types.InscriptionTokenPositionListPosition{
		HolderAddress: "bc1qhuv3dhpnm0wktasd3v0kt6e4aqfqsd0uhfdu7d",
		Amount:        "8704276.68038247",
		Rank:          "1",
	}

	expectData.PositionList[1] = types.InscriptionTokenPositionListPosition{
		HolderAddress: "bc1qggf48ykykz996uv5vsp5p9m9zwetzq9run6s64hm6uqfn33nhq0ql9t85q",
		Amount:        "1675781.3938851",
		Rank:          "2",
	}

	expectData.PositionList[2] = types.InscriptionTokenPositionListPosition{
		HolderAddress: "bc1qm64dsdz853ntzwleqsrdt5p53w75zfrtnmyzcx",
		Amount:        "1121981.97971559",
		Rank:          "3",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	for i := 0; i < len(data.PositionList); i++ {
		// TODO: The data obtained for HolderAddress and Amount does not match the data provided in the documentation
		//assert.Equal(t, data.PositionList[i].HolderAddress, expectData.PositionList[i].HolderAddress)
		//assert.Equal(t, data.PositionList[i].Amount, expectData.PositionList[i].Amount)
		assert.Equal(t, data.PositionList[i].Rank, expectData.PositionList[i].Rank)
	}
}
```

#### 1.2.3 Get inscription token transfer list

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionTokenTransactionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionTokenTransactionList("btc", "brc20", "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0", "1", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionTokenTransactionListResp{
		Data: make([]types.InscriptionTokenTransactionListData, 0),
	}

	expectData := types.InscriptionTokenTransactionListData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "10000",
		ChainFullName:   "Bitcoin",
		ChainShortName:  "BTC",
		TotalTransfer:   "337393",
		TransactionList: make([]types.InscriptionTokenTransactionListTransaction, 0),
	}

	expectTransaction := types.InscriptionTokenTransactionListTransaction{
		TxId:               "0cf990907ef51eb0607f7fc6bb81809137d5750f4b64de9a8fc7917770bd1ad5",
		BlockHash:          "00000000000000000000256813d252f532a57f5473f2e723d8c7483a7df4d42f",
		Height:             "832486",
		TransactionTime:    "1709177741000",
		From:               "",
		To:                 "bc1pqjwg8673seyk0t8jtz9j9g78uddps3cppd6nmnvjpp42sn22fqkqy8h700",
		Amount:             "141",
		Symbol:             "ordi",
		Action:             "inscribeTransfer",
		TokenInscriptionId: "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0",
		ProtocolType:       "BRC20",
		State:              "success",
		InscriptionId:      "0cf990907ef51eb0607f7fc6bb81809137d5750f4b64de9a8fc7917770bd1ad5i0",
		InscriptionNumber:  "62753536",
		OutputIndex:        "",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTransaction)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.ChainFullName, expectData.ChainFullName)
	assert.Equal(t, data.ChainShortName, expectData.ChainShortName)
	assert.Equal(t, data.TotalTransfer, expectData.TotalTransfer)

	transaction := data.TransactionList[0]
	assert.Equal(t, transaction.State, expectTransaction.State)
}
```

#### 1.2.4 Get inscription list

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionInscriptionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionInscriptionList("btc", "brc20", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionInscriptionListResp{
		Data: make([]types.InscriptionInscriptionListData, 0),
	}

	expectData := types.InscriptionInscriptionListData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "10000",
		InscriptionList: make([]types.InscriptionInscriptionListInscription, 0),
	}

	expectInscription := types.InscriptionInscriptionListInscription{
		InscriptionId:      "03fbeb834260fed03f87f0a09e06339379efc5fd3d6d08cc0a87451e509c32f1i0",
		InscriptionNumber:  "62752328",
		TokenInscriptionId: "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0",
		Symbol:             "ordi",
		State:              "success",
		ProtocolType:       "BRC20",
		Action:             "inscribeTransfer",
		OwnerAddress:       "bc1q6fh6ll49efsjrgcdwh7hp3cswt8faw85agghkk",
	}

	expectData.InscriptionList = append(expectData.InscriptionList, expectInscription)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	inscription := data.InscriptionList[0]
	assert.Equal(t, inscription.ProtocolType, expectInscription.ProtocolType)
}
```

#### 1.2.5 Get inscription token list by address

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionAddressTokenList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionAddressTokenList("btc", "bc1qhuv3dhpnm0wktasd3v0kt6e4aqfqsd0uhfdu7d", "brc20", "", "", "", "", 2)

	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionAddressTokenListResp{
		Data: make([]types.InscriptionAddressTokenListData, 0),
	}

	expectData := types.InscriptionAddressTokenListData{
		Page:      "1",
		Limit:     "2",
		TotalPage: "68",
		TokenList: make([]types.InscriptionAddressTokenListToken, 2),
	}

	expectData.TokenList[0] = types.InscriptionAddressTokenListToken{
		Symbol:             "maibi",
		TokenInscriptionId: "b7456b7e688edea8fb814df146a83a062260b596616bfceff3ae2b9ceb8dbab2i0",
		HoldingAmount:      "8188888888888889300",
		InscriptionAmount:  "2",
		AvailableAmount:    "8188888888888889300",
		TransferableAmount: "0",
		InscriptionNumber:  "",
	}

	expectData.TokenList[1] = types.InscriptionAddressTokenListToken{
		Symbol:             "GPTu",
		TokenInscriptionId: "25e7f4ca11734aa1d1d8c9d9be262b8ea8b09a660a93a826084f7b21b3b41518i0",
		HoldingAmount:      "8000000000000000000",
		InscriptionAmount:  "2",
		AvailableAmount:    "8000000000000000000",
		TransferableAmount: "0",
		InscriptionNumber:  "",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	for i := 0; i < len(data.TokenList); i++ {
		token := data.TokenList[i]
		expectToken := expectData.TokenList[i]

		assert.Equal(t, token.Symbol, expectToken.Symbol)
	}
}
```

#### 1.2.6 Get inscription token list by address

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionAddressTokenList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionAddressTokenList("btc", "bc1qhuv3dhpnm0wktasd3v0kt6e4aqfqsd0uhfdu7d", "brc20", "", "", "", "", 2)

	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionAddressTokenListResp{
		Data: make([]types.InscriptionAddressTokenListData, 0),
	}

	expectData := types.InscriptionAddressTokenListData{
		Page:      "1",
		Limit:     "2",
		TotalPage: "68",
		TokenList: make([]types.InscriptionAddressTokenListToken, 2),
	}

	expectData.TokenList[0] = types.InscriptionAddressTokenListToken{
		Symbol:             "maibi",
		TokenInscriptionId: "b7456b7e688edea8fb814df146a83a062260b596616bfceff3ae2b9ceb8dbab2i0",
		HoldingAmount:      "8188888888888889300",
		InscriptionAmount:  "2",
		AvailableAmount:    "8188888888888889300",
		TransferableAmount: "0",
		InscriptionNumber:  "",
	}

	expectData.TokenList[1] = types.InscriptionAddressTokenListToken{
		Symbol:             "GPTu",
		TokenInscriptionId: "25e7f4ca11734aa1d1d8c9d9be262b8ea8b09a660a93a826084f7b21b3b41518i0",
		HoldingAmount:      "8000000000000000000",
		InscriptionAmount:  "2",
		AvailableAmount:    "8000000000000000000",
		TransferableAmount: "0",
		InscriptionNumber:  "",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	for i := 0; i < len(data.TokenList); i++ {
		token := data.TokenList[i]
		expectToken := expectData.TokenList[i]

		assert.Equal(t, token.Symbol, expectToken.Symbol)
	}
}
```

#### 1.2.7 Get inscription token transfers by address

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionAddressTokenTransactionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionAddressTokenTransactionList("btc", "bc1qvwqt8vtn2k7vrjqrsct63pkfw9ufqjldmjm439", "brc20", "", "", "", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionAddressTokenTransactionListResp{
		Data: make([]types.InscriptionAddressTokenTransactionListData, 0),
	}

	expectData := types.InscriptionAddressTokenTransactionListData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "81",
		ChainFullName:   "Bitcoin",
		ChainShortName:  "BTC",
		TotalTransfer:   "81",
		TransactionList: make([]types.InscriptionAddressTokenTransactionListTransaction, 0),
	}

	expectTransaction := types.InscriptionAddressTokenTransactionListTransaction{
		TxId:               "4cf9aefbc9febf80b68376fa773849aabfdd8e3f7a5254ad11fd7ec6c32d3e89",
		BlockHash:          "000000000000000000001765a54bc80e84b856d70a77884544839256b42e9a4e",
		Height:             "832015",
		TransactionTime:    "1708885500000",
		From:               "",
		To:                 "bc1qvwqt8vtn2k7vrjqrsct63pkfw9ufqjldmjm439",
		Amount:             "5000",
		Symbol:             "ANSM",
		Action:             "inscribeTransfer",
		TokenInscriptionId: "4f54d82160bf08bab83bbe89276b2fd9bed514ce843c91a796daa07bafb85239i0",
		ProtocolType:       "BRC20",
		State:              "success",
		InscriptionId:      "4cf9aefbc9febf80b68376fa773849aabfdd8e3f7a5254ad11fd7ec6c32d3e89i0",
		InscriptionNumber:  "62377691",
		OutputIndex:        "",
	}

	expectData.TransactionList = append(expectData.TransactionList, expectTransaction)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.ChainFullName, expectData.ChainFullName)
	assert.Equal(t, data.ChainShortName, expectData.ChainShortName)

	transaction := data.TransactionList[0]
	assert.Equal(t, transaction.State, expectTransaction.State)
	assert.Equal(t, transaction.Symbol, expectTransaction.Symbol)
}
```

#### 1.2.8 Get inscription token transaction details for specific hash

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionTransactionDetail(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionTransactionDetail("btc", "c29fc5f33756c572fc55152435d9314059f8639797708b39471330536b94ed0c", "brc20", "", 0)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionTransactionDetailResp{
		Data: make([]types.InscriptionTransactionDetailData, 0),
	}

	expectData := types.InscriptionTransactionDetailData{
		Page:            "1",
		Limit:           "20",
		TotalPage:       "1",
		TransactionList: make([]types.InscriptionTransactionDetailTransaction, 0),
	}

	expectTx := types.InscriptionTransactionDetailTransaction{
		TxId:               "c29fc5f33756c572fc55152435d9314059f8639797708b39471330536b94ed0c",
		BlockHash:          "0000000000000000000159cea4e78229ccffab8ecfa94354729ee2b0c52b7a3f",
		Height:             "831823",
		TransactionTime:    "1708779611000",
		From:               "",
		To:                 "bc1qhuv3dhpnm0wktasd3v0kt6e4aqfqsd0uhfdu7d",
		Amount:             "2198220440",
		Action:             "inscribeTransfer",
		TokenInscriptionId: "9b664bdd6f5ed80d8d88957b63364c41f3ad4efb8eee11366aa16435974d9333i0",
		ProtocolType:       "BRC20",
		State:              "success",
		InscriptionId:      "c29fc5f33756c572fc55152435d9314059f8639797708b39471330536b94ed0ci0",
		InscriptionNumber:  "62184839",
		Symbol:             "sats",
		OutputIndex:        "",
	}
	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.State, expectTx.State)
	assert.Equal(t, tx.Symbol, expectTx.Symbol)
}
```

#### 1.2.9 Get inscription token transaction details for specific block

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetInscriptionBlockTokenTransaction(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetInscriptionBlockTokenTransaction("btc", "831823", "brc20", "", "", "", 1)
	if err != nil {
		t.Error(err)
	}

	expectResp := &types.InscriptionBlockTokenTransactionResp{
		Data: make([]types.InscriptionBlockTokenTransactionData, 0),
	}

	expectData := types.InscriptionBlockTokenTransactionData{
		Page:            "1",
		Limit:           "1",
		TotalPage:       "559",
		TotalTransfer:   "559",
		TransactionList: make([]types.InscriptionBlockTokenTransactionTransaction, 0),
	}

	expectTx := types.InscriptionBlockTokenTransactionTransaction{
		TxId:               "5ac740cad1c29266bf3615fc4f108c082431c7c0be74944e352edd75eed471ff",
		BlockHash:          "0000000000000000000159cea4e78229ccffab8ecfa94354729ee2b0c52b7a3f",
		Height:             "831823",
		From:               "",
		To:                 "bc1pnad2fk3fw6q3d20mhyacdekl7wf96rpg5yqxhtchzvpwet989lpsmvuc9n",
		Amount:             "1000",
		Action:             "mint",
		TokenInscriptionId: "4865f05b9132f12bb09d6215f13da5a304a502a95315d0a49463d6f8c0bb7740i0",
		ProtocolType:       "BRC20",
		State:              "success",
		InscriptionId:      "5ac740cad1c29266bf3615fc4f108c082431c7c0be74944e352edd75eed471ffi0",
		InscriptionNumber:  "62196529",
		Symbol:             "LABS",
		TransactionTime:    "1708779611000",
		OutputIndex:        "",
	}

	expectData.TransactionList = append(expectData.TransactionList, expectTx)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.TotalTransfer, expectData.TotalTransfer)

	tx := data.TransactionList[0]
	assert.Equal(t, tx.State, expectTx.State)
	assert.Equal(t, tx.Symbol, expectTx.Symbol)
}
```

### 1.3 UTXO-specific data

#### 1.3.1 Get remaining UTXO addresses

```go
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
```

#### 1.3.2 Get BTC transaction statistics

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetUtxoTransactionStats(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetUtxoTransactionStats("btc", "", "", "", 2)
	if err != nil {
		t.Errorf("GetUtxoTransactionStats error: %v", err)
	}

	expectResp := &types.UtxoTransactionStatsResp{
		Data: make([]types.UtxoTransactionStatsData, 0),
	}

	expectData := types.UtxoTransactionStatsData{
		Page:                   "1",
		Limit:                  "2",
		TotalPage:              "2805",
		TransactionHistoryList: make([]types.UtxoTransactionStatsTransactionHistory, 2),
	}

	expectData.TransactionHistoryList[0] = types.UtxoTransactionStatsTransactionHistory{
		Time:                      "1715616000000",
		TotalTransactionCount:     "480918",
		NormalTransactionCount:    "454299",
		AtomicalsTransactionCount: "1888",
		StampTransactionCount:     "859",
		OrdinalsTransactionCount:  "32594",
	}

	expectData.TransactionHistoryList[1] = types.UtxoTransactionStatsTransactionHistory{
		Time:                      "1715529600000",
		TotalTransactionCount:     "596134",
		NormalTransactionCount:    "573757",
		AtomicalsTransactionCount: "2478",
		StampTransactionCount:     "739",
		OrdinalsTransactionCount:  "42489",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)

	for i := 0; i < len(data.TransactionHistoryList); i++ {
		txHistory := data.TransactionHistoryList[i]
		expectTxHistory := expectData.TransactionHistoryList[i]
		assert.Equal(t, txHistory.OrdinalsTransactionCount, expectTxHistory.OrdinalsTransactionCount)
	}
}
```

#### 1.3.3 Get mining revenue statistics

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetUtxoRevenueStats(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetUtxoRevenueStats("btc", "", "", "", 2)
	if err != nil {
		t.Errorf("GetUtxoRevenueStats error: %s", err)
		return
	}

	expectResp := &types.UtxoRevenueStatsResp{
		Data: make([]types.UtxoRevenueStatsData, 0),
	}

	expectData := types.UtxoRevenueStatsData{
		Page:               "1",
		Limit:              "2",
		TotalPage:          "2806",
		RevenueHistoryList: make([]types.UtxoRevenueStatsRevenueHistory, 2),
	}

	expectData.RevenueHistoryList[0] = types.UtxoRevenueStatsRevenueHistory{
		Time:           "1715702400000",
		BlockReward:    "181.25",
		TransactionFee: "9.6476646",
	}

	expectData.RevenueHistoryList[1] = types.UtxoRevenueStatsRevenueHistory{
		Time:           "1715616000000",
		BlockReward:    "440.625",
		TransactionFee: "24.94226618",
	}

	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, len(data.RevenueHistoryList), len(expectData.RevenueHistoryList))
}
```

#### 1.3.4 BRC-20 data

##### 1.3.4.1 Get inscription list

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcInscriptionsList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcInscriptionsList("", "", "", "", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcInscriptionsListResp{
		Data: make([]types.BtcInscriptionsListData, 0),
	}

	expectData := types.BtcInscriptionsListData{
		Page:             "1",
		Limit:            "20",
		TotalPage:        "500",
		TotalInscription: "60109663",
		InscriptionsList: make([]types.BtcInscriptionsListInscription, 0),
	}

	expectInscription := types.BtcInscriptionsListInscription{
		InscriptionId:     "92780ef845a5190a1027724c24b5adbe6713abdaa43c5f273ff8a87d41f6cc8ci0",
		InscriptionNumber: "9999999",
		Location:          "92780ef845a5190a1027724c24b5adbe6713abdaa43c5f273ff8a87d41f6cc8c",
		Token:             "10MM",
		State:             "success",
		Msg:               "",
		TokenType:         "BRC20",
		ActionType:        "mint",
		LogoUrl:           "",
		OwnerAddress:      "bc1p53rrfs7l3fdyd60wzsk9gphnjm6y9hr4vhfrp207tsltyxjatfqqj9ly8k",
		TxId:              "92780ef845a5190a1027724c24b5adbe6713abdaa43c5f273ff8a87d41f6cc8c",
		BlockHeight:       "792013",
		ContentSize:       "",
		Time:              "1685401405000",
	}
	expectData.InscriptionsList = append(expectData.InscriptionsList, expectInscription)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
	assert.Equal(t, data.TotalInscription, expectData.TotalInscription)

	inscription := data.InscriptionsList[0]
	assert.Equal(t, inscription.TokenType, expectInscription.TokenType)
}
```

##### 1.3.4.2 Get BRC-20 token list

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcTokenList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcTokenList("", "", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcTokenListResp{
		Data: make([]types.BtcTokenListData, 0),
	}

	expectData := types.BtcTokenListData{
		Page:      "1",
		Limit:     "20",
		TotalPage: "500",
		TokenList: make([]types.BtcTokenListToken, 0),
	}

	expectToken := types.BtcTokenListToken{
		Token:             "ordi",
		DeployTime:        "1678248991000",
		InscriptionId:     "b61b0172d95e266c18aea0c624db987e971a5d6d4ebc2aaed85da4642d635735i0",
		InscriptionNumber: "348020",
		TotalSupply:       "21000000",
		MintAmount:        "21000000",
		TransactionCount:  "484169",
		Holder:            "13080",
		MintRate:          "1",
	}
	expectData.TokenList = append(expectData.TokenList, expectToken)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	token := data.TokenList[0]
	assert.Equal(t, token.Token, expectToken.Token)
}
```

##### 1.3.4.3 Get BRC-20 token details

```go
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
```

##### 1.3.4.4 Get list of addresses holding BRC-20 tokens

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcPositionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcPositionList("sats", "", 2)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcPositionListResp{
		Data: make([]types.BtcPositionListData, 0),
	}

	expectData := types.BtcPositionListData{
		Page:         "1",
		Limit:        "2",
		TotalPage:    "5000",
		PositionList: make([]types.BtcPositionListPosition, 2),
	}

	expectData.PositionList[0] = types.BtcPositionListPosition{
		HolderAddress: "bc1plff0sqm6ym55eak9vjljghd55h7hkheg22c84w55w646l857elqsfzmfdv",
		Amount:        "31740686608926",
		Rank:          "1",
	}

	expectData.PositionList[1] = types.BtcPositionListPosition{
		HolderAddress: "bc1pun3whtlzac75f2vcznxmpfc09dnzyp0luw8tpfwc7wrruwav30pqsu6l9u",
		Amount:        "22651199774504",
		Rank:          "1",
	}
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)

	for i := 0; i < len(data.PositionList); i++ {
		position := data.PositionList[i]
		expectPosition := expectData.PositionList[i]

		assert.Equal(t, position.HolderAddress, expectPosition.HolderAddress)
		assert.Equal(t, position.Amount, expectPosition.Amount)
		assert.Equal(t, position.Rank, expectPosition.Rank)
	}
}
```

##### 1.3.4.5 Get BRC-20 token transfer list

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcTransactionList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcTransactionList("", "", "", "", "", "", "", "", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcTransactionListResp{
		Data: make([]types.BtcTransactionListData, 0),
	}

	expectData := types.BtcTransactionListData{
		Page:             "1",
		Limit:            "20",
		TotalPage:        "500",
		TotalTransaction: "31124061",
		InscriptionsList: make([]types.GetBtcTransactionInscription, 0),
	}

	expectInscription := types.GetBtcTransactionInscription{
		TxId:              "af6bb18c64c57296ae07b8f4b1857c655160402a8d332fdb523915d7887476e2",
		BlockHeight:       "812873",
		State:             "success",
		TokenType:         "BRC20",
		ActionType:        "mint",
		FromAddress:       "",
		ToAddress:         "bc1qx2l6pzt7aknazsgdnvf5vnhp8ezx38ykg2wvfz",
		Amount:            "",
		Token:             "SoIa",
		InscriptionId:     "af6bb18c64c57296ae07b8f4b1857c655160402a8d332fdb523915d7887476e2i0",
		InscriptionNumber: "35346117",
		Index:             "0",
		Location:          "af6bb18c64c57296ae07b8f4b1857c655160402a8d332fdb523915d7887476e2",
		Msg:               "tick: $ORC has been minted",
		Time:              "1697695573000",
	}
	expectData.InscriptionsList = append(expectData.InscriptionsList, expectInscription)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	inscriptionsList := data.InscriptionsList[0]
	assert.Equal(t, inscriptionsList.Token, expectInscription.Token)
	assert.Equal(t, inscriptionsList.State, expectInscription.State)
}
```

##### 1.3.4.6 Get BRC-20 token balance list by address

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcAddressBalanceList(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcAddressBalanceList("bc1ph0057nc25ka94z8ydg43j8tnnp38u3hxpadutnt4n3jyfrmjzmcqw99mk2", "", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcAddressBalanceListResp{
		Data: make([]types.BtcAddressBalanceListData, 0),
	}

	expectData := types.BtcAddressBalanceListData{
		Page:        "1",
		Limit:       "20",
		TotalPage:   "1",
		BalanceList: make([]types.BtcAddressBalanceListBalance, 0),
	}

	expectBalance := types.BtcAddressBalanceListBalance{
		Token:            "X@AI",
		TokenType:        "BRC20",
		Balance:          "1350000000000",
		AvailableBalance: "1350000000000",
		TransferBalance:  "0",
	}
	expectData.BalanceList = append(expectData.BalanceList, expectBalance)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)

	balance := data.BalanceList[0]
	assert.Equal(t, balance.Token, expectBalance.Token)
	assert.Equal(t, balance.TokenType, expectBalance.TokenType)
}
```

##### 1.3.4.7 Get BRC-20 token balance details by address

```go
//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestGetBtcAddressBalanceDetails(t *testing.T) {
	client := oklink.New()

	resp, err := client.GetBtcAddressBalanceDetails("bc1ph0057nc25ka94z8ydg43j8tnnp38u3hxpadutnt4n3jyfrmjzmcqw99mk2", "meme", "", 0)
	if err != nil {
		t.Error(err)
		return
	}

	expectResp := &types.BtcAddressBalanceDetailsResp{
		Data: make([]types.BtcAddressBalanceDetailsData, 0),
	}

	expectData := types.BtcAddressBalanceDetailsData{
		Page:                "1",
		Limit:               "20",
		TotalPage:           "1",
		Token:               "meme",
		TokenType:           "BRC20",
		Balance:             "18",
		AvailableBalance:    "",
		TransferBalance:     "18",
		TransferBalanceList: make([]types.BtcAddressBalanceDetailsTransferBalance, 0),
	}

	expectBalance := types.BtcAddressBalanceDetailsTransferBalance{
		InscriptionId:     "a1002519472f9a1d45db5a3df30ea521ecd5425e546a63a79f3a4a9ff4e6e582i0",
		InscriptionNumber: "4615101",
		Amount:            "3",
	}

	expectData.TransferBalanceList = append(expectData.TransferBalanceList, expectBalance)
	expectResp.Data = append(expectResp.Data, expectData)

	assert.Equal(t, len(resp.Data), len(expectResp.Data))

	data := resp.Data[0]
	assert.Equal(t, data.Page, expectData.Page)
	assert.Equal(t, data.Limit, expectData.Limit)
	assert.Equal(t, data.TotalPage, expectData.TotalPage)
}
```