# RecordRpzAaaaIpaddressAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpzaaaaipaddressGet**](RecordRpzAaaaIpaddressAPI.md#RecordrpzaaaaipaddressGet) | **Get** /record:rpz:aaaa:ipaddress | Retrieve record:rpz:aaaa:ipaddress objects
[**RecordrpzaaaaipaddressPost**](RecordRpzAaaaIpaddressAPI.md#RecordrpzaaaaipaddressPost) | **Post** /record:rpz:aaaa:ipaddress | Create a record:rpz:aaaa:ipaddress object
[**RecordrpzaaaaipaddressReferenceDelete**](RecordRpzAaaaIpaddressAPI.md#RecordrpzaaaaipaddressReferenceDelete) | **Delete** /record:rpz:aaaa:ipaddress/{reference} | Delete a record:rpz:aaaa:ipaddress object
[**RecordrpzaaaaipaddressReferenceGet**](RecordRpzAaaaIpaddressAPI.md#RecordrpzaaaaipaddressReferenceGet) | **Get** /record:rpz:aaaa:ipaddress/{reference} | Get a specific record:rpz:aaaa:ipaddress object
[**RecordrpzaaaaipaddressReferencePut**](RecordRpzAaaaIpaddressAPI.md#RecordrpzaaaaipaddressReferencePut) | **Put** /record:rpz:aaaa:ipaddress/{reference} | Update a record:rpz:aaaa:ipaddress object



## RecordrpzaaaaipaddressGet

> ListRecordRpzAaaaIpaddressResponse RecordrpzaaaaipaddressGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:aaaa:ipaddress objects



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
	resp, r, err := apiClient.RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaaaaipaddressGet`: ListRecordRpzAaaaIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaIpaddressAPIRecordrpzaaaaipaddressGetRequest` struct via the builder pattern


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

[**ListRecordRpzAaaaIpaddressResponse**](ListRecordRpzAaaaIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzaaaaipaddressPost

> CreateRecordRpzAaaaIpaddressResponse RecordrpzaaaaipaddressPost(ctx).RecordRpzAaaaIpaddress(recordRpzAaaaIpaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:aaaa:ipaddress object



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
	recordRpzAaaaIpaddress := *rpz.NewRecordRpzAaaaIpaddress() // RecordRpzAaaaIpaddress | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressPost(context.Background()).RecordRpzAaaaIpaddress(recordRpzAaaaIpaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaaaaipaddressPost`: CreateRecordRpzAaaaIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaIpaddressAPIRecordrpzaaaaipaddressPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzAaaaIpaddress** | [**RecordRpzAaaaIpaddress**](RecordRpzAaaaIpaddress.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzAaaaIpaddressResponse**](CreateRecordRpzAaaaIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzaaaaipaddressReferenceDelete

> RecordrpzaaaaipaddressReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:aaaa:ipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:aaaa:ipaddress object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:aaaa:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaIpaddressAPIRecordrpzaaaaipaddressReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpzaaaaipaddressReferenceGet

> GetRecordRpzAaaaIpaddressResponse RecordrpzaaaaipaddressReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:aaaa:ipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:aaaa:ipaddress object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaaaaipaddressReferenceGet`: GetRecordRpzAaaaIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:aaaa:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaIpaddressAPIRecordrpzaaaaipaddressReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzAaaaIpaddressResponse**](GetRecordRpzAaaaIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzaaaaipaddressReferencePut

> UpdateRecordRpzAaaaIpaddressResponse RecordrpzaaaaipaddressReferencePut(ctx, reference).RecordRpzAaaaIpaddress(recordRpzAaaaIpaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:aaaa:ipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:aaaa:ipaddress object
	recordRpzAaaaIpaddress := *rpz.NewRecordRpzAaaaIpaddress() // RecordRpzAaaaIpaddress | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressReferencePut(context.Background(), reference).RecordRpzAaaaIpaddress(recordRpzAaaaIpaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaaaaipaddressReferencePut`: UpdateRecordRpzAaaaIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAaaaIpaddressAPI.RecordrpzaaaaipaddressReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:aaaa:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaIpaddressAPIRecordrpzaaaaipaddressReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzAaaaIpaddress** | [**RecordRpzAaaaIpaddress**](RecordRpzAaaaIpaddress.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzAaaaIpaddressResponse**](UpdateRecordRpzAaaaIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

