# SmartfolderGlobalAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](SmartfolderGlobalAPI.md#Create) | **Post** /smartfolder:global | Create a smartfolder:global object
[**Delete**](SmartfolderGlobalAPI.md#Delete) | **Delete** /smartfolder:global/{reference} | Delete a smartfolder:global object
[**List**](SmartfolderGlobalAPI.md#List) | **Get** /smartfolder:global | Retrieve smartfolder:global objects
[**Read**](SmartfolderGlobalAPI.md#Read) | **Get** /smartfolder:global/{reference} | Get a specific smartfolder:global object
[**Update**](SmartfolderGlobalAPI.md#Update) | **Put** /smartfolder:global/{reference} | Update a smartfolder:global object



## Create

> CreateSmartfolderGlobalResponse Create(ctx).SmartfolderGlobal(smartfolderGlobal).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a smartfolder:global object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
)

func main() {
	smartfolderGlobal := *smartfolder.NewSmartfolderGlobal() // SmartfolderGlobal | Object data to create

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderGlobalAPI.Create(context.Background()).SmartfolderGlobal(smartfolderGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateSmartfolderGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderGlobalAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**smartfolderGlobal** | [**SmartfolderGlobal**](SmartfolderGlobal.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSmartfolderGlobalResponse**](CreateSmartfolderGlobalResponse.md)

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

Delete a smartfolder:global object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
)

func main() {
	reference := "reference_example" // string | Reference of the smartfolder:global object

	apiClient := smartfolder.NewAPIClient()
	r, err := apiClient.SmartfolderGlobalAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:global object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPIDeleteRequest` struct via the builder pattern


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

> ListSmartfolderGlobalResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve smartfolder:global objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
)

func main() {

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderGlobalAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListSmartfolderGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderGlobalAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPIListRequest` struct via the builder pattern


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

[**ListSmartfolderGlobalResponse**](ListSmartfolderGlobalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetSmartfolderGlobalResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific smartfolder:global object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
)

func main() {
	reference := "reference_example" // string | Reference of the smartfolder:global object

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderGlobalAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetSmartfolderGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderGlobalAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:global object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetSmartfolderGlobalResponse**](GetSmartfolderGlobalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateSmartfolderGlobalResponse Update(ctx, reference).SmartfolderGlobal(smartfolderGlobal).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a smartfolder:global object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
)

func main() {
	reference := "reference_example" // string | Reference of the smartfolder:global object
	smartfolderGlobal := *smartfolder.NewSmartfolderGlobal() // SmartfolderGlobal | Object data to update

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderGlobalAPI.Update(context.Background(), reference).SmartfolderGlobal(smartfolderGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateSmartfolderGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderGlobalAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:global object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**smartfolderGlobal** | [**SmartfolderGlobal**](SmartfolderGlobal.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSmartfolderGlobalResponse**](UpdateSmartfolderGlobalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

