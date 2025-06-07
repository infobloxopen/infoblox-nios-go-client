# OrderedresponsepolicyzonesAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](OrderedresponsepolicyzonesAPI.md#Get) | **Get** /orderedresponsepolicyzones | Retrieve orderedresponsepolicyzones objects
[**ReferenceGet**](OrderedresponsepolicyzonesAPI.md#ReferenceGet) | **Get** /orderedresponsepolicyzones/{reference} | Get a specific orderedresponsepolicyzones object
[**ReferencePut**](OrderedresponsepolicyzonesAPI.md#ReferencePut) | **Put** /orderedresponsepolicyzones/{reference} | Update a orderedresponsepolicyzones object



## Get

> ListOrderedresponsepolicyzonesResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve orderedresponsepolicyzones objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rpz"
)

func main() {

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.OrderedresponsepolicyzonesAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderedresponsepolicyzonesAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListOrderedresponsepolicyzonesResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderedresponsepolicyzonesAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `OrderedresponsepolicyzonesAPIGetRequest` struct via the builder pattern


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

[**ListOrderedresponsepolicyzonesResponse**](ListOrderedresponsepolicyzonesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceGet

> GetOrderedresponsepolicyzonesResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific orderedresponsepolicyzones object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rpz"
)

func main() {
	reference := "reference_example" // string | Reference of the orderedresponsepolicyzones object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.OrderedresponsepolicyzonesAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderedresponsepolicyzonesAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetOrderedresponsepolicyzonesResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderedresponsepolicyzonesAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the orderedresponsepolicyzones object | 

### Other Parameters

Other parameters are passed through a pointer to a `OrderedresponsepolicyzonesAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetOrderedresponsepolicyzonesResponse**](GetOrderedresponsepolicyzonesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateOrderedresponsepolicyzonesResponse ReferencePut(ctx, reference).Orderedresponsepolicyzones(orderedresponsepolicyzones).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a orderedresponsepolicyzones object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rpz"
)

func main() {
	reference := "reference_example" // string | Reference of the orderedresponsepolicyzones object
	orderedresponsepolicyzones := *rpz.NewOrderedresponsepolicyzones() // Orderedresponsepolicyzones | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.OrderedresponsepolicyzonesAPI.ReferencePut(context.Background(), reference).Orderedresponsepolicyzones(orderedresponsepolicyzones).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderedresponsepolicyzonesAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateOrderedresponsepolicyzonesResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderedresponsepolicyzonesAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the orderedresponsepolicyzones object | 

### Other Parameters

Other parameters are passed through a pointer to a `OrderedresponsepolicyzonesAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**orderedresponsepolicyzones** | [**Orderedresponsepolicyzones**](Orderedresponsepolicyzones.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateOrderedresponsepolicyzonesResponse**](UpdateOrderedresponsepolicyzonesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

