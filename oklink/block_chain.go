package oklink

import (
	"encoding/json"
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"net/http"
)

type BlockChain struct {
	URL string // URL The full address of the OkLink API
}

func (b *BlockChain) GetGas(client *Client, chainShortName string) (*types.GetGasResp, error) {
	b.URL = fmt.Sprintf(client.baseURL+"api/v5/explorer/blockchain/fee?chainShortName=%s", chainShortName)

	body, err := client.SendRequest(http.MethodGet, b.URL, nil)
	if err != nil {
		return nil, err
	}

	resp := &types.GetGasResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
