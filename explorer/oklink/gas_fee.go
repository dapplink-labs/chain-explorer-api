package oklink

import (
	"errors"
	"fmt"

	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
)

func (cea *ChainExplorerAdaptor) GetEstimateGasFee(request *gas_fee.GasEstimateFeeRequest) (*gas_fee.GasEstimateFeeResponse, error) {
	apiUrl := fmt.Sprintf("api/v5/explorer/blockchain/fee?chainShortName=%s", request.ChainShortName)
	var resp []gas_fee.GasEstimateFeeResponse

	err := cea.baseClient.Call(ChainExplorerName, "", "", apiUrl, nil, &resp)

	if err != nil {
		return nil, err
	}
	// 检查切片的长度
	if len(resp) == 0 {
		return nil, errors.New("no data returned from API")
	}

	return &resp[0], nil
}
