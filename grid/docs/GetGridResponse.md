# GetGridResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AllowRecursiveDeletion** | Pointer to **string** | The property to allow recursive deletion. Determines the users who can choose to perform recursive deletion on networks or zones from the GUI only. | [optional] 
**AuditLogFormat** | Pointer to **string** | Determines the audit log format. | [optional] 
**AuditToSyslogEnable** | Pointer to **bool** | If set to True, audit log messages are also copied to the syslog. | [optional] 
**AutomatedTrafficCaptureSetting** | Pointer to [**GridAutomatedTrafficCaptureSetting**](GridAutomatedTrafficCaptureSetting.md) |  | [optional] 
**ConsentBannerSetting** | Pointer to [**GridConsentBannerSetting**](GridConsentBannerSetting.md) |  | [optional] 
**CspApiConfig** | Pointer to [**GridCspApiConfig**](GridCspApiConfig.md) |  | [optional] 
**CspGridSetting** | Pointer to [**GridCspGridSetting**](GridCspGridSetting.md) |  | [optional] 
**DenyMgmSnapshots** | Pointer to **bool** | If set to True, the managed Grid will not send snapshots to the Multi-Grid Master. | [optional] 
**DescendantsAction** | Pointer to [**GridDescendantsAction**](GridDescendantsAction.md) |  | [optional] 
**DnsResolverSetting** | Pointer to [**GridDnsResolverSetting**](GridDnsResolverSetting.md) |  | [optional] 
**Dscp** | Pointer to **int64** | The DSCP value. Valid values are integers between 0 and 63 inclusive. | [optional] 
**EmailSetting** | Pointer to [**GridEmailSetting**](GridEmailSetting.md) |  | [optional] 
**EnableFederation** | Pointer to **bool** | Determines if the federation feature is enabled or not. Test Setting will be performed for any change in enable_federation. | [optional] 
**EnableForceSyncJoinTokenToGmc** | Pointer to **bool** | Determines if the force sync join token from GM to GMC is enabled or not. | [optional] 
**EnableGuiApiForLanVip** | Pointer to **bool** | If set to True, GUI and API access are enabled on the LAN/VIP port and MGMT port (if configured). | [optional] 
**EnableLom** | Pointer to **bool** | Determines if the LOM functionality is enabled or not. | [optional] 
**EnableMemberRedirect** | Pointer to **bool** | Determines redirections is enabled or not for members. | [optional] 
**EnableRecycleBin** | Pointer to **bool** | Determines if the Recycle Bin is enabled or not. | [optional] 
**EnableRirSwip** | Pointer to **bool** | Determines if the RIR/SWIP support is enabled or not. | [optional] 
**ExternalSyslogBackupServers** | Pointer to [**[]GridExternalSyslogBackupServers**](GridExternalSyslogBackupServers.md) | The list of external backup syslog servers. | [optional] 
**ExternalSyslogServerEnable** | Pointer to **bool** | If set to True, external syslog servers are enabled. | [optional] 
**HttpProxyServerSetting** | Pointer to [**GridHttpProxyServerSetting**](GridHttpProxyServerSetting.md) |  | [optional] 
**InformationalBannerSetting** | Pointer to [**GridInformationalBannerSetting**](GridInformationalBannerSetting.md) |  | [optional] 
**IsGridVisualizationVisible** | Pointer to **bool** | If set to True, graphical visualization of the Grid is enabled. | [optional] 
**LockoutSetting** | Pointer to [**GridLockoutSetting**](GridLockoutSetting.md) |  | [optional] 
**LomUsers** | Pointer to [**[]GridLomUsers**](GridLomUsers.md) | The list of LOM users. | [optional] 
**MgmStrictDelegateMode** | Pointer to **bool** | Determines if strict delegate mode for the Grid managed by the Master Grid is enabled or not. | [optional] 
**MsSetting** | Pointer to [**GridMsSetting**](GridMsSetting.md) |  | [optional] 
**Name** | Pointer to **string** | The grid name. | [optional] 
**NatGroups** | Pointer to **[]string** | The list of all Network Address Translation (NAT) groups configured on the Grid. | [optional] 
**NtpSetting** | Pointer to [**GridNtpSetting**](GridNtpSetting.md) |  | [optional] 
**ObjectsChangesTrackingSetting** | Pointer to [**GridObjectsChangesTrackingSetting**](GridObjectsChangesTrackingSetting.md) |  | [optional] 
**PasswordSetting** | Pointer to [**GridPasswordSetting**](GridPasswordSetting.md) |  | [optional] 
**RestartBannerSetting** | Pointer to [**GridRestartBannerSetting**](GridRestartBannerSetting.md) |  | [optional] 
**RestartStatus** | Pointer to **string** | The restart status for the Grid. | [optional] [readonly] 
**RpzHitRateInterval** | Pointer to **int64** | The time interval (in seconds) that determines how often the appliance calculates the RPZ hit rate. | [optional] 
**RpzHitRateMaxQuery** | Pointer to **int64** | The maximum number of incoming queries between the RPZ hit rate checks. | [optional] 
**RpzHitRateMinQuery** | Pointer to **int64** | The minimum number of incoming queries between the RPZ hit rate checks. | [optional] 
**ScheduledBackup** | Pointer to [**GridScheduledBackup**](GridScheduledBackup.md) |  | [optional] 
**Secret** | Pointer to **string** | The shared secret of the Grid. This is a write-only attribute. | [optional] 
**SecurityBannerSetting** | Pointer to [**GridSecurityBannerSetting**](GridSecurityBannerSetting.md) |  | [optional] 
**SecuritySetting** | Pointer to [**GridSecuritySetting**](GridSecuritySetting.md) |  | [optional] 
**ServiceStatus** | Pointer to **string** | Determines overall service status of the Grid. | [optional] [readonly] 
**SnmpSetting** | Pointer to [**GridSnmpSetting**](GridSnmpSetting.md) |  | [optional] 
**SupportBundleDownloadTimeout** | Pointer to **int64** | Support bundle download timeout in seconds. | [optional] 
**SyslogFacility** | Pointer to **string** | If &#39;audit_to_syslog_enable&#39; is set to True, the facility that determines the processes and daemons from which the log messages are generated. | [optional] 
**SyslogServers** | Pointer to [**[]GridSyslogServers**](GridSyslogServers.md) | The list of external syslog servers. | [optional] 
**SyslogSize** | Pointer to **int64** | The maximum size for the syslog file expressed in megabytes. | [optional] 
**ThresholdTraps** | Pointer to [**[]GridThresholdTraps**](GridThresholdTraps.md) | Determines the list of threshold traps. The user can only change the values for each trap or remove traps. | [optional] 
**TimeZone** | Pointer to **string** | The time zone of the Grid. The UTC string that represents the time zone, such as \&quot;US/Eastern\&quot;. | [optional] 
**TokenUsageDelay** | Pointer to **int64** | The delayed usage (in minutes) of a permission token. | [optional] 
**TrafficCaptureAuthDnsSetting** | Pointer to [**GridTrafficCaptureAuthDnsSetting**](GridTrafficCaptureAuthDnsSetting.md) |  | [optional] 
**TrafficCaptureChrSetting** | Pointer to [**GridTrafficCaptureChrSetting**](GridTrafficCaptureChrSetting.md) |  | [optional] 
**TrafficCaptureQpsSetting** | Pointer to [**GridTrafficCaptureQpsSetting**](GridTrafficCaptureQpsSetting.md) |  | [optional] 
**TrafficCaptureRecDnsSetting** | Pointer to [**GridTrafficCaptureRecDnsSetting**](GridTrafficCaptureRecDnsSetting.md) |  | [optional] 
**TrafficCaptureRecQueriesSetting** | Pointer to [**GridTrafficCaptureRecQueriesSetting**](GridTrafficCaptureRecQueriesSetting.md) |  | [optional] 
**TrapNotifications** | Pointer to [**[]GridTrapNotifications**](GridTrapNotifications.md) | Determines configuration of the trap notifications. | [optional] 
**UpdatesDownloadMemberConfig** | Pointer to [**[]GridUpdatesDownloadMemberConfig**](GridUpdatesDownloadMemberConfig.md) | The list of member configuration structures, which provides information and settings for configuring the member that is responsible for downloading updates. | [optional] 
**VpnPort** | Pointer to **int64** | The VPN port. | [optional] 
**Result** | Pointer to [**Grid**](Grid.md) |  | [optional] 

## Methods

### NewGetGridResponse

`func NewGetGridResponse() *GetGridResponse`

NewGetGridResponse instantiates a new GetGridResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridResponseWithDefaults

`func NewGetGridResponseWithDefaults() *GetGridResponse`

NewGetGridResponseWithDefaults instantiates a new GetGridResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowRecursiveDeletion

`func (o *GetGridResponse) GetAllowRecursiveDeletion() string`

GetAllowRecursiveDeletion returns the AllowRecursiveDeletion field if non-nil, zero value otherwise.

### GetAllowRecursiveDeletionOk

`func (o *GetGridResponse) GetAllowRecursiveDeletionOk() (*string, bool)`

GetAllowRecursiveDeletionOk returns a tuple with the AllowRecursiveDeletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRecursiveDeletion

`func (o *GetGridResponse) SetAllowRecursiveDeletion(v string)`

SetAllowRecursiveDeletion sets AllowRecursiveDeletion field to given value.

### HasAllowRecursiveDeletion

`func (o *GetGridResponse) HasAllowRecursiveDeletion() bool`

HasAllowRecursiveDeletion returns a boolean if a field has been set.

### GetAuditLogFormat

`func (o *GetGridResponse) GetAuditLogFormat() string`

GetAuditLogFormat returns the AuditLogFormat field if non-nil, zero value otherwise.

### GetAuditLogFormatOk

`func (o *GetGridResponse) GetAuditLogFormatOk() (*string, bool)`

GetAuditLogFormatOk returns a tuple with the AuditLogFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditLogFormat

`func (o *GetGridResponse) SetAuditLogFormat(v string)`

SetAuditLogFormat sets AuditLogFormat field to given value.

### HasAuditLogFormat

`func (o *GetGridResponse) HasAuditLogFormat() bool`

HasAuditLogFormat returns a boolean if a field has been set.

### GetAuditToSyslogEnable

`func (o *GetGridResponse) GetAuditToSyslogEnable() bool`

GetAuditToSyslogEnable returns the AuditToSyslogEnable field if non-nil, zero value otherwise.

### GetAuditToSyslogEnableOk

`func (o *GetGridResponse) GetAuditToSyslogEnableOk() (*bool, bool)`

GetAuditToSyslogEnableOk returns a tuple with the AuditToSyslogEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditToSyslogEnable

`func (o *GetGridResponse) SetAuditToSyslogEnable(v bool)`

SetAuditToSyslogEnable sets AuditToSyslogEnable field to given value.

### HasAuditToSyslogEnable

`func (o *GetGridResponse) HasAuditToSyslogEnable() bool`

HasAuditToSyslogEnable returns a boolean if a field has been set.

### GetAutomatedTrafficCaptureSetting

`func (o *GetGridResponse) GetAutomatedTrafficCaptureSetting() GridAutomatedTrafficCaptureSetting`

GetAutomatedTrafficCaptureSetting returns the AutomatedTrafficCaptureSetting field if non-nil, zero value otherwise.

### GetAutomatedTrafficCaptureSettingOk

`func (o *GetGridResponse) GetAutomatedTrafficCaptureSettingOk() (*GridAutomatedTrafficCaptureSetting, bool)`

GetAutomatedTrafficCaptureSettingOk returns a tuple with the AutomatedTrafficCaptureSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedTrafficCaptureSetting

`func (o *GetGridResponse) SetAutomatedTrafficCaptureSetting(v GridAutomatedTrafficCaptureSetting)`

SetAutomatedTrafficCaptureSetting sets AutomatedTrafficCaptureSetting field to given value.

### HasAutomatedTrafficCaptureSetting

`func (o *GetGridResponse) HasAutomatedTrafficCaptureSetting() bool`

HasAutomatedTrafficCaptureSetting returns a boolean if a field has been set.

### GetConsentBannerSetting

`func (o *GetGridResponse) GetConsentBannerSetting() GridConsentBannerSetting`

GetConsentBannerSetting returns the ConsentBannerSetting field if non-nil, zero value otherwise.

### GetConsentBannerSettingOk

`func (o *GetGridResponse) GetConsentBannerSettingOk() (*GridConsentBannerSetting, bool)`

GetConsentBannerSettingOk returns a tuple with the ConsentBannerSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentBannerSetting

`func (o *GetGridResponse) SetConsentBannerSetting(v GridConsentBannerSetting)`

SetConsentBannerSetting sets ConsentBannerSetting field to given value.

### HasConsentBannerSetting

`func (o *GetGridResponse) HasConsentBannerSetting() bool`

HasConsentBannerSetting returns a boolean if a field has been set.

### GetCspApiConfig

`func (o *GetGridResponse) GetCspApiConfig() GridCspApiConfig`

GetCspApiConfig returns the CspApiConfig field if non-nil, zero value otherwise.

### GetCspApiConfigOk

`func (o *GetGridResponse) GetCspApiConfigOk() (*GridCspApiConfig, bool)`

GetCspApiConfigOk returns a tuple with the CspApiConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspApiConfig

`func (o *GetGridResponse) SetCspApiConfig(v GridCspApiConfig)`

SetCspApiConfig sets CspApiConfig field to given value.

### HasCspApiConfig

`func (o *GetGridResponse) HasCspApiConfig() bool`

HasCspApiConfig returns a boolean if a field has been set.

### GetCspGridSetting

`func (o *GetGridResponse) GetCspGridSetting() GridCspGridSetting`

GetCspGridSetting returns the CspGridSetting field if non-nil, zero value otherwise.

### GetCspGridSettingOk

`func (o *GetGridResponse) GetCspGridSettingOk() (*GridCspGridSetting, bool)`

GetCspGridSettingOk returns a tuple with the CspGridSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspGridSetting

`func (o *GetGridResponse) SetCspGridSetting(v GridCspGridSetting)`

SetCspGridSetting sets CspGridSetting field to given value.

### HasCspGridSetting

`func (o *GetGridResponse) HasCspGridSetting() bool`

HasCspGridSetting returns a boolean if a field has been set.

### GetDenyMgmSnapshots

`func (o *GetGridResponse) GetDenyMgmSnapshots() bool`

GetDenyMgmSnapshots returns the DenyMgmSnapshots field if non-nil, zero value otherwise.

### GetDenyMgmSnapshotsOk

`func (o *GetGridResponse) GetDenyMgmSnapshotsOk() (*bool, bool)`

GetDenyMgmSnapshotsOk returns a tuple with the DenyMgmSnapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenyMgmSnapshots

`func (o *GetGridResponse) SetDenyMgmSnapshots(v bool)`

SetDenyMgmSnapshots sets DenyMgmSnapshots field to given value.

### HasDenyMgmSnapshots

`func (o *GetGridResponse) HasDenyMgmSnapshots() bool`

HasDenyMgmSnapshots returns a boolean if a field has been set.

### GetDescendantsAction

`func (o *GetGridResponse) GetDescendantsAction() GridDescendantsAction`

GetDescendantsAction returns the DescendantsAction field if non-nil, zero value otherwise.

### GetDescendantsActionOk

`func (o *GetGridResponse) GetDescendantsActionOk() (*GridDescendantsAction, bool)`

GetDescendantsActionOk returns a tuple with the DescendantsAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescendantsAction

`func (o *GetGridResponse) SetDescendantsAction(v GridDescendantsAction)`

SetDescendantsAction sets DescendantsAction field to given value.

### HasDescendantsAction

`func (o *GetGridResponse) HasDescendantsAction() bool`

HasDescendantsAction returns a boolean if a field has been set.

### GetDnsResolverSetting

`func (o *GetGridResponse) GetDnsResolverSetting() GridDnsResolverSetting`

GetDnsResolverSetting returns the DnsResolverSetting field if non-nil, zero value otherwise.

### GetDnsResolverSettingOk

`func (o *GetGridResponse) GetDnsResolverSettingOk() (*GridDnsResolverSetting, bool)`

GetDnsResolverSettingOk returns a tuple with the DnsResolverSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolverSetting

`func (o *GetGridResponse) SetDnsResolverSetting(v GridDnsResolverSetting)`

SetDnsResolverSetting sets DnsResolverSetting field to given value.

### HasDnsResolverSetting

`func (o *GetGridResponse) HasDnsResolverSetting() bool`

HasDnsResolverSetting returns a boolean if a field has been set.

### GetDscp

`func (o *GetGridResponse) GetDscp() int64`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *GetGridResponse) GetDscpOk() (*int64, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *GetGridResponse) SetDscp(v int64)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *GetGridResponse) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetEmailSetting

`func (o *GetGridResponse) GetEmailSetting() GridEmailSetting`

GetEmailSetting returns the EmailSetting field if non-nil, zero value otherwise.

### GetEmailSettingOk

`func (o *GetGridResponse) GetEmailSettingOk() (*GridEmailSetting, bool)`

GetEmailSettingOk returns a tuple with the EmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSetting

`func (o *GetGridResponse) SetEmailSetting(v GridEmailSetting)`

SetEmailSetting sets EmailSetting field to given value.

### HasEmailSetting

`func (o *GetGridResponse) HasEmailSetting() bool`

HasEmailSetting returns a boolean if a field has been set.

### GetEnableFederation

`func (o *GetGridResponse) GetEnableFederation() bool`

GetEnableFederation returns the EnableFederation field if non-nil, zero value otherwise.

### GetEnableFederationOk

`func (o *GetGridResponse) GetEnableFederationOk() (*bool, bool)`

GetEnableFederationOk returns a tuple with the EnableFederation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFederation

`func (o *GetGridResponse) SetEnableFederation(v bool)`

SetEnableFederation sets EnableFederation field to given value.

### HasEnableFederation

`func (o *GetGridResponse) HasEnableFederation() bool`

HasEnableFederation returns a boolean if a field has been set.

### GetEnableForceSyncJoinTokenToGmc

`func (o *GetGridResponse) GetEnableForceSyncJoinTokenToGmc() bool`

GetEnableForceSyncJoinTokenToGmc returns the EnableForceSyncJoinTokenToGmc field if non-nil, zero value otherwise.

### GetEnableForceSyncJoinTokenToGmcOk

`func (o *GetGridResponse) GetEnableForceSyncJoinTokenToGmcOk() (*bool, bool)`

GetEnableForceSyncJoinTokenToGmcOk returns a tuple with the EnableForceSyncJoinTokenToGmc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableForceSyncJoinTokenToGmc

`func (o *GetGridResponse) SetEnableForceSyncJoinTokenToGmc(v bool)`

SetEnableForceSyncJoinTokenToGmc sets EnableForceSyncJoinTokenToGmc field to given value.

### HasEnableForceSyncJoinTokenToGmc

`func (o *GetGridResponse) HasEnableForceSyncJoinTokenToGmc() bool`

HasEnableForceSyncJoinTokenToGmc returns a boolean if a field has been set.

### GetEnableGuiApiForLanVip

`func (o *GetGridResponse) GetEnableGuiApiForLanVip() bool`

GetEnableGuiApiForLanVip returns the EnableGuiApiForLanVip field if non-nil, zero value otherwise.

### GetEnableGuiApiForLanVipOk

`func (o *GetGridResponse) GetEnableGuiApiForLanVipOk() (*bool, bool)`

GetEnableGuiApiForLanVipOk returns a tuple with the EnableGuiApiForLanVip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableGuiApiForLanVip

`func (o *GetGridResponse) SetEnableGuiApiForLanVip(v bool)`

SetEnableGuiApiForLanVip sets EnableGuiApiForLanVip field to given value.

### HasEnableGuiApiForLanVip

`func (o *GetGridResponse) HasEnableGuiApiForLanVip() bool`

HasEnableGuiApiForLanVip returns a boolean if a field has been set.

### GetEnableLom

`func (o *GetGridResponse) GetEnableLom() bool`

GetEnableLom returns the EnableLom field if non-nil, zero value otherwise.

### GetEnableLomOk

`func (o *GetGridResponse) GetEnableLomOk() (*bool, bool)`

GetEnableLomOk returns a tuple with the EnableLom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLom

`func (o *GetGridResponse) SetEnableLom(v bool)`

SetEnableLom sets EnableLom field to given value.

### HasEnableLom

`func (o *GetGridResponse) HasEnableLom() bool`

HasEnableLom returns a boolean if a field has been set.

### GetEnableMemberRedirect

`func (o *GetGridResponse) GetEnableMemberRedirect() bool`

GetEnableMemberRedirect returns the EnableMemberRedirect field if non-nil, zero value otherwise.

### GetEnableMemberRedirectOk

`func (o *GetGridResponse) GetEnableMemberRedirectOk() (*bool, bool)`

GetEnableMemberRedirectOk returns a tuple with the EnableMemberRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMemberRedirect

`func (o *GetGridResponse) SetEnableMemberRedirect(v bool)`

SetEnableMemberRedirect sets EnableMemberRedirect field to given value.

### HasEnableMemberRedirect

`func (o *GetGridResponse) HasEnableMemberRedirect() bool`

HasEnableMemberRedirect returns a boolean if a field has been set.

### GetEnableRecycleBin

`func (o *GetGridResponse) GetEnableRecycleBin() bool`

GetEnableRecycleBin returns the EnableRecycleBin field if non-nil, zero value otherwise.

### GetEnableRecycleBinOk

`func (o *GetGridResponse) GetEnableRecycleBinOk() (*bool, bool)`

GetEnableRecycleBinOk returns a tuple with the EnableRecycleBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecycleBin

`func (o *GetGridResponse) SetEnableRecycleBin(v bool)`

SetEnableRecycleBin sets EnableRecycleBin field to given value.

### HasEnableRecycleBin

`func (o *GetGridResponse) HasEnableRecycleBin() bool`

HasEnableRecycleBin returns a boolean if a field has been set.

### GetEnableRirSwip

`func (o *GetGridResponse) GetEnableRirSwip() bool`

GetEnableRirSwip returns the EnableRirSwip field if non-nil, zero value otherwise.

### GetEnableRirSwipOk

`func (o *GetGridResponse) GetEnableRirSwipOk() (*bool, bool)`

GetEnableRirSwipOk returns a tuple with the EnableRirSwip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRirSwip

`func (o *GetGridResponse) SetEnableRirSwip(v bool)`

SetEnableRirSwip sets EnableRirSwip field to given value.

### HasEnableRirSwip

`func (o *GetGridResponse) HasEnableRirSwip() bool`

HasEnableRirSwip returns a boolean if a field has been set.

### GetExternalSyslogBackupServers

`func (o *GetGridResponse) GetExternalSyslogBackupServers() []GridExternalSyslogBackupServers`

GetExternalSyslogBackupServers returns the ExternalSyslogBackupServers field if non-nil, zero value otherwise.

### GetExternalSyslogBackupServersOk

`func (o *GetGridResponse) GetExternalSyslogBackupServersOk() (*[]GridExternalSyslogBackupServers, bool)`

GetExternalSyslogBackupServersOk returns a tuple with the ExternalSyslogBackupServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSyslogBackupServers

`func (o *GetGridResponse) SetExternalSyslogBackupServers(v []GridExternalSyslogBackupServers)`

SetExternalSyslogBackupServers sets ExternalSyslogBackupServers field to given value.

### HasExternalSyslogBackupServers

`func (o *GetGridResponse) HasExternalSyslogBackupServers() bool`

HasExternalSyslogBackupServers returns a boolean if a field has been set.

### GetExternalSyslogServerEnable

`func (o *GetGridResponse) GetExternalSyslogServerEnable() bool`

GetExternalSyslogServerEnable returns the ExternalSyslogServerEnable field if non-nil, zero value otherwise.

### GetExternalSyslogServerEnableOk

`func (o *GetGridResponse) GetExternalSyslogServerEnableOk() (*bool, bool)`

GetExternalSyslogServerEnableOk returns a tuple with the ExternalSyslogServerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSyslogServerEnable

`func (o *GetGridResponse) SetExternalSyslogServerEnable(v bool)`

SetExternalSyslogServerEnable sets ExternalSyslogServerEnable field to given value.

### HasExternalSyslogServerEnable

`func (o *GetGridResponse) HasExternalSyslogServerEnable() bool`

HasExternalSyslogServerEnable returns a boolean if a field has been set.

### GetHttpProxyServerSetting

`func (o *GetGridResponse) GetHttpProxyServerSetting() GridHttpProxyServerSetting`

GetHttpProxyServerSetting returns the HttpProxyServerSetting field if non-nil, zero value otherwise.

### GetHttpProxyServerSettingOk

`func (o *GetGridResponse) GetHttpProxyServerSettingOk() (*GridHttpProxyServerSetting, bool)`

GetHttpProxyServerSettingOk returns a tuple with the HttpProxyServerSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpProxyServerSetting

`func (o *GetGridResponse) SetHttpProxyServerSetting(v GridHttpProxyServerSetting)`

SetHttpProxyServerSetting sets HttpProxyServerSetting field to given value.

### HasHttpProxyServerSetting

`func (o *GetGridResponse) HasHttpProxyServerSetting() bool`

HasHttpProxyServerSetting returns a boolean if a field has been set.

### GetInformationalBannerSetting

`func (o *GetGridResponse) GetInformationalBannerSetting() GridInformationalBannerSetting`

GetInformationalBannerSetting returns the InformationalBannerSetting field if non-nil, zero value otherwise.

### GetInformationalBannerSettingOk

`func (o *GetGridResponse) GetInformationalBannerSettingOk() (*GridInformationalBannerSetting, bool)`

GetInformationalBannerSettingOk returns a tuple with the InformationalBannerSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformationalBannerSetting

`func (o *GetGridResponse) SetInformationalBannerSetting(v GridInformationalBannerSetting)`

SetInformationalBannerSetting sets InformationalBannerSetting field to given value.

### HasInformationalBannerSetting

`func (o *GetGridResponse) HasInformationalBannerSetting() bool`

HasInformationalBannerSetting returns a boolean if a field has been set.

### GetIsGridVisualizationVisible

`func (o *GetGridResponse) GetIsGridVisualizationVisible() bool`

GetIsGridVisualizationVisible returns the IsGridVisualizationVisible field if non-nil, zero value otherwise.

### GetIsGridVisualizationVisibleOk

`func (o *GetGridResponse) GetIsGridVisualizationVisibleOk() (*bool, bool)`

GetIsGridVisualizationVisibleOk returns a tuple with the IsGridVisualizationVisible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGridVisualizationVisible

`func (o *GetGridResponse) SetIsGridVisualizationVisible(v bool)`

SetIsGridVisualizationVisible sets IsGridVisualizationVisible field to given value.

### HasIsGridVisualizationVisible

`func (o *GetGridResponse) HasIsGridVisualizationVisible() bool`

HasIsGridVisualizationVisible returns a boolean if a field has been set.

### GetLockoutSetting

`func (o *GetGridResponse) GetLockoutSetting() GridLockoutSetting`

GetLockoutSetting returns the LockoutSetting field if non-nil, zero value otherwise.

### GetLockoutSettingOk

`func (o *GetGridResponse) GetLockoutSettingOk() (*GridLockoutSetting, bool)`

GetLockoutSettingOk returns a tuple with the LockoutSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockoutSetting

`func (o *GetGridResponse) SetLockoutSetting(v GridLockoutSetting)`

SetLockoutSetting sets LockoutSetting field to given value.

### HasLockoutSetting

`func (o *GetGridResponse) HasLockoutSetting() bool`

HasLockoutSetting returns a boolean if a field has been set.

### GetLomUsers

`func (o *GetGridResponse) GetLomUsers() []GridLomUsers`

GetLomUsers returns the LomUsers field if non-nil, zero value otherwise.

### GetLomUsersOk

`func (o *GetGridResponse) GetLomUsersOk() (*[]GridLomUsers, bool)`

GetLomUsersOk returns a tuple with the LomUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomUsers

`func (o *GetGridResponse) SetLomUsers(v []GridLomUsers)`

SetLomUsers sets LomUsers field to given value.

### HasLomUsers

`func (o *GetGridResponse) HasLomUsers() bool`

HasLomUsers returns a boolean if a field has been set.

### GetMgmStrictDelegateMode

`func (o *GetGridResponse) GetMgmStrictDelegateMode() bool`

GetMgmStrictDelegateMode returns the MgmStrictDelegateMode field if non-nil, zero value otherwise.

### GetMgmStrictDelegateModeOk

`func (o *GetGridResponse) GetMgmStrictDelegateModeOk() (*bool, bool)`

GetMgmStrictDelegateModeOk returns a tuple with the MgmStrictDelegateMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmStrictDelegateMode

`func (o *GetGridResponse) SetMgmStrictDelegateMode(v bool)`

SetMgmStrictDelegateMode sets MgmStrictDelegateMode field to given value.

### HasMgmStrictDelegateMode

`func (o *GetGridResponse) HasMgmStrictDelegateMode() bool`

HasMgmStrictDelegateMode returns a boolean if a field has been set.

### GetMsSetting

`func (o *GetGridResponse) GetMsSetting() GridMsSetting`

GetMsSetting returns the MsSetting field if non-nil, zero value otherwise.

### GetMsSettingOk

`func (o *GetGridResponse) GetMsSettingOk() (*GridMsSetting, bool)`

GetMsSettingOk returns a tuple with the MsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSetting

`func (o *GetGridResponse) SetMsSetting(v GridMsSetting)`

SetMsSetting sets MsSetting field to given value.

### HasMsSetting

`func (o *GetGridResponse) HasMsSetting() bool`

HasMsSetting returns a boolean if a field has been set.

### GetName

`func (o *GetGridResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGridResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGridResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGridResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNatGroups

`func (o *GetGridResponse) GetNatGroups() []string`

GetNatGroups returns the NatGroups field if non-nil, zero value otherwise.

### GetNatGroupsOk

`func (o *GetGridResponse) GetNatGroupsOk() (*[]string, bool)`

GetNatGroupsOk returns a tuple with the NatGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatGroups

`func (o *GetGridResponse) SetNatGroups(v []string)`

SetNatGroups sets NatGroups field to given value.

### HasNatGroups

`func (o *GetGridResponse) HasNatGroups() bool`

HasNatGroups returns a boolean if a field has been set.

### GetNtpSetting

`func (o *GetGridResponse) GetNtpSetting() GridNtpSetting`

GetNtpSetting returns the NtpSetting field if non-nil, zero value otherwise.

### GetNtpSettingOk

`func (o *GetGridResponse) GetNtpSettingOk() (*GridNtpSetting, bool)`

GetNtpSettingOk returns a tuple with the NtpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpSetting

`func (o *GetGridResponse) SetNtpSetting(v GridNtpSetting)`

SetNtpSetting sets NtpSetting field to given value.

### HasNtpSetting

`func (o *GetGridResponse) HasNtpSetting() bool`

HasNtpSetting returns a boolean if a field has been set.

### GetObjectsChangesTrackingSetting

`func (o *GetGridResponse) GetObjectsChangesTrackingSetting() GridObjectsChangesTrackingSetting`

GetObjectsChangesTrackingSetting returns the ObjectsChangesTrackingSetting field if non-nil, zero value otherwise.

### GetObjectsChangesTrackingSettingOk

`func (o *GetGridResponse) GetObjectsChangesTrackingSettingOk() (*GridObjectsChangesTrackingSetting, bool)`

GetObjectsChangesTrackingSettingOk returns a tuple with the ObjectsChangesTrackingSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectsChangesTrackingSetting

`func (o *GetGridResponse) SetObjectsChangesTrackingSetting(v GridObjectsChangesTrackingSetting)`

SetObjectsChangesTrackingSetting sets ObjectsChangesTrackingSetting field to given value.

### HasObjectsChangesTrackingSetting

`func (o *GetGridResponse) HasObjectsChangesTrackingSetting() bool`

HasObjectsChangesTrackingSetting returns a boolean if a field has been set.

### GetPasswordSetting

`func (o *GetGridResponse) GetPasswordSetting() GridPasswordSetting`

GetPasswordSetting returns the PasswordSetting field if non-nil, zero value otherwise.

### GetPasswordSettingOk

`func (o *GetGridResponse) GetPasswordSettingOk() (*GridPasswordSetting, bool)`

GetPasswordSettingOk returns a tuple with the PasswordSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordSetting

`func (o *GetGridResponse) SetPasswordSetting(v GridPasswordSetting)`

SetPasswordSetting sets PasswordSetting field to given value.

### HasPasswordSetting

`func (o *GetGridResponse) HasPasswordSetting() bool`

HasPasswordSetting returns a boolean if a field has been set.

### GetRestartBannerSetting

`func (o *GetGridResponse) GetRestartBannerSetting() GridRestartBannerSetting`

GetRestartBannerSetting returns the RestartBannerSetting field if non-nil, zero value otherwise.

### GetRestartBannerSettingOk

`func (o *GetGridResponse) GetRestartBannerSettingOk() (*GridRestartBannerSetting, bool)`

GetRestartBannerSettingOk returns a tuple with the RestartBannerSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartBannerSetting

`func (o *GetGridResponse) SetRestartBannerSetting(v GridRestartBannerSetting)`

SetRestartBannerSetting sets RestartBannerSetting field to given value.

### HasRestartBannerSetting

`func (o *GetGridResponse) HasRestartBannerSetting() bool`

HasRestartBannerSetting returns a boolean if a field has been set.

### GetRestartStatus

`func (o *GetGridResponse) GetRestartStatus() string`

GetRestartStatus returns the RestartStatus field if non-nil, zero value otherwise.

### GetRestartStatusOk

`func (o *GetGridResponse) GetRestartStatusOk() (*string, bool)`

GetRestartStatusOk returns a tuple with the RestartStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartStatus

`func (o *GetGridResponse) SetRestartStatus(v string)`

SetRestartStatus sets RestartStatus field to given value.

### HasRestartStatus

`func (o *GetGridResponse) HasRestartStatus() bool`

HasRestartStatus returns a boolean if a field has been set.

### GetRpzHitRateInterval

`func (o *GetGridResponse) GetRpzHitRateInterval() int64`

GetRpzHitRateInterval returns the RpzHitRateInterval field if non-nil, zero value otherwise.

### GetRpzHitRateIntervalOk

`func (o *GetGridResponse) GetRpzHitRateIntervalOk() (*int64, bool)`

GetRpzHitRateIntervalOk returns a tuple with the RpzHitRateInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzHitRateInterval

`func (o *GetGridResponse) SetRpzHitRateInterval(v int64)`

SetRpzHitRateInterval sets RpzHitRateInterval field to given value.

### HasRpzHitRateInterval

`func (o *GetGridResponse) HasRpzHitRateInterval() bool`

HasRpzHitRateInterval returns a boolean if a field has been set.

### GetRpzHitRateMaxQuery

`func (o *GetGridResponse) GetRpzHitRateMaxQuery() int64`

GetRpzHitRateMaxQuery returns the RpzHitRateMaxQuery field if non-nil, zero value otherwise.

### GetRpzHitRateMaxQueryOk

`func (o *GetGridResponse) GetRpzHitRateMaxQueryOk() (*int64, bool)`

GetRpzHitRateMaxQueryOk returns a tuple with the RpzHitRateMaxQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzHitRateMaxQuery

`func (o *GetGridResponse) SetRpzHitRateMaxQuery(v int64)`

SetRpzHitRateMaxQuery sets RpzHitRateMaxQuery field to given value.

### HasRpzHitRateMaxQuery

`func (o *GetGridResponse) HasRpzHitRateMaxQuery() bool`

HasRpzHitRateMaxQuery returns a boolean if a field has been set.

### GetRpzHitRateMinQuery

`func (o *GetGridResponse) GetRpzHitRateMinQuery() int64`

GetRpzHitRateMinQuery returns the RpzHitRateMinQuery field if non-nil, zero value otherwise.

### GetRpzHitRateMinQueryOk

`func (o *GetGridResponse) GetRpzHitRateMinQueryOk() (*int64, bool)`

GetRpzHitRateMinQueryOk returns a tuple with the RpzHitRateMinQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzHitRateMinQuery

`func (o *GetGridResponse) SetRpzHitRateMinQuery(v int64)`

SetRpzHitRateMinQuery sets RpzHitRateMinQuery field to given value.

### HasRpzHitRateMinQuery

`func (o *GetGridResponse) HasRpzHitRateMinQuery() bool`

HasRpzHitRateMinQuery returns a boolean if a field has been set.

### GetScheduledBackup

`func (o *GetGridResponse) GetScheduledBackup() GridScheduledBackup`

GetScheduledBackup returns the ScheduledBackup field if non-nil, zero value otherwise.

### GetScheduledBackupOk

`func (o *GetGridResponse) GetScheduledBackupOk() (*GridScheduledBackup, bool)`

GetScheduledBackupOk returns a tuple with the ScheduledBackup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledBackup

`func (o *GetGridResponse) SetScheduledBackup(v GridScheduledBackup)`

SetScheduledBackup sets ScheduledBackup field to given value.

### HasScheduledBackup

`func (o *GetGridResponse) HasScheduledBackup() bool`

HasScheduledBackup returns a boolean if a field has been set.

### GetSecret

`func (o *GetGridResponse) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GetGridResponse) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GetGridResponse) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GetGridResponse) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetSecurityBannerSetting

`func (o *GetGridResponse) GetSecurityBannerSetting() GridSecurityBannerSetting`

GetSecurityBannerSetting returns the SecurityBannerSetting field if non-nil, zero value otherwise.

### GetSecurityBannerSettingOk

`func (o *GetGridResponse) GetSecurityBannerSettingOk() (*GridSecurityBannerSetting, bool)`

GetSecurityBannerSettingOk returns a tuple with the SecurityBannerSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityBannerSetting

`func (o *GetGridResponse) SetSecurityBannerSetting(v GridSecurityBannerSetting)`

SetSecurityBannerSetting sets SecurityBannerSetting field to given value.

### HasSecurityBannerSetting

`func (o *GetGridResponse) HasSecurityBannerSetting() bool`

HasSecurityBannerSetting returns a boolean if a field has been set.

### GetSecuritySetting

`func (o *GetGridResponse) GetSecuritySetting() GridSecuritySetting`

GetSecuritySetting returns the SecuritySetting field if non-nil, zero value otherwise.

### GetSecuritySettingOk

`func (o *GetGridResponse) GetSecuritySettingOk() (*GridSecuritySetting, bool)`

GetSecuritySettingOk returns a tuple with the SecuritySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecuritySetting

`func (o *GetGridResponse) SetSecuritySetting(v GridSecuritySetting)`

SetSecuritySetting sets SecuritySetting field to given value.

### HasSecuritySetting

`func (o *GetGridResponse) HasSecuritySetting() bool`

HasSecuritySetting returns a boolean if a field has been set.

### GetServiceStatus

`func (o *GetGridResponse) GetServiceStatus() string`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *GetGridResponse) GetServiceStatusOk() (*string, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *GetGridResponse) SetServiceStatus(v string)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *GetGridResponse) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.

### GetSnmpSetting

`func (o *GetGridResponse) GetSnmpSetting() GridSnmpSetting`

GetSnmpSetting returns the SnmpSetting field if non-nil, zero value otherwise.

### GetSnmpSettingOk

`func (o *GetGridResponse) GetSnmpSettingOk() (*GridSnmpSetting, bool)`

GetSnmpSettingOk returns a tuple with the SnmpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpSetting

`func (o *GetGridResponse) SetSnmpSetting(v GridSnmpSetting)`

SetSnmpSetting sets SnmpSetting field to given value.

### HasSnmpSetting

`func (o *GetGridResponse) HasSnmpSetting() bool`

HasSnmpSetting returns a boolean if a field has been set.

### GetSupportBundleDownloadTimeout

`func (o *GetGridResponse) GetSupportBundleDownloadTimeout() int64`

GetSupportBundleDownloadTimeout returns the SupportBundleDownloadTimeout field if non-nil, zero value otherwise.

### GetSupportBundleDownloadTimeoutOk

`func (o *GetGridResponse) GetSupportBundleDownloadTimeoutOk() (*int64, bool)`

GetSupportBundleDownloadTimeoutOk returns a tuple with the SupportBundleDownloadTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBundleDownloadTimeout

`func (o *GetGridResponse) SetSupportBundleDownloadTimeout(v int64)`

SetSupportBundleDownloadTimeout sets SupportBundleDownloadTimeout field to given value.

### HasSupportBundleDownloadTimeout

`func (o *GetGridResponse) HasSupportBundleDownloadTimeout() bool`

HasSupportBundleDownloadTimeout returns a boolean if a field has been set.

### GetSyslogFacility

`func (o *GetGridResponse) GetSyslogFacility() string`

GetSyslogFacility returns the SyslogFacility field if non-nil, zero value otherwise.

### GetSyslogFacilityOk

`func (o *GetGridResponse) GetSyslogFacilityOk() (*string, bool)`

GetSyslogFacilityOk returns a tuple with the SyslogFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogFacility

`func (o *GetGridResponse) SetSyslogFacility(v string)`

SetSyslogFacility sets SyslogFacility field to given value.

### HasSyslogFacility

`func (o *GetGridResponse) HasSyslogFacility() bool`

HasSyslogFacility returns a boolean if a field has been set.

### GetSyslogServers

`func (o *GetGridResponse) GetSyslogServers() []GridSyslogServers`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *GetGridResponse) GetSyslogServersOk() (*[]GridSyslogServers, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *GetGridResponse) SetSyslogServers(v []GridSyslogServers)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *GetGridResponse) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### GetSyslogSize

`func (o *GetGridResponse) GetSyslogSize() int64`

GetSyslogSize returns the SyslogSize field if non-nil, zero value otherwise.

### GetSyslogSizeOk

`func (o *GetGridResponse) GetSyslogSizeOk() (*int64, bool)`

GetSyslogSizeOk returns a tuple with the SyslogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSize

`func (o *GetGridResponse) SetSyslogSize(v int64)`

SetSyslogSize sets SyslogSize field to given value.

### HasSyslogSize

`func (o *GetGridResponse) HasSyslogSize() bool`

HasSyslogSize returns a boolean if a field has been set.

### GetThresholdTraps

`func (o *GetGridResponse) GetThresholdTraps() []GridThresholdTraps`

GetThresholdTraps returns the ThresholdTraps field if non-nil, zero value otherwise.

### GetThresholdTrapsOk

`func (o *GetGridResponse) GetThresholdTrapsOk() (*[]GridThresholdTraps, bool)`

GetThresholdTrapsOk returns a tuple with the ThresholdTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdTraps

`func (o *GetGridResponse) SetThresholdTraps(v []GridThresholdTraps)`

SetThresholdTraps sets ThresholdTraps field to given value.

### HasThresholdTraps

`func (o *GetGridResponse) HasThresholdTraps() bool`

HasThresholdTraps returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetGridResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetGridResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetGridResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetGridResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTokenUsageDelay

`func (o *GetGridResponse) GetTokenUsageDelay() int64`

GetTokenUsageDelay returns the TokenUsageDelay field if non-nil, zero value otherwise.

### GetTokenUsageDelayOk

`func (o *GetGridResponse) GetTokenUsageDelayOk() (*int64, bool)`

GetTokenUsageDelayOk returns a tuple with the TokenUsageDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUsageDelay

`func (o *GetGridResponse) SetTokenUsageDelay(v int64)`

SetTokenUsageDelay sets TokenUsageDelay field to given value.

### HasTokenUsageDelay

`func (o *GetGridResponse) HasTokenUsageDelay() bool`

HasTokenUsageDelay returns a boolean if a field has been set.

### GetTrafficCaptureAuthDnsSetting

`func (o *GetGridResponse) GetTrafficCaptureAuthDnsSetting() GridTrafficCaptureAuthDnsSetting`

GetTrafficCaptureAuthDnsSetting returns the TrafficCaptureAuthDnsSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureAuthDnsSettingOk

`func (o *GetGridResponse) GetTrafficCaptureAuthDnsSettingOk() (*GridTrafficCaptureAuthDnsSetting, bool)`

GetTrafficCaptureAuthDnsSettingOk returns a tuple with the TrafficCaptureAuthDnsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureAuthDnsSetting

`func (o *GetGridResponse) SetTrafficCaptureAuthDnsSetting(v GridTrafficCaptureAuthDnsSetting)`

SetTrafficCaptureAuthDnsSetting sets TrafficCaptureAuthDnsSetting field to given value.

### HasTrafficCaptureAuthDnsSetting

`func (o *GetGridResponse) HasTrafficCaptureAuthDnsSetting() bool`

HasTrafficCaptureAuthDnsSetting returns a boolean if a field has been set.

### GetTrafficCaptureChrSetting

`func (o *GetGridResponse) GetTrafficCaptureChrSetting() GridTrafficCaptureChrSetting`

GetTrafficCaptureChrSetting returns the TrafficCaptureChrSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureChrSettingOk

`func (o *GetGridResponse) GetTrafficCaptureChrSettingOk() (*GridTrafficCaptureChrSetting, bool)`

GetTrafficCaptureChrSettingOk returns a tuple with the TrafficCaptureChrSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureChrSetting

`func (o *GetGridResponse) SetTrafficCaptureChrSetting(v GridTrafficCaptureChrSetting)`

SetTrafficCaptureChrSetting sets TrafficCaptureChrSetting field to given value.

### HasTrafficCaptureChrSetting

`func (o *GetGridResponse) HasTrafficCaptureChrSetting() bool`

HasTrafficCaptureChrSetting returns a boolean if a field has been set.

### GetTrafficCaptureQpsSetting

`func (o *GetGridResponse) GetTrafficCaptureQpsSetting() GridTrafficCaptureQpsSetting`

GetTrafficCaptureQpsSetting returns the TrafficCaptureQpsSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureQpsSettingOk

`func (o *GetGridResponse) GetTrafficCaptureQpsSettingOk() (*GridTrafficCaptureQpsSetting, bool)`

GetTrafficCaptureQpsSettingOk returns a tuple with the TrafficCaptureQpsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureQpsSetting

`func (o *GetGridResponse) SetTrafficCaptureQpsSetting(v GridTrafficCaptureQpsSetting)`

SetTrafficCaptureQpsSetting sets TrafficCaptureQpsSetting field to given value.

### HasTrafficCaptureQpsSetting

`func (o *GetGridResponse) HasTrafficCaptureQpsSetting() bool`

HasTrafficCaptureQpsSetting returns a boolean if a field has been set.

### GetTrafficCaptureRecDnsSetting

`func (o *GetGridResponse) GetTrafficCaptureRecDnsSetting() GridTrafficCaptureRecDnsSetting`

GetTrafficCaptureRecDnsSetting returns the TrafficCaptureRecDnsSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureRecDnsSettingOk

`func (o *GetGridResponse) GetTrafficCaptureRecDnsSettingOk() (*GridTrafficCaptureRecDnsSetting, bool)`

GetTrafficCaptureRecDnsSettingOk returns a tuple with the TrafficCaptureRecDnsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureRecDnsSetting

`func (o *GetGridResponse) SetTrafficCaptureRecDnsSetting(v GridTrafficCaptureRecDnsSetting)`

SetTrafficCaptureRecDnsSetting sets TrafficCaptureRecDnsSetting field to given value.

### HasTrafficCaptureRecDnsSetting

`func (o *GetGridResponse) HasTrafficCaptureRecDnsSetting() bool`

HasTrafficCaptureRecDnsSetting returns a boolean if a field has been set.

### GetTrafficCaptureRecQueriesSetting

`func (o *GetGridResponse) GetTrafficCaptureRecQueriesSetting() GridTrafficCaptureRecQueriesSetting`

GetTrafficCaptureRecQueriesSetting returns the TrafficCaptureRecQueriesSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureRecQueriesSettingOk

`func (o *GetGridResponse) GetTrafficCaptureRecQueriesSettingOk() (*GridTrafficCaptureRecQueriesSetting, bool)`

GetTrafficCaptureRecQueriesSettingOk returns a tuple with the TrafficCaptureRecQueriesSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureRecQueriesSetting

`func (o *GetGridResponse) SetTrafficCaptureRecQueriesSetting(v GridTrafficCaptureRecQueriesSetting)`

SetTrafficCaptureRecQueriesSetting sets TrafficCaptureRecQueriesSetting field to given value.

### HasTrafficCaptureRecQueriesSetting

`func (o *GetGridResponse) HasTrafficCaptureRecQueriesSetting() bool`

HasTrafficCaptureRecQueriesSetting returns a boolean if a field has been set.

### GetTrapNotifications

`func (o *GetGridResponse) GetTrapNotifications() []GridTrapNotifications`

GetTrapNotifications returns the TrapNotifications field if non-nil, zero value otherwise.

### GetTrapNotificationsOk

`func (o *GetGridResponse) GetTrapNotificationsOk() (*[]GridTrapNotifications, bool)`

GetTrapNotificationsOk returns a tuple with the TrapNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapNotifications

`func (o *GetGridResponse) SetTrapNotifications(v []GridTrapNotifications)`

SetTrapNotifications sets TrapNotifications field to given value.

### HasTrapNotifications

`func (o *GetGridResponse) HasTrapNotifications() bool`

HasTrapNotifications returns a boolean if a field has been set.

### GetUpdatesDownloadMemberConfig

`func (o *GetGridResponse) GetUpdatesDownloadMemberConfig() []GridUpdatesDownloadMemberConfig`

GetUpdatesDownloadMemberConfig returns the UpdatesDownloadMemberConfig field if non-nil, zero value otherwise.

### GetUpdatesDownloadMemberConfigOk

`func (o *GetGridResponse) GetUpdatesDownloadMemberConfigOk() (*[]GridUpdatesDownloadMemberConfig, bool)`

GetUpdatesDownloadMemberConfigOk returns a tuple with the UpdatesDownloadMemberConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatesDownloadMemberConfig

`func (o *GetGridResponse) SetUpdatesDownloadMemberConfig(v []GridUpdatesDownloadMemberConfig)`

SetUpdatesDownloadMemberConfig sets UpdatesDownloadMemberConfig field to given value.

### HasUpdatesDownloadMemberConfig

`func (o *GetGridResponse) HasUpdatesDownloadMemberConfig() bool`

HasUpdatesDownloadMemberConfig returns a boolean if a field has been set.

### GetVpnPort

`func (o *GetGridResponse) GetVpnPort() int64`

GetVpnPort returns the VpnPort field if non-nil, zero value otherwise.

### GetVpnPortOk

`func (o *GetGridResponse) GetVpnPortOk() (*int64, bool)`

GetVpnPortOk returns a tuple with the VpnPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnPort

`func (o *GetGridResponse) SetVpnPort(v int64)`

SetVpnPort sets VpnPort field to given value.

### HasVpnPort

`func (o *GetGridResponse) HasVpnPort() bool`

HasVpnPort returns a boolean if a field has been set.

### GetResult

`func (o *GetGridResponse) GetResult() Grid`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridResponse) GetResultOk() (*Grid, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridResponse) SetResult(v Grid)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


