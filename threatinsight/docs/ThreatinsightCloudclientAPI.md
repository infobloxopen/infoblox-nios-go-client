# ThreatinsightCloudclientAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](ThreatinsightCloudclientAPI.md#List) | **Get** /threatinsight:cloudclient | Retrieve threatinsight:cloudclient objects
[**Read**](ThreatinsightCloudclientAPI.md#Read) | **Get** /threatinsight:cloudclient/{reference} | Get a specific threatinsight:cloudclient object
[**Update**](ThreatinsightCloudclientAPI.md#Update) | **Put** /threatinsight:cloudclient/{reference} | Update a threatinsight:cloudclient object



## List

> ListThreatinsightCloudclientResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve threatinsight:cloudclient objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/threatinsight"
)

func main() {

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightCloudclientAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightCloudclientAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListThreatinsightCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightCloudclientAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightCloudclientAPIListRequest` struct via the builder pattern


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

[**ListThreatinsightCloudclientResponse**](ListThreatinsightCloudclientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetThreatinsightCloudclientResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific threatinsight:cloudclient object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/threatinsight"
)

func main() {
	reference := "reference_example" // string | Reference of the threatinsight:cloudclient object

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightCloudclientAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightCloudclientAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetThreatinsightCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightCloudclientAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatinsight:cloudclient object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightCloudclientAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetThreatinsightCloudclientResponse**](GetThreatinsightCloudclientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateThreatinsightCloudclientResponse Update(ctx, reference).ThreatinsightCloudclient(threatinsightCloudclient).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a threatinsight:cloudclient object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/threatinsight"
)

func main() {
	reference := "reference_example" // string | Reference of the threatinsight:cloudclient object
	threatinsightCloudclient := *threatinsight.NewThreatinsightCloudclient() // ThreatinsightCloudclient | Object data to update

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightCloudclientAPI.Update(context.Background(), reference).ThreatinsightCloudclient(threatinsightCloudclient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightCloudclientAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateThreatinsightCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightCloudclientAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatinsight:cloudclient object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightCloudclientAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatinsightCloudclient** | [**ThreatinsightCloudclient**](ThreatinsightCloudclient.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateThreatinsightCloudclientResponse**](UpdateThreatinsightCloudclientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

