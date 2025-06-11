# RecordAliasAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordaliasGet**](RecordAliasAPI.md#RecordaliasGet) | **Get** /record:alias | Retrieve record:alias objects
[**RecordaliasPost**](RecordAliasAPI.md#RecordaliasPost) | **Post** /record:alias | Create a record:alias object
[**RecordaliasReferenceDelete**](RecordAliasAPI.md#RecordaliasReferenceDelete) | **Delete** /record:alias/{reference} | Delete a record:alias object
[**RecordaliasReferenceGet**](RecordAliasAPI.md#RecordaliasReferenceGet) | **Get** /record:alias/{reference} | Get a specific record:alias object
[**RecordaliasReferencePut**](RecordAliasAPI.md#RecordaliasReferencePut) | **Put** /record:alias/{reference} | Update a record:alias object



## RecordaliasGet

> ListRecordAliasResponse RecordaliasGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:alias objects



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
	resp, r, err := apiClient.RecordAliasAPI.RecordaliasGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAliasAPI.RecordaliasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaliasGet`: ListRecordAliasResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAliasAPI.RecordaliasGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordAliasAPIRecordaliasGetRequest` struct via the builder pattern


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

[**ListRecordAliasResponse**](ListRecordAliasResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordaliasPost

> CreateRecordAliasResponse RecordaliasPost(ctx).RecordAlias(recordAlias).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:alias object



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
	recordAlias := *dns.NewRecordAlias() // RecordAlias | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordAliasAPI.RecordaliasPost(context.Background()).RecordAlias(recordAlias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAliasAPI.RecordaliasPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaliasPost`: CreateRecordAliasResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAliasAPI.RecordaliasPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordAliasAPIRecordaliasPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordAlias** | [**RecordAlias**](RecordAlias.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordAliasResponse**](CreateRecordAliasResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordaliasReferenceDelete

> RecordaliasReferenceDelete(ctx, reference).Execute()

Delete a record:alias object



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
	reference := "reference_example" // string | Reference of the record:alias object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordAliasAPI.RecordaliasReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAliasAPI.RecordaliasReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:alias object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordAliasAPIRecordaliasReferenceDeleteRequest` struct via the builder pattern


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


## RecordaliasReferenceGet

> GetRecordAliasResponse RecordaliasReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:alias object



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
	reference := "reference_example" // string | Reference of the record:alias object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordAliasAPI.RecordaliasReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAliasAPI.RecordaliasReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaliasReferenceGet`: GetRecordAliasResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAliasAPI.RecordaliasReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:alias object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordAliasAPIRecordaliasReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordAliasResponse**](GetRecordAliasResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordaliasReferencePut

> UpdateRecordAliasResponse RecordaliasReferencePut(ctx, reference).RecordAlias(recordAlias).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:alias object



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
	reference := "reference_example" // string | Reference of the record:alias object
	recordAlias := *dns.NewRecordAlias() // RecordAlias | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordAliasAPI.RecordaliasReferencePut(context.Background(), reference).RecordAlias(recordAlias).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAliasAPI.RecordaliasReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaliasReferencePut`: UpdateRecordAliasResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAliasAPI.RecordaliasReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:alias object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordAliasAPIRecordaliasReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordAlias** | [**RecordAlias**](RecordAlias.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordAliasResponse**](UpdateRecordAliasResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

