# RecordHostAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](RecordHostAPI.md#Create) | **Post** /record:host | Create a record:host object
[**Delete**](RecordHostAPI.md#Delete) | **Delete** /record:host/{reference} | Delete a record:host object
[**List**](RecordHostAPI.md#List) | **Get** /record:host | Retrieve record:host objects
[**Read**](RecordHostAPI.md#Read) | **Get** /record:host/{reference} | Get a specific record:host object
[**Update**](RecordHostAPI.md#Update) | **Put** /record:host/{reference} | Update a record:host object



## Create

> CreateRecordHostResponse Create(ctx).RecordHost(recordHost).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a record:host object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	recordHost := *dns.NewRecordHost() // RecordHost | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordHostAPI.Create(context.Background()).RecordHost(recordHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateRecordHostResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordHostAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordHost** | [**RecordHost**](RecordHost.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordHostResponse**](CreateRecordHostResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, reference).Execute()

Delete a record:host object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the record:host object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordHostAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:host object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordHostAPIDeleteRequest` struct via the builder pattern


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


## List

> ListRecordHostResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve record:host objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordHostAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListRecordHostResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordHostAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**paging** | **int32** | Control paging of results | 
**pageId** | **string** | Page id for retrieving next page of results | 
**filters** | **map[string]interface{}** |  | 
**extattrfilter** | **map[string]interface{}** |  | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**ListRecordHostResponse**](ListRecordHostResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetRecordHostResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific record:host object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the record:host object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordHostAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetRecordHostResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:host object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordHostAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetRecordHostResponse**](GetRecordHostResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateRecordHostResponse Update(ctx, reference).RecordHost(recordHost).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a record:host object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the record:host object
	recordHost := *dns.NewRecordHost() // RecordHost | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordHostAPI.Update(context.Background(), reference).RecordHost(recordHost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateRecordHostResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:host object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordHostAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordHost** | [**RecordHost**](RecordHost.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordHostResponse**](UpdateRecordHostResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

