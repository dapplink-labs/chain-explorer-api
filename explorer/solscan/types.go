package solscan

import (
	"github.com/dapplink-labs/chain-explorer-api/common/chain"
	"net/url"
	"strconv"
	"time"
)

// AddressSummaryData The Data field within the Response structure of
// Fundamental blockchain data -> Address Data -> Get basic address details

type AddressSummaryData struct {
	TokenAccount string `json:"tokenAccount"`
	TokenAddress string `json:"tokenAddress"`
	TokenAmount  struct {
		Amount         string `json:"amount"`
		Decimals       int64  `json:"decimals"`
		UiAmount       int64  `json:"uiAmount"`
		UiAmountString string `json:"uiAmountString"`
	} `json:"tokenAmount"`
	Decimals    int64   `json:"decimals"`
	RentEpoch   float64 `json:"rentEpoch"`
	Lamports    int64   `json:"lamports"`
	TokenSymbol string  `json:"tokenSymbol"`
	TokenName   string  `json:"tokenName"`
	TokenIcon   string  `json:"tokenIcon"`
}

type AddressSolTransactionResp struct {
	Total uint64 `json:"total"`
	Data  []struct {
		Slot      int64  `json:"slot"`
		BlockTime int64  `json:"blockTime"`
		TxHash    string `json:"txHash"`
		Src       string `json:"src"`
		Decimals  int64  `json:"decimals"`
		Dst       string `json:"dst"`
		Lamport   int64  `json:"lamport"`
		Status    string `json:"status"`
		Fee       int64  `json:"fee"`
	} `json:"data"`
}

type AddressSplTransactionResp struct {
	Total uint64 `json:"total"`
	Data  []struct {
		Slot         int64    `json:"slot"`
		BlockTime    int64    `json:"blockTime"`
		Signature    []string `json:"signature"`
		ChangeType   string   `json:"changeType"`
		ChangeAmount string   `json:"changeAmount"`
		Decimals     int64    `json:"decimals"`
		PostBalance  string   `json:"postBalance"`
		PreBalance   string   `json:"preBalance"`
		TokenAddress string   `json:"tokenAddress"`
		Owner        string   `json:"owner"`
		Fee          int64    `json:"fee"`
		Address      string   `json:"address"`
		Symbol       string   `json:"symbol"`
		TokenName    string   `json:"tokenName"`
	} `json:"data"`
}

type AddressTransactionRequest struct {
	Page       uint64         `json:"page"`
	Offset     uint64         `json:"offset"`
	Address    string         `url:"address"`
	StartBlock uint64         `url:"startblock"`
	EndBlock   uint64         `url:"endblock"`
	Sort       chain.SortType `url:"sort"`
}

func (req AddressTransactionRequest) ToQueryUrl() string {
	values := url.Values{}
	if req.Page != 0 {
		values.Set("page", strconv.FormatUint(req.Page, 10))
	}
	if req.Offset != 0 {
		values.Set("offset", strconv.FormatUint(req.Offset, 10))
	}
	if req.Address != "" {
		values.Set("address", req.Address)
	}
	if req.StartBlock != 0 {
		values.Set("startblock", strconv.FormatUint(req.StartBlock, 10))
	}
	if req.EndBlock != 0 {
		values.Set("endblock", strconv.FormatUint(req.EndBlock, 10))
	}
	if req.Sort != "" {
		values.Set("sort", string(req.Sort))
	}
	return values.Encode()
}

func (req AddressTransactionRequest) ToQueryParamMap() map[string]any {
	result := make(map[string]any)

	if req.Page != 0 {
		result["page"] = strconv.FormatUint(req.Page, 10)
	}
	if req.Offset != 0 {
		result["offset"] = strconv.FormatUint(req.Offset, 10)
	}
	if req.Address != "" {
		result["address"] = req.Address
	}
	if req.StartBlock != 0 {
		result["startblock"] = strconv.FormatUint(req.StartBlock, 10)
	}
	if req.EndBlock != 0 {
		result["endblock"] = strconv.FormatUint(req.EndBlock, 10)
	}
	if req.Sort != "" {
		result["sort"] = string(req.Sort)
	}

	return result
}

type TokenListResp struct {
	Data []struct {
		Address       string `json:"address"`
		CoingeckoInfo *struct {
			CoingeckoRank int `json:"coingeckoRank"`
			MarketCapRank int `json:"marketCapRank"`
			MarketData    struct {
				CurrentPrice                 float64   `json:"currentPrice"`
				Ath                          float64   `json:"ath"`
				AthChangePercentage          float64   `json:"athChangePercentage"`
				AthDate                      time.Time `json:"athDate"`
				Atl                          float64   `json:"atl"`
				AtlChangePercentage          float64   `json:"atlChangePercentage"`
				AtlDate                      time.Time `json:"atlDate"`
				MarketCap                    int64     `json:"marketCap"`
				MarketCapRank                int       `json:"marketCapRank"`
				FullyDilutedValuation        int64     `json:"fullyDilutedValuation"`
				TotalVolume                  int64     `json:"totalVolume"`
				PriceHigh24H                 float64   `json:"priceHigh24h"`
				PriceLow24H                  float64   `json:"priceLow24h"`
				PriceChange24H               float64   `json:"priceChange24h"`
				PriceChangePercentage24H     float64   `json:"priceChangePercentage24h"`
				PriceChangePercentage7D      float64   `json:"priceChangePercentage7d"`
				PriceChangePercentage14D     float64   `json:"priceChangePercentage14d"`
				PriceChangePercentage30D     float64   `json:"priceChangePercentage30d"`
				PriceChangePercentage60D     int       `json:"priceChangePercentage60d"`
				PriceChangePercentage200D    *float64  `json:"priceChangePercentage200d"`
				PriceChangePercentage1Y      *float64  `json:"priceChangePercentage1y"`
				MarketCapChange24H           float64   `json:"marketCapChange24h"`
				MarketCapChangePercentage24H float64   `json:"marketCapChangePercentage24h"`
				TotalSupply                  float64   `json:"totalSupply"`
				MaxSupply                    *int64    `json:"maxSupply"`
				CirculatingSupply            float64   `json:"circulatingSupply"`
				LastUpdated                  time.Time `json:"lastUpdated"`
			} `json:"marketData"`
		} `json:"coingeckoInfo"`
		Decimals   int `json:"decimals"`
		Extensions struct {
			CoingeckoId     string `json:"coingeckoId"`
			SerumV3Usdc     string `json:"serumV3Usdc,omitempty"`
			SerumV3Usdt     string `json:"serumV3Usdt,omitempty"`
			Website         string `json:"website,omitempty"`
			CoinMarketcapId string `json:"coinMarketcapId,omitempty"`
			Discord         string `json:"discord,omitempty"`
			Medium          string `json:"medium,omitempty"`
			Telegram        string `json:"telegram,omitempty"`
			Twitter         string `json:"twitter,omitempty"`
			Description     string `json:"description,omitempty"`
		} `json:"extensions"`
		Holder            int         `json:"holder"`
		Icon              string      `json:"icon"`
		IsViolate         bool        `json:"isViolate"`
		MarketCapFD       float64     `json:"marketCapFD"`
		MarketCapRank     *int        `json:"marketCapRank,omitempty"`
		MintAddress       string      `json:"mintAddress"`
		PriceUst          interface{} `json:"priceUst"`
		SolAlphaVolume    *float64    `json:"solAlphaVolume"`
		Tags              []string    `json:"tags,omitempty"`
		TokenName         string      `json:"tokenName"`
		TokenSymbol       string      `json:"tokenSymbol"`
		Reputation        string      `json:"reputation"`
		Twitter           string      `json:"twitter,omitempty"`
		Website           string      `json:"website,omitempty"`
		OnChainExtensions string      `json:"onChainExtensions"`
		Supply            struct {
			Amount         float64 `json:"amount"`
			UiAmount       float64 `json:"uiAmount"`
			UiAmountString string  `json:"uiAmountString"`
		} `json:"supply"`
		ChainId int    `json:"chainId"`
		Image   string `json:"image,omitempty"`
		Source  string `json:"source,omitempty"`
	} `json:"data"`
	Total int `json:"total"`
}

type Transaction struct {
	BlockTime    int      `json:"blockTime"`
	Slot         int      `json:"slot"`
	TxHash       string   `json:"txHash"`
	Fee          int      `json:"fee"`
	Status       string   `json:"status"`
	Lamport      int      `json:"lamport"`
	Signer       []string `json:"signer"`
	LogMessage   []string `json:"logMessage"`
	InputAccount []struct {
		Account     string `json:"account"`
		Signer      bool   `json:"signer"`
		Writable    bool   `json:"writable"`
		PreBalance  int    `json:"preBalance"`
		PostBalance int    `json:"postBalance"`
	} `json:"inputAccount"`
	RecentBlockhash   string `json:"recentBlockhash"`
	InnerInstructions []struct {
		Index              int `json:"index"`
		ParsedInstructions []struct {
			ProgramId string `json:"programId"`
			Program   string `json:"program"`
			Type      string `json:"type"`
			Name      string `json:"name"`
			Params    struct {
				ExtensionTypes    []string `json:"extensionTypes,omitempty"`
				Mint              string   `json:"mint,omitempty"`
				NewAccount        string   `json:"newAccount,omitempty"`
				Source            string   `json:"source,omitempty"`
				TransferAmountSOL float64  `json:"transferAmount(SOL),omitempty"`
				ProgramOwner      string   `json:"programOwner,omitempty"`
				Account           string   `json:"account,omitempty"`
				Owner             string   `json:"owner,omitempty"`
			} `json:"params"`
		} `json:"parsedInstructions"`
	} `json:"innerInstructions"`
	TokenBalances []struct {
		Account string `json:"account"`
		Amount  struct {
			PostAmount string `json:"postAmount"`
			PreAmount  int    `json:"preAmount"`
		} `json:"amount"`
		Token struct {
			Decimals     int    `json:"decimals"`
			TokenAddress string `json:"tokenAddress"`
			Name         string `json:"name"`
			Symbol       string `json:"symbol"`
			Icon         string `json:"icon"`
		} `json:"token"`
	} `json:"tokenBalances"`
	ParsedInstruction []struct {
		ProgramId  string `json:"programId"`
		Type       string `json:"type"`
		Data       string `json:"data"`
		DataEncode string `json:"dataEncode"`
		Name       string `json:"name"`
		Params     struct {
			Authority         string `json:"authority,omitempty"`
			AssociatedAccount string `json:"associatedAccount,omitempty"`
			TokenAddress      string `json:"tokenAddress,omitempty"`
			TokenProgramId    string `json:"tokenProgramId,omitempty"`
			Source            string `json:"source"`
			Destination       string `json:"destination"`
			Amount            int    `json:"amount"`
		} `json:"params"`
		Program string `json:"program,omitempty"`
	} `json:"parsedInstruction"`
	Confirmations       interface{}   `json:"confirmations"`
	Version             int           `json:"version"`
	TokenTransfers      []interface{} `json:"tokenTransfers"`
	SolTransfers        []interface{} `json:"solTransfers"`
	SerumTransactions   []interface{} `json:"serumTransactions"`
	RaydiumTransactions []interface{} `json:"raydiumTransactions"`
	UnknownTransfers    []interface{} `json:"unknownTransfers"`
}
