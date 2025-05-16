# MemberdhcppropertiesAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](MemberdhcppropertiesAPI.md#Get) | **Get** /member:dhcpproperties | Retrieve member:dhcpproperties objects
[**ReferenceGet**](MemberdhcppropertiesAPI.md#ReferenceGet) | **Get** /member:dhcpproperties/{reference} | Get a specific member:dhcpproperties object
[**ReferencePut**](MemberdhcppropertiesAPI.md#ReferencePut) | **Put** /member:dhcpproperties/{reference} | Update a member:dhcpproperties object



## Get

> ListMemberDhcppropertiesResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve member:dhcpproperties objects



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
	resp, r, err := apiClient.MemberdhcppropertiesAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberdhcppropertiesAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListMemberDhcppropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberdhcppropertiesAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MemberdhcppropertiesAPIGetRequest` struct via the builder pattern


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


## ReferenceGet

> GetMemberDhcppropertiesResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific member:dhcpproperties object



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
	reference := "reference_example" // string | Reference of the member:dhcpproperties object

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.MemberdhcppropertiesAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberdhcppropertiesAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetMemberDhcppropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberdhcppropertiesAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:dhcpproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberdhcppropertiesAPIReferenceGetRequest` struct via the builder pattern


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


## ReferencePut

> UpdateMemberDhcppropertiesResponse ReferencePut(ctx, reference).MemberDhcpproperties(memberDhcpproperties).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a member:dhcpproperties object



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
	reference := "reference_example" // string | Reference of the member:dhcpproperties object
	memberDhcpproperties := *dhcp.NewMemberDhcpproperties() // MemberDhcpproperties | Object data to update

	apiClient := dhcp.NewAPIClient()
	resp, r, err := apiClient.MemberdhcppropertiesAPI.ReferencePut(context.Background(), reference).MemberDhcpproperties(memberDhcpproperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberdhcppropertiesAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateMemberDhcppropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberdhcppropertiesAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:dhcpproperties object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberdhcppropertiesAPIReferencePutRequest` struct via the builder pattern


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

