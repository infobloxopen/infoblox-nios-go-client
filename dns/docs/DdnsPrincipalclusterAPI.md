# DdnsPrincipalclusterAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DdnsPrincipalclusterAPI.md#Create) | **Post** /ddns:principalcluster | Create a ddns:principalcluster object
[**Delete**](DdnsPrincipalclusterAPI.md#Delete) | **Delete** /ddns:principalcluster/{reference} | Delete a ddns:principalcluster object
[**List**](DdnsPrincipalclusterAPI.md#List) | **Get** /ddns:principalcluster | Retrieve ddns:principalcluster objects
[**Read**](DdnsPrincipalclusterAPI.md#Read) | **Get** /ddns:principalcluster/{reference} | Get a specific ddns:principalcluster object
[**Update**](DdnsPrincipalclusterAPI.md#Update) | **Put** /ddns:principalcluster/{reference} | Update a ddns:principalcluster object



## Create

> CreateDdnsPrincipalclusterResponse Create(ctx).DdnsPrincipalcluster(ddnsPrincipalcluster).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a ddns:principalcluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	ddnsPrincipalcluster := *dns.NewDdnsPrincipalcluster() // DdnsPrincipalcluster | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterAPI.Create(context.Background()).DdnsPrincipalcluster(ddnsPrincipalcluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateDdnsPrincipalclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ddnsPrincipalcluster** | [**DdnsPrincipalcluster**](DdnsPrincipalcluster.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDdnsPrincipalclusterResponse**](CreateDdnsPrincipalclusterResponse.md)

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

Delete a ddns:principalcluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the ddns:principalcluster object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.DdnsPrincipalclusterAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPIDeleteRequest` struct via the builder pattern


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

> ListDdnsPrincipalclusterResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve ddns:principalcluster objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListDdnsPrincipalclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPIListRequest` struct via the builder pattern


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

[**ListDdnsPrincipalclusterResponse**](ListDdnsPrincipalclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetDdnsPrincipalclusterResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific ddns:principalcluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the ddns:principalcluster object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetDdnsPrincipalclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetDdnsPrincipalclusterResponse**](GetDdnsPrincipalclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateDdnsPrincipalclusterResponse Update(ctx, reference).DdnsPrincipalcluster(ddnsPrincipalcluster).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a ddns:principalcluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the ddns:principalcluster object
	ddnsPrincipalcluster := *dns.NewDdnsPrincipalcluster() // DdnsPrincipalcluster | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterAPI.Update(context.Background(), reference).DdnsPrincipalcluster(ddnsPrincipalcluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateDdnsPrincipalclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ddnsPrincipalcluster** | [**DdnsPrincipalcluster**](DdnsPrincipalcluster.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDdnsPrincipalclusterResponse**](UpdateDdnsPrincipalclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

