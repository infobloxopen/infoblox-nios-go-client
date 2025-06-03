# GetMemberResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ActivePosition** | Pointer to **string** | The active server of a Grid member. | [optional] [readonly] 
**AdditionalIpList** | Pointer to [**[]MemberAdditionalIpList**](MemberAdditionalIpList.md) | The additional IP list of a Grid member. This list contains additional interface information that can be used at the member level. Note that interface structure(s) with interface type set to &#39;MGMT&#39; are not supported. | [optional] 
**AutomatedTrafficCaptureSetting** | Pointer to [**MemberAutomatedTrafficCaptureSetting**](MemberAutomatedTrafficCaptureSetting.md) |  | [optional] 
**BgpAs** | Pointer to [**[]MemberBgpAs**](MemberBgpAs.md) | The BGP configuration for anycast for a Grid member. | [optional] 
**CaptureTrafficControl** | Pointer to **map[string]interface{}** |  | [optional] 
**CaptureTrafficStatus** | Pointer to **map[string]interface{}** |  | [optional] 
**Comment** | Pointer to **string** | A descriptive comment of the Grid member. | [optional] 
**ConfigAddrType** | Pointer to **string** | Address configuration type. | [optional] 
**CreateToken** | Pointer to **map[string]interface{}** |  | [optional] 
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
**HostName** | Pointer to **string** | The host name of the Grid member. | [optional] 
**Ipv6Setting** | Pointer to [**MemberIpv6Setting**](MemberIpv6Setting.md) |  | [optional] 
**Ipv6StaticRoutes** | Pointer to [**[]MemberIpv6StaticRoutes**](MemberIpv6StaticRoutes.md) | List of IPv6 static routes. | [optional] 
**IsDscpCapable** | Pointer to **bool** | Determines if a Grid member supports DSCP (Differentiated Services Code Point). | [optional] [readonly] 
**Lan2Enabled** | Pointer to **bool** | If this is set to \&quot;true\&quot;, the LAN2 port is enabled as an independent port or as a port for failover purposes. | [optional] 
**Lan2PortSetting** | Pointer to [**MemberLan2PortSetting**](MemberLan2PortSetting.md) |  | [optional] 
**LomNetworkConfig** | Pointer to [**[]MemberLomNetworkConfig**](MemberLomNetworkConfig.md) | The Network configurations for LOM. | [optional] 
**LomUsers** | Pointer to [**[]MemberLomUsers**](MemberLomUsers.md) | The list of LOM users. | [optional] 
**MasterCandidate** | Pointer to **bool** | Determines if a Grid member is a Grid Master Candidate or not. This flag enables the Grid member to assume the role of the Grid Master as a disaster recovery measure. | [optional] 
**MemberAdminOperation** | Pointer to **map[string]interface{}** |  | [optional] 
**MemberServiceCommunication** | Pointer to [**[]MemberMemberServiceCommunication**](MemberMemberServiceCommunication.md) | Configure communication type for various services. | [optional] 
**MgmtPortSetting** | Pointer to [**MemberMgmtPortSetting**](MemberMgmtPortSetting.md) |  | [optional] 
**MmdbEaBuildTime** | Pointer to **int64** | Extensible attributes Topology database build time. | [optional] [readonly] 
**MmdbGeoipBuildTime** | Pointer to **int64** | GeoIP Topology database build time. | [optional] [readonly] 
**NatSetting** | Pointer to [**MemberNatSetting**](MemberNatSetting.md) |  | [optional] 
**NodeInfo** | Pointer to [**[]MemberNodeInfo**](MemberNodeInfo.md) | The node information list with detailed status report on the operations of the Grid Member. | [optional] 
**NtpSetting** | Pointer to [**MemberNtpSetting**](MemberNtpSetting.md) |  | [optional] 
**OspfList** | Pointer to [**[]MemberOspfList**](MemberOspfList.md) | The OSPF area configuration (for anycast) list for a Grid member. | [optional] 
**PassiveHaArpEnabled** | Pointer to **bool** | The ARP protocol setting on the passive node of an HA pair. If you do not specify a value, the default value is \&quot;false\&quot;. You can only set this value to \&quot;true\&quot; if the member is an HA pair. | [optional] 
**Platform** | Pointer to **string** | Hardware Platform. | [optional] 
**PreProvisioning** | Pointer to [**MemberPreProvisioning**](MemberPreProvisioning.md) |  | [optional] 
**PreserveIfOwnsDelegation** | Pointer to **bool** | Set this flag to \&quot;true\&quot; to prevent the deletion of the member if any delegated object remains attached to it. | [optional] 
**ReadToken** | Pointer to **map[string]interface{}** |  | [optional] 
**RemoteConsoleAccessEnable** | Pointer to **bool** | If set to True, superuser admins can access the Infoblox CLI from a remote location using an SSH (Secure Shell) v2 client. | [optional] 
**Requestrestartservicestatus** | Pointer to **map[string]interface{}** |  | [optional] 
**Restartservices** | Pointer to **map[string]interface{}** |  | [optional] 
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
**Result** | Pointer to [**Member**](Member.md) |  | [optional] 

## Methods

### NewGetMemberResponse

`func NewGetMemberResponse() *GetMemberResponse`

NewGetMemberResponse instantiates a new GetMemberResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMemberResponseWithDefaults

`func NewGetMemberResponseWithDefaults() *GetMemberResponse`

NewGetMemberResponseWithDefaults instantiates a new GetMemberResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMemberResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMemberResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMemberResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMemberResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActivePosition

`func (o *GetMemberResponse) GetActivePosition() string`

GetActivePosition returns the ActivePosition field if non-nil, zero value otherwise.

### GetActivePositionOk

`func (o *GetMemberResponse) GetActivePositionOk() (*string, bool)`

GetActivePositionOk returns a tuple with the ActivePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePosition

`func (o *GetMemberResponse) SetActivePosition(v string)`

SetActivePosition sets ActivePosition field to given value.

### HasActivePosition

`func (o *GetMemberResponse) HasActivePosition() bool`

HasActivePosition returns a boolean if a field has been set.

### GetAdditionalIpList

`func (o *GetMemberResponse) GetAdditionalIpList() []MemberAdditionalIpList`

GetAdditionalIpList returns the AdditionalIpList field if non-nil, zero value otherwise.

### GetAdditionalIpListOk

`func (o *GetMemberResponse) GetAdditionalIpListOk() (*[]MemberAdditionalIpList, bool)`

GetAdditionalIpListOk returns a tuple with the AdditionalIpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalIpList

`func (o *GetMemberResponse) SetAdditionalIpList(v []MemberAdditionalIpList)`

SetAdditionalIpList sets AdditionalIpList field to given value.

### HasAdditionalIpList

`func (o *GetMemberResponse) HasAdditionalIpList() bool`

HasAdditionalIpList returns a boolean if a field has been set.

### GetAutomatedTrafficCaptureSetting

`func (o *GetMemberResponse) GetAutomatedTrafficCaptureSetting() MemberAutomatedTrafficCaptureSetting`

GetAutomatedTrafficCaptureSetting returns the AutomatedTrafficCaptureSetting field if non-nil, zero value otherwise.

### GetAutomatedTrafficCaptureSettingOk

`func (o *GetMemberResponse) GetAutomatedTrafficCaptureSettingOk() (*MemberAutomatedTrafficCaptureSetting, bool)`

GetAutomatedTrafficCaptureSettingOk returns a tuple with the AutomatedTrafficCaptureSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedTrafficCaptureSetting

`func (o *GetMemberResponse) SetAutomatedTrafficCaptureSetting(v MemberAutomatedTrafficCaptureSetting)`

SetAutomatedTrafficCaptureSetting sets AutomatedTrafficCaptureSetting field to given value.

### HasAutomatedTrafficCaptureSetting

`func (o *GetMemberResponse) HasAutomatedTrafficCaptureSetting() bool`

HasAutomatedTrafficCaptureSetting returns a boolean if a field has been set.

### GetBgpAs

`func (o *GetMemberResponse) GetBgpAs() []MemberBgpAs`

GetBgpAs returns the BgpAs field if non-nil, zero value otherwise.

### GetBgpAsOk

`func (o *GetMemberResponse) GetBgpAsOk() (*[]MemberBgpAs, bool)`

GetBgpAsOk returns a tuple with the BgpAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAs

`func (o *GetMemberResponse) SetBgpAs(v []MemberBgpAs)`

SetBgpAs sets BgpAs field to given value.

### HasBgpAs

`func (o *GetMemberResponse) HasBgpAs() bool`

HasBgpAs returns a boolean if a field has been set.

### GetCaptureTrafficControl

`func (o *GetMemberResponse) GetCaptureTrafficControl() map[string]interface{}`

GetCaptureTrafficControl returns the CaptureTrafficControl field if non-nil, zero value otherwise.

### GetCaptureTrafficControlOk

`func (o *GetMemberResponse) GetCaptureTrafficControlOk() (*map[string]interface{}, bool)`

GetCaptureTrafficControlOk returns a tuple with the CaptureTrafficControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureTrafficControl

`func (o *GetMemberResponse) SetCaptureTrafficControl(v map[string]interface{})`

SetCaptureTrafficControl sets CaptureTrafficControl field to given value.

### HasCaptureTrafficControl

`func (o *GetMemberResponse) HasCaptureTrafficControl() bool`

HasCaptureTrafficControl returns a boolean if a field has been set.

### GetCaptureTrafficStatus

`func (o *GetMemberResponse) GetCaptureTrafficStatus() map[string]interface{}`

GetCaptureTrafficStatus returns the CaptureTrafficStatus field if non-nil, zero value otherwise.

### GetCaptureTrafficStatusOk

`func (o *GetMemberResponse) GetCaptureTrafficStatusOk() (*map[string]interface{}, bool)`

GetCaptureTrafficStatusOk returns a tuple with the CaptureTrafficStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureTrafficStatus

`func (o *GetMemberResponse) SetCaptureTrafficStatus(v map[string]interface{})`

SetCaptureTrafficStatus sets CaptureTrafficStatus field to given value.

### HasCaptureTrafficStatus

`func (o *GetMemberResponse) HasCaptureTrafficStatus() bool`

HasCaptureTrafficStatus returns a boolean if a field has been set.

### GetComment

`func (o *GetMemberResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetMemberResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetMemberResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetMemberResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConfigAddrType

`func (o *GetMemberResponse) GetConfigAddrType() string`

GetConfigAddrType returns the ConfigAddrType field if non-nil, zero value otherwise.

### GetConfigAddrTypeOk

`func (o *GetMemberResponse) GetConfigAddrTypeOk() (*string, bool)`

GetConfigAddrTypeOk returns a tuple with the ConfigAddrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigAddrType

`func (o *GetMemberResponse) SetConfigAddrType(v string)`

SetConfigAddrType sets ConfigAddrType field to given value.

### HasConfigAddrType

`func (o *GetMemberResponse) HasConfigAddrType() bool`

HasConfigAddrType returns a boolean if a field has been set.

### GetCreateToken

`func (o *GetMemberResponse) GetCreateToken() map[string]interface{}`

GetCreateToken returns the CreateToken field if non-nil, zero value otherwise.

### GetCreateTokenOk

`func (o *GetMemberResponse) GetCreateTokenOk() (*map[string]interface{}, bool)`

GetCreateTokenOk returns a tuple with the CreateToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateToken

`func (o *GetMemberResponse) SetCreateToken(v map[string]interface{})`

SetCreateToken sets CreateToken field to given value.

### HasCreateToken

`func (o *GetMemberResponse) HasCreateToken() bool`

HasCreateToken returns a boolean if a field has been set.

### GetCspAccessKey

`func (o *GetMemberResponse) GetCspAccessKey() []string`

GetCspAccessKey returns the CspAccessKey field if non-nil, zero value otherwise.

### GetCspAccessKeyOk

`func (o *GetMemberResponse) GetCspAccessKeyOk() (*[]string, bool)`

GetCspAccessKeyOk returns a tuple with the CspAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspAccessKey

`func (o *GetMemberResponse) SetCspAccessKey(v []string)`

SetCspAccessKey sets CspAccessKey field to given value.

### HasCspAccessKey

`func (o *GetMemberResponse) HasCspAccessKey() bool`

HasCspAccessKey returns a boolean if a field has been set.

### GetCspMemberSetting

`func (o *GetMemberResponse) GetCspMemberSetting() MemberCspMemberSetting`

GetCspMemberSetting returns the CspMemberSetting field if non-nil, zero value otherwise.

### GetCspMemberSettingOk

`func (o *GetMemberResponse) GetCspMemberSettingOk() (*MemberCspMemberSetting, bool)`

GetCspMemberSettingOk returns a tuple with the CspMemberSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspMemberSetting

`func (o *GetMemberResponse) SetCspMemberSetting(v MemberCspMemberSetting)`

SetCspMemberSetting sets CspMemberSetting field to given value.

### HasCspMemberSetting

`func (o *GetMemberResponse) HasCspMemberSetting() bool`

HasCspMemberSetting returns a boolean if a field has been set.

### GetDnsResolverSetting

`func (o *GetMemberResponse) GetDnsResolverSetting() MemberDnsResolverSetting`

GetDnsResolverSetting returns the DnsResolverSetting field if non-nil, zero value otherwise.

### GetDnsResolverSettingOk

`func (o *GetMemberResponse) GetDnsResolverSettingOk() (*MemberDnsResolverSetting, bool)`

GetDnsResolverSettingOk returns a tuple with the DnsResolverSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsResolverSetting

`func (o *GetMemberResponse) SetDnsResolverSetting(v MemberDnsResolverSetting)`

SetDnsResolverSetting sets DnsResolverSetting field to given value.

### HasDnsResolverSetting

`func (o *GetMemberResponse) HasDnsResolverSetting() bool`

HasDnsResolverSetting returns a boolean if a field has been set.

### GetDscp

`func (o *GetMemberResponse) GetDscp() int64`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *GetMemberResponse) GetDscpOk() (*int64, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *GetMemberResponse) SetDscp(v int64)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *GetMemberResponse) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetEmailSetting

`func (o *GetMemberResponse) GetEmailSetting() MemberEmailSetting`

GetEmailSetting returns the EmailSetting field if non-nil, zero value otherwise.

### GetEmailSettingOk

`func (o *GetMemberResponse) GetEmailSettingOk() (*MemberEmailSetting, bool)`

GetEmailSettingOk returns a tuple with the EmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailSetting

`func (o *GetMemberResponse) SetEmailSetting(v MemberEmailSetting)`

SetEmailSetting sets EmailSetting field to given value.

### HasEmailSetting

`func (o *GetMemberResponse) HasEmailSetting() bool`

HasEmailSetting returns a boolean if a field has been set.

### GetEnableHa

`func (o *GetMemberResponse) GetEnableHa() bool`

GetEnableHa returns the EnableHa field if non-nil, zero value otherwise.

### GetEnableHaOk

`func (o *GetMemberResponse) GetEnableHaOk() (*bool, bool)`

GetEnableHaOk returns a tuple with the EnableHa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHa

`func (o *GetMemberResponse) SetEnableHa(v bool)`

SetEnableHa sets EnableHa field to given value.

### HasEnableHa

`func (o *GetMemberResponse) HasEnableHa() bool`

HasEnableHa returns a boolean if a field has been set.

### GetEnableLom

`func (o *GetMemberResponse) GetEnableLom() bool`

GetEnableLom returns the EnableLom field if non-nil, zero value otherwise.

### GetEnableLomOk

`func (o *GetMemberResponse) GetEnableLomOk() (*bool, bool)`

GetEnableLomOk returns a tuple with the EnableLom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableLom

`func (o *GetMemberResponse) SetEnableLom(v bool)`

SetEnableLom sets EnableLom field to given value.

### HasEnableLom

`func (o *GetMemberResponse) HasEnableLom() bool`

HasEnableLom returns a boolean if a field has been set.

### GetEnableMemberRedirect

`func (o *GetMemberResponse) GetEnableMemberRedirect() bool`

GetEnableMemberRedirect returns the EnableMemberRedirect field if non-nil, zero value otherwise.

### GetEnableMemberRedirectOk

`func (o *GetMemberResponse) GetEnableMemberRedirectOk() (*bool, bool)`

GetEnableMemberRedirectOk returns a tuple with the EnableMemberRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableMemberRedirect

`func (o *GetMemberResponse) SetEnableMemberRedirect(v bool)`

SetEnableMemberRedirect sets EnableMemberRedirect field to given value.

### HasEnableMemberRedirect

`func (o *GetMemberResponse) HasEnableMemberRedirect() bool`

HasEnableMemberRedirect returns a boolean if a field has been set.

### GetEnableRoApiAccess

`func (o *GetMemberResponse) GetEnableRoApiAccess() bool`

GetEnableRoApiAccess returns the EnableRoApiAccess field if non-nil, zero value otherwise.

### GetEnableRoApiAccessOk

`func (o *GetMemberResponse) GetEnableRoApiAccessOk() (*bool, bool)`

GetEnableRoApiAccessOk returns a tuple with the EnableRoApiAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRoApiAccess

`func (o *GetMemberResponse) SetEnableRoApiAccess(v bool)`

SetEnableRoApiAccess sets EnableRoApiAccess field to given value.

### HasEnableRoApiAccess

`func (o *GetMemberResponse) HasEnableRoApiAccess() bool`

HasEnableRoApiAccess returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetMemberResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetMemberResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetMemberResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetMemberResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetExternalSyslogBackupServers

`func (o *GetMemberResponse) GetExternalSyslogBackupServers() []MemberExternalSyslogBackupServers`

GetExternalSyslogBackupServers returns the ExternalSyslogBackupServers field if non-nil, zero value otherwise.

### GetExternalSyslogBackupServersOk

`func (o *GetMemberResponse) GetExternalSyslogBackupServersOk() (*[]MemberExternalSyslogBackupServers, bool)`

GetExternalSyslogBackupServersOk returns a tuple with the ExternalSyslogBackupServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSyslogBackupServers

`func (o *GetMemberResponse) SetExternalSyslogBackupServers(v []MemberExternalSyslogBackupServers)`

SetExternalSyslogBackupServers sets ExternalSyslogBackupServers field to given value.

### HasExternalSyslogBackupServers

`func (o *GetMemberResponse) HasExternalSyslogBackupServers() bool`

HasExternalSyslogBackupServers returns a boolean if a field has been set.

### GetExternalSyslogServerEnable

`func (o *GetMemberResponse) GetExternalSyslogServerEnable() bool`

GetExternalSyslogServerEnable returns the ExternalSyslogServerEnable field if non-nil, zero value otherwise.

### GetExternalSyslogServerEnableOk

`func (o *GetMemberResponse) GetExternalSyslogServerEnableOk() (*bool, bool)`

GetExternalSyslogServerEnableOk returns a tuple with the ExternalSyslogServerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSyslogServerEnable

`func (o *GetMemberResponse) SetExternalSyslogServerEnable(v bool)`

SetExternalSyslogServerEnable sets ExternalSyslogServerEnable field to given value.

### HasExternalSyslogServerEnable

`func (o *GetMemberResponse) HasExternalSyslogServerEnable() bool`

HasExternalSyslogServerEnable returns a boolean if a field has been set.

### GetHostName

`func (o *GetMemberResponse) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *GetMemberResponse) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *GetMemberResponse) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *GetMemberResponse) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetIpv6Setting

`func (o *GetMemberResponse) GetIpv6Setting() MemberIpv6Setting`

GetIpv6Setting returns the Ipv6Setting field if non-nil, zero value otherwise.

### GetIpv6SettingOk

`func (o *GetMemberResponse) GetIpv6SettingOk() (*MemberIpv6Setting, bool)`

GetIpv6SettingOk returns a tuple with the Ipv6Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Setting

`func (o *GetMemberResponse) SetIpv6Setting(v MemberIpv6Setting)`

SetIpv6Setting sets Ipv6Setting field to given value.

### HasIpv6Setting

`func (o *GetMemberResponse) HasIpv6Setting() bool`

HasIpv6Setting returns a boolean if a field has been set.

### GetIpv6StaticRoutes

`func (o *GetMemberResponse) GetIpv6StaticRoutes() []MemberIpv6StaticRoutes`

GetIpv6StaticRoutes returns the Ipv6StaticRoutes field if non-nil, zero value otherwise.

### GetIpv6StaticRoutesOk

`func (o *GetMemberResponse) GetIpv6StaticRoutesOk() (*[]MemberIpv6StaticRoutes, bool)`

GetIpv6StaticRoutesOk returns a tuple with the Ipv6StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6StaticRoutes

`func (o *GetMemberResponse) SetIpv6StaticRoutes(v []MemberIpv6StaticRoutes)`

SetIpv6StaticRoutes sets Ipv6StaticRoutes field to given value.

### HasIpv6StaticRoutes

`func (o *GetMemberResponse) HasIpv6StaticRoutes() bool`

HasIpv6StaticRoutes returns a boolean if a field has been set.

### GetIsDscpCapable

`func (o *GetMemberResponse) GetIsDscpCapable() bool`

GetIsDscpCapable returns the IsDscpCapable field if non-nil, zero value otherwise.

### GetIsDscpCapableOk

`func (o *GetMemberResponse) GetIsDscpCapableOk() (*bool, bool)`

GetIsDscpCapableOk returns a tuple with the IsDscpCapable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDscpCapable

`func (o *GetMemberResponse) SetIsDscpCapable(v bool)`

SetIsDscpCapable sets IsDscpCapable field to given value.

### HasIsDscpCapable

`func (o *GetMemberResponse) HasIsDscpCapable() bool`

HasIsDscpCapable returns a boolean if a field has been set.

### GetLan2Enabled

`func (o *GetMemberResponse) GetLan2Enabled() bool`

GetLan2Enabled returns the Lan2Enabled field if non-nil, zero value otherwise.

### GetLan2EnabledOk

`func (o *GetMemberResponse) GetLan2EnabledOk() (*bool, bool)`

GetLan2EnabledOk returns a tuple with the Lan2Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan2Enabled

`func (o *GetMemberResponse) SetLan2Enabled(v bool)`

SetLan2Enabled sets Lan2Enabled field to given value.

### HasLan2Enabled

`func (o *GetMemberResponse) HasLan2Enabled() bool`

HasLan2Enabled returns a boolean if a field has been set.

### GetLan2PortSetting

`func (o *GetMemberResponse) GetLan2PortSetting() MemberLan2PortSetting`

GetLan2PortSetting returns the Lan2PortSetting field if non-nil, zero value otherwise.

### GetLan2PortSettingOk

`func (o *GetMemberResponse) GetLan2PortSettingOk() (*MemberLan2PortSetting, bool)`

GetLan2PortSettingOk returns a tuple with the Lan2PortSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan2PortSetting

`func (o *GetMemberResponse) SetLan2PortSetting(v MemberLan2PortSetting)`

SetLan2PortSetting sets Lan2PortSetting field to given value.

### HasLan2PortSetting

`func (o *GetMemberResponse) HasLan2PortSetting() bool`

HasLan2PortSetting returns a boolean if a field has been set.

### GetLomNetworkConfig

`func (o *GetMemberResponse) GetLomNetworkConfig() []MemberLomNetworkConfig`

GetLomNetworkConfig returns the LomNetworkConfig field if non-nil, zero value otherwise.

### GetLomNetworkConfigOk

`func (o *GetMemberResponse) GetLomNetworkConfigOk() (*[]MemberLomNetworkConfig, bool)`

GetLomNetworkConfigOk returns a tuple with the LomNetworkConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomNetworkConfig

`func (o *GetMemberResponse) SetLomNetworkConfig(v []MemberLomNetworkConfig)`

SetLomNetworkConfig sets LomNetworkConfig field to given value.

### HasLomNetworkConfig

`func (o *GetMemberResponse) HasLomNetworkConfig() bool`

HasLomNetworkConfig returns a boolean if a field has been set.

### GetLomUsers

`func (o *GetMemberResponse) GetLomUsers() []MemberLomUsers`

GetLomUsers returns the LomUsers field if non-nil, zero value otherwise.

### GetLomUsersOk

`func (o *GetMemberResponse) GetLomUsersOk() (*[]MemberLomUsers, bool)`

GetLomUsersOk returns a tuple with the LomUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLomUsers

`func (o *GetMemberResponse) SetLomUsers(v []MemberLomUsers)`

SetLomUsers sets LomUsers field to given value.

### HasLomUsers

`func (o *GetMemberResponse) HasLomUsers() bool`

HasLomUsers returns a boolean if a field has been set.

### GetMasterCandidate

`func (o *GetMemberResponse) GetMasterCandidate() bool`

GetMasterCandidate returns the MasterCandidate field if non-nil, zero value otherwise.

### GetMasterCandidateOk

`func (o *GetMemberResponse) GetMasterCandidateOk() (*bool, bool)`

GetMasterCandidateOk returns a tuple with the MasterCandidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterCandidate

`func (o *GetMemberResponse) SetMasterCandidate(v bool)`

SetMasterCandidate sets MasterCandidate field to given value.

### HasMasterCandidate

`func (o *GetMemberResponse) HasMasterCandidate() bool`

HasMasterCandidate returns a boolean if a field has been set.

### GetMemberAdminOperation

`func (o *GetMemberResponse) GetMemberAdminOperation() map[string]interface{}`

GetMemberAdminOperation returns the MemberAdminOperation field if non-nil, zero value otherwise.

### GetMemberAdminOperationOk

`func (o *GetMemberResponse) GetMemberAdminOperationOk() (*map[string]interface{}, bool)`

GetMemberAdminOperationOk returns a tuple with the MemberAdminOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberAdminOperation

`func (o *GetMemberResponse) SetMemberAdminOperation(v map[string]interface{})`

SetMemberAdminOperation sets MemberAdminOperation field to given value.

### HasMemberAdminOperation

`func (o *GetMemberResponse) HasMemberAdminOperation() bool`

HasMemberAdminOperation returns a boolean if a field has been set.

### GetMemberServiceCommunication

`func (o *GetMemberResponse) GetMemberServiceCommunication() []MemberMemberServiceCommunication`

GetMemberServiceCommunication returns the MemberServiceCommunication field if non-nil, zero value otherwise.

### GetMemberServiceCommunicationOk

`func (o *GetMemberResponse) GetMemberServiceCommunicationOk() (*[]MemberMemberServiceCommunication, bool)`

GetMemberServiceCommunicationOk returns a tuple with the MemberServiceCommunication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberServiceCommunication

`func (o *GetMemberResponse) SetMemberServiceCommunication(v []MemberMemberServiceCommunication)`

SetMemberServiceCommunication sets MemberServiceCommunication field to given value.

### HasMemberServiceCommunication

`func (o *GetMemberResponse) HasMemberServiceCommunication() bool`

HasMemberServiceCommunication returns a boolean if a field has been set.

### GetMgmtPortSetting

`func (o *GetMemberResponse) GetMgmtPortSetting() MemberMgmtPortSetting`

GetMgmtPortSetting returns the MgmtPortSetting field if non-nil, zero value otherwise.

### GetMgmtPortSettingOk

`func (o *GetMemberResponse) GetMgmtPortSettingOk() (*MemberMgmtPortSetting, bool)`

GetMgmtPortSettingOk returns a tuple with the MgmtPortSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPortSetting

`func (o *GetMemberResponse) SetMgmtPortSetting(v MemberMgmtPortSetting)`

SetMgmtPortSetting sets MgmtPortSetting field to given value.

### HasMgmtPortSetting

`func (o *GetMemberResponse) HasMgmtPortSetting() bool`

HasMgmtPortSetting returns a boolean if a field has been set.

### GetMmdbEaBuildTime

`func (o *GetMemberResponse) GetMmdbEaBuildTime() int64`

GetMmdbEaBuildTime returns the MmdbEaBuildTime field if non-nil, zero value otherwise.

### GetMmdbEaBuildTimeOk

`func (o *GetMemberResponse) GetMmdbEaBuildTimeOk() (*int64, bool)`

GetMmdbEaBuildTimeOk returns a tuple with the MmdbEaBuildTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmdbEaBuildTime

`func (o *GetMemberResponse) SetMmdbEaBuildTime(v int64)`

SetMmdbEaBuildTime sets MmdbEaBuildTime field to given value.

### HasMmdbEaBuildTime

`func (o *GetMemberResponse) HasMmdbEaBuildTime() bool`

HasMmdbEaBuildTime returns a boolean if a field has been set.

### GetMmdbGeoipBuildTime

`func (o *GetMemberResponse) GetMmdbGeoipBuildTime() int64`

GetMmdbGeoipBuildTime returns the MmdbGeoipBuildTime field if non-nil, zero value otherwise.

### GetMmdbGeoipBuildTimeOk

`func (o *GetMemberResponse) GetMmdbGeoipBuildTimeOk() (*int64, bool)`

GetMmdbGeoipBuildTimeOk returns a tuple with the MmdbGeoipBuildTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMmdbGeoipBuildTime

`func (o *GetMemberResponse) SetMmdbGeoipBuildTime(v int64)`

SetMmdbGeoipBuildTime sets MmdbGeoipBuildTime field to given value.

### HasMmdbGeoipBuildTime

`func (o *GetMemberResponse) HasMmdbGeoipBuildTime() bool`

HasMmdbGeoipBuildTime returns a boolean if a field has been set.

### GetNatSetting

`func (o *GetMemberResponse) GetNatSetting() MemberNatSetting`

GetNatSetting returns the NatSetting field if non-nil, zero value otherwise.

### GetNatSettingOk

`func (o *GetMemberResponse) GetNatSettingOk() (*MemberNatSetting, bool)`

GetNatSettingOk returns a tuple with the NatSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatSetting

`func (o *GetMemberResponse) SetNatSetting(v MemberNatSetting)`

SetNatSetting sets NatSetting field to given value.

### HasNatSetting

`func (o *GetMemberResponse) HasNatSetting() bool`

HasNatSetting returns a boolean if a field has been set.

### GetNodeInfo

`func (o *GetMemberResponse) GetNodeInfo() []MemberNodeInfo`

GetNodeInfo returns the NodeInfo field if non-nil, zero value otherwise.

### GetNodeInfoOk

`func (o *GetMemberResponse) GetNodeInfoOk() (*[]MemberNodeInfo, bool)`

GetNodeInfoOk returns a tuple with the NodeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeInfo

`func (o *GetMemberResponse) SetNodeInfo(v []MemberNodeInfo)`

SetNodeInfo sets NodeInfo field to given value.

### HasNodeInfo

`func (o *GetMemberResponse) HasNodeInfo() bool`

HasNodeInfo returns a boolean if a field has been set.

### GetNtpSetting

`func (o *GetMemberResponse) GetNtpSetting() MemberNtpSetting`

GetNtpSetting returns the NtpSetting field if non-nil, zero value otherwise.

### GetNtpSettingOk

`func (o *GetMemberResponse) GetNtpSettingOk() (*MemberNtpSetting, bool)`

GetNtpSettingOk returns a tuple with the NtpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtpSetting

`func (o *GetMemberResponse) SetNtpSetting(v MemberNtpSetting)`

SetNtpSetting sets NtpSetting field to given value.

### HasNtpSetting

`func (o *GetMemberResponse) HasNtpSetting() bool`

HasNtpSetting returns a boolean if a field has been set.

### GetOspfList

`func (o *GetMemberResponse) GetOspfList() []MemberOspfList`

GetOspfList returns the OspfList field if non-nil, zero value otherwise.

### GetOspfListOk

`func (o *GetMemberResponse) GetOspfListOk() (*[]MemberOspfList, bool)`

GetOspfListOk returns a tuple with the OspfList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOspfList

`func (o *GetMemberResponse) SetOspfList(v []MemberOspfList)`

SetOspfList sets OspfList field to given value.

### HasOspfList

`func (o *GetMemberResponse) HasOspfList() bool`

HasOspfList returns a boolean if a field has been set.

### GetPassiveHaArpEnabled

`func (o *GetMemberResponse) GetPassiveHaArpEnabled() bool`

GetPassiveHaArpEnabled returns the PassiveHaArpEnabled field if non-nil, zero value otherwise.

### GetPassiveHaArpEnabledOk

`func (o *GetMemberResponse) GetPassiveHaArpEnabledOk() (*bool, bool)`

GetPassiveHaArpEnabledOk returns a tuple with the PassiveHaArpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveHaArpEnabled

`func (o *GetMemberResponse) SetPassiveHaArpEnabled(v bool)`

SetPassiveHaArpEnabled sets PassiveHaArpEnabled field to given value.

### HasPassiveHaArpEnabled

`func (o *GetMemberResponse) HasPassiveHaArpEnabled() bool`

HasPassiveHaArpEnabled returns a boolean if a field has been set.

### GetPlatform

`func (o *GetMemberResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetMemberResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetMemberResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetMemberResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPreProvisioning

`func (o *GetMemberResponse) GetPreProvisioning() MemberPreProvisioning`

GetPreProvisioning returns the PreProvisioning field if non-nil, zero value otherwise.

### GetPreProvisioningOk

`func (o *GetMemberResponse) GetPreProvisioningOk() (*MemberPreProvisioning, bool)`

GetPreProvisioningOk returns a tuple with the PreProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreProvisioning

`func (o *GetMemberResponse) SetPreProvisioning(v MemberPreProvisioning)`

SetPreProvisioning sets PreProvisioning field to given value.

### HasPreProvisioning

`func (o *GetMemberResponse) HasPreProvisioning() bool`

HasPreProvisioning returns a boolean if a field has been set.

### GetPreserveIfOwnsDelegation

`func (o *GetMemberResponse) GetPreserveIfOwnsDelegation() bool`

GetPreserveIfOwnsDelegation returns the PreserveIfOwnsDelegation field if non-nil, zero value otherwise.

### GetPreserveIfOwnsDelegationOk

`func (o *GetMemberResponse) GetPreserveIfOwnsDelegationOk() (*bool, bool)`

GetPreserveIfOwnsDelegationOk returns a tuple with the PreserveIfOwnsDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveIfOwnsDelegation

`func (o *GetMemberResponse) SetPreserveIfOwnsDelegation(v bool)`

SetPreserveIfOwnsDelegation sets PreserveIfOwnsDelegation field to given value.

### HasPreserveIfOwnsDelegation

`func (o *GetMemberResponse) HasPreserveIfOwnsDelegation() bool`

HasPreserveIfOwnsDelegation returns a boolean if a field has been set.

### GetReadToken

`func (o *GetMemberResponse) GetReadToken() map[string]interface{}`

GetReadToken returns the ReadToken field if non-nil, zero value otherwise.

### GetReadTokenOk

`func (o *GetMemberResponse) GetReadTokenOk() (*map[string]interface{}, bool)`

GetReadTokenOk returns a tuple with the ReadToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadToken

`func (o *GetMemberResponse) SetReadToken(v map[string]interface{})`

SetReadToken sets ReadToken field to given value.

### HasReadToken

`func (o *GetMemberResponse) HasReadToken() bool`

HasReadToken returns a boolean if a field has been set.

### GetRemoteConsoleAccessEnable

`func (o *GetMemberResponse) GetRemoteConsoleAccessEnable() bool`

GetRemoteConsoleAccessEnable returns the RemoteConsoleAccessEnable field if non-nil, zero value otherwise.

### GetRemoteConsoleAccessEnableOk

`func (o *GetMemberResponse) GetRemoteConsoleAccessEnableOk() (*bool, bool)`

GetRemoteConsoleAccessEnableOk returns a tuple with the RemoteConsoleAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteConsoleAccessEnable

`func (o *GetMemberResponse) SetRemoteConsoleAccessEnable(v bool)`

SetRemoteConsoleAccessEnable sets RemoteConsoleAccessEnable field to given value.

### HasRemoteConsoleAccessEnable

`func (o *GetMemberResponse) HasRemoteConsoleAccessEnable() bool`

HasRemoteConsoleAccessEnable returns a boolean if a field has been set.

### GetRequestrestartservicestatus

`func (o *GetMemberResponse) GetRequestrestartservicestatus() map[string]interface{}`

GetRequestrestartservicestatus returns the Requestrestartservicestatus field if non-nil, zero value otherwise.

### GetRequestrestartservicestatusOk

`func (o *GetMemberResponse) GetRequestrestartservicestatusOk() (*map[string]interface{}, bool)`

GetRequestrestartservicestatusOk returns a tuple with the Requestrestartservicestatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestrestartservicestatus

`func (o *GetMemberResponse) SetRequestrestartservicestatus(v map[string]interface{})`

SetRequestrestartservicestatus sets Requestrestartservicestatus field to given value.

### HasRequestrestartservicestatus

`func (o *GetMemberResponse) HasRequestrestartservicestatus() bool`

HasRequestrestartservicestatus returns a boolean if a field has been set.

### GetRestartservices

`func (o *GetMemberResponse) GetRestartservices() map[string]interface{}`

GetRestartservices returns the Restartservices field if non-nil, zero value otherwise.

### GetRestartservicesOk

`func (o *GetMemberResponse) GetRestartservicesOk() (*map[string]interface{}, bool)`

GetRestartservicesOk returns a tuple with the Restartservices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartservices

`func (o *GetMemberResponse) SetRestartservices(v map[string]interface{})`

SetRestartservices sets Restartservices field to given value.

### HasRestartservices

`func (o *GetMemberResponse) HasRestartservices() bool`

HasRestartservices returns a boolean if a field has been set.

### GetRouterId

`func (o *GetMemberResponse) GetRouterId() int64`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *GetMemberResponse) GetRouterIdOk() (*int64, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *GetMemberResponse) SetRouterId(v int64)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *GetMemberResponse) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetServiceStatus

`func (o *GetMemberResponse) GetServiceStatus() []MemberServiceStatus`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *GetMemberResponse) GetServiceStatusOk() (*[]MemberServiceStatus, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *GetMemberResponse) SetServiceStatus(v []MemberServiceStatus)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *GetMemberResponse) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.

### GetServiceTypeConfiguration

`func (o *GetMemberResponse) GetServiceTypeConfiguration() string`

GetServiceTypeConfiguration returns the ServiceTypeConfiguration field if non-nil, zero value otherwise.

### GetServiceTypeConfigurationOk

`func (o *GetMemberResponse) GetServiceTypeConfigurationOk() (*string, bool)`

GetServiceTypeConfigurationOk returns a tuple with the ServiceTypeConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeConfiguration

`func (o *GetMemberResponse) SetServiceTypeConfiguration(v string)`

SetServiceTypeConfiguration sets ServiceTypeConfiguration field to given value.

### HasServiceTypeConfiguration

`func (o *GetMemberResponse) HasServiceTypeConfiguration() bool`

HasServiceTypeConfiguration returns a boolean if a field has been set.

### GetSnmpSetting

`func (o *GetMemberResponse) GetSnmpSetting() MemberSnmpSetting`

GetSnmpSetting returns the SnmpSetting field if non-nil, zero value otherwise.

### GetSnmpSettingOk

`func (o *GetMemberResponse) GetSnmpSettingOk() (*MemberSnmpSetting, bool)`

GetSnmpSettingOk returns a tuple with the SnmpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpSetting

`func (o *GetMemberResponse) SetSnmpSetting(v MemberSnmpSetting)`

SetSnmpSetting sets SnmpSetting field to given value.

### HasSnmpSetting

`func (o *GetMemberResponse) HasSnmpSetting() bool`

HasSnmpSetting returns a boolean if a field has been set.

### GetStaticRoutes

`func (o *GetMemberResponse) GetStaticRoutes() []MemberStaticRoutes`

GetStaticRoutes returns the StaticRoutes field if non-nil, zero value otherwise.

### GetStaticRoutesOk

`func (o *GetMemberResponse) GetStaticRoutesOk() (*[]MemberStaticRoutes, bool)`

GetStaticRoutesOk returns a tuple with the StaticRoutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticRoutes

`func (o *GetMemberResponse) SetStaticRoutes(v []MemberStaticRoutes)`

SetStaticRoutes sets StaticRoutes field to given value.

### HasStaticRoutes

`func (o *GetMemberResponse) HasStaticRoutes() bool`

HasStaticRoutes returns a boolean if a field has been set.

### GetSupportAccessEnable

`func (o *GetMemberResponse) GetSupportAccessEnable() bool`

GetSupportAccessEnable returns the SupportAccessEnable field if non-nil, zero value otherwise.

### GetSupportAccessEnableOk

`func (o *GetMemberResponse) GetSupportAccessEnableOk() (*bool, bool)`

GetSupportAccessEnableOk returns a tuple with the SupportAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAccessEnable

`func (o *GetMemberResponse) SetSupportAccessEnable(v bool)`

SetSupportAccessEnable sets SupportAccessEnable field to given value.

### HasSupportAccessEnable

`func (o *GetMemberResponse) HasSupportAccessEnable() bool`

HasSupportAccessEnable returns a boolean if a field has been set.

### GetSupportAccessInfo

`func (o *GetMemberResponse) GetSupportAccessInfo() string`

GetSupportAccessInfo returns the SupportAccessInfo field if non-nil, zero value otherwise.

### GetSupportAccessInfoOk

`func (o *GetMemberResponse) GetSupportAccessInfoOk() (*string, bool)`

GetSupportAccessInfoOk returns a tuple with the SupportAccessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAccessInfo

`func (o *GetMemberResponse) SetSupportAccessInfo(v string)`

SetSupportAccessInfo sets SupportAccessInfo field to given value.

### HasSupportAccessInfo

`func (o *GetMemberResponse) HasSupportAccessInfo() bool`

HasSupportAccessInfo returns a boolean if a field has been set.

### GetSyslogProxySetting

`func (o *GetMemberResponse) GetSyslogProxySetting() MemberSyslogProxySetting`

GetSyslogProxySetting returns the SyslogProxySetting field if non-nil, zero value otherwise.

### GetSyslogProxySettingOk

`func (o *GetMemberResponse) GetSyslogProxySettingOk() (*MemberSyslogProxySetting, bool)`

GetSyslogProxySettingOk returns a tuple with the SyslogProxySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogProxySetting

`func (o *GetMemberResponse) SetSyslogProxySetting(v MemberSyslogProxySetting)`

SetSyslogProxySetting sets SyslogProxySetting field to given value.

### HasSyslogProxySetting

`func (o *GetMemberResponse) HasSyslogProxySetting() bool`

HasSyslogProxySetting returns a boolean if a field has been set.

### GetSyslogServers

`func (o *GetMemberResponse) GetSyslogServers() []MemberSyslogServers`

GetSyslogServers returns the SyslogServers field if non-nil, zero value otherwise.

### GetSyslogServersOk

`func (o *GetMemberResponse) GetSyslogServersOk() (*[]MemberSyslogServers, bool)`

GetSyslogServersOk returns a tuple with the SyslogServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogServers

`func (o *GetMemberResponse) SetSyslogServers(v []MemberSyslogServers)`

SetSyslogServers sets SyslogServers field to given value.

### HasSyslogServers

`func (o *GetMemberResponse) HasSyslogServers() bool`

HasSyslogServers returns a boolean if a field has been set.

### GetSyslogSize

`func (o *GetMemberResponse) GetSyslogSize() int64`

GetSyslogSize returns the SyslogSize field if non-nil, zero value otherwise.

### GetSyslogSizeOk

`func (o *GetMemberResponse) GetSyslogSizeOk() (*int64, bool)`

GetSyslogSizeOk returns a tuple with the SyslogSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyslogSize

`func (o *GetMemberResponse) SetSyslogSize(v int64)`

SetSyslogSize sets SyslogSize field to given value.

### HasSyslogSize

`func (o *GetMemberResponse) HasSyslogSize() bool`

HasSyslogSize returns a boolean if a field has been set.

### GetThresholdTraps

`func (o *GetMemberResponse) GetThresholdTraps() []MemberThresholdTraps`

GetThresholdTraps returns the ThresholdTraps field if non-nil, zero value otherwise.

### GetThresholdTrapsOk

`func (o *GetMemberResponse) GetThresholdTrapsOk() (*[]MemberThresholdTraps, bool)`

GetThresholdTrapsOk returns a tuple with the ThresholdTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdTraps

`func (o *GetMemberResponse) SetThresholdTraps(v []MemberThresholdTraps)`

SetThresholdTraps sets ThresholdTraps field to given value.

### HasThresholdTraps

`func (o *GetMemberResponse) HasThresholdTraps() bool`

HasThresholdTraps returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetMemberResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetMemberResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetMemberResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetMemberResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetTrafficCaptureAuthDnsSetting

`func (o *GetMemberResponse) GetTrafficCaptureAuthDnsSetting() MemberTrafficCaptureAuthDnsSetting`

GetTrafficCaptureAuthDnsSetting returns the TrafficCaptureAuthDnsSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureAuthDnsSettingOk

`func (o *GetMemberResponse) GetTrafficCaptureAuthDnsSettingOk() (*MemberTrafficCaptureAuthDnsSetting, bool)`

GetTrafficCaptureAuthDnsSettingOk returns a tuple with the TrafficCaptureAuthDnsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureAuthDnsSetting

`func (o *GetMemberResponse) SetTrafficCaptureAuthDnsSetting(v MemberTrafficCaptureAuthDnsSetting)`

SetTrafficCaptureAuthDnsSetting sets TrafficCaptureAuthDnsSetting field to given value.

### HasTrafficCaptureAuthDnsSetting

`func (o *GetMemberResponse) HasTrafficCaptureAuthDnsSetting() bool`

HasTrafficCaptureAuthDnsSetting returns a boolean if a field has been set.

### GetTrafficCaptureChrSetting

`func (o *GetMemberResponse) GetTrafficCaptureChrSetting() MemberTrafficCaptureChrSetting`

GetTrafficCaptureChrSetting returns the TrafficCaptureChrSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureChrSettingOk

`func (o *GetMemberResponse) GetTrafficCaptureChrSettingOk() (*MemberTrafficCaptureChrSetting, bool)`

GetTrafficCaptureChrSettingOk returns a tuple with the TrafficCaptureChrSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureChrSetting

`func (o *GetMemberResponse) SetTrafficCaptureChrSetting(v MemberTrafficCaptureChrSetting)`

SetTrafficCaptureChrSetting sets TrafficCaptureChrSetting field to given value.

### HasTrafficCaptureChrSetting

`func (o *GetMemberResponse) HasTrafficCaptureChrSetting() bool`

HasTrafficCaptureChrSetting returns a boolean if a field has been set.

### GetTrafficCaptureQpsSetting

`func (o *GetMemberResponse) GetTrafficCaptureQpsSetting() MemberTrafficCaptureQpsSetting`

GetTrafficCaptureQpsSetting returns the TrafficCaptureQpsSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureQpsSettingOk

`func (o *GetMemberResponse) GetTrafficCaptureQpsSettingOk() (*MemberTrafficCaptureQpsSetting, bool)`

GetTrafficCaptureQpsSettingOk returns a tuple with the TrafficCaptureQpsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureQpsSetting

`func (o *GetMemberResponse) SetTrafficCaptureQpsSetting(v MemberTrafficCaptureQpsSetting)`

SetTrafficCaptureQpsSetting sets TrafficCaptureQpsSetting field to given value.

### HasTrafficCaptureQpsSetting

`func (o *GetMemberResponse) HasTrafficCaptureQpsSetting() bool`

HasTrafficCaptureQpsSetting returns a boolean if a field has been set.

### GetTrafficCaptureRecDnsSetting

`func (o *GetMemberResponse) GetTrafficCaptureRecDnsSetting() MemberTrafficCaptureRecDnsSetting`

GetTrafficCaptureRecDnsSetting returns the TrafficCaptureRecDnsSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureRecDnsSettingOk

`func (o *GetMemberResponse) GetTrafficCaptureRecDnsSettingOk() (*MemberTrafficCaptureRecDnsSetting, bool)`

GetTrafficCaptureRecDnsSettingOk returns a tuple with the TrafficCaptureRecDnsSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureRecDnsSetting

`func (o *GetMemberResponse) SetTrafficCaptureRecDnsSetting(v MemberTrafficCaptureRecDnsSetting)`

SetTrafficCaptureRecDnsSetting sets TrafficCaptureRecDnsSetting field to given value.

### HasTrafficCaptureRecDnsSetting

`func (o *GetMemberResponse) HasTrafficCaptureRecDnsSetting() bool`

HasTrafficCaptureRecDnsSetting returns a boolean if a field has been set.

### GetTrafficCaptureRecQueriesSetting

`func (o *GetMemberResponse) GetTrafficCaptureRecQueriesSetting() MemberTrafficCaptureRecQueriesSetting`

GetTrafficCaptureRecQueriesSetting returns the TrafficCaptureRecQueriesSetting field if non-nil, zero value otherwise.

### GetTrafficCaptureRecQueriesSettingOk

`func (o *GetMemberResponse) GetTrafficCaptureRecQueriesSettingOk() (*MemberTrafficCaptureRecQueriesSetting, bool)`

GetTrafficCaptureRecQueriesSettingOk returns a tuple with the TrafficCaptureRecQueriesSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficCaptureRecQueriesSetting

`func (o *GetMemberResponse) SetTrafficCaptureRecQueriesSetting(v MemberTrafficCaptureRecQueriesSetting)`

SetTrafficCaptureRecQueriesSetting sets TrafficCaptureRecQueriesSetting field to given value.

### HasTrafficCaptureRecQueriesSetting

`func (o *GetMemberResponse) HasTrafficCaptureRecQueriesSetting() bool`

HasTrafficCaptureRecQueriesSetting returns a boolean if a field has been set.

### GetTrapNotifications

`func (o *GetMemberResponse) GetTrapNotifications() []MemberTrapNotifications`

GetTrapNotifications returns the TrapNotifications field if non-nil, zero value otherwise.

### GetTrapNotificationsOk

`func (o *GetMemberResponse) GetTrapNotificationsOk() (*[]MemberTrapNotifications, bool)`

GetTrapNotificationsOk returns a tuple with the TrapNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrapNotifications

`func (o *GetMemberResponse) SetTrapNotifications(v []MemberTrapNotifications)`

SetTrapNotifications sets TrapNotifications field to given value.

### HasTrapNotifications

`func (o *GetMemberResponse) HasTrapNotifications() bool`

HasTrapNotifications returns a boolean if a field has been set.

### GetUpgradeGroup

`func (o *GetMemberResponse) GetUpgradeGroup() string`

GetUpgradeGroup returns the UpgradeGroup field if non-nil, zero value otherwise.

### GetUpgradeGroupOk

`func (o *GetMemberResponse) GetUpgradeGroupOk() (*string, bool)`

GetUpgradeGroupOk returns a tuple with the UpgradeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeGroup

`func (o *GetMemberResponse) SetUpgradeGroup(v string)`

SetUpgradeGroup sets UpgradeGroup field to given value.

### HasUpgradeGroup

`func (o *GetMemberResponse) HasUpgradeGroup() bool`

HasUpgradeGroup returns a boolean if a field has been set.

### GetUseAutomatedTrafficCapture

`func (o *GetMemberResponse) GetUseAutomatedTrafficCapture() bool`

GetUseAutomatedTrafficCapture returns the UseAutomatedTrafficCapture field if non-nil, zero value otherwise.

### GetUseAutomatedTrafficCaptureOk

`func (o *GetMemberResponse) GetUseAutomatedTrafficCaptureOk() (*bool, bool)`

GetUseAutomatedTrafficCaptureOk returns a tuple with the UseAutomatedTrafficCapture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAutomatedTrafficCapture

`func (o *GetMemberResponse) SetUseAutomatedTrafficCapture(v bool)`

SetUseAutomatedTrafficCapture sets UseAutomatedTrafficCapture field to given value.

### HasUseAutomatedTrafficCapture

`func (o *GetMemberResponse) HasUseAutomatedTrafficCapture() bool`

HasUseAutomatedTrafficCapture returns a boolean if a field has been set.

### GetUseDnsResolverSetting

`func (o *GetMemberResponse) GetUseDnsResolverSetting() bool`

GetUseDnsResolverSetting returns the UseDnsResolverSetting field if non-nil, zero value otherwise.

### GetUseDnsResolverSettingOk

`func (o *GetMemberResponse) GetUseDnsResolverSettingOk() (*bool, bool)`

GetUseDnsResolverSettingOk returns a tuple with the UseDnsResolverSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnsResolverSetting

`func (o *GetMemberResponse) SetUseDnsResolverSetting(v bool)`

SetUseDnsResolverSetting sets UseDnsResolverSetting field to given value.

### HasUseDnsResolverSetting

`func (o *GetMemberResponse) HasUseDnsResolverSetting() bool`

HasUseDnsResolverSetting returns a boolean if a field has been set.

### GetUseDscp

`func (o *GetMemberResponse) GetUseDscp() bool`

GetUseDscp returns the UseDscp field if non-nil, zero value otherwise.

### GetUseDscpOk

`func (o *GetMemberResponse) GetUseDscpOk() (*bool, bool)`

GetUseDscpOk returns a tuple with the UseDscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDscp

`func (o *GetMemberResponse) SetUseDscp(v bool)`

SetUseDscp sets UseDscp field to given value.

### HasUseDscp

`func (o *GetMemberResponse) HasUseDscp() bool`

HasUseDscp returns a boolean if a field has been set.

### GetUseEmailSetting

`func (o *GetMemberResponse) GetUseEmailSetting() bool`

GetUseEmailSetting returns the UseEmailSetting field if non-nil, zero value otherwise.

### GetUseEmailSettingOk

`func (o *GetMemberResponse) GetUseEmailSettingOk() (*bool, bool)`

GetUseEmailSettingOk returns a tuple with the UseEmailSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEmailSetting

`func (o *GetMemberResponse) SetUseEmailSetting(v bool)`

SetUseEmailSetting sets UseEmailSetting field to given value.

### HasUseEmailSetting

`func (o *GetMemberResponse) HasUseEmailSetting() bool`

HasUseEmailSetting returns a boolean if a field has been set.

### GetUseEnableLom

`func (o *GetMemberResponse) GetUseEnableLom() bool`

GetUseEnableLom returns the UseEnableLom field if non-nil, zero value otherwise.

### GetUseEnableLomOk

`func (o *GetMemberResponse) GetUseEnableLomOk() (*bool, bool)`

GetUseEnableLomOk returns a tuple with the UseEnableLom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableLom

`func (o *GetMemberResponse) SetUseEnableLom(v bool)`

SetUseEnableLom sets UseEnableLom field to given value.

### HasUseEnableLom

`func (o *GetMemberResponse) HasUseEnableLom() bool`

HasUseEnableLom returns a boolean if a field has been set.

### GetUseEnableMemberRedirect

`func (o *GetMemberResponse) GetUseEnableMemberRedirect() bool`

GetUseEnableMemberRedirect returns the UseEnableMemberRedirect field if non-nil, zero value otherwise.

### GetUseEnableMemberRedirectOk

`func (o *GetMemberResponse) GetUseEnableMemberRedirectOk() (*bool, bool)`

GetUseEnableMemberRedirectOk returns a tuple with the UseEnableMemberRedirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseEnableMemberRedirect

`func (o *GetMemberResponse) SetUseEnableMemberRedirect(v bool)`

SetUseEnableMemberRedirect sets UseEnableMemberRedirect field to given value.

### HasUseEnableMemberRedirect

`func (o *GetMemberResponse) HasUseEnableMemberRedirect() bool`

HasUseEnableMemberRedirect returns a boolean if a field has been set.

### GetUseExternalSyslogBackupServers

`func (o *GetMemberResponse) GetUseExternalSyslogBackupServers() bool`

GetUseExternalSyslogBackupServers returns the UseExternalSyslogBackupServers field if non-nil, zero value otherwise.

### GetUseExternalSyslogBackupServersOk

`func (o *GetMemberResponse) GetUseExternalSyslogBackupServersOk() (*bool, bool)`

GetUseExternalSyslogBackupServersOk returns a tuple with the UseExternalSyslogBackupServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExternalSyslogBackupServers

`func (o *GetMemberResponse) SetUseExternalSyslogBackupServers(v bool)`

SetUseExternalSyslogBackupServers sets UseExternalSyslogBackupServers field to given value.

### HasUseExternalSyslogBackupServers

`func (o *GetMemberResponse) HasUseExternalSyslogBackupServers() bool`

HasUseExternalSyslogBackupServers returns a boolean if a field has been set.

### GetUseRemoteConsoleAccessEnable

`func (o *GetMemberResponse) GetUseRemoteConsoleAccessEnable() bool`

GetUseRemoteConsoleAccessEnable returns the UseRemoteConsoleAccessEnable field if non-nil, zero value otherwise.

### GetUseRemoteConsoleAccessEnableOk

`func (o *GetMemberResponse) GetUseRemoteConsoleAccessEnableOk() (*bool, bool)`

GetUseRemoteConsoleAccessEnableOk returns a tuple with the UseRemoteConsoleAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRemoteConsoleAccessEnable

`func (o *GetMemberResponse) SetUseRemoteConsoleAccessEnable(v bool)`

SetUseRemoteConsoleAccessEnable sets UseRemoteConsoleAccessEnable field to given value.

### HasUseRemoteConsoleAccessEnable

`func (o *GetMemberResponse) HasUseRemoteConsoleAccessEnable() bool`

HasUseRemoteConsoleAccessEnable returns a boolean if a field has been set.

### GetUseSnmpSetting

`func (o *GetMemberResponse) GetUseSnmpSetting() bool`

GetUseSnmpSetting returns the UseSnmpSetting field if non-nil, zero value otherwise.

### GetUseSnmpSettingOk

`func (o *GetMemberResponse) GetUseSnmpSettingOk() (*bool, bool)`

GetUseSnmpSettingOk returns a tuple with the UseSnmpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpSetting

`func (o *GetMemberResponse) SetUseSnmpSetting(v bool)`

SetUseSnmpSetting sets UseSnmpSetting field to given value.

### HasUseSnmpSetting

`func (o *GetMemberResponse) HasUseSnmpSetting() bool`

HasUseSnmpSetting returns a boolean if a field has been set.

### GetUseSupportAccessEnable

`func (o *GetMemberResponse) GetUseSupportAccessEnable() bool`

GetUseSupportAccessEnable returns the UseSupportAccessEnable field if non-nil, zero value otherwise.

### GetUseSupportAccessEnableOk

`func (o *GetMemberResponse) GetUseSupportAccessEnableOk() (*bool, bool)`

GetUseSupportAccessEnableOk returns a tuple with the UseSupportAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSupportAccessEnable

`func (o *GetMemberResponse) SetUseSupportAccessEnable(v bool)`

SetUseSupportAccessEnable sets UseSupportAccessEnable field to given value.

### HasUseSupportAccessEnable

`func (o *GetMemberResponse) HasUseSupportAccessEnable() bool`

HasUseSupportAccessEnable returns a boolean if a field has been set.

### GetUseSyslogProxySetting

`func (o *GetMemberResponse) GetUseSyslogProxySetting() bool`

GetUseSyslogProxySetting returns the UseSyslogProxySetting field if non-nil, zero value otherwise.

### GetUseSyslogProxySettingOk

`func (o *GetMemberResponse) GetUseSyslogProxySettingOk() (*bool, bool)`

GetUseSyslogProxySettingOk returns a tuple with the UseSyslogProxySetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSyslogProxySetting

`func (o *GetMemberResponse) SetUseSyslogProxySetting(v bool)`

SetUseSyslogProxySetting sets UseSyslogProxySetting field to given value.

### HasUseSyslogProxySetting

`func (o *GetMemberResponse) HasUseSyslogProxySetting() bool`

HasUseSyslogProxySetting returns a boolean if a field has been set.

### GetUseThresholdTraps

`func (o *GetMemberResponse) GetUseThresholdTraps() bool`

GetUseThresholdTraps returns the UseThresholdTraps field if non-nil, zero value otherwise.

### GetUseThresholdTrapsOk

`func (o *GetMemberResponse) GetUseThresholdTrapsOk() (*bool, bool)`

GetUseThresholdTrapsOk returns a tuple with the UseThresholdTraps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseThresholdTraps

`func (o *GetMemberResponse) SetUseThresholdTraps(v bool)`

SetUseThresholdTraps sets UseThresholdTraps field to given value.

### HasUseThresholdTraps

`func (o *GetMemberResponse) HasUseThresholdTraps() bool`

HasUseThresholdTraps returns a boolean if a field has been set.

### GetUseTimeZone

`func (o *GetMemberResponse) GetUseTimeZone() bool`

GetUseTimeZone returns the UseTimeZone field if non-nil, zero value otherwise.

### GetUseTimeZoneOk

`func (o *GetMemberResponse) GetUseTimeZoneOk() (*bool, bool)`

GetUseTimeZoneOk returns a tuple with the UseTimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTimeZone

`func (o *GetMemberResponse) SetUseTimeZone(v bool)`

SetUseTimeZone sets UseTimeZone field to given value.

### HasUseTimeZone

`func (o *GetMemberResponse) HasUseTimeZone() bool`

HasUseTimeZone returns a boolean if a field has been set.

### GetUseTrafficCaptureAuthDns

`func (o *GetMemberResponse) GetUseTrafficCaptureAuthDns() bool`

GetUseTrafficCaptureAuthDns returns the UseTrafficCaptureAuthDns field if non-nil, zero value otherwise.

### GetUseTrafficCaptureAuthDnsOk

`func (o *GetMemberResponse) GetUseTrafficCaptureAuthDnsOk() (*bool, bool)`

GetUseTrafficCaptureAuthDnsOk returns a tuple with the UseTrafficCaptureAuthDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureAuthDns

`func (o *GetMemberResponse) SetUseTrafficCaptureAuthDns(v bool)`

SetUseTrafficCaptureAuthDns sets UseTrafficCaptureAuthDns field to given value.

### HasUseTrafficCaptureAuthDns

`func (o *GetMemberResponse) HasUseTrafficCaptureAuthDns() bool`

HasUseTrafficCaptureAuthDns returns a boolean if a field has been set.

### GetUseTrafficCaptureChr

`func (o *GetMemberResponse) GetUseTrafficCaptureChr() bool`

GetUseTrafficCaptureChr returns the UseTrafficCaptureChr field if non-nil, zero value otherwise.

### GetUseTrafficCaptureChrOk

`func (o *GetMemberResponse) GetUseTrafficCaptureChrOk() (*bool, bool)`

GetUseTrafficCaptureChrOk returns a tuple with the UseTrafficCaptureChr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureChr

`func (o *GetMemberResponse) SetUseTrafficCaptureChr(v bool)`

SetUseTrafficCaptureChr sets UseTrafficCaptureChr field to given value.

### HasUseTrafficCaptureChr

`func (o *GetMemberResponse) HasUseTrafficCaptureChr() bool`

HasUseTrafficCaptureChr returns a boolean if a field has been set.

### GetUseTrafficCaptureQps

`func (o *GetMemberResponse) GetUseTrafficCaptureQps() bool`

GetUseTrafficCaptureQps returns the UseTrafficCaptureQps field if non-nil, zero value otherwise.

### GetUseTrafficCaptureQpsOk

`func (o *GetMemberResponse) GetUseTrafficCaptureQpsOk() (*bool, bool)`

GetUseTrafficCaptureQpsOk returns a tuple with the UseTrafficCaptureQps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureQps

`func (o *GetMemberResponse) SetUseTrafficCaptureQps(v bool)`

SetUseTrafficCaptureQps sets UseTrafficCaptureQps field to given value.

### HasUseTrafficCaptureQps

`func (o *GetMemberResponse) HasUseTrafficCaptureQps() bool`

HasUseTrafficCaptureQps returns a boolean if a field has been set.

### GetUseTrafficCaptureRecDns

`func (o *GetMemberResponse) GetUseTrafficCaptureRecDns() bool`

GetUseTrafficCaptureRecDns returns the UseTrafficCaptureRecDns field if non-nil, zero value otherwise.

### GetUseTrafficCaptureRecDnsOk

`func (o *GetMemberResponse) GetUseTrafficCaptureRecDnsOk() (*bool, bool)`

GetUseTrafficCaptureRecDnsOk returns a tuple with the UseTrafficCaptureRecDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureRecDns

`func (o *GetMemberResponse) SetUseTrafficCaptureRecDns(v bool)`

SetUseTrafficCaptureRecDns sets UseTrafficCaptureRecDns field to given value.

### HasUseTrafficCaptureRecDns

`func (o *GetMemberResponse) HasUseTrafficCaptureRecDns() bool`

HasUseTrafficCaptureRecDns returns a boolean if a field has been set.

### GetUseTrafficCaptureRecQueries

`func (o *GetMemberResponse) GetUseTrafficCaptureRecQueries() bool`

GetUseTrafficCaptureRecQueries returns the UseTrafficCaptureRecQueries field if non-nil, zero value otherwise.

### GetUseTrafficCaptureRecQueriesOk

`func (o *GetMemberResponse) GetUseTrafficCaptureRecQueriesOk() (*bool, bool)`

GetUseTrafficCaptureRecQueriesOk returns a tuple with the UseTrafficCaptureRecQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrafficCaptureRecQueries

`func (o *GetMemberResponse) SetUseTrafficCaptureRecQueries(v bool)`

SetUseTrafficCaptureRecQueries sets UseTrafficCaptureRecQueries field to given value.

### HasUseTrafficCaptureRecQueries

`func (o *GetMemberResponse) HasUseTrafficCaptureRecQueries() bool`

HasUseTrafficCaptureRecQueries returns a boolean if a field has been set.

### GetUseTrapNotifications

`func (o *GetMemberResponse) GetUseTrapNotifications() bool`

GetUseTrapNotifications returns the UseTrapNotifications field if non-nil, zero value otherwise.

### GetUseTrapNotificationsOk

`func (o *GetMemberResponse) GetUseTrapNotificationsOk() (*bool, bool)`

GetUseTrapNotificationsOk returns a tuple with the UseTrapNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTrapNotifications

`func (o *GetMemberResponse) SetUseTrapNotifications(v bool)`

SetUseTrapNotifications sets UseTrapNotifications field to given value.

### HasUseTrapNotifications

`func (o *GetMemberResponse) HasUseTrapNotifications() bool`

HasUseTrapNotifications returns a boolean if a field has been set.

### GetUseV4Vrrp

`func (o *GetMemberResponse) GetUseV4Vrrp() bool`

GetUseV4Vrrp returns the UseV4Vrrp field if non-nil, zero value otherwise.

### GetUseV4VrrpOk

`func (o *GetMemberResponse) GetUseV4VrrpOk() (*bool, bool)`

GetUseV4VrrpOk returns a tuple with the UseV4Vrrp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseV4Vrrp

`func (o *GetMemberResponse) SetUseV4Vrrp(v bool)`

SetUseV4Vrrp sets UseV4Vrrp field to given value.

### HasUseV4Vrrp

`func (o *GetMemberResponse) HasUseV4Vrrp() bool`

HasUseV4Vrrp returns a boolean if a field has been set.

### GetVipSetting

`func (o *GetMemberResponse) GetVipSetting() MemberVipSetting`

GetVipSetting returns the VipSetting field if non-nil, zero value otherwise.

### GetVipSettingOk

`func (o *GetMemberResponse) GetVipSettingOk() (*MemberVipSetting, bool)`

GetVipSettingOk returns a tuple with the VipSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSetting

`func (o *GetMemberResponse) SetVipSetting(v MemberVipSetting)`

SetVipSetting sets VipSetting field to given value.

### HasVipSetting

`func (o *GetMemberResponse) HasVipSetting() bool`

HasVipSetting returns a boolean if a field has been set.

### GetVpnMtu

`func (o *GetMemberResponse) GetVpnMtu() int64`

GetVpnMtu returns the VpnMtu field if non-nil, zero value otherwise.

### GetVpnMtuOk

`func (o *GetMemberResponse) GetVpnMtuOk() (*int64, bool)`

GetVpnMtuOk returns a tuple with the VpnMtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnMtu

`func (o *GetMemberResponse) SetVpnMtu(v int64)`

SetVpnMtu sets VpnMtu field to given value.

### HasVpnMtu

`func (o *GetMemberResponse) HasVpnMtu() bool`

HasVpnMtu returns a boolean if a field has been set.

### GetResult

`func (o *GetMemberResponse) GetResult() Member`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMemberResponse) GetResultOk() (*Member, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMemberResponse) SetResult(v Member)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMemberResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


