# GridNtpSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableNtp** | Pointer to **bool** | Determines whether NTP is enabled on the Grid. | [optional] 
**NtpServers** | Pointer to [**GridntpsettingNtpServers**](GridntpsettingNtpServers.md) |  | [optional] 
**NtpKeys** | Pointer to [**GridntpsettingNtpKeys**](GridntpsettingNtpKeys.md) |  | [optional] 
**NtpAcl** | Pointer to [**GridntpsettingNtpAcl**](GridntpsettingNtpAcl.md) |  | [optional] 
**NtpKod** | Pointer to **bool** | Determines whether the Kiss-o&#39;-Death packets are enabled. | [optional] 
**GmLocalNtpStratum** | Pointer to **int64** | Grid level GM local NTP stratum. | [optional] 
**LocalNtpStratum** | Pointer to **int64** | Local NTP stratum for non-GM members. | [optional] 
**UseDefaultStratum** | Pointer to **bool** | This flag controls whether gm_local_ntp_stratum value be set to a default value | [optional] 

## Methods

### NewGridNtpSetting

`func NewGridNtpSetting() *GridNtpSetting`

NewGridNtpSetting instantiates a new GridNtpSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridNtpSettingWithDefaults

`func NewGridNtpSettingWithDefaults() *GridNtpSetting`

NewGridNtpSettingWithDefaults instantiates a new GridNtpSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableNtp

`func (o *GridNtpSetting) GetEnableNtp() bool`

GetEnableNtp returns the EnableNtp field if non-nil, zero value otherwise.

### GetEnableNtpOk

`func (o *GridNtpSetting) GetEnableNtpOk() (*bool, bool)`

GetEnableNtpOk returns a tuple with the EnableNtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNtp

`func (o *GridNtpSetting) SetEnableNtp(v bool)`

SetEnableNtp sets EnableNtp field to given value.

### HasEnableNtp

`func (o *GridNtpSetting) HasEnableNtp() bool`

HasEnableNtp returns a boolean if a field has been set.

### GetNtpServers

`func (o *GridNtpSetting) GetNtpServers() GridntpsettingNtpServers`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *GridNtpSetting) GetNtpServersOk() (*GridntpsettingNtpServers, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *GridNtpSetting) SetNtpServers(v GridntpsettingNtpServers)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *GridNtpSetting) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetNtpKeys

`func (o *GridNtpSetting) GetNtpKeys() GridntpsettingNtpKeys`

GetNtpKeys returns the NtpKeys field if non-nil, zero value otherwise.

### GetNtpKeysOk

`func (o *GridNtpSetting) GetNtpKeysOk() (*GridntpsettingNtpKeys, bool)`

GetNtpKeysOk returns a tuple with the NtpKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpKeys

`func (o *GridNtpSetting) SetNtpKeys(v GridntpsettingNtpKeys)`

SetNtpKeys sets NtpKeys field to given value.

### HasNtpKeys

`func (o *GridNtpSetting) HasNtpKeys() bool`

HasNtpKeys returns a boolean if a field has been set.

### GetNtpAcl

`func (o *GridNtpSetting) GetNtpAcl() GridntpsettingNtpAcl`

GetNtpAcl returns the NtpAcl field if non-nil, zero value otherwise.

### GetNtpAclOk

`func (o *GridNtpSetting) GetNtpAclOk() (*GridntpsettingNtpAcl, bool)`

GetNtpAclOk returns a tuple with the NtpAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpAcl

`func (o *GridNtpSetting) SetNtpAcl(v GridntpsettingNtpAcl)`

SetNtpAcl sets NtpAcl field to given value.

### HasNtpAcl

`func (o *GridNtpSetting) HasNtpAcl() bool`

HasNtpAcl returns a boolean if a field has been set.

### GetNtpKod

`func (o *GridNtpSetting) GetNtpKod() bool`

GetNtpKod returns the NtpKod field if non-nil, zero value otherwise.

### GetNtpKodOk

`func (o *GridNtpSetting) GetNtpKodOk() (*bool, bool)`

GetNtpKodOk returns a tuple with the NtpKod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpKod

`func (o *GridNtpSetting) SetNtpKod(v bool)`

SetNtpKod sets NtpKod field to given value.

### HasNtpKod

`func (o *GridNtpSetting) HasNtpKod() bool`

HasNtpKod returns a boolean if a field has been set.

### GetGmLocalNtpStratum

`func (o *GridNtpSetting) GetGmLocalNtpStratum() int64`

GetGmLocalNtpStratum returns the GmLocalNtpStratum field if non-nil, zero value otherwise.

### GetGmLocalNtpStratumOk

`func (o *GridNtpSetting) GetGmLocalNtpStratumOk() (*int64, bool)`

GetGmLocalNtpStratumOk returns a tuple with the GmLocalNtpStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmLocalNtpStratum

`func (o *GridNtpSetting) SetGmLocalNtpStratum(v int64)`

SetGmLocalNtpStratum sets GmLocalNtpStratum field to given value.

### HasGmLocalNtpStratum

`func (o *GridNtpSetting) HasGmLocalNtpStratum() bool`

HasGmLocalNtpStratum returns a boolean if a field has been set.

### GetLocalNtpStratum

`func (o *GridNtpSetting) GetLocalNtpStratum() int64`

GetLocalNtpStratum returns the LocalNtpStratum field if non-nil, zero value otherwise.

### GetLocalNtpStratumOk

`func (o *GridNtpSetting) GetLocalNtpStratumOk() (*int64, bool)`

GetLocalNtpStratumOk returns a tuple with the LocalNtpStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNtpStratum

`func (o *GridNtpSetting) SetLocalNtpStratum(v int64)`

SetLocalNtpStratum sets LocalNtpStratum field to given value.

### HasLocalNtpStratum

`func (o *GridNtpSetting) HasLocalNtpStratum() bool`

HasLocalNtpStratum returns a boolean if a field has been set.

### GetUseDefaultStratum

`func (o *GridNtpSetting) GetUseDefaultStratum() bool`

GetUseDefaultStratum returns the UseDefaultStratum field if non-nil, zero value otherwise.

### GetUseDefaultStratumOk

`func (o *GridNtpSetting) GetUseDefaultStratumOk() (*bool, bool)`

GetUseDefaultStratumOk returns a tuple with the UseDefaultStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultStratum

`func (o *GridNtpSetting) SetUseDefaultStratum(v bool)`

SetUseDefaultStratum sets UseDefaultStratum field to given value.

### HasUseDefaultStratum

`func (o *GridNtpSetting) HasUseDefaultStratum() bool`

HasUseDefaultStratum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


