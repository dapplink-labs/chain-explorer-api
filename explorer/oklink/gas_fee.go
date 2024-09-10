package oklink

import (
	"fmt"
	"github.com/dapplink-labs/chain-explorer-api/common/gas_fee"
)

func (cea *ChainExplorerAdaptor) GetEstimateGasFee(request *gas_fee.GasEstimateFeeRequest) (*gas_fee.GasEstimateFeeResponse, error) {
	apiUrl := fmt.Sprintf("api/v5/explorer/blockchain/fee?chainShortName=%s", request.ChainShortName)
	resp := &gas_fee.GasEstimateFeeResponse{}

	err := cea.baseClient.Call("oklink", "", "", apiUrl, nil, resp)

	if err != nil {
		return nil, err
	}
	return resp, nil
}
