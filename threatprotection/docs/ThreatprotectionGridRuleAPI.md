# ThreatprotectionGridRuleAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ThreatprotectionGridRuleAPI.md#Create) | **Post** /threatprotection:grid:rule | Create a threatprotection:grid:rule object
[**Delete**](ThreatprotectionGridRuleAPI.md#Delete) | **Delete** /threatprotection:grid:rule/{reference} | Delete a threatprotection:grid:rule object
[**List**](ThreatprotectionGridRuleAPI.md#List) | **Get** /threatprotection:grid:rule | Retrieve threatprotection:grid:rule objects
[**Read**](ThreatprotectionGridRuleAPI.md#Read) | **Get** /threatprotection:grid:rule/{reference} | Get a specific threatprotection:grid:rule object
[**Update**](ThreatprotectionGridRuleAPI.md#Update) | **Put** /threatprotection:grid:rule/{reference} | Update a threatprotection:grid:rule object



## Create

> CreateThreatprotectionGridRuleResponse Create(ctx).ThreatprotectionGridRule(threatprotectionGridRule).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a threatprotection:grid:rule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
)

func main() {
	threatprotectionGridRule := *threatprotection.NewThreatprotectionGridRule() // ThreatprotectionGridRule | Object data to create

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionGridRuleAPI.Create(context.Background()).ThreatprotectionGridRule(threatprotectionGridRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateThreatprotectionGridRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionGridRuleAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionGridRule** | [**ThreatprotectionGridRule**](ThreatprotectionGridRule.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## Delete

> Delete(ctx, reference).Execute()

Delete a threatprotection:grid:rule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
)

func main() {
	reference := "reference_example" // string | Reference of the threatprotection:grid:rule object

	apiClient := threatprotection.NewAPIClient()
	r, err := apiClient.ThreatprotectionGridRuleAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPIDeleteRequest` struct via the builder pattern


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


## List

> ListThreatprotectionGridRuleResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve threatprotection:grid:rule objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
)

func main() {

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionGridRuleAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListThreatprotectionGridRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionGridRuleAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPIListRequest` struct via the builder pattern


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
**proxySearch** | **string** | Search Grid members for data | 

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


## Read

> GetThreatprotectionGridRuleResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific threatprotection:grid:rule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
)

func main() {
	reference := "reference_example" // string | Reference of the threatprotection:grid:rule object

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionGridRuleAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetThreatprotectionGridRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionGridRuleAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:grid:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

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


## Update

> UpdateThreatprotectionGridRuleResponse Update(ctx, reference).ThreatprotectionGridRule(threatprotectionGridRule).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a threatprotection:grid:rule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
)

func main() {
	reference := "reference_example" // string | Reference of the threatprotection:grid:rule object
	threatprotectionGridRule := *threatprotection.NewThreatprotectionGridRule() // ThreatprotectionGridRule | Object data to update

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionGridRuleAPI.Update(context.Background(), reference).ThreatprotectionGridRule(threatprotectionGridRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionGridRuleAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateThreatprotectionGridRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionGridRuleAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:grid:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionGridRuleAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionGridRule** | [**ThreatprotectionGridRule**](ThreatprotectionGridRule.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

