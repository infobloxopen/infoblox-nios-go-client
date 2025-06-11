# RecordTxtAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordtxtGet**](RecordTxtAPI.md#RecordtxtGet) | **Get** /record:txt | Retrieve record:txt objects
[**RecordtxtPost**](RecordTxtAPI.md#RecordtxtPost) | **Post** /record:txt | Create a record:txt object
[**RecordtxtReferenceDelete**](RecordTxtAPI.md#RecordtxtReferenceDelete) | **Delete** /record:txt/{reference} | Delete a record:txt object
[**RecordtxtReferenceGet**](RecordTxtAPI.md#RecordtxtReferenceGet) | **Get** /record:txt/{reference} | Get a specific record:txt object
[**RecordtxtReferencePut**](RecordTxtAPI.md#RecordtxtReferencePut) | **Put** /record:txt/{reference} | Update a record:txt object



## RecordtxtGet

> ListRecordTxtResponse RecordtxtGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:txt objects



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
	resp, r, err := apiClient.RecordTxtAPI.RecordtxtGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTxtAPI.RecordtxtGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordtxtGet`: ListRecordTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordTxtAPI.RecordtxtGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordTxtAPIRecordtxtGetRequest` struct via the builder pattern


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

[**ListRecordTxtResponse**](ListRecordTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordtxtPost

> CreateRecordTxtResponse RecordtxtPost(ctx).RecordTxt(recordTxt).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:txt object



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
	recordTxt := *dns.NewRecordTxt() // RecordTxt | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordTxtAPI.RecordtxtPost(context.Background()).RecordTxt(recordTxt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTxtAPI.RecordtxtPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordtxtPost`: CreateRecordTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordTxtAPI.RecordtxtPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordTxtAPIRecordtxtPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordTxt** | [**RecordTxt**](RecordTxt.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordTxtResponse**](CreateRecordTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordtxtReferenceDelete

> RecordtxtReferenceDelete(ctx, reference).Execute()

Delete a record:txt object



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
	reference := "reference_example" // string | Reference of the record:txt object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordTxtAPI.RecordtxtReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTxtAPI.RecordtxtReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:txt object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordTxtAPIRecordtxtReferenceDeleteRequest` struct via the builder pattern


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


## RecordtxtReferenceGet

> GetRecordTxtResponse RecordtxtReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:txt object



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
	reference := "reference_example" // string | Reference of the record:txt object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordTxtAPI.RecordtxtReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTxtAPI.RecordtxtReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordtxtReferenceGet`: GetRecordTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordTxtAPI.RecordtxtReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:txt object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordTxtAPIRecordtxtReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordTxtResponse**](GetRecordTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordtxtReferencePut

> UpdateRecordTxtResponse RecordtxtReferencePut(ctx, reference).RecordTxt(recordTxt).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:txt object



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
	reference := "reference_example" // string | Reference of the record:txt object
	recordTxt := *dns.NewRecordTxt() // RecordTxt | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordTxtAPI.RecordtxtReferencePut(context.Background(), reference).RecordTxt(recordTxt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTxtAPI.RecordtxtReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordtxtReferencePut`: UpdateRecordTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordTxtAPI.RecordtxtReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:txt object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordTxtAPIRecordtxtReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordTxt** | [**RecordTxt**](RecordTxt.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordTxtResponse**](UpdateRecordTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

