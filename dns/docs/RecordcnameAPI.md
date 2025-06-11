# RecordCnameAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordcnameGet**](RecordCnameAPI.md#RecordcnameGet) | **Get** /record:cname | Retrieve record:cname objects
[**RecordcnamePost**](RecordCnameAPI.md#RecordcnamePost) | **Post** /record:cname | Create a record:cname object
[**RecordcnameReferenceDelete**](RecordCnameAPI.md#RecordcnameReferenceDelete) | **Delete** /record:cname/{reference} | Delete a record:cname object
[**RecordcnameReferenceGet**](RecordCnameAPI.md#RecordcnameReferenceGet) | **Get** /record:cname/{reference} | Get a specific record:cname object
[**RecordcnameReferencePut**](RecordCnameAPI.md#RecordcnameReferencePut) | **Put** /record:cname/{reference} | Update a record:cname object



## RecordcnameGet

> ListRecordCnameResponse RecordcnameGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:cname objects



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
	resp, r, err := apiClient.RecordCnameAPI.RecordcnameGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCnameAPI.RecordcnameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordcnameGet`: ListRecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordCnameAPI.RecordcnameGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordCnameAPIRecordcnameGetRequest` struct via the builder pattern


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

[**ListRecordCnameResponse**](ListRecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordcnamePost

> CreateRecordCnameResponse RecordcnamePost(ctx).RecordCname(recordCname).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:cname object



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
	recordCname := *dns.NewRecordCname() // RecordCname | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordCnameAPI.RecordcnamePost(context.Background()).RecordCname(recordCname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCnameAPI.RecordcnamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordcnamePost`: CreateRecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordCnameAPI.RecordcnamePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordCnameAPIRecordcnamePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordCname** | [**RecordCname**](RecordCname.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordCnameResponse**](CreateRecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordcnameReferenceDelete

> RecordcnameReferenceDelete(ctx, reference).Execute()

Delete a record:cname object



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
	reference := "reference_example" // string | Reference of the record:cname object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordCnameAPI.RecordcnameReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCnameAPI.RecordcnameReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:cname object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordCnameAPIRecordcnameReferenceDeleteRequest` struct via the builder pattern


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


## RecordcnameReferenceGet

> GetRecordCnameResponse RecordcnameReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:cname object



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
	reference := "reference_example" // string | Reference of the record:cname object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordCnameAPI.RecordcnameReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCnameAPI.RecordcnameReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordcnameReferenceGet`: GetRecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordCnameAPI.RecordcnameReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:cname object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordCnameAPIRecordcnameReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordCnameResponse**](GetRecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordcnameReferencePut

> UpdateRecordCnameResponse RecordcnameReferencePut(ctx, reference).RecordCname(recordCname).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:cname object



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
	reference := "reference_example" // string | Reference of the record:cname object
	recordCname := *dns.NewRecordCname() // RecordCname | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordCnameAPI.RecordcnameReferencePut(context.Background(), reference).RecordCname(recordCname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordCnameAPI.RecordcnameReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordcnameReferencePut`: UpdateRecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordCnameAPI.RecordcnameReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:cname object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordCnameAPIRecordcnameReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordCname** | [**RecordCname**](RecordCname.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordCnameResponse**](UpdateRecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

