# BulkhostnametemplateAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](BulkhostnametemplateAPI.md#Create) | **Post** /bulkhostnametemplate | Create a bulkhostnametemplate object
[**Delete**](BulkhostnametemplateAPI.md#Delete) | **Delete** /bulkhostnametemplate/{reference} | Delete a bulkhostnametemplate object
[**List**](BulkhostnametemplateAPI.md#List) | **Get** /bulkhostnametemplate | Retrieve bulkhostnametemplate objects
[**Read**](BulkhostnametemplateAPI.md#Read) | **Get** /bulkhostnametemplate/{reference} | Get a specific bulkhostnametemplate object
[**Update**](BulkhostnametemplateAPI.md#Update) | **Put** /bulkhostnametemplate/{reference} | Update a bulkhostnametemplate object



## Create

> CreateBulkhostnametemplateResponse Create(ctx).Bulkhostnametemplate(bulkhostnametemplate).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a bulkhostnametemplate object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {
	bulkhostnametemplate := *ipam.NewBulkhostnametemplate() // Bulkhostnametemplate | Object data to create

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.BulkhostnametemplateAPI.Create(context.Background()).Bulkhostnametemplate(bulkhostnametemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkhostnametemplateAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateBulkhostnametemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkhostnametemplateAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `BulkhostnametemplateAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**bulkhostnametemplate** | [**Bulkhostnametemplate**](Bulkhostnametemplate.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateBulkhostnametemplateResponse**](CreateBulkhostnametemplateResponse.md)

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

Delete a bulkhostnametemplate object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the bulkhostnametemplate object

	apiClient := ipam.NewAPIClient()
	r, err := apiClient.BulkhostnametemplateAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkhostnametemplateAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the bulkhostnametemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `BulkhostnametemplateAPIDeleteRequest` struct via the builder pattern


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

> ListBulkhostnametemplateResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve bulkhostnametemplate objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.BulkhostnametemplateAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkhostnametemplateAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListBulkhostnametemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkhostnametemplateAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `BulkhostnametemplateAPIListRequest` struct via the builder pattern


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

[**ListBulkhostnametemplateResponse**](ListBulkhostnametemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetBulkhostnametemplateResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific bulkhostnametemplate object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the bulkhostnametemplate object

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.BulkhostnametemplateAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkhostnametemplateAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetBulkhostnametemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkhostnametemplateAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the bulkhostnametemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `BulkhostnametemplateAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetBulkhostnametemplateResponse**](GetBulkhostnametemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateBulkhostnametemplateResponse Update(ctx, reference).Bulkhostnametemplate(bulkhostnametemplate).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a bulkhostnametemplate object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

func main() {
	reference := "reference_example" // string | Reference of the bulkhostnametemplate object
	bulkhostnametemplate := *ipam.NewBulkhostnametemplate() // Bulkhostnametemplate | Object data to update

	apiClient := ipam.NewAPIClient()
	resp, r, err := apiClient.BulkhostnametemplateAPI.Update(context.Background(), reference).Bulkhostnametemplate(bulkhostnametemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkhostnametemplateAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateBulkhostnametemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `BulkhostnametemplateAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the bulkhostnametemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `BulkhostnametemplateAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**bulkhostnametemplate** | [**Bulkhostnametemplate**](Bulkhostnametemplate.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateBulkhostnametemplateResponse**](UpdateBulkhostnametemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

