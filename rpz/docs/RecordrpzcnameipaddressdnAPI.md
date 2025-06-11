# RecordRpzCnameIpaddressdnAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpzcnameipaddressdnGet**](RecordRpzCnameIpaddressdnAPI.md#RecordrpzcnameipaddressdnGet) | **Get** /record:rpz:cname:ipaddressdn | Retrieve record:rpz:cname:ipaddressdn objects
[**RecordrpzcnameipaddressdnPost**](RecordRpzCnameIpaddressdnAPI.md#RecordrpzcnameipaddressdnPost) | **Post** /record:rpz:cname:ipaddressdn | Create a record:rpz:cname:ipaddressdn object
[**RecordrpzcnameipaddressdnReferenceDelete**](RecordRpzCnameIpaddressdnAPI.md#RecordrpzcnameipaddressdnReferenceDelete) | **Delete** /record:rpz:cname:ipaddressdn/{reference} | Delete a record:rpz:cname:ipaddressdn object
[**RecordrpzcnameipaddressdnReferenceGet**](RecordRpzCnameIpaddressdnAPI.md#RecordrpzcnameipaddressdnReferenceGet) | **Get** /record:rpz:cname:ipaddressdn/{reference} | Get a specific record:rpz:cname:ipaddressdn object
[**RecordrpzcnameipaddressdnReferencePut**](RecordRpzCnameIpaddressdnAPI.md#RecordrpzcnameipaddressdnReferencePut) | **Put** /record:rpz:cname:ipaddressdn/{reference} | Update a record:rpz:cname:ipaddressdn object



## RecordrpzcnameipaddressdnGet

> ListRecordRpzCnameIpaddressdnResponse RecordrpzcnameipaddressdnGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:cname:ipaddressdn objects



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
	resp, r, err := apiClient.RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameipaddressdnGet`: ListRecordRpzCnameIpaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressdnAPIRecordrpzcnameipaddressdnGetRequest` struct via the builder pattern


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

[**ListRecordRpzCnameIpaddressdnResponse**](ListRecordRpzCnameIpaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameipaddressdnPost

> CreateRecordRpzCnameIpaddressdnResponse RecordrpzcnameipaddressdnPost(ctx).RecordRpzCnameIpaddressdn(recordRpzCnameIpaddressdn).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:cname:ipaddressdn object



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
	recordRpzCnameIpaddressdn := *rpz.NewRecordRpzCnameIpaddressdn() // RecordRpzCnameIpaddressdn | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnPost(context.Background()).RecordRpzCnameIpaddressdn(recordRpzCnameIpaddressdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameipaddressdnPost`: CreateRecordRpzCnameIpaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressdnAPIRecordrpzcnameipaddressdnPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameIpaddressdn** | [**RecordRpzCnameIpaddressdn**](RecordRpzCnameIpaddressdn.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzCnameIpaddressdnResponse**](CreateRecordRpzCnameIpaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameipaddressdnReferenceDelete

> RecordrpzcnameipaddressdnReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:cname:ipaddressdn object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:ipaddressdn object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:ipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressdnAPIRecordrpzcnameipaddressdnReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpzcnameipaddressdnReferenceGet

> GetRecordRpzCnameIpaddressdnResponse RecordrpzcnameipaddressdnReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:cname:ipaddressdn object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:ipaddressdn object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameipaddressdnReferenceGet`: GetRecordRpzCnameIpaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:ipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressdnAPIRecordrpzcnameipaddressdnReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzCnameIpaddressdnResponse**](GetRecordRpzCnameIpaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameipaddressdnReferencePut

> UpdateRecordRpzCnameIpaddressdnResponse RecordrpzcnameipaddressdnReferencePut(ctx, reference).RecordRpzCnameIpaddressdn(recordRpzCnameIpaddressdn).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:cname:ipaddressdn object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:ipaddressdn object
	recordRpzCnameIpaddressdn := *rpz.NewRecordRpzCnameIpaddressdn() // RecordRpzCnameIpaddressdn | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnReferencePut(context.Background(), reference).RecordRpzCnameIpaddressdn(recordRpzCnameIpaddressdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameipaddressdnReferencePut`: UpdateRecordRpzCnameIpaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameIpaddressdnAPI.RecordrpzcnameipaddressdnReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:ipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressdnAPIRecordrpzcnameipaddressdnReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameIpaddressdn** | [**RecordRpzCnameIpaddressdn**](RecordRpzCnameIpaddressdn.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzCnameIpaddressdnResponse**](UpdateRecordRpzCnameIpaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

