# HostnamerewritepolicyAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](HostnamerewritepolicyAPI.md#Get) | **Get** /hostnamerewritepolicy | Retrieve hostnamerewritepolicy objects
[**ReferenceGet**](HostnamerewritepolicyAPI.md#ReferenceGet) | **Get** /hostnamerewritepolicy/{reference} | Get a specific hostnamerewritepolicy object
[**ReferencePut**](HostnamerewritepolicyAPI.md#ReferencePut) | **Put** /hostnamerewritepolicy/{reference} | Update a hostnamerewritepolicy object



## Get

> ListHostnamerewritepolicyResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve hostnamerewritepolicy objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
)

func main() {

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.HostnamerewritepolicyAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostnamerewritepolicyAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListHostnamerewritepolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `HostnamerewritepolicyAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HostnamerewritepolicyAPIGetRequest` struct via the builder pattern


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

[**ListHostnamerewritepolicyResponse**](ListHostnamerewritepolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceGet

> GetHostnamerewritepolicyResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific hostnamerewritepolicy object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the hostnamerewritepolicy object

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.HostnamerewritepolicyAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostnamerewritepolicyAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetHostnamerewritepolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `HostnamerewritepolicyAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hostnamerewritepolicy object | 

### Other Parameters

Other parameters are passed through a pointer to a `HostnamerewritepolicyAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetHostnamerewritepolicyResponse**](GetHostnamerewritepolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateHostnamerewritepolicyResponse ReferencePut(ctx, reference).Hostnamerewritepolicy(hostnamerewritepolicy).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a hostnamerewritepolicy object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the hostnamerewritepolicy object
	hostnamerewritepolicy := *ipam.NewHostnamerewritepolicy() // Hostnamerewritepolicy | Object data to update

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.HostnamerewritepolicyAPI.ReferencePut(context.Background(), reference).Hostnamerewritepolicy(hostnamerewritepolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostnamerewritepolicyAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateHostnamerewritepolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `HostnamerewritepolicyAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hostnamerewritepolicy object | 

### Other Parameters

Other parameters are passed through a pointer to a `HostnamerewritepolicyAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hostnamerewritepolicy** | [**Hostnamerewritepolicy**](Hostnamerewritepolicy.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateHostnamerewritepolicyResponse**](UpdateHostnamerewritepolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

