# NotificationRestEndpointAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](NotificationRestEndpointAPI.md#Create) | **Post** /notification:rest:endpoint | Create a notification:rest:endpoint object
[**Delete**](NotificationRestEndpointAPI.md#Delete) | **Delete** /notification:rest:endpoint/{reference} | Delete a notification:rest:endpoint object
[**List**](NotificationRestEndpointAPI.md#List) | **Get** /notification:rest:endpoint | Retrieve notification:rest:endpoint objects
[**Read**](NotificationRestEndpointAPI.md#Read) | **Get** /notification:rest:endpoint/{reference} | Get a specific notification:rest:endpoint object
[**Update**](NotificationRestEndpointAPI.md#Update) | **Put** /notification:rest:endpoint/{reference} | Update a notification:rest:endpoint object



## Create

> CreateNotificationRestEndpointResponse Create(ctx).NotificationRestEndpoint(notificationRestEndpoint).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a notification:rest:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"
)

func main() {
	notificationRestEndpoint := *notification.NewNotificationRestEndpoint() // NotificationRestEndpoint | Object data to create

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestEndpointAPI.Create(context.Background()).NotificationRestEndpoint(notificationRestEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateNotificationRestEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestEndpointAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestEndpointAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**notificationRestEndpoint** | [**NotificationRestEndpoint**](NotificationRestEndpoint.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateNotificationRestEndpointResponse**](CreateNotificationRestEndpointResponse.md)

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

Delete a notification:rest:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"
)

func main() {
	reference := "reference_example" // string | Reference of the notification:rest:endpoint object

	apiClient := notification.NewAPIClient()
	r, err := apiClient.NotificationRestEndpointAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rest:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestEndpointAPIDeleteRequest` struct via the builder pattern


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

> ListNotificationRestEndpointResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve notification:rest:endpoint objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"
)

func main() {

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestEndpointAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListNotificationRestEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestEndpointAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestEndpointAPIListRequest` struct via the builder pattern


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

[**ListNotificationRestEndpointResponse**](ListNotificationRestEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetNotificationRestEndpointResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific notification:rest:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"
)

func main() {
	reference := "reference_example" // string | Reference of the notification:rest:endpoint object

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestEndpointAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetNotificationRestEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestEndpointAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rest:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestEndpointAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetNotificationRestEndpointResponse**](GetNotificationRestEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateNotificationRestEndpointResponse Update(ctx, reference).NotificationRestEndpoint(notificationRestEndpoint).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a notification:rest:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/notification"
)

func main() {
	reference := "reference_example" // string | Reference of the notification:rest:endpoint object
	notificationRestEndpoint := *notification.NewNotificationRestEndpoint() // NotificationRestEndpoint | Object data to update

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestEndpointAPI.Update(context.Background(), reference).NotificationRestEndpoint(notificationRestEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateNotificationRestEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestEndpointAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rest:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestEndpointAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**notificationRestEndpoint** | [**NotificationRestEndpoint**](NotificationRestEndpoint.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateNotificationRestEndpointResponse**](UpdateNotificationRestEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

