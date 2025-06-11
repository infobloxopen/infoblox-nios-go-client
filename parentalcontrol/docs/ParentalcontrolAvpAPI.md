# ParentalcontrolAvpAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ParentalcontrolavpGet**](ParentalcontrolAvpAPI.md#ParentalcontrolavpGet) | **Get** /parentalcontrol:avp | Retrieve parentalcontrol:avp objects
[**ParentalcontrolavpPost**](ParentalcontrolAvpAPI.md#ParentalcontrolavpPost) | **Post** /parentalcontrol:avp | Create a parentalcontrol:avp object
[**ParentalcontrolavpReferenceDelete**](ParentalcontrolAvpAPI.md#ParentalcontrolavpReferenceDelete) | **Delete** /parentalcontrol:avp/{reference} | Delete a parentalcontrol:avp object
[**ParentalcontrolavpReferenceGet**](ParentalcontrolAvpAPI.md#ParentalcontrolavpReferenceGet) | **Get** /parentalcontrol:avp/{reference} | Get a specific parentalcontrol:avp object
[**ParentalcontrolavpReferencePut**](ParentalcontrolAvpAPI.md#ParentalcontrolavpReferencePut) | **Put** /parentalcontrol:avp/{reference} | Update a parentalcontrol:avp object



## ParentalcontrolavpGet

> ListParentalcontrolAvpResponse ParentalcontrolavpGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve parentalcontrol:avp objects



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
	resp, r, err := apiClient.ParentalcontrolAvpAPI.ParentalcontrolavpGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolAvpAPI.ParentalcontrolavpGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolavpGet`: ListParentalcontrolAvpResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolAvpAPI.ParentalcontrolavpGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolAvpAPIParentalcontrolavpGetRequest` struct via the builder pattern


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

[**ListParentalcontrolAvpResponse**](ListParentalcontrolAvpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentalcontrolavpPost

> CreateParentalcontrolAvpResponse ParentalcontrolavpPost(ctx).ParentalcontrolAvp(parentalcontrolAvp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a parentalcontrol:avp object



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
	parentalcontrolAvp := *parentalcontrol.NewParentalcontrolAvp() // ParentalcontrolAvp | Object data to create

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolAvpAPI.ParentalcontrolavpPost(context.Background()).ParentalcontrolAvp(parentalcontrolAvp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolAvpAPI.ParentalcontrolavpPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolavpPost`: CreateParentalcontrolAvpResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolAvpAPI.ParentalcontrolavpPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolAvpAPIParentalcontrolavpPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolAvp** | [**ParentalcontrolAvp**](ParentalcontrolAvp.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateParentalcontrolAvpResponse**](CreateParentalcontrolAvpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentalcontrolavpReferenceDelete

> ParentalcontrolavpReferenceDelete(ctx, reference).Execute()

Delete a parentalcontrol:avp object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:avp object

	apiClient := parentalcontrol.NewAPIClient()
	r, err := apiClient.ParentalcontrolAvpAPI.ParentalcontrolavpReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolAvpAPI.ParentalcontrolavpReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:avp object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolAvpAPIParentalcontrolavpReferenceDeleteRequest` struct via the builder pattern


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


## ParentalcontrolavpReferenceGet

> GetParentalcontrolAvpResponse ParentalcontrolavpReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific parentalcontrol:avp object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:avp object

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolAvpAPI.ParentalcontrolavpReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolAvpAPI.ParentalcontrolavpReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolavpReferenceGet`: GetParentalcontrolAvpResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolAvpAPI.ParentalcontrolavpReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:avp object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolAvpAPIParentalcontrolavpReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetParentalcontrolAvpResponse**](GetParentalcontrolAvpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentalcontrolavpReferencePut

> UpdateParentalcontrolAvpResponse ParentalcontrolavpReferencePut(ctx, reference).ParentalcontrolAvp(parentalcontrolAvp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a parentalcontrol:avp object



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
	reference := "reference_example" // string | Reference of the parentalcontrol:avp object
	parentalcontrolAvp := *parentalcontrol.NewParentalcontrolAvp() // ParentalcontrolAvp | Object data to update

	apiClient := parentalcontrol.NewAPIClient()
	resp, r, err := apiClient.ParentalcontrolAvpAPI.ParentalcontrolavpReferencePut(context.Background(), reference).ParentalcontrolAvp(parentalcontrolAvp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParentalcontrolAvpAPI.ParentalcontrolavpReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentalcontrolavpReferencePut`: UpdateParentalcontrolAvpResponse
	fmt.Fprintf(os.Stdout, "Response from `ParentalcontrolAvpAPI.ParentalcontrolavpReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the parentalcontrol:avp object | 

### Other Parameters

Other parameters are passed through a pointer to a `ParentalcontrolAvpAPIParentalcontrolavpReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**parentalcontrolAvp** | [**ParentalcontrolAvp**](ParentalcontrolAvp.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateParentalcontrolAvpResponse**](UpdateParentalcontrolAvpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

