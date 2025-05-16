# DtcmonitoricmpAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](DtcmonitoricmpAPI.md#Get) | **Get** /dtc:monitor:icmp | Retrieve dtc:monitor:icmp objects
[**Post**](DtcmonitoricmpAPI.md#Post) | **Post** /dtc:monitor:icmp | Create a dtc:monitor:icmp object
[**ReferenceDelete**](DtcmonitoricmpAPI.md#ReferenceDelete) | **Delete** /dtc:monitor:icmp/{reference} | Delete a dtc:monitor:icmp object
[**ReferenceGet**](DtcmonitoricmpAPI.md#ReferenceGet) | **Get** /dtc:monitor:icmp/{reference} | Get a specific dtc:monitor:icmp object
[**ReferencePut**](DtcmonitoricmpAPI.md#ReferencePut) | **Put** /dtc:monitor:icmp/{reference} | Update a dtc:monitor:icmp object



## Get

> ListDtcMonitorIcmpResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

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
	resp, r, err := apiClient.DtcmonitoricmpAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcmonitoricmpAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListDtcMonitorIcmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcmonitoricmpAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcmonitoricmpAPIGetRequest` struct via the builder pattern


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


## Post

> CreateDtcMonitorIcmpResponse Post(ctx).DtcMonitorIcmp(dtcMonitorIcmp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.DtcmonitoricmpAPI.Post(context.Background()).DtcMonitorIcmp(dtcMonitorIcmp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcmonitoricmpAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateDtcMonitorIcmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcmonitoricmpAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcmonitoricmpAPIPostRequest` struct via the builder pattern


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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

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
	r, err := apiClient.DtcmonitoricmpAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcmonitoricmpAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `DtcmonitoricmpAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetDtcMonitorIcmpResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.DtcmonitoricmpAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcmonitoricmpAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetDtcMonitorIcmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcmonitoricmpAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:icmp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcmonitoricmpAPIReferenceGetRequest` struct via the builder pattern


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


## ReferencePut

> UpdateDtcMonitorIcmpResponse ReferencePut(ctx, reference).DtcMonitorIcmp(dtcMonitorIcmp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.DtcmonitoricmpAPI.ReferencePut(context.Background(), reference).DtcMonitorIcmp(dtcMonitorIcmp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcmonitoricmpAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateDtcMonitorIcmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcmonitoricmpAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:icmp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcmonitoricmpAPIReferencePutRequest` struct via the builder pattern


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

