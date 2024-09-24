package explorer

import (
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
	"github.com/dapplink-labs/chain-explorer-api/common/token"
	"github.com/dapplink-labs/chain-explorer-api/common/transaction"
)

type ChainExplorerAdaptor interface {
	// GetChainExplorer 获取链信息
	GetChainExplorer(req *chain.SupportChainExplorerRequest) (*chain.SupportChainExplorerResponse, error)
	// GetAccountBalance 获取账户余额
	GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error)
	// GetMultiAccountBalance 获取多个账户余额
	GetMultiAccountBalance(req *account.AccountBalanceRequest) ([]*account.AccountBalanceResponse, error)
	// GetAccountUtxo 获取账户utxo
	GetAccountUtxo(req *account.AccountUtxoRequest) ([]account.AccountUtxoResponse, error)
	// GetTxByAddress 根据地址获取交易信息
	GetTxByAddress(request *account.AccountTxRequest) (*account.TransactionResponse[account.AccountTxResponse], error)
	// GetTokenList 获取代币列表
	GetTokenList(request *token.TokenRequest) ([]token.TokenResponse, error)
	// GetEstimateGasFee 获取预估手续费
	GetEstimateGasFee(req *gas_fee.GasEstimateFeeRequest) (*gas_fee.GasEstimateFeeResponse, error)
	// GetTxByHash 根据交易hash获取交易信息
	GetTxByHash(request *transaction.TxRequest) (*transaction.TxResponse, error)
}
