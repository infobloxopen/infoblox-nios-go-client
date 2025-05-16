# Ipv6fixedaddressAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](Ipv6fixedaddressAPI.md#Get) | **Get** /ipv6fixedaddress | Retrieve ipv6fixedaddress objects
[**Post**](Ipv6fixedaddressAPI.md#Post) | **Post** /ipv6fixedaddress | Create a ipv6fixedaddress object
[**ReferenceDelete**](Ipv6fixedaddressAPI.md#ReferenceDelete) | **Delete** /ipv6fixedaddress/{reference} | Delete a ipv6fixedaddress object
[**ReferenceGet**](Ipv6fixedaddressAPI.md#ReferenceGet) | **Get** /ipv6fixedaddress/{reference} | Get a specific ipv6fixedaddress object
[**ReferencePut**](Ipv6fixedaddressAPI.md#ReferencePut) | **Put** /ipv6fixedaddress/{reference} | Update a ipv6fixedaddress object



## Get

> ListIpv6fixedaddressResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve ipv6fixedaddress objects



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
	resp, r, err := apiClient.Ipv6fixedaddressAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddressAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListIpv6fixedaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddressAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddressAPIGetRequest` struct via the builder pattern


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

[**ListIpv6fixedaddressResponse**](ListIpv6fixedaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateIpv6fixedaddressResponse Post(ctx).Ipv6fixedaddress(ipv6fixedaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a ipv6fixedaddress object



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
	ipv6fixedaddress := *dhcp.NewIpv6fixedaddress() // Ipv6fixedaddress | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6fixedaddressAPI.Post(context.Background()).Ipv6fixedaddress(ipv6fixedaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddressAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateIpv6fixedaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddressAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddressAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6fixedaddress** | [**Ipv6fixedaddress**](Ipv6fixedaddress.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateIpv6fixedaddressResponse**](CreateIpv6fixedaddressResponse.md)

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

Delete a ipv6fixedaddress object



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
	reference := "reference_example" // string | Reference of the ipv6fixedaddress object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.Ipv6fixedaddressAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddressAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6fixedaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddressAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetIpv6fixedaddressResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific ipv6fixedaddress object



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
	reference := "reference_example" // string | Reference of the ipv6fixedaddress object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6fixedaddressAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddressAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetIpv6fixedaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddressAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6fixedaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddressAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetIpv6fixedaddressResponse**](GetIpv6fixedaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateIpv6fixedaddressResponse ReferencePut(ctx, reference).Ipv6fixedaddress(ipv6fixedaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a ipv6fixedaddress object



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
	reference := "reference_example" // string | Reference of the ipv6fixedaddress object
	ipv6fixedaddress := *dhcp.NewIpv6fixedaddress() // Ipv6fixedaddress | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6fixedaddressAPI.ReferencePut(context.Background(), reference).Ipv6fixedaddress(ipv6fixedaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddressAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateIpv6fixedaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddressAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6fixedaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddressAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6fixedaddress** | [**Ipv6fixedaddress**](Ipv6fixedaddress.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateIpv6fixedaddressResponse**](UpdateIpv6fixedaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

