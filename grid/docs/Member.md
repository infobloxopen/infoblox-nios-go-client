# Member

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ActivePosition** | Pointer to **string** | The active server of a Grid member. | [optional] [readonly] 
**AdditionalIpList** | Pointer to [**[]MemberAdditionalIpList**](MemberAdditionalIpList.md) | The additional IP list of a Grid member. This list contains additional interface information that can be used at the member level. Note that interface structure(s) with interface type set to &#39;MGMT&#39; are not supported. | [optional] 
**AutomatedTrafficCaptureSetting** | Pointer to [**MemberAutomatedTrafficCaptureSetting**](MemberAutomatedTrafficCaptureSetting.md) |  | [optional] 
**BgpAs** | Pointer to [**[]MemberBgpAs**](MemberBgpAs.md) | The BGP configuration for anycast for a Grid member. | [optional] 
**Comment** | Pointer to **string** | A descriptive comment of the Grid member. | [optional] 
**ConfigAddrType** | Pointer to **string** | Address configuration type. | [optional] 
**CspAccessKey** | Pointer to **[]string** | CSP portal on-prem host access key | [optional] 
**CspMemberSetting** | Pointer to [**MemberCspMemberSetting**](MemberCspMemberSetting.md) |  | [optional] 
**DnsResolverSetting** | Pointer to [**MemberDnsResolverSetting**](MemberDnsResolverSetting.md) |  | [optional] 
**Dscp** | Pointer to **int64** | The DSCP (Differentiated Services Code Point) value. | [optional] 
**EmailSetting** | Pointer to [**MemberEmailSetting**](MemberEmailSetting.md) |  | [optional] 
**EnableHa** | Pointer to **bool** | If set to True, the member has two physical nodes (HA pair). | [optional] 
**EnableLom** | Pointer to **bool** | Determines if the LOM functionality is enabled or not. | [optional] 
**EnableMemberRedirect** | Pointer to **bool** | Determines if the member will redirect GUI connections to the Grid Master or not. | [optional] 
**EnableRoApiAccess** | Pointer to **bool** | If set to True and the member object is a Grid Master Candidate, then read-only API access is enabled. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExternalSyslogBackupServers** | Pointer to [**[]MemberExternalSyslogBackupServers**](MemberExternalSyslogBackupServers.md) | The list of external syslog backup servers. | [optional] 
**ExternalSyslogServerEnable** | Pointer to **bool** | Determines if external syslog servers should be enabled. | [optional] 
**HaCloudPlatform** | Pointer to **string** | Cloud platform for HA. | [optional] 
**HaOnCloud** | Pointer to **bool** | True: HA on cloud. False: HA not on cloud. | [optional] 
**HostName** | Pointer to **string** | The host name of the Grid member. | [optional] 
**Ipv6Setting** | Pointer to [**MemberIpv6Setting**](MemberIpv6Setting.md) |  | [optional] 
**Ipv6StaticRoutes** | Pointer to [**[]MemberIpv6StaticRoutes**](MemberIpv6StaticRoutes.md) | List of IPv6 static routes. | [optional] 
**IsDscpCapable** | Pointer to **bool** | Determines if a Grid member supports DSCP (Differentiated Services Code Point). | [optional] [readonly] 
**Lan2Enabled** | Pointer to **bool** | If this is set to \&quot;true\&quot;, the LAN2 port is enabled as an independent port or as a port for failover purposes. | [optional] 
**Lan2PortSetting** | Pointer to [**MemberLan2PortSetting**](MemberLan2PortSetting.md) |  | [optional] 
**LomNetworkConfig** | Pointer to [**[]MemberLomNetworkConfig**](MemberLomNetworkConfig.md) | The Network configurations for LOM. | [optional] 
**LomUsers** | Pointer to [**[]MemberLomUsers**](MemberLomUsers.md) | The list of LOM users. | [optional] 
**MasterCandidate** | Pointer to **bool** | Determines if a Grid member is a Grid Master Candidate or not. This flag enables the Grid member to assume the role of the Grid Master as a disaster recovery measure. | [optional] 
**MemberServiceCommunication** | Pointer to [**[]MemberMemberServiceCommunication**](MemberMemberServiceCommunication.md) | Configure communication type for various services. | [optional] 
**MgmtPortSetting** | Pointer to [**MemberMgmtPortSetting**](MemberMgmtPortSetting.md) |  | [optional] 
**MmdbEaBuildTime** | Pointer to **int64** | Extensible attributes Topology database build time. | [optional] [readonly] 
**MmdbGeoipBuildTime** | Pointer to **int64** | GeoIP Topology database build time. | [optional] [readonly] 
**NatSetting** | Pointer to [**MemberNatSetting**](MemberNatSetting.md) |  | [optional] 
**NodeInfo** | Pointer to [**[]MemberNodeInfo**](MemberNodeInfo.md) | The node information list with detailed status report on the operations of the Grid Member, mgmt_port_setting must be enabled when configuring the MGMT Port using the node_info field. | [optional] 
**NtpSetting** | Pointer to [**MemberNtpSetting**](MemberNtpSetting.md) |  | [optional] 
**OspfList** | Pointer to [**[]MemberOspfList**](MemberOspfList.md) | The OSPF area configuration (for anycast) list for a Grid member. | [optional] 
**PassiveHaArpEnabled** | Pointer to **bool** | The ARP protocol setting on the passive node of an HA pair. If you do not specify a value, the default value is \&quot;false\&quot;. You can only set this value to \&quot;true\&quot; if the member is an HA pair. | [optional] 
**Platform** | Pointer to **string** | Hardware Platform. | [optional] 
**PreProvisioning** | Pointer to [**MemberPreProvisioning**](MemberPreProvisioning.md) |  | [optional] 
**PreserveIfOwnsDelegation** | Pointer to **bool** | Set this flag to \&quot;true\&quot; to prevent the deletion of the member if any delegated object remains attached to it. | [optional] 
**RemoteConsoleAccessEnable** | Pointer to **bool** | If set to True, superuser admins can access the Infoblox CLI from a remote location using an SSH (Secure Shell) v2 client. | [optional] 
**RouterId** | Pointer to **int64** | Virutal router identifier. Provide this ID if \&quot;ha_enabled\&quot; is set to \&quot;true\&quot;. This is a unique VRID number (from 1 to 255) for the local subnet. | [optional] 
**ServiceStatus** | Pointer to [**[]MemberServiceStatus**](MemberServiceStatus.md) | The service status list of a grid member. | [optional] [readonly] 
**ServiceTypeConfiguration** | Pointer to **string** | Configure all services to the given type. | [optional] 
**SnmpSetting** | Pointer to [**MemberSnmpSetting**](MemberSnmpSetting.md) |  | [optional] 
**StaticRoutes** | Pointer to [**[]MemberStaticRoutes**](MemberStaticRoutes.md) | List of static routes. | [optional] 
**SupportAccessEnable** | Pointer to **bool** | Determines if support access for the Grid member should be enabled. | [optional] 
**SupportAccessInfo** | Pointer to **string** | The information string for support access. | [optional] [readonly] 
**SyslogProxySetting** | Pointer to [**MemberSyslogProxySetting**](MemberSyslogProxySetting.md) |  | [optional] 
**SyslogServers** | Pointer to [**[]MemberSyslogServers**](MemberSyslogServers.md) | The list of external syslog servers. | [optional] 
**SyslogSize** | Pointer to **int64** | The maximum size for the syslog file expressed in megabytes. | [optional] 
**ThresholdTraps** | Pointer to [**[]MemberThresholdTraps**](MemberThresholdTraps.md) | Determines the list of threshold traps. The user can only change the values for each trap or remove traps. | [optional] 
**TimeZone** | Pointer to **string** | The time zone of the Grid member. The UTC string that represents the time zone, such as \&quot;Asia/Kolkata\&quot;. | [optional] 
**TrafficCaptureAuthDnsSetting** | Pointer to [**MemberTrafficCaptureAuthDnsSetting**](MemberTrafficCaptureAuthDnsSetting.md) |  | [optional] 
**TrafficCaptureChrSetting** | Pointer to [**MemberTrafficCaptureChrSetting**](MemberTrafficCaptureChrSetting.md) |  | [optional] 
**TrafficCaptureQpsSetting** | Pointer to [**MemberTrafficCaptureQpsSetting**](MemberTrafficCaptureQpsSetting.md) |  | [optional] 
**TrafficCaptureRecDnsSetting** | Pointer to [**MemberTrafficCaptureRecDnsSetting**](MemberTrafficCaptureRecDnsSetting.md) |  | [optional] 
**TrafficCaptureRecQueriesSetting** | Pointer to [**MemberTrafficCaptureRecQueriesSetting**](MemberTrafficCaptureRecQueriesSetting.md) |  | [optional] 
**TrapNotifications** | Pointer to [**[]MemberTrapNotifications**](MemberTrapNotifications.md) | Determines configuration of the trap notifications. | [optional] 
**UpgradeGroup** | Pointer to **string** | The name of the upgrade group to which this Grid member belongs. | [optional] 
**UseAutomatedTrafficCapture** | Pointer to **bool** | This flag is the use flag for enabling automated traffic capture based on DNS cache ratio thresholds. | [optional] 
**UseDnsResolverSetting** | Pointer to **bool** | Use flag for: dns_resolver_setting | [optional] 
**UseDscp** | Pointer to **bool** | Use flag for: dscp | [optional] 
**UseEmailSetting** | Pointer to **bool** | Use flag for: email_setting | [optional] 
**UseEnableLom** | Pointer to **bool** | Use flag for: enable_lom | [optional] 
**UseEnableMemberRedirect** | Pointer to **bool** | Use flag for: enable_member_redirect | [optional] 
**UseExternalSyslogBackupServers** | Pointer to **bool** | Use flag for: external_syslog_backup_servers | [optional] 
**UseRemoteConsoleAccessEnable** | Pointer to **bool** | Use flag for: remote_console_access_enable | [optional] 
**UseSnmpSetting** | Pointer to **bool** | Use flag for: snmp_setting | [optional] 
**UseSupportAccessEnable** | Pointer to **bool** | Use flag for: support_access_enable | [optional] 
**UseSyslogProxySetting** | Pointer to **bool** | Use flag for: external_syslog_server_enable , syslog_servers, syslog_proxy_setting, syslog_size | [optional] 
**UseThresholdTraps** | Pointer to **bool** | Use flag for: threshold_traps | [optional] 
**UseTimeZone** | Pointer to **bool** | Use flag for: time_zone | [optional] 
**UseTrafficCaptureAuthDns** | Pointer to **bool** | This flag is the use flag for enabling automated traffic capture based on authorative DNS latency. | [optional] 
**UseTrafficCaptureChr** | Pointer to **bool** | This flag is the use flag for automated traffic capture settings at member level. | [optional] 
**UseTrafficCaptureQps** | Pointer to **bool** | This flag is the use flag for enabling automated traffic capture based on DNS querie per second thresholds. | [optional] 
**UseTrafficCaptureRecDns** | Pointer to **bool** | This flag is the use flag for enabling automated traffic capture based on recursive DNS latency. | [optional] 
**UseTrafficCaptureRecQueries** | Pointer to **bool** | This flag is the use flag for enabling automated traffic capture based on outgoing recursive queries. | [optional] 
**UseTrapNotifications** | Pointer to **bool** | Use flag for: trap_notifications | [optional] 
**UseV4Vrrp** | Pointer to **bool** | Specify \&quot;true\&quot; to use VRRPv4 or \&quot;false\&quot; to use VRRPv6. | [optional] 
**VipSetting** | Pointer to [**MemberVipSetting**](MemberVipSetting.md) |  | [optional] 
**VpnMtu** | Pointer to **int64** | The VPN maximum transmission unit (MTU). | [optional] 

## Methods

### NewMember

`func NewMember() *Member`

NewMember instantiates a new Member object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberWithDefaults

`func NewMemberWithDefaults() *Member`

NewMemberWithDefaults instantiates a new Member object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Member) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Member) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Member) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Member) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActivePosition

`func (o *Member) GetActivePosition() string`

GetActivePosition returns the ActivePosition field if non-nil, zero value otherwise.

### GetActivePositionOk

`func (o *Member) GetActivePositionOk() (*string, bool)`

GetActivePositionOk returns a tuple with the ActivePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePosition

`func (o *Member) SetActivePosition(v string)`

SetActivePosition sets ActivePosition field to given value.

### HasActivePosition

`func (o *Member) HasActivePosition() bool`

HasActivePosition returns a boolean if a field has been set.

### GetAdditionalIpList

`func (o *Member) GetAdditionalIpList() []MemberAdditionalIpList`

GetAdditionalIpList returns the AdditionalIpList field if non-nil, zero value otherwise.

### GetAdditionalIpListOk

`func (o *Member) GetAdditionalIpListOk() (*[]MemberAdditionalIpList, bool)`

GetAdditionalIpListOk returns a tuple with the AdditionalIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIpList

`func (o *Member) SetAdditionalIpList(v []MemberAdditionalIpList)`

SetAdditionalIpList sets AdditionalIpList field to given value.

### HasAdditionalIpList

`func (o *Member) HasAdditionalIpList() bool`

HasAdditionalIpList returns a boolean if a field has been set.

### GetAutomatedTrafficCaptureSetting

`func (o *Member) GetAutomatedTrafficCaptureSetting() MemberAutomatedTrafficCaptureSetting`

GetAutomatedTrafficCaptureSetting returns the AutomatedTrafficCaptureSetting field if non-nil, zero value otherwise.

### GetAutomatedTrafficCaptureSettingOk

`func (o *Member) GetAutomatedTrafficCaptureSettingOk() (*MemberAutomatedTrafficCaptureSetting, bool)`

GetAutomatedTrafficCaptureSettingOk returns a tuple with the AutomatedTrafficCaptureSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedTrafficCaptureSetting

`func (o *Member) SetAutomatedTrafficCaptureSetting(v MemberAutomatedTrafficCaptureSetting)`

SetAutomatedTrafficCaptureSetting sets AutomatedTrafficCaptureSetting field to given value.

### HasAutomatedTrafficCaptureSetting

`func (o *Member) HasAutomatedTrafficCaptureSetting() bool`

HasAutomatedTrafficCaptureSetting returns a boolean if a field has been set.

### GetBgpAs

`func (o *Member) GetBgpAs() []MemberBgpAs`

GetBgpAs returns the BgpAs field if non-nil, zero value otherwise.

### GetBgpAsOk

`func (o *Member) GetBgpAsOk() (*[]MemberBgpAs, bool)`

GetBgpAsOk returns a tuple with the BgpAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAs

`func (o *Member) SetBgpAs(v []MemberBgpAs)`

SetBgpAs sets BgpAs field to given value.

### HasBgpAs

`func (o *Member) HasBgpAs() bool`

HasBgpAs returns a boolean if a field has been set.

### GetComment

`func (o *Member) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Member) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Member) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Member) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConfigAddrType

`func (o *Member) GetConfigAddrType() string`

GetConfigAddrType returns the ConfigAddrType field if non-nil, zero value otherwise.

### GetConfigAddrTypeOk

`func (o *Member) GetConfigAddrTypeOk() (*string, bool)`

GetConfigAddrTypeOk returns a tuple with the ConfigAddrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAddrType

`func (o *Member) SetConfigAddrType(v string)`

SetConfigAddrType sets ConfigAddrType field to given value.

### HasConfigAddrType

`func (o *Member) HasConfigAddrType() bool`

HasConfigAddrType returns a boolean if a field has been set.

### GetCspAccessKey

`func (o *Member) GetCspAccessKey() []string`

GetCspAccessKey returns the CspAccessKey field if non-nil, zero value otherwise.

### GetCspAccessKeyOk

`func (o *Member) GetCspAccessKeyOk() (*[]string, bool)`

GetCspAccessKeyOk returns a tuple with the CspAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspAccessKey

`func (o *Member) SetCspAccessKey(v []string)`

SetCspAccessKey sets CspAccessKey field to given value.

### HasCspAccessKey

`func (o *Member) HasCspAccessKey() bool`

HasCspAccessKey returns a boolean if a field has been set.

### GetCspMemberSetting

`func (o *Member) GetCspMemberSetting() MemberCspMemberSetting`

GetCspMemberSetting returns the CspMemberSetting field if non-nil, zero value otherwise.

### GetCspMemberSettingOk

`func (o *Member) GetCspMemberSettingOk() (*MemberCspMemberSetting, bool)`

GetCspMemberSettingOk returns a tuple with the CspMemberSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspMemberSetting

`func (o *Member) SetCspMemberSetting(v MemberCspMemberSetting)`

SetCspMemberSetting sets CspMemberSetting field to given value.

### HasCspMemberSetting

`func (o *Member) HasCspMemberSetting() bool`

HasCspMemberSetting returns a boolean if a field has been set.

### GetDnsResolverSetting

`func (o *Member) GetDnsResolverSetting() MemberDnsResolverSetting`

GetDnsResolverSetting returns the DnsResolverSetting field if non-nil, zero value otherwise.

### GetDnsResolverSettingOk

`func (o *Member) GetDnsResolverSettingOk() (*MemberDnsResolverSetting, bool)`

GetDnsResolverSettingOk returns a tuple with the DnsResolverSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolverSetting

`func (o *Member) SetDnsResolverSetting(v MemberDnsResolverSetting)`

SetDnsResolverSetting sets DnsResolverSetting field to given value.

### HasDnsResolverSetting

`func (o *Member) HasDnsResolverSetting() bool`

HasDnsResolverSetting returns a boolean if a field has been set.

### GetDscp

`func (o *Member) GetDscp() int64`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *Member) GetDscpOk() (*int64, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *Member) SetDscp(v int64)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *Member) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetEmailSetting

`func (o *Member) GetEmailSetting() MemberEmailSetting`

GetEmailSetting returns the EmailSetting field if non-nil, zero value otherwise.

### GetEmailSettingOk

`func (o *Member) GetEmailSettingOk() (*MemberEmailSetting, bool)`

GetEmailSettingOk returns a tuple with the EmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSetting

`func (o *Member) SetEmailSetting(v MemberEmailSetting)`

SetEmailSetting sets EmailSetting field to given value.

### HasEmailSetting

`func (o *Member) HasEmailSetting() bool`

HasEmailSetting returns a boolean if a field has been set.

### GetEnableHa

`func (o *Member) GetEnableHa() bool`

GetEnableHa returns the EnableHa field if non-nil, zero value otherwise.

### GetEnableHaOk

`func (o *Member) GetEnableHaOk() (*bool, bool)`

GetEnableHaOk returns a tuple with the EnableHa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHa

`func (o *Member) SetEnableHa(v bool)`

SetEnableHa sets EnableHa field to given value.

### HasEnableHa

`func (o *Member) HasEnableHa() bool`

HasEnableHa returns a boolean if a field has been set.

### GetEnableLom

`func (o *Member) GetEnableLom() bool`

GetEnableLom returns the EnableLom field if non-nil, zero value otherwise.

### GetEnableLomOk

`func (o *Member) GetEnableLomOk() (*bool, bool)`

GetEnableLomOk returns a tuple with the EnableLom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLom

`func (o *Member) SetEnableLom(v bool)`

SetEnableLom sets EnableLom field to given value.

### HasEnableLom

`func (o *Member) HasEnableLom() bool`

HasEnableLom returns a boolean if a field has been set.

### GetEnableMemberRedirect

`func (o *Member) GetEnableMemberRedirect() bool`

GetEnableMemberRedirect returns the EnableMemberRedirect field if non-nil, zero value otherwise.

### GetEnableMemberRedirectOk

`func (o *Member) GetEnableMemberRedirectOk() (*bool, bool)`

GetEnableMemberRedirectOk returns a tuple with the EnableMemberRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMemberRedirect

`func (o *Member) SetEnableMemberRedirect(v bool)`

SetEnableMemberRedirect sets EnableMemberRedirect field to given value.

### HasEnableMemberRedirect

`func (o *Member) HasEnableMemberRedirect() bool`

HasEnableMemberRedirect returns a boolean if a field has been set.

### GetEnableRoApiAccess

`func (o *Member) GetEnableRoApiAccess() bool`

GetEnableRoApiAccess returns the EnableRoApiAccess field if non-nil, zero value otherwise.

### GetEnableRoApiAccessOk

`func (o *Member) GetEnableRoApiAccessOk() (*bool, bool)`

GetEnableRoApiAccessOk returns a tuple with the EnableRoApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRoApiAccess

`func (o *Member) SetEnableRoApiAccess(v bool)`

SetEnableRoApiAccess sets EnableRoApiAccess field to given value.

### HasEnableRoApiAccess

`func (o *Member) HasEnableRoApiAccess() bool`

HasEnableRoApiAccess returns a boolean if a field has been set.

### GetExtattrs

`func (o *Member) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Member) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Member) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Member) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetExternalSyslogBackupServers

`func (o *Member) GetExternalSyslogBackupServers() []MemberExternalSyslogBackupServers`

GetExternalSyslogBackupServers returns the ExternalSyslogBackupServers field if non-nil, zero value otherwise.

### GetExternalSyslogBackupServersOk

`func (o *Member) GetExternalSyslogBackupServersOk() (*[]MemberExternalSyslogBackupServers, bool)`

GetExternalSyslogBackupServersOk returns a tuple with the ExternalSyslogBackupServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSyslogBackupServers

`func (o *Member) SetExternalSyslogBackupServers(v []MemberExternalSyslogBackupServers)`

SetExternalSyslogBackupServers sets ExternalSyslogBackupServers field to given value.

### HasExternalSyslogBackupServers

`func (o *Member) HasExternalSyslogBackupServers() bool`

HasExternalSyslogBackupServers returns a boolean if a field has been set.

### GetExternalSyslogServerEnable

`func (o *Member) GetExternalSyslogServerEnable() bool`

GetExternalSyslogServerEnable returns the ExternalSyslogServerEnable field if non-nil, zero value otherwise.

### GetExternalSyslogServerEnableOk

`func (o *Member) GetExternalSyslogServerEnableOk() (*bool, bool)`

GetExternalSyslogServerEnableOk returns a tuple with the ExternalSyslogServerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSyslogServerEnable

`func (o *Member) SetExternalSyslogServerEnable(v bool)`

SetExternalSyslogServerEnable sets ExternalSyslogServerEnable field to given value.

### HasExternalSyslogServerEnable

`func (o *Member) HasExternalSyslogServerEnable() bool`

HasExternalSyslogServerEnable returns a boolean if a field has been set.

### GetHaCloudPlatform

`func (o *Member) GetHaCloudPlatform() string`

GetHaCloudPlatform returns the HaCloudPlatform field if non-nil, zero value otherwise.

### GetHaCloudPlatformOk

`func (o *Member) GetHaCloudPlatformOk() (*string, bool)`

GetHaCloudPlatformOk returns a tuple with the HaCloudPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaCloudPlatform

`func (o *Member) SetHaCloudPlatform(v string)`

SetHaCloudPlatform sets HaCloudPlatform field to given value.

### HasHaCloudPlatform

`func (o *Member) HasHaCloudPlatform() bool`

HasHaCloudPlatform returns a boolean if a field has been set.

### GetHaOnCloud

`func (o *Member) GetHaOnCloud() bool`

GetHaOnCloud returns the HaOnCloud field if non-nil, zero value otherwise.

### GetHaOnCloudOk

`func (o *Member) GetHaOnCloudOk() (*bool, bool)`

GetHaOnCloudOk returns a tuple with the HaOnCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaOnCloud

`func (o *Member) SetHaOnCloud(v bool)`

SetHaOnCloud sets HaOnCloud field to given value.

### HasHaOnCloud

`func (o *Member) HasHaOnCloud() bool`

HasHaOnCloud returns a boolean if a field has been set.

### GetHostName

`func (o *Member) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *Member) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *Member) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *Member) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpv6Setting

`func (o *Member) GetIpv6Setting() MemberIpv6Setting`

GetIpv6Setting returns the Ipv6Setting field if non-nil, zero value otherwise.

### GetIpv6SettingOk

`func (o *Member) GetIpv6SettingOk() (*MemberIpv6Setting, bool)`

GetIpv6SettingOk returns a tuple with the Ipv6Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Setting

`func (o *Member) SetIpv6Setting(v MemberIpv6Setting)`

SetIpv6Setting sets Ipv6Setting field to given value.

### HasIpv6Setting

`func (o *Member) HasIpv6Setting() bool`

HasIpv6Setting returns a boolean if a field has been set.

### GetIpv6StaticRoutes

`func (o *Member) GetIpv6StaticRoutes() []MemberIpv6StaticRoutes`

GetIpv6StaticRoutes returns the Ipv6StaticRoutes field if non-nil, zero value otherwise.

### GetIpv6StaticRoutesOk

`func (o *Member) GetIpv6StaticRoutesOk() (*[]MemberIpv6StaticRoutes, bool)`

GetIpv6StaticRoutesOk returns a tuple with the Ipv6StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6StaticRoutes

`func (o *Member) SetIpv6StaticRoutes(v []MemberIpv6StaticRoutes)`

SetIpv6StaticRoutes sets Ipv6StaticRoutes field to given value.

### HasIpv6StaticRoutes

`func (o *Member) HasIpv6StaticRoutes() bool`

HasIpv6StaticRoutes returns a boolean if a field has been set.

### GetIsDscpCapable

`func (o *Member) GetIsDscpCapable() bool`

GetIsDscpCapable returns the IsDscpCapable field if non-nil, zero value otherwise.

### GetIsDscpCapableOk

`func (o *Member) GetIsDscpCapableOk() (*bool, bool)`

GetIsDscpCapableOk returns a tuple with the IsDscpCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDscpCapable

`func (o *Member) SetIsDscpCapable(v bool)`

SetIsDscpCapable sets IsDscpCapable field to given value.

### HasIsDscpCapable

`func (o *Member) HasIsDscpCapable() bool`

HasIsDscpCapable returns a boolean if a field has been set.

### GetLan2Enabled

`func (o *Member) GetLan2Enabled() bool`

GetLan2Enabled returns the Lan2Enabled field if non-nil, zero value otherwise.

### GetLan2EnabledOk

`func (o *Member) GetLan2EnabledOk() (*bool, bool)`

GetLan2EnabledOk returns a tuple with the Lan2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan2Enabled

`func (o *Member) SetLan2Enabled(v bool)`

SetLan2Enabled sets Lan2Enabled field to given value.

### HasLan2Enabled

`func (o *Member) HasLan2Enabled() bool`

HasLan2Enabled returns a boolean if a field has been set.

### GetLan2PortSetting

`func (o *Member) GetLan2PortSetting() MemberLan2PortSetting`

GetLan2PortSetting returns the Lan2PortSetting field if non-nil, zero value otherwise.

### GetLan2PortSettingOk

`func (o *Member) GetLan2PortSettingOk() (*MemberLan2PortSetting, bool)`

GetLan2PortSettingOk returns a tuple with the Lan2PortSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan2PortSetting

`func (o *Member) SetLan2PortSetting(v MemberLan2PortSetting)`

SetLan2PortSetting sets Lan2PortSetting field to given value.

### HasLan2PortSetting

`func (o *Member) HasLan2PortSetting() bool`

HasLan2PortSetting returns a boolean if a field has been set.

### GetLomNetworkConfig

`func (o *Member) GetLomNetworkConfig() []MemberLomNetworkConfig`

GetLomNetworkConfig returns the LomNetworkConfig field if non-nil, zero value otherwise.

### GetLomNetworkConfigOk

`func (o *Member) GetLomNetworkConfigOk() (*[]MemberLomNetworkConfig, bool)`

GetLomNetworkConfigOk returns a tuple with the LomNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomNetworkConfig

`func (o *Member) SetLomNetworkConfig(v []MemberLomNetworkConfig)`

SetLomNetworkConfig sets LomNetworkConfig field to given value.

### HasLomNetworkConfig

`func (o *Member) HasLomNetworkConfig() bool`

HasLomNetworkConfig returns a boolean if a field has been set.

### GetLomUsers

`func (o *Member) GetLomUsers() []MemberLomUsers`

GetLomUsers returns the LomUsers field if non-nil, zero value otherwise.

### GetLomUsersOk

`func (o *Member) GetLomUsersOk() (*[]MemberLomUsers, bool)`

GetLomUsersOk returns a tuple with the LomUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomUsers

`func (o *Member) SetLomUsers(v []MemberLomUsers)`

SetLomUsers sets LomUsers field to given value.

### HasLomUsers

`func (o *Member) HasLomUsers() bool`

HasLomUsers returns a boolean if a field has been set.

### GetMasterCandidate

`func (o *Member) GetMasterCandidate() bool`

GetMasterCandidate returns the MasterCandidate field if non-nil, zero value otherwise.

### GetMasterCandidateOk

`func (o *Member) GetMasterCandidateOk() (*bool, bool)`

GetMasterCandidateOk returns a tuple with the MasterCandidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterCandidate

`func (o *Member) SetMasterCandidate(v bool)`

SetMasterCandidate sets MasterCandidate field to given value.

### HasMasterCandidate

`func (o *Member) HasMasterCandidate() bool`

HasMasterCandidate returns a boolean if a field has been set.

### GetMemberServiceCommunication

`func (o *Member) GetMemberServiceCommunication() []MemberMemberServiceCommunication`

GetMemberServiceCommunication returns the MemberServiceCommunication field if non-nil, zero value otherwise.

### GetMemberServiceCommunicationOk

`func (o *Member) GetMemberServiceCommunicationOk() (*[]MemberMemberServiceCommunication, bool)`

GetMemberServiceCommunicationOk returns a tuple with the MemberServiceCommunication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberServiceCommunication

`func (o *Member) SetMemberServiceCommunication(v []MemberMemberServiceCommunication)`

SetMemberServiceCommunication sets MemberServiceCommunication field to given value.

### HasMemberServiceCommunication

`func (o *Member) HasMemberServiceCommunication() bool`

HasMemberServiceCommunication returns a boolean if a field has been set.

### GetMgmtPortSetting

`func (o *Member) GetMgmtPortSetting() MemberMgmtPortSetting`

GetMgmtPortSetting returns the MgmtPortSetting field if non-nil, zero value otherwise.

### GetMgmtPortSettingOk

`func (o *Member) GetMgmtPortSettingOk() (*MemberMgmtPortSetting, bool)`

GetMgmtPortSettingOk returns a tuple with the MgmtPortSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPortSetting

`func (o *Member) SetMgmtPortSetting(v MemberMgmtPortSetting)`

SetMgmtPortSetting sets MgmtPortSetting field to given value.

### HasMgmtPortSetting

`func (o *Member) HasMgmtPortSetting() bool`

HasMgmtPortSetting returns a boolean if a field has been set.

### GetMmdbEaBuildTime

`func (o *Member) GetMmdbEaBuildTime() int64`

GetMmdbEaBuildTime returns the MmdbEaBuildTime field if non-nil, zero value otherwise.

### GetMmdbEaBuildTimeOk

`func (o *Member) GetMmdbEaBuildTimeOk() (*int64, bool)`

GetMmdbEaBuildTimeOk returns a tuple with the MmdbEaBuildTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmdbEaBuildTime

`func (o *Member) SetMmdbEaBuildTime(v int64)`

SetMmdbEaBuildTime sets MmdbEaBuildTime field to given value.

### HasMmdbEaBuildTime

`func (o *Member) HasMmdbEaBuildTime() bool`

HasMmdbEaBuildTime returns a boolean if a field has been set.

### GetMmdbGeoipBuildTime

`func (o *Member) GetMmdbGeoipBuildTime() int64`

GetMmdbGeoipBuildTime returns the MmdbGeoipBuildTime field if non-nil, zero value otherwise.

### GetMmdbGeoipBuildTimeOk

`func (o *Member) GetMmdbGeoipBuildTimeOk() (*int64, bool)`

GetMmdbGeoipBuildTimeOk returns a tuple with the MmdbGeoipBuildTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmdbGeoipBuildTime

`func (o *Member) SetMmdbGeoipBuildTime(v int64)`

SetMmdbGeoipBuildTime sets MmdbGeoipBuildTime field to given value.

### HasMmdbGeoipBuildTime

`func (o *Member) HasMmdbGeoipBuildTime() bool`

HasMmdbGeoipBuildTime returns a boolean if a field has been set.

### GetNatSetting

`func (o *Member) GetNatSetting() MemberNatSetting`

GetNatSetting returns the NatSetting field if non-nil, zero value otherwise.

### GetNatSettingOk

`func (o *Member) GetNatSettingOk() (*MemberNatSetting, bool)`

GetNatSettingOk returns a tuple with the NatSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatSetting

`func (o *Member) SetNatSetting(v MemberNatSetting)`

SetNatSetting sets NatSetting field to given value.

### HasNatSetting

`func (o *Member) HasNatSetting() bool`

HasNatSetting returns a boolean if a field has been set.

### GetNodeInfo

`func (o *Member) GetNodeInfo() []MemberNodeInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *Member) GetNodeInfoOk() (*[]MemberNodeInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *Member) SetNodeInfo(v []MemberNodeInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *Member) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetNtpSetting

`func (o *Member) GetNtpSetting() MemberNtpSetting`

GetNtpSetting returns the NtpSetting field if non-nil, zero value otherwise.

### GetNtpSettingOk

`func (o *Member) GetNtpSettingOk() (*MemberNtpSetting, bool)`

GetNtpSettingOk returns a tuple with the NtpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpSetting

`func (o *Member) SetNtpSetting(v MemberNtpSetting)`

SetNtpSetting sets NtpSetting field to given value.

### HasNtpSetting

`func (o *Member) HasNtpSetting() bool`

HasNtpSetting returns a boolean if a field has been set.

### GetOspfList

`func (o *Member) GetOspfList() []MemberOspfList`

GetOspfList returns the OspfList field if non-nil, zero value otherwise.

### GetOspfListOk

`func (o *Member) GetOspfListOk() (*[]MemberOspfList, bool)`

GetOspfListOk returns a tuple with the OspfList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfList

`func (o *Member) SetOspfList(v []MemberOspfList)`

SetOspfList sets OspfList field to given value.

### HasOspfList

`func (o *Member) HasOspfList() bool`

HasOspfList returns a boolean if a field has been set.

### GetPassiveHaArpEnabled

`func (o *Member) GetPassiveHaArpEnabled() bool`

GetPassiveHaArpEnabled returns the PassiveHaArpEnabled field if non-nil, zero value otherwise.

### GetPassiveHaArpEnabledOk

`func (o *Member) GetPassiveHaArpEnabledOk() (*bool, bool)`

GetPassiveHaArpEnabledOk returns a tuple with the PassiveHaArpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveHaArpEnabled

`func (o *Member) SetPassiveHaArpEnabled(v bool)`

SetPassiveHaArpEnabled sets PassiveHaArpEnabled field to given value.

### HasPassiveHaArpEnabled

`func (o *Member) HasPassiveHaArpEnabled() bool`

HasPassiveHaArpEnabled returns a boolean if a field has been set.

### GetPlatform

`func (o *Member) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Member) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Member) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Member) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPreProvisioning

`func (o *Member) GetPreProvisioning() MemberPreProvisioning`

GetPreProvisioning returns the PreProvisioning field if non-nil, zero value otherwise.

### GetPreProvisioningOk

`func (o *Member) GetPreProvisioningOk() (*MemberPreProvisioning, bool)`

GetPreProvisioningOk returns a tuple with the PreProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProvisioning

`func (o *Member) SetPreProvisioning(v MemberPreProvisioning)`

SetPreProvisioning sets PreProvisioning field to given value.

### HasPreProvisioning

`func (o *Member) HasPreProvisioning() bool`

HasPreProvisioning returns a boolean if a field has been set.

### GetPreserveIfOwnsDelegation

`func (o *Member) GetPreserveIfOwnsDelegation() bool`

GetPreserveIfOwnsDelegation returns the PreserveIfOwnsDelegation field if non-nil, zero value otherwise.

### GetPreserveIfOwnsDelegationOk

`func (o *Member) GetPreserveIfOwnsDelegationOk() (*bool, bool)`

GetPreserveIfOwnsDelegationOk returns a tuple with the PreserveIfOwnsDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveIfOwnsDelegation

`func (o *Member) SetPreserveIfOwnsDelegation(v bool)`

SetPreserveIfOwnsDelegation sets PreserveIfOwnsDelegation field to given value.

### HasPreserveIfOwnsDelegation

`func (o *Member) HasPreserveIfOwnsDelegation() bool`

HasPreserveIfOwnsDelegation returns a boolean if a field has been set.

### GetRemoteConsoleAccessEnable

`func (o *Member) GetRemoteConsoleAccessEnable() bool`

GetRemoteConsoleAccessEnable returns the RemoteConsoleAccessEnable field if non-nil, zero value otherwise.

### GetRemoteConsoleAccessEnableOk

`func (o *Member) GetRemoteConsoleAccessEnableOk() (*bool, bool)`

GetRemoteConsoleAccessEnableOk returns a tuple with the RemoteConsoleAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConsoleAccessEnable

`func (o *Member) SetRemoteConsoleAccessEnable(v bool)`

SetRemoteConsoleAccessEnable sets RemoteConsoleAccessEnable field to given value.

### HasRemoteConsoleAccessEnable

`func (o *Member) HasRemoteConsoleAccessEnable() bool`

HasRemoteConsoleAccessEnable returns a boolean if a field has been set.

### GetRouterId

`func (o *Member) GetRouterId() int64`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *Member) GetRouterIdOk() (*int64, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *Member) SetRouterId(v int64)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *Member) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetServiceStatus

`func (o *Member) GetServiceStatus() []MemberServiceStatus`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *Member) GetServiceStatusOk() (*[]MemberServiceStatus, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *Member) SetServiceStatus(v []MemberServiceStatus)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *Member) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.

### GetServiceTypeConfiguration

`func (o *Member) GetServiceTypeConfiguration() string`

GetServiceTypeConfiguration returns the ServiceTypeConfiguration field if non-nil, zero value otherwise.

### GetServiceTypeConfigurationOk

`func (o *Member) GetServiceTypeConfigurationOk() (*string, bool)`

GetServiceTypeConfigurationOk returns a tuple with the ServiceTypeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeConfiguration

`func (o *Member) SetServiceTypeConfiguration(v string)`

SetServiceTypeConfiguration sets ServiceTypeConfiguration field to given value.

### HasServiceTypeConfiguration

`func (o *Member) HasServiceTypeConfiguration() bool`

HasServiceTypeConfiguration returns a boolean if a field has been set.

### GetSnmpSetting

`func (o *Member) GetSnmpSetting() MemberSnmpSetting`

GetSnmpSetting returns the SnmpSetting field if non-nil, zero value otherwise.

### GetSnmpSettingOk

`func (o *Member) GetSnmpSettingOk() (*MemberSnmpSetting, bool)`

GetSnmpSettingOk returns a tuple with the SnmpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpSetting

`func (o *Member) SetSnmpSetting(v MemberSnmpSetting)`

SetSnmpSetting sets SnmpSetting field to given value.

### HasSnmpSetting

`func (o *Member) HasSnmpSetting() bool`

HasSnmpSetting returns a boolean if a field has been set.

### GetStaticRoutes

`func (o *Member) GetStaticRoutes() []MemberStaticRoutes`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *Member) GetStaticRoutesOk() (*[]MemberStaticRoutes, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *Member) SetStaticRoutes(v []MemberStaticRoutes)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *Member) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.

### GetSupportAccessEnable

`func (o *Member) GetSupportAccessEnable() bool`

GetSupportAccessEnable returns the SupportAccessEnable field if non-nil, zero value otherwise.

### GetSupportAccessEnableOk

`func (o *Member) GetSupportAccessEnableOk() (*bool, bool)`

GetSupportAccessEnableOk returns a tuple with the SupportAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAccessEnable

`func (o *Member) SetSupportAccessEnable(v bool)`

SetSupportAccessEnable sets SupportAccessEnable field to given value.

### HasSupportAccessEnable

`func (o *Member) HasSupportAccessEnable() bool`

HasSupportAccessEnable returns a boolean if a field has been set.

### GetSupportAccessInfo

`func (o *Member) GetSupportAccessInfo() string`

GetSupportAccessInfo returns the SupportAccessInfo field if non-nil, zero value otherwise.

### GetSupportAccessInfoOk

`func (o *Member) GetSupportAccessInfoOk() (*string, bool)`

GetSupportAccessInfoOk returns a tuple with the SupportAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAccessInfo

`func (o *Member) SetSupportAccessInfo(v string)`

SetSupportAccessInfo sets SupportAccessInfo field to given value.

### HasSupportAccessInfo

`func (o *Member) HasSupportAccessInfo() bool`

HasSupportAccessInfo returns a boolean if a field has been set.

### GetSyslogProxySetting

`func (o *Member) GetSyslogProxySetting() MemberSyslogProxySetting`

GetSyslogProxySetting returns the SyslogProxySetting field if non-nil, zero value otherwise.

### GetSyslogProxySettingOk

`func (o *Member) GetSyslogProxySettingOk() (*MemberSyslogProxySetting, bool)`

GetSyslogProxySettingOk returns a tuple with the SyslogProxySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogProxySetting

`func (o *Member) SetSyslogProxySetting(v MemberSyslogProxySetting)`

SetSyslogProxySetting sets SyslogProxySetting field to given value.

### HasSyslogProxySetting

`func (o *Member) HasSyslogProxySetting() bool`

HasSyslogProxySetting returns a boolean if a field has been set.

### GetSyslogServers

`func (o *Member) GetSyslogServers() []MemberSyslogServers`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *Member) GetSyslogServersOk() (*[]MemberSyslogServers, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *Member) SetSyslogServers(v []MemberSyslogServers)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *Member) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### GetSyslogSize

`func (o *Member) GetSyslogSize() int64`

GetSyslogSize returns the SyslogSize field if non-nil, zero value otherwise.

### GetSyslogSizeOk

`func (o *Member) GetSyslogSizeOk() (*int64, bool)`

GetSyslogSizeOk returns a tuple with the SyslogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSize

`func (o *Member) SetSyslogSize(v int64)`

SetSyslogSize sets SyslogSize field to given value.

### HasSyslogSize

`func (o *Member) HasSyslogSize() bool`

HasSyslogSize returns a boolean if a field has been set.

### GetThresholdTraps

`func (o *Member) GetThresholdTraps() []MemberThresholdTraps`

GetThresholdTraps returns the ThresholdTraps field if non-nil, zero value otherwise.

### GetThresholdTrapsOk

`func (o *Member) GetThresholdTrapsOk() (*[]MemberThresholdTraps, bool)`

GetThresholdTrapsOk returns a tuple with the ThresholdTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdTraps

`func (o *Member) SetThresholdTraps(v []MemberThresholdTraps)`

SetThresholdTraps sets ThresholdTraps field to given value.

### HasThresholdTraps

`func (o *Member) HasThresholdTraps() bool`

HasThresholdTraps returns a boolean if a field has been set.

### GetTimeZone

`func (o *Member) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Member) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Member) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Member) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTrafficCaptureAuthDnsSetting

`func (o *Member) GetTrafficCaptureAuthDnsSetting() MemberTrafficCaptureAuthDnsSetting`

GetTrafficCaptureAuthDnsSetting returns the TrafficCaptureAuthDnsSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureAuthDnsSettingOk

`func (o *Member) GetTrafficCaptureAuthDnsSettingOk() (*MemberTrafficCaptureAuthDnsSetting, bool)`

GetTrafficCaptureAuthDnsSettingOk returns a tuple with the TrafficCaptureAuthDnsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureAuthDnsSetting

`func (o *Member) SetTrafficCaptureAuthDnsSetting(v MemberTrafficCaptureAuthDnsSetting)`

SetTrafficCaptureAuthDnsSetting sets TrafficCaptureAuthDnsSetting field to given value.

### HasTrafficCaptureAuthDnsSetting

`func (o *Member) HasTrafficCaptureAuthDnsSetting() bool`

HasTrafficCaptureAuthDnsSetting returns a boolean if a field has been set.

### GetTrafficCaptureChrSetting

`func (o *Member) GetTrafficCaptureChrSetting() MemberTrafficCaptureChrSetting`

GetTrafficCaptureChrSetting returns the TrafficCaptureChrSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureChrSettingOk

`func (o *Member) GetTrafficCaptureChrSettingOk() (*MemberTrafficCaptureChrSetting, bool)`

GetTrafficCaptureChrSettingOk returns a tuple with the TrafficCaptureChrSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureChrSetting

`func (o *Member) SetTrafficCaptureChrSetting(v MemberTrafficCaptureChrSetting)`

SetTrafficCaptureChrSetting sets TrafficCaptureChrSetting field to given value.

### HasTrafficCaptureChrSetting

`func (o *Member) HasTrafficCaptureChrSetting() bool`

HasTrafficCaptureChrSetting returns a boolean if a field has been set.

### GetTrafficCaptureQpsSetting

`func (o *Member) GetTrafficCaptureQpsSetting() MemberTrafficCaptureQpsSetting`

GetTrafficCaptureQpsSetting returns the TrafficCaptureQpsSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureQpsSettingOk

`func (o *Member) GetTrafficCaptureQpsSettingOk() (*MemberTrafficCaptureQpsSetting, bool)`

GetTrafficCaptureQpsSettingOk returns a tuple with the TrafficCaptureQpsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureQpsSetting

`func (o *Member) SetTrafficCaptureQpsSetting(v MemberTrafficCaptureQpsSetting)`

SetTrafficCaptureQpsSetting sets TrafficCaptureQpsSetting field to given value.

### HasTrafficCaptureQpsSetting

`func (o *Member) HasTrafficCaptureQpsSetting() bool`

HasTrafficCaptureQpsSetting returns a boolean if a field has been set.

### GetTrafficCaptureRecDnsSetting

`func (o *Member) GetTrafficCaptureRecDnsSetting() MemberTrafficCaptureRecDnsSetting`

GetTrafficCaptureRecDnsSetting returns the TrafficCaptureRecDnsSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureRecDnsSettingOk

`func (o *Member) GetTrafficCaptureRecDnsSettingOk() (*MemberTrafficCaptureRecDnsSetting, bool)`

GetTrafficCaptureRecDnsSettingOk returns a tuple with the TrafficCaptureRecDnsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureRecDnsSetting

`func (o *Member) SetTrafficCaptureRecDnsSetting(v MemberTrafficCaptureRecDnsSetting)`

SetTrafficCaptureRecDnsSetting sets TrafficCaptureRecDnsSetting field to given value.

### HasTrafficCaptureRecDnsSetting

`func (o *Member) HasTrafficCaptureRecDnsSetting() bool`

HasTrafficCaptureRecDnsSetting returns a boolean if a field has been set.

### GetTrafficCaptureRecQueriesSetting

`func (o *Member) GetTrafficCaptureRecQueriesSetting() MemberTrafficCaptureRecQueriesSetting`

GetTrafficCaptureRecQueriesSetting returns the TrafficCaptureRecQueriesSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureRecQueriesSettingOk

`func (o *Member) GetTrafficCaptureRecQueriesSettingOk() (*MemberTrafficCaptureRecQueriesSetting, bool)`

GetTrafficCaptureRecQueriesSettingOk returns a tuple with the TrafficCaptureRecQueriesSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureRecQueriesSetting

`func (o *Member) SetTrafficCaptureRecQueriesSetting(v MemberTrafficCaptureRecQueriesSetting)`

SetTrafficCaptureRecQueriesSetting sets TrafficCaptureRecQueriesSetting field to given value.

### HasTrafficCaptureRecQueriesSetting

`func (o *Member) HasTrafficCaptureRecQueriesSetting() bool`

HasTrafficCaptureRecQueriesSetting returns a boolean if a field has been set.

### GetTrapNotifications

`func (o *Member) GetTrapNotifications() []MemberTrapNotifications`

GetTrapNotifications returns the TrapNotifications field if non-nil, zero value otherwise.

### GetTrapNotificationsOk

`func (o *Member) GetTrapNotificationsOk() (*[]MemberTrapNotifications, bool)`

GetTrapNotificationsOk returns a tuple with the TrapNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapNotifications

`func (o *Member) SetTrapNotifications(v []MemberTrapNotifications)`

SetTrapNotifications sets TrapNotifications field to given value.

### HasTrapNotifications

`func (o *Member) HasTrapNotifications() bool`

HasTrapNotifications returns a boolean if a field has been set.

### GetUpgradeGroup

`func (o *Member) GetUpgradeGroup() string`

GetUpgradeGroup returns the UpgradeGroup field if non-nil, zero value otherwise.

### GetUpgradeGroupOk

`func (o *Member) GetUpgradeGroupOk() (*string, bool)`

GetUpgradeGroupOk returns a tuple with the UpgradeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeGroup

`func (o *Member) SetUpgradeGroup(v string)`

SetUpgradeGroup sets UpgradeGroup field to given value.

### HasUpgradeGroup

`func (o *Member) HasUpgradeGroup() bool`

HasUpgradeGroup returns a boolean if a field has been set.

### GetUseAutomatedTrafficCapture

`func (o *Member) GetUseAutomatedTrafficCapture() bool`

GetUseAutomatedTrafficCapture returns the UseAutomatedTrafficCapture field if non-nil, zero value otherwise.

### GetUseAutomatedTrafficCaptureOk

`func (o *Member) GetUseAutomatedTrafficCaptureOk() (*bool, bool)`

GetUseAutomatedTrafficCaptureOk returns a tuple with the UseAutomatedTrafficCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAutomatedTrafficCapture

`func (o *Member) SetUseAutomatedTrafficCapture(v bool)`

SetUseAutomatedTrafficCapture sets UseAutomatedTrafficCapture field to given value.

### HasUseAutomatedTrafficCapture

`func (o *Member) HasUseAutomatedTrafficCapture() bool`

HasUseAutomatedTrafficCapture returns a boolean if a field has been set.

### GetUseDnsResolverSetting

`func (o *Member) GetUseDnsResolverSetting() bool`

GetUseDnsResolverSetting returns the UseDnsResolverSetting field if non-nil, zero value otherwise.

### GetUseDnsResolverSettingOk

`func (o *Member) GetUseDnsResolverSettingOk() (*bool, bool)`

GetUseDnsResolverSettingOk returns a tuple with the UseDnsResolverSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnsResolverSetting

`func (o *Member) SetUseDnsResolverSetting(v bool)`

SetUseDnsResolverSetting sets UseDnsResolverSetting field to given value.

### HasUseDnsResolverSetting

`func (o *Member) HasUseDnsResolverSetting() bool`

HasUseDnsResolverSetting returns a boolean if a field has been set.

### GetUseDscp

`func (o *Member) GetUseDscp() bool`

GetUseDscp returns the UseDscp field if non-nil, zero value otherwise.

### GetUseDscpOk

`func (o *Member) GetUseDscpOk() (*bool, bool)`

GetUseDscpOk returns a tuple with the UseDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDscp

`func (o *Member) SetUseDscp(v bool)`

SetUseDscp sets UseDscp field to given value.

### HasUseDscp

`func (o *Member) HasUseDscp() bool`

HasUseDscp returns a boolean if a field has been set.

### GetUseEmailSetting

`func (o *Member) GetUseEmailSetting() bool`

GetUseEmailSetting returns the UseEmailSetting field if non-nil, zero value otherwise.

### GetUseEmailSettingOk

`func (o *Member) GetUseEmailSettingOk() (*bool, bool)`

GetUseEmailSettingOk returns a tuple with the UseEmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailSetting

`func (o *Member) SetUseEmailSetting(v bool)`

SetUseEmailSetting sets UseEmailSetting field to given value.

### HasUseEmailSetting

`func (o *Member) HasUseEmailSetting() bool`

HasUseEmailSetting returns a boolean if a field has been set.

### GetUseEnableLom

`func (o *Member) GetUseEnableLom() bool`

GetUseEnableLom returns the UseEnableLom field if non-nil, zero value otherwise.

### GetUseEnableLomOk

`func (o *Member) GetUseEnableLomOk() (*bool, bool)`

GetUseEnableLomOk returns a tuple with the UseEnableLom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableLom

`func (o *Member) SetUseEnableLom(v bool)`

SetUseEnableLom sets UseEnableLom field to given value.

### HasUseEnableLom

`func (o *Member) HasUseEnableLom() bool`

HasUseEnableLom returns a boolean if a field has been set.

### GetUseEnableMemberRedirect

`func (o *Member) GetUseEnableMemberRedirect() bool`

GetUseEnableMemberRedirect returns the UseEnableMemberRedirect field if non-nil, zero value otherwise.

### GetUseEnableMemberRedirectOk

`func (o *Member) GetUseEnableMemberRedirectOk() (*bool, bool)`

GetUseEnableMemberRedirectOk returns a tuple with the UseEnableMemberRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableMemberRedirect

`func (o *Member) SetUseEnableMemberRedirect(v bool)`

SetUseEnableMemberRedirect sets UseEnableMemberRedirect field to given value.

### HasUseEnableMemberRedirect

`func (o *Member) HasUseEnableMemberRedirect() bool`

HasUseEnableMemberRedirect returns a boolean if a field has been set.

### GetUseExternalSyslogBackupServers

`func (o *Member) GetUseExternalSyslogBackupServers() bool`

GetUseExternalSyslogBackupServers returns the UseExternalSyslogBackupServers field if non-nil, zero value otherwise.

### GetUseExternalSyslogBackupServersOk

`func (o *Member) GetUseExternalSyslogBackupServersOk() (*bool, bool)`

GetUseExternalSyslogBackupServersOk returns a tuple with the UseExternalSyslogBackupServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExternalSyslogBackupServers

`func (o *Member) SetUseExternalSyslogBackupServers(v bool)`

SetUseExternalSyslogBackupServers sets UseExternalSyslogBackupServers field to given value.

### HasUseExternalSyslogBackupServers

`func (o *Member) HasUseExternalSyslogBackupServers() bool`

HasUseExternalSyslogBackupServers returns a boolean if a field has been set.

### GetUseRemoteConsoleAccessEnable

`func (o *Member) GetUseRemoteConsoleAccessEnable() bool`

GetUseRemoteConsoleAccessEnable returns the UseRemoteConsoleAccessEnable field if non-nil, zero value otherwise.

### GetUseRemoteConsoleAccessEnableOk

`func (o *Member) GetUseRemoteConsoleAccessEnableOk() (*bool, bool)`

GetUseRemoteConsoleAccessEnableOk returns a tuple with the UseRemoteConsoleAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRemoteConsoleAccessEnable

`func (o *Member) SetUseRemoteConsoleAccessEnable(v bool)`

SetUseRemoteConsoleAccessEnable sets UseRemoteConsoleAccessEnable field to given value.

### HasUseRemoteConsoleAccessEnable

`func (o *Member) HasUseRemoteConsoleAccessEnable() bool`

HasUseRemoteConsoleAccessEnable returns a boolean if a field has been set.

### GetUseSnmpSetting

`func (o *Member) GetUseSnmpSetting() bool`

GetUseSnmpSetting returns the UseSnmpSetting field if non-nil, zero value otherwise.

### GetUseSnmpSettingOk

`func (o *Member) GetUseSnmpSettingOk() (*bool, bool)`

GetUseSnmpSettingOk returns a tuple with the UseSnmpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpSetting

`func (o *Member) SetUseSnmpSetting(v bool)`

SetUseSnmpSetting sets UseSnmpSetting field to given value.

### HasUseSnmpSetting

`func (o *Member) HasUseSnmpSetting() bool`

HasUseSnmpSetting returns a boolean if a field has been set.

### GetUseSupportAccessEnable

`func (o *Member) GetUseSupportAccessEnable() bool`

GetUseSupportAccessEnable returns the UseSupportAccessEnable field if non-nil, zero value otherwise.

### GetUseSupportAccessEnableOk

`func (o *Member) GetUseSupportAccessEnableOk() (*bool, bool)`

GetUseSupportAccessEnableOk returns a tuple with the UseSupportAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSupportAccessEnable

`func (o *Member) SetUseSupportAccessEnable(v bool)`

SetUseSupportAccessEnable sets UseSupportAccessEnable field to given value.

### HasUseSupportAccessEnable

`func (o *Member) HasUseSupportAccessEnable() bool`

HasUseSupportAccessEnable returns a boolean if a field has been set.

### GetUseSyslogProxySetting

`func (o *Member) GetUseSyslogProxySetting() bool`

GetUseSyslogProxySetting returns the UseSyslogProxySetting field if non-nil, zero value otherwise.

### GetUseSyslogProxySettingOk

`func (o *Member) GetUseSyslogProxySettingOk() (*bool, bool)`

GetUseSyslogProxySettingOk returns a tuple with the UseSyslogProxySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSyslogProxySetting

`func (o *Member) SetUseSyslogProxySetting(v bool)`

SetUseSyslogProxySetting sets UseSyslogProxySetting field to given value.

### HasUseSyslogProxySetting

`func (o *Member) HasUseSyslogProxySetting() bool`

HasUseSyslogProxySetting returns a boolean if a field has been set.

### GetUseThresholdTraps

`func (o *Member) GetUseThresholdTraps() bool`

GetUseThresholdTraps returns the UseThresholdTraps field if non-nil, zero value otherwise.

### GetUseThresholdTrapsOk

`func (o *Member) GetUseThresholdTrapsOk() (*bool, bool)`

GetUseThresholdTrapsOk returns a tuple with the UseThresholdTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseThresholdTraps

`func (o *Member) SetUseThresholdTraps(v bool)`

SetUseThresholdTraps sets UseThresholdTraps field to given value.

### HasUseThresholdTraps

`func (o *Member) HasUseThresholdTraps() bool`

HasUseThresholdTraps returns a boolean if a field has been set.

### GetUseTimeZone

`func (o *Member) GetUseTimeZone() bool`

GetUseTimeZone returns the UseTimeZone field if non-nil, zero value otherwise.

### GetUseTimeZoneOk

`func (o *Member) GetUseTimeZoneOk() (*bool, bool)`

GetUseTimeZoneOk returns a tuple with the UseTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeZone

`func (o *Member) SetUseTimeZone(v bool)`

SetUseTimeZone sets UseTimeZone field to given value.

### HasUseTimeZone

`func (o *Member) HasUseTimeZone() bool`

HasUseTimeZone returns a boolean if a field has been set.

### GetUseTrafficCaptureAuthDns

`func (o *Member) GetUseTrafficCaptureAuthDns() bool`

GetUseTrafficCaptureAuthDns returns the UseTrafficCaptureAuthDns field if non-nil, zero value otherwise.

### GetUseTrafficCaptureAuthDnsOk

`func (o *Member) GetUseTrafficCaptureAuthDnsOk() (*bool, bool)`

GetUseTrafficCaptureAuthDnsOk returns a tuple with the UseTrafficCaptureAuthDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureAuthDns

`func (o *Member) SetUseTrafficCaptureAuthDns(v bool)`

SetUseTrafficCaptureAuthDns sets UseTrafficCaptureAuthDns field to given value.

### HasUseTrafficCaptureAuthDns

`func (o *Member) HasUseTrafficCaptureAuthDns() bool`

HasUseTrafficCaptureAuthDns returns a boolean if a field has been set.

### GetUseTrafficCaptureChr

`func (o *Member) GetUseTrafficCaptureChr() bool`

GetUseTrafficCaptureChr returns the UseTrafficCaptureChr field if non-nil, zero value otherwise.

### GetUseTrafficCaptureChrOk

`func (o *Member) GetUseTrafficCaptureChrOk() (*bool, bool)`

GetUseTrafficCaptureChrOk returns a tuple with the UseTrafficCaptureChr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureChr

`func (o *Member) SetUseTrafficCaptureChr(v bool)`

SetUseTrafficCaptureChr sets UseTrafficCaptureChr field to given value.

### HasUseTrafficCaptureChr

`func (o *Member) HasUseTrafficCaptureChr() bool`

HasUseTrafficCaptureChr returns a boolean if a field has been set.

### GetUseTrafficCaptureQps

`func (o *Member) GetUseTrafficCaptureQps() bool`

GetUseTrafficCaptureQps returns the UseTrafficCaptureQps field if non-nil, zero value otherwise.

### GetUseTrafficCaptureQpsOk

`func (o *Member) GetUseTrafficCaptureQpsOk() (*bool, bool)`

GetUseTrafficCaptureQpsOk returns a tuple with the UseTrafficCaptureQps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureQps

`func (o *Member) SetUseTrafficCaptureQps(v bool)`

SetUseTrafficCaptureQps sets UseTrafficCaptureQps field to given value.

### HasUseTrafficCaptureQps

`func (o *Member) HasUseTrafficCaptureQps() bool`

HasUseTrafficCaptureQps returns a boolean if a field has been set.

### GetUseTrafficCaptureRecDns

`func (o *Member) GetUseTrafficCaptureRecDns() bool`

GetUseTrafficCaptureRecDns returns the UseTrafficCaptureRecDns field if non-nil, zero value otherwise.

### GetUseTrafficCaptureRecDnsOk

`func (o *Member) GetUseTrafficCaptureRecDnsOk() (*bool, bool)`

GetUseTrafficCaptureRecDnsOk returns a tuple with the UseTrafficCaptureRecDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureRecDns

`func (o *Member) SetUseTrafficCaptureRecDns(v bool)`

SetUseTrafficCaptureRecDns sets UseTrafficCaptureRecDns field to given value.

### HasUseTrafficCaptureRecDns

`func (o *Member) HasUseTrafficCaptureRecDns() bool`

HasUseTrafficCaptureRecDns returns a boolean if a field has been set.

### GetUseTrafficCaptureRecQueries

`func (o *Member) GetUseTrafficCaptureRecQueries() bool`

GetUseTrafficCaptureRecQueries returns the UseTrafficCaptureRecQueries field if non-nil, zero value otherwise.

### GetUseTrafficCaptureRecQueriesOk

`func (o *Member) GetUseTrafficCaptureRecQueriesOk() (*bool, bool)`

GetUseTrafficCaptureRecQueriesOk returns a tuple with the UseTrafficCaptureRecQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureRecQueries

`func (o *Member) SetUseTrafficCaptureRecQueries(v bool)`

SetUseTrafficCaptureRecQueries sets UseTrafficCaptureRecQueries field to given value.

### HasUseTrafficCaptureRecQueries

`func (o *Member) HasUseTrafficCaptureRecQueries() bool`

HasUseTrafficCaptureRecQueries returns a boolean if a field has been set.

### GetUseTrapNotifications

`func (o *Member) GetUseTrapNotifications() bool`

GetUseTrapNotifications returns the UseTrapNotifications field if non-nil, zero value otherwise.

### GetUseTrapNotificationsOk

`func (o *Member) GetUseTrapNotificationsOk() (*bool, bool)`

GetUseTrapNotificationsOk returns a tuple with the UseTrapNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrapNotifications

`func (o *Member) SetUseTrapNotifications(v bool)`

SetUseTrapNotifications sets UseTrapNotifications field to given value.

### HasUseTrapNotifications

`func (o *Member) HasUseTrapNotifications() bool`

HasUseTrapNotifications returns a boolean if a field has been set.

### GetUseV4Vrrp

`func (o *Member) GetUseV4Vrrp() bool`

GetUseV4Vrrp returns the UseV4Vrrp field if non-nil, zero value otherwise.

### GetUseV4VrrpOk

`func (o *Member) GetUseV4VrrpOk() (*bool, bool)`

GetUseV4VrrpOk returns a tuple with the UseV4Vrrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseV4Vrrp

`func (o *Member) SetUseV4Vrrp(v bool)`

SetUseV4Vrrp sets UseV4Vrrp field to given value.

### HasUseV4Vrrp

`func (o *Member) HasUseV4Vrrp() bool`

HasUseV4Vrrp returns a boolean if a field has been set.

### GetVipSetting

`func (o *Member) GetVipSetting() MemberVipSetting`

GetVipSetting returns the VipSetting field if non-nil, zero value otherwise.

### GetVipSettingOk

`func (o *Member) GetVipSettingOk() (*MemberVipSetting, bool)`

GetVipSettingOk returns a tuple with the VipSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSetting

`func (o *Member) SetVipSetting(v MemberVipSetting)`

SetVipSetting sets VipSetting field to given value.

### HasVipSetting

`func (o *Member) HasVipSetting() bool`

HasVipSetting returns a boolean if a field has been set.

### GetVpnMtu

`func (o *Member) GetVpnMtu() int64`

GetVpnMtu returns the VpnMtu field if non-nil, zero value otherwise.

### GetVpnMtuOk

`func (o *Member) GetVpnMtuOk() (*int64, bool)`

GetVpnMtuOk returns a tuple with the VpnMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnMtu

`func (o *Member) SetVpnMtu(v int64)`

SetVpnMtu sets VpnMtu field to given value.

### HasVpnMtu

`func (o *Member) HasVpnMtu() bool`

HasVpnMtu returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


