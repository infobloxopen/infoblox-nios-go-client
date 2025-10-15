# RecordRpzCnameClientipaddressdnAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](RecordRpzCnameClientipaddressdnAPI.md#Create) | **Post** /record:rpz:cname:clientipaddressdn | Create a record:rpz:cname:clientipaddressdn object
[**Delete**](RecordRpzCnameClientipaddressdnAPI.md#Delete) | **Delete** /record:rpz:cname:clientipaddressdn/{reference} | Delete a record:rpz:cname:clientipaddressdn object
[**List**](RecordRpzCnameClientipaddressdnAPI.md#List) | **Get** /record:rpz:cname:clientipaddressdn | Retrieve record:rpz:cname:clientipaddressdn objects
[**Read**](RecordRpzCnameClientipaddressdnAPI.md#Read) | **Get** /record:rpz:cname:clientipaddressdn/{reference} | Get a specific record:rpz:cname:clientipaddressdn object
[**Update**](RecordRpzCnameClientipaddressdnAPI.md#Update) | **Put** /record:rpz:cname:clientipaddressdn/{reference} | Update a record:rpz:cname:clientipaddressdn object



## Create

> CreateRecordRpzCnameClientipaddressdnResponse Create(ctx).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:cname:clientipaddressdn object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
)

func main() {
	recordRpzCnameClientipaddressdn := *rpz.NewRecordRpzCnameClientipaddressdn() // RecordRpzCnameClientipaddressdn | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.Create(context.Background()).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressdnAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameClientipaddressdn** | [**RecordRpzCnameClientipaddressdn**](RecordRpzCnameClientipaddressdn.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzCnameClientipaddressdnResponse**](CreateRecordRpzCnameClientipaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, reference).Execute()

Delete a record:rpz:cname:clientipaddressdn object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
)

func main() {
	reference := "reference_example" // string | Reference of the record:rpz:cname:clientipaddressdn object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPIDeleteRequest` struct via the builder pattern


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


## List

> ListRecordRpzCnameClientipaddressdnResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve record:rpz:cname:clientipaddressdn objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
)

func main() {

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressdnAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPIListRequest` struct via the builder pattern


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
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**ListRecordRpzCnameClientipaddressdnResponse**](ListRecordRpzCnameClientipaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetRecordRpzCnameClientipaddressdnResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific record:rpz:cname:clientipaddressdn object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
)

func main() {
	reference := "reference_example" // string | Reference of the record:rpz:cname:clientipaddressdn object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressdnAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetRecordRpzCnameClientipaddressdnResponse**](GetRecordRpzCnameClientipaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateRecordRpzCnameClientipaddressdnResponse Update(ctx, reference).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:cname:clientipaddressdn object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
)

func main() {
	reference := "reference_example" // string | Reference of the record:rpz:cname:clientipaddressdn object
	recordRpzCnameClientipaddressdn := *rpz.NewRecordRpzCnameClientipaddressdn() // RecordRpzCnameClientipaddressdn | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzCnameClientipaddressdnAPI.Update(context.Background(), reference).RecordRpzCnameClientipaddressdn(recordRpzCnameClientipaddressdn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzCnameClientipaddressdnAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateRecordRpzCnameClientipaddressdnResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzCnameClientipaddressdnAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:clientipaddressdn object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzCnameClientipaddressdnAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameClientipaddressdn** | [**RecordRpzCnameClientipaddressdn**](RecordRpzCnameClientipaddressdn.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzCnameClientipaddressdnResponse**](UpdateRecordRpzCnameClientipaddressdnResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

