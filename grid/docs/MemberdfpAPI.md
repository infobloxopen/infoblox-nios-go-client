# MemberdfpAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](MemberdfpAPI.md#Get) | **Get** /memberdfp | Retrieve memberdfp objects
[**Post**](MemberdfpAPI.md#Post) | **Post** /memberdfp | Create a memberdfp object
[**ReferenceDelete**](MemberdfpAPI.md#ReferenceDelete) | **Delete** /memberdfp/{reference} | Delete a memberdfp object
[**ReferenceGet**](MemberdfpAPI.md#ReferenceGet) | **Get** /memberdfp/{reference} | Get a specific memberdfp object
[**ReferencePut**](MemberdfpAPI.md#ReferencePut) | **Put** /memberdfp/{reference} | Update a memberdfp object



## Get

> ListMemberdfpResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve memberdfp objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberdfpAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberdfpAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListMemberdfpResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberdfpAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MemberdfpAPIGetRequest` struct via the builder pattern


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

[**ListMemberdfpResponse**](ListMemberdfpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateMemberdfpResponse Post(ctx).Memberdfp(memberdfp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a memberdfp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {
	memberdfp := *grid.NewMemberdfp() // Memberdfp | Object data to create

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberdfpAPI.Post(context.Background()).Memberdfp(memberdfp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberdfpAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateMemberdfpResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberdfpAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MemberdfpAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**memberdfp** | [**Memberdfp**](Memberdfp.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateMemberdfpResponse**](CreateMemberdfpResponse.md)

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

Delete a memberdfp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {
	reference := "reference_example" // string | Reference of the memberdfp object

	apiClient := grid.NewAPIClient()
	r, err := apiClient.MemberdfpAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberdfpAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the memberdfp object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberdfpAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetMemberdfpResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific memberdfp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {
	reference := "reference_example" // string | Reference of the memberdfp object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberdfpAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberdfpAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetMemberdfpResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberdfpAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the memberdfp object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberdfpAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMemberdfpResponse**](GetMemberdfpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateMemberdfpResponse ReferencePut(ctx, reference).Memberdfp(memberdfp).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a memberdfp object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/grid"
)

func main() {
	reference := "reference_example" // string | Reference of the memberdfp object
	memberdfp := *grid.NewMemberdfp() // Memberdfp | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberdfpAPI.ReferencePut(context.Background(), reference).Memberdfp(memberdfp).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberdfpAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateMemberdfpResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberdfpAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the memberdfp object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberdfpAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**memberdfp** | [**Memberdfp**](Memberdfp.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMemberdfpResponse**](UpdateMemberdfpResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

