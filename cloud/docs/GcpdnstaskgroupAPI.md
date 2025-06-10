# GcpdnstaskgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](GcpdnstaskgroupAPI.md#Get) | **Get** /gcpdnstaskgroup | Retrieve gcpdnstaskgroup objects
[**Post**](GcpdnstaskgroupAPI.md#Post) | **Post** /gcpdnstaskgroup | Create a gcpdnstaskgroup object
[**ReferenceDelete**](GcpdnstaskgroupAPI.md#ReferenceDelete) | **Delete** /gcpdnstaskgroup/{reference} | Delete a gcpdnstaskgroup object
[**ReferenceGet**](GcpdnstaskgroupAPI.md#ReferenceGet) | **Get** /gcpdnstaskgroup/{reference} | Get a specific gcpdnstaskgroup object
[**ReferencePut**](GcpdnstaskgroupAPI.md#ReferencePut) | **Put** /gcpdnstaskgroup/{reference} | Update a gcpdnstaskgroup object



## Get

> ListGcpdnstaskgroupResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve gcpdnstaskgroup objects



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
	resp, r, err := apiClient.GcpdnstaskgroupAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListGcpdnstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GcpdnstaskgroupAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPIGetRequest` struct via the builder pattern


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

[**ListGcpdnstaskgroupResponse**](ListGcpdnstaskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateGcpdnstaskgroupResponse Post(ctx).Gcpdnstaskgroup(gcpdnstaskgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a gcpdnstaskgroup object



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
	gcpdnstaskgroup := *cloud.NewGcpdnstaskgroup() // Gcpdnstaskgroup | Object data to create

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GcpdnstaskgroupAPI.Post(context.Background()).Gcpdnstaskgroup(gcpdnstaskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateGcpdnstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GcpdnstaskgroupAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gcpdnstaskgroup** | [**Gcpdnstaskgroup**](Gcpdnstaskgroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

Delete a gcpdnstaskgroup object



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
	reference := "reference_example" // string | Reference of the gcpdnstaskgroup object

	apiClient := cloud.NewAPIClient()
	r, err := apiClient.GcpdnstaskgroupAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPIReferenceDeleteRequest` struct via the builder pattern


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


## ReferenceGet

> GetGcpdnstaskgroupResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific gcpdnstaskgroup object



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
	reference := "reference_example" // string | Reference of the gcpdnstaskgroup object

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GcpdnstaskgroupAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetGcpdnstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GcpdnstaskgroupAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the gcpdnstaskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## ReferencePut

> UpdateGcpdnstaskgroupResponse ReferencePut(ctx, reference).Gcpdnstaskgroup(gcpdnstaskgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a gcpdnstaskgroup object



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
	reference := "reference_example" // string | Reference of the gcpdnstaskgroup object
	gcpdnstaskgroup := *cloud.NewGcpdnstaskgroup() // Gcpdnstaskgroup | Object data to update

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.GcpdnstaskgroupAPI.ReferencePut(context.Background(), reference).Gcpdnstaskgroup(gcpdnstaskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GcpdnstaskgroupAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateGcpdnstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GcpdnstaskgroupAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the gcpdnstaskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `GcpdnstaskgroupAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gcpdnstaskgroup** | [**Gcpdnstaskgroup**](Gcpdnstaskgroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

