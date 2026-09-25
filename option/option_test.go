package option

import (
	"net/http"
	"os"
	"path/filepath"
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

func TestWithSslVerify(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithSslVerify(true)
	opt(config)
	assert.Equal(t, true, config.SslVerify)
}

func TestWithCACert(t *testing.T) {
	config := &internal.Configuration{}
	caCertPEM := []byte("test-ca-cert-pem")
	opt := WithCACert(caCertPEM)
	opt(config)
	assert.Equal(t, caCertPEM, config.CACert)
}

func TestWithCACert_Empty(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithCACert(nil)
	opt(config)
	assert.Nil(t, config.CACert)
}

func TestWithCACertPath(t *testing.T) {
	caCertPEM := []byte("test-ca-cert-pem")
	caCertPath := filepath.Join(t.TempDir(), "ca.cert.pem")
	require := assert.New(t)
	require.NoError(os.WriteFile(caCertPath, caCertPEM, 0o600))

	config := &internal.Configuration{}
	opt := WithCACertPath(caCertPath)
	opt(config)
	assert.Equal(t, caCertPEM, config.CACert)
}

func TestWithCACertPath_MissingFile(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithCACertPath(filepath.Join(t.TempDir(), "does-not-exist.pem"))
	opt(config)
	assert.Nil(t, config.CACert)
}
