package oklink

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/btc"
)

const ChainExplorerModel = "btc"

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

	err := cea.baseClient.Call(ChainExplorerName, ChainExplorerModel,
		"/api/v5/explorer/btc/inscriptions-list", reqMap, resp)

	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}
	return resp, nil
}
