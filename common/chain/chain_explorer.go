package chain

type SupportChainExplorerRequest struct {
	Name string `json:"name"`
}

type SupportChainExplorerResponse struct {
	Ok bool `json:"ok"`
}

type PageRequest struct {
	Page  string `json:"page"`
	Limit string `json:"limit"`
}

type PageResponse struct {
	Page      string `json:"page"`
	Limit     string `json:"limit"`
	TotalPage string `json:"totalPage"`
}
