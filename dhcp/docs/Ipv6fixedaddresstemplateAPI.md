# Ipv6fixedaddresstemplateAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Ipv6fixedaddresstemplateAPI.md#Create) | **Post** /ipv6fixedaddresstemplate | Create a ipv6fixedaddresstemplate object
[**Delete**](Ipv6fixedaddresstemplateAPI.md#Delete) | **Delete** /ipv6fixedaddresstemplate/{reference} | Delete a ipv6fixedaddresstemplate object
[**List**](Ipv6fixedaddresstemplateAPI.md#List) | **Get** /ipv6fixedaddresstemplate | Retrieve ipv6fixedaddresstemplate objects
[**Read**](Ipv6fixedaddresstemplateAPI.md#Read) | **Get** /ipv6fixedaddresstemplate/{reference} | Get a specific ipv6fixedaddresstemplate object
[**Update**](Ipv6fixedaddresstemplateAPI.md#Update) | **Put** /ipv6fixedaddresstemplate/{reference} | Update a ipv6fixedaddresstemplate object



## Create

> CreateIpv6fixedaddresstemplateResponse Create(ctx).Ipv6fixedaddresstemplate(ipv6fixedaddresstemplate).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a ipv6fixedaddresstemplate object



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
	ipv6fixedaddresstemplate := *dhcp.NewIpv6fixedaddresstemplate() // Ipv6fixedaddresstemplate | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6fixedaddresstemplateAPI.Create(context.Background()).Ipv6fixedaddresstemplate(ipv6fixedaddresstemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateIpv6fixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddresstemplateAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6fixedaddresstemplate** | [**Ipv6fixedaddresstemplate**](Ipv6fixedaddresstemplate.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateIpv6fixedaddresstemplateResponse**](CreateIpv6fixedaddresstemplateResponse.md)

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

Delete a ipv6fixedaddresstemplate object



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
	reference := "reference_example" // string | Reference of the ipv6fixedaddresstemplate object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.Ipv6fixedaddresstemplateAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6fixedaddresstemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPIDeleteRequest` struct via the builder pattern


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

> ListIpv6fixedaddresstemplateResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve ipv6fixedaddresstemplate objects



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
	resp, r, err := apiClient.Ipv6fixedaddresstemplateAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListIpv6fixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddresstemplateAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPIListRequest` struct via the builder pattern


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

[**ListIpv6fixedaddresstemplateResponse**](ListIpv6fixedaddresstemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetIpv6fixedaddresstemplateResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific ipv6fixedaddresstemplate object



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
	reference := "reference_example" // string | Reference of the ipv6fixedaddresstemplate object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6fixedaddresstemplateAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetIpv6fixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddresstemplateAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6fixedaddresstemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetIpv6fixedaddresstemplateResponse**](GetIpv6fixedaddresstemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateIpv6fixedaddresstemplateResponse Update(ctx, reference).Ipv6fixedaddresstemplate(ipv6fixedaddresstemplate).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a ipv6fixedaddresstemplate object



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
	reference := "reference_example" // string | Reference of the ipv6fixedaddresstemplate object
	ipv6fixedaddresstemplate := *dhcp.NewIpv6fixedaddresstemplate() // Ipv6fixedaddresstemplate | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6fixedaddresstemplateAPI.Update(context.Background(), reference).Ipv6fixedaddresstemplate(ipv6fixedaddresstemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateIpv6fixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddresstemplateAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6fixedaddresstemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6fixedaddresstemplate** | [**Ipv6fixedaddresstemplate**](Ipv6fixedaddresstemplate.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateIpv6fixedaddresstemplateResponse**](UpdateIpv6fixedaddresstemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

