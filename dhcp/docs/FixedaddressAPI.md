# FixedaddressAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](FixedaddressAPI.md#Create) | **Post** /fixedaddress | Create a fixedaddress object
[**Delete**](FixedaddressAPI.md#Delete) | **Delete** /fixedaddress/{reference} | Delete a fixedaddress object
[**List**](FixedaddressAPI.md#List) | **Get** /fixedaddress | Retrieve fixedaddress objects
[**MsServerUpdate**](FixedaddressAPI.md#MsServerUpdate) | **Put** /fixedaddress | Update a fixedaddress object
[**Read**](FixedaddressAPI.md#Read) | **Get** /fixedaddress/{reference} | Get a specific fixedaddress object
[**Update**](FixedaddressAPI.md#Update) | **Put** /fixedaddress/{reference} | Update a fixedaddress object



## Create

> CreateFixedaddressResponse Create(ctx).Fixedaddress(fixedaddress).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a fixedaddress object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {
	fixedaddress := *dhcp.NewFixedaddress() // Fixedaddress | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.FixedaddressAPI.Create(context.Background()).Fixedaddress(fixedaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddressAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateFixedaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `FixedaddressAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `FixedaddressAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fixedaddress** | [**Fixedaddress**](Fixedaddress.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateFixedaddressResponse**](CreateFixedaddressResponse.md)

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

Delete a fixedaddress object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the fixedaddress object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.FixedaddressAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddressAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the fixedaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `FixedaddressAPIDeleteRequest` struct via the builder pattern


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

> ListFixedaddressResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve fixedaddress objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.FixedaddressAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddressAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListFixedaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `FixedaddressAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `FixedaddressAPIListRequest` struct via the builder pattern


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

[**ListFixedaddressResponse**](ListFixedaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsServerUpdate

> ListFixedaddressResponse MsServerUpdate(ctx).FixedAddressStruct(fixedAddressStruct).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Method(method).Paging(paging).PageId(pageId).MaxResults(maxResults).Execute()

Update a fixedaddress object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {
	fixedAddressStruct := *dhcp.NewFixedAddressStruct(*dhcp.NewFixedAddressStructMsServer()) // FixedAddressStruct | Update Call for Search by Microsoft Server

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.FixedaddressAPI.MsServerUpdate(context.Background()).FixedAddressStruct(fixedAddressStruct).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddressAPI.MsServerUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsServerUpdate`: ListFixedaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `FixedaddressAPI.MsServerUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `FixedaddressAPIMsServerUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fixedAddressStruct** | [**FixedAddressStruct**](FixedAddressStruct.md) | Update Call for Search by Microsoft Server | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**method** | **string** | Enter the method type for the request | 
**paging** | **int32** | Control paging of results | 
**pageId** | **string** | Page id for retrieving next page of results | 
**maxResults** | **int32** | Enter the number of results to be fetched | 

### Return type

[**ListFixedaddressResponse**](ListFixedaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetFixedaddressResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific fixedaddress object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the fixedaddress object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.FixedaddressAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddressAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetFixedaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `FixedaddressAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the fixedaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `FixedaddressAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetFixedaddressResponse**](GetFixedaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateFixedaddressResponse Update(ctx, reference).Fixedaddress(fixedaddress).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a fixedaddress object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the fixedaddress object
	fixedaddress := *dhcp.NewFixedaddress() // Fixedaddress | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.FixedaddressAPI.Update(context.Background(), reference).Fixedaddress(fixedaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddressAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateFixedaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `FixedaddressAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the fixedaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `FixedaddressAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fixedaddress** | [**Fixedaddress**](Fixedaddress.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateFixedaddressResponse**](UpdateFixedaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

