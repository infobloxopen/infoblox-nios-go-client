# RecordRpzCnameClientipaddressAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpzcnameclientipaddressGet**](RecordRpzCnameClientipaddressAPI.md#RecordrpzcnameclientipaddressGet) | **Get** /record:rpz:cname:clientipaddress | Retrieve record:rpz:cname:clientipaddress objects
[**RecordrpzcnameclientipaddressPost**](RecordRpzCnameClientipaddressAPI.md#RecordrpzcnameclientipaddressPost) | **Post** /record:rpz:cname:clientipaddress | Create a record:rpz:cname:clientipaddress object
[**RecordrpzcnameclientipaddressReferenceDelete**](RecordRpzCnameClientipaddressAPI.md#RecordrpzcnameclientipaddressReferenceDelete) | **Delete** /record:rpz:cname:clientipaddress/{reference} | Delete a record:rpz:cname:clientipaddress object
[**RecordrpzcnameclientipaddressReferenceGet**](RecordRpzCnameClientipaddressAPI.md#RecordrpzcnameclientipaddressReferenceGet) | **Get** /record:rpz:cname:clientipaddress/{reference} | Get a specific record:rpz:cname:clientipaddress object
[**RecordrpzcnameclientipaddressReferencePut**](RecordRpzCnameClientipaddressAPI.md#RecordrpzcnameclientipaddressReferencePut) | **Put** /record:rpz:cname:clientipaddress/{reference} | Update a record:rpz:cname:clientipaddress object



## RecordrpzcnameclientipaddressGet

> ListRecordRpzCnameClientipaddressResponse RecordrpzcnameclientipaddressGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:cname:clientipaddress objects



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
	resp, r, err := apiClient.RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameclientipaddressGet`: ListRecordRpzCnameClientipaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressAPIRecordrpzcnameclientipaddressGetRequest` struct via the builder pattern


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

[**ListRecordRpzCnameClientipaddressResponse**](ListRecordRpzCnameClientipaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameclientipaddressPost

> CreateRecordRpzCnameClientipaddressResponse RecordrpzcnameclientipaddressPost(ctx).RecordRpzCnameClientipaddress(recordRpzCnameClientipaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:cname:clientipaddress object



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
	recordRpzCnameClientipaddress := *rpz.NewRecordRpzCnameClientipaddress() // RecordRpzCnameClientipaddress | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressPost(context.Background()).RecordRpzCnameClientipaddress(recordRpzCnameClientipaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameclientipaddressPost`: CreateRecordRpzCnameClientipaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressAPIRecordrpzcnameclientipaddressPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameClientipaddress** | [**RecordRpzCnameClientipaddress**](RecordRpzCnameClientipaddress.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzCnameClientipaddressResponse**](CreateRecordRpzCnameClientipaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameclientipaddressReferenceDelete

> RecordrpzcnameclientipaddressReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:cname:clientipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:clientipaddress object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressAPIRecordrpzcnameclientipaddressReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpzcnameclientipaddressReferenceGet

> GetRecordRpzCnameClientipaddressResponse RecordrpzcnameclientipaddressReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:cname:clientipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:clientipaddress object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameclientipaddressReferenceGet`: GetRecordRpzCnameClientipaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressAPIRecordrpzcnameclientipaddressReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzCnameClientipaddressResponse**](GetRecordRpzCnameClientipaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameclientipaddressReferencePut

> UpdateRecordRpzCnameClientipaddressResponse RecordrpzcnameclientipaddressReferencePut(ctx, reference).RecordRpzCnameClientipaddress(recordRpzCnameClientipaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:cname:clientipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:clientipaddress object
	recordRpzCnameClientipaddress := *rpz.NewRecordRpzCnameClientipaddress() // RecordRpzCnameClientipaddress | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressReferencePut(context.Background(), reference).RecordRpzCnameClientipaddress(recordRpzCnameClientipaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameclientipaddressReferencePut`: UpdateRecordRpzCnameClientipaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressAPI.RecordrpzcnameclientipaddressReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressAPIRecordrpzcnameclientipaddressReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameClientipaddress** | [**RecordRpzCnameClientipaddress**](RecordRpzCnameClientipaddress.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzCnameClientipaddressResponse**](UpdateRecordRpzCnameClientipaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

