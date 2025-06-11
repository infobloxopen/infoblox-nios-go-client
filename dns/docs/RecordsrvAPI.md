# RecordSrvAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordsrvGet**](RecordSrvAPI.md#RecordsrvGet) | **Get** /record:srv | Retrieve record:srv objects
[**RecordsrvPost**](RecordSrvAPI.md#RecordsrvPost) | **Post** /record:srv | Create a record:srv object
[**RecordsrvReferenceDelete**](RecordSrvAPI.md#RecordsrvReferenceDelete) | **Delete** /record:srv/{reference} | Delete a record:srv object
[**RecordsrvReferenceGet**](RecordSrvAPI.md#RecordsrvReferenceGet) | **Get** /record:srv/{reference} | Get a specific record:srv object
[**RecordsrvReferencePut**](RecordSrvAPI.md#RecordsrvReferencePut) | **Put** /record:srv/{reference} | Update a record:srv object



## RecordsrvGet

> ListRecordSrvResponse RecordsrvGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:srv objects



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
	resp, r, err := apiClient.RecordSrvAPI.RecordsrvGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordSrvAPI.RecordsrvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordsrvGet`: ListRecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordSrvAPI.RecordsrvGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordSrvAPIRecordsrvGetRequest` struct via the builder pattern


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

[**ListRecordSrvResponse**](ListRecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordsrvPost

> CreateRecordSrvResponse RecordsrvPost(ctx).RecordSrv(recordSrv).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:srv object



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
	recordSrv := *dns.NewRecordSrv() // RecordSrv | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordSrvAPI.RecordsrvPost(context.Background()).RecordSrv(recordSrv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordSrvAPI.RecordsrvPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordsrvPost`: CreateRecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordSrvAPI.RecordsrvPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordSrvAPIRecordsrvPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordSrv** | [**RecordSrv**](RecordSrv.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordSrvResponse**](CreateRecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordsrvReferenceDelete

> RecordsrvReferenceDelete(ctx, reference).Execute()

Delete a record:srv object



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
	reference := "reference_example" // string | Reference of the record:srv object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordSrvAPI.RecordsrvReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordSrvAPI.RecordsrvReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordSrvAPIRecordsrvReferenceDeleteRequest` struct via the builder pattern


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


## RecordsrvReferenceGet

> GetRecordSrvResponse RecordsrvReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:srv object



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
	reference := "reference_example" // string | Reference of the record:srv object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordSrvAPI.RecordsrvReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordSrvAPI.RecordsrvReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordsrvReferenceGet`: GetRecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordSrvAPI.RecordsrvReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordSrvAPIRecordsrvReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordSrvResponse**](GetRecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordsrvReferencePut

> UpdateRecordSrvResponse RecordsrvReferencePut(ctx, reference).RecordSrv(recordSrv).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:srv object



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
	reference := "reference_example" // string | Reference of the record:srv object
	recordSrv := *dns.NewRecordSrv() // RecordSrv | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordSrvAPI.RecordsrvReferencePut(context.Background(), reference).RecordSrv(recordSrv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordSrvAPI.RecordsrvReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordsrvReferencePut`: UpdateRecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordSrvAPI.RecordsrvReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordSrvAPIRecordsrvReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordSrv** | [**RecordSrv**](RecordSrv.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordSrvResponse**](UpdateRecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

