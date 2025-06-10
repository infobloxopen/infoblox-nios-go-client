# RecordDnameAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecorddnameGet**](RecordDnameAPI.md#RecorddnameGet) | **Get** /record:dname | Retrieve record:dname objects
[**RecorddnamePost**](RecordDnameAPI.md#RecorddnamePost) | **Post** /record:dname | Create a record:dname object
[**RecorddnameReferenceDelete**](RecordDnameAPI.md#RecorddnameReferenceDelete) | **Delete** /record:dname/{reference} | Delete a record:dname object
[**RecorddnameReferenceGet**](RecordDnameAPI.md#RecorddnameReferenceGet) | **Get** /record:dname/{reference} | Get a specific record:dname object
[**RecorddnameReferencePut**](RecordDnameAPI.md#RecorddnameReferencePut) | **Put** /record:dname/{reference} | Update a record:dname object



## RecorddnameGet

> ListRecordDnameResponse RecorddnameGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:dname objects



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
	resp, r, err := apiClient.RecordDnameAPI.RecorddnameGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordDnameAPI.RecorddnameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecorddnameGet`: ListRecordDnameResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordDnameAPI.RecorddnameGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordDnameAPIRecorddnameGetRequest` struct via the builder pattern


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

[**ListRecordDnameResponse**](ListRecordDnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecorddnamePost

> CreateRecordDnameResponse RecorddnamePost(ctx).RecordDname(recordDname).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:dname object



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
	recordDname := *dns.NewRecordDname() // RecordDname | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordDnameAPI.RecorddnamePost(context.Background()).RecordDname(recordDname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordDnameAPI.RecorddnamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecorddnamePost`: CreateRecordDnameResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordDnameAPI.RecorddnamePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordDnameAPIRecorddnamePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordDname** | [**RecordDname**](RecordDname.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordDnameResponse**](CreateRecordDnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecorddnameReferenceDelete

> RecorddnameReferenceDelete(ctx, reference).Execute()

Delete a record:dname object



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
	reference := "reference_example" // string | Reference of the record:dname object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordDnameAPI.RecorddnameReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordDnameAPI.RecorddnameReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:dname object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordDnameAPIRecorddnameReferenceDeleteRequest` struct via the builder pattern


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


## RecorddnameReferenceGet

> GetRecordDnameResponse RecorddnameReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:dname object



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
	reference := "reference_example" // string | Reference of the record:dname object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordDnameAPI.RecorddnameReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordDnameAPI.RecorddnameReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecorddnameReferenceGet`: GetRecordDnameResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordDnameAPI.RecorddnameReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:dname object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordDnameAPIRecorddnameReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordDnameResponse**](GetRecordDnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecorddnameReferencePut

> UpdateRecordDnameResponse RecorddnameReferencePut(ctx, reference).RecordDname(recordDname).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:dname object



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
	reference := "reference_example" // string | Reference of the record:dname object
	recordDname := *dns.NewRecordDname() // RecordDname | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordDnameAPI.RecorddnameReferencePut(context.Background(), reference).RecordDname(recordDname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordDnameAPI.RecorddnameReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecorddnameReferencePut`: UpdateRecordDnameResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordDnameAPI.RecorddnameReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:dname object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordDnameAPIRecorddnameReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordDname** | [**RecordDname**](RecordDname.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordDnameResponse**](UpdateRecordDnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

