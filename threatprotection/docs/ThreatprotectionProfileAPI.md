# ThreatprotectionProfileAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ThreatprotectionProfileAPI.md#Create) | **Post** /threatprotection:profile | Create a threatprotection:profile object
[**Delete**](ThreatprotectionProfileAPI.md#Delete) | **Delete** /threatprotection:profile/{reference} | Delete a threatprotection:profile object
[**List**](ThreatprotectionProfileAPI.md#List) | **Get** /threatprotection:profile | Retrieve threatprotection:profile objects
[**Read**](ThreatprotectionProfileAPI.md#Read) | **Get** /threatprotection:profile/{reference} | Get a specific threatprotection:profile object
[**Update**](ThreatprotectionProfileAPI.md#Update) | **Put** /threatprotection:profile/{reference} | Update a threatprotection:profile object



## Create

> CreateThreatprotectionProfileResponse Create(ctx).ThreatprotectionProfile(threatprotectionProfile).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a threatprotection:profile object



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
	threatprotectionProfile := *threatprotection.NewThreatprotectionProfile() // ThreatprotectionProfile | Object data to create

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionProfileAPI.Create(context.Background()).ThreatprotectionProfile(threatprotectionProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateThreatprotectionProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionProfile** | [**ThreatprotectionProfile**](ThreatprotectionProfile.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateThreatprotectionProfileResponse**](CreateThreatprotectionProfileResponse.md)

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

Delete a threatprotection:profile object



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
	reference := "reference_example" // string | Reference of the threatprotection:profile object

	apiClient := threatprotection.NewAPIClient()
	r, err := apiClient.ThreatprotectionProfileAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:profile object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPIDeleteRequest` struct via the builder pattern


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

> ListThreatprotectionProfileResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve threatprotection:profile objects



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
	resp, r, err := apiClient.ThreatprotectionProfileAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListThreatprotectionProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPIListRequest` struct via the builder pattern


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

[**ListThreatprotectionProfileResponse**](ListThreatprotectionProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetThreatprotectionProfileResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific threatprotection:profile object



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
	reference := "reference_example" // string | Reference of the threatprotection:profile object

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionProfileAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetThreatprotectionProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:profile object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetThreatprotectionProfileResponse**](GetThreatprotectionProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateThreatprotectionProfileResponse Update(ctx, reference).ThreatprotectionProfile(threatprotectionProfile).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a threatprotection:profile object



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
	reference := "reference_example" // string | Reference of the threatprotection:profile object
	threatprotectionProfile := *threatprotection.NewThreatprotectionProfile() // ThreatprotectionProfile | Object data to update

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionProfileAPI.Update(context.Background(), reference).ThreatprotectionProfile(threatprotectionProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateThreatprotectionProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:profile object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionProfile** | [**ThreatprotectionProfile**](ThreatprotectionProfile.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateThreatprotectionProfileResponse**](UpdateThreatprotectionProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

