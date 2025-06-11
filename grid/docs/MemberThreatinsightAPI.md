# MemberThreatinsightAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberthreatinsightGet**](MemberThreatinsightAPI.md#MemberthreatinsightGet) | **Get** /member:threatinsight | Retrieve member:threatinsight objects
[**MemberthreatinsightReferenceGet**](MemberThreatinsightAPI.md#MemberthreatinsightReferenceGet) | **Get** /member:threatinsight/{reference} | Get a specific member:threatinsight object
[**MemberthreatinsightReferencePut**](MemberThreatinsightAPI.md#MemberthreatinsightReferencePut) | **Put** /member:threatinsight/{reference} | Update a member:threatinsight object



## MemberthreatinsightGet

> ListMemberThreatinsightResponse MemberthreatinsightGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve member:threatinsight objects



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
	resp, r, err := apiClient.MemberThreatinsightAPI.MemberthreatinsightGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberThreatinsightAPI.MemberthreatinsightGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberthreatinsightGet`: ListMemberThreatinsightResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberThreatinsightAPI.MemberthreatinsightGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MemberThreatinsightAPIMemberthreatinsightGetRequest` struct via the builder pattern


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

[**ListMemberThreatinsightResponse**](ListMemberThreatinsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberthreatinsightReferenceGet

> GetMemberThreatinsightResponse MemberthreatinsightReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific member:threatinsight object



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
	reference := "reference_example" // string | Reference of the member:threatinsight object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberThreatinsightAPI.MemberthreatinsightReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberThreatinsightAPI.MemberthreatinsightReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberthreatinsightReferenceGet`: GetMemberThreatinsightResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberThreatinsightAPI.MemberthreatinsightReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:threatinsight object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberThreatinsightAPIMemberthreatinsightReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMemberThreatinsightResponse**](GetMemberThreatinsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberthreatinsightReferencePut

> UpdateMemberThreatinsightResponse MemberthreatinsightReferencePut(ctx, reference).MemberThreatinsight(memberThreatinsight).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a member:threatinsight object



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
	reference := "reference_example" // string | Reference of the member:threatinsight object
	memberThreatinsight := *grid.NewMemberThreatinsight() // MemberThreatinsight | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberThreatinsightAPI.MemberthreatinsightReferencePut(context.Background(), reference).MemberThreatinsight(memberThreatinsight).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberThreatinsightAPI.MemberthreatinsightReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberthreatinsightReferencePut`: UpdateMemberThreatinsightResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberThreatinsightAPI.MemberthreatinsightReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:threatinsight object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberThreatinsightAPIMemberthreatinsightReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**memberThreatinsight** | [**MemberThreatinsight**](MemberThreatinsight.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMemberThreatinsightResponse**](UpdateMemberThreatinsightResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

