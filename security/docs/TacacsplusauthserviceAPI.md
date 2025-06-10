# TacacsplusAuthserviceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TacacsplusauthserviceGet**](TacacsplusAuthserviceAPI.md#TacacsplusauthserviceGet) | **Get** /tacacsplus:authservice | Retrieve tacacsplus:authservice objects
[**TacacsplusauthservicePost**](TacacsplusAuthserviceAPI.md#TacacsplusauthservicePost) | **Post** /tacacsplus:authservice | Create a tacacsplus:authservice object
[**TacacsplusauthserviceReferenceDelete**](TacacsplusAuthserviceAPI.md#TacacsplusauthserviceReferenceDelete) | **Delete** /tacacsplus:authservice/{reference} | Delete a tacacsplus:authservice object
[**TacacsplusauthserviceReferenceGet**](TacacsplusAuthserviceAPI.md#TacacsplusauthserviceReferenceGet) | **Get** /tacacsplus:authservice/{reference} | Get a specific tacacsplus:authservice object
[**TacacsplusauthserviceReferencePut**](TacacsplusAuthserviceAPI.md#TacacsplusauthserviceReferencePut) | **Put** /tacacsplus:authservice/{reference} | Update a tacacsplus:authservice object



## TacacsplusauthserviceGet

> ListTacacsplusAuthserviceResponse TacacsplusauthserviceGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve tacacsplus:authservice objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.TacacsplusAuthserviceAPI.TacacsplusauthserviceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.TacacsplusauthserviceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TacacsplusauthserviceGet`: ListTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusAuthserviceAPI.TacacsplusauthserviceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPITacacsplusauthserviceGetRequest` struct via the builder pattern


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

[**ListTacacsplusAuthserviceResponse**](ListTacacsplusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TacacsplusauthservicePost

> CreateTacacsplusAuthserviceResponse TacacsplusauthservicePost(ctx).TacacsplusAuthservice(tacacsplusAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a tacacsplus:authservice object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {
	tacacsplusAuthservice := *security.NewTacacsplusAuthservice() // TacacsplusAuthservice | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.TacacsplusAuthserviceAPI.TacacsplusauthservicePost(context.Background()).TacacsplusAuthservice(tacacsplusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.TacacsplusauthservicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TacacsplusauthservicePost`: CreateTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusAuthserviceAPI.TacacsplusauthservicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPITacacsplusauthservicePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**tacacsplusAuthservice** | [**TacacsplusAuthservice**](TacacsplusAuthservice.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateTacacsplusAuthserviceResponse**](CreateTacacsplusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TacacsplusauthserviceReferenceDelete

> TacacsplusauthserviceReferenceDelete(ctx, reference).Execute()

Delete a tacacsplus:authservice object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the tacacsplus:authservice object

	apiClient := security.NewAPIClient()
	r, err := apiClient.TacacsplusAuthserviceAPI.TacacsplusauthserviceReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.TacacsplusauthserviceReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the tacacsplus:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPITacacsplusauthserviceReferenceDeleteRequest` struct via the builder pattern


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


## TacacsplusauthserviceReferenceGet

> GetTacacsplusAuthserviceResponse TacacsplusauthserviceReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific tacacsplus:authservice object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the tacacsplus:authservice object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.TacacsplusAuthserviceAPI.TacacsplusauthserviceReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.TacacsplusauthserviceReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TacacsplusauthserviceReferenceGet`: GetTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusAuthserviceAPI.TacacsplusauthserviceReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the tacacsplus:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPITacacsplusauthserviceReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetTacacsplusAuthserviceResponse**](GetTacacsplusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TacacsplusauthserviceReferencePut

> UpdateTacacsplusAuthserviceResponse TacacsplusauthserviceReferencePut(ctx, reference).TacacsplusAuthservice(tacacsplusAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a tacacsplus:authservice object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the tacacsplus:authservice object
	tacacsplusAuthservice := *security.NewTacacsplusAuthservice() // TacacsplusAuthservice | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.TacacsplusAuthserviceAPI.TacacsplusauthserviceReferencePut(context.Background(), reference).TacacsplusAuthservice(tacacsplusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusAuthserviceAPI.TacacsplusauthserviceReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TacacsplusauthserviceReferencePut`: UpdateTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusAuthserviceAPI.TacacsplusauthserviceReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the tacacsplus:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusAuthserviceAPITacacsplusauthserviceReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**tacacsplusAuthservice** | [**TacacsplusAuthservice**](TacacsplusAuthservice.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateTacacsplusAuthserviceResponse**](UpdateTacacsplusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

