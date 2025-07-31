package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/client"
	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/infoblox-nios-go-client/option"
) //"os"

func main() {

	NIOS_HOST_URL := "https://172.28.83.91"
	// CLIENT_CERT_PATH := "client.cert.pem"
	// CLIENT_KEY_PATH := "client.key.pem"
	//NIOS_AUTH := "admin:infoblox"

	os.Setenv("CLIENT_CERT_PATH", "client.cert.pem")
	os.Setenv("CLIENT_KEY_PATH", "client.key.pem")

	apiClient := client.NewAPIClient(
		option.WithNIOSHostUrl(NIOS_HOST_URL),
	)

	networkcontainerobj := apiClient.IPAMAPI

	reqBody := ipam.Networkcontainer{
		Network: &ipam.NetworkcontainerNetwork{String: ipam.PtrString("31.0.0.0/24")},
		Comment: ipam.PtrString("this is comment"),
	}

	resp, httpRes, err := networkcontainerobj.NetworkcontainerAPI.Create(context.Background()).Networkcontainer(reqBody).Execute()

	if err != nil {
		fmt.Printf("Error getting auth zone: %v\nFull HTTP response: %v\n", err, httpRes)
	}

	fmt.Println(resp)

	// return_fields := []string{"cloud_info", "comment", "discovery_basic_poll_settings", "discovery_blackout_setting", "extattrs"}
	// return_fields_str := strings.Join(return_fields, ",")

	// resp1, httpRes, err := networkcontainerobj.NetworkcontainerAPI.Get(context.Background()).MaxResults(20).ReturnAsObject(0).ReturnFields2(return_fields_str).Execute()
	// if err != nil {
	// 	fmt.Printf("Error getting network containers: %v\nFull HTTP response: %v\n", err, httpRes)
	// }
	// fmt.Println(resp1)

	// ref := "ZG5zLm5ldHdvcmtfY29udGFpbmVyJDMxLjAuMC4wLzI0LzA:31.0.0.0/24/default"

	// resp2, httpRes, err := networkcontainerobj.NetworkcontainerAPI.ReferenceGet(context.Background(), ref).ReturnAsObject(0).Execute()
	// if err != nil {
	// 	fmt.Printf("Error getting network container: %v\nFull HTTP response: %v\n", err, httpRes)
	// }
	// fmt.Println(resp2)

	// //update with reference
	// reqBody1 := ipam.Networkcontainer{
	// 	Comment: ipam.PtrString("this is updated comment"),
	// }

	// ref2 := "ZG5zLm5ldHdvcmtfY29udGFpbmVyJDMxLjAuMC4wLzI0LzA:31.0.0.0/24/default"

	// resp3, httpRes, err := networkcontainerobj.NetworkcontainerAPI.ReferencePut(context.Background(), ref2).Networkcontainer(reqBody1).Execute()
	// if err != nil {
	// 	fmt.Printf("Error getting network container: %v\nFull HTTP response: %v\n", err, httpRes)
	// }
	// fmt.Println(resp3)

	// //Delete with reference

	// ref3 := "ZG5zLm5ldHdvcmtfY29udGFpbmVyJDIxLjAuMC4wLzI0LzA:31.0.0.0/24/default"
	// httpRes, err = networkcontainerobj.NetworkcontainerAPI.ReferenceDelete(context.Background(), ref3).Execute()
	// if err != nil {
	// 	fmt.Printf("Error getting network container: %v\nFull HTTP response: %v\n", err, httpRes)
	// }

}
