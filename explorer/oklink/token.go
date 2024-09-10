package oklink

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common/token"
)

// GET /api/v5/explorer/token/token-list
func (cea *ChainExplorerAdaptor) GetTokenList(request *token.TokenRequest) (*token.TokenResponse, error) {
	var trps *token.TokenResponse
	_protocolType := request.ProtocolType
	if token.IsStandardProtocol(_protocolType) {
		apiUrl := fmt.Sprintf("api/v5/explorer/token/token-list?chainShortName=%s&protocolType=%s&tokenContractAddress=%s&page=%s&limit=%s",
			request.ChainShortName, _protocolType, request.ContractAddress, request.Page, request.Limit)
		response := TokenListData{}
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, response)
		if err != nil {
			return nil, err
		}
		for i, _token := range response.TokenList {
			trps.TokenList[i] = token.TokenInfo{
				Symbol:               _token.Token,
				TokenContractAddress: _token.TokenContractAddress,
				TokenId:              _token.TokenId,
				TotalSupply:          _token.TotalSupply,
				Decimal:              _token.Precision,
			}
		}
	} else if token.IsBTCEcosystemProtocol(_protocolType) {
		apiUrl := fmt.Sprintf("api/v5/explorer/inscription/token-list?chainShortName=%s&protocolType=%s&tokenInscriptionId=%s&symbol=%s&projectId=%s&page=%s&limit=%s",
			request.ChainShortName, _protocolType, request.TokenInscriptionId, request.Symbol, request.ProjectId, request.Page, request.Limit)
		response := TokenListData{}
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, response)
		if err != nil {
			return nil, err
		}
		for i, _token := range response.TokenList {
			trps.TokenList[i] = token.TokenInfo{
				Symbol:               _token.Symbol,
				TokenContractAddress: _token.TokenContractAddress,
				TokenId:              _token.TokenInscriptionId,
				TotalSupply:          _token.TotalSupply,
				Decimal:              _token.Precision,
			}
		}
	}
	return trps, nil
}
