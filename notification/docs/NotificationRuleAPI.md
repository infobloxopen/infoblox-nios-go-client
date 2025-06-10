# NotificationRuleAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NotificationruleGet**](NotificationRuleAPI.md#NotificationruleGet) | **Get** /notification:rule | Retrieve notification:rule objects
[**NotificationrulePost**](NotificationRuleAPI.md#NotificationrulePost) | **Post** /notification:rule | Create a notification:rule object
[**NotificationruleReferenceDelete**](NotificationRuleAPI.md#NotificationruleReferenceDelete) | **Delete** /notification:rule/{reference} | Delete a notification:rule object
[**NotificationruleReferenceGet**](NotificationRuleAPI.md#NotificationruleReferenceGet) | **Get** /notification:rule/{reference} | Get a specific notification:rule object
[**NotificationruleReferencePut**](NotificationRuleAPI.md#NotificationruleReferencePut) | **Put** /notification:rule/{reference} | Update a notification:rule object



## NotificationruleGet

> ListNotificationRuleResponse NotificationruleGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve notification:rule objects



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
	resp, r, err := apiClient.NotificationRuleAPI.NotificationruleGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRuleAPI.NotificationruleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationruleGet`: ListNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRuleAPI.NotificationruleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NotificationRuleAPINotificationruleGetRequest` struct via the builder pattern


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

[**ListNotificationRuleResponse**](ListNotificationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationrulePost

> CreateNotificationRuleResponse NotificationrulePost(ctx).NotificationRule(notificationRule).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a notification:rule object



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
	notificationRule := *notification.NewNotificationRule() // NotificationRule | Object data to create

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRuleAPI.NotificationrulePost(context.Background()).NotificationRule(notificationRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRuleAPI.NotificationrulePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationrulePost`: CreateNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRuleAPI.NotificationrulePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NotificationRuleAPINotificationrulePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**notificationRule** | [**NotificationRule**](NotificationRule.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateNotificationRuleResponse**](CreateNotificationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationruleReferenceDelete

> NotificationruleReferenceDelete(ctx, reference).Execute()

Delete a notification:rule object



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
	reference := "reference_example" // string | Reference of the notification:rule object

	apiClient := notification.NewAPIClient()
	r, err := apiClient.NotificationRuleAPI.NotificationruleReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRuleAPI.NotificationruleReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRuleAPINotificationruleReferenceDeleteRequest` struct via the builder pattern


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


## NotificationruleReferenceGet

> GetNotificationRuleResponse NotificationruleReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific notification:rule object



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
	reference := "reference_example" // string | Reference of the notification:rule object

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRuleAPI.NotificationruleReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRuleAPI.NotificationruleReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationruleReferenceGet`: GetNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRuleAPI.NotificationruleReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRuleAPINotificationruleReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetNotificationRuleResponse**](GetNotificationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NotificationruleReferencePut

> UpdateNotificationRuleResponse NotificationruleReferencePut(ctx, reference).NotificationRule(notificationRule).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a notification:rule object



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
	reference := "reference_example" // string | Reference of the notification:rule object
	notificationRule := *notification.NewNotificationRule() // NotificationRule | Object data to update

	apiClient := notification.NewAPIClient()
	resp, r, err := apiClient.NotificationRuleAPI.NotificationruleReferencePut(context.Background(), reference).NotificationRule(notificationRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationRuleAPI.NotificationruleReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NotificationruleReferencePut`: UpdateNotificationRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationRuleAPI.NotificationruleReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the notification:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `NotificationRuleAPINotificationruleReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**notificationRule** | [**NotificationRule**](NotificationRule.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateNotificationRuleResponse**](UpdateNotificationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

