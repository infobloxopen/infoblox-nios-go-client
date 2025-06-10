# NsgroupStubmemberAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NsgroupstubmemberGet**](NsgroupStubmemberAPI.md#NsgroupstubmemberGet) | **Get** /nsgroup:stubmember | Retrieve nsgroup:stubmember objects
[**NsgroupstubmemberPost**](NsgroupStubmemberAPI.md#NsgroupstubmemberPost) | **Post** /nsgroup:stubmember | Create a nsgroup:stubmember object
[**NsgroupstubmemberReferenceDelete**](NsgroupStubmemberAPI.md#NsgroupstubmemberReferenceDelete) | **Delete** /nsgroup:stubmember/{reference} | Delete a nsgroup:stubmember object
[**NsgroupstubmemberReferenceGet**](NsgroupStubmemberAPI.md#NsgroupstubmemberReferenceGet) | **Get** /nsgroup:stubmember/{reference} | Get a specific nsgroup:stubmember object
[**NsgroupstubmemberReferencePut**](NsgroupStubmemberAPI.md#NsgroupstubmemberReferencePut) | **Put** /nsgroup:stubmember/{reference} | Update a nsgroup:stubmember object



## NsgroupstubmemberGet

> ListNsgroupStubmemberResponse NsgroupstubmemberGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

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
	resp, r, err := apiClient.NsgroupStubmemberAPI.NsgroupstubmemberGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupStubmemberAPI.NsgroupstubmemberGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupstubmemberGet`: ListNsgroupStubmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupStubmemberAPI.NsgroupstubmemberGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupStubmemberAPINsgroupstubmemberGetRequest` struct via the builder pattern


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


## NsgroupstubmemberPost

> CreateNsgroupStubmemberResponse NsgroupstubmemberPost(ctx).NsgroupStubmember(nsgroupStubmember).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.NsgroupStubmemberAPI.NsgroupstubmemberPost(context.Background()).NsgroupStubmember(nsgroupStubmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupStubmemberAPI.NsgroupstubmemberPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupstubmemberPost`: CreateNsgroupStubmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupStubmemberAPI.NsgroupstubmemberPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupStubmemberAPINsgroupstubmemberPostRequest` struct via the builder pattern


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


## NsgroupstubmemberReferenceDelete

> NsgroupstubmemberReferenceDelete(ctx, reference).Execute()

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
	r, err := apiClient.NsgroupStubmemberAPI.NsgroupstubmemberReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupStubmemberAPI.NsgroupstubmemberReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `NsgroupStubmemberAPINsgroupstubmemberReferenceDeleteRequest` struct via the builder pattern


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


## NsgroupstubmemberReferenceGet

> GetNsgroupStubmemberResponse NsgroupstubmemberReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.NsgroupStubmemberAPI.NsgroupstubmemberReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupStubmemberAPI.NsgroupstubmemberReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupstubmemberReferenceGet`: GetNsgroupStubmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupStubmemberAPI.NsgroupstubmemberReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:stubmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupStubmemberAPINsgroupstubmemberReferenceGetRequest` struct via the builder pattern


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


## NsgroupstubmemberReferencePut

> UpdateNsgroupStubmemberResponse NsgroupstubmemberReferencePut(ctx, reference).NsgroupStubmember(nsgroupStubmember).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.NsgroupStubmemberAPI.NsgroupstubmemberReferencePut(context.Background(), reference).NsgroupStubmember(nsgroupStubmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupStubmemberAPI.NsgroupstubmemberReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NsgroupstubmemberReferencePut`: UpdateNsgroupStubmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupStubmemberAPI.NsgroupstubmemberReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:stubmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupStubmemberAPINsgroupstubmemberReferencePutRequest` struct via the builder pattern


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

