package solscan

import (
	"fmt"
	"strconv"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/token"
)

var pageSizes = []string{"10", "20", "30", "40", "60", "100"}

// 判断字符串是否在数组中
func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
func (cea *ChainExplorerAdaptor) GetTokenList(req *token.TokenRequest) ([]token.TokenResponse, error) {
	var responseList []token.TokenResponse
	var responseData TokenListResp

	// 判断req.Limit是否在pageSizes内
	if !contains(pageSizes, req.Limit) {
		return responseList, fmt.Errorf("limit must be one of %v", pageSizes)
	}
	params := common.M{
		"page":      req.Page,
		"page_size": req.Limit,
	}
	err := cea.baseClient.Call(ChainExplorerName, "", "", "/v1.0/account/splTransfers", params, &responseData)
	if err != nil {
		return responseList, err
	}
	for _, t := range responseData.Data {
		responseList = append(responseList, token.TokenResponse{
			Symbol:               t.TokenSymbol,
			TokenContractAddress: t.MintAddress,
			TokenId:              t.Address,
			TotalSupply:          t.Supply.UiAmountString,
			Decimal:              strconv.Itoa(t.Decimals),
		})
	}
	return responseList, nil
}

func (cea *ChainExplorerAdaptor) GetNFTDetails(request *token.GetNFTDetailsRequest) (*token.GetNFTDetailsResponse, error) {
	//todo implement
	return nil, fmt.Errorf("not implemented")
}
