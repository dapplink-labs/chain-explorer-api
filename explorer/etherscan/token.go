package etherscan

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/token"
)

func (cea *ChainExplorerAdaptor) GetTokenList(req *token.TokenRequest) ([]token.TokenResponse, error) {
	var tokenList []token.TokenResponse
	var responseData []TokensResp
	param := common.M{
		"contractaddress": req.ContractAddress,
	}
	err := cea.baseClient.Call("etherscan", "token", "tokeninfo", "", param, &responseData)
	if err != nil {
		fmt.Println("call token list for etherscan fail", "err", err)
	}
	for _, tokenValue := range responseData {
		tokenItem := token.TokenResponse{
			Symbol:               tokenValue.Symbol,
			TokenContractAddress: tokenValue.ContractAddress,
			TokenId:              tokenValue.TokenId,
			TotalSupply:          tokenValue.TotalSupply,
			Decimal:              tokenValue.Divisor,
		}
		tokenList = append(tokenList, tokenItem)
	}
	return tokenList, nil
}
