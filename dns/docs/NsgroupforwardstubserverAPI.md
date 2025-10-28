# NsgroupForwardstubserverAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](NsgroupForwardstubserverAPI.md#Create) | **Post** /nsgroup:forwardstubserver | Create a nsgroup:forwardstubserver object
[**Delete**](NsgroupForwardstubserverAPI.md#Delete) | **Delete** /nsgroup:forwardstubserver/{reference} | Delete a nsgroup:forwardstubserver object
[**List**](NsgroupForwardstubserverAPI.md#List) | **Get** /nsgroup:forwardstubserver | Retrieve nsgroup:forwardstubserver objects
[**Read**](NsgroupForwardstubserverAPI.md#Read) | **Get** /nsgroup:forwardstubserver/{reference} | Get a specific nsgroup:forwardstubserver object
[**Update**](NsgroupForwardstubserverAPI.md#Update) | **Put** /nsgroup:forwardstubserver/{reference} | Update a nsgroup:forwardstubserver object



## Create

> CreateNsgroupForwardstubserverResponse Create(ctx).NsgroupForwardstubserver(nsgroupForwardstubserver).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a nsgroup:forwardstubserver object



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
	nsgroupForwardstubserver := *dns.NewNsgroupForwardstubserver() // NsgroupForwardstubserver | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardstubserverAPI.Create(context.Background()).NsgroupForwardstubserver(nsgroupForwardstubserver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardstubserverAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupForwardstubserver** | [**NsgroupForwardstubserver**](NsgroupForwardstubserver.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateNsgroupForwardstubserverResponse**](CreateNsgroupForwardstubserverResponse.md)

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

Delete a nsgroup:forwardstubserver object



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
	reference := "reference_example" // string | Reference of the nsgroup:forwardstubserver object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.NsgroupForwardstubserverAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardstubserver object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPIDeleteRequest` struct via the builder pattern


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

> ListNsgroupForwardstubserverResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve nsgroup:forwardstubserver objects



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
	resp, r, err := apiClient.NsgroupForwardstubserverAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardstubserverAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPIListRequest` struct via the builder pattern


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

[**ListNsgroupForwardstubserverResponse**](ListNsgroupForwardstubserverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetNsgroupForwardstubserverResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific nsgroup:forwardstubserver object



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
	reference := "reference_example" // string | Reference of the nsgroup:forwardstubserver object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardstubserverAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardstubserverAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardstubserver object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetNsgroupForwardstubserverResponse**](GetNsgroupForwardstubserverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateNsgroupForwardstubserverResponse Update(ctx, reference).NsgroupForwardstubserver(nsgroupForwardstubserver).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a nsgroup:forwardstubserver object



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
	reference := "reference_example" // string | Reference of the nsgroup:forwardstubserver object
	nsgroupForwardstubserver := *dns.NewNsgroupForwardstubserver() // NsgroupForwardstubserver | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardstubserverAPI.Update(context.Background(), reference).NsgroupForwardstubserver(nsgroupForwardstubserver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardstubserverAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardstubserver object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupForwardstubserver** | [**NsgroupForwardstubserver**](NsgroupForwardstubserver.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateNsgroupForwardstubserverResponse**](UpdateNsgroupForwardstubserverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

