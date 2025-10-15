# DhcpoptionspaceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DhcpoptionspaceAPI.md#Create) | **Post** /dhcpoptionspace | Create a dhcpoptionspace object
[**Delete**](DhcpoptionspaceAPI.md#Delete) | **Delete** /dhcpoptionspace/{reference} | Delete a dhcpoptionspace object
[**List**](DhcpoptionspaceAPI.md#List) | **Get** /dhcpoptionspace | Retrieve dhcpoptionspace objects
[**Read**](DhcpoptionspaceAPI.md#Read) | **Get** /dhcpoptionspace/{reference} | Get a specific dhcpoptionspace object
[**Update**](DhcpoptionspaceAPI.md#Update) | **Put** /dhcpoptionspace/{reference} | Update a dhcpoptionspace object



## Create

> CreateDhcpoptionspaceResponse Create(ctx).Dhcpoptionspace(dhcpoptionspace).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a dhcpoptionspace object



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
	dhcpoptionspace := *dhcp.NewDhcpoptionspace() // Dhcpoptionspace | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptionspaceAPI.Create(context.Background()).Dhcpoptionspace(dhcpoptionspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateDhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptionspaceAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptionspaceAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpoptionspace** | [**Dhcpoptionspace**](Dhcpoptionspace.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDhcpoptionspaceResponse**](CreateDhcpoptionspaceResponse.md)

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

Delete a dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the dhcpoptionspace object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.DhcpoptionspaceAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptionspaceAPIDeleteRequest` struct via the builder pattern


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

> ListDhcpoptionspaceResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve dhcpoptionspace objects



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
	resp, r, err := apiClient.DhcpoptionspaceAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListDhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptionspaceAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptionspaceAPIListRequest` struct via the builder pattern


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

[**ListDhcpoptionspaceResponse**](ListDhcpoptionspaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetDhcpoptionspaceResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the dhcpoptionspace object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptionspaceAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetDhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptionspaceAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptionspaceAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetDhcpoptionspaceResponse**](GetDhcpoptionspaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateDhcpoptionspaceResponse Update(ctx, reference).Dhcpoptionspace(dhcpoptionspace).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the dhcpoptionspace object
	dhcpoptionspace := *dhcp.NewDhcpoptionspace() // Dhcpoptionspace | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptionspaceAPI.Update(context.Background(), reference).Dhcpoptionspace(dhcpoptionspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateDhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptionspaceAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptionspaceAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpoptionspace** | [**Dhcpoptionspace**](Dhcpoptionspace.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDhcpoptionspaceResponse**](UpdateDhcpoptionspaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

