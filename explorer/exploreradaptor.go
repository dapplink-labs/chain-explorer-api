package explorer

import (
	"github.com/dapplink-labs/chain-explorer-api/common/account"
	"github.com/dapplink-labs/chain-explorer-api/common/block"
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"github.com/dapplink-labs/chain-explorer-api/common/event"
	"github.com/dapplink-labs/chain-explorer-api/common/gas_tracker"
)

type ChainExplorerAdaptor interface {
	GetChainExplorer(req *chain.SupportChainExplorerRequest) (*chain.SupportChainExplorerResponse, error)
	GetAccountBalance(req *account.AccountBalanceRequest) (*account.AccountBalanceResponse, error)
	GetBlockReward(req *block.BlockRewardRequest) (*block.BlockRewardResponse, error)
	GetBlockNumber(req *block.BlockNumberRequest) (*block.BlockNumberResponse, error)
	GetGasEstimate(req *gas_tracker.GasEstimateRequest) (*gas_tracker.GasEstimateResponse, error)
	GetGasOracle(req *gas_tracker.GasOracleRequest) (*gas_tracker.GasOracleResponse, error)
	GetLogs(req *event.EventRequest) (*event.EventResponse, error)
	GetBlockAddressBalanceHistory(req *block.BlockAddressBalanceHistoryRequest) (*block.BlockAddressBalanceHistoryResponse, error)
	GetBlockByNumber(req *block.BlockByNumberRequest) (*block.BlockByNumberResponse, error)
	GetBlockByHash(req *block.BlockByHashRequest) (*block.BlockByHashResponse, error)
	GetMultiAccountBalance(req *account.AccountBalanceRequest) ([]account.AccountBalanceResponse, error)
	GetTxByAddress(request *account.AccountTxRequest) (*account.AccountTxResponse, error)
}
