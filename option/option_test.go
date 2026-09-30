package option

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

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
	config := &internal.Configuration{}
	opt := WithCACertPath(" /path/to/ca.pem ")
	opt(config)
	assert.Equal(t, "/path/to/ca.pem", config.CACertPath)
}

func TestWithCACertPath_Empty(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithCACertPath("")
	opt(config)
	assert.Empty(t, config.CACertPath)
}

// testCACert returns a self-signed CA certificate in DER and PEM form.
func testCACert(t *testing.T) (der, pemBytes []byte) {
	t.Helper()

	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)

	template := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "test-ca"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	der, err = x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	assert.NoError(t, err)

	return der, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
}

func TestValidateCACert(t *testing.T) {
	t.Setenv("CA_CERT_PATH", "")

	der, pemBytes := testCACert(t)
	dir := t.TempDir()

	pemPath := filepath.Join(dir, "ca.pem")
	assert.NoError(t, os.WriteFile(pemPath, pemBytes, 0o600))

	derPath := filepath.Join(dir, "ca.crt")
	assert.NoError(t, os.WriteFile(derPath, der, 0o600))

	missingPath := filepath.Join(dir, "does-not-exist.pem")

	assert.NoError(t, ValidateCACert())
	assert.NoError(t, ValidateCACert(WithCACertPath(pemPath)))
	assert.NoError(t, ValidateCACert(WithCACert(pemBytes)))

	assert.ErrorIs(t, ValidateCACert(WithCACertPath(missingPath)), os.ErrNotExist)
	assert.Error(t, ValidateCACert(WithCACertPath(derPath)), "DER-encoded certificate is not PEM")
	assert.Error(t, ValidateCACert(WithCACert([]byte("not a certificate"))))
}

func TestValidateCACert_EnvPath(t *testing.T) {
	t.Setenv("CA_CERT_PATH", filepath.Join(t.TempDir(), "does-not-exist.pem"))
	assert.ErrorIs(t, ValidateCACert(), os.ErrNotExist)
}
