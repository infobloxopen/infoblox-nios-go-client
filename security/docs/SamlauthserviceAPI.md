# SamlAuthserviceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SamlauthserviceGet**](SamlAuthserviceAPI.md#SamlauthserviceGet) | **Get** /saml:authservice | Retrieve saml:authservice objects
[**SamlauthservicePost**](SamlAuthserviceAPI.md#SamlauthservicePost) | **Post** /saml:authservice | Create a saml:authservice object
[**SamlauthserviceReferenceDelete**](SamlAuthserviceAPI.md#SamlauthserviceReferenceDelete) | **Delete** /saml:authservice/{reference} | Delete a saml:authservice object
[**SamlauthserviceReferenceGet**](SamlAuthserviceAPI.md#SamlauthserviceReferenceGet) | **Get** /saml:authservice/{reference} | Get a specific saml:authservice object
[**SamlauthserviceReferencePut**](SamlAuthserviceAPI.md#SamlauthserviceReferencePut) | **Put** /saml:authservice/{reference} | Update a saml:authservice object



## SamlauthserviceGet

> ListSamlAuthserviceResponse SamlauthserviceGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve saml:authservice objects



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
	resp, r, err := apiClient.SamlAuthserviceAPI.SamlauthserviceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAuthserviceAPI.SamlauthserviceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SamlauthserviceGet`: ListSamlAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `SamlAuthserviceAPI.SamlauthserviceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SamlAuthserviceAPISamlauthserviceGetRequest` struct via the builder pattern


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

[**ListSamlAuthserviceResponse**](ListSamlAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamlauthservicePost

> CreateSamlAuthserviceResponse SamlauthservicePost(ctx).SamlAuthservice(samlAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a saml:authservice object



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
	samlAuthservice := *security.NewSamlAuthservice() // SamlAuthservice | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.SamlAuthserviceAPI.SamlauthservicePost(context.Background()).SamlAuthservice(samlAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAuthserviceAPI.SamlauthservicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SamlauthservicePost`: CreateSamlAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `SamlAuthserviceAPI.SamlauthservicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SamlAuthserviceAPISamlauthservicePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**samlAuthservice** | [**SamlAuthservice**](SamlAuthservice.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSamlAuthserviceResponse**](CreateSamlAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamlauthserviceReferenceDelete

> SamlauthserviceReferenceDelete(ctx, reference).Execute()

Delete a saml:authservice object



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
	reference := "reference_example" // string | Reference of the saml:authservice object

	apiClient := security.NewAPIClient()
	r, err := apiClient.SamlAuthserviceAPI.SamlauthserviceReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAuthserviceAPI.SamlauthserviceReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the saml:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `SamlAuthserviceAPISamlauthserviceReferenceDeleteRequest` struct via the builder pattern


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


## SamlauthserviceReferenceGet

> GetSamlAuthserviceResponse SamlauthserviceReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific saml:authservice object



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
	reference := "reference_example" // string | Reference of the saml:authservice object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.SamlAuthserviceAPI.SamlauthserviceReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAuthserviceAPI.SamlauthserviceReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SamlauthserviceReferenceGet`: GetSamlAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `SamlAuthserviceAPI.SamlauthserviceReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the saml:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `SamlAuthserviceAPISamlauthserviceReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSamlAuthserviceResponse**](GetSamlAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SamlauthserviceReferencePut

> UpdateSamlAuthserviceResponse SamlauthserviceReferencePut(ctx, reference).SamlAuthservice(samlAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a saml:authservice object



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
	reference := "reference_example" // string | Reference of the saml:authservice object
	samlAuthservice := *security.NewSamlAuthservice() // SamlAuthservice | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.SamlAuthserviceAPI.SamlauthserviceReferencePut(context.Background(), reference).SamlAuthservice(samlAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SamlAuthserviceAPI.SamlauthserviceReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SamlauthserviceReferencePut`: UpdateSamlAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `SamlAuthserviceAPI.SamlauthserviceReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the saml:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `SamlAuthserviceAPISamlauthserviceReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**samlAuthservice** | [**SamlAuthservice**](SamlAuthservice.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSamlAuthserviceResponse**](UpdateSamlAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

