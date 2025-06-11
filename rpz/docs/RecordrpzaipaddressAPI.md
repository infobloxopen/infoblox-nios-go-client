# RecordRpzAIpaddressAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpzaipaddressGet**](RecordRpzAIpaddressAPI.md#RecordrpzaipaddressGet) | **Get** /record:rpz:a:ipaddress | Retrieve record:rpz:a:ipaddress objects
[**RecordrpzaipaddressPost**](RecordRpzAIpaddressAPI.md#RecordrpzaipaddressPost) | **Post** /record:rpz:a:ipaddress | Create a record:rpz:a:ipaddress object
[**RecordrpzaipaddressReferenceDelete**](RecordRpzAIpaddressAPI.md#RecordrpzaipaddressReferenceDelete) | **Delete** /record:rpz:a:ipaddress/{reference} | Delete a record:rpz:a:ipaddress object
[**RecordrpzaipaddressReferenceGet**](RecordRpzAIpaddressAPI.md#RecordrpzaipaddressReferenceGet) | **Get** /record:rpz:a:ipaddress/{reference} | Get a specific record:rpz:a:ipaddress object
[**RecordrpzaipaddressReferencePut**](RecordRpzAIpaddressAPI.md#RecordrpzaipaddressReferencePut) | **Put** /record:rpz:a:ipaddress/{reference} | Update a record:rpz:a:ipaddress object



## RecordrpzaipaddressGet

> ListRecordRpzAIpaddressResponse RecordrpzaipaddressGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:a:ipaddress objects



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
	resp, r, err := apiClient.RecordRpzAIpaddressAPI.RecordrpzaipaddressGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAIpaddressAPI.RecordrpzaipaddressGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaipaddressGet`: ListRecordRpzAIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAIpaddressAPI.RecordrpzaipaddressGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAIpaddressAPIRecordrpzaipaddressGetRequest` struct via the builder pattern


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

[**ListRecordRpzAIpaddressResponse**](ListRecordRpzAIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzaipaddressPost

> CreateRecordRpzAIpaddressResponse RecordrpzaipaddressPost(ctx).RecordRpzAIpaddress(recordRpzAIpaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:a:ipaddress object



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
	recordRpzAIpaddress := *rpz.NewRecordRpzAIpaddress() // RecordRpzAIpaddress | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzAIpaddressAPI.RecordrpzaipaddressPost(context.Background()).RecordRpzAIpaddress(recordRpzAIpaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAIpaddressAPI.RecordrpzaipaddressPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaipaddressPost`: CreateRecordRpzAIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAIpaddressAPI.RecordrpzaipaddressPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAIpaddressAPIRecordrpzaipaddressPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzAIpaddress** | [**RecordRpzAIpaddress**](RecordRpzAIpaddress.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzAIpaddressResponse**](CreateRecordRpzAIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzaipaddressReferenceDelete

> RecordrpzaipaddressReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:a:ipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:a:ipaddress object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzAIpaddressAPI.RecordrpzaipaddressReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAIpaddressAPI.RecordrpzaipaddressReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:a:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAIpaddressAPIRecordrpzaipaddressReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpzaipaddressReferenceGet

> GetRecordRpzAIpaddressResponse RecordrpzaipaddressReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:a:ipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:a:ipaddress object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzAIpaddressAPI.RecordrpzaipaddressReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAIpaddressAPI.RecordrpzaipaddressReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaipaddressReferenceGet`: GetRecordRpzAIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAIpaddressAPI.RecordrpzaipaddressReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:a:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAIpaddressAPIRecordrpzaipaddressReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzAIpaddressResponse**](GetRecordRpzAIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzaipaddressReferencePut

> UpdateRecordRpzAIpaddressResponse RecordrpzaipaddressReferencePut(ctx, reference).RecordRpzAIpaddress(recordRpzAIpaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:a:ipaddress object



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
	reference := "reference_example" // string | Reference of the record:rpz:a:ipaddress object
	recordRpzAIpaddress := *rpz.NewRecordRpzAIpaddress() // RecordRpzAIpaddress | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzAIpaddressAPI.RecordrpzaipaddressReferencePut(context.Background(), reference).RecordRpzAIpaddress(recordRpzAIpaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAIpaddressAPI.RecordrpzaipaddressReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaipaddressReferencePut`: UpdateRecordRpzAIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAIpaddressAPI.RecordrpzaipaddressReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:a:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAIpaddressAPIRecordrpzaipaddressReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzAIpaddress** | [**RecordRpzAIpaddress**](RecordRpzAIpaddress.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzAIpaddressResponse**](UpdateRecordRpzAIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

