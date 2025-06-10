# DtcTopologyAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtctopologyGet**](DtcTopologyAPI.md#DtctopologyGet) | **Get** /dtc:topology | Retrieve dtc:topology objects
[**DtctopologyPost**](DtcTopologyAPI.md#DtctopologyPost) | **Post** /dtc:topology | Create a dtc:topology object
[**DtctopologyReferenceDelete**](DtcTopologyAPI.md#DtctopologyReferenceDelete) | **Delete** /dtc:topology/{reference} | Delete a dtc:topology object
[**DtctopologyReferenceGet**](DtcTopologyAPI.md#DtctopologyReferenceGet) | **Get** /dtc:topology/{reference} | Get a specific dtc:topology object
[**DtctopologyReferencePut**](DtcTopologyAPI.md#DtctopologyReferencePut) | **Put** /dtc:topology/{reference} | Update a dtc:topology object



## DtctopologyGet

> ListDtcTopologyResponse DtctopologyGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:topology objects



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
	resp, r, err := apiClient.DtcTopologyAPI.DtctopologyGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyAPI.DtctopologyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtctopologyGet`: ListDtcTopologyResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyAPI.DtctopologyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyAPIDtctopologyGetRequest` struct via the builder pattern


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

[**ListDtcTopologyResponse**](ListDtcTopologyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtctopologyPost

> CreateDtcTopologyResponse DtctopologyPost(ctx).DtcTopology(dtcTopology).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:topology object



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
	dtcTopology := *dtc.NewDtcTopology() // DtcTopology | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyAPI.DtctopologyPost(context.Background()).DtcTopology(dtcTopology).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyAPI.DtctopologyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtctopologyPost`: CreateDtcTopologyResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyAPI.DtctopologyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyAPIDtctopologyPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcTopology** | [**DtcTopology**](DtcTopology.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcTopologyResponse**](CreateDtcTopologyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtctopologyReferenceDelete

> DtctopologyReferenceDelete(ctx, reference).Execute()

Delete a dtc:topology object



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
	reference := "reference_example" // string | Reference of the dtc:topology object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcTopologyAPI.DtctopologyReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyAPI.DtctopologyReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:topology object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyAPIDtctopologyReferenceDeleteRequest` struct via the builder pattern


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


## DtctopologyReferenceGet

> GetDtcTopologyResponse DtctopologyReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:topology object



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
	reference := "reference_example" // string | Reference of the dtc:topology object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyAPI.DtctopologyReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyAPI.DtctopologyReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtctopologyReferenceGet`: GetDtcTopologyResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyAPI.DtctopologyReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:topology object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyAPIDtctopologyReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcTopologyResponse**](GetDtcTopologyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtctopologyReferencePut

> UpdateDtcTopologyResponse DtctopologyReferencePut(ctx, reference).DtcTopology(dtcTopology).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:topology object



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
	reference := "reference_example" // string | Reference of the dtc:topology object
	dtcTopology := *dtc.NewDtcTopology() // DtcTopology | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyAPI.DtctopologyReferencePut(context.Background(), reference).DtcTopology(dtcTopology).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyAPI.DtctopologyReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtctopologyReferencePut`: UpdateDtcTopologyResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyAPI.DtctopologyReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:topology object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyAPIDtctopologyReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcTopology** | [**DtcTopology**](DtcTopology.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcTopologyResponse**](UpdateDtcTopologyResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

