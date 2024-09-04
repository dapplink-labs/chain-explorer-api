package base

const (
	Mainnet  Network = "api"
	Ropsten  Network = "api-ropsten"
	Kovan    Network = "api-kovan"
	Rinkby   Network = "api-rinkeby"
	Goerli   Network = "api-goerli"
	Tobalaba Network = "api-tobalaba"
)

type Network string

func (n Network) SubDomain() (sub string) {
	return string(n)
}
