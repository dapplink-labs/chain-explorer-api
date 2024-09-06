package oklink

import (
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"net/http"
)

type Inscription struct {
	URL string // URL The full address of the OkLink API
}

func (i *Inscription) GetInscriptionTokenList(client *Client, chainShortName, protocolType, tokenInscriptionId, symbol, projectId, startTime, endTime, page string, limit int) (*types.InscriptionTokenListResp, error) {
	i.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/inscription/token-list?chainShortName=%s&protocolType=%s", chainShortName, protocolType)

	if limit != 0 {
		i.URL += fmt.Sprintf("&limit=%d", limit)
	}

	if tokenInscriptionId != "" {
		i.URL += fmt.Sprintf("&tokenInscriptionId=%s", tokenInscriptionId)
	}

	if symbol != "" {
		i.URL += fmt.Sprintf("&symbol=%s", symbol)
	}

	if projectId != "" {
		i.URL += fmt.Sprintf("&projectId=%s", projectId)
	}

	if startTime != "" {
		i.URL += fmt.Sprintf("&startTime=%s", startTime)
	}

	if endTime != "" {
		i.URL += fmt.Sprintf("&endTime=%s", endTime)
	}

	if page != "" {
		i.URL += fmt.Sprintf("&page=%s", page)
	}

	body, err := client.SendRequest(http.MethodGet, i.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.InscriptionTokenListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Inscription) GetInscriptionTokenPositionList(client *Client, chainShortName, protocolType, tokenInscriptionId, symbol, projectId, holderAddress, page string, limit int) (*types.InscriptionTokenPositionListResp, error) {
	i.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/inscription/token-position-list?chainShortName=%s&protocolType=%s", chainShortName, protocolType)

	if tokenInscriptionId != "" {
		i.URL += fmt.Sprintf("&tokenInscriptionId=%s", tokenInscriptionId)
	}

	if symbol != "" {
		i.URL += fmt.Sprintf("&symbol=%s", symbol)
	}

	if projectId != "" {
		i.URL += fmt.Sprintf("&projectId=%s", projectId)
	}

	if holderAddress != "" {
		i.URL += fmt.Sprintf("&holderAddress=%s", holderAddress)
	}

	if page != "" {
		i.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		i.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, i.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.InscriptionTokenPositionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Inscription) GetInscriptionTokenTransactionList(client *Client, chainShortName, protocolType, tokenInscriptionId, symbol, projectId, startTime, endTime, page string, limit int) (*types.InscriptionTokenTransactionListResp, error) {
	i.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/inscription/token-transaction-list?chainShortName=%s&protocolType=%s", chainShortName, protocolType)

	if tokenInscriptionId != "" {
		i.URL += fmt.Sprintf("&tokenInscriptionId=%s", tokenInscriptionId)
	}

	if symbol != "" {
		i.URL += fmt.Sprintf("&symbol=%s", symbol)
	}

	if projectId != "" {
		i.URL += fmt.Sprintf("&projectId=%s", projectId)
	}

	if startTime != "" {
		i.URL += fmt.Sprintf("&startTime=%s", startTime)
	}

	if endTime != "" {
		i.URL += fmt.Sprintf("&endTime=%s", endTime)
	}

	if page != "" {
		i.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		i.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, i.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.InscriptionTokenTransactionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Inscription) GetInscriptionInscriptionList(client *Client, chainShortName, protocolType, tokenInscriptionId, symbol, projectId, page string, limit int) (*types.InscriptionInscriptionListResp, error) {
	i.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/inscription/inscription-list?chainShortName=%s&protocolType=%s", chainShortName, protocolType)

	if tokenInscriptionId != "" {
		i.URL += fmt.Sprintf("&tokenInscriptionId=%s", tokenInscriptionId)
	}

	if symbol != "" {
		i.URL += fmt.Sprintf("&symbol=%s", symbol)
	}

	if projectId != "" {
		i.URL += fmt.Sprintf("&projectId=%s", projectId)
	}

	if page != "" {
		i.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		i.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, i.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.InscriptionInscriptionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Inscription) GetInscriptionAddressTokenList(client *Client, chainShortName, address, protocolType, tokenInscriptionId, symbol, projectId, page string, limit int) (*types.InscriptionAddressTokenListResp, error) {
	i.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/inscription/address-token-list?chainShortName=%s&address=%s&protocolType=%s", chainShortName, address, protocolType)

	if tokenInscriptionId != "" {
		i.URL += fmt.Sprintf("&tokenInscriptionId=%s", tokenInscriptionId)
	}

	if symbol != "" {
		i.URL += fmt.Sprintf("&symbol=%s", symbol)
	}

	if projectId != "" {
		i.URL += fmt.Sprintf("&projectId=%s", projectId)
	}

	if page != "" {
		i.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		i.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, i.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.InscriptionAddressTokenListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Inscription) GetInscriptionAddressInscriptionList(client *Client, chainShortName, address, protocolType, tokenInscriptionId, symbol, projectId, page string, limit int) (*types.InscriptionAddressInscriptionListResp, error) {
	i.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/inscription/address-inscription-list?chainShortName=%s&address=%s&protocolType=%s", chainShortName, address, protocolType)

	if tokenInscriptionId != "" {
		i.URL += fmt.Sprintf("&tokenInscriptionId=%s", tokenInscriptionId)
	}

	if symbol != "" {
		i.URL += fmt.Sprintf("&symbol=%s", symbol)
	}

	if projectId != "" {
		i.URL += fmt.Sprintf("&projectId=%s", projectId)
	}

	if page != "" {
		i.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		i.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, i.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.InscriptionAddressInscriptionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Inscription) GetInscriptionAddressTokenTransactionList(client *Client, chainShortName, address, protocolType, tokenInscriptionId, symbol, projectId, startTime, endTime, page string, limit int) (*types.InscriptionAddressTokenTransactionListResp, error) {
	i.URL = fmt.Sprintf(client.baseURL+"api/v5/explorer/inscription/address-token-transaction-list?chainShortName=%s&protocolType=%s&address=%s", chainShortName, protocolType, address)

	if tokenInscriptionId != "" {
		i.URL += fmt.Sprintf("&tokenInscriptionId=%s", tokenInscriptionId)
	}

	if symbol != "" {
		i.URL += fmt.Sprintf("&symbol=%s", symbol)
	}

	if projectId != "" {
		i.URL += fmt.Sprintf("&projectId=%s", projectId)
	}

	if startTime != "" {
		i.URL += fmt.Sprintf("&startTime=%s", startTime)
	}

	if endTime != "" {
		i.URL += fmt.Sprintf("&endTime=%s", endTime)
	}

	if page != "" {
		i.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		i.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, i.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.InscriptionAddressTokenTransactionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Inscription) GetInscriptionTransactionDetail(client *Client, chainShortName, txId, protocolType, page string, limit int) (*types.InscriptionTransactionDetailResp, error) {
	i.URL = fmt.Sprintf(client.baseURL+"api/v5/explorer/inscription/transaction-detail?chainShortName=%s&protocolType=%s&txId=%s", chainShortName, protocolType, txId)

	if page != "" {
		i.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		i.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, i.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.InscriptionTransactionDetailResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (i *Inscription) GetInscriptionBlockTokenTransaction(client *Client, chainShortName, height, protocolType, txnStartIndex, txnEndIndex, page string, limit int) (*types.InscriptionBlockTokenTransactionResp, error) {
	i.URL = fmt.Sprintf(client.baseURL+"api/v5/explorer/inscription/block-token-transaction?chainShortName=%s&protocolType=%s&height=%s", chainShortName, protocolType, height)

	if txnStartIndex != "" {
		i.URL += fmt.Sprintf("&txnStartIndex=%s", txnStartIndex)
	}

	if txnEndIndex != "" {
		i.URL += fmt.Sprintf("&txnEndIndex=%s", txnEndIndex)
	}

	if page != "" {
		i.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		i.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, i.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.InscriptionBlockTokenTransactionResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
