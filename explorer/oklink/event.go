package oklink

import "github.com/dapplink-labs/chain-explorer-api/common/event"

func (cea *ChainExplorerAdaptor) GetLogs(req *event.EventRequest) (*event.EventResponse, error) {
	reqMap := map[string]interface{}{
		// ?????
		"chainShortName": req.ExplorerName,
		"address":        req.Address,
		"topic0":         req.Topic0,
		"fromBlock":      req.FromBlock,
		"toBlock":        req.ToBlock,
	}

	resp := &event.EventResponse{}

	err := cea.baseClient.Call(ChainExplorerName, "event",
		"/api/v5/explorer/event/logs", reqMap, resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
