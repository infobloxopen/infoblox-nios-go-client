# DhcpoptionspaceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](DhcpoptionspaceAPI.md#Get) | **Get** /dhcpoptionspace | Retrieve dhcpoptionspace objects
[**Post**](DhcpoptionspaceAPI.md#Post) | **Post** /dhcpoptionspace | Create a dhcpoptionspace object
[**ReferenceDelete**](DhcpoptionspaceAPI.md#ReferenceDelete) | **Delete** /dhcpoptionspace/{reference} | Delete a dhcpoptionspace object
[**ReferenceGet**](DhcpoptionspaceAPI.md#ReferenceGet) | **Get** /dhcpoptionspace/{reference} | Get a specific dhcpoptionspace object
[**ReferencePut**](DhcpoptionspaceAPI.md#ReferencePut) | **Put** /dhcpoptionspace/{reference} | Update a dhcpoptionspace object



## Get

> ListDhcpoptionspaceResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dhcpoptionspace objects



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
	resp, r, err := apiClient.DhcpoptionspaceAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListDhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptionspaceAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptionspaceAPIGetRequest` struct via the builder pattern


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

[**ListDhcpoptionspaceResponse**](ListDhcpoptionspaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateDhcpoptionspaceResponse Post(ctx).Dhcpoptionspace(dhcpoptionspace).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dhcpoptionspace object



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
	dhcpoptionspace := *dhcp.NewDhcpoptionspace() // Dhcpoptionspace | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptionspaceAPI.Post(context.Background()).Dhcpoptionspace(dhcpoptionspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateDhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptionspaceAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptionspaceAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpoptionspace** | [**Dhcpoptionspace**](Dhcpoptionspace.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

Delete a dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the dhcpoptionspace object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.DhcpoptionspaceAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `DhcpoptionspaceAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetDhcpoptionspaceResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the dhcpoptionspace object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptionspaceAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetDhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptionspaceAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptionspaceAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## ReferencePut

> UpdateDhcpoptionspaceResponse ReferencePut(ctx, reference).Dhcpoptionspace(dhcpoptionspace).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dhcpoptionspace object



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
	reference := "reference_example" // string | Reference of the dhcpoptionspace object
	dhcpoptionspace := *dhcp.NewDhcpoptionspace() // Dhcpoptionspace | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptionspaceAPI.ReferencePut(context.Background(), reference).Dhcpoptionspace(dhcpoptionspace).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptionspaceAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateDhcpoptionspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptionspaceAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptionspace object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptionspaceAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpoptionspace** | [**Dhcpoptionspace**](Dhcpoptionspace.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

