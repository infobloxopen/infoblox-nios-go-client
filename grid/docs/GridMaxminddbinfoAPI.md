# GridMaxminddbinfoAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GridmaxminddbinfoGet**](GridMaxminddbinfoAPI.md#GridmaxminddbinfoGet) | **Get** /grid:maxminddbinfo | Retrieve grid:maxminddbinfo objects
[**GridmaxminddbinfoReferenceGet**](GridMaxminddbinfoAPI.md#GridmaxminddbinfoReferenceGet) | **Get** /grid:maxminddbinfo/{reference} | Get a specific grid:maxminddbinfo object



## GridmaxminddbinfoGet

> ListGridMaxminddbinfoResponse GridmaxminddbinfoGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:maxminddbinfo objects



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
	resp, r, err := apiClient.GridMaxminddbinfoAPI.GridmaxminddbinfoGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridMaxminddbinfoAPI.GridmaxminddbinfoGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridmaxminddbinfoGet`: ListGridMaxminddbinfoResponse
	fmt.Fprintf(os.Stdout, "Response from `GridMaxminddbinfoAPI.GridmaxminddbinfoGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridMaxminddbinfoAPIGridmaxminddbinfoGetRequest` struct via the builder pattern


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

[**ListGridMaxminddbinfoResponse**](ListGridMaxminddbinfoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridmaxminddbinfoReferenceGet

> GetGridMaxminddbinfoResponse GridmaxminddbinfoReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:maxminddbinfo object



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
	reference := "reference_example" // string | Reference of the grid:maxminddbinfo object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridMaxminddbinfoAPI.GridmaxminddbinfoReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridMaxminddbinfoAPI.GridmaxminddbinfoReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridmaxminddbinfoReferenceGet`: GetGridMaxminddbinfoResponse
	fmt.Fprintf(os.Stdout, "Response from `GridMaxminddbinfoAPI.GridmaxminddbinfoReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:maxminddbinfo object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridMaxminddbinfoAPIGridmaxminddbinfoReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridMaxminddbinfoResponse**](GetGridMaxminddbinfoResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

