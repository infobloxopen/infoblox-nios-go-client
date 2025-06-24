# GetAdmingroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**Result** | Pointer to [**Admingroup**](Admingroup.md) |  | [optional] 

## Methods

### NewGetAdmingroupResponse

`func NewGetAdmingroupResponse() *GetAdmingroupResponse`

NewGetAdmingroupResponse instantiates a new GetAdmingroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdmingroupResponseWithDefaults

`func NewGetAdmingroupResponseWithDefaults() *GetAdmingroupResponse`

NewGetAdmingroupResponseWithDefaults instantiates a new GetAdmingroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAdmingroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAdmingroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAdmingroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAdmingroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAccessMethod

`func (o *GetAdmingroupResponse) GetAccessMethod() []string`

GetAccessMethod returns the AccessMethod field if non-nil, zero value otherwise.

### GetAccessMethodOk

`func (o *GetAdmingroupResponse) GetAccessMethodOk() (*[]string, bool)`

GetAccessMethodOk returns a tuple with the AccessMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessMethod

`func (o *GetAdmingroupResponse) SetAccessMethod(v []string)`

SetAccessMethod sets AccessMethod field to given value.

### HasAccessMethod

`func (o *GetAdmingroupResponse) HasAccessMethod() bool`

HasAccessMethod returns a boolean if a field has been set.

### GetAdminSetCommands

`func (o *GetAdmingroupResponse) GetAdminSetCommands() AdmingroupAdminSetCommands`

GetAdminSetCommands returns the AdminSetCommands field if non-nil, zero value otherwise.

### GetAdminSetCommandsOk

`func (o *GetAdmingroupResponse) GetAdminSetCommandsOk() (*AdmingroupAdminSetCommands, bool)`

GetAdminSetCommandsOk returns a tuple with the AdminSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminSetCommands

`func (o *GetAdmingroupResponse) SetAdminSetCommands(v AdmingroupAdminSetCommands)`

SetAdminSetCommands sets AdminSetCommands field to given value.

### HasAdminSetCommands

`func (o *GetAdmingroupResponse) HasAdminSetCommands() bool`

HasAdminSetCommands returns a boolean if a field has been set.

### GetAdminShowCommands

`func (o *GetAdmingroupResponse) GetAdminShowCommands() AdmingroupAdminShowCommands`

GetAdminShowCommands returns the AdminShowCommands field if non-nil, zero value otherwise.

### GetAdminShowCommandsOk

`func (o *GetAdmingroupResponse) GetAdminShowCommandsOk() (*AdmingroupAdminShowCommands, bool)`

GetAdminShowCommandsOk returns a tuple with the AdminShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminShowCommands

`func (o *GetAdmingroupResponse) SetAdminShowCommands(v AdmingroupAdminShowCommands)`

SetAdminShowCommands sets AdminShowCommands field to given value.

### HasAdminShowCommands

`func (o *GetAdmingroupResponse) HasAdminShowCommands() bool`

HasAdminShowCommands returns a boolean if a field has been set.

### GetAdminToplevelCommands

`func (o *GetAdmingroupResponse) GetAdminToplevelCommands() AdmingroupAdminToplevelCommands`

GetAdminToplevelCommands returns the AdminToplevelCommands field if non-nil, zero value otherwise.

### GetAdminToplevelCommandsOk

`func (o *GetAdmingroupResponse) GetAdminToplevelCommandsOk() (*AdmingroupAdminToplevelCommands, bool)`

GetAdminToplevelCommandsOk returns a tuple with the AdminToplevelCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminToplevelCommands

`func (o *GetAdmingroupResponse) SetAdminToplevelCommands(v AdmingroupAdminToplevelCommands)`

SetAdminToplevelCommands sets AdminToplevelCommands field to given value.

### HasAdminToplevelCommands

`func (o *GetAdmingroupResponse) HasAdminToplevelCommands() bool`

HasAdminToplevelCommands returns a boolean if a field has been set.

### GetCloudSetCommands

`func (o *GetAdmingroupResponse) GetCloudSetCommands() AdmingroupCloudSetCommands`

GetCloudSetCommands returns the CloudSetCommands field if non-nil, zero value otherwise.

### GetCloudSetCommandsOk

`func (o *GetAdmingroupResponse) GetCloudSetCommandsOk() (*AdmingroupCloudSetCommands, bool)`

GetCloudSetCommandsOk returns a tuple with the CloudSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudSetCommands

`func (o *GetAdmingroupResponse) SetCloudSetCommands(v AdmingroupCloudSetCommands)`

SetCloudSetCommands sets CloudSetCommands field to given value.

### HasCloudSetCommands

`func (o *GetAdmingroupResponse) HasCloudSetCommands() bool`

HasCloudSetCommands returns a boolean if a field has been set.

### GetCloudShowCommands

`func (o *GetAdmingroupResponse) GetCloudShowCommands() AdmingroupCloudShowCommands`

GetCloudShowCommands returns the CloudShowCommands field if non-nil, zero value otherwise.

### GetCloudShowCommandsOk

`func (o *GetAdmingroupResponse) GetCloudShowCommandsOk() (*AdmingroupCloudShowCommands, bool)`

GetCloudShowCommandsOk returns a tuple with the CloudShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudShowCommands

`func (o *GetAdmingroupResponse) SetCloudShowCommands(v AdmingroupCloudShowCommands)`

SetCloudShowCommands sets CloudShowCommands field to given value.

### HasCloudShowCommands

`func (o *GetAdmingroupResponse) HasCloudShowCommands() bool`

HasCloudShowCommands returns a boolean if a field has been set.

### GetComment

`func (o *GetAdmingroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAdmingroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAdmingroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAdmingroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDatabaseSetCommands

`func (o *GetAdmingroupResponse) GetDatabaseSetCommands() AdmingroupDatabaseSetCommands`

GetDatabaseSetCommands returns the DatabaseSetCommands field if non-nil, zero value otherwise.

### GetDatabaseSetCommandsOk

`func (o *GetAdmingroupResponse) GetDatabaseSetCommandsOk() (*AdmingroupDatabaseSetCommands, bool)`

GetDatabaseSetCommandsOk returns a tuple with the DatabaseSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseSetCommands

`func (o *GetAdmingroupResponse) SetDatabaseSetCommands(v AdmingroupDatabaseSetCommands)`

SetDatabaseSetCommands sets DatabaseSetCommands field to given value.

### HasDatabaseSetCommands

`func (o *GetAdmingroupResponse) HasDatabaseSetCommands() bool`

HasDatabaseSetCommands returns a boolean if a field has been set.

### GetDatabaseShowCommands

`func (o *GetAdmingroupResponse) GetDatabaseShowCommands() AdmingroupDatabaseShowCommands`

GetDatabaseShowCommands returns the DatabaseShowCommands field if non-nil, zero value otherwise.

### GetDatabaseShowCommandsOk

`func (o *GetAdmingroupResponse) GetDatabaseShowCommandsOk() (*AdmingroupDatabaseShowCommands, bool)`

GetDatabaseShowCommandsOk returns a tuple with the DatabaseShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseShowCommands

`func (o *GetAdmingroupResponse) SetDatabaseShowCommands(v AdmingroupDatabaseShowCommands)`

SetDatabaseShowCommands sets DatabaseShowCommands field to given value.

### HasDatabaseShowCommands

`func (o *GetAdmingroupResponse) HasDatabaseShowCommands() bool`

HasDatabaseShowCommands returns a boolean if a field has been set.

### GetDhcpSetCommands

`func (o *GetAdmingroupResponse) GetDhcpSetCommands() AdmingroupDhcpSetCommands`

GetDhcpSetCommands returns the DhcpSetCommands field if non-nil, zero value otherwise.

### GetDhcpSetCommandsOk

`func (o *GetAdmingroupResponse) GetDhcpSetCommandsOk() (*AdmingroupDhcpSetCommands, bool)`

GetDhcpSetCommandsOk returns a tuple with the DhcpSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpSetCommands

`func (o *GetAdmingroupResponse) SetDhcpSetCommands(v AdmingroupDhcpSetCommands)`

SetDhcpSetCommands sets DhcpSetCommands field to given value.

### HasDhcpSetCommands

`func (o *GetAdmingroupResponse) HasDhcpSetCommands() bool`

HasDhcpSetCommands returns a boolean if a field has been set.

### GetDhcpShowCommands

`func (o *GetAdmingroupResponse) GetDhcpShowCommands() AdmingroupDhcpShowCommands`

GetDhcpShowCommands returns the DhcpShowCommands field if non-nil, zero value otherwise.

### GetDhcpShowCommandsOk

`func (o *GetAdmingroupResponse) GetDhcpShowCommandsOk() (*AdmingroupDhcpShowCommands, bool)`

GetDhcpShowCommandsOk returns a tuple with the DhcpShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpShowCommands

`func (o *GetAdmingroupResponse) SetDhcpShowCommands(v AdmingroupDhcpShowCommands)`

SetDhcpShowCommands sets DhcpShowCommands field to given value.

### HasDhcpShowCommands

`func (o *GetAdmingroupResponse) HasDhcpShowCommands() bool`

HasDhcpShowCommands returns a boolean if a field has been set.

### GetDisable

`func (o *GetAdmingroupResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetAdmingroupResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetAdmingroupResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetAdmingroupResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisableConcurrentLogin

`func (o *GetAdmingroupResponse) GetDisableConcurrentLogin() bool`

GetDisableConcurrentLogin returns the DisableConcurrentLogin field if non-nil, zero value otherwise.

### GetDisableConcurrentLoginOk

`func (o *GetAdmingroupResponse) GetDisableConcurrentLoginOk() (*bool, bool)`

GetDisableConcurrentLoginOk returns a tuple with the DisableConcurrentLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConcurrentLogin

`func (o *GetAdmingroupResponse) SetDisableConcurrentLogin(v bool)`

SetDisableConcurrentLogin sets DisableConcurrentLogin field to given value.

### HasDisableConcurrentLogin

`func (o *GetAdmingroupResponse) HasDisableConcurrentLogin() bool`

HasDisableConcurrentLogin returns a boolean if a field has been set.

### GetDnsSetCommands

`func (o *GetAdmingroupResponse) GetDnsSetCommands() AdmingroupDnsSetCommands`

GetDnsSetCommands returns the DnsSetCommands field if non-nil, zero value otherwise.

### GetDnsSetCommandsOk

`func (o *GetAdmingroupResponse) GetDnsSetCommandsOk() (*AdmingroupDnsSetCommands, bool)`

GetDnsSetCommandsOk returns a tuple with the DnsSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSetCommands

`func (o *GetAdmingroupResponse) SetDnsSetCommands(v AdmingroupDnsSetCommands)`

SetDnsSetCommands sets DnsSetCommands field to given value.

### HasDnsSetCommands

`func (o *GetAdmingroupResponse) HasDnsSetCommands() bool`

HasDnsSetCommands returns a boolean if a field has been set.

### GetDnsShowCommands

`func (o *GetAdmingroupResponse) GetDnsShowCommands() AdmingroupDnsShowCommands`

GetDnsShowCommands returns the DnsShowCommands field if non-nil, zero value otherwise.

### GetDnsShowCommandsOk

`func (o *GetAdmingroupResponse) GetDnsShowCommandsOk() (*AdmingroupDnsShowCommands, bool)`

GetDnsShowCommandsOk returns a tuple with the DnsShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsShowCommands

`func (o *GetAdmingroupResponse) SetDnsShowCommands(v AdmingroupDnsShowCommands)`

SetDnsShowCommands sets DnsShowCommands field to given value.

### HasDnsShowCommands

`func (o *GetAdmingroupResponse) HasDnsShowCommands() bool`

HasDnsShowCommands returns a boolean if a field has been set.

### GetDnsToplevelCommands

`func (o *GetAdmingroupResponse) GetDnsToplevelCommands() AdmingroupDnsToplevelCommands`

GetDnsToplevelCommands returns the DnsToplevelCommands field if non-nil, zero value otherwise.

### GetDnsToplevelCommandsOk

`func (o *GetAdmingroupResponse) GetDnsToplevelCommandsOk() (*AdmingroupDnsToplevelCommands, bool)`

GetDnsToplevelCommandsOk returns a tuple with the DnsToplevelCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsToplevelCommands

`func (o *GetAdmingroupResponse) SetDnsToplevelCommands(v AdmingroupDnsToplevelCommands)`

SetDnsToplevelCommands sets DnsToplevelCommands field to given value.

### HasDnsToplevelCommands

`func (o *GetAdmingroupResponse) HasDnsToplevelCommands() bool`

HasDnsToplevelCommands returns a boolean if a field has been set.

### GetDockerSetCommands

`func (o *GetAdmingroupResponse) GetDockerSetCommands() AdmingroupDockerSetCommands`

GetDockerSetCommands returns the DockerSetCommands field if non-nil, zero value otherwise.

### GetDockerSetCommandsOk

`func (o *GetAdmingroupResponse) GetDockerSetCommandsOk() (*AdmingroupDockerSetCommands, bool)`

GetDockerSetCommandsOk returns a tuple with the DockerSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerSetCommands

`func (o *GetAdmingroupResponse) SetDockerSetCommands(v AdmingroupDockerSetCommands)`

SetDockerSetCommands sets DockerSetCommands field to given value.

### HasDockerSetCommands

`func (o *GetAdmingroupResponse) HasDockerSetCommands() bool`

HasDockerSetCommands returns a boolean if a field has been set.

### GetDockerShowCommands

`func (o *GetAdmingroupResponse) GetDockerShowCommands() AdmingroupDockerShowCommands`

GetDockerShowCommands returns the DockerShowCommands field if non-nil, zero value otherwise.

### GetDockerShowCommandsOk

`func (o *GetAdmingroupResponse) GetDockerShowCommandsOk() (*AdmingroupDockerShowCommands, bool)`

GetDockerShowCommandsOk returns a tuple with the DockerShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerShowCommands

`func (o *GetAdmingroupResponse) SetDockerShowCommands(v AdmingroupDockerShowCommands)`

SetDockerShowCommands sets DockerShowCommands field to given value.

### HasDockerShowCommands

`func (o *GetAdmingroupResponse) HasDockerShowCommands() bool`

HasDockerShowCommands returns a boolean if a field has been set.

### GetEmailAddresses

`func (o *GetAdmingroupResponse) GetEmailAddresses() []string`

GetEmailAddresses returns the EmailAddresses field if non-nil, zero value otherwise.

### GetEmailAddressesOk

`func (o *GetAdmingroupResponse) GetEmailAddressesOk() (*[]string, bool)`

GetEmailAddressesOk returns a tuple with the EmailAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddresses

`func (o *GetAdmingroupResponse) SetEmailAddresses(v []string)`

SetEmailAddresses sets EmailAddresses field to given value.

### HasEmailAddresses

`func (o *GetAdmingroupResponse) HasEmailAddresses() bool`

HasEmailAddresses returns a boolean if a field has been set.

### GetEnableRestrictedUserAccess

`func (o *GetAdmingroupResponse) GetEnableRestrictedUserAccess() bool`

GetEnableRestrictedUserAccess returns the EnableRestrictedUserAccess field if non-nil, zero value otherwise.

### GetEnableRestrictedUserAccessOk

`func (o *GetAdmingroupResponse) GetEnableRestrictedUserAccessOk() (*bool, bool)`

GetEnableRestrictedUserAccessOk returns a tuple with the EnableRestrictedUserAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRestrictedUserAccess

`func (o *GetAdmingroupResponse) SetEnableRestrictedUserAccess(v bool)`

SetEnableRestrictedUserAccess sets EnableRestrictedUserAccess field to given value.

### HasEnableRestrictedUserAccess

`func (o *GetAdmingroupResponse) HasEnableRestrictedUserAccess() bool`

HasEnableRestrictedUserAccess returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetAdmingroupResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetAdmingroupResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetAdmingroupResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetAdmingroupResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetGridSetCommands

`func (o *GetAdmingroupResponse) GetGridSetCommands() AdmingroupGridSetCommands`

GetGridSetCommands returns the GridSetCommands field if non-nil, zero value otherwise.

### GetGridSetCommandsOk

`func (o *GetAdmingroupResponse) GetGridSetCommandsOk() (*AdmingroupGridSetCommands, bool)`

GetGridSetCommandsOk returns a tuple with the GridSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridSetCommands

`func (o *GetAdmingroupResponse) SetGridSetCommands(v AdmingroupGridSetCommands)`

SetGridSetCommands sets GridSetCommands field to given value.

### HasGridSetCommands

`func (o *GetAdmingroupResponse) HasGridSetCommands() bool`

HasGridSetCommands returns a boolean if a field has been set.

### GetGridShowCommands

`func (o *GetAdmingroupResponse) GetGridShowCommands() AdmingroupGridShowCommands`

GetGridShowCommands returns the GridShowCommands field if non-nil, zero value otherwise.

### GetGridShowCommandsOk

`func (o *GetAdmingroupResponse) GetGridShowCommandsOk() (*AdmingroupGridShowCommands, bool)`

GetGridShowCommandsOk returns a tuple with the GridShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridShowCommands

`func (o *GetAdmingroupResponse) SetGridShowCommands(v AdmingroupGridShowCommands)`

SetGridShowCommands sets GridShowCommands field to given value.

### HasGridShowCommands

`func (o *GetAdmingroupResponse) HasGridShowCommands() bool`

HasGridShowCommands returns a boolean if a field has been set.

### GetInactivityLockoutSetting

`func (o *GetAdmingroupResponse) GetInactivityLockoutSetting() AdmingroupInactivityLockoutSetting`

GetInactivityLockoutSetting returns the InactivityLockoutSetting field if non-nil, zero value otherwise.

### GetInactivityLockoutSettingOk

`func (o *GetAdmingroupResponse) GetInactivityLockoutSettingOk() (*AdmingroupInactivityLockoutSetting, bool)`

GetInactivityLockoutSettingOk returns a tuple with the InactivityLockoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactivityLockoutSetting

`func (o *GetAdmingroupResponse) SetInactivityLockoutSetting(v AdmingroupInactivityLockoutSetting)`

SetInactivityLockoutSetting sets InactivityLockoutSetting field to given value.

### HasInactivityLockoutSetting

`func (o *GetAdmingroupResponse) HasInactivityLockoutSetting() bool`

HasInactivityLockoutSetting returns a boolean if a field has been set.

### GetLicensingSetCommands

`func (o *GetAdmingroupResponse) GetLicensingSetCommands() AdmingroupLicensingSetCommands`

GetLicensingSetCommands returns the LicensingSetCommands field if non-nil, zero value otherwise.

### GetLicensingSetCommandsOk

`func (o *GetAdmingroupResponse) GetLicensingSetCommandsOk() (*AdmingroupLicensingSetCommands, bool)`

GetLicensingSetCommandsOk returns a tuple with the LicensingSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingSetCommands

`func (o *GetAdmingroupResponse) SetLicensingSetCommands(v AdmingroupLicensingSetCommands)`

SetLicensingSetCommands sets LicensingSetCommands field to given value.

### HasLicensingSetCommands

`func (o *GetAdmingroupResponse) HasLicensingSetCommands() bool`

HasLicensingSetCommands returns a boolean if a field has been set.

### GetLicensingShowCommands

`func (o *GetAdmingroupResponse) GetLicensingShowCommands() AdmingroupLicensingShowCommands`

GetLicensingShowCommands returns the LicensingShowCommands field if non-nil, zero value otherwise.

### GetLicensingShowCommandsOk

`func (o *GetAdmingroupResponse) GetLicensingShowCommandsOk() (*AdmingroupLicensingShowCommands, bool)`

GetLicensingShowCommandsOk returns a tuple with the LicensingShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensingShowCommands

`func (o *GetAdmingroupResponse) SetLicensingShowCommands(v AdmingroupLicensingShowCommands)`

SetLicensingShowCommands sets LicensingShowCommands field to given value.

### HasLicensingShowCommands

`func (o *GetAdmingroupResponse) HasLicensingShowCommands() bool`

HasLicensingShowCommands returns a boolean if a field has been set.

### GetLockoutSetting

`func (o *GetAdmingroupResponse) GetLockoutSetting() AdmingroupLockoutSetting`

GetLockoutSetting returns the LockoutSetting field if non-nil, zero value otherwise.

### GetLockoutSettingOk

`func (o *GetAdmingroupResponse) GetLockoutSettingOk() (*AdmingroupLockoutSetting, bool)`

GetLockoutSettingOk returns a tuple with the LockoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutSetting

`func (o *GetAdmingroupResponse) SetLockoutSetting(v AdmingroupLockoutSetting)`

SetLockoutSetting sets LockoutSetting field to given value.

### HasLockoutSetting

`func (o *GetAdmingroupResponse) HasLockoutSetting() bool`

HasLockoutSetting returns a boolean if a field has been set.

### GetMachineControlToplevelCommands

`func (o *GetAdmingroupResponse) GetMachineControlToplevelCommands() AdmingroupMachineControlToplevelCommands`

GetMachineControlToplevelCommands returns the MachineControlToplevelCommands field if non-nil, zero value otherwise.

### GetMachineControlToplevelCommandsOk

`func (o *GetAdmingroupResponse) GetMachineControlToplevelCommandsOk() (*AdmingroupMachineControlToplevelCommands, bool)`

GetMachineControlToplevelCommandsOk returns a tuple with the MachineControlToplevelCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineControlToplevelCommands

`func (o *GetAdmingroupResponse) SetMachineControlToplevelCommands(v AdmingroupMachineControlToplevelCommands)`

SetMachineControlToplevelCommands sets MachineControlToplevelCommands field to given value.

### HasMachineControlToplevelCommands

`func (o *GetAdmingroupResponse) HasMachineControlToplevelCommands() bool`

HasMachineControlToplevelCommands returns a boolean if a field has been set.

### GetName

`func (o *GetAdmingroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAdmingroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAdmingroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAdmingroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkingSetCommands

`func (o *GetAdmingroupResponse) GetNetworkingSetCommands() AdmingroupNetworkingSetCommands`

GetNetworkingSetCommands returns the NetworkingSetCommands field if non-nil, zero value otherwise.

### GetNetworkingSetCommandsOk

`func (o *GetAdmingroupResponse) GetNetworkingSetCommandsOk() (*AdmingroupNetworkingSetCommands, bool)`

GetNetworkingSetCommandsOk returns a tuple with the NetworkingSetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingSetCommands

`func (o *GetAdmingroupResponse) SetNetworkingSetCommands(v AdmingroupNetworkingSetCommands)`

SetNetworkingSetCommands sets NetworkingSetCommands field to given value.

### HasNetworkingSetCommands

`func (o *GetAdmingroupResponse) HasNetworkingSetCommands() bool`

HasNetworkingSetCommands returns a boolean if a field has been set.

### GetNetworkingShowCommands

`func (o *GetAdmingroupResponse) GetNetworkingShowCommands() AdmingroupNetworkingShowCommands`

GetNetworkingShowCommands returns the NetworkingShowCommands field if non-nil, zero value otherwise.

### GetNetworkingShowCommandsOk

`func (o *GetAdmingroupResponse) GetNetworkingShowCommandsOk() (*AdmingroupNetworkingShowCommands, bool)`

GetNetworkingShowCommandsOk returns a tuple with the NetworkingShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkingShowCommands

`func (o *GetAdmingroupResponse) SetNetworkingShowCommands(v AdmingroupNetworkingShowCommands)`

SetNetworkingShowCommands sets NetworkingShowCommands field to given value.

### HasNetworkingShowCommands

`func (o *GetAdmingroupResponse) HasNetworkingShowCommands() bool`

HasNetworkingShowCommands returns a boolean if a field has been set.

### GetPasswordSetting

`func (o *GetAdmingroupResponse) GetPasswordSetting() AdmingroupPasswordSetting`

GetPasswordSetting returns the PasswordSetting field if non-nil, zero value otherwise.

### GetPasswordSettingOk

`func (o *GetAdmingroupResponse) GetPasswordSettingOk() (*AdmingroupPasswordSetting, bool)`

GetPasswordSettingOk returns a tuple with the PasswordSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSetting

`func (o *GetAdmingroupResponse) SetPasswordSetting(v AdmingroupPasswordSetting)`

SetPasswordSetting sets PasswordSetting field to given value.

### HasPasswordSetting

`func (o *GetAdmingroupResponse) HasPasswordSetting() bool`

HasPasswordSetting returns a boolean if a field has been set.

### GetRoles

`func (o *GetAdmingroupResponse) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *GetAdmingroupResponse) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *GetAdmingroupResponse) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *GetAdmingroupResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSamlSetting

`func (o *GetAdmingroupResponse) GetSamlSetting() AdmingroupSamlSetting`

GetSamlSetting returns the SamlSetting field if non-nil, zero value otherwise.

### GetSamlSettingOk

`func (o *GetAdmingroupResponse) GetSamlSettingOk() (*AdmingroupSamlSetting, bool)`

GetSamlSettingOk returns a tuple with the SamlSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlSetting

`func (o *GetAdmingroupResponse) SetSamlSetting(v AdmingroupSamlSetting)`

SetSamlSetting sets SamlSetting field to given value.

### HasSamlSetting

`func (o *GetAdmingroupResponse) HasSamlSetting() bool`

HasSamlSetting returns a boolean if a field has been set.

### GetSecuritySetCommands

`func (o *GetAdmingroupResponse) GetSecuritySetCommands() AdmingroupSecuritySetCommands`

GetSecuritySetCommands returns the SecuritySetCommands field if non-nil, zero value otherwise.

### GetSecuritySetCommandsOk

`func (o *GetAdmingroupResponse) GetSecuritySetCommandsOk() (*AdmingroupSecuritySetCommands, bool)`

GetSecuritySetCommandsOk returns a tuple with the SecuritySetCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySetCommands

`func (o *GetAdmingroupResponse) SetSecuritySetCommands(v AdmingroupSecuritySetCommands)`

SetSecuritySetCommands sets SecuritySetCommands field to given value.

### HasSecuritySetCommands

`func (o *GetAdmingroupResponse) HasSecuritySetCommands() bool`

HasSecuritySetCommands returns a boolean if a field has been set.

### GetSecurityShowCommands

`func (o *GetAdmingroupResponse) GetSecurityShowCommands() AdmingroupSecurityShowCommands`

GetSecurityShowCommands returns the SecurityShowCommands field if non-nil, zero value otherwise.

### GetSecurityShowCommandsOk

`func (o *GetAdmingroupResponse) GetSecurityShowCommandsOk() (*AdmingroupSecurityShowCommands, bool)`

GetSecurityShowCommandsOk returns a tuple with the SecurityShowCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityShowCommands

`func (o *GetAdmingroupResponse) SetSecurityShowCommands(v AdmingroupSecurityShowCommands)`

SetSecurityShowCommands sets SecurityShowCommands field to given value.

### HasSecurityShowCommands

`func (o *GetAdmingroupResponse) HasSecurityShowCommands() bool`

HasSecurityShowCommands returns a boolean if a field has been set.

### GetSuperuser

`func (o *GetAdmingroupResponse) GetSuperuser() bool`

GetSuperuser returns the Superuser field if non-nil, zero value otherwise.

### GetSuperuserOk

`func (o *GetAdmingroupResponse) GetSuperuserOk() (*bool, bool)`

GetSuperuserOk returns a tuple with the Superuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuperuser

`func (o *GetAdmingroupResponse) SetSuperuser(v bool)`

SetSuperuser sets Superuser field to given value.

### HasSuperuser

`func (o *GetAdmingroupResponse) HasSuperuser() bool`

HasSuperuser returns a boolean if a field has been set.

### GetTroubleShootingToplevelCommands

`func (o *GetAdmingroupResponse) GetTroubleShootingToplevelCommands() AdmingroupTroubleShootingToplevelCommands`

GetTroubleShootingToplevelCommands returns the TroubleShootingToplevelCommands field if non-nil, zero value otherwise.

### GetTroubleShootingToplevelCommandsOk

`func (o *GetAdmingroupResponse) GetTroubleShootingToplevelCommandsOk() (*AdmingroupTroubleShootingToplevelCommands, bool)`

GetTroubleShootingToplevelCommandsOk returns a tuple with the TroubleShootingToplevelCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTroubleShootingToplevelCommands

`func (o *GetAdmingroupResponse) SetTroubleShootingToplevelCommands(v AdmingroupTroubleShootingToplevelCommands)`

SetTroubleShootingToplevelCommands sets TroubleShootingToplevelCommands field to given value.

### HasTroubleShootingToplevelCommands

`func (o *GetAdmingroupResponse) HasTroubleShootingToplevelCommands() bool`

HasTroubleShootingToplevelCommands returns a boolean if a field has been set.

### GetUseAccountInactivityLockoutEnable

`func (o *GetAdmingroupResponse) GetUseAccountInactivityLockoutEnable() bool`

GetUseAccountInactivityLockoutEnable returns the UseAccountInactivityLockoutEnable field if non-nil, zero value otherwise.

### GetUseAccountInactivityLockoutEnableOk

`func (o *GetAdmingroupResponse) GetUseAccountInactivityLockoutEnableOk() (*bool, bool)`

GetUseAccountInactivityLockoutEnableOk returns a tuple with the UseAccountInactivityLockoutEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAccountInactivityLockoutEnable

`func (o *GetAdmingroupResponse) SetUseAccountInactivityLockoutEnable(v bool)`

SetUseAccountInactivityLockoutEnable sets UseAccountInactivityLockoutEnable field to given value.

### HasUseAccountInactivityLockoutEnable

`func (o *GetAdmingroupResponse) HasUseAccountInactivityLockoutEnable() bool`

HasUseAccountInactivityLockoutEnable returns a boolean if a field has been set.

### GetUseDisableConcurrentLogin

`func (o *GetAdmingroupResponse) GetUseDisableConcurrentLogin() bool`

GetUseDisableConcurrentLogin returns the UseDisableConcurrentLogin field if non-nil, zero value otherwise.

### GetUseDisableConcurrentLoginOk

`func (o *GetAdmingroupResponse) GetUseDisableConcurrentLoginOk() (*bool, bool)`

GetUseDisableConcurrentLoginOk returns a tuple with the UseDisableConcurrentLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDisableConcurrentLogin

`func (o *GetAdmingroupResponse) SetUseDisableConcurrentLogin(v bool)`

SetUseDisableConcurrentLogin sets UseDisableConcurrentLogin field to given value.

### HasUseDisableConcurrentLogin

`func (o *GetAdmingroupResponse) HasUseDisableConcurrentLogin() bool`

HasUseDisableConcurrentLogin returns a boolean if a field has been set.

### GetUseLockoutSetting

`func (o *GetAdmingroupResponse) GetUseLockoutSetting() bool`

GetUseLockoutSetting returns the UseLockoutSetting field if non-nil, zero value otherwise.

### GetUseLockoutSettingOk

`func (o *GetAdmingroupResponse) GetUseLockoutSettingOk() (*bool, bool)`

GetUseLockoutSettingOk returns a tuple with the UseLockoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLockoutSetting

`func (o *GetAdmingroupResponse) SetUseLockoutSetting(v bool)`

SetUseLockoutSetting sets UseLockoutSetting field to given value.

### HasUseLockoutSetting

`func (o *GetAdmingroupResponse) HasUseLockoutSetting() bool`

HasUseLockoutSetting returns a boolean if a field has been set.

### GetUsePasswordSetting

`func (o *GetAdmingroupResponse) GetUsePasswordSetting() bool`

GetUsePasswordSetting returns the UsePasswordSetting field if non-nil, zero value otherwise.

### GetUsePasswordSettingOk

`func (o *GetAdmingroupResponse) GetUsePasswordSettingOk() (*bool, bool)`

GetUsePasswordSettingOk returns a tuple with the UsePasswordSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsePasswordSetting

`func (o *GetAdmingroupResponse) SetUsePasswordSetting(v bool)`

SetUsePasswordSetting sets UsePasswordSetting field to given value.

### HasUsePasswordSetting

`func (o *GetAdmingroupResponse) HasUsePasswordSetting() bool`

HasUsePasswordSetting returns a boolean if a field has been set.

### GetUserAccess

`func (o *GetAdmingroupResponse) GetUserAccess() []AdmingroupUserAccess`

GetUserAccess returns the UserAccess field if non-nil, zero value otherwise.

### GetUserAccessOk

`func (o *GetAdmingroupResponse) GetUserAccessOk() (*[]AdmingroupUserAccess, bool)`

GetUserAccessOk returns a tuple with the UserAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccess

`func (o *GetAdmingroupResponse) SetUserAccess(v []AdmingroupUserAccess)`

SetUserAccess sets UserAccess field to given value.

### HasUserAccess

`func (o *GetAdmingroupResponse) HasUserAccess() bool`

HasUserAccess returns a boolean if a field has been set.

### GetResult

`func (o *GetAdmingroupResponse) GetResult() Admingroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAdmingroupResponse) GetResultOk() (*Admingroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAdmingroupResponse) SetResult(v Admingroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAdmingroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


