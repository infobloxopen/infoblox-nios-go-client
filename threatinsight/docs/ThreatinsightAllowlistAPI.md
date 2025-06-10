# ThreatinsightAllowlistAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreatinsightallowlistGet**](ThreatinsightAllowlistAPI.md#ThreatinsightallowlistGet) | **Get** /threatinsight:allowlist | Retrieve threatinsight:allowlist objects
[**ThreatinsightallowlistPost**](ThreatinsightAllowlistAPI.md#ThreatinsightallowlistPost) | **Post** /threatinsight:allowlist | Create a threatinsight:allowlist object
[**ThreatinsightallowlistReferenceDelete**](ThreatinsightAllowlistAPI.md#ThreatinsightallowlistReferenceDelete) | **Delete** /threatinsight:allowlist/{reference} | Delete a threatinsight:allowlist object
[**ThreatinsightallowlistReferenceGet**](ThreatinsightAllowlistAPI.md#ThreatinsightallowlistReferenceGet) | **Get** /threatinsight:allowlist/{reference} | Get a specific threatinsight:allowlist object
[**ThreatinsightallowlistReferencePut**](ThreatinsightAllowlistAPI.md#ThreatinsightallowlistReferencePut) | **Put** /threatinsight:allowlist/{reference} | Update a threatinsight:allowlist object



## ThreatinsightallowlistGet

> ListThreatinsightAllowlistResponse ThreatinsightallowlistGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve threatinsight:allowlist objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightAllowlistAPI.ThreatinsightallowlistGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightAllowlistAPI.ThreatinsightallowlistGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightallowlistGet`: ListThreatinsightAllowlistResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightAllowlistAPI.ThreatinsightallowlistGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightAllowlistAPIThreatinsightallowlistGetRequest` struct via the builder pattern


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

[**ListThreatinsightAllowlistResponse**](ListThreatinsightAllowlistResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatinsightallowlistPost

> CreateThreatinsightAllowlistResponse ThreatinsightallowlistPost(ctx).ThreatinsightAllowlist(threatinsightAllowlist).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a threatinsight:allowlist object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {
	threatinsightAllowlist := *threatinsight.NewThreatinsightAllowlist() // ThreatinsightAllowlist | Object data to create

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightAllowlistAPI.ThreatinsightallowlistPost(context.Background()).ThreatinsightAllowlist(threatinsightAllowlist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightAllowlistAPI.ThreatinsightallowlistPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightallowlistPost`: CreateThreatinsightAllowlistResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightAllowlistAPI.ThreatinsightallowlistPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightAllowlistAPIThreatinsightallowlistPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatinsightAllowlist** | [**ThreatinsightAllowlist**](ThreatinsightAllowlist.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateThreatinsightAllowlistResponse**](CreateThreatinsightAllowlistResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatinsightallowlistReferenceDelete

> ThreatinsightallowlistReferenceDelete(ctx, reference).Execute()

Delete a threatinsight:allowlist object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {
	reference := "reference_example" // string | Reference of the threatinsight:allowlist object

	apiClient := threatinsight.NewAPIClient()
	r, err := apiClient.ThreatinsightAllowlistAPI.ThreatinsightallowlistReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightAllowlistAPI.ThreatinsightallowlistReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatinsight:allowlist object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightAllowlistAPIThreatinsightallowlistReferenceDeleteRequest` struct via the builder pattern


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


## ThreatinsightallowlistReferenceGet

> GetThreatinsightAllowlistResponse ThreatinsightallowlistReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific threatinsight:allowlist object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {
	reference := "reference_example" // string | Reference of the threatinsight:allowlist object

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightAllowlistAPI.ThreatinsightallowlistReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightAllowlistAPI.ThreatinsightallowlistReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightallowlistReferenceGet`: GetThreatinsightAllowlistResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightAllowlistAPI.ThreatinsightallowlistReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatinsight:allowlist object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightAllowlistAPIThreatinsightallowlistReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetThreatinsightAllowlistResponse**](GetThreatinsightAllowlistResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatinsightallowlistReferencePut

> UpdateThreatinsightAllowlistResponse ThreatinsightallowlistReferencePut(ctx, reference).ThreatinsightAllowlist(threatinsightAllowlist).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a threatinsight:allowlist object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {
	reference := "reference_example" // string | Reference of the threatinsight:allowlist object
	threatinsightAllowlist := *threatinsight.NewThreatinsightAllowlist() // ThreatinsightAllowlist | Object data to update

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightAllowlistAPI.ThreatinsightallowlistReferencePut(context.Background(), reference).ThreatinsightAllowlist(threatinsightAllowlist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightAllowlistAPI.ThreatinsightallowlistReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightallowlistReferencePut`: UpdateThreatinsightAllowlistResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightAllowlistAPI.ThreatinsightallowlistReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatinsight:allowlist object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightAllowlistAPIThreatinsightallowlistReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatinsightAllowlist** | [**ThreatinsightAllowlist**](ThreatinsightAllowlist.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateThreatinsightAllowlistResponse**](UpdateThreatinsightAllowlistResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

