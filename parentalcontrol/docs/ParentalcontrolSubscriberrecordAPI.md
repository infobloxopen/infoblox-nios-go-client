# ParentalcontrolSubscriberrecordAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ParentalcontrolsubscriberrecordGet**](ParentalcontrolSubscriberrecordAPI.md#ParentalcontrolsubscriberrecordGet) | **Get** /parentalcontrol:subscriberrecord | Retrieve parentalcontrol:subscriberrecord objects
[**ParentalcontrolsubscriberrecordPost**](ParentalcontrolSubscriberrecordAPI.md#ParentalcontrolsubscriberrecordPost) | **Post** /parentalcontrol:subscriberrecord | Create a parentalcontrol:subscriberrecord object
[**ParentalcontrolsubscriberrecordReferenceDelete**](ParentalcontrolSubscriberrecordAPI.md#ParentalcontrolsubscriberrecordReferenceDelete) | **Delete** /parentalcontrol:subscriberrecord/{reference} | Delete a parentalcontrol:subscriberrecord object
[**ParentalcontrolsubscriberrecordReferenceGet**](ParentalcontrolSubscriberrecordAPI.md#ParentalcontrolsubscriberrecordReferenceGet) | **Get** /parentalcontrol:subscriberrecord/{reference} | Get a specific parentalcontrol:subscriberrecord object
[**ParentalcontrolsubscriberrecordReferencePut**](ParentalcontrolSubscriberrecordAPI.md#ParentalcontrolsubscriberrecordReferencePut) | **Put** /parentalcontrol:subscriberrecord/{reference} | Update a parentalcontrol:subscriberrecord object



## ParentalcontrolsubscriberrecordGet

> ListParentalcontrolSubscriberrecordResponse ParentalcontrolsubscriberrecordGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve parentalcontrol:subscriberrecord objects



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
	resp, r, err := apiClient.ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscriberrecordGet`: ListParentalcontrolSubscriberrecordResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPIParentalcontrolsubscriberrecordGetRequest` struct via the builder pattern


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

[**ListParentalcontrolSubscriberrecordResponse**](ListParentalcontrolSubscriberrecordResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentalcontrolsubscriberrecordPost

> CreateParentalcontrolSubscriberrecordResponse ParentalcontrolsubscriberrecordPost(ctx).ParentalcontrolSubscriberrecord(parentalcontrolSubscriberrecord).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a parentalcontrol:subscriberrecord object



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
	parentalcontrolSubscriberrecord := *parentalcontrol.NewParentalcontrolSubscriberrecord() // ParentalcontrolSubscriberrecord | Object data to create

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordPost(context.Background()).ParentalcontrolSubscriberrecord(parentalcontrolSubscriberrecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscriberrecordPost`: CreateParentalcontrolSubscriberrecordResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPIParentalcontrolsubscriberrecordPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolSubscriberrecord** | [**ParentalcontrolSubscriberrecord**](ParentalcontrolSubscriberrecord.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## ParentalcontrolsubscriberrecordReferenceDelete

> ParentalcontrolsubscriberrecordReferenceDelete(ctx, reference).Execute()

Delete a parentalcontrol:subscriberrecord object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriberrecord object

	apiClient := parentalcontrol.NewAPIClient()
	r, err := apiClient.ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPIParentalcontrolsubscriberrecordReferenceDeleteRequest` struct via the builder pattern


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


## ParentalcontrolsubscriberrecordReferenceGet

> GetParentalcontrolSubscriberrecordResponse ParentalcontrolsubscriberrecordReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific parentalcontrol:subscriberrecord object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriberrecord object

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscriberrecordReferenceGet`: GetParentalcontrolSubscriberrecordResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscriberrecord object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPIParentalcontrolsubscriberrecordReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## ParentalcontrolsubscriberrecordReferencePut

> UpdateParentalcontrolSubscriberrecordResponse ParentalcontrolsubscriberrecordReferencePut(ctx, reference).ParentalcontrolSubscriberrecord(parentalcontrolSubscriberrecord).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a parentalcontrol:subscriberrecord object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriberrecord object
	parentalcontrolSubscriberrecord := *parentalcontrol.NewParentalcontrolSubscriberrecord() // ParentalcontrolSubscriberrecord | Object data to update

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordReferencePut(context.Background(), reference).ParentalcontrolSubscriberrecord(parentalcontrolSubscriberrecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscriberrecordReferencePut`: UpdateParentalcontrolSubscriberrecordResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberrecordAPI.ParentalcontrolsubscriberrecordReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscriberrecord object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberrecordAPIParentalcontrolsubscriberrecordReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolSubscriberrecord** | [**ParentalcontrolSubscriberrecord**](ParentalcontrolSubscriberrecord.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

