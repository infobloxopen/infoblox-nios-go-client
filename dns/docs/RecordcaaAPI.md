# RecordCaaAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordcaaGet**](RecordCaaAPI.md#RecordcaaGet) | **Get** /record:caa | Retrieve record:caa objects
[**RecordcaaPost**](RecordCaaAPI.md#RecordcaaPost) | **Post** /record:caa | Create a record:caa object
[**RecordcaaReferenceDelete**](RecordCaaAPI.md#RecordcaaReferenceDelete) | **Delete** /record:caa/{reference} | Delete a record:caa object
[**RecordcaaReferenceGet**](RecordCaaAPI.md#RecordcaaReferenceGet) | **Get** /record:caa/{reference} | Get a specific record:caa object
[**RecordcaaReferencePut**](RecordCaaAPI.md#RecordcaaReferencePut) | **Put** /record:caa/{reference} | Update a record:caa object



## RecordcaaGet

> ListRecordCaaResponse RecordcaaGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:caa objects



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
	resp, r, err := apiClient.RecordCaaAPI.RecordcaaGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCaaAPI.RecordcaaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordcaaGet`: ListRecordCaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordCaaAPI.RecordcaaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordCaaAPIRecordcaaGetRequest` struct via the builder pattern


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

[**ListRecordCaaResponse**](ListRecordCaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordcaaPost

> CreateRecordCaaResponse RecordcaaPost(ctx).RecordCaa(recordCaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:caa object



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
	recordCaa := *dns.NewRecordCaa() // RecordCaa | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordCaaAPI.RecordcaaPost(context.Background()).RecordCaa(recordCaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCaaAPI.RecordcaaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordcaaPost`: CreateRecordCaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordCaaAPI.RecordcaaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordCaaAPIRecordcaaPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordCaa** | [**RecordCaa**](RecordCaa.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordCaaResponse**](CreateRecordCaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordcaaReferenceDelete

> RecordcaaReferenceDelete(ctx, reference).Execute()

Delete a record:caa object



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
	reference := "reference_example" // string | Reference of the record:caa object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordCaaAPI.RecordcaaReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCaaAPI.RecordcaaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:caa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordCaaAPIRecordcaaReferenceDeleteRequest` struct via the builder pattern


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


## RecordcaaReferenceGet

> GetRecordCaaResponse RecordcaaReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:caa object



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
	reference := "reference_example" // string | Reference of the record:caa object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordCaaAPI.RecordcaaReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCaaAPI.RecordcaaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordcaaReferenceGet`: GetRecordCaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordCaaAPI.RecordcaaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:caa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordCaaAPIRecordcaaReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordCaaResponse**](GetRecordCaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordcaaReferencePut

> UpdateRecordCaaResponse RecordcaaReferencePut(ctx, reference).RecordCaa(recordCaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:caa object



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
	reference := "reference_example" // string | Reference of the record:caa object
	recordCaa := *dns.NewRecordCaa() // RecordCaa | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordCaaAPI.RecordcaaReferencePut(context.Background(), reference).RecordCaa(recordCaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCaaAPI.RecordcaaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordcaaReferencePut`: UpdateRecordCaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordCaaAPI.RecordcaaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:caa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordCaaAPIRecordcaaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordCaa** | [**RecordCaa**](RecordCaa.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordCaaResponse**](UpdateRecordCaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

