# SharedrecordCnameAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SharedrecordcnameGet**](SharedrecordCnameAPI.md#SharedrecordcnameGet) | **Get** /sharedrecord:cname | Retrieve sharedrecord:cname objects
[**SharedrecordcnamePost**](SharedrecordCnameAPI.md#SharedrecordcnamePost) | **Post** /sharedrecord:cname | Create a sharedrecord:cname object
[**SharedrecordcnameReferenceDelete**](SharedrecordCnameAPI.md#SharedrecordcnameReferenceDelete) | **Delete** /sharedrecord:cname/{reference} | Delete a sharedrecord:cname object
[**SharedrecordcnameReferenceGet**](SharedrecordCnameAPI.md#SharedrecordcnameReferenceGet) | **Get** /sharedrecord:cname/{reference} | Get a specific sharedrecord:cname object
[**SharedrecordcnameReferencePut**](SharedrecordCnameAPI.md#SharedrecordcnameReferencePut) | **Put** /sharedrecord:cname/{reference} | Update a sharedrecord:cname object



## SharedrecordcnameGet

> ListSharedrecordCnameResponse SharedrecordcnameGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve sharedrecord:cname objects



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
	resp, r, err := apiClient.SharedrecordCnameAPI.SharedrecordcnameGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordCnameAPI.SharedrecordcnameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordcnameGet`: ListSharedrecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordCnameAPI.SharedrecordcnameGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordCnameAPISharedrecordcnameGetRequest` struct via the builder pattern


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

[**ListSharedrecordCnameResponse**](ListSharedrecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordcnamePost

> CreateSharedrecordCnameResponse SharedrecordcnamePost(ctx).SharedrecordCname(sharedrecordCname).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a sharedrecord:cname object



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
	sharedrecordCname := *dns.NewSharedrecordCname() // SharedrecordCname | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordCnameAPI.SharedrecordcnamePost(context.Background()).SharedrecordCname(sharedrecordCname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordCnameAPI.SharedrecordcnamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordcnamePost`: CreateSharedrecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordCnameAPI.SharedrecordcnamePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordCnameAPISharedrecordcnamePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordCname** | [**SharedrecordCname**](SharedrecordCname.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSharedrecordCnameResponse**](CreateSharedrecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordcnameReferenceDelete

> SharedrecordcnameReferenceDelete(ctx, reference).Execute()

Delete a sharedrecord:cname object



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
	reference := "reference_example" // string | Reference of the sharedrecord:cname object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.SharedrecordCnameAPI.SharedrecordcnameReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordCnameAPI.SharedrecordcnameReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:cname object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordCnameAPISharedrecordcnameReferenceDeleteRequest` struct via the builder pattern


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


## SharedrecordcnameReferenceGet

> GetSharedrecordCnameResponse SharedrecordcnameReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific sharedrecord:cname object



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
	reference := "reference_example" // string | Reference of the sharedrecord:cname object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordCnameAPI.SharedrecordcnameReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordCnameAPI.SharedrecordcnameReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordcnameReferenceGet`: GetSharedrecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordCnameAPI.SharedrecordcnameReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:cname object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordCnameAPISharedrecordcnameReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSharedrecordCnameResponse**](GetSharedrecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SharedrecordcnameReferencePut

> UpdateSharedrecordCnameResponse SharedrecordcnameReferencePut(ctx, reference).SharedrecordCname(sharedrecordCname).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a sharedrecord:cname object



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
	reference := "reference_example" // string | Reference of the sharedrecord:cname object
	sharedrecordCname := *dns.NewSharedrecordCname() // SharedrecordCname | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordCnameAPI.SharedrecordcnameReferencePut(context.Background(), reference).SharedrecordCname(sharedrecordCname).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordCnameAPI.SharedrecordcnameReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SharedrecordcnameReferencePut`: UpdateSharedrecordCnameResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordCnameAPI.SharedrecordcnameReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecord:cname object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordCnameAPISharedrecordcnameReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordCname** | [**SharedrecordCname**](SharedrecordCname.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSharedrecordCnameResponse**](UpdateSharedrecordCnameResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

