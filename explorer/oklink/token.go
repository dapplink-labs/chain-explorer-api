package oklink

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common/token"
)

// GET /api/v5/explorer/token/token-list
func (cea *ChainExplorerAdaptor) GetTokenList(request *token.TokenRequest) ([]token.TokenResponse, error) {
	var responseList []token.TokenResponse
	_protocolType := request.ProtocolType
	if token.IsBTCEcosystemProtocol(_protocolType) {
		apiUrl := fmt.Sprintf("api/v5/explorer/inscription/token-list?chainShortName=%s&protocolType=%s&tokenInscriptionId=%s&symbol=%s&projectId=%s&page=%s&limit=%s",
			request.ChainShortName, _protocolType, request.TokenInscriptionId, request.Symbol, request.ProjectId, request.Page, request.Limit)
		var responseData []TokenListData
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		for _, tokenValue := range responseData[0].TokenList {
			responseItem := token.TokenResponse{
				Symbol:               tokenValue.Symbol,
				TokenContractAddress: tokenValue.TokenContractAddress,
				TokenId:              tokenValue.TokenInscriptionId,
				TotalSupply:          tokenValue.TotalSupply,
				Decimal:              tokenValue.Precision,
			}
			responseList = append(responseList, responseItem)
		}
	} else {
		apiUrl := fmt.Sprintf("api/v5/explorer/token/token-list?chainShortName=%s&protocolType=%s&tokenContractAddress=%s&page=%s&limit=%s",
			request.ChainShortName, _protocolType, request.ContractAddress, request.Page, request.Limit)
		var responseData []TokenListData
		err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, &responseData)
		if err != nil {
			return nil, err
		}
		for _, tokenValue := range responseData[0].TokenList {
			responseItem := token.TokenResponse{
				Symbol:               tokenValue.Token,
				TokenContractAddress: tokenValue.TokenContractAddress,
				TokenId:              tokenValue.TokenId,
				TotalSupply:          tokenValue.TotalSupply,
				Decimal:              tokenValue.Precision,
			}
			responseList = append(responseList, responseItem)
		}
	}
	return responseList, nil
}

// GET /api/v5/explorer/nft/nft-details
func (cea *ChainExplorerAdaptor) GetNFTDetails(request *token.GetNFTDetailsRequest) (*token.GetNFTDetailsResponse, error) {
	var response []*NFTDetailsData
	apiUrl := fmt.Sprintf("api/v5/explorer/nft/nft-details?chainShortName=%s&tokenContractAddress=%s&tokenId=%s",
		request.ChainShortName, request.ContractAddress, request.TokenId)
	err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, &response)
	if err != nil {
		return nil, err
	}
	if len(response) > 0 {
		data := response[0]
		attributes := make([]token.Attributes, len(data.Attributes))
		for i, attr := range data.Attributes {
			attributes[i] = token.Attributes{
				TraitType:  attr.TraitType,
				Value:      attr.Value,
				Prevalence: attr.Prevalence,
			}
		}
		result := token.GetNFTDetailsResponse{
			CollectionName:       data.CollectionName,
			TokenContractAddress: data.TokenContractAddress,
			TokenId:              data.TokenId,
			ProtocolType:         data.ProtocolType,
			Token:                data.Token,
			OwnerAddress:         data.OwnerAddress,
			LogoUrl:              data.LogoUrl,
			LastPrice:            data.LastPrice,
			FloorPrice:           data.FloorPrice,
			LastPriceUnit:        data.LastPriceUnit,
			LastTransactionTime:  data.LastTransactionTime,
			LastHeight:           data.LastHeight,
			LastTxid:             data.LastTxid,
			TransactionCount:     data.TransactionCount,
			MinterAddress:        data.MinterAddress,
			StorageMethod:        data.StorageMethod,
			MintTime:             data.MintTime,
			Title:                data.Title,
			Attributes:           attributes,
		}
		return &result, nil
	}
	return nil, nil
}
