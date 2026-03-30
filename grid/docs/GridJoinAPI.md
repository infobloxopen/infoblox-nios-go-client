# GridJoinAPI

All URIs are relative to *http://localhost/wapi/v2.14*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](GridJoinAPI.md#Create) | **Post** /grid/join | Join an Infoblox appliance to an existing grid



## Create

> CreateGridJoinResponse Create(ctx).GridJoin(gridJoin).ReturnAsObject(returnAsObject).Execute()

Join an Infoblox appliance to an existing grid



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/infobloxopen/infoblox-nios-go-client/grid"
)

func main() {
	gridJoin := *grid.NewGridJoin() // GridJoin | Object data to create

	apiClient := grid.NewAPIClient()
	resp, r, err := apiClient.GridJoinAPI.Create(context.Background()).GridJoin(gridJoin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GridJoinAPI.Create``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Create`: CreateGridJoinResponse
	fmt.Fprintf(os.Stdout, "Response from `GridJoinAPI.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a `GridJoinAPICreateRequest` struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**gridJoin** | [**GridJoin**](GridJoin.md) | Object data to create | 
**returnAsObject** | **int32** | Select 1 if result is required as an object | 

### Return type

[**CreateGridJoinResponse**](CreateGridJoinResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

