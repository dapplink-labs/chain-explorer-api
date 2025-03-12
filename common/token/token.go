package token

// type TokenConfig struct {
// 	ChainShortName string `json:"chainShortName"`
// 	ExplorerName   string `json:"explorerName"`
// }

// StandardTokenProtocols 20 tokens: token_20; 721 tokens: token_721; 1155 tokens: token_1155;
// Inscription token protocol type
// Runes: runes
// BRC-20 tokens: brc20
// SRC-20 tokens: src20
// ARC-20 tokens: arc20
// Ordinals NFT: ordinals_nft
type StandardTokenProtocols struct {
	Token20   string
	Token721  string
	Token1155 string
}

type BTCEcosystemProtocols struct {
	Runes       string
	BRC20       string
	SRC20       string
	ARC20       string
	OrdinalsNFT string
}

type Protocols struct {
	StandardTokenProtocols
	BTCEcosystemProtocols
}

var Protocol = Protocols{
	StandardTokenProtocols{
		Token20:   "token_20",
		Token721:  "token_721",
		Token1155: "token_1155",
	},
	BTCEcosystemProtocols{
		Runes:       "runes",
		BRC20:       "brc20",
		SRC20:       "src20",
		ARC20:       "arc20",
		OrdinalsNFT: "ordinals_nft",
	},
}

type TokenRequest struct {
	ChainShortName     string `json:"chainShortName"`
	ExplorerName       string `json:"explorerName"`
	ContractAddress    string `json:"contractAddress"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:"tokenInscriptionId"`
	Symbol             string `json:"symbol"`
	ProjectId          string `json:"projectId"`
	Page               string `json:"page"`
	Limit              string `json:"limit"`
}

type TokenResponse struct {
	Symbol               string `json:"symbol"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenId              string `json:"tokenId"`
	TotalSupply          string `json:"totalSupply"` // TotalSupply
	Decimal              string `json:"decimal"`     //  default 1
}

type GetNFTDetailsRequest struct {
	ExplorerName    string `json:"explorerName"`
	ChainShortName  string `json:"chainShortName"`
	ContractAddress string `json:"contractAddress"`
	TokenId         string `json:"tokenId"`
}

type GetNFTDetailsResponse struct {
	CollectionName       string       `json:"collectionName"`       // Project name (full name)
	TokenContractAddress string       `json:"tokenContractAddress"` // Project contract address
	TokenId              string       `json:"tokenId"`              // NFT token ID
	ProtocolType         string       `json:"protocolType"`         // NFT contract type
	Token                string       `json:"token"`                // Token name
	OwnerAddress         string       `json:"ownerAddress"`         // NFT holder address
	LogoUrl              string       `json:"logoUrl"`              // NFT image URL
	LastPrice            string       `json:"lastPrice"`            // Latest price of the NFT
	FloorPrice           string       `json:"floorPrice"`           // Floor price of the NFT
	LastPriceUnit        string       `json:"lastPriceUnit"`        // Price unit
	LastTransactionTime  string       `json:"lastTransactionTime"`  // Latest transaction time, Unix timestamp in milliseconds, e.g., 1597026383085
	LastHeight           string       `json:"lastHeight"`           // Block height of the latest transaction
	LastTxid             string       `json:"lastTxid"`             // Hash of the latest transaction
	TransactionCount     string       `json:"transactionCount"`     // Number of transactions for this NFT
	MinterAddress        string       `json:"minterAddress"`        // NFT creator address
	StorageMethod        string       `json:"storageMethod"`        // NFT storage method
	MintTime             string       `json:"mintTime"`             // NFT minting time
	Title                string       `json:"title"`                // Name property in NFT metadata, can be ENS or DID
	Attributes           []Attributes `json:"attributes"`           // Attributes
}

type Attributes struct {
	TraitType  string `json:"traitType"`  // Attribute category
	Value      string `json:"value"`      // Attribute value
	Prevalence string `json:"prevalence"` // Attribute rarity, expressed as decimal, 0.1=10%
}

// IsStandardProtocol Determine whether the given protocol type is a standard token protocol
func IsStandardProtocol(protocolType string) bool {
	switch protocolType {
	case Protocol.StandardTokenProtocols.Token20, Protocol.StandardTokenProtocols.Token721, Protocol.StandardTokenProtocols.Token1155:
		return true
	default:
		return false
	}
}

// IsBTCEcosystemProtocol Determine whether the given protocol type is a special token protocol for the Bitcoin ecosystem
func IsBTCEcosystemProtocol(protocolType string) bool {
	switch protocolType {
	case Protocol.BTCEcosystemProtocols.Runes, Protocol.BTCEcosystemProtocols.BRC20, Protocol.BTCEcosystemProtocols.SRC20, Protocol.BTCEcosystemProtocols.ARC20, Protocol.BTCEcosystemProtocols.OrdinalsNFT:
		return true
	default:
		return false
	}
}
