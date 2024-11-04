package account

import (
	"net/url"
	"strconv"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
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

type ActionType string

const (
	OkLinkActionNormal   ActionType = "normal"
	OkLinkActionInternal ActionType = "internal"
	OkLinkActionToken    ActionType = "token"
	OkLinkActionUtxo     ActionType = "utxo"

	EtherscanActionTxList         ActionType = "txlist"
	EtherscanActionTxListInternal ActionType = "txlistinternal"
	EtherscanActionTokenTx        ActionType = "tokentx"
	EtherscanActionTokenNftTx     ActionType = "tokennfttx"
	EtherscanActionToken1155Tx    ActionType = "token1155tx"

	SolScanActionSol ActionType = "sol"
	SolScanActionSpl ActionType = "spl"
)

type ProtocolType string

const (
	ProtocolTypeToken20   ProtocolType = "token_20"
	ProtocolTypeToken721  ProtocolType = "token_721"
	ProtocolTypeToken1155 ProtocolType = "token_1155"
)

type IsFromToType string

const (
	From IsFromToType = "from"
	To   IsFromToType = "to"
)

type PageRequest struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}
type AccountTxRequest struct {
	chain.PageRequest
	ExplorerName   string `json:"explorerName"`
	ChainShortName string `json:"chainShortName"`
	// normal internal token
	Action ActionType     `json:"action"`
	Sort   chain.SortType `url:"sort"`

	Address          string `json:"address"`
	StartBlockHeight uint64 `json:"startBlockHeight"`
	EndBlockHeight   uint64 `json:"endBlockHeight"`
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
	if req.StartBlockHeight != 0 {
		params.Add("startBlockHeight", strconv.FormatUint(req.StartBlockHeight, 10))
	}
	if req.EndBlockHeight != 0 {
		params.Add("endBlockHeight", strconv.FormatUint(req.EndBlockHeight, 10))
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
	if req.Limit != 0 {
		params.Add("limit", strconv.FormatUint(req.Limit, 10))
	}
	if req.Page != 0 {
		params.Add("page", strconv.FormatUint(req.Page, 10))
	}
	return params.Encode()
}

type AccountTxResponse struct {
	TxId             string `json:"txId"`
	BlockHash        string `json:"blockHash"`
	Height           string `json:"height"`
	TransactionTime  string `json:"transactionTime"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Nonce            string `json:"nonce"`

	Amount string `json:"amount"`
	Symbol string `json:"symbol"`

	Operation       string `json:"operation"`
	GasPrice        string `json:"gasPrice"`
	GasLimit        string `json:"gasLimit"`
	GasUsed         string `json:"gasUsed"`
	TxFee           string `json:"txFee"`
	State           string `json:"state"`
	TransactionType string `json:"transactionType"`
	Confirmations   string `json:"confirmations"`
	IsError         string `json:"isError"`
	TraceId         string `json:"traceId"`

	Input                string `json:"input"`
	MethodId             string `json:"methodId"`
	FunctionName         string `json:"functionName"`
	TokenContractAddress string `json:"tokenContractAddress"`

	IsFromContract bool `json:"isFromContract"`
	IsToContract   bool `json:"isToContract"`

	TokenId      string `json:"tokenId"`
	TokenName    string `json:"tokenName"`
	TokenSymbol  string `json:"tokenSymbol"`
	TokenDecimal string `json:"tokenDecimal"`
	TokenValue   string `json:"tokenValue"`
}

type TransactionResponse[T any] struct {
	chain.PageResponse
	TransactionList []T `json:"transactionList"`
}
