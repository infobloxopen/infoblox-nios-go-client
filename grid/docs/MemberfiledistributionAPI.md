# MemberFiledistributionAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberfiledistributionGet**](MemberFiledistributionAPI.md#MemberfiledistributionGet) | **Get** /member:filedistribution | Retrieve member:filedistribution objects
[**MemberfiledistributionReferenceGet**](MemberFiledistributionAPI.md#MemberfiledistributionReferenceGet) | **Get** /member:filedistribution/{reference} | Get a specific member:filedistribution object
[**MemberfiledistributionReferencePut**](MemberFiledistributionAPI.md#MemberfiledistributionReferencePut) | **Put** /member:filedistribution/{reference} | Update a member:filedistribution object



## MemberfiledistributionGet

> ListMemberFiledistributionResponse MemberfiledistributionGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve member:filedistribution objects



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
	resp, r, err := apiClient.MemberFiledistributionAPI.MemberfiledistributionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberFiledistributionAPI.MemberfiledistributionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberfiledistributionGet`: ListMemberFiledistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberFiledistributionAPI.MemberfiledistributionGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MemberFiledistributionAPIMemberfiledistributionGetRequest` struct via the builder pattern


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

[**ListMemberFiledistributionResponse**](ListMemberFiledistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberfiledistributionReferenceGet

> GetMemberFiledistributionResponse MemberfiledistributionReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific member:filedistribution object



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
	reference := "reference_example" // string | Reference of the member:filedistribution object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberFiledistributionAPI.MemberfiledistributionReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberFiledistributionAPI.MemberfiledistributionReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberfiledistributionReferenceGet`: GetMemberFiledistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberFiledistributionAPI.MemberfiledistributionReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:filedistribution object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberFiledistributionAPIMemberfiledistributionReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMemberFiledistributionResponse**](GetMemberFiledistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberfiledistributionReferencePut

> UpdateMemberFiledistributionResponse MemberfiledistributionReferencePut(ctx, reference).MemberFiledistribution(memberFiledistribution).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a member:filedistribution object



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
	reference := "reference_example" // string | Reference of the member:filedistribution object
	memberFiledistribution := *grid.NewMemberFiledistribution() // MemberFiledistribution | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberFiledistributionAPI.MemberfiledistributionReferencePut(context.Background(), reference).MemberFiledistribution(memberFiledistribution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberFiledistributionAPI.MemberfiledistributionReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberfiledistributionReferencePut`: UpdateMemberFiledistributionResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberFiledistributionAPI.MemberfiledistributionReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:filedistribution object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberFiledistributionAPIMemberfiledistributionReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**memberFiledistribution** | [**MemberFiledistribution**](MemberFiledistribution.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMemberFiledistributionResponse**](UpdateMemberFiledistributionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

