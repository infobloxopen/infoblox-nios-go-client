# GridCloudapiGatewayConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableProxyService** | Pointer to **bool** | Enable Gateway Service. | [optional] 
**Port** | Pointer to **int64** | Gateway port | [optional] 
**EndpointMapping** | Pointer to [**GridcloudapigatewayconfigEndpointMapping**](GridcloudapigatewayconfigEndpointMapping.md) |  | [optional] 

## Methods

### NewGridCloudapiGatewayConfig

`func NewGridCloudapiGatewayConfig() *GridCloudapiGatewayConfig`

NewGridCloudapiGatewayConfig instantiates a new GridCloudapiGatewayConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridCloudapiGatewayConfigWithDefaults

`func NewGridCloudapiGatewayConfigWithDefaults() *GridCloudapiGatewayConfig`

NewGridCloudapiGatewayConfigWithDefaults instantiates a new GridCloudapiGatewayConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableProxyService

`func (o *GridCloudapiGatewayConfig) GetEnableProxyService() bool`

GetEnableProxyService returns the EnableProxyService field if non-nil, zero value otherwise.

### GetEnableProxyServiceOk

`func (o *GridCloudapiGatewayConfig) GetEnableProxyServiceOk() (*bool, bool)`

GetEnableProxyServiceOk returns a tuple with the EnableProxyService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProxyService

`func (o *GridCloudapiGatewayConfig) SetEnableProxyService(v bool)`

SetEnableProxyService sets EnableProxyService field to given value.

### HasEnableProxyService

`func (o *GridCloudapiGatewayConfig) HasEnableProxyService() bool`

HasEnableProxyService returns a boolean if a field has been set.

### GetPort

`func (o *GridCloudapiGatewayConfig) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GridCloudapiGatewayConfig) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GridCloudapiGatewayConfig) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *GridCloudapiGatewayConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetEndpointMapping

`func (o *GridCloudapiGatewayConfig) GetEndpointMapping() GridcloudapigatewayconfigEndpointMapping`

GetEndpointMapping returns the EndpointMapping field if non-nil, zero value otherwise.

### GetEndpointMappingOk

`func (o *GridCloudapiGatewayConfig) GetEndpointMappingOk() (*GridcloudapigatewayconfigEndpointMapping, bool)`

GetEndpointMappingOk returns a tuple with the EndpointMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointMapping

`func (o *GridCloudapiGatewayConfig) SetEndpointMapping(v GridcloudapigatewayconfigEndpointMapping)`

SetEndpointMapping sets EndpointMapping field to given value.

### HasEndpointMapping

`func (o *GridCloudapiGatewayConfig) HasEndpointMapping() bool`

HasEndpointMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


