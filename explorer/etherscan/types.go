package etherscan

import (
	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"net/url"
	"strconv"
)

type TransactionResponse[T any] struct {
	Page             string `json:"page"`
	Limit            string `json:"limit"`
	TotalPage        string `json:"totalPage"`
	ChainFullName    string `json:"chainFullName"`
	ChainShortName   string `json:"chainShortName"`
	TransactionList  []T    `json:"transactionList"`
	TransactionLists []T    `json:"transactionLists"`
}

type AccountBalance struct {
	Account string         `json:"account"`
	Balance *common.BigInt `json:"balance"`
}

type GasTrackerGasOracleResp struct {
	LastBlock       string `json:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}

type AddressTransactionRequest struct {
	Page       uint64         `json:"page"`
	Offset     uint64         `json:"offset"`
	Address    string         `url:"address"`
	StartBlock uint64         `url:"startblock"`
	EndBlock   uint64         `url:"endblock"`
	Sort       chain.SortType `url:"sort"`
}

func (req AddressTransactionRequest) ToQueryUrl() string {
	values := url.Values{}
	if req.Page != 0 {
		values.Set("page", strconv.FormatUint(req.Page, 10))
	}
	if req.Offset != 0 {
		values.Set("offset", strconv.FormatUint(req.Offset, 10))
	}
	if req.Address != "" {
		values.Set("address", req.Address)
	}
	if req.StartBlock != 0 {
		values.Set("startblock", strconv.FormatUint(req.StartBlock, 10))
	}
	if req.EndBlock != 0 {
		values.Set("endblock", strconv.FormatUint(req.EndBlock, 10))
	}
	if req.Sort != "" {
		values.Set("sort", string(req.Sort))
	}
	return values.Encode()
}

func (req AddressTransactionRequest) ToQueryParamMap() map[string]any {
	result := make(map[string]any)

	if req.Page != 0 {
		result["page"] = strconv.FormatUint(req.Page, 10)
	}
	if req.Offset != 0 {
		result["offset"] = strconv.FormatUint(req.Offset, 10)
	}
	if req.Address != "" {
		result["address"] = req.Address
	}
	if req.StartBlock != 0 {
		result["startblock"] = strconv.FormatUint(req.StartBlock, 10)
	}
	if req.EndBlock != 0 {
		result["endblock"] = strconv.FormatUint(req.EndBlock, 10)
	}
	if req.Sort != "" {
		result["sort"] = string(req.Sort)
	}

	return result
}

type AddressTransactionResp struct {
	BlockNumber      string `json:"blockNumber"`
	BlockHash        string `json:"blockHash"`
	TimeStamp        string `json:"timeStamp"`
	Hash             string `json:"hash"`
	Nonce            string `json:"nonce"`
	TransactionIndex string `json:"transactionIndex"`
	From             string `json:"from"`
	To               string `json:"to"`
	Value            string `json:"value"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Input            string `json:"input"`
	MethodId         string `json:"methodId"`
	FunctionName     string `json:"functionName"`
	Type             string `json:"type"`

	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	IsError           string `json:"isError"`
	TraceId           string `json:"traceId"`

	TokenID      string `json:"tokenID,omitempty"`
	TokenName    string `json:"tokenName,omitempty"`
	TokenSymbol  string `json:"tokenSymbol,omitempty"`
	TokenDecimal string `json:"tokenDecimal,omitempty"`
	TokenValue   string `json:"tokenValue,omitempty"`
}

type TokensResp struct {
	Symbol          string `json:"symbol"`
	ContractAddress string `json:"contractAddress"`
	TokenId         string `json:"tokenId"`
	TotalSupply     string `json:"totalSupply"`
	Divisor         string `json:"divisor"`
	Decimal         string `json:"decimal"`
}
