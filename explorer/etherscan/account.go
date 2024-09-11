package etherscan

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"

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
	resp := &ApiResponse[[]AddressTransactionResp]{}

	// normal transaction
	if request.Action == account.EtherscanActionTxList {
		err := cea.baseClient.Call(ChainExplorerName, "account", "txlist", "", nil, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// internal transaction
	if request.Action == account.EtherscanActionTxListInternal {
		err := cea.baseClient.Call(ChainExplorerName, "account", "txlistinternal", "", nil, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// token transaction
	if request.Action == account.EtherscanActionTokenTx {
		err := cea.baseClient.Call(ChainExplorerName, "account", "tokentx", "", nil, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// nft transaction
	if request.Action == account.EtherscanActionTokenNftTx {
		err := cea.baseClient.Call(ChainExplorerName, "account", "tokennfttx", "", nil, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// nft 1155 transaction
	if request.Action == account.EtherscanActionToken1155Tx {
		err := cea.baseClient.Call(ChainExplorerName, "account", "token1155tx", "", nil, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	pageResponse := chain.PageResponse{
		Page:      request.Page,
		Limit:     request.Limit,
		TotalPage: "",
	}

	var transactionList []account.AccountTxResponse
	for _, tx := range resp.Result {
		tempOkTx := account.AccountTxResponse{
			TxId:                 tx.Hash,
			BlockHash:            tx.BlockHash,
			Height:               tx.BlockNumber,
			TransactionTime:      tx.TimeStamp,
			From:                 tx.From,
			To:                   tx.To,
			TokenContractAddress: tx.ContractAddress,
			TokenId:              tx.TokenID,
			Amount:               tx.Value,
			Symbol:               tx.TokenSymbol,
			IsFromContract:       false,
			IsToContract:         false,
			Operation:            tx.FunctionName,
			MethodId:             tx.MethodId,
			Nonce:                tx.Nonce,
			GasPrice:             tx.GasPrice,
			GasLimit:             tx.Gas,
			GasUsed:              tx.GasUsed,
			TxFee:                "",
			State:                tx.TxReceiptStatus,
			TransactionType:      "",
		}
		transactionList = append(transactionList, tempOkTx)
	}
	return &account.TransactionResponse[account.AccountTxResponse]{
		PageResponse:    pageResponse,
		TransactionList: transactionList,
	}, nil
}
