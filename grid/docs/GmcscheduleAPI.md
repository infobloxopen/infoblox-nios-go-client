# GmcscheduleAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](GmcscheduleAPI.md#Get) | **Get** /gmcschedule | Retrieve gmcschedule objects
[**ReferenceGet**](GmcscheduleAPI.md#ReferenceGet) | **Get** /gmcschedule/{reference} | Get a specific gmcschedule object
[**ReferencePut**](GmcscheduleAPI.md#ReferencePut) | **Put** /gmcschedule/{reference} | Update a gmcschedule object



## Get

> ListGmcscheduleResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve gmcschedule objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GmcscheduleAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GmcscheduleAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListGmcscheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `GmcscheduleAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GmcscheduleAPIGetRequest` struct via the builder pattern


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

[**ListGmcscheduleResponse**](ListGmcscheduleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceGet

> GetGmcscheduleResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific gmcschedule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {
	reference := "reference_example" // string | Reference of the gmcschedule object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GmcscheduleAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GmcscheduleAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetGmcscheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `GmcscheduleAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the gmcschedule object | 

### Other Parameters

Other parameters are passed through a pointer to a `GmcscheduleAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGmcscheduleResponse**](GetGmcscheduleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateGmcscheduleResponse ReferencePut(ctx, reference).Gmcschedule(gmcschedule).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a gmcschedule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {
	reference := "reference_example" // string | Reference of the gmcschedule object
	gmcschedule := *grid.NewGmcschedule() // Gmcschedule | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GmcscheduleAPI.ReferencePut(context.Background(), reference).Gmcschedule(gmcschedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GmcscheduleAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateGmcscheduleResponse
	fmt.Fprintf(os.Stdout, "Response from `GmcscheduleAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the gmcschedule object | 

### Other Parameters

Other parameters are passed through a pointer to a `GmcscheduleAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gmcschedule** | [**Gmcschedule**](Gmcschedule.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGmcscheduleResponse**](UpdateGmcscheduleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

