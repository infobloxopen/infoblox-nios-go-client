package client

import (
	"github.com/Infoblox-CTO/infoblox-nios-go-client/internal"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/option"
	v2_12_3_cloud "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.12.3/cloud"
	v2_12_3_dns "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.12.3/dns"
	v2_12_3_dtc "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.12.3/dtc"
	v2_13_6_dtc "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/cloud"
	v2_13_6_dns "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/dns"
	v2_13_6_cloud "github.com/Infoblox-CTO/infoblox-nios-go-client/v2.13.6/dtc"
)

type APIClient struct {
	WAPIV_2_12_3 WAPIV_2_12_3
	WAPIV_2_13_6 WAPIV_2_13_6
}

type WAPIV_2_12_3 struct {
	DNS   *v2_12_3_dns.APIClient
	DTC   *v2_12_3_dtc.APIClient
	Cloud *v2_12_3_cloud.APIClient // Assuming Cloud is also part of v2.12.3
}

type WAPIV_2_13_6 struct {
	DNS   *v2_13_6_dns.APIClient
	DTC   *v2_13_6_dtc.APIClient
	Cloud *v2_13_6_cloud.APIClient // Assuming Cloud is also part of v2.13.6
}

// NewAPIClient creates a new NIOS WAPI Client.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	cfg := internal.NewConfiguration()
	for _, o := range options {
		o(cfg)
	}

	client := &APIClient{}

	initV2_12_3(client, options...)
	initV2_13_6(client, options...)

	return client
}

// initV2_12_3 initializes all services for the 2.12.3 API version
func initV2_12_3(client *APIClient, options ...option.ClientOption) {
	client.WAPIV_2_12_3.DNS = v2_12_3_dns.NewAPIClient(options...)
	client.WAPIV_2_12_3.DTC = v2_12_3_dtc.NewAPIClient(options...)
	client.WAPIV_2_12_3.Cloud = v2_12_3_cloud.NewAPIClient(options...)
}

// initV2_13_5 initializes all services for the 2.13.5 API version
func initV2_13_6(client *APIClient, options ...option.ClientOption) {
	client.WAPIV_2_13_6.DNS = v2_13_6_dns.NewAPIClient(options...)
	client.WAPIV_2_13_6.DTC = v2_13_6_dtc.NewAPIClient(options...)
	client.WAPIV_2_13_6.Cloud = v2_13_6_cloud.NewAPIClient(options...)
}
