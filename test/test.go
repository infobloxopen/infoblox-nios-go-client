package main

import (
	"context"
	"fmt"
	"github.com/infobloxopen/infoblox-nios-go-client/option"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

var readableAttributes = "comment,name,ttl,use_ttl"

func CreateARecord() {
	// Configure the API client with necessary options
	apiClient := dns.NewAPIClient(
		option.WithNIOSHostUrl("NIOS_HOST_URL"),
		option.WithNIOSUsername("NIOS_USERNAME"),
		option.WithNIOSPassword("NIOS_PASSWORD"),
	)
	RecordA := dns.RecordA{
		Comment: dns.PtrString("Example comment"),
		UseTtl:  dns.PtrBool(true),
		Ttl:     dns.PtrInt64(0),
		Name:    dns.PtrString("example_record.example.com"),
		Ipv4addr: &dns.RecordAIpv4addr{
			String: dns.PtrString("127.0.0.1"),
		},
	}
	resp, _, err := apiClient.RecordAAPI.Create(context.Background()).RecordA(RecordA).ReturnFieldsPlus(readableAttributes).Execute()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Response: %v\n", resp)
}

func main() {
	CreateARecord()
}
