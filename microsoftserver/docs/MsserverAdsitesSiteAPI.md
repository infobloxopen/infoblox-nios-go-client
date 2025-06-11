# MsserverAdsitesSiteAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MsserveradsitessiteGet**](MsserverAdsitesSiteAPI.md#MsserveradsitessiteGet) | **Get** /msserver:adsites:site | Retrieve msserver:adsites:site objects
[**MsserveradsitessitePost**](MsserverAdsitesSiteAPI.md#MsserveradsitessitePost) | **Post** /msserver:adsites:site | Create a msserver:adsites:site object
[**MsserveradsitessiteReferenceDelete**](MsserverAdsitesSiteAPI.md#MsserveradsitessiteReferenceDelete) | **Delete** /msserver:adsites:site/{reference} | Delete a msserver:adsites:site object
[**MsserveradsitessiteReferenceGet**](MsserverAdsitesSiteAPI.md#MsserveradsitessiteReferenceGet) | **Get** /msserver:adsites:site/{reference} | Get a specific msserver:adsites:site object
[**MsserveradsitessiteReferencePut**](MsserverAdsitesSiteAPI.md#MsserveradsitessiteReferencePut) | **Put** /msserver:adsites:site/{reference} | Update a msserver:adsites:site object



## MsserveradsitessiteGet

> ListMsserverAdsitesSiteResponse MsserveradsitessiteGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve msserver:adsites:site objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverAdsitesSiteAPI.MsserveradsitessiteGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverAdsitesSiteAPI.MsserveradsitessiteGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserveradsitessiteGet`: ListMsserverAdsitesSiteResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverAdsitesSiteAPI.MsserveradsitessiteGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MsserverAdsitesSiteAPIMsserveradsitessiteGetRequest` struct via the builder pattern


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

[**ListMsserverAdsitesSiteResponse**](ListMsserverAdsitesSiteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserveradsitessitePost

> CreateMsserverAdsitesSiteResponse MsserveradsitessitePost(ctx).MsserverAdsitesSite(msserverAdsitesSite).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a msserver:adsites:site object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {
	msserverAdsitesSite := *microsoftserver.NewMsserverAdsitesSite() // MsserverAdsitesSite | Object data to create

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverAdsitesSiteAPI.MsserveradsitessitePost(context.Background()).MsserverAdsitesSite(msserverAdsitesSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverAdsitesSiteAPI.MsserveradsitessitePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserveradsitessitePost`: CreateMsserverAdsitesSiteResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverAdsitesSiteAPI.MsserveradsitessitePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MsserverAdsitesSiteAPIMsserveradsitessitePostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**msserverAdsitesSite** | [**MsserverAdsitesSite**](MsserverAdsitesSite.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateMsserverAdsitesSiteResponse**](CreateMsserverAdsitesSiteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserveradsitessiteReferenceDelete

> MsserveradsitessiteReferenceDelete(ctx, reference).Execute()

Delete a msserver:adsites:site object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {
	reference := "reference_example" // string | Reference of the msserver:adsites:site object

	apiClient := microsoftserver.NewAPIClient()
	r, err := apiClient.MsserverAdsitesSiteAPI.MsserveradsitessiteReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverAdsitesSiteAPI.MsserveradsitessiteReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:adsites:site object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverAdsitesSiteAPIMsserveradsitessiteReferenceDeleteRequest` struct via the builder pattern


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


## MsserveradsitessiteReferenceGet

> GetMsserverAdsitesSiteResponse MsserveradsitessiteReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific msserver:adsites:site object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {
	reference := "reference_example" // string | Reference of the msserver:adsites:site object

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverAdsitesSiteAPI.MsserveradsitessiteReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverAdsitesSiteAPI.MsserveradsitessiteReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserveradsitessiteReferenceGet`: GetMsserverAdsitesSiteResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverAdsitesSiteAPI.MsserveradsitessiteReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:adsites:site object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverAdsitesSiteAPIMsserveradsitessiteReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMsserverAdsitesSiteResponse**](GetMsserverAdsitesSiteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserveradsitessiteReferencePut

> UpdateMsserverAdsitesSiteResponse MsserveradsitessiteReferencePut(ctx, reference).MsserverAdsitesSite(msserverAdsitesSite).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a msserver:adsites:site object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/microsoftserver"
)

func main() {
	reference := "reference_example" // string | Reference of the msserver:adsites:site object
	msserverAdsitesSite := *microsoftserver.NewMsserverAdsitesSite() // MsserverAdsitesSite | Object data to update

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverAdsitesSiteAPI.MsserveradsitessiteReferencePut(context.Background(), reference).MsserverAdsitesSite(msserverAdsitesSite).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverAdsitesSiteAPI.MsserveradsitessiteReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserveradsitessiteReferencePut`: UpdateMsserverAdsitesSiteResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverAdsitesSiteAPI.MsserveradsitessiteReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:adsites:site object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverAdsitesSiteAPIMsserveradsitessiteReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**msserverAdsitesSite** | [**MsserverAdsitesSite**](MsserverAdsitesSite.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateMsserverAdsitesSiteResponse**](UpdateMsserverAdsitesSiteResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

