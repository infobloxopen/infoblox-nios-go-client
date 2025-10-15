# Ipv6dhcpoptionspaceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Ipv6dhcpoptionspaceAPI.md#Create) | **Post** /ipv6dhcpoptionspace | Create a ipv6dhcpoptionspace object
[**Delete**](Ipv6dhcpoptionspaceAPI.md#Delete) | **Delete** /ipv6dhcpoptionspace/{reference} | Delete a ipv6dhcpoptionspace object
[**List**](Ipv6dhcpoptionspaceAPI.md#List) | **Get** /ipv6dhcpoptionspace | Retrieve ipv6dhcpoptionspace objects
[**Read**](Ipv6dhcpoptionspaceAPI.md#Read) | **Get** /ipv6dhcpoptionspace/{reference} | Get a specific ipv6dhcpoptionspace object
[**Update**](Ipv6dhcpoptionspaceAPI.md#Update) | **Put** /ipv6dhcpoptionspace/{reference} | Update a ipv6dhcpoptionspace object



## Create

> CreateIpv6dhcpoptionspaceResponse Create(ctx).Ipv6dhcpoptionspace(ipv6dhcpoptionspace).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a ipv6dhcpoptionspace object



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
	ipv6dhcpoptionspace := *dhcp.NewIpv6dhcpoptionspace() // Ipv6dhcpoptionspace | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptionspaceAPI.Create(context.Background()).Ipv6dhcpoptionspace(ipv6dhcpoptionspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateIpv6dhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptionspaceAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6dhcpoptionspace** | [**Ipv6dhcpoptionspace**](Ipv6dhcpoptionspace.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## Delete

> Delete(ctx, reference).Execute()

Delete a ipv6dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the ipv6dhcpoptionspace object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.Ipv6dhcpoptionspaceAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPIDeleteRequest` struct via the builder pattern


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

> ListIpv6dhcpoptionspaceResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve ipv6dhcpoptionspace objects



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
	resp, r, err := apiClient.Ipv6dhcpoptionspaceAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListIpv6dhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptionspaceAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPIListRequest` struct via the builder pattern


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

[**ListIpv6dhcpoptionspaceResponse**](ListIpv6dhcpoptionspaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetIpv6dhcpoptionspaceResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific ipv6dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the ipv6dhcpoptionspace object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptionspaceAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetIpv6dhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptionspaceAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

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


## Update

> UpdateIpv6dhcpoptionspaceResponse Update(ctx, reference).Ipv6dhcpoptionspace(ipv6dhcpoptionspace).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a ipv6dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the ipv6dhcpoptionspace object
	ipv6dhcpoptionspace := *dhcp.NewIpv6dhcpoptionspace() // Ipv6dhcpoptionspace | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6dhcpoptionspaceAPI.Update(context.Background(), reference).Ipv6dhcpoptionspace(ipv6dhcpoptionspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6dhcpoptionspaceAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateIpv6dhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6dhcpoptionspaceAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6dhcpoptionspaceAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6dhcpoptionspace** | [**Ipv6dhcpoptionspace**](Ipv6dhcpoptionspace.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

