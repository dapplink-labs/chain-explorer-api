package etherscan

import "github.com/dapplink-labs/chain-explorer-api/common"

type ApiResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  T      `json:"result"`
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
	Address    string `url:"address"`
	StartBlock int    `url:"startblock"`
	EndBlock   int    `url:"endblock"`
	Page       int    `url:"page"`
	Offset     int    `url:"offset"`
	Sort       string `url:"sort"`
}

type AddressTransactionResp struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxReceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	MethodId          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
	TokenID           string `json:"tokenID,omitempty"`
	TokenName         string `json:"tokenName,omitempty"`
	TokenSymbol       string `json:"tokenSymbol,omitempty"`
	TokenDecimal      string `json:"tokenDecimal,omitempty"`
	TokenValue        string `json:"tokenValue,omitempty"`
}
