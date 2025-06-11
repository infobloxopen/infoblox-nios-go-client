# DhcpStatisticsAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DhcpstatisticsGet**](DhcpStatisticsAPI.md#DhcpstatisticsGet) | **Get** /dhcp:statistics | Retrieve dhcp:statistics objects
[**DhcpstatisticsReferenceGet**](DhcpStatisticsAPI.md#DhcpstatisticsReferenceGet) | **Get** /dhcp:statistics/{reference} | Get a specific dhcp:statistics object



## DhcpstatisticsGet

> ListDhcpStatisticsResponse DhcpstatisticsGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dhcp:statistics objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpStatisticsAPI.DhcpstatisticsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpStatisticsAPI.DhcpstatisticsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DhcpstatisticsGet`: ListDhcpStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpStatisticsAPI.DhcpstatisticsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpStatisticsAPIDhcpstatisticsGetRequest` struct via the builder pattern


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

[**ListDhcpStatisticsResponse**](ListDhcpStatisticsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DhcpstatisticsReferenceGet

> GetDhcpStatisticsResponse DhcpstatisticsReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dhcp:statistics object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the dhcp:statistics object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpStatisticsAPI.DhcpstatisticsReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpStatisticsAPI.DhcpstatisticsReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DhcpstatisticsReferenceGet`: GetDhcpStatisticsResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpStatisticsAPI.DhcpstatisticsReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcp:statistics object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpStatisticsAPIDhcpstatisticsReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDhcpStatisticsResponse**](GetDhcpStatisticsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

