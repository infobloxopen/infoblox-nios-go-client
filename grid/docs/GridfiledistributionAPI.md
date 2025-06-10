# GridFiledistributionAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GridfiledistributionGet**](GridFiledistributionAPI.md#GridfiledistributionGet) | **Get** /grid:filedistribution | Retrieve grid:filedistribution objects
[**GridfiledistributionReferenceGet**](GridFiledistributionAPI.md#GridfiledistributionReferenceGet) | **Get** /grid:filedistribution/{reference} | Get a specific grid:filedistribution object
[**GridfiledistributionReferencePut**](GridFiledistributionAPI.md#GridfiledistributionReferencePut) | **Put** /grid:filedistribution/{reference} | Update a grid:filedistribution object



## GridfiledistributionGet

> ListGridFiledistributionResponse GridfiledistributionGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:filedistribution objects



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
	resp, r, err := apiClient.GridFiledistributionAPI.GridfiledistributionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridFiledistributionAPI.GridfiledistributionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridfiledistributionGet`: ListGridFiledistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `GridFiledistributionAPI.GridfiledistributionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridFiledistributionAPIGridfiledistributionGetRequest` struct via the builder pattern


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

[**ListGridFiledistributionResponse**](ListGridFiledistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridfiledistributionReferenceGet

> GetGridFiledistributionResponse GridfiledistributionReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:filedistribution object



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
	reference := "reference_example" // string | Reference of the grid:filedistribution object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridFiledistributionAPI.GridfiledistributionReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridFiledistributionAPI.GridfiledistributionReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridfiledistributionReferenceGet`: GetGridFiledistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `GridFiledistributionAPI.GridfiledistributionReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:filedistribution object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridFiledistributionAPIGridfiledistributionReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridFiledistributionResponse**](GetGridFiledistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridfiledistributionReferencePut

> UpdateGridFiledistributionResponse GridfiledistributionReferencePut(ctx, reference).GridFiledistribution(gridFiledistribution).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:filedistribution object



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
	reference := "reference_example" // string | Reference of the grid:filedistribution object
	gridFiledistribution := *grid.NewGridFiledistribution() // GridFiledistribution | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridFiledistributionAPI.GridfiledistributionReferencePut(context.Background(), reference).GridFiledistribution(gridFiledistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridFiledistributionAPI.GridfiledistributionReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridfiledistributionReferencePut`: UpdateGridFiledistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `GridFiledistributionAPI.GridfiledistributionReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:filedistribution object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridFiledistributionAPIGridfiledistributionReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridFiledistribution** | [**GridFiledistribution**](GridFiledistribution.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridFiledistributionResponse**](UpdateGridFiledistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

