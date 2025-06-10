# SharedrecordMxAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharedrecordmxGet**](SharedrecordMxAPI.md#SharedrecordmxGet) | **Get** /sharedrecord:mx | Retrieve sharedrecord:mx objects
[**SharedrecordmxPost**](SharedrecordMxAPI.md#SharedrecordmxPost) | **Post** /sharedrecord:mx | Create a sharedrecord:mx object
[**SharedrecordmxReferenceDelete**](SharedrecordMxAPI.md#SharedrecordmxReferenceDelete) | **Delete** /sharedrecord:mx/{reference} | Delete a sharedrecord:mx object
[**SharedrecordmxReferenceGet**](SharedrecordMxAPI.md#SharedrecordmxReferenceGet) | **Get** /sharedrecord:mx/{reference} | Get a specific sharedrecord:mx object
[**SharedrecordmxReferencePut**](SharedrecordMxAPI.md#SharedrecordmxReferencePut) | **Put** /sharedrecord:mx/{reference} | Update a sharedrecord:mx object



## SharedrecordmxGet

> ListSharedrecordMxResponse SharedrecordmxGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve sharedrecord:mx objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordMxAPI.SharedrecordmxGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordMxAPI.SharedrecordmxGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordmxGet`: ListSharedrecordMxResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordMxAPI.SharedrecordmxGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordMxAPISharedrecordmxGetRequest` struct via the builder pattern


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

[**ListSharedrecordMxResponse**](ListSharedrecordMxResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordmxPost

> CreateSharedrecordMxResponse SharedrecordmxPost(ctx).SharedrecordMx(sharedrecordMx).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a sharedrecord:mx object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	sharedrecordMx := *dns.NewSharedrecordMx() // SharedrecordMx | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordMxAPI.SharedrecordmxPost(context.Background()).SharedrecordMx(sharedrecordMx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordMxAPI.SharedrecordmxPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordmxPost`: CreateSharedrecordMxResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordMxAPI.SharedrecordmxPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordMxAPISharedrecordmxPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordMx** | [**SharedrecordMx**](SharedrecordMx.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSharedrecordMxResponse**](CreateSharedrecordMxResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordmxReferenceDelete

> SharedrecordmxReferenceDelete(ctx, reference).Execute()

Delete a sharedrecord:mx object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the sharedrecord:mx object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.SharedrecordMxAPI.SharedrecordmxReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordMxAPI.SharedrecordmxReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:mx object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordMxAPISharedrecordmxReferenceDeleteRequest` struct via the builder pattern


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


## SharedrecordmxReferenceGet

> GetSharedrecordMxResponse SharedrecordmxReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific sharedrecord:mx object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the sharedrecord:mx object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordMxAPI.SharedrecordmxReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordMxAPI.SharedrecordmxReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordmxReferenceGet`: GetSharedrecordMxResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordMxAPI.SharedrecordmxReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:mx object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordMxAPISharedrecordmxReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSharedrecordMxResponse**](GetSharedrecordMxResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordmxReferencePut

> UpdateSharedrecordMxResponse SharedrecordmxReferencePut(ctx, reference).SharedrecordMx(sharedrecordMx).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a sharedrecord:mx object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the sharedrecord:mx object
	sharedrecordMx := *dns.NewSharedrecordMx() // SharedrecordMx | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordMxAPI.SharedrecordmxReferencePut(context.Background(), reference).SharedrecordMx(sharedrecordMx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordMxAPI.SharedrecordmxReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordmxReferencePut`: UpdateSharedrecordMxResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordMxAPI.SharedrecordmxReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:mx object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordMxAPISharedrecordmxReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordMx** | [**SharedrecordMx**](SharedrecordMx.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSharedrecordMxResponse**](UpdateSharedrecordMxResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

