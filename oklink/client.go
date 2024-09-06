package oklink

import (
	"errors"
	"github.com/dapplink-labs/chain-explorer-api/oklink/conf"
	"github.com/dapplink-labs/chain-explorer-api/oklink/types"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	conn          *http.Client                                                                                     // conn The client used to send HTTP requests
	key           string                                                                                           // key Used to store the credentials for calling the OkLink API
	baseURL       string                                                                                           // baseURL The scheme and domain parts of all OkLink APIs
	Verbose       bool                                                                                             // Verbose The filed is used to specify whether to print logs
	BeforeRequest func(module, action string, param map[string]interface{}) error                                  // BeforeRequest Used to specify the function to be executed before the request is sent
	AfterRequest  func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error) // AfterRequest Used to specify the function to be executed after processing the response
}

func New() *Client {
	config := &conf.Config{}
	config.New()

	client := &Client{
		conn:          &http.Client{Timeout: config.Timeout},
		key:           config.ApiKey,
		baseURL:       config.BaseUrl,
		Verbose:       config.Verbose,
		BeforeRequest: nil,
		AfterRequest:  nil,
	}

	return client
}

// GetGas Fundamental blockchain data -> Fundamental data -> Get transaction or Gas fee
func (c *Client) GetGas(chainShortName string) (resp *types.GetGasResp, err error) {
	blockChain := &BlockChain{}
	resp, err = blockChain.GetGas(c, chainShortName)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAddressSummary Fundamental blockchain data -> Address Data -> Get basic address details
func (c *Client) GetAddressSummary(chainShortName, address string) (resp *types.AddressSummaryResp, err error) {
	addr := &Address{}
	resp, err = addr.GetAddressSummary(c, chainShortName, address)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAddressTokenBalance Fundamental blockchain data -> Address data -> Get token balance details by address
func (c *Client) GetAddressTokenBalance(chainShortName, address, protocolType string, limit int) (resp *types.AddressTokenBalanceResp, err error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	addr := &Address{}
	resp, err = addr.GetAddressTokenBalance(c, chainShortName, address, protocolType, limit)
	if err != nil {
		return nil, err
	}

	return
}

// GetAddressBalanceFills Fundamental blockchain data -> Address data -> Get token balance details
func (c *Client) GetAddressBalanceFills(chainShortName, address string, limit int) (*types.AddressBalanceFillsResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	addr := &Address{}
	resp, err := addr.GetAddressBalanceFills(c, chainShortName, address, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetBlockAddressBalanceHistory Fundamental blockchain data -> Address data -> Get address balance at specific block height
func (c *Client) GetBlockAddressBalanceHistory(chainShortName, height, address, tokenContractAddress string) (*types.BlockAddressBalanceHistoryResp, error) {
	block := &Block{}
	resp, err := block.GetBlockAddressBalanceHistory(c, chainShortName, height, address, tokenContractAddress)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAddressTransactionList Fundamental blockchain data -> Address data -> Get address transaction list
func (c *Client) GetAddressTransactionList(chainShortName, address, protocolType, tokenContractAddress, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressTransactionListResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	addr := &Address{}
	resp, err := addr.GetAddressTransactionList(c, chainShortName, address, protocolType, tokenContractAddress, startBlockHeight, endBlockHeight, isFromOrTo, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAddressNormalTransactionList Fundamental blockchain data -> Address data -> Get standard transaction list by address
func (c *Client) GetAddressNormalTransactionList(chainShortName, address, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressNormalTransactionListResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	addr := &Address{}
	resp, err := addr.GetAddressNormalTransactionList(c, chainShortName, address, startBlockHeight, endBlockHeight, isFromOrTo, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAddressInternalTransactionList Fundamental blockchain data -> Address data -> Get internal transaction list by address
func (c *Client) GetAddressInternalTransactionList(chainShortName, address, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressInternalTransactionListResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	addr := &Address{}
	resp, err := addr.GetAddressInternalTransactionList(c, chainShortName, address, startBlockHeight, endBlockHeight, isFromOrTo, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAddressTokenTransactionList Fundamental blockchain data -> Address data -> Get token transaction list by address
func (c *Client) GetAddressTokenTransactionList(chainShortName, address, protocolType, tokenContractAddress, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressTokenTransactionListResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	addr := &Address{}
	resp, err := addr.GetAddressTokenTransactionList(c, chainShortName, address, protocolType, tokenContractAddress, startBlockHeight, endBlockHeight, isFromOrTo, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAddressBalanceMulti Fundamental blockchain data -> Address data -> Get native token balance in batches
func (c *Client) GetAddressBalanceMulti(chainShortName, address string) (*types.AddressBalanceMultiResp, error) {
	addr := &Address{}
	resp, err := addr.GetAddressBalanceMulti(c, chainShortName, address)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAddressTokenBalanceMulti Fundamental blockchain data -> Address data -> Get token balance in batches
func (c *Client) GetAddressTokenBalanceMulti(chainShortName, address, protocolType, page string, limit int) (*types.AddressTokenBalanceMultiResp, error) {
	if limit != 0 && (limit < 0 || limit > 2000) {
		return nil, errors.New("the limit must be between 1 and 2000")
	}

	addr := &Address{}
	resp, err := addr.GetAddressTokenBalanceMulti(c, chainShortName, address, protocolType, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetAddressNormalTransactionListMulti Fundamental blockchain data -> Address data -> Get standard transaction list for specific block
func (c *Client) GetAddressNormalTransactionListMulti(chainShortName, address, startBlockHeight, endBlockHeight, isFromOrTo, page string, limit int) (*types.AddressNormalTransactionListMultiResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	addr := &Address{}
	resp, err := addr.GetAddressNormalTransactionListMulti(c, chainShortName, address, startBlockHeight, endBlockHeight, isFromOrTo, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAddressInternalTransactionListMulti Fundamental blockchain data -> Address data -> Get internal transaction list for specific block
func (c *Client) GetAddressInternalTransactionListMulti(chainShortName, address, startBlockHeight, endBlockHeight, page string, limit int) (*types.AddressInternalTransactionListMultiResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	addr := &Address{}
	resp, err := addr.GetAddressInternalTransactionListMulti(c, chainShortName, address, startBlockHeight, endBlockHeight, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAddressTokenTransactionListMulti Fundamental blockchain data -> Address data -> Get token transaction list for specific block
func (c *Client) GetAddressTokenTransactionListMulti(chainShortName, address, startBlockHeight, endBlockHeight, protocolType, tokenContractAddress, isFromOrTo, page string, limit int) (*types.AddressTokenTransactionListMultiResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	addr := &Address{}
	resp, err := addr.GetAddressTokenTransactionListMulti(c, chainShortName, address, startBlockHeight, endBlockHeight, protocolType, tokenContractAddress, isFromOrTo, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInscriptionTokenList Btc inscription data -> Get list of inscription tokens
func (c *Client) GetInscriptionTokenList(chainShortName, protocolType, tokenInscriptionId, symbol, projectId, startTime, endTime, page string, limit int) (*types.InscriptionTokenListResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	inscription := &Inscription{}
	resp, err := inscription.GetInscriptionTokenList(c, chainShortName, protocolType, tokenInscriptionId, symbol, projectId, startTime, endTime, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInscriptionTokenPositionList Btc inscription data -> Get list of addresses holding inscription tokens
func (c *Client) GetInscriptionTokenPositionList(chainShortName, protocolType, tokenInscriptionId, symbol, projectId, holderAddress, page string, limit int) (*types.InscriptionTokenPositionListResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	inscription := &Inscription{}
	resp, err := inscription.GetInscriptionTokenPositionList(c, chainShortName, protocolType, tokenInscriptionId, symbol, projectId, holderAddress, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInscriptionTokenTransactionList Btc inscription data -> Get inscription token transfer list
func (c *Client) GetInscriptionTokenTransactionList(chainShortName, protocolType, tokenInscriptionId, symbol, projectId, startTime, endTime, page string, limit int) (*types.InscriptionTokenTransactionListResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	inscription := &Inscription{}
	resp, err := inscription.GetInscriptionTokenTransactionList(c, chainShortName, protocolType, tokenInscriptionId, symbol, projectId, startTime, endTime, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInscriptionInscriptionList Btc inscription data -> Get inscription list
func (c *Client) GetInscriptionInscriptionList(chainShortName, protocolType, tokenInscriptionId, symbol, projectId, page string, limit int) (*types.InscriptionInscriptionListResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	inscription := &Inscription{}
	resp, err := inscription.GetInscriptionInscriptionList(c, chainShortName, protocolType, tokenInscriptionId, symbol, projectId, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInscriptionAddressTokenList Btc inscription data -> Get inscription token list by address
func (c *Client) GetInscriptionAddressTokenList(chainShortName, address, protocolType, tokenInscriptionId, symbol, projectId, page string, limit int) (*types.InscriptionAddressTokenListResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	inscription := &Inscription{}
	resp, err := inscription.GetInscriptionAddressTokenList(c, chainShortName, address, protocolType, tokenInscriptionId, symbol, projectId, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInscriptionAddressInscriptionList Btc inscription data -> Get inscription list by address
func (c *Client) GetInscriptionAddressInscriptionList(chainShortName, address, protocolType, tokenInscriptionId, symbol, projectId, page string, limit int) (*types.InscriptionAddressInscriptionListResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	inscription := &Inscription{}
	resp, err := inscription.GetInscriptionAddressInscriptionList(c, chainShortName, address, protocolType, tokenInscriptionId, symbol, projectId, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInscriptionAddressTokenTransactionList Btc inscription data -> Get inscription token transfers by address
func (c *Client) GetInscriptionAddressTokenTransactionList(chainShortName, address, protocolType, tokenInscriptionId, symbol, projectId, startTime, endTime, page string, limit int) (*types.InscriptionAddressTokenTransactionListResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	inscription := &Inscription{}
	resp, err := inscription.GetInscriptionAddressTokenTransactionList(c, chainShortName, address, protocolType, tokenInscriptionId, symbol, projectId, startTime, endTime, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInscriptionTransactionDetail Btc inscription data -> Get inscription token transaction details for specific hash
func (c *Client) GetInscriptionTransactionDetail(chainShortName, txId, protocolType, page string, limit int) (*types.InscriptionTransactionDetailResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	inscription := &Inscription{}
	resp, err := inscription.GetInscriptionTransactionDetail(c, chainShortName, txId, protocolType, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetInscriptionBlockTokenTransaction Btc inscription data -> Get inscription token transaction details for specific block
func (c *Client) GetInscriptionBlockTokenTransaction(chainShortName, height, protocolType, txnStartIndex, txnEndIndex, page string, limit int) (*types.InscriptionBlockTokenTransactionResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	inscription := &Inscription{}
	resp, err := inscription.GetInscriptionBlockTokenTransaction(c, chainShortName, height, protocolType, txnStartIndex, txnEndIndex, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetAddressUtxo UTXO-specific data -> Get remaining UTXO addresses
func (c *Client) GetAddressUtxo(chainShortName, address, page string, limit int) (*types.AddressUtxoResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	addr := &Address{}
	resp, err := addr.GetAddressUtxo(c, chainShortName, address, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetUtxoTransactionStats UTXO-specific data -> Get BTC transaction statistics
func (c *Client) GetUtxoTransactionStats(chainShortName, startTime, endTime, page string, limit int) (*types.UtxoTransactionStatsResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	utxo := &Utxo{}
	resp, err := utxo.GetUtxoTransactionStats(c, chainShortName, startTime, endTime, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetUtxoRevenueStats UTXO-specific data -> Get mining revenue statistics
func (c *Client) GetUtxoRevenueStats(chainShortName, startTime, endTime, page string, limit int) (*types.UtxoRevenueStatsResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	utxo := &Utxo{}
	resp, err := utxo.GetUtxoRevenueStats(c, chainShortName, startTime, endTime, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetBtcInscriptionsList UTXO-specific data -> BRC-20 data -> Get inscription list
func (c *Client) GetBtcInscriptionsList(token, inscriptionId, inscriptionNumber, state, page string, limit int) (*types.BtcInscriptionsListResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	btc := &Btc{}
	resp, err := btc.GetBtcInscriptionsList(c, token, inscriptionId, inscriptionNumber, state, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetBtcTokenList UTXO-specific data -> BRC-20 data -> Get BRC-20 token list
func (c *Client) GetBtcTokenList(token, orderBy, page string, limit int) (*types.BtcTokenListResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	btc := &Btc{}
	resp, err := btc.GetBtcTokenList(c, token, orderBy, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetBtcTokenDetails UTXO-specific data -> BRC-20 data -> Get BRC-20 token details
func (c *Client) GetBtcTokenDetails(token string) (*types.BtcTokenDetailsResp, error) {
	btc := &Btc{}
	resp, err := btc.GetBtcTokenDetails(c, token)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetBtcPositionList UTXO-specific data -> BRC-20 data -> Get list of addresses holding BRC-20 tokens
func (c *Client) GetBtcPositionList(token, page string, limit int) (*types.BtcPositionListResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	btc := &Btc{}
	resp, err := btc.GetBtcPositionList(c, token, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetBtcTransactionList UTXO-specific data -> BRC-20 data -> Get BRC-20 token transfer list
func (c *Client) GetBtcTransactionList(address, token, inscriptionNumber, actionType, toAddress, fromAddress, txId, blockHeight, page string, limit int) (*types.BtcTransactionListResp, error) {
	if limit != 0 && (limit < 1 || limit > 100) {
		return nil, errors.New("the limit must be between 1 and 100")
	}

	btc := &Btc{}
	resp, err := btc.GetBtcTransactionList(c, address, token, inscriptionNumber, actionType, toAddress, fromAddress, txId, blockHeight, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetBtcAddressBalanceList UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance list by address
func (c *Client) GetBtcAddressBalanceList(address, token, page string, limit int) (*types.BtcAddressBalanceListResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	btc := &Btc{}
	resp, err := btc.GetBtcAddressBalanceList(c, address, token, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// GetBtcAddressBalanceDetails UTXO-specific data -> BRC-20 data -> Get BRC-20 token balance details by address
func (c *Client) GetBtcAddressBalanceDetails(address, token, page string, limit int) (*types.BtcAddressBalanceDetailsResp, error) {
	if !c.checkLimit(limit) {
		return nil, errors.New("the limit must be between 1 and 50")
	}

	btc := &Btc{}
	resp, err := btc.GetBtcAddressBalanceDetails(c, address, token, page, limit)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *Client) SendRequest(method string, url string, reqBody io.Reader) (respBody []byte, err error) {
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Ok-Access-Key", c.key)

	res, err := c.conn.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = res.Body.Close()
		if err != nil {
			log.Fatalf("Failed to close the response body: %v", err)
		}
	}()

	respBody, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return respBody, err
}

func (c *Client) checkLimit(limit int) bool {
	if limit != 0 && (limit < 1 || limit > 50) {
		return false
	}

	return true
}
