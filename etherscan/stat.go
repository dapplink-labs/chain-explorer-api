package etherscan

import "github.com/dapplink-labs/chain-explorer-api/etherscan/common"

func (c *Client) EtherTotalSupply() (totalSupply *common.BigInt, err error) {
	err = c.call("stats", "ethsupply", nil, &totalSupply)
	return
}

func (c *Client) EtherLatestPrice() (price common.LatestPrice, err error) {
	err = c.call("stats", "ethprice", nil, &price)
	return
}

func (c *Client) TokenTotalSupply(contractAddress string) (totalSupply *common.BigInt, err error) {
	param := common.M{
		"contractaddress": contractAddress,
	}

	err = c.call("stats", "tokensupply", param, &totalSupply)
	return
}
