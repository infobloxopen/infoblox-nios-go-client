# SyslogEndpointAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SyslogendpointGet**](SyslogEndpointAPI.md#SyslogendpointGet) | **Get** /syslog:endpoint | Retrieve syslog:endpoint objects
[**SyslogendpointPost**](SyslogEndpointAPI.md#SyslogendpointPost) | **Post** /syslog:endpoint | Create a syslog:endpoint object
[**SyslogendpointReferenceDelete**](SyslogEndpointAPI.md#SyslogendpointReferenceDelete) | **Delete** /syslog:endpoint/{reference} | Delete a syslog:endpoint object
[**SyslogendpointReferenceGet**](SyslogEndpointAPI.md#SyslogendpointReferenceGet) | **Get** /syslog:endpoint/{reference} | Get a specific syslog:endpoint object
[**SyslogendpointReferencePut**](SyslogEndpointAPI.md#SyslogendpointReferencePut) | **Put** /syslog:endpoint/{reference} | Update a syslog:endpoint object



## SyslogendpointGet

> ListSyslogEndpointResponse SyslogendpointGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve syslog:endpoint objects



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
	resp, r, err := apiClient.SyslogEndpointAPI.SyslogendpointGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogEndpointAPI.SyslogendpointGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyslogendpointGet`: ListSyslogEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `SyslogEndpointAPI.SyslogendpointGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SyslogEndpointAPISyslogendpointGetRequest` struct via the builder pattern


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

[**ListSyslogEndpointResponse**](ListSyslogEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyslogendpointPost

> CreateSyslogEndpointResponse SyslogendpointPost(ctx).SyslogEndpoint(syslogEndpoint).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a syslog:endpoint object



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
	syslogEndpoint := *misc.NewSyslogEndpoint() // SyslogEndpoint | Object data to create

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.SyslogEndpointAPI.SyslogendpointPost(context.Background()).SyslogEndpoint(syslogEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogEndpointAPI.SyslogendpointPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyslogendpointPost`: CreateSyslogEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `SyslogEndpointAPI.SyslogendpointPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `SyslogEndpointAPISyslogendpointPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**syslogEndpoint** | [**SyslogEndpoint**](SyslogEndpoint.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateSyslogEndpointResponse**](CreateSyslogEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyslogendpointReferenceDelete

> SyslogendpointReferenceDelete(ctx, reference).Execute()

Delete a syslog:endpoint object



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
	reference := "reference_example" // string | Reference of the syslog:endpoint object

	apiClient := misc.NewAPIClient()
	r, err := apiClient.SyslogEndpointAPI.SyslogendpointReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogEndpointAPI.SyslogendpointReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the syslog:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `SyslogEndpointAPISyslogendpointReferenceDeleteRequest` struct via the builder pattern


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


## SyslogendpointReferenceGet

> GetSyslogEndpointResponse SyslogendpointReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific syslog:endpoint object



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
	reference := "reference_example" // string | Reference of the syslog:endpoint object

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.SyslogEndpointAPI.SyslogendpointReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogEndpointAPI.SyslogendpointReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyslogendpointReferenceGet`: GetSyslogEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `SyslogEndpointAPI.SyslogendpointReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the syslog:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `SyslogEndpointAPISyslogendpointReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetSyslogEndpointResponse**](GetSyslogEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyslogendpointReferencePut

> UpdateSyslogEndpointResponse SyslogendpointReferencePut(ctx, reference).SyslogEndpoint(syslogEndpoint).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a syslog:endpoint object



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
	reference := "reference_example" // string | Reference of the syslog:endpoint object
	syslogEndpoint := *misc.NewSyslogEndpoint() // SyslogEndpoint | Object data to update

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.SyslogEndpointAPI.SyslogendpointReferencePut(context.Background(), reference).SyslogEndpoint(syslogEndpoint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SyslogEndpointAPI.SyslogendpointReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyslogendpointReferencePut`: UpdateSyslogEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `SyslogEndpointAPI.SyslogendpointReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the syslog:endpoint object | 

### Other Parameters

Other parameters are passed through a pointer to a `SyslogEndpointAPISyslogendpointReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**syslogEndpoint** | [**SyslogEndpoint**](SyslogEndpoint.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateSyslogEndpointResponse**](UpdateSyslogEndpointResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

