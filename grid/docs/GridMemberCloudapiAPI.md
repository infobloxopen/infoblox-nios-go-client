# GridMemberCloudapiAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GridmembercloudapiGet**](GridMemberCloudapiAPI.md#GridmembercloudapiGet) | **Get** /grid:member:cloudapi | Retrieve grid:member:cloudapi objects
[**GridmembercloudapiReferenceGet**](GridMemberCloudapiAPI.md#GridmembercloudapiReferenceGet) | **Get** /grid:member:cloudapi/{reference} | Get a specific grid:member:cloudapi object
[**GridmembercloudapiReferencePut**](GridMemberCloudapiAPI.md#GridmembercloudapiReferencePut) | **Put** /grid:member:cloudapi/{reference} | Update a grid:member:cloudapi object



## GridmembercloudapiGet

> ListGridMemberCloudapiResponse GridmembercloudapiGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:member:cloudapi objects



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
	resp, r, err := apiClient.GridMemberCloudapiAPI.GridmembercloudapiGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridMemberCloudapiAPI.GridmembercloudapiGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridmembercloudapiGet`: ListGridMemberCloudapiResponse
	fmt.Fprintf(os.Stdout, "Response from `GridMemberCloudapiAPI.GridmembercloudapiGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridMemberCloudapiAPIGridmembercloudapiGetRequest` struct via the builder pattern


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

[**ListGridMemberCloudapiResponse**](ListGridMemberCloudapiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridmembercloudapiReferenceGet

> GetGridMemberCloudapiResponse GridmembercloudapiReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:member:cloudapi object



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
	reference := "reference_example" // string | Reference of the grid:member:cloudapi object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridMemberCloudapiAPI.GridmembercloudapiReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridMemberCloudapiAPI.GridmembercloudapiReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridmembercloudapiReferenceGet`: GetGridMemberCloudapiResponse
	fmt.Fprintf(os.Stdout, "Response from `GridMemberCloudapiAPI.GridmembercloudapiReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:member:cloudapi object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridMemberCloudapiAPIGridmembercloudapiReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridMemberCloudapiResponse**](GetGridMemberCloudapiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridmembercloudapiReferencePut

> UpdateGridMemberCloudapiResponse GridmembercloudapiReferencePut(ctx, reference).GridMemberCloudapi(gridMemberCloudapi).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:member:cloudapi object



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
	reference := "reference_example" // string | Reference of the grid:member:cloudapi object
	gridMemberCloudapi := *grid.NewGridMemberCloudapi() // GridMemberCloudapi | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridMemberCloudapiAPI.GridmembercloudapiReferencePut(context.Background(), reference).GridMemberCloudapi(gridMemberCloudapi).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridMemberCloudapiAPI.GridmembercloudapiReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridmembercloudapiReferencePut`: UpdateGridMemberCloudapiResponse
	fmt.Fprintf(os.Stdout, "Response from `GridMemberCloudapiAPI.GridmembercloudapiReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:member:cloudapi object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridMemberCloudapiAPIGridmembercloudapiReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridMemberCloudapi** | [**GridMemberCloudapi**](GridMemberCloudapi.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridMemberCloudapiResponse**](UpdateGridMemberCloudapiResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

