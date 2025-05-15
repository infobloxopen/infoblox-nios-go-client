# NsgroupforwardingmemberAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](NsgroupforwardingmemberAPI.md#Get) | **Get** /nsgroup:forwardingmember | Retrieve nsgroup:forwardingmember objects
[**Post**](NsgroupforwardingmemberAPI.md#Post) | **Post** /nsgroup:forwardingmember | Create a nsgroup:forwardingmember object
[**ReferenceDelete**](NsgroupforwardingmemberAPI.md#ReferenceDelete) | **Delete** /nsgroup:forwardingmember/{reference} | Delete a nsgroup:forwardingmember object
[**ReferenceGet**](NsgroupforwardingmemberAPI.md#ReferenceGet) | **Get** /nsgroup:forwardingmember/{reference} | Get a specific nsgroup:forwardingmember object
[**ReferencePut**](NsgroupforwardingmemberAPI.md#ReferencePut) | **Put** /nsgroup:forwardingmember/{reference} | Update a nsgroup:forwardingmember object



## Get

> ListNsgroupForwardingmemberResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

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
	resp, r, err := apiClient.NsgroupforwardingmemberAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardingmemberAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupforwardingmemberAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupforwardingmemberAPIGetRequest` struct via the builder pattern


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


## Post

> CreateNsgroupForwardingmemberResponse Post(ctx).NsgroupForwardingmember(nsgroupForwardingmember).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.NsgroupforwardingmemberAPI.Post(context.Background()).NsgroupForwardingmember(nsgroupForwardingmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardingmemberAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupforwardingmemberAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupforwardingmemberAPIPostRequest` struct via the builder pattern


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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

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
	r, err := apiClient.NsgroupforwardingmemberAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardingmemberAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `NsgroupforwardingmemberAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetNsgroupForwardingmemberResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.NsgroupforwardingmemberAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardingmemberAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupforwardingmemberAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardingmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupforwardingmemberAPIReferenceGetRequest` struct via the builder pattern


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


## ReferencePut

> UpdateNsgroupForwardingmemberResponse ReferencePut(ctx, reference).NsgroupForwardingmember(nsgroupForwardingmember).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.NsgroupforwardingmemberAPI.ReferencePut(context.Background(), reference).NsgroupForwardingmember(nsgroupForwardingmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupforwardingmemberAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupforwardingmemberAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardingmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupforwardingmemberAPIReferencePutRequest` struct via the builder pattern


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

