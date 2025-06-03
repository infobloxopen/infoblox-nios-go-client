# GridSnmpSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | Pointer to **[]string** | The engine ID of the appliance that manages the SNMP agent. | [optional] [readonly] 
**QueriesCommunityString** | Pointer to **string** | The community string for SNMP queries. | [optional] 
**QueriesEnable** | Pointer to **bool** | If set to True, SNMP queries are enabled. | [optional] 
**Snmpv3QueriesEnable** | Pointer to **bool** | If set to True, SNMPv3 queries are enabled. | [optional] 
**Snmpv3QueriesUsers** | Pointer to [**GridsnmpsettingSnmpv3QueriesUsers**](GridsnmpsettingSnmpv3QueriesUsers.md) |  | [optional] 
**Snmpv3TrapsEnable** | Pointer to **bool** | If set to True, SNMPv3 traps are enabled. | [optional] 
**Syscontact** | Pointer to **[]string** | The name of the contact person for the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored. | [optional] 
**Sysdescr** | Pointer to **[]string** | Useful information about the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored. | [optional] 
**Syslocation** | Pointer to **[]string** | The physical location of the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored. | [optional] 
**Sysname** | Pointer to **[]string** | The FQDN (Fully Qualified Domain Name) of the appliance. Second value is applicable only for HA pair. Otherwise second value is ignored. | [optional] 
**TrapReceivers** | Pointer to [**GridsnmpsettingTrapReceivers**](GridsnmpsettingTrapReceivers.md) |  | [optional] 
**TrapsCommunityString** | Pointer to **string** | A string the NIOS appliance sends to the management system together with its traps. Note that this community string must match exactly what you enter in the management system. | [optional] 
**TrapsEnable** | Pointer to **bool** | If set to True, SNMP traps are enabled. | [optional] 

## Methods

### NewGridSnmpSetting

`func NewGridSnmpSetting() *GridSnmpSetting`

NewGridSnmpSetting instantiates a new GridSnmpSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridSnmpSettingWithDefaults

`func NewGridSnmpSettingWithDefaults() *GridSnmpSetting`

NewGridSnmpSettingWithDefaults instantiates a new GridSnmpSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *GridSnmpSetting) GetEngineId() []string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *GridSnmpSetting) GetEngineIdOk() (*[]string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *GridSnmpSetting) SetEngineId(v []string)`

SetEngineId sets EngineId field to given value.

### HasEngineId

`func (o *GridSnmpSetting) HasEngineId() bool`

HasEngineId returns a boolean if a field has been set.

### GetQueriesCommunityString

`func (o *GridSnmpSetting) GetQueriesCommunityString() string`

GetQueriesCommunityString returns the QueriesCommunityString field if non-nil, zero value otherwise.

### GetQueriesCommunityStringOk

`func (o *GridSnmpSetting) GetQueriesCommunityStringOk() (*string, bool)`

GetQueriesCommunityStringOk returns a tuple with the QueriesCommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueriesCommunityString

`func (o *GridSnmpSetting) SetQueriesCommunityString(v string)`

SetQueriesCommunityString sets QueriesCommunityString field to given value.

### HasQueriesCommunityString

`func (o *GridSnmpSetting) HasQueriesCommunityString() bool`

HasQueriesCommunityString returns a boolean if a field has been set.

### GetQueriesEnable

`func (o *GridSnmpSetting) GetQueriesEnable() bool`

GetQueriesEnable returns the QueriesEnable field if non-nil, zero value otherwise.

### GetQueriesEnableOk

`func (o *GridSnmpSetting) GetQueriesEnableOk() (*bool, bool)`

GetQueriesEnableOk returns a tuple with the QueriesEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueriesEnable

`func (o *GridSnmpSetting) SetQueriesEnable(v bool)`

SetQueriesEnable sets QueriesEnable field to given value.

### HasQueriesEnable

`func (o *GridSnmpSetting) HasQueriesEnable() bool`

HasQueriesEnable returns a boolean if a field has been set.

### GetSnmpv3QueriesEnable

`func (o *GridSnmpSetting) GetSnmpv3QueriesEnable() bool`

GetSnmpv3QueriesEnable returns the Snmpv3QueriesEnable field if non-nil, zero value otherwise.

### GetSnmpv3QueriesEnableOk

`func (o *GridSnmpSetting) GetSnmpv3QueriesEnableOk() (*bool, bool)`

GetSnmpv3QueriesEnableOk returns a tuple with the Snmpv3QueriesEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3QueriesEnable

`func (o *GridSnmpSetting) SetSnmpv3QueriesEnable(v bool)`

SetSnmpv3QueriesEnable sets Snmpv3QueriesEnable field to given value.

### HasSnmpv3QueriesEnable

`func (o *GridSnmpSetting) HasSnmpv3QueriesEnable() bool`

HasSnmpv3QueriesEnable returns a boolean if a field has been set.

### GetSnmpv3QueriesUsers

`func (o *GridSnmpSetting) GetSnmpv3QueriesUsers() GridsnmpsettingSnmpv3QueriesUsers`

GetSnmpv3QueriesUsers returns the Snmpv3QueriesUsers field if non-nil, zero value otherwise.

### GetSnmpv3QueriesUsersOk

`func (o *GridSnmpSetting) GetSnmpv3QueriesUsersOk() (*GridsnmpsettingSnmpv3QueriesUsers, bool)`

GetSnmpv3QueriesUsersOk returns a tuple with the Snmpv3QueriesUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3QueriesUsers

`func (o *GridSnmpSetting) SetSnmpv3QueriesUsers(v GridsnmpsettingSnmpv3QueriesUsers)`

SetSnmpv3QueriesUsers sets Snmpv3QueriesUsers field to given value.

### HasSnmpv3QueriesUsers

`func (o *GridSnmpSetting) HasSnmpv3QueriesUsers() bool`

HasSnmpv3QueriesUsers returns a boolean if a field has been set.

### GetSnmpv3TrapsEnable

`func (o *GridSnmpSetting) GetSnmpv3TrapsEnable() bool`

GetSnmpv3TrapsEnable returns the Snmpv3TrapsEnable field if non-nil, zero value otherwise.

### GetSnmpv3TrapsEnableOk

`func (o *GridSnmpSetting) GetSnmpv3TrapsEnableOk() (*bool, bool)`

GetSnmpv3TrapsEnableOk returns a tuple with the Snmpv3TrapsEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpv3TrapsEnable

`func (o *GridSnmpSetting) SetSnmpv3TrapsEnable(v bool)`

SetSnmpv3TrapsEnable sets Snmpv3TrapsEnable field to given value.

### HasSnmpv3TrapsEnable

`func (o *GridSnmpSetting) HasSnmpv3TrapsEnable() bool`

HasSnmpv3TrapsEnable returns a boolean if a field has been set.

### GetSyscontact

`func (o *GridSnmpSetting) GetSyscontact() []string`

GetSyscontact returns the Syscontact field if non-nil, zero value otherwise.

### GetSyscontactOk

`func (o *GridSnmpSetting) GetSyscontactOk() (*[]string, bool)`

GetSyscontactOk returns a tuple with the Syscontact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyscontact

`func (o *GridSnmpSetting) SetSyscontact(v []string)`

SetSyscontact sets Syscontact field to given value.

### HasSyscontact

`func (o *GridSnmpSetting) HasSyscontact() bool`

HasSyscontact returns a boolean if a field has been set.

### GetSysdescr

`func (o *GridSnmpSetting) GetSysdescr() []string`

GetSysdescr returns the Sysdescr field if non-nil, zero value otherwise.

### GetSysdescrOk

`func (o *GridSnmpSetting) GetSysdescrOk() (*[]string, bool)`

GetSysdescrOk returns a tuple with the Sysdescr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysdescr

`func (o *GridSnmpSetting) SetSysdescr(v []string)`

SetSysdescr sets Sysdescr field to given value.

### HasSysdescr

`func (o *GridSnmpSetting) HasSysdescr() bool`

HasSysdescr returns a boolean if a field has been set.

### GetSyslocation

`func (o *GridSnmpSetting) GetSyslocation() []string`

GetSyslocation returns the Syslocation field if non-nil, zero value otherwise.

### GetSyslocationOk

`func (o *GridSnmpSetting) GetSyslocationOk() (*[]string, bool)`

GetSyslocationOk returns a tuple with the Syslocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslocation

`func (o *GridSnmpSetting) SetSyslocation(v []string)`

SetSyslocation sets Syslocation field to given value.

### HasSyslocation

`func (o *GridSnmpSetting) HasSyslocation() bool`

HasSyslocation returns a boolean if a field has been set.

### GetSysname

`func (o *GridSnmpSetting) GetSysname() []string`

GetSysname returns the Sysname field if non-nil, zero value otherwise.

### GetSysnameOk

`func (o *GridSnmpSetting) GetSysnameOk() (*[]string, bool)`

GetSysnameOk returns a tuple with the Sysname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSysname

`func (o *GridSnmpSetting) SetSysname(v []string)`

SetSysname sets Sysname field to given value.

### HasSysname

`func (o *GridSnmpSetting) HasSysname() bool`

HasSysname returns a boolean if a field has been set.

### GetTrapReceivers

`func (o *GridSnmpSetting) GetTrapReceivers() GridsnmpsettingTrapReceivers`

GetTrapReceivers returns the TrapReceivers field if non-nil, zero value otherwise.

### GetTrapReceiversOk

`func (o *GridSnmpSetting) GetTrapReceiversOk() (*GridsnmpsettingTrapReceivers, bool)`

GetTrapReceiversOk returns a tuple with the TrapReceivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapReceivers

`func (o *GridSnmpSetting) SetTrapReceivers(v GridsnmpsettingTrapReceivers)`

SetTrapReceivers sets TrapReceivers field to given value.

### HasTrapReceivers

`func (o *GridSnmpSetting) HasTrapReceivers() bool`

HasTrapReceivers returns a boolean if a field has been set.

### GetTrapsCommunityString

`func (o *GridSnmpSetting) GetTrapsCommunityString() string`

GetTrapsCommunityString returns the TrapsCommunityString field if non-nil, zero value otherwise.

### GetTrapsCommunityStringOk

`func (o *GridSnmpSetting) GetTrapsCommunityStringOk() (*string, bool)`

GetTrapsCommunityStringOk returns a tuple with the TrapsCommunityString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapsCommunityString

`func (o *GridSnmpSetting) SetTrapsCommunityString(v string)`

SetTrapsCommunityString sets TrapsCommunityString field to given value.

### HasTrapsCommunityString

`func (o *GridSnmpSetting) HasTrapsCommunityString() bool`

HasTrapsCommunityString returns a boolean if a field has been set.

### GetTrapsEnable

`func (o *GridSnmpSetting) GetTrapsEnable() bool`

GetTrapsEnable returns the TrapsEnable field if non-nil, zero value otherwise.

### GetTrapsEnableOk

`func (o *GridSnmpSetting) GetTrapsEnableOk() (*bool, bool)`

GetTrapsEnableOk returns a tuple with the TrapsEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapsEnable

`func (o *GridSnmpSetting) SetTrapsEnable(v bool)`

SetTrapsEnable sets TrapsEnable field to given value.

### HasTrapsEnable

`func (o *GridSnmpSetting) HasTrapsEnable() bool`

HasTrapsEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


