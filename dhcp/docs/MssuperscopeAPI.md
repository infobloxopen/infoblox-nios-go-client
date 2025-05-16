# MssuperscopeAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](MssuperscopeAPI.md#Get) | **Get** /mssuperscope | Retrieve mssuperscope objects
[**Post**](MssuperscopeAPI.md#Post) | **Post** /mssuperscope | Create a mssuperscope object
[**ReferenceDelete**](MssuperscopeAPI.md#ReferenceDelete) | **Delete** /mssuperscope/{reference} | Delete a mssuperscope object
[**ReferenceGet**](MssuperscopeAPI.md#ReferenceGet) | **Get** /mssuperscope/{reference} | Get a specific mssuperscope object
[**ReferencePut**](MssuperscopeAPI.md#ReferencePut) | **Put** /mssuperscope/{reference} | Update a mssuperscope object



## Get

> ListMssuperscopeResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve mssuperscope objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.MssuperscopeAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListMssuperscopeResponse
	fmt.Fprintf(os.Stdout, "Response from `MssuperscopeAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPIGetRequest` struct via the builder pattern


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

[**ListMssuperscopeResponse**](ListMssuperscopeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateMssuperscopeResponse Post(ctx).Mssuperscope(mssuperscope).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a mssuperscope object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {
	mssuperscope := *dhcp.NewMssuperscope() // Mssuperscope | Object data to create

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.MssuperscopeAPI.Post(context.Background()).Mssuperscope(mssuperscope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateMssuperscopeResponse
	fmt.Fprintf(os.Stdout, "Response from `MssuperscopeAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**mssuperscope** | [**Mssuperscope**](Mssuperscope.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateMssuperscopeResponse**](CreateMssuperscopeResponse.md)

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

Delete a mssuperscope object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the mssuperscope object

	apiClient := dhcp.NewAPIClient()
	r, err := apiClient.MssuperscopeAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the mssuperscope object | 

### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetMssuperscopeResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific mssuperscope object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the mssuperscope object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.MssuperscopeAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetMssuperscopeResponse
	fmt.Fprintf(os.Stdout, "Response from `MssuperscopeAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the mssuperscope object | 

### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMssuperscopeResponse**](GetMssuperscopeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateMssuperscopeResponse ReferencePut(ctx, reference).Mssuperscope(mssuperscope).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a mssuperscope object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dhcp"
)

func main() {
	reference := "reference_example" // string | Reference of the mssuperscope object
	mssuperscope := *dhcp.NewMssuperscope() // Mssuperscope | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.MssuperscopeAPI.ReferencePut(context.Background(), reference).Mssuperscope(mssuperscope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MssuperscopeAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateMssuperscopeResponse
	fmt.Fprintf(os.Stdout, "Response from `MssuperscopeAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the mssuperscope object | 

### Other Parameters

Other parameters are passed through a pointer to a `MssuperscopeAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**mssuperscope** | [**Mssuperscope**](Mssuperscope.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMssuperscopeResponse**](UpdateMssuperscopeResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

