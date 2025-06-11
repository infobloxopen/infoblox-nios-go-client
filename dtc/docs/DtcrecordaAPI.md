# DtcRecordAAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcrecordaGet**](DtcRecordAAPI.md#DtcrecordaGet) | **Get** /dtc:record:a | Retrieve dtc:record:a objects
[**DtcrecordaPost**](DtcRecordAAPI.md#DtcrecordaPost) | **Post** /dtc:record:a | Create a dtc:record:a object
[**DtcrecordaReferenceDelete**](DtcRecordAAPI.md#DtcrecordaReferenceDelete) | **Delete** /dtc:record:a/{reference} | Delete a dtc:record:a object
[**DtcrecordaReferenceGet**](DtcRecordAAPI.md#DtcrecordaReferenceGet) | **Get** /dtc:record:a/{reference} | Get a specific dtc:record:a object
[**DtcrecordaReferencePut**](DtcRecordAAPI.md#DtcrecordaReferencePut) | **Put** /dtc:record:a/{reference} | Update a dtc:record:a object



## DtcrecordaGet

> ListDtcRecordAResponse DtcrecordaGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:record:a objects



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
	resp, r, err := apiClient.DtcRecordAAPI.DtcrecordaGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAAPI.DtcrecordaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordaGet`: ListDtcRecordAResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordAAPI.DtcrecordaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAAPIDtcrecordaGetRequest` struct via the builder pattern


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

[**ListDtcRecordAResponse**](ListDtcRecordAResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordaPost

> CreateDtcRecordAResponse DtcrecordaPost(ctx).DtcRecordA(dtcRecordA).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:record:a object



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
	dtcRecordA := *dtc.NewDtcRecordA() // DtcRecordA | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordAAPI.DtcrecordaPost(context.Background()).DtcRecordA(dtcRecordA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAAPI.DtcrecordaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordaPost`: CreateDtcRecordAResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordAAPI.DtcrecordaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAAPIDtcrecordaPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordA** | [**DtcRecordA**](DtcRecordA.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcRecordAResponse**](CreateDtcRecordAResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordaReferenceDelete

> DtcrecordaReferenceDelete(ctx, reference).Execute()

Delete a dtc:record:a object



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
	reference := "reference_example" // string | Reference of the dtc:record:a object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcRecordAAPI.DtcrecordaReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAAPI.DtcrecordaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:a object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAAPIDtcrecordaReferenceDeleteRequest` struct via the builder pattern


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


## DtcrecordaReferenceGet

> GetDtcRecordAResponse DtcrecordaReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:record:a object



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
	reference := "reference_example" // string | Reference of the dtc:record:a object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordAAPI.DtcrecordaReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAAPI.DtcrecordaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordaReferenceGet`: GetDtcRecordAResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordAAPI.DtcrecordaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:a object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAAPIDtcrecordaReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcRecordAResponse**](GetDtcRecordAResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordaReferencePut

> UpdateDtcRecordAResponse DtcrecordaReferencePut(ctx, reference).DtcRecordA(dtcRecordA).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:record:a object



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
	reference := "reference_example" // string | Reference of the dtc:record:a object
	dtcRecordA := *dtc.NewDtcRecordA() // DtcRecordA | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordAAPI.DtcrecordaReferencePut(context.Background(), reference).DtcRecordA(dtcRecordA).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAAPI.DtcrecordaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordaReferencePut`: UpdateDtcRecordAResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordAAPI.DtcrecordaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:a object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAAPIDtcrecordaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordA** | [**DtcRecordA**](DtcRecordA.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcRecordAResponse**](UpdateDtcRecordAResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

