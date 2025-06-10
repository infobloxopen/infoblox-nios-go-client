# DiscoveryGridpropertiesAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverygridpropertiesGet**](DiscoveryGridpropertiesAPI.md#DiscoverygridpropertiesGet) | **Get** /discovery:gridproperties | Retrieve discovery:gridproperties objects
[**DiscoverygridpropertiesReferenceGet**](DiscoveryGridpropertiesAPI.md#DiscoverygridpropertiesReferenceGet) | **Get** /discovery:gridproperties/{reference} | Get a specific discovery:gridproperties object
[**DiscoverygridpropertiesReferencePut**](DiscoveryGridpropertiesAPI.md#DiscoverygridpropertiesReferencePut) | **Put** /discovery:gridproperties/{reference} | Update a discovery:gridproperties object



## DiscoverygridpropertiesGet

> ListDiscoveryGridpropertiesResponse DiscoverygridpropertiesGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve discovery:gridproperties objects



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
	resp, r, err := apiClient.DiscoveryGridpropertiesAPI.DiscoverygridpropertiesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryGridpropertiesAPI.DiscoverygridpropertiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverygridpropertiesGet`: ListDiscoveryGridpropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryGridpropertiesAPI.DiscoverygridpropertiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryGridpropertiesAPIDiscoverygridpropertiesGetRequest` struct via the builder pattern


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

[**ListDiscoveryGridpropertiesResponse**](ListDiscoveryGridpropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverygridpropertiesReferenceGet

> GetDiscoveryGridpropertiesResponse DiscoverygridpropertiesReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific discovery:gridproperties object



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
	reference := "reference_example" // string | Reference of the discovery:gridproperties object

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryGridpropertiesAPI.DiscoverygridpropertiesReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryGridpropertiesAPI.DiscoverygridpropertiesReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverygridpropertiesReferenceGet`: GetDiscoveryGridpropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryGridpropertiesAPI.DiscoverygridpropertiesReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:gridproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryGridpropertiesAPIDiscoverygridpropertiesReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDiscoveryGridpropertiesResponse**](GetDiscoveryGridpropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverygridpropertiesReferencePut

> UpdateDiscoveryGridpropertiesResponse DiscoverygridpropertiesReferencePut(ctx, reference).DiscoveryGridproperties(discoveryGridproperties).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a discovery:gridproperties object



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
	reference := "reference_example" // string | Reference of the discovery:gridproperties object
	discoveryGridproperties := *discovery.NewDiscoveryGridproperties() // DiscoveryGridproperties | Object data to update

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryGridpropertiesAPI.DiscoverygridpropertiesReferencePut(context.Background(), reference).DiscoveryGridproperties(discoveryGridproperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryGridpropertiesAPI.DiscoverygridpropertiesReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverygridpropertiesReferencePut`: UpdateDiscoveryGridpropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryGridpropertiesAPI.DiscoverygridpropertiesReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:gridproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryGridpropertiesAPIDiscoverygridpropertiesReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**discoveryGridproperties** | [**DiscoveryGridproperties**](DiscoveryGridproperties.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDiscoveryGridpropertiesResponse**](UpdateDiscoveryGridpropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

