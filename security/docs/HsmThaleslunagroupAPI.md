# HsmThaleslunagroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](HsmThaleslunagroupAPI.md#Create) | **Post** /hsm:thaleslunagroup | Create a hsm:thaleslunagroup object
[**Delete**](HsmThaleslunagroupAPI.md#Delete) | **Delete** /hsm:thaleslunagroup/{reference} | Delete a hsm:thaleslunagroup object
[**List**](HsmThaleslunagroupAPI.md#List) | **Get** /hsm:thaleslunagroup | Retrieve hsm:thaleslunagroup objects
[**Read**](HsmThaleslunagroupAPI.md#Read) | **Get** /hsm:thaleslunagroup/{reference} | Get a specific hsm:thaleslunagroup object
[**Update**](HsmThaleslunagroupAPI.md#Update) | **Put** /hsm:thaleslunagroup/{reference} | Update a hsm:thaleslunagroup object



## Create

> CreateHsmThaleslunagroupResponse Create(ctx).HsmThaleslunagroup(hsmThaleslunagroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a hsm:thaleslunagroup object



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
	hsmThaleslunagroup := *security.NewHsmThaleslunagroup() // HsmThaleslunagroup | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmThaleslunagroupAPI.Create(context.Background()).HsmThaleslunagroup(hsmThaleslunagroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateHsmThaleslunagroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmThaleslunagroupAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hsmThaleslunagroup** | [**HsmThaleslunagroup**](HsmThaleslunagroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateHsmThaleslunagroupResponse**](CreateHsmThaleslunagroupResponse.md)

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

Delete a hsm:thaleslunagroup object



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
	reference := "reference_example" // string | Reference of the hsm:thaleslunagroup object

	apiClient := security.NewAPIClient()
	r, err := apiClient.HsmThaleslunagroupAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:thaleslunagroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPIDeleteRequest` struct via the builder pattern


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

> ListHsmThaleslunagroupResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve hsm:thaleslunagroup objects



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
	resp, r, err := apiClient.HsmThaleslunagroupAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListHsmThaleslunagroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmThaleslunagroupAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPIListRequest` struct via the builder pattern


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

[**ListHsmThaleslunagroupResponse**](ListHsmThaleslunagroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetHsmThaleslunagroupResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific hsm:thaleslunagroup object



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
	reference := "reference_example" // string | Reference of the hsm:thaleslunagroup object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmThaleslunagroupAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetHsmThaleslunagroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmThaleslunagroupAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:thaleslunagroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetHsmThaleslunagroupResponse**](GetHsmThaleslunagroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateHsmThaleslunagroupResponse Update(ctx, reference).HsmThaleslunagroup(hsmThaleslunagroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a hsm:thaleslunagroup object



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
	reference := "reference_example" // string | Reference of the hsm:thaleslunagroup object
	hsmThaleslunagroup := *security.NewHsmThaleslunagroup() // HsmThaleslunagroup | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmThaleslunagroupAPI.Update(context.Background(), reference).HsmThaleslunagroup(hsmThaleslunagroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateHsmThaleslunagroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmThaleslunagroupAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:thaleslunagroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hsmThaleslunagroup** | [**HsmThaleslunagroup**](HsmThaleslunagroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateHsmThaleslunagroupResponse**](UpdateHsmThaleslunagroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

