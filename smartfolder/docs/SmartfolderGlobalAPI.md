# SmartfolderGlobalAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartfolderglobalGet**](SmartfolderGlobalAPI.md#SmartfolderglobalGet) | **Get** /smartfolder:global | Retrieve smartfolder:global objects
[**SmartfolderglobalPost**](SmartfolderGlobalAPI.md#SmartfolderglobalPost) | **Post** /smartfolder:global | Create a smartfolder:global object
[**SmartfolderglobalReferenceDelete**](SmartfolderGlobalAPI.md#SmartfolderglobalReferenceDelete) | **Delete** /smartfolder:global/{reference} | Delete a smartfolder:global object
[**SmartfolderglobalReferenceGet**](SmartfolderGlobalAPI.md#SmartfolderglobalReferenceGet) | **Get** /smartfolder:global/{reference} | Get a specific smartfolder:global object
[**SmartfolderglobalReferencePut**](SmartfolderGlobalAPI.md#SmartfolderglobalReferencePut) | **Put** /smartfolder:global/{reference} | Update a smartfolder:global object



## SmartfolderglobalGet

> ListSmartfolderGlobalResponse SmartfolderglobalGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve smartfolder:global objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/smartfolder"
)

func main() {

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderGlobalAPI.SmartfolderglobalGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.SmartfolderglobalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmartfolderglobalGet`: ListSmartfolderGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderGlobalAPI.SmartfolderglobalGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPISmartfolderglobalGetRequest` struct via the builder pattern


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

[**ListSmartfolderGlobalResponse**](ListSmartfolderGlobalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartfolderglobalPost

> CreateSmartfolderGlobalResponse SmartfolderglobalPost(ctx).SmartfolderGlobal(smartfolderGlobal).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a smartfolder:global object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/smartfolder"
)

func main() {
	smartfolderGlobal := *smartfolder.NewSmartfolderGlobal() // SmartfolderGlobal | Object data to create

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderGlobalAPI.SmartfolderglobalPost(context.Background()).SmartfolderGlobal(smartfolderGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.SmartfolderglobalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmartfolderglobalPost`: CreateSmartfolderGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderGlobalAPI.SmartfolderglobalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPISmartfolderglobalPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**smartfolderGlobal** | [**SmartfolderGlobal**](SmartfolderGlobal.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSmartfolderGlobalResponse**](CreateSmartfolderGlobalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartfolderglobalReferenceDelete

> SmartfolderglobalReferenceDelete(ctx, reference).Execute()

Delete a smartfolder:global object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/smartfolder"
)

func main() {
	reference := "reference_example" // string | Reference of the smartfolder:global object

	apiClient := smartfolder.NewAPIClient()
	r, err := apiClient.SmartfolderGlobalAPI.SmartfolderglobalReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.SmartfolderglobalReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:global object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPISmartfolderglobalReferenceDeleteRequest` struct via the builder pattern


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


## SmartfolderglobalReferenceGet

> GetSmartfolderGlobalResponse SmartfolderglobalReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific smartfolder:global object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/smartfolder"
)

func main() {
	reference := "reference_example" // string | Reference of the smartfolder:global object

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderGlobalAPI.SmartfolderglobalReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.SmartfolderglobalReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmartfolderglobalReferenceGet`: GetSmartfolderGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderGlobalAPI.SmartfolderglobalReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:global object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPISmartfolderglobalReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSmartfolderGlobalResponse**](GetSmartfolderGlobalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartfolderglobalReferencePut

> UpdateSmartfolderGlobalResponse SmartfolderglobalReferencePut(ctx, reference).SmartfolderGlobal(smartfolderGlobal).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a smartfolder:global object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/smartfolder"
)

func main() {
	reference := "reference_example" // string | Reference of the smartfolder:global object
	smartfolderGlobal := *smartfolder.NewSmartfolderGlobal() // SmartfolderGlobal | Object data to update

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderGlobalAPI.SmartfolderglobalReferencePut(context.Background(), reference).SmartfolderGlobal(smartfolderGlobal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderGlobalAPI.SmartfolderglobalReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmartfolderglobalReferencePut`: UpdateSmartfolderGlobalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderGlobalAPI.SmartfolderglobalReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:global object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderGlobalAPISmartfolderglobalReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**smartfolderGlobal** | [**SmartfolderGlobal**](SmartfolderGlobal.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSmartfolderGlobalResponse**](UpdateSmartfolderGlobalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

