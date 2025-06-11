# MemberParentalcontrolAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MemberparentalcontrolGet**](MemberParentalcontrolAPI.md#MemberparentalcontrolGet) | **Get** /member:parentalcontrol | Retrieve member:parentalcontrol objects
[**MemberparentalcontrolReferenceGet**](MemberParentalcontrolAPI.md#MemberparentalcontrolReferenceGet) | **Get** /member:parentalcontrol/{reference} | Get a specific member:parentalcontrol object
[**MemberparentalcontrolReferencePut**](MemberParentalcontrolAPI.md#MemberparentalcontrolReferencePut) | **Put** /member:parentalcontrol/{reference} | Update a member:parentalcontrol object



## MemberparentalcontrolGet

> ListMemberParentalcontrolResponse MemberparentalcontrolGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve member:parentalcontrol objects



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
	resp, r, err := apiClient.MemberParentalcontrolAPI.MemberparentalcontrolGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberParentalcontrolAPI.MemberparentalcontrolGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberparentalcontrolGet`: ListMemberParentalcontrolResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberParentalcontrolAPI.MemberparentalcontrolGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MemberParentalcontrolAPIMemberparentalcontrolGetRequest` struct via the builder pattern


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

[**ListMemberParentalcontrolResponse**](ListMemberParentalcontrolResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberparentalcontrolReferenceGet

> GetMemberParentalcontrolResponse MemberparentalcontrolReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific member:parentalcontrol object



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
	reference := "reference_example" // string | Reference of the member:parentalcontrol object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberParentalcontrolAPI.MemberparentalcontrolReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberParentalcontrolAPI.MemberparentalcontrolReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberparentalcontrolReferenceGet`: GetMemberParentalcontrolResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberParentalcontrolAPI.MemberparentalcontrolReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:parentalcontrol object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberParentalcontrolAPIMemberparentalcontrolReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMemberParentalcontrolResponse**](GetMemberParentalcontrolResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MemberparentalcontrolReferencePut

> UpdateMemberParentalcontrolResponse MemberparentalcontrolReferencePut(ctx, reference).MemberParentalcontrol(memberParentalcontrol).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a member:parentalcontrol object



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
	reference := "reference_example" // string | Reference of the member:parentalcontrol object
	memberParentalcontrol := *grid.NewMemberParentalcontrol() // MemberParentalcontrol | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.MemberParentalcontrolAPI.MemberparentalcontrolReferencePut(context.Background(), reference).MemberParentalcontrol(memberParentalcontrol).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MemberParentalcontrolAPI.MemberparentalcontrolReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MemberparentalcontrolReferencePut`: UpdateMemberParentalcontrolResponse
	fmt.Fprintf(os.Stdout, "Response from `MemberParentalcontrolAPI.MemberparentalcontrolReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the member:parentalcontrol object | 

### Other Parameters

Other parameters are passed through a pointer to a `MemberParentalcontrolAPIMemberparentalcontrolReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**memberParentalcontrol** | [**MemberParentalcontrol**](MemberParentalcontrol.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMemberParentalcontrolResponse**](UpdateMemberParentalcontrolResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

