package oklink

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/transaction"
)

func (cea *ChainExplorerAdaptor) GetTxByHash(request *transaction.TxRequest) (*transaction.TxResponse, error) {
	apiUrl := fmt.Sprintf("api/v5/explorer/transaction/transaction-fills?chainShortName=%s&txid=%s", request.ChainShortName, request.Txid)

	fmt.Println("apiUrl===", apiUrl)

	var response []transaction.TxResponse
	err := cea.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, &response)

	if err != nil {
		return nil, err
	}
	return &response[0], nil
}
