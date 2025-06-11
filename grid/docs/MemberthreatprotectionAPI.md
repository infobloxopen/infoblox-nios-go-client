# MemberThreatprotectionAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberthreatprotectionGet**](MemberThreatprotectionAPI.md#MemberthreatprotectionGet) | **Get** /member:threatprotection | Retrieve member:threatprotection objects
[**MemberthreatprotectionReferenceGet**](MemberThreatprotectionAPI.md#MemberthreatprotectionReferenceGet) | **Get** /member:threatprotection/{reference} | Get a specific member:threatprotection object
[**MemberthreatprotectionReferencePut**](MemberThreatprotectionAPI.md#MemberthreatprotectionReferencePut) | **Put** /member:threatprotection/{reference} | Update a member:threatprotection object



## MemberthreatprotectionGet

> ListMemberThreatprotectionResponse MemberthreatprotectionGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve member:threatprotection objects



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
	resp, r, err := apiClient.MemberThreatprotectionAPI.MemberthreatprotectionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberThreatprotectionAPI.MemberthreatprotectionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberthreatprotectionGet`: ListMemberThreatprotectionResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberThreatprotectionAPI.MemberthreatprotectionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MemberThreatprotectionAPIMemberthreatprotectionGetRequest` struct via the builder pattern


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

[**ListMemberThreatprotectionResponse**](ListMemberThreatprotectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberthreatprotectionReferenceGet

> GetMemberThreatprotectionResponse MemberthreatprotectionReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific member:threatprotection object



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
	reference := "reference_example" // string | Reference of the member:threatprotection object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberThreatprotectionAPI.MemberthreatprotectionReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberThreatprotectionAPI.MemberthreatprotectionReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberthreatprotectionReferenceGet`: GetMemberThreatprotectionResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberThreatprotectionAPI.MemberthreatprotectionReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:threatprotection object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberThreatprotectionAPIMemberthreatprotectionReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMemberThreatprotectionResponse**](GetMemberThreatprotectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberthreatprotectionReferencePut

> UpdateMemberThreatprotectionResponse MemberthreatprotectionReferencePut(ctx, reference).MemberThreatprotection(memberThreatprotection).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a member:threatprotection object



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
	reference := "reference_example" // string | Reference of the member:threatprotection object
	memberThreatprotection := *grid.NewMemberThreatprotection() // MemberThreatprotection | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberThreatprotectionAPI.MemberthreatprotectionReferencePut(context.Background(), reference).MemberThreatprotection(memberThreatprotection).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberThreatprotectionAPI.MemberthreatprotectionReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberthreatprotectionReferencePut`: UpdateMemberThreatprotectionResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberThreatprotectionAPI.MemberthreatprotectionReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:threatprotection object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberThreatprotectionAPIMemberthreatprotectionReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**memberThreatprotection** | [**MemberThreatprotection**](MemberThreatprotection.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMemberThreatprotectionResponse**](UpdateMemberThreatprotectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

