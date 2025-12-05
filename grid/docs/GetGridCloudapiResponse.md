# GetGridCloudapiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AllowApiAdmins** | Pointer to **string** | Defines administrators who can perform cloud API requests on the Grid Master. The valid value is NONE (no administrator), ALL (all administrators), or LIST (administrators on the ACL). | [optional] 
**AllowedApiAdmins** | Pointer to [**[]GridCloudapiAllowedApiAdmins**](GridCloudapiAllowedApiAdmins.md) | The list of administrators who can perform cloud API requests on the Cloud Platform Appliance. | [optional] 
**EnableRecycleBin** | Pointer to **bool** | Determines whether the recycle bin for deleted cloud objects is enabled or not on the Grid Master. | [optional] 
**GatewayConfig** | Pointer to [**GridCloudapiGatewayConfig**](GridCloudapiGatewayConfig.md) |  | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**GridCloudapi**](GridCloudapi.md) |  | [optional] 

## Methods

### NewGetGridCloudapiResponse

`func NewGetGridCloudapiResponse() *GetGridCloudapiResponse`

NewGetGridCloudapiResponse instantiates a new GetGridCloudapiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridCloudapiResponseWithDefaults

`func NewGetGridCloudapiResponseWithDefaults() *GetGridCloudapiResponse`

NewGetGridCloudapiResponseWithDefaults instantiates a new GetGridCloudapiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridCloudapiResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridCloudapiResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridCloudapiResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridCloudapiResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowApiAdmins

`func (o *GetGridCloudapiResponse) GetAllowApiAdmins() string`

GetAllowApiAdmins returns the AllowApiAdmins field if non-nil, zero value otherwise.

### GetAllowApiAdminsOk

`func (o *GetGridCloudapiResponse) GetAllowApiAdminsOk() (*string, bool)`

GetAllowApiAdminsOk returns a tuple with the AllowApiAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowApiAdmins

`func (o *GetGridCloudapiResponse) SetAllowApiAdmins(v string)`

SetAllowApiAdmins sets AllowApiAdmins field to given value.

### HasAllowApiAdmins

`func (o *GetGridCloudapiResponse) HasAllowApiAdmins() bool`

HasAllowApiAdmins returns a boolean if a field has been set.

### GetAllowedApiAdmins

`func (o *GetGridCloudapiResponse) GetAllowedApiAdmins() []GridCloudapiAllowedApiAdmins`

GetAllowedApiAdmins returns the AllowedApiAdmins field if non-nil, zero value otherwise.

### GetAllowedApiAdminsOk

`func (o *GetGridCloudapiResponse) GetAllowedApiAdminsOk() (*[]GridCloudapiAllowedApiAdmins, bool)`

GetAllowedApiAdminsOk returns a tuple with the AllowedApiAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedApiAdmins

`func (o *GetGridCloudapiResponse) SetAllowedApiAdmins(v []GridCloudapiAllowedApiAdmins)`

SetAllowedApiAdmins sets AllowedApiAdmins field to given value.

### HasAllowedApiAdmins

`func (o *GetGridCloudapiResponse) HasAllowedApiAdmins() bool`

HasAllowedApiAdmins returns a boolean if a field has been set.

### GetEnableRecycleBin

`func (o *GetGridCloudapiResponse) GetEnableRecycleBin() bool`

GetEnableRecycleBin returns the EnableRecycleBin field if non-nil, zero value otherwise.

### GetEnableRecycleBinOk

`func (o *GetGridCloudapiResponse) GetEnableRecycleBinOk() (*bool, bool)`

GetEnableRecycleBinOk returns a tuple with the EnableRecycleBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecycleBin

`func (o *GetGridCloudapiResponse) SetEnableRecycleBin(v bool)`

SetEnableRecycleBin sets EnableRecycleBin field to given value.

### HasEnableRecycleBin

`func (o *GetGridCloudapiResponse) HasEnableRecycleBin() bool`

HasEnableRecycleBin returns a boolean if a field has been set.

### GetGatewayConfig

`func (o *GetGridCloudapiResponse) GetGatewayConfig() GridCloudapiGatewayConfig`

GetGatewayConfig returns the GatewayConfig field if non-nil, zero value otherwise.

### GetGatewayConfigOk

`func (o *GetGridCloudapiResponse) GetGatewayConfigOk() (*GridCloudapiGatewayConfig, bool)`

GetGatewayConfigOk returns a tuple with the GatewayConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayConfig

`func (o *GetGridCloudapiResponse) SetGatewayConfig(v GridCloudapiGatewayConfig)`

SetGatewayConfig sets GatewayConfig field to given value.

### HasGatewayConfig

`func (o *GetGridCloudapiResponse) HasGatewayConfig() bool`

HasGatewayConfig returns a boolean if a field has been set.

### GetUuid

`func (o *GetGridCloudapiResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGridCloudapiResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGridCloudapiResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGridCloudapiResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetGridCloudapiResponse) GetResult() GridCloudapi`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridCloudapiResponse) GetResultOk() (*GridCloudapi, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridCloudapiResponse) SetResult(v GridCloudapi)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridCloudapiResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


