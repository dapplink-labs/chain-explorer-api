package gas_fee

type GasEstimateFeeRequest struct {
	ExplorerName   string `json:"explorerName"`
	ChainShortName string `json:"chainShortName"`
	GasLimit       int    `json:"gasLimit"`
}

type GasEstimateFeeResponse struct {
	ChainName             string `json:"chainName"`
	BestTransactionFee    string `json:"bestTransactionFee"`
	BestTransactionFeeSat string `json:"bestTransactionFeeSat"`
	RecommendedGasPrice   string `json:"recommendedGasPrice"`
	RapidGasPrice         string `json:"rapidGasPrice"`
	StandardGasPrice      string `json:"standardGasPrice"`
	SlowGasPrice          string `json:"slowGasPrice"`
	BaseFee               string `json:"baseFee"`
	GasUsedRatio          string `json:"gasUsedRatio"`
}
