# DtcRecordCnameAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcrecordcnameGet**](DtcRecordCnameAPI.md#DtcrecordcnameGet) | **Get** /dtc:record:cname | Retrieve dtc:record:cname objects
[**DtcrecordcnamePost**](DtcRecordCnameAPI.md#DtcrecordcnamePost) | **Post** /dtc:record:cname | Create a dtc:record:cname object
[**DtcrecordcnameReferenceDelete**](DtcRecordCnameAPI.md#DtcrecordcnameReferenceDelete) | **Delete** /dtc:record:cname/{reference} | Delete a dtc:record:cname object
[**DtcrecordcnameReferenceGet**](DtcRecordCnameAPI.md#DtcrecordcnameReferenceGet) | **Get** /dtc:record:cname/{reference} | Get a specific dtc:record:cname object
[**DtcrecordcnameReferencePut**](DtcRecordCnameAPI.md#DtcrecordcnameReferencePut) | **Put** /dtc:record:cname/{reference} | Update a dtc:record:cname object



## DtcrecordcnameGet

> ListDtcRecordCnameResponse DtcrecordcnameGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:record:cname objects



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
	resp, r, err := apiClient.DtcRecordCnameAPI.DtcrecordcnameGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordCnameAPI.DtcrecordcnameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordcnameGet`: ListDtcRecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordCnameAPI.DtcrecordcnameGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordCnameAPIDtcrecordcnameGetRequest` struct via the builder pattern


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

[**ListDtcRecordCnameResponse**](ListDtcRecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordcnamePost

> CreateDtcRecordCnameResponse DtcrecordcnamePost(ctx).DtcRecordCname(dtcRecordCname).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:record:cname object



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
	dtcRecordCname := *dtc.NewDtcRecordCname() // DtcRecordCname | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordCnameAPI.DtcrecordcnamePost(context.Background()).DtcRecordCname(dtcRecordCname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordCnameAPI.DtcrecordcnamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordcnamePost`: CreateDtcRecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordCnameAPI.DtcrecordcnamePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordCnameAPIDtcrecordcnamePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordCname** | [**DtcRecordCname**](DtcRecordCname.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcRecordCnameResponse**](CreateDtcRecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordcnameReferenceDelete

> DtcrecordcnameReferenceDelete(ctx, reference).Execute()

Delete a dtc:record:cname object



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
	reference := "reference_example" // string | Reference of the dtc:record:cname object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcRecordCnameAPI.DtcrecordcnameReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordCnameAPI.DtcrecordcnameReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:cname object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordCnameAPIDtcrecordcnameReferenceDeleteRequest` struct via the builder pattern


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


## DtcrecordcnameReferenceGet

> GetDtcRecordCnameResponse DtcrecordcnameReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:record:cname object



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
	reference := "reference_example" // string | Reference of the dtc:record:cname object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordCnameAPI.DtcrecordcnameReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordCnameAPI.DtcrecordcnameReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordcnameReferenceGet`: GetDtcRecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordCnameAPI.DtcrecordcnameReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:cname object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordCnameAPIDtcrecordcnameReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcRecordCnameResponse**](GetDtcRecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordcnameReferencePut

> UpdateDtcRecordCnameResponse DtcrecordcnameReferencePut(ctx, reference).DtcRecordCname(dtcRecordCname).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:record:cname object



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
	reference := "reference_example" // string | Reference of the dtc:record:cname object
	dtcRecordCname := *dtc.NewDtcRecordCname() // DtcRecordCname | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordCnameAPI.DtcrecordcnameReferencePut(context.Background(), reference).DtcRecordCname(dtcRecordCname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordCnameAPI.DtcrecordcnameReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordcnameReferencePut`: UpdateDtcRecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordCnameAPI.DtcrecordcnameReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:cname object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordCnameAPIDtcrecordcnameReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordCname** | [**DtcRecordCname**](DtcRecordCname.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcRecordCnameResponse**](UpdateDtcRecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

