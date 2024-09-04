package etherscan

import "github.com/dapplink-labs/chain-explorer-api/etherscan/common"

func (c *Client) ContractABI(address string) (abi string, err error) {
	param := common.M{
		"address": address,
	}

	err = c.call("contract", "getabi", param, &abi)
	return
}

func (c *Client) ContractSource(address string) (source []common.ContractSource, err error) {
	param := common.M{
		"address": address,
	}

	err = c.call("contract", "getsourcecode", param, &source)
	return
}
