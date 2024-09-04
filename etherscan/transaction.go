package etherscan

import (
	"errors"

	"github.com/dapplink-labs/chain-explorer-api/etherscan/common"
)

var ErrPreByzantiumTx = errors.New("pre-byzantium transaction does not support receipt status check")

func (c *Client) ExecutionStatus(txHash string) (status common.ExecutionStatus, err error) {
	param := common.M{
		"txhash": txHash,
	}
	err = c.call("transaction", "getstatus", param, &status)
	return
}

func (c *Client) ReceiptStatus(txHash string) (receiptStatus int, err error) {
	param := common.M{
		"txhash": txHash,
	}

	var rawStatus = struct {
		Status string `json:"status"`
	}{}

	err = c.call("transaction", "gettxreceiptstatus", param, &rawStatus)
	if err != nil {
		return
	}

	switch rawStatus.Status {
	case "0":
		receiptStatus = 0
	case "1":
		receiptStatus = 1
	default:
		receiptStatus = -1
		err = ErrPreByzantiumTx
	}

	return
}
