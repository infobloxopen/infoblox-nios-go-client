# Awsrte53taskgroupAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](Awsrte53taskgroupAPI.md#Create) | **Post** /awsrte53taskgroup | Create a awsrte53taskgroup object
[**Delete**](Awsrte53taskgroupAPI.md#Delete) | **Delete** /awsrte53taskgroup/{reference} | Delete a awsrte53taskgroup object
[**List**](Awsrte53taskgroupAPI.md#List) | **Get** /awsrte53taskgroup | Retrieve awsrte53taskgroup objects
[**Read**](Awsrte53taskgroupAPI.md#Read) | **Get** /awsrte53taskgroup/{reference} | Get a specific awsrte53taskgroup object
[**Update**](Awsrte53taskgroupAPI.md#Update) | **Put** /awsrte53taskgroup/{reference} | Update a awsrte53taskgroup object



## Create

> CreateAwsrte53taskgroupResponse Create(ctx).Awsrte53taskgroup(awsrte53taskgroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a awsrte53taskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	awsrte53taskgroup := *cloud.NewAwsrte53taskgroup() // Awsrte53taskgroup | Object data to create

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.Awsrte53taskgroupAPI.Create(context.Background()).Awsrte53taskgroup(awsrte53taskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateAwsrte53taskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `Awsrte53taskgroupAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**awsrte53taskgroup** | [**Awsrte53taskgroup**](Awsrte53taskgroup.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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


## Delete

> Delete(ctx, reference).Execute()

Delete a awsrte53taskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the awsrte53taskgroup object

	apiClient := cloud.NewAPIClient()
	r, err := apiClient.Awsrte53taskgroupAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPIDeleteRequest` struct via the builder pattern


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

> ListAwsrte53taskgroupResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve awsrte53taskgroup objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.Awsrte53taskgroupAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListAwsrte53taskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `Awsrte53taskgroupAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPIListRequest` struct via the builder pattern


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

[**ListAwsrte53taskgroupResponse**](ListAwsrte53taskgroupResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetAwsrte53taskgroupResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific awsrte53taskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the awsrte53taskgroup object

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.Awsrte53taskgroupAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetAwsrte53taskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `Awsrte53taskgroupAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the awsrte53taskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

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


## Update

> UpdateAwsrte53taskgroupResponse Update(ctx, reference).Awsrte53taskgroup(awsrte53taskgroup).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a awsrte53taskgroup object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
)

func main() {
	reference := "reference_example" // string | Reference of the awsrte53taskgroup object
	awsrte53taskgroup := *cloud.NewAwsrte53taskgroup() // Awsrte53taskgroup | Object data to update

	apiClient := cloud.NewAPIClient()
	resp, r, err := apiClient.Awsrte53taskgroupAPI.Update(context.Background(), reference).Awsrte53taskgroup(awsrte53taskgroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Awsrte53taskgroupAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateAwsrte53taskgroupResponse
	fmt.Fprintf(os.Stdout, "Response from `Awsrte53taskgroupAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the awsrte53taskgroup object | 

### Other Parameters

Other parameters are passed through a pointer to a `Awsrte53taskgroupAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**awsrte53taskgroup** | [**Awsrte53taskgroup**](Awsrte53taskgroup.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
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

