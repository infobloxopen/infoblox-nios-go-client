# HsmThaleslunagroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HsmthaleslunagroupGet**](HsmThaleslunagroupAPI.md#HsmthaleslunagroupGet) | **Get** /hsm:thaleslunagroup | Retrieve hsm:thaleslunagroup objects
[**HsmthaleslunagroupPost**](HsmThaleslunagroupAPI.md#HsmthaleslunagroupPost) | **Post** /hsm:thaleslunagroup | Create a hsm:thaleslunagroup object
[**HsmthaleslunagroupReferenceDelete**](HsmThaleslunagroupAPI.md#HsmthaleslunagroupReferenceDelete) | **Delete** /hsm:thaleslunagroup/{reference} | Delete a hsm:thaleslunagroup object
[**HsmthaleslunagroupReferenceGet**](HsmThaleslunagroupAPI.md#HsmthaleslunagroupReferenceGet) | **Get** /hsm:thaleslunagroup/{reference} | Get a specific hsm:thaleslunagroup object
[**HsmthaleslunagroupReferencePut**](HsmThaleslunagroupAPI.md#HsmthaleslunagroupReferencePut) | **Put** /hsm:thaleslunagroup/{reference} | Update a hsm:thaleslunagroup object



## HsmthaleslunagroupGet

> ListHsmThaleslunagroupResponse HsmthaleslunagroupGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve hsm:thaleslunagroup objects



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
	resp, r, err := apiClient.HsmThaleslunagroupAPI.HsmthaleslunagroupGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.HsmthaleslunagroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmthaleslunagroupGet`: ListHsmThaleslunagroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmThaleslunagroupAPI.HsmthaleslunagroupGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPIHsmthaleslunagroupGetRequest` struct via the builder pattern


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

[**ListHsmThaleslunagroupResponse**](ListHsmThaleslunagroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HsmthaleslunagroupPost

> CreateHsmThaleslunagroupResponse HsmthaleslunagroupPost(ctx).HsmThaleslunagroup(hsmThaleslunagroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a hsm:thaleslunagroup object



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
	hsmThaleslunagroup := *security.NewHsmThaleslunagroup() // HsmThaleslunagroup | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmThaleslunagroupAPI.HsmthaleslunagroupPost(context.Background()).HsmThaleslunagroup(hsmThaleslunagroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.HsmthaleslunagroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmthaleslunagroupPost`: CreateHsmThaleslunagroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmThaleslunagroupAPI.HsmthaleslunagroupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPIHsmthaleslunagroupPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hsmThaleslunagroup** | [**HsmThaleslunagroup**](HsmThaleslunagroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateHsmThaleslunagroupResponse**](CreateHsmThaleslunagroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HsmthaleslunagroupReferenceDelete

> HsmthaleslunagroupReferenceDelete(ctx, reference).Execute()

Delete a hsm:thaleslunagroup object



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
	reference := "reference_example" // string | Reference of the hsm:thaleslunagroup object

	apiClient := security.NewAPIClient()
	r, err := apiClient.HsmThaleslunagroupAPI.HsmthaleslunagroupReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.HsmthaleslunagroupReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:thaleslunagroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPIHsmthaleslunagroupReferenceDeleteRequest` struct via the builder pattern


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


## HsmthaleslunagroupReferenceGet

> GetHsmThaleslunagroupResponse HsmthaleslunagroupReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific hsm:thaleslunagroup object



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
	reference := "reference_example" // string | Reference of the hsm:thaleslunagroup object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmThaleslunagroupAPI.HsmthaleslunagroupReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.HsmthaleslunagroupReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmthaleslunagroupReferenceGet`: GetHsmThaleslunagroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmThaleslunagroupAPI.HsmthaleslunagroupReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:thaleslunagroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPIHsmthaleslunagroupReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetHsmThaleslunagroupResponse**](GetHsmThaleslunagroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HsmthaleslunagroupReferencePut

> UpdateHsmThaleslunagroupResponse HsmthaleslunagroupReferencePut(ctx, reference).HsmThaleslunagroup(hsmThaleslunagroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a hsm:thaleslunagroup object



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
	reference := "reference_example" // string | Reference of the hsm:thaleslunagroup object
	hsmThaleslunagroup := *security.NewHsmThaleslunagroup() // HsmThaleslunagroup | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.HsmThaleslunagroupAPI.HsmthaleslunagroupReferencePut(context.Background(), reference).HsmThaleslunagroup(hsmThaleslunagroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HsmThaleslunagroupAPI.HsmthaleslunagroupReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HsmthaleslunagroupReferencePut`: UpdateHsmThaleslunagroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HsmThaleslunagroupAPI.HsmthaleslunagroupReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the hsm:thaleslunagroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `HsmThaleslunagroupAPIHsmthaleslunagroupReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**hsmThaleslunagroup** | [**HsmThaleslunagroup**](HsmThaleslunagroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateHsmThaleslunagroupResponse**](UpdateHsmThaleslunagroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

