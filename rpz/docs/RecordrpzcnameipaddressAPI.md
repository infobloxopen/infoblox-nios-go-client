# RecordRpzCnameIpaddressAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpzcnameipaddressGet**](RecordRpzCnameIpaddressAPI.md#RecordrpzcnameipaddressGet) | **Get** /record:rpz:cname:ipaddress | Retrieve record:rpz:cname:ipaddress objects
[**RecordrpzcnameipaddressPost**](RecordRpzCnameIpaddressAPI.md#RecordrpzcnameipaddressPost) | **Post** /record:rpz:cname:ipaddress | Create a record:rpz:cname:ipaddress object
[**RecordrpzcnameipaddressReferenceDelete**](RecordRpzCnameIpaddressAPI.md#RecordrpzcnameipaddressReferenceDelete) | **Delete** /record:rpz:cname:ipaddress/{reference} | Delete a record:rpz:cname:ipaddress object
[**RecordrpzcnameipaddressReferenceGet**](RecordRpzCnameIpaddressAPI.md#RecordrpzcnameipaddressReferenceGet) | **Get** /record:rpz:cname:ipaddress/{reference} | Get a specific record:rpz:cname:ipaddress object
[**RecordrpzcnameipaddressReferencePut**](RecordRpzCnameIpaddressAPI.md#RecordrpzcnameipaddressReferencePut) | **Put** /record:rpz:cname:ipaddress/{reference} | Update a record:rpz:cname:ipaddress object



## RecordrpzcnameipaddressGet

> ListRecordRpzCnameIpaddressResponse RecordrpzcnameipaddressGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:cname:ipaddress objects



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
	resp, r, err := apiClient.RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameipaddressGet`: ListRecordRpzCnameIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressAPIRecordrpzcnameipaddressGetRequest` struct via the builder pattern


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

[**ListRecordRpzCnameIpaddressResponse**](ListRecordRpzCnameIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameipaddressPost

> CreateRecordRpzCnameIpaddressResponse RecordrpzcnameipaddressPost(ctx).RecordRpzCnameIpaddress(recordRpzCnameIpaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:cname:ipaddress object



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
	recordRpzCnameIpaddress := *rpz.NewRecordRpzCnameIpaddress() // RecordRpzCnameIpaddress | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressPost(context.Background()).RecordRpzCnameIpaddress(recordRpzCnameIpaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameipaddressPost`: CreateRecordRpzCnameIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressAPIRecordrpzcnameipaddressPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameIpaddress** | [**RecordRpzCnameIpaddress**](RecordRpzCnameIpaddress.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzCnameIpaddressResponse**](CreateRecordRpzCnameIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameipaddressReferenceDelete

> RecordrpzcnameipaddressReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:cname:ipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:ipaddress object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressAPIRecordrpzcnameipaddressReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpzcnameipaddressReferenceGet

> GetRecordRpzCnameIpaddressResponse RecordrpzcnameipaddressReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:cname:ipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:ipaddress object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameipaddressReferenceGet`: GetRecordRpzCnameIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressAPIRecordrpzcnameipaddressReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzCnameIpaddressResponse**](GetRecordRpzCnameIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzcnameipaddressReferencePut

> UpdateRecordRpzCnameIpaddressResponse RecordrpzcnameipaddressReferencePut(ctx, reference).RecordRpzCnameIpaddress(recordRpzCnameIpaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:cname:ipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:cname:ipaddress object
	recordRpzCnameIpaddress := *rpz.NewRecordRpzCnameIpaddress() // RecordRpzCnameIpaddress | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressReferencePut(context.Background(), reference).RecordRpzCnameIpaddress(recordRpzCnameIpaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzcnameipaddressReferencePut`: UpdateRecordRpzCnameIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameIpaddressAPI.RecordrpzcnameipaddressReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameIpaddressAPIRecordrpzcnameipaddressReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameIpaddress** | [**RecordRpzCnameIpaddress**](RecordRpzCnameIpaddress.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzCnameIpaddressResponse**](UpdateRecordRpzCnameIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

