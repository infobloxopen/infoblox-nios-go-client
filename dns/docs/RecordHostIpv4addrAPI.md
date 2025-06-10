# RecordHostIpv4addrAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordhostIpv4addrGet**](RecordHostIpv4addrAPI.md#RecordhostIpv4addrGet) | **Get** /record:host_ipv4addr | Retrieve record:host_ipv4addr objects
[**RecordhostIpv4addrReferenceGet**](RecordHostIpv4addrAPI.md#RecordhostIpv4addrReferenceGet) | **Get** /record:host_ipv4addr/{reference} | Get a specific record:host_ipv4addr object
[**RecordhostIpv4addrReferencePut**](RecordHostIpv4addrAPI.md#RecordhostIpv4addrReferencePut) | **Put** /record:host_ipv4addr/{reference} | Update a record:host_ipv4addr object



## RecordhostIpv4addrGet

> ListRecordHostIpv4addrResponse RecordhostIpv4addrGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:host_ipv4addr objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordHostIpv4addrAPI.RecordhostIpv4addrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostIpv4addrAPI.RecordhostIpv4addrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordhostIpv4addrGet`: ListRecordHostIpv4addrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostIpv4addrAPI.RecordhostIpv4addrGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordHostIpv4addrAPIRecordhostIpv4addrGetRequest` struct via the builder pattern


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

[**ListRecordHostIpv4addrResponse**](ListRecordHostIpv4addrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordhostIpv4addrReferenceGet

> GetRecordHostIpv4addrResponse RecordhostIpv4addrReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:host_ipv4addr object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the record:host_ipv4addr object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordHostIpv4addrAPI.RecordhostIpv4addrReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostIpv4addrAPI.RecordhostIpv4addrReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordhostIpv4addrReferenceGet`: GetRecordHostIpv4addrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostIpv4addrAPI.RecordhostIpv4addrReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:host_ipv4addr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordHostIpv4addrAPIRecordhostIpv4addrReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordHostIpv4addrResponse**](GetRecordHostIpv4addrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordhostIpv4addrReferencePut

> UpdateRecordHostIpv4addrResponse RecordhostIpv4addrReferencePut(ctx, reference).RecordHostIpv4addr(recordHostIpv4addr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:host_ipv4addr object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dns"
)

func main() {
	reference := "reference_example" // string | Reference of the record:host_ipv4addr object
	recordHostIpv4addr := *dns.NewRecordHostIpv4addr() // RecordHostIpv4addr | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordHostIpv4addrAPI.RecordhostIpv4addrReferencePut(context.Background(), reference).RecordHostIpv4addr(recordHostIpv4addr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostIpv4addrAPI.RecordhostIpv4addrReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordhostIpv4addrReferencePut`: UpdateRecordHostIpv4addrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostIpv4addrAPI.RecordhostIpv4addrReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:host_ipv4addr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordHostIpv4addrAPIRecordhostIpv4addrReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordHostIpv4addr** | [**RecordHostIpv4addr**](RecordHostIpv4addr.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordHostIpv4addrResponse**](UpdateRecordHostIpv4addrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

