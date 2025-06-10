# NsgroupForwardstubserverAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NsgroupforwardstubserverGet**](NsgroupForwardstubserverAPI.md#NsgroupforwardstubserverGet) | **Get** /nsgroup:forwardstubserver | Retrieve nsgroup:forwardstubserver objects
[**NsgroupforwardstubserverPost**](NsgroupForwardstubserverAPI.md#NsgroupforwardstubserverPost) | **Post** /nsgroup:forwardstubserver | Create a nsgroup:forwardstubserver object
[**NsgroupforwardstubserverReferenceDelete**](NsgroupForwardstubserverAPI.md#NsgroupforwardstubserverReferenceDelete) | **Delete** /nsgroup:forwardstubserver/{reference} | Delete a nsgroup:forwardstubserver object
[**NsgroupforwardstubserverReferenceGet**](NsgroupForwardstubserverAPI.md#NsgroupforwardstubserverReferenceGet) | **Get** /nsgroup:forwardstubserver/{reference} | Get a specific nsgroup:forwardstubserver object
[**NsgroupforwardstubserverReferencePut**](NsgroupForwardstubserverAPI.md#NsgroupforwardstubserverReferencePut) | **Put** /nsgroup:forwardstubserver/{reference} | Update a nsgroup:forwardstubserver object



## NsgroupforwardstubserverGet

> ListNsgroupForwardstubserverResponse NsgroupforwardstubserverGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

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
	resp, r, err := apiClient.NsgroupForwardstubserverAPI.NsgroupforwardstubserverGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.NsgroupforwardstubserverGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupforwardstubserverGet`: ListNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardstubserverAPI.NsgroupforwardstubserverGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPINsgroupforwardstubserverGetRequest` struct via the builder pattern


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


## NsgroupforwardstubserverPost

> CreateNsgroupForwardstubserverResponse NsgroupforwardstubserverPost(ctx).NsgroupForwardstubserver(nsgroupForwardstubserver).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.NsgroupForwardstubserverAPI.NsgroupforwardstubserverPost(context.Background()).NsgroupForwardstubserver(nsgroupForwardstubserver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.NsgroupforwardstubserverPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupforwardstubserverPost`: CreateNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardstubserverAPI.NsgroupforwardstubserverPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPINsgroupforwardstubserverPostRequest` struct via the builder pattern


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


## NsgroupforwardstubserverReferenceDelete

> NsgroupforwardstubserverReferenceDelete(ctx, reference).Execute()

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
	r, err := apiClient.NsgroupForwardstubserverAPI.NsgroupforwardstubserverReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.NsgroupforwardstubserverReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPINsgroupforwardstubserverReferenceDeleteRequest` struct via the builder pattern


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


## NsgroupforwardstubserverReferenceGet

> GetNsgroupForwardstubserverResponse NsgroupforwardstubserverReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.NsgroupForwardstubserverAPI.NsgroupforwardstubserverReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.NsgroupforwardstubserverReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupforwardstubserverReferenceGet`: GetNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardstubserverAPI.NsgroupforwardstubserverReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardstubserver object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPINsgroupforwardstubserverReferenceGetRequest` struct via the builder pattern


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


## NsgroupforwardstubserverReferencePut

> UpdateNsgroupForwardstubserverResponse NsgroupforwardstubserverReferencePut(ctx, reference).NsgroupForwardstubserver(nsgroupForwardstubserver).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.NsgroupForwardstubserverAPI.NsgroupforwardstubserverReferencePut(context.Background(), reference).NsgroupForwardstubserver(nsgroupForwardstubserver).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardstubserverAPI.NsgroupforwardstubserverReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupforwardstubserverReferencePut`: UpdateNsgroupForwardstubserverResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardstubserverAPI.NsgroupforwardstubserverReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardstubserver object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardstubserverAPINsgroupforwardstubserverReferencePutRequest` struct via the builder pattern


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

