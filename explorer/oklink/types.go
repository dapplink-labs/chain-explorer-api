package oklink

// AddressSummaryData The Data field within the Response structure of
// Fundamental blockchain data -> Address Data -> Get basic address details
type AddressSummaryData struct {
	ChainFullName                 string `json:"chainFullName"`
	ChainShortName                string `json:"chainShortName"`
	Address                       string `json:"address"`
	ContractAddress               string `json:"contractAddress"`
	Balance                       string `json:"balance"`
	BalanceSymbol                 string `json:"balanceSymbol"`
	TransactionCount              string `json:"transactionCount"`
	Verifying                     string `json:"verifying"`
	SendAmount                    string `json:"sendAmount"`
	ReceiveAmount                 string `json:"receiveAmount"`
	TokenAmount                   string `json:"tokenAmount"`
	TotalTokenValue               string `json:"totalTokenValue"`
	CreateContractAddress         string `json:"createContractAddress"`
	CreateContractTransactionHash string `json:"createContractTransactionHash"`
	FirstTransactionTime          string `json:"firstTransactionTime"`
	LastTransactionTime           string `json:"lastTransactionTime"`
	Token                         string `json:"token"`
	Bandwidth                     string `json:"bandwidth"`
	Energy                        string `json:"energy"`
	VotingRights                  string `json:"votingRights"`
	UnclaimedVotingRewards        string `json:"unclaimedVotingRewards"`
	IsAaAddress                   bool   `json:"isAaAddress"`
}

// AddressSummaryResp The Response structure of
// Fundamental blockchain data -> Address Data -> Get basic address details
type AddressSummaryResp struct {
	Data []AddressSummaryData `json:"data"`
}

// AddressTokenBalanceToken The TokenList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details by address
type AddressTokenBalanceToken struct {
	Symbol               string `json:"symbol"`
	TokenContractAddress string `json:"tokenContractAddress"`
	HoldingAmount        string `json:"holdingAmount"`
	PriceUsd             string `json:"priceUsd"`
	ValueUsd             string `json:"valueUsd"`
	TokenId              string `json:"tokenId"`
}

// AddressTokenBalanceData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details by address
type AddressTokenBalanceData struct {
	Limit     string                     `json:"limit"`
	Page      string                     `json:"page"`
	TotalPage string                     `json:"totalPage"`
	TokenList []AddressTokenBalanceToken `json:"tokenList"`
}

// AddressTokenBalanceResp The Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details by address
type AddressTokenBalanceResp struct {
	Data []AddressTokenBalanceData `json:"data"`
}
