# DtcMonitorTcpAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcmonitortcpGet**](DtcMonitorTcpAPI.md#DtcmonitortcpGet) | **Get** /dtc:monitor:tcp | Retrieve dtc:monitor:tcp objects
[**DtcmonitortcpPost**](DtcMonitorTcpAPI.md#DtcmonitortcpPost) | **Post** /dtc:monitor:tcp | Create a dtc:monitor:tcp object
[**DtcmonitortcpReferenceDelete**](DtcMonitorTcpAPI.md#DtcmonitortcpReferenceDelete) | **Delete** /dtc:monitor:tcp/{reference} | Delete a dtc:monitor:tcp object
[**DtcmonitortcpReferenceGet**](DtcMonitorTcpAPI.md#DtcmonitortcpReferenceGet) | **Get** /dtc:monitor:tcp/{reference} | Get a specific dtc:monitor:tcp object
[**DtcmonitortcpReferencePut**](DtcMonitorTcpAPI.md#DtcmonitortcpReferencePut) | **Put** /dtc:monitor:tcp/{reference} | Update a dtc:monitor:tcp object



## DtcmonitortcpGet

> ListDtcMonitorTcpResponse DtcmonitortcpGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:monitor:tcp objects



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
	resp, r, err := apiClient.DtcMonitorTcpAPI.DtcmonitortcpGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorTcpAPI.DtcmonitortcpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitortcpGet`: ListDtcMonitorTcpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorTcpAPI.DtcmonitortcpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorTcpAPIDtcmonitortcpGetRequest` struct via the builder pattern


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

[**ListDtcMonitorTcpResponse**](ListDtcMonitorTcpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitortcpPost

> CreateDtcMonitorTcpResponse DtcmonitortcpPost(ctx).DtcMonitorTcp(dtcMonitorTcp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:monitor:tcp object



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
	dtcMonitorTcp := *dtc.NewDtcMonitorTcp() // DtcMonitorTcp | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorTcpAPI.DtcmonitortcpPost(context.Background()).DtcMonitorTcp(dtcMonitorTcp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorTcpAPI.DtcmonitortcpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitortcpPost`: CreateDtcMonitorTcpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorTcpAPI.DtcmonitortcpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorTcpAPIDtcmonitortcpPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorTcp** | [**DtcMonitorTcp**](DtcMonitorTcp.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcMonitorTcpResponse**](CreateDtcMonitorTcpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitortcpReferenceDelete

> DtcmonitortcpReferenceDelete(ctx, reference).Execute()

Delete a dtc:monitor:tcp object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:tcp object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcMonitorTcpAPI.DtcmonitortcpReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorTcpAPI.DtcmonitortcpReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:tcp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorTcpAPIDtcmonitortcpReferenceDeleteRequest` struct via the builder pattern


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


## DtcmonitortcpReferenceGet

> GetDtcMonitorTcpResponse DtcmonitortcpReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:monitor:tcp object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:tcp object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorTcpAPI.DtcmonitortcpReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorTcpAPI.DtcmonitortcpReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitortcpReferenceGet`: GetDtcMonitorTcpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorTcpAPI.DtcmonitortcpReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:tcp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorTcpAPIDtcmonitortcpReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcMonitorTcpResponse**](GetDtcMonitorTcpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitortcpReferencePut

> UpdateDtcMonitorTcpResponse DtcmonitortcpReferencePut(ctx, reference).DtcMonitorTcp(dtcMonitorTcp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:monitor:tcp object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:tcp object
	dtcMonitorTcp := *dtc.NewDtcMonitorTcp() // DtcMonitorTcp | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorTcpAPI.DtcmonitortcpReferencePut(context.Background(), reference).DtcMonitorTcp(dtcMonitorTcp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorTcpAPI.DtcmonitortcpReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitortcpReferencePut`: UpdateDtcMonitorTcpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorTcpAPI.DtcmonitortcpReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:tcp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorTcpAPIDtcmonitortcpReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorTcp** | [**DtcMonitorTcp**](DtcMonitorTcp.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcMonitorTcpResponse**](UpdateDtcMonitorTcpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

