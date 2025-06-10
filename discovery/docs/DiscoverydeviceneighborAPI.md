# DiscoveryDeviceneighborAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverydeviceneighborGet**](DiscoveryDeviceneighborAPI.md#DiscoverydeviceneighborGet) | **Get** /discovery:deviceneighbor | Retrieve discovery:deviceneighbor objects
[**DiscoverydeviceneighborReferenceGet**](DiscoveryDeviceneighborAPI.md#DiscoverydeviceneighborReferenceGet) | **Get** /discovery:deviceneighbor/{reference} | Get a specific discovery:deviceneighbor object



## DiscoverydeviceneighborGet

> ListDiscoveryDeviceneighborResponse DiscoverydeviceneighborGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve discovery:deviceneighbor objects



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
	resp, r, err := apiClient.DiscoveryDeviceneighborAPI.DiscoverydeviceneighborGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDeviceneighborAPI.DiscoverydeviceneighborGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydeviceneighborGet`: ListDiscoveryDeviceneighborResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDeviceneighborAPI.DiscoverydeviceneighborGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDeviceneighborAPIDiscoverydeviceneighborGetRequest` struct via the builder pattern


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

[**ListDiscoveryDeviceneighborResponse**](ListDiscoveryDeviceneighborResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverydeviceneighborReferenceGet

> GetDiscoveryDeviceneighborResponse DiscoverydeviceneighborReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific discovery:deviceneighbor object



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
	reference := "reference_example" // string | Reference of the discovery:deviceneighbor object

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryDeviceneighborAPI.DiscoverydeviceneighborReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDeviceneighborAPI.DiscoverydeviceneighborReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydeviceneighborReferenceGet`: GetDiscoveryDeviceneighborResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDeviceneighborAPI.DiscoverydeviceneighborReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:deviceneighbor object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDeviceneighborAPIDiscoverydeviceneighborReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDiscoveryDeviceneighborResponse**](GetDiscoveryDeviceneighborResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

