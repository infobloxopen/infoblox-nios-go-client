# PxgridEndpointAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PxgridendpointGet**](PxgridEndpointAPI.md#PxgridendpointGet) | **Get** /pxgrid:endpoint | Retrieve pxgrid:endpoint objects
[**PxgridendpointPost**](PxgridEndpointAPI.md#PxgridendpointPost) | **Post** /pxgrid:endpoint | Create a pxgrid:endpoint object
[**PxgridendpointReferenceDelete**](PxgridEndpointAPI.md#PxgridendpointReferenceDelete) | **Delete** /pxgrid:endpoint/{reference} | Delete a pxgrid:endpoint object
[**PxgridendpointReferenceGet**](PxgridEndpointAPI.md#PxgridendpointReferenceGet) | **Get** /pxgrid:endpoint/{reference} | Get a specific pxgrid:endpoint object
[**PxgridendpointReferencePut**](PxgridEndpointAPI.md#PxgridendpointReferencePut) | **Put** /pxgrid:endpoint/{reference} | Update a pxgrid:endpoint object



## PxgridendpointGet

> ListPxgridEndpointResponse PxgridendpointGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve pxgrid:endpoint objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.PxgridEndpointAPI.PxgridendpointGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PxgridEndpointAPI.PxgridendpointGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PxgridendpointGet`: ListPxgridEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `PxgridEndpointAPI.PxgridendpointGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `PxgridEndpointAPIPxgridendpointGetRequest` struct via the builder pattern


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

[**ListPxgridEndpointResponse**](ListPxgridEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PxgridendpointPost

> CreatePxgridEndpointResponse PxgridendpointPost(ctx).PxgridEndpoint(pxgridEndpoint).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a pxgrid:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	pxgridEndpoint := *misc.NewPxgridEndpoint() // PxgridEndpoint | Object data to create

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.PxgridEndpointAPI.PxgridendpointPost(context.Background()).PxgridEndpoint(pxgridEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PxgridEndpointAPI.PxgridendpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PxgridendpointPost`: CreatePxgridEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `PxgridEndpointAPI.PxgridendpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `PxgridEndpointAPIPxgridendpointPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pxgridEndpoint** | [**PxgridEndpoint**](PxgridEndpoint.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreatePxgridEndpointResponse**](CreatePxgridEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PxgridendpointReferenceDelete

> PxgridendpointReferenceDelete(ctx, reference).Execute()

Delete a pxgrid:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the pxgrid:endpoint object

	apiClient := misc.NewAPIClient()
	r, err := apiClient.PxgridEndpointAPI.PxgridendpointReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PxgridEndpointAPI.PxgridendpointReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the pxgrid:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `PxgridEndpointAPIPxgridendpointReferenceDeleteRequest` struct via the builder pattern


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


## PxgridendpointReferenceGet

> GetPxgridEndpointResponse PxgridendpointReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific pxgrid:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the pxgrid:endpoint object

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.PxgridEndpointAPI.PxgridendpointReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PxgridEndpointAPI.PxgridendpointReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PxgridendpointReferenceGet`: GetPxgridEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `PxgridEndpointAPI.PxgridendpointReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the pxgrid:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `PxgridEndpointAPIPxgridendpointReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetPxgridEndpointResponse**](GetPxgridEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PxgridendpointReferencePut

> UpdatePxgridEndpointResponse PxgridendpointReferencePut(ctx, reference).PxgridEndpoint(pxgridEndpoint).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a pxgrid:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the pxgrid:endpoint object
	pxgridEndpoint := *misc.NewPxgridEndpoint() // PxgridEndpoint | Object data to update

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.PxgridEndpointAPI.PxgridendpointReferencePut(context.Background(), reference).PxgridEndpoint(pxgridEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PxgridEndpointAPI.PxgridendpointReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PxgridendpointReferencePut`: UpdatePxgridEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `PxgridEndpointAPI.PxgridendpointReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the pxgrid:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `PxgridEndpointAPIPxgridendpointReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**pxgridEndpoint** | [**PxgridEndpoint**](PxgridEndpoint.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdatePxgridEndpointResponse**](UpdatePxgridEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

