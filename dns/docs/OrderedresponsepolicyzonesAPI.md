# OrderedresponsepolicyzonesAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](OrderedresponsepolicyzonesAPI.md#List) | **Get** /orderedresponsepolicyzones | Retrieve orderedresponsepolicyzones objects
[**Read**](OrderedresponsepolicyzonesAPI.md#Read) | **Get** /orderedresponsepolicyzones/{reference} | Get a specific orderedresponsepolicyzones object
[**Update**](OrderedresponsepolicyzonesAPI.md#Update) | **Put** /orderedresponsepolicyzones/{reference} | Update a orderedresponsepolicyzones object



## List

> ListOrderedresponsepolicyzonesResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve orderedresponsepolicyzones objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.OrderedresponsepolicyzonesAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderedresponsepolicyzonesAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListOrderedresponsepolicyzonesResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderedresponsepolicyzonesAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `OrderedresponsepolicyzonesAPIListRequest` struct via the builder pattern


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

[**ListOrderedresponsepolicyzonesResponse**](ListOrderedresponsepolicyzonesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetOrderedresponsepolicyzonesResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific orderedresponsepolicyzones object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the orderedresponsepolicyzones object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.OrderedresponsepolicyzonesAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderedresponsepolicyzonesAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetOrderedresponsepolicyzonesResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderedresponsepolicyzonesAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the orderedresponsepolicyzones object | 

### Other Parameters

Other parameters are passed through a pointer to a `OrderedresponsepolicyzonesAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

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


## Update

> UpdateOrderedresponsepolicyzonesResponse Update(ctx, reference).Orderedresponsepolicyzones(orderedresponsepolicyzones).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a orderedresponsepolicyzones object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the orderedresponsepolicyzones object
	orderedresponsepolicyzones := *dns.NewOrderedresponsepolicyzones() // Orderedresponsepolicyzones | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.OrderedresponsepolicyzonesAPI.Update(context.Background(), reference).Orderedresponsepolicyzones(orderedresponsepolicyzones).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderedresponsepolicyzonesAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateOrderedresponsepolicyzonesResponse
	fmt.Fprintf(os.Stdout, "Response from `OrderedresponsepolicyzonesAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the orderedresponsepolicyzones object | 

### Other Parameters

Other parameters are passed through a pointer to a `OrderedresponsepolicyzonesAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**orderedresponsepolicyzones** | [**Orderedresponsepolicyzones**](Orderedresponsepolicyzones.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

