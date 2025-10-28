# GridServicerestartGroupOrderAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](GridServicerestartGroupOrderAPI.md#Create) | **Post** /grid:servicerestart:group:order | Create a grid:servicerestart:group:order object
[**List**](GridServicerestartGroupOrderAPI.md#List) | **Get** /grid:servicerestart:group:order | Retrieve grid:servicerestart:group:order objects
[**Read**](GridServicerestartGroupOrderAPI.md#Read) | **Get** /grid:servicerestart:group:order/{reference} | Get a specific grid:servicerestart:group:order object
[**Update**](GridServicerestartGroupOrderAPI.md#Update) | **Put** /grid:servicerestart:group:order/{reference} | Update a grid:servicerestart:group:order object



## Create

> CreateGridServicerestartGroupOrderResponse Create(ctx).GridServicerestartGroupOrder(gridServicerestartGroupOrder).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a grid:servicerestart:group:order object



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
	gridServicerestartGroupOrder := *grid.NewGridServicerestartGroupOrder() // GridServicerestartGroupOrder | Object data to create

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridServicerestartGroupOrderAPI.Create(context.Background()).GridServicerestartGroupOrder(gridServicerestartGroupOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupOrderAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateGridServicerestartGroupOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupOrderAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupOrderAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridServicerestartGroupOrder** | [**GridServicerestartGroupOrder**](GridServicerestartGroupOrder.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## List

> ListGridServicerestartGroupOrderResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve grid:servicerestart:group:order objects



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
	resp, r, err := apiClient.GridServicerestartGroupOrderAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupOrderAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListGridServicerestartGroupOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupOrderAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupOrderAPIListRequest` struct via the builder pattern


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
**proxySearch** | **string** | Search Grid members for data | 

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


## Read

> GetGridServicerestartGroupOrderResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific grid:servicerestart:group:order object



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
	reference := "reference_example" // string | Reference of the grid:servicerestart:group:order object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridServicerestartGroupOrderAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupOrderAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetGridServicerestartGroupOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupOrderAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:servicerestart:group:order object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupOrderAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

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


## Update

> UpdateGridServicerestartGroupOrderResponse Update(ctx, reference).GridServicerestartGroupOrder(gridServicerestartGroupOrder).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a grid:servicerestart:group:order object



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
	reference := "reference_example" // string | Reference of the grid:servicerestart:group:order object
	gridServicerestartGroupOrder := *grid.NewGridServicerestartGroupOrder() // GridServicerestartGroupOrder | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridServicerestartGroupOrderAPI.Update(context.Background(), reference).GridServicerestartGroupOrder(gridServicerestartGroupOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupOrderAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateGridServicerestartGroupOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupOrderAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:servicerestart:group:order object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupOrderAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridServicerestartGroupOrder** | [**GridServicerestartGroupOrder**](GridServicerestartGroupOrder.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

