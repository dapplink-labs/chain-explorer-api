package account

import "github.com/dapplink-labs/chain-explorer-api/common"

type AccountBalanceRequest struct {
	ExplorerName string   `json:"explorerName"`
	Account      []string `json:"account"`
}

type AccountBalanceResponse struct {
	Account string         `json:"account"`
	Balance *common.BigInt `json:"balance"`
}

type AccountUtxoRequest struct {
	ExplorerName   string `json:"explorerName"`
	ChainShortName string `json:"chainShortName"`
	Address        string `json:"address"`
	Page           string `json:"page"`
	Limit          string `json:"limit"`
}

type AccountUtxoResponse struct {
	TotalUtxo     string `json:"totalUtxo"`
	Txid          string `json:"txid"`
	Height        string `json:"height"`
	BlockTime     string `json:"blockTime"`
	Address       string `json:"address"`
	UnspentAmount string `json:"unspentAmount"`
	Index         string `json:"index"`
}

type AccountTxRequest struct {
	ExplorerName string `json:"explorerName"`
	Action       string `json:"action"`
	Address      string `json:"address"`
	StartBlock   *int   `json:"startBlock"`
	EndBlock     *int   `json:"endBlock"`
	Page         int    `json:"page"`
	Offset       int    `json:"offset"`
	Desc         bool   `json:"desc"`
}

type AccountTxResponse struct {
	BlockNumber       int            `json:"blockNumber,string"`
	TimeStamp         common.Time    `json:"timeStamp"`
	Hash              string         `json:"hash"`
	From              string         `json:"from"`
	To                string         `json:"to"`
	Value             *common.BigInt `json:"value"`
	Nonce             int            `json:"nonce,string"`
	BlockHash         string         `json:"blockHash"`
	TransactionIndex  int            `json:"transactionIndex,string"`
	ContractAddress   string         `json:"contract_address"`
	Gas               int            `json:"gas,string"`
	GasPrice          *common.BigInt `json:"gasPrice"`
	GasUsed           int            `json:"gasUsed,string"`
	CumulativeGasUsed int            `json:"cumulativeGasUsed,string"`
	Input             string         `json:"input"`
	Confirmations     int            `json:"confirmations,string"`
	IsError           int            `json:"isError,string"`
	TxReceiptStatus   string         `json:"txreceipt_status"`
	TxType            string         `json:"tx_type"`
	MethodId          string         `json:"methodId"`
	FunctionName      string         `json:"functionName"`
	TokenID           *common.BigInt `json:"tokenID"`
	TokenName         string         `json:"tokenName"`
	TokenSymbol       string         `json:"tokenSymbol"`
	TokenDecimal      uint8          `json:"tokenDecimal,string"`
	TokenValue        uint8          `json:"tokenValue,string"`
	TraceID           string         `json:"traceId"`
	ErrCode           string         `json:"errCode"`
}
