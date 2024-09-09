package utxo

type TransactionStatsRequest struct {
	ChainShortName string `json:"chainShortName"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	Page           string `json:"page"`
	Limit          int    `json:"limit"`
}

type RevenueStatsRequest struct {
	ChainShortName string `json:"chainShortName"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	Page           string `json:"page"`
	Limit          int    `json:"limit"`
}

// TransactionStatsResp The Response structure of
// UTXO-specific data -> Get BTC transaction statistics
type TransactionStatsResp struct {
	Data []TransactionStatsData `json:"data"`
}

// TransactionStatsData The Data field within the Response structure of
// UTXO-specific data -> Get BTC transaction statistics
type TransactionStatsData struct {
	Page                   string                               `json:"page"`
	Limit                  string                               `json:"limit"`
	TotalPage              string                               `json:"totalPage"`
	TransactionHistoryList []TransactionStatsTransactionHistory `json:"transactionHistoryList"`
}

// TransactionStatsTransactionHistory The TransactionHistoryList field within the Data field within the Response structure of
// UTXO-specific data -> Get BTC transaction statistics
type TransactionStatsTransactionHistory struct {
	Time                      string `json:"time"`
	TotalTransactionCount     string `json:"totalTransactionCount"`
	NormalTransactionCount    string `json:"normalTransactionCount"`
	AtomicalsTransactionCount string `json:"atomicalsTransactionCount"`
	StampTransactionCount     string `json:"stampTransactionCount"`
	OrdinalsTransactionCount  string `json:"ordinalsTransactionCount"`
}

// RevenueStatsResp The Response structure of
// UTXO-specific data -> Get mining revenue statistics
type RevenueStatsResp struct {
	Data []RevenueStatsData `json:"data"`
}

// RevenueStatsData The Data field within the Response structure of
// UTXO-specific data -> Get mining revenue statistics
type RevenueStatsData struct {
	Page               string                       `json:"page"`
	Limit              string                       `json:"limit"`
	TotalPage          string                       `json:"totalPage"`
	RevenueHistoryList []RevenueStatsRevenueHistory `json:"revenueHistoryList"`
}

// RevenueStatsRevenueHistory The RevenueHistoryList field within the Data field within the Response structure of
// UTXO-specific data -> Get mining revenue statistics
type RevenueStatsRevenueHistory struct {
	Time           string `json:"time"`
	BlockReward    string `json:"blockReward"`
	TransactionFee string `json:"transactionFee"`
}
