# GridThreatprotectionAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GridthreatprotectionGet**](GridThreatprotectionAPI.md#GridthreatprotectionGet) | **Get** /grid:threatprotection | Retrieve grid:threatprotection objects
[**GridthreatprotectionReferenceGet**](GridThreatprotectionAPI.md#GridthreatprotectionReferenceGet) | **Get** /grid:threatprotection/{reference} | Get a specific grid:threatprotection object
[**GridthreatprotectionReferencePut**](GridThreatprotectionAPI.md#GridthreatprotectionReferencePut) | **Put** /grid:threatprotection/{reference} | Update a grid:threatprotection object



## GridthreatprotectionGet

> ListGridThreatprotectionResponse GridthreatprotectionGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:threatprotection objects



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
	resp, r, err := apiClient.GridThreatprotectionAPI.GridthreatprotectionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridThreatprotectionAPI.GridthreatprotectionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridthreatprotectionGet`: ListGridThreatprotectionResponse
	fmt.Fprintf(os.Stdout, "Response from `GridThreatprotectionAPI.GridthreatprotectionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridThreatprotectionAPIGridthreatprotectionGetRequest` struct via the builder pattern


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

[**ListGridThreatprotectionResponse**](ListGridThreatprotectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridthreatprotectionReferenceGet

> GetGridThreatprotectionResponse GridthreatprotectionReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:threatprotection object



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
	reference := "reference_example" // string | Reference of the grid:threatprotection object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridThreatprotectionAPI.GridthreatprotectionReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridThreatprotectionAPI.GridthreatprotectionReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridthreatprotectionReferenceGet`: GetGridThreatprotectionResponse
	fmt.Fprintf(os.Stdout, "Response from `GridThreatprotectionAPI.GridthreatprotectionReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:threatprotection object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridThreatprotectionAPIGridthreatprotectionReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## GridthreatprotectionReferencePut

> UpdateGridThreatprotectionResponse GridthreatprotectionReferencePut(ctx, reference).GridThreatprotection(gridThreatprotection).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:threatprotection object



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
	reference := "reference_example" // string | Reference of the grid:threatprotection object
	gridThreatprotection := *grid.NewGridThreatprotection() // GridThreatprotection | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridThreatprotectionAPI.GridthreatprotectionReferencePut(context.Background(), reference).GridThreatprotection(gridThreatprotection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridThreatprotectionAPI.GridthreatprotectionReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridthreatprotectionReferencePut`: UpdateGridThreatprotectionResponse
	fmt.Fprintf(os.Stdout, "Response from `GridThreatprotectionAPI.GridthreatprotectionReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:threatprotection object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridThreatprotectionAPIGridthreatprotectionReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridThreatprotection** | [**GridThreatprotection**](GridThreatprotection.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

