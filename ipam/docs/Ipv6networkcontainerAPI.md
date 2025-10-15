# Ipv6networkcontainerAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Ipv6networkcontainerAPI.md#Create) | **Post** /ipv6networkcontainer | Create a ipv6networkcontainer object
[**Delete**](Ipv6networkcontainerAPI.md#Delete) | **Delete** /ipv6networkcontainer/{reference} | Delete a ipv6networkcontainer object
[**List**](Ipv6networkcontainerAPI.md#List) | **Get** /ipv6networkcontainer | Retrieve ipv6networkcontainer objects
[**Read**](Ipv6networkcontainerAPI.md#Read) | **Get** /ipv6networkcontainer/{reference} | Get a specific ipv6networkcontainer object
[**Update**](Ipv6networkcontainerAPI.md#Update) | **Put** /ipv6networkcontainer/{reference} | Update a ipv6networkcontainer object



## Create

> CreateIpv6networkcontainerResponse Create(ctx).Ipv6networkcontainer(ipv6networkcontainer).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a ipv6networkcontainer object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {
	ipv6networkcontainer := *ipam.NewIpv6networkcontainer() // Ipv6networkcontainer | Object data to create

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.Ipv6networkcontainerAPI.Create(context.Background()).Ipv6networkcontainer(ipv6networkcontainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateIpv6networkcontainerResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6networkcontainerAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6networkcontainer** | [**Ipv6networkcontainer**](Ipv6networkcontainer.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateIpv6networkcontainerResponse**](CreateIpv6networkcontainerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, reference).RemoveSubnets(removeSubnets).Execute()

Delete a ipv6networkcontainer object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6networkcontainer object

	apiClient := ipam.NewAPIClient()
	r, err := apiClient.Ipv6networkcontainerAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6networkcontainer object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPIDeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeSubnets** | **bool** | Remove subnets delete option. Determines whether all child objects should be removed alongside with the IPv6 network container or child objects should be assigned to another parental container. By default child objects are deleted with this network container. | 

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

> ListIpv6networkcontainerResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve ipv6networkcontainer objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.Ipv6networkcontainerAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListIpv6networkcontainerResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6networkcontainerAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPIListRequest` struct via the builder pattern


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

[**ListIpv6networkcontainerResponse**](ListIpv6networkcontainerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetIpv6networkcontainerResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific ipv6networkcontainer object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6networkcontainer object

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.Ipv6networkcontainerAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetIpv6networkcontainerResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6networkcontainerAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6networkcontainer object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetIpv6networkcontainerResponse**](GetIpv6networkcontainerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateIpv6networkcontainerResponse Update(ctx, reference).Ipv6networkcontainer(ipv6networkcontainer).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a ipv6networkcontainer object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6networkcontainer object
	ipv6networkcontainer := *ipam.NewIpv6networkcontainer() // Ipv6networkcontainer | Object data to update

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.Ipv6networkcontainerAPI.Update(context.Background(), reference).Ipv6networkcontainer(ipv6networkcontainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateIpv6networkcontainerResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6networkcontainerAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6networkcontainer object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6networkcontainer** | [**Ipv6networkcontainer**](Ipv6networkcontainer.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateIpv6networkcontainerResponse**](UpdateIpv6networkcontainerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

