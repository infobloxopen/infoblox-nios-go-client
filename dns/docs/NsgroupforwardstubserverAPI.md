# NsgroupforwardstubserverAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](NsgroupforwardstubserverAPI.md#Get) | **Get** /nsgroup:forwardstubserver | Retrieve nsgroup:forwardstubserver objects
[**Post**](NsgroupforwardstubserverAPI.md#Post) | **Post** /nsgroup:forwardstubserver | Create a nsgroup:forwardstubserver object
[**ReferenceDelete**](NsgroupforwardstubserverAPI.md#ReferenceDelete) | **Delete** /nsgroup:forwardstubserver/{reference} | Delete a nsgroup:forwardstubserver object
[**ReferenceGet**](NsgroupforwardstubserverAPI.md#ReferenceGet) | **Get** /nsgroup:forwardstubserver/{reference} | Get a specific nsgroup:forwardstubserver object
[**ReferencePut**](NsgroupforwardstubserverAPI.md#ReferencePut) | **Put** /nsgroup:forwardstubserver/{reference} | Update a nsgroup:forwardstubserver object



## Get

> ListNsgroupForwardstubserverResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve nsgroup:forwardstubserver objects



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
	resp, r, err := apiClient.NsgroupforwardstubserverAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardstubserverAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupforwardstubserverAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupforwardstubserverAPIGetRequest` struct via the builder pattern


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

[**ListNsgroupForwardstubserverResponse**](ListNsgroupForwardstubserverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateNsgroupForwardstubserverResponse Post(ctx).NsgroupForwardstubserver(nsgroupForwardstubserver).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a nsgroup:forwardstubserver object



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
	nsgroupForwardstubserver := *dns.NewNsgroupForwardstubserver() // NsgroupForwardstubserver | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupforwardstubserverAPI.Post(context.Background()).NsgroupForwardstubserver(nsgroupForwardstubserver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardstubserverAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupforwardstubserverAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupforwardstubserverAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupForwardstubserver** | [**NsgroupForwardstubserver**](NsgroupForwardstubserver.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

Delete a nsgroup:forwardstubserver object



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
	reference := "reference_example" // string | Reference of the nsgroup:forwardstubserver object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.NsgroupforwardstubserverAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardstubserverAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `NsgroupforwardstubserverAPIReferenceDeleteRequest` struct via the builder pattern


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


## ReferenceGet

> GetNsgroupForwardstubserverResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific nsgroup:forwardstubserver object



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
	reference := "reference_example" // string | Reference of the nsgroup:forwardstubserver object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupforwardstubserverAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardstubserverAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupforwardstubserverAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardstubserver object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupforwardstubserverAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## ReferencePut

> UpdateNsgroupForwardstubserverResponse ReferencePut(ctx, reference).NsgroupForwardstubserver(nsgroupForwardstubserver).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a nsgroup:forwardstubserver object



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
	reference := "reference_example" // string | Reference of the nsgroup:forwardstubserver object
	nsgroupForwardstubserver := *dns.NewNsgroupForwardstubserver() // NsgroupForwardstubserver | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupforwardstubserverAPI.ReferencePut(context.Background(), reference).NsgroupForwardstubserver(nsgroupForwardstubserver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardstubserverAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupforwardstubserverAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardstubserver object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupforwardstubserverAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupForwardstubserver** | [**NsgroupForwardstubserver**](NsgroupForwardstubserver.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

