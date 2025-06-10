# ThreatprotectionRulecategoryAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreatprotectionrulecategoryGet**](ThreatprotectionRulecategoryAPI.md#ThreatprotectionrulecategoryGet) | **Get** /threatprotection:rulecategory | Retrieve threatprotection:rulecategory objects
[**ThreatprotectionrulecategoryReferenceGet**](ThreatprotectionRulecategoryAPI.md#ThreatprotectionrulecategoryReferenceGet) | **Get** /threatprotection:rulecategory/{reference} | Get a specific threatprotection:rulecategory object



## ThreatprotectionrulecategoryGet

> ListThreatprotectionRulecategoryResponse ThreatprotectionrulecategoryGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve threatprotection:rulecategory objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatprotection"
)

func main() {

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionRulecategoryAPI.ThreatprotectionrulecategoryGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionRulecategoryAPI.ThreatprotectionrulecategoryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionrulecategoryGet`: ListThreatprotectionRulecategoryResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionRulecategoryAPI.ThreatprotectionrulecategoryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionRulecategoryAPIThreatprotectionrulecategoryGetRequest` struct via the builder pattern


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

[**ListThreatprotectionRulecategoryResponse**](ListThreatprotectionRulecategoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectionrulecategoryReferenceGet

> GetThreatprotectionRulecategoryResponse ThreatprotectionrulecategoryReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific threatprotection:rulecategory object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatprotection"
)

func main() {
	reference := "reference_example" // string | Reference of the threatprotection:rulecategory object

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionRulecategoryAPI.ThreatprotectionrulecategoryReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionRulecategoryAPI.ThreatprotectionrulecategoryReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionrulecategoryReferenceGet`: GetThreatprotectionRulecategoryResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionRulecategoryAPI.ThreatprotectionrulecategoryReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:rulecategory object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionRulecategoryAPIThreatprotectionrulecategoryReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetThreatprotectionRulecategoryResponse**](GetThreatprotectionRulecategoryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

