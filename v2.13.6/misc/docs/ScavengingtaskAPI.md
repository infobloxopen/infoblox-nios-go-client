# ScavengingtaskAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](ScavengingtaskAPI.md#List) | **Get** /scavengingtask | Retrieve scavengingtask objects
[**Read**](ScavengingtaskAPI.md#Read) | **Get** /scavengingtask/{reference} | Get a specific scavengingtask object



## List

> ListScavengingtaskResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve scavengingtask objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
)

func main() {

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.ScavengingtaskAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScavengingtaskAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListScavengingtaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ScavengingtaskAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ScavengingtaskAPIListRequest` struct via the builder pattern


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

[**ListScavengingtaskResponse**](ListScavengingtaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetScavengingtaskResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Get a specific scavengingtask object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the scavengingtask object

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.ScavengingtaskAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScavengingtaskAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetScavengingtaskResponse
	fmt.Fprintf(os.Stdout, "Response from `ScavengingtaskAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the scavengingtask object | 

### Other Parameters

Other parameters are passed through a pointer to a `ScavengingtaskAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetScavengingtaskResponse**](GetScavengingtaskResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

