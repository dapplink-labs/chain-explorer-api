package account

import "github.com/dapplink-labs/chain-explorer-api/common"

type AccountBalanceRequest struct {
	ExplorerName string `json:"explorer_name"`
	Account      string `json:"account"`
}

type AccountBalanceResponse struct {
	Account string         `json:"account"`
	Balance *common.BigInt `json:"balance"`
}
