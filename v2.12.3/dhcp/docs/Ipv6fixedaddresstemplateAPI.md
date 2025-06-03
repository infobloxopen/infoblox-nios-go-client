# Ipv6fixedaddresstemplateAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](Ipv6fixedaddresstemplateAPI.md#Get) | **Get** /ipv6fixedaddresstemplate | Retrieve ipv6fixedaddresstemplate objects
[**Post**](Ipv6fixedaddresstemplateAPI.md#Post) | **Post** /ipv6fixedaddresstemplate | Create a ipv6fixedaddresstemplate object
[**ReferenceDelete**](Ipv6fixedaddresstemplateAPI.md#ReferenceDelete) | **Delete** /ipv6fixedaddresstemplate/{reference} | Delete a ipv6fixedaddresstemplate object
[**ReferenceGet**](Ipv6fixedaddresstemplateAPI.md#ReferenceGet) | **Get** /ipv6fixedaddresstemplate/{reference} | Get a specific ipv6fixedaddresstemplate object
[**ReferencePut**](Ipv6fixedaddresstemplateAPI.md#ReferencePut) | **Put** /ipv6fixedaddresstemplate/{reference} | Update a ipv6fixedaddresstemplate object



## Get

> ListIpv6fixedaddresstemplateResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve ipv6fixedaddresstemplate objects



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
	resp, r, err := apiClient.Ipv6fixedaddresstemplateAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListIpv6fixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddresstemplateAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPIGetRequest` struct via the builder pattern


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

[**ListIpv6fixedaddresstemplateResponse**](ListIpv6fixedaddresstemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateIpv6fixedaddresstemplateResponse Post(ctx).Ipv6fixedaddresstemplate(ipv6fixedaddresstemplate).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a ipv6fixedaddresstemplate object



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
	ipv6fixedaddresstemplate := *dhcp.NewIpv6fixedaddresstemplate() // Ipv6fixedaddresstemplate | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6fixedaddresstemplateAPI.Post(context.Background()).Ipv6fixedaddresstemplate(ipv6fixedaddresstemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateIpv6fixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddresstemplateAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6fixedaddresstemplate** | [**Ipv6fixedaddresstemplate**](Ipv6fixedaddresstemplate.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

Delete a ipv6fixedaddresstemplate object



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
	reference := "reference_example" // string | Reference of the ipv6fixedaddresstemplate object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.Ipv6fixedaddresstemplateAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetIpv6fixedaddresstemplateResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific ipv6fixedaddresstemplate object



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
	reference := "reference_example" // string | Reference of the ipv6fixedaddresstemplate object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6fixedaddresstemplateAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetIpv6fixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddresstemplateAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6fixedaddresstemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## ReferencePut

> UpdateIpv6fixedaddresstemplateResponse ReferencePut(ctx, reference).Ipv6fixedaddresstemplate(ipv6fixedaddresstemplate).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a ipv6fixedaddresstemplate object



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
	reference := "reference_example" // string | Reference of the ipv6fixedaddresstemplate object
	ipv6fixedaddresstemplate := *dhcp.NewIpv6fixedaddresstemplate() // Ipv6fixedaddresstemplate | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.Ipv6fixedaddresstemplateAPI.ReferencePut(context.Background(), reference).Ipv6fixedaddresstemplate(ipv6fixedaddresstemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Ipv6fixedaddresstemplateAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateIpv6fixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `Ipv6fixedaddresstemplateAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ipv6fixedaddresstemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `Ipv6fixedaddresstemplateAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ipv6fixedaddresstemplate** | [**Ipv6fixedaddresstemplate**](Ipv6fixedaddresstemplate.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

