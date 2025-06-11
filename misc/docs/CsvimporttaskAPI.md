# CsvimporttaskAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](CsvimporttaskAPI.md#Get) | **Get** /csvimporttask | Retrieve csvimporttask objects
[**ReferenceGet**](CsvimporttaskAPI.md#ReferenceGet) | **Get** /csvimporttask/{reference} | Get a specific csvimporttask object
[**ReferencePut**](CsvimporttaskAPI.md#ReferencePut) | **Put** /csvimporttask/{reference} | Update a csvimporttask object



## Get

> ListCsvimporttaskResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve csvimporttask objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.CsvimporttaskAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CsvimporttaskAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListCsvimporttaskResponse
	fmt.Fprintf(os.Stdout, "Response from `CsvimporttaskAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CsvimporttaskAPIGetRequest` struct via the builder pattern


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

[**ListCsvimporttaskResponse**](ListCsvimporttaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceGet

> GetCsvimporttaskResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific csvimporttask object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the csvimporttask object

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.CsvimporttaskAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CsvimporttaskAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetCsvimporttaskResponse
	fmt.Fprintf(os.Stdout, "Response from `CsvimporttaskAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the csvimporttask object | 

### Other Parameters

Other parameters are passed through a pointer to a `CsvimporttaskAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetCsvimporttaskResponse**](GetCsvimporttaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateCsvimporttaskResponse ReferencePut(ctx, reference).Csvimporttask(csvimporttask).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a csvimporttask object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the csvimporttask object
	csvimporttask := *misc.NewCsvimporttask() // Csvimporttask | Object data to update

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.CsvimporttaskAPI.ReferencePut(context.Background(), reference).Csvimporttask(csvimporttask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CsvimporttaskAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateCsvimporttaskResponse
	fmt.Fprintf(os.Stdout, "Response from `CsvimporttaskAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the csvimporttask object | 

### Other Parameters

Other parameters are passed through a pointer to a `CsvimporttaskAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**csvimporttask** | [**Csvimporttask**](Csvimporttask.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateCsvimporttaskResponse**](UpdateCsvimporttaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

