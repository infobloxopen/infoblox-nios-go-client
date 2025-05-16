# DiscoverymemberpropertiesAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](DiscoverymemberpropertiesAPI.md#Get) | **Get** /discovery:memberproperties | Retrieve discovery:memberproperties objects
[**ReferenceGet**](DiscoverymemberpropertiesAPI.md#ReferenceGet) | **Get** /discovery:memberproperties/{reference} | Get a specific discovery:memberproperties object
[**ReferencePut**](DiscoverymemberpropertiesAPI.md#ReferencePut) | **Put** /discovery:memberproperties/{reference} | Update a discovery:memberproperties object



## Get

> ListDiscoveryMemberpropertiesResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

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
	resp, r, err := apiClient.DiscoverymemberpropertiesAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoverymemberpropertiesAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListDiscoveryMemberpropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoverymemberpropertiesAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoverymemberpropertiesAPIGetRequest` struct via the builder pattern


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


## ReferenceGet

> GetDiscoveryMemberpropertiesResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.DiscoverymemberpropertiesAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoverymemberpropertiesAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetDiscoveryMemberpropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoverymemberpropertiesAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:memberproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoverymemberpropertiesAPIReferenceGetRequest` struct via the builder pattern


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


## ReferencePut

> UpdateDiscoveryMemberpropertiesResponse ReferencePut(ctx, reference).DiscoveryMemberproperties(discoveryMemberproperties).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.DiscoverymemberpropertiesAPI.ReferencePut(context.Background(), reference).DiscoveryMemberproperties(discoveryMemberproperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoverymemberpropertiesAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateDiscoveryMemberpropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoverymemberpropertiesAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:memberproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoverymemberpropertiesAPIReferencePutRequest` struct via the builder pattern


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

