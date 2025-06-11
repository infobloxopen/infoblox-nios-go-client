# RecordPtrAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordptrGet**](RecordPtrAPI.md#RecordptrGet) | **Get** /record:ptr | Retrieve record:ptr objects
[**RecordptrPost**](RecordPtrAPI.md#RecordptrPost) | **Post** /record:ptr | Create a record:ptr object
[**RecordptrReferenceDelete**](RecordPtrAPI.md#RecordptrReferenceDelete) | **Delete** /record:ptr/{reference} | Delete a record:ptr object
[**RecordptrReferenceGet**](RecordPtrAPI.md#RecordptrReferenceGet) | **Get** /record:ptr/{reference} | Get a specific record:ptr object
[**RecordptrReferencePut**](RecordPtrAPI.md#RecordptrReferencePut) | **Put** /record:ptr/{reference} | Update a record:ptr object



## RecordptrGet

> ListRecordPtrResponse RecordptrGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:ptr objects



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
	resp, r, err := apiClient.RecordPtrAPI.RecordptrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordPtrAPI.RecordptrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordptrGet`: ListRecordPtrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordPtrAPI.RecordptrGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordPtrAPIRecordptrGetRequest` struct via the builder pattern


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

[**ListRecordPtrResponse**](ListRecordPtrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordptrPost

> CreateRecordPtrResponse RecordptrPost(ctx).RecordPtr(recordPtr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:ptr object



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
	recordPtr := *dns.NewRecordPtr() // RecordPtr | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordPtrAPI.RecordptrPost(context.Background()).RecordPtr(recordPtr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordPtrAPI.RecordptrPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordptrPost`: CreateRecordPtrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordPtrAPI.RecordptrPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordPtrAPIRecordptrPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordPtr** | [**RecordPtr**](RecordPtr.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordPtrResponse**](CreateRecordPtrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordptrReferenceDelete

> RecordptrReferenceDelete(ctx, reference).Execute()

Delete a record:ptr object



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
	reference := "reference_example" // string | Reference of the record:ptr object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordPtrAPI.RecordptrReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordPtrAPI.RecordptrReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:ptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordPtrAPIRecordptrReferenceDeleteRequest` struct via the builder pattern


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


## RecordptrReferenceGet

> GetRecordPtrResponse RecordptrReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:ptr object



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
	reference := "reference_example" // string | Reference of the record:ptr object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordPtrAPI.RecordptrReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordPtrAPI.RecordptrReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordptrReferenceGet`: GetRecordPtrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordPtrAPI.RecordptrReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:ptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordPtrAPIRecordptrReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordPtrResponse**](GetRecordPtrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordptrReferencePut

> UpdateRecordPtrResponse RecordptrReferencePut(ctx, reference).RecordPtr(recordPtr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:ptr object



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
	reference := "reference_example" // string | Reference of the record:ptr object
	recordPtr := *dns.NewRecordPtr() // RecordPtr | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordPtrAPI.RecordptrReferencePut(context.Background(), reference).RecordPtr(recordPtr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordPtrAPI.RecordptrReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordptrReferencePut`: UpdateRecordPtrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordPtrAPI.RecordptrReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:ptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordPtrAPIRecordptrReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordPtr** | [**RecordPtr**](RecordPtr.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordPtrResponse**](UpdateRecordPtrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

