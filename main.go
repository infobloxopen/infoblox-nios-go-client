package main

import (
	"context"
	"fmt"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/client"
)

func main() {
	// Create a client using v2.12.3 (default)
	apiClient := client.NewAPIClient()

	dnsClient := apiClient.WAPIV_2_12_3.DNS

	// resp := apiClient.DNSAPI.Post(context.Background())
	resp, _, _ := dnsClient.RecordaAPI.Post(context.Background()).Execute()

	fmt.Println(resp)

}
