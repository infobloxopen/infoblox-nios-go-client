# DatacollectionclusterAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](DatacollectionclusterAPI.md#Get) | **Get** /datacollectioncluster | Retrieve datacollectioncluster objects
[**Post**](DatacollectionclusterAPI.md#Post) | **Post** /datacollectioncluster | Create a datacollectioncluster object
[**ReferenceDelete**](DatacollectionclusterAPI.md#ReferenceDelete) | **Delete** /datacollectioncluster/{reference} | Delete a datacollectioncluster object
[**ReferenceGet**](DatacollectionclusterAPI.md#ReferenceGet) | **Get** /datacollectioncluster/{reference} | Get a specific datacollectioncluster object
[**ReferencePut**](DatacollectionclusterAPI.md#ReferencePut) | **Put** /datacollectioncluster/{reference} | Update a datacollectioncluster object



## Get

> ListDatacollectionclusterResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve datacollectioncluster objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.DatacollectionclusterAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatacollectionclusterAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListDatacollectionclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DatacollectionclusterAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DatacollectionclusterAPIGetRequest` struct via the builder pattern


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

[**ListDatacollectionclusterResponse**](ListDatacollectionclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateDatacollectionclusterResponse Post(ctx).Datacollectioncluster(datacollectioncluster).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a datacollectioncluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	datacollectioncluster := *misc.NewDatacollectioncluster() // Datacollectioncluster | Object data to create

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.DatacollectionclusterAPI.Post(context.Background()).Datacollectioncluster(datacollectioncluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatacollectionclusterAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateDatacollectionclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DatacollectionclusterAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DatacollectionclusterAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**datacollectioncluster** | [**Datacollectioncluster**](Datacollectioncluster.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDatacollectionclusterResponse**](CreateDatacollectionclusterResponse.md)

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

Delete a datacollectioncluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the datacollectioncluster object

	apiClient := misc.NewAPIClient()
	r, err := apiClient.DatacollectionclusterAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatacollectionclusterAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the datacollectioncluster object | 

### Other Parameters

Other parameters are passed through a pointer to a `DatacollectionclusterAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetDatacollectionclusterResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific datacollectioncluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the datacollectioncluster object

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.DatacollectionclusterAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatacollectionclusterAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetDatacollectionclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DatacollectionclusterAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the datacollectioncluster object | 

### Other Parameters

Other parameters are passed through a pointer to a `DatacollectionclusterAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDatacollectionclusterResponse**](GetDatacollectionclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateDatacollectionclusterResponse ReferencePut(ctx, reference).Datacollectioncluster(datacollectioncluster).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a datacollectioncluster object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/misc"
)

func main() {
	reference := "reference_example" // string | Reference of the datacollectioncluster object
	datacollectioncluster := *misc.NewDatacollectioncluster() // Datacollectioncluster | Object data to update

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.DatacollectionclusterAPI.ReferencePut(context.Background(), reference).Datacollectioncluster(datacollectioncluster).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatacollectionclusterAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateDatacollectionclusterResponse
	fmt.Fprintf(os.Stdout, "Response from `DatacollectionclusterAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the datacollectioncluster object | 

### Other Parameters

Other parameters are passed through a pointer to a `DatacollectionclusterAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**datacollectioncluster** | [**Datacollectioncluster**](Datacollectioncluster.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDatacollectionclusterResponse**](UpdateDatacollectionclusterResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

