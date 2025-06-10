# ThreatinsightInsightAllowlistAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreatinsightinsightAllowlistGet**](ThreatinsightInsightAllowlistAPI.md#ThreatinsightinsightAllowlistGet) | **Get** /threatinsight:insight_allowlist | Retrieve threatinsight:insight_allowlist objects
[**ThreatinsightinsightAllowlistReferenceGet**](ThreatinsightInsightAllowlistAPI.md#ThreatinsightinsightAllowlistReferenceGet) | **Get** /threatinsight:insight_allowlist/{reference} | Get a specific threatinsight:insight_allowlist object



## ThreatinsightinsightAllowlistGet

> ListThreatinsightInsightAllowlistResponse ThreatinsightinsightAllowlistGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve threatinsight:insight_allowlist objects



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
	resp, r, err := apiClient.ThreatinsightInsightAllowlistAPI.ThreatinsightinsightAllowlistGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightInsightAllowlistAPI.ThreatinsightinsightAllowlistGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightinsightAllowlistGet`: ListThreatinsightInsightAllowlistResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightInsightAllowlistAPI.ThreatinsightinsightAllowlistGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightInsightAllowlistAPIThreatinsightinsightAllowlistGetRequest` struct via the builder pattern


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

[**ListThreatinsightInsightAllowlistResponse**](ListThreatinsightInsightAllowlistResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatinsightinsightAllowlistReferenceGet

> GetThreatinsightInsightAllowlistResponse ThreatinsightinsightAllowlistReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific threatinsight:insight_allowlist object



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
	reference := "reference_example" // string | Reference of the threatinsight:insight_allowlist object

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightInsightAllowlistAPI.ThreatinsightinsightAllowlistReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightInsightAllowlistAPI.ThreatinsightinsightAllowlistReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightinsightAllowlistReferenceGet`: GetThreatinsightInsightAllowlistResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightInsightAllowlistAPI.ThreatinsightinsightAllowlistReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatinsight:insight_allowlist object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightInsightAllowlistAPIThreatinsightinsightAllowlistReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetThreatinsightInsightAllowlistResponse**](GetThreatinsightInsightAllowlistResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

