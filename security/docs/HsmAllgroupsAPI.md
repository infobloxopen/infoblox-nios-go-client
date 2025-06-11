# HsmAllgroupsAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HsmallgroupsGet**](HsmAllgroupsAPI.md#HsmallgroupsGet) | **Get** /hsm:allgroups | Retrieve hsm:allgroups objects
[**HsmallgroupsReferenceGet**](HsmAllgroupsAPI.md#HsmallgroupsReferenceGet) | **Get** /hsm:allgroups/{reference} | Get a specific hsm:allgroups object
[**HsmallgroupsReferencePut**](HsmAllgroupsAPI.md#HsmallgroupsReferencePut) | **Put** /hsm:allgroups/{reference} | Update a hsm:allgroups object



## HsmallgroupsGet

> ListHsmAllgroupsResponse HsmallgroupsGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve hsm:allgroups objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmAllgroupsAPI.HsmallgroupsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAllgroupsAPI.HsmallgroupsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmallgroupsGet`: ListHsmAllgroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmAllgroupsAPI.HsmallgroupsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HsmAllgroupsAPIHsmallgroupsGetRequest` struct via the builder pattern


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

[**ListHsmAllgroupsResponse**](ListHsmAllgroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HsmallgroupsReferenceGet

> GetHsmAllgroupsResponse HsmallgroupsReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific hsm:allgroups object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the hsm:allgroups object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmAllgroupsAPI.HsmallgroupsReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAllgroupsAPI.HsmallgroupsReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmallgroupsReferenceGet`: GetHsmAllgroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmAllgroupsAPI.HsmallgroupsReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:allgroups object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmAllgroupsAPIHsmallgroupsReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetHsmAllgroupsResponse**](GetHsmAllgroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HsmallgroupsReferencePut

> UpdateHsmAllgroupsResponse HsmallgroupsReferencePut(ctx, reference).HsmAllgroups(hsmAllgroups).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a hsm:allgroups object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the hsm:allgroups object
	hsmAllgroups := *security.NewHsmAllgroups() // HsmAllgroups | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmAllgroupsAPI.HsmallgroupsReferencePut(context.Background(), reference).HsmAllgroups(hsmAllgroups).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmAllgroupsAPI.HsmallgroupsReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmallgroupsReferencePut`: UpdateHsmAllgroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmAllgroupsAPI.HsmallgroupsReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:allgroups object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmAllgroupsAPIHsmallgroupsReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hsmAllgroups** | [**HsmAllgroups**](HsmAllgroups.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateHsmAllgroupsResponse**](UpdateHsmAllgroupsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

