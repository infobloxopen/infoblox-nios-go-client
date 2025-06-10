# HsmEntrustnshieldgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HsmentrustnshieldgroupGet**](HsmEntrustnshieldgroupAPI.md#HsmentrustnshieldgroupGet) | **Get** /hsm:entrustnshieldgroup | Retrieve hsm:entrustnshieldgroup objects
[**HsmentrustnshieldgroupPost**](HsmEntrustnshieldgroupAPI.md#HsmentrustnshieldgroupPost) | **Post** /hsm:entrustnshieldgroup | Create a hsm:entrustnshieldgroup object
[**HsmentrustnshieldgroupReferenceDelete**](HsmEntrustnshieldgroupAPI.md#HsmentrustnshieldgroupReferenceDelete) | **Delete** /hsm:entrustnshieldgroup/{reference} | Delete a hsm:entrustnshieldgroup object
[**HsmentrustnshieldgroupReferenceGet**](HsmEntrustnshieldgroupAPI.md#HsmentrustnshieldgroupReferenceGet) | **Get** /hsm:entrustnshieldgroup/{reference} | Get a specific hsm:entrustnshieldgroup object
[**HsmentrustnshieldgroupReferencePut**](HsmEntrustnshieldgroupAPI.md#HsmentrustnshieldgroupReferencePut) | **Put** /hsm:entrustnshieldgroup/{reference} | Update a hsm:entrustnshieldgroup object



## HsmentrustnshieldgroupGet

> ListHsmEntrustnshieldgroupResponse HsmentrustnshieldgroupGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve hsm:entrustnshieldgroup objects



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
	resp, r, err := apiClient.HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmentrustnshieldgroupGet`: ListHsmEntrustnshieldgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPIHsmentrustnshieldgroupGetRequest` struct via the builder pattern


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

[**ListHsmEntrustnshieldgroupResponse**](ListHsmEntrustnshieldgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HsmentrustnshieldgroupPost

> CreateHsmEntrustnshieldgroupResponse HsmentrustnshieldgroupPost(ctx).HsmEntrustnshieldgroup(hsmEntrustnshieldgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a hsm:entrustnshieldgroup object



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
	hsmEntrustnshieldgroup := *security.NewHsmEntrustnshieldgroup() // HsmEntrustnshieldgroup | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupPost(context.Background()).HsmEntrustnshieldgroup(hsmEntrustnshieldgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmentrustnshieldgroupPost`: CreateHsmEntrustnshieldgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPIHsmentrustnshieldgroupPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hsmEntrustnshieldgroup** | [**HsmEntrustnshieldgroup**](HsmEntrustnshieldgroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateHsmEntrustnshieldgroupResponse**](CreateHsmEntrustnshieldgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HsmentrustnshieldgroupReferenceDelete

> HsmentrustnshieldgroupReferenceDelete(ctx, reference).Execute()

Delete a hsm:entrustnshieldgroup object



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
	reference := "reference_example" // string | Reference of the hsm:entrustnshieldgroup object

	apiClient := security.NewAPIClient()
	r, err := apiClient.HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:entrustnshieldgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPIHsmentrustnshieldgroupReferenceDeleteRequest` struct via the builder pattern


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


## HsmentrustnshieldgroupReferenceGet

> GetHsmEntrustnshieldgroupResponse HsmentrustnshieldgroupReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific hsm:entrustnshieldgroup object



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
	reference := "reference_example" // string | Reference of the hsm:entrustnshieldgroup object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmentrustnshieldgroupReferenceGet`: GetHsmEntrustnshieldgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:entrustnshieldgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPIHsmentrustnshieldgroupReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetHsmEntrustnshieldgroupResponse**](GetHsmEntrustnshieldgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HsmentrustnshieldgroupReferencePut

> UpdateHsmEntrustnshieldgroupResponse HsmentrustnshieldgroupReferencePut(ctx, reference).HsmEntrustnshieldgroup(hsmEntrustnshieldgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a hsm:entrustnshieldgroup object



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
	reference := "reference_example" // string | Reference of the hsm:entrustnshieldgroup object
	hsmEntrustnshieldgroup := *security.NewHsmEntrustnshieldgroup() // HsmEntrustnshieldgroup | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupReferencePut(context.Background(), reference).HsmEntrustnshieldgroup(hsmEntrustnshieldgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmentrustnshieldgroupReferencePut`: UpdateHsmEntrustnshieldgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmEntrustnshieldgroupAPI.HsmentrustnshieldgroupReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:entrustnshieldgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmEntrustnshieldgroupAPIHsmentrustnshieldgroupReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hsmEntrustnshieldgroup** | [**HsmEntrustnshieldgroup**](HsmEntrustnshieldgroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateHsmEntrustnshieldgroupResponse**](UpdateHsmEntrustnshieldgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

