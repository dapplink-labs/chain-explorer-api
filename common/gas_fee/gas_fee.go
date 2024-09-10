package gas_fee

type GasEstimateFeeRequest struct {
	ExplorerName   string `json:"explorerName"`
	ChainShortName string `json:"chainShortName"`
	GasLimit       int    `json:"gasLimit"`
}

type GasEstimateFeeResponse struct {
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
