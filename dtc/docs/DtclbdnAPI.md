# DtcLbdnAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtclbdnGet**](DtcLbdnAPI.md#DtclbdnGet) | **Get** /dtc:lbdn | Retrieve dtc:lbdn objects
[**DtclbdnPost**](DtcLbdnAPI.md#DtclbdnPost) | **Post** /dtc:lbdn | Create a dtc:lbdn object
[**DtclbdnReferenceDelete**](DtcLbdnAPI.md#DtclbdnReferenceDelete) | **Delete** /dtc:lbdn/{reference} | Delete a dtc:lbdn object
[**DtclbdnReferenceGet**](DtcLbdnAPI.md#DtclbdnReferenceGet) | **Get** /dtc:lbdn/{reference} | Get a specific dtc:lbdn object
[**DtclbdnReferencePut**](DtcLbdnAPI.md#DtclbdnReferencePut) | **Put** /dtc:lbdn/{reference} | Update a dtc:lbdn object



## DtclbdnGet

> ListDtcLbdnResponse DtclbdnGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:lbdn objects



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
	resp, r, err := apiClient.DtcLbdnAPI.DtclbdnGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcLbdnAPI.DtclbdnGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtclbdnGet`: ListDtcLbdnResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcLbdnAPI.DtclbdnGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcLbdnAPIDtclbdnGetRequest` struct via the builder pattern


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

[**ListDtcLbdnResponse**](ListDtcLbdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtclbdnPost

> CreateDtcLbdnResponse DtclbdnPost(ctx).DtcLbdn(dtcLbdn).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:lbdn object



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
	dtcLbdn := *dtc.NewDtcLbdn() // DtcLbdn | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcLbdnAPI.DtclbdnPost(context.Background()).DtcLbdn(dtcLbdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcLbdnAPI.DtclbdnPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtclbdnPost`: CreateDtcLbdnResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcLbdnAPI.DtclbdnPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcLbdnAPIDtclbdnPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcLbdn** | [**DtcLbdn**](DtcLbdn.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcLbdnResponse**](CreateDtcLbdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtclbdnReferenceDelete

> DtclbdnReferenceDelete(ctx, reference).Execute()

Delete a dtc:lbdn object



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
	reference := "reference_example" // string | Reference of the dtc:lbdn object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcLbdnAPI.DtclbdnReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcLbdnAPI.DtclbdnReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:lbdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcLbdnAPIDtclbdnReferenceDeleteRequest` struct via the builder pattern


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


## DtclbdnReferenceGet

> GetDtcLbdnResponse DtclbdnReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:lbdn object



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
	reference := "reference_example" // string | Reference of the dtc:lbdn object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcLbdnAPI.DtclbdnReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcLbdnAPI.DtclbdnReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtclbdnReferenceGet`: GetDtcLbdnResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcLbdnAPI.DtclbdnReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:lbdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcLbdnAPIDtclbdnReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcLbdnResponse**](GetDtcLbdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtclbdnReferencePut

> UpdateDtcLbdnResponse DtclbdnReferencePut(ctx, reference).DtcLbdn(dtcLbdn).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:lbdn object



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
	reference := "reference_example" // string | Reference of the dtc:lbdn object
	dtcLbdn := *dtc.NewDtcLbdn() // DtcLbdn | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcLbdnAPI.DtclbdnReferencePut(context.Background(), reference).DtcLbdn(dtcLbdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcLbdnAPI.DtclbdnReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtclbdnReferencePut`: UpdateDtcLbdnResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcLbdnAPI.DtclbdnReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:lbdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcLbdnAPIDtclbdnReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcLbdn** | [**DtcLbdn**](DtcLbdn.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcLbdnResponse**](UpdateDtcLbdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

