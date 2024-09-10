package etherscan

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/token"
)

func (cea *ChainExplorerAdaptor) GetTokenList(req *token.TokenRequest) (*token.TokenResponse, error) {
	var trps *token.TokenResponse
	tokenList := []TokensResp{}
	param := common.M{
		"contractaddress": req.ContractAddress,
	}
	err := cea.baseClient.Call("etherscan", "token", "tokeninfo", "", param, &tokenList)
	if err != nil {
		fmt.Println("err", err)
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
