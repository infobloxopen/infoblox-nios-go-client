# Awsrte53taskgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](Awsrte53taskgroupAPI.md#Get) | **Get** /awsrte53taskgroup | Retrieve awsrte53taskgroup objects
[**Post**](Awsrte53taskgroupAPI.md#Post) | **Post** /awsrte53taskgroup | Create a awsrte53taskgroup object
[**ReferenceDelete**](Awsrte53taskgroupAPI.md#ReferenceDelete) | **Delete** /awsrte53taskgroup/{reference} | Delete a awsrte53taskgroup object
[**ReferenceGet**](Awsrte53taskgroupAPI.md#ReferenceGet) | **Get** /awsrte53taskgroup/{reference} | Get a specific awsrte53taskgroup object
[**ReferencePut**](Awsrte53taskgroupAPI.md#ReferencePut) | **Put** /awsrte53taskgroup/{reference} | Update a awsrte53taskgroup object



## Get

> ListAwsrte53taskgroupResponse Get(ctx).ReturnFields(returnFields).ReturnFields2(returnFields2).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).Execute()

Retrieve awsrte53taskgroup objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.Awsrte53taskgroupAPI.Get(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.Get``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Get`: ListAwsrte53taskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `Awsrte53taskgroupAPI.Get`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPIGetRequest` struct via the builder pattern


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

[**ListAwsrte53taskgroupResponse**](ListAwsrte53taskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> CreateAwsrte53taskgroupResponse Post(ctx).Awsrte53taskgroup(awsrte53taskgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Create a awsrte53taskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {
	awsrte53taskgroup := *cloud.NewAwsrte53taskgroup() // Awsrte53taskgroup | Object data to create

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.Awsrte53taskgroupAPI.Post(context.Background()).Awsrte53taskgroup(awsrte53taskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.Post``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Post`: CreateAwsrte53taskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `Awsrte53taskgroupAPI.Post`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPIPostRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**awsrte53taskgroup** | [**Awsrte53taskgroup**](Awsrte53taskgroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateAwsrte53taskgroupResponse**](CreateAwsrte53taskgroupResponse.md)

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

Delete a awsrte53taskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the awsrte53taskgroup object

	apiClient := cloud.NewAPIClient()
	r, err := apiClient.Awsrte53taskgroupAPI.ReferenceDelete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.ReferenceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the awsrte53taskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPIReferenceDeleteRequest` struct via the builder pattern


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

> GetAwsrte53taskgroupResponse ReferenceGet(ctx, reference).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Get a specific awsrte53taskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the awsrte53taskgroup object

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.Awsrte53taskgroupAPI.ReferenceGet(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.ReferenceGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferenceGet`: GetAwsrte53taskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `Awsrte53taskgroupAPI.ReferenceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the awsrte53taskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPIReferenceGetRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**GetAwsrte53taskgroupResponse**](GetAwsrte53taskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReferencePut

> UpdateAwsrte53taskgroupResponse ReferencePut(ctx, reference).Awsrte53taskgroup(awsrte53taskgroup).ReturnFields(returnFields).ReturnFields2(returnFields2).ReturnAsObject(returnAsObject).Execute()

Update a awsrte53taskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Infoblox-CTO/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the awsrte53taskgroup object
	awsrte53taskgroup := *cloud.NewAwsrte53taskgroup() // Awsrte53taskgroup | Object data to update

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.Awsrte53taskgroupAPI.ReferencePut(context.Background(), reference).Awsrte53taskgroup(awsrte53taskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.ReferencePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReferencePut`: UpdateAwsrte53taskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `Awsrte53taskgroupAPI.ReferencePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the awsrte53taskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPIReferencePutRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**awsrte53taskgroup** | [**Awsrte53taskgroup**](Awsrte53taskgroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFields2** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateAwsrte53taskgroupResponse**](UpdateAwsrte53taskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

