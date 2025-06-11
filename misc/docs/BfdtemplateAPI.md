# BfdtemplateAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](BfdtemplateAPI.md#Get) | **Get** /bfdtemplate | Retrieve bfdtemplate objects
[**Post**](BfdtemplateAPI.md#Post) | **Post** /bfdtemplate | Create a bfdtemplate object
[**ReferenceDelete**](BfdtemplateAPI.md#ReferenceDelete) | **Delete** /bfdtemplate/{reference} | Delete a bfdtemplate object
[**ReferenceGet**](BfdtemplateAPI.md#ReferenceGet) | **Get** /bfdtemplate/{reference} | Get a specific bfdtemplate object
[**ReferencePut**](BfdtemplateAPI.md#ReferencePut) | **Put** /bfdtemplate/{reference} | Update a bfdtemplate object



## Get

> ListBfdtemplateResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve bfdtemplate objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.BfdtemplateAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BfdtemplateAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListBfdtemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `BfdtemplateAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `BfdtemplateAPIGetRequest` struct via the builder pattern


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

[**ListBfdtemplateResponse**](ListBfdtemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateBfdtemplateResponse Post(ctx).Bfdtemplate(bfdtemplate).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a bfdtemplate object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	bfdtemplate := *misc.NewBfdtemplate() // Bfdtemplate | Object data to create

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.BfdtemplateAPI.Post(context.Background()).Bfdtemplate(bfdtemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BfdtemplateAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateBfdtemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `BfdtemplateAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `BfdtemplateAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**bfdtemplate** | [**Bfdtemplate**](Bfdtemplate.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateBfdtemplateResponse**](CreateBfdtemplateResponse.md)

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

Delete a bfdtemplate object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the bfdtemplate object

	apiClient := misc.NewAPIClient()
	r, err := apiClient.BfdtemplateAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BfdtemplateAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the bfdtemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `BfdtemplateAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetBfdtemplateResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific bfdtemplate object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the bfdtemplate object

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.BfdtemplateAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BfdtemplateAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetBfdtemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `BfdtemplateAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the bfdtemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `BfdtemplateAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetBfdtemplateResponse**](GetBfdtemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateBfdtemplateResponse ReferencePut(ctx, reference).Bfdtemplate(bfdtemplate).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a bfdtemplate object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the bfdtemplate object
	bfdtemplate := *misc.NewBfdtemplate() // Bfdtemplate | Object data to update

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.BfdtemplateAPI.ReferencePut(context.Background(), reference).Bfdtemplate(bfdtemplate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BfdtemplateAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateBfdtemplateResponse
	fmt.Fprintf(os.Stdout, "Response from `BfdtemplateAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the bfdtemplate object | 

### Other Parameters

Other parameters are passed through a pointer to a `BfdtemplateAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**bfdtemplate** | [**Bfdtemplate**](Bfdtemplate.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateBfdtemplateResponse**](UpdateBfdtemplateResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

