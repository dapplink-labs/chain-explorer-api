package btc

type InscriptionsListRequest struct {
	Token             string `json:"token"`
	InscriptionId     string `json:"inscription_id"`
	InscriptionNumber string `json:"inscription_number"`
	State             string `json:"state"`
	Page              string `json:"page"`
	Limit             int    `json:"limit"`
}

// InscriptionsListResponse The Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get inscription list
type InscriptionsListResponse struct {
	Page             string                        `json:"page"`
	Limit            string                        `json:"limit"`
	TotalPage        string                        `json:"totalPage"`
	TotalInscription string                        `json:"totalInscription"`
	InscriptionsList []InscriptionsListInscription `json:"inscriptionsList"`
}

// InscriptionsListInscription The InscriptionsList field within the Data field within the Response structure of
// UTXO-specific data -> BRC-20 data -> Get inscription list
type InscriptionsListInscription struct {
	InscriptionId     string `json:"inscriptionId"`
	InscriptionNumber string `json:"inscriptionNumber"`
	Location          string `json:"location"`
	Token             string `json:"token"`
	State             string `json:"state"`
	Msg               string `json:"msg"`
	TokenType         string `json:"tokenType"`
	ActionType        string `json:"actionType"`
	LogoUrl           string `json:"logoUrl"`
	OwnerAddress      string `json:"ownerAddress"`
	TxId              string `json:"txId"`
	BlockHeight       string `json:"blockHeight"`
	ContentSize       string `json:"contentSize"`
	Time              string `json:"time"`
}
