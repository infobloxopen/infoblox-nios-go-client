# SharedrecordSrvAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharedrecordsrvGet**](SharedrecordSrvAPI.md#SharedrecordsrvGet) | **Get** /sharedrecord:srv | Retrieve sharedrecord:srv objects
[**SharedrecordsrvPost**](SharedrecordSrvAPI.md#SharedrecordsrvPost) | **Post** /sharedrecord:srv | Create a sharedrecord:srv object
[**SharedrecordsrvReferenceDelete**](SharedrecordSrvAPI.md#SharedrecordsrvReferenceDelete) | **Delete** /sharedrecord:srv/{reference} | Delete a sharedrecord:srv object
[**SharedrecordsrvReferenceGet**](SharedrecordSrvAPI.md#SharedrecordsrvReferenceGet) | **Get** /sharedrecord:srv/{reference} | Get a specific sharedrecord:srv object
[**SharedrecordsrvReferencePut**](SharedrecordSrvAPI.md#SharedrecordsrvReferencePut) | **Put** /sharedrecord:srv/{reference} | Update a sharedrecord:srv object



## SharedrecordsrvGet

> ListSharedrecordSrvResponse SharedrecordsrvGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve sharedrecord:srv objects



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
	resp, r, err := apiClient.SharedrecordSrvAPI.SharedrecordsrvGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordSrvAPI.SharedrecordsrvGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordsrvGet`: ListSharedrecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordSrvAPI.SharedrecordsrvGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordSrvAPISharedrecordsrvGetRequest` struct via the builder pattern


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

[**ListSharedrecordSrvResponse**](ListSharedrecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordsrvPost

> CreateSharedrecordSrvResponse SharedrecordsrvPost(ctx).SharedrecordSrv(sharedrecordSrv).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a sharedrecord:srv object



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
	sharedrecordSrv := *dns.NewSharedrecordSrv() // SharedrecordSrv | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordSrvAPI.SharedrecordsrvPost(context.Background()).SharedrecordSrv(sharedrecordSrv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordSrvAPI.SharedrecordsrvPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordsrvPost`: CreateSharedrecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordSrvAPI.SharedrecordsrvPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordSrvAPISharedrecordsrvPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordSrv** | [**SharedrecordSrv**](SharedrecordSrv.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSharedrecordSrvResponse**](CreateSharedrecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordsrvReferenceDelete

> SharedrecordsrvReferenceDelete(ctx, reference).Execute()

Delete a sharedrecord:srv object



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
	reference := "reference_example" // string | Reference of the sharedrecord:srv object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.SharedrecordSrvAPI.SharedrecordsrvReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordSrvAPI.SharedrecordsrvReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordSrvAPISharedrecordsrvReferenceDeleteRequest` struct via the builder pattern


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


## SharedrecordsrvReferenceGet

> GetSharedrecordSrvResponse SharedrecordsrvReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific sharedrecord:srv object



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
	reference := "reference_example" // string | Reference of the sharedrecord:srv object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordSrvAPI.SharedrecordsrvReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordSrvAPI.SharedrecordsrvReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordsrvReferenceGet`: GetSharedrecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordSrvAPI.SharedrecordsrvReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordSrvAPISharedrecordsrvReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSharedrecordSrvResponse**](GetSharedrecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordsrvReferencePut

> UpdateSharedrecordSrvResponse SharedrecordsrvReferencePut(ctx, reference).SharedrecordSrv(sharedrecordSrv).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a sharedrecord:srv object



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
	reference := "reference_example" // string | Reference of the sharedrecord:srv object
	sharedrecordSrv := *dns.NewSharedrecordSrv() // SharedrecordSrv | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordSrvAPI.SharedrecordsrvReferencePut(context.Background(), reference).SharedrecordSrv(sharedrecordSrv).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordSrvAPI.SharedrecordsrvReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordsrvReferencePut`: UpdateSharedrecordSrvResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordSrvAPI.SharedrecordsrvReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:srv object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordSrvAPISharedrecordsrvReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordSrv** | [**SharedrecordSrv**](SharedrecordSrv.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSharedrecordSrvResponse**](UpdateSharedrecordSrvResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

