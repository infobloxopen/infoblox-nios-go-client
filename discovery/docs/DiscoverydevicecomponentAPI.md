# DiscoveryDevicecomponentAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverydevicecomponentGet**](DiscoveryDevicecomponentAPI.md#DiscoverydevicecomponentGet) | **Get** /discovery:devicecomponent | Retrieve discovery:devicecomponent objects
[**DiscoverydevicecomponentReferenceGet**](DiscoveryDevicecomponentAPI.md#DiscoverydevicecomponentReferenceGet) | **Get** /discovery:devicecomponent/{reference} | Get a specific discovery:devicecomponent object



## DiscoverydevicecomponentGet

> ListDiscoveryDevicecomponentResponse DiscoverydevicecomponentGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve discovery:devicecomponent objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryDevicecomponentAPI.DiscoverydevicecomponentGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDevicecomponentAPI.DiscoverydevicecomponentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydevicecomponentGet`: ListDiscoveryDevicecomponentResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDevicecomponentAPI.DiscoverydevicecomponentGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDevicecomponentAPIDiscoverydevicecomponentGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**paging** | **int32** | Control paging of results | 
**pageId** | **string** | Page id for retrieving next page of results | 
**filters** | **map[string]interface{}** |  | 
**extattrfilter** | **map[string]interface{}** |  | 

### Return type

[**ListDiscoveryDevicecomponentResponse**](ListDiscoveryDevicecomponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverydevicecomponentReferenceGet

> GetDiscoveryDevicecomponentResponse DiscoverydevicecomponentReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific discovery:devicecomponent object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {
	reference := "reference_example" // string | Reference of the discovery:devicecomponent object

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryDevicecomponentAPI.DiscoverydevicecomponentReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDevicecomponentAPI.DiscoverydevicecomponentReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydevicecomponentReferenceGet`: GetDiscoveryDevicecomponentResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDevicecomponentAPI.DiscoverydevicecomponentReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:devicecomponent object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDevicecomponentAPIDiscoverydevicecomponentReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDiscoveryDevicecomponentResponse**](GetDiscoveryDevicecomponentResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

