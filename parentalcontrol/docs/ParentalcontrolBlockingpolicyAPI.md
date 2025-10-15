# ParentalcontrolBlockingpolicyAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ParentalcontrolBlockingpolicyAPI.md#Create) | **Post** /parentalcontrol:blockingpolicy | Create a parentalcontrol:blockingpolicy object
[**Delete**](ParentalcontrolBlockingpolicyAPI.md#Delete) | **Delete** /parentalcontrol:blockingpolicy/{reference} | Delete a parentalcontrol:blockingpolicy object
[**List**](ParentalcontrolBlockingpolicyAPI.md#List) | **Get** /parentalcontrol:blockingpolicy | Retrieve parentalcontrol:blockingpolicy objects
[**Read**](ParentalcontrolBlockingpolicyAPI.md#Read) | **Get** /parentalcontrol:blockingpolicy/{reference} | Get a specific parentalcontrol:blockingpolicy object
[**Update**](ParentalcontrolBlockingpolicyAPI.md#Update) | **Put** /parentalcontrol:blockingpolicy/{reference} | Update a parentalcontrol:blockingpolicy object



## Create

> CreateParentalcontrolBlockingpolicyResponse Create(ctx).ParentalcontrolBlockingpolicy(parentalcontrolBlockingpolicy).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a parentalcontrol:blockingpolicy object



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
	parentalcontrolBlockingpolicy := *parentalcontrol.NewParentalcontrolBlockingpolicy() // ParentalcontrolBlockingpolicy | Object data to create

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolBlockingpolicyAPI.Create(context.Background()).ParentalcontrolBlockingpolicy(parentalcontrolBlockingpolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateParentalcontrolBlockingpolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolBlockingpolicyAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolBlockingpolicy** | [**ParentalcontrolBlockingpolicy**](ParentalcontrolBlockingpolicy.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateParentalcontrolBlockingpolicyResponse**](CreateParentalcontrolBlockingpolicyResponse.md)

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

Delete a parentalcontrol:blockingpolicy object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:blockingpolicy object

	apiClient := parentalcontrol.NewAPIClient()
	r, err := apiClient.ParentalcontrolBlockingpolicyAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:blockingpolicy object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPIDeleteRequest` struct via the builder pattern


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

> ListParentalcontrolBlockingpolicyResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve parentalcontrol:blockingpolicy objects



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
	resp, r, err := apiClient.ParentalcontrolBlockingpolicyAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListParentalcontrolBlockingpolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolBlockingpolicyAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPIListRequest` struct via the builder pattern


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

[**ListParentalcontrolBlockingpolicyResponse**](ListParentalcontrolBlockingpolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetParentalcontrolBlockingpolicyResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific parentalcontrol:blockingpolicy object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:blockingpolicy object

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolBlockingpolicyAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetParentalcontrolBlockingpolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolBlockingpolicyAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:blockingpolicy object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetParentalcontrolBlockingpolicyResponse**](GetParentalcontrolBlockingpolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateParentalcontrolBlockingpolicyResponse Update(ctx, reference).ParentalcontrolBlockingpolicy(parentalcontrolBlockingpolicy).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a parentalcontrol:blockingpolicy object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:blockingpolicy object
	parentalcontrolBlockingpolicy := *parentalcontrol.NewParentalcontrolBlockingpolicy() // ParentalcontrolBlockingpolicy | Object data to update

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolBlockingpolicyAPI.Update(context.Background(), reference).ParentalcontrolBlockingpolicy(parentalcontrolBlockingpolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateParentalcontrolBlockingpolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolBlockingpolicyAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:blockingpolicy object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolBlockingpolicy** | [**ParentalcontrolBlockingpolicy**](ParentalcontrolBlockingpolicy.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateParentalcontrolBlockingpolicyResponse**](UpdateParentalcontrolBlockingpolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

