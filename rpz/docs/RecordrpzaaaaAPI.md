# RecordRpzAaaaAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordrpzaaaaGet**](RecordRpzAaaaAPI.md#RecordrpzaaaaGet) | **Get** /record:rpz:aaaa | Retrieve record:rpz:aaaa objects
[**RecordrpzaaaaPost**](RecordRpzAaaaAPI.md#RecordrpzaaaaPost) | **Post** /record:rpz:aaaa | Create a record:rpz:aaaa object
[**RecordrpzaaaaReferenceDelete**](RecordRpzAaaaAPI.md#RecordrpzaaaaReferenceDelete) | **Delete** /record:rpz:aaaa/{reference} | Delete a record:rpz:aaaa object
[**RecordrpzaaaaReferenceGet**](RecordRpzAaaaAPI.md#RecordrpzaaaaReferenceGet) | **Get** /record:rpz:aaaa/{reference} | Get a specific record:rpz:aaaa object
[**RecordrpzaaaaReferencePut**](RecordRpzAaaaAPI.md#RecordrpzaaaaReferencePut) | **Put** /record:rpz:aaaa/{reference} | Update a record:rpz:aaaa object



## RecordrpzaaaaGet

> ListRecordRpzAaaaResponse RecordrpzaaaaGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:rpz:aaaa objects



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
	resp, r, err := apiClient.RecordRpzAaaaAPI.RecordrpzaaaaGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaAPI.RecordrpzaaaaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaaaaGet`: ListRecordRpzAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAaaaAPI.RecordrpzaaaaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaAPIRecordrpzaaaaGetRequest` struct via the builder pattern


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

[**ListRecordRpzAaaaResponse**](ListRecordRpzAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzaaaaPost

> CreateRecordRpzAaaaResponse RecordrpzaaaaPost(ctx).RecordRpzAaaa(recordRpzAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:rpz:aaaa object



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
	recordRpzAaaa := *rpz.NewRecordRpzAaaa() // RecordRpzAaaa | Object data to create

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzAaaaAPI.RecordrpzaaaaPost(context.Background()).RecordRpzAaaa(recordRpzAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaAPI.RecordrpzaaaaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaaaaPost`: CreateRecordRpzAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAaaaAPI.RecordrpzaaaaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaAPIRecordrpzaaaaPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzAaaa** | [**RecordRpzAaaa**](RecordRpzAaaa.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordRpzAaaaResponse**](CreateRecordRpzAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzaaaaReferenceDelete

> RecordrpzaaaaReferenceDelete(ctx, reference).Execute()

Delete a record:rpz:aaaa object



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
	reference := "reference_example" // string | Reference of the record:rpz:aaaa object

	apiClient := rpz.NewAPIClient()
	r, err := apiClient.RecordRpzAaaaAPI.RecordrpzaaaaReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaAPI.RecordrpzaaaaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaAPIRecordrpzaaaaReferenceDeleteRequest` struct via the builder pattern


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


## RecordrpzaaaaReferenceGet

> GetRecordRpzAaaaResponse RecordrpzaaaaReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:rpz:aaaa object



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
	reference := "reference_example" // string | Reference of the record:rpz:aaaa object

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzAaaaAPI.RecordrpzaaaaReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaAPI.RecordrpzaaaaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaaaaReferenceGet`: GetRecordRpzAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAaaaAPI.RecordrpzaaaaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaAPIRecordrpzaaaaReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordRpzAaaaResponse**](GetRecordRpzAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordrpzaaaaReferencePut

> UpdateRecordRpzAaaaResponse RecordrpzaaaaReferencePut(ctx, reference).RecordRpzAaaa(recordRpzAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:rpz:aaaa object



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
	reference := "reference_example" // string | Reference of the record:rpz:aaaa object
	recordRpzAaaa := *rpz.NewRecordRpzAaaa() // RecordRpzAaaa | Object data to update

	apiClient := rpz.NewAPIClient()
	resp, r, err := apiClient.RecordRpzAaaaAPI.RecordrpzaaaaReferencePut(context.Background(), reference).RecordRpzAaaa(recordRpzAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordRpzAaaaAPI.RecordrpzaaaaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordrpzaaaaReferencePut`: UpdateRecordRpzAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordRpzAaaaAPI.RecordrpzaaaaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:rpz:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordRpzAaaaAPIRecordrpzaaaaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordRpzAaaa** | [**RecordRpzAaaa**](RecordRpzAaaa.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordRpzAaaaResponse**](UpdateRecordRpzAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

