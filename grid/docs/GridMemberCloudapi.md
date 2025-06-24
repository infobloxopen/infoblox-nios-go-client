# GridMemberCloudapi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllowApiAdmins** | Pointer to **string** | Defines which administrators are allowed to perform Cloud API request on the Grid Member: no administrators (NONE), any administrators (ALL) or administrators in the ACL list (LIST). Default is ALL. | [optional] 
**AllowedApiAdmins** | Pointer to [**[]GridMemberCloudapiAllowedApiAdmins**](GridMemberCloudapiAllowedApiAdmins.md) | List of administrators allowed to perform Cloud Platform API requests on that member. | [optional] 
**EnableService** | Pointer to **bool** | Controls whether the Cloud API service runs on the member or not. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] [readonly] 
**GatewayConfig** | Pointer to [**GridMemberCloudapiGatewayConfig**](GridMemberCloudapiGatewayConfig.md) |  | [optional] 
**Member** | Pointer to [**GridMemberCloudapiMember**](GridMemberCloudapiMember.md) |  | [optional] 
**Status** | Pointer to **string** | Status of Cloud API service on the member. | [optional] [readonly] 

## Methods

### NewGridMemberCloudapi

`func NewGridMemberCloudapi() *GridMemberCloudapi`

NewGridMemberCloudapi instantiates a new GridMemberCloudapi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridMemberCloudapiWithDefaults

`func NewGridMemberCloudapiWithDefaults() *GridMemberCloudapi`

NewGridMemberCloudapiWithDefaults instantiates a new GridMemberCloudapi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridMemberCloudapi) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridMemberCloudapi) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridMemberCloudapi) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridMemberCloudapi) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowApiAdmins

`func (o *GridMemberCloudapi) GetAllowApiAdmins() string`

GetAllowApiAdmins returns the AllowApiAdmins field if non-nil, zero value otherwise.

### GetAllowApiAdminsOk

`func (o *GridMemberCloudapi) GetAllowApiAdminsOk() (*string, bool)`

GetAllowApiAdminsOk returns a tuple with the AllowApiAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApiAdmins

`func (o *GridMemberCloudapi) SetAllowApiAdmins(v string)`

SetAllowApiAdmins sets AllowApiAdmins field to given value.

### HasAllowApiAdmins

`func (o *GridMemberCloudapi) HasAllowApiAdmins() bool`

HasAllowApiAdmins returns a boolean if a field has been set.

### GetAllowedApiAdmins

`func (o *GridMemberCloudapi) GetAllowedApiAdmins() []GridMemberCloudapiAllowedApiAdmins`

GetAllowedApiAdmins returns the AllowedApiAdmins field if non-nil, zero value otherwise.

### GetAllowedApiAdminsOk

`func (o *GridMemberCloudapi) GetAllowedApiAdminsOk() (*[]GridMemberCloudapiAllowedApiAdmins, bool)`

GetAllowedApiAdminsOk returns a tuple with the AllowedApiAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedApiAdmins

`func (o *GridMemberCloudapi) SetAllowedApiAdmins(v []GridMemberCloudapiAllowedApiAdmins)`

SetAllowedApiAdmins sets AllowedApiAdmins field to given value.

### HasAllowedApiAdmins

`func (o *GridMemberCloudapi) HasAllowedApiAdmins() bool`

HasAllowedApiAdmins returns a boolean if a field has been set.

### GetEnableService

`func (o *GridMemberCloudapi) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *GridMemberCloudapi) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *GridMemberCloudapi) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *GridMemberCloudapi) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GridMemberCloudapi) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GridMemberCloudapi) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GridMemberCloudapi) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GridMemberCloudapi) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetGatewayConfig

`func (o *GridMemberCloudapi) GetGatewayConfig() GridMemberCloudapiGatewayConfig`

GetGatewayConfig returns the GatewayConfig field if non-nil, zero value otherwise.

### GetGatewayConfigOk

`func (o *GridMemberCloudapi) GetGatewayConfigOk() (*GridMemberCloudapiGatewayConfig, bool)`

GetGatewayConfigOk returns a tuple with the GatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConfig

`func (o *GridMemberCloudapi) SetGatewayConfig(v GridMemberCloudapiGatewayConfig)`

SetGatewayConfig sets GatewayConfig field to given value.

### HasGatewayConfig

`func (o *GridMemberCloudapi) HasGatewayConfig() bool`

HasGatewayConfig returns a boolean if a field has been set.

### GetMember

`func (o *GridMemberCloudapi) GetMember() GridMemberCloudapiMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GridMemberCloudapi) GetMemberOk() (*GridMemberCloudapiMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GridMemberCloudapi) SetMember(v GridMemberCloudapiMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *GridMemberCloudapi) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetStatus

`func (o *GridMemberCloudapi) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GridMemberCloudapi) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GridMemberCloudapi) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GridMemberCloudapi) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


