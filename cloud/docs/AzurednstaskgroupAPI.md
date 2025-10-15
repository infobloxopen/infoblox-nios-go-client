# AzurednstaskgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](AzurednstaskgroupAPI.md#Create) | **Post** /azurednstaskgroup | Create a azurednstaskgroup object
[**Delete**](AzurednstaskgroupAPI.md#Delete) | **Delete** /azurednstaskgroup/{reference} | Delete a azurednstaskgroup object
[**List**](AzurednstaskgroupAPI.md#List) | **Get** /azurednstaskgroup | Retrieve azurednstaskgroup objects
[**Read**](AzurednstaskgroupAPI.md#Read) | **Get** /azurednstaskgroup/{reference} | Get a specific azurednstaskgroup object
[**Update**](AzurednstaskgroupAPI.md#Update) | **Put** /azurednstaskgroup/{reference} | Update a azurednstaskgroup object



## Create

> CreateAzurednstaskgroupResponse Create(ctx).Azurednstaskgroup(azurednstaskgroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a azurednstaskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	azurednstaskgroup := *cloud.NewAzurednstaskgroup() // Azurednstaskgroup | Object data to create

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.AzurednstaskgroupAPI.Create(context.Background()).Azurednstaskgroup(azurednstaskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzurednstaskgroupAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateAzurednstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `AzurednstaskgroupAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `AzurednstaskgroupAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**azurednstaskgroup** | [**Azurednstaskgroup**](Azurednstaskgroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateAzurednstaskgroupResponse**](CreateAzurednstaskgroupResponse.md)

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

Delete a azurednstaskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the azurednstaskgroup object

	apiClient := cloud.NewAPIClient()
	r, err := apiClient.AzurednstaskgroupAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzurednstaskgroupAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the azurednstaskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `AzurednstaskgroupAPIDeleteRequest` struct via the builder pattern


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

> ListAzurednstaskgroupResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve azurednstaskgroup objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.AzurednstaskgroupAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzurednstaskgroupAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListAzurednstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `AzurednstaskgroupAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `AzurednstaskgroupAPIListRequest` struct via the builder pattern


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

[**ListAzurednstaskgroupResponse**](ListAzurednstaskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetAzurednstaskgroupResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific azurednstaskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the azurednstaskgroup object

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.AzurednstaskgroupAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzurednstaskgroupAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetAzurednstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `AzurednstaskgroupAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the azurednstaskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `AzurednstaskgroupAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetAzurednstaskgroupResponse**](GetAzurednstaskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateAzurednstaskgroupResponse Update(ctx, reference).Azurednstaskgroup(azurednstaskgroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a azurednstaskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the azurednstaskgroup object
	azurednstaskgroup := *cloud.NewAzurednstaskgroup() // Azurednstaskgroup | Object data to update

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.AzurednstaskgroupAPI.Update(context.Background(), reference).Azurednstaskgroup(azurednstaskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AzurednstaskgroupAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateAzurednstaskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `AzurednstaskgroupAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the azurednstaskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `AzurednstaskgroupAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**azurednstaskgroup** | [**Azurednstaskgroup**](Azurednstaskgroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateAzurednstaskgroupResponse**](UpdateAzurednstaskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

