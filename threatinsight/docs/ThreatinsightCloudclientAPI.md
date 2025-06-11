# ThreatinsightCloudclientAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThreatinsightcloudclientGet**](ThreatinsightCloudclientAPI.md#ThreatinsightcloudclientGet) | **Get** /threatinsight:cloudclient | Retrieve threatinsight:cloudclient objects
[**ThreatinsightcloudclientReferenceGet**](ThreatinsightCloudclientAPI.md#ThreatinsightcloudclientReferenceGet) | **Get** /threatinsight:cloudclient/{reference} | Get a specific threatinsight:cloudclient object
[**ThreatinsightcloudclientReferencePut**](ThreatinsightCloudclientAPI.md#ThreatinsightcloudclientReferencePut) | **Put** /threatinsight:cloudclient/{reference} | Update a threatinsight:cloudclient object



## ThreatinsightcloudclientGet

> ListThreatinsightCloudclientResponse ThreatinsightcloudclientGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve threatinsight:cloudclient objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightCloudclientAPI.ThreatinsightcloudclientGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightCloudclientAPI.ThreatinsightcloudclientGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightcloudclientGet`: ListThreatinsightCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightCloudclientAPI.ThreatinsightcloudclientGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightCloudclientAPIThreatinsightcloudclientGetRequest` struct via the builder pattern


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

[**ListThreatinsightCloudclientResponse**](ListThreatinsightCloudclientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatinsightcloudclientReferenceGet

> GetThreatinsightCloudclientResponse ThreatinsightcloudclientReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific threatinsight:cloudclient object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {
	reference := "reference_example" // string | Reference of the threatinsight:cloudclient object

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightCloudclientAPI.ThreatinsightcloudclientReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightCloudclientAPI.ThreatinsightcloudclientReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightcloudclientReferenceGet`: GetThreatinsightCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightCloudclientAPI.ThreatinsightcloudclientReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatinsight:cloudclient object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightCloudclientAPIThreatinsightcloudclientReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetThreatinsightCloudclientResponse**](GetThreatinsightCloudclientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThreatinsightcloudclientReferencePut

> UpdateThreatinsightCloudclientResponse ThreatinsightcloudclientReferencePut(ctx, reference).ThreatinsightCloudclient(threatinsightCloudclient).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a threatinsight:cloudclient object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/threatinsight"
)

func main() {
	reference := "reference_example" // string | Reference of the threatinsight:cloudclient object
	threatinsightCloudclient := *threatinsight.NewThreatinsightCloudclient() // ThreatinsightCloudclient | Object data to update

	apiClient := threatinsight.NewAPIClient()
	resp, r, err := apiClient.ThreatinsightCloudclientAPI.ThreatinsightcloudclientReferencePut(context.Background(), reference).ThreatinsightCloudclient(threatinsightCloudclient).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThreatinsightCloudclientAPI.ThreatinsightcloudclientReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ThreatinsightcloudclientReferencePut`: UpdateThreatinsightCloudclientResponse
	fmt.Fprintf(os.Stdout, "Response from `ThreatinsightCloudclientAPI.ThreatinsightcloudclientReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the threatinsight:cloudclient object | 

### Other Parameters

Other parameters are passed through a pointer to a `ThreatinsightCloudclientAPIThreatinsightcloudclientReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**threatinsightCloudclient** | [**ThreatinsightCloudclient**](ThreatinsightCloudclient.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateThreatinsightCloudclientResponse**](UpdateThreatinsightCloudclientResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

