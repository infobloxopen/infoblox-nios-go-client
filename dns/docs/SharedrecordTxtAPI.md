# SharedrecordTxtAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharedrecordtxtGet**](SharedrecordTxtAPI.md#SharedrecordtxtGet) | **Get** /sharedrecord:txt | Retrieve sharedrecord:txt objects
[**SharedrecordtxtPost**](SharedrecordTxtAPI.md#SharedrecordtxtPost) | **Post** /sharedrecord:txt | Create a sharedrecord:txt object
[**SharedrecordtxtReferenceDelete**](SharedrecordTxtAPI.md#SharedrecordtxtReferenceDelete) | **Delete** /sharedrecord:txt/{reference} | Delete a sharedrecord:txt object
[**SharedrecordtxtReferenceGet**](SharedrecordTxtAPI.md#SharedrecordtxtReferenceGet) | **Get** /sharedrecord:txt/{reference} | Get a specific sharedrecord:txt object
[**SharedrecordtxtReferencePut**](SharedrecordTxtAPI.md#SharedrecordtxtReferencePut) | **Put** /sharedrecord:txt/{reference} | Update a sharedrecord:txt object



## SharedrecordtxtGet

> ListSharedrecordTxtResponse SharedrecordtxtGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve sharedrecord:txt objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordTxtAPI.SharedrecordtxtGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordTxtAPI.SharedrecordtxtGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordtxtGet`: ListSharedrecordTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordTxtAPI.SharedrecordtxtGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordTxtAPISharedrecordtxtGetRequest` struct via the builder pattern


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

[**ListSharedrecordTxtResponse**](ListSharedrecordTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordtxtPost

> CreateSharedrecordTxtResponse SharedrecordtxtPost(ctx).SharedrecordTxt(sharedrecordTxt).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a sharedrecord:txt object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	sharedrecordTxt := *dns.NewSharedrecordTxt() // SharedrecordTxt | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordTxtAPI.SharedrecordtxtPost(context.Background()).SharedrecordTxt(sharedrecordTxt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordTxtAPI.SharedrecordtxtPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordtxtPost`: CreateSharedrecordTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordTxtAPI.SharedrecordtxtPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordTxtAPISharedrecordtxtPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordTxt** | [**SharedrecordTxt**](SharedrecordTxt.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSharedrecordTxtResponse**](CreateSharedrecordTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordtxtReferenceDelete

> SharedrecordtxtReferenceDelete(ctx, reference).Execute()

Delete a sharedrecord:txt object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the sharedrecord:txt object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.SharedrecordTxtAPI.SharedrecordtxtReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordTxtAPI.SharedrecordtxtReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:txt object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordTxtAPISharedrecordtxtReferenceDeleteRequest` struct via the builder pattern


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


## SharedrecordtxtReferenceGet

> GetSharedrecordTxtResponse SharedrecordtxtReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific sharedrecord:txt object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the sharedrecord:txt object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordTxtAPI.SharedrecordtxtReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordTxtAPI.SharedrecordtxtReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordtxtReferenceGet`: GetSharedrecordTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordTxtAPI.SharedrecordtxtReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:txt object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordTxtAPISharedrecordtxtReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSharedrecordTxtResponse**](GetSharedrecordTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordtxtReferencePut

> UpdateSharedrecordTxtResponse SharedrecordtxtReferencePut(ctx, reference).SharedrecordTxt(sharedrecordTxt).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a sharedrecord:txt object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the sharedrecord:txt object
	sharedrecordTxt := *dns.NewSharedrecordTxt() // SharedrecordTxt | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordTxtAPI.SharedrecordtxtReferencePut(context.Background(), reference).SharedrecordTxt(sharedrecordTxt).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordTxtAPI.SharedrecordtxtReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordtxtReferencePut`: UpdateSharedrecordTxtResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordTxtAPI.SharedrecordtxtReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:txt object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordTxtAPISharedrecordtxtReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordTxt** | [**SharedrecordTxt**](SharedrecordTxt.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSharedrecordTxtResponse**](UpdateSharedrecordTxtResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

