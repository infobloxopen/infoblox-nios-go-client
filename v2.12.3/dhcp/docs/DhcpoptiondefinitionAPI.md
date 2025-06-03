# DhcpoptiondefinitionAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](DhcpoptiondefinitionAPI.md#Get) | **Get** /dhcpoptiondefinition | Retrieve dhcpoptiondefinition objects
[**Post**](DhcpoptiondefinitionAPI.md#Post) | **Post** /dhcpoptiondefinition | Create a dhcpoptiondefinition object
[**ReferenceDelete**](DhcpoptiondefinitionAPI.md#ReferenceDelete) | **Delete** /dhcpoptiondefinition/{reference} | Delete a dhcpoptiondefinition object
[**ReferenceGet**](DhcpoptiondefinitionAPI.md#ReferenceGet) | **Get** /dhcpoptiondefinition/{reference} | Get a specific dhcpoptiondefinition object
[**ReferencePut**](DhcpoptiondefinitionAPI.md#ReferencePut) | **Put** /dhcpoptiondefinition/{reference} | Update a dhcpoptiondefinition object



## Get

> ListDhcpoptiondefinitionResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dhcpoptiondefinition objects



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
	resp, r, err := apiClient.DhcpoptiondefinitionAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListDhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptiondefinitionAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPIGetRequest` struct via the builder pattern


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

[**ListDhcpoptiondefinitionResponse**](ListDhcpoptiondefinitionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateDhcpoptiondefinitionResponse Post(ctx).Dhcpoptiondefinition(dhcpoptiondefinition).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dhcpoptiondefinition object



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
	dhcpoptiondefinition := *dhcp.NewDhcpoptiondefinition() // Dhcpoptiondefinition | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptiondefinitionAPI.Post(context.Background()).Dhcpoptiondefinition(dhcpoptiondefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateDhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptiondefinitionAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpoptiondefinition** | [**Dhcpoptiondefinition**](Dhcpoptiondefinition.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

Delete a dhcpoptiondefinition object



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
	reference := "reference_example" // string | Reference of the dhcpoptiondefinition object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.DhcpoptiondefinitionAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetDhcpoptiondefinitionResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dhcpoptiondefinition object



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
	reference := "reference_example" // string | Reference of the dhcpoptiondefinition object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptiondefinitionAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetDhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptiondefinitionAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## ReferencePut

> UpdateDhcpoptiondefinitionResponse ReferencePut(ctx, reference).Dhcpoptiondefinition(dhcpoptiondefinition).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dhcpoptiondefinition object



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
	reference := "reference_example" // string | Reference of the dhcpoptiondefinition object
	dhcpoptiondefinition := *dhcp.NewDhcpoptiondefinition() // Dhcpoptiondefinition | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpoptiondefinitionAPI.ReferencePut(context.Background(), reference).Dhcpoptiondefinition(dhcpoptiondefinition).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpoptiondefinitionAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateDhcpoptiondefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpoptiondefinitionAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpoptiondefinition object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpoptiondefinitionAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpoptiondefinition** | [**Dhcpoptiondefinition**](Dhcpoptiondefinition.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

