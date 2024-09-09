package base

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
)

type Envelope struct {
	Status  int             `json:"status,string"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type BaseClient struct {
	coon          *http.Client
	key           string
	baseURL       string
	Verbose       bool
	BeforeRequest func(name, module, action string, param map[string]interface{}) error
	AfterRequest  func(name, module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

func NewBaseClient(apiKey string, baseUrl string, verbose bool, timeout time.Duration) *BaseClient {
	return NewCustomizedClient(CustomizationClient{
		Timeout: timeout,
		Key:     apiKey,
		BaseURL: baseUrl,
		Verbose: verbose,
	})
}

type CustomizationClient struct {
	Timeout       time.Duration
	Key           string
	BaseURL       string
	Verbose       bool
	Client        *http.Client
	BeforeRequest func(name, module, action string, param map[string]interface{}) error
	AfterRequest  func(name, module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

func NewCustomizedClient(cc CustomizationClient) *BaseClient {
	var httpClient *http.Client
	if cc.Client != nil {
		httpClient = cc.Client
	} else {
		httpClient = &http.Client{
			Timeout: cc.Timeout,
		}
	}
	return &BaseClient{
		coon:          httpClient,
		key:           cc.Key,
		baseURL:       cc.BaseURL,
		Verbose:       cc.Verbose,
		BeforeRequest: cc.BeforeRequest,
		AfterRequest:  cc.AfterRequest,
	}
}

func (bc *BaseClient) Call(name, module, action string, param map[string]interface{}, outcome interface{}) (err error) {
	if bc.BeforeRequest != nil {
		err = bc.BeforeRequest(name, module, action, param)
		if err != nil {
			err = common.WrapErr(err, "beforeRequest")
			return
		}
	}
	if bc.AfterRequest != nil {
		defer bc.AfterRequest(name, module, action, param, outcome, err)
	}

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[ouch! panic recovered] please report this with what you did and what you expected, panic detail: %v", r)
		}
	}()

	req, err := http.NewRequest(http.MethodGet, bc.CraftURL(module, action, param), http.NoBody)
	if err != nil {
		err = common.WrapErr(err, "http.NewRequest")
		return
	}

	req.Header.Set("User-Agent", "DappLinkOpenSource")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if bc.Verbose {
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

	res, err := bc.coon.Do(req)
	if err != nil {
		err = common.WrapErr(err, "sending request")
		return
	}
	defer res.Body.Close()

	if bc.Verbose {
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

	var envelope Envelope
	err = json.Unmarshal(content.Bytes(), &envelope)
	if err != nil {
		err = common.WrapErr(err, "json unmarshal envelope")
		return
	}
	if envelope.Status != 1 {
		err = fmt.Errorf("etherscan server: %s", envelope.Message)
		return
	}

	// only for etherscan handle
	if name == "etherscan" && action == "tokentx" {
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

func (bc *BaseClient) CraftURL(module, action string, param map[string]interface{}) (URL string) {
	q := url.Values{
		"module": []string{module},
		"action": []string{action},
		"apikey": []string{bc.key},
	}

	for k, v := range param {
		q[k] = common.ExtractValue(v)
	}

	URL = bc.baseURL + q.Encode()
	return
}
