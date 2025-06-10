# ParentalcontrolSubscribersiteAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ParentalcontrolsubscribersiteGet**](ParentalcontrolSubscribersiteAPI.md#ParentalcontrolsubscribersiteGet) | **Get** /parentalcontrol:subscribersite | Retrieve parentalcontrol:subscribersite objects
[**ParentalcontrolsubscribersitePost**](ParentalcontrolSubscribersiteAPI.md#ParentalcontrolsubscribersitePost) | **Post** /parentalcontrol:subscribersite | Create a parentalcontrol:subscribersite object
[**ParentalcontrolsubscribersiteReferenceDelete**](ParentalcontrolSubscribersiteAPI.md#ParentalcontrolsubscribersiteReferenceDelete) | **Delete** /parentalcontrol:subscribersite/{reference} | Delete a parentalcontrol:subscribersite object
[**ParentalcontrolsubscribersiteReferenceGet**](ParentalcontrolSubscribersiteAPI.md#ParentalcontrolsubscribersiteReferenceGet) | **Get** /parentalcontrol:subscribersite/{reference} | Get a specific parentalcontrol:subscribersite object
[**ParentalcontrolsubscribersiteReferencePut**](ParentalcontrolSubscribersiteAPI.md#ParentalcontrolsubscribersiteReferencePut) | **Put** /parentalcontrol:subscribersite/{reference} | Update a parentalcontrol:subscribersite object



## ParentalcontrolsubscribersiteGet

> ListParentalcontrolSubscribersiteResponse ParentalcontrolsubscribersiteGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve parentalcontrol:subscribersite objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/parentalcontrol"
)

func main() {

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscribersiteGet`: ListParentalcontrolSubscribersiteResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscribersiteAPIParentalcontrolsubscribersiteGetRequest` struct via the builder pattern


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

[**ListParentalcontrolSubscribersiteResponse**](ListParentalcontrolSubscribersiteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentalcontrolsubscribersitePost

> CreateParentalcontrolSubscribersiteResponse ParentalcontrolsubscribersitePost(ctx).ParentalcontrolSubscribersite(parentalcontrolSubscribersite).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a parentalcontrol:subscribersite object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/parentalcontrol"
)

func main() {
	parentalcontrolSubscribersite := *parentalcontrol.NewParentalcontrolSubscribersite() // ParentalcontrolSubscribersite | Object data to create

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersitePost(context.Background()).ParentalcontrolSubscribersite(parentalcontrolSubscribersite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersitePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscribersitePost`: CreateParentalcontrolSubscribersiteResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersitePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscribersiteAPIParentalcontrolsubscribersitePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolSubscribersite** | [**ParentalcontrolSubscribersite**](ParentalcontrolSubscribersite.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateParentalcontrolSubscribersiteResponse**](CreateParentalcontrolSubscribersiteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentalcontrolsubscribersiteReferenceDelete

> ParentalcontrolsubscribersiteReferenceDelete(ctx, reference).Execute()

Delete a parentalcontrol:subscribersite object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/parentalcontrol"
)

func main() {
	reference := "reference_example" // string | Reference of the parentalcontrol:subscribersite object

	apiClient := parentalcontrol.NewAPIClient()
	r, err := apiClient.ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscribersite object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscribersiteAPIParentalcontrolsubscribersiteReferenceDeleteRequest` struct via the builder pattern


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


## ParentalcontrolsubscribersiteReferenceGet

> GetParentalcontrolSubscribersiteResponse ParentalcontrolsubscribersiteReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific parentalcontrol:subscribersite object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/parentalcontrol"
)

func main() {
	reference := "reference_example" // string | Reference of the parentalcontrol:subscribersite object

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscribersiteReferenceGet`: GetParentalcontrolSubscribersiteResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscribersite object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscribersiteAPIParentalcontrolsubscribersiteReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetParentalcontrolSubscribersiteResponse**](GetParentalcontrolSubscribersiteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentalcontrolsubscribersiteReferencePut

> UpdateParentalcontrolSubscribersiteResponse ParentalcontrolsubscribersiteReferencePut(ctx, reference).ParentalcontrolSubscribersite(parentalcontrolSubscribersite).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a parentalcontrol:subscribersite object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/parentalcontrol"
)

func main() {
	reference := "reference_example" // string | Reference of the parentalcontrol:subscribersite object
	parentalcontrolSubscribersite := *parentalcontrol.NewParentalcontrolSubscribersite() // ParentalcontrolSubscribersite | Object data to update

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteReferencePut(context.Background(), reference).ParentalcontrolSubscribersite(parentalcontrolSubscribersite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscribersiteReferencePut`: UpdateParentalcontrolSubscribersiteResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscribersiteAPI.ParentalcontrolsubscribersiteReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscribersite object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscribersiteAPIParentalcontrolsubscribersiteReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolSubscribersite** | [**ParentalcontrolSubscribersite**](ParentalcontrolSubscribersite.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateParentalcontrolSubscribersiteResponse**](UpdateParentalcontrolSubscribersiteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

