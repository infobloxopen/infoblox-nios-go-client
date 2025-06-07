# GridservicerestartgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](GridservicerestartgroupAPI.md#Get) | **Get** /grid:servicerestart:group | Retrieve grid:servicerestart:group objects
[**Post**](GridservicerestartgroupAPI.md#Post) | **Post** /grid:servicerestart:group | Create a grid:servicerestart:group object
[**ReferenceDelete**](GridservicerestartgroupAPI.md#ReferenceDelete) | **Delete** /grid:servicerestart:group/{reference} | Delete a grid:servicerestart:group object
[**ReferenceGet**](GridservicerestartgroupAPI.md#ReferenceGet) | **Get** /grid:servicerestart:group/{reference} | Get a specific grid:servicerestart:group object
[**ReferencePut**](GridservicerestartgroupAPI.md#ReferencePut) | **Put** /grid:servicerestart:group/{reference} | Update a grid:servicerestart:group object



## Get

> ListGridServicerestartGroupResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

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
	resp, r, err := apiClient.GridservicerestartgroupAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridservicerestartgroupAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListGridServicerestartGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GridservicerestartgroupAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridservicerestartgroupAPIGetRequest` struct via the builder pattern


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


## Post

> CreateGridServicerestartGroupResponse Post(ctx).GridServicerestartGroup(gridServicerestartGroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.GridservicerestartgroupAPI.Post(context.Background()).GridServicerestartGroup(gridServicerestartGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridservicerestartgroupAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateGridServicerestartGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GridservicerestartgroupAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridservicerestartgroupAPIPostRequest` struct via the builder pattern


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


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

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
	r, err := apiClient.GridservicerestartgroupAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridservicerestartgroupAPI.ReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `GridservicerestartgroupAPIReferenceDeleteRequest` struct via the builder pattern


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


## ReferenceGet

> GetGridServicerestartGroupResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.GridservicerestartgroupAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridservicerestartgroupAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetGridServicerestartGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GridservicerestartgroupAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:servicerestart:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridservicerestartgroupAPIReferenceGetRequest` struct via the builder pattern


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


## ReferencePut

> UpdateGridServicerestartGroupResponse ReferencePut(ctx, reference).GridServicerestartGroup(gridServicerestartGroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.GridservicerestartgroupAPI.ReferencePut(context.Background(), reference).GridServicerestartGroup(gridServicerestartGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridservicerestartgroupAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateGridServicerestartGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `GridservicerestartgroupAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the grid:servicerestart:group object | 

### Other Parameters

Other parameters are passed through a pointer to a `GridservicerestartgroupAPIReferencePutRequest` struct via the builder pattern


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

