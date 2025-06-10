# RecordTlsaAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordtlsaGet**](RecordTlsaAPI.md#RecordtlsaGet) | **Get** /record:tlsa | Retrieve record:tlsa objects
[**RecordtlsaPost**](RecordTlsaAPI.md#RecordtlsaPost) | **Post** /record:tlsa | Create a record:tlsa object
[**RecordtlsaReferenceDelete**](RecordTlsaAPI.md#RecordtlsaReferenceDelete) | **Delete** /record:tlsa/{reference} | Delete a record:tlsa object
[**RecordtlsaReferenceGet**](RecordTlsaAPI.md#RecordtlsaReferenceGet) | **Get** /record:tlsa/{reference} | Get a specific record:tlsa object
[**RecordtlsaReferencePut**](RecordTlsaAPI.md#RecordtlsaReferencePut) | **Put** /record:tlsa/{reference} | Update a record:tlsa object



## RecordtlsaGet

> ListRecordTlsaResponse RecordtlsaGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:tlsa objects



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
	resp, r, err := apiClient.RecordTlsaAPI.RecordtlsaGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTlsaAPI.RecordtlsaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordtlsaGet`: ListRecordTlsaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordTlsaAPI.RecordtlsaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordTlsaAPIRecordtlsaGetRequest` struct via the builder pattern


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

[**ListRecordTlsaResponse**](ListRecordTlsaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordtlsaPost

> CreateRecordTlsaResponse RecordtlsaPost(ctx).RecordTlsa(recordTlsa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:tlsa object



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
	recordTlsa := *dns.NewRecordTlsa() // RecordTlsa | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordTlsaAPI.RecordtlsaPost(context.Background()).RecordTlsa(recordTlsa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTlsaAPI.RecordtlsaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordtlsaPost`: CreateRecordTlsaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordTlsaAPI.RecordtlsaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordTlsaAPIRecordtlsaPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordTlsa** | [**RecordTlsa**](RecordTlsa.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordTlsaResponse**](CreateRecordTlsaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordtlsaReferenceDelete

> RecordtlsaReferenceDelete(ctx, reference).Execute()

Delete a record:tlsa object



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
	reference := "reference_example" // string | Reference of the record:tlsa object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordTlsaAPI.RecordtlsaReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTlsaAPI.RecordtlsaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:tlsa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordTlsaAPIRecordtlsaReferenceDeleteRequest` struct via the builder pattern


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


## RecordtlsaReferenceGet

> GetRecordTlsaResponse RecordtlsaReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:tlsa object



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
	reference := "reference_example" // string | Reference of the record:tlsa object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordTlsaAPI.RecordtlsaReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTlsaAPI.RecordtlsaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordtlsaReferenceGet`: GetRecordTlsaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordTlsaAPI.RecordtlsaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:tlsa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordTlsaAPIRecordtlsaReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordTlsaResponse**](GetRecordTlsaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordtlsaReferencePut

> UpdateRecordTlsaResponse RecordtlsaReferencePut(ctx, reference).RecordTlsa(recordTlsa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:tlsa object



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
	reference := "reference_example" // string | Reference of the record:tlsa object
	recordTlsa := *dns.NewRecordTlsa() // RecordTlsa | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordTlsaAPI.RecordtlsaReferencePut(context.Background(), reference).RecordTlsa(recordTlsa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordTlsaAPI.RecordtlsaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordtlsaReferencePut`: UpdateRecordTlsaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordTlsaAPI.RecordtlsaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:tlsa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordTlsaAPIRecordtlsaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordTlsa** | [**RecordTlsa**](RecordTlsa.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordTlsaResponse**](UpdateRecordTlsaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

