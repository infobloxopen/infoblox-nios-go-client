# RecordAaaaAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RecordaaaaGet**](RecordAaaaAPI.md#RecordaaaaGet) | **Get** /record:aaaa | Retrieve record:aaaa objects
[**RecordaaaaPost**](RecordAaaaAPI.md#RecordaaaaPost) | **Post** /record:aaaa | Create a record:aaaa object
[**RecordaaaaReferenceDelete**](RecordAaaaAPI.md#RecordaaaaReferenceDelete) | **Delete** /record:aaaa/{reference} | Delete a record:aaaa object
[**RecordaaaaReferenceGet**](RecordAaaaAPI.md#RecordaaaaReferenceGet) | **Get** /record:aaaa/{reference} | Get a specific record:aaaa object
[**RecordaaaaReferencePut**](RecordAaaaAPI.md#RecordaaaaReferencePut) | **Put** /record:aaaa/{reference} | Update a record:aaaa object



## RecordaaaaGet

> ListRecordAaaaResponse RecordaaaaGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve record:aaaa objects



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
	resp, r, err := apiClient.RecordAaaaAPI.RecordaaaaGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAaaaAPI.RecordaaaaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaaaaGet`: ListRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAaaaAPI.RecordaaaaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordAaaaAPIRecordaaaaGetRequest` struct via the builder pattern


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

[**ListRecordAaaaResponse**](ListRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordaaaaPost

> CreateRecordAaaaResponse RecordaaaaPost(ctx).RecordAaaa(recordAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a record:aaaa object



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
	recordAaaa := *dns.NewRecordAaaa() // RecordAaaa | Object data to create

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordAaaaAPI.RecordaaaaPost(context.Background()).RecordAaaa(recordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAaaaAPI.RecordaaaaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaaaaPost`: CreateRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAaaaAPI.RecordaaaaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `RecordAaaaAPIRecordaaaaPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordAaaa** | [**RecordAaaa**](RecordAaaa.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateRecordAaaaResponse**](CreateRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordaaaaReferenceDelete

> RecordaaaaReferenceDelete(ctx, reference).RemoveAssociatedPtr(removeAssociatedPtr).Execute()

Delete a record:aaaa object



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
	reference := "reference_example" // string | Reference of the record:aaaa object

	apiClient := dns.NewAPIClient()
	r, err := apiClient.RecordAaaaAPI.RecordaaaaReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAaaaAPI.RecordaaaaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordAaaaAPIRecordaaaaReferenceDeleteRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**removeAssociatedPtr** | **bool** | Delete option that indicates whether the associated PTR records should be removed while deleting the specified A record. | 

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


## RecordaaaaReferenceGet

> GetRecordAaaaResponse RecordaaaaReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific record:aaaa object



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
	reference := "reference_example" // string | Reference of the record:aaaa object

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordAaaaAPI.RecordaaaaReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAaaaAPI.RecordaaaaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaaaaReferenceGet`: GetRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAaaaAPI.RecordaaaaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordAaaaAPIRecordaaaaReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetRecordAaaaResponse**](GetRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecordaaaaReferencePut

> UpdateRecordAaaaResponse RecordaaaaReferencePut(ctx, reference).RecordAaaa(recordAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a record:aaaa object



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
	reference := "reference_example" // string | Reference of the record:aaaa object
	recordAaaa := *dns.NewRecordAaaa() // RecordAaaa | Object data to update

	apiClient := dns.NewAPIClient()
	resp, r, err := apiClient.RecordAaaaAPI.RecordaaaaReferencePut(context.Background(), reference).RecordAaaa(recordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RecordAaaaAPI.RecordaaaaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RecordaaaaReferencePut`: UpdateRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `RecordAaaaAPI.RecordaaaaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the record:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `RecordAaaaAPIRecordaaaaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**recordAaaa** | [**RecordAaaa**](RecordAaaa.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateRecordAaaaResponse**](UpdateRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

