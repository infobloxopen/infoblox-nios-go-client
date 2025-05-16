# Ipv6dhcpoptionspaceAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](Ipv6dhcpoptionspaceAPI.md#Get) | **Get** /ipv6dhcpoptionspace | Retrieve ipv6dhcpoptionspace objects
[**Post**](Ipv6dhcpoptionspaceAPI.md#Post) | **Post** /ipv6dhcpoptionspace | Create a ipv6dhcpoptionspace object
[**ReferenceDelete**](Ipv6dhcpoptionspaceAPI.md#ReferenceDelete) | **Delete** /ipv6dhcpoptionspace/{reference} | Delete a ipv6dhcpoptionspace object
[**ReferenceGet**](Ipv6dhcpoptionspaceAPI.md#ReferenceGet) | **Get** /ipv6dhcpoptionspace/{reference} | Get a specific ipv6dhcpoptionspace object
[**ReferencePut**](Ipv6dhcpoptionspaceAPI.md#ReferencePut) | **Put** /ipv6dhcpoptionspace/{reference} | Update a ipv6dhcpoptionspace object



## Get

> ListIpv6dhcpoptionspaceResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve ipv6dhcpoptionspace objects



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
	resp, r, err := apiClient.Ipv6dhcpoptionspaceAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListIpv6dhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptionspaceAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPIGetRequest` struct via the builder pattern


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

[**ListIpv6dhcpoptionspaceResponse**](ListIpv6dhcpoptionspaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateIpv6dhcpoptionspaceResponse Post(ctx).Ipv6dhcpoptionspace(ipv6dhcpoptionspace).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a ipv6dhcpoptionspace object



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
	ipv6dhcpoptionspace := *dhcp.NewIpv6dhcpoptionspace() // Ipv6dhcpoptionspace | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptionspaceAPI.Post(context.Background()).Ipv6dhcpoptionspace(ipv6dhcpoptionspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateIpv6dhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptionspaceAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6dhcpoptionspace** | [**Ipv6dhcpoptionspace**](Ipv6dhcpoptionspace.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateIpv6dhcpoptionspaceResponse**](CreateIpv6dhcpoptionspaceResponse.md)

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

Delete a ipv6dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the ipv6dhcpoptionspace object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.Ipv6dhcpoptionspaceAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetIpv6dhcpoptionspaceResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific ipv6dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the ipv6dhcpoptionspace object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptionspaceAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetIpv6dhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptionspaceAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetIpv6dhcpoptionspaceResponse**](GetIpv6dhcpoptionspaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateIpv6dhcpoptionspaceResponse ReferencePut(ctx, reference).Ipv6dhcpoptionspace(ipv6dhcpoptionspace).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a ipv6dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the ipv6dhcpoptionspace object
	ipv6dhcpoptionspace := *dhcp.NewIpv6dhcpoptionspace() // Ipv6dhcpoptionspace | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptionspaceAPI.ReferencePut(context.Background(), reference).Ipv6dhcpoptionspace(ipv6dhcpoptionspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateIpv6dhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptionspaceAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6dhcpoptionspace** | [**Ipv6dhcpoptionspace**](Ipv6dhcpoptionspace.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateIpv6dhcpoptionspaceResponse**](UpdateIpv6dhcpoptionspaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

