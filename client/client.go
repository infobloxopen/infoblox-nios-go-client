package client

import (
	"github.com/Infoblox-CTO/infoblox-nios-go-client/internal"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/option"
	v2_13_6_acl "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/acl"
	v2_13_6_cloud "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/cloud"
	v2_13_6_dhcp "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/dhcp"
	v2_13_6_discovery "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/discovery"
	v2_13_6_dns "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/dns"
	v2_13_6_dtc "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/dtc"
	v2_13_6_federatedrealms "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/federatedrealms"
	v2_13_6_grid "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/grid"
	v2_13_6_ipam "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/ipam"
	v2_13_6_microsoftserver "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/microsoftserver"
	v2_13_6_misc "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/misc"
	v2_13_6_notification "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/notification"
	v2_13_6_parentalcontrol "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/parentalcontrol"
	v2_13_6_rir "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/rir"
	v2_13_6_rpz "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/rpz"
	v2_13_6_security "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/security"
	v2_13_6_smartfolder "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/smartfolder"
	v2_13_6_threatinsight "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/threatinsight"
	v2_13_6_threatprotection "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/threatprotection"
)

type APIClient struct {
	WAPIV_2_13_6 WAPIV_2_13_6
}

type WAPIV_2_13_6 struct {
	ACL              *v2_13_6_acl.APIClient
	Cloud            *v2_13_6_cloud.APIClient
	DHCP             *v2_13_6_dhcp.APIClient
	Discovery        *v2_13_6_discovery.APIClient
	DNS              *v2_13_6_dns.APIClient
	DTC              *v2_13_6_dtc.APIClient
	FederatedRealms  *v2_13_6_federatedrealms.APIClient
	Grid             *v2_13_6_grid.APIClient
	IPAM             *v2_13_6_ipam.APIClient
	MicrosoftServer  *v2_13_6_microsoftserver.APIClient
	Misc             *v2_13_6_misc.APIClient
	Notification     *v2_13_6_notification.APIClient
	ParentalControl  *v2_13_6_parentalcontrol.APIClient
	RIR              *v2_13_6_rir.APIClient
	RPZ              *v2_13_6_rpz.APIClient
	Security         *v2_13_6_security.APIClient
	SmartFolder      *v2_13_6_smartfolder.APIClient
	ThreatInsight    *v2_13_6_threatinsight.APIClient
	ThreatProtection *v2_13_6_threatprotection.APIClient
}

// NewAPIClient creates a new NIOS WAPI Client.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	cfg := internal.NewConfiguration()
	for _, o := range options {
		o(cfg)
	}

	client := &APIClient{}

	initV2_13_6(client, options...)

	return client
}

// initV2_13_6 initializes all services for the 2.13.6 API version
func initV2_13_6(client *APIClient, options ...option.ClientOption) {
	client.WAPIV_2_13_6.ACL = v2_13_6_acl.NewAPIClient(options...)
	client.WAPIV_2_13_6.Cloud = v2_13_6_cloud.NewAPIClient(options...)
	client.WAPIV_2_13_6.DHCP = v2_13_6_dhcp.NewAPIClient(options...)
	client.WAPIV_2_13_6.Discovery = v2_13_6_discovery.NewAPIClient(options...)
	client.WAPIV_2_13_6.DNS = v2_13_6_dns.NewAPIClient(options...)
	client.WAPIV_2_13_6.DTC = v2_13_6_dtc.NewAPIClient(options...)
	client.WAPIV_2_13_6.FederatedRealms = v2_13_6_federatedrealms.NewAPIClient(options...)
	client.WAPIV_2_13_6.Grid = v2_13_6_grid.NewAPIClient(options...)
	client.WAPIV_2_13_6.IPAM = v2_13_6_ipam.NewAPIClient(options...)
	client.WAPIV_2_13_6.MicrosoftServer = v2_13_6_microsoftserver.NewAPIClient(options...)
	client.WAPIV_2_13_6.Misc = v2_13_6_misc.NewAPIClient(options...)
	client.WAPIV_2_13_6.Notification = v2_13_6_notification.NewAPIClient(options...)
	client.WAPIV_2_13_6.ParentalControl = v2_13_6_parentalcontrol.NewAPIClient(options...)
	client.WAPIV_2_13_6.RIR = v2_13_6_rir.NewAPIClient(options...)
	client.WAPIV_2_13_6.RPZ = v2_13_6_rpz.NewAPIClient(options...)
	client.WAPIV_2_13_6.Security = v2_13_6_security.NewAPIClient(options...)
	client.WAPIV_2_13_6.SmartFolder = v2_13_6_smartfolder.NewAPIClient(options...)
	client.WAPIV_2_13_6.ThreatInsight = v2_13_6_threatinsight.NewAPIClient(options...)
	client.WAPIV_2_13_6.ThreatProtection = v2_13_6_threatprotection.NewAPIClient(options...)
}
