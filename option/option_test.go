package option

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/internal"
)

func TestWithNIOSAuthUrl(t *testing.T) {
	config := &internal.Configuration{}
	url := "http://test.com"
	opt := WithNIOSHostUrl(url)
	opt(config)
	assert.Equal(t, url, config.NIOSHostURL)
}

func TestWithNIOSAuth(t *testing.T) {
	config := &internal.Configuration{}
	apiKey := "testKey"
	opt := WithNIOSAuth(apiKey)
	opt(config)
	assert.Equal(t, apiKey, config.NIOSAuth)
}

func TestWithHTTPClient(t *testing.T) {
	config := &internal.Configuration{}
	client := &http.Client{}
	opt := WithHTTPClient(client)
	opt(config)
	assert.Equal(t, client, config.HTTPClient)
}

func TestWithDefaultExtAttrs(t *testing.T) {
	config := &internal.Configuration{}
	extAttrs := map[string]struct{ Value string }{"tag1": {Value: "value1"}}
	opt := WithDefaultExtAttrs(extAttrs)
	opt(config)
	assert.Equal(t, extAttrs, config.DefaultExtAttrs)
}

func TestWithClientName(t *testing.T) {
	config := &internal.Configuration{}
	name := "testClient"
	opt := WithClientName(name)
	opt(config)
	assert.Equal(t, name, config.ClientName)
}

func TestWithDebug(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithDebug(true)
	opt(config)
	assert.Equal(t, true, config.Debug)
}
