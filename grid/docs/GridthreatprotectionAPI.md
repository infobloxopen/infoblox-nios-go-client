# GridThreatprotectionAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](GridThreatprotectionAPI.md#List) | **Get** /grid:threatprotection | Retrieve grid:threatprotection objects
[**Read**](GridThreatprotectionAPI.md#Read) | **Get** /grid:threatprotection/{reference} | Get a specific grid:threatprotection object
[**Update**](GridThreatprotectionAPI.md#Update) | **Put** /grid:threatprotection/{reference} | Update a grid:threatprotection object



## List

> ListGridThreatprotectionResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:threatprotection objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
)

func main() {

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridThreatprotectionAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridThreatprotectionAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListGridThreatprotectionResponse
	fmt.Fprintf(os.Stdout, "Response from `GridThreatprotectionAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridThreatprotectionAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**paging** | **int32** | Control paging of results | 
**pageId** | **string** | Page id for retrieving next page of results | 
**filters** | **map[string]interface{}** |  | 
**extattrfilter** | **map[string]interface{}** |  | 

### Return type

[**ListGridThreatprotectionResponse**](ListGridThreatprotectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetGridThreatprotectionResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:threatprotection object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
)

func main() {
	reference := "reference_example" // string | Reference of the grid:threatprotection object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridThreatprotectionAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridThreatprotectionAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetGridThreatprotectionResponse
	fmt.Fprintf(os.Stdout, "Response from `GridThreatprotectionAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:threatprotection object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridThreatprotectionAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridThreatprotectionResponse**](GetGridThreatprotectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateGridThreatprotectionResponse Update(ctx, reference).GridThreatprotection(gridThreatprotection).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a grid:threatprotection object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
)

func main() {
	reference := "reference_example" // string | Reference of the grid:threatprotection object
	gridThreatprotection := *grid.NewGridThreatprotection() // GridThreatprotection | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridThreatprotectionAPI.Update(context.Background(), reference).GridThreatprotection(gridThreatprotection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridThreatprotectionAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateGridThreatprotectionResponse
	fmt.Fprintf(os.Stdout, "Response from `GridThreatprotectionAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:threatprotection object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridThreatprotectionAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridThreatprotection** | [**GridThreatprotection**](GridThreatprotection.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridThreatprotectionResponse**](UpdateGridThreatprotectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

