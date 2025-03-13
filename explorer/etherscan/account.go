package etherscan

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"github.com/dapplink-labs/chain-explorer-api/common/token"
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

func (cea *ChainExplorerAdaptor) GetAccountBalanceV2(request *account.GetAccountBalanceRequest) (*account.GetAccountBalanceResponse, error) {
	var result *account.GetAccountBalanceResponse
	if request.TokenType == "contract" {
		if len(request.Account) > 1 {
			return nil, errors.New("etherscan not support multi account contract token balance")
		} else {
			switch request.ProtocolType {
			case "token_20":
				balance := new(common.BigInt)
				param := common.M{
					"contractaddress": request.ContractAddress,
					"address":         request.Account[0],
					"tag":             "latest",
				}
				err := cea.baseClient.Call("etherscan", "account", "tokenbalance", "", param, &balance)
				if err != nil {
					fmt.Println("err", err)
					return nil, err
				}
				symbol := ""
				tokenInfo, _ := cea.GetTokenList(&token.TokenRequest{
					ContractAddress: request.ContractAddress,
				})
				if len(tokenInfo) > 0 {
					symbol = tokenInfo[0].Symbol
				}
				var balanceItems []account.Balance
				balanceItems = append(balanceItems, account.Balance{
					Account:         request.Account[0],
					Balance:         balance,
					BalanceStr:      balance.Int().String(),
					Symbol:          symbol,
					ContractAddress: request.ContractAddress,
				})
				result = &account.GetAccountBalanceResponse{
					Page:        "1",
					Limit:       "1",
					TotalPage:   "1",
					BalanceList: balanceItems,
				}
			case "token_721":
				var token721List []token.Token721Response
				param := common.M{
					"address": request.Account[0],
					"page":    request.Page,
					"offset":  request.Limit,
				}
				err := cea.baseClient.Call("etherscan", "account", "addresstokennftbalance", "", param, &token721List)
				if err != nil {
					fmt.Println("err", err)
					return nil, err
				}
				var balanceItems []account.Balance
				for _, token721 := range token721List {
					balance, _ := new(big.Int).SetString(token721.TokenQuantity, 10)
					balanceItems = append(balanceItems, account.Balance{
						Account:         request.Account[0],
						Balance:         (*common.BigInt)(balance),
						BalanceStr:      token721.TokenQuantity,
						Symbol:          token721.TokenSymbol,
						ContractAddress: token721.TokenAddress,
					})
				}
				result = &account.GetAccountBalanceResponse{
					Page:        request.Page,
					Limit:       request.Limit,
					TotalPage:   "",
					BalanceList: balanceItems,
				}
			case "token_1155":
				return nil, errors.New("etherscan not support account contract token_1155 balance")
			}
		}
	} else {
		var balanceItems []account.Balance
		if len(request.Account) > 1 {
			param := common.M{
				"tag":     "latest",
				"address": request.Account,
			}
			balances := make([]AccountBalance, 0, len(request.Account))
			err := cea.baseClient.Call("etherscan", "account", "balancemulti", "", param, &balances)
			if err != nil {
				fmt.Println("err", err)
				return nil, err
			}
			if len(balances) > 0 {
				for _, balance := range balances {
					balanceItems = append(balanceItems,
						account.Balance{
							Account:    balance.Account,
							Balance:    balance.Balance,
							BalanceStr: balance.Balance.Int().String(),
							Symbol:     "ETH"})
				}
			} else {
				balanceItems = append(balanceItems,
					account.Balance{
						Account:    "",
						Balance:    (*common.BigInt)(big.NewInt(0)),
						BalanceStr: "0",
						Symbol:     "ETH"})
			}
		} else {
			balance := new(common.BigInt)
			param := common.M{
				"tag":     "latest",
				"address": request.Account[0],
			}
			err := cea.baseClient.Call("etherscan", "account", "balance", "", param, balance)
			if err != nil {
				fmt.Println("err", err)
				return nil, err
			}
			balanceItems = append(balanceItems,
				account.Balance{
					Account:    request.Account[0],
					Balance:    balance,
					BalanceStr: balance.Int().String(),
					Symbol:     "ETH"})
		}
		result = &account.GetAccountBalanceResponse{
			BalanceList: balanceItems,
		}
	}
	return result, nil
}
