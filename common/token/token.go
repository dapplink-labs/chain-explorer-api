package token

type TokenConfig struct {
	ChainShortName string `json:"chainShortName"`
	ExplorerName   string `json:"explorerName"`
}

// 20代币：token_20; 721代币：token_721; 1155代币：token_1155;
// 铭文代币协议类型
// Runes符文：runes
// BRC-20代币：brc20
// SRC-20代币：src20
// ARC-20代币：arc20
// Ordinals NFT：ordinals_nft
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
	TokenConfig
	ContractAddress    string `json:"contractAddress"`
	ProtocolType       string `json:"protocolType"`
	TokenInscriptionId string `json:tokenInscriptionId`
	Symbol             string `json:symbol`
	ProjectId          string `json:projectId`
	Page               string `json:"page"`
	Limit              string `json:"limit"`
}

type TokenInfo struct {
	Symbol               string `json:"symbol"`
	TokenContractAddress string `json:"tokenContractAddress"`
	TokenId              string `json:"tokenId"`
	TotalSupply          string `json:"totalSupply"` // 最大供应量
	Decimal              string `json:"decimal"`     // 精度 默认为1
}

type TokenResponse struct {
	TokenList []TokenInfo `json:"tokenInfo"`
}

// IsStandardProtocol 判断给定的协议类型是否为标准token协议
func IsStandardProtocol(protocolType string) bool {
	switch protocolType {
	case Protocol.StandardTokenProtocols.Token20, Protocol.StandardTokenProtocols.Token721, Protocol.StandardTokenProtocols.Token1155:
		return true
	default:
		return false
	}
}

// IsBTCEcosystemProtocol 判断给定的协议类型是否为比特币生态的特殊token协议
func IsBTCEcosystemProtocol(protocolType string) bool {
	switch protocolType {
	case Protocol.BTCEcosystemProtocols.Runes, Protocol.BTCEcosystemProtocols.BRC20, Protocol.BTCEcosystemProtocols.SRC20, Protocol.BTCEcosystemProtocols.ARC20, Protocol.BTCEcosystemProtocols.OrdinalsNFT:
		return true
	default:
		return false
	}
}
