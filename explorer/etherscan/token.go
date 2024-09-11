package etherscan

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/token"
)

func (cea *ChainExplorerAdaptor) GetTokenList(req *token.TokenRequest) (*token.TokenResponse, error) {
	var trps *token.TokenResponse
	resData := &ApiResponse[[]TokensResp]{}
	tokenList := []TokensResp{}
	param := common.M{
		"contractaddress": req.ContractAddress,
	}
	fmt.Println(param, "GetTokenList param ")

	err := cea.baseClient.Call("etherscan", "token", "tokeninfo", "", param, &resData)
	//err := cea.baseClient.Call("etherscan", "stats", "tokensupply", "", param, &tokenList)
	if err != nil {
		fmt.Println("err 99999999", err)
	}
	for i, _token := range tokenList {
		trps.TokenList[i] = token.TokenInfo{
			Symbol:               _token.Symbol,
			TokenContractAddress: _token.ContractAddress,
			TokenId:              _token.TokenId,
			TotalSupply:          _token.TotalSupply,
			Decimal:              _token.Divisor,
		}
	}
	return trps, nil
}
