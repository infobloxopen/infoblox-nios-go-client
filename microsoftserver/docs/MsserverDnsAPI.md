# MsserverDnsAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MsserverdnsGet**](MsserverDnsAPI.md#MsserverdnsGet) | **Get** /msserver:dns | Retrieve msserver:dns objects
[**MsserverdnsPost**](MsserverDnsAPI.md#MsserverdnsPost) | **Post** /msserver:dns | Create a msserver:dns object
[**MsserverdnsReferenceDelete**](MsserverDnsAPI.md#MsserverdnsReferenceDelete) | **Delete** /msserver:dns/{reference} | Delete a msserver:dns object
[**MsserverdnsReferenceGet**](MsserverDnsAPI.md#MsserverdnsReferenceGet) | **Get** /msserver:dns/{reference} | Get a specific msserver:dns object
[**MsserverdnsReferencePut**](MsserverDnsAPI.md#MsserverdnsReferencePut) | **Put** /msserver:dns/{reference} | Update a msserver:dns object



## MsserverdnsGet

> ListMsserverDnsResponse MsserverdnsGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve msserver:dns objects



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
	resp, r, err := apiClient.MsserverDnsAPI.MsserverdnsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDnsAPI.MsserverdnsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserverdnsGet`: ListMsserverDnsResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverDnsAPI.MsserverdnsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MsserverDnsAPIMsserverdnsGetRequest` struct via the builder pattern


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

[**ListMsserverDnsResponse**](ListMsserverDnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserverdnsPost

> CreateMsserverDnsResponse MsserverdnsPost(ctx).MsserverDns(msserverDns).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a msserver:dns object



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
	msserverDns := *microsoftserver.NewMsserverDns() // MsserverDns | Object data to create

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverDnsAPI.MsserverdnsPost(context.Background()).MsserverDns(msserverDns).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDnsAPI.MsserverdnsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserverdnsPost`: CreateMsserverDnsResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverDnsAPI.MsserverdnsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MsserverDnsAPIMsserverdnsPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**msserverDns** | [**MsserverDns**](MsserverDns.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateMsserverDnsResponse**](CreateMsserverDnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserverdnsReferenceDelete

> MsserverdnsReferenceDelete(ctx, reference).Execute()

Delete a msserver:dns object



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
	reference := "reference_example" // string | Reference of the msserver:dns object

	apiClient := microsoftserver.NewAPIClient()
	r, err := apiClient.MsserverDnsAPI.MsserverdnsReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDnsAPI.MsserverdnsReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:dns object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverDnsAPIMsserverdnsReferenceDeleteRequest` struct via the builder pattern


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


## MsserverdnsReferenceGet

> GetMsserverDnsResponse MsserverdnsReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific msserver:dns object



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
	reference := "reference_example" // string | Reference of the msserver:dns object

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverDnsAPI.MsserverdnsReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDnsAPI.MsserverdnsReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserverdnsReferenceGet`: GetMsserverDnsResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverDnsAPI.MsserverdnsReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:dns object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverDnsAPIMsserverdnsReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMsserverDnsResponse**](GetMsserverDnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserverdnsReferencePut

> UpdateMsserverDnsResponse MsserverdnsReferencePut(ctx, reference).MsserverDns(msserverDns).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a msserver:dns object



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
	reference := "reference_example" // string | Reference of the msserver:dns object
	msserverDns := *microsoftserver.NewMsserverDns() // MsserverDns | Object data to update

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverDnsAPI.MsserverdnsReferencePut(context.Background(), reference).MsserverDns(msserverDns).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverDnsAPI.MsserverdnsReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserverdnsReferencePut`: UpdateMsserverDnsResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverDnsAPI.MsserverdnsReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:dns object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverDnsAPIMsserverdnsReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**msserverDns** | [**MsserverDns**](MsserverDns.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMsserverDnsResponse**](UpdateMsserverDnsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

