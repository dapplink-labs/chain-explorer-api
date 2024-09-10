package gas_tracker

import (
	"time"
)

type GasEstimateRequest struct {
	ExplorerName string `json:"explorer_name"`
	GasLimit     int    `json:"gas_limit"`
}

type GasEstimateResponse struct {
	ConfirmationTimes GasConfirmationTimes `json:"confirmation_times"`
}

type GasConfirmationTimes struct {
	Slow   time.Duration `json:"slow"`
	Medium time.Duration `json:"medium"`
	Fast   time.Duration `json:"fast"`
}

type GasOracleRequest struct {
	ExplorerName string `json:"explorer_name"`
}

type GasOracleResponse struct {
	LastBlock            int
	SafeGasPrice         float64
	ProposeGasPrice      float64
	FastGasPrice         float64
	SuggestBaseFeeInGwei float64   `json:"suggestBaseFee"`
	GasUsedRatio         []float64 `json:"gasUsedRatio"`
}
