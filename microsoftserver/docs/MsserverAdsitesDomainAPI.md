# MsserverAdsitesDomainAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MsserveradsitesdomainGet**](MsserverAdsitesDomainAPI.md#MsserveradsitesdomainGet) | **Get** /msserver:adsites:domain | Retrieve msserver:adsites:domain objects
[**MsserveradsitesdomainReferenceGet**](MsserverAdsitesDomainAPI.md#MsserveradsitesdomainReferenceGet) | **Get** /msserver:adsites:domain/{reference} | Get a specific msserver:adsites:domain object



## MsserveradsitesdomainGet

> ListMsserverAdsitesDomainResponse MsserveradsitesdomainGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve msserver:adsites:domain objects



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
	resp, r, err := apiClient.MsserverAdsitesDomainAPI.MsserveradsitesdomainGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverAdsitesDomainAPI.MsserveradsitesdomainGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserveradsitesdomainGet`: ListMsserverAdsitesDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverAdsitesDomainAPI.MsserveradsitesdomainGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `MsserverAdsitesDomainAPIMsserveradsitesdomainGetRequest` struct via the builder pattern


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

[**ListMsserverAdsitesDomainResponse**](ListMsserverAdsitesDomainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MsserveradsitesdomainReferenceGet

> GetMsserverAdsitesDomainResponse MsserveradsitesdomainReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific msserver:adsites:domain object



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
	reference := "reference_example" // string | Reference of the msserver:adsites:domain object

	apiClient := microsoftserver.NewAPIClient()
	resp, r, err := apiClient.MsserverAdsitesDomainAPI.MsserveradsitesdomainReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MsserverAdsitesDomainAPI.MsserveradsitesdomainReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MsserveradsitesdomainReferenceGet`: GetMsserverAdsitesDomainResponse
	fmt.Fprintf(os.Stdout, "Response from `MsserverAdsitesDomainAPI.MsserveradsitesdomainReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the msserver:adsites:domain object | 

### Other Parameters

Other parameters are passed through a pointer to a `MsserverAdsitesDomainAPIMsserveradsitesdomainReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetMsserverAdsitesDomainResponse**](GetMsserverAdsitesDomainResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

