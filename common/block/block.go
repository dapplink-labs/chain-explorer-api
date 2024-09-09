package block

import (
	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type BlockAddressBalanceHistoryRequest struct {
	ExplorerName         string `json:"explorer_name"`
	ChainShortName       string `json:"chainShortName"`
	Height               string `json:"height"`
	Address              string `json:"address"`
	TokenContractAddress string `json:"tokenContractAddress"`
}

type BlockAddressBalanceHistoryData struct {
	Address              string `json:"address"`
	Height               string `json:"height"`
	Balance              string `json:"balance"`
	BalanceSymbol        string `json:"balanceSymbol"`
	TokenContractAddress string `json:"tokenContractAddress"`
	BlockTime            string `json:"blockTime"`
}

type BlockAddressBalanceHistoryResponse struct {
	Data []BlockAddressBalanceHistoryData `json:"data"`
}

type BlockRewardRequest struct {
	ExplorerName string `json:"explorer_name"`
	BlockNum     int    `json:"block_num"`
}

type BlockRewards struct {
	BlockNumber int            `json:"blockNumber,string"`
	TimeStamp   common.Time    `json:"timeStamp"`
	BlockMiner  string         `json:"blockMiner"`
	BlockReward *common.BigInt `json:"blockReward"`
	Uncles      []struct {
		Miner         string         `json:"miner"`
		UnclePosition int            `json:"unclePosition,string"`
		BlockReward   *common.BigInt `json:"blockreward"`
	} `json:"uncles"`
	UncleInclusionReward *common.BigInt `json:"uncleInclusionReward"`
}
type BlockRewardResponse struct {
	Rewards BlockRewards `json:"rewards"`
}

type BlockNumberRequest struct {
	ExplorerName string `json:"explorer_name"`
	Timestamp    int64  `json:"timestamp"`
	Closest      string `json:"closest"`
}

type BlockNumberResponse struct {
	BlockNumber int `json:"block_number"`
}

type BlockByNumberRequest struct {
	ExplorerName   string `json:"explorer_name"`
	ChainShortName string `json:"chainShortName"`
	Height         string `json:"height"`
}

type BlockByNumberResponse struct {
	Block *types.Block `json:"block"`
}

type BlockByHashRequest struct {
	ExplorerName   string `json:"explorer_name"`
	ChainShortName string `json:"chainShortName"`
	Hash           string `json:"hash"`
}

type BlockByHashResponse struct {
	Block *types.Block `json:"block"`
}
