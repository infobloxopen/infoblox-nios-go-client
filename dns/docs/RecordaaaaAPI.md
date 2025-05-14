# RecordaaaaAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](RecordaaaaAPI.md#Get) | **Get** /record:aaaa | Retrieve record:aaaa objects
[**Post**](RecordaaaaAPI.md#Post) | **Post** /record:aaaa | Create a record:aaaa object
[**Put**](RecordaaaaAPI.md#Put) | **Put** /record:aaaa | Use PUT call as GET operation with _method for a Struct field of a record:aaaa object
[**ReferenceDelete**](RecordaaaaAPI.md#ReferenceDelete) | **Delete** /record:aaaa/{reference} | Delete a record:aaaa object
[**ReferenceGet**](RecordaaaaAPI.md#ReferenceGet) | **Get** /record:aaaa/{reference} | Get a specific record:aaaa object
[**ReferencePut**](RecordaaaaAPI.md#ReferencePut) | **Put** /record:aaaa/{reference} | Update a record:aaaa object



## Get

> ListRecordAaaaResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:aaaa objects



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
	resp, r, err := apiClient.RecordaaaaAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIGetRequest` struct via the builder pattern


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

[**ListRecordAaaaResponse**](ListRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateRecordAaaaResponse Post(ctx).RecordAaaa(recordAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:aaaa object



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
	recordAaaa := *dns.NewRecordAaaa() // RecordAaaa | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaaaaAPI.Post(context.Background()).RecordAaaa(recordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordAaaa** | [**RecordAaaa**](RecordAaaa.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordAaaaResponse**](CreateRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Put

> ListRecordAaaaResponse Put(ctx).RecordAaaa(recordAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).MaxResults(maxResults).Method(method).Execute()

Use PUT call as GET operation with _method for a Struct field of a record:aaaa object



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
	recordAaaa := *dns.NewRecordAaaa() // RecordAaaa | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaaaaAPI.Put(context.Background()).RecordAaaa(recordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.Put``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Put`: ListRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.Put`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIPutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordAaaa** | [**RecordAaaa**](RecordAaaa.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**method** | **string** | Enter the method type for the request | 

### Return type

[**ListRecordAaaaResponse**](ListRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceDelete

> ReferenceDelete(ctx, reference).RemoveAssociatedPtr(removeAssociatedPtr).Execute()

Delete a record:aaaa object



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
	reference := "reference_example" // string | Reference of the record:aaaa object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordaaaaAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIReferenceDeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeAssociatedPtr** | **bool** | Delete option that indicates whether the associated PTR records should be removed while deleting the specified A record. | 

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


## ReferenceGet

> GetRecordAaaaResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:aaaa object



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
	reference := "reference_example" // string | Reference of the record:aaaa object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaaaaAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordAaaaResponse**](GetRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateRecordAaaaResponse ReferencePut(ctx, reference).RecordAaaa(recordAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:aaaa object



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
	reference := "reference_example" // string | Reference of the record:aaaa object
	recordAaaa := *dns.NewRecordAaaa() // RecordAaaa | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordaaaaAPI.ReferencePut(context.Background(), reference).RecordAaaa(recordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordaaaaAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordaaaaAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordaaaaAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordAaaa** | [**RecordAaaa**](RecordAaaa.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordAaaaResponse**](UpdateRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

