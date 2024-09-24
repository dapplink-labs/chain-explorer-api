package solscan

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/transaction"
	"strconv"
)

func (cea *ChainExplorerAdaptor) GetTxByHash(request *transaction.TxRequest) (*transaction.TxResponse, error) {
	apiUrl := fmt.Sprintf("v1.0/transaction/%s", request.Txid)
	var resp Transaction
	err := cea.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, &resp)

	if err != nil {
		return nil, err
	}
	tokenTransferDetails := make([]transaction.TokenTransferDetail, 0, len(resp.ParsedInstruction))
	for _, tx := range resp.ParsedInstruction {
		tokenTransferDetails = append(tokenTransferDetails, transaction.TokenTransferDetail{
			//Index      : ,
			//Token              : ,
			//TokenContractAddress : ,
			//Symbol              : ,
			From: tx.Params.Source,
			To:   tx.Params.Destination,
			//IsFromContract      : ,
			//IsToContract        : ,
			//TokenId             : ,
			Amount: strconv.Itoa(tx.Params.Amount),
		})
	}

	return &transaction.TxResponse{
		Txid:                 resp.TxHash,
		Height:               strconv.Itoa(resp.Slot),
		Amount:               strconv.Itoa(resp.Lamport),
		Txfee:                strconv.Itoa(resp.Fee),
		State:                resp.Status,
		TokenTransferDetails: tokenTransferDetails,
	}, nil
}
