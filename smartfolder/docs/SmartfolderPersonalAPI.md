# SmartfolderPersonalAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](SmartfolderPersonalAPI.md#Create) | **Post** /smartfolder:personal | Create a smartfolder:personal object
[**Delete**](SmartfolderPersonalAPI.md#Delete) | **Delete** /smartfolder:personal/{reference} | Delete a smartfolder:personal object
[**List**](SmartfolderPersonalAPI.md#List) | **Get** /smartfolder:personal | Retrieve smartfolder:personal objects
[**Read**](SmartfolderPersonalAPI.md#Read) | **Get** /smartfolder:personal/{reference} | Get a specific smartfolder:personal object
[**Update**](SmartfolderPersonalAPI.md#Update) | **Put** /smartfolder:personal/{reference} | Update a smartfolder:personal object



## Create

> CreateSmartfolderPersonalResponse Create(ctx).SmartfolderPersonal(smartfolderPersonal).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a smartfolder:personal object



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
	smartfolderPersonal := *smartfolder.NewSmartfolderPersonal() // SmartfolderPersonal | Object data to create

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderPersonalAPI.Create(context.Background()).SmartfolderPersonal(smartfolderPersonal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateSmartfolderPersonalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderPersonalAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**smartfolderPersonal** | [**SmartfolderPersonal**](SmartfolderPersonal.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSmartfolderPersonalResponse**](CreateSmartfolderPersonalResponse.md)

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

Delete a smartfolder:personal object



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
	reference := "reference_example" // string | Reference of the smartfolder:personal object

	apiClient := smartfolder.NewAPIClient()
	r, err := apiClient.SmartfolderPersonalAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:personal object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPIDeleteRequest` struct via the builder pattern


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

> ListSmartfolderPersonalResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve smartfolder:personal objects



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
	resp, r, err := apiClient.SmartfolderPersonalAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListSmartfolderPersonalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderPersonalAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPIListRequest` struct via the builder pattern


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

[**ListSmartfolderPersonalResponse**](ListSmartfolderPersonalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetSmartfolderPersonalResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific smartfolder:personal object



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
	reference := "reference_example" // string | Reference of the smartfolder:personal object

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderPersonalAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetSmartfolderPersonalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderPersonalAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:personal object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetSmartfolderPersonalResponse**](GetSmartfolderPersonalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateSmartfolderPersonalResponse Update(ctx, reference).SmartfolderPersonal(smartfolderPersonal).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a smartfolder:personal object



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
	reference := "reference_example" // string | Reference of the smartfolder:personal object
	smartfolderPersonal := *smartfolder.NewSmartfolderPersonal() // SmartfolderPersonal | Object data to update

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderPersonalAPI.Update(context.Background(), reference).SmartfolderPersonal(smartfolderPersonal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateSmartfolderPersonalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderPersonalAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:personal object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**smartfolderPersonal** | [**SmartfolderPersonal**](SmartfolderPersonal.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSmartfolderPersonalResponse**](UpdateSmartfolderPersonalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

