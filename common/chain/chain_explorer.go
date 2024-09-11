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
	Page      uint64 `json:"page,string"`
	Limit     uint64 `json:"limit,string"`
	TotalPage uint64 `json:"totalPage,string"`
}

type SortType string

const (
	SortTypeAsc  SortType = "asc"
	SortTypeDesc SortType = "desc"
)
