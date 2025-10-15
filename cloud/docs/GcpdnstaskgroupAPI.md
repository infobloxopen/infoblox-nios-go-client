# GcpdnstaskgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](GcpdnstaskgroupAPI.md#Create) | **Post** /gcpdnstaskgroup | Create a gcpdnstaskgroup object
[**Delete**](GcpdnstaskgroupAPI.md#Delete) | **Delete** /gcpdnstaskgroup/{reference} | Delete a gcpdnstaskgroup object
[**List**](GcpdnstaskgroupAPI.md#List) | **Get** /gcpdnstaskgroup | Retrieve gcpdnstaskgroup objects
[**Read**](GcpdnstaskgroupAPI.md#Read) | **Get** /gcpdnstaskgroup/{reference} | Get a specific gcpdnstaskgroup object
[**Update**](GcpdnstaskgroupAPI.md#Update) | **Put** /gcpdnstaskgroup/{reference} | Update a gcpdnstaskgroup object



## Create

> CreateGcpdnstaskgroupResponse Create(ctx).Gcpdnstaskgroup(gcpdnstaskgroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a gcpdnstaskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	gcpdnstaskgroup := *cloud.NewGcpdnstaskgroup() // Gcpdnstaskgroup | Object data to create

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GcpdnstaskgroupAPI.Create(context.Background()).Gcpdnstaskgroup(gcpdnstaskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateGcpdnstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GcpdnstaskgroupAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gcpdnstaskgroup** | [**Gcpdnstaskgroup**](Gcpdnstaskgroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateGcpdnstaskgroupResponse**](CreateGcpdnstaskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, reference).Execute()

Delete a gcpdnstaskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the gcpdnstaskgroup object

	apiClient := cloud.NewAPIClient()
	r, err := apiClient.GcpdnstaskgroupAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the gcpdnstaskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPIDeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListGcpdnstaskgroupResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve gcpdnstaskgroup objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GcpdnstaskgroupAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListGcpdnstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GcpdnstaskgroupAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPIListRequest` struct via the builder pattern


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

[**ListGcpdnstaskgroupResponse**](ListGcpdnstaskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetGcpdnstaskgroupResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific gcpdnstaskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the gcpdnstaskgroup object

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GcpdnstaskgroupAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetGcpdnstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GcpdnstaskgroupAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the gcpdnstaskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetGcpdnstaskgroupResponse**](GetGcpdnstaskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateGcpdnstaskgroupResponse Update(ctx, reference).Gcpdnstaskgroup(gcpdnstaskgroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a gcpdnstaskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the gcpdnstaskgroup object
	gcpdnstaskgroup := *cloud.NewGcpdnstaskgroup() // Gcpdnstaskgroup | Object data to update

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GcpdnstaskgroupAPI.Update(context.Background(), reference).Gcpdnstaskgroup(gcpdnstaskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateGcpdnstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GcpdnstaskgroupAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the gcpdnstaskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gcpdnstaskgroup** | [**Gcpdnstaskgroup**](Gcpdnstaskgroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGcpdnstaskgroupResponse**](UpdateGcpdnstaskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

