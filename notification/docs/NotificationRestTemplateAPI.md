# NotificationRestTemplateAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationresttemplateGet**](NotificationRestTemplateAPI.md#NotificationresttemplateGet) | **Get** /notification:rest:template | Retrieve notification:rest:template objects
[**NotificationresttemplateReferenceDelete**](NotificationRestTemplateAPI.md#NotificationresttemplateReferenceDelete) | **Delete** /notification:rest:template/{reference} | Delete a notification:rest:template object
[**NotificationresttemplateReferenceGet**](NotificationRestTemplateAPI.md#NotificationresttemplateReferenceGet) | **Get** /notification:rest:template/{reference} | Get a specific notification:rest:template object
[**NotificationresttemplateReferencePut**](NotificationRestTemplateAPI.md#NotificationresttemplateReferencePut) | **Put** /notification:rest:template/{reference} | Update a notification:rest:template object



## NotificationresttemplateGet

> ListNotificationRestTemplateResponse NotificationresttemplateGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve notification:rest:template objects



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
	resp, r, err := apiClient.NotificationRestTemplateAPI.NotificationresttemplateGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestTemplateAPI.NotificationresttemplateGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationresttemplateGet`: ListNotificationRestTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestTemplateAPI.NotificationresttemplateGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestTemplateAPINotificationresttemplateGetRequest` struct via the builder pattern


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

[**ListNotificationRestTemplateResponse**](ListNotificationRestTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationresttemplateReferenceDelete

> NotificationresttemplateReferenceDelete(ctx, reference).Execute()

Delete a notification:rest:template object



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
	reference := "reference_example" // string | Reference of the notification:rest:template object

	apiClient := notification.NewAPIClient()
	r, err := apiClient.NotificationRestTemplateAPI.NotificationresttemplateReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestTemplateAPI.NotificationresttemplateReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rest:template object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestTemplateAPINotificationresttemplateReferenceDeleteRequest` struct via the builder pattern


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


## NotificationresttemplateReferenceGet

> GetNotificationRestTemplateResponse NotificationresttemplateReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific notification:rest:template object



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
	reference := "reference_example" // string | Reference of the notification:rest:template object

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestTemplateAPI.NotificationresttemplateReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestTemplateAPI.NotificationresttemplateReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationresttemplateReferenceGet`: GetNotificationRestTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestTemplateAPI.NotificationresttemplateReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rest:template object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestTemplateAPINotificationresttemplateReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetNotificationRestTemplateResponse**](GetNotificationRestTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationresttemplateReferencePut

> UpdateNotificationRestTemplateResponse NotificationresttemplateReferencePut(ctx, reference).NotificationRestTemplate(notificationRestTemplate).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a notification:rest:template object



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
	reference := "reference_example" // string | Reference of the notification:rest:template object
	notificationRestTemplate := *notification.NewNotificationRestTemplate() // NotificationRestTemplate | Object data to update

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRestTemplateAPI.NotificationresttemplateReferencePut(context.Background(), reference).NotificationRestTemplate(notificationRestTemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRestTemplateAPI.NotificationresttemplateReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationresttemplateReferencePut`: UpdateNotificationRestTemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRestTemplateAPI.NotificationresttemplateReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rest:template object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRestTemplateAPINotificationresttemplateReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**notificationRestTemplate** | [**NotificationRestTemplate**](NotificationRestTemplate.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateNotificationRestTemplateResponse**](UpdateNotificationRestTemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

