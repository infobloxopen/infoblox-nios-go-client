# DiscoveryMemberpropertiesAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverymemberpropertiesGet**](DiscoveryMemberpropertiesAPI.md#DiscoverymemberpropertiesGet) | **Get** /discovery:memberproperties | Retrieve discovery:memberproperties objects
[**DiscoverymemberpropertiesReferenceGet**](DiscoveryMemberpropertiesAPI.md#DiscoverymemberpropertiesReferenceGet) | **Get** /discovery:memberproperties/{reference} | Get a specific discovery:memberproperties object
[**DiscoverymemberpropertiesReferencePut**](DiscoveryMemberpropertiesAPI.md#DiscoverymemberpropertiesReferencePut) | **Put** /discovery:memberproperties/{reference} | Update a discovery:memberproperties object



## DiscoverymemberpropertiesGet

> ListDiscoveryMemberpropertiesResponse DiscoverymemberpropertiesGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve discovery:memberproperties objects



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
	resp, r, err := apiClient.DiscoveryMemberpropertiesAPI.DiscoverymemberpropertiesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryMemberpropertiesAPI.DiscoverymemberpropertiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverymemberpropertiesGet`: ListDiscoveryMemberpropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryMemberpropertiesAPI.DiscoverymemberpropertiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryMemberpropertiesAPIDiscoverymemberpropertiesGetRequest` struct via the builder pattern


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

[**ListDiscoveryMemberpropertiesResponse**](ListDiscoveryMemberpropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverymemberpropertiesReferenceGet

> GetDiscoveryMemberpropertiesResponse DiscoverymemberpropertiesReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific discovery:memberproperties object



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
	reference := "reference_example" // string | Reference of the discovery:memberproperties object

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryMemberpropertiesAPI.DiscoverymemberpropertiesReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryMemberpropertiesAPI.DiscoverymemberpropertiesReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverymemberpropertiesReferenceGet`: GetDiscoveryMemberpropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryMemberpropertiesAPI.DiscoverymemberpropertiesReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:memberproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryMemberpropertiesAPIDiscoverymemberpropertiesReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDiscoveryMemberpropertiesResponse**](GetDiscoveryMemberpropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverymemberpropertiesReferencePut

> UpdateDiscoveryMemberpropertiesResponse DiscoverymemberpropertiesReferencePut(ctx, reference).DiscoveryMemberproperties(discoveryMemberproperties).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a discovery:memberproperties object



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
	reference := "reference_example" // string | Reference of the discovery:memberproperties object
	discoveryMemberproperties := *discovery.NewDiscoveryMemberproperties() // DiscoveryMemberproperties | Object data to update

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryMemberpropertiesAPI.DiscoverymemberpropertiesReferencePut(context.Background(), reference).DiscoveryMemberproperties(discoveryMemberproperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryMemberpropertiesAPI.DiscoverymemberpropertiesReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverymemberpropertiesReferencePut`: UpdateDiscoveryMemberpropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryMemberpropertiesAPI.DiscoverymemberpropertiesReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:memberproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryMemberpropertiesAPIDiscoverymemberpropertiesReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**discoveryMemberproperties** | [**DiscoveryMemberproperties**](DiscoveryMemberproperties.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDiscoveryMemberpropertiesResponse**](UpdateDiscoveryMemberpropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

