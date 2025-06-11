# NsgroupForwardingmemberAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NsgroupforwardingmemberGet**](NsgroupForwardingmemberAPI.md#NsgroupforwardingmemberGet) | **Get** /nsgroup:forwardingmember | Retrieve nsgroup:forwardingmember objects
[**NsgroupforwardingmemberPost**](NsgroupForwardingmemberAPI.md#NsgroupforwardingmemberPost) | **Post** /nsgroup:forwardingmember | Create a nsgroup:forwardingmember object
[**NsgroupforwardingmemberReferenceDelete**](NsgroupForwardingmemberAPI.md#NsgroupforwardingmemberReferenceDelete) | **Delete** /nsgroup:forwardingmember/{reference} | Delete a nsgroup:forwardingmember object
[**NsgroupforwardingmemberReferenceGet**](NsgroupForwardingmemberAPI.md#NsgroupforwardingmemberReferenceGet) | **Get** /nsgroup:forwardingmember/{reference} | Get a specific nsgroup:forwardingmember object
[**NsgroupforwardingmemberReferencePut**](NsgroupForwardingmemberAPI.md#NsgroupforwardingmemberReferencePut) | **Put** /nsgroup:forwardingmember/{reference} | Update a nsgroup:forwardingmember object



## NsgroupforwardingmemberGet

> ListNsgroupForwardingmemberResponse NsgroupforwardingmemberGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve nsgroup:forwardingmember objects



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
	resp, r, err := apiClient.NsgroupForwardingmemberAPI.NsgroupforwardingmemberGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.NsgroupforwardingmemberGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupforwardingmemberGet`: ListNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardingmemberAPI.NsgroupforwardingmemberGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPINsgroupforwardingmemberGetRequest` struct via the builder pattern


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

[**ListNsgroupForwardingmemberResponse**](ListNsgroupForwardingmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NsgroupforwardingmemberPost

> CreateNsgroupForwardingmemberResponse NsgroupforwardingmemberPost(ctx).NsgroupForwardingmember(nsgroupForwardingmember).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a nsgroup:forwardingmember object



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
	nsgroupForwardingmember := *dns.NewNsgroupForwardingmember() // NsgroupForwardingmember | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardingmemberAPI.NsgroupforwardingmemberPost(context.Background()).NsgroupForwardingmember(nsgroupForwardingmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.NsgroupforwardingmemberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupforwardingmemberPost`: CreateNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardingmemberAPI.NsgroupforwardingmemberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPINsgroupforwardingmemberPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupForwardingmember** | [**NsgroupForwardingmember**](NsgroupForwardingmember.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateNsgroupForwardingmemberResponse**](CreateNsgroupForwardingmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NsgroupforwardingmemberReferenceDelete

> NsgroupforwardingmemberReferenceDelete(ctx, reference).Execute()

Delete a nsgroup:forwardingmember object



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
	reference := "reference_example" // string | Reference of the nsgroup:forwardingmember object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.NsgroupForwardingmemberAPI.NsgroupforwardingmemberReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.NsgroupforwardingmemberReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardingmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPINsgroupforwardingmemberReferenceDeleteRequest` struct via the builder pattern


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


## NsgroupforwardingmemberReferenceGet

> GetNsgroupForwardingmemberResponse NsgroupforwardingmemberReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific nsgroup:forwardingmember object



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
	reference := "reference_example" // string | Reference of the nsgroup:forwardingmember object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardingmemberAPI.NsgroupforwardingmemberReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.NsgroupforwardingmemberReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupforwardingmemberReferenceGet`: GetNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardingmemberAPI.NsgroupforwardingmemberReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardingmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPINsgroupforwardingmemberReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetNsgroupForwardingmemberResponse**](GetNsgroupForwardingmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NsgroupforwardingmemberReferencePut

> UpdateNsgroupForwardingmemberResponse NsgroupforwardingmemberReferencePut(ctx, reference).NsgroupForwardingmember(nsgroupForwardingmember).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a nsgroup:forwardingmember object



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
	reference := "reference_example" // string | Reference of the nsgroup:forwardingmember object
	nsgroupForwardingmember := *dns.NewNsgroupForwardingmember() // NsgroupForwardingmember | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardingmemberAPI.NsgroupforwardingmemberReferencePut(context.Background(), reference).NsgroupForwardingmember(nsgroupForwardingmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.NsgroupforwardingmemberReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupforwardingmemberReferencePut`: UpdateNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardingmemberAPI.NsgroupforwardingmemberReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardingmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPINsgroupforwardingmemberReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupForwardingmember** | [**NsgroupForwardingmember**](NsgroupForwardingmember.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateNsgroupForwardingmemberResponse**](UpdateNsgroupForwardingmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

