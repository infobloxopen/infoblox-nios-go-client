# TacacsplusAuthserviceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](TacacsplusAuthserviceAPI.md#Create) | **Post** /tacacsplus:authservice | Create a tacacsplus:authservice object
[**Delete**](TacacsplusAuthserviceAPI.md#Delete) | **Delete** /tacacsplus:authservice/{reference} | Delete a tacacsplus:authservice object
[**List**](TacacsplusAuthserviceAPI.md#List) | **Get** /tacacsplus:authservice | Retrieve tacacsplus:authservice objects
[**Read**](TacacsplusAuthserviceAPI.md#Read) | **Get** /tacacsplus:authservice/{reference} | Get a specific tacacsplus:authservice object
[**Update**](TacacsplusAuthserviceAPI.md#Update) | **Put** /tacacsplus:authservice/{reference} | Update a tacacsplus:authservice object



## Create

> CreateTacacsplusAuthserviceResponse Create(ctx).TacacsplusAuthservice(tacacsplusAuthservice).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a tacacsplus:authservice object



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
	tacacsplusAuthservice := *security.NewTacacsplusAuthservice() // TacacsplusAuthservice | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.TacacsplusAuthserviceAPI.Create(context.Background()).TacacsplusAuthservice(tacacsplusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusAuthserviceAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**tacacsplusAuthservice** | [**TacacsplusAuthservice**](TacacsplusAuthservice.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateTacacsplusAuthserviceResponse**](CreateTacacsplusAuthserviceResponse.md)

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

Delete a tacacsplus:authservice object



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
	reference := "reference_example" // string | Reference of the tacacsplus:authservice object

	apiClient := security.NewAPIClient()
	r, err := apiClient.TacacsplusAuthserviceAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the tacacsplus:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPIDeleteRequest` struct via the builder pattern


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

> ListTacacsplusAuthserviceResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve tacacsplus:authservice objects



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
	resp, r, err := apiClient.TacacsplusAuthserviceAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusAuthserviceAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPIListRequest` struct via the builder pattern


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

[**ListTacacsplusAuthserviceResponse**](ListTacacsplusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetTacacsplusAuthserviceResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific tacacsplus:authservice object



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
	reference := "reference_example" // string | Reference of the tacacsplus:authservice object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.TacacsplusAuthserviceAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusAuthserviceAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the tacacsplus:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetTacacsplusAuthserviceResponse**](GetTacacsplusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateTacacsplusAuthserviceResponse Update(ctx, reference).TacacsplusAuthservice(tacacsplusAuthservice).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a tacacsplus:authservice object



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
	reference := "reference_example" // string | Reference of the tacacsplus:authservice object
	tacacsplusAuthservice := *security.NewTacacsplusAuthservice() // TacacsplusAuthservice | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.TacacsplusAuthserviceAPI.Update(context.Background(), reference).TacacsplusAuthservice(tacacsplusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusAuthserviceAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the tacacsplus:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**tacacsplusAuthservice** | [**TacacsplusAuthservice**](TacacsplusAuthservice.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateTacacsplusAuthserviceResponse**](UpdateTacacsplusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

