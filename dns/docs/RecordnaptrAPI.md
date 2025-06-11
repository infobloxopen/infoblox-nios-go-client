# RecordNaptrAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordnaptrGet**](RecordNaptrAPI.md#RecordnaptrGet) | **Get** /record:naptr | Retrieve record:naptr objects
[**RecordnaptrPost**](RecordNaptrAPI.md#RecordnaptrPost) | **Post** /record:naptr | Create a record:naptr object
[**RecordnaptrReferenceDelete**](RecordNaptrAPI.md#RecordnaptrReferenceDelete) | **Delete** /record:naptr/{reference} | Delete a record:naptr object
[**RecordnaptrReferenceGet**](RecordNaptrAPI.md#RecordnaptrReferenceGet) | **Get** /record:naptr/{reference} | Get a specific record:naptr object
[**RecordnaptrReferencePut**](RecordNaptrAPI.md#RecordnaptrReferencePut) | **Put** /record:naptr/{reference} | Update a record:naptr object



## RecordnaptrGet

> ListRecordNaptrResponse RecordnaptrGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:naptr objects



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
	resp, r, err := apiClient.RecordNaptrAPI.RecordnaptrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordNaptrAPI.RecordnaptrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordnaptrGet`: ListRecordNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordNaptrAPI.RecordnaptrGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordNaptrAPIRecordnaptrGetRequest` struct via the builder pattern


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

[**ListRecordNaptrResponse**](ListRecordNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordnaptrPost

> CreateRecordNaptrResponse RecordnaptrPost(ctx).RecordNaptr(recordNaptr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:naptr object



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
	recordNaptr := *dns.NewRecordNaptr() // RecordNaptr | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordNaptrAPI.RecordnaptrPost(context.Background()).RecordNaptr(recordNaptr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordNaptrAPI.RecordnaptrPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordnaptrPost`: CreateRecordNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordNaptrAPI.RecordnaptrPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordNaptrAPIRecordnaptrPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordNaptr** | [**RecordNaptr**](RecordNaptr.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordNaptrResponse**](CreateRecordNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordnaptrReferenceDelete

> RecordnaptrReferenceDelete(ctx, reference).Execute()

Delete a record:naptr object



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
	reference := "reference_example" // string | Reference of the record:naptr object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordNaptrAPI.RecordnaptrReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordNaptrAPI.RecordnaptrReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:naptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordNaptrAPIRecordnaptrReferenceDeleteRequest` struct via the builder pattern


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


## RecordnaptrReferenceGet

> GetRecordNaptrResponse RecordnaptrReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:naptr object



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
	reference := "reference_example" // string | Reference of the record:naptr object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordNaptrAPI.RecordnaptrReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordNaptrAPI.RecordnaptrReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordnaptrReferenceGet`: GetRecordNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordNaptrAPI.RecordnaptrReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:naptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordNaptrAPIRecordnaptrReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordNaptrResponse**](GetRecordNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordnaptrReferencePut

> UpdateRecordNaptrResponse RecordnaptrReferencePut(ctx, reference).RecordNaptr(recordNaptr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:naptr object



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
	reference := "reference_example" // string | Reference of the record:naptr object
	recordNaptr := *dns.NewRecordNaptr() // RecordNaptr | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordNaptrAPI.RecordnaptrReferencePut(context.Background(), reference).RecordNaptr(recordNaptr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordNaptrAPI.RecordnaptrReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordnaptrReferencePut`: UpdateRecordNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordNaptrAPI.RecordnaptrReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:naptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordNaptrAPIRecordnaptrReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordNaptr** | [**RecordNaptr**](RecordNaptr.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordNaptrResponse**](UpdateRecordNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

