package oklink

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/inscription"
)

const ChainExplorerModelInscription = "inscription"

func (cea *ChainExplorerAdaptor) GetInscriptionTokenList(req *inscription.TokenListRequest) (*inscription.TokenListResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName":     req.ChainShortName,
		"protocolType":       req.ProtocolType,
		"tokenInscriptionId": req.TokenInscriptionId,
		"symbol":             req.Symbol,
		"projectId":          req.ProjectId,
		"startTime":          req.StartTime,
		"endTime":            req.EndTime,
		"page":               req.Page,
		"limit":              req.Limit,
	}
	resp := &inscription.TokenListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelInscription,
		"/api/v5/explorer/inscription/token-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetInscriptionTokenPositionList(req *inscription.TokenPositionListRequest) (*inscription.TokenPositionListResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName":     req.ChainShortName,
		"protocolType":       req.ProtocolType,
		"tokenInscriptionId": req.TokenInscriptionId,
		"symbol":             req.Symbol,
		"projectId":          req.ProjectId,
		"holderAddress":      req.HolderAddress,
		"page":               req.Page,
		"limit":              req.Limit,
	}
	resp := &inscription.TokenPositionListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/inscription/token-position-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetInscriptionTokenTransactionList(req *inscription.TokenTransactionListRequest) (*inscription.TokenTransactionListResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName":     req.ChainShortName,
		"protocolType":       req.ProtocolType,
		"tokenInscriptionId": req.TokenInscriptionId,
		"symbol":             req.Symbol,
		"projectId":          req.ProjectId,
		"startTime":          req.StartTime,
		"endTime":            req.EndTime,
		"page":               req.Page,
		"limit":              req.Limit,
	}
	resp := &inscription.TokenTransactionListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/inscription/token-transaction-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetInscriptionInscriptionList(req *inscription.InscriptionListRequest) (*inscription.InscriptionListResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName":     req.ChainShortName,
		"protocolType":       req.ProtocolType,
		"tokenInscriptionId": req.TokenInscriptionId,
		"symbol":             req.Symbol,
		"projectId":          req.ProjectId,
		"page":               req.Page,
		"limit":              req.Limit,
	}
	resp := &inscription.InscriptionListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/inscription/inscription-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetInscriptionAddressTokenList(req *inscription.AddressTokenListRequest) (*inscription.AddressTokenListResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName":     req.ChainShortName,
		"address":            req.Address,
		"protocolType":       req.ProtocolType,
		"tokenInscriptionId": req.TokenInscriptionId,
		"symbol":             req.Symbol,
		"projectId":          req.ProjectId,
		"page":               req.Page,
		"limit":              req.Limit,
	}
	resp := &inscription.AddressTokenListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/inscription/address-token-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetInscriptionAddressInscriptionList(req *inscription.AddressInscriptionListRequest) (*inscription.AddressInscriptionListResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName":     req.ChainShortName,
		"address":            req.Address,
		"protocolType":       req.ProtocolType,
		"tokenInscriptionId": req.TokenInscriptionId,
		"symbol":             req.Symbol,
		"projectId":          req.ProjectId,
		"page":               req.Page,
		"limit":              req.Limit,
	}
	resp := &inscription.AddressInscriptionListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/inscription/address-inscription-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetInscriptionAddressTokenTransactionList(req *inscription.AddressTokenTransactionListRequest) (*inscription.AddressTokenTransactionListResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName":     req.ChainShortName,
		"address":            req.Address,
		"protocolType":       req.ProtocolType,
		"tokenInscriptionId": req.TokenInscriptionId,
		"symbol":             req.Symbol,
		"projectId":          req.ProjectId,
		"startTime":          req.StartTime,
		"endTime":            req.EndTime,
		"page":               req.Page,
		"limit":              req.Limit,
	}
	resp := &inscription.AddressTokenTransactionListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/inscription/address-token-transaction-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetInscriptionTransactionDetail(req *inscription.TransactionDetailRequest) (*inscription.TransactionDetailResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName": req.ChainShortName,
		"txId":           req.TxId,
		"protocolType":   req.ProtocolType,
		"page":           req.Page,
		"limit":          req.Limit,
	}
	resp := &inscription.TransactionDetailResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/inscription/transaction-detail", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetInscriptionBlockTokenTransaction(req *inscription.BlockTokenTransactionRequest) (*inscription.BlockTokenTransactionResp, error) {
	reqMap := map[string]interface{}{
		"chainShortName": req.ChainShortName,
		"height":         req.Height,
		"protocolType":   req.ProtocolType,
		"txnStartIndex":  req.TxnStartIndex,
		"txnEndIndex":    req.TxnEndIndex,
		"page":           req.Page,
		"limit":          req.Limit,
	}
	resp := &inscription.BlockTokenTransactionResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/inscription/block-token-transaction", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}
