# ThreatinsightModulesetAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreatinsightmodulesetGet**](ThreatinsightModulesetAPI.md#ThreatinsightmodulesetGet) | **Get** /threatinsight:moduleset | Retrieve threatinsight:moduleset objects
[**ThreatinsightmodulesetReferenceGet**](ThreatinsightModulesetAPI.md#ThreatinsightmodulesetReferenceGet) | **Get** /threatinsight:moduleset/{reference} | Get a specific threatinsight:moduleset object



## ThreatinsightmodulesetGet

> ListThreatinsightModulesetResponse ThreatinsightmodulesetGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve threatinsight:moduleset objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightModulesetAPI.ThreatinsightmodulesetGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightModulesetAPI.ThreatinsightmodulesetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightmodulesetGet`: ListThreatinsightModulesetResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightModulesetAPI.ThreatinsightmodulesetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightModulesetAPIThreatinsightmodulesetGetRequest` struct via the builder pattern


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

[**ListThreatinsightModulesetResponse**](ListThreatinsightModulesetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatinsightmodulesetReferenceGet

> GetThreatinsightModulesetResponse ThreatinsightmodulesetReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific threatinsight:moduleset object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {
	reference := "reference_example" // string | Reference of the threatinsight:moduleset object

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightModulesetAPI.ThreatinsightmodulesetReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightModulesetAPI.ThreatinsightmodulesetReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightmodulesetReferenceGet`: GetThreatinsightModulesetResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightModulesetAPI.ThreatinsightmodulesetReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatinsight:moduleset object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightModulesetAPIThreatinsightmodulesetReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetThreatinsightModulesetResponse**](GetThreatinsightModulesetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

