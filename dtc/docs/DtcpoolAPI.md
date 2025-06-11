# DtcPoolAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcpoolGet**](DtcPoolAPI.md#DtcpoolGet) | **Get** /dtc:pool | Retrieve dtc:pool objects
[**DtcpoolPost**](DtcPoolAPI.md#DtcpoolPost) | **Post** /dtc:pool | Create a dtc:pool object
[**DtcpoolReferenceDelete**](DtcPoolAPI.md#DtcpoolReferenceDelete) | **Delete** /dtc:pool/{reference} | Delete a dtc:pool object
[**DtcpoolReferenceGet**](DtcPoolAPI.md#DtcpoolReferenceGet) | **Get** /dtc:pool/{reference} | Get a specific dtc:pool object
[**DtcpoolReferencePut**](DtcPoolAPI.md#DtcpoolReferencePut) | **Put** /dtc:pool/{reference} | Update a dtc:pool object



## DtcpoolGet

> ListDtcPoolResponse DtcpoolGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:pool objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcPoolAPI.DtcpoolGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcPoolAPI.DtcpoolGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcpoolGet`: ListDtcPoolResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcPoolAPI.DtcpoolGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcPoolAPIDtcpoolGetRequest` struct via the builder pattern


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

[**ListDtcPoolResponse**](ListDtcPoolResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcpoolPost

> CreateDtcPoolResponse DtcpoolPost(ctx).DtcPool(dtcPool).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:pool object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	dtcPool := *dtc.NewDtcPool() // DtcPool | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcPoolAPI.DtcpoolPost(context.Background()).DtcPool(dtcPool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcPoolAPI.DtcpoolPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcpoolPost`: CreateDtcPoolResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcPoolAPI.DtcpoolPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcPoolAPIDtcpoolPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcPool** | [**DtcPool**](DtcPool.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcPoolResponse**](CreateDtcPoolResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcpoolReferenceDelete

> DtcpoolReferenceDelete(ctx, reference).Execute()

Delete a dtc:pool object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:pool object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcPoolAPI.DtcpoolReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcPoolAPI.DtcpoolReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:pool object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcPoolAPIDtcpoolReferenceDeleteRequest` struct via the builder pattern


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


## DtcpoolReferenceGet

> GetDtcPoolResponse DtcpoolReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:pool object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:pool object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcPoolAPI.DtcpoolReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcPoolAPI.DtcpoolReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcpoolReferenceGet`: GetDtcPoolResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcPoolAPI.DtcpoolReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:pool object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcPoolAPIDtcpoolReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcPoolResponse**](GetDtcPoolResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcpoolReferencePut

> UpdateDtcPoolResponse DtcpoolReferencePut(ctx, reference).DtcPool(dtcPool).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:pool object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:pool object
	dtcPool := *dtc.NewDtcPool() // DtcPool | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcPoolAPI.DtcpoolReferencePut(context.Background(), reference).DtcPool(dtcPool).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcPoolAPI.DtcpoolReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcpoolReferencePut`: UpdateDtcPoolResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcPoolAPI.DtcpoolReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:pool object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcPoolAPIDtcpoolReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcPool** | [**DtcPool**](DtcPool.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcPoolResponse**](UpdateDtcPoolResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

