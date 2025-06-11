# RecordUnknownAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordunknownGet**](RecordUnknownAPI.md#RecordunknownGet) | **Get** /record:unknown | Retrieve record:unknown objects
[**RecordunknownPost**](RecordUnknownAPI.md#RecordunknownPost) | **Post** /record:unknown | Create a record:unknown object
[**RecordunknownReferenceDelete**](RecordUnknownAPI.md#RecordunknownReferenceDelete) | **Delete** /record:unknown/{reference} | Delete a record:unknown object
[**RecordunknownReferenceGet**](RecordUnknownAPI.md#RecordunknownReferenceGet) | **Get** /record:unknown/{reference} | Get a specific record:unknown object
[**RecordunknownReferencePut**](RecordUnknownAPI.md#RecordunknownReferencePut) | **Put** /record:unknown/{reference} | Update a record:unknown object



## RecordunknownGet

> ListRecordUnknownResponse RecordunknownGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:unknown objects



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
	resp, r, err := apiClient.RecordUnknownAPI.RecordunknownGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordUnknownAPI.RecordunknownGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordunknownGet`: ListRecordUnknownResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordUnknownAPI.RecordunknownGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordUnknownAPIRecordunknownGetRequest` struct via the builder pattern


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

[**ListRecordUnknownResponse**](ListRecordUnknownResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordunknownPost

> CreateRecordUnknownResponse RecordunknownPost(ctx).RecordUnknown(recordUnknown).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:unknown object



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
	recordUnknown := *dns.NewRecordUnknown() // RecordUnknown | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordUnknownAPI.RecordunknownPost(context.Background()).RecordUnknown(recordUnknown).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordUnknownAPI.RecordunknownPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordunknownPost`: CreateRecordUnknownResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordUnknownAPI.RecordunknownPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordUnknownAPIRecordunknownPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordUnknown** | [**RecordUnknown**](RecordUnknown.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordUnknownResponse**](CreateRecordUnknownResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordunknownReferenceDelete

> RecordunknownReferenceDelete(ctx, reference).Execute()

Delete a record:unknown object



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
	reference := "reference_example" // string | Reference of the record:unknown object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordUnknownAPI.RecordunknownReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordUnknownAPI.RecordunknownReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:unknown object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordUnknownAPIRecordunknownReferenceDeleteRequest` struct via the builder pattern


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


## RecordunknownReferenceGet

> GetRecordUnknownResponse RecordunknownReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:unknown object



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
	reference := "reference_example" // string | Reference of the record:unknown object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordUnknownAPI.RecordunknownReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordUnknownAPI.RecordunknownReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordunknownReferenceGet`: GetRecordUnknownResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordUnknownAPI.RecordunknownReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:unknown object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordUnknownAPIRecordunknownReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordUnknownResponse**](GetRecordUnknownResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordunknownReferencePut

> UpdateRecordUnknownResponse RecordunknownReferencePut(ctx, reference).RecordUnknown(recordUnknown).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:unknown object



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
	reference := "reference_example" // string | Reference of the record:unknown object
	recordUnknown := *dns.NewRecordUnknown() // RecordUnknown | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordUnknownAPI.RecordunknownReferencePut(context.Background(), reference).RecordUnknown(recordUnknown).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordUnknownAPI.RecordunknownReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordunknownReferencePut`: UpdateRecordUnknownResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordUnknownAPI.RecordunknownReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:unknown object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordUnknownAPIRecordunknownReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordUnknown** | [**RecordUnknown**](RecordUnknown.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordUnknownResponse**](UpdateRecordUnknownResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

