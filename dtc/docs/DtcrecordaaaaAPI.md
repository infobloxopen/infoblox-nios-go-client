# DtcRecordAaaaAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DtcrecordaaaaGet**](DtcRecordAaaaAPI.md#DtcrecordaaaaGet) | **Get** /dtc:record:aaaa | Retrieve dtc:record:aaaa objects
[**DtcrecordaaaaPost**](DtcRecordAaaaAPI.md#DtcrecordaaaaPost) | **Post** /dtc:record:aaaa | Create a dtc:record:aaaa object
[**DtcrecordaaaaReferenceDelete**](DtcRecordAaaaAPI.md#DtcrecordaaaaReferenceDelete) | **Delete** /dtc:record:aaaa/{reference} | Delete a dtc:record:aaaa object
[**DtcrecordaaaaReferenceGet**](DtcRecordAaaaAPI.md#DtcrecordaaaaReferenceGet) | **Get** /dtc:record:aaaa/{reference} | Get a specific dtc:record:aaaa object
[**DtcrecordaaaaReferencePut**](DtcRecordAaaaAPI.md#DtcrecordaaaaReferencePut) | **Put** /dtc:record:aaaa/{reference} | Update a dtc:record:aaaa object



## DtcrecordaaaaGet

> ListDtcRecordAaaaResponse DtcrecordaaaaGet(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve dtc:record:aaaa objects



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
	resp, r, err := apiClient.DtcRecordAaaaAPI.DtcrecordaaaaGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAaaaAPI.DtcrecordaaaaGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordaaaaGet`: ListDtcRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordAaaaAPI.DtcrecordaaaaGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAaaaAPIDtcrecordaaaaGetRequest` struct via the builder pattern


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

[**ListDtcRecordAaaaResponse**](ListDtcRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordaaaaPost

> CreateDtcRecordAaaaResponse DtcrecordaaaaPost(ctx).DtcRecordAaaa(dtcRecordAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a dtc:record:aaaa object



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
	dtcRecordAaaa := *dtc.NewDtcRecordAaaa() // DtcRecordAaaa | Object data to create

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordAaaaAPI.DtcrecordaaaaPost(context.Background()).DtcRecordAaaa(dtcRecordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAaaaAPI.DtcrecordaaaaPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordaaaaPost`: CreateDtcRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordAaaaAPI.DtcrecordaaaaPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAaaaAPIDtcrecordaaaaPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordAaaa** | [**DtcRecordAaaa**](DtcRecordAaaa.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateDtcRecordAaaaResponse**](CreateDtcRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordaaaaReferenceDelete

> DtcrecordaaaaReferenceDelete(ctx, reference).Execute()

Delete a dtc:record:aaaa object



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
	reference := "reference_example" // string | Reference of the dtc:record:aaaa object

	apiClient := dtc.NewAPIClient()
	r, err := apiClient.DtcRecordAaaaAPI.DtcrecordaaaaReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAaaaAPI.DtcrecordaaaaReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAaaaAPIDtcrecordaaaaReferenceDeleteRequest` struct via the builder pattern


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


## DtcrecordaaaaReferenceGet

> GetDtcRecordAaaaResponse DtcrecordaaaaReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific dtc:record:aaaa object



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
	reference := "reference_example" // string | Reference of the dtc:record:aaaa object

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordAaaaAPI.DtcrecordaaaaReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAaaaAPI.DtcrecordaaaaReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordaaaaReferenceGet`: GetDtcRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordAaaaAPI.DtcrecordaaaaReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAaaaAPIDtcrecordaaaaReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetDtcRecordAaaaResponse**](GetDtcRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DtcrecordaaaaReferencePut

> UpdateDtcRecordAaaaResponse DtcrecordaaaaReferencePut(ctx, reference).DtcRecordAaaa(dtcRecordAaaa).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a dtc:record:aaaa object



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
	reference := "reference_example" // string | Reference of the dtc:record:aaaa object
	dtcRecordAaaa := *dtc.NewDtcRecordAaaa() // DtcRecordAaaa | Object data to update

	apiClient := dtc.NewAPIClient()
	resp, r, err := apiClient.DtcRecordAaaaAPI.DtcrecordaaaaReferencePut(context.Background(), reference).DtcRecordAaaa(dtcRecordAaaa).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DtcRecordAaaaAPI.DtcrecordaaaaReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DtcrecordaaaaReferencePut`: UpdateDtcRecordAaaaResponse
	fmt.Fprintf(os.Stdout, "Response from `DtcRecordAaaaAPI.DtcrecordaaaaReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the dtc:record:aaaa object | 

### Other Parameters

Other parameters are passed through a pointer to a `DtcRecordAaaaAPIDtcrecordaaaaReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**dtcRecordAaaa** | [**DtcRecordAaaa**](DtcRecordAaaa.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateDtcRecordAaaaResponse**](UpdateDtcRecordAaaaResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

