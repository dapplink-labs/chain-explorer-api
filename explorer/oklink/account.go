package oklink

import "github.com/dapplink-labs/chain-explorer-api/common/account"

func (cea *ChainExplorerAdaptor) GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error) {
	return nil, nil
}

func (cea *ChainExplorerAdaptor) GetMultiAccountBalance(req *account.AccountBalanceRequest) ([]account.AccountBalanceResponse, error) {
	return nil, nil
}

func (cea *ChainExplorerAdaptor) GetAccountUtxo(req *account.AccountUtxoRequest) (*account.AccountUtxoResponse, error) {
	return &account.AccountUtxoResponse{}, nil
}

func (cea *ChainExplorerAdaptor) GetTxByAddress(request *account.AccountTxRequest) (*account.AccountTxResponse, error) {
	return nil, nil
}
