# DdnsPrincipalclusterGroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DdnsprincipalclustergroupGet**](DdnsPrincipalclusterGroupAPI.md#DdnsprincipalclustergroupGet) | **Get** /ddns:principalcluster:group | Retrieve ddns:principalcluster:group objects
[**DdnsprincipalclustergroupPost**](DdnsPrincipalclusterGroupAPI.md#DdnsprincipalclustergroupPost) | **Post** /ddns:principalcluster:group | Create a ddns:principalcluster:group object
[**DdnsprincipalclustergroupReferenceDelete**](DdnsPrincipalclusterGroupAPI.md#DdnsprincipalclustergroupReferenceDelete) | **Delete** /ddns:principalcluster:group/{reference} | Delete a ddns:principalcluster:group object
[**DdnsprincipalclustergroupReferenceGet**](DdnsPrincipalclusterGroupAPI.md#DdnsprincipalclustergroupReferenceGet) | **Get** /ddns:principalcluster:group/{reference} | Get a specific ddns:principalcluster:group object
[**DdnsprincipalclustergroupReferencePut**](DdnsPrincipalclusterGroupAPI.md#DdnsprincipalclustergroupReferencePut) | **Put** /ddns:principalcluster:group/{reference} | Update a ddns:principalcluster:group object



## DdnsprincipalclustergroupGet

> ListDdnsPrincipalclusterGroupResponse DdnsprincipalclustergroupGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve ddns:principalcluster:group objects



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
	resp, r, err := apiClient.DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdnsprincipalclustergroupGet`: ListDdnsPrincipalclusterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPIDdnsprincipalclustergroupGetRequest` struct via the builder pattern


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

[**ListDdnsPrincipalclusterGroupResponse**](ListDdnsPrincipalclusterGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DdnsprincipalclustergroupPost

> CreateDdnsPrincipalclusterGroupResponse DdnsprincipalclustergroupPost(ctx).DdnsPrincipalclusterGroup(ddnsPrincipalclusterGroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a ddns:principalcluster:group object



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
	ddnsPrincipalclusterGroup := *dns.NewDdnsPrincipalclusterGroup() // DdnsPrincipalclusterGroup | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupPost(context.Background()).DdnsPrincipalclusterGroup(ddnsPrincipalclusterGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdnsprincipalclustergroupPost`: CreateDdnsPrincipalclusterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPIDdnsprincipalclustergroupPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ddnsPrincipalclusterGroup** | [**DdnsPrincipalclusterGroup**](DdnsPrincipalclusterGroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDdnsPrincipalclusterGroupResponse**](CreateDdnsPrincipalclusterGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DdnsprincipalclustergroupReferenceDelete

> DdnsprincipalclustergroupReferenceDelete(ctx, reference).Execute()

Delete a ddns:principalcluster:group object



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
	reference := "reference_example" // string | Reference of the ddns:principalcluster:group object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPIDdnsprincipalclustergroupReferenceDeleteRequest` struct via the builder pattern


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


## DdnsprincipalclustergroupReferenceGet

> GetDdnsPrincipalclusterGroupResponse DdnsprincipalclustergroupReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific ddns:principalcluster:group object



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
	reference := "reference_example" // string | Reference of the ddns:principalcluster:group object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdnsprincipalclustergroupReferenceGet`: GetDdnsPrincipalclusterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPIDdnsprincipalclustergroupReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDdnsPrincipalclusterGroupResponse**](GetDdnsPrincipalclusterGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DdnsprincipalclustergroupReferencePut

> UpdateDdnsPrincipalclusterGroupResponse DdnsprincipalclustergroupReferencePut(ctx, reference).DdnsPrincipalclusterGroup(ddnsPrincipalclusterGroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a ddns:principalcluster:group object



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
	reference := "reference_example" // string | Reference of the ddns:principalcluster:group object
	ddnsPrincipalclusterGroup := *dns.NewDdnsPrincipalclusterGroup() // DdnsPrincipalclusterGroup | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupReferencePut(context.Background(), reference).DdnsPrincipalclusterGroup(ddnsPrincipalclusterGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DdnsprincipalclustergroupReferencePut`: UpdateDdnsPrincipalclusterGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DdnsPrincipalclusterGroupAPI.DdnsprincipalclustergroupReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ddns:principalcluster:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `DdnsPrincipalclusterGroupAPIDdnsprincipalclustergroupReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ddnsPrincipalclusterGroup** | [**DdnsPrincipalclusterGroup**](DdnsPrincipalclusterGroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDdnsPrincipalclusterGroupResponse**](UpdateDdnsPrincipalclusterGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

