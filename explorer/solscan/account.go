package solscan

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"math/big"
	"strconv"
)

// GetAccountBalance 获取账户余额
func (cea *ChainExplorerAdaptor) GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error) {
	var responseData []AddressSummaryData
	cea.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/tokens", common.M{
		"account": req.Account[0],
	}, &responseData)

	return &account.AccountBalanceResponse{
		Account:         req.Account[0],                                                     // 账户地址
		Balance:         (*common.BigInt)(big.NewInt(responseData[0].TokenAmount.UiAmount)), // 余额
		BalanceStr:      responseData[0].TokenAmount.UiAmountString,                         // 余额字符串
		Symbol:          responseData[0].TokenSymbol,                                        // 代币符号
		ContractAddress: responseData[0].TokenAddress,                                       // 代币合约地址
		TokenId:         responseData[0].TokenAccount,                                       // 代币ID
	}, nil
}

// GetMultiAccountBalance 获取多个账户余额
func (cea *ChainExplorerAdaptor) GetMultiAccountBalance(req *account.AccountBalanceRequest) ([]*account.AccountBalanceResponse, error) {

	var responseData []AddressSummaryData
	cea.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/tokens", common.M{
		"account": req.Account[0],
	}, &responseData)

	balanceList := make([]*account.AccountBalanceResponse, 0, len(req.Account))

	for _, balance := range responseData {
		balanceList = append(balanceList, &account.AccountBalanceResponse{
			Account:         req.Account[0],                                             // 账户地址
			Balance:         (*common.BigInt)(big.NewInt(balance.TokenAmount.UiAmount)), // 余额
			BalanceStr:      balance.TokenAmount.UiAmountString,                         // 余额字符串
			Symbol:          balance.TokenSymbol,                                        // 代币符号
			ContractAddress: balance.TokenAddress,                                       // 代币合约地址
			TokenId:         balance.TokenAccount,                                       // 代币ID
		})
	}

	return balanceList, nil
}

func (cea *ChainExplorerAdaptor) GetTxByAddress(request *account.AccountTxRequest) (*account.TransactionResponse[account.AccountTxResponse], error) {
	if request.Limit > 50 {
		return nil, fmt.Errorf("limit must be less than or equal to 50")
	}
	params := common.M{
		"offset":  request.Page,
		"limit":   request.Limit,
		"account": request.Address,
	}
	if request.Action == account.SolScanActionSol {
		resp := &AddressSolTransactionResp{}
		err := cea.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/solTransfers", params, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{
				PageResponse: chain.PageResponse{
					Page:      request.Page,
					Limit:     request.Limit,
					TotalPage: 0,
				},
				TransactionList: []account.AccountTxResponse{},
			}, nil
		}

		transactionList := make([]account.AccountTxResponse, 0, len(resp.Data))

		for _, tx := range resp.Data {
			var operation string
			if tx.Src == request.Address {
				operation = "dec"
			} else {
				operation = "inc"
			}
			transactionList = append(transactionList, account.AccountTxResponse{
				TxId:         tx.TxHash,
				From:         tx.Src,
				To:           tx.Dst,
				Height:       strconv.FormatInt(tx.Slot, 10),
				Amount:       strconv.FormatInt(tx.Lamport, 10),
				TokenDecimal: strconv.FormatInt(tx.Decimals, 10),
				State:        tx.Status,
				TxFee:        strconv.FormatInt(tx.Fee, 10),
				Operation:    operation,
			})
		}
		return &account.TransactionResponse[account.AccountTxResponse]{
			PageResponse: chain.PageResponse{
				Page:      request.Page,
				Limit:     request.Limit,
				TotalPage: resp.Total,
			},
			TransactionList: transactionList,
		}, nil
	} else {
		resp := &AddressSplTransactionResp{}
		err := cea.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/splTransfers", params, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{
				PageResponse: chain.PageResponse{
					Page:      request.Page,
					Limit:     request.Limit,
					TotalPage: 0,
				},
				TransactionList: []account.AccountTxResponse{},
			}, nil
		}

		transactionList := make([]account.AccountTxResponse, 0, len(resp.Data))

		for _, tx := range resp.Data {
			transactionList = append(transactionList, account.AccountTxResponse{
				Amount:       tx.ChangeAmount,
				Height:       strconv.FormatInt(tx.Slot, 10),
				TokenDecimal: strconv.FormatInt(tx.Decimals, 10),
				TxFee:        strconv.FormatInt(tx.Fee, 10),
				TokenId:      tx.TokenAddress,
				TokenSymbol:  tx.Symbol,
				TokenName:    tx.TokenName,
				Operation:    tx.ChangeType,
			})
		}
		return &account.TransactionResponse[account.AccountTxResponse]{
			PageResponse: chain.PageResponse{
				Page:      request.Page,
				Limit:     request.Limit,
				TotalPage: resp.Total,
			},
			TransactionList: transactionList,
		}, nil
	}
}
