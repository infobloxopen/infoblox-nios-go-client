# DtcMonitorSipAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcmonitorsipGet**](DtcMonitorSipAPI.md#DtcmonitorsipGet) | **Get** /dtc:monitor:sip | Retrieve dtc:monitor:sip objects
[**DtcmonitorsipPost**](DtcMonitorSipAPI.md#DtcmonitorsipPost) | **Post** /dtc:monitor:sip | Create a dtc:monitor:sip object
[**DtcmonitorsipReferenceDelete**](DtcMonitorSipAPI.md#DtcmonitorsipReferenceDelete) | **Delete** /dtc:monitor:sip/{reference} | Delete a dtc:monitor:sip object
[**DtcmonitorsipReferenceGet**](DtcMonitorSipAPI.md#DtcmonitorsipReferenceGet) | **Get** /dtc:monitor:sip/{reference} | Get a specific dtc:monitor:sip object
[**DtcmonitorsipReferencePut**](DtcMonitorSipAPI.md#DtcmonitorsipReferencePut) | **Put** /dtc:monitor:sip/{reference} | Update a dtc:monitor:sip object



## DtcmonitorsipGet

> ListDtcMonitorSipResponse DtcmonitorsipGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:monitor:sip objects



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
	resp, r, err := apiClient.DtcMonitorSipAPI.DtcmonitorsipGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSipAPI.DtcmonitorsipGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorsipGet`: ListDtcMonitorSipResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorSipAPI.DtcmonitorsipGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSipAPIDtcmonitorsipGetRequest` struct via the builder pattern


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

[**ListDtcMonitorSipResponse**](ListDtcMonitorSipResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorsipPost

> CreateDtcMonitorSipResponse DtcmonitorsipPost(ctx).DtcMonitorSip(dtcMonitorSip).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:monitor:sip object



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
	dtcMonitorSip := *dtc.NewDtcMonitorSip() // DtcMonitorSip | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorSipAPI.DtcmonitorsipPost(context.Background()).DtcMonitorSip(dtcMonitorSip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSipAPI.DtcmonitorsipPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorsipPost`: CreateDtcMonitorSipResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorSipAPI.DtcmonitorsipPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSipAPIDtcmonitorsipPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorSip** | [**DtcMonitorSip**](DtcMonitorSip.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcMonitorSipResponse**](CreateDtcMonitorSipResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorsipReferenceDelete

> DtcmonitorsipReferenceDelete(ctx, reference).Execute()

Delete a dtc:monitor:sip object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:sip object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcMonitorSipAPI.DtcmonitorsipReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSipAPI.DtcmonitorsipReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:sip object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSipAPIDtcmonitorsipReferenceDeleteRequest` struct via the builder pattern


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


## DtcmonitorsipReferenceGet

> GetDtcMonitorSipResponse DtcmonitorsipReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:monitor:sip object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:sip object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorSipAPI.DtcmonitorsipReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSipAPI.DtcmonitorsipReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorsipReferenceGet`: GetDtcMonitorSipResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorSipAPI.DtcmonitorsipReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:sip object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSipAPIDtcmonitorsipReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcMonitorSipResponse**](GetDtcMonitorSipResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorsipReferencePut

> UpdateDtcMonitorSipResponse DtcmonitorsipReferencePut(ctx, reference).DtcMonitorSip(dtcMonitorSip).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:monitor:sip object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:sip object
	dtcMonitorSip := *dtc.NewDtcMonitorSip() // DtcMonitorSip | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorSipAPI.DtcmonitorsipReferencePut(context.Background(), reference).DtcMonitorSip(dtcMonitorSip).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSipAPI.DtcmonitorsipReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorsipReferencePut`: UpdateDtcMonitorSipResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorSipAPI.DtcmonitorsipReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:sip object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSipAPIDtcmonitorsipReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorSip** | [**DtcMonitorSip**](DtcMonitorSip.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcMonitorSipResponse**](UpdateDtcMonitorSipResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

