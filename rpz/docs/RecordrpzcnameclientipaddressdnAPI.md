# RecordRpzCnameClientipaddressdnAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpzcnameclientipaddressdnGet**](RecordRpzCnameClientipaddressdnAPI.md#RecordrpzcnameclientipaddressdnGet) | **Get** /record:rpz:cname:clientipaddressdn | Retrieve record:rpz:cname:clientipaddressdn objects
[**RecordrpzcnameclientipaddressdnPost**](RecordRpzCnameClientipaddressdnAPI.md#RecordrpzcnameclientipaddressdnPost) | **Post** /record:rpz:cname:clientipaddressdn | Create a record:rpz:cname:clientipaddressdn object
[**RecordrpzcnameclientipaddressdnReferenceDelete**](RecordRpzCnameClientipaddressdnAPI.md#RecordrpzcnameclientipaddressdnReferenceDelete) | **Delete** /record:rpz:cname:clientipaddressdn/{reference} | Delete a record:rpz:cname:clientipaddressdn object
[**RecordrpzcnameclientipaddressdnReferenceGet**](RecordRpzCnameClientipaddressdnAPI.md#RecordrpzcnameclientipaddressdnReferenceGet) | **Get** /record:rpz:cname:clientipaddressdn/{reference} | Get a specific record:rpz:cname:clientipaddressdn object
[**RecordrpzcnameclientipaddressdnReferencePut**](RecordRpzCnameClientipaddressdnAPI.md#RecordrpzcnameclientipaddressdnReferencePut) | **Put** /record:rpz:cname:clientipaddressdn/{reference} | Update a record:rpz:cname:clientipaddressdn object



## RecordrpzcnameclientipaddressdnGet

> ListRecordRpzCnameClientipaddressdnResponse RecordrpzcnameclientipaddressdnGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:cname:clientipaddressdn objects



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
	resp, r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameclientipaddressdnGet`: ListRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPIRecordrpzcnameclientipaddressdnGetRequest` struct via the builder pattern


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

[**ListRecordRpzCnameClientipaddressdnResponse**](ListRecordRpzCnameClientipaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameclientipaddressdnPost

> CreateRecordRpzCnameClientipaddressdnResponse RecordrpzcnameclientipaddressdnPost(ctx).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:cname:clientipaddressdn object



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
	recordRpzCnameClientipaddressdn := *rpz.NewRecordRpzCnameClientipaddressdn() // RecordRpzCnameClientipaddressdn | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnPost(context.Background()).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameclientipaddressdnPost`: CreateRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPIRecordrpzcnameclientipaddressdnPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameClientipaddressdn** | [**RecordRpzCnameClientipaddressdn**](RecordRpzCnameClientipaddressdn.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzCnameClientipaddressdnResponse**](CreateRecordRpzCnameClientipaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameclientipaddressdnReferenceDelete

> RecordrpzcnameclientipaddressdnReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:cname:clientipaddressdn object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:clientipaddressdn object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPIRecordrpzcnameclientipaddressdnReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpzcnameclientipaddressdnReferenceGet

> GetRecordRpzCnameClientipaddressdnResponse RecordrpzcnameclientipaddressdnReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:cname:clientipaddressdn object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:clientipaddressdn object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameclientipaddressdnReferenceGet`: GetRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPIRecordrpzcnameclientipaddressdnReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzCnameClientipaddressdnResponse**](GetRecordRpzCnameClientipaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameclientipaddressdnReferencePut

> UpdateRecordRpzCnameClientipaddressdnResponse RecordrpzcnameclientipaddressdnReferencePut(ctx, reference).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:cname:clientipaddressdn object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:clientipaddressdn object
	recordRpzCnameClientipaddressdn := *rpz.NewRecordRpzCnameClientipaddressdn() // RecordRpzCnameClientipaddressdn | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnReferencePut(context.Background(), reference).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameclientipaddressdnReferencePut`: UpdateRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressdnAPI.RecordrpzcnameclientipaddressdnReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPIRecordrpzcnameclientipaddressdnReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameClientipaddressdn** | [**RecordRpzCnameClientipaddressdn**](RecordRpzCnameClientipaddressdn.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzCnameClientipaddressdnResponse**](UpdateRecordRpzCnameClientipaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

