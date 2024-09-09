package btc

type InscriptionsListRequest struct {
	Token             string `json:"token"`
	InscriptionId     string `json:"inscription_id"`
	InscriptionNumber string `json:"inscription_number"`
	State             string `json:"state"`
	Page              string `json:"page"`
	Limit             int    `json:"limit"`
}

// InscriptionsListResponse The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get inscription list
type InscriptionsListResponse struct {
	Page             string                        `json:"page"`
	Limit            string                        `json:"limit"`
	TotalPage        string                        `json:"totalPage"`
	TotalInscription string                        `json:"totalInscription"`
	InscriptionsList []InscriptionsListInscription `json:"inscriptionsList"`
}

// InscriptionsListInscription The InscriptionsList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get inscription list
type InscriptionsListInscription struct {
	InscriptionId     string `json:"inscriptionId"`
	InscriptionNumber string `json:"inscriptionNumber"`
	Location          string `json:"location"`
	Token             string `json:"token"`
	State             string `json:"state"`
	Msg               string `json:"msg"`
	TokenType         string `json:"tokenType"`
	ActionType        string `json:"actionType"`
	LogoUrl           string `json:"logoUrl"`
	OwnerAddress      string `json:"ownerAddress"`
	TxId              string `json:"txId"`
	BlockHeight       string `json:"blockHeight"`
	ContentSize       string `json:"contentSize"`
	Time              string `json:"time"`
}

type TokenListRequest struct {
	Token   string `json:"token"`
	OrderBy string `json:"orderBy"`
	Page    string `json:"page"`
	Limit   int    `json:"limit"`
}

// TokenListResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token list
type TokenListResp struct {
	Data []TokenListData `json:"data"`
}

// TokenListData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token list
type TokenListData struct {
	Page      string           `json:"page"`
	Limit     string           `json:"limit"`
	TotalPage string           `json:"totalPage"`
	TokenList []TokenListToken `json:"tokenList"`
}

// TokenListToken The TokenList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token list
type TokenListToken struct {
	Token             string `json:"token"`
	DeployTime        string `json:"deployTime"`
	InscriptionId     string `json:"inscriptionId"`
	InscriptionNumber string `json:"inscriptionNumber"`
	TotalSupply       string `json:"totalSupply"`
	MintAmount        string `json:"mintAmount"`
	TransactionCount  string `json:"transactionCount"`
	Holder            string `json:"holder"`
	MintRate          string `json:"mintRate"`
}

type TokenDetailsRequest struct {
	Token string `json:"token"`
}

// TokenDetailsResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token details
type TokenDetailsResp struct {
	Data []TokenDetailsTokenData `json:"data"`
}

// TokenDetailsTokenData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token details
type TokenDetailsTokenData struct {
	Token             string `json:"token"`
	Precision         string `json:"precision"`
	TotalSupply       string `json:"totalSupply"`
	MintAmount        string `json:"mintAmount"`
	LimitPerMint      string `json:"limitPerMint"`
	Holder            string `json:"holder"`
	DeployAddress     string `json:"deployAddress"`
	LogoUrl           string `json:"logoUrl"`
	TxId              string `json:"txId"`
	InscriptionId     string `json:"inscriptionId"`
	DeployHeight      string `json:"deployHeight"`
	DeployTime        string `json:"deployTime"`
	InscriptionNumber string `json:"inscriptionNumber"`
	State             string `json:"state"`
	TokenType         string `json:"tokenType"`
	Msg               string `json:"msg"`
}

type PositionListRequest struct {
	Token string `json:"token"`
	Page  string `json:"page"`
	Limit int    `json:"limit"`
}

// PositionListResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get list of addresses holding BRC-20 tokens
type PositionListResp struct {
	Data []PositionListData `json:"data"`
}

// PositionListData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get list of addresses holding BRC-20 tokens
type PositionListData struct {
	Page         string                 `json:"page"`
	Limit        string                 `json:"limit"`
	TotalPage    string                 `json:"totalPage"`
	PositionList []PositionListPosition `json:"positionList"`
}

// PositionListPosition The PositionList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get list of addresses holding BRC-20 tokens
type PositionListPosition struct {
	HolderAddress string `json:"holderAddress"`
	Amount        string `json:"amount"`
	Rank          string `json:"rank"`
}

type TransactionListRequest struct {
	Address           string `json:"address"`
	Token             string `json:"token"`
	InscriptionNumber string `json:"inscriptionNumber"`
	ActionType        string `json:"actionType"`
	ToAddress         string `json:"toAddress"`
	FromAddress       string `json:"fromAddress"`
	TxId              string `json:"txId"`
	BlockHeight       string `json:"blockHeight"`
	Page              string `json:"page"`
	Limit             int    `json:"limit"`
}

// TransactionListResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token transfer list
type TransactionListResp struct {
	Data []TransactionListData `json:"data"`
}

// TransactionListData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token transfer list
type TransactionListData struct {
	Page             string                         `json:"page"`
	Limit            string                         `json:"limit"`
	TotalPage        string                         `json:"totalPage"`
	TotalTransaction string                         `json:"totalTransaction"`
	InscriptionsList []GetBtcTransactionInscription `json:"inscriptionsList"`
}

// GetBtcTransactionInscription The InscriptionsList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token transfer list
type GetBtcTransactionInscription struct {
	TxId              string `json:"txId"`
	BlockHeight       string `json:"blockHeight"`
	State             string `json:"state"`
	TokenType         string `json:"tokenType"`
	ActionType        string `json:"actionType"`
	FromAddress       string `json:"fromAddress"`
	ToAddress         string `json:"toAddress"`
	Amount            string `json:"amount"`
	Token             string `json:"token"`
	InscriptionId     string `json:"inscriptionId"`
	InscriptionNumber string `json:"inscriptionNumber"`
	Index             string `json:"index"`
	Location          string `json:"location"`
	Msg               string `json:"msg"`
	Time              string `json:"time"`
}

type AddressBalanceListRequest struct {
	Address string `json:"address"`
	Token   string `json:"token"`
	Page    string `json:"page"`
	Limit   int    `json:"limit"`
}

// AddressBalanceListResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance list by address
type AddressBalanceListResp struct {
	Data []AddressBalanceListData `json:"data"`
}

// AddressBalanceListData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance list by address
type AddressBalanceListData struct {
	Page        string                      `json:"page"`
	Limit       string                      `json:"limit"`
	TotalPage   string                      `json:"totalPage"`
	BalanceList []AddressBalanceListBalance `json:"balanceList"`
}

// AddressBalanceListBalance The BalanceList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance list by address
type AddressBalanceListBalance struct {
	Token            string `json:"token"`
	TokenType        string `json:"tokenType"`
	Balance          string `json:"balance"`
	AvailableBalance string `json:"availableBalance"`
	TransferBalance  string `json:"transferBalance"`
}

type AddressBalanceDetailsRequest struct {
	Address string `json:"address"`
	Token   string `json:"token"`
	Page    string `json:"page"`
	Limit   int    `json:"limit"`
}

// AddressBalanceDetailsResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance details by address
type AddressBalanceDetailsResp struct {
	Data []AddressBalanceDetailsData `json:"data"`
}

// AddressBalanceDetailsData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance details by address
type AddressBalanceDetailsData struct {
	Page                string                                 `json:"page"`
	Limit               string                                 `json:"limit"`
	TotalPage           string                                 `json:"totalPage"`
	Token               string                                 `json:"token"`
	TokenType           string                                 `json:"tokenType"`
	Balance             string                                 `json:"balance"`
	AvailableBalance    string                                 `json:"availableBalance"`
	TransferBalance     string                                 `json:"transferBalance"`
	TransferBalanceList []AddressBalanceDetailsTransferBalance `json:"transferBalanceList"`
}

// AddressBalanceDetailsTransferBalance The TransferBalanceList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance details by address
type AddressBalanceDetailsTransferBalance struct {
	InscriptionId     string `json:"inscriptionId"`
	InscriptionNumber string `json:"inscriptionNumber"`
	Amount            string `json:"amount"`
}
