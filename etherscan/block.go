package etherscan

import (
	"fmt"
	"strconv"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/etherscan/base"
)

func (c *Client) BlockReward(blockNum int) (rewards base.BlockRewards, err error) {
	param := common.M{
		"blockno": blockNum,
	}
	err = c.call("block", "getblockreward", param, &rewards)
	return
}

func (c *Client) BlockNumber(timestamp int64, closest string) (blockNumber int, err error) {
	var blockNumberStr string
	param := common.M{
		"timestamp": strconv.FormatInt(timestamp, 10),
		"closest":   closest,
	}

	err = c.call("block", "getblocknobytime", param, &blockNumberStr)

	if err != nil {
		return
	}

	blockNumber, err = strconv.Atoi(blockNumberStr)
	if err != nil {
		err = fmt.Errorf("parsing block number %q: %w", blockNumberStr, err)
		return
	}
	return
}
