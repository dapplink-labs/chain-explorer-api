package etherscan

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
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
		BalanceStr:      balance.Int().String(),
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

func (cea *ChainExplorerAdaptor) GetTxByAddress(request *account.AccountTxRequest) (*account.TransactionResponse[account.AccountTxResponse], error) {
	resp := &[]AddressTransactionResp{}

	tempRequest := AddressTransactionRequest{
		Address:    request.Address,
		StartBlock: request.StartBlockHeight,
		EndBlock:   request.EndBlockHeight,
		Page:       request.Page,
		Offset:     request.Limit,
		Sort:       request.Sort,
	}

	// normal transaction
	if request.Action == account.EtherscanActionTxList {
		err := cea.baseClient.Call(ChainExplorerName, "account", "txlist", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// internal transaction
	if request.Action == account.EtherscanActionTxListInternal {
		err := cea.baseClient.Call(ChainExplorerName, "account", "txlistinternal", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// token transaction
	if request.Action == account.EtherscanActionTokenTx {
		err := cea.baseClient.Call(ChainExplorerName, "account", "tokentx", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// nft transaction
	if request.Action == account.EtherscanActionTokenNftTx {
		err := cea.baseClient.Call(ChainExplorerName, "account", "tokennfttx", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// nft 1155 transaction
	if request.Action == account.EtherscanActionToken1155Tx {
		err := cea.baseClient.Call(ChainExplorerName, "account", "token1155tx", "", tempRequest.ToQueryParamMap(), &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	var transactionList []account.AccountTxResponse
	for _, tx := range *resp {
		tempOkTx := account.AccountTxResponse{
			TxId:                 tx.Hash,
			BlockHash:            tx.BlockHash,
			Height:               tx.BlockNumber,
			TransactionTime:      tx.TimeStamp,
			TransactionIndex:     tx.TransactionIndex,
			From:                 tx.From,
			To:                   tx.To,
			Nonce:                tx.Nonce,
			Amount:               tx.Value,
			Symbol:               tx.TokenSymbol,
			Operation:            tx.Type,
			GasPrice:             tx.GasPrice,
			GasLimit:             tx.Gas,
			GasUsed:              tx.GasUsed,
			TxFee:                tx.CumulativeGasUsed,
			State:                tx.TxReceiptStatus,
			TransactionType:      "",
			Confirmations:        tx.Confirmations,
			IsError:              tx.IsError,
			TraceId:              tx.TraceId,
			Input:                tx.Input,
			MethodId:             tx.MethodId,
			FunctionName:         tx.FunctionName,
			TokenContractAddress: tx.ContractAddress,
			IsFromContract:       false,
			IsToContract:         false,
			TokenId:              tx.TokenID,
			TokenName:            tx.TokenName,
			TokenSymbol:          tx.TokenSymbol,
			TokenDecimal:         tx.TokenDecimal,
			TokenValue:           tx.TokenValue,
		}
		transactionList = append(transactionList, tempOkTx)
	}

	var tempTotal uint64 = request.Page
	if len(transactionList) >= int(request.Limit) {
		tempTotal++
	}
	pageResponse := chain.PageResponse{
		Page:      request.Page,
		Limit:     request.Limit,
		TotalPage: tempTotal,
	}

	if transactionList == nil || len(transactionList) == 0 {
		transactionList = []account.AccountTxResponse{}
	}

	return &account.TransactionResponse[account.AccountTxResponse]{
		PageResponse:    pageResponse,
		TransactionList: transactionList,
	}, nil
}
