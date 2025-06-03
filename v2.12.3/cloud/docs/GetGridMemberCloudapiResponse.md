# GetGridMemberCloudapiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AllowApiAdmins** | Pointer to **string** | Defines which administrators are allowed to perform Cloud API request on the Grid Member: no administrators (NONE), any administrators (ALL) or administrators in the ACL list (LIST). Default is ALL. | [optional] 
**AllowedApiAdmins** | Pointer to [**[]GridMemberCloudapiAllowedApiAdmins**](GridMemberCloudapiAllowedApiAdmins.md) | List of administrators allowed to perform Cloud Platform API requests on that member. | [optional] 
**EnableService** | Pointer to **bool** | Controls whether the Cloud API service runs on the member or not. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] [readonly] 
**GatewayConfig** | Pointer to [**GridMemberCloudapiGatewayConfig**](GridMemberCloudapiGatewayConfig.md) |  | [optional] 
**Member** | Pointer to [**GridMemberCloudapiMember**](GridMemberCloudapiMember.md) |  | [optional] 
**Status** | Pointer to **string** | Status of Cloud API service on the member. | [optional] [readonly] 
**Result** | Pointer to [**GridMemberCloudapi**](GridMemberCloudapi.md) |  | [optional] 

## Methods

### NewGetGridMemberCloudapiResponse

`func NewGetGridMemberCloudapiResponse() *GetGridMemberCloudapiResponse`

NewGetGridMemberCloudapiResponse instantiates a new GetGridMemberCloudapiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridMemberCloudapiResponseWithDefaults

`func NewGetGridMemberCloudapiResponseWithDefaults() *GetGridMemberCloudapiResponse`

NewGetGridMemberCloudapiResponseWithDefaults instantiates a new GetGridMemberCloudapiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridMemberCloudapiResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridMemberCloudapiResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridMemberCloudapiResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridMemberCloudapiResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowApiAdmins

`func (o *GetGridMemberCloudapiResponse) GetAllowApiAdmins() string`

GetAllowApiAdmins returns the AllowApiAdmins field if non-nil, zero value otherwise.

### GetAllowApiAdminsOk

`func (o *GetGridMemberCloudapiResponse) GetAllowApiAdminsOk() (*string, bool)`

GetAllowApiAdminsOk returns a tuple with the AllowApiAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApiAdmins

`func (o *GetGridMemberCloudapiResponse) SetAllowApiAdmins(v string)`

SetAllowApiAdmins sets AllowApiAdmins field to given value.

### HasAllowApiAdmins

`func (o *GetGridMemberCloudapiResponse) HasAllowApiAdmins() bool`

HasAllowApiAdmins returns a boolean if a field has been set.

### GetAllowedApiAdmins

`func (o *GetGridMemberCloudapiResponse) GetAllowedApiAdmins() []GridMemberCloudapiAllowedApiAdmins`

GetAllowedApiAdmins returns the AllowedApiAdmins field if non-nil, zero value otherwise.

### GetAllowedApiAdminsOk

`func (o *GetGridMemberCloudapiResponse) GetAllowedApiAdminsOk() (*[]GridMemberCloudapiAllowedApiAdmins, bool)`

GetAllowedApiAdminsOk returns a tuple with the AllowedApiAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedApiAdmins

`func (o *GetGridMemberCloudapiResponse) SetAllowedApiAdmins(v []GridMemberCloudapiAllowedApiAdmins)`

SetAllowedApiAdmins sets AllowedApiAdmins field to given value.

### HasAllowedApiAdmins

`func (o *GetGridMemberCloudapiResponse) HasAllowedApiAdmins() bool`

HasAllowedApiAdmins returns a boolean if a field has been set.

### GetEnableService

`func (o *GetGridMemberCloudapiResponse) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *GetGridMemberCloudapiResponse) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *GetGridMemberCloudapiResponse) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *GetGridMemberCloudapiResponse) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetGridMemberCloudapiResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetGridMemberCloudapiResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetGridMemberCloudapiResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetGridMemberCloudapiResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetGatewayConfig

`func (o *GetGridMemberCloudapiResponse) GetGatewayConfig() GridMemberCloudapiGatewayConfig`

GetGatewayConfig returns the GatewayConfig field if non-nil, zero value otherwise.

### GetGatewayConfigOk

`func (o *GetGridMemberCloudapiResponse) GetGatewayConfigOk() (*GridMemberCloudapiGatewayConfig, bool)`

GetGatewayConfigOk returns a tuple with the GatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConfig

`func (o *GetGridMemberCloudapiResponse) SetGatewayConfig(v GridMemberCloudapiGatewayConfig)`

SetGatewayConfig sets GatewayConfig field to given value.

### HasGatewayConfig

`func (o *GetGridMemberCloudapiResponse) HasGatewayConfig() bool`

HasGatewayConfig returns a boolean if a field has been set.

### GetMember

`func (o *GetGridMemberCloudapiResponse) GetMember() GridMemberCloudapiMember`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *GetGridMemberCloudapiResponse) GetMemberOk() (*GridMemberCloudapiMember, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *GetGridMemberCloudapiResponse) SetMember(v GridMemberCloudapiMember)`

SetMember sets Member field to given value.

### HasMember

`func (o *GetGridMemberCloudapiResponse) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetStatus

`func (o *GetGridMemberCloudapiResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetGridMemberCloudapiResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetGridMemberCloudapiResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetGridMemberCloudapiResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *GetGridMemberCloudapiResponse) GetResult() GridMemberCloudapi`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridMemberCloudapiResponse) GetResultOk() (*GridMemberCloudapi, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridMemberCloudapiResponse) SetResult(v GridMemberCloudapi)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridMemberCloudapiResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


