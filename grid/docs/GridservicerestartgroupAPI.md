# GridServicerestartGroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GridservicerestartgroupGet**](GridServicerestartGroupAPI.md#GridservicerestartgroupGet) | **Get** /grid:servicerestart:group | Retrieve grid:servicerestart:group objects
[**GridservicerestartgroupPost**](GridServicerestartGroupAPI.md#GridservicerestartgroupPost) | **Post** /grid:servicerestart:group | Create a grid:servicerestart:group object
[**GridservicerestartgroupReferenceDelete**](GridServicerestartGroupAPI.md#GridservicerestartgroupReferenceDelete) | **Delete** /grid:servicerestart:group/{reference} | Delete a grid:servicerestart:group object
[**GridservicerestartgroupReferenceGet**](GridServicerestartGroupAPI.md#GridservicerestartgroupReferenceGet) | **Get** /grid:servicerestart:group/{reference} | Get a specific grid:servicerestart:group object
[**GridservicerestartgroupReferencePut**](GridServicerestartGroupAPI.md#GridservicerestartgroupReferencePut) | **Put** /grid:servicerestart:group/{reference} | Update a grid:servicerestart:group object



## GridservicerestartgroupGet

> ListGridServicerestartGroupResponse GridservicerestartgroupGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve grid:servicerestart:group objects



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
	resp, r, err := apiClient.GridServicerestartGroupAPI.GridservicerestartgroupGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupAPI.GridservicerestartgroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridservicerestartgroupGet`: ListGridServicerestartGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupAPI.GridservicerestartgroupGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupAPIGridservicerestartgroupGetRequest` struct via the builder pattern


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

[**ListGridServicerestartGroupResponse**](ListGridServicerestartGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridservicerestartgroupPost

> CreateGridServicerestartGroupResponse GridservicerestartgroupPost(ctx).GridServicerestartGroup(gridServicerestartGroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a grid:servicerestart:group object



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
	gridServicerestartGroup := *grid.NewGridServicerestartGroup() // GridServicerestartGroup | Object data to create

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridServicerestartGroupAPI.GridservicerestartgroupPost(context.Background()).GridServicerestartGroup(gridServicerestartGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupAPI.GridservicerestartgroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridservicerestartgroupPost`: CreateGridServicerestartGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupAPI.GridservicerestartgroupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupAPIGridservicerestartgroupPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridServicerestartGroup** | [**GridServicerestartGroup**](GridServicerestartGroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateGridServicerestartGroupResponse**](CreateGridServicerestartGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridservicerestartgroupReferenceDelete

> GridservicerestartgroupReferenceDelete(ctx, reference).Execute()

Delete a grid:servicerestart:group object



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
	reference := "reference_example" // string | Reference of the grid:servicerestart:group object

	apiClient := grid.NewAPIClient()
	r, err := apiClient.GridServicerestartGroupAPI.GridservicerestartgroupReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupAPI.GridservicerestartgroupReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:servicerestart:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupAPIGridservicerestartgroupReferenceDeleteRequest` struct via the builder pattern


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


## GridservicerestartgroupReferenceGet

> GetGridServicerestartGroupResponse GridservicerestartgroupReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific grid:servicerestart:group object



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
	reference := "reference_example" // string | Reference of the grid:servicerestart:group object

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridServicerestartGroupAPI.GridservicerestartgroupReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupAPI.GridservicerestartgroupReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridservicerestartgroupReferenceGet`: GetGridServicerestartGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupAPI.GridservicerestartgroupReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:servicerestart:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupAPIGridservicerestartgroupReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetGridServicerestartGroupResponse**](GetGridServicerestartGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GridservicerestartgroupReferencePut

> UpdateGridServicerestartGroupResponse GridservicerestartgroupReferencePut(ctx, reference).GridServicerestartGroup(gridServicerestartGroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a grid:servicerestart:group object



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
	reference := "reference_example" // string | Reference of the grid:servicerestart:group object
	gridServicerestartGroup := *grid.NewGridServicerestartGroup() // GridServicerestartGroup | Object data to update

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridServicerestartGroupAPI.GridservicerestartgroupReferencePut(context.Background(), reference).GridServicerestartGroup(gridServicerestartGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridServicerestartGroupAPI.GridservicerestartgroupReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GridservicerestartgroupReferencePut`: UpdateGridServicerestartGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GridServicerestartGroupAPI.GridservicerestartgroupReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:servicerestart:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridServicerestartGroupAPIGridservicerestartgroupReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridServicerestartGroup** | [**GridServicerestartGroup**](GridServicerestartGroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateGridServicerestartGroupResponse**](UpdateGridServicerestartGroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

