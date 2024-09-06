package oklink

import (
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"net/http"
)

type Block struct {
	URL string // URL The full address of the OkLink API
}

func (b *Block) GetBlockAddressBalanceHistory(client *Client, chainShortName, height, address, tokenContractAddress string) (*types.BlockAddressBalanceHistoryResp, error) {
	b.URL = fmt.Sprintf(client.baseURL+"api/v5/explorer/block/address-balance-history?chainShortName=%s&height=%s&address=%s", chainShortName, height, address)
	if tokenContractAddress != "" {
		b.URL += fmt.Sprintf("&tokenContractAddress=%s", tokenContractAddress)
	}

	body, err := client.SendRequest(http.MethodGet, b.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.BlockAddressBalanceHistoryResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
