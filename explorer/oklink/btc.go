package oklink

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/btc"
)

const ChainExplorerModelBtc = "btc"

func (cea *ChainExplorerAdaptor) GetBtcInscriptionsList(req *btc.InscriptionsListRequest) (*btc.InscriptionsListResponse, error) {
	reqMap := map[string]interface{}{
		"token":              req.Token,
		"inscription_id":     req.InscriptionId,
		"inscription_number": req.InscriptionNumber,
		"state":              req.State,
		"page":               req.Page,
		"limit":              req.Limit,
	}
	resp := &btc.InscriptionsListResponse{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/btc/inscriptions-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetBtcTokenList(req *btc.TokenListRequest) (*btc.TokenListResp, error) {
	reqMap := map[string]interface{}{
		"token":   req.Token,
		"orderBy": req.OrderBy,
		"page":    req.Page,
		"limit":   req.Limit,
	}
	resp := &btc.TokenListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/btc/token-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetBtcTokenDetails(req *btc.TokenDetailsRequest) (*btc.TokenDetailsResp, error) {
	reqMap := map[string]interface{}{
		"token": req.Token,
	}
	resp := &btc.TokenDetailsResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/btc/token-details", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetBtcPositionList(req *btc.PositionListRequest) (*btc.PositionListResp, error) {
	reqMap := map[string]interface{}{
		"token": req.Token,
		"page":  req.Page,
		"limit": req.Limit,
	}
	resp := &btc.PositionListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/btc/position-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetBtcTransactionList(req *btc.TransactionListRequest) (*btc.TransactionListResp, error) {
	reqMap := map[string]interface{}{
		"address":           req.Address,
		"token":             req.Token,
		"inscriptionNumber": req.InscriptionNumber,
		"actionType":        req.ActionType,
		"toAddress":         req.ToAddress,
		"fromAddress":       req.FromAddress,
		"txId":              req.TxId,
		"blockHeight":       req.BlockHeight,
		"page":              req.Page,
		"limit":             req.Limit,
	}
	resp := &btc.TransactionListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/btc/transaction-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetBtcAddressBalanceList(req *btc.AddressBalanceListRequest) (*btc.AddressBalanceListResp, error) {
	reqMap := map[string]interface{}{
		"address": req.Address,
		"token":   req.Token,
		"page":    req.Page,
		"limit":   req.Limit,
	}
	resp := &btc.AddressBalanceListResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/btc/address-balance-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}

func (cea *ChainExplorerAdaptor) GetBtcAddressBalanceDetails(req *btc.AddressBalanceDetailsRequest) (*btc.AddressBalanceDetailsResp, error) {
	reqMap := map[string]interface{}{
		"address": req.Address,
		"token":   req.Token,
		"page":    req.Page,
		"limit":   req.Limit,
	}
	resp := &btc.AddressBalanceDetailsResp{}

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModelBtc,
		"/api/v5/explorer/btc/address-balance-details", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}
