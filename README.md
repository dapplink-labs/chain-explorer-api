<!--
parent:
  order: false
-->

<div align="center">
  <h1> chain-explorer-api repo </h1>
</div>

<div align="center">
  <a href="https://github.com/dapplink-labs/chain-explorer-api/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/dapplink-labs/chain-explorer-api.svg" />
  </a>
  <a href="https://github.com/dapplink-labs/chain-explorer-api/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/dapplink-labs/chain-explorer-api.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/dapplink-labs/chain-explorer-api">
    <img alt="GoDoc" src="https://godoc.org/github.com/dapplink-labs/chain-explorer-api?status.svg" />
  </a>
</div>

Chain Explorer API is a statistical browser interface gateway that simplifies the calling process of browser APIs, provides a unified imported parameter exported parameter underlying call library, and helps third-party projects simplify the development process.

**Tips**: need [Go 1.18+](https://golang.org/dl/)

## Install

### Use this Library

#### 1.New ChainExplorerAdaptor

```
var (
	EtherscanBaseUrl = "https://api.etherscan.io/api?"
	EtherscanApiKey  = ""
	EtherscanTimeout = time.Second * 20

	OklinkBaseUrl = "https://www.oklink.com/"
	OklinkApiKey  = ""
	OkTimeout     = time.Second * 20
)

func NewChainExplorerAdaptor() (*oklink.ChainExplorerAdaptor, *etherscan.ChainExplorerAdaptor, error) {
	var err error
	oklinkCli, err := oklink.NewChainExplorerAdaptor(OklinkApiKey, OklinkBaseUrl, false, time.Duration(OkTimeout))
	if err != nil {
		fmt.Println("Mock oklink client fail", "err", err)
		return nil, nil, err
	}

	etherscanCli, err := etherscan.NewChainExplorerAdaptor(EtherscanApiKey, EtherscanBaseUrl, false, time.Duration(EtherscanTimeout))
	if err != nil {
		fmt.Println("Mock etherscan client fail", "err", err)
		return nil, nil, err
	}
	return oklinkCli, etherscanCli, nil
}
```

#### 2. Example Case

```
func GetAccountBalance() {
	oklinkClient, etherscanClient, err := NewChainExplorerAdaptor()
	if err != nil {
		fmt.Println("new chain explorer adaptor fail", "err", err)
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
	etherscanResp, err := etherscanClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========etherscanResp============")
	fmt.Println(etherscanResp.Balance.Int())
	fmt.Println(etherscanResp.Account)
	fmt.Println(etherscanResp.Symbol)
	fmt.Println("===========etherscanResp===========")

	oklinkResp, err := oklinkClient.GetAccountBalance(acbr)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("==========oklinkResp============")
	fmt.Println(oklinkResp.BalanceStr)
	fmt.Println(oklinkResp.Account)
	fmt.Println(oklinkResp.Symbol)
	fmt.Println("==========oklinkResp============")
}
```

## Contribute

### 1.fork repo

fork chain-explorer-api to your github

### 2.clone repo

```bash
git@github.com:guoshijiang/chain-explorer-api.git
```

### 3. create new branch and commit code

```bash
git branch -C xxx
git checkout xxx

coding

git add .
git commit -m "xxx"
git push origin xxx
```

### 4.commit PR

Have a pr on your github and submit it to the chain-explorer-api repository

### 5.review

After the chain-explorer-api code maintainer has passed the review, the code will be merged into the chain-explorer-api library. At this point, your PR submission is complete

### 6.Disclaimer

This code has not yet been audited, and should not be used in any production systems.
