package oklink

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common/account"
)

func (cea *ChainExplorerAdaptor) GetAccountUtxo(request *account.AccountUtxoRequest) ([]AddressUtxoData, error) {
	apiUrl := fmt.Sprintf("api/v5/explorer/address/utxo?chainShortName=%s&address=%s", request.ChainShortName, request.Address)
	if request.Page != "" {
		apiUrl += fmt.Sprintf("&page=%s", request.Page)
	}
	if request.Limit != "" {
		apiUrl += fmt.Sprintf("&limit=%d", request.Limit)
	}
	fmt.Println("apiUrl===", apiUrl)

	var responseData []AddressUtxoData

	err := cea.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, &responseData)
	if err != nil {
		return nil, err
	}

	return responseData, nil
}
