package client

import (
	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/option"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/rpz"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

// APIClient is an aggregation of different NIOS WAPI clients.
type APIClient struct {
	CloudAPI     *cloud.APIClient
	DHCPAPI      *dhcp.APIClient
	DiscoveryAPI *discovery.APIClient
	DNSAPI       *dns.APIClient
	DTCAPI       *dtc.APIClient
	GridAPI      *grid.APIClient
	IPAMAPI      *ipam.APIClient
	RPZAPI       *rpz.APIClient
	SecurityAPI  *security.APIClient
}

// NewAPIClient creates a new NIOS WAPI Client.
// This is an aggregation of different NIOS WAPI clients.
// The following clients are available:
// - CloudAPI
// - DHCPAPI
// - DiscoveryAPI
// - DNSAPI
// - DTCAPI
// - GridAPI
// - IPAMAPI
// - RPZAPI
// - SecurityAPI
// The client can be configured with a variadic option. The following options are available:
// - WithClientName(string) sets the name of the client using the SDK.
// - WithNIOSHostUrl(string) sets the URL for NIOS Portal.
// - WithNIOSUsername(string) sets the Username for the NIOS Portal.
// - WithNIOSPassword(string) sets the Password for the NIOS Portal.
// - WithHTTPClient(*http.Client) sets the HTTPClient to use for the SDK.
// - WithDefaultExtAttrs(map[string]struct{ Value String }) sets the Extensible Attributes the client can set by default for objects that has Extensible Attributes support.
// - WithDebug() sets the debug mode.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	return &APIClient{
		CloudAPI:     cloud.NewAPIClient(options...),
		DHCPAPI:      dhcp.NewAPIClient(options...),
		DiscoveryAPI: discovery.NewAPIClient(options...),
		DNSAPI:       dns.NewAPIClient(options...),
		DTCAPI:       dtc.NewAPIClient(options...),
		GridAPI:      grid.NewAPIClient(options...),
		IPAMAPI:      ipam.NewAPIClient(options...),
		RPZAPI:       rpz.NewAPIClient(options...),
		SecurityAPI:  security.NewAPIClient(options...),
	}
}
