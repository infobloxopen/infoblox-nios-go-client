# DiscoveryCredentialgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DiscoverycredentialgroupGet**](DiscoveryCredentialgroupAPI.md#DiscoverycredentialgroupGet) | **Get** /discovery:credentialgroup | Retrieve discovery:credentialgroup objects
[**DiscoverycredentialgroupPost**](DiscoveryCredentialgroupAPI.md#DiscoverycredentialgroupPost) | **Post** /discovery:credentialgroup | Create a discovery:credentialgroup object
[**DiscoverycredentialgroupReferenceDelete**](DiscoveryCredentialgroupAPI.md#DiscoverycredentialgroupReferenceDelete) | **Delete** /discovery:credentialgroup/{reference} | Delete a discovery:credentialgroup object
[**DiscoverycredentialgroupReferenceGet**](DiscoveryCredentialgroupAPI.md#DiscoverycredentialgroupReferenceGet) | **Get** /discovery:credentialgroup/{reference} | Get a specific discovery:credentialgroup object
[**DiscoverycredentialgroupReferencePut**](DiscoveryCredentialgroupAPI.md#DiscoverycredentialgroupReferencePut) | **Put** /discovery:credentialgroup/{reference} | Update a discovery:credentialgroup object



## DiscoverycredentialgroupGet

> ListDiscoveryCredentialgroupResponse DiscoverycredentialgroupGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

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
	resp, r, err := apiClient.DiscoveryCredentialgroupAPI.DiscoverycredentialgroupGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCredentialgroupAPI.DiscoverycredentialgroupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverycredentialgroupGet`: ListDiscoveryCredentialgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryCredentialgroupAPI.DiscoverycredentialgroupGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryCredentialgroupAPIDiscoverycredentialgroupGetRequest` struct via the builder pattern


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


## DiscoverycredentialgroupPost

> CreateDiscoveryCredentialgroupResponse DiscoverycredentialgroupPost(ctx).DiscoveryCredentialgroup(discoveryCredentialgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.DiscoveryCredentialgroupAPI.DiscoverycredentialgroupPost(context.Background()).DiscoveryCredentialgroup(discoveryCredentialgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCredentialgroupAPI.DiscoverycredentialgroupPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverycredentialgroupPost`: CreateDiscoveryCredentialgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryCredentialgroupAPI.DiscoverycredentialgroupPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryCredentialgroupAPIDiscoverycredentialgroupPostRequest` struct via the builder pattern


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


## DiscoverycredentialgroupReferenceDelete

> DiscoverycredentialgroupReferenceDelete(ctx, reference).Execute()

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
	r, err := apiClient.DiscoveryCredentialgroupAPI.DiscoverycredentialgroupReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCredentialgroupAPI.DiscoverycredentialgroupReferenceDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a `DiscoveryCredentialgroupAPIDiscoverycredentialgroupReferenceDeleteRequest` struct via the builder pattern


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


## DiscoverycredentialgroupReferenceGet

> GetDiscoveryCredentialgroupResponse DiscoverycredentialgroupReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.DiscoveryCredentialgroupAPI.DiscoverycredentialgroupReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCredentialgroupAPI.DiscoverycredentialgroupReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverycredentialgroupReferenceGet`: GetDiscoveryCredentialgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryCredentialgroupAPI.DiscoverycredentialgroupReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:credentialgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryCredentialgroupAPIDiscoverycredentialgroupReferenceGetRequest` struct via the builder pattern


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


## DiscoverycredentialgroupReferencePut

> UpdateDiscoveryCredentialgroupResponse DiscoverycredentialgroupReferencePut(ctx, reference).DiscoveryCredentialgroup(discoveryCredentialgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

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
	resp, r, err := apiClient.DiscoveryCredentialgroupAPI.DiscoverycredentialgroupReferencePut(context.Background(), reference).DiscoveryCredentialgroup(discoveryCredentialgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DiscoveryCredentialgroupAPI.DiscoverycredentialgroupReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DiscoverycredentialgroupReferencePut`: UpdateDiscoveryCredentialgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `DiscoveryCredentialgroupAPI.DiscoverycredentialgroupReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the discovery:credentialgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `DiscoveryCredentialgroupAPIDiscoverycredentialgroupReferencePutRequest` struct via the builder pattern


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

