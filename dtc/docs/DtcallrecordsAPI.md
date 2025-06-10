# DtcAllrecordsAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcallrecordsGet**](DtcAllrecordsAPI.md#DtcallrecordsGet) | **Get** /dtc:allrecords | Retrieve dtc:allrecords objects
[**DtcallrecordsReferenceGet**](DtcAllrecordsAPI.md#DtcallrecordsReferenceGet) | **Get** /dtc:allrecords/{reference} | Get a specific dtc:allrecords object
[**DtcallrecordsReferencePut**](DtcAllrecordsAPI.md#DtcallrecordsReferencePut) | **Put** /dtc:allrecords/{reference} | Update a dtc:allrecords object



## DtcallrecordsGet

> ListDtcAllrecordsResponse DtcallrecordsGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:allrecords objects



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
	resp, r, err := apiClient.DtcAllrecordsAPI.DtcallrecordsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcAllrecordsAPI.DtcallrecordsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcallrecordsGet`: ListDtcAllrecordsResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcAllrecordsAPI.DtcallrecordsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcAllrecordsAPIDtcallrecordsGetRequest` struct via the builder pattern


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

[**ListDtcAllrecordsResponse**](ListDtcAllrecordsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcallrecordsReferenceGet

> GetDtcAllrecordsResponse DtcallrecordsReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:allrecords object



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
	reference := "reference_example" // string | Reference of the dtc:allrecords object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcAllrecordsAPI.DtcallrecordsReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcAllrecordsAPI.DtcallrecordsReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcallrecordsReferenceGet`: GetDtcAllrecordsResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcAllrecordsAPI.DtcallrecordsReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:allrecords object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcAllrecordsAPIDtcallrecordsReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcAllrecordsResponse**](GetDtcAllrecordsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcallrecordsReferencePut

> UpdateDtcAllrecordsResponse DtcallrecordsReferencePut(ctx, reference).DtcAllrecords(dtcAllrecords).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:allrecords object



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
	reference := "reference_example" // string | Reference of the dtc:allrecords object
	dtcAllrecords := *dtc.NewDtcAllrecords() // DtcAllrecords | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcAllrecordsAPI.DtcallrecordsReferencePut(context.Background(), reference).DtcAllrecords(dtcAllrecords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcAllrecordsAPI.DtcallrecordsReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcallrecordsReferencePut`: UpdateDtcAllrecordsResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcAllrecordsAPI.DtcallrecordsReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:allrecords object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcAllrecordsAPIDtcallrecordsReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcAllrecords** | [**DtcAllrecords**](DtcAllrecords.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcAllrecordsResponse**](UpdateDtcAllrecordsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

