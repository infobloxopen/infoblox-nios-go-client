# ExtensibleattributedefAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](ExtensibleattributedefAPI.md#Get) | **Get** /extensibleattributedef | Retrieve extensibleattributedef objects
[**Post**](ExtensibleattributedefAPI.md#Post) | **Post** /extensibleattributedef | Create a extensibleattributedef object
[**ReferenceDelete**](ExtensibleattributedefAPI.md#ReferenceDelete) | **Delete** /extensibleattributedef/{reference} | Delete a extensibleattributedef object
[**ReferenceGet**](ExtensibleattributedefAPI.md#ReferenceGet) | **Get** /extensibleattributedef/{reference} | Get a specific extensibleattributedef object
[**ReferencePut**](ExtensibleattributedefAPI.md#ReferencePut) | **Put** /extensibleattributedef/{reference} | Update a extensibleattributedef object



## Get

> ListExtensibleattributedefResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve extensibleattributedef objects



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
	resp, r, err := apiClient.ExtensibleattributedefAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensibleattributedefAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListExtensibleattributedefResponse
	fmt.Fprintf(os.Stdout, "Response from `ExtensibleattributedefAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ExtensibleattributedefAPIGetRequest` struct via the builder pattern


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

[**ListExtensibleattributedefResponse**](ListExtensibleattributedefResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateExtensibleattributedefResponse Post(ctx).Extensibleattributedef(extensibleattributedef).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a extensibleattributedef object



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
	extensibleattributedef := *grid.NewExtensibleattributedef() // Extensibleattributedef | Object data to create

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.ExtensibleattributedefAPI.Post(context.Background()).Extensibleattributedef(extensibleattributedef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensibleattributedefAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateExtensibleattributedefResponse
	fmt.Fprintf(os.Stdout, "Response from `ExtensibleattributedefAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ExtensibleattributedefAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**extensibleattributedef** | [**Extensibleattributedef**](Extensibleattributedef.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateExtensibleattributedefResponse**](CreateExtensibleattributedefResponse.md)

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

Delete a extensibleattributedef object



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
	reference := "reference_example" // string | Reference of the extensibleattributedef object

	apiClient := grid.NewAPIClient()
	r, err := apiClient.ExtensibleattributedefAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensibleattributedefAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the extensibleattributedef object | 

### Other Parameters

Other parameters are passed through a pointer to a `ExtensibleattributedefAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetExtensibleattributedefResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific extensibleattributedef object



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
	reference := "reference_example" // string | Reference of the extensibleattributedef object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.ExtensibleattributedefAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensibleattributedefAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetExtensibleattributedefResponse
	fmt.Fprintf(os.Stdout, "Response from `ExtensibleattributedefAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the extensibleattributedef object | 

### Other Parameters

Other parameters are passed through a pointer to a `ExtensibleattributedefAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetExtensibleattributedefResponse**](GetExtensibleattributedefResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateExtensibleattributedefResponse ReferencePut(ctx, reference).Extensibleattributedef(extensibleattributedef).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a extensibleattributedef object



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
	reference := "reference_example" // string | Reference of the extensibleattributedef object
	extensibleattributedef := *grid.NewExtensibleattributedef() // Extensibleattributedef | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.ExtensibleattributedefAPI.ReferencePut(context.Background(), reference).Extensibleattributedef(extensibleattributedef).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExtensibleattributedefAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateExtensibleattributedefResponse
	fmt.Fprintf(os.Stdout, "Response from `ExtensibleattributedefAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the extensibleattributedef object | 

### Other Parameters

Other parameters are passed through a pointer to a `ExtensibleattributedefAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**extensibleattributedef** | [**Extensibleattributedef**](Extensibleattributedef.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateExtensibleattributedefResponse**](UpdateExtensibleattributedefResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

