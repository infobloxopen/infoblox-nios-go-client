# ParentalcontrolSubscriberAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ParentalcontrolsubscriberGet**](ParentalcontrolSubscriberAPI.md#ParentalcontrolsubscriberGet) | **Get** /parentalcontrol:subscriber | Retrieve parentalcontrol:subscriber objects
[**ParentalcontrolsubscriberReferenceGet**](ParentalcontrolSubscriberAPI.md#ParentalcontrolsubscriberReferenceGet) | **Get** /parentalcontrol:subscriber/{reference} | Get a specific parentalcontrol:subscriber object
[**ParentalcontrolsubscriberReferencePut**](ParentalcontrolSubscriberAPI.md#ParentalcontrolsubscriberReferencePut) | **Put** /parentalcontrol:subscriber/{reference} | Update a parentalcontrol:subscriber object



## ParentalcontrolsubscriberGet

> ListParentalcontrolSubscriberResponse ParentalcontrolsubscriberGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve parentalcontrol:subscriber objects



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
	resp, r, err := apiClient.ParentalcontrolSubscriberAPI.ParentalcontrolsubscriberGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberAPI.ParentalcontrolsubscriberGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscriberGet`: ListParentalcontrolSubscriberResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberAPI.ParentalcontrolsubscriberGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberAPIParentalcontrolsubscriberGetRequest` struct via the builder pattern


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

[**ListParentalcontrolSubscriberResponse**](ListParentalcontrolSubscriberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentalcontrolsubscriberReferenceGet

> GetParentalcontrolSubscriberResponse ParentalcontrolsubscriberReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific parentalcontrol:subscriber object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriber object

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberAPI.ParentalcontrolsubscriberReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberAPI.ParentalcontrolsubscriberReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscriberReferenceGet`: GetParentalcontrolSubscriberResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberAPI.ParentalcontrolsubscriberReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscriber object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberAPIParentalcontrolsubscriberReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

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


## ParentalcontrolsubscriberReferencePut

> UpdateParentalcontrolSubscriberResponse ParentalcontrolsubscriberReferencePut(ctx, reference).ParentalcontrolSubscriber(parentalcontrolSubscriber).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a parentalcontrol:subscriber object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:subscriber object
	parentalcontrolSubscriber := *parentalcontrol.NewParentalcontrolSubscriber() // ParentalcontrolSubscriber | Object data to update

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolSubscriberAPI.ParentalcontrolsubscriberReferencePut(context.Background(), reference).ParentalcontrolSubscriber(parentalcontrolSubscriber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolSubscriberAPI.ParentalcontrolsubscriberReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolsubscriberReferencePut`: UpdateParentalcontrolSubscriberResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolSubscriberAPI.ParentalcontrolsubscriberReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:subscriber object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolSubscriberAPIParentalcontrolsubscriberReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolSubscriber** | [**ParentalcontrolSubscriber**](ParentalcontrolSubscriber.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

