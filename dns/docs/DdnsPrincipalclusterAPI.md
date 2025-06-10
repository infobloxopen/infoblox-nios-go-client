# DdnsPrincipalclusterAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DdnsprincipalclusterGet**](DdnsPrincipalclusterAPI.md#DdnsprincipalclusterGet) | **Get** /ddns:principalcluster | Retrieve ddns:principalcluster objects
[**DdnsprincipalclusterPost**](DdnsPrincipalclusterAPI.md#DdnsprincipalclusterPost) | **Post** /ddns:principalcluster | Create a ddns:principalcluster object
[**DdnsprincipalclusterReferenceDelete**](DdnsPrincipalclusterAPI.md#DdnsprincipalclusterReferenceDelete) | **Delete** /ddns:principalcluster/{reference} | Delete a ddns:principalcluster object
[**DdnsprincipalclusterReferenceGet**](DdnsPrincipalclusterAPI.md#DdnsprincipalclusterReferenceGet) | **Get** /ddns:principalcluster/{reference} | Get a specific ddns:principalcluster object
[**DdnsprincipalclusterReferencePut**](DdnsPrincipalclusterAPI.md#DdnsprincipalclusterReferencePut) | **Put** /ddns:principalcluster/{reference} | Update a ddns:principalcluster object



## DdnsprincipalclusterGet

> ListDdnsPrincipalclusterResponse DdnsprincipalclusterGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve ddns:principalcluster objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterAPI.DdnsprincipalclusterGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.DdnsprincipalclusterGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdnsprincipalclusterGet`: ListDdnsPrincipalclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterAPI.DdnsprincipalclusterGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPIDdnsprincipalclusterGetRequest` struct via the builder pattern


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

[**ListDdnsPrincipalclusterResponse**](ListDdnsPrincipalclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DdnsprincipalclusterPost

> CreateDdnsPrincipalclusterResponse DdnsprincipalclusterPost(ctx).DdnsPrincipalcluster(ddnsPrincipalcluster).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a ddns:principalcluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	ddnsPrincipalcluster := *dns.NewDdnsPrincipalcluster() // DdnsPrincipalcluster | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterAPI.DdnsprincipalclusterPost(context.Background()).DdnsPrincipalcluster(ddnsPrincipalcluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.DdnsprincipalclusterPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdnsprincipalclusterPost`: CreateDdnsPrincipalclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterAPI.DdnsprincipalclusterPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPIDdnsprincipalclusterPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ddnsPrincipalcluster** | [**DdnsPrincipalcluster**](DdnsPrincipalcluster.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDdnsPrincipalclusterResponse**](CreateDdnsPrincipalclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DdnsprincipalclusterReferenceDelete

> DdnsprincipalclusterReferenceDelete(ctx, reference).Execute()

Delete a ddns:principalcluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the ddns:principalcluster object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.DdnsPrincipalclusterAPI.DdnsprincipalclusterReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.DdnsprincipalclusterReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPIDdnsprincipalclusterReferenceDeleteRequest` struct via the builder pattern


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


## DdnsprincipalclusterReferenceGet

> GetDdnsPrincipalclusterResponse DdnsprincipalclusterReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific ddns:principalcluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the ddns:principalcluster object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterAPI.DdnsprincipalclusterReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.DdnsprincipalclusterReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdnsprincipalclusterReferenceGet`: GetDdnsPrincipalclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterAPI.DdnsprincipalclusterReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPIDdnsprincipalclusterReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDdnsPrincipalclusterResponse**](GetDdnsPrincipalclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DdnsprincipalclusterReferencePut

> UpdateDdnsPrincipalclusterResponse DdnsprincipalclusterReferencePut(ctx, reference).DdnsPrincipalcluster(ddnsPrincipalcluster).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a ddns:principalcluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the ddns:principalcluster object
	ddnsPrincipalcluster := *dns.NewDdnsPrincipalcluster() // DdnsPrincipalcluster | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterAPI.DdnsprincipalclusterReferencePut(context.Background(), reference).DdnsPrincipalcluster(ddnsPrincipalcluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterAPI.DdnsprincipalclusterReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdnsprincipalclusterReferencePut`: UpdateDdnsPrincipalclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterAPI.DdnsprincipalclusterReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterAPIDdnsprincipalclusterReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ddnsPrincipalcluster** | [**DdnsPrincipalcluster**](DdnsPrincipalcluster.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDdnsPrincipalclusterResponse**](UpdateDdnsPrincipalclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

