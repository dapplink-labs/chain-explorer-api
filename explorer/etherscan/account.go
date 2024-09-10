package etherscan

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/explorer/base"
)

func (cea *ChainExplorerAdaptor) GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error) {
	param := common.M{
		"tag":     "latest",
		"address": req.Account,
	}
	balance := new(common.BigInt)
	err := cea.baseClient.Call("etherscan", "account", "balance", "", param, balance)
	if err != nil {
		fmt.Println("err", err)
	}
	return &account.AccountBalanceResponse{
		Account: req.Account[0],
		Balance: balance,
	}, nil
}

func (cea *ChainExplorerAdaptor) GetMultiAccountBalance(req *account.AccountBalanceRequest) ([]account.AccountBalanceResponse, error) {
	param := common.M{
		"tag":     "latest",
		"address": req.Account,
	}
	balances := make([]account.AccountBalanceResponse, 0, len(req.Account))
	err := cea.baseClient.Call("etherscan", "account", "balancemulti", "", param, &balances)
	if err != nil {
		fmt.Println("err", err)
	}
	return balances, nil
}

// GetAccountUtxo etherscan can not support this function, so return empty struct
func (cea *ChainExplorerAdaptor) GetAccountUtxo(req *account.AccountUtxoRequest) (*account.AccountUtxoResponse, error) {
	return &account.AccountUtxoResponse{}, nil
}

func (cea *ChainExplorerAdaptor) GetTxByAddress(request *account.AccountTxRequest) (*account.AccountTxResponse, error) {
	param := common.M{
		"address": request.Address,
		"page":    request.Page,
		"offset":  request.Offset,
	}
	common.Compose(param, "startblock", request.StartBlock)
	common.Compose(param, "endblock", request.EndBlock)
	if request.Desc {
		param["sort"] = "desc"
	} else {
		param["sort"] = "asc"
	}

	var err error
	// normal transaction
	if request.Action == "txlist" {
		var txs []base.NormalTx
		err = cea.baseClient.Call("etherscan", "account", "txlist", "", param, &txs)
	}

	// internal transaction
	if request.Action == "txlistinternal" {
		var txs []base.InternalTx
		err = cea.baseClient.Call("etherscan", "account", "txlistinternal", "", param, &txs)
	}

	// token transaction
	if request.Action == "tokentx" {
		var txs []base.ERC20Transfer
		err = cea.baseClient.Call("etherscan", "account", "tokentx", "", param, &txs)
	}

	// nft transaction
	if request.Action == "tokennfttx" {
		var txs []base.ERC721Transfer
		err = cea.baseClient.Call("etherscan", "account", "tokennfttx", "", param, &txs)
	}

	// nft 1155 transaction
	if request.Action == "token1155tx" {
		var txs []base.ERC1155Transfer
		err = cea.baseClient.Call("etherscan", "account", "token1155tx", "", param, &txs)
	}

	if err != nil {
		fmt.Println("err", err)
	}

	return nil, nil
}
