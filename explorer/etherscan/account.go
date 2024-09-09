package etherscan

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
)

func (cea *ChainExplorerAdaptor) GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error) {
	param := common.M{
		"tag":     "latest",
		"address": req.Account,
	}
	balance := new(common.BigInt)
	err := cea.baseClient.Call("etherscan", "account", "balance", param, balance)
	if err != nil {
		fmt.Println("err", err)
	}
	return &account.AccountBalanceResponse{
		Account: req.Account,
		Balance: balance,
	}, nil
}
