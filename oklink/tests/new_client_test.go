//go:build e2e

package tests

import (
	"github.com/dapplink-labs/chain-explorer-api/oklink"
	"github.com/dapplink-labs/chain-explorer-api/oklink/conf"
	"github.com/magiconair/properties/assert"
	"reflect"
	"testing"
)

// TestNewClient Test client.New() function
func TestNewClient(t *testing.T) {
	client := oklink.New()

	// use reflect to check the fields of the client
	val := reflect.ValueOf(client).Elem()

	apiKey := val.FieldByName("key").String()

	config := &conf.Config{}
	config.New()
	exceptApiKey := config.ApiKey
	assert.Equal(t, apiKey, exceptApiKey)

	baseUrl := val.FieldByName("baseURL").String()
	exceptBaseUrl := config.BaseUrl
	assert.Equal(t, baseUrl, exceptBaseUrl)

	verbose := val.FieldByName("Verbose").Bool()
	assert.Equal(t, verbose, true)
}
