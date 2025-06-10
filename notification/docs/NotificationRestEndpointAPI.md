# NotificationRestEndpointAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationrestendpointGet**](NotificationRestEndpointAPI.md#NotificationrestendpointGet) | **Get** /notification:rest:endpoint | Retrieve notification:rest:endpoint objects
[**NotificationrestendpointPost**](NotificationRestEndpointAPI.md#NotificationrestendpointPost) | **Post** /notification:rest:endpoint | Create a notification:rest:endpoint object
[**NotificationrestendpointReferenceDelete**](NotificationRestEndpointAPI.md#NotificationrestendpointReferenceDelete) | **Delete** /notification:rest:endpoint/{reference} | Delete a notification:rest:endpoint object
[**NotificationrestendpointReferenceGet**](NotificationRestEndpointAPI.md#NotificationrestendpointReferenceGet) | **Get** /notification:rest:endpoint/{reference} | Get a specific notification:rest:endpoint object
[**NotificationrestendpointReferencePut**](NotificationRestEndpointAPI.md#NotificationrestendpointReferencePut) | **Put** /notification:rest:endpoint/{reference} | Update a notification:rest:endpoint object



## NotificationrestendpointGet

> ListNotificationRestEndpointResponse NotificationrestendpointGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve notification:rest:endpoint objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/notification"
)

func main() {

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestEndpointAPI.NotificationrestendpointGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.NotificationrestendpointGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationrestendpointGet`: ListNotificationRestEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestEndpointAPI.NotificationrestendpointGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestEndpointAPINotificationrestendpointGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**paging** | **int32** | Control paging of results | 
**pageId** | **string** | Page id for retrieving next page of results | 
**filters** | **map[string]interface{}** |  | 
**extattrfilter** | **map[string]interface{}** |  | 

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


## NotificationrestendpointPost

> CreateNotificationRestEndpointResponse NotificationrestendpointPost(ctx).NotificationRestEndpoint(notificationRestEndpoint).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a notification:rest:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/notification"
)

func main() {
	notificationRestEndpoint := *notification.NewNotificationRestEndpoint() // NotificationRestEndpoint | Object data to create

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestEndpointAPI.NotificationrestendpointPost(context.Background()).NotificationRestEndpoint(notificationRestEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.NotificationrestendpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationrestendpointPost`: CreateNotificationRestEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestEndpointAPI.NotificationrestendpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestEndpointAPINotificationrestendpointPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**notificationRestEndpoint** | [**NotificationRestEndpoint**](NotificationRestEndpoint.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## NotificationrestendpointReferenceDelete

> NotificationrestendpointReferenceDelete(ctx, reference).Execute()

Delete a notification:rest:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/notification"
)

func main() {
	reference := "reference_example" // string | Reference of the notification:rest:endpoint object

	apiClient := notification.NewAPIClient()
	r, err := apiClient.NotificationRestEndpointAPI.NotificationrestendpointReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.NotificationrestendpointReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `NotificationRestEndpointAPINotificationrestendpointReferenceDeleteRequest` struct via the builder pattern


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


## NotificationrestendpointReferenceGet

> GetNotificationRestEndpointResponse NotificationrestendpointReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific notification:rest:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/notification"
)

func main() {
	reference := "reference_example" // string | Reference of the notification:rest:endpoint object

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestEndpointAPI.NotificationrestendpointReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.NotificationrestendpointReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationrestendpointReferenceGet`: GetNotificationRestEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestEndpointAPI.NotificationrestendpointReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rest:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestEndpointAPINotificationrestendpointReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## NotificationrestendpointReferencePut

> UpdateNotificationRestEndpointResponse NotificationrestendpointReferencePut(ctx, reference).NotificationRestEndpoint(notificationRestEndpoint).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a notification:rest:endpoint object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/notification"
)

func main() {
	reference := "reference_example" // string | Reference of the notification:rest:endpoint object
	notificationRestEndpoint := *notification.NewNotificationRestEndpoint() // NotificationRestEndpoint | Object data to update

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestEndpointAPI.NotificationrestendpointReferencePut(context.Background(), reference).NotificationRestEndpoint(notificationRestEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestEndpointAPI.NotificationrestendpointReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationrestendpointReferencePut`: UpdateNotificationRestEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestEndpointAPI.NotificationrestendpointReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rest:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestEndpointAPINotificationrestendpointReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**notificationRestEndpoint** | [**NotificationRestEndpoint**](NotificationRestEndpoint.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

