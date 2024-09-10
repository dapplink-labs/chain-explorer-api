package token

type TokenRequest struct {
	ChainShortName string `json:"chainShortName"`
	ExplorerName   string `json:"explorerName"`
}

type TokenResponse struct {
}
