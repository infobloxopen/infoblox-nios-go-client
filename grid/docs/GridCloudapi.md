# GridCloudapi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AllowApiAdmins** | Pointer to **string** | Defines administrators who can perform cloud API requests on the Grid Master. The valid value is NONE (no administrator), ALL (all administrators), or LIST (administrators on the ACL). | [optional] 
**AllowedApiAdmins** | Pointer to [**[]GridCloudapiAllowedApiAdmins**](GridCloudapiAllowedApiAdmins.md) | The list of administrators who can perform cloud API requests on the Cloud Platform Appliance. | [optional] 
**EnableRecycleBin** | Pointer to **bool** | Determines whether the recycle bin for deleted cloud objects is enabled or not on the Grid Master. | [optional] 
**GatewayConfig** | Pointer to [**GridCloudapiGatewayConfig**](GridCloudapiGatewayConfig.md) |  | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewGridCloudapi

`func NewGridCloudapi() *GridCloudapi`

NewGridCloudapi instantiates a new GridCloudapi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridCloudapiWithDefaults

`func NewGridCloudapiWithDefaults() *GridCloudapi`

NewGridCloudapiWithDefaults instantiates a new GridCloudapi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridCloudapi) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridCloudapi) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridCloudapi) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridCloudapi) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowApiAdmins

`func (o *GridCloudapi) GetAllowApiAdmins() string`

GetAllowApiAdmins returns the AllowApiAdmins field if non-nil, zero value otherwise.

### GetAllowApiAdminsOk

`func (o *GridCloudapi) GetAllowApiAdminsOk() (*string, bool)`

GetAllowApiAdminsOk returns a tuple with the AllowApiAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApiAdmins

`func (o *GridCloudapi) SetAllowApiAdmins(v string)`

SetAllowApiAdmins sets AllowApiAdmins field to given value.

### HasAllowApiAdmins

`func (o *GridCloudapi) HasAllowApiAdmins() bool`

HasAllowApiAdmins returns a boolean if a field has been set.

### GetAllowedApiAdmins

`func (o *GridCloudapi) GetAllowedApiAdmins() []GridCloudapiAllowedApiAdmins`

GetAllowedApiAdmins returns the AllowedApiAdmins field if non-nil, zero value otherwise.

### GetAllowedApiAdminsOk

`func (o *GridCloudapi) GetAllowedApiAdminsOk() (*[]GridCloudapiAllowedApiAdmins, bool)`

GetAllowedApiAdminsOk returns a tuple with the AllowedApiAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedApiAdmins

`func (o *GridCloudapi) SetAllowedApiAdmins(v []GridCloudapiAllowedApiAdmins)`

SetAllowedApiAdmins sets AllowedApiAdmins field to given value.

### HasAllowedApiAdmins

`func (o *GridCloudapi) HasAllowedApiAdmins() bool`

HasAllowedApiAdmins returns a boolean if a field has been set.

### GetEnableRecycleBin

`func (o *GridCloudapi) GetEnableRecycleBin() bool`

GetEnableRecycleBin returns the EnableRecycleBin field if non-nil, zero value otherwise.

### GetEnableRecycleBinOk

`func (o *GridCloudapi) GetEnableRecycleBinOk() (*bool, bool)`

GetEnableRecycleBinOk returns a tuple with the EnableRecycleBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecycleBin

`func (o *GridCloudapi) SetEnableRecycleBin(v bool)`

SetEnableRecycleBin sets EnableRecycleBin field to given value.

### HasEnableRecycleBin

`func (o *GridCloudapi) HasEnableRecycleBin() bool`

HasEnableRecycleBin returns a boolean if a field has been set.

### GetGatewayConfig

`func (o *GridCloudapi) GetGatewayConfig() GridCloudapiGatewayConfig`

GetGatewayConfig returns the GatewayConfig field if non-nil, zero value otherwise.

### GetGatewayConfigOk

`func (o *GridCloudapi) GetGatewayConfigOk() (*GridCloudapiGatewayConfig, bool)`

GetGatewayConfigOk returns a tuple with the GatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConfig

`func (o *GridCloudapi) SetGatewayConfig(v GridCloudapiGatewayConfig)`

SetGatewayConfig sets GatewayConfig field to given value.

### HasGatewayConfig

`func (o *GridCloudapi) HasGatewayConfig() bool`

HasGatewayConfig returns a boolean if a field has been set.

### GetUuid

`func (o *GridCloudapi) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GridCloudapi) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GridCloudapi) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GridCloudapi) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


