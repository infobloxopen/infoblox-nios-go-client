# RadiusAuthserviceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](RadiusAuthserviceAPI.md#Create) | **Post** /radius:authservice | Create a radius:authservice object
[**Delete**](RadiusAuthserviceAPI.md#Delete) | **Delete** /radius:authservice/{reference} | Delete a radius:authservice object
[**List**](RadiusAuthserviceAPI.md#List) | **Get** /radius:authservice | Retrieve radius:authservice objects
[**Read**](RadiusAuthserviceAPI.md#Read) | **Get** /radius:authservice/{reference} | Get a specific radius:authservice object
[**Update**](RadiusAuthserviceAPI.md#Update) | **Put** /radius:authservice/{reference} | Update a radius:authservice object



## Create

> CreateRadiusAuthserviceResponse Create(ctx).RadiusAuthservice(radiusAuthservice).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a radius:authservice object



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
	radiusAuthservice := *security.NewRadiusAuthservice() // RadiusAuthservice | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.RadiusAuthserviceAPI.Create(context.Background()).RadiusAuthservice(radiusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateRadiusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `RadiusAuthserviceAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RadiusAuthserviceAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**radiusAuthservice** | [**RadiusAuthservice**](RadiusAuthservice.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## Delete

> Delete(ctx, reference).Execute()

Delete a radius:authservice object



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
	reference := "reference_example" // string | Reference of the radius:authservice object

	apiClient := security.NewAPIClient()
	r, err := apiClient.RadiusAuthserviceAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a `RadiusAuthserviceAPIDeleteRequest` struct via the builder pattern


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

> ListRadiusAuthserviceResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve radius:authservice objects



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
	resp, r, err := apiClient.RadiusAuthserviceAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListRadiusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `RadiusAuthserviceAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RadiusAuthserviceAPIListRequest` struct via the builder pattern


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

[**ListRadiusAuthserviceResponse**](ListRadiusAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetRadiusAuthserviceResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific radius:authservice object



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
	reference := "reference_example" // string | Reference of the radius:authservice object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.RadiusAuthserviceAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetRadiusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `RadiusAuthserviceAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the radius:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `RadiusAuthserviceAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

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


## Update

> UpdateRadiusAuthserviceResponse Update(ctx, reference).RadiusAuthservice(radiusAuthservice).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a radius:authservice object



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
	reference := "reference_example" // string | Reference of the radius:authservice object
	radiusAuthservice := *security.NewRadiusAuthservice() // RadiusAuthservice | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.RadiusAuthserviceAPI.Update(context.Background(), reference).RadiusAuthservice(radiusAuthservice).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RadiusAuthserviceAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateRadiusAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `RadiusAuthserviceAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the radius:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `RadiusAuthserviceAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**radiusAuthservice** | [**RadiusAuthservice**](RadiusAuthservice.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

