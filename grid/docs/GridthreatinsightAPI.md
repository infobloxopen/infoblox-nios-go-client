# GridThreatinsightAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GridthreatinsightGet**](GridThreatinsightAPI.md#GridthreatinsightGet) | **Get** /grid:threatinsight | Retrieve grid:threatinsight objects
[**GridthreatinsightReferenceGet**](GridThreatinsightAPI.md#GridthreatinsightReferenceGet) | **Get** /grid:threatinsight/{reference} | Get a specific grid:threatinsight object
[**GridthreatinsightReferencePut**](GridThreatinsightAPI.md#GridthreatinsightReferencePut) | **Put** /grid:threatinsight/{reference} | Update a grid:threatinsight object



## GridthreatinsightGet

> ListGridThreatinsightResponse GridthreatinsightGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:threatinsight objects



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
	resp, r, err := apiClient.GridThreatinsightAPI.GridthreatinsightGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridThreatinsightAPI.GridthreatinsightGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridthreatinsightGet`: ListGridThreatinsightResponse
	fmt.Fprintf(os.Stdout, "Response from `GridThreatinsightAPI.GridthreatinsightGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridThreatinsightAPIGridthreatinsightGetRequest` struct via the builder pattern


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

[**ListGridThreatinsightResponse**](ListGridThreatinsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridthreatinsightReferenceGet

> GetGridThreatinsightResponse GridthreatinsightReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:threatinsight object



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
	reference := "reference_example" // string | Reference of the grid:threatinsight object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridThreatinsightAPI.GridthreatinsightReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridThreatinsightAPI.GridthreatinsightReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridthreatinsightReferenceGet`: GetGridThreatinsightResponse
	fmt.Fprintf(os.Stdout, "Response from `GridThreatinsightAPI.GridthreatinsightReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:threatinsight object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridThreatinsightAPIGridthreatinsightReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridThreatinsightResponse**](GetGridThreatinsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridthreatinsightReferencePut

> UpdateGridThreatinsightResponse GridthreatinsightReferencePut(ctx, reference).GridThreatinsight(gridThreatinsight).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:threatinsight object



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
	reference := "reference_example" // string | Reference of the grid:threatinsight object
	gridThreatinsight := *grid.NewGridThreatinsight() // GridThreatinsight | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridThreatinsightAPI.GridthreatinsightReferencePut(context.Background(), reference).GridThreatinsight(gridThreatinsight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridThreatinsightAPI.GridthreatinsightReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridthreatinsightReferencePut`: UpdateGridThreatinsightResponse
	fmt.Fprintf(os.Stdout, "Response from `GridThreatinsightAPI.GridthreatinsightReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:threatinsight object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridThreatinsightAPIGridthreatinsightReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridThreatinsight** | [**GridThreatinsight**](GridThreatinsight.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridThreatinsightResponse**](UpdateGridThreatinsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

