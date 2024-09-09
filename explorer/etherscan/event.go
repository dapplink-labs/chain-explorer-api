package etherscan

import (
	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/event"
)

func (cea *ChainExplorerAdaptor) GetLogs(req *event.EventRequest) (*event.EventResponse, error) {
	param := common.M{
		"address":   req.Address,
		"topic0":    req.Topic0,
		"fromBlock": req.FromBlock,
		"toBlock":   req.ToBlock,
	}

	var logs []event.Log
	err := cea.baseClient.Call("logs", "getlogs", "GET", param, &logs)
	if err != nil {
		return nil, err
	}

	return &event.EventResponse{
		Logs: logs,
	}, nil
}
