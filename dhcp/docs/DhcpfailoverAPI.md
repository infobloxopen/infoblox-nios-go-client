# DhcpfailoverAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](DhcpfailoverAPI.md#Get) | **Get** /dhcpfailover | Retrieve dhcpfailover objects
[**Post**](DhcpfailoverAPI.md#Post) | **Post** /dhcpfailover | Create a dhcpfailover object
[**ReferenceDelete**](DhcpfailoverAPI.md#ReferenceDelete) | **Delete** /dhcpfailover/{reference} | Delete a dhcpfailover object
[**ReferenceGet**](DhcpfailoverAPI.md#ReferenceGet) | **Get** /dhcpfailover/{reference} | Get a specific dhcpfailover object
[**ReferencePut**](DhcpfailoverAPI.md#ReferencePut) | **Put** /dhcpfailover/{reference} | Update a dhcpfailover object



## Get

> ListDhcpfailoverResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dhcpfailover objects



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
	resp, r, err := apiClient.DhcpfailoverAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpfailoverAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListDhcpfailoverResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpfailoverAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpfailoverAPIGetRequest` struct via the builder pattern


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

[**ListDhcpfailoverResponse**](ListDhcpfailoverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateDhcpfailoverResponse Post(ctx).Dhcpfailover(dhcpfailover).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dhcpfailover object



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
	dhcpfailover := *dhcp.NewDhcpfailover() // Dhcpfailover | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpfailoverAPI.Post(context.Background()).Dhcpfailover(dhcpfailover).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpfailoverAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateDhcpfailoverResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpfailoverAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DhcpfailoverAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpfailover** | [**Dhcpfailover**](Dhcpfailover.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDhcpfailoverResponse**](CreateDhcpfailoverResponse.md)

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

Delete a dhcpfailover object



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
	reference := "reference_example" // string | Reference of the dhcpfailover object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.DhcpfailoverAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpfailoverAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpfailover object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpfailoverAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetDhcpfailoverResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dhcpfailover object



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
	reference := "reference_example" // string | Reference of the dhcpfailover object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpfailoverAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpfailoverAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetDhcpfailoverResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpfailoverAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpfailover object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpfailoverAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDhcpfailoverResponse**](GetDhcpfailoverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateDhcpfailoverResponse ReferencePut(ctx, reference).Dhcpfailover(dhcpfailover).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dhcpfailover object



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
	reference := "reference_example" // string | Reference of the dhcpfailover object
	dhcpfailover := *dhcp.NewDhcpfailover() // Dhcpfailover | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.DhcpfailoverAPI.ReferencePut(context.Background(), reference).Dhcpfailover(dhcpfailover).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DhcpfailoverAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateDhcpfailoverResponse
	fmt.Fprintf(os.Stdout, "Response from `DhcpfailoverAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dhcpfailover object | 

### Other Parameters

Other parameters are passed through a pointer to a `DhcpfailoverAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dhcpfailover** | [**Dhcpfailover**](Dhcpfailover.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDhcpfailoverResponse**](UpdateDhcpfailoverResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

