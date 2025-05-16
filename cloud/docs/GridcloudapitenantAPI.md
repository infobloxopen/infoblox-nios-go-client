# GridcloudapitenantAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](GridcloudapitenantAPI.md#Get) | **Get** /grid:cloudapi:tenant | Retrieve grid:cloudapi:tenant objects
[**ReferenceGet**](GridcloudapitenantAPI.md#ReferenceGet) | **Get** /grid:cloudapi:tenant/{reference} | Get a specific grid:cloudapi:tenant object
[**ReferencePut**](GridcloudapitenantAPI.md#ReferencePut) | **Put** /grid:cloudapi:tenant/{reference} | Update a grid:cloudapi:tenant object



## Get

> ListGridCloudapiTenantResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:cloudapi:tenant objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GridcloudapitenantAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridcloudapitenantAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListGridCloudapiTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `GridcloudapitenantAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridcloudapitenantAPIGetRequest` struct via the builder pattern


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

[**ListGridCloudapiTenantResponse**](ListGridCloudapiTenantResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceGet

> GetGridCloudapiTenantResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:cloudapi:tenant object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the grid:cloudapi:tenant object

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GridcloudapitenantAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridcloudapitenantAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetGridCloudapiTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `GridcloudapitenantAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:cloudapi:tenant object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridcloudapitenantAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridCloudapiTenantResponse**](GetGridCloudapiTenantResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateGridCloudapiTenantResponse ReferencePut(ctx, reference).GridCloudapiTenant(gridCloudapiTenant).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:cloudapi:tenant object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the grid:cloudapi:tenant object
	gridCloudapiTenant := *cloud.NewGridCloudapiTenant() // GridCloudapiTenant | Object data to update

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GridcloudapitenantAPI.ReferencePut(context.Background(), reference).GridCloudapiTenant(gridCloudapiTenant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridcloudapitenantAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateGridCloudapiTenantResponse
	fmt.Fprintf(os.Stdout, "Response from `GridcloudapitenantAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:cloudapi:tenant object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridcloudapitenantAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridCloudapiTenant** | [**GridCloudapiTenant**](GridCloudapiTenant.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridCloudapiTenantResponse**](UpdateGridCloudapiTenantResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

