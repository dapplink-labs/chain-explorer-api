package oklink

import (
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"net/http"
)

type Btc struct {
	URL string // URL The full address of the OkLink API
}

func (b *Btc) GetBtcInscriptionsList(client *Client, token, inscriptionId, inscriptionNumber, state, page string, limit int) (*types.BtcInscriptionsListResp, error) {
	b.URL = client.baseURL + "/api/v5/explorer/btc/inscriptions-list"

	if token != "" {
		b.URL += fmt.Sprintf(b.URL+"?token=%s", token)
	}

	if inscriptionId != "" {
		b.URL += fmt.Sprintf(b.URL+"&inscriptionId=%s", inscriptionId)
	}

	if inscriptionNumber != "" {
		b.URL += fmt.Sprintf(b.URL+"&inscriptionNumber=%s", inscriptionNumber)
	}

	if state != "" {
		b.URL += fmt.Sprintf(b.URL+"&state=%s", state)
	}

	if page != "" {
		b.URL += fmt.Sprintf(b.URL+"&page=%s", page)
	}

	if limit != 0 {
		b.URL += fmt.Sprintf(b.URL+"&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, b.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.BtcInscriptionsListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *Btc) GetBtcTokenList(client *Client, token, orderBy, page string, limit int) (*types.BtcTokenListResp, error) {
	b.URL = client.baseURL + "/api/v5/explorer/btc/token-list"

	if token != "" {
		b.URL += fmt.Sprintf(b.URL+"?token=%s", token)
	}

	if orderBy != "" {
		b.URL += fmt.Sprintf(b.URL+"&orderBy=%s", orderBy)
	}

	if page != "" {
		b.URL += fmt.Sprintf(b.URL+"&page=%s", page)
	}

	if limit != 0 {
		b.URL += fmt.Sprintf(b.URL+"&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, b.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.BtcTokenListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *Btc) GetBtcTokenDetails(client *Client, token string) (*types.BtcTokenDetailsResp, error) {
	b.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/btc/token-details?token=%s", token)

	body, err := client.SendRequest(http.MethodGet, b.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.BtcTokenDetailsResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *Btc) GetBtcPositionList(client *Client, token, page string, limit int) (*types.BtcPositionListResp, error) {
	b.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/btc/position-list?token=%s", token)

	if page != "" {
		b.URL += fmt.Sprintf(b.URL+"&page=%s", page)
	}

	if limit != 0 {
		b.URL += fmt.Sprintf(b.URL+"&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, b.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.BtcPositionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *Btc) GetBtcTransactionList(client *Client, address, token, inscriptionNumber, actionType, toAddress, fromAddress, txId, blockHeight, page string, limit int) (*types.BtcTransactionListResp, error) {
	b.URL = client.baseURL + "/api/v5/explorer/btc/transaction-list"

	if address != "" {
		b.URL += fmt.Sprintf(b.URL+"?&address=%s", address)
	}

	if token != "" {
		b.URL += fmt.Sprintf(b.URL+"&token=%s", token)
	}

	if inscriptionNumber != "" {
		b.URL += fmt.Sprintf(b.URL+"&inscriptionNumber=%s", inscriptionNumber)
	}

	if actionType != "" {
		b.URL += fmt.Sprintf(b.URL+"&actionType=%s", actionType)
	}

	if toAddress != "" {
		b.URL += fmt.Sprintf(b.URL+"&toAddress=%s", toAddress)
	}

	if fromAddress != "" {
		b.URL += fmt.Sprintf(b.URL+"&fromAddress=%s", fromAddress)
	}

	if txId != "" {
		b.URL += fmt.Sprintf(b.URL+"&txId=%s", txId)
	}

	if blockHeight != "" {
		b.URL += fmt.Sprintf(b.URL+"&blockHeight=%s", blockHeight)
	}

	if page != "" {
		b.URL += fmt.Sprintf(b.URL+"&page=%s", page)
	}

	if limit != 0 {
		b.URL += fmt.Sprintf(b.URL+"&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, b.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.BtcTransactionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *Btc) GetBtcAddressBalanceList(client *Client, address, token, page string, limit int) (*types.BtcAddressBalanceListResp, error) {
	b.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/btc/address-balance-list?address=%s", address)

	if token != "" {
		b.URL += fmt.Sprintf(b.URL+"&token=%s", token)
	}

	if page != "" {
		b.URL += fmt.Sprintf(b.URL+"&page=%s", page)
	}

	if limit != 0 {
		b.URL += fmt.Sprintf(b.URL+"&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, b.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.BtcAddressBalanceListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (b *Btc) GetBtcAddressBalanceDetails(client *Client, address, token, page string, limit int) (*types.BtcAddressBalanceDetailsResp, error) {
	b.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/btc/address-balance-details?address=%s&token=%s", address, token)

	if page != "" {
		b.URL += fmt.Sprintf(b.URL+"&page=%s", page)
	}

	if limit != 0 {
		b.URL += fmt.Sprintf(b.URL+"&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, b.URL, nil)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%s\n", string(body))

	resp := &types.BtcAddressBalanceDetailsResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
