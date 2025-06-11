# DtcServerAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcserverGet**](DtcServerAPI.md#DtcserverGet) | **Get** /dtc:server | Retrieve dtc:server objects
[**DtcserverPost**](DtcServerAPI.md#DtcserverPost) | **Post** /dtc:server | Create a dtc:server object
[**DtcserverReferenceDelete**](DtcServerAPI.md#DtcserverReferenceDelete) | **Delete** /dtc:server/{reference} | Delete a dtc:server object
[**DtcserverReferenceGet**](DtcServerAPI.md#DtcserverReferenceGet) | **Get** /dtc:server/{reference} | Get a specific dtc:server object
[**DtcserverReferencePut**](DtcServerAPI.md#DtcserverReferencePut) | **Put** /dtc:server/{reference} | Update a dtc:server object



## DtcserverGet

> ListDtcServerResponse DtcserverGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:server objects



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
	resp, r, err := apiClient.DtcServerAPI.DtcserverGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcServerAPI.DtcserverGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcserverGet`: ListDtcServerResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcServerAPI.DtcserverGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcServerAPIDtcserverGetRequest` struct via the builder pattern


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

[**ListDtcServerResponse**](ListDtcServerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcserverPost

> CreateDtcServerResponse DtcserverPost(ctx).DtcServer(dtcServer).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:server object



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
	dtcServer := *dtc.NewDtcServer() // DtcServer | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcServerAPI.DtcserverPost(context.Background()).DtcServer(dtcServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcServerAPI.DtcserverPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcserverPost`: CreateDtcServerResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcServerAPI.DtcserverPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcServerAPIDtcserverPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcServer** | [**DtcServer**](DtcServer.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcServerResponse**](CreateDtcServerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcserverReferenceDelete

> DtcserverReferenceDelete(ctx, reference).Execute()

Delete a dtc:server object



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
	reference := "reference_example" // string | Reference of the dtc:server object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcServerAPI.DtcserverReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcServerAPI.DtcserverReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:server object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcServerAPIDtcserverReferenceDeleteRequest` struct via the builder pattern


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


## DtcserverReferenceGet

> GetDtcServerResponse DtcserverReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:server object



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
	reference := "reference_example" // string | Reference of the dtc:server object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcServerAPI.DtcserverReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcServerAPI.DtcserverReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcserverReferenceGet`: GetDtcServerResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcServerAPI.DtcserverReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:server object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcServerAPIDtcserverReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcServerResponse**](GetDtcServerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcserverReferencePut

> UpdateDtcServerResponse DtcserverReferencePut(ctx, reference).DtcServer(dtcServer).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:server object



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
	reference := "reference_example" // string | Reference of the dtc:server object
	dtcServer := *dtc.NewDtcServer() // DtcServer | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcServerAPI.DtcserverReferencePut(context.Background(), reference).DtcServer(dtcServer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcServerAPI.DtcserverReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcserverReferencePut`: UpdateDtcServerResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcServerAPI.DtcserverReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:server object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcServerAPIDtcserverReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcServer** | [**DtcServer**](DtcServer.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcServerResponse**](UpdateDtcServerResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

