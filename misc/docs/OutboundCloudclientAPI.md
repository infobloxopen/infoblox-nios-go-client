# OutboundCloudclientAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OutboundcloudclientGet**](OutboundCloudclientAPI.md#OutboundcloudclientGet) | **Get** /outbound:cloudclient | Retrieve outbound:cloudclient objects
[**OutboundcloudclientReferenceGet**](OutboundCloudclientAPI.md#OutboundcloudclientReferenceGet) | **Get** /outbound:cloudclient/{reference} | Get a specific outbound:cloudclient object
[**OutboundcloudclientReferencePut**](OutboundCloudclientAPI.md#OutboundcloudclientReferencePut) | **Put** /outbound:cloudclient/{reference} | Update a outbound:cloudclient object



## OutboundcloudclientGet

> ListOutboundCloudclientResponse OutboundcloudclientGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve outbound:cloudclient objects



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
	resp, r, err := apiClient.OutboundCloudclientAPI.OutboundcloudclientGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundCloudclientAPI.OutboundcloudclientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutboundcloudclientGet`: ListOutboundCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundCloudclientAPI.OutboundcloudclientGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `OutboundCloudclientAPIOutboundcloudclientGetRequest` struct via the builder pattern


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

[**ListOutboundCloudclientResponse**](ListOutboundCloudclientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutboundcloudclientReferenceGet

> GetOutboundCloudclientResponse OutboundcloudclientReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific outbound:cloudclient object



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
	reference := "reference_example" // string | Reference of the outbound:cloudclient object

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.OutboundCloudclientAPI.OutboundcloudclientReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundCloudclientAPI.OutboundcloudclientReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutboundcloudclientReferenceGet`: GetOutboundCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundCloudclientAPI.OutboundcloudclientReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the outbound:cloudclient object | 

### Other Parameters

Other parameters are passed through a pointer to a `OutboundCloudclientAPIOutboundcloudclientReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetOutboundCloudclientResponse**](GetOutboundCloudclientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OutboundcloudclientReferencePut

> UpdateOutboundCloudclientResponse OutboundcloudclientReferencePut(ctx, reference).OutboundCloudclient(outboundCloudclient).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a outbound:cloudclient object



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
	reference := "reference_example" // string | Reference of the outbound:cloudclient object
	outboundCloudclient := *misc.NewOutboundCloudclient() // OutboundCloudclient | Object data to update

	apiClient := misc.NewAPIClient()
	resp, r, err := apiClient.OutboundCloudclientAPI.OutboundcloudclientReferencePut(context.Background(), reference).OutboundCloudclient(outboundCloudclient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundCloudclientAPI.OutboundcloudclientReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OutboundcloudclientReferencePut`: UpdateOutboundCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundCloudclientAPI.OutboundcloudclientReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the outbound:cloudclient object | 

### Other Parameters

Other parameters are passed through a pointer to a `OutboundCloudclientAPIOutboundcloudclientReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**outboundCloudclient** | [**OutboundCloudclient**](OutboundCloudclient.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateOutboundCloudclientResponse**](UpdateOutboundCloudclientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

