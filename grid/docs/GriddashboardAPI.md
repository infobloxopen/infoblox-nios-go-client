# GridDashboardAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GriddashboardGet**](GridDashboardAPI.md#GriddashboardGet) | **Get** /grid:dashboard | Retrieve grid:dashboard objects
[**GriddashboardReferenceGet**](GridDashboardAPI.md#GriddashboardReferenceGet) | **Get** /grid:dashboard/{reference} | Get a specific grid:dashboard object
[**GriddashboardReferencePut**](GridDashboardAPI.md#GriddashboardReferencePut) | **Put** /grid:dashboard/{reference} | Update a grid:dashboard object



## GriddashboardGet

> ListGridDashboardResponse GriddashboardGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:dashboard objects



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
	resp, r, err := apiClient.GridDashboardAPI.GriddashboardGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridDashboardAPI.GriddashboardGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GriddashboardGet`: ListGridDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `GridDashboardAPI.GriddashboardGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridDashboardAPIGriddashboardGetRequest` struct via the builder pattern


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

[**ListGridDashboardResponse**](ListGridDashboardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GriddashboardReferenceGet

> GetGridDashboardResponse GriddashboardReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:dashboard object



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
	reference := "reference_example" // string | Reference of the grid:dashboard object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridDashboardAPI.GriddashboardReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridDashboardAPI.GriddashboardReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GriddashboardReferenceGet`: GetGridDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `GridDashboardAPI.GriddashboardReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:dashboard object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridDashboardAPIGriddashboardReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridDashboardResponse**](GetGridDashboardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GriddashboardReferencePut

> UpdateGridDashboardResponse GriddashboardReferencePut(ctx, reference).GridDashboard(gridDashboard).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:dashboard object



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
	reference := "reference_example" // string | Reference of the grid:dashboard object
	gridDashboard := *grid.NewGridDashboard() // GridDashboard | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridDashboardAPI.GriddashboardReferencePut(context.Background(), reference).GridDashboard(gridDashboard).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridDashboardAPI.GriddashboardReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GriddashboardReferencePut`: UpdateGridDashboardResponse
	fmt.Fprintf(os.Stdout, "Response from `GridDashboardAPI.GriddashboardReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:dashboard object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridDashboardAPIGriddashboardReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridDashboard** | [**GridDashboard**](GridDashboard.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridDashboardResponse**](UpdateGridDashboardResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

