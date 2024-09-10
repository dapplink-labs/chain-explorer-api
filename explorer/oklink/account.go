package oklink

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/account"
)

// GET /api/v5/explorer/address/address-summary?chainShortName=eth&address=0x85c6627c4ed773cb7c32644b041f58a058b00d30
func (cea *ChainExplorerAdaptor) GetAccountBalance(request *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error) {
	var abrps *account.AccountBalanceResponse
	if request.ContractAddress[0] != "0x00" {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/address-summary?chainShortName=%s&address=%s", request.ChainShortName, request.Account[0])
		response := AddressSummaryResp{}
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, response)
		if err != nil {
			return nil, err
		}
		balance, _ := new(big.Int).SetString(response.Data[0].Balance, 10)
		abrps = &account.AccountBalanceResponse{
			Account:         response.Data[0].Address,
			Balance:         (*common.BigInt)(balance),
			Symbol:          response.Data[0].BalanceSymbol,
			ContractAddress: response.Data[0].CreateContractAddress,
			TokenId:         "0x00",
		}
	} else {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/token-balance?chainShortName=%s&address=%s&protocolType=%s&limit=%d", request.ChainShortName, request.Account[0], request.ProtocolType[0], request.Limit[0])
		response := AddressTokenBalanceResp{}
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, response)
		if err != nil {
			return nil, err
		}
		balance, _ := new(big.Int).SetString(response.Data[0].TokenList[0].HoldingAmount, 10)
		abrps = &account.AccountBalanceResponse{
			Account:         request.Account[0],
			Balance:         (*common.BigInt)(balance),
			Symbol:          response.Data[0].TokenList[0].Symbol,
			ContractAddress: response.Data[0].TokenList[0].TokenContractAddress,
			TokenId:         response.Data[0].TokenList[0].TokenId,
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
	if request.ContractAddress[0] != "0x00" {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/balance-multi?chainShortName=%s&address=%s", request.ChainShortName, result)
		response := AddressSummaryResp{}
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, response)
		if err != nil {
			return nil, err
		}
		for _, value := range response.Data {
			balance, _ := new(big.Int).SetString(value.Balance, 10)
			abrps := &account.AccountBalanceResponse{
				Account:         value.Address,
				Balance:         (*common.BigInt)(balance),
				Symbol:          value.BalanceSymbol,
				ContractAddress: value.CreateContractAddress,
				TokenId:         "0x00",
			}
			abrpsList = append(abrpsList, abrps)
		}
	} else {
		apiUrl := fmt.Sprintf("api/v5/explorer/address/token-balance?chainShortName=%s&address=%s&protocolType=%s&limit=%d", request.ChainShortName, request.Account[0], request.ProtocolType[0], request.Limit[0])
		response := AddressTokenBalanceResp{}
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, response)
		if err != nil {
			return nil, err
		}
		for _, dataValue := range response.Data {
			for _, token := range dataValue.TokenList {
				balance, _ := new(big.Int).SetString(token.HoldingAmount, 10)
				abrps := &account.AccountBalanceResponse{
					Account:         result,
					Balance:         (*common.BigInt)(balance),
					Symbol:          token.Symbol,
					ContractAddress: token.TokenContractAddress,
					TokenId:         token.TokenId,
				}
				abrpsList = append(abrpsList, abrps)
			}
		}
	}
	return abrpsList, nil
}

func (cea *ChainExplorerAdaptor) GetAccountUtxo(req *account.AccountUtxoRequest) (*account.AccountUtxoResponse, error) {
	return &account.AccountUtxoResponse{}, nil
}

func (cea *ChainExplorerAdaptor) GetTxByAddress(request *account.AccountTxRequest) (*account.AccountTxResponse, error) {
	return nil, nil
}
