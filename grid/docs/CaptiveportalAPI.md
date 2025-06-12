# CaptiveportalAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](CaptiveportalAPI.md#List) | **Get** /captiveportal | Retrieve captiveportal objects
[**Read**](CaptiveportalAPI.md#Read) | **Get** /captiveportal/{reference} | Get a specific captiveportal object
[**Update**](CaptiveportalAPI.md#Update) | **Put** /captiveportal/{reference} | Update a captiveportal object



## List

> ListCaptiveportalResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve captiveportal objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.CaptiveportalAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptiveportalAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListCaptiveportalResponse
	fmt.Fprintf(os.Stdout, "Response from `CaptiveportalAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CaptiveportalAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**paging** | **int32** | Control paging of results | 
**pageId** | **string** | Page id for retrieving next page of results | 
**filters** | **map[string]interface{}** |  | 
**extattrfilter** | **map[string]interface{}** |  | 

### Return type

[**ListCaptiveportalResponse**](ListCaptiveportalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetCaptiveportalResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Get a specific captiveportal object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {
	reference := "reference_example" // string | Reference of the captiveportal object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.CaptiveportalAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptiveportalAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetCaptiveportalResponse
	fmt.Fprintf(os.Stdout, "Response from `CaptiveportalAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the captiveportal object | 

### Other Parameters

Other parameters are passed through a pointer to a `CaptiveportalAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetCaptiveportalResponse**](GetCaptiveportalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateCaptiveportalResponse Update(ctx, reference).Captiveportal(captiveportal).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a captiveportal object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {
	reference := "reference_example" // string | Reference of the captiveportal object
	captiveportal := *grid.NewCaptiveportal() // Captiveportal | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.CaptiveportalAPI.Update(context.Background(), reference).Captiveportal(captiveportal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CaptiveportalAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateCaptiveportalResponse
	fmt.Fprintf(os.Stdout, "Response from `CaptiveportalAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the captiveportal object | 

### Other Parameters

Other parameters are passed through a pointer to a `CaptiveportalAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**captiveportal** | [**Captiveportal**](Captiveportal.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateCaptiveportalResponse**](UpdateCaptiveportalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

