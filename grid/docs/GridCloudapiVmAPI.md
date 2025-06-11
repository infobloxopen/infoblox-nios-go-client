# GridCloudapiVmAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GridcloudapivmGet**](GridCloudapiVmAPI.md#GridcloudapivmGet) | **Get** /grid:cloudapi:vm | Retrieve grid:cloudapi:vm objects
[**GridcloudapivmReferenceGet**](GridCloudapiVmAPI.md#GridcloudapivmReferenceGet) | **Get** /grid:cloudapi:vm/{reference} | Get a specific grid:cloudapi:vm object
[**GridcloudapivmReferencePut**](GridCloudapiVmAPI.md#GridcloudapivmReferencePut) | **Put** /grid:cloudapi:vm/{reference} | Update a grid:cloudapi:vm object



## GridcloudapivmGet

> ListGridCloudapiVmResponse GridcloudapivmGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:cloudapi:vm objects



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
	resp, r, err := apiClient.GridCloudapiVmAPI.GridcloudapivmGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridCloudapiVmAPI.GridcloudapivmGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridcloudapivmGet`: ListGridCloudapiVmResponse
	fmt.Fprintf(os.Stdout, "Response from `GridCloudapiVmAPI.GridcloudapivmGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridCloudapiVmAPIGridcloudapivmGetRequest` struct via the builder pattern


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

[**ListGridCloudapiVmResponse**](ListGridCloudapiVmResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridcloudapivmReferenceGet

> GetGridCloudapiVmResponse GridcloudapivmReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:cloudapi:vm object



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
	reference := "reference_example" // string | Reference of the grid:cloudapi:vm object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridCloudapiVmAPI.GridcloudapivmReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridCloudapiVmAPI.GridcloudapivmReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridcloudapivmReferenceGet`: GetGridCloudapiVmResponse
	fmt.Fprintf(os.Stdout, "Response from `GridCloudapiVmAPI.GridcloudapivmReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:cloudapi:vm object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridCloudapiVmAPIGridcloudapivmReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridCloudapiVmResponse**](GetGridCloudapiVmResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridcloudapivmReferencePut

> UpdateGridCloudapiVmResponse GridcloudapivmReferencePut(ctx, reference).GridCloudapiVm(gridCloudapiVm).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:cloudapi:vm object



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
	reference := "reference_example" // string | Reference of the grid:cloudapi:vm object
	gridCloudapiVm := *grid.NewGridCloudapiVm() // GridCloudapiVm | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridCloudapiVmAPI.GridcloudapivmReferencePut(context.Background(), reference).GridCloudapiVm(gridCloudapiVm).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridCloudapiVmAPI.GridcloudapivmReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridcloudapivmReferencePut`: UpdateGridCloudapiVmResponse
	fmt.Fprintf(os.Stdout, "Response from `GridCloudapiVmAPI.GridcloudapivmReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:cloudapi:vm object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridCloudapiVmAPIGridcloudapivmReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridCloudapiVm** | [**GridCloudapiVm**](GridCloudapiVm.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridCloudapiVmResponse**](UpdateGridCloudapiVmResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

