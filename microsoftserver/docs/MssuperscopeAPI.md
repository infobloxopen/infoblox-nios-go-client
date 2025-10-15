# MssuperscopeAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](MssuperscopeAPI.md#Create) | **Post** /mssuperscope | Create a mssuperscope object
[**Delete**](MssuperscopeAPI.md#Delete) | **Delete** /mssuperscope/{reference} | Delete a mssuperscope object
[**List**](MssuperscopeAPI.md#List) | **Get** /mssuperscope | Retrieve mssuperscope objects
[**Read**](MssuperscopeAPI.md#Read) | **Get** /mssuperscope/{reference} | Get a specific mssuperscope object
[**Update**](MssuperscopeAPI.md#Update) | **Put** /mssuperscope/{reference} | Update a mssuperscope object



## Create

> CreateMssuperscopeResponse Create(ctx).Mssuperscope(mssuperscope).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a mssuperscope object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
)

func main() {
	mssuperscope := *microsoftserver.NewMssuperscope() // Mssuperscope | Object data to create

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MssuperscopeAPI.Create(context.Background()).Mssuperscope(mssuperscope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateMssuperscopeResponse
	fmt.Fprintf(os.Stdout, "Response from `MssuperscopeAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**mssuperscope** | [**Mssuperscope**](Mssuperscope.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateMssuperscopeResponse**](CreateMssuperscopeResponse.md)

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

Delete a mssuperscope object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
)

func main() {
	reference := "reference_example" // string | Reference of the mssuperscope object

	apiClient := microsoftserver.NewAPIClient()
	r, err := apiClient.MssuperscopeAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the mssuperscope object | 

### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPIDeleteRequest` struct via the builder pattern


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

> ListMssuperscopeResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve mssuperscope objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
)

func main() {

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MssuperscopeAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListMssuperscopeResponse
	fmt.Fprintf(os.Stdout, "Response from `MssuperscopeAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPIListRequest` struct via the builder pattern


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

[**ListMssuperscopeResponse**](ListMssuperscopeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetMssuperscopeResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific mssuperscope object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
)

func main() {
	reference := "reference_example" // string | Reference of the mssuperscope object

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MssuperscopeAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetMssuperscopeResponse
	fmt.Fprintf(os.Stdout, "Response from `MssuperscopeAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the mssuperscope object | 

### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetMssuperscopeResponse**](GetMssuperscopeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateMssuperscopeResponse Update(ctx, reference).Mssuperscope(mssuperscope).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a mssuperscope object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
)

func main() {
	reference := "reference_example" // string | Reference of the mssuperscope object
	mssuperscope := *microsoftserver.NewMssuperscope() // Mssuperscope | Object data to update

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MssuperscopeAPI.Update(context.Background(), reference).Mssuperscope(mssuperscope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateMssuperscopeResponse
	fmt.Fprintf(os.Stdout, "Response from `MssuperscopeAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the mssuperscope object | 

### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**mssuperscope** | [**Mssuperscope**](Mssuperscope.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMssuperscopeResponse**](UpdateMssuperscopeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

