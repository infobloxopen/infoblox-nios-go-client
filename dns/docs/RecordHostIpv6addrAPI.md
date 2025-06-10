# RecordHostIpv6addrAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordhostIpv6addrGet**](RecordHostIpv6addrAPI.md#RecordhostIpv6addrGet) | **Get** /record:host_ipv6addr | Retrieve record:host_ipv6addr objects
[**RecordhostIpv6addrReferenceGet**](RecordHostIpv6addrAPI.md#RecordhostIpv6addrReferenceGet) | **Get** /record:host_ipv6addr/{reference} | Get a specific record:host_ipv6addr object
[**RecordhostIpv6addrReferencePut**](RecordHostIpv6addrAPI.md#RecordhostIpv6addrReferencePut) | **Put** /record:host_ipv6addr/{reference} | Update a record:host_ipv6addr object



## RecordhostIpv6addrGet

> ListRecordHostIpv6addrResponse RecordhostIpv6addrGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:host_ipv6addr objects



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
	resp, r, err := apiClient.RecordHostIpv6addrAPI.RecordhostIpv6addrGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostIpv6addrAPI.RecordhostIpv6addrGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordhostIpv6addrGet`: ListRecordHostIpv6addrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostIpv6addrAPI.RecordhostIpv6addrGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordHostIpv6addrAPIRecordhostIpv6addrGetRequest` struct via the builder pattern


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

[**ListRecordHostIpv6addrResponse**](ListRecordHostIpv6addrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordhostIpv6addrReferenceGet

> GetRecordHostIpv6addrResponse RecordhostIpv6addrReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:host_ipv6addr object



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
	reference := "reference_example" // string | Reference of the record:host_ipv6addr object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordHostIpv6addrAPI.RecordhostIpv6addrReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostIpv6addrAPI.RecordhostIpv6addrReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordhostIpv6addrReferenceGet`: GetRecordHostIpv6addrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostIpv6addrAPI.RecordhostIpv6addrReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:host_ipv6addr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordHostIpv6addrAPIRecordhostIpv6addrReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordHostIpv6addrResponse**](GetRecordHostIpv6addrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordhostIpv6addrReferencePut

> UpdateRecordHostIpv6addrResponse RecordhostIpv6addrReferencePut(ctx, reference).RecordHostIpv6addr(recordHostIpv6addr).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:host_ipv6addr object



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
	reference := "reference_example" // string | Reference of the record:host_ipv6addr object
	recordHostIpv6addr := *dns.NewRecordHostIpv6addr() // RecordHostIpv6addr | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordHostIpv6addrAPI.RecordhostIpv6addrReferencePut(context.Background(), reference).RecordHostIpv6addr(recordHostIpv6addr).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordHostIpv6addrAPI.RecordhostIpv6addrReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordhostIpv6addrReferencePut`: UpdateRecordHostIpv6addrResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordHostIpv6addrAPI.RecordhostIpv6addrReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:host_ipv6addr object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordHostIpv6addrAPIRecordhostIpv6addrReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordHostIpv6addr** | [**RecordHostIpv6addr**](RecordHostIpv6addr.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordHostIpv6addrResponse**](UpdateRecordHostIpv6addrResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

