# MembercloudsyncAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](MembercloudsyncAPI.md#Get) | **Get** /membercloudsync | Retrieve membercloudsync objects
[**ReferenceGet**](MembercloudsyncAPI.md#ReferenceGet) | **Get** /membercloudsync/{reference} | Get a specific membercloudsync object
[**ReferencePut**](MembercloudsyncAPI.md#ReferencePut) | **Put** /membercloudsync/{reference} | Update a membercloudsync object



## Get

> ListMembercloudsyncResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve membercloudsync objects



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
	resp, r, err := apiClient.MembercloudsyncAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembercloudsyncAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListMembercloudsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MembercloudsyncAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MembercloudsyncAPIGetRequest` struct via the builder pattern


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

[**ListMembercloudsyncResponse**](ListMembercloudsyncResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceGet

> GetMembercloudsyncResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific membercloudsync object



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
	reference := "reference_example" // string | Reference of the membercloudsync object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MembercloudsyncAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembercloudsyncAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetMembercloudsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MembercloudsyncAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the membercloudsync object | 

### Other Parameters

Other parameters are passed through a pointer to a `MembercloudsyncAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMembercloudsyncResponse**](GetMembercloudsyncResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateMembercloudsyncResponse ReferencePut(ctx, reference).Membercloudsync(membercloudsync).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a membercloudsync object



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
	reference := "reference_example" // string | Reference of the membercloudsync object
	membercloudsync := *grid.NewMembercloudsync() // Membercloudsync | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MembercloudsyncAPI.ReferencePut(context.Background(), reference).Membercloudsync(membercloudsync).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MembercloudsyncAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateMembercloudsyncResponse
	fmt.Fprintf(os.Stdout, "Response from `MembercloudsyncAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the membercloudsync object | 

### Other Parameters

Other parameters are passed through a pointer to a `MembercloudsyncAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**membercloudsync** | [**Membercloudsync**](Membercloudsync.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMembercloudsyncResponse**](UpdateMembercloudsyncResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

