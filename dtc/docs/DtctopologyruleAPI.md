# DtcTopologyRuleAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](DtcTopologyRuleAPI.md#Create) | **Post** /dtc:topology:rule | Create a dtc:topology:rule object
[**List**](DtcTopologyRuleAPI.md#List) | **Get** /dtc:topology:rule | Retrieve dtc:topology:rule objects
[**Read**](DtcTopologyRuleAPI.md#Read) | **Get** /dtc:topology:rule/{reference} | Get a specific dtc:topology:rule object
[**Update**](DtcTopologyRuleAPI.md#Update) | **Put** /dtc:topology:rule/{reference} | Update a dtc:topology:rule object



## Create

> CreateDtcTopologyRuleResponse Create(ctx).DtcTopologyRule(dtcTopologyRule).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a dtc:topology:rule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
)

func main() {
	dtcTopologyRule := *dtc.NewDtcTopologyRule() // DtcTopologyRule | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyRuleAPI.Create(context.Background()).DtcTopologyRule(dtcTopologyRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyRuleAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateDtcTopologyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyRuleAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyRuleAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcTopologyRule** | [**DtcTopologyRule**](DtcTopologyRule.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcTopologyRuleResponse**](CreateDtcTopologyRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListDtcTopologyRuleResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:topology:rule objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
)

func main() {

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyRuleAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyRuleAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListDtcTopologyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyRuleAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyRuleAPIListRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## Read

> GetDtcTopologyRuleResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:topology:rule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:topology:rule object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyRuleAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyRuleAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetDtcTopologyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyRuleAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:topology:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyRuleAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## Update

> UpdateDtcTopologyRuleResponse Update(ctx, reference).DtcTopologyRule(dtcTopologyRule).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a dtc:topology:rule object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
)

func main() {
	reference := "reference_example" // string | Reference of the dtc:topology:rule object
	dtcTopologyRule := *dtc.NewDtcTopologyRule() // DtcTopologyRule | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcTopologyRuleAPI.Update(context.Background(), reference).DtcTopologyRule(dtcTopologyRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcTopologyRuleAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateDtcTopologyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcTopologyRuleAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:topology:rule object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcTopologyRuleAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcTopologyRule** | [**DtcTopologyRule**](DtcTopologyRule.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

