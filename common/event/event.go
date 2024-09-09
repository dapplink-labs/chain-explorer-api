package event

type EventRequest struct {
	ExplorerName string `json:"explorer_name"`
	FromBlock    string `json:"fromBlock"`
	ToBlock      string `json:"toBlock"`
	Topic0       string `json:"topic0"`
	Address      string `json:"address"`
}

type Log struct {
	Address         string   `json:"address"`
	Topics          []string `json:"topics"`
	Data            string   `json:"data"`
	BlockNumber     string   `json:"blockNumber"`
	TransactionHash string   `json:"transactionHash"`
	BlockHash       string   `json:"blockHash"`
	LogIndex        string   `json:"logIndex"`
	Removed         bool     `json:"removed"`
}

type EventResponse struct {
	Logs []Log `json:"logs"`
}
