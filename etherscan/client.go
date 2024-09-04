package etherscan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/dapplink-labs/chain-explorer-api/common"
	"github.com/dapplink-labs/chain-explorer-api/etherscan/base"
)

type Client struct {
	coon          *http.Client
	key           string
	baseURL       string
	Verbose       bool
	BeforeRequest func(module, action string, param map[string]interface{}) error
	AfterRequest  func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

func New(network base.Network, APIKey string) *Client {
	return NewCustomized(Customization{
		Timeout: 30 * time.Second,
		Key:     APIKey,
		BaseURL: fmt.Sprintf(`https://%s.etherscan.io/api?`, network.SubDomain()),
	})
}

type Customization struct {
	Timeout       time.Duration
	Key           string
	BaseURL       string
	Verbose       bool
	Client        *http.Client
	BeforeRequest func(module, action string, param map[string]interface{}) error
	AfterRequest  func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

func NewCustomized(config Customization) *Client {
	var httpClient *http.Client
	if config.Client != nil {
		httpClient = config.Client
	} else {
		httpClient = &http.Client{
			Timeout: config.Timeout,
		}
	}
	return &Client{
		coon:          httpClient,
		key:           config.Key,
		baseURL:       config.BaseURL,
		Verbose:       config.Verbose,
		BeforeRequest: config.BeforeRequest,
		AfterRequest:  config.AfterRequest,
	}
}

func (c *Client) call(module, action string, param map[string]interface{}, outcome interface{}) (err error) {
	if c.BeforeRequest != nil {
		err = c.BeforeRequest(module, action, param)
		if err != nil {
			err = common.WrapErr(err, "beforeRequest")
			return
		}
	}
	if c.AfterRequest != nil {
		defer c.AfterRequest(module, action, param, outcome, err)
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[ouch! panic recovered] please report this with what you did and what you expected, panic detail: %v", r)
		}
	}()

	req, err := http.NewRequest(http.MethodGet, c.craftURL(module, action, param), http.NoBody)
	if err != nil {
		err = common.WrapErr(err, "http.NewRequest")
		return
	}
	req.Header.Set("User-Agent", "etherscan-api(Go)")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if c.Verbose {
		var reqDump []byte
		reqDump, err = httputil.DumpRequestOut(req, false)
		if err != nil {
			err = common.WrapErr(err, "verbose mode req dump failed")
			return
		}

		fmt.Printf("\n%s\n", reqDump)
		defer func() {
			if err != nil {
				fmt.Printf("[Error] %v\n", err)
			}
		}()
	}

	res, err := c.coon.Do(req)
	if err != nil {
		err = common.WrapErr(err, "sending request")
		return
	}
	defer res.Body.Close()

	if c.Verbose {
		var resDump []byte
		resDump, err = httputil.DumpResponse(res, true)
		if err != nil {
			err = common.WrapErr(err, "verbose mode res dump failed")
			return
		}

		fmt.Printf("%s\n", resDump)
	}

	var content bytes.Buffer
	if _, err = io.Copy(&content, res.Body); err != nil {
		err = common.WrapErr(err, "reading response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status %v %s, response body: %s", res.StatusCode, res.Status, content.String())
		return
	}

	var envelope base.Envelope
	err = json.Unmarshal(content.Bytes(), &envelope)
	if err != nil {
		err = common.WrapErr(err, "json unmarshal envelope")
		return
	}
	if envelope.Status != 1 {
		err = fmt.Errorf("etherscan server: %s", envelope.Message)
		return
	}

	if action == "tokentx" {
		err = json.Unmarshal(bytes.Replace(envelope.Result, []byte(`"tokenDecimal":""`), []byte(`"tokenDecimal":"0"`), -1), outcome)
	} else {
		err = json.Unmarshal(envelope.Result, outcome)
	}
	if err != nil {
		err = common.WrapErr(err, "json unmarshal outcome")
		return
	}

	return
}

func (c *Client) craftURL(module, action string, param map[string]interface{}) (URL string) {
	q := url.Values{
		"module": []string{module},
		"action": []string{action},
		"apikey": []string{c.key},
	}

	for k, v := range param {
		q[k] = common.ExtractValue(v)
	}

	URL = c.baseURL + q.Encode()
	return
}
