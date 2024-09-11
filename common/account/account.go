package account

import (
	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"net/url"
)

type AccountBalanceRequest struct {
	ChainShortName  string   `json:"chainShortName"`
	ExplorerName    string   `json:"explorerName"`
	Account         []string `json:"account"`
	Symbol          []string `json:"symbol"`
	ContractAddress []string `json:"contractAddress"`
	ProtocolType    []string `json:"protocolType"` // 20代币：token_20; 721代币：token_721; 1155代币：token_1155;
	Page            []string `json:"page"`
	Limit           []string `json:"limit"`
}

type AccountBalanceResponse struct {
	Account         string         `json:"account"`
	Balance         *common.BigInt `json:"balance"`
	BalanceStr      string         `json:"balanceStr"`
	Symbol          string         `json:"symbol"`
	ContractAddress string         `json:"contractAddress"`
	TokenId         string         `json:"token_id"`
}

type AccountUtxoRequest struct {
	ExplorerName   string `json:"explorerName"`
	ChainShortName string `json:"chainShortName"`
	Address        string `json:"address"`
	Page           string `json:"page"`
	Limit          string `json:"limit"`
}

type AccountUtxoResponse struct {
	TxId          string `json:"txid"`
	Height        string `json:"height"`
	BlockTime     string `json:"blockTime"`
	Address       string `json:"address"`
	UnspentAmount string `json:"unspentAmount"`
	Index         string `json:"index"`
}

type Action string

const (
	OkLinkActionNormal   Action = "normal"
	OkLinkActionInternal Action = "internal"
	OkLinkActionToken    Action = "token"

	EtherscanActionTxList         Action = "txlist"
	EtherscanActionTxListInternal Action = "txlistinternal"
	EtherscanActionTokenTx        Action = "tokentx"
	EtherscanActionTokenNftTx     Action = "tokennfttx"
	EtherscanActionToken1155Tx    Action = "token1155tx"
)

type ProtocolType string

const (
	ProtocolTypeToken20   ProtocolType = "token_20"
	ProtocolTypeToken721  ProtocolType = "token_721"
	ProtocolTypeToken1155 ProtocolType = "token_1155"
)

type AccountTxRequest struct {
	chain.PageRequest

	ExplorerName   string `json:"explorerName"`
	ChainShortName string `json:"chainShortName"`
	// normal internal token
	Action Action `json:"action"`

	Address          string `json:"address"`
	StartBlockHeight string `json:"startBlockHeight"`
	EndBlockHeight   string `json:"endBlockHeight"`
	IsFromOrTo       string `json:"isFromOrTo"`
	// token_20 token_721 token_1155
	ProtocolType         ProtocolType `json:"protocolType"`
	TokenContractAddress string       `json:"tokenContractAddress"`
}

func (req *AccountTxRequest) ToQueryUrl() string {
	params := url.Values{}

	if req.ExplorerName != "" {
		params.Add("explorerName", req.ExplorerName)
	}
	if req.ChainShortName != "" {
		params.Add("chainShortName", req.ChainShortName)
	}
	if req.Action != "" {
		params.Add("action", string(req.Action))
	}
	if req.Address != "" {
		params.Add("address", req.Address)
	}
	if req.StartBlockHeight != "" {
		params.Add("startBlockHeight", req.StartBlockHeight)
	}
	if req.EndBlockHeight != "" {
		params.Add("endBlockHeight", req.EndBlockHeight)
	}
	if req.IsFromOrTo != "" {
		params.Add("isFromOrTo", req.IsFromOrTo)
	}
	if req.ProtocolType != "" {
		params.Add("protocolType", string(req.ProtocolType))
	}
	if req.TokenContractAddress != "" {
		params.Add("tokenContractAddress", req.TokenContractAddress)
	}
	if req.Limit != "" {
		params.Add("limit", req.Limit)
	}
	if req.Page != "" {
		params.Add("page", req.Page)
	}
	return params.Encode()
}

type AccountTxResponse struct {
	TxId                 string `json:"txId"`
	BlockHash            string `json:"blockHash"`
	Height               string `json:"height"`
	TransactionTime      string `json:"transactionTime"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenId              string `json:"tokenId"`
	Amount               string `json:"amount"`
	Symbol               string `json:"symbol"`
	IsFromContract       bool   `json:"isFromContract"`
	IsToContract         bool   `json:"isToContract"`
	Operation            string `json:"operation"`
	MethodId             string `json:"methodId"`
	Nonce                string `json:"nonce"`
	GasPrice             string `json:"gasPrice"`
	GasLimit             string `json:"gasLimit"`
	GasUsed              string `json:"gasUsed"`
	TxFee                string `json:"txFee"`
	State                string `json:"state"`
	TransactionType      string `json:"transactionType"`
}

type TransactionResponse[T any] struct {
	chain.PageResponse
	TransactionList []T `json:"transactionList"`
}
