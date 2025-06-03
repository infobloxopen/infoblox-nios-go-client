# MemberNtpSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableNtp** | Pointer to **bool** | Determines whether the NTP service is enabled on the member. | [optional] 
**NtpServers** | Pointer to [**MemberntpsettingNtpServers**](MemberntpsettingNtpServers.md) |  | [optional] 
**NtpKeys** | Pointer to [**MemberntpsettingNtpKeys**](MemberntpsettingNtpKeys.md) |  | [optional] 
**NtpAcl** | Pointer to [**MemberntpsettingNtpAcl**](MemberntpsettingNtpAcl.md) |  | [optional] 
**NtpKod** | Pointer to **bool** | Determines whether the Kiss-o&#39;-Death packets are enabled or disabled. | [optional] 
**EnableExternalNtpServers** | Pointer to **bool** | Determines whether the use of the external NTP servers is enabled for the member. | [optional] 
**ExcludeGridMasterNtpServer** | Pointer to **bool** | Determines whether the Grid Master is excluded as an NTP server. | [optional] 
**UseLocalNtpStratum** | Pointer to **bool** | Override Grid level NTP stratum. | [optional] 
**LocalNtpStratum** | Pointer to **int64** | Vnode level local NTP stratum. | [optional] 
**UseDefaultStratum** | Pointer to **bool** | Vnode level default stratum. | [optional] 
**UseNtpServers** | Pointer to **bool** | Use flag for: ntp_servers | [optional] 
**UseNtpKeys** | Pointer to **bool** | Use flag for: ntp_keys | [optional] 
**UseNtpAcl** | Pointer to **bool** | Use flag for: ntp_acl | [optional] 
**UseNtpKod** | Pointer to **bool** | Use flag for: ntp_kod | [optional] 

## Methods

### NewMemberNtpSetting

`func NewMemberNtpSetting() *MemberNtpSetting`

NewMemberNtpSetting instantiates a new MemberNtpSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberNtpSettingWithDefaults

`func NewMemberNtpSettingWithDefaults() *MemberNtpSetting`

NewMemberNtpSettingWithDefaults instantiates a new MemberNtpSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableNtp

`func (o *MemberNtpSetting) GetEnableNtp() bool`

GetEnableNtp returns the EnableNtp field if non-nil, zero value otherwise.

### GetEnableNtpOk

`func (o *MemberNtpSetting) GetEnableNtpOk() (*bool, bool)`

GetEnableNtpOk returns a tuple with the EnableNtp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNtp

`func (o *MemberNtpSetting) SetEnableNtp(v bool)`

SetEnableNtp sets EnableNtp field to given value.

### HasEnableNtp

`func (o *MemberNtpSetting) HasEnableNtp() bool`

HasEnableNtp returns a boolean if a field has been set.

### GetNtpServers

`func (o *MemberNtpSetting) GetNtpServers() MemberntpsettingNtpServers`

GetNtpServers returns the NtpServers field if non-nil, zero value otherwise.

### GetNtpServersOk

`func (o *MemberNtpSetting) GetNtpServersOk() (*MemberntpsettingNtpServers, bool)`

GetNtpServersOk returns a tuple with the NtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpServers

`func (o *MemberNtpSetting) SetNtpServers(v MemberntpsettingNtpServers)`

SetNtpServers sets NtpServers field to given value.

### HasNtpServers

`func (o *MemberNtpSetting) HasNtpServers() bool`

HasNtpServers returns a boolean if a field has been set.

### GetNtpKeys

`func (o *MemberNtpSetting) GetNtpKeys() MemberntpsettingNtpKeys`

GetNtpKeys returns the NtpKeys field if non-nil, zero value otherwise.

### GetNtpKeysOk

`func (o *MemberNtpSetting) GetNtpKeysOk() (*MemberntpsettingNtpKeys, bool)`

GetNtpKeysOk returns a tuple with the NtpKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpKeys

`func (o *MemberNtpSetting) SetNtpKeys(v MemberntpsettingNtpKeys)`

SetNtpKeys sets NtpKeys field to given value.

### HasNtpKeys

`func (o *MemberNtpSetting) HasNtpKeys() bool`

HasNtpKeys returns a boolean if a field has been set.

### GetNtpAcl

`func (o *MemberNtpSetting) GetNtpAcl() MemberntpsettingNtpAcl`

GetNtpAcl returns the NtpAcl field if non-nil, zero value otherwise.

### GetNtpAclOk

`func (o *MemberNtpSetting) GetNtpAclOk() (*MemberntpsettingNtpAcl, bool)`

GetNtpAclOk returns a tuple with the NtpAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpAcl

`func (o *MemberNtpSetting) SetNtpAcl(v MemberntpsettingNtpAcl)`

SetNtpAcl sets NtpAcl field to given value.

### HasNtpAcl

`func (o *MemberNtpSetting) HasNtpAcl() bool`

HasNtpAcl returns a boolean if a field has been set.

### GetNtpKod

`func (o *MemberNtpSetting) GetNtpKod() bool`

GetNtpKod returns the NtpKod field if non-nil, zero value otherwise.

### GetNtpKodOk

`func (o *MemberNtpSetting) GetNtpKodOk() (*bool, bool)`

GetNtpKodOk returns a tuple with the NtpKod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpKod

`func (o *MemberNtpSetting) SetNtpKod(v bool)`

SetNtpKod sets NtpKod field to given value.

### HasNtpKod

`func (o *MemberNtpSetting) HasNtpKod() bool`

HasNtpKod returns a boolean if a field has been set.

### GetEnableExternalNtpServers

`func (o *MemberNtpSetting) GetEnableExternalNtpServers() bool`

GetEnableExternalNtpServers returns the EnableExternalNtpServers field if non-nil, zero value otherwise.

### GetEnableExternalNtpServersOk

`func (o *MemberNtpSetting) GetEnableExternalNtpServersOk() (*bool, bool)`

GetEnableExternalNtpServersOk returns a tuple with the EnableExternalNtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableExternalNtpServers

`func (o *MemberNtpSetting) SetEnableExternalNtpServers(v bool)`

SetEnableExternalNtpServers sets EnableExternalNtpServers field to given value.

### HasEnableExternalNtpServers

`func (o *MemberNtpSetting) HasEnableExternalNtpServers() bool`

HasEnableExternalNtpServers returns a boolean if a field has been set.

### GetExcludeGridMasterNtpServer

`func (o *MemberNtpSetting) GetExcludeGridMasterNtpServer() bool`

GetExcludeGridMasterNtpServer returns the ExcludeGridMasterNtpServer field if non-nil, zero value otherwise.

### GetExcludeGridMasterNtpServerOk

`func (o *MemberNtpSetting) GetExcludeGridMasterNtpServerOk() (*bool, bool)`

GetExcludeGridMasterNtpServerOk returns a tuple with the ExcludeGridMasterNtpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGridMasterNtpServer

`func (o *MemberNtpSetting) SetExcludeGridMasterNtpServer(v bool)`

SetExcludeGridMasterNtpServer sets ExcludeGridMasterNtpServer field to given value.

### HasExcludeGridMasterNtpServer

`func (o *MemberNtpSetting) HasExcludeGridMasterNtpServer() bool`

HasExcludeGridMasterNtpServer returns a boolean if a field has been set.

### GetUseLocalNtpStratum

`func (o *MemberNtpSetting) GetUseLocalNtpStratum() bool`

GetUseLocalNtpStratum returns the UseLocalNtpStratum field if non-nil, zero value otherwise.

### GetUseLocalNtpStratumOk

`func (o *MemberNtpSetting) GetUseLocalNtpStratumOk() (*bool, bool)`

GetUseLocalNtpStratumOk returns a tuple with the UseLocalNtpStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLocalNtpStratum

`func (o *MemberNtpSetting) SetUseLocalNtpStratum(v bool)`

SetUseLocalNtpStratum sets UseLocalNtpStratum field to given value.

### HasUseLocalNtpStratum

`func (o *MemberNtpSetting) HasUseLocalNtpStratum() bool`

HasUseLocalNtpStratum returns a boolean if a field has been set.

### GetLocalNtpStratum

`func (o *MemberNtpSetting) GetLocalNtpStratum() int64`

GetLocalNtpStratum returns the LocalNtpStratum field if non-nil, zero value otherwise.

### GetLocalNtpStratumOk

`func (o *MemberNtpSetting) GetLocalNtpStratumOk() (*int64, bool)`

GetLocalNtpStratumOk returns a tuple with the LocalNtpStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalNtpStratum

`func (o *MemberNtpSetting) SetLocalNtpStratum(v int64)`

SetLocalNtpStratum sets LocalNtpStratum field to given value.

### HasLocalNtpStratum

`func (o *MemberNtpSetting) HasLocalNtpStratum() bool`

HasLocalNtpStratum returns a boolean if a field has been set.

### GetUseDefaultStratum

`func (o *MemberNtpSetting) GetUseDefaultStratum() bool`

GetUseDefaultStratum returns the UseDefaultStratum field if non-nil, zero value otherwise.

### GetUseDefaultStratumOk

`func (o *MemberNtpSetting) GetUseDefaultStratumOk() (*bool, bool)`

GetUseDefaultStratumOk returns a tuple with the UseDefaultStratum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultStratum

`func (o *MemberNtpSetting) SetUseDefaultStratum(v bool)`

SetUseDefaultStratum sets UseDefaultStratum field to given value.

### HasUseDefaultStratum

`func (o *MemberNtpSetting) HasUseDefaultStratum() bool`

HasUseDefaultStratum returns a boolean if a field has been set.

### GetUseNtpServers

`func (o *MemberNtpSetting) GetUseNtpServers() bool`

GetUseNtpServers returns the UseNtpServers field if non-nil, zero value otherwise.

### GetUseNtpServersOk

`func (o *MemberNtpSetting) GetUseNtpServersOk() (*bool, bool)`

GetUseNtpServersOk returns a tuple with the UseNtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNtpServers

`func (o *MemberNtpSetting) SetUseNtpServers(v bool)`

SetUseNtpServers sets UseNtpServers field to given value.

### HasUseNtpServers

`func (o *MemberNtpSetting) HasUseNtpServers() bool`

HasUseNtpServers returns a boolean if a field has been set.

### GetUseNtpKeys

`func (o *MemberNtpSetting) GetUseNtpKeys() bool`

GetUseNtpKeys returns the UseNtpKeys field if non-nil, zero value otherwise.

### GetUseNtpKeysOk

`func (o *MemberNtpSetting) GetUseNtpKeysOk() (*bool, bool)`

GetUseNtpKeysOk returns a tuple with the UseNtpKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNtpKeys

`func (o *MemberNtpSetting) SetUseNtpKeys(v bool)`

SetUseNtpKeys sets UseNtpKeys field to given value.

### HasUseNtpKeys

`func (o *MemberNtpSetting) HasUseNtpKeys() bool`

HasUseNtpKeys returns a boolean if a field has been set.

### GetUseNtpAcl

`func (o *MemberNtpSetting) GetUseNtpAcl() bool`

GetUseNtpAcl returns the UseNtpAcl field if non-nil, zero value otherwise.

### GetUseNtpAclOk

`func (o *MemberNtpSetting) GetUseNtpAclOk() (*bool, bool)`

GetUseNtpAclOk returns a tuple with the UseNtpAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNtpAcl

`func (o *MemberNtpSetting) SetUseNtpAcl(v bool)`

SetUseNtpAcl sets UseNtpAcl field to given value.

### HasUseNtpAcl

`func (o *MemberNtpSetting) HasUseNtpAcl() bool`

HasUseNtpAcl returns a boolean if a field has been set.

### GetUseNtpKod

`func (o *MemberNtpSetting) GetUseNtpKod() bool`

GetUseNtpKod returns the UseNtpKod field if non-nil, zero value otherwise.

### GetUseNtpKodOk

`func (o *MemberNtpSetting) GetUseNtpKodOk() (*bool, bool)`

GetUseNtpKodOk returns a tuple with the UseNtpKod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNtpKod

`func (o *MemberNtpSetting) SetUseNtpKod(v bool)`

SetUseNtpKod sets UseNtpKod field to given value.

### HasUseNtpKod

`func (o *MemberNtpSetting) HasUseNtpKod() bool`

HasUseNtpKod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


