# NsgroupDelegationAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NsgroupdelegationGet**](NsgroupDelegationAPI.md#NsgroupdelegationGet) | **Get** /nsgroup:delegation | Retrieve nsgroup:delegation objects
[**NsgroupdelegationPost**](NsgroupDelegationAPI.md#NsgroupdelegationPost) | **Post** /nsgroup:delegation | Create a nsgroup:delegation object
[**NsgroupdelegationReferenceDelete**](NsgroupDelegationAPI.md#NsgroupdelegationReferenceDelete) | **Delete** /nsgroup:delegation/{reference} | Delete a nsgroup:delegation object
[**NsgroupdelegationReferenceGet**](NsgroupDelegationAPI.md#NsgroupdelegationReferenceGet) | **Get** /nsgroup:delegation/{reference} | Get a specific nsgroup:delegation object
[**NsgroupdelegationReferencePut**](NsgroupDelegationAPI.md#NsgroupdelegationReferencePut) | **Put** /nsgroup:delegation/{reference} | Update a nsgroup:delegation object



## NsgroupdelegationGet

> ListNsgroupDelegationResponse NsgroupdelegationGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve nsgroup:delegation objects



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
	resp, r, err := apiClient.NsgroupDelegationAPI.NsgroupdelegationGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupDelegationAPI.NsgroupdelegationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupdelegationGet`: ListNsgroupDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupDelegationAPI.NsgroupdelegationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupDelegationAPINsgroupdelegationGetRequest` struct via the builder pattern


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

[**ListNsgroupDelegationResponse**](ListNsgroupDelegationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NsgroupdelegationPost

> CreateNsgroupDelegationResponse NsgroupdelegationPost(ctx).NsgroupDelegation(nsgroupDelegation).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a nsgroup:delegation object



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
	nsgroupDelegation := *dns.NewNsgroupDelegation() // NsgroupDelegation | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupDelegationAPI.NsgroupdelegationPost(context.Background()).NsgroupDelegation(nsgroupDelegation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupDelegationAPI.NsgroupdelegationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupdelegationPost`: CreateNsgroupDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupDelegationAPI.NsgroupdelegationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupDelegationAPINsgroupdelegationPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupDelegation** | [**NsgroupDelegation**](NsgroupDelegation.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateNsgroupDelegationResponse**](CreateNsgroupDelegationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NsgroupdelegationReferenceDelete

> NsgroupdelegationReferenceDelete(ctx, reference).Execute()

Delete a nsgroup:delegation object



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
	reference := "reference_example" // string | Reference of the nsgroup:delegation object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.NsgroupDelegationAPI.NsgroupdelegationReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupDelegationAPI.NsgroupdelegationReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:delegation object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupDelegationAPINsgroupdelegationReferenceDeleteRequest` struct via the builder pattern


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


## NsgroupdelegationReferenceGet

> GetNsgroupDelegationResponse NsgroupdelegationReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific nsgroup:delegation object



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
	reference := "reference_example" // string | Reference of the nsgroup:delegation object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupDelegationAPI.NsgroupdelegationReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupDelegationAPI.NsgroupdelegationReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupdelegationReferenceGet`: GetNsgroupDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupDelegationAPI.NsgroupdelegationReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:delegation object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupDelegationAPINsgroupdelegationReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetNsgroupDelegationResponse**](GetNsgroupDelegationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NsgroupdelegationReferencePut

> UpdateNsgroupDelegationResponse NsgroupdelegationReferencePut(ctx, reference).NsgroupDelegation(nsgroupDelegation).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a nsgroup:delegation object



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
	reference := "reference_example" // string | Reference of the nsgroup:delegation object
	nsgroupDelegation := *dns.NewNsgroupDelegation() // NsgroupDelegation | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupDelegationAPI.NsgroupdelegationReferencePut(context.Background(), reference).NsgroupDelegation(nsgroupDelegation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupDelegationAPI.NsgroupdelegationReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupdelegationReferencePut`: UpdateNsgroupDelegationResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupDelegationAPI.NsgroupdelegationReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:delegation object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupDelegationAPINsgroupdelegationReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupDelegation** | [**NsgroupDelegation**](NsgroupDelegation.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateNsgroupDelegationResponse**](UpdateNsgroupDelegationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

