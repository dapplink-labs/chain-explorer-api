package oklink

import (
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common/account"
)

func (cea *ChainExplorerAdaptor) GetAccountUtxo(request *account.AccountUtxoRequest) ([]account.AccountUtxoResponse, error) {
	var aurList []account.AccountUtxoResponse
	apiUrl := fmt.Sprintf("api/v5/explorer/address/utxo?chainShortName=%s&address=%s", request.ChainShortName, request.Address)
	if request.Page != "" {
		apiUrl += fmt.Sprintf("&page=%s", request.Page)
	}
	if request.Limit != "" {
		apiUrl += fmt.Sprintf("&limit=%d", request.Limit)
	}
	var responseData []AddressUtxoData

	err := cea.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, &responseData)
	if err != nil {
		return nil, err
	}
	for _, responseValue := range responseData {
		for _, utxoItem := range responseValue.UtxoList {
			aur := account.AccountUtxoResponse{
				TxId:          utxoItem.TxId,
				Height:        utxoItem.Height,
				BlockTime:     utxoItem.BlockTime,
				Address:       utxoItem.Address,
				UnspentAmount: utxoItem.UnspentAmount,
				Index:         utxoItem.Index,
			}
			aurList = append(aurList, aur)
		}
	}
	return aurList, nil
}
