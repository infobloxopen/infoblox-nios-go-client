# Ipv6dhcpoptiondefinitionAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Ipv6dhcpoptiondefinitionAPI.md#Create) | **Post** /ipv6dhcpoptiondefinition | Create a ipv6dhcpoptiondefinition object
[**Delete**](Ipv6dhcpoptiondefinitionAPI.md#Delete) | **Delete** /ipv6dhcpoptiondefinition/{reference} | Delete a ipv6dhcpoptiondefinition object
[**List**](Ipv6dhcpoptiondefinitionAPI.md#List) | **Get** /ipv6dhcpoptiondefinition | Retrieve ipv6dhcpoptiondefinition objects
[**Read**](Ipv6dhcpoptiondefinitionAPI.md#Read) | **Get** /ipv6dhcpoptiondefinition/{reference} | Get a specific ipv6dhcpoptiondefinition object
[**Update**](Ipv6dhcpoptiondefinitionAPI.md#Update) | **Put** /ipv6dhcpoptiondefinition/{reference} | Update a ipv6dhcpoptiondefinition object



## Create

> CreateIpv6dhcpoptiondefinitionResponse Create(ctx).Ipv6dhcpoptiondefinition(ipv6dhcpoptiondefinition).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a ipv6dhcpoptiondefinition object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {
	ipv6dhcpoptiondefinition := *dhcp.NewIpv6dhcpoptiondefinition() // Ipv6dhcpoptiondefinition | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.Create(context.Background()).Ipv6dhcpoptiondefinition(ipv6dhcpoptiondefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateIpv6dhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptiondefinitionAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6dhcpoptiondefinition** | [**Ipv6dhcpoptiondefinition**](Ipv6dhcpoptiondefinition.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## Delete

> Delete(ctx, reference).Execute()

Delete a ipv6dhcpoptiondefinition object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6dhcpoptiondefinition object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPIDeleteRequest` struct via the builder pattern


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

> ListIpv6dhcpoptiondefinitionResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve ipv6dhcpoptiondefinition objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListIpv6dhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptiondefinitionAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPIListRequest` struct via the builder pattern


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

[**ListIpv6dhcpoptiondefinitionResponse**](ListIpv6dhcpoptiondefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetIpv6dhcpoptiondefinitionResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific ipv6dhcpoptiondefinition object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6dhcpoptiondefinition object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetIpv6dhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptiondefinitionAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

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


## Update

> UpdateIpv6dhcpoptiondefinitionResponse Update(ctx, reference).Ipv6dhcpoptiondefinition(ipv6dhcpoptiondefinition).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a ipv6dhcpoptiondefinition object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6dhcpoptiondefinition object
	ipv6dhcpoptiondefinition := *dhcp.NewIpv6dhcpoptiondefinition() // Ipv6dhcpoptiondefinition | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptiondefinitionAPI.Update(context.Background(), reference).Ipv6dhcpoptiondefinition(ipv6dhcpoptiondefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptiondefinitionAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateIpv6dhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptiondefinitionAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptiondefinitionAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6dhcpoptiondefinition** | [**Ipv6dhcpoptiondefinition**](Ipv6dhcpoptiondefinition.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

