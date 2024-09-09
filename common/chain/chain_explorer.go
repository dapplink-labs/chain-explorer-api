package chain

type SupportChainExplorerRequest struct {
	Name string `json:"name"`
}

type SupportChainExplorerResponse struct {
	Ok bool `json:"ok"`
}
