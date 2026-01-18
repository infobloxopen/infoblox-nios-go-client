# Admingroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AccessMethod** | Pointer to **[]string** | Access methods specify whether an admin group can use the GUI and the API to access the appliance or to send Taxii messages to the appliance. Note that API includes both the Perl API and RESTful API. | [optional] 
**AdminSetCommands** | Pointer to [**AdmingroupAdminSetCommands**](AdmingroupAdminSetCommands.md) |  | [optional] 
**AdminShowCommands** | Pointer to [**AdmingroupAdminShowCommands**](AdmingroupAdminShowCommands.md) |  | [optional] 
**AdminToplevelCommands** | Pointer to [**AdmingroupAdminToplevelCommands**](AdmingroupAdminToplevelCommands.md) |  | [optional] 
**CloudSetCommands** | Pointer to [**AdmingroupCloudSetCommands**](AdmingroupCloudSetCommands.md) |  | [optional] 
**CloudShowCommands** | Pointer to [**AdmingroupCloudShowCommands**](AdmingroupCloudShowCommands.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the Admin Group; maximum 256 characters. | [optional] 
**DatabaseSetCommands** | Pointer to [**AdmingroupDatabaseSetCommands**](AdmingroupDatabaseSetCommands.md) |  | [optional] 
**DatabaseShowCommands** | Pointer to [**AdmingroupDatabaseShowCommands**](AdmingroupDatabaseShowCommands.md) |  | [optional] 
**DhcpSetCommands** | Pointer to [**AdmingroupDhcpSetCommands**](AdmingroupDhcpSetCommands.md) |  | [optional] 
**DhcpShowCommands** | Pointer to [**AdmingroupDhcpShowCommands**](AdmingroupDhcpShowCommands.md) |  | [optional] 
**Disable** | Pointer to **bool** | Determines whether the Admin Group is disabled or not. When this is set to False, the Admin Group is enabled. | [optional] 
**DisableConcurrentLogin** | Pointer to **bool** | Disable concurrent login feature | [optional] 
**DnsSetCommands** | Pointer to [**AdmingroupDnsSetCommands**](AdmingroupDnsSetCommands.md) |  | [optional] 
**DnsShowCommands** | Pointer to [**AdmingroupDnsShowCommands**](AdmingroupDnsShowCommands.md) |  | [optional] 
**DnsToplevelCommands** | Pointer to [**AdmingroupDnsToplevelCommands**](AdmingroupDnsToplevelCommands.md) |  | [optional] 
**DockerSetCommands** | Pointer to [**AdmingroupDockerSetCommands**](AdmingroupDockerSetCommands.md) |  | [optional] 
**DockerShowCommands** | Pointer to [**AdmingroupDockerShowCommands**](AdmingroupDockerShowCommands.md) |  | [optional] 
**EmailAddresses** | Pointer to **[]string** | The e-mail addresses for the Admin Group. | [optional] 
**EnableRestrictedUserAccess** | Pointer to **bool** | Determines whether the restrictions will be applied to the admin connector level for users of this Admin Group. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**GridSetCommands** | Pointer to [**AdmingroupGridSetCommands**](AdmingroupGridSetCommands.md) |  | [optional] 
**GridShowCommands** | Pointer to [**AdmingroupGridShowCommands**](AdmingroupGridShowCommands.md) |  | [optional] 
**InactivityLockoutSetting** | Pointer to [**AdmingroupInactivityLockoutSetting**](AdmingroupInactivityLockoutSetting.md) |  | [optional] 
**LicensingSetCommands** | Pointer to [**AdmingroupLicensingSetCommands**](AdmingroupLicensingSetCommands.md) |  | [optional] 
**LicensingShowCommands** | Pointer to [**AdmingroupLicensingShowCommands**](AdmingroupLicensingShowCommands.md) |  | [optional] 
**LockoutSetting** | Pointer to [**AdmingroupLockoutSetting**](AdmingroupLockoutSetting.md) |  | [optional] 
**MachineControlToplevelCommands** | Pointer to [**AdmingroupMachineControlToplevelCommands**](AdmingroupMachineControlToplevelCommands.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the Admin Group. | [optional] 
**NetworkingSetCommands** | Pointer to [**AdmingroupNetworkingSetCommands**](AdmingroupNetworkingSetCommands.md) |  | [optional] 
**NetworkingShowCommands** | Pointer to [**AdmingroupNetworkingShowCommands**](AdmingroupNetworkingShowCommands.md) |  | [optional] 
**PasswordSetting** | Pointer to [**AdmingroupPasswordSetting**](AdmingroupPasswordSetting.md) |  | [optional] 
**Roles** | Pointer to **[]string** | The names of roles this Admin Group applies to. | [optional] 
**SamlSetting** | Pointer to [**AdmingroupSamlSetting**](AdmingroupSamlSetting.md) |  | [optional] 
**SecuritySetCommands** | Pointer to [**AdmingroupSecuritySetCommands**](AdmingroupSecuritySetCommands.md) |  | [optional] 
**SecurityShowCommands** | Pointer to [**AdmingroupSecurityShowCommands**](AdmingroupSecurityShowCommands.md) |  | [optional] 
**Superuser** | Pointer to **bool** | Determines whether this Admin Group is a superuser group. A superuser group can perform all operations on the appliance, and can view and configure all types of data. | [optional] 
**TroubleShootingToplevelCommands** | Pointer to [**AdmingroupTroubleShootingToplevelCommands**](AdmingroupTroubleShootingToplevelCommands.md) |  | [optional] 
**UseAccountInactivityLockoutEnable** | Pointer to **bool** | This is the use flag for account inactivity lockout settings. | [optional] 
**UseDisableConcurrentLogin** | Pointer to **bool** | Whether to override grid concurrent login | [optional] 
**UseLockoutSetting** | Pointer to **bool** | Whether to override grid sequential lockout setting | [optional] 
**UsePasswordSetting** | Pointer to **bool** | Whether grid password expiry setting should be override. | [optional] 
**UserAccess** | Pointer to [**[]AdmingroupUserAccess**](AdmingroupUserAccess.md) | The access control items for this Admin Group. | [optional] 

## Methods

### NewAdmingroup

`func NewAdmingroup() *Admingroup`

NewAdmingroup instantiates a new Admingroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdmingroupWithDefaults

`func NewAdmingroupWithDefaults() *Admingroup`

NewAdmingroupWithDefaults instantiates a new Admingroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Admingroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Admingroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Admingroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Admingroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAccessMethod

`func (o *Admingroup) GetAccessMethod() []string`

GetAccessMethod returns the AccessMethod field if non-nil, zero value otherwise.

### GetAccessMethodOk

`func (o *Admingroup) GetAccessMethodOk() (*[]string, bool)`

GetAccessMethodOk returns a tuple with the AccessMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMethod

`func (o *Admingroup) SetAccessMethod(v []string)`

SetAccessMethod sets AccessMethod field to given value.

### HasAccessMethod

`func (o *Admingroup) HasAccessMethod() bool`

HasAccessMethod returns a boolean if a field has been set.

### GetAdminSetCommands

`func (o *Admingroup) GetAdminSetCommands() AdmingroupAdminSetCommands`

GetAdminSetCommands returns the AdminSetCommands field if non-nil, zero value otherwise.

### GetAdminSetCommandsOk

`func (o *Admingroup) GetAdminSetCommandsOk() (*AdmingroupAdminSetCommands, bool)`

GetAdminSetCommandsOk returns a tuple with the AdminSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSetCommands

`func (o *Admingroup) SetAdminSetCommands(v AdmingroupAdminSetCommands)`

SetAdminSetCommands sets AdminSetCommands field to given value.

### HasAdminSetCommands

`func (o *Admingroup) HasAdminSetCommands() bool`

HasAdminSetCommands returns a boolean if a field has been set.

### GetAdminShowCommands

`func (o *Admingroup) GetAdminShowCommands() AdmingroupAdminShowCommands`

GetAdminShowCommands returns the AdminShowCommands field if non-nil, zero value otherwise.

### GetAdminShowCommandsOk

`func (o *Admingroup) GetAdminShowCommandsOk() (*AdmingroupAdminShowCommands, bool)`

GetAdminShowCommandsOk returns a tuple with the AdminShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminShowCommands

`func (o *Admingroup) SetAdminShowCommands(v AdmingroupAdminShowCommands)`

SetAdminShowCommands sets AdminShowCommands field to given value.

### HasAdminShowCommands

`func (o *Admingroup) HasAdminShowCommands() bool`

HasAdminShowCommands returns a boolean if a field has been set.

### GetAdminToplevelCommands

`func (o *Admingroup) GetAdminToplevelCommands() AdmingroupAdminToplevelCommands`

GetAdminToplevelCommands returns the AdminToplevelCommands field if non-nil, zero value otherwise.

### GetAdminToplevelCommandsOk

`func (o *Admingroup) GetAdminToplevelCommandsOk() (*AdmingroupAdminToplevelCommands, bool)`

GetAdminToplevelCommandsOk returns a tuple with the AdminToplevelCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminToplevelCommands

`func (o *Admingroup) SetAdminToplevelCommands(v AdmingroupAdminToplevelCommands)`

SetAdminToplevelCommands sets AdminToplevelCommands field to given value.

### HasAdminToplevelCommands

`func (o *Admingroup) HasAdminToplevelCommands() bool`

HasAdminToplevelCommands returns a boolean if a field has been set.

### GetCloudSetCommands

`func (o *Admingroup) GetCloudSetCommands() AdmingroupCloudSetCommands`

GetCloudSetCommands returns the CloudSetCommands field if non-nil, zero value otherwise.

### GetCloudSetCommandsOk

`func (o *Admingroup) GetCloudSetCommandsOk() (*AdmingroupCloudSetCommands, bool)`

GetCloudSetCommandsOk returns a tuple with the CloudSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSetCommands

`func (o *Admingroup) SetCloudSetCommands(v AdmingroupCloudSetCommands)`

SetCloudSetCommands sets CloudSetCommands field to given value.

### HasCloudSetCommands

`func (o *Admingroup) HasCloudSetCommands() bool`

HasCloudSetCommands returns a boolean if a field has been set.

### GetCloudShowCommands

`func (o *Admingroup) GetCloudShowCommands() AdmingroupCloudShowCommands`

GetCloudShowCommands returns the CloudShowCommands field if non-nil, zero value otherwise.

### GetCloudShowCommandsOk

`func (o *Admingroup) GetCloudShowCommandsOk() (*AdmingroupCloudShowCommands, bool)`

GetCloudShowCommandsOk returns a tuple with the CloudShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudShowCommands

`func (o *Admingroup) SetCloudShowCommands(v AdmingroupCloudShowCommands)`

SetCloudShowCommands sets CloudShowCommands field to given value.

### HasCloudShowCommands

`func (o *Admingroup) HasCloudShowCommands() bool`

HasCloudShowCommands returns a boolean if a field has been set.

### GetComment

`func (o *Admingroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Admingroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Admingroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Admingroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDatabaseSetCommands

`func (o *Admingroup) GetDatabaseSetCommands() AdmingroupDatabaseSetCommands`

GetDatabaseSetCommands returns the DatabaseSetCommands field if non-nil, zero value otherwise.

### GetDatabaseSetCommandsOk

`func (o *Admingroup) GetDatabaseSetCommandsOk() (*AdmingroupDatabaseSetCommands, bool)`

GetDatabaseSetCommandsOk returns a tuple with the DatabaseSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSetCommands

`func (o *Admingroup) SetDatabaseSetCommands(v AdmingroupDatabaseSetCommands)`

SetDatabaseSetCommands sets DatabaseSetCommands field to given value.

### HasDatabaseSetCommands

`func (o *Admingroup) HasDatabaseSetCommands() bool`

HasDatabaseSetCommands returns a boolean if a field has been set.

### GetDatabaseShowCommands

`func (o *Admingroup) GetDatabaseShowCommands() AdmingroupDatabaseShowCommands`

GetDatabaseShowCommands returns the DatabaseShowCommands field if non-nil, zero value otherwise.

### GetDatabaseShowCommandsOk

`func (o *Admingroup) GetDatabaseShowCommandsOk() (*AdmingroupDatabaseShowCommands, bool)`

GetDatabaseShowCommandsOk returns a tuple with the DatabaseShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseShowCommands

`func (o *Admingroup) SetDatabaseShowCommands(v AdmingroupDatabaseShowCommands)`

SetDatabaseShowCommands sets DatabaseShowCommands field to given value.

### HasDatabaseShowCommands

`func (o *Admingroup) HasDatabaseShowCommands() bool`

HasDatabaseShowCommands returns a boolean if a field has been set.

### GetDhcpSetCommands

`func (o *Admingroup) GetDhcpSetCommands() AdmingroupDhcpSetCommands`

GetDhcpSetCommands returns the DhcpSetCommands field if non-nil, zero value otherwise.

### GetDhcpSetCommandsOk

`func (o *Admingroup) GetDhcpSetCommandsOk() (*AdmingroupDhcpSetCommands, bool)`

GetDhcpSetCommandsOk returns a tuple with the DhcpSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSetCommands

`func (o *Admingroup) SetDhcpSetCommands(v AdmingroupDhcpSetCommands)`

SetDhcpSetCommands sets DhcpSetCommands field to given value.

### HasDhcpSetCommands

`func (o *Admingroup) HasDhcpSetCommands() bool`

HasDhcpSetCommands returns a boolean if a field has been set.

### GetDhcpShowCommands

`func (o *Admingroup) GetDhcpShowCommands() AdmingroupDhcpShowCommands`

GetDhcpShowCommands returns the DhcpShowCommands field if non-nil, zero value otherwise.

### GetDhcpShowCommandsOk

`func (o *Admingroup) GetDhcpShowCommandsOk() (*AdmingroupDhcpShowCommands, bool)`

GetDhcpShowCommandsOk returns a tuple with the DhcpShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpShowCommands

`func (o *Admingroup) SetDhcpShowCommands(v AdmingroupDhcpShowCommands)`

SetDhcpShowCommands sets DhcpShowCommands field to given value.

### HasDhcpShowCommands

`func (o *Admingroup) HasDhcpShowCommands() bool`

HasDhcpShowCommands returns a boolean if a field has been set.

### GetDisable

`func (o *Admingroup) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *Admingroup) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *Admingroup) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *Admingroup) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisableConcurrentLogin

`func (o *Admingroup) GetDisableConcurrentLogin() bool`

GetDisableConcurrentLogin returns the DisableConcurrentLogin field if non-nil, zero value otherwise.

### GetDisableConcurrentLoginOk

`func (o *Admingroup) GetDisableConcurrentLoginOk() (*bool, bool)`

GetDisableConcurrentLoginOk returns a tuple with the DisableConcurrentLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConcurrentLogin

`func (o *Admingroup) SetDisableConcurrentLogin(v bool)`

SetDisableConcurrentLogin sets DisableConcurrentLogin field to given value.

### HasDisableConcurrentLogin

`func (o *Admingroup) HasDisableConcurrentLogin() bool`

HasDisableConcurrentLogin returns a boolean if a field has been set.

### GetDnsSetCommands

`func (o *Admingroup) GetDnsSetCommands() AdmingroupDnsSetCommands`

GetDnsSetCommands returns the DnsSetCommands field if non-nil, zero value otherwise.

### GetDnsSetCommandsOk

`func (o *Admingroup) GetDnsSetCommandsOk() (*AdmingroupDnsSetCommands, bool)`

GetDnsSetCommandsOk returns a tuple with the DnsSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSetCommands

`func (o *Admingroup) SetDnsSetCommands(v AdmingroupDnsSetCommands)`

SetDnsSetCommands sets DnsSetCommands field to given value.

### HasDnsSetCommands

`func (o *Admingroup) HasDnsSetCommands() bool`

HasDnsSetCommands returns a boolean if a field has been set.

### GetDnsShowCommands

`func (o *Admingroup) GetDnsShowCommands() AdmingroupDnsShowCommands`

GetDnsShowCommands returns the DnsShowCommands field if non-nil, zero value otherwise.

### GetDnsShowCommandsOk

`func (o *Admingroup) GetDnsShowCommandsOk() (*AdmingroupDnsShowCommands, bool)`

GetDnsShowCommandsOk returns a tuple with the DnsShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsShowCommands

`func (o *Admingroup) SetDnsShowCommands(v AdmingroupDnsShowCommands)`

SetDnsShowCommands sets DnsShowCommands field to given value.

### HasDnsShowCommands

`func (o *Admingroup) HasDnsShowCommands() bool`

HasDnsShowCommands returns a boolean if a field has been set.

### GetDnsToplevelCommands

`func (o *Admingroup) GetDnsToplevelCommands() AdmingroupDnsToplevelCommands`

GetDnsToplevelCommands returns the DnsToplevelCommands field if non-nil, zero value otherwise.

### GetDnsToplevelCommandsOk

`func (o *Admingroup) GetDnsToplevelCommandsOk() (*AdmingroupDnsToplevelCommands, bool)`

GetDnsToplevelCommandsOk returns a tuple with the DnsToplevelCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsToplevelCommands

`func (o *Admingroup) SetDnsToplevelCommands(v AdmingroupDnsToplevelCommands)`

SetDnsToplevelCommands sets DnsToplevelCommands field to given value.

### HasDnsToplevelCommands

`func (o *Admingroup) HasDnsToplevelCommands() bool`

HasDnsToplevelCommands returns a boolean if a field has been set.

### GetDockerSetCommands

`func (o *Admingroup) GetDockerSetCommands() AdmingroupDockerSetCommands`

GetDockerSetCommands returns the DockerSetCommands field if non-nil, zero value otherwise.

### GetDockerSetCommandsOk

`func (o *Admingroup) GetDockerSetCommandsOk() (*AdmingroupDockerSetCommands, bool)`

GetDockerSetCommandsOk returns a tuple with the DockerSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerSetCommands

`func (o *Admingroup) SetDockerSetCommands(v AdmingroupDockerSetCommands)`

SetDockerSetCommands sets DockerSetCommands field to given value.

### HasDockerSetCommands

`func (o *Admingroup) HasDockerSetCommands() bool`

HasDockerSetCommands returns a boolean if a field has been set.

### GetDockerShowCommands

`func (o *Admingroup) GetDockerShowCommands() AdmingroupDockerShowCommands`

GetDockerShowCommands returns the DockerShowCommands field if non-nil, zero value otherwise.

### GetDockerShowCommandsOk

`func (o *Admingroup) GetDockerShowCommandsOk() (*AdmingroupDockerShowCommands, bool)`

GetDockerShowCommandsOk returns a tuple with the DockerShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerShowCommands

`func (o *Admingroup) SetDockerShowCommands(v AdmingroupDockerShowCommands)`

SetDockerShowCommands sets DockerShowCommands field to given value.

### HasDockerShowCommands

`func (o *Admingroup) HasDockerShowCommands() bool`

HasDockerShowCommands returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *Admingroup) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *Admingroup) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *Admingroup) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *Admingroup) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetEnableRestrictedUserAccess

`func (o *Admingroup) GetEnableRestrictedUserAccess() bool`

GetEnableRestrictedUserAccess returns the EnableRestrictedUserAccess field if non-nil, zero value otherwise.

### GetEnableRestrictedUserAccessOk

`func (o *Admingroup) GetEnableRestrictedUserAccessOk() (*bool, bool)`

GetEnableRestrictedUserAccessOk returns a tuple with the EnableRestrictedUserAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRestrictedUserAccess

`func (o *Admingroup) SetEnableRestrictedUserAccess(v bool)`

SetEnableRestrictedUserAccess sets EnableRestrictedUserAccess field to given value.

### HasEnableRestrictedUserAccess

`func (o *Admingroup) HasEnableRestrictedUserAccess() bool`

HasEnableRestrictedUserAccess returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Admingroup) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Admingroup) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Admingroup) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Admingroup) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Admingroup) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Admingroup) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Admingroup) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Admingroup) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Admingroup) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Admingroup) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Admingroup) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Admingroup) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetGridSetCommands

`func (o *Admingroup) GetGridSetCommands() AdmingroupGridSetCommands`

GetGridSetCommands returns the GridSetCommands field if non-nil, zero value otherwise.

### GetGridSetCommandsOk

`func (o *Admingroup) GetGridSetCommandsOk() (*AdmingroupGridSetCommands, bool)`

GetGridSetCommandsOk returns a tuple with the GridSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridSetCommands

`func (o *Admingroup) SetGridSetCommands(v AdmingroupGridSetCommands)`

SetGridSetCommands sets GridSetCommands field to given value.

### HasGridSetCommands

`func (o *Admingroup) HasGridSetCommands() bool`

HasGridSetCommands returns a boolean if a field has been set.

### GetGridShowCommands

`func (o *Admingroup) GetGridShowCommands() AdmingroupGridShowCommands`

GetGridShowCommands returns the GridShowCommands field if non-nil, zero value otherwise.

### GetGridShowCommandsOk

`func (o *Admingroup) GetGridShowCommandsOk() (*AdmingroupGridShowCommands, bool)`

GetGridShowCommandsOk returns a tuple with the GridShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridShowCommands

`func (o *Admingroup) SetGridShowCommands(v AdmingroupGridShowCommands)`

SetGridShowCommands sets GridShowCommands field to given value.

### HasGridShowCommands

`func (o *Admingroup) HasGridShowCommands() bool`

HasGridShowCommands returns a boolean if a field has been set.

### GetInactivityLockoutSetting

`func (o *Admingroup) GetInactivityLockoutSetting() AdmingroupInactivityLockoutSetting`

GetInactivityLockoutSetting returns the InactivityLockoutSetting field if non-nil, zero value otherwise.

### GetInactivityLockoutSettingOk

`func (o *Admingroup) GetInactivityLockoutSettingOk() (*AdmingroupInactivityLockoutSetting, bool)`

GetInactivityLockoutSettingOk returns a tuple with the InactivityLockoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityLockoutSetting

`func (o *Admingroup) SetInactivityLockoutSetting(v AdmingroupInactivityLockoutSetting)`

SetInactivityLockoutSetting sets InactivityLockoutSetting field to given value.

### HasInactivityLockoutSetting

`func (o *Admingroup) HasInactivityLockoutSetting() bool`

HasInactivityLockoutSetting returns a boolean if a field has been set.

### GetLicensingSetCommands

`func (o *Admingroup) GetLicensingSetCommands() AdmingroupLicensingSetCommands`

GetLicensingSetCommands returns the LicensingSetCommands field if non-nil, zero value otherwise.

### GetLicensingSetCommandsOk

`func (o *Admingroup) GetLicensingSetCommandsOk() (*AdmingroupLicensingSetCommands, bool)`

GetLicensingSetCommandsOk returns a tuple with the LicensingSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingSetCommands

`func (o *Admingroup) SetLicensingSetCommands(v AdmingroupLicensingSetCommands)`

SetLicensingSetCommands sets LicensingSetCommands field to given value.

### HasLicensingSetCommands

`func (o *Admingroup) HasLicensingSetCommands() bool`

HasLicensingSetCommands returns a boolean if a field has been set.

### GetLicensingShowCommands

`func (o *Admingroup) GetLicensingShowCommands() AdmingroupLicensingShowCommands`

GetLicensingShowCommands returns the LicensingShowCommands field if non-nil, zero value otherwise.

### GetLicensingShowCommandsOk

`func (o *Admingroup) GetLicensingShowCommandsOk() (*AdmingroupLicensingShowCommands, bool)`

GetLicensingShowCommandsOk returns a tuple with the LicensingShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingShowCommands

`func (o *Admingroup) SetLicensingShowCommands(v AdmingroupLicensingShowCommands)`

SetLicensingShowCommands sets LicensingShowCommands field to given value.

### HasLicensingShowCommands

`func (o *Admingroup) HasLicensingShowCommands() bool`

HasLicensingShowCommands returns a boolean if a field has been set.

### GetLockoutSetting

`func (o *Admingroup) GetLockoutSetting() AdmingroupLockoutSetting`

GetLockoutSetting returns the LockoutSetting field if non-nil, zero value otherwise.

### GetLockoutSettingOk

`func (o *Admingroup) GetLockoutSettingOk() (*AdmingroupLockoutSetting, bool)`

GetLockoutSettingOk returns a tuple with the LockoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutSetting

`func (o *Admingroup) SetLockoutSetting(v AdmingroupLockoutSetting)`

SetLockoutSetting sets LockoutSetting field to given value.

### HasLockoutSetting

`func (o *Admingroup) HasLockoutSetting() bool`

HasLockoutSetting returns a boolean if a field has been set.

### GetMachineControlToplevelCommands

`func (o *Admingroup) GetMachineControlToplevelCommands() AdmingroupMachineControlToplevelCommands`

GetMachineControlToplevelCommands returns the MachineControlToplevelCommands field if non-nil, zero value otherwise.

### GetMachineControlToplevelCommandsOk

`func (o *Admingroup) GetMachineControlToplevelCommandsOk() (*AdmingroupMachineControlToplevelCommands, bool)`

GetMachineControlToplevelCommandsOk returns a tuple with the MachineControlToplevelCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineControlToplevelCommands

`func (o *Admingroup) SetMachineControlToplevelCommands(v AdmingroupMachineControlToplevelCommands)`

SetMachineControlToplevelCommands sets MachineControlToplevelCommands field to given value.

### HasMachineControlToplevelCommands

`func (o *Admingroup) HasMachineControlToplevelCommands() bool`

HasMachineControlToplevelCommands returns a boolean if a field has been set.

### GetName

`func (o *Admingroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Admingroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Admingroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Admingroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkingSetCommands

`func (o *Admingroup) GetNetworkingSetCommands() AdmingroupNetworkingSetCommands`

GetNetworkingSetCommands returns the NetworkingSetCommands field if non-nil, zero value otherwise.

### GetNetworkingSetCommandsOk

`func (o *Admingroup) GetNetworkingSetCommandsOk() (*AdmingroupNetworkingSetCommands, bool)`

GetNetworkingSetCommandsOk returns a tuple with the NetworkingSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingSetCommands

`func (o *Admingroup) SetNetworkingSetCommands(v AdmingroupNetworkingSetCommands)`

SetNetworkingSetCommands sets NetworkingSetCommands field to given value.

### HasNetworkingSetCommands

`func (o *Admingroup) HasNetworkingSetCommands() bool`

HasNetworkingSetCommands returns a boolean if a field has been set.

### GetNetworkingShowCommands

`func (o *Admingroup) GetNetworkingShowCommands() AdmingroupNetworkingShowCommands`

GetNetworkingShowCommands returns the NetworkingShowCommands field if non-nil, zero value otherwise.

### GetNetworkingShowCommandsOk

`func (o *Admingroup) GetNetworkingShowCommandsOk() (*AdmingroupNetworkingShowCommands, bool)`

GetNetworkingShowCommandsOk returns a tuple with the NetworkingShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingShowCommands

`func (o *Admingroup) SetNetworkingShowCommands(v AdmingroupNetworkingShowCommands)`

SetNetworkingShowCommands sets NetworkingShowCommands field to given value.

### HasNetworkingShowCommands

`func (o *Admingroup) HasNetworkingShowCommands() bool`

HasNetworkingShowCommands returns a boolean if a field has been set.

### GetPasswordSetting

`func (o *Admingroup) GetPasswordSetting() AdmingroupPasswordSetting`

GetPasswordSetting returns the PasswordSetting field if non-nil, zero value otherwise.

### GetPasswordSettingOk

`func (o *Admingroup) GetPasswordSettingOk() (*AdmingroupPasswordSetting, bool)`

GetPasswordSettingOk returns a tuple with the PasswordSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSetting

`func (o *Admingroup) SetPasswordSetting(v AdmingroupPasswordSetting)`

SetPasswordSetting sets PasswordSetting field to given value.

### HasPasswordSetting

`func (o *Admingroup) HasPasswordSetting() bool`

HasPasswordSetting returns a boolean if a field has been set.

### GetRoles

`func (o *Admingroup) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *Admingroup) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *Admingroup) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *Admingroup) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSamlSetting

`func (o *Admingroup) GetSamlSetting() AdmingroupSamlSetting`

GetSamlSetting returns the SamlSetting field if non-nil, zero value otherwise.

### GetSamlSettingOk

`func (o *Admingroup) GetSamlSettingOk() (*AdmingroupSamlSetting, bool)`

GetSamlSettingOk returns a tuple with the SamlSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSetting

`func (o *Admingroup) SetSamlSetting(v AdmingroupSamlSetting)`

SetSamlSetting sets SamlSetting field to given value.

### HasSamlSetting

`func (o *Admingroup) HasSamlSetting() bool`

HasSamlSetting returns a boolean if a field has been set.

### GetSecuritySetCommands

`func (o *Admingroup) GetSecuritySetCommands() AdmingroupSecuritySetCommands`

GetSecuritySetCommands returns the SecuritySetCommands field if non-nil, zero value otherwise.

### GetSecuritySetCommandsOk

`func (o *Admingroup) GetSecuritySetCommandsOk() (*AdmingroupSecuritySetCommands, bool)`

GetSecuritySetCommandsOk returns a tuple with the SecuritySetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySetCommands

`func (o *Admingroup) SetSecuritySetCommands(v AdmingroupSecuritySetCommands)`

SetSecuritySetCommands sets SecuritySetCommands field to given value.

### HasSecuritySetCommands

`func (o *Admingroup) HasSecuritySetCommands() bool`

HasSecuritySetCommands returns a boolean if a field has been set.

### GetSecurityShowCommands

`func (o *Admingroup) GetSecurityShowCommands() AdmingroupSecurityShowCommands`

GetSecurityShowCommands returns the SecurityShowCommands field if non-nil, zero value otherwise.

### GetSecurityShowCommandsOk

`func (o *Admingroup) GetSecurityShowCommandsOk() (*AdmingroupSecurityShowCommands, bool)`

GetSecurityShowCommandsOk returns a tuple with the SecurityShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityShowCommands

`func (o *Admingroup) SetSecurityShowCommands(v AdmingroupSecurityShowCommands)`

SetSecurityShowCommands sets SecurityShowCommands field to given value.

### HasSecurityShowCommands

`func (o *Admingroup) HasSecurityShowCommands() bool`

HasSecurityShowCommands returns a boolean if a field has been set.

### GetSuperuser

`func (o *Admingroup) GetSuperuser() bool`

GetSuperuser returns the Superuser field if non-nil, zero value otherwise.

### GetSuperuserOk

`func (o *Admingroup) GetSuperuserOk() (*bool, bool)`

GetSuperuserOk returns a tuple with the Superuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperuser

`func (o *Admingroup) SetSuperuser(v bool)`

SetSuperuser sets Superuser field to given value.

### HasSuperuser

`func (o *Admingroup) HasSuperuser() bool`

HasSuperuser returns a boolean if a field has been set.

### GetTroubleShootingToplevelCommands

`func (o *Admingroup) GetTroubleShootingToplevelCommands() AdmingroupTroubleShootingToplevelCommands`

GetTroubleShootingToplevelCommands returns the TroubleShootingToplevelCommands field if non-nil, zero value otherwise.

### GetTroubleShootingToplevelCommandsOk

`func (o *Admingroup) GetTroubleShootingToplevelCommandsOk() (*AdmingroupTroubleShootingToplevelCommands, bool)`

GetTroubleShootingToplevelCommandsOk returns a tuple with the TroubleShootingToplevelCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroubleShootingToplevelCommands

`func (o *Admingroup) SetTroubleShootingToplevelCommands(v AdmingroupTroubleShootingToplevelCommands)`

SetTroubleShootingToplevelCommands sets TroubleShootingToplevelCommands field to given value.

### HasTroubleShootingToplevelCommands

`func (o *Admingroup) HasTroubleShootingToplevelCommands() bool`

HasTroubleShootingToplevelCommands returns a boolean if a field has been set.

### GetUseAccountInactivityLockoutEnable

`func (o *Admingroup) GetUseAccountInactivityLockoutEnable() bool`

GetUseAccountInactivityLockoutEnable returns the UseAccountInactivityLockoutEnable field if non-nil, zero value otherwise.

### GetUseAccountInactivityLockoutEnableOk

`func (o *Admingroup) GetUseAccountInactivityLockoutEnableOk() (*bool, bool)`

GetUseAccountInactivityLockoutEnableOk returns a tuple with the UseAccountInactivityLockoutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAccountInactivityLockoutEnable

`func (o *Admingroup) SetUseAccountInactivityLockoutEnable(v bool)`

SetUseAccountInactivityLockoutEnable sets UseAccountInactivityLockoutEnable field to given value.

### HasUseAccountInactivityLockoutEnable

`func (o *Admingroup) HasUseAccountInactivityLockoutEnable() bool`

HasUseAccountInactivityLockoutEnable returns a boolean if a field has been set.

### GetUseDisableConcurrentLogin

`func (o *Admingroup) GetUseDisableConcurrentLogin() bool`

GetUseDisableConcurrentLogin returns the UseDisableConcurrentLogin field if non-nil, zero value otherwise.

### GetUseDisableConcurrentLoginOk

`func (o *Admingroup) GetUseDisableConcurrentLoginOk() (*bool, bool)`

GetUseDisableConcurrentLoginOk returns a tuple with the UseDisableConcurrentLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDisableConcurrentLogin

`func (o *Admingroup) SetUseDisableConcurrentLogin(v bool)`

SetUseDisableConcurrentLogin sets UseDisableConcurrentLogin field to given value.

### HasUseDisableConcurrentLogin

`func (o *Admingroup) HasUseDisableConcurrentLogin() bool`

HasUseDisableConcurrentLogin returns a boolean if a field has been set.

### GetUseLockoutSetting

`func (o *Admingroup) GetUseLockoutSetting() bool`

GetUseLockoutSetting returns the UseLockoutSetting field if non-nil, zero value otherwise.

### GetUseLockoutSettingOk

`func (o *Admingroup) GetUseLockoutSettingOk() (*bool, bool)`

GetUseLockoutSettingOk returns a tuple with the UseLockoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLockoutSetting

`func (o *Admingroup) SetUseLockoutSetting(v bool)`

SetUseLockoutSetting sets UseLockoutSetting field to given value.

### HasUseLockoutSetting

`func (o *Admingroup) HasUseLockoutSetting() bool`

HasUseLockoutSetting returns a boolean if a field has been set.

### GetUsePasswordSetting

`func (o *Admingroup) GetUsePasswordSetting() bool`

GetUsePasswordSetting returns the UsePasswordSetting field if non-nil, zero value otherwise.

### GetUsePasswordSettingOk

`func (o *Admingroup) GetUsePasswordSettingOk() (*bool, bool)`

GetUsePasswordSettingOk returns a tuple with the UsePasswordSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePasswordSetting

`func (o *Admingroup) SetUsePasswordSetting(v bool)`

SetUsePasswordSetting sets UsePasswordSetting field to given value.

### HasUsePasswordSetting

`func (o *Admingroup) HasUsePasswordSetting() bool`

HasUsePasswordSetting returns a boolean if a field has been set.

### GetUserAccess

`func (o *Admingroup) GetUserAccess() []AdmingroupUserAccess`

GetUserAccess returns the UserAccess field if non-nil, zero value otherwise.

### GetUserAccessOk

`func (o *Admingroup) GetUserAccessOk() (*[]AdmingroupUserAccess, bool)`

GetUserAccessOk returns a tuple with the UserAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccess

`func (o *Admingroup) SetUserAccess(v []AdmingroupUserAccess)`

SetUserAccess sets UserAccess field to given value.

### HasUserAccess

`func (o *Admingroup) HasUserAccess() bool`

HasUserAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


