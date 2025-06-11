# RecordRpzTxtAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpztxtGet**](RecordRpzTxtAPI.md#RecordrpztxtGet) | **Get** /record:rpz:txt | Retrieve record:rpz:txt objects
[**RecordrpztxtPost**](RecordRpzTxtAPI.md#RecordrpztxtPost) | **Post** /record:rpz:txt | Create a record:rpz:txt object
[**RecordrpztxtReferenceDelete**](RecordRpzTxtAPI.md#RecordrpztxtReferenceDelete) | **Delete** /record:rpz:txt/{reference} | Delete a record:rpz:txt object
[**RecordrpztxtReferenceGet**](RecordRpzTxtAPI.md#RecordrpztxtReferenceGet) | **Get** /record:rpz:txt/{reference} | Get a specific record:rpz:txt object
[**RecordrpztxtReferencePut**](RecordRpzTxtAPI.md#RecordrpztxtReferencePut) | **Put** /record:rpz:txt/{reference} | Update a record:rpz:txt object



## RecordrpztxtGet

> ListRecordRpzTxtResponse RecordrpztxtGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:txt objects



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
	resp, r, err := apiClient.RecordRpzTxtAPI.RecordrpztxtGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzTxtAPI.RecordrpztxtGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpztxtGet`: ListRecordRpzTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzTxtAPI.RecordrpztxtGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzTxtAPIRecordrpztxtGetRequest` struct via the builder pattern


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

[**ListRecordRpzTxtResponse**](ListRecordRpzTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpztxtPost

> CreateRecordRpzTxtResponse RecordrpztxtPost(ctx).RecordRpzTxt(recordRpzTxt).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:txt object



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
	recordRpzTxt := *rpz.NewRecordRpzTxt() // RecordRpzTxt | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzTxtAPI.RecordrpztxtPost(context.Background()).RecordRpzTxt(recordRpzTxt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzTxtAPI.RecordrpztxtPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpztxtPost`: CreateRecordRpzTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzTxtAPI.RecordrpztxtPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzTxtAPIRecordrpztxtPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzTxt** | [**RecordRpzTxt**](RecordRpzTxt.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzTxtResponse**](CreateRecordRpzTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpztxtReferenceDelete

> RecordrpztxtReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:txt object



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
	reference := "reference_example" // string | Reference of the record:rpz:txt object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzTxtAPI.RecordrpztxtReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzTxtAPI.RecordrpztxtReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:txt object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzTxtAPIRecordrpztxtReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpztxtReferenceGet

> GetRecordRpzTxtResponse RecordrpztxtReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:txt object



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
	reference := "reference_example" // string | Reference of the record:rpz:txt object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzTxtAPI.RecordrpztxtReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzTxtAPI.RecordrpztxtReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpztxtReferenceGet`: GetRecordRpzTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzTxtAPI.RecordrpztxtReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:txt object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzTxtAPIRecordrpztxtReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzTxtResponse**](GetRecordRpzTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpztxtReferencePut

> UpdateRecordRpzTxtResponse RecordrpztxtReferencePut(ctx, reference).RecordRpzTxt(recordRpzTxt).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:txt object



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
	reference := "reference_example" // string | Reference of the record:rpz:txt object
	recordRpzTxt := *rpz.NewRecordRpzTxt() // RecordRpzTxt | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzTxtAPI.RecordrpztxtReferencePut(context.Background(), reference).RecordRpzTxt(recordRpzTxt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzTxtAPI.RecordrpztxtReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpztxtReferencePut`: UpdateRecordRpzTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzTxtAPI.RecordrpztxtReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:txt object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzTxtAPIRecordrpztxtReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzTxt** | [**RecordRpzTxt**](RecordRpzTxt.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzTxtResponse**](UpdateRecordRpzTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

