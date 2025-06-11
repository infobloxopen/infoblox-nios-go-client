# MemberDhcppropertiesAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberdhcppropertiesGet**](MemberDhcppropertiesAPI.md#MemberdhcppropertiesGet) | **Get** /member:dhcpproperties | Retrieve member:dhcpproperties objects
[**MemberdhcppropertiesReferenceGet**](MemberDhcppropertiesAPI.md#MemberdhcppropertiesReferenceGet) | **Get** /member:dhcpproperties/{reference} | Get a specific member:dhcpproperties object
[**MemberdhcppropertiesReferencePut**](MemberDhcppropertiesAPI.md#MemberdhcppropertiesReferencePut) | **Put** /member:dhcpproperties/{reference} | Update a member:dhcpproperties object



## MemberdhcppropertiesGet

> ListMemberDhcppropertiesResponse MemberdhcppropertiesGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve member:dhcpproperties objects



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
	resp, r, err := apiClient.MemberDhcppropertiesAPI.MemberdhcppropertiesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDhcppropertiesAPI.MemberdhcppropertiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberdhcppropertiesGet`: ListMemberDhcppropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberDhcppropertiesAPI.MemberdhcppropertiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MemberDhcppropertiesAPIMemberdhcppropertiesGetRequest` struct via the builder pattern


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

[**ListMemberDhcppropertiesResponse**](ListMemberDhcppropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberdhcppropertiesReferenceGet

> GetMemberDhcppropertiesResponse MemberdhcppropertiesReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific member:dhcpproperties object



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
	reference := "reference_example" // string | Reference of the member:dhcpproperties object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberDhcppropertiesAPI.MemberdhcppropertiesReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDhcppropertiesAPI.MemberdhcppropertiesReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberdhcppropertiesReferenceGet`: GetMemberDhcppropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberDhcppropertiesAPI.MemberdhcppropertiesReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:dhcpproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberDhcppropertiesAPIMemberdhcppropertiesReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMemberDhcppropertiesResponse**](GetMemberDhcppropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberdhcppropertiesReferencePut

> UpdateMemberDhcppropertiesResponse MemberdhcppropertiesReferencePut(ctx, reference).MemberDhcpproperties(memberDhcpproperties).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a member:dhcpproperties object



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
	reference := "reference_example" // string | Reference of the member:dhcpproperties object
	memberDhcpproperties := *grid.NewMemberDhcpproperties() // MemberDhcpproperties | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberDhcppropertiesAPI.MemberdhcppropertiesReferencePut(context.Background(), reference).MemberDhcpproperties(memberDhcpproperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberDhcppropertiesAPI.MemberdhcppropertiesReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberdhcppropertiesReferencePut`: UpdateMemberDhcppropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberDhcppropertiesAPI.MemberdhcppropertiesReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:dhcpproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberDhcppropertiesAPIMemberdhcppropertiesReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**memberDhcpproperties** | [**MemberDhcpproperties**](MemberDhcpproperties.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMemberDhcppropertiesResponse**](UpdateMemberDhcppropertiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

