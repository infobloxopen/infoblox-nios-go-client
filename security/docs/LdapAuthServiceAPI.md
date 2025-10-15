# LdapAuthServiceAPI

All URIs are relative to *http://localhost/wapi/v2.13.6*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](LdapAuthServiceAPI.md#Create) | **Post** /ldap_auth_service | Create a ldap_auth_service object
[**Delete**](LdapAuthServiceAPI.md#Delete) | **Delete** /ldap_auth_service/{reference} | Delete a ldap_auth_service object
[**List**](LdapAuthServiceAPI.md#List) | **Get** /ldap_auth_service | Retrieve ldap_auth_service objects
[**Read**](LdapAuthServiceAPI.md#Read) | **Get** /ldap_auth_service/{reference} | Get a specific ldap_auth_service object
[**Update**](LdapAuthServiceAPI.md#Update) | **Put** /ldap_auth_service/{reference} | Update a ldap_auth_service object



## Create

> CreateLdapAuthServiceResponse Create(ctx).LdapAuthService(ldapAuthService).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Create a ldap_auth_service object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	ldapAuthService := *security.NewLdapAuthService() // LdapAuthService | Object data to create

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.LdapAuthServiceAPI.Create(context.Background()).LdapAuthService(ldapAuthService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapAuthServiceAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateLdapAuthServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `LdapAuthServiceAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `LdapAuthServiceAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ldapAuthService** | [**LdapAuthService**](LdapAuthService.md) | Object data to create | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateLdapAuthServiceResponse**](CreateLdapAuthServiceResponse.md)

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

Delete a ldap_auth_service object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the ldap_auth_service object

	apiClient := security.NewAPIClient()
	r, err := apiClient.LdapAuthServiceAPI.Delete(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapAuthServiceAPI.Delete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ldap_auth_service object | 

### Other Parameters

Other parameters are passed through a pointer to a `LdapAuthServiceAPIDeleteRequest` struct via the builder pattern


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

> ListLdapAuthServiceResponse List(ctx).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).MaxResults(maxResults).ReturnAsObject(returnAsObject).Paging(paging).PageId(pageId).Filters(filters).Extattrfilter(extattrfilter).ProxySearch(proxySearch).Execute()

Retrieve ldap_auth_service objects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.LdapAuthServiceAPI.List(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapAuthServiceAPI.List``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `List`: ListLdapAuthServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `LdapAuthServiceAPI.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `LdapAuthServiceAPIListRequest` struct via the builder pattern


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

[**ListLdapAuthServiceResponse**](ListLdapAuthServiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Read

> GetLdapAuthServiceResponse Read(ctx, reference).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).ProxySearch(proxySearch).Execute()

Get a specific ldap_auth_service object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the ldap_auth_service object

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.LdapAuthServiceAPI.Read(context.Background(), reference).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapAuthServiceAPI.Read``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Read`: GetLdapAuthServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `LdapAuthServiceAPI.Read`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ldap_auth_service object | 

### Other Parameters

Other parameters are passed through a pointer to a `LdapAuthServiceAPIReadRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 
**proxySearch** | **string** | Search Grid members for data | 

### Return type

[**GetLdapAuthServiceResponse**](GetLdapAuthServiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> UpdateLdapAuthServiceResponse Update(ctx, reference).LdapAuthService(ldapAuthService).ReturnFields(returnFields).ReturnFieldsPlus(returnFieldsPlus).ReturnAsObject(returnAsObject).Execute()

Update a ldap_auth_service object



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/security"
)

func main() {
	reference := "reference_example" // string | Reference of the ldap_auth_service object
	ldapAuthService := *security.NewLdapAuthService() // LdapAuthService | Object data to update

	apiClient := security.NewAPIClient()
	resp, r, err := apiClient.LdapAuthServiceAPI.Update(context.Background(), reference).LdapAuthService(ldapAuthService).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LdapAuthServiceAPI.Update``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Update`: UpdateLdapAuthServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `LdapAuthServiceAPI.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reference** | **string** | Reference of the ldap_auth_service object | 

### Other Parameters

Other parameters are passed through a pointer to a `LdapAuthServiceAPIUpdateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ldapAuthService** | [**LdapAuthService**](LdapAuthService.md) | Object data to update | 
**returnFields** | **string** | Enter the field names followed by comma | 
**returnFieldsPlus** | **string** | Enter the field names followed by comma, this returns the required fields along with the default fields | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**UpdateLdapAuthServiceResponse**](UpdateLdapAuthServiceResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

