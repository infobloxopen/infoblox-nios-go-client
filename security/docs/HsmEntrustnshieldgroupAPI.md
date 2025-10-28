# HsmEntrustnshieldgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](HsmEntrustnshieldgroupAPI.md#Create) | **Post** /hsm:entrustnshieldgroup | Create a hsm:entrustnshieldgroup object
[**Delete**](HsmEntrustnshieldgroupAPI.md#Delete) | **Delete** /hsm:entrustnshieldgroup/{reference} | Delete a hsm:entrustnshieldgroup object
[**List**](HsmEntrustnshieldgroupAPI.md#List) | **Get** /hsm:entrustnshieldgroup | Retrieve hsm:entrustnshieldgroup objects
[**Read**](HsmEntrustnshieldgroupAPI.md#Read) | **Get** /hsm:entrustnshieldgroup/{reference} | Get a specific hsm:entrustnshieldgroup object
[**Update**](HsmEntrustnshieldgroupAPI.md#Update) | **Put** /hsm:entrustnshieldgroup/{reference} | Update a hsm:entrustnshieldgroup object



## Create

> CreateHsmEntrustnshieldgroupResponse Create(ctx).HsmEntrustnshieldgroup(hsmEntrustnshieldgroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a hsm:entrustnshieldgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	hsmEntrustnshieldgroup := *security.NewHsmEntrustnshieldgroup() // HsmEntrustnshieldgroup | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmEntrustnshieldgroupAPI.Create(context.Background()).HsmEntrustnshieldgroup(hsmEntrustnshieldgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateHsmEntrustnshieldgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmEntrustnshieldgroupAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hsmEntrustnshieldgroup** | [**HsmEntrustnshieldgroup**](HsmEntrustnshieldgroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateHsmEntrustnshieldgroupResponse**](CreateHsmEntrustnshieldgroupResponse.md)

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

Delete a hsm:entrustnshieldgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the hsm:entrustnshieldgroup object

	apiClient := security.NewAPIClient()
	r, err := apiClient.HsmEntrustnshieldgroupAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:entrustnshieldgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPIDeleteRequest` struct via the builder pattern


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

> ListHsmEntrustnshieldgroupResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve hsm:entrustnshieldgroup objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmEntrustnshieldgroupAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListHsmEntrustnshieldgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmEntrustnshieldgroupAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPIListRequest` struct via the builder pattern


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

[**ListHsmEntrustnshieldgroupResponse**](ListHsmEntrustnshieldgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetHsmEntrustnshieldgroupResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific hsm:entrustnshieldgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the hsm:entrustnshieldgroup object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmEntrustnshieldgroupAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetHsmEntrustnshieldgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmEntrustnshieldgroupAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:entrustnshieldgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetHsmEntrustnshieldgroupResponse**](GetHsmEntrustnshieldgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateHsmEntrustnshieldgroupResponse Update(ctx, reference).HsmEntrustnshieldgroup(hsmEntrustnshieldgroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a hsm:entrustnshieldgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the hsm:entrustnshieldgroup object
	hsmEntrustnshieldgroup := *security.NewHsmEntrustnshieldgroup() // HsmEntrustnshieldgroup | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmEntrustnshieldgroupAPI.Update(context.Background(), reference).HsmEntrustnshieldgroup(hsmEntrustnshieldgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateHsmEntrustnshieldgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmEntrustnshieldgroupAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:entrustnshieldgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hsmEntrustnshieldgroup** | [**HsmEntrustnshieldgroup**](HsmEntrustnshieldgroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateHsmEntrustnshieldgroupResponse**](UpdateHsmEntrustnshieldgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

