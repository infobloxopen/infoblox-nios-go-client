# ParentalcontrolSubscriberAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](ParentalcontrolSubscriberAPI.md#List) | **Get** /parentalcontrol:subscriber | Retrieve parentalcontrol:subscriber objects
[**Read**](ParentalcontrolSubscriberAPI.md#Read) | **Get** /parentalcontrol:subscriber/{reference} | Get a specific parentalcontrol:subscriber object
[**Update**](ParentalcontrolSubscriberAPI.md#Update) | **Put** /parentalcontrol:subscriber/{reference} | Update a parentalcontrol:subscriber object



## List

> ListParentalcontrolSubscriberResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve parentalcontrol:subscriber objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
)

func main() {

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListParentalcontrolSubscriberResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberAPIListRequest` struct via the builder pattern


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

[**ListParentalcontrolSubscriberResponse**](ListParentalcontrolSubscriberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetParentalcontrolSubscriberResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific parentalcontrol:subscriber object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
)

func main() {
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriber object

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetParentalcontrolSubscriberResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscriber object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetParentalcontrolSubscriberResponse**](GetParentalcontrolSubscriberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateParentalcontrolSubscriberResponse Update(ctx, reference).ParentalcontrolSubscriber(parentalcontrolSubscriber).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a parentalcontrol:subscriber object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
)

func main() {
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriber object
	parentalcontrolSubscriber := *parentalcontrol.NewParentalcontrolSubscriber() // ParentalcontrolSubscriber | Object data to update

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberAPI.Update(context.Background(), reference).ParentalcontrolSubscriber(parentalcontrolSubscriber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateParentalcontrolSubscriberResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscriber object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolSubscriber** | [**ParentalcontrolSubscriber**](ParentalcontrolSubscriber.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateParentalcontrolSubscriberResponse**](UpdateParentalcontrolSubscriberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

