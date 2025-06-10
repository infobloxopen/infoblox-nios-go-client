# ThreatprotectionProfileRuleAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreatprotectionprofileruleGet**](ThreatprotectionProfileRuleAPI.md#ThreatprotectionprofileruleGet) | **Get** /threatprotection:profile:rule | Retrieve threatprotection:profile:rule objects
[**ThreatprotectionprofileruleReferenceGet**](ThreatprotectionProfileRuleAPI.md#ThreatprotectionprofileruleReferenceGet) | **Get** /threatprotection:profile:rule/{reference} | Get a specific threatprotection:profile:rule object
[**ThreatprotectionprofileruleReferencePut**](ThreatprotectionProfileRuleAPI.md#ThreatprotectionprofileruleReferencePut) | **Put** /threatprotection:profile:rule/{reference} | Update a threatprotection:profile:rule object



## ThreatprotectionprofileruleGet

> ListThreatprotectionProfileRuleResponse ThreatprotectionprofileruleGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve threatprotection:profile:rule objects



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
	resp, r, err := apiClient.ThreatprotectionProfileRuleAPI.ThreatprotectionprofileruleGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileRuleAPI.ThreatprotectionprofileruleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionprofileruleGet`: ListThreatprotectionProfileRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileRuleAPI.ThreatprotectionprofileruleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileRuleAPIThreatprotectionprofileruleGetRequest` struct via the builder pattern


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

[**ListThreatprotectionProfileRuleResponse**](ListThreatprotectionProfileRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectionprofileruleReferenceGet

> GetThreatprotectionProfileRuleResponse ThreatprotectionprofileruleReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific threatprotection:profile:rule object



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
	reference := "reference_example" // string | Reference of the threatprotection:profile:rule object

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionProfileRuleAPI.ThreatprotectionprofileruleReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileRuleAPI.ThreatprotectionprofileruleReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionprofileruleReferenceGet`: GetThreatprotectionProfileRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileRuleAPI.ThreatprotectionprofileruleReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:profile:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileRuleAPIThreatprotectionprofileruleReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetThreatprotectionProfileRuleResponse**](GetThreatprotectionProfileRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectionprofileruleReferencePut

> UpdateThreatprotectionProfileRuleResponse ThreatprotectionprofileruleReferencePut(ctx, reference).ThreatprotectionProfileRule(threatprotectionProfileRule).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a threatprotection:profile:rule object



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
	reference := "reference_example" // string | Reference of the threatprotection:profile:rule object
	threatprotectionProfileRule := *threatprotection.NewThreatprotectionProfileRule() // ThreatprotectionProfileRule | Object data to update

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionProfileRuleAPI.ThreatprotectionprofileruleReferencePut(context.Background(), reference).ThreatprotectionProfileRule(threatprotectionProfileRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileRuleAPI.ThreatprotectionprofileruleReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionprofileruleReferencePut`: UpdateThreatprotectionProfileRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileRuleAPI.ThreatprotectionprofileruleReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:profile:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileRuleAPIThreatprotectionprofileruleReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionProfileRule** | [**ThreatprotectionProfileRule**](ThreatprotectionProfileRule.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateThreatprotectionProfileRuleResponse**](UpdateThreatprotectionProfileRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

