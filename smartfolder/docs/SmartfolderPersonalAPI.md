# SmartfolderPersonalAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SmartfolderpersonalGet**](SmartfolderPersonalAPI.md#SmartfolderpersonalGet) | **Get** /smartfolder:personal | Retrieve smartfolder:personal objects
[**SmartfolderpersonalPost**](SmartfolderPersonalAPI.md#SmartfolderpersonalPost) | **Post** /smartfolder:personal | Create a smartfolder:personal object
[**SmartfolderpersonalReferenceDelete**](SmartfolderPersonalAPI.md#SmartfolderpersonalReferenceDelete) | **Delete** /smartfolder:personal/{reference} | Delete a smartfolder:personal object
[**SmartfolderpersonalReferenceGet**](SmartfolderPersonalAPI.md#SmartfolderpersonalReferenceGet) | **Get** /smartfolder:personal/{reference} | Get a specific smartfolder:personal object
[**SmartfolderpersonalReferencePut**](SmartfolderPersonalAPI.md#SmartfolderpersonalReferencePut) | **Put** /smartfolder:personal/{reference} | Update a smartfolder:personal object



## SmartfolderpersonalGet

> ListSmartfolderPersonalResponse SmartfolderpersonalGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve smartfolder:personal objects



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
	resp, r, err := apiClient.SmartfolderPersonalAPI.SmartfolderpersonalGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.SmartfolderpersonalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmartfolderpersonalGet`: ListSmartfolderPersonalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderPersonalAPI.SmartfolderpersonalGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPISmartfolderpersonalGetRequest` struct via the builder pattern


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

[**ListSmartfolderPersonalResponse**](ListSmartfolderPersonalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartfolderpersonalPost

> CreateSmartfolderPersonalResponse SmartfolderpersonalPost(ctx).SmartfolderPersonal(smartfolderPersonal).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a smartfolder:personal object



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
	smartfolderPersonal := *smartfolder.NewSmartfolderPersonal() // SmartfolderPersonal | Object data to create

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderPersonalAPI.SmartfolderpersonalPost(context.Background()).SmartfolderPersonal(smartfolderPersonal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.SmartfolderpersonalPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmartfolderpersonalPost`: CreateSmartfolderPersonalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderPersonalAPI.SmartfolderpersonalPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPISmartfolderpersonalPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**smartfolderPersonal** | [**SmartfolderPersonal**](SmartfolderPersonal.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSmartfolderPersonalResponse**](CreateSmartfolderPersonalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartfolderpersonalReferenceDelete

> SmartfolderpersonalReferenceDelete(ctx, reference).Execute()

Delete a smartfolder:personal object



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
	reference := "reference_example" // string | Reference of the smartfolder:personal object

	apiClient := smartfolder.NewAPIClient()
	r, err := apiClient.SmartfolderPersonalAPI.SmartfolderpersonalReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.SmartfolderpersonalReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:personal object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPISmartfolderpersonalReferenceDeleteRequest` struct via the builder pattern


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


## SmartfolderpersonalReferenceGet

> GetSmartfolderPersonalResponse SmartfolderpersonalReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific smartfolder:personal object



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
	reference := "reference_example" // string | Reference of the smartfolder:personal object

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderPersonalAPI.SmartfolderpersonalReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.SmartfolderpersonalReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmartfolderpersonalReferenceGet`: GetSmartfolderPersonalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderPersonalAPI.SmartfolderpersonalReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:personal object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPISmartfolderpersonalReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSmartfolderPersonalResponse**](GetSmartfolderPersonalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SmartfolderpersonalReferencePut

> UpdateSmartfolderPersonalResponse SmartfolderpersonalReferencePut(ctx, reference).SmartfolderPersonal(smartfolderPersonal).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a smartfolder:personal object



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
	reference := "reference_example" // string | Reference of the smartfolder:personal object
	smartfolderPersonal := *smartfolder.NewSmartfolderPersonal() // SmartfolderPersonal | Object data to update

	apiClient := smartfolder.NewAPIClient()
	resp, r, err := apiClient.SmartfolderPersonalAPI.SmartfolderpersonalReferencePut(context.Background(), reference).SmartfolderPersonal(smartfolderPersonal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SmartfolderPersonalAPI.SmartfolderpersonalReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SmartfolderpersonalReferencePut`: UpdateSmartfolderPersonalResponse
	fmt.Fprintf(os.Stdout, "Response from `SmartfolderPersonalAPI.SmartfolderpersonalReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the smartfolder:personal object | 

### Other Parameters

Other parameters are passed through a pointer to a `SmartfolderPersonalAPISmartfolderpersonalReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**smartfolderPersonal** | [**SmartfolderPersonal**](SmartfolderPersonal.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSmartfolderPersonalResponse**](UpdateSmartfolderPersonalResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

