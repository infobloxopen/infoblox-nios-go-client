# NsgroupstubmemberAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](NsgroupstubmemberAPI.md#Get) | **Get** /nsgroup:stubmember | Retrieve nsgroup:stubmember objects
[**Post**](NsgroupstubmemberAPI.md#Post) | **Post** /nsgroup:stubmember | Create a nsgroup:stubmember object
[**Put**](NsgroupstubmemberAPI.md#Put) | **Put** /nsgroup:stubmember | Use PUT call as GET operation with _method for a Struct field of a nsgroup:stubmember object
[**ReferenceDelete**](NsgroupstubmemberAPI.md#ReferenceDelete) | **Delete** /nsgroup:stubmember/{reference} | Delete a nsgroup:stubmember object
[**ReferenceGet**](NsgroupstubmemberAPI.md#ReferenceGet) | **Get** /nsgroup:stubmember/{reference} | Get a specific nsgroup:stubmember object
[**ReferencePut**](NsgroupstubmemberAPI.md#ReferencePut) | **Put** /nsgroup:stubmember/{reference} | Update a nsgroup:stubmember object



## Get

> ListNsgroupStubmemberResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve nsgroup:stubmember objects



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
	resp, r, err := apiClient.NsgroupstubmemberAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupstubmemberAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListNsgroupStubmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupstubmemberAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupstubmemberAPIGetRequest` struct via the builder pattern


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

[**ListNsgroupStubmemberResponse**](ListNsgroupStubmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateNsgroupStubmemberResponse Post(ctx).NsgroupStubmember(nsgroupStubmember).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a nsgroup:stubmember object



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
	nsgroupStubmember := *dns.NewNsgroupStubmember() // NsgroupStubmember | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupstubmemberAPI.Post(context.Background()).NsgroupStubmember(nsgroupStubmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupstubmemberAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateNsgroupStubmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupstubmemberAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupstubmemberAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupStubmember** | [**NsgroupStubmember**](NsgroupStubmember.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateNsgroupStubmemberResponse**](CreateNsgroupStubmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Put

> ListNsgroupStubmemberResponse Put(ctx).NsgroupStubmember(nsgroupStubmember).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).MaxResults(maxResults).Method(method).Execute()

Use PUT call as GET operation with _method for a Struct field of a nsgroup:stubmember object



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
	nsgroupStubmember := *dns.NewNsgroupStubmember() // NsgroupStubmember | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupstubmemberAPI.Put(context.Background()).NsgroupStubmember(nsgroupStubmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupstubmemberAPI.Put``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Put`: ListNsgroupStubmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupstubmemberAPI.Put`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupstubmemberAPIPutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupStubmember** | [**NsgroupStubmember**](NsgroupStubmember.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**method** | **string** | Enter the method type for the request | 

### Return type

[**ListNsgroupStubmemberResponse**](ListNsgroupStubmemberResponse.md)

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

Delete a nsgroup:stubmember object



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
	reference := "reference_example" // string | Reference of the nsgroup:stubmember object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.NsgroupstubmemberAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupstubmemberAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:stubmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupstubmemberAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetNsgroupStubmemberResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific nsgroup:stubmember object



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
	reference := "reference_example" // string | Reference of the nsgroup:stubmember object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupstubmemberAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupstubmemberAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetNsgroupStubmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupstubmemberAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:stubmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupstubmemberAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetNsgroupStubmemberResponse**](GetNsgroupStubmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateNsgroupStubmemberResponse ReferencePut(ctx, reference).NsgroupStubmember(nsgroupStubmember).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a nsgroup:stubmember object



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
	reference := "reference_example" // string | Reference of the nsgroup:stubmember object
	nsgroupStubmember := *dns.NewNsgroupStubmember() // NsgroupStubmember | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupstubmemberAPI.ReferencePut(context.Background(), reference).NsgroupStubmember(nsgroupStubmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupstubmemberAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateNsgroupStubmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupstubmemberAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:stubmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupstubmemberAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupStubmember** | [**NsgroupStubmember**](NsgroupStubmember.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateNsgroupStubmemberResponse**](UpdateNsgroupStubmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

