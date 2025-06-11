# SharedrecordAaaaAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharedrecordaaaaGet**](SharedrecordAaaaAPI.md#SharedrecordaaaaGet) | **Get** /sharedrecord:aaaa | Retrieve sharedrecord:aaaa objects
[**SharedrecordaaaaPost**](SharedrecordAaaaAPI.md#SharedrecordaaaaPost) | **Post** /sharedrecord:aaaa | Create a sharedrecord:aaaa object
[**SharedrecordaaaaReferenceDelete**](SharedrecordAaaaAPI.md#SharedrecordaaaaReferenceDelete) | **Delete** /sharedrecord:aaaa/{reference} | Delete a sharedrecord:aaaa object
[**SharedrecordaaaaReferenceGet**](SharedrecordAaaaAPI.md#SharedrecordaaaaReferenceGet) | **Get** /sharedrecord:aaaa/{reference} | Get a specific sharedrecord:aaaa object
[**SharedrecordaaaaReferencePut**](SharedrecordAaaaAPI.md#SharedrecordaaaaReferencePut) | **Put** /sharedrecord:aaaa/{reference} | Update a sharedrecord:aaaa object



## SharedrecordaaaaGet

> ListSharedrecordAaaaResponse SharedrecordaaaaGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve sharedrecord:aaaa objects



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
	resp, r, err := apiClient.SharedrecordAaaaAPI.SharedrecordaaaaGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.SharedrecordaaaaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordaaaaGet`: ListSharedrecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordAaaaAPI.SharedrecordaaaaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPISharedrecordaaaaGetRequest` struct via the builder pattern


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

[**ListSharedrecordAaaaResponse**](ListSharedrecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordaaaaPost

> CreateSharedrecordAaaaResponse SharedrecordaaaaPost(ctx).SharedrecordAaaa(sharedrecordAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a sharedrecord:aaaa object



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
	sharedrecordAaaa := *dns.NewSharedrecordAaaa() // SharedrecordAaaa | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordAaaaAPI.SharedrecordaaaaPost(context.Background()).SharedrecordAaaa(sharedrecordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.SharedrecordaaaaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordaaaaPost`: CreateSharedrecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordAaaaAPI.SharedrecordaaaaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPISharedrecordaaaaPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordAaaa** | [**SharedrecordAaaa**](SharedrecordAaaa.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSharedrecordAaaaResponse**](CreateSharedrecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordaaaaReferenceDelete

> SharedrecordaaaaReferenceDelete(ctx, reference).Execute()

Delete a sharedrecord:aaaa object



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
	reference := "reference_example" // string | Reference of the sharedrecord:aaaa object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.SharedrecordAaaaAPI.SharedrecordaaaaReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.SharedrecordaaaaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPISharedrecordaaaaReferenceDeleteRequest` struct via the builder pattern


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


## SharedrecordaaaaReferenceGet

> GetSharedrecordAaaaResponse SharedrecordaaaaReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific sharedrecord:aaaa object



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
	reference := "reference_example" // string | Reference of the sharedrecord:aaaa object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordAaaaAPI.SharedrecordaaaaReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.SharedrecordaaaaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordaaaaReferenceGet`: GetSharedrecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordAaaaAPI.SharedrecordaaaaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPISharedrecordaaaaReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSharedrecordAaaaResponse**](GetSharedrecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordaaaaReferencePut

> UpdateSharedrecordAaaaResponse SharedrecordaaaaReferencePut(ctx, reference).SharedrecordAaaa(sharedrecordAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a sharedrecord:aaaa object



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
	reference := "reference_example" // string | Reference of the sharedrecord:aaaa object
	sharedrecordAaaa := *dns.NewSharedrecordAaaa() // SharedrecordAaaa | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordAaaaAPI.SharedrecordaaaaReferencePut(context.Background(), reference).SharedrecordAaaa(sharedrecordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.SharedrecordaaaaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordaaaaReferencePut`: UpdateSharedrecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordAaaaAPI.SharedrecordaaaaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPISharedrecordaaaaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordAaaa** | [**SharedrecordAaaa**](SharedrecordAaaa.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSharedrecordAaaaResponse**](UpdateSharedrecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

