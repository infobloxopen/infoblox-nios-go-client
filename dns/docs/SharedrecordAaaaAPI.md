# SharedrecordAaaaAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](SharedrecordAaaaAPI.md#Create) | **Post** /sharedrecord:aaaa | Create a sharedrecord:aaaa object
[**Delete**](SharedrecordAaaaAPI.md#Delete) | **Delete** /sharedrecord:aaaa/{reference} | Delete a sharedrecord:aaaa object
[**List**](SharedrecordAaaaAPI.md#List) | **Get** /sharedrecord:aaaa | Retrieve sharedrecord:aaaa objects
[**Read**](SharedrecordAaaaAPI.md#Read) | **Get** /sharedrecord:aaaa/{reference} | Get a specific sharedrecord:aaaa object
[**Update**](SharedrecordAaaaAPI.md#Update) | **Put** /sharedrecord:aaaa/{reference} | Update a sharedrecord:aaaa object



## Create

> CreateSharedrecordAaaaResponse Create(ctx).SharedrecordAaaa(sharedrecordAaaa).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a sharedrecord:aaaa object



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
	sharedrecordAaaa := *dns.NewSharedrecordAaaa() // SharedrecordAaaa | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordAaaaAPI.Create(context.Background()).SharedrecordAaaa(sharedrecordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateSharedrecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordAaaaAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordAaaa** | [**SharedrecordAaaa**](SharedrecordAaaa.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSharedrecordAaaaResponse**](CreateSharedrecordAaaaResponse.md)

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

Delete a sharedrecord:aaaa object



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
	reference := "reference_example" // string | Reference of the sharedrecord:aaaa object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.SharedrecordAaaaAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPIDeleteRequest` struct via the builder pattern


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

> ListSharedrecordAaaaResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve sharedrecord:aaaa objects



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
	resp, r, err := apiClient.SharedrecordAaaaAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListSharedrecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordAaaaAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPIListRequest` struct via the builder pattern


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

[**ListSharedrecordAaaaResponse**](ListSharedrecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetSharedrecordAaaaResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific sharedrecord:aaaa object



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
	reference := "reference_example" // string | Reference of the sharedrecord:aaaa object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordAaaaAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetSharedrecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordAaaaAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetSharedrecordAaaaResponse**](GetSharedrecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateSharedrecordAaaaResponse Update(ctx, reference).SharedrecordAaaa(sharedrecordAaaa).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a sharedrecord:aaaa object



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
	reference := "reference_example" // string | Reference of the sharedrecord:aaaa object
	sharedrecordAaaa := *dns.NewSharedrecordAaaa() // SharedrecordAaaa | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordAaaaAPI.Update(context.Background(), reference).SharedrecordAaaa(sharedrecordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordAaaaAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateSharedrecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordAaaaAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordAaaaAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordAaaa** | [**SharedrecordAaaa**](SharedrecordAaaa.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSharedrecordAaaaResponse**](UpdateSharedrecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

