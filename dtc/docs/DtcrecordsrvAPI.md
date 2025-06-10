# DtcRecordSrvAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcrecordsrvGet**](DtcRecordSrvAPI.md#DtcrecordsrvGet) | **Get** /dtc:record:srv | Retrieve dtc:record:srv objects
[**DtcrecordsrvPost**](DtcRecordSrvAPI.md#DtcrecordsrvPost) | **Post** /dtc:record:srv | Create a dtc:record:srv object
[**DtcrecordsrvReferenceDelete**](DtcRecordSrvAPI.md#DtcrecordsrvReferenceDelete) | **Delete** /dtc:record:srv/{reference} | Delete a dtc:record:srv object
[**DtcrecordsrvReferenceGet**](DtcRecordSrvAPI.md#DtcrecordsrvReferenceGet) | **Get** /dtc:record:srv/{reference} | Get a specific dtc:record:srv object
[**DtcrecordsrvReferencePut**](DtcRecordSrvAPI.md#DtcrecordsrvReferencePut) | **Put** /dtc:record:srv/{reference} | Update a dtc:record:srv object



## DtcrecordsrvGet

> ListDtcRecordSrvResponse DtcrecordsrvGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:record:srv objects



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
	resp, r, err := apiClient.DtcRecordSrvAPI.DtcrecordsrvGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordSrvAPI.DtcrecordsrvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordsrvGet`: ListDtcRecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordSrvAPI.DtcrecordsrvGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordSrvAPIDtcrecordsrvGetRequest` struct via the builder pattern


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

[**ListDtcRecordSrvResponse**](ListDtcRecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordsrvPost

> CreateDtcRecordSrvResponse DtcrecordsrvPost(ctx).DtcRecordSrv(dtcRecordSrv).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:record:srv object



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
	dtcRecordSrv := *dtc.NewDtcRecordSrv() // DtcRecordSrv | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordSrvAPI.DtcrecordsrvPost(context.Background()).DtcRecordSrv(dtcRecordSrv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordSrvAPI.DtcrecordsrvPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordsrvPost`: CreateDtcRecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordSrvAPI.DtcrecordsrvPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordSrvAPIDtcrecordsrvPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordSrv** | [**DtcRecordSrv**](DtcRecordSrv.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcRecordSrvResponse**](CreateDtcRecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordsrvReferenceDelete

> DtcrecordsrvReferenceDelete(ctx, reference).Execute()

Delete a dtc:record:srv object



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
	reference := "reference_example" // string | Reference of the dtc:record:srv object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcRecordSrvAPI.DtcrecordsrvReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordSrvAPI.DtcrecordsrvReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordSrvAPIDtcrecordsrvReferenceDeleteRequest` struct via the builder pattern


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


## DtcrecordsrvReferenceGet

> GetDtcRecordSrvResponse DtcrecordsrvReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:record:srv object



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
	reference := "reference_example" // string | Reference of the dtc:record:srv object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordSrvAPI.DtcrecordsrvReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordSrvAPI.DtcrecordsrvReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordsrvReferenceGet`: GetDtcRecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordSrvAPI.DtcrecordsrvReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordSrvAPIDtcrecordsrvReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcRecordSrvResponse**](GetDtcRecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordsrvReferencePut

> UpdateDtcRecordSrvResponse DtcrecordsrvReferencePut(ctx, reference).DtcRecordSrv(dtcRecordSrv).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:record:srv object



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
	reference := "reference_example" // string | Reference of the dtc:record:srv object
	dtcRecordSrv := *dtc.NewDtcRecordSrv() // DtcRecordSrv | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordSrvAPI.DtcrecordsrvReferencePut(context.Background(), reference).DtcRecordSrv(dtcRecordSrv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordSrvAPI.DtcrecordsrvReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordsrvReferencePut`: UpdateDtcRecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordSrvAPI.DtcrecordsrvReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordSrvAPIDtcrecordsrvReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordSrv** | [**DtcRecordSrv**](DtcRecordSrv.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcRecordSrvResponse**](UpdateDtcRecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

