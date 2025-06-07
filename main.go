package main

import (
	"context"
	"fmt"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/client"
	"github.com/Infoblox-CTO/infoblox-nios-go-client/option"
)

func main() {
	// Create a client using v2.12.3 (default)

	NIOS_HOST_URL := "https://172.28.83.87"
	NIOS_AUTH := "admin:Infoblox@123"
	apiClient := client.NewAPIClient(
		option.WithNIOSHostUrl(NIOS_HOST_URL),
		option.WithNIOSAuth(NIOS_AUTH),
		option.WithDebug(true),
	)

	dnsClient := apiClient.WAPIV_2_12_3.DNS

	// resp := apiClient.DNSAPI.Post(context.Background())
	resp, _, _ := dnsClient.RecordaAPI.Get(context.Background()).MaxResults(10).Execute()

	fmt.Println(resp)

}
