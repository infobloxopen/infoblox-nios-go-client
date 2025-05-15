package client

import (
	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/option"
)

// APIClient is an aggregation of different NIOS WAPI clients.
type APIClient struct {
	DNSAPI *dns.APIClient
}

// NewAPIClient creates a new NIOS WAPI Client.
// This is an aggregation of different NIOS WAPI clients.
// The following clients are available:
// The client can be configured with a variadic option. The following options are available:
// - WithClientName(string) sets the name of the client using the SDK.
// - WithNIOSHostUrl(string) sets the URL for NIOS Portal.
// - WithNIOSAuth(string) sets the NIOSAuth for accessing the NIOS Portal.
// - WithHTTPClient(*http.Client) sets the HTTPClient to use for the SDK.
// - WithDefaultExtAttrs(map[string]string) sets the Extensible Attributes the client can set by default for objects that has Extensible Attributes support.
// - WithDebug() sets the debug mode.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	return &APIClient{
		DNSAPI: dns.NewAPIClient(options...),
	}
}
