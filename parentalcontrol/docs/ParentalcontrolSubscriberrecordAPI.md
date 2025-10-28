# ParentalcontrolSubscriberrecordAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ParentalcontrolSubscriberrecordAPI.md#Create) | **Post** /parentalcontrol:subscriberrecord | Create a parentalcontrol:subscriberrecord object
[**Delete**](ParentalcontrolSubscriberrecordAPI.md#Delete) | **Delete** /parentalcontrol:subscriberrecord/{reference} | Delete a parentalcontrol:subscriberrecord object
[**List**](ParentalcontrolSubscriberrecordAPI.md#List) | **Get** /parentalcontrol:subscriberrecord | Retrieve parentalcontrol:subscriberrecord objects
[**Read**](ParentalcontrolSubscriberrecordAPI.md#Read) | **Get** /parentalcontrol:subscriberrecord/{reference} | Get a specific parentalcontrol:subscriberrecord object
[**Update**](ParentalcontrolSubscriberrecordAPI.md#Update) | **Put** /parentalcontrol:subscriberrecord/{reference} | Update a parentalcontrol:subscriberrecord object



## Create

> CreateParentalcontrolSubscriberrecordResponse Create(ctx).ParentalcontrolSubscriberrecord(parentalcontrolSubscriberrecord).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a parentalcontrol:subscriberrecord object



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
	parentalcontrolSubscriberrecord := *parentalcontrol.NewParentalcontrolSubscriberrecord() // ParentalcontrolSubscriberrecord | Object data to create

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberrecordAPI.Create(context.Background()).ParentalcontrolSubscriberrecord(parentalcontrolSubscriberrecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateParentalcontrolSubscriberrecordResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberrecordAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolSubscriberrecord** | [**ParentalcontrolSubscriberrecord**](ParentalcontrolSubscriberrecord.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateParentalcontrolSubscriberrecordResponse**](CreateParentalcontrolSubscriberrecordResponse.md)

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

Delete a parentalcontrol:subscriberrecord object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriberrecord object

	apiClient := parentalcontrol.NewAPIClient()
	r, err := apiClient.ParentalcontrolSubscriberrecordAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscriberrecord object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPIDeleteRequest` struct via the builder pattern


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

> ListParentalcontrolSubscriberrecordResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve parentalcontrol:subscriberrecord objects



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
	resp, r, err := apiClient.ParentalcontrolSubscriberrecordAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListParentalcontrolSubscriberrecordResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberrecordAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPIListRequest` struct via the builder pattern


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

[**ListParentalcontrolSubscriberrecordResponse**](ListParentalcontrolSubscriberrecordResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetParentalcontrolSubscriberrecordResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific parentalcontrol:subscriberrecord object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriberrecord object

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberrecordAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetParentalcontrolSubscriberrecordResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberrecordAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscriberrecord object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetParentalcontrolSubscriberrecordResponse**](GetParentalcontrolSubscriberrecordResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateParentalcontrolSubscriberrecordResponse Update(ctx, reference).ParentalcontrolSubscriberrecord(parentalcontrolSubscriberrecord).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a parentalcontrol:subscriberrecord object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriberrecord object
	parentalcontrolSubscriberrecord := *parentalcontrol.NewParentalcontrolSubscriberrecord() // ParentalcontrolSubscriberrecord | Object data to update

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberrecordAPI.Update(context.Background(), reference).ParentalcontrolSubscriberrecord(parentalcontrolSubscriberrecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateParentalcontrolSubscriberrecordResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberrecordAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscriberrecord object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolSubscriberrecord** | [**ParentalcontrolSubscriberrecord**](ParentalcontrolSubscriberrecord.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateParentalcontrolSubscriberrecordResponse**](UpdateParentalcontrolSubscriberrecordResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

