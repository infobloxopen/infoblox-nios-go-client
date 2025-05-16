# DiscoverycredentialgroupAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](DiscoverycredentialgroupAPI.md#Get) | **Get** /discovery:credentialgroup | Retrieve discovery:credentialgroup objects
[**Post**](DiscoverycredentialgroupAPI.md#Post) | **Post** /discovery:credentialgroup | Create a discovery:credentialgroup object
[**ReferenceDelete**](DiscoverycredentialgroupAPI.md#ReferenceDelete) | **Delete** /discovery:credentialgroup/{reference} | Delete a discovery:credentialgroup object
[**ReferenceGet**](DiscoverycredentialgroupAPI.md#ReferenceGet) | **Get** /discovery:credentialgroup/{reference} | Get a specific discovery:credentialgroup object
[**ReferencePut**](DiscoverycredentialgroupAPI.md#ReferencePut) | **Put** /discovery:credentialgroup/{reference} | Update a discovery:credentialgroup object



## Get

> ListDiscoveryCredentialgroupResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve discovery:credentialgroup objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoverycredentialgroupAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoverycredentialgroupAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListDiscoveryCredentialgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoverycredentialgroupAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoverycredentialgroupAPIGetRequest` struct via the builder pattern


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

[**ListDiscoveryCredentialgroupResponse**](ListDiscoveryCredentialgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateDiscoveryCredentialgroupResponse Post(ctx).DiscoveryCredentialgroup(discoveryCredentialgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a discovery:credentialgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {
	discoveryCredentialgroup := *discovery.NewDiscoveryCredentialgroup() // DiscoveryCredentialgroup | Object data to create

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoverycredentialgroupAPI.Post(context.Background()).DiscoveryCredentialgroup(discoveryCredentialgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoverycredentialgroupAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateDiscoveryCredentialgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoverycredentialgroupAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoverycredentialgroupAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**discoveryCredentialgroup** | [**DiscoveryCredentialgroup**](DiscoveryCredentialgroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDiscoveryCredentialgroupResponse**](CreateDiscoveryCredentialgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

Delete a discovery:credentialgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {
	reference := "reference_example" // string | Reference of the discovery:credentialgroup object

	apiClient := discovery.NewAPIClient()
	r, err := apiClient.DiscoverycredentialgroupAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoverycredentialgroupAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:credentialgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoverycredentialgroupAPIReferenceDeleteRequest` struct via the builder pattern


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


## ReferenceGet

> GetDiscoveryCredentialgroupResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific discovery:credentialgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {
	reference := "reference_example" // string | Reference of the discovery:credentialgroup object

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoverycredentialgroupAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoverycredentialgroupAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetDiscoveryCredentialgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoverycredentialgroupAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:credentialgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoverycredentialgroupAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDiscoveryCredentialgroupResponse**](GetDiscoveryCredentialgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateDiscoveryCredentialgroupResponse ReferencePut(ctx, reference).DiscoveryCredentialgroup(discoveryCredentialgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a discovery:credentialgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/discovery"
)

func main() {
	reference := "reference_example" // string | Reference of the discovery:credentialgroup object
	discoveryCredentialgroup := *discovery.NewDiscoveryCredentialgroup() // DiscoveryCredentialgroup | Object data to update

	apiClient := discovery.NewAPIClient()
	resp, r, err := apiClient.DiscoverycredentialgroupAPI.ReferencePut(context.Background(), reference).DiscoveryCredentialgroup(discoveryCredentialgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoverycredentialgroupAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateDiscoveryCredentialgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoverycredentialgroupAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:credentialgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoverycredentialgroupAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**discoveryCredentialgroup** | [**DiscoveryCredentialgroup**](DiscoveryCredentialgroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDiscoveryCredentialgroupResponse**](UpdateDiscoveryCredentialgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

