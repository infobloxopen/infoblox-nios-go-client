# NsgroupForwardingmemberAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](NsgroupForwardingmemberAPI.md#Create) | **Post** /nsgroup:forwardingmember | Create a nsgroup:forwardingmember object
[**Delete**](NsgroupForwardingmemberAPI.md#Delete) | **Delete** /nsgroup:forwardingmember/{reference} | Delete a nsgroup:forwardingmember object
[**List**](NsgroupForwardingmemberAPI.md#List) | **Get** /nsgroup:forwardingmember | Retrieve nsgroup:forwardingmember objects
[**Read**](NsgroupForwardingmemberAPI.md#Read) | **Get** /nsgroup:forwardingmember/{reference} | Get a specific nsgroup:forwardingmember object
[**Update**](NsgroupForwardingmemberAPI.md#Update) | **Put** /nsgroup:forwardingmember/{reference} | Update a nsgroup:forwardingmember object



## Create

> CreateNsgroupForwardingmemberResponse Create(ctx).NsgroupForwardingmember(nsgroupForwardingmember).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a nsgroup:forwardingmember object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	nsgroupForwardingmember := *dns.NewNsgroupForwardingmember() // NsgroupForwardingmember | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardingmemberAPI.Create(context.Background()).NsgroupForwardingmember(nsgroupForwardingmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardingmemberAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupForwardingmember** | [**NsgroupForwardingmember**](NsgroupForwardingmember.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateNsgroupForwardingmemberResponse**](CreateNsgroupForwardingmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, reference).Execute()

Delete a nsgroup:forwardingmember object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the nsgroup:forwardingmember object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.NsgroupForwardingmemberAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardingmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPIDeleteRequest` struct via the builder pattern


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


## List

> ListNsgroupForwardingmemberResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve nsgroup:forwardingmember objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardingmemberAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardingmemberAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**maxResults** | **int32** | Enter the number of results to be fetched | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**paging** | **int32** | Control paging of results | 
**pageId** | **string** | Page id for retrieving next page of results | 
**filters** | **map[string]interface{}** |  | 
**extattrfilter** | **map[string]interface{}** |  | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**ListNsgroupForwardingmemberResponse**](ListNsgroupForwardingmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetNsgroupForwardingmemberResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific nsgroup:forwardingmember object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the nsgroup:forwardingmember object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardingmemberAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardingmemberAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardingmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetNsgroupForwardingmemberResponse**](GetNsgroupForwardingmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateNsgroupForwardingmemberResponse Update(ctx, reference).NsgroupForwardingmember(nsgroupForwardingmember).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a nsgroup:forwardingmember object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the nsgroup:forwardingmember object
	nsgroupForwardingmember := *dns.NewNsgroupForwardingmember() // NsgroupForwardingmember | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.NsgroupForwardingmemberAPI.Update(context.Background(), reference).NsgroupForwardingmember(nsgroupForwardingmember).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NsgroupForwardingmemberAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateNsgroupForwardingmemberResponse
	fmt.Fprintf(os.Stdout, "Response from `NsgroupForwardingmemberAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the nsgroup:forwardingmember object | 

### Other Parameters

Other parameters are passed through a pointer to a `NsgroupForwardingmemberAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**nsgroupForwardingmember** | [**NsgroupForwardingmember**](NsgroupForwardingmember.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateNsgroupForwardingmemberResponse**](UpdateNsgroupForwardingmemberResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

