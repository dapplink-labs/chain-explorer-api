package solscan

import (
	"time"

	"github.com/dapplink-labs/chain-explorer-api/explorer"
	"github.com/dapplink-labs/chain-explorer-api/explorer/base"
)

const ChainExplorerName = "solscan"

type ChainExplorerAdaptor struct {
	explorer.ChainExplorerAdaptor
	baseClient *base.BaseClient
}

func NewChainExplorerAdaptor(key string, baseURL string, verbose bool, timeout time.Duration) (*ChainExplorerAdaptor, error) {
	baseClient := base.NewBaseClient(key, baseURL, verbose, timeout)
	return &ChainExplorerAdaptor{
		baseClient: baseClient,
	}, nil
}
