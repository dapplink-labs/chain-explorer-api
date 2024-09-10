package oklink

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/transaction"
)

func (cea *ChainExplorerAdaptor) GetTxByHash(request *transaction.TxRequest) (*transaction.TxResponse, error) {
	apiUrl := fmt.Sprintf("api/v5/explorer/transaction/transaction-fills?chainShortName=%s&txid=%s", request.ChainShortName, request.Txid)

	resp := &transaction.TxResponse{}
	err := cea.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
