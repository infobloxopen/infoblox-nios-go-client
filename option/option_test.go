package option

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

func TestWithNIOSAuthUrl(t *testing.T) {
	config := &internal.Configuration{}
	url := "http://test.com"
	opt := WithNIOSHostUrl(url)
	opt(config)
	assert.Equal(t, url, config.NIOSHostURL)
}

func TestWithNIOSUsername(t *testing.T) {
	config := &internal.Configuration{}
	username := "testUser"
	opt := WithNIOSUsername(username)
	opt(config)
	assert.Equal(t, username, config.NIOSUsername)
}

func TestWithNIOSPassword(t *testing.T) {
	config := &internal.Configuration{}
	password := "testPassword"
	opt := WithNIOSPassword(password)
	opt(config)
	assert.Equal(t, password, config.NIOSPassword)
}

func TestWithNIOSWAPIVersion(t *testing.T) {
	config := &internal.Configuration{}
	wapiVersion := "v2.13.6"
	opt := WithNIOSWapiVersion(wapiVersion)
	opt(config)
	assert.Equal(t, wapiVersion, config.NIOSWAPIVersion)
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
