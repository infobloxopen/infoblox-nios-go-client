# ThreatprotectionProfileAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreatprotectionprofileGet**](ThreatprotectionProfileAPI.md#ThreatprotectionprofileGet) | **Get** /threatprotection:profile | Retrieve threatprotection:profile objects
[**ThreatprotectionprofilePost**](ThreatprotectionProfileAPI.md#ThreatprotectionprofilePost) | **Post** /threatprotection:profile | Create a threatprotection:profile object
[**ThreatprotectionprofileReferenceDelete**](ThreatprotectionProfileAPI.md#ThreatprotectionprofileReferenceDelete) | **Delete** /threatprotection:profile/{reference} | Delete a threatprotection:profile object
[**ThreatprotectionprofileReferenceGet**](ThreatprotectionProfileAPI.md#ThreatprotectionprofileReferenceGet) | **Get** /threatprotection:profile/{reference} | Get a specific threatprotection:profile object
[**ThreatprotectionprofileReferencePut**](ThreatprotectionProfileAPI.md#ThreatprotectionprofileReferencePut) | **Put** /threatprotection:profile/{reference} | Update a threatprotection:profile object



## ThreatprotectionprofileGet

> ListThreatprotectionProfileResponse ThreatprotectionprofileGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve threatprotection:profile objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatprotection"
)

func main() {

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionProfileAPI.ThreatprotectionprofileGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.ThreatprotectionprofileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionprofileGet`: ListThreatprotectionProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileAPI.ThreatprotectionprofileGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPIThreatprotectionprofileGetRequest` struct via the builder pattern


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

[**ListThreatprotectionProfileResponse**](ListThreatprotectionProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectionprofilePost

> CreateThreatprotectionProfileResponse ThreatprotectionprofilePost(ctx).ThreatprotectionProfile(threatprotectionProfile).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a threatprotection:profile object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatprotection"
)

func main() {
	threatprotectionProfile := *threatprotection.NewThreatprotectionProfile() // ThreatprotectionProfile | Object data to create

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionProfileAPI.ThreatprotectionprofilePost(context.Background()).ThreatprotectionProfile(threatprotectionProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.ThreatprotectionprofilePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionprofilePost`: CreateThreatprotectionProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileAPI.ThreatprotectionprofilePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPIThreatprotectionprofilePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionProfile** | [**ThreatprotectionProfile**](ThreatprotectionProfile.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateThreatprotectionProfileResponse**](CreateThreatprotectionProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectionprofileReferenceDelete

> ThreatprotectionprofileReferenceDelete(ctx, reference).Execute()

Delete a threatprotection:profile object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatprotection"
)

func main() {
	reference := "reference_example" // string | Reference of the threatprotection:profile object

	apiClient := threatprotection.NewAPIClient()
	r, err := apiClient.ThreatprotectionProfileAPI.ThreatprotectionprofileReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.ThreatprotectionprofileReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:profile object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPIThreatprotectionprofileReferenceDeleteRequest` struct via the builder pattern


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


## ThreatprotectionprofileReferenceGet

> GetThreatprotectionProfileResponse ThreatprotectionprofileReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific threatprotection:profile object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatprotection"
)

func main() {
	reference := "reference_example" // string | Reference of the threatprotection:profile object

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionProfileAPI.ThreatprotectionprofileReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.ThreatprotectionprofileReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionprofileReferenceGet`: GetThreatprotectionProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileAPI.ThreatprotectionprofileReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:profile object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPIThreatprotectionprofileReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetThreatprotectionProfileResponse**](GetThreatprotectionProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatprotectionprofileReferencePut

> UpdateThreatprotectionProfileResponse ThreatprotectionprofileReferencePut(ctx, reference).ThreatprotectionProfile(threatprotectionProfile).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a threatprotection:profile object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatprotection"
)

func main() {
	reference := "reference_example" // string | Reference of the threatprotection:profile object
	threatprotectionProfile := *threatprotection.NewThreatprotectionProfile() // ThreatprotectionProfile | Object data to update

	apiClient := threatprotection.NewAPIClient()
	resp, r, err := apiClient.ThreatprotectionProfileAPI.ThreatprotectionprofileReferencePut(context.Background(), reference).ThreatprotectionProfile(threatprotectionProfile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatprotectionProfileAPI.ThreatprotectionprofileReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatprotectionprofileReferencePut`: UpdateThreatprotectionProfileResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatprotectionProfileAPI.ThreatprotectionprofileReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatprotection:profile object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatprotectionProfileAPIThreatprotectionprofileReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatprotectionProfile** | [**ThreatprotectionProfile**](ThreatprotectionProfile.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateThreatprotectionProfileResponse**](UpdateThreatprotectionProfileResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

