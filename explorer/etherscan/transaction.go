package etherscan

import "github.com/dapplink-labs/chain-explorer-api/common/transaction"

func (cea *ChainExplorerAdaptor) GetTxByHash(request *transaction.TxRequest) (*transaction.TxResponse, error) {
	return &transaction.TxResponse{}, nil
}
