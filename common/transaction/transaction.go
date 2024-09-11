package transaction

type TxRequest struct {
	ExplorerName   string `json:"explorerName"`
	ChainShortName string `json:"chainShortName"`
	Txid           string `json:"txid"`
}

type TxResponse struct {
	ChainFullName        string                `json:"chainFullName"`
	ChainShortName       string                `json:"chainShortName"`
	Txid                 string                `json:"txid"`
	Height               string                `json:"height"`
	TransactionTime      string                `json:"transactionTime"`
	TransactionType      string                `json:"transactionType"`
	Amount               string                `json:"amount"`
	TransactionSymbol    string                `json:"transactionSymbol"`
	MethodId             string                `json:"methodId"`
	ErrorLog             string                `json:"errorLog"`
	InputData            string                `json:"inputData"`
	Txfee                string                `json:"txfee"`
	Index                string                `json:"index"`
	Confirm              string                `json:"confirm"`
	InputDetails         []InputDetail         `json:"inputDetails"`
	OutputDetails        []OutputDetail        `json:"outputDetails"`
	State                string                `json:"state"`
	GasLimit             string                `json:"gasLimit"`
	GasUsed              string                `json:"gasUsed"`
	GasPrice             string                `json:"gasPrice"`
	TotalTransactionSize string                `json:"totalTransactionSize"`
	VirtualSize          string                `json:"virtualSize"`
	Weight               string                `json:"weight"`
	Nonce                string                `json:"nonce"`
	IsAaTransaction      bool                  `json:"isAaTransaction"`
	TokenTransferDetails []TokenTransferDetail `json:"tokenTransferDetails"`
	ContractDetails      []ContractDetail      `json:"contractDetails"`
}

type InputDetail struct {
	InputHash  string `json:"inputHash"`
	IsContract bool   `json:"isContract"`
	Amount     string `json:"amount"`
}

type OutputDetail struct {
	OutputHash string `json:"outputHash"`
	IsContract bool   `json:"isContract"`
	Amount     string `json:"amount"`
}

type TokenTransferDetail struct {
	Index                string `json:"index"`
	Token                string `json:"token"`
	TokenContractAddress string `json:"tokenContractAddress"`
	Symbol               string `json:"symbol"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	IsFromContract       bool   `json:"isFromContract"`
	IsToContract         bool   `json:"isToContract"`
	TokenId              string `json:"tokenId"`
	Amount               string `json:"amount"`
}

type ContractDetail struct {
	Index          string `json:"index"`
	From           string `json:"from"`
	To             string `json:"to"`
	IsFromContract bool   `json:"isFromContract"`
	IsToContract   bool   `json:"isToContract"`
	Amount         string `json:"amount"`
	GasLimit       string `json:"gasLimit"`
}
