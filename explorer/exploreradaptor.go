package explorer

import (
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
	"github.com/dapplink-labs/chain-explorer-api/common/token"
	"github.com/dapplink-labs/chain-explorer-api/common/transaction"
)

type ChainExplorerAdaptor interface {
	// GetChainExplorer account
	GetChainExplorer(req *chain.SupportChainExplorerRequest) (*chain.SupportChainExplorerResponse, error)
	GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error)
	GetMultiAccountBalance(req *account.AccountBalanceRequest) ([]*account.AccountBalanceResponse, error)
	GetAccountUtxo(req *account.AccountUtxoRequest) (*account.AccountUtxoResponse, error)
	GetTxByAddress(request *account.AccountTxRequest) (*account.TransactionResponse[account.AccountTxResponse], error)

	// GetTokenList token
	GetTokenList(request *token.TokenRequest) (*token.TokenResponse, error)

	// GetEstimateGasFee gas fee
	GetEstimateGasFee(req *gas_fee.GasEstimateFeeRequest) (*gas_fee.GasEstimateFeeResponse, error)

	// GetTxByHash tx
	GetTxByHash(request *transaction.TxRequest) (*transaction.TxResponse, error)
}
