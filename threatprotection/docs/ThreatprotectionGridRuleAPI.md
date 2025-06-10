# ThreatprotectionGridRuleAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreatprotectiongridruleGet**](ThreatprotectionGridRuleAPI.md#ThreatprotectiongridruleGet) | **Get** /threatprotection:grid:rule | Retrieve threatprotection:grid:rule objects
[**ThreatprotectiongridrulePost**](ThreatprotectionGridRuleAPI.md#ThreatprotectiongridrulePost) | **Post** /threatprotection:grid:rule | Create a threatprotection:grid:rule object
[**ThreatprotectiongridruleReferenceDelete**](ThreatprotectionGridRuleAPI.md#ThreatprotectiongridruleReferenceDelete) | **Delete** /threatprotection:grid:rule/{reference} | Delete a threatprotection:grid:rule object
[**ThreatprotectiongridruleReferenceGet**](ThreatprotectionGridRuleAPI.md#ThreatprotectiongridruleReferenceGet) | **Get** /threatprotection:grid:rule/{reference} | Get a specific threatprotection:grid:rule object
[**ThreatprotectiongridruleReferencePut**](ThreatprotectionGridRuleAPI.md#ThreatprotectiongridruleReferencePut) | **Put** /threatprotection:grid:rule/{reference} | Update a threatprotection:grid:rule object



## ThreatprotectiongridruleGet

> ListThreatprotectionGridRuleResponse ThreatprotectiongridruleGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve threatprotection:grid:rule objects



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
	resp, r, err := apiClient.ThreatprotectionGridRuleAPI.ThreatprotectiongridruleGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.ThreatprotectiongridruleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectiongridruleGet`: ListThreatprotectionGridRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionGridRuleAPI.ThreatprotectiongridruleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPIThreatprotectiongridruleGetRequest` struct via the builder pattern


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

[**ListThreatprotectionGridRuleResponse**](ListThreatprotectionGridRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectiongridrulePost

> CreateThreatprotectionGridRuleResponse ThreatprotectiongridrulePost(ctx).ThreatprotectionGridRule(threatprotectionGridRule).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a threatprotection:grid:rule object



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
	threatprotectionGridRule := *threatprotection.NewThreatprotectionGridRule() // ThreatprotectionGridRule | Object data to create

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionGridRuleAPI.ThreatprotectiongridrulePost(context.Background()).ThreatprotectionGridRule(threatprotectionGridRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.ThreatprotectiongridrulePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectiongridrulePost`: CreateThreatprotectionGridRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionGridRuleAPI.ThreatprotectiongridrulePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPIThreatprotectiongridrulePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionGridRule** | [**ThreatprotectionGridRule**](ThreatprotectionGridRule.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateThreatprotectionGridRuleResponse**](CreateThreatprotectionGridRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectiongridruleReferenceDelete

> ThreatprotectiongridruleReferenceDelete(ctx, reference).Execute()

Delete a threatprotection:grid:rule object



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
	reference := "reference_example" // string | Reference of the threatprotection:grid:rule object

	apiClient := threatprotection.NewAPIClient()
	r, err := apiClient.ThreatprotectionGridRuleAPI.ThreatprotectiongridruleReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.ThreatprotectiongridruleReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:grid:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPIThreatprotectiongridruleReferenceDeleteRequest` struct via the builder pattern


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


## ThreatprotectiongridruleReferenceGet

> GetThreatprotectionGridRuleResponse ThreatprotectiongridruleReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific threatprotection:grid:rule object



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
	reference := "reference_example" // string | Reference of the threatprotection:grid:rule object

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionGridRuleAPI.ThreatprotectiongridruleReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.ThreatprotectiongridruleReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectiongridruleReferenceGet`: GetThreatprotectionGridRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionGridRuleAPI.ThreatprotectiongridruleReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:grid:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPIThreatprotectiongridruleReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetThreatprotectionGridRuleResponse**](GetThreatprotectionGridRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectiongridruleReferencePut

> UpdateThreatprotectionGridRuleResponse ThreatprotectiongridruleReferencePut(ctx, reference).ThreatprotectionGridRule(threatprotectionGridRule).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a threatprotection:grid:rule object



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
	reference := "reference_example" // string | Reference of the threatprotection:grid:rule object
	threatprotectionGridRule := *threatprotection.NewThreatprotectionGridRule() // ThreatprotectionGridRule | Object data to update

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionGridRuleAPI.ThreatprotectiongridruleReferencePut(context.Background(), reference).ThreatprotectionGridRule(threatprotectionGridRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.ThreatprotectiongridruleReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectiongridruleReferencePut`: UpdateThreatprotectionGridRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionGridRuleAPI.ThreatprotectiongridruleReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:grid:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPIThreatprotectiongridruleReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionGridRule** | [**ThreatprotectionGridRule**](ThreatprotectionGridRule.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateThreatprotectionGridRuleResponse**](UpdateThreatprotectionGridRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

