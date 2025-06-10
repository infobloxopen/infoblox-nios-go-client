# DtcMonitorHttpAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcmonitorhttpGet**](DtcMonitorHttpAPI.md#DtcmonitorhttpGet) | **Get** /dtc:monitor:http | Retrieve dtc:monitor:http objects
[**DtcmonitorhttpPost**](DtcMonitorHttpAPI.md#DtcmonitorhttpPost) | **Post** /dtc:monitor:http | Create a dtc:monitor:http object
[**DtcmonitorhttpReferenceDelete**](DtcMonitorHttpAPI.md#DtcmonitorhttpReferenceDelete) | **Delete** /dtc:monitor:http/{reference} | Delete a dtc:monitor:http object
[**DtcmonitorhttpReferenceGet**](DtcMonitorHttpAPI.md#DtcmonitorhttpReferenceGet) | **Get** /dtc:monitor:http/{reference} | Get a specific dtc:monitor:http object
[**DtcmonitorhttpReferencePut**](DtcMonitorHttpAPI.md#DtcmonitorhttpReferencePut) | **Put** /dtc:monitor:http/{reference} | Update a dtc:monitor:http object



## DtcmonitorhttpGet

> ListDtcMonitorHttpResponse DtcmonitorhttpGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:monitor:http objects



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
	resp, r, err := apiClient.DtcMonitorHttpAPI.DtcmonitorhttpGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorHttpAPI.DtcmonitorhttpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorhttpGet`: ListDtcMonitorHttpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorHttpAPI.DtcmonitorhttpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorHttpAPIDtcmonitorhttpGetRequest` struct via the builder pattern


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

[**ListDtcMonitorHttpResponse**](ListDtcMonitorHttpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorhttpPost

> CreateDtcMonitorHttpResponse DtcmonitorhttpPost(ctx).DtcMonitorHttp(dtcMonitorHttp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:monitor:http object



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
	dtcMonitorHttp := *dtc.NewDtcMonitorHttp() // DtcMonitorHttp | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorHttpAPI.DtcmonitorhttpPost(context.Background()).DtcMonitorHttp(dtcMonitorHttp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorHttpAPI.DtcmonitorhttpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorhttpPost`: CreateDtcMonitorHttpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorHttpAPI.DtcmonitorhttpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorHttpAPIDtcmonitorhttpPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorHttp** | [**DtcMonitorHttp**](DtcMonitorHttp.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcMonitorHttpResponse**](CreateDtcMonitorHttpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorhttpReferenceDelete

> DtcmonitorhttpReferenceDelete(ctx, reference).Execute()

Delete a dtc:monitor:http object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:http object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcMonitorHttpAPI.DtcmonitorhttpReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorHttpAPI.DtcmonitorhttpReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:http object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorHttpAPIDtcmonitorhttpReferenceDeleteRequest` struct via the builder pattern


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


## DtcmonitorhttpReferenceGet

> GetDtcMonitorHttpResponse DtcmonitorhttpReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:monitor:http object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:http object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorHttpAPI.DtcmonitorhttpReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorHttpAPI.DtcmonitorhttpReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorhttpReferenceGet`: GetDtcMonitorHttpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorHttpAPI.DtcmonitorhttpReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:http object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorHttpAPIDtcmonitorhttpReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcMonitorHttpResponse**](GetDtcMonitorHttpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcmonitorhttpReferencePut

> UpdateDtcMonitorHttpResponse DtcmonitorhttpReferencePut(ctx, reference).DtcMonitorHttp(dtcMonitorHttp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:monitor:http object



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
	reference := "reference_example" // string | Reference of the dtc:monitor:http object
	dtcMonitorHttp := *dtc.NewDtcMonitorHttp() // DtcMonitorHttp | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcMonitorHttpAPI.DtcmonitorhttpReferencePut(context.Background(), reference).DtcMonitorHttp(dtcMonitorHttp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcMonitorHttpAPI.DtcmonitorhttpReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcmonitorhttpReferencePut`: UpdateDtcMonitorHttpResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcMonitorHttpAPI.DtcmonitorhttpReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:monitor:http object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcMonitorHttpAPIDtcmonitorhttpReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcMonitorHttp** | [**DtcMonitorHttp**](DtcMonitorHttp.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcMonitorHttpResponse**](UpdateDtcMonitorHttpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

