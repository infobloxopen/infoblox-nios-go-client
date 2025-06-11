# DtcMonitorSnmpAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcmonitorsnmpGet**](DtcMonitorSnmpAPI.md#DtcmonitorsnmpGet) | **Get** /dtc:monitor:snmp | Retrieve dtc:monitor:snmp objects
[**DtcmonitorsnmpPost**](DtcMonitorSnmpAPI.md#DtcmonitorsnmpPost) | **Post** /dtc:monitor:snmp | Create a dtc:monitor:snmp object
[**DtcmonitorsnmpReferenceDelete**](DtcMonitorSnmpAPI.md#DtcmonitorsnmpReferenceDelete) | **Delete** /dtc:monitor:snmp/{reference} | Delete a dtc:monitor:snmp object
[**DtcmonitorsnmpReferenceGet**](DtcMonitorSnmpAPI.md#DtcmonitorsnmpReferenceGet) | **Get** /dtc:monitor:snmp/{reference} | Get a specific dtc:monitor:snmp object
[**DtcmonitorsnmpReferencePut**](DtcMonitorSnmpAPI.md#DtcmonitorsnmpReferencePut) | **Put** /dtc:monitor:snmp/{reference} | Update a dtc:monitor:snmp object



## DtcmonitorsnmpGet

> ListDtcMonitorSnmpResponse DtcmonitorsnmpGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:monitor:snmp objects



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
	resp, r, err := apiClient.DtcMonitorSnmpAPI.DtcmonitorsnmpGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSnmpAPI.DtcmonitorsnmpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorsnmpGet`: ListDtcMonitorSnmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorSnmpAPI.DtcmonitorsnmpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSnmpAPIDtcmonitorsnmpGetRequest` struct via the builder pattern


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

[**ListDtcMonitorSnmpResponse**](ListDtcMonitorSnmpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorsnmpPost

> CreateDtcMonitorSnmpResponse DtcmonitorsnmpPost(ctx).DtcMonitorSnmp(dtcMonitorSnmp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:monitor:snmp object



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
	dtcMonitorSnmp := *dtc.NewDtcMonitorSnmp() // DtcMonitorSnmp | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorSnmpAPI.DtcmonitorsnmpPost(context.Background()).DtcMonitorSnmp(dtcMonitorSnmp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSnmpAPI.DtcmonitorsnmpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorsnmpPost`: CreateDtcMonitorSnmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorSnmpAPI.DtcmonitorsnmpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSnmpAPIDtcmonitorsnmpPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorSnmp** | [**DtcMonitorSnmp**](DtcMonitorSnmp.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcMonitorSnmpResponse**](CreateDtcMonitorSnmpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorsnmpReferenceDelete

> DtcmonitorsnmpReferenceDelete(ctx, reference).Execute()

Delete a dtc:monitor:snmp object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:snmp object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcMonitorSnmpAPI.DtcmonitorsnmpReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSnmpAPI.DtcmonitorsnmpReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:snmp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSnmpAPIDtcmonitorsnmpReferenceDeleteRequest` struct via the builder pattern


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


## DtcmonitorsnmpReferenceGet

> GetDtcMonitorSnmpResponse DtcmonitorsnmpReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:monitor:snmp object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:snmp object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorSnmpAPI.DtcmonitorsnmpReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSnmpAPI.DtcmonitorsnmpReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorsnmpReferenceGet`: GetDtcMonitorSnmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorSnmpAPI.DtcmonitorsnmpReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:snmp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSnmpAPIDtcmonitorsnmpReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcMonitorSnmpResponse**](GetDtcMonitorSnmpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorsnmpReferencePut

> UpdateDtcMonitorSnmpResponse DtcmonitorsnmpReferencePut(ctx, reference).DtcMonitorSnmp(dtcMonitorSnmp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:monitor:snmp object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:snmp object
	dtcMonitorSnmp := *dtc.NewDtcMonitorSnmp() // DtcMonitorSnmp | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorSnmpAPI.DtcmonitorsnmpReferencePut(context.Background(), reference).DtcMonitorSnmp(dtcMonitorSnmp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorSnmpAPI.DtcmonitorsnmpReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorsnmpReferencePut`: UpdateDtcMonitorSnmpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorSnmpAPI.DtcmonitorsnmpReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:snmp object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorSnmpAPIDtcmonitorsnmpReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorSnmp** | [**DtcMonitorSnmp**](DtcMonitorSnmp.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcMonitorSnmpResponse**](UpdateDtcMonitorSnmpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

