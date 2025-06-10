# DiscoveryDevicesupportbundleAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverydevicesupportbundleGet**](DiscoveryDevicesupportbundleAPI.md#DiscoverydevicesupportbundleGet) | **Get** /discovery:devicesupportbundle | Retrieve discovery:devicesupportbundle objects
[**DiscoverydevicesupportbundleReferenceDelete**](DiscoveryDevicesupportbundleAPI.md#DiscoverydevicesupportbundleReferenceDelete) | **Delete** /discovery:devicesupportbundle/{reference} | Delete a discovery:devicesupportbundle object
[**DiscoverydevicesupportbundleReferenceGet**](DiscoveryDevicesupportbundleAPI.md#DiscoverydevicesupportbundleReferenceGet) | **Get** /discovery:devicesupportbundle/{reference} | Get a specific discovery:devicesupportbundle object



## DiscoverydevicesupportbundleGet

> ListDiscoveryDevicesupportbundleResponse DiscoverydevicesupportbundleGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve discovery:devicesupportbundle objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryDevicesupportbundleAPI.DiscoverydevicesupportbundleGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDevicesupportbundleAPI.DiscoverydevicesupportbundleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydevicesupportbundleGet`: ListDiscoveryDevicesupportbundleResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDevicesupportbundleAPI.DiscoverydevicesupportbundleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDevicesupportbundleAPIDiscoverydevicesupportbundleGetRequest` struct via the builder pattern


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

[**ListDiscoveryDevicesupportbundleResponse**](ListDiscoveryDevicesupportbundleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverydevicesupportbundleReferenceDelete

> DiscoverydevicesupportbundleReferenceDelete(ctx, reference).Execute()

Delete a discovery:devicesupportbundle object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {
	reference := "reference_example" // string | Reference of the discovery:devicesupportbundle object

	apiClient := discovery.NewAPIClient()
	r, err := apiClient.DiscoveryDevicesupportbundleAPI.DiscoverydevicesupportbundleReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDevicesupportbundleAPI.DiscoverydevicesupportbundleReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:devicesupportbundle object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDevicesupportbundleAPIDiscoverydevicesupportbundleReferenceDeleteRequest` struct via the builder pattern


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


## DiscoverydevicesupportbundleReferenceGet

> GetDiscoveryDevicesupportbundleResponse DiscoverydevicesupportbundleReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific discovery:devicesupportbundle object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {
	reference := "reference_example" // string | Reference of the discovery:devicesupportbundle object

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryDevicesupportbundleAPI.DiscoverydevicesupportbundleReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDevicesupportbundleAPI.DiscoverydevicesupportbundleReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydevicesupportbundleReferenceGet`: GetDiscoveryDevicesupportbundleResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDevicesupportbundleAPI.DiscoverydevicesupportbundleReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:devicesupportbundle object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDevicesupportbundleAPIDiscoverydevicesupportbundleReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDiscoveryDevicesupportbundleResponse**](GetDiscoveryDevicesupportbundleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

