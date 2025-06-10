# GridServicerestartGroupOrderAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GridservicerestartgrouporderGet**](GridServicerestartGroupOrderAPI.md#GridservicerestartgrouporderGet) | **Get** /grid:servicerestart:group:order | Retrieve grid:servicerestart:group:order objects
[**GridservicerestartgrouporderPost**](GridServicerestartGroupOrderAPI.md#GridservicerestartgrouporderPost) | **Post** /grid:servicerestart:group:order | Create a grid:servicerestart:group:order object
[**GridservicerestartgrouporderReferenceGet**](GridServicerestartGroupOrderAPI.md#GridservicerestartgrouporderReferenceGet) | **Get** /grid:servicerestart:group:order/{reference} | Get a specific grid:servicerestart:group:order object
[**GridservicerestartgrouporderReferencePut**](GridServicerestartGroupOrderAPI.md#GridservicerestartgrouporderReferencePut) | **Put** /grid:servicerestart:group:order/{reference} | Update a grid:servicerestart:group:order object



## GridservicerestartgrouporderGet

> ListGridServicerestartGroupOrderResponse GridservicerestartgrouporderGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:servicerestart:group:order objects



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
	resp, r, err := apiClient.GridServicerestartGroupOrderAPI.GridservicerestartgrouporderGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupOrderAPI.GridservicerestartgrouporderGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridservicerestartgrouporderGet`: ListGridServicerestartGroupOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupOrderAPI.GridservicerestartgrouporderGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupOrderAPIGridservicerestartgrouporderGetRequest` struct via the builder pattern


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

[**ListGridServicerestartGroupOrderResponse**](ListGridServicerestartGroupOrderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridservicerestartgrouporderPost

> CreateGridServicerestartGroupOrderResponse GridservicerestartgrouporderPost(ctx).GridServicerestartGroupOrder(gridServicerestartGroupOrder).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a grid:servicerestart:group:order object



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
	gridServicerestartGroupOrder := *grid.NewGridServicerestartGroupOrder() // GridServicerestartGroupOrder | Object data to create

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridServicerestartGroupOrderAPI.GridservicerestartgrouporderPost(context.Background()).GridServicerestartGroupOrder(gridServicerestartGroupOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupOrderAPI.GridservicerestartgrouporderPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridservicerestartgrouporderPost`: CreateGridServicerestartGroupOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupOrderAPI.GridservicerestartgrouporderPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupOrderAPIGridservicerestartgrouporderPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridServicerestartGroupOrder** | [**GridServicerestartGroupOrder**](GridServicerestartGroupOrder.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateGridServicerestartGroupOrderResponse**](CreateGridServicerestartGroupOrderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridservicerestartgrouporderReferenceGet

> GetGridServicerestartGroupOrderResponse GridservicerestartgrouporderReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:servicerestart:group:order object



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
	reference := "reference_example" // string | Reference of the grid:servicerestart:group:order object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridServicerestartGroupOrderAPI.GridservicerestartgrouporderReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupOrderAPI.GridservicerestartgrouporderReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridservicerestartgrouporderReferenceGet`: GetGridServicerestartGroupOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupOrderAPI.GridservicerestartgrouporderReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:servicerestart:group:order object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupOrderAPIGridservicerestartgrouporderReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridServicerestartGroupOrderResponse**](GetGridServicerestartGroupOrderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridservicerestartgrouporderReferencePut

> UpdateGridServicerestartGroupOrderResponse GridservicerestartgrouporderReferencePut(ctx, reference).GridServicerestartGroupOrder(gridServicerestartGroupOrder).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:servicerestart:group:order object



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
	reference := "reference_example" // string | Reference of the grid:servicerestart:group:order object
	gridServicerestartGroupOrder := *grid.NewGridServicerestartGroupOrder() // GridServicerestartGroupOrder | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridServicerestartGroupOrderAPI.GridservicerestartgrouporderReferencePut(context.Background(), reference).GridServicerestartGroupOrder(gridServicerestartGroupOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupOrderAPI.GridservicerestartgrouporderReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridservicerestartgrouporderReferencePut`: UpdateGridServicerestartGroupOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupOrderAPI.GridservicerestartgrouporderReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:servicerestart:group:order object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupOrderAPIGridservicerestartgrouporderReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridServicerestartGroupOrder** | [**GridServicerestartGroupOrder**](GridServicerestartGroupOrder.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridServicerestartGroupOrderResponse**](UpdateGridServicerestartGroupOrderResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

