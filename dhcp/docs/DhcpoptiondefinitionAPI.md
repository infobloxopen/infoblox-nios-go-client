# DhcpoptiondefinitionAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DhcpoptiondefinitionAPI.md#Create) | **Post** /dhcpoptiondefinition | Create a dhcpoptiondefinition object
[**Delete**](DhcpoptiondefinitionAPI.md#Delete) | **Delete** /dhcpoptiondefinition/{reference} | Delete a dhcpoptiondefinition object
[**List**](DhcpoptiondefinitionAPI.md#List) | **Get** /dhcpoptiondefinition | Retrieve dhcpoptiondefinition objects
[**Read**](DhcpoptiondefinitionAPI.md#Read) | **Get** /dhcpoptiondefinition/{reference} | Get a specific dhcpoptiondefinition object
[**Update**](DhcpoptiondefinitionAPI.md#Update) | **Put** /dhcpoptiondefinition/{reference} | Update a dhcpoptiondefinition object



## Create

> CreateDhcpoptiondefinitionResponse Create(ctx).Dhcpoptiondefinition(dhcpoptiondefinition).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a dhcpoptiondefinition object



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
	dhcpoptiondefinition := *dhcp.NewDhcpoptiondefinition() // Dhcpoptiondefinition | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptiondefinitionAPI.Create(context.Background()).Dhcpoptiondefinition(dhcpoptiondefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateDhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptiondefinitionAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpoptiondefinition** | [**Dhcpoptiondefinition**](Dhcpoptiondefinition.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDhcpoptiondefinitionResponse**](CreateDhcpoptiondefinitionResponse.md)

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

Delete a dhcpoptiondefinition object



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
	reference := "reference_example" // string | Reference of the dhcpoptiondefinition object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.DhcpoptiondefinitionAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPIDeleteRequest` struct via the builder pattern


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

> ListDhcpoptiondefinitionResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve dhcpoptiondefinition objects



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
	resp, r, err := apiClient.DhcpoptiondefinitionAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListDhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptiondefinitionAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPIListRequest` struct via the builder pattern


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

[**ListDhcpoptiondefinitionResponse**](ListDhcpoptiondefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetDhcpoptiondefinitionResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific dhcpoptiondefinition object



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
	reference := "reference_example" // string | Reference of the dhcpoptiondefinition object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptiondefinitionAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetDhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptiondefinitionAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetDhcpoptiondefinitionResponse**](GetDhcpoptiondefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateDhcpoptiondefinitionResponse Update(ctx, reference).Dhcpoptiondefinition(dhcpoptiondefinition).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a dhcpoptiondefinition object



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
	reference := "reference_example" // string | Reference of the dhcpoptiondefinition object
	dhcpoptiondefinition := *dhcp.NewDhcpoptiondefinition() // Dhcpoptiondefinition | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptiondefinitionAPI.Update(context.Background(), reference).Dhcpoptiondefinition(dhcpoptiondefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateDhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptiondefinitionAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpoptiondefinition** | [**Dhcpoptiondefinition**](Dhcpoptiondefinition.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDhcpoptiondefinitionResponse**](UpdateDhcpoptiondefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

