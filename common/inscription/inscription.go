package inscription

type TokenListRequest struct {
	ChainShortName     string `json:"chainShortName"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	ProjectId          string `json:"projectId"`
	StartTime          string `json:"startTime"`
	EndTime            string `json:"endTime"`
	Page               string `json:"page"`
	Limit              int    `json:"limit"`
}

type TokenPositionListRequest struct {
	ChainShortName     string `json:"chainShortName"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	ProjectId          string `json:"projectId"`
	HolderAddress      string `json:"holderAddress"`
	Page               string `json:"page"`
	Limit              int    `json:"limit"`
}

type TokenTransactionListRequest struct {
	ChainShortName     string `json:"chainShortName"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	ProjectId          string `json:"projectId"`
	StartTime          string `json:"startTime"`
	EndTime            string `json:"endTime"`
	Page               string `json:"page"`
	Limit              int    `json:"limit"`
}

type InscriptionListRequest struct {
	ChainShortName     string `json:"chainShortName"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	ProjectId          string `json:"projectId"`
	Page               string `json:"page"`
	Limit              int    `json:"limit"`
}

type AddressTokenListRequest struct {
	ChainShortName     string `json:"chainShortName"`
	Address            string `json:"address"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	ProjectId          string `json:"projectId"`
	Page               string `json:"page"`
	Limit              int    `json:"limit"`
}

type AddressInscriptionListRequest struct {
	ChainShortName     string `json:"chainShortName"`
	Address            string `json:"address"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	ProjectId          string `json:"projectId"`
	Page               string `json:"page"`
	Limit              int    `json:"limit"`
}

type AddressTokenTransactionListRequest struct {
	ChainShortName     string `json:"chainShortName"`
	Address            string `json:"address"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	ProjectId          string `json:"projectId"`
	StartTime          string `json:"startTime"`
	EndTime            string `json:"endTime"`
	Page               string `json:"page"`
	Limit              int    `json:"limit"`
}

type TransactionDetailRequest struct {
	ChainShortName string `json:"chainShortName"`
	TxId           string `json:"txId"`
	ProtocolType   string `json:"protocolType"`
	Page           string `json:"page"`
	Limit          int    `json:"limit"`
}

type BlockTokenTransactionRequest struct {
	ChainShortName string `json:"chainShortName"`
	Height         string `json:"height"`
	ProtocolType   string `json:"protocolType"`
	TxnStartIndex  string `json:"txnStartIndex"`
	TxnEndIndex    string `json:"txnEndIndex"`
	Page           string `json:"page"`
	Limit          int    `json:"limit"`
}

// TokenListResp The Response structure of
// Btc inscription data -> Get list of inscription tokens
type TokenListResp struct {
	Data []TokenListData `json:"data"`
}

// TokenListData The Data field within the Response structure of
// Btc inscription data -> Get list of inscription tokens
type TokenListData struct {
	Page      string           `json:"page"`
	Limit     string           `json:"limit"`
	TotalPage string           `json:"totalPage"`
	TokenList []TokenListToken `json:"tokenList"`
}

// TokenListToken The TokenList field within the Data field within the Response structure of
// Btc inscription data -> Get list of inscription tokens
type TokenListToken struct {
	Symbol             string `json:"symbol"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	ProtocolType       string `json:"protocolType"`
	TotalSupply        string `json:"totalSupply"`
	MintAmount         string `json:"mintAmount"`
	DeployTime         string `json:"deployTime"`
	Holder             string `json:"holder"`
	TransactionCount   string `json:"transactionCount"`
	CirculatingSupply  string `json:"circulatingSupply"`
	MintBitwork        string `json:"mintBitwork"`
	LimitPerMint       string `json:"limitPerMint"`
	RunesSymbol        string `json:"runesSymbol"`
	Divisibility       string `json:"divisibility"`
	MintStatus         string `json:"mintStatus"`
	Premint            string `json:"premint"`
	Burn               string `json:"burn"`
	MintStartBlock     string `json:"mintStartBlock"`
	MintEndBlock       string `json:"mintEndBlock"`
	MintCap            string `json:"mintCap"`
}

// TokenPositionListResp The Response structure of
// Btc inscription data -> Get list of addresses holding inscription tokens
type TokenPositionListResp struct {
	Data []TokenPositionListData `json:"data"`
}

// TokenPositionListData The Data field within the Response structure of
// Btc inscription data -> Get list of addresses holding inscription tokens
type TokenPositionListData struct {
	Page         string                      `json:"page"`
	Limit        string                      `json:"limit"`
	TotalPage    string                      `json:"totalPage"`
	PositionList []TokenPositionListPosition `json:"positionList"`
}

// TokenPositionListPosition The PositionList field within the Data field within the Response structure of
// Btc inscription data -> Get list of addresses holding inscription tokens
type TokenPositionListPosition struct {
	HolderAddress string `json:"holderAddress"`
	Amount        string `json:"amount"`
	Rank          string `json:"rank"`
}

// TokenTransactionListResp The Response structure of
// Btc inscription data -> Get inscription token transfer list
type TokenTransactionListResp struct {
	Data []TokenTransactionListData `json:"data"`
}

// TokenTransactionListData The Data field within the Response structure of
// Btc inscription data -> Get inscription token transfer list
type TokenTransactionListData struct {
	Page            string                            `json:"page"`
	Limit           string                            `json:"limit"`
	TotalPage       string                            `json:"totalPage"`
	ChainFullName   string                            `json:"chainFullName"`
	ChainShortName  string                            `json:"chainShortName"`
	TotalTransfer   string                            `json:"totalTransfer"`
	TransactionList []TokenTransactionListTransaction `json:"transactionList"`
}

// TokenTransactionListTransaction The TransactionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token transfer list
type TokenTransactionListTransaction struct {
	TxId               string `json:"txId"`
	BlockHash          string `json:"blockHash"`
	Height             string `json:"height"`
	TransactionTime    string `json:"transactionTime"`
	From               string `json:"from"`
	To                 string `json:"to"`
	Amount             string `json:"amount"`
	Symbol             string `json:"symbol"`
	Action             string `json:"action"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	ProtocolType       string `json:"protocolType"`
	State              string `json:"state"`
	InscriptionId      string `json:"inscriptionId"`
	InscriptionNumber  string `json:"inscriptionNumber"`
	OutputIndex        string `json:"outputIndex"`
}

// InscriptionListResp The Response structure of
// Btc inscription data -> Get inscription list
type InscriptionListResp struct {
	Data []InscriptionListData `json:"data"`
}

// InscriptionListData The Data field within the Response structure of
// Btc inscription data -> Get inscription list
type InscriptionListData struct {
	Page            string                       `json:"page"`
	Limit           string                       `json:"limit"`
	TotalPage       string                       `json:"totalPage"`
	InscriptionList []InscriptionListInscription `json:"inscriptionList"`
}

// InscriptionListInscription The TokenList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription list
type InscriptionListInscription struct {
	InscriptionId      string `json:"inscriptionId"`
	InscriptionNumber  string `json:"inscriptionNumber"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	State              string `json:"state"`
	ProtocolType       string `json:"protocolType"`
	Action             string `json:"action"`
	OwnerAddress       string `json:"ownerAddress"`
}

// AddressTokenListResp The Response structure of
// Btc inscription data -> Get inscription token list by address
type AddressTokenListResp struct {
	Data []AddressTokenListData `json:"data"`
}

// AddressTokenListData The Data field within the Response structure of
// Btc inscription data -> Get inscription token list by address
type AddressTokenListData struct {
	Page      string                  `json:"page"`
	Limit     string                  `json:"limit"`
	TotalPage string                  `json:"totalPage"`
	TokenList []AddressTokenListToken `json:"tokenList"`
}

// AddressTokenListToken The TokenList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token list by address
type AddressTokenListToken struct {
	Symbol             string `json:"symbol"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	HoldingAmount      string `json:"holdingAmount"`
	InscriptionAmount  string `json:"inscriptionAmount"`
	AvailableAmount    string `json:"availableAmount"`
	TransferableAmount string `json:"transferableAmount"`
	InscriptionNumber  string `json:"inscriptionNumber"`
}

// AddressInscriptionListResp The Response structure of
// Btc inscription data -> Get inscription list by address
type AddressInscriptionListResp struct {
	Data []AddressInscriptionListData `json:"data"`
}

// AddressInscriptionListData The Data field within the Response structure of
// Btc inscription data -> Get inscription list by address
type AddressInscriptionListData struct {
	Page             string                              `json:"page"`
	Limit            string                              `json:"limit"`
	TotalPage        string                              `json:"totalPage"`
	TotalInscription string                              `json:"totalInscription"`
	InscriptionList  []AddressInscriptionListInscription `json:"inscriptionList"`
}

// AddressInscriptionListInscription The InscriptionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription list by address
type AddressInscriptionListInscription struct {
	InscriptionId      string `json:"inscriptionId"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	InscriptionNumber  string `json:"inscriptionNumber"`
	Symbol             string `json:"symbol"`
	State              string `json:"state"`
	ProtocolType       string `json:"protocolType"`
	Action             string `json:"action"`
}

// AddressTokenTransactionListResp The Response structure of
// Btc inscription data -> Get inscription token transfers by address
type AddressTokenTransactionListResp struct {
	Data []AddressTokenTransactionListData `json:"data"`
}

// AddressTokenTransactionListData The Data field within the Response structure of
// Btc inscription data -> Get inscription token transfers by address
type AddressTokenTransactionListData struct {
	Page            string                                   `json:"page"`
	Limit           string                                   `json:"limit"`
	TotalPage       string                                   `json:"totalPage"`
	ChainFullName   string                                   `json:"chainFullName"`
	ChainShortName  string                                   `json:"chainShortName"`
	TotalTransfer   string                                   `json:"totalTransfer"`
	TransactionList []AddressTokenTransactionListTransaction `json:"transactionList"`
}

// AddressTokenTransactionListTransaction The TransactionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token transfers by address
type AddressTokenTransactionListTransaction struct {
	TxId               string `json:"txId"`
	BlockHash          string `json:"blockHash"`
	Height             string `json:"height"`
	TransactionTime    string `json:"transactionTime"`
	From               string `json:"from"`
	To                 string `json:"to"`
	Amount             string `json:"amount"`
	Symbol             string `json:"symbol"`
	Action             string `json:"action"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	ProtocolType       string `json:"protocolType"`
	State              string `json:"state"`
	InscriptionId      string `json:"inscriptionId"`
	InscriptionNumber  string `json:"inscriptionNumber"`
	OutputIndex        string `json:"outputIndex"`
}

// TransactionDetailResp The Response structure of
// Btc inscription data -> Get inscription token transaction details for specific hash
type TransactionDetailResp struct {
	Data []TransactionDetailData `json:"data"`
}

// TransactionDetailData The Data field within the Response structure of
// Btc inscription data -> Get inscription token transaction details for specific hash
type TransactionDetailData struct {
	Page            string                         `json:"page"`
	Limit           string                         `json:"limit"`
	TotalPage       string                         `json:"totalPage"`
	TransactionList []TransactionDetailTransaction `json:"transactionList"`
}

// TransactionDetailTransaction The TransactionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token transaction details for specific hash
type TransactionDetailTransaction struct {
	TxId               string `json:"txId"`
	BlockHash          string `json:"blockHash"`
	Height             string `json:"height"`
	TransactionTime    string `json:"transactionTime"`
	From               string `json:"from"`
	To                 string `json:"to"`
	Amount             string `json:"amount"`
	Action             string `json:"action"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	ProtocolType       string `json:"protocolType"`
	State              string `json:"state"`
	InscriptionId      string `json:"inscriptionId"`
	InscriptionNumber  string `json:"inscriptionNumber"`
	Symbol             string `json:"symbol"`
	OutputIndex        string `json:"outputIndex"`
}

// BlockTokenTransactionResp The Response structure of
// Btc inscription data -> Get inscription token transaction details for specific block
type BlockTokenTransactionResp struct {
	Data []BlockTokenTransactionData `json:"data"`
}

// BlockTokenTransactionData The Data field within the Response structure of
// Btc inscription data -> Get inscription token transaction details for specific block
type BlockTokenTransactionData struct {
	Page            string                             `json:"page"`
	Limit           string                             `json:"limit"`
	TotalPage       string                             `json:"totalPage"`
	TotalTransfer   string                             `json:"totalTransfer"`
	TransactionList []BlockTokenTransactionTransaction `json:"transactionList"`
}

// BlockTokenTransactionTransaction The TransactionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token transaction details for specific block
type BlockTokenTransactionTransaction struct {
	TxId               string `json:"txId"`
	BlockHash          string `json:"blockHash"`
	Height             string `json:"height"`
	From               string `json:"from"`
	To                 string `json:"to"`
	Amount             string `json:"amount"`
	Action             string `json:"action"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	ProtocolType       string `json:"protocolType"`
	State              string `json:"state"`
	InscriptionId      string `json:"inscriptionId"`
	InscriptionNumber  string `json:"inscriptionNumber"`
	Symbol             string `json:"symbol"`
	TransactionTime    string `json:"transactionTime"`
	OutputIndex        string `json:"outputIndex"`
}
