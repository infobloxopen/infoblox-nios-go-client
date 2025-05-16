# TacacsplusauthserviceAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](TacacsplusauthserviceAPI.md#Get) | **Get** /tacacsplus:authservice | Retrieve tacacsplus:authservice objects
[**Post**](TacacsplusauthserviceAPI.md#Post) | **Post** /tacacsplus:authservice | Create a tacacsplus:authservice object
[**ReferenceDelete**](TacacsplusauthserviceAPI.md#ReferenceDelete) | **Delete** /tacacsplus:authservice/{reference} | Delete a tacacsplus:authservice object
[**ReferenceGet**](TacacsplusauthserviceAPI.md#ReferenceGet) | **Get** /tacacsplus:authservice/{reference} | Get a specific tacacsplus:authservice object
[**ReferencePut**](TacacsplusauthserviceAPI.md#ReferencePut) | **Put** /tacacsplus:authservice/{reference} | Update a tacacsplus:authservice object



## Get

> ListTacacsplusAuthserviceResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

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
	resp, r, err := apiClient.TacacsplusauthserviceAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusauthserviceAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusauthserviceAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusauthserviceAPIGetRequest` struct via the builder pattern


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


## Post

> CreateTacacsplusAuthserviceResponse Post(ctx).TacacsplusAuthservice(tacacsplusAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.TacacsplusauthserviceAPI.Post(context.Background()).TacacsplusAuthservice(tacacsplusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusauthserviceAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusauthserviceAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusauthserviceAPIPostRequest` struct via the builder pattern


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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

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
	r, err := apiClient.TacacsplusauthserviceAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusauthserviceAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `TacacsplusauthserviceAPIReferenceDeleteRequest` struct via the builder pattern


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


## ReferenceGet

> GetTacacsplusAuthserviceResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.TacacsplusauthserviceAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusauthserviceAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusauthserviceAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the tacacsplus:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusauthserviceAPIReferenceGetRequest` struct via the builder pattern


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


## ReferencePut

> UpdateTacacsplusAuthserviceResponse ReferencePut(ctx, reference).TacacsplusAuthservice(tacacsplusAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.TacacsplusauthserviceAPI.ReferencePut(context.Background(), reference).TacacsplusAuthservice(tacacsplusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TacacsplusauthserviceAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateTacacsplusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `TacacsplusauthserviceAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the tacacsplus:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `TacacsplusauthserviceAPIReferencePutRequest` struct via the builder pattern


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

