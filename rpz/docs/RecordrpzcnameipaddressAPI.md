# RecordrpzcnameipaddressAPI

All URIs are relative to *http://localhost/wapi/v2.12.3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](RecordrpzcnameipaddressAPI.md#Get) | **Get** /record:rpz:cname:ipaddress | Retrieve record:rpz:cname:ipaddress objects
[**Post**](RecordrpzcnameipaddressAPI.md#Post) | **Post** /record:rpz:cname:ipaddress | Create a record:rpz:cname:ipaddress object
[**ReferenceDelete**](RecordrpzcnameipaddressAPI.md#ReferenceDelete) | **Delete** /record:rpz:cname:ipaddress/{reference} | Delete a record:rpz:cname:ipaddress object
[**ReferenceGet**](RecordrpzcnameipaddressAPI.md#ReferenceGet) | **Get** /record:rpz:cname:ipaddress/{reference} | Get a specific record:rpz:cname:ipaddress object
[**ReferencePut**](RecordrpzcnameipaddressAPI.md#ReferencePut) | **Put** /record:rpz:cname:ipaddress/{reference} | Update a record:rpz:cname:ipaddress object



## Get

> ListRecordRpzCnameIpaddressResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:cname:ipaddress objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rpz"
)

func main() {

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordrpzcnameipaddressAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameipaddressAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListRecordRpzCnameIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordrpzcnameipaddressAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordrpzcnameipaddressAPIGetRequest` struct via the builder pattern


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

[**ListRecordRpzCnameIpaddressResponse**](ListRecordRpzCnameIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateRecordRpzCnameIpaddressResponse Post(ctx).RecordRpzCnameIpaddress(recordRpzCnameIpaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:cname:ipaddress object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rpz"
)

func main() {
	recordRpzCnameIpaddress := *rpz.NewRecordRpzCnameIpaddress() // RecordRpzCnameIpaddress | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordrpzcnameipaddressAPI.Post(context.Background()).RecordRpzCnameIpaddress(recordRpzCnameIpaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameipaddressAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateRecordRpzCnameIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordrpzcnameipaddressAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordrpzcnameipaddressAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameIpaddress** | [**RecordRpzCnameIpaddress**](RecordRpzCnameIpaddress.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzCnameIpaddressResponse**](CreateRecordRpzCnameIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferenceDelete

> ReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:cname:ipaddress object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rpz"
)

func main() {
	reference := "reference_example" // string | Reference of the record:rpz:cname:ipaddress object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordrpzcnameipaddressAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameipaddressAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordrpzcnameipaddressAPIReferenceDeleteRequest` struct via the builder pattern


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


## ReferenceGet

> GetRecordRpzCnameIpaddressResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:cname:ipaddress object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rpz"
)

func main() {
	reference := "reference_example" // string | Reference of the record:rpz:cname:ipaddress object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordrpzcnameipaddressAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameipaddressAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetRecordRpzCnameIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordrpzcnameipaddressAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordrpzcnameipaddressAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzCnameIpaddressResponse**](GetRecordRpzCnameIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateRecordRpzCnameIpaddressResponse ReferencePut(ctx, reference).RecordRpzCnameIpaddress(recordRpzCnameIpaddress).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:cname:ipaddress object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/rpz"
)

func main() {
	reference := "reference_example" // string | Reference of the record:rpz:cname:ipaddress object
	recordRpzCnameIpaddress := *rpz.NewRecordRpzCnameIpaddress() // RecordRpzCnameIpaddress | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordrpzcnameipaddressAPI.ReferencePut(context.Background(), reference).RecordRpzCnameIpaddress(recordRpzCnameIpaddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordrpzcnameipaddressAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateRecordRpzCnameIpaddressResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordrpzcnameipaddressAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:cname:ipaddress object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordrpzcnameipaddressAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzCnameIpaddress** | [**RecordRpzCnameIpaddress**](RecordRpzCnameIpaddress.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzCnameIpaddressResponse**](UpdateRecordRpzCnameIpaddressResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

