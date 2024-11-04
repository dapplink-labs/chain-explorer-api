package oklink

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"github.com/dapplink-labs/chain-explorer-api/explorer/etherscan"
)

// GET /api/v5/explorer/address/address-summary?chainShortName=eth&address=0x85c6627c4ed773cb7c32644b041f58a058b00d30
func (cea *ChainExplorerAdaptor) GetAccountBalance(request *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error) {
	var abrps *account.AccountBalanceResponse
	if request.ContractAddress[0] == "0x00" {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/address-summary?chainShortName=%s&address=%s", request.ChainShortName, request.Account[0])
		var responseData []AddressSummaryData
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			fmt.Println("error is:", err)
			return nil, err
		}
		balance, _ := new(big.Int).SetString(responseData[0].Balance, 10)
		abrps = &account.AccountBalanceResponse{
			Account:         responseData[0].Address,
			Balance:         (*common.BigInt)(balance),
			BalanceStr:      responseData[0].Balance,
			Symbol:          responseData[0].BalanceSymbol,
			ContractAddress: responseData[0].CreateContractAddress,
			TokenId:         "0x00",
		}
	} else {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/token-balance?chainShortName=%s&address=%s&tokenContractAddress=%s&protocolType=%s&limit=%d", request.ChainShortName, request.Account[0], request.ContractAddress[0], request.ProtocolType[0], request.Limit[0])
		fmt.Println(apiUrl)
		var responseData []AddressTokenBalanceData
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		fmt.Println(responseData)
		fmt.Println(responseData[0].TokenList)
		balance, _ := new(big.Int).SetString(responseData[0].TokenList[0].HoldingAmount, 10)
		abrps = &account.AccountBalanceResponse{
			Account:         request.Account[0],
			Balance:         (*common.BigInt)(balance),
			BalanceStr:      responseData[0].TokenList[0].HoldingAmount,
			Symbol:          responseData[0].TokenList[0].Symbol,
			ContractAddress: responseData[0].TokenList[0].TokenContractAddress,
			TokenId:         responseData[0].TokenList[0].TokenId,
		}
	}
	return abrps, nil
}

func (cea *ChainExplorerAdaptor) GetMultiAccountBalance(request *account.AccountBalanceRequest) ([]*account.AccountBalanceResponse, error) {
	var abrpsList []*account.AccountBalanceResponse
	addressStr := make([]string, len(request.Account))
	for i, v := range request.Account {
		addressStr[i] = fmt.Sprintf("%s", v)
	}
	result := strings.Join(addressStr, ",")
	if request.ContractAddress[0] == "0x00" {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/balance-multi?chainShortName=%s&address=%s", request.ChainShortName, result)
		var responseData []AddressBalanceMultiData
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		for _, dataValue := range responseData {
			for _, value := range dataValue.BalanceList {
				balance, _ := new(big.Int).SetString(value.Balance, 10)
				abrps := &account.AccountBalanceResponse{
					Account:         value.Address,
					Balance:         (*common.BigInt)(balance),
					BalanceStr:      value.Balance,
					Symbol:          dataValue.Symbol,
					ContractAddress: "0x00",
					TokenId:         "0x00",
				}
				abrpsList = append(abrpsList, abrps)
			}
		}
	} else {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/token-balance-multi?chainShortName=%s&address=%s&protocolType=%s&page=%s&limit=%s", request.ChainShortName, result, request.ProtocolType[0], request.Page[0], request.Limit[0])
		fmt.Println("apiUrlapiUrl===", apiUrl)
		var responseData []AddressTokenBalanceMultiData
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		for _, dataValue := range responseData {
			for _, token := range dataValue.BalanceList {
				balance, _ := new(big.Int).SetString(token.HoldingAmount, 10)
				abrps := &account.AccountBalanceResponse{
					Account:         token.Address,
					Balance:         (*common.BigInt)(balance),
					BalanceStr:      token.HoldingAmount,
					Symbol:          "unknown",
					ContractAddress: token.TokenContractAddress,
					TokenId:         "0x00",
				}
				abrpsList = append(abrpsList, abrps)
			}
		}
	}
	return abrpsList, nil
}

func (cea *ChainExplorerAdaptor) GetTxByAddress(request *account.AccountTxRequest) (*account.TransactionResponse[account.AccountTxResponse], error) {
	type TransactionType = account.AccountTxResponse
	type TransactionResponseType = etherscan.TransactionResponse[TransactionType]
	var resp []TransactionResponseType
	// utxo
	if request.Action == account.OkLinkActionUtxo {
		baseURL := "/api/v5/explorer/address/transaction-list"
		fullURL := fmt.Sprintf("%s?%s", baseURL, request.ToQueryUrl())
		err := cea.baseClient.Call(ChainExplorerName, "", "", fullURL, nil, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
		// utxo is transactionLists, not transactionList
		if len(resp) > 0 {
			resp[0].TransactionList = resp[0].TransactionLists
		}
	}
	// normal transaction
	if request.Action == account.OkLinkActionNormal {
		baseURL := "/api/v5/explorer/address/normal-transaction-list"
		fullURL := fmt.Sprintf("%s?%s", baseURL, request.ToQueryUrl())

		//resp := &ApiResponse[[]account.TransactionResponse[account.NormalTransaction]]{}
		err := cea.baseClient.Call(ChainExplorerName, "", "", fullURL, nil, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// internal transaction
	if request.Action == account.OkLinkActionInternal {
		baseURL := "/api/v5/explorer/address/internal-transaction-list"
		fullURL := fmt.Sprintf("%s?%s", baseURL, request.ToQueryUrl())

		err := cea.baseClient.Call(ChainExplorerName, "", "", fullURL, nil, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}

	// token transaction
	if request.Action == account.OkLinkActionToken {
		baseURL := "/api/v5/explorer/address/token-transaction-list"
		fullURL := fmt.Sprintf("%s?%s", baseURL, request.ToQueryUrl())

		err := cea.baseClient.Call(ChainExplorerName, "", "", fullURL, nil, &resp)
		if err != nil {
			fmt.Println("err", err)
			return &account.TransactionResponse[account.AccountTxResponse]{}, nil
		}
	}
	if len(resp) == 0 || len(resp[0].TransactionList) == 0 {
		response := account.TransactionResponse[account.AccountTxResponse]{
			PageResponse: chain.PageResponse{
				Page:      request.Page,
				Limit:     request.Limit,
				TotalPage: request.Page,
			},
			TransactionList: []account.AccountTxResponse{},
		}
		return &response, nil
	}

	tempResp := resp[0]
	response := account.TransactionResponse[account.AccountTxResponse]{
		PageResponse: chain.PageResponse{
			Page:      StringToUint64(tempResp.Page, request.Page),
			Limit:     StringToUint64(tempResp.Limit, request.Limit),
			TotalPage: StringToUint64(tempResp.TotalPage, request.Page+1),
		},
		TransactionList: tempResp.TransactionList,
	}
	return &response, nil
}

func StringToUint64(s string, defaultValue uint64) uint64 {
	value, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		fmt.Println("err", err)
		return defaultValue
	}
	return value
}
