package etherscan

import (
	"strings"

	"github.com/dapplink-labs/chain-explorer-api/etherscan/common"
)

func (c *Client) AccountBalance(address string) (balance *common.BigInt, err error) {
	param := common.M{
		"tag":     "latest",
		"address": address,
	}
	balance = new(common.BigInt)
	err = c.call("account", "balance", param, balance)
	return
}

func (c *Client) MultiAccountBalance(addresses ...string) (balances []common.AccountBalance, err error) {
	param := common.M{
		"tag":     "latest",
		"address": addresses,
	}
	balances = make([]common.AccountBalance, 0, len(addresses))
	err = c.call("account", "balancemulti", param, &balances)
	return
}

func (c *Client) NormalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []common.NormalTx, err error) {
	param := common.M{
		"address": address,
		"page":    page,
		"offset":  offset,
	}
	common.Compose(param, "startblock", startBlock)
	common.Compose(param, "endblock", endBlock)
	if desc {
		param["sort"] = "desc"
	} else {
		param["sort"] = "asc"
	}

	err = c.call("account", "txlist", param, &txs)
	return
}

func (c *Client) InternalTxByAddress(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []common.InternalTx, err error) {
	param := common.M{
		"address": address,
		"page":    page,
		"offset":  offset,
	}
	common.Compose(param, "startblock", startBlock)
	common.Compose(param, "endblock", endBlock)
	if desc {
		param["sort"] = "desc"
	} else {
		param["sort"] = "asc"
	}

	err = c.call("account", "txlistinternal", param, &txs)
	return
}

func (c *Client) ERC20Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []common.ERC20Transfer, err error) {
	param := common.M{
		"page":   page,
		"offset": offset,
	}
	common.Compose(param, "contractaddress", contractAddress)
	common.Compose(param, "address", address)
	common.Compose(param, "startblock", startBlock)
	common.Compose(param, "endblock", endBlock)

	if desc {
		param["sort"] = "desc"
	} else {
		param["sort"] = "asc"
	}

	err = c.call("account", "tokentx", param, &txs)
	return
}

func (c *Client) ERC721Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []common.ERC721Transfer, err error) {
	param := common.M{
		"page":   page,
		"offset": offset,
	}
	common.Compose(param, "contractaddress", contractAddress)
	common.Compose(param, "address", address)
	common.Compose(param, "startblock", startBlock)
	common.Compose(param, "endblock", endBlock)

	if desc {
		param["sort"] = "desc"
	} else {
		param["sort"] = "asc"
	}

	err = c.call("account", "tokennfttx", param, &txs)
	return
}

func (c *Client) ERC1155Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []common.ERC1155Transfer, err error) {
	param := common.M{
		"page":   page,
		"offset": offset,
	}
	common.Compose(param, "contractaddress", contractAddress)
	common.Compose(param, "address", address)
	common.Compose(param, "startblock", startBlock)
	common.Compose(param, "endblock", endBlock)

	if desc {
		param["sort"] = "desc"
	} else {
		param["sort"] = "asc"
	}

	err = c.call("account", "token1155tx", param, &txs)
	return
}

func (c *Client) SwapTransactions(address string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []common.SwapTransaction, err error) {
	param := common.M{
		"page":   page,
		"offset": offset,
	}
	common.Compose(param, "address", address)
	common.Compose(param, "startblock", startBlock)
	common.Compose(param, "endblock", endBlock)
	if desc {
		param["sort"] = "desc"
	} else {
		param["sort"] = "asc"
	}
	err = c.call("account", "txlist", param, &txs)
	var txList []common.SwapTransaction
	for _, tx := range txs {
		if tx.MethodId != "0x" && strings.ContainsAny(tx.FunctionName, "execute") {
			txList = append(txList, tx)
		}
	}
	return txList, nil
}

func (c *Client) BlocksMinedByAddress(address string, page int, offset int) (mined []common.MinedBlock, err error) {
	param := common.M{
		"address":   address,
		"blocktype": "blocks",
		"page":      page,
		"offset":    offset,
	}

	err = c.call("account", "getminedblocks", param, &mined)
	return
}

func (c *Client) UnclesMinedByAddress(address string, page int, offset int) (mined []common.MinedBlock, err error) {
	param := common.M{
		"address":   address,
		"blocktype": "uncles",
		"page":      page,
		"offset":    offset,
	}

	err = c.call("account", "getminedblocks", param, &mined)
	return
}

func (c *Client) TokenBalance(contractAddress, address string) (balance *common.BigInt, err error) {
	param := common.M{
		"contractaddress": contractAddress,
		"address":         address,
		"tag":             "latest",
	}

	err = c.call("account", "tokenbalance", param, &balance)
	return
}
