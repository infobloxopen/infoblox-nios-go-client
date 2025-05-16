# OutboundcloudclientAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](OutboundcloudclientAPI.md#Get) | **Get** /outbound:cloudclient | Retrieve outbound:cloudclient objects
[**ReferenceGet**](OutboundcloudclientAPI.md#ReferenceGet) | **Get** /outbound:cloudclient/{reference} | Get a specific outbound:cloudclient object
[**ReferencePut**](OutboundcloudclientAPI.md#ReferencePut) | **Put** /outbound:cloudclient/{reference} | Update a outbound:cloudclient object



## Get

> ListOutboundCloudclientResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve outbound:cloudclient objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.OutboundcloudclientAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundcloudclientAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListOutboundCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundcloudclientAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `OutboundcloudclientAPIGetRequest` struct via the builder pattern


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


## ReferenceGet

> GetOutboundCloudclientResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific outbound:cloudclient object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the outbound:cloudclient object

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.OutboundcloudclientAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundcloudclientAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetOutboundCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundcloudclientAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the outbound:cloudclient object | 

### Other Parameters

Other parameters are passed through a pointer to a `OutboundcloudclientAPIReferenceGetRequest` struct via the builder pattern


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


## ReferencePut

> UpdateOutboundCloudclientResponse ReferencePut(ctx, reference).OutboundCloudclient(outboundCloudclient).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a outbound:cloudclient object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the outbound:cloudclient object
	outboundCloudclient := *cloud.NewOutboundCloudclient() // OutboundCloudclient | Object data to update

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.OutboundcloudclientAPI.ReferencePut(context.Background(), reference).OutboundCloudclient(outboundCloudclient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OutboundcloudclientAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateOutboundCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `OutboundcloudclientAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the outbound:cloudclient object | 

### Other Parameters

Other parameters are passed through a pointer to a `OutboundcloudclientAPIReferencePutRequest` struct via the builder pattern


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

