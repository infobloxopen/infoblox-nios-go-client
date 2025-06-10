# CertificateAuthserviceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CertificateauthserviceGet**](CertificateAuthserviceAPI.md#CertificateauthserviceGet) | **Get** /certificate:authservice | Retrieve certificate:authservice objects
[**CertificateauthservicePost**](CertificateAuthserviceAPI.md#CertificateauthservicePost) | **Post** /certificate:authservice | Create a certificate:authservice object
[**CertificateauthserviceReferenceDelete**](CertificateAuthserviceAPI.md#CertificateauthserviceReferenceDelete) | **Delete** /certificate:authservice/{reference} | Delete a certificate:authservice object
[**CertificateauthserviceReferenceGet**](CertificateAuthserviceAPI.md#CertificateauthserviceReferenceGet) | **Get** /certificate:authservice/{reference} | Get a specific certificate:authservice object
[**CertificateauthserviceReferencePut**](CertificateAuthserviceAPI.md#CertificateauthserviceReferencePut) | **Put** /certificate:authservice/{reference} | Update a certificate:authservice object



## CertificateauthserviceGet

> ListCertificateAuthserviceResponse CertificateauthserviceGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve certificate:authservice objects



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
	resp, r, err := apiClient.CertificateAuthserviceAPI.CertificateauthserviceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.CertificateauthserviceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateauthserviceGet`: ListCertificateAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthserviceAPI.CertificateauthserviceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CertificateAuthserviceAPICertificateauthserviceGetRequest` struct via the builder pattern


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

[**ListCertificateAuthserviceResponse**](ListCertificateAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateauthservicePost

> CreateCertificateAuthserviceResponse CertificateauthservicePost(ctx).CertificateAuthservice(certificateAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a certificate:authservice object



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
	certificateAuthservice := *security.NewCertificateAuthservice() // CertificateAuthservice | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.CertificateAuthserviceAPI.CertificateauthservicePost(context.Background()).CertificateAuthservice(certificateAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.CertificateauthservicePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateauthservicePost`: CreateCertificateAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthserviceAPI.CertificateauthservicePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CertificateAuthserviceAPICertificateauthservicePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**certificateAuthservice** | [**CertificateAuthservice**](CertificateAuthservice.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateCertificateAuthserviceResponse**](CreateCertificateAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateauthserviceReferenceDelete

> CertificateauthserviceReferenceDelete(ctx, reference).Execute()

Delete a certificate:authservice object



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
	reference := "reference_example" // string | Reference of the certificate:authservice object

	apiClient := security.NewAPIClient()
	r, err := apiClient.CertificateAuthserviceAPI.CertificateauthserviceReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.CertificateauthserviceReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the certificate:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `CertificateAuthserviceAPICertificateauthserviceReferenceDeleteRequest` struct via the builder pattern


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


## CertificateauthserviceReferenceGet

> GetCertificateAuthserviceResponse CertificateauthserviceReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific certificate:authservice object



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
	reference := "reference_example" // string | Reference of the certificate:authservice object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.CertificateAuthserviceAPI.CertificateauthserviceReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.CertificateauthserviceReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateauthserviceReferenceGet`: GetCertificateAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthserviceAPI.CertificateauthserviceReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the certificate:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `CertificateAuthserviceAPICertificateauthserviceReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetCertificateAuthserviceResponse**](GetCertificateAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CertificateauthserviceReferencePut

> UpdateCertificateAuthserviceResponse CertificateauthserviceReferencePut(ctx, reference).CertificateAuthservice(certificateAuthservice).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a certificate:authservice object



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
	reference := "reference_example" // string | Reference of the certificate:authservice object
	certificateAuthservice := *security.NewCertificateAuthservice() // CertificateAuthservice | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.CertificateAuthserviceAPI.CertificateauthserviceReferencePut(context.Background(), reference).CertificateAuthservice(certificateAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.CertificateauthserviceReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CertificateauthserviceReferencePut`: UpdateCertificateAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthserviceAPI.CertificateauthserviceReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the certificate:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `CertificateAuthserviceAPICertificateauthserviceReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**certificateAuthservice** | [**CertificateAuthservice**](CertificateAuthservice.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateCertificateAuthserviceResponse**](UpdateCertificateAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

