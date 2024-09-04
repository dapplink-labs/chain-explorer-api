package etherscan

import (
	"time"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/etherscan/base"
)

func (c *Client) GasEstimate(gasPrice int) (confirmationTimeInSec time.Duration, err error) {
	params := common.M{"gasPrice": gasPrice}
	var confTime string
	err = c.call("gastracker", "gasestimate", params, &confTime)
	if err != nil {
		return
	}
	return time.ParseDuration(confTime + "s")
}

func (c *Client) GasOracle() (gasPrices base.GasPrices, err error) {
	err = c.call("gastracker", "gasoracle", common.M{}, &gasPrices)
	return
}
