# DtcMonitorPdpAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcmonitorpdpGet**](DtcMonitorPdpAPI.md#DtcmonitorpdpGet) | **Get** /dtc:monitor:pdp | Retrieve dtc:monitor:pdp objects
[**DtcmonitorpdpPost**](DtcMonitorPdpAPI.md#DtcmonitorpdpPost) | **Post** /dtc:monitor:pdp | Create a dtc:monitor:pdp object
[**DtcmonitorpdpReferenceDelete**](DtcMonitorPdpAPI.md#DtcmonitorpdpReferenceDelete) | **Delete** /dtc:monitor:pdp/{reference} | Delete a dtc:monitor:pdp object
[**DtcmonitorpdpReferenceGet**](DtcMonitorPdpAPI.md#DtcmonitorpdpReferenceGet) | **Get** /dtc:monitor:pdp/{reference} | Get a specific dtc:monitor:pdp object
[**DtcmonitorpdpReferencePut**](DtcMonitorPdpAPI.md#DtcmonitorpdpReferencePut) | **Put** /dtc:monitor:pdp/{reference} | Update a dtc:monitor:pdp object



## DtcmonitorpdpGet

> ListDtcMonitorPdpResponse DtcmonitorpdpGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:monitor:pdp objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorPdpAPI.DtcmonitorpdpGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorPdpAPI.DtcmonitorpdpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorpdpGet`: ListDtcMonitorPdpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorPdpAPI.DtcmonitorpdpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorPdpAPIDtcmonitorpdpGetRequest` struct via the builder pattern


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

[**ListDtcMonitorPdpResponse**](ListDtcMonitorPdpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorpdpPost

> CreateDtcMonitorPdpResponse DtcmonitorpdpPost(ctx).DtcMonitorPdp(dtcMonitorPdp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:monitor:pdp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	dtcMonitorPdp := *dtc.NewDtcMonitorPdp() // DtcMonitorPdp | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorPdpAPI.DtcmonitorpdpPost(context.Background()).DtcMonitorPdp(dtcMonitorPdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorPdpAPI.DtcmonitorpdpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorpdpPost`: CreateDtcMonitorPdpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorPdpAPI.DtcmonitorpdpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorPdpAPIDtcmonitorpdpPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorPdp** | [**DtcMonitorPdp**](DtcMonitorPdp.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcMonitorPdpResponse**](CreateDtcMonitorPdpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorpdpReferenceDelete

> DtcmonitorpdpReferenceDelete(ctx, reference).Execute()

Delete a dtc:monitor:pdp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:monitor:pdp object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcMonitorPdpAPI.DtcmonitorpdpReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorPdpAPI.DtcmonitorpdpReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:pdp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorPdpAPIDtcmonitorpdpReferenceDeleteRequest` struct via the builder pattern


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


## DtcmonitorpdpReferenceGet

> GetDtcMonitorPdpResponse DtcmonitorpdpReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:monitor:pdp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:monitor:pdp object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorPdpAPI.DtcmonitorpdpReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorPdpAPI.DtcmonitorpdpReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorpdpReferenceGet`: GetDtcMonitorPdpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorPdpAPI.DtcmonitorpdpReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:pdp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorPdpAPIDtcmonitorpdpReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcMonitorPdpResponse**](GetDtcMonitorPdpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorpdpReferencePut

> UpdateDtcMonitorPdpResponse DtcmonitorpdpReferencePut(ctx, reference).DtcMonitorPdp(dtcMonitorPdp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:monitor:pdp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:monitor:pdp object
	dtcMonitorPdp := *dtc.NewDtcMonitorPdp() // DtcMonitorPdp | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorPdpAPI.DtcmonitorpdpReferencePut(context.Background(), reference).DtcMonitorPdp(dtcMonitorPdp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorPdpAPI.DtcmonitorpdpReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorpdpReferencePut`: UpdateDtcMonitorPdpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorPdpAPI.DtcmonitorpdpReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:pdp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorPdpAPIDtcmonitorpdpReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorPdp** | [**DtcMonitorPdp**](DtcMonitorPdp.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcMonitorPdpResponse**](UpdateDtcMonitorPdpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

