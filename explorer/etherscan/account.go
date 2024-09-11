package etherscan

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
)

func (cea *ChainExplorerAdaptor) GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error) {
	balance := new(common.BigInt)
	if req.ContractAddress[0] == "0x00" {
		param := common.M{
			"tag":     "latest",
			"address": req.Account[0],
		}
		err := cea.baseClient.Call("etherscan", "account", "balance", "", param, balance)
		if err != nil {
			fmt.Println("err", err)
			return &account.AccountBalanceResponse{}, nil
		}
	} else {
		param := common.M{
			"contractaddress": req.ContractAddress[0],
			"address":         req.Account[0],
			"tag":             "latest",
		}
		err := cea.baseClient.Call("etherscan", "account", "tokenbalance", "", param, &balance)
		if err != nil {
			fmt.Println("err", err)
		}
	}
	return &account.AccountBalanceResponse{
		Account:         req.Account[0],
		Balance:         balance,
		Symbol:          req.Symbol[0],
		ContractAddress: req.ContractAddress[0],
		TokenId:         "0x0",
	}, nil
}

func (cea *ChainExplorerAdaptor) GetMultiAccountBalance(req *account.AccountBalanceRequest) ([]*account.AccountBalanceResponse, error) {
	var abrList []*account.AccountBalanceResponse
	if req.ContractAddress[0] == "0x00" {
		param := common.M{
			"tag":     "latest",
			"address": req.Account,
		}
		balances := make([]AccountBalance, 0, len(req.Account))
		err := cea.baseClient.Call("etherscan", "account", "balancemulti", "", param, &balances)
		if err != nil {
			fmt.Println("err", err)
			return []*account.AccountBalanceResponse{}, nil
		}
		for _, balance := range balances {
			abr := &account.AccountBalanceResponse{
				Account:         balance.Account,
				Balance:         balance.Balance,
				Symbol:          "ETH",
				ContractAddress: "0x0",
				TokenId:         "0x0",
			}
			abrList = append(abrList, abr)
		}
	} else {
		return []*account.AccountBalanceResponse{}, nil
	}
	return abrList, nil
}

// GetAccountUtxo etherscan can not support this function, so return empty struct
func (cea *ChainExplorerAdaptor) GetAccountUtxo(req *account.AccountUtxoRequest) (*account.AccountUtxoResponse, error) {
	return &account.AccountUtxoResponse{}, nil
}

func (cea *ChainExplorerAdaptor) GetTxByAddress(request *account.AccountTxRequest) (*account.TransactionResponse[account.AccountTxResponse], error) {
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

	// normal transaction
	if request.Action == "txlist" {
		resp := &ApiResponse[[]NormalTransaction]{}
		err := cea.baseClient.Call("etherscan", "account", "txlist", "", param, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.AccountTxResponse{}, nil
		}
	}

	// internal transaction
	if request.Action == "txlistinternal" {
		resp := &ApiResponse[[]InternalTransaction]{}
		err := cea.baseClient.Call("etherscan", "account", "txlistinternal", "", param, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.AccountTxResponse{}, nil
		}
	}

	// token transaction
	if request.Action == "tokentx" {
		resp := &ApiResponse[[]TokenErc20Transaction]{}
		err := cea.baseClient.Call("etherscan", "account", "tokentx", "", param, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.AccountTxResponse{}, nil
		}
	}

	// nft transaction
	if request.Action == "tokennfttx" {
		resp := &ApiResponse[[]TokenErc721Transaction]{}
		err := cea.baseClient.Call("etherscan", "account", "tokennfttx", "", param, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.AccountTxResponse{}, nil
		}
	}

	// nft 1155 transaction
	if request.Action == "token1155tx" {
		resp := &ApiResponse[[]TokenErc721Transaction]{}
		err := cea.baseClient.Call("etherscan", "account", "token1155tx", "", param, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.AccountTxResponse{}, nil
		}
	}

	return nil, nil
}
