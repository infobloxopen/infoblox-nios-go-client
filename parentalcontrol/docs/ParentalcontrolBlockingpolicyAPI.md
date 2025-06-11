# ParentalcontrolBlockingpolicyAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ParentalcontrolblockingpolicyGet**](ParentalcontrolBlockingpolicyAPI.md#ParentalcontrolblockingpolicyGet) | **Get** /parentalcontrol:blockingpolicy | Retrieve parentalcontrol:blockingpolicy objects
[**ParentalcontrolblockingpolicyPost**](ParentalcontrolBlockingpolicyAPI.md#ParentalcontrolblockingpolicyPost) | **Post** /parentalcontrol:blockingpolicy | Create a parentalcontrol:blockingpolicy object
[**ParentalcontrolblockingpolicyReferenceDelete**](ParentalcontrolBlockingpolicyAPI.md#ParentalcontrolblockingpolicyReferenceDelete) | **Delete** /parentalcontrol:blockingpolicy/{reference} | Delete a parentalcontrol:blockingpolicy object
[**ParentalcontrolblockingpolicyReferenceGet**](ParentalcontrolBlockingpolicyAPI.md#ParentalcontrolblockingpolicyReferenceGet) | **Get** /parentalcontrol:blockingpolicy/{reference} | Get a specific parentalcontrol:blockingpolicy object
[**ParentalcontrolblockingpolicyReferencePut**](ParentalcontrolBlockingpolicyAPI.md#ParentalcontrolblockingpolicyReferencePut) | **Put** /parentalcontrol:blockingpolicy/{reference} | Update a parentalcontrol:blockingpolicy object



## ParentalcontrolblockingpolicyGet

> ListParentalcontrolBlockingpolicyResponse ParentalcontrolblockingpolicyGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve parentalcontrol:blockingpolicy objects



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
	resp, r, err := apiClient.ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolblockingpolicyGet`: ListParentalcontrolBlockingpolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPIParentalcontrolblockingpolicyGetRequest` struct via the builder pattern


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

[**ListParentalcontrolBlockingpolicyResponse**](ListParentalcontrolBlockingpolicyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentalcontrolblockingpolicyPost

> CreateParentalcontrolBlockingpolicyResponse ParentalcontrolblockingpolicyPost(ctx).ParentalcontrolBlockingpolicy(parentalcontrolBlockingpolicy).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a parentalcontrol:blockingpolicy object



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
	parentalcontrolBlockingpolicy := *parentalcontrol.NewParentalcontrolBlockingpolicy() // ParentalcontrolBlockingpolicy | Object data to create

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyPost(context.Background()).ParentalcontrolBlockingpolicy(parentalcontrolBlockingpolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolblockingpolicyPost`: CreateParentalcontrolBlockingpolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPIParentalcontrolblockingpolicyPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolBlockingpolicy** | [**ParentalcontrolBlockingpolicy**](ParentalcontrolBlockingpolicy.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## ParentalcontrolblockingpolicyReferenceDelete

> ParentalcontrolblockingpolicyReferenceDelete(ctx, reference).Execute()

Delete a parentalcontrol:blockingpolicy object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:blockingpolicy object

	apiClient := parentalcontrol.NewAPIClient()
	r, err := apiClient.ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPIParentalcontrolblockingpolicyReferenceDeleteRequest` struct via the builder pattern


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


## ParentalcontrolblockingpolicyReferenceGet

> GetParentalcontrolBlockingpolicyResponse ParentalcontrolblockingpolicyReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific parentalcontrol:blockingpolicy object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:blockingpolicy object

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolblockingpolicyReferenceGet`: GetParentalcontrolBlockingpolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:blockingpolicy object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPIParentalcontrolblockingpolicyReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## ParentalcontrolblockingpolicyReferencePut

> UpdateParentalcontrolBlockingpolicyResponse ParentalcontrolblockingpolicyReferencePut(ctx, reference).ParentalcontrolBlockingpolicy(parentalcontrolBlockingpolicy).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a parentalcontrol:blockingpolicy object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:blockingpolicy object
	parentalcontrolBlockingpolicy := *parentalcontrol.NewParentalcontrolBlockingpolicy() // ParentalcontrolBlockingpolicy | Object data to update

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyReferencePut(context.Background(), reference).ParentalcontrolBlockingpolicy(parentalcontrolBlockingpolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolblockingpolicyReferencePut`: UpdateParentalcontrolBlockingpolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolBlockingpolicyAPI.ParentalcontrolblockingpolicyReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:blockingpolicy object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolBlockingpolicyAPIParentalcontrolblockingpolicyReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolBlockingpolicy** | [**ParentalcontrolBlockingpolicy**](ParentalcontrolBlockingpolicy.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

