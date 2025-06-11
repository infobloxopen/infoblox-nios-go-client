# AdAuthServiceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](AdAuthServiceAPI.md#Get) | **Get** /ad_auth_service | Retrieve ad_auth_service objects
[**Post**](AdAuthServiceAPI.md#Post) | **Post** /ad_auth_service | Create a ad_auth_service object
[**ReferenceDelete**](AdAuthServiceAPI.md#ReferenceDelete) | **Delete** /ad_auth_service/{reference} | Delete a ad_auth_service object
[**ReferenceGet**](AdAuthServiceAPI.md#ReferenceGet) | **Get** /ad_auth_service/{reference} | Get a specific ad_auth_service object
[**ReferencePut**](AdAuthServiceAPI.md#ReferencePut) | **Put** /ad_auth_service/{reference} | Update a ad_auth_service object



## Get

> ListAdAuthServiceResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve ad_auth_service objects



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
	resp, r, err := apiClient.AdAuthServiceAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAuthServiceAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListAdAuthServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `AdAuthServiceAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `AdAuthServiceAPIGetRequest` struct via the builder pattern


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

[**ListAdAuthServiceResponse**](ListAdAuthServiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateAdAuthServiceResponse Post(ctx).AdAuthService(adAuthService).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a ad_auth_service object



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
	adAuthService := *security.NewAdAuthService() // AdAuthService | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.AdAuthServiceAPI.Post(context.Background()).AdAuthService(adAuthService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAuthServiceAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateAdAuthServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `AdAuthServiceAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `AdAuthServiceAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**adAuthService** | [**AdAuthService**](AdAuthService.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateAdAuthServiceResponse**](CreateAdAuthServiceResponse.md)

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

Delete a ad_auth_service object



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
	reference := "reference_example" // string | Reference of the ad_auth_service object

	apiClient := security.NewAPIClient()
	r, err := apiClient.AdAuthServiceAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAuthServiceAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ad_auth_service object | 

### Other Parameters

Other parameters are passed through a pointer to a `AdAuthServiceAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetAdAuthServiceResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific ad_auth_service object



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
	reference := "reference_example" // string | Reference of the ad_auth_service object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.AdAuthServiceAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAuthServiceAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetAdAuthServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `AdAuthServiceAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ad_auth_service object | 

### Other Parameters

Other parameters are passed through a pointer to a `AdAuthServiceAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetAdAuthServiceResponse**](GetAdAuthServiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateAdAuthServiceResponse ReferencePut(ctx, reference).AdAuthService(adAuthService).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a ad_auth_service object



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
	reference := "reference_example" // string | Reference of the ad_auth_service object
	adAuthService := *security.NewAdAuthService() // AdAuthService | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.AdAuthServiceAPI.ReferencePut(context.Background(), reference).AdAuthService(adAuthService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdAuthServiceAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateAdAuthServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `AdAuthServiceAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ad_auth_service object | 

### Other Parameters

Other parameters are passed through a pointer to a `AdAuthServiceAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**adAuthService** | [**AdAuthService**](AdAuthService.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateAdAuthServiceResponse**](UpdateAdAuthServiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

