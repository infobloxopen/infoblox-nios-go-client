# Ipv6dhcpoptiondefinitionAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](Ipv6dhcpoptiondefinitionAPI.md#Get) | **Get** /ipv6dhcpoptiondefinition | Retrieve ipv6dhcpoptiondefinition objects
[**Post**](Ipv6dhcpoptiondefinitionAPI.md#Post) | **Post** /ipv6dhcpoptiondefinition | Create a ipv6dhcpoptiondefinition object
[**ReferenceDelete**](Ipv6dhcpoptiondefinitionAPI.md#ReferenceDelete) | **Delete** /ipv6dhcpoptiondefinition/{reference} | Delete a ipv6dhcpoptiondefinition object
[**ReferenceGet**](Ipv6dhcpoptiondefinitionAPI.md#ReferenceGet) | **Get** /ipv6dhcpoptiondefinition/{reference} | Get a specific ipv6dhcpoptiondefinition object
[**ReferencePut**](Ipv6dhcpoptiondefinitionAPI.md#ReferencePut) | **Put** /ipv6dhcpoptiondefinition/{reference} | Update a ipv6dhcpoptiondefinition object



## Get

> ListIpv6dhcpoptiondefinitionResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve ipv6dhcpoptiondefinition objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListIpv6dhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptiondefinitionAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPIGetRequest` struct via the builder pattern


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

[**ListIpv6dhcpoptiondefinitionResponse**](ListIpv6dhcpoptiondefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateIpv6dhcpoptiondefinitionResponse Post(ctx).Ipv6dhcpoptiondefinition(ipv6dhcpoptiondefinition).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a ipv6dhcpoptiondefinition object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {
	ipv6dhcpoptiondefinition := *dhcp.NewIpv6dhcpoptiondefinition() // Ipv6dhcpoptiondefinition | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.Post(context.Background()).Ipv6dhcpoptiondefinition(ipv6dhcpoptiondefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateIpv6dhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptiondefinitionAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6dhcpoptiondefinition** | [**Ipv6dhcpoptiondefinition**](Ipv6dhcpoptiondefinition.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateIpv6dhcpoptiondefinitionResponse**](CreateIpv6dhcpoptiondefinitionResponse.md)

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

Delete a ipv6dhcpoptiondefinition object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6dhcpoptiondefinition object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetIpv6dhcpoptiondefinitionResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific ipv6dhcpoptiondefinition object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6dhcpoptiondefinition object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetIpv6dhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptiondefinitionAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetIpv6dhcpoptiondefinitionResponse**](GetIpv6dhcpoptiondefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateIpv6dhcpoptiondefinitionResponse ReferencePut(ctx, reference).Ipv6dhcpoptiondefinition(ipv6dhcpoptiondefinition).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a ipv6dhcpoptiondefinition object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6dhcpoptiondefinition object
	ipv6dhcpoptiondefinition := *dhcp.NewIpv6dhcpoptiondefinition() // Ipv6dhcpoptiondefinition | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.ReferencePut(context.Background(), reference).Ipv6dhcpoptiondefinition(ipv6dhcpoptiondefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateIpv6dhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptiondefinitionAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6dhcpoptiondefinition** | [**Ipv6dhcpoptiondefinition**](Ipv6dhcpoptiondefinition.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateIpv6dhcpoptiondefinitionResponse**](UpdateIpv6dhcpoptiondefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

