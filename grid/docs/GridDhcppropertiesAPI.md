# GridDhcppropertiesAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GriddhcppropertiesGet**](GridDhcppropertiesAPI.md#GriddhcppropertiesGet) | **Get** /grid:dhcpproperties | Retrieve grid:dhcpproperties objects
[**GriddhcppropertiesReferenceGet**](GridDhcppropertiesAPI.md#GriddhcppropertiesReferenceGet) | **Get** /grid:dhcpproperties/{reference} | Get a specific grid:dhcpproperties object
[**GriddhcppropertiesReferencePut**](GridDhcppropertiesAPI.md#GriddhcppropertiesReferencePut) | **Put** /grid:dhcpproperties/{reference} | Update a grid:dhcpproperties object



## GriddhcppropertiesGet

> ListGridDhcppropertiesResponse GriddhcppropertiesGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:dhcpproperties objects



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
	resp, r, err := apiClient.GridDhcppropertiesAPI.GriddhcppropertiesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridDhcppropertiesAPI.GriddhcppropertiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GriddhcppropertiesGet`: ListGridDhcppropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `GridDhcppropertiesAPI.GriddhcppropertiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridDhcppropertiesAPIGriddhcppropertiesGetRequest` struct via the builder pattern


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

[**ListGridDhcppropertiesResponse**](ListGridDhcppropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GriddhcppropertiesReferenceGet

> GetGridDhcppropertiesResponse GriddhcppropertiesReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:dhcpproperties object



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
	reference := "reference_example" // string | Reference of the grid:dhcpproperties object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridDhcppropertiesAPI.GriddhcppropertiesReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridDhcppropertiesAPI.GriddhcppropertiesReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GriddhcppropertiesReferenceGet`: GetGridDhcppropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `GridDhcppropertiesAPI.GriddhcppropertiesReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:dhcpproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridDhcppropertiesAPIGriddhcppropertiesReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridDhcppropertiesResponse**](GetGridDhcppropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GriddhcppropertiesReferencePut

> UpdateGridDhcppropertiesResponse GriddhcppropertiesReferencePut(ctx, reference).GridDhcpproperties(gridDhcpproperties).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:dhcpproperties object



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
	reference := "reference_example" // string | Reference of the grid:dhcpproperties object
	gridDhcpproperties := *grid.NewGridDhcpproperties() // GridDhcpproperties | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridDhcppropertiesAPI.GriddhcppropertiesReferencePut(context.Background(), reference).GridDhcpproperties(gridDhcpproperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridDhcppropertiesAPI.GriddhcppropertiesReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GriddhcppropertiesReferencePut`: UpdateGridDhcppropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `GridDhcppropertiesAPI.GriddhcppropertiesReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:dhcpproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridDhcppropertiesAPIGriddhcppropertiesReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridDhcpproperties** | [**GridDhcpproperties**](GridDhcpproperties.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridDhcppropertiesResponse**](UpdateGridDhcppropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

