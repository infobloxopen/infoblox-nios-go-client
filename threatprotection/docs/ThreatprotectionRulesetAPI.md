# ThreatprotectionRulesetAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreatprotectionrulesetGet**](ThreatprotectionRulesetAPI.md#ThreatprotectionrulesetGet) | **Get** /threatprotection:ruleset | Retrieve threatprotection:ruleset objects
[**ThreatprotectionrulesetReferenceDelete**](ThreatprotectionRulesetAPI.md#ThreatprotectionrulesetReferenceDelete) | **Delete** /threatprotection:ruleset/{reference} | Delete a threatprotection:ruleset object
[**ThreatprotectionrulesetReferenceGet**](ThreatprotectionRulesetAPI.md#ThreatprotectionrulesetReferenceGet) | **Get** /threatprotection:ruleset/{reference} | Get a specific threatprotection:ruleset object
[**ThreatprotectionrulesetReferencePut**](ThreatprotectionRulesetAPI.md#ThreatprotectionrulesetReferencePut) | **Put** /threatprotection:ruleset/{reference} | Update a threatprotection:ruleset object



## ThreatprotectionrulesetGet

> ListThreatprotectionRulesetResponse ThreatprotectionrulesetGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve threatprotection:ruleset objects



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
	resp, r, err := apiClient.ThreatprotectionRulesetAPI.ThreatprotectionrulesetGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionRulesetAPI.ThreatprotectionrulesetGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionrulesetGet`: ListThreatprotectionRulesetResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionRulesetAPI.ThreatprotectionrulesetGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionRulesetAPIThreatprotectionrulesetGetRequest` struct via the builder pattern


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

[**ListThreatprotectionRulesetResponse**](ListThreatprotectionRulesetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectionrulesetReferenceDelete

> ThreatprotectionrulesetReferenceDelete(ctx, reference).Execute()

Delete a threatprotection:ruleset object



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
	reference := "reference_example" // string | Reference of the threatprotection:ruleset object

	apiClient := threatprotection.NewAPIClient()
	r, err := apiClient.ThreatprotectionRulesetAPI.ThreatprotectionrulesetReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionRulesetAPI.ThreatprotectionrulesetReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:ruleset object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionRulesetAPIThreatprotectionrulesetReferenceDeleteRequest` struct via the builder pattern


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


## ThreatprotectionrulesetReferenceGet

> GetThreatprotectionRulesetResponse ThreatprotectionrulesetReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific threatprotection:ruleset object



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
	reference := "reference_example" // string | Reference of the threatprotection:ruleset object

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionRulesetAPI.ThreatprotectionrulesetReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionRulesetAPI.ThreatprotectionrulesetReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionrulesetReferenceGet`: GetThreatprotectionRulesetResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionRulesetAPI.ThreatprotectionrulesetReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:ruleset object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionRulesetAPIThreatprotectionrulesetReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetThreatprotectionRulesetResponse**](GetThreatprotectionRulesetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectionrulesetReferencePut

> UpdateThreatprotectionRulesetResponse ThreatprotectionrulesetReferencePut(ctx, reference).ThreatprotectionRuleset(threatprotectionRuleset).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a threatprotection:ruleset object



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
	reference := "reference_example" // string | Reference of the threatprotection:ruleset object
	threatprotectionRuleset := *threatprotection.NewThreatprotectionRuleset() // ThreatprotectionRuleset | Object data to update

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionRulesetAPI.ThreatprotectionrulesetReferencePut(context.Background(), reference).ThreatprotectionRuleset(threatprotectionRuleset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionRulesetAPI.ThreatprotectionrulesetReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionrulesetReferencePut`: UpdateThreatprotectionRulesetResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionRulesetAPI.ThreatprotectionrulesetReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:ruleset object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionRulesetAPIThreatprotectionrulesetReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionRuleset** | [**ThreatprotectionRuleset**](ThreatprotectionRuleset.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateThreatprotectionRulesetResponse**](UpdateThreatprotectionRulesetResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

