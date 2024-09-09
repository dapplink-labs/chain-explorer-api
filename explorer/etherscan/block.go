package etherscan

import (
	"fmt"
	"strconv"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/common/block"
)

func (cea *ChainExplorerAdaptor) BlockReward(req *block.BlockRewardRequest) (*block.BlockRewardResponse, error) {
	param := common.M{
		"blockno": req.BlockNum,
	}

	var rewards block.BlockRewards
	err := cea.baseClient.Call("block", "getblockreward", "GET", param, &rewards)
	if err != nil {
		return nil, err
	}

	return &block.BlockRewardResponse{
		Rewards: rewards,
	}, nil
}

func (cea *ChainExplorerAdaptor) GetBlockNumber(req *block.BlockNumberRequest) (resp *block.BlockNumberResponse, err error) {
	var blockNumberStr string
	param := common.M{
		"timestamp": strconv.FormatInt(req.Timestamp, 10),
		"closest":   req.Closest,
	}

	err = cea.baseClient.Call("block", "getblocknobytime", "GET", param, &blockNumberStr)
	if err != nil {
		return nil, err
	}

	blockNumber, err := strconv.Atoi(blockNumberStr)
	if err != nil {
		return nil, fmt.Errorf("parsing block number %q: %w", blockNumberStr, err)
	}

	return &block.BlockNumberResponse{
		BlockNumber: blockNumber,
	}, nil
}
