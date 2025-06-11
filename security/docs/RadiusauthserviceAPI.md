# RadiusAuthserviceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RadiusauthserviceGet**](RadiusAuthserviceAPI.md#RadiusauthserviceGet) | **Get** /radius:authservice | Retrieve radius:authservice objects
[**RadiusauthservicePost**](RadiusAuthserviceAPI.md#RadiusauthservicePost) | **Post** /radius:authservice | Create a radius:authservice object
[**RadiusauthserviceReferenceDelete**](RadiusAuthserviceAPI.md#RadiusauthserviceReferenceDelete) | **Delete** /radius:authservice/{reference} | Delete a radius:authservice object
[**RadiusauthserviceReferenceGet**](RadiusAuthserviceAPI.md#RadiusauthserviceReferenceGet) | **Get** /radius:authservice/{reference} | Get a specific radius:authservice object
[**RadiusauthserviceReferencePut**](RadiusAuthserviceAPI.md#RadiusauthserviceReferencePut) | **Put** /radius:authservice/{reference} | Update a radius:authservice object



## RadiusauthserviceGet

> ListRadiusAuthserviceResponse RadiusauthserviceGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve radius:authservice objects



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
	resp, r, err := apiClient.RadiusAuthserviceAPI.RadiusauthserviceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.RadiusauthserviceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadiusauthserviceGet`: ListRadiusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `RadiusAuthserviceAPI.RadiusauthserviceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RadiusAuthserviceAPIRadiusauthserviceGetRequest` struct via the builder pattern


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

[**ListRadiusAuthserviceResponse**](ListRadiusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadiusauthservicePost

> CreateRadiusAuthserviceResponse RadiusauthservicePost(ctx).RadiusAuthservice(radiusAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a radius:authservice object



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
	radiusAuthservice := *security.NewRadiusAuthservice() // RadiusAuthservice | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.RadiusAuthserviceAPI.RadiusauthservicePost(context.Background()).RadiusAuthservice(radiusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.RadiusauthservicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadiusauthservicePost`: CreateRadiusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `RadiusAuthserviceAPI.RadiusauthservicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RadiusAuthserviceAPIRadiusauthservicePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**radiusAuthservice** | [**RadiusAuthservice**](RadiusAuthservice.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRadiusAuthserviceResponse**](CreateRadiusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadiusauthserviceReferenceDelete

> RadiusauthserviceReferenceDelete(ctx, reference).Execute()

Delete a radius:authservice object



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
	reference := "reference_example" // string | Reference of the radius:authservice object

	apiClient := security.NewAPIClient()
	r, err := apiClient.RadiusAuthserviceAPI.RadiusauthserviceReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.RadiusauthserviceReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the radius:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `RadiusAuthserviceAPIRadiusauthserviceReferenceDeleteRequest` struct via the builder pattern


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


## RadiusauthserviceReferenceGet

> GetRadiusAuthserviceResponse RadiusauthserviceReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific radius:authservice object



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
	reference := "reference_example" // string | Reference of the radius:authservice object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.RadiusAuthserviceAPI.RadiusauthserviceReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.RadiusauthserviceReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadiusauthserviceReferenceGet`: GetRadiusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `RadiusAuthserviceAPI.RadiusauthserviceReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the radius:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `RadiusAuthserviceAPIRadiusauthserviceReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRadiusAuthserviceResponse**](GetRadiusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RadiusauthserviceReferencePut

> UpdateRadiusAuthserviceResponse RadiusauthserviceReferencePut(ctx, reference).RadiusAuthservice(radiusAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a radius:authservice object



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
	reference := "reference_example" // string | Reference of the radius:authservice object
	radiusAuthservice := *security.NewRadiusAuthservice() // RadiusAuthservice | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.RadiusAuthserviceAPI.RadiusauthserviceReferencePut(context.Background(), reference).RadiusAuthservice(radiusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.RadiusauthserviceReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RadiusauthserviceReferencePut`: UpdateRadiusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `RadiusAuthserviceAPI.RadiusauthserviceReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the radius:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `RadiusAuthserviceAPIRadiusauthserviceReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**radiusAuthservice** | [**RadiusAuthservice**](RadiusAuthservice.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRadiusAuthserviceResponse**](UpdateRadiusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

