# DtcTopologyRuleAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtctopologyruleGet**](DtcTopologyRuleAPI.md#DtctopologyruleGet) | **Get** /dtc:topology:rule | Retrieve dtc:topology:rule objects
[**DtctopologyruleReferenceGet**](DtcTopologyRuleAPI.md#DtctopologyruleReferenceGet) | **Get** /dtc:topology:rule/{reference} | Get a specific dtc:topology:rule object
[**DtctopologyruleReferencePut**](DtcTopologyRuleAPI.md#DtctopologyruleReferencePut) | **Put** /dtc:topology:rule/{reference} | Update a dtc:topology:rule object



## DtctopologyruleGet

> ListDtcTopologyRuleResponse DtctopologyruleGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:topology:rule objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyRuleAPI.DtctopologyruleGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyRuleAPI.DtctopologyruleGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtctopologyruleGet`: ListDtcTopologyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyRuleAPI.DtctopologyruleGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyRuleAPIDtctopologyruleGetRequest` struct via the builder pattern


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

[**ListDtcTopologyRuleResponse**](ListDtcTopologyRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtctopologyruleReferenceGet

> GetDtcTopologyRuleResponse DtctopologyruleReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:topology:rule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:topology:rule object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyRuleAPI.DtctopologyruleReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyRuleAPI.DtctopologyruleReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtctopologyruleReferenceGet`: GetDtcTopologyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyRuleAPI.DtctopologyruleReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:topology:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyRuleAPIDtctopologyruleReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcTopologyRuleResponse**](GetDtcTopologyRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtctopologyruleReferencePut

> UpdateDtcTopologyRuleResponse DtctopologyruleReferencePut(ctx, reference).DtcTopologyRule(dtcTopologyRule).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:topology:rule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:topology:rule object
	dtcTopologyRule := *dtc.NewDtcTopologyRule() // DtcTopologyRule | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyRuleAPI.DtctopologyruleReferencePut(context.Background(), reference).DtcTopologyRule(dtcTopologyRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyRuleAPI.DtctopologyruleReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtctopologyruleReferencePut`: UpdateDtcTopologyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyRuleAPI.DtctopologyruleReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:topology:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyRuleAPIDtctopologyruleReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcTopologyRule** | [**DtcTopologyRule**](DtcTopologyRule.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcTopologyRuleResponse**](UpdateDtcTopologyRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

