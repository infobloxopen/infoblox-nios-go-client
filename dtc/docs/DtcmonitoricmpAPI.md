# DtcMonitorIcmpAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcmonitoricmpGet**](DtcMonitorIcmpAPI.md#DtcmonitoricmpGet) | **Get** /dtc:monitor:icmp | Retrieve dtc:monitor:icmp objects
[**DtcmonitoricmpPost**](DtcMonitorIcmpAPI.md#DtcmonitoricmpPost) | **Post** /dtc:monitor:icmp | Create a dtc:monitor:icmp object
[**DtcmonitoricmpReferenceDelete**](DtcMonitorIcmpAPI.md#DtcmonitoricmpReferenceDelete) | **Delete** /dtc:monitor:icmp/{reference} | Delete a dtc:monitor:icmp object
[**DtcmonitoricmpReferenceGet**](DtcMonitorIcmpAPI.md#DtcmonitoricmpReferenceGet) | **Get** /dtc:monitor:icmp/{reference} | Get a specific dtc:monitor:icmp object
[**DtcmonitoricmpReferencePut**](DtcMonitorIcmpAPI.md#DtcmonitoricmpReferencePut) | **Put** /dtc:monitor:icmp/{reference} | Update a dtc:monitor:icmp object



## DtcmonitoricmpGet

> ListDtcMonitorIcmpResponse DtcmonitoricmpGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:monitor:icmp objects



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
	resp, r, err := apiClient.DtcMonitorIcmpAPI.DtcmonitoricmpGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorIcmpAPI.DtcmonitoricmpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitoricmpGet`: ListDtcMonitorIcmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorIcmpAPI.DtcmonitoricmpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorIcmpAPIDtcmonitoricmpGetRequest` struct via the builder pattern


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

[**ListDtcMonitorIcmpResponse**](ListDtcMonitorIcmpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitoricmpPost

> CreateDtcMonitorIcmpResponse DtcmonitoricmpPost(ctx).DtcMonitorIcmp(dtcMonitorIcmp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:monitor:icmp object



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
	dtcMonitorIcmp := *dtc.NewDtcMonitorIcmp() // DtcMonitorIcmp | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorIcmpAPI.DtcmonitoricmpPost(context.Background()).DtcMonitorIcmp(dtcMonitorIcmp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorIcmpAPI.DtcmonitoricmpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitoricmpPost`: CreateDtcMonitorIcmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorIcmpAPI.DtcmonitoricmpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorIcmpAPIDtcmonitoricmpPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorIcmp** | [**DtcMonitorIcmp**](DtcMonitorIcmp.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcMonitorIcmpResponse**](CreateDtcMonitorIcmpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitoricmpReferenceDelete

> DtcmonitoricmpReferenceDelete(ctx, reference).Execute()

Delete a dtc:monitor:icmp object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:icmp object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcMonitorIcmpAPI.DtcmonitoricmpReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorIcmpAPI.DtcmonitoricmpReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:icmp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorIcmpAPIDtcmonitoricmpReferenceDeleteRequest` struct via the builder pattern


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


## DtcmonitoricmpReferenceGet

> GetDtcMonitorIcmpResponse DtcmonitoricmpReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:monitor:icmp object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:icmp object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorIcmpAPI.DtcmonitoricmpReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorIcmpAPI.DtcmonitoricmpReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitoricmpReferenceGet`: GetDtcMonitorIcmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorIcmpAPI.DtcmonitoricmpReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:icmp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorIcmpAPIDtcmonitoricmpReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcMonitorIcmpResponse**](GetDtcMonitorIcmpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitoricmpReferencePut

> UpdateDtcMonitorIcmpResponse DtcmonitoricmpReferencePut(ctx, reference).DtcMonitorIcmp(dtcMonitorIcmp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:monitor:icmp object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:icmp object
	dtcMonitorIcmp := *dtc.NewDtcMonitorIcmp() // DtcMonitorIcmp | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorIcmpAPI.DtcmonitoricmpReferencePut(context.Background(), reference).DtcMonitorIcmp(dtcMonitorIcmp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorIcmpAPI.DtcmonitoricmpReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitoricmpReferencePut`: UpdateDtcMonitorIcmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorIcmpAPI.DtcmonitoricmpReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:icmp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorIcmpAPIDtcmonitoricmpReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorIcmp** | [**DtcMonitorIcmp**](DtcMonitorIcmp.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcMonitorIcmpResponse**](UpdateDtcMonitorIcmpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

