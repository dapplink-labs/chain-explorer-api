package transaction

type TxRequest struct {
	ExplorerName string `json:"explorerName"`
}

type TxResponse struct {
	ChainName string `json:"chainName"`
}
