# DtcRecordNaptrAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcrecordnaptrGet**](DtcRecordNaptrAPI.md#DtcrecordnaptrGet) | **Get** /dtc:record:naptr | Retrieve dtc:record:naptr objects
[**DtcrecordnaptrPost**](DtcRecordNaptrAPI.md#DtcrecordnaptrPost) | **Post** /dtc:record:naptr | Create a dtc:record:naptr object
[**DtcrecordnaptrReferenceDelete**](DtcRecordNaptrAPI.md#DtcrecordnaptrReferenceDelete) | **Delete** /dtc:record:naptr/{reference} | Delete a dtc:record:naptr object
[**DtcrecordnaptrReferenceGet**](DtcRecordNaptrAPI.md#DtcrecordnaptrReferenceGet) | **Get** /dtc:record:naptr/{reference} | Get a specific dtc:record:naptr object
[**DtcrecordnaptrReferencePut**](DtcRecordNaptrAPI.md#DtcrecordnaptrReferencePut) | **Put** /dtc:record:naptr/{reference} | Update a dtc:record:naptr object



## DtcrecordnaptrGet

> ListDtcRecordNaptrResponse DtcrecordnaptrGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:record:naptr objects



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
	resp, r, err := apiClient.DtcRecordNaptrAPI.DtcrecordnaptrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordNaptrAPI.DtcrecordnaptrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordnaptrGet`: ListDtcRecordNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordNaptrAPI.DtcrecordnaptrGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordNaptrAPIDtcrecordnaptrGetRequest` struct via the builder pattern


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

[**ListDtcRecordNaptrResponse**](ListDtcRecordNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordnaptrPost

> CreateDtcRecordNaptrResponse DtcrecordnaptrPost(ctx).DtcRecordNaptr(dtcRecordNaptr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:record:naptr object



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
	dtcRecordNaptr := *dtc.NewDtcRecordNaptr() // DtcRecordNaptr | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordNaptrAPI.DtcrecordnaptrPost(context.Background()).DtcRecordNaptr(dtcRecordNaptr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordNaptrAPI.DtcrecordnaptrPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordnaptrPost`: CreateDtcRecordNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordNaptrAPI.DtcrecordnaptrPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordNaptrAPIDtcrecordnaptrPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordNaptr** | [**DtcRecordNaptr**](DtcRecordNaptr.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcRecordNaptrResponse**](CreateDtcRecordNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordnaptrReferenceDelete

> DtcrecordnaptrReferenceDelete(ctx, reference).Execute()

Delete a dtc:record:naptr object



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
	reference := "reference_example" // string | Reference of the dtc:record:naptr object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcRecordNaptrAPI.DtcrecordnaptrReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordNaptrAPI.DtcrecordnaptrReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:naptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordNaptrAPIDtcrecordnaptrReferenceDeleteRequest` struct via the builder pattern


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


## DtcrecordnaptrReferenceGet

> GetDtcRecordNaptrResponse DtcrecordnaptrReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:record:naptr object



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
	reference := "reference_example" // string | Reference of the dtc:record:naptr object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordNaptrAPI.DtcrecordnaptrReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordNaptrAPI.DtcrecordnaptrReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordnaptrReferenceGet`: GetDtcRecordNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordNaptrAPI.DtcrecordnaptrReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:naptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordNaptrAPIDtcrecordnaptrReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcRecordNaptrResponse**](GetDtcRecordNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordnaptrReferencePut

> UpdateDtcRecordNaptrResponse DtcrecordnaptrReferencePut(ctx, reference).DtcRecordNaptr(dtcRecordNaptr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:record:naptr object



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
	reference := "reference_example" // string | Reference of the dtc:record:naptr object
	dtcRecordNaptr := *dtc.NewDtcRecordNaptr() // DtcRecordNaptr | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordNaptrAPI.DtcrecordnaptrReferencePut(context.Background(), reference).DtcRecordNaptr(dtcRecordNaptr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordNaptrAPI.DtcrecordnaptrReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordnaptrReferencePut`: UpdateDtcRecordNaptrResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordNaptrAPI.DtcrecordnaptrReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:naptr object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordNaptrAPIDtcrecordnaptrReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordNaptr** | [**DtcRecordNaptr**](DtcRecordNaptr.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcRecordNaptrResponse**](UpdateDtcRecordNaptrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

