# DdnsPrincipalclusterGroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DdnsPrincipalclusterGroupAPI.md#Create) | **Post** /ddns:principalcluster:group | Create a ddns:principalcluster:group object
[**Delete**](DdnsPrincipalclusterGroupAPI.md#Delete) | **Delete** /ddns:principalcluster:group/{reference} | Delete a ddns:principalcluster:group object
[**List**](DdnsPrincipalclusterGroupAPI.md#List) | **Get** /ddns:principalcluster:group | Retrieve ddns:principalcluster:group objects
[**Read**](DdnsPrincipalclusterGroupAPI.md#Read) | **Get** /ddns:principalcluster:group/{reference} | Get a specific ddns:principalcluster:group object
[**Update**](DdnsPrincipalclusterGroupAPI.md#Update) | **Put** /ddns:principalcluster:group/{reference} | Update a ddns:principalcluster:group object



## Create

> CreateDdnsPrincipalclusterGroupResponse Create(ctx).DdnsPrincipalclusterGroup(ddnsPrincipalclusterGroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a ddns:principalcluster:group object



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
	ddnsPrincipalclusterGroup := *dns.NewDdnsPrincipalclusterGroup() // DdnsPrincipalclusterGroup | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterGroupAPI.Create(context.Background()).DdnsPrincipalclusterGroup(ddnsPrincipalclusterGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateDdnsPrincipalclusterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterGroupAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ddnsPrincipalclusterGroup** | [**DdnsPrincipalclusterGroup**](DdnsPrincipalclusterGroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDdnsPrincipalclusterGroupResponse**](CreateDdnsPrincipalclusterGroupResponse.md)

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

Delete a ddns:principalcluster:group object



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
	reference := "reference_example" // string | Reference of the ddns:principalcluster:group object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.DdnsPrincipalclusterGroupAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPIDeleteRequest` struct via the builder pattern


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

> ListDdnsPrincipalclusterGroupResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve ddns:principalcluster:group objects



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
	resp, r, err := apiClient.DdnsPrincipalclusterGroupAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListDdnsPrincipalclusterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterGroupAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPIListRequest` struct via the builder pattern


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

[**ListDdnsPrincipalclusterGroupResponse**](ListDdnsPrincipalclusterGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetDdnsPrincipalclusterGroupResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific ddns:principalcluster:group object



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
	reference := "reference_example" // string | Reference of the ddns:principalcluster:group object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterGroupAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetDdnsPrincipalclusterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterGroupAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetDdnsPrincipalclusterGroupResponse**](GetDdnsPrincipalclusterGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateDdnsPrincipalclusterGroupResponse Update(ctx, reference).DdnsPrincipalclusterGroup(ddnsPrincipalclusterGroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a ddns:principalcluster:group object



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
	reference := "reference_example" // string | Reference of the ddns:principalcluster:group object
	ddnsPrincipalclusterGroup := *dns.NewDdnsPrincipalclusterGroup() // DdnsPrincipalclusterGroup | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterGroupAPI.Update(context.Background(), reference).DdnsPrincipalclusterGroup(ddnsPrincipalclusterGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateDdnsPrincipalclusterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterGroupAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ddnsPrincipalclusterGroup** | [**DdnsPrincipalclusterGroup**](DdnsPrincipalclusterGroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDdnsPrincipalclusterGroupResponse**](UpdateDdnsPrincipalclusterGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

