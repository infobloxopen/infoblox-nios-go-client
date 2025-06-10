# DiscoveryDeviceinterfaceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverydeviceinterfaceGet**](DiscoveryDeviceinterfaceAPI.md#DiscoverydeviceinterfaceGet) | **Get** /discovery:deviceinterface | Retrieve discovery:deviceinterface objects
[**DiscoverydeviceinterfaceReferenceGet**](DiscoveryDeviceinterfaceAPI.md#DiscoverydeviceinterfaceReferenceGet) | **Get** /discovery:deviceinterface/{reference} | Get a specific discovery:deviceinterface object
[**DiscoverydeviceinterfaceReferencePut**](DiscoveryDeviceinterfaceAPI.md#DiscoverydeviceinterfaceReferencePut) | **Put** /discovery:deviceinterface/{reference} | Update a discovery:deviceinterface object



## DiscoverydeviceinterfaceGet

> ListDiscoveryDeviceinterfaceResponse DiscoverydeviceinterfaceGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve discovery:deviceinterface objects



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
	resp, r, err := apiClient.DiscoveryDeviceinterfaceAPI.DiscoverydeviceinterfaceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDeviceinterfaceAPI.DiscoverydeviceinterfaceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydeviceinterfaceGet`: ListDiscoveryDeviceinterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDeviceinterfaceAPI.DiscoverydeviceinterfaceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDeviceinterfaceAPIDiscoverydeviceinterfaceGetRequest` struct via the builder pattern


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

[**ListDiscoveryDeviceinterfaceResponse**](ListDiscoveryDeviceinterfaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverydeviceinterfaceReferenceGet

> GetDiscoveryDeviceinterfaceResponse DiscoverydeviceinterfaceReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific discovery:deviceinterface object



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
	reference := "reference_example" // string | Reference of the discovery:deviceinterface object

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryDeviceinterfaceAPI.DiscoverydeviceinterfaceReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDeviceinterfaceAPI.DiscoverydeviceinterfaceReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydeviceinterfaceReferenceGet`: GetDiscoveryDeviceinterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDeviceinterfaceAPI.DiscoverydeviceinterfaceReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:deviceinterface object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDeviceinterfaceAPIDiscoverydeviceinterfaceReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDiscoveryDeviceinterfaceResponse**](GetDiscoveryDeviceinterfaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DiscoverydeviceinterfaceReferencePut

> UpdateDiscoveryDeviceinterfaceResponse DiscoverydeviceinterfaceReferencePut(ctx, reference).DiscoveryDeviceinterface(discoveryDeviceinterface).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a discovery:deviceinterface object



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
	reference := "reference_example" // string | Reference of the discovery:deviceinterface object
	discoveryDeviceinterface := *discovery.NewDiscoveryDeviceinterface() // DiscoveryDeviceinterface | Object data to update

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoveryDeviceinterfaceAPI.DiscoverydeviceinterfaceReferencePut(context.Background(), reference).DiscoveryDeviceinterface(discoveryDeviceinterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryDeviceinterfaceAPI.DiscoverydeviceinterfaceReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverydeviceinterfaceReferencePut`: UpdateDiscoveryDeviceinterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryDeviceinterfaceAPI.DiscoverydeviceinterfaceReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:deviceinterface object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryDeviceinterfaceAPIDiscoverydeviceinterfaceReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**discoveryDeviceinterface** | [**DiscoveryDeviceinterface**](DiscoveryDeviceinterface.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDiscoveryDeviceinterfaceResponse**](UpdateDiscoveryDeviceinterfaceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

