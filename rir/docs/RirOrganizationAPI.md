# RirOrganizationAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RirorganizationGet**](RirOrganizationAPI.md#RirorganizationGet) | **Get** /rir:organization | Retrieve rir:organization objects
[**RirorganizationPost**](RirOrganizationAPI.md#RirorganizationPost) | **Post** /rir:organization | Create a rir:organization object
[**RirorganizationReferenceDelete**](RirOrganizationAPI.md#RirorganizationReferenceDelete) | **Delete** /rir:organization/{reference} | Delete a rir:organization object
[**RirorganizationReferenceGet**](RirOrganizationAPI.md#RirorganizationReferenceGet) | **Get** /rir:organization/{reference} | Get a specific rir:organization object
[**RirorganizationReferencePut**](RirOrganizationAPI.md#RirorganizationReferencePut) | **Put** /rir:organization/{reference} | Update a rir:organization object



## RirorganizationGet

> ListRirOrganizationResponse RirorganizationGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve rir:organization objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rir"
)

func main() {

	apiClient := rir.NewAPIClient()
	resp, r, err := apiClient.RirOrganizationAPI.RirorganizationGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RirOrganizationAPI.RirorganizationGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RirorganizationGet`: ListRirOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `RirOrganizationAPI.RirorganizationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RirOrganizationAPIRirorganizationGetRequest` struct via the builder pattern


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

[**ListRirOrganizationResponse**](ListRirOrganizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RirorganizationPost

> CreateRirOrganizationResponse RirorganizationPost(ctx).RirOrganization(rirOrganization).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a rir:organization object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rir"
)

func main() {
	rirOrganization := *rir.NewRirOrganization() // RirOrganization | Object data to create

	apiClient := rir.NewAPIClient()
	resp, r, err := apiClient.RirOrganizationAPI.RirorganizationPost(context.Background()).RirOrganization(rirOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RirOrganizationAPI.RirorganizationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RirorganizationPost`: CreateRirOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `RirOrganizationAPI.RirorganizationPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RirOrganizationAPIRirorganizationPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**rirOrganization** | [**RirOrganization**](RirOrganization.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRirOrganizationResponse**](CreateRirOrganizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RirorganizationReferenceDelete

> RirorganizationReferenceDelete(ctx, reference).Execute()

Delete a rir:organization object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rir"
)

func main() {
	reference := "reference_example" // string | Reference of the rir:organization object

	apiClient := rir.NewAPIClient()
	r, err := apiClient.RirOrganizationAPI.RirorganizationReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RirOrganizationAPI.RirorganizationReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the rir:organization object | 

### Other Parameters

Other parameters are passed through a pointer to a `RirOrganizationAPIRirorganizationReferenceDeleteRequest` struct via the builder pattern


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


## RirorganizationReferenceGet

> GetRirOrganizationResponse RirorganizationReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific rir:organization object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rir"
)

func main() {
	reference := "reference_example" // string | Reference of the rir:organization object

	apiClient := rir.NewAPIClient()
	resp, r, err := apiClient.RirOrganizationAPI.RirorganizationReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RirOrganizationAPI.RirorganizationReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RirorganizationReferenceGet`: GetRirOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `RirOrganizationAPI.RirorganizationReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the rir:organization object | 

### Other Parameters

Other parameters are passed through a pointer to a `RirOrganizationAPIRirorganizationReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRirOrganizationResponse**](GetRirOrganizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RirorganizationReferencePut

> UpdateRirOrganizationResponse RirorganizationReferencePut(ctx, reference).RirOrganization(rirOrganization).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a rir:organization object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rir"
)

func main() {
	reference := "reference_example" // string | Reference of the rir:organization object
	rirOrganization := *rir.NewRirOrganization() // RirOrganization | Object data to update

	apiClient := rir.NewAPIClient()
	resp, r, err := apiClient.RirOrganizationAPI.RirorganizationReferencePut(context.Background(), reference).RirOrganization(rirOrganization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RirOrganizationAPI.RirorganizationReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RirorganizationReferencePut`: UpdateRirOrganizationResponse
	fmt.Fprintf(os.Stdout, "Response from `RirOrganizationAPI.RirorganizationReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the rir:organization object | 

### Other Parameters

Other parameters are passed through a pointer to a `RirOrganizationAPIRirorganizationReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**rirOrganization** | [**RirOrganization**](RirOrganization.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRirOrganizationResponse**](UpdateRirOrganizationResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

