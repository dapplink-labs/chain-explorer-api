package etherscan

import "github.com/dapplink-labs/chain-explorer-api/common/account"

func (cea *ChainExplorerAdaptor) GetAccountUtxo(req *account.AccountUtxoRequest) (*account.AccountUtxoResponse, error) {
	return &account.AccountUtxoResponse{}, nil
}
