# MsserverAdSites

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseDefaultIpSiteLink** | Pointer to **bool** | Flag to override MS Server default IP site link | [optional] 
**DefaultIpSiteLink** | Pointer to **string** | Default IP site link for sites created from NIOS | [optional] 
**UseLogin** | Pointer to **bool** | Flag to override login name and password from the MS Server | [optional] 
**LoginName** | Pointer to **string** | Microsoft Server login name, with optional | [optional] 
**LoginPassword** | Pointer to **string** | Microsoft Server login password. | [optional] 
**UseSynchronizationMinDelay** | Pointer to **bool** | Flag to override synchronization interval from the MS Server | [optional] 
**SynchronizationMinDelay** | Pointer to **int64** | Minimum number of minutes between two synchronizations | [optional] 
**UseLdapTimeout** | Pointer to **bool** | Flag to override cluster LDAP timeoutMS Server | [optional] 
**LdapTimeout** | Pointer to **int64** | Timeout in seconds for LDAP connections for this MS Server | [optional] 
**LdapAuthPort** | Pointer to **int64** | TCP port for LDAP connections for this | [optional] 
**LdapEncryption** | Pointer to **string** | Encryption for LDAP connections for this MS Server | [optional] 
**Managed** | Pointer to **bool** | Controls whether the Sites of this MS Server are to be synchronized by the assigned managing member or not | [optional] 
**ReadOnly** | Pointer to **bool** | Enable/disable read-only synchronization of Sites for this Active Directory domain | [optional] 
**LastSyncTs** | Pointer to **int64** | Timestamp of the last synchronization attempt | [optional] [readonly] 
**LastSyncStatus** | Pointer to **string** | Status of the last synchronization attempt | [optional] [readonly] 
**LastSyncDetail** | Pointer to **string** | The detailed status of the last synchronization attempt. | [optional] [readonly] 
**SupportsIpv6** | Pointer to **bool** | Flag indicating if the server supports IPv6 | [optional] [readonly] 

## Methods

### NewMsserverAdSites

`func NewMsserverAdSites() *MsserverAdSites`

NewMsserverAdSites instantiates a new MsserverAdSites object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMsserverAdSitesWithDefaults

`func NewMsserverAdSitesWithDefaults() *MsserverAdSites`

NewMsserverAdSitesWithDefaults instantiates a new MsserverAdSites object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseDefaultIpSiteLink

`func (o *MsserverAdSites) GetUseDefaultIpSiteLink() bool`

GetUseDefaultIpSiteLink returns the UseDefaultIpSiteLink field if non-nil, zero value otherwise.

### GetUseDefaultIpSiteLinkOk

`func (o *MsserverAdSites) GetUseDefaultIpSiteLinkOk() (*bool, bool)`

GetUseDefaultIpSiteLinkOk returns a tuple with the UseDefaultIpSiteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDefaultIpSiteLink

`func (o *MsserverAdSites) SetUseDefaultIpSiteLink(v bool)`

SetUseDefaultIpSiteLink sets UseDefaultIpSiteLink field to given value.

### HasUseDefaultIpSiteLink

`func (o *MsserverAdSites) HasUseDefaultIpSiteLink() bool`

HasUseDefaultIpSiteLink returns a boolean if a field has been set.

### GetDefaultIpSiteLink

`func (o *MsserverAdSites) GetDefaultIpSiteLink() string`

GetDefaultIpSiteLink returns the DefaultIpSiteLink field if non-nil, zero value otherwise.

### GetDefaultIpSiteLinkOk

`func (o *MsserverAdSites) GetDefaultIpSiteLinkOk() (*string, bool)`

GetDefaultIpSiteLinkOk returns a tuple with the DefaultIpSiteLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIpSiteLink

`func (o *MsserverAdSites) SetDefaultIpSiteLink(v string)`

SetDefaultIpSiteLink sets DefaultIpSiteLink field to given value.

### HasDefaultIpSiteLink

`func (o *MsserverAdSites) HasDefaultIpSiteLink() bool`

HasDefaultIpSiteLink returns a boolean if a field has been set.

### GetUseLogin

`func (o *MsserverAdSites) GetUseLogin() bool`

GetUseLogin returns the UseLogin field if non-nil, zero value otherwise.

### GetUseLoginOk

`func (o *MsserverAdSites) GetUseLoginOk() (*bool, bool)`

GetUseLoginOk returns a tuple with the UseLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLogin

`func (o *MsserverAdSites) SetUseLogin(v bool)`

SetUseLogin sets UseLogin field to given value.

### HasUseLogin

`func (o *MsserverAdSites) HasUseLogin() bool`

HasUseLogin returns a boolean if a field has been set.

### GetLoginName

`func (o *MsserverAdSites) GetLoginName() string`

GetLoginName returns the LoginName field if non-nil, zero value otherwise.

### GetLoginNameOk

`func (o *MsserverAdSites) GetLoginNameOk() (*string, bool)`

GetLoginNameOk returns a tuple with the LoginName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginName

`func (o *MsserverAdSites) SetLoginName(v string)`

SetLoginName sets LoginName field to given value.

### HasLoginName

`func (o *MsserverAdSites) HasLoginName() bool`

HasLoginName returns a boolean if a field has been set.

### GetLoginPassword

`func (o *MsserverAdSites) GetLoginPassword() string`

GetLoginPassword returns the LoginPassword field if non-nil, zero value otherwise.

### GetLoginPasswordOk

`func (o *MsserverAdSites) GetLoginPasswordOk() (*string, bool)`

GetLoginPasswordOk returns a tuple with the LoginPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginPassword

`func (o *MsserverAdSites) SetLoginPassword(v string)`

SetLoginPassword sets LoginPassword field to given value.

### HasLoginPassword

`func (o *MsserverAdSites) HasLoginPassword() bool`

HasLoginPassword returns a boolean if a field has been set.

### GetUseSynchronizationMinDelay

`func (o *MsserverAdSites) GetUseSynchronizationMinDelay() bool`

GetUseSynchronizationMinDelay returns the UseSynchronizationMinDelay field if non-nil, zero value otherwise.

### GetUseSynchronizationMinDelayOk

`func (o *MsserverAdSites) GetUseSynchronizationMinDelayOk() (*bool, bool)`

GetUseSynchronizationMinDelayOk returns a tuple with the UseSynchronizationMinDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSynchronizationMinDelay

`func (o *MsserverAdSites) SetUseSynchronizationMinDelay(v bool)`

SetUseSynchronizationMinDelay sets UseSynchronizationMinDelay field to given value.

### HasUseSynchronizationMinDelay

`func (o *MsserverAdSites) HasUseSynchronizationMinDelay() bool`

HasUseSynchronizationMinDelay returns a boolean if a field has been set.

### GetSynchronizationMinDelay

`func (o *MsserverAdSites) GetSynchronizationMinDelay() int64`

GetSynchronizationMinDelay returns the SynchronizationMinDelay field if non-nil, zero value otherwise.

### GetSynchronizationMinDelayOk

`func (o *MsserverAdSites) GetSynchronizationMinDelayOk() (*int64, bool)`

GetSynchronizationMinDelayOk returns a tuple with the SynchronizationMinDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynchronizationMinDelay

`func (o *MsserverAdSites) SetSynchronizationMinDelay(v int64)`

SetSynchronizationMinDelay sets SynchronizationMinDelay field to given value.

### HasSynchronizationMinDelay

`func (o *MsserverAdSites) HasSynchronizationMinDelay() bool`

HasSynchronizationMinDelay returns a boolean if a field has been set.

### GetUseLdapTimeout

`func (o *MsserverAdSites) GetUseLdapTimeout() bool`

GetUseLdapTimeout returns the UseLdapTimeout field if non-nil, zero value otherwise.

### GetUseLdapTimeoutOk

`func (o *MsserverAdSites) GetUseLdapTimeoutOk() (*bool, bool)`

GetUseLdapTimeoutOk returns a tuple with the UseLdapTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLdapTimeout

`func (o *MsserverAdSites) SetUseLdapTimeout(v bool)`

SetUseLdapTimeout sets UseLdapTimeout field to given value.

### HasUseLdapTimeout

`func (o *MsserverAdSites) HasUseLdapTimeout() bool`

HasUseLdapTimeout returns a boolean if a field has been set.

### GetLdapTimeout

`func (o *MsserverAdSites) GetLdapTimeout() int64`

GetLdapTimeout returns the LdapTimeout field if non-nil, zero value otherwise.

### GetLdapTimeoutOk

`func (o *MsserverAdSites) GetLdapTimeoutOk() (*int64, bool)`

GetLdapTimeoutOk returns a tuple with the LdapTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapTimeout

`func (o *MsserverAdSites) SetLdapTimeout(v int64)`

SetLdapTimeout sets LdapTimeout field to given value.

### HasLdapTimeout

`func (o *MsserverAdSites) HasLdapTimeout() bool`

HasLdapTimeout returns a boolean if a field has been set.

### GetLdapAuthPort

`func (o *MsserverAdSites) GetLdapAuthPort() int64`

GetLdapAuthPort returns the LdapAuthPort field if non-nil, zero value otherwise.

### GetLdapAuthPortOk

`func (o *MsserverAdSites) GetLdapAuthPortOk() (*int64, bool)`

GetLdapAuthPortOk returns a tuple with the LdapAuthPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAuthPort

`func (o *MsserverAdSites) SetLdapAuthPort(v int64)`

SetLdapAuthPort sets LdapAuthPort field to given value.

### HasLdapAuthPort

`func (o *MsserverAdSites) HasLdapAuthPort() bool`

HasLdapAuthPort returns a boolean if a field has been set.

### GetLdapEncryption

`func (o *MsserverAdSites) GetLdapEncryption() string`

GetLdapEncryption returns the LdapEncryption field if non-nil, zero value otherwise.

### GetLdapEncryptionOk

`func (o *MsserverAdSites) GetLdapEncryptionOk() (*string, bool)`

GetLdapEncryptionOk returns a tuple with the LdapEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapEncryption

`func (o *MsserverAdSites) SetLdapEncryption(v string)`

SetLdapEncryption sets LdapEncryption field to given value.

### HasLdapEncryption

`func (o *MsserverAdSites) HasLdapEncryption() bool`

HasLdapEncryption returns a boolean if a field has been set.

### GetManaged

`func (o *MsserverAdSites) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *MsserverAdSites) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *MsserverAdSites) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *MsserverAdSites) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetReadOnly

`func (o *MsserverAdSites) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MsserverAdSites) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *MsserverAdSites) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *MsserverAdSites) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetLastSyncTs

`func (o *MsserverAdSites) GetLastSyncTs() int64`

GetLastSyncTs returns the LastSyncTs field if non-nil, zero value otherwise.

### GetLastSyncTsOk

`func (o *MsserverAdSites) GetLastSyncTsOk() (*int64, bool)`

GetLastSyncTsOk returns a tuple with the LastSyncTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTs

`func (o *MsserverAdSites) SetLastSyncTs(v int64)`

SetLastSyncTs sets LastSyncTs field to given value.

### HasLastSyncTs

`func (o *MsserverAdSites) HasLastSyncTs() bool`

HasLastSyncTs returns a boolean if a field has been set.

### GetLastSyncStatus

`func (o *MsserverAdSites) GetLastSyncStatus() string`

GetLastSyncStatus returns the LastSyncStatus field if non-nil, zero value otherwise.

### GetLastSyncStatusOk

`func (o *MsserverAdSites) GetLastSyncStatusOk() (*string, bool)`

GetLastSyncStatusOk returns a tuple with the LastSyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncStatus

`func (o *MsserverAdSites) SetLastSyncStatus(v string)`

SetLastSyncStatus sets LastSyncStatus field to given value.

### HasLastSyncStatus

`func (o *MsserverAdSites) HasLastSyncStatus() bool`

HasLastSyncStatus returns a boolean if a field has been set.

### GetLastSyncDetail

`func (o *MsserverAdSites) GetLastSyncDetail() string`

GetLastSyncDetail returns the LastSyncDetail field if non-nil, zero value otherwise.

### GetLastSyncDetailOk

`func (o *MsserverAdSites) GetLastSyncDetailOk() (*string, bool)`

GetLastSyncDetailOk returns a tuple with the LastSyncDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDetail

`func (o *MsserverAdSites) SetLastSyncDetail(v string)`

SetLastSyncDetail sets LastSyncDetail field to given value.

### HasLastSyncDetail

`func (o *MsserverAdSites) HasLastSyncDetail() bool`

HasLastSyncDetail returns a boolean if a field has been set.

### GetSupportsIpv6

`func (o *MsserverAdSites) GetSupportsIpv6() bool`

GetSupportsIpv6 returns the SupportsIpv6 field if non-nil, zero value otherwise.

### GetSupportsIpv6Ok

`func (o *MsserverAdSites) GetSupportsIpv6Ok() (*bool, bool)`

GetSupportsIpv6Ok returns a tuple with the SupportsIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsIpv6

`func (o *MsserverAdSites) SetSupportsIpv6(v bool)`

SetSupportsIpv6 sets SupportsIpv6 field to given value.

### HasSupportsIpv6

`func (o *MsserverAdSites) HasSupportsIpv6() bool`

HasSupportsIpv6 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


