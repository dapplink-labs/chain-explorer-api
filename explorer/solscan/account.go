package solscan

import (
	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"math/big"
	"strconv"
)

func (cea *ChainExplorerAdaptor) GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error) {

	var responseData []AddressSummaryData

	cea.baseClient.Call("solscan", "", "", "/v1.0/account/tokens", map[string]any{
		"account": req.Account[0],
	}, &responseData)

	return &account.AccountBalanceResponse{
		Account:         req.Account[0],
		Balance:         (*common.BigInt)(big.NewInt(responseData[0].Amount)),
		BalanceStr:      strconv.FormatInt(responseData[0].Amount, 10),
		Symbol:          responseData[0].TokenAccount,
		ContractAddress: responseData[0].TokenAddress,
		TokenId:         responseData[0].TokenAccount,
	}, nil
}
