# Ipv6networkcontainerAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](Ipv6networkcontainerAPI.md#Get) | **Get** /ipv6networkcontainer | Retrieve ipv6networkcontainer objects
[**Post**](Ipv6networkcontainerAPI.md#Post) | **Post** /ipv6networkcontainer | Create a ipv6networkcontainer object
[**ReferenceDelete**](Ipv6networkcontainerAPI.md#ReferenceDelete) | **Delete** /ipv6networkcontainer/{reference} | Delete a ipv6networkcontainer object
[**ReferenceGet**](Ipv6networkcontainerAPI.md#ReferenceGet) | **Get** /ipv6networkcontainer/{reference} | Get a specific ipv6networkcontainer object
[**ReferencePut**](Ipv6networkcontainerAPI.md#ReferencePut) | **Put** /ipv6networkcontainer/{reference} | Update a ipv6networkcontainer object



## Get

> ListIpv6networkcontainerResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve ipv6networkcontainer objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
)

func main() {

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.Ipv6networkcontainerAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListIpv6networkcontainerResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6networkcontainerAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPIGetRequest` struct via the builder pattern


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

[**ListIpv6networkcontainerResponse**](ListIpv6networkcontainerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateIpv6networkcontainerResponse Post(ctx).Ipv6networkcontainer(ipv6networkcontainer).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a ipv6networkcontainer object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
)

func main() {
	ipv6networkcontainer := *ipam.NewIpv6networkcontainer() // Ipv6networkcontainer | Object data to create

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.Ipv6networkcontainerAPI.Post(context.Background()).Ipv6networkcontainer(ipv6networkcontainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateIpv6networkcontainerResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6networkcontainerAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6networkcontainer** | [**Ipv6networkcontainer**](Ipv6networkcontainer.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## ReferenceDelete

> ReferenceDelete(ctx, reference).RemoveSubnets(removeSubnets).Execute()

Delete a ipv6networkcontainer object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6networkcontainer object

	apiClient := ipam.NewAPIClient()
	r, err := apiClient.Ipv6networkcontainerAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPIReferenceDeleteRequest` struct via the builder pattern


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


## ReferenceGet

> GetIpv6networkcontainerResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific ipv6networkcontainer object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6networkcontainer object

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.Ipv6networkcontainerAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetIpv6networkcontainerResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6networkcontainerAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6networkcontainer object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## ReferencePut

> UpdateIpv6networkcontainerResponse ReferencePut(ctx, reference).Ipv6networkcontainer(ipv6networkcontainer).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a ipv6networkcontainer object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the ipv6networkcontainer object
	ipv6networkcontainer := *ipam.NewIpv6networkcontainer() // Ipv6networkcontainer | Object data to update

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.Ipv6networkcontainerAPI.ReferencePut(context.Background(), reference).Ipv6networkcontainer(ipv6networkcontainer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6networkcontainerAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateIpv6networkcontainerResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6networkcontainerAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6networkcontainer object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6networkcontainerAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6networkcontainer** | [**Ipv6networkcontainer**](Ipv6networkcontainer.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

