package oklink

import (
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"net/http"
)

type Utxo struct {
	URL string // URL The full address of the OkLink API
}

func (u *Utxo) GetUtxoTransactionStats(client *Client, chainShortName, startTime, endTime, page string, limit int) (*types.UtxoTransactionStatsResp, error) {
	u.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/utxo/transaction-stats?chainShortName=%s", chainShortName)

	if startTime != "" {
		u.URL = fmt.Sprintf(u.URL+"&startTime=%s", startTime)
	}

	if endTime != "" {
		u.URL = fmt.Sprintf(u.URL+"&endTime=%s", endTime)
	}

	if page != "" {
		u.URL = fmt.Sprintf(u.URL+"&page=%s", page)
	}

	if limit != 0 {
		u.URL = fmt.Sprintf(u.URL+"&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, u.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.UtxoTransactionStatsResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (u *Utxo) GetUtxoRevenueStats(client *Client, chainShortName, startTime, endTime, page string, limit int) (*types.UtxoRevenueStatsResp, error) {
	u.URL = fmt.Sprintf(client.baseURL+"/api/v5/explorer/utxo/revenue-stats?chainShortName=%s", chainShortName)

	if startTime != "" {
		u.URL = fmt.Sprintf(u.URL+"&startTime=%s", startTime)
	}

	if endTime != "" {
		u.URL = fmt.Sprintf(u.URL+"&endTime=%s", endTime)
	}

	if page != "" {
		u.URL = fmt.Sprintf(u.URL+"&page=%s", page)
	}

	if limit != 0 {
		u.URL = fmt.Sprintf(u.URL+"&limit=%d", limit)
	}

	body, err := client.SendRequest(http.MethodGet, u.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.UtxoRevenueStatsResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
