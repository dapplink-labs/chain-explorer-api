package chain

type SupportChainExplorerRequest struct {
	Name string `json:"name"`
}

type SupportChainExplorerResponse struct {
	Ok bool `json:"ok"`
}

type PageRequest struct {
	Page  uint64 `json:"page"`
	Limit uint64 `json:"limit"`
}

type PageResponse struct {
	Page      uint64 `json:"page"`
	Limit     uint64 `json:"limit"`
	TotalPage uint64 `json:"totalPage"`
}

type SortType string

const (
	SortTypeAsc  SortType = "asc"
	SortTypeDesc SortType = "desc"
)
