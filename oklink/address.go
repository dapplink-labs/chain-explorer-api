package oklink

import (
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"net/http"
)

type Address struct {
	URL string // URL The full address of the OkLink API
}

func (a *Address) GetAddressSummary(client *Client, chainShortName, address string) (*types.AddressSummaryResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/address-summary?chainShortName=%s&address=%s", chainShortName, address)

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressSummaryResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressTokenBalance(client *Client, chainShortName, address, protocolType string, limit int) (*types.AddressTokenBalanceResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/token-balance?chainShortName=%s&address=%s&protocolType=%s&limit=%d", chainShortName, address, protocolType, limit)

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressTokenBalanceResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressBalanceFills(client *Client, chainShortName, address string, limit int) (*types.AddressBalanceFillsResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/address-balance-fills?chainShortName=%s&address=%s&limit=%d", chainShortName, address, limit)

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressBalanceFillsResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressTransactionList(client *Client, chainShortName, address, protocolType, tokenContractAddress, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressTransactionListResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/transaction-list?chainShortName=%s&address=%s", chainShortName, address)
	if protocolType != "" {
		a.URL += fmt.Sprintf("&protocolType=%s", protocolType)
	}

	if tokenContractAddress != "" {
		a.URL += fmt.Sprintf("&tokenContractAddress=%s", tokenContractAddress)
	}

	if startBlockHeight != "" {
		a.URL += fmt.Sprintf("&startBlockHeight=%s", startBlockHeight)
	}

	if endBlockHeight != "" {
		a.URL += fmt.Sprintf("&endBlockHeight=%s", endBlockHeight)
	}

	if isFromOrTo != "" {
		a.URL += fmt.Sprintf("&isFromOrTo=%s", isFromOrTo)
	}

	if page != "" {
		a.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		a.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressTransactionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressNormalTransactionList(client *Client, chainShortName, address, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressNormalTransactionListResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/normal-transaction-list?chainShortName=%s&address=%s", chainShortName, address)
	if startBlockHeight != "" {
		a.URL += fmt.Sprintf("&startBlockHeight=%s", startBlockHeight)
	}

	if endBlockHeight != "" {
		a.URL += fmt.Sprintf("&endBlockHeight=%s", endBlockHeight)
	}

	if isFromOrTo != "" {
		a.URL += fmt.Sprintf("&isFromOrTo=%s", isFromOrTo)
	}

	if page != "" {
		a.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		a.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressNormalTransactionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressInternalTransactionList(client *Client, chainShortName, address, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressInternalTransactionListResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/internal-transaction-list?chainShortName=%s&address=%s", chainShortName, address)
	if startBlockHeight != "" {
		a.URL += fmt.Sprintf("&startBlockHeight=%s", startBlockHeight)
	}

	if endBlockHeight != "" {
		a.URL += fmt.Sprintf("&endBlockHeight=%s", endBlockHeight)
	}

	if isFromOrTo != "" {
		a.URL += fmt.Sprintf("&isFromOrTo=%s", isFromOrTo)
	}

	if page != "" {
		a.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		a.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressInternalTransactionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressTokenTransactionList(client *Client, chainShortName, address, protocolType, tokenContractAddress, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressTokenTransactionListResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/token-transaction-list?chainShortName=%s&address=%s", chainShortName, address)
	if protocolType != "" {
		a.URL += fmt.Sprintf("&protocolType=%s", protocolType)
	}

	if tokenContractAddress != "" {
		a.URL += fmt.Sprintf("&tokenContractAddress=%s", tokenContractAddress)
	}

	if startBlockHeight != "" {
		a.URL += fmt.Sprintf("&startBlockHeight=%s", startBlockHeight)
	}

	if endBlockHeight != "" {
		a.URL += fmt.Sprintf("&endBlockHeight=%s", endBlockHeight)
	}

	if isFromOrTo != "" {
		a.URL += fmt.Sprintf("&isFromOrTo=%s", isFromOrTo)
	}

	if page != "" {
		a.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		a.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressTokenTransactionListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressBalanceMulti(client *Client, chainShortName, address string) (*types.AddressBalanceMultiResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/balance-multi?chainShortName=%s&address=%s", chainShortName, address)

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressBalanceMultiResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressTokenBalanceMulti(client *Client, chainShortName, address, protocolType, page string, limit int) (*types.AddressTokenBalanceMultiResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/token-balance-multi?chainShortName=%s&address=%s", chainShortName, address)
	if protocolType != "" {
		a.URL += fmt.Sprintf("&protocolType=%s", protocolType)
	}

	if page != "" {
		a.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		a.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressTokenBalanceMultiResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressNormalTransactionListMulti(client *Client, chainShortName, address, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressNormalTransactionListMultiResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/normal-transaction-list-multi?chainShortName=%s&address=%s&startBlockHeight=%s&endBlockHeight=%s", chainShortName, address, startBlockHeight, endBlockHeight)

	if isFromOrTo != "" {
		a.URL += fmt.Sprintf("&isFromOrTo=%s", isFromOrTo)
	}

	if page != "" {
		a.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		a.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressNormalTransactionListMultiResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressInternalTransactionListMulti(client *Client, chainShortName, address, startBlockHeight, endBlockHeight, page string, limit int) (*types.AddressInternalTransactionListMultiResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/internal-transaction-list-multi?chainShortName=%s&address=%s&startBlockHeight=%s&endBlockHeight=%s", chainShortName, address, startBlockHeight, endBlockHeight)
	if page != "" {
		a.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		a.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressInternalTransactionListMultiResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressTokenTransactionListMulti(client *Client, chainShortName, address, startBlockHeight, endBlockHeight, protocolType, tokenContractAddress, isFromOrTo, page string, limit int) (*types.AddressTokenTransactionListMultiResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/token-transaction-list-multi?chainShortName=%s&address=%s&startBlockHeight=%s&endBlockHeight=%s", chainShortName, address, startBlockHeight, endBlockHeight)
	if protocolType != "" {
		a.URL += fmt.Sprintf("&protocolType=%s", protocolType)
	}

	if tokenContractAddress != "" {
		a.URL += fmt.Sprintf("&tokenContractAddress=%s", tokenContractAddress)
	}

	if isFromOrTo != "" {
		a.URL += fmt.Sprintf("&isFromOrTo=%s", isFromOrTo)
	}

	if page != "" {
		a.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		a.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressTokenTransactionListMultiResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *Address) GetAddressUtxo(client *Client, chainShortName, address, page string, limit int) (*types.AddressUtxoResp, error) {
	a.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/address/utxo?chainShortName=%s&address=%s", chainShortName, address)

	if page != "" {
		a.URL += fmt.Sprintf("&page=%s", page)
	}

	if limit != 0 {
		a.URL += fmt.Sprintf("&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.AddressUtxoResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
