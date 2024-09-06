package types

// GetGasResp The Response structure of
// Fundamental blockchain data -> Fundamental data -> Get transaction or Gas fee
type GetGasResp struct {
	Data []GetGasData `json:"data"`
}

// GetGasData The Data field within the Response structure of
// Fundamental blockchain data -> Fundamental data -> Get transaction or Gas fee
type GetGasData struct {
	ChainFullName         string `json:"chainFullName"`
	ChainShortName        string `json:"chainShortName"`
	Symbol                string `json:"symbol"`
	BestTransactionFee    string `json:"bestTransactionFee"`
	BestTransactionFeeSat string `json:"bestTransactionFeeSat"`
	RecommendedGasPrice   string `json:"recommendedGasPrice"`
	RapidGasPrice         string `json:"rapidGasPrice"`
	StandardGasPrice      string `json:"standardGasPrice"`
	SlowGasPrice          string `json:"slowGasPrice"`
	BaseFee               string `json:"baseFee"`
	GasUsedRatio          string `json:"gasUsedRatio"`
}

// AddressSummaryResp The Response structure of
// Fundamental blockchain data -> Address Data -> Get basic address details
type AddressSummaryResp struct {
	Data []AddressSummaryData `json:"data"`
}

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

// AddressTokenBalanceResp The Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details by address
type AddressTokenBalanceResp struct {
	Data []AddressTokenBalanceData `json:"data"`
}

// AddressTokenBalanceData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details by address
type AddressTokenBalanceData struct {
	Limit     string                     `json:"limit"`
	Page      string                     `json:"page"`
	TotalPage string                     `json:"totalPage"`
	TokenList []AddressTokenBalanceToken `json:"tokenList"`
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

// AddressBalanceFillsResp The Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details
type AddressBalanceFillsResp struct {
	Data []AddressBalanceFillsData `json:"data"`
}

// AddressBalanceFillsData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details
type AddressBalanceFillsData struct {
	Page           string                     `json:"page"`
	Limit          string                     `json:"limit"`
	TotalPage      string                     `json:"totalPage"`
	ChainFullName  string                     `json:"chainFullName"`
	ChainShortName string                     `json:"chainShortName"`
	TokenList      []AddressBalanceFillsToken `json:"tokenList"`
}

// AddressBalanceFillsToken The TokenList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance details
type AddressBalanceFillsToken struct {
	Token                string `json:"token"`
	TokenId              string `json:"tokenId"`
	HoldingAmount        string `json:"holdingAmount"`
	TotalTokenValue      string `json:"totalTokenValue"`
	Change24h            string `json:"change24h"`
	PriceUsd             string `json:"priceUsd"`
	ValueUsd             string `json:"valueUsd"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

// BlockAddressBalanceHistoryResp The Response structure of
// Fundamental blockchain data -> Address data -> Get address balance at specific block height
type BlockAddressBalanceHistoryResp struct {
	Data []BlockAddressBalanceHistoryData `json:"data"`
}

// BlockAddressBalanceHistoryData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get address balance at specific block height
type BlockAddressBalanceHistoryData struct {
	Address              string `json:"address"`
	Height               string `json:"height"`
	Balance              string `json:"balance"`
	BalanceSymbol        string `json:"balanceSymbol"`
	TokenContractAddress string `json:"tokenContractAddress"`
	BlockTime            string `json:"blockTime"`
}

// AddressTransactionListResp The Response structure of
// Fundamental blockchain data -> Address data -> Get address transaction list
type AddressTransactionListResp struct {
	Data []AddressTransactionListData `json:"data"`
}

// AddressTransactionListData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get address transaction list
type AddressTransactionListData struct {
	Page            string               `json:"page"`
	Limit           string               `json:"limit"`
	TotalPage       string               `json:"totalPage"`
	ChainFullName   string               `json:"chainFullName"`
	ChainShortName  string               `json:"chainShortName"`
	TransactionList []AddressTransaction `json:"transactionLists"`
}

// AddressTransaction The TransactionLists field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get address transaction list
type AddressTransaction struct {
	TxId                 string `json:"txId"`
	MethodId             string `json:"methodId"`
	BlockHash            string `json:"blockHash"`
	Height               string `json:"height"`
	TransactionTime      string `json:"transactionTime"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	IsFromContract       bool   `json:"isFromContract"`
	IsToContract         bool   `json:"isToContract"`
	Amount               string `json:"amount"`
	TransactionSymbol    string `json:"transactionSymbol"`
	TxFee                string `json:"txFee"`
	State                string `json:"state"`
	TokenId              string `json:"tokenId"`
	TokenContractAddress string `json:"tokenContractAddress"`
	ChallengeStatus      string `json:"challengeStatus"`
	L1OriginHash         string `json:"l1OriginHash"`
}

// AddressNormalTransactionListResp The Response structure of
// Fundamental blockchain data -> Address data -> Get standard transaction list by address
type AddressNormalTransactionListResp struct {
	Data []AddressNormalTransactionListData `json:"data"`
}

// AddressNormalTransactionListData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get standard transaction list by address
type AddressNormalTransactionListData struct {
	Limit           string                                    `json:"limit"`
	Page            string                                    `json:"page"`
	TotalPage       string                                    `json:"totalPage"`
	TransactionList []AddressNormalTransactionListTransaction `json:"transactionList"`
}

// AddressNormalTransactionListTransaction The TransactionList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get standard transaction list by address
type AddressNormalTransactionListTransaction struct {
	TxId            string `json:"txId"`
	MethodId        string `json:"methodId"`
	Nonce           string `json:"nonce"`
	GasPrice        string `json:"gasPrice"`
	GasLimit        string `json:"gasLimit"`
	GasUsed         string `json:"gasUsed"`
	BlockHash       string `json:"blockHash"`
	Height          string `json:"height"`
	TransactionTime string `json:"transactionTime"`
	From            string `json:"from"`
	To              string `json:"to"`
	IsFromContract  bool   `json:"isFromContract"`
	IsToContract    bool   `json:"isToContract"`
	Amount          string `json:"amount"`
	Symbol          string `json:"symbol"`
	TxFee           string `json:"txFee"`
	State           string `json:"state"`
	TransactionType string `json:"transactionType"`
}

// AddressInternalTransactionListResp The Response structure of
// Fundamental blockchain data -> Address data -> Get internal transaction list by address
type AddressInternalTransactionListResp struct {
	Data []AddressInternalTransactionListData `json:"data"`
}

// AddressInternalTransactionListData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get internal transaction list by address
type AddressInternalTransactionListData struct {
	Limit           string                                      `json:"limit"`
	Page            string                                      `json:"page"`
	TotalPage       string                                      `json:"totalPage"`
	TransactionList []AddressInternalTransactionListTransaction `json:"transactionList"`
}

// AddressInternalTransactionListTransaction The TransactionList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get internal transaction list by address
type AddressInternalTransactionListTransaction struct {
	TxId            string `json:"txId"`
	Operation       string `json:"operation"`
	BlockHash       string `json:"blockHash"`
	Height          string `json:"height"`
	TransactionTime string `json:"transactionTime"`
	From            string `json:"from"`
	To              string `json:"to"`
	IsFromContract  bool   `json:"isFromContract"`
	IsToContract    bool   `json:"isToContract"`
	Amount          string `json:"amount"`
	Symbol          string `json:"symbol"`
}

// AddressTokenTransactionListResp The Response structure of
// Fundamental blockchain data -> Address data -> Get token transaction list by address
type AddressTokenTransactionListResp struct {
	Data []AddressTokenTransactionListData `json:"data"`
}

// AddressTokenTransactionListData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token transaction list by address
type AddressTokenTransactionListData struct {
	Limit           string                                   `json:"limit"`
	Page            string                                   `json:"page"`
	TotalPage       string                                   `json:"totalPage"`
	TransactionList []AddressTokenTransactionListTransaction `json:"transactionList"`
}

// AddressTokenTransactionListTransaction The TransactionList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token transaction list by address
type AddressTokenTransactionListTransaction struct {
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
}

// AddressBalanceMultiResp The Response structure of
// Fundamental blockchain data -> Address data -> Get native token balance in batches
type AddressBalanceMultiResp struct {
	Data []AddressBalanceMultiData `json:"data"`
}

// AddressBalanceMultiData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get native token balance in batches
type AddressBalanceMultiData struct {
	Symbol      string                       `json:"symbol"`
	BalanceList []AddressBalanceMultiBalance `json:"balanceList"`
}

// AddressBalanceMultiBalance The BalanceList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get native token balance in batches
type AddressBalanceMultiBalance struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

// AddressTokenBalanceMultiResp The Response structure of
// Fundamental blockchain data -> Address data -> Get token balance in batches
type AddressTokenBalanceMultiResp struct {
	Data []AddressTokenBalanceMultiData `json:"data"`
}

// AddressTokenBalanceMultiData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance in batches
type AddressTokenBalanceMultiData struct {
	Page        string                            `json:"page"`
	Limit       string                            `json:"limit"`
	TotalPage   string                            `json:"totalPage"`
	BalanceList []AddressTokenBalanceMultiBalance `json:"balanceList"`
}

// AddressTokenBalanceMultiBalance The BalanceList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token balance in batches
type AddressTokenBalanceMultiBalance struct {
	Address              string `json:"address"`
	HoldingAmount        string `json:"holdingAmount"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

// AddressNormalTransactionListMultiResp The Response structure of
// Fundamental blockchain data -> Address data -> Get standard transaction list for specific block
type AddressNormalTransactionListMultiResp struct {
	Data []AddressNormalTransactionListMultiData `json:"data"`
}

// AddressNormalTransactionListMultiData 地址数据 -> The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get standard transaction list for specific block
type AddressNormalTransactionListMultiData struct {
	Page            string                                         `json:"page"`
	Limit           string                                         `json:"limit"`
	TotalPage       string                                         `json:"totalPage"`
	TransactionList []AddressNormalTransactionListMultiTransaction `json:"transactionList"`
}

// AddressNormalTransactionListMultiTransaction The TransactionList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get standard transaction list for specific block
type AddressNormalTransactionListMultiTransaction struct {
	TxId            string `json:"txId"`
	MethodId        string `json:"methodId"`
	BlockHash       string `json:"blockHash"`
	Height          string `json:"height"`
	TransactionTime string `json:"transactionTime"`
	From            string `json:"from"`
	To              string `json:"to"`
	IsFromContract  bool   `json:"isFromContract"`
	IsToContract    bool   `json:"isToContract"`
	Amount          string `json:"amount"`
	Symbol          string `json:"symbol"`
	TxFee           string `json:"txFee"`
	GasLimit        string `json:"gasLimit"`
	GasUsed         string `json:"gasUsed"`
	GasPrice        string `json:"gasPrice"`
	Nonce           string `json:"nonce"`
	TransactionType string `json:"transactionType"`
}

// AddressInternalTransactionListMultiResp The Response structure of
// Fundamental blockchain data -> Address data -> Get internal transaction list for specific block
type AddressInternalTransactionListMultiResp struct {
	Data []AddressInternalTransactionListMultiData `json:"data"`
}

// AddressInternalTransactionListMultiData The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get internal transaction list for specific block
type AddressInternalTransactionListMultiData struct {
	Page            string                                           `json:"page"`
	Limit           string                                           `json:"limit"`
	TotalPage       string                                           `json:"totalPage"`
	TransactionList []AddressInternalTransactionListMultiTransaction `json:"transactionList"`
}

// AddressInternalTransactionListMultiTransaction The TransactionList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get internal transaction list for specific block
type AddressInternalTransactionListMultiTransaction struct {
	TxId            string `json:"txId"`
	BlockHash       string `json:"blockHash"`
	Height          string `json:"height"`
	TransactionTime string `json:"transactionTime"`
	Operation       string `json:"operation"`
	From            string `json:"from"`
	To              string `json:"to"`
	IsFromContract  bool   `json:"isFromContract"`
	IsToContract    bool   `json:"isToContract"`
	Amount          string `json:"amount"`
}

// AddressTokenTransactionListMultiResp The Response structure of
// Fundamental blockchain data -> Address data -> Get token transaction list for specific block
type AddressTokenTransactionListMultiResp struct {
	Data []AddressTokenTransactionListMultiData `json:"data"`
}

// AddressTokenTransactionListMultiData 地址数据 -> The Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token transaction list for specific block
type AddressTokenTransactionListMultiData struct {
	Page            string                                        `json:"page"`
	Limit           string                                        `json:"limit"`
	TotalPage       string                                        `json:"totalPage"`
	TransactionList []AddressTokenTransactionListMultiTransaction `json:"transactionList"`
}

// AddressTokenTransactionListMultiTransaction The TransactionList field within the Data field within the Response structure of
// Fundamental blockchain data -> Address data -> Get token transaction list for specific block
type AddressTokenTransactionListMultiTransaction struct {
	TxId                 string `json:"txId"`
	BlockHash            string `json:"blockHash"`
	Height               string `json:"height"`
	TransactionTime      string `json:"transactionTime"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	IsFromContract       bool   `json:"isFromContract"`
	IsToContract         bool   `json:"isToContract"`
	Amount               string `json:"amount"`
	TokenId              string `json:"tokenId"`
	Symbol               string `json:"symbol"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

// InscriptionTokenListResp The Response structure of
// Btc inscription data -> Get list of inscription tokens
type InscriptionTokenListResp struct {
	Data []InscriptionTokenListData `json:"data"`
}

// InscriptionTokenListData The Data field within the Response structure of
// Btc inscription data -> Get list of inscription tokens
type InscriptionTokenListData struct {
	Page      string                      `json:"page"`
	Limit     string                      `json:"limit"`
	TotalPage string                      `json:"totalPage"`
	TokenList []InscriptionTokenListToken `json:"tokenList"`
}

// InscriptionTokenListToken The TokenList field within the Data field within the Response structure of
// Btc inscription data -> Get list of inscription tokens
type InscriptionTokenListToken struct {
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

// InscriptionTokenPositionListResp The Response structure of
// Btc inscription data -> Get list of addresses holding inscription tokens
type InscriptionTokenPositionListResp struct {
	Data []InscriptionTokenPositionListData `json:"data"`
}

// InscriptionTokenPositionListData The Data field within the Response structure of
// Btc inscription data -> Get list of addresses holding inscription tokens
type InscriptionTokenPositionListData struct {
	Page         string                                 `json:"page"`
	Limit        string                                 `json:"limit"`
	TotalPage    string                                 `json:"totalPage"`
	PositionList []InscriptionTokenPositionListPosition `json:"positionList"`
}

// InscriptionTokenPositionListPosition The PositionList field within the Data field within the Response structure of
// Btc inscription data -> Get list of addresses holding inscription tokens
type InscriptionTokenPositionListPosition struct {
	HolderAddress string `json:"holderAddress"`
	Amount        string `json:"amount"`
	Rank          string `json:"rank"`
}

// InscriptionTokenTransactionListResp The Response structure of
// Btc inscription data -> Get inscription token transfer list
type InscriptionTokenTransactionListResp struct {
	Data []InscriptionTokenTransactionListData `json:"data"`
}

// InscriptionTokenTransactionListData The Data field within the Response structure of
// Btc inscription data -> Get inscription token transfer list
type InscriptionTokenTransactionListData struct {
	Page            string                                       `json:"page"`
	Limit           string                                       `json:"limit"`
	TotalPage       string                                       `json:"totalPage"`
	ChainFullName   string                                       `json:"chainFullName"`
	ChainShortName  string                                       `json:"chainShortName"`
	TotalTransfer   string                                       `json:"totalTransfer"`
	TransactionList []InscriptionTokenTransactionListTransaction `json:"transactionList"`
}

// InscriptionTokenTransactionListTransaction The TransactionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token transfer list
type InscriptionTokenTransactionListTransaction struct {
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

// InscriptionInscriptionListResp The Response structure of
// Btc inscription data -> Get inscription list
type InscriptionInscriptionListResp struct {
	Data []InscriptionInscriptionListData `json:"data"`
}

// InscriptionInscriptionListData The Data field within the Response structure of
// Btc inscription data -> Get inscription list
type InscriptionInscriptionListData struct {
	Page            string                                  `json:"page"`
	Limit           string                                  `json:"limit"`
	TotalPage       string                                  `json:"totalPage"`
	InscriptionList []InscriptionInscriptionListInscription `json:"inscriptionList"`
}

// InscriptionInscriptionListInscription The TokenList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription list
type InscriptionInscriptionListInscription struct {
	InscriptionId      string `json:"inscriptionId"`
	InscriptionNumber  string `json:"inscriptionNumber"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	State              string `json:"state"`
	ProtocolType       string `json:"protocolType"`
	Action             string `json:"action"`
	OwnerAddress       string `json:"ownerAddress"`
}

// InscriptionAddressTokenListResp The Response structure of
// Btc inscription data -> Get inscription token list by address
type InscriptionAddressTokenListResp struct {
	Data []InscriptionAddressTokenListData `json:"data"`
}

// InscriptionAddressTokenListData The Data field within the Response structure of
// Btc inscription data -> Get inscription token list by address
type InscriptionAddressTokenListData struct {
	Page      string                             `json:"page"`
	Limit     string                             `json:"limit"`
	TotalPage string                             `json:"totalPage"`
	TokenList []InscriptionAddressTokenListToken `json:"tokenList"`
}

// InscriptionAddressTokenListToken The TokenList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token list by address
type InscriptionAddressTokenListToken struct {
	Symbol             string `json:"symbol"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	HoldingAmount      string `json:"holdingAmount"`
	InscriptionAmount  string `json:"inscriptionAmount"`
	AvailableAmount    string `json:"availableAmount"`
	TransferableAmount string `json:"transferableAmount"`
	InscriptionNumber  string `json:"inscriptionNumber"`
}

// InscriptionAddressInscriptionListResp The Response structure of
// Btc inscription data -> Get inscription list by address
type InscriptionAddressInscriptionListResp struct {
	Data []InscriptionAddressInscriptionListData `json:"data"`
}

// InscriptionAddressInscriptionListData The Data field within the Response structure of
// Btc inscription data -> Get inscription list by address
type InscriptionAddressInscriptionListData struct {
	Page             string                                         `json:"page"`
	Limit            string                                         `json:"limit"`
	TotalPage        string                                         `json:"totalPage"`
	TotalInscription string                                         `json:"totalInscription"`
	InscriptionList  []InscriptionAddressInscriptionListInscription `json:"inscriptionList"`
}

// InscriptionAddressInscriptionListInscription The InscriptionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription list by address
type InscriptionAddressInscriptionListInscription struct {
	InscriptionId      string `json:"inscriptionId"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	InscriptionNumber  string `json:"inscriptionNumber"`
	Symbol             string `json:"symbol"`
	State              string `json:"state"`
	ProtocolType       string `json:"protocolType"`
	Action             string `json:"action"`
}

// InscriptionAddressTokenTransactionListResp The Response structure of
// Btc inscription data -> Get inscription token transfers by address
type InscriptionAddressTokenTransactionListResp struct {
	Data []InscriptionAddressTokenTransactionListData `json:"data"`
}

// InscriptionAddressTokenTransactionListData The Data field within the Response structure of
// Btc inscription data -> Get inscription token transfers by address
type InscriptionAddressTokenTransactionListData struct {
	Page            string                                              `json:"page"`
	Limit           string                                              `json:"limit"`
	TotalPage       string                                              `json:"totalPage"`
	ChainFullName   string                                              `json:"chainFullName"`
	ChainShortName  string                                              `json:"chainShortName"`
	TotalTransfer   string                                              `json:"totalTransfer"`
	TransactionList []InscriptionAddressTokenTransactionListTransaction `json:"transactionList"`
}

// InscriptionAddressTokenTransactionListTransaction The TransactionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token transfers by address
type InscriptionAddressTokenTransactionListTransaction struct {
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

// InscriptionTransactionDetailResp The Response structure of
// Btc inscription data -> Get inscription token transaction details for specific hash
type InscriptionTransactionDetailResp struct {
	Data []InscriptionTransactionDetailData `json:"data"`
}

// InscriptionTransactionDetailData The Data field within the Response structure of
// Btc inscription data -> Get inscription token transaction details for specific hash
type InscriptionTransactionDetailData struct {
	Page            string                                    `json:"page"`
	Limit           string                                    `json:"limit"`
	TotalPage       string                                    `json:"totalPage"`
	TransactionList []InscriptionTransactionDetailTransaction `json:"transactionList"`
}

// InscriptionTransactionDetailTransaction The TransactionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token transaction details for specific hash
type InscriptionTransactionDetailTransaction struct {
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

// InscriptionBlockTokenTransactionResp The Response structure of
// Btc inscription data -> Get inscription token transaction details for specific block
type InscriptionBlockTokenTransactionResp struct {
	Data []InscriptionBlockTokenTransactionData `json:"data"`
}

// InscriptionBlockTokenTransactionData The Data field within the Response structure of
// Btc inscription data -> Get inscription token transaction details for specific block
type InscriptionBlockTokenTransactionData struct {
	Page            string                                        `json:"page"`
	Limit           string                                        `json:"limit"`
	TotalPage       string                                        `json:"totalPage"`
	TotalTransfer   string                                        `json:"totalTransfer"`
	TransactionList []InscriptionBlockTokenTransactionTransaction `json:"transactionList"`
}

// InscriptionBlockTokenTransactionTransaction The TransactionList field within the Data field within the Response structure of
// Btc inscription data -> Get inscription token transaction details for specific block
type InscriptionBlockTokenTransactionTransaction struct {
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

// AddressUtxoResp The Response structure of
// UTXO-specific data -> Get remaining UTXO addresses
type AddressUtxoResp struct {
	Data []AddressUtxoData `json:"data"`
}

// AddressUtxoData The Data field within the Response structure of
// UTXO-specific data -> Get remaining UTXO addresses
type AddressUtxoData struct {
	Page      string            `json:"page"`
	Limit     string            `json:"limit"`
	TotalPage string            `json:"totalPage"`
	UtxoList  []AddressUtxoUtxo `json:"utxoList"`
}

// AddressUtxoUtxo The UtxoList field within the Data field within the Response structure of
// UTXO-specific data -> Get remaining UTXO addresses
type AddressUtxoUtxo struct {
	TxId          string `json:"txid"`
	Height        string `json:"height"`
	BlockTime     string `json:"blockTime"`
	Address       string `json:"address"`
	UnspentAmount string `json:"unspentAmount"`
	Index         string `json:"index"`
}

// UtxoTransactionStatsResp The Response structure of
// UTXO-specific data -> Get BTC transaction statistics
type UtxoTransactionStatsResp struct {
	Data []UtxoTransactionStatsData `json:"data"`
}

// UtxoTransactionStatsData The Data field within the Response structure of
// UTXO-specific data -> Get BTC transaction statistics
type UtxoTransactionStatsData struct {
	Page                   string                                   `json:"page"`
	Limit                  string                                   `json:"limit"`
	TotalPage              string                                   `json:"totalPage"`
	TransactionHistoryList []UtxoTransactionStatsTransactionHistory `json:"transactionHistoryList"`
}

// UtxoTransactionStatsTransactionHistory The TransactionHistoryList field within the Data field within the Response structure of
// UTXO-specific data -> Get BTC transaction statistics
type UtxoTransactionStatsTransactionHistory struct {
	Time                      string `json:"time"`
	TotalTransactionCount     string `json:"totalTransactionCount"`
	NormalTransactionCount    string `json:"normalTransactionCount"`
	AtomicalsTransactionCount string `json:"atomicalsTransactionCount"`
	StampTransactionCount     string `json:"stampTransactionCount"`
	OrdinalsTransactionCount  string `json:"ordinalsTransactionCount"`
}

// UtxoRevenueStatsResp The Response structure of
// UTXO-specific data -> Get mining revenue statistics
type UtxoRevenueStatsResp struct {
	Data []UtxoRevenueStatsData `json:"data"`
}

// UtxoRevenueStatsData The Data field within the Response structure of
// UTXO-specific data -> Get mining revenue statistics
type UtxoRevenueStatsData struct {
	Page               string                           `json:"page"`
	Limit              string                           `json:"limit"`
	TotalPage          string                           `json:"totalPage"`
	RevenueHistoryList []UtxoRevenueStatsRevenueHistory `json:"revenueHistoryList"`
}

// UtxoRevenueStatsRevenueHistory The RevenueHistoryList field within the Data field within the Response structure of
// UTXO-specific data -> Get mining revenue statistics
type UtxoRevenueStatsRevenueHistory struct {
	Time           string `json:"time"`
	BlockReward    string `json:"blockReward"`
	TransactionFee string `json:"transactionFee"`
}

// BtcInscriptionsListResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get inscription list
type BtcInscriptionsListResp struct {
	Data []BtcInscriptionsListData `json:"data"`
}

// BtcInscriptionsListData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get inscription list
type BtcInscriptionsListData struct {
	Page             string                           `json:"page"`
	Limit            string                           `json:"limit"`
	TotalPage        string                           `json:"totalPage"`
	TotalInscription string                           `json:"totalInscription"`
	InscriptionsList []BtcInscriptionsListInscription `json:"inscriptionsList"`
}

// BtcInscriptionsListInscription The InscriptionsList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get inscription list
type BtcInscriptionsListInscription struct {
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

// BtcTokenListResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token list
type BtcTokenListResp struct {
	Data []BtcTokenListData `json:"data"`
}

// BtcTokenListData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token list
type BtcTokenListData struct {
	Page      string              `json:"page"`
	Limit     string              `json:"limit"`
	TotalPage string              `json:"totalPage"`
	TokenList []BtcTokenListToken `json:"tokenList"`
}

// BtcTokenListToken The TokenList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token list
type BtcTokenListToken struct {
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

// BtcTokenDetailsResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token details
type BtcTokenDetailsResp struct {
	Data []BtcTokenDetailsTokenData `json:"data"`
}

// BtcTokenDetailsTokenData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token details
type BtcTokenDetailsTokenData struct {
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

// BtcPositionListResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get list of addresses holding BRC-20 tokens
type BtcPositionListResp struct {
	Data []BtcPositionListData `json:"data"`
}

// BtcPositionListData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get list of addresses holding BRC-20 tokens
type BtcPositionListData struct {
	Page         string                    `json:"page"`
	Limit        string                    `json:"limit"`
	TotalPage    string                    `json:"totalPage"`
	PositionList []BtcPositionListPosition `json:"positionList"`
}

// BtcPositionListPosition The PositionList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get list of addresses holding BRC-20 tokens
type BtcPositionListPosition struct {
	HolderAddress string `json:"holderAddress"`
	Amount        string `json:"amount"`
	Rank          string `json:"rank"`
}

// BtcTransactionListResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token transfer list
type BtcTransactionListResp struct {
	Data []BtcTransactionListData `json:"data"`
}

// BtcTransactionListData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token transfer list
type BtcTransactionListData struct {
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

// BtcAddressBalanceListResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance list by address
type BtcAddressBalanceListResp struct {
	Data []BtcAddressBalanceListData `json:"data"`
}

// BtcAddressBalanceListData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance list by address
type BtcAddressBalanceListData struct {
	Page        string                         `json:"page"`
	Limit       string                         `json:"limit"`
	TotalPage   string                         `json:"totalPage"`
	BalanceList []BtcAddressBalanceListBalance `json:"balanceList"`
}

// BtcAddressBalanceListBalance The BalanceList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance list by address
type BtcAddressBalanceListBalance struct {
	Token            string `json:"token"`
	TokenType        string `json:"tokenType"`
	Balance          string `json:"balance"`
	AvailableBalance string `json:"availableBalance"`
	TransferBalance  string `json:"transferBalance"`
}

// BtcAddressBalanceDetailsResp The Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance details by address
type BtcAddressBalanceDetailsResp struct {
	Data []BtcAddressBalanceDetailsData `json:"data"`
}

// BtcAddressBalanceDetailsData The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance details by address
type BtcAddressBalanceDetailsData struct {
	Page                string                                    `json:"page"`
	Limit               string                                    `json:"limit"`
	TotalPage           string                                    `json:"totalPage"`
	Token               string                                    `json:"token"`
	TokenType           string                                    `json:"tokenType"`
	Balance             string                                    `json:"balance"`
	AvailableBalance    string                                    `json:"availableBalance"`
	TransferBalance     string                                    `json:"transferBalance"`
	TransferBalanceList []BtcAddressBalanceDetailsTransferBalance `json:"transferBalanceList"`
}

// BtcAddressBalanceDetailsTransferBalance The TransferBalanceList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance details by address
type BtcAddressBalanceDetailsTransferBalance struct {
	InscriptionId     string `json:"inscriptionId"`
	InscriptionNumber string `json:"inscriptionNumber"`
	Amount            string `json:"amount"`
}
