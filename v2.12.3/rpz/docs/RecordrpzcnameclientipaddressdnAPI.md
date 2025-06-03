# RecordrpzcnameclientipaddressdnAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](RecordrpzcnameclientipaddressdnAPI.md#Get) | **Get** /record:rpz:cname:clientipaddressdn | Retrieve record:rpz:cname:clientipaddressdn objects
[**Post**](RecordrpzcnameclientipaddressdnAPI.md#Post) | **Post** /record:rpz:cname:clientipaddressdn | Create a record:rpz:cname:clientipaddressdn object
[**ReferenceDelete**](RecordrpzcnameclientipaddressdnAPI.md#ReferenceDelete) | **Delete** /record:rpz:cname:clientipaddressdn/{reference} | Delete a record:rpz:cname:clientipaddressdn object
[**ReferenceGet**](RecordrpzcnameclientipaddressdnAPI.md#ReferenceGet) | **Get** /record:rpz:cname:clientipaddressdn/{reference} | Get a specific record:rpz:cname:clientipaddressdn object
[**ReferencePut**](RecordrpzcnameclientipaddressdnAPI.md#ReferencePut) | **Put** /record:rpz:cname:clientipaddressdn/{reference} | Update a record:rpz:cname:clientipaddressdn object



## Get

> ListRecordRpzCnameClientipaddressdnResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

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
	resp, r, err := apiClient.RecordrpzcnameclientipaddressdnAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameclientipaddressdnAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordrpzcnameclientipaddressdnAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordrpzcnameclientipaddressdnAPIGetRequest` struct via the builder pattern


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


## Post

> CreateRecordRpzCnameClientipaddressdnResponse Post(ctx).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.RecordrpzcnameclientipaddressdnAPI.Post(context.Background()).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameclientipaddressdnAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordrpzcnameclientipaddressdnAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordrpzcnameclientipaddressdnAPIPostRequest` struct via the builder pattern


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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

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
	r, err := apiClient.RecordrpzcnameclientipaddressdnAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameclientipaddressdnAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `RecordrpzcnameclientipaddressdnAPIReferenceDeleteRequest` struct via the builder pattern


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


## ReferenceGet

> GetRecordRpzCnameClientipaddressdnResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.RecordrpzcnameclientipaddressdnAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameclientipaddressdnAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordrpzcnameclientipaddressdnAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordrpzcnameclientipaddressdnAPIReferenceGetRequest` struct via the builder pattern


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


## ReferencePut

> UpdateRecordRpzCnameClientipaddressdnResponse ReferencePut(ctx, reference).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.RecordrpzcnameclientipaddressdnAPI.ReferencePut(context.Background(), reference).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameclientipaddressdnAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordrpzcnameclientipaddressdnAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordrpzcnameclientipaddressdnAPIReferencePutRequest` struct via the builder pattern


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

