# MsserverDhcpAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MsserverdhcpGet**](MsserverDhcpAPI.md#MsserverdhcpGet) | **Get** /msserver:dhcp | Retrieve msserver:dhcp objects
[**MsserverdhcpPost**](MsserverDhcpAPI.md#MsserverdhcpPost) | **Post** /msserver:dhcp | Create a msserver:dhcp object
[**MsserverdhcpReferenceDelete**](MsserverDhcpAPI.md#MsserverdhcpReferenceDelete) | **Delete** /msserver:dhcp/{reference} | Delete a msserver:dhcp object
[**MsserverdhcpReferenceGet**](MsserverDhcpAPI.md#MsserverdhcpReferenceGet) | **Get** /msserver:dhcp/{reference} | Get a specific msserver:dhcp object
[**MsserverdhcpReferencePut**](MsserverDhcpAPI.md#MsserverdhcpReferencePut) | **Put** /msserver:dhcp/{reference} | Update a msserver:dhcp object



## MsserverdhcpGet

> ListMsserverDhcpResponse MsserverdhcpGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve msserver:dhcp objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverDhcpAPI.MsserverdhcpGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDhcpAPI.MsserverdhcpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserverdhcpGet`: ListMsserverDhcpResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverDhcpAPI.MsserverdhcpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MsserverDhcpAPIMsserverdhcpGetRequest` struct via the builder pattern


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

[**ListMsserverDhcpResponse**](ListMsserverDhcpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserverdhcpPost

> CreateMsserverDhcpResponse MsserverdhcpPost(ctx).MsserverDhcp(msserverDhcp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a msserver:dhcp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {
	msserverDhcp := *microsoftserver.NewMsserverDhcp() // MsserverDhcp | Object data to create

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverDhcpAPI.MsserverdhcpPost(context.Background()).MsserverDhcp(msserverDhcp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDhcpAPI.MsserverdhcpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserverdhcpPost`: CreateMsserverDhcpResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverDhcpAPI.MsserverdhcpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MsserverDhcpAPIMsserverdhcpPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**msserverDhcp** | [**MsserverDhcp**](MsserverDhcp.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateMsserverDhcpResponse**](CreateMsserverDhcpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserverdhcpReferenceDelete

> MsserverdhcpReferenceDelete(ctx, reference).Execute()

Delete a msserver:dhcp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {
	reference := "reference_example" // string | Reference of the msserver:dhcp object

	apiClient := microsoftserver.NewAPIClient()
	r, err := apiClient.MsserverDhcpAPI.MsserverdhcpReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDhcpAPI.MsserverdhcpReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:dhcp object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverDhcpAPIMsserverdhcpReferenceDeleteRequest` struct via the builder pattern


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


## MsserverdhcpReferenceGet

> GetMsserverDhcpResponse MsserverdhcpReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific msserver:dhcp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {
	reference := "reference_example" // string | Reference of the msserver:dhcp object

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverDhcpAPI.MsserverdhcpReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDhcpAPI.MsserverdhcpReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserverdhcpReferenceGet`: GetMsserverDhcpResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverDhcpAPI.MsserverdhcpReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:dhcp object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverDhcpAPIMsserverdhcpReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMsserverDhcpResponse**](GetMsserverDhcpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserverdhcpReferencePut

> UpdateMsserverDhcpResponse MsserverdhcpReferencePut(ctx, reference).MsserverDhcp(msserverDhcp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a msserver:dhcp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {
	reference := "reference_example" // string | Reference of the msserver:dhcp object
	msserverDhcp := *microsoftserver.NewMsserverDhcp() // MsserverDhcp | Object data to update

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverDhcpAPI.MsserverdhcpReferencePut(context.Background(), reference).MsserverDhcp(msserverDhcp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDhcpAPI.MsserverdhcpReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserverdhcpReferencePut`: UpdateMsserverDhcpResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverDhcpAPI.MsserverdhcpReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:dhcp object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverDhcpAPIMsserverdhcpReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**msserverDhcp** | [**MsserverDhcp**](MsserverDhcp.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMsserverDhcpResponse**](UpdateMsserverDhcpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

