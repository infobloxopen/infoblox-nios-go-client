# RecordRpzNaptrAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpznaptrGet**](RecordRpzNaptrAPI.md#RecordrpznaptrGet) | **Get** /record:rpz:naptr | Retrieve record:rpz:naptr objects
[**RecordrpznaptrPost**](RecordRpzNaptrAPI.md#RecordrpznaptrPost) | **Post** /record:rpz:naptr | Create a record:rpz:naptr object
[**RecordrpznaptrReferenceDelete**](RecordRpzNaptrAPI.md#RecordrpznaptrReferenceDelete) | **Delete** /record:rpz:naptr/{reference} | Delete a record:rpz:naptr object
[**RecordrpznaptrReferenceGet**](RecordRpzNaptrAPI.md#RecordrpznaptrReferenceGet) | **Get** /record:rpz:naptr/{reference} | Get a specific record:rpz:naptr object
[**RecordrpznaptrReferencePut**](RecordRpzNaptrAPI.md#RecordrpznaptrReferencePut) | **Put** /record:rpz:naptr/{reference} | Update a record:rpz:naptr object



## RecordrpznaptrGet

> ListRecordRpzNaptrResponse RecordrpznaptrGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:naptr objects



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
	resp, r, err := apiClient.RecordRpzNaptrAPI.RecordrpznaptrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzNaptrAPI.RecordrpznaptrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpznaptrGet`: ListRecordRpzNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzNaptrAPI.RecordrpznaptrGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzNaptrAPIRecordrpznaptrGetRequest` struct via the builder pattern


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

[**ListRecordRpzNaptrResponse**](ListRecordRpzNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpznaptrPost

> CreateRecordRpzNaptrResponse RecordrpznaptrPost(ctx).RecordRpzNaptr(recordRpzNaptr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:naptr object



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
	recordRpzNaptr := *rpz.NewRecordRpzNaptr() // RecordRpzNaptr | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzNaptrAPI.RecordrpznaptrPost(context.Background()).RecordRpzNaptr(recordRpzNaptr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzNaptrAPI.RecordrpznaptrPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpznaptrPost`: CreateRecordRpzNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzNaptrAPI.RecordrpznaptrPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzNaptrAPIRecordrpznaptrPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzNaptr** | [**RecordRpzNaptr**](RecordRpzNaptr.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzNaptrResponse**](CreateRecordRpzNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpznaptrReferenceDelete

> RecordrpznaptrReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:naptr object



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
	reference := "reference_example" // string | Reference of the record:rpz:naptr object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzNaptrAPI.RecordrpznaptrReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzNaptrAPI.RecordrpznaptrReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:naptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzNaptrAPIRecordrpznaptrReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpznaptrReferenceGet

> GetRecordRpzNaptrResponse RecordrpznaptrReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:naptr object



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
	reference := "reference_example" // string | Reference of the record:rpz:naptr object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzNaptrAPI.RecordrpznaptrReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzNaptrAPI.RecordrpznaptrReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpznaptrReferenceGet`: GetRecordRpzNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzNaptrAPI.RecordrpznaptrReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:naptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzNaptrAPIRecordrpznaptrReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzNaptrResponse**](GetRecordRpzNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpznaptrReferencePut

> UpdateRecordRpzNaptrResponse RecordrpznaptrReferencePut(ctx, reference).RecordRpzNaptr(recordRpzNaptr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:naptr object



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
	reference := "reference_example" // string | Reference of the record:rpz:naptr object
	recordRpzNaptr := *rpz.NewRecordRpzNaptr() // RecordRpzNaptr | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzNaptrAPI.RecordrpznaptrReferencePut(context.Background(), reference).RecordRpzNaptr(recordRpzNaptr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzNaptrAPI.RecordrpznaptrReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpznaptrReferencePut`: UpdateRecordRpzNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzNaptrAPI.RecordrpznaptrReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:naptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzNaptrAPIRecordrpznaptrReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzNaptr** | [**RecordRpzNaptr**](RecordRpzNaptr.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzNaptrResponse**](UpdateRecordRpzNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

