# RecordRpzSrvAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpzsrvGet**](RecordRpzSrvAPI.md#RecordrpzsrvGet) | **Get** /record:rpz:srv | Retrieve record:rpz:srv objects
[**RecordrpzsrvPost**](RecordRpzSrvAPI.md#RecordrpzsrvPost) | **Post** /record:rpz:srv | Create a record:rpz:srv object
[**RecordrpzsrvReferenceDelete**](RecordRpzSrvAPI.md#RecordrpzsrvReferenceDelete) | **Delete** /record:rpz:srv/{reference} | Delete a record:rpz:srv object
[**RecordrpzsrvReferenceGet**](RecordRpzSrvAPI.md#RecordrpzsrvReferenceGet) | **Get** /record:rpz:srv/{reference} | Get a specific record:rpz:srv object
[**RecordrpzsrvReferencePut**](RecordRpzSrvAPI.md#RecordrpzsrvReferencePut) | **Put** /record:rpz:srv/{reference} | Update a record:rpz:srv object



## RecordrpzsrvGet

> ListRecordRpzSrvResponse RecordrpzsrvGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:srv objects



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
	resp, r, err := apiClient.RecordRpzSrvAPI.RecordrpzsrvGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzSrvAPI.RecordrpzsrvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzsrvGet`: ListRecordRpzSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzSrvAPI.RecordrpzsrvGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzSrvAPIRecordrpzsrvGetRequest` struct via the builder pattern


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

[**ListRecordRpzSrvResponse**](ListRecordRpzSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzsrvPost

> CreateRecordRpzSrvResponse RecordrpzsrvPost(ctx).RecordRpzSrv(recordRpzSrv).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:srv object



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
	recordRpzSrv := *rpz.NewRecordRpzSrv() // RecordRpzSrv | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzSrvAPI.RecordrpzsrvPost(context.Background()).RecordRpzSrv(recordRpzSrv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzSrvAPI.RecordrpzsrvPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzsrvPost`: CreateRecordRpzSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzSrvAPI.RecordrpzsrvPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzSrvAPIRecordrpzsrvPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzSrv** | [**RecordRpzSrv**](RecordRpzSrv.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzSrvResponse**](CreateRecordRpzSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzsrvReferenceDelete

> RecordrpzsrvReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:srv object



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
	reference := "reference_example" // string | Reference of the record:rpz:srv object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzSrvAPI.RecordrpzsrvReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzSrvAPI.RecordrpzsrvReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzSrvAPIRecordrpzsrvReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpzsrvReferenceGet

> GetRecordRpzSrvResponse RecordrpzsrvReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:srv object



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
	reference := "reference_example" // string | Reference of the record:rpz:srv object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzSrvAPI.RecordrpzsrvReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzSrvAPI.RecordrpzsrvReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzsrvReferenceGet`: GetRecordRpzSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzSrvAPI.RecordrpzsrvReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzSrvAPIRecordrpzsrvReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzSrvResponse**](GetRecordRpzSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzsrvReferencePut

> UpdateRecordRpzSrvResponse RecordrpzsrvReferencePut(ctx, reference).RecordRpzSrv(recordRpzSrv).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:srv object



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
	reference := "reference_example" // string | Reference of the record:rpz:srv object
	recordRpzSrv := *rpz.NewRecordRpzSrv() // RecordRpzSrv | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzSrvAPI.RecordrpzsrvReferencePut(context.Background(), reference).RecordRpzSrv(recordRpzSrv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzSrvAPI.RecordrpzsrvReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzsrvReferencePut`: UpdateRecordRpzSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzSrvAPI.RecordrpzsrvReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzSrvAPIRecordrpzsrvReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzSrv** | [**RecordRpzSrv**](RecordRpzSrv.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzSrvResponse**](UpdateRecordRpzSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

