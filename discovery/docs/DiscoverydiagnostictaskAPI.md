# DiscoveryDiagnostictaskAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverydiagnostictaskGet**](DiscoveryDiagnostictaskAPI.md#DiscoverydiagnostictaskGet) | **Get** /discovery:diagnostictask | Retrieve discovery:diagnostictask objects
[**DiscoverydiagnostictaskReferenceGet**](DiscoveryDiagnostictaskAPI.md#DiscoverydiagnostictaskReferenceGet) | **Get** /discovery:diagnostictask/{reference} | Get a specific discovery:diagnostictask object
[**DiscoverydiagnostictaskReferencePut**](DiscoveryDiagnostictaskAPI.md#DiscoverydiagnostictaskReferencePut) | **Put** /discovery:diagnostictask/{reference} | Update a discovery:diagnostictask object



## DiscoverydiagnostictaskGet

> ListDiscoveryDiagnostictaskResponse DiscoverydiagnostictaskGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve discovery:diagnostictask objects



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
	resp, r, err := apiClient.DiscoveryDiagnostictaskAPI.DiscoverydiagnostictaskGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDiagnostictaskAPI.DiscoverydiagnostictaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydiagnostictaskGet`: ListDiscoveryDiagnostictaskResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDiagnostictaskAPI.DiscoverydiagnostictaskGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDiagnostictaskAPIDiscoverydiagnostictaskGetRequest` struct via the builder pattern


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

[**ListDiscoveryDiagnostictaskResponse**](ListDiscoveryDiagnostictaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverydiagnostictaskReferenceGet

> GetDiscoveryDiagnostictaskResponse DiscoverydiagnostictaskReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific discovery:diagnostictask object



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
	reference := "reference_example" // string | Reference of the discovery:diagnostictask object

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryDiagnostictaskAPI.DiscoverydiagnostictaskReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDiagnostictaskAPI.DiscoverydiagnostictaskReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydiagnostictaskReferenceGet`: GetDiscoveryDiagnostictaskResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDiagnostictaskAPI.DiscoverydiagnostictaskReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:diagnostictask object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDiagnostictaskAPIDiscoverydiagnostictaskReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDiscoveryDiagnostictaskResponse**](GetDiscoveryDiagnostictaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverydiagnostictaskReferencePut

> UpdateDiscoveryDiagnostictaskResponse DiscoverydiagnostictaskReferencePut(ctx, reference).DiscoveryDiagnostictask(discoveryDiagnostictask).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a discovery:diagnostictask object



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
	reference := "reference_example" // string | Reference of the discovery:diagnostictask object
	discoveryDiagnostictask := *discovery.NewDiscoveryDiagnostictask() // DiscoveryDiagnostictask | Object data to update

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryDiagnostictaskAPI.DiscoverydiagnostictaskReferencePut(context.Background(), reference).DiscoveryDiagnostictask(discoveryDiagnostictask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDiagnostictaskAPI.DiscoverydiagnostictaskReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydiagnostictaskReferencePut`: UpdateDiscoveryDiagnostictaskResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDiagnostictaskAPI.DiscoverydiagnostictaskReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:diagnostictask object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDiagnostictaskAPIDiscoverydiagnostictaskReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**discoveryDiagnostictask** | [**DiscoveryDiagnostictask**](DiscoveryDiagnostictask.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDiscoveryDiagnostictaskResponse**](UpdateDiscoveryDiagnostictaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

