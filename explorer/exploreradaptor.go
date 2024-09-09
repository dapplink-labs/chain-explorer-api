package explorer

import (
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
)

type ChainExplorerAdaptor interface {
	GetChainExplorer(req *chain.SupportChainExplorerRequest) (*chain.SupportChainExplorerResponse, error)
	GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error)
	GetMultiAccountBalance(req *account.AccountBalanceRequest) ([]account.AccountBalanceResponse, error)
	GetTxByAddress(request *account.AccountTxRequest) (*account.AccountTxResponse, error)
}
