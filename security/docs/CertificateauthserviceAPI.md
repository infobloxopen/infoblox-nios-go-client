# CertificateAuthserviceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CertificateAuthserviceAPI.md#Create) | **Post** /certificate:authservice | Create a certificate:authservice object
[**Delete**](CertificateAuthserviceAPI.md#Delete) | **Delete** /certificate:authservice/{reference} | Delete a certificate:authservice object
[**List**](CertificateAuthserviceAPI.md#List) | **Get** /certificate:authservice | Retrieve certificate:authservice objects
[**Read**](CertificateAuthserviceAPI.md#Read) | **Get** /certificate:authservice/{reference} | Get a specific certificate:authservice object
[**Update**](CertificateAuthserviceAPI.md#Update) | **Put** /certificate:authservice/{reference} | Update a certificate:authservice object



## Create

> CreateCertificateAuthserviceResponse Create(ctx).CertificateAuthservice(certificateAuthservice).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a certificate:authservice object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	certificateAuthservice := *security.NewCertificateAuthservice() // CertificateAuthservice | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.CertificateAuthserviceAPI.Create(context.Background()).CertificateAuthservice(certificateAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateCertificateAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthserviceAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CertificateAuthserviceAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**certificateAuthservice** | [**CertificateAuthservice**](CertificateAuthservice.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## Delete

> Delete(ctx, reference).Execute()

Delete a certificate:authservice object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the certificate:authservice object

	apiClient := security.NewAPIClient()
	r, err := apiClient.CertificateAuthserviceAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a `CertificateAuthserviceAPIDeleteRequest` struct via the builder pattern


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


## List

> ListCertificateAuthserviceResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve certificate:authservice objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.CertificateAuthserviceAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListCertificateAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthserviceAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `CertificateAuthserviceAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**paging** | **int32** | Control paging of results | 
**pageId** | **string** | Page id for retrieving next page of results | 
**filters** | **map[string]interface{}** |  | 
**extattrfilter** | **map[string]interface{}** |  | 
**proxySearch** | **string** | Search Grid members for data | 

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


## Read

> GetCertificateAuthserviceResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific certificate:authservice object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the certificate:authservice object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.CertificateAuthserviceAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetCertificateAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthserviceAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the certificate:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `CertificateAuthserviceAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

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


## Update

> UpdateCertificateAuthserviceResponse Update(ctx, reference).CertificateAuthservice(certificateAuthservice).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a certificate:authservice object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the certificate:authservice object
	certificateAuthservice := *security.NewCertificateAuthservice() // CertificateAuthservice | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.CertificateAuthserviceAPI.Update(context.Background(), reference).CertificateAuthservice(certificateAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CertificateAuthserviceAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateCertificateAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `CertificateAuthserviceAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the certificate:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `CertificateAuthserviceAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**certificateAuthservice** | [**CertificateAuthservice**](CertificateAuthservice.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

