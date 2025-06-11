# LocaluserAuthserviceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocaluserauthserviceGet**](LocaluserAuthserviceAPI.md#LocaluserauthserviceGet) | **Get** /localuser:authservice | Retrieve localuser:authservice objects
[**LocaluserauthserviceReferenceGet**](LocaluserAuthserviceAPI.md#LocaluserauthserviceReferenceGet) | **Get** /localuser:authservice/{reference} | Get a specific localuser:authservice object



## LocaluserauthserviceGet

> ListLocaluserAuthserviceResponse LocaluserauthserviceGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve localuser:authservice objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.LocaluserAuthserviceAPI.LocaluserauthserviceGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocaluserAuthserviceAPI.LocaluserauthserviceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocaluserauthserviceGet`: ListLocaluserAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `LocaluserAuthserviceAPI.LocaluserauthserviceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `LocaluserAuthserviceAPILocaluserauthserviceGetRequest` struct via the builder pattern


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

[**ListLocaluserAuthserviceResponse**](ListLocaluserAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocaluserauthserviceReferenceGet

> GetLocaluserAuthserviceResponse LocaluserauthserviceReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific localuser:authservice object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the localuser:authservice object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.LocaluserAuthserviceAPI.LocaluserauthserviceReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LocaluserAuthserviceAPI.LocaluserauthserviceReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LocaluserauthserviceReferenceGet`: GetLocaluserAuthserviceResponse
	fmt.Fprintf(os.Stdout, "Response from `LocaluserAuthserviceAPI.LocaluserauthserviceReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the localuser:authservice object | 

### Other Parameters

Other parameters are passed through a pointer to a `LocaluserAuthserviceAPILocaluserauthserviceReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetLocaluserAuthserviceResponse**](GetLocaluserAuthserviceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

