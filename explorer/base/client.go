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

type EtherscanEnvelope struct {
	Status  int             `json:"status,string"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

type OklinkEnvelope struct {
	Code string          `json:"code"`
	Msg  string          `json:"msg"`
	Data json.RawMessage `json:"data"`
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

func (bc *BaseClient) Call(name, module, action, apiUrl string, param map[string]interface{}, outcome interface{}) (err error) {
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

	var req *http.Request
	var httpErr error

	if name == "etherscan" {
		fmt.Println(" apiKey==", bc.key, "name=", name)
		req, httpErr = http.NewRequest(http.MethodGet, bc.CraftEtherScanURL(module, action, param), http.NoBody)
		if httpErr != nil {
			err = common.WrapErr(httpErr, "http.NewRequest")
			return
		}
		req.Header.Set("User-Agent", "DappLinkOpenSource")
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	if name == "oklink" {
		fmt.Println("apiKey and name", "apiKey:", bc.key, "name:", name)
		req, httpErr = http.NewRequest(http.MethodGet, bc.CraftOkLinkURL(apiUrl), http.NoBody)
		if httpErr != nil {
			err = common.WrapErr(httpErr, "http.NewRequest")
			return
		}
		req.Header.Set("Ok-Access-Key", bc.key)
	}

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
	if name == "etherscan" {
		err = bc.HandleEtherscanResponse(action, content.Bytes(), outcome)
		if err != nil {
			fmt.Printf("handle etherscan err", "err", err)
		}
	} else if name == "oklink" {
		err = bc.HandleOklinkResponse(content.Bytes(), outcome)
		if err != nil {
			fmt.Printf("handle oklink err", "err", err)
		}
	} else {
		fmt.Printf("unsuport type")
	}
	return
}

func (bc *BaseClient) CraftEtherScanURL(module, action string, param map[string]interface{}) (URL string) {
	q := url.Values{
		"module": []string{module},
		"action": []string{action},
		"apikey": []string{bc.key},
	}

	for k, v := range param {
		q[k] = common.ExtractValue(v)
	}

	URL = bc.baseURL + q.Encode()
	fmt.Println("CraftEtherScanURL", URL)
	return
}

func (bc *BaseClient) CraftOkLinkURL(apiUrl string) (URL string) {
	URL = bc.baseURL + apiUrl
	fmt.Println("CraftEtherScanURL", URL)
	return
}

func (bc *BaseClient) HandleEtherscanResponse(action string, data []byte, outcome interface{}) error {
	var etherscanEnvelope EtherscanEnvelope
	err := json.Unmarshal(data, &etherscanEnvelope)
	if err != nil {
		err = common.WrapErr(err, "json unmarshal etherscan envelope")
		return err
	}
	if etherscanEnvelope.Status != 1 {
		err = fmt.Errorf("etherscan server: %s", etherscanEnvelope.Message)
		return err
	}
	if action == "tokentx" {
		err = json.Unmarshal(bytes.Replace(etherscanEnvelope.Result, []byte(`"tokenDecimal":""`), []byte(`"tokenDecimal":"0"`), -1), outcome)
	} else {
		err = json.Unmarshal(etherscanEnvelope.Result, outcome)
	}
	if err != nil {
		err = common.WrapErr(err, "json unmarshal etherscan outcome")
		return err
	}
	return nil
}

func (bc *BaseClient) HandleOklinkResponse(data []byte, outcome interface{}) error {
	var oklink OklinkEnvelope
	err := json.Unmarshal(data, &oklink)
	if err != nil {
		err = common.WrapErr(err, "json unmarshal oklink envelope")
		return err
	}
	fmt.Println("Parse oklink data success", "code", oklink.Code, "msg", oklink.Msg)
	if oklink.Code != "0" {
		err = fmt.Errorf("oklink scan server: %s", oklink.Msg)
		return err
	}
	err = json.Unmarshal(oklink.Data, outcome)
	if err != nil {
		err = common.WrapErr(err, "json unmarshal oklink outcome")
		return err
	}
	return nil
}
