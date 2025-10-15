# FixedaddresstemplateAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](FixedaddresstemplateAPI.md#Create) | **Post** /fixedaddresstemplate | Create a fixedaddresstemplate object
[**Delete**](FixedaddresstemplateAPI.md#Delete) | **Delete** /fixedaddresstemplate/{reference} | Delete a fixedaddresstemplate object
[**List**](FixedaddresstemplateAPI.md#List) | **Get** /fixedaddresstemplate | Retrieve fixedaddresstemplate objects
[**Read**](FixedaddresstemplateAPI.md#Read) | **Get** /fixedaddresstemplate/{reference} | Get a specific fixedaddresstemplate object
[**Update**](FixedaddresstemplateAPI.md#Update) | **Put** /fixedaddresstemplate/{reference} | Update a fixedaddresstemplate object



## Create

> CreateFixedaddresstemplateResponse Create(ctx).Fixedaddresstemplate(fixedaddresstemplate).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a fixedaddresstemplate object



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
	fixedaddresstemplate := *dhcp.NewFixedaddresstemplate() // Fixedaddresstemplate | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.FixedaddresstemplateAPI.Create(context.Background()).Fixedaddresstemplate(fixedaddresstemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddresstemplateAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateFixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `FixedaddresstemplateAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `FixedaddresstemplateAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fixedaddresstemplate** | [**Fixedaddresstemplate**](Fixedaddresstemplate.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateFixedaddresstemplateResponse**](CreateFixedaddresstemplateResponse.md)

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

Delete a fixedaddresstemplate object



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
	reference := "reference_example" // string | Reference of the fixedaddresstemplate object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.FixedaddresstemplateAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddresstemplateAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the fixedaddresstemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `FixedaddresstemplateAPIDeleteRequest` struct via the builder pattern


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

> ListFixedaddresstemplateResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve fixedaddresstemplate objects



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
	resp, r, err := apiClient.FixedaddresstemplateAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddresstemplateAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListFixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `FixedaddresstemplateAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `FixedaddresstemplateAPIListRequest` struct via the builder pattern


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

[**ListFixedaddresstemplateResponse**](ListFixedaddresstemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetFixedaddresstemplateResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific fixedaddresstemplate object



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
	reference := "reference_example" // string | Reference of the fixedaddresstemplate object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.FixedaddresstemplateAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddresstemplateAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetFixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `FixedaddresstemplateAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the fixedaddresstemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `FixedaddresstemplateAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetFixedaddresstemplateResponse**](GetFixedaddresstemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateFixedaddresstemplateResponse Update(ctx, reference).Fixedaddresstemplate(fixedaddresstemplate).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a fixedaddresstemplate object



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
	reference := "reference_example" // string | Reference of the fixedaddresstemplate object
	fixedaddresstemplate := *dhcp.NewFixedaddresstemplate() // Fixedaddresstemplate | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.FixedaddresstemplateAPI.Update(context.Background(), reference).Fixedaddresstemplate(fixedaddresstemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FixedaddresstemplateAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateFixedaddresstemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `FixedaddresstemplateAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the fixedaddresstemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `FixedaddresstemplateAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**fixedaddresstemplate** | [**Fixedaddresstemplate**](Fixedaddresstemplate.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateFixedaddresstemplateResponse**](UpdateFixedaddresstemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

