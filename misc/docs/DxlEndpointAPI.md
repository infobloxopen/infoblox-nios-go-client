# DxlEndpointAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DxlendpointGet**](DxlEndpointAPI.md#DxlendpointGet) | **Get** /dxl:endpoint | Retrieve dxl:endpoint objects
[**DxlendpointPost**](DxlEndpointAPI.md#DxlendpointPost) | **Post** /dxl:endpoint | Create a dxl:endpoint object
[**DxlendpointReferenceDelete**](DxlEndpointAPI.md#DxlendpointReferenceDelete) | **Delete** /dxl:endpoint/{reference} | Delete a dxl:endpoint object
[**DxlendpointReferenceGet**](DxlEndpointAPI.md#DxlendpointReferenceGet) | **Get** /dxl:endpoint/{reference} | Get a specific dxl:endpoint object
[**DxlendpointReferencePut**](DxlEndpointAPI.md#DxlendpointReferencePut) | **Put** /dxl:endpoint/{reference} | Update a dxl:endpoint object



## DxlendpointGet

> ListDxlEndpointResponse DxlendpointGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dxl:endpoint objects



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
	resp, r, err := apiClient.DxlEndpointAPI.DxlendpointGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DxlEndpointAPI.DxlendpointGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DxlendpointGet`: ListDxlEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `DxlEndpointAPI.DxlendpointGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DxlEndpointAPIDxlendpointGetRequest` struct via the builder pattern


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

[**ListDxlEndpointResponse**](ListDxlEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DxlendpointPost

> CreateDxlEndpointResponse DxlendpointPost(ctx).DxlEndpoint(dxlEndpoint).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dxl:endpoint object



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
	dxlEndpoint := *misc.NewDxlEndpoint() // DxlEndpoint | Object data to create

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.DxlEndpointAPI.DxlendpointPost(context.Background()).DxlEndpoint(dxlEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DxlEndpointAPI.DxlendpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DxlendpointPost`: CreateDxlEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `DxlEndpointAPI.DxlendpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DxlEndpointAPIDxlendpointPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dxlEndpoint** | [**DxlEndpoint**](DxlEndpoint.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDxlEndpointResponse**](CreateDxlEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DxlendpointReferenceDelete

> DxlendpointReferenceDelete(ctx, reference).Execute()

Delete a dxl:endpoint object



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
	reference := "reference_example" // string | Reference of the dxl:endpoint object

	apiClient := misc.NewAPIClient()
	r, err := apiClient.DxlEndpointAPI.DxlendpointReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DxlEndpointAPI.DxlendpointReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dxl:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `DxlEndpointAPIDxlendpointReferenceDeleteRequest` struct via the builder pattern


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


## DxlendpointReferenceGet

> GetDxlEndpointResponse DxlendpointReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dxl:endpoint object



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
	reference := "reference_example" // string | Reference of the dxl:endpoint object

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.DxlEndpointAPI.DxlendpointReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DxlEndpointAPI.DxlendpointReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DxlendpointReferenceGet`: GetDxlEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `DxlEndpointAPI.DxlendpointReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dxl:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `DxlEndpointAPIDxlendpointReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDxlEndpointResponse**](GetDxlEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DxlendpointReferencePut

> UpdateDxlEndpointResponse DxlendpointReferencePut(ctx, reference).DxlEndpoint(dxlEndpoint).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dxl:endpoint object



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
	reference := "reference_example" // string | Reference of the dxl:endpoint object
	dxlEndpoint := *misc.NewDxlEndpoint() // DxlEndpoint | Object data to update

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.DxlEndpointAPI.DxlendpointReferencePut(context.Background(), reference).DxlEndpoint(dxlEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DxlEndpointAPI.DxlendpointReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DxlendpointReferencePut`: UpdateDxlEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `DxlEndpointAPI.DxlendpointReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dxl:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `DxlEndpointAPIDxlendpointReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dxlEndpoint** | [**DxlEndpoint**](DxlEndpoint.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDxlEndpointResponse**](UpdateDxlEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

