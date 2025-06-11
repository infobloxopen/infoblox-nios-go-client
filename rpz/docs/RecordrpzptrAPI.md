# RecordRpzPtrAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpzptrGet**](RecordRpzPtrAPI.md#RecordrpzptrGet) | **Get** /record:rpz:ptr | Retrieve record:rpz:ptr objects
[**RecordrpzptrPost**](RecordRpzPtrAPI.md#RecordrpzptrPost) | **Post** /record:rpz:ptr | Create a record:rpz:ptr object
[**RecordrpzptrReferenceDelete**](RecordRpzPtrAPI.md#RecordrpzptrReferenceDelete) | **Delete** /record:rpz:ptr/{reference} | Delete a record:rpz:ptr object
[**RecordrpzptrReferenceGet**](RecordRpzPtrAPI.md#RecordrpzptrReferenceGet) | **Get** /record:rpz:ptr/{reference} | Get a specific record:rpz:ptr object
[**RecordrpzptrReferencePut**](RecordRpzPtrAPI.md#RecordrpzptrReferencePut) | **Put** /record:rpz:ptr/{reference} | Update a record:rpz:ptr object



## RecordrpzptrGet

> ListRecordRpzPtrResponse RecordrpzptrGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:ptr objects



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
	resp, r, err := apiClient.RecordRpzPtrAPI.RecordrpzptrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzPtrAPI.RecordrpzptrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzptrGet`: ListRecordRpzPtrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzPtrAPI.RecordrpzptrGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzPtrAPIRecordrpzptrGetRequest` struct via the builder pattern


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

[**ListRecordRpzPtrResponse**](ListRecordRpzPtrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzptrPost

> CreateRecordRpzPtrResponse RecordrpzptrPost(ctx).RecordRpzPtr(recordRpzPtr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:ptr object



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
	recordRpzPtr := *rpz.NewRecordRpzPtr() // RecordRpzPtr | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzPtrAPI.RecordrpzptrPost(context.Background()).RecordRpzPtr(recordRpzPtr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzPtrAPI.RecordrpzptrPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzptrPost`: CreateRecordRpzPtrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzPtrAPI.RecordrpzptrPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzPtrAPIRecordrpzptrPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzPtr** | [**RecordRpzPtr**](RecordRpzPtr.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzPtrResponse**](CreateRecordRpzPtrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzptrReferenceDelete

> RecordrpzptrReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:ptr object



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
	reference := "reference_example" // string | Reference of the record:rpz:ptr object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzPtrAPI.RecordrpzptrReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzPtrAPI.RecordrpzptrReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:ptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzPtrAPIRecordrpzptrReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpzptrReferenceGet

> GetRecordRpzPtrResponse RecordrpzptrReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:ptr object



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
	reference := "reference_example" // string | Reference of the record:rpz:ptr object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzPtrAPI.RecordrpzptrReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzPtrAPI.RecordrpzptrReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzptrReferenceGet`: GetRecordRpzPtrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzPtrAPI.RecordrpzptrReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:ptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzPtrAPIRecordrpzptrReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzPtrResponse**](GetRecordRpzPtrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzptrReferencePut

> UpdateRecordRpzPtrResponse RecordrpzptrReferencePut(ctx, reference).RecordRpzPtr(recordRpzPtr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:ptr object



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
	reference := "reference_example" // string | Reference of the record:rpz:ptr object
	recordRpzPtr := *rpz.NewRecordRpzPtr() // RecordRpzPtr | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzPtrAPI.RecordrpzptrReferencePut(context.Background(), reference).RecordRpzPtr(recordRpzPtr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzPtrAPI.RecordrpzptrReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzptrReferencePut`: UpdateRecordRpzPtrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzPtrAPI.RecordrpzptrReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:ptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzPtrAPIRecordrpzptrReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzPtr** | [**RecordRpzPtr**](RecordRpzPtr.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzPtrResponse**](UpdateRecordRpzPtrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

