# SharedrecordgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](SharedrecordgroupAPI.md#Get) | **Get** /sharedrecordgroup | Retrieve sharedrecordgroup objects
[**Post**](SharedrecordgroupAPI.md#Post) | **Post** /sharedrecordgroup | Create a sharedrecordgroup object
[**ReferenceDelete**](SharedrecordgroupAPI.md#ReferenceDelete) | **Delete** /sharedrecordgroup/{reference} | Delete a sharedrecordgroup object
[**ReferenceGet**](SharedrecordgroupAPI.md#ReferenceGet) | **Get** /sharedrecordgroup/{reference} | Get a specific sharedrecordgroup object
[**ReferencePut**](SharedrecordgroupAPI.md#ReferencePut) | **Put** /sharedrecordgroup/{reference} | Update a sharedrecordgroup object



## Get

> ListSharedrecordgroupResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve sharedrecordgroup objects



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
	resp, r, err := apiClient.SharedrecordgroupAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordgroupAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListSharedrecordgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordgroupAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordgroupAPIGetRequest` struct via the builder pattern


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

[**ListSharedrecordgroupResponse**](ListSharedrecordgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateSharedrecordgroupResponse Post(ctx).Sharedrecordgroup(sharedrecordgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a sharedrecordgroup object



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
	sharedrecordgroup := *dns.NewSharedrecordgroup() // Sharedrecordgroup | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordgroupAPI.Post(context.Background()).Sharedrecordgroup(sharedrecordgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordgroupAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateSharedrecordgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordgroupAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordgroupAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordgroup** | [**Sharedrecordgroup**](Sharedrecordgroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSharedrecordgroupResponse**](CreateSharedrecordgroupResponse.md)

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

Delete a sharedrecordgroup object



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
	reference := "reference_example" // string | Reference of the sharedrecordgroup object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.SharedrecordgroupAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordgroupAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecordgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordgroupAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetSharedrecordgroupResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific sharedrecordgroup object



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
	reference := "reference_example" // string | Reference of the sharedrecordgroup object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordgroupAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordgroupAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetSharedrecordgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordgroupAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecordgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordgroupAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSharedrecordgroupResponse**](GetSharedrecordgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateSharedrecordgroupResponse ReferencePut(ctx, reference).Sharedrecordgroup(sharedrecordgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a sharedrecordgroup object



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
	reference := "reference_example" // string | Reference of the sharedrecordgroup object
	sharedrecordgroup := *dns.NewSharedrecordgroup() // Sharedrecordgroup | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.SharedrecordgroupAPI.ReferencePut(context.Background(), reference).Sharedrecordgroup(sharedrecordgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedrecordgroupAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateSharedrecordgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SharedrecordgroupAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the sharedrecordgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `SharedrecordgroupAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**sharedrecordgroup** | [**Sharedrecordgroup**](Sharedrecordgroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSharedrecordgroupResponse**](UpdateSharedrecordgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

