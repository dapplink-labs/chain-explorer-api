package solscan

// AddressSummaryData The Data field within the Response structure of
// Fundamental blockchain data -> Address Data -> Get basic address details

type AddressSummaryData struct {
	TokenAccount  string `json:"token_account"`
	TokenAddress  string `json:"token_address"`
	Amount        int64  `json:"amount"`
	TokenDecimals int    `json:"token_decimals"`
	Owner         string `json:"owner"`
}
