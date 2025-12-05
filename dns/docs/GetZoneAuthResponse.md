# GetZoneAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The IP address of the server that is serving this zone. | [optional] [readonly] 
**AllowActiveDir** | Pointer to [**[]ZoneAuthAllowActiveDir**](ZoneAuthAllowActiveDir.md) | This field allows the zone to receive GSS-TSIG authenticated DDNS updates from DHCP clients and servers in an AD domain. Note that addresses specified in this field ignore the permission set in the struct which will be set to &#39;ALLOW&#39;. | [optional] 
**AllowFixedRrsetOrder** | Pointer to **bool** | The flag that allows to enable or disable fixed RRset ordering for authoritative forward-mapping zones. | [optional] 
**AllowGssTsigForUnderscoreZone** | Pointer to **bool** | The flag that allows DHCP clients to perform GSS-TSIG signed updates for underscore zones. | [optional] 
**AllowGssTsigZoneUpdates** | Pointer to **bool** | The flag that enables or disables the zone for GSS-TSIG updates. | [optional] 
**AllowQuery** | Pointer to [**[]ZoneAuthAllowQuery**](ZoneAuthAllowQuery.md) | Determines whether DNS queries are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts. | [optional] 
**AllowTransfer** | Pointer to [**[]ZoneAuthAllowTransfer**](ZoneAuthAllowTransfer.md) | Determines whether zone transfers are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts. | [optional] 
**AllowUpdate** | Pointer to [**[]ZoneAuthAllowUpdate**](ZoneAuthAllowUpdate.md) | Determines whether dynamic DNS updates are allowed from a named ACL, or from a list of IPv4/IPv6 addresses, networks, and TSIG keys for the hosts. | [optional] 
**AllowUpdateForwarding** | Pointer to **bool** | The list with IP addresses, networks or TSIG keys for clients, from which forwarded dynamic updates are allowed. | [optional] 
**AwsRte53ZoneInfo** | Pointer to [**ZoneAuthAwsRte53ZoneInfo**](ZoneAuthAwsRte53ZoneInfo.md) |  | [optional] 
**CloudInfo** | Pointer to [**ZoneAuthCloudInfo**](ZoneAuthCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the zone; maximum 256 characters. | [optional] 
**CopyXferToNotify** | Pointer to **bool** | If this flag is set to True then copy allowed IPs from Allow Transfer to Also Notify. | [optional] 
**CreatePtrForBulkHosts** | Pointer to **bool** | Determines if PTR records are created for hosts automatically, if necessary, when the zone data is imported. This field is meaningful only when import_from is set. | [optional] 
**CreatePtrForHosts** | Pointer to **bool** | Determines if PTR records are created for hosts automatically, if necessary, when the zone data is imported. This field is meaningful only when import_from is set. | [optional] 
**CreateUnderscoreZones** | Pointer to **bool** | Determines whether automatic creation of subzones is enabled or not. | [optional] 
**DdnsForceCreationTimestampUpdate** | Pointer to **bool** | Defines whether creation timestamp of RR should be updated &#39; when DDNS update happens even if there is no change to &#39; the RR. | [optional] 
**DdnsPrincipalGroup** | Pointer to **string** | The DDNS Principal cluster group name. | [optional] 
**DdnsPrincipalTracking** | Pointer to **bool** | The flag that indicates whether the DDNS principal track is enabled or disabled. | [optional] 
**DdnsRestrictPatterns** | Pointer to **bool** | The flag that indicates whether an option to restrict DDNS update request based on FQDN patterns is enabled or disabled. | [optional] 
**DdnsRestrictPatternsList** | Pointer to **[]string** | The unordered list of restriction patterns for an option of to restrict DDNS updates based on FQDN patterns. | [optional] 
**DdnsRestrictProtected** | Pointer to **bool** | The flag that indicates whether an option to restrict DDNS update request to protected resource records is enabled or disabled. | [optional] 
**DdnsRestrictSecure** | Pointer to **bool** | The flag that indicates whether DDNS update request for principal other than target resource record&#39;s principal is restricted. | [optional] 
**DdnsRestrictStatic** | Pointer to **bool** | The flag that indicates whether an option to restrict DDNS update request to resource records which are marked as &#39;STATIC&#39; is enabled or disabled. | [optional] 
**Disable** | Pointer to **bool** | Determines whether a zone is disabled or not. When this is set to False, the zone is enabled. | [optional] 
**DisableForwarding** | Pointer to **bool** | Determines whether the name servers that host the zone should forward queries (ended with the domain name of the zone) to any configured forwarders. | [optional] 
**DisplayDomain** | Pointer to **string** | The displayed name of the DNS zone. | [optional] [readonly] 
**DnsFqdn** | Pointer to **string** | The name of this DNS zone in punycode format. For a reverse zone, this is in \&quot;address/cidr\&quot; format. For other zones, this is in FQDN format in punycode format. | [optional] [readonly] 
**DnsIntegrityEnable** | Pointer to **bool** | If this is set to True, DNS integrity check is enabled for this zone. | [optional] 
**DnsIntegrityFrequency** | Pointer to **int64** | The frequency, in seconds, of DNS integrity checks for this zone. | [optional] 
**DnsIntegrityMember** | Pointer to **string** | The Grid member that performs DNS integrity checks for this zone. | [optional] 
**DnsIntegrityVerboseLogging** | Pointer to **bool** | If this is set to True, more information is logged for DNS integrity checks for this zone. | [optional] 
**DnsSoaEmail** | Pointer to **string** | The SOA email for the zone in punycode format. | [optional] [readonly] 
**DnssecKeyParams** | Pointer to [**ZoneAuthDnssecKeyParams**](ZoneAuthDnssecKeyParams.md) |  | [optional] 
**DnssecKeys** | Pointer to [**[]ZoneAuthDnssecKeys**](ZoneAuthDnssecKeys.md) | A list of DNSSEC keys for the zone. | [optional] 
**DnssecKskRolloverDate** | Pointer to **int64** | The rollover date for the Key Signing Key. | [optional] [readonly] 
**DnssecZskRolloverDate** | Pointer to **int64** | The rollover date for the Zone Signing Key. | [optional] [readonly] 
**DoHostAbstraction** | Pointer to **bool** | Determines if hosts and bulk hosts are automatically created when the zone data is imported. This field is meaningful only when import_from is set. | [optional] 
**EffectiveCheckNamesPolicy** | Pointer to **string** | The value of the check names policy, which indicates the action the appliance takes when it encounters host names that do not comply with the Strict Hostname Checking policy. This value applies only if the host name restriction policy is set to \&quot;Strict Hostname Checking\&quot;. | [optional] 
**EffectiveRecordNamePolicy** | Pointer to **string** | The selected hostname policy for records under this zone. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExternalPrimaries** | Pointer to [**[]ZoneAuthExternalPrimaries**](ZoneAuthExternalPrimaries.md) | The list of external primary servers. | [optional] 
**ExternalSecondaries** | Pointer to [**[]ZoneAuthExternalSecondaries**](ZoneAuthExternalSecondaries.md) | The list of external secondary servers. | [optional] 
**Fqdn** | Pointer to **string** | The name of this DNS zone. For a reverse zone, this is in \&quot;address/cidr\&quot; format. For other zones, this is in FQDN format. This value can be in unicode format. Note that for a reverse zone, the corresponding zone_format value should be set. | [optional] 
**GridPrimary** | Pointer to [**[]ZoneAuthGridPrimary**](ZoneAuthGridPrimary.md) | The grid primary servers for this zone. | [optional] 
**GridPrimarySharedWithMsParentDelegation** | Pointer to **bool** | Determines if the server is duplicated with parent delegation. | [optional] [readonly] 
**GridSecondaries** | Pointer to [**[]ZoneAuthGridSecondaries**](ZoneAuthGridSecondaries.md) | The list with Grid members that are secondary servers for this zone. | [optional] 
**ImportFrom** | Pointer to **string** | The IP address of the Infoblox appliance from which zone data is imported. Setting this address to &#39;255.255.255.255&#39; and do_host_abstraction to &#39;true&#39; will create Host records from A records in this zone without importing zone data. | [optional] 
**IsDnssecEnabled** | Pointer to **bool** | This flag is set to True if DNSSEC is enabled for the zone. | [optional] [readonly] 
**IsDnssecSigned** | Pointer to **bool** | Determines if the zone is DNSSEC signed. | [optional] [readonly] 
**IsMultimaster** | Pointer to **bool** | Determines if multi-master DNS is enabled for the zone. | [optional] [readonly] 
**LastQueried** | Pointer to **int64** | The time the zone was last queried on. | [optional] [readonly] 
**LastQueriedAcl** | Pointer to [**[]ZoneAuthLastQueriedAcl**](ZoneAuthLastQueriedAcl.md) | Determines last queried ACL for the specified IPv4 or IPv6 addresses and networks in scavenging settings. | [optional] 
**Locked** | Pointer to **bool** | If you enable this flag, other administrators cannot make conflicting changes. This is for administration purposes only. The zone will continue to serve DNS data even when it is locked. | [optional] 
**LockedBy** | Pointer to **string** | The name of a superuser or the administrator who locked this zone. | [optional] [readonly] 
**MaskPrefix** | Pointer to **string** | IPv4 Netmask or IPv6 prefix for this zone. | [optional] [readonly] 
**MemberSoaMnames** | Pointer to [**[]ZoneAuthMemberSoaMnames**](ZoneAuthMemberSoaMnames.md) | The list of per-member SOA MNAME information. | [optional] 
**MemberSoaSerials** | Pointer to [**[]ZoneAuthMemberSoaSerials**](ZoneAuthMemberSoaSerials.md) | The list of per-member SOA serial information. | [optional] [readonly] 
**MgmPrivate** | Pointer to **bool** | This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized. | [optional] 
**MgmPrivateOverridable** | Pointer to **bool** | This field is assumed to be True unless filled by any conforming objects, such as Network, IPv6 Network, Network Container, IPv6 Network Container, and Network View. This value is set to False if mgm_private is set to True in the parent object. | [optional] [readonly] 
**MsAdIntegrated** | Pointer to **bool** | The flag that determines whether Active Directory is integrated or not. This field is valid only when ms_managed is \&quot;STUB\&quot;, \&quot;AUTH_PRIMARY\&quot;, or \&quot;AUTH_BOTH\&quot;. | [optional] 
**MsAllowTransfer** | Pointer to [**[]ZoneAuthMsAllowTransfer**](ZoneAuthMsAllowTransfer.md) | The list of DNS clients that are allowed to perform zone transfers from a Microsoft DNS server. This setting applies only to zones with Microsoft DNS servers that are either primary or secondary servers. This setting does not inherit any value from the Grid or from any member that defines an allow_transfer value. This setting does not apply to any grid member. Use the allow_transfer field to control which DNS clients are allowed to perform zone transfers on Grid members. | [optional] 
**MsAllowTransferMode** | Pointer to **string** | Determines which DNS clients are allowed to perform zone transfers from a Microsoft DNS server. Valid values are: \&quot;ADDRESS_AC\&quot;, to use ms_allow_transfer field for specifying IP addresses, networks and Transaction Signature (TSIG) keys for clients that are allowed to do zone transfers. \&quot;ANY\&quot;, to allow any client. \&quot;ANY_NS\&quot;, to allow only the nameservers listed in this zone. \&quot;NONE\&quot;, to deny all zone transfer requests. | [optional] 
**MsDcNsRecordCreation** | Pointer to [**[]ZoneAuthMsDcNsRecordCreation**](ZoneAuthMsDcNsRecordCreation.md) | The list of domain controllers that are allowed to create NS records for authoritative zones. | [optional] 
**MsDdnsMode** | Pointer to **string** | Determines whether an Active Directory-integrated zone with a Microsoft DNS server as primary allows dynamic updates. Valid values are: \&quot;SECURE\&quot; if the zone allows secure updates only. \&quot;NONE\&quot; if the zone forbids dynamic updates. \&quot;ANY\&quot; if the zone accepts both secure and nonsecure updates. This field is valid only if ms_managed is either \&quot;AUTH_PRIMARY\&quot; or \&quot;AUTH_BOTH\&quot;. If the flag ms_ad_integrated is false, the value \&quot;SECURE\&quot; is not allowed. | [optional] 
**MsManaged** | Pointer to **string** | The flag that indicates whether the zone is assigned to a Microsoft DNS server. This flag returns the authoritative name server type of the Microsoft DNS server. Valid values are: \&quot;NONE\&quot; if the zone is not assigned to any Microsoft DNS server. \&quot;STUB\&quot; if the zone is assigned to a Microsoft DNS server as a stub zone. \&quot;AUTH_PRIMARY\&quot; if only the primary server of the zone is a Microsoft DNS server. \&quot;AUTH_SECONDARY\&quot; if only the secondary server of the zone is a Microsoft DNS server. \&quot;AUTH_BOTH\&quot; if both the primary and secondary servers of the zone are Microsoft DNS servers. | [optional] [readonly] 
**MsPrimaries** | Pointer to [**[]ZoneAuthMsPrimaries**](ZoneAuthMsPrimaries.md) | The list with the Microsoft DNS servers that are primary servers for the zone. Although a zone typically has just one primary name server, you can specify up to ten independent servers for a single zone. | [optional] 
**MsReadOnly** | Pointer to **bool** | Determines if a Grid member manages the zone served by a Microsoft DNS server in read-only mode. This flag is true when a Grid member manages the zone in read-only mode, false otherwise. When the zone has the ms_read_only flag set to True, no changes can be made to this zone. | [optional] [readonly] 
**MsSecondaries** | Pointer to [**[]ZoneAuthMsSecondaries**](ZoneAuthMsSecondaries.md) | The list with the Microsoft DNS servers that are secondary servers for the zone. | [optional] 
**MsSyncDisabled** | Pointer to **bool** | This flag controls whether this zone is synchronized with Microsoft DNS servers. | [optional] 
**MsSyncMasterName** | Pointer to **string** | The name of MS synchronization master for this zone. | [optional] [readonly] 
**NetworkAssociations** | Pointer to **[]string** | The list with the associated network/network container information. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this zone resides. | [optional] [readonly] 
**NotifyDelay** | Pointer to **int64** | The number of seconds in delay with which notify messages are sent to secondaries. | [optional] 
**NsGroup** | Pointer to **string** | The name server group that serves DNS for this zone. | [optional] 
**Parent** | Pointer to **string** | The parent zone of this zone. Note that when searching for reverse zones, the \&quot;in-addr.arpa\&quot; notation should be used. | [optional] [readonly] 
**Prefix** | Pointer to **string** | The RFC2317 prefix value of this DNS zone. Use this field only when the netmask is greater than 24 bits; that is, for a mask between 25 and 31 bits. Enter a prefix, such as the name of the allocated address block. The prefix can be alphanumeric characters, such as 128/26 , 128-189 , or sub-B. | [optional] 
**PrimaryType** | Pointer to **string** | The type of the primary server. | [optional] [readonly] 
**RecordNamePolicy** | Pointer to **string** | The hostname policy for records under this zone. | [optional] 
**RecordsMonitored** | Pointer to **bool** | Determines if this zone is also monitoring resource records. | [optional] [readonly] 
**RemoveSubzones** | Pointer to **bool** | Remove subzones delete option. Determines whether all child objects should be removed alongside with the parent zone or child objects should be assigned to another parental zone. By default child objects are deleted with the parent zone. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. | [optional] 
**RrNotQueriedEnabledTime** | Pointer to **int64** | The time data collection for Not Queried Resource Record was enabled for this zone. | [optional] [readonly] 
**ScavengingSettings** | Pointer to [**ZoneAuthScavengingSettings**](ZoneAuthScavengingSettings.md) |  | [optional] 
**SetSoaSerialNumber** | Pointer to **bool** | The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \&quot;soa_serial_number\&quot; and pass \&quot;set_soa_serial_number\&quot; as True. | [optional] 
**SoaDefaultTtl** | Pointer to **int64** | The Time to Live (TTL) value of the SOA record of this zone. This value is the number of seconds that data is cached. | [optional] 
**SoaEmail** | Pointer to **string** | The SOA email value for this zone. This value can be in unicode format. | [optional] 
**SoaExpire** | Pointer to **int64** | This setting defines the amount of time, in seconds, after which the secondary server stops giving out answers about the zone because the zone data is too old to be useful. The default is one week. | [optional] 
**SoaNegativeTtl** | Pointer to **int64** | The negative Time to Live (TTL) value of the SOA of the zone indicates how long a secondary server can cache data for \&quot;Does Not Respond\&quot; responses. | [optional] 
**SoaRefresh** | Pointer to **int64** | This indicates the interval at which a secondary server sends a message to the primary server for a zone to check that its data is current, and retrieve fresh data if it is not. | [optional] 
**SoaRetry** | Pointer to **int64** | This indicates how long a secondary server must wait before attempting to recontact the primary server after a connection failure between the two servers occurs. | [optional] 
**SoaSerial** | Pointer to **int64** | The serial number in the SOA record incrementally changes every time the record is modified. The Infoblox appliance allows you to change the serial number (in the SOA record) for the primary server so it is higher than the secondary server, thereby ensuring zone transfers come from the primary server (as they should). To change the serial number you need to set a new value at \&quot;soa_serial_number\&quot; and pass \&quot;set_soa_serial_number\&quot; as True. | [optional] 
**Srgs** | Pointer to **[]string** | The associated shared record groups of a DNS zone. If a shared record group is associated with a zone, then all shared records in a shared record group will be shared in the zone. | [optional] 
**UpdateForwarding** | Pointer to [**[]ZoneAuthUpdateForwarding**](ZoneAuthUpdateForwarding.md) | Use this field to allow or deny dynamic DNS updates that are forwarded from specific IPv4/IPv6 addresses, networks, or a named ACL. You can also provide TSIG keys for clients that are allowed or denied to perform zone updates. This setting overrides the member-level setting. | [optional] 
**UseAllowActiveDir** | Pointer to **bool** | Use flag for: allow_active_dir | [optional] 
**UseAllowQuery** | Pointer to **bool** | Use flag for: allow_query | [optional] 
**UseAllowTransfer** | Pointer to **bool** | Use flag for: allow_transfer | [optional] 
**UseAllowUpdate** | Pointer to **bool** | Use flag for: allow_update | [optional] 
**UseAllowUpdateForwarding** | Pointer to **bool** | Use flag for: allow_update_forwarding | [optional] 
**UseCheckNamesPolicy** | Pointer to **bool** | Apply policy to dynamic updates and inbound zone transfers (This value applies only if the host name restriction policy is set to \&quot;Strict Hostname Checking\&quot;.) | [optional] 
**UseCopyXferToNotify** | Pointer to **bool** | Use flag for: copy_xfer_to_notify | [optional] 
**UseDdnsForceCreationTimestampUpdate** | Pointer to **bool** | Use flag for: ddns_force_creation_timestamp_update | [optional] 
**UseDdnsPatternsRestriction** | Pointer to **bool** | Use flag for: ddns_restrict_patterns_list , ddns_restrict_patterns | [optional] 
**UseDdnsPrincipalSecurity** | Pointer to **bool** | Use flag for: ddns_restrict_secure , ddns_principal_tracking, ddns_principal_group | [optional] 
**UseDdnsRestrictProtected** | Pointer to **bool** | Use flag for: ddns_restrict_protected | [optional] 
**UseDdnsRestrictStatic** | Pointer to **bool** | Use flag for: ddns_restrict_static | [optional] 
**UseDnssecKeyParams** | Pointer to **bool** | Use flag for: dnssec_key_params | [optional] 
**UseExternalPrimary** | Pointer to **bool** | This flag controls whether the zone is using an external primary. | [optional] 
**UseGridZoneTimer** | Pointer to **bool** | Use flag for: soa_default_ttl , soa_expire, soa_negative_ttl, soa_refresh, soa_retry | [optional] 
**UseImportFrom** | Pointer to **bool** | Use flag for: import_from | [optional] 
**UseNotifyDelay** | Pointer to **bool** | Use flag for: notify_delay | [optional] 
**UseRecordNamePolicy** | Pointer to **bool** | Use flag for: record_name_policy | [optional] 
**UseScavengingSettings** | Pointer to **bool** | Use flag for: scavenging_settings , last_queried_acl | [optional] 
**UseSoaEmail** | Pointer to **bool** | Use flag for: soa_email | [optional] 
**UsingSrgAssociations** | Pointer to **bool** | This is true if the zone is associated with a shared record group. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**View** | Pointer to **string** | The name of the DNS view in which the zone resides. Example \&quot;external\&quot;. | [optional] 
**ZoneFormat** | Pointer to **string** | Determines the format of this zone. | [optional] 
**ZoneNotQueriedEnabledTime** | Pointer to **int64** | The time when \&quot;DNS Zones Last Queried\&quot; was turned on for this zone. | [optional] [readonly] 
**Result** | Pointer to [**ZoneAuth**](ZoneAuth.md) |  | [optional] 

## Methods

### NewGetZoneAuthResponse

`func NewGetZoneAuthResponse() *GetZoneAuthResponse`

NewGetZoneAuthResponse instantiates a new GetZoneAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneAuthResponseWithDefaults

`func NewGetZoneAuthResponseWithDefaults() *GetZoneAuthResponse`

NewGetZoneAuthResponseWithDefaults instantiates a new GetZoneAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetZoneAuthResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetZoneAuthResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetZoneAuthResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetZoneAuthResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetZoneAuthResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetZoneAuthResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetZoneAuthResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetZoneAuthResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAllowActiveDir

`func (o *GetZoneAuthResponse) GetAllowActiveDir() []ZoneAuthAllowActiveDir`

GetAllowActiveDir returns the AllowActiveDir field if non-nil, zero value otherwise.

### GetAllowActiveDirOk

`func (o *GetZoneAuthResponse) GetAllowActiveDirOk() (*[]ZoneAuthAllowActiveDir, bool)`

GetAllowActiveDirOk returns a tuple with the AllowActiveDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowActiveDir

`func (o *GetZoneAuthResponse) SetAllowActiveDir(v []ZoneAuthAllowActiveDir)`

SetAllowActiveDir sets AllowActiveDir field to given value.

### HasAllowActiveDir

`func (o *GetZoneAuthResponse) HasAllowActiveDir() bool`

HasAllowActiveDir returns a boolean if a field has been set.

### GetAllowFixedRrsetOrder

`func (o *GetZoneAuthResponse) GetAllowFixedRrsetOrder() bool`

GetAllowFixedRrsetOrder returns the AllowFixedRrsetOrder field if non-nil, zero value otherwise.

### GetAllowFixedRrsetOrderOk

`func (o *GetZoneAuthResponse) GetAllowFixedRrsetOrderOk() (*bool, bool)`

GetAllowFixedRrsetOrderOk returns a tuple with the AllowFixedRrsetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFixedRrsetOrder

`func (o *GetZoneAuthResponse) SetAllowFixedRrsetOrder(v bool)`

SetAllowFixedRrsetOrder sets AllowFixedRrsetOrder field to given value.

### HasAllowFixedRrsetOrder

`func (o *GetZoneAuthResponse) HasAllowFixedRrsetOrder() bool`

HasAllowFixedRrsetOrder returns a boolean if a field has been set.

### GetAllowGssTsigForUnderscoreZone

`func (o *GetZoneAuthResponse) GetAllowGssTsigForUnderscoreZone() bool`

GetAllowGssTsigForUnderscoreZone returns the AllowGssTsigForUnderscoreZone field if non-nil, zero value otherwise.

### GetAllowGssTsigForUnderscoreZoneOk

`func (o *GetZoneAuthResponse) GetAllowGssTsigForUnderscoreZoneOk() (*bool, bool)`

GetAllowGssTsigForUnderscoreZoneOk returns a tuple with the AllowGssTsigForUnderscoreZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGssTsigForUnderscoreZone

`func (o *GetZoneAuthResponse) SetAllowGssTsigForUnderscoreZone(v bool)`

SetAllowGssTsigForUnderscoreZone sets AllowGssTsigForUnderscoreZone field to given value.

### HasAllowGssTsigForUnderscoreZone

`func (o *GetZoneAuthResponse) HasAllowGssTsigForUnderscoreZone() bool`

HasAllowGssTsigForUnderscoreZone returns a boolean if a field has been set.

### GetAllowGssTsigZoneUpdates

`func (o *GetZoneAuthResponse) GetAllowGssTsigZoneUpdates() bool`

GetAllowGssTsigZoneUpdates returns the AllowGssTsigZoneUpdates field if non-nil, zero value otherwise.

### GetAllowGssTsigZoneUpdatesOk

`func (o *GetZoneAuthResponse) GetAllowGssTsigZoneUpdatesOk() (*bool, bool)`

GetAllowGssTsigZoneUpdatesOk returns a tuple with the AllowGssTsigZoneUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowGssTsigZoneUpdates

`func (o *GetZoneAuthResponse) SetAllowGssTsigZoneUpdates(v bool)`

SetAllowGssTsigZoneUpdates sets AllowGssTsigZoneUpdates field to given value.

### HasAllowGssTsigZoneUpdates

`func (o *GetZoneAuthResponse) HasAllowGssTsigZoneUpdates() bool`

HasAllowGssTsigZoneUpdates returns a boolean if a field has been set.

### GetAllowQuery

`func (o *GetZoneAuthResponse) GetAllowQuery() []ZoneAuthAllowQuery`

GetAllowQuery returns the AllowQuery field if non-nil, zero value otherwise.

### GetAllowQueryOk

`func (o *GetZoneAuthResponse) GetAllowQueryOk() (*[]ZoneAuthAllowQuery, bool)`

GetAllowQueryOk returns a tuple with the AllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuery

`func (o *GetZoneAuthResponse) SetAllowQuery(v []ZoneAuthAllowQuery)`

SetAllowQuery sets AllowQuery field to given value.

### HasAllowQuery

`func (o *GetZoneAuthResponse) HasAllowQuery() bool`

HasAllowQuery returns a boolean if a field has been set.

### GetAllowTransfer

`func (o *GetZoneAuthResponse) GetAllowTransfer() []ZoneAuthAllowTransfer`

GetAllowTransfer returns the AllowTransfer field if non-nil, zero value otherwise.

### GetAllowTransferOk

`func (o *GetZoneAuthResponse) GetAllowTransferOk() (*[]ZoneAuthAllowTransfer, bool)`

GetAllowTransferOk returns a tuple with the AllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTransfer

`func (o *GetZoneAuthResponse) SetAllowTransfer(v []ZoneAuthAllowTransfer)`

SetAllowTransfer sets AllowTransfer field to given value.

### HasAllowTransfer

`func (o *GetZoneAuthResponse) HasAllowTransfer() bool`

HasAllowTransfer returns a boolean if a field has been set.

### GetAllowUpdate

`func (o *GetZoneAuthResponse) GetAllowUpdate() []ZoneAuthAllowUpdate`

GetAllowUpdate returns the AllowUpdate field if non-nil, zero value otherwise.

### GetAllowUpdateOk

`func (o *GetZoneAuthResponse) GetAllowUpdateOk() (*[]ZoneAuthAllowUpdate, bool)`

GetAllowUpdateOk returns a tuple with the AllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdate

`func (o *GetZoneAuthResponse) SetAllowUpdate(v []ZoneAuthAllowUpdate)`

SetAllowUpdate sets AllowUpdate field to given value.

### HasAllowUpdate

`func (o *GetZoneAuthResponse) HasAllowUpdate() bool`

HasAllowUpdate returns a boolean if a field has been set.

### GetAllowUpdateForwarding

`func (o *GetZoneAuthResponse) GetAllowUpdateForwarding() bool`

GetAllowUpdateForwarding returns the AllowUpdateForwarding field if non-nil, zero value otherwise.

### GetAllowUpdateForwardingOk

`func (o *GetZoneAuthResponse) GetAllowUpdateForwardingOk() (*bool, bool)`

GetAllowUpdateForwardingOk returns a tuple with the AllowUpdateForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUpdateForwarding

`func (o *GetZoneAuthResponse) SetAllowUpdateForwarding(v bool)`

SetAllowUpdateForwarding sets AllowUpdateForwarding field to given value.

### HasAllowUpdateForwarding

`func (o *GetZoneAuthResponse) HasAllowUpdateForwarding() bool`

HasAllowUpdateForwarding returns a boolean if a field has been set.

### GetAwsRte53ZoneInfo

`func (o *GetZoneAuthResponse) GetAwsRte53ZoneInfo() ZoneAuthAwsRte53ZoneInfo`

GetAwsRte53ZoneInfo returns the AwsRte53ZoneInfo field if non-nil, zero value otherwise.

### GetAwsRte53ZoneInfoOk

`func (o *GetZoneAuthResponse) GetAwsRte53ZoneInfoOk() (*ZoneAuthAwsRte53ZoneInfo, bool)`

GetAwsRte53ZoneInfoOk returns a tuple with the AwsRte53ZoneInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRte53ZoneInfo

`func (o *GetZoneAuthResponse) SetAwsRte53ZoneInfo(v ZoneAuthAwsRte53ZoneInfo)`

SetAwsRte53ZoneInfo sets AwsRte53ZoneInfo field to given value.

### HasAwsRte53ZoneInfo

`func (o *GetZoneAuthResponse) HasAwsRte53ZoneInfo() bool`

HasAwsRte53ZoneInfo returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetZoneAuthResponse) GetCloudInfo() ZoneAuthCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetZoneAuthResponse) GetCloudInfoOk() (*ZoneAuthCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetZoneAuthResponse) SetCloudInfo(v ZoneAuthCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetZoneAuthResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetZoneAuthResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetZoneAuthResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetZoneAuthResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetZoneAuthResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCopyXferToNotify

`func (o *GetZoneAuthResponse) GetCopyXferToNotify() bool`

GetCopyXferToNotify returns the CopyXferToNotify field if non-nil, zero value otherwise.

### GetCopyXferToNotifyOk

`func (o *GetZoneAuthResponse) GetCopyXferToNotifyOk() (*bool, bool)`

GetCopyXferToNotifyOk returns a tuple with the CopyXferToNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyXferToNotify

`func (o *GetZoneAuthResponse) SetCopyXferToNotify(v bool)`

SetCopyXferToNotify sets CopyXferToNotify field to given value.

### HasCopyXferToNotify

`func (o *GetZoneAuthResponse) HasCopyXferToNotify() bool`

HasCopyXferToNotify returns a boolean if a field has been set.

### GetCreatePtrForBulkHosts

`func (o *GetZoneAuthResponse) GetCreatePtrForBulkHosts() bool`

GetCreatePtrForBulkHosts returns the CreatePtrForBulkHosts field if non-nil, zero value otherwise.

### GetCreatePtrForBulkHostsOk

`func (o *GetZoneAuthResponse) GetCreatePtrForBulkHostsOk() (*bool, bool)`

GetCreatePtrForBulkHostsOk returns a tuple with the CreatePtrForBulkHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePtrForBulkHosts

`func (o *GetZoneAuthResponse) SetCreatePtrForBulkHosts(v bool)`

SetCreatePtrForBulkHosts sets CreatePtrForBulkHosts field to given value.

### HasCreatePtrForBulkHosts

`func (o *GetZoneAuthResponse) HasCreatePtrForBulkHosts() bool`

HasCreatePtrForBulkHosts returns a boolean if a field has been set.

### GetCreatePtrForHosts

`func (o *GetZoneAuthResponse) GetCreatePtrForHosts() bool`

GetCreatePtrForHosts returns the CreatePtrForHosts field if non-nil, zero value otherwise.

### GetCreatePtrForHostsOk

`func (o *GetZoneAuthResponse) GetCreatePtrForHostsOk() (*bool, bool)`

GetCreatePtrForHostsOk returns a tuple with the CreatePtrForHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatePtrForHosts

`func (o *GetZoneAuthResponse) SetCreatePtrForHosts(v bool)`

SetCreatePtrForHosts sets CreatePtrForHosts field to given value.

### HasCreatePtrForHosts

`func (o *GetZoneAuthResponse) HasCreatePtrForHosts() bool`

HasCreatePtrForHosts returns a boolean if a field has been set.

### GetCreateUnderscoreZones

`func (o *GetZoneAuthResponse) GetCreateUnderscoreZones() bool`

GetCreateUnderscoreZones returns the CreateUnderscoreZones field if non-nil, zero value otherwise.

### GetCreateUnderscoreZonesOk

`func (o *GetZoneAuthResponse) GetCreateUnderscoreZonesOk() (*bool, bool)`

GetCreateUnderscoreZonesOk returns a tuple with the CreateUnderscoreZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUnderscoreZones

`func (o *GetZoneAuthResponse) SetCreateUnderscoreZones(v bool)`

SetCreateUnderscoreZones sets CreateUnderscoreZones field to given value.

### HasCreateUnderscoreZones

`func (o *GetZoneAuthResponse) HasCreateUnderscoreZones() bool`

HasCreateUnderscoreZones returns a boolean if a field has been set.

### GetDdnsForceCreationTimestampUpdate

`func (o *GetZoneAuthResponse) GetDdnsForceCreationTimestampUpdate() bool`

GetDdnsForceCreationTimestampUpdate returns the DdnsForceCreationTimestampUpdate field if non-nil, zero value otherwise.

### GetDdnsForceCreationTimestampUpdateOk

`func (o *GetZoneAuthResponse) GetDdnsForceCreationTimestampUpdateOk() (*bool, bool)`

GetDdnsForceCreationTimestampUpdateOk returns a tuple with the DdnsForceCreationTimestampUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsForceCreationTimestampUpdate

`func (o *GetZoneAuthResponse) SetDdnsForceCreationTimestampUpdate(v bool)`

SetDdnsForceCreationTimestampUpdate sets DdnsForceCreationTimestampUpdate field to given value.

### HasDdnsForceCreationTimestampUpdate

`func (o *GetZoneAuthResponse) HasDdnsForceCreationTimestampUpdate() bool`

HasDdnsForceCreationTimestampUpdate returns a boolean if a field has been set.

### GetDdnsPrincipalGroup

`func (o *GetZoneAuthResponse) GetDdnsPrincipalGroup() string`

GetDdnsPrincipalGroup returns the DdnsPrincipalGroup field if non-nil, zero value otherwise.

### GetDdnsPrincipalGroupOk

`func (o *GetZoneAuthResponse) GetDdnsPrincipalGroupOk() (*string, bool)`

GetDdnsPrincipalGroupOk returns a tuple with the DdnsPrincipalGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalGroup

`func (o *GetZoneAuthResponse) SetDdnsPrincipalGroup(v string)`

SetDdnsPrincipalGroup sets DdnsPrincipalGroup field to given value.

### HasDdnsPrincipalGroup

`func (o *GetZoneAuthResponse) HasDdnsPrincipalGroup() bool`

HasDdnsPrincipalGroup returns a boolean if a field has been set.

### GetDdnsPrincipalTracking

`func (o *GetZoneAuthResponse) GetDdnsPrincipalTracking() bool`

GetDdnsPrincipalTracking returns the DdnsPrincipalTracking field if non-nil, zero value otherwise.

### GetDdnsPrincipalTrackingOk

`func (o *GetZoneAuthResponse) GetDdnsPrincipalTrackingOk() (*bool, bool)`

GetDdnsPrincipalTrackingOk returns a tuple with the DdnsPrincipalTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsPrincipalTracking

`func (o *GetZoneAuthResponse) SetDdnsPrincipalTracking(v bool)`

SetDdnsPrincipalTracking sets DdnsPrincipalTracking field to given value.

### HasDdnsPrincipalTracking

`func (o *GetZoneAuthResponse) HasDdnsPrincipalTracking() bool`

HasDdnsPrincipalTracking returns a boolean if a field has been set.

### GetDdnsRestrictPatterns

`func (o *GetZoneAuthResponse) GetDdnsRestrictPatterns() bool`

GetDdnsRestrictPatterns returns the DdnsRestrictPatterns field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsOk

`func (o *GetZoneAuthResponse) GetDdnsRestrictPatternsOk() (*bool, bool)`

GetDdnsRestrictPatternsOk returns a tuple with the DdnsRestrictPatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatterns

`func (o *GetZoneAuthResponse) SetDdnsRestrictPatterns(v bool)`

SetDdnsRestrictPatterns sets DdnsRestrictPatterns field to given value.

### HasDdnsRestrictPatterns

`func (o *GetZoneAuthResponse) HasDdnsRestrictPatterns() bool`

HasDdnsRestrictPatterns returns a boolean if a field has been set.

### GetDdnsRestrictPatternsList

`func (o *GetZoneAuthResponse) GetDdnsRestrictPatternsList() []string`

GetDdnsRestrictPatternsList returns the DdnsRestrictPatternsList field if non-nil, zero value otherwise.

### GetDdnsRestrictPatternsListOk

`func (o *GetZoneAuthResponse) GetDdnsRestrictPatternsListOk() (*[]string, bool)`

GetDdnsRestrictPatternsListOk returns a tuple with the DdnsRestrictPatternsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictPatternsList

`func (o *GetZoneAuthResponse) SetDdnsRestrictPatternsList(v []string)`

SetDdnsRestrictPatternsList sets DdnsRestrictPatternsList field to given value.

### HasDdnsRestrictPatternsList

`func (o *GetZoneAuthResponse) HasDdnsRestrictPatternsList() bool`

HasDdnsRestrictPatternsList returns a boolean if a field has been set.

### GetDdnsRestrictProtected

`func (o *GetZoneAuthResponse) GetDdnsRestrictProtected() bool`

GetDdnsRestrictProtected returns the DdnsRestrictProtected field if non-nil, zero value otherwise.

### GetDdnsRestrictProtectedOk

`func (o *GetZoneAuthResponse) GetDdnsRestrictProtectedOk() (*bool, bool)`

GetDdnsRestrictProtectedOk returns a tuple with the DdnsRestrictProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictProtected

`func (o *GetZoneAuthResponse) SetDdnsRestrictProtected(v bool)`

SetDdnsRestrictProtected sets DdnsRestrictProtected field to given value.

### HasDdnsRestrictProtected

`func (o *GetZoneAuthResponse) HasDdnsRestrictProtected() bool`

HasDdnsRestrictProtected returns a boolean if a field has been set.

### GetDdnsRestrictSecure

`func (o *GetZoneAuthResponse) GetDdnsRestrictSecure() bool`

GetDdnsRestrictSecure returns the DdnsRestrictSecure field if non-nil, zero value otherwise.

### GetDdnsRestrictSecureOk

`func (o *GetZoneAuthResponse) GetDdnsRestrictSecureOk() (*bool, bool)`

GetDdnsRestrictSecureOk returns a tuple with the DdnsRestrictSecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictSecure

`func (o *GetZoneAuthResponse) SetDdnsRestrictSecure(v bool)`

SetDdnsRestrictSecure sets DdnsRestrictSecure field to given value.

### HasDdnsRestrictSecure

`func (o *GetZoneAuthResponse) HasDdnsRestrictSecure() bool`

HasDdnsRestrictSecure returns a boolean if a field has been set.

### GetDdnsRestrictStatic

`func (o *GetZoneAuthResponse) GetDdnsRestrictStatic() bool`

GetDdnsRestrictStatic returns the DdnsRestrictStatic field if non-nil, zero value otherwise.

### GetDdnsRestrictStaticOk

`func (o *GetZoneAuthResponse) GetDdnsRestrictStaticOk() (*bool, bool)`

GetDdnsRestrictStaticOk returns a tuple with the DdnsRestrictStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsRestrictStatic

`func (o *GetZoneAuthResponse) SetDdnsRestrictStatic(v bool)`

SetDdnsRestrictStatic sets DdnsRestrictStatic field to given value.

### HasDdnsRestrictStatic

`func (o *GetZoneAuthResponse) HasDdnsRestrictStatic() bool`

HasDdnsRestrictStatic returns a boolean if a field has been set.

### GetDisable

`func (o *GetZoneAuthResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetZoneAuthResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetZoneAuthResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetZoneAuthResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisableForwarding

`func (o *GetZoneAuthResponse) GetDisableForwarding() bool`

GetDisableForwarding returns the DisableForwarding field if non-nil, zero value otherwise.

### GetDisableForwardingOk

`func (o *GetZoneAuthResponse) GetDisableForwardingOk() (*bool, bool)`

GetDisableForwardingOk returns a tuple with the DisableForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableForwarding

`func (o *GetZoneAuthResponse) SetDisableForwarding(v bool)`

SetDisableForwarding sets DisableForwarding field to given value.

### HasDisableForwarding

`func (o *GetZoneAuthResponse) HasDisableForwarding() bool`

HasDisableForwarding returns a boolean if a field has been set.

### GetDisplayDomain

`func (o *GetZoneAuthResponse) GetDisplayDomain() string`

GetDisplayDomain returns the DisplayDomain field if non-nil, zero value otherwise.

### GetDisplayDomainOk

`func (o *GetZoneAuthResponse) GetDisplayDomainOk() (*string, bool)`

GetDisplayDomainOk returns a tuple with the DisplayDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDomain

`func (o *GetZoneAuthResponse) SetDisplayDomain(v string)`

SetDisplayDomain sets DisplayDomain field to given value.

### HasDisplayDomain

`func (o *GetZoneAuthResponse) HasDisplayDomain() bool`

HasDisplayDomain returns a boolean if a field has been set.

### GetDnsFqdn

`func (o *GetZoneAuthResponse) GetDnsFqdn() string`

GetDnsFqdn returns the DnsFqdn field if non-nil, zero value otherwise.

### GetDnsFqdnOk

`func (o *GetZoneAuthResponse) GetDnsFqdnOk() (*string, bool)`

GetDnsFqdnOk returns a tuple with the DnsFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsFqdn

`func (o *GetZoneAuthResponse) SetDnsFqdn(v string)`

SetDnsFqdn sets DnsFqdn field to given value.

### HasDnsFqdn

`func (o *GetZoneAuthResponse) HasDnsFqdn() bool`

HasDnsFqdn returns a boolean if a field has been set.

### GetDnsIntegrityEnable

`func (o *GetZoneAuthResponse) GetDnsIntegrityEnable() bool`

GetDnsIntegrityEnable returns the DnsIntegrityEnable field if non-nil, zero value otherwise.

### GetDnsIntegrityEnableOk

`func (o *GetZoneAuthResponse) GetDnsIntegrityEnableOk() (*bool, bool)`

GetDnsIntegrityEnableOk returns a tuple with the DnsIntegrityEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrityEnable

`func (o *GetZoneAuthResponse) SetDnsIntegrityEnable(v bool)`

SetDnsIntegrityEnable sets DnsIntegrityEnable field to given value.

### HasDnsIntegrityEnable

`func (o *GetZoneAuthResponse) HasDnsIntegrityEnable() bool`

HasDnsIntegrityEnable returns a boolean if a field has been set.

### GetDnsIntegrityFrequency

`func (o *GetZoneAuthResponse) GetDnsIntegrityFrequency() int64`

GetDnsIntegrityFrequency returns the DnsIntegrityFrequency field if non-nil, zero value otherwise.

### GetDnsIntegrityFrequencyOk

`func (o *GetZoneAuthResponse) GetDnsIntegrityFrequencyOk() (*int64, bool)`

GetDnsIntegrityFrequencyOk returns a tuple with the DnsIntegrityFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrityFrequency

`func (o *GetZoneAuthResponse) SetDnsIntegrityFrequency(v int64)`

SetDnsIntegrityFrequency sets DnsIntegrityFrequency field to given value.

### HasDnsIntegrityFrequency

`func (o *GetZoneAuthResponse) HasDnsIntegrityFrequency() bool`

HasDnsIntegrityFrequency returns a boolean if a field has been set.

### GetDnsIntegrityMember

`func (o *GetZoneAuthResponse) GetDnsIntegrityMember() string`

GetDnsIntegrityMember returns the DnsIntegrityMember field if non-nil, zero value otherwise.

### GetDnsIntegrityMemberOk

`func (o *GetZoneAuthResponse) GetDnsIntegrityMemberOk() (*string, bool)`

GetDnsIntegrityMemberOk returns a tuple with the DnsIntegrityMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrityMember

`func (o *GetZoneAuthResponse) SetDnsIntegrityMember(v string)`

SetDnsIntegrityMember sets DnsIntegrityMember field to given value.

### HasDnsIntegrityMember

`func (o *GetZoneAuthResponse) HasDnsIntegrityMember() bool`

HasDnsIntegrityMember returns a boolean if a field has been set.

### GetDnsIntegrityVerboseLogging

`func (o *GetZoneAuthResponse) GetDnsIntegrityVerboseLogging() bool`

GetDnsIntegrityVerboseLogging returns the DnsIntegrityVerboseLogging field if non-nil, zero value otherwise.

### GetDnsIntegrityVerboseLoggingOk

`func (o *GetZoneAuthResponse) GetDnsIntegrityVerboseLoggingOk() (*bool, bool)`

GetDnsIntegrityVerboseLoggingOk returns a tuple with the DnsIntegrityVerboseLogging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsIntegrityVerboseLogging

`func (o *GetZoneAuthResponse) SetDnsIntegrityVerboseLogging(v bool)`

SetDnsIntegrityVerboseLogging sets DnsIntegrityVerboseLogging field to given value.

### HasDnsIntegrityVerboseLogging

`func (o *GetZoneAuthResponse) HasDnsIntegrityVerboseLogging() bool`

HasDnsIntegrityVerboseLogging returns a boolean if a field has been set.

### GetDnsSoaEmail

`func (o *GetZoneAuthResponse) GetDnsSoaEmail() string`

GetDnsSoaEmail returns the DnsSoaEmail field if non-nil, zero value otherwise.

### GetDnsSoaEmailOk

`func (o *GetZoneAuthResponse) GetDnsSoaEmailOk() (*string, bool)`

GetDnsSoaEmailOk returns a tuple with the DnsSoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsSoaEmail

`func (o *GetZoneAuthResponse) SetDnsSoaEmail(v string)`

SetDnsSoaEmail sets DnsSoaEmail field to given value.

### HasDnsSoaEmail

`func (o *GetZoneAuthResponse) HasDnsSoaEmail() bool`

HasDnsSoaEmail returns a boolean if a field has been set.

### GetDnssecKeyParams

`func (o *GetZoneAuthResponse) GetDnssecKeyParams() ZoneAuthDnssecKeyParams`

GetDnssecKeyParams returns the DnssecKeyParams field if non-nil, zero value otherwise.

### GetDnssecKeyParamsOk

`func (o *GetZoneAuthResponse) GetDnssecKeyParamsOk() (*ZoneAuthDnssecKeyParams, bool)`

GetDnssecKeyParamsOk returns a tuple with the DnssecKeyParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecKeyParams

`func (o *GetZoneAuthResponse) SetDnssecKeyParams(v ZoneAuthDnssecKeyParams)`

SetDnssecKeyParams sets DnssecKeyParams field to given value.

### HasDnssecKeyParams

`func (o *GetZoneAuthResponse) HasDnssecKeyParams() bool`

HasDnssecKeyParams returns a boolean if a field has been set.

### GetDnssecKeys

`func (o *GetZoneAuthResponse) GetDnssecKeys() []ZoneAuthDnssecKeys`

GetDnssecKeys returns the DnssecKeys field if non-nil, zero value otherwise.

### GetDnssecKeysOk

`func (o *GetZoneAuthResponse) GetDnssecKeysOk() (*[]ZoneAuthDnssecKeys, bool)`

GetDnssecKeysOk returns a tuple with the DnssecKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecKeys

`func (o *GetZoneAuthResponse) SetDnssecKeys(v []ZoneAuthDnssecKeys)`

SetDnssecKeys sets DnssecKeys field to given value.

### HasDnssecKeys

`func (o *GetZoneAuthResponse) HasDnssecKeys() bool`

HasDnssecKeys returns a boolean if a field has been set.

### GetDnssecKskRolloverDate

`func (o *GetZoneAuthResponse) GetDnssecKskRolloverDate() int64`

GetDnssecKskRolloverDate returns the DnssecKskRolloverDate field if non-nil, zero value otherwise.

### GetDnssecKskRolloverDateOk

`func (o *GetZoneAuthResponse) GetDnssecKskRolloverDateOk() (*int64, bool)`

GetDnssecKskRolloverDateOk returns a tuple with the DnssecKskRolloverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecKskRolloverDate

`func (o *GetZoneAuthResponse) SetDnssecKskRolloverDate(v int64)`

SetDnssecKskRolloverDate sets DnssecKskRolloverDate field to given value.

### HasDnssecKskRolloverDate

`func (o *GetZoneAuthResponse) HasDnssecKskRolloverDate() bool`

HasDnssecKskRolloverDate returns a boolean if a field has been set.

### GetDnssecZskRolloverDate

`func (o *GetZoneAuthResponse) GetDnssecZskRolloverDate() int64`

GetDnssecZskRolloverDate returns the DnssecZskRolloverDate field if non-nil, zero value otherwise.

### GetDnssecZskRolloverDateOk

`func (o *GetZoneAuthResponse) GetDnssecZskRolloverDateOk() (*int64, bool)`

GetDnssecZskRolloverDateOk returns a tuple with the DnssecZskRolloverDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecZskRolloverDate

`func (o *GetZoneAuthResponse) SetDnssecZskRolloverDate(v int64)`

SetDnssecZskRolloverDate sets DnssecZskRolloverDate field to given value.

### HasDnssecZskRolloverDate

`func (o *GetZoneAuthResponse) HasDnssecZskRolloverDate() bool`

HasDnssecZskRolloverDate returns a boolean if a field has been set.

### GetDoHostAbstraction

`func (o *GetZoneAuthResponse) GetDoHostAbstraction() bool`

GetDoHostAbstraction returns the DoHostAbstraction field if non-nil, zero value otherwise.

### GetDoHostAbstractionOk

`func (o *GetZoneAuthResponse) GetDoHostAbstractionOk() (*bool, bool)`

GetDoHostAbstractionOk returns a tuple with the DoHostAbstraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoHostAbstraction

`func (o *GetZoneAuthResponse) SetDoHostAbstraction(v bool)`

SetDoHostAbstraction sets DoHostAbstraction field to given value.

### HasDoHostAbstraction

`func (o *GetZoneAuthResponse) HasDoHostAbstraction() bool`

HasDoHostAbstraction returns a boolean if a field has been set.

### GetEffectiveCheckNamesPolicy

`func (o *GetZoneAuthResponse) GetEffectiveCheckNamesPolicy() string`

GetEffectiveCheckNamesPolicy returns the EffectiveCheckNamesPolicy field if non-nil, zero value otherwise.

### GetEffectiveCheckNamesPolicyOk

`func (o *GetZoneAuthResponse) GetEffectiveCheckNamesPolicyOk() (*string, bool)`

GetEffectiveCheckNamesPolicyOk returns a tuple with the EffectiveCheckNamesPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCheckNamesPolicy

`func (o *GetZoneAuthResponse) SetEffectiveCheckNamesPolicy(v string)`

SetEffectiveCheckNamesPolicy sets EffectiveCheckNamesPolicy field to given value.

### HasEffectiveCheckNamesPolicy

`func (o *GetZoneAuthResponse) HasEffectiveCheckNamesPolicy() bool`

HasEffectiveCheckNamesPolicy returns a boolean if a field has been set.

### GetEffectiveRecordNamePolicy

`func (o *GetZoneAuthResponse) GetEffectiveRecordNamePolicy() string`

GetEffectiveRecordNamePolicy returns the EffectiveRecordNamePolicy field if non-nil, zero value otherwise.

### GetEffectiveRecordNamePolicyOk

`func (o *GetZoneAuthResponse) GetEffectiveRecordNamePolicyOk() (*string, bool)`

GetEffectiveRecordNamePolicyOk returns a tuple with the EffectiveRecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveRecordNamePolicy

`func (o *GetZoneAuthResponse) SetEffectiveRecordNamePolicy(v string)`

SetEffectiveRecordNamePolicy sets EffectiveRecordNamePolicy field to given value.

### HasEffectiveRecordNamePolicy

`func (o *GetZoneAuthResponse) HasEffectiveRecordNamePolicy() bool`

HasEffectiveRecordNamePolicy returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetZoneAuthResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetZoneAuthResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetZoneAuthResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetZoneAuthResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetZoneAuthResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetZoneAuthResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetZoneAuthResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetZoneAuthResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetZoneAuthResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetZoneAuthResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetZoneAuthResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetZoneAuthResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetExternalPrimaries

`func (o *GetZoneAuthResponse) GetExternalPrimaries() []ZoneAuthExternalPrimaries`

GetExternalPrimaries returns the ExternalPrimaries field if non-nil, zero value otherwise.

### GetExternalPrimariesOk

`func (o *GetZoneAuthResponse) GetExternalPrimariesOk() (*[]ZoneAuthExternalPrimaries, bool)`

GetExternalPrimariesOk returns a tuple with the ExternalPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPrimaries

`func (o *GetZoneAuthResponse) SetExternalPrimaries(v []ZoneAuthExternalPrimaries)`

SetExternalPrimaries sets ExternalPrimaries field to given value.

### HasExternalPrimaries

`func (o *GetZoneAuthResponse) HasExternalPrimaries() bool`

HasExternalPrimaries returns a boolean if a field has been set.

### GetExternalSecondaries

`func (o *GetZoneAuthResponse) GetExternalSecondaries() []ZoneAuthExternalSecondaries`

GetExternalSecondaries returns the ExternalSecondaries field if non-nil, zero value otherwise.

### GetExternalSecondariesOk

`func (o *GetZoneAuthResponse) GetExternalSecondariesOk() (*[]ZoneAuthExternalSecondaries, bool)`

GetExternalSecondariesOk returns a tuple with the ExternalSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSecondaries

`func (o *GetZoneAuthResponse) SetExternalSecondaries(v []ZoneAuthExternalSecondaries)`

SetExternalSecondaries sets ExternalSecondaries field to given value.

### HasExternalSecondaries

`func (o *GetZoneAuthResponse) HasExternalSecondaries() bool`

HasExternalSecondaries returns a boolean if a field has been set.

### GetFqdn

`func (o *GetZoneAuthResponse) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *GetZoneAuthResponse) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *GetZoneAuthResponse) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *GetZoneAuthResponse) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetGridPrimary

`func (o *GetZoneAuthResponse) GetGridPrimary() []ZoneAuthGridPrimary`

GetGridPrimary returns the GridPrimary field if non-nil, zero value otherwise.

### GetGridPrimaryOk

`func (o *GetZoneAuthResponse) GetGridPrimaryOk() (*[]ZoneAuthGridPrimary, bool)`

GetGridPrimaryOk returns a tuple with the GridPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimary

`func (o *GetZoneAuthResponse) SetGridPrimary(v []ZoneAuthGridPrimary)`

SetGridPrimary sets GridPrimary field to given value.

### HasGridPrimary

`func (o *GetZoneAuthResponse) HasGridPrimary() bool`

HasGridPrimary returns a boolean if a field has been set.

### GetGridPrimarySharedWithMsParentDelegation

`func (o *GetZoneAuthResponse) GetGridPrimarySharedWithMsParentDelegation() bool`

GetGridPrimarySharedWithMsParentDelegation returns the GridPrimarySharedWithMsParentDelegation field if non-nil, zero value otherwise.

### GetGridPrimarySharedWithMsParentDelegationOk

`func (o *GetZoneAuthResponse) GetGridPrimarySharedWithMsParentDelegationOk() (*bool, bool)`

GetGridPrimarySharedWithMsParentDelegationOk returns a tuple with the GridPrimarySharedWithMsParentDelegation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridPrimarySharedWithMsParentDelegation

`func (o *GetZoneAuthResponse) SetGridPrimarySharedWithMsParentDelegation(v bool)`

SetGridPrimarySharedWithMsParentDelegation sets GridPrimarySharedWithMsParentDelegation field to given value.

### HasGridPrimarySharedWithMsParentDelegation

`func (o *GetZoneAuthResponse) HasGridPrimarySharedWithMsParentDelegation() bool`

HasGridPrimarySharedWithMsParentDelegation returns a boolean if a field has been set.

### GetGridSecondaries

`func (o *GetZoneAuthResponse) GetGridSecondaries() []ZoneAuthGridSecondaries`

GetGridSecondaries returns the GridSecondaries field if non-nil, zero value otherwise.

### GetGridSecondariesOk

`func (o *GetZoneAuthResponse) GetGridSecondariesOk() (*[]ZoneAuthGridSecondaries, bool)`

GetGridSecondariesOk returns a tuple with the GridSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridSecondaries

`func (o *GetZoneAuthResponse) SetGridSecondaries(v []ZoneAuthGridSecondaries)`

SetGridSecondaries sets GridSecondaries field to given value.

### HasGridSecondaries

`func (o *GetZoneAuthResponse) HasGridSecondaries() bool`

HasGridSecondaries returns a boolean if a field has been set.

### GetImportFrom

`func (o *GetZoneAuthResponse) GetImportFrom() string`

GetImportFrom returns the ImportFrom field if non-nil, zero value otherwise.

### GetImportFromOk

`func (o *GetZoneAuthResponse) GetImportFromOk() (*string, bool)`

GetImportFromOk returns a tuple with the ImportFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportFrom

`func (o *GetZoneAuthResponse) SetImportFrom(v string)`

SetImportFrom sets ImportFrom field to given value.

### HasImportFrom

`func (o *GetZoneAuthResponse) HasImportFrom() bool`

HasImportFrom returns a boolean if a field has been set.

### GetIsDnssecEnabled

`func (o *GetZoneAuthResponse) GetIsDnssecEnabled() bool`

GetIsDnssecEnabled returns the IsDnssecEnabled field if non-nil, zero value otherwise.

### GetIsDnssecEnabledOk

`func (o *GetZoneAuthResponse) GetIsDnssecEnabledOk() (*bool, bool)`

GetIsDnssecEnabledOk returns a tuple with the IsDnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDnssecEnabled

`func (o *GetZoneAuthResponse) SetIsDnssecEnabled(v bool)`

SetIsDnssecEnabled sets IsDnssecEnabled field to given value.

### HasIsDnssecEnabled

`func (o *GetZoneAuthResponse) HasIsDnssecEnabled() bool`

HasIsDnssecEnabled returns a boolean if a field has been set.

### GetIsDnssecSigned

`func (o *GetZoneAuthResponse) GetIsDnssecSigned() bool`

GetIsDnssecSigned returns the IsDnssecSigned field if non-nil, zero value otherwise.

### GetIsDnssecSignedOk

`func (o *GetZoneAuthResponse) GetIsDnssecSignedOk() (*bool, bool)`

GetIsDnssecSignedOk returns a tuple with the IsDnssecSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDnssecSigned

`func (o *GetZoneAuthResponse) SetIsDnssecSigned(v bool)`

SetIsDnssecSigned sets IsDnssecSigned field to given value.

### HasIsDnssecSigned

`func (o *GetZoneAuthResponse) HasIsDnssecSigned() bool`

HasIsDnssecSigned returns a boolean if a field has been set.

### GetIsMultimaster

`func (o *GetZoneAuthResponse) GetIsMultimaster() bool`

GetIsMultimaster returns the IsMultimaster field if non-nil, zero value otherwise.

### GetIsMultimasterOk

`func (o *GetZoneAuthResponse) GetIsMultimasterOk() (*bool, bool)`

GetIsMultimasterOk returns a tuple with the IsMultimaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMultimaster

`func (o *GetZoneAuthResponse) SetIsMultimaster(v bool)`

SetIsMultimaster sets IsMultimaster field to given value.

### HasIsMultimaster

`func (o *GetZoneAuthResponse) HasIsMultimaster() bool`

HasIsMultimaster returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetZoneAuthResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetZoneAuthResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetZoneAuthResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetZoneAuthResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetLastQueriedAcl

`func (o *GetZoneAuthResponse) GetLastQueriedAcl() []ZoneAuthLastQueriedAcl`

GetLastQueriedAcl returns the LastQueriedAcl field if non-nil, zero value otherwise.

### GetLastQueriedAclOk

`func (o *GetZoneAuthResponse) GetLastQueriedAclOk() (*[]ZoneAuthLastQueriedAcl, bool)`

GetLastQueriedAclOk returns a tuple with the LastQueriedAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueriedAcl

`func (o *GetZoneAuthResponse) SetLastQueriedAcl(v []ZoneAuthLastQueriedAcl)`

SetLastQueriedAcl sets LastQueriedAcl field to given value.

### HasLastQueriedAcl

`func (o *GetZoneAuthResponse) HasLastQueriedAcl() bool`

HasLastQueriedAcl returns a boolean if a field has been set.

### GetLocked

`func (o *GetZoneAuthResponse) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *GetZoneAuthResponse) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *GetZoneAuthResponse) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *GetZoneAuthResponse) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLockedBy

`func (o *GetZoneAuthResponse) GetLockedBy() string`

GetLockedBy returns the LockedBy field if non-nil, zero value otherwise.

### GetLockedByOk

`func (o *GetZoneAuthResponse) GetLockedByOk() (*string, bool)`

GetLockedByOk returns a tuple with the LockedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedBy

`func (o *GetZoneAuthResponse) SetLockedBy(v string)`

SetLockedBy sets LockedBy field to given value.

### HasLockedBy

`func (o *GetZoneAuthResponse) HasLockedBy() bool`

HasLockedBy returns a boolean if a field has been set.

### GetMaskPrefix

`func (o *GetZoneAuthResponse) GetMaskPrefix() string`

GetMaskPrefix returns the MaskPrefix field if non-nil, zero value otherwise.

### GetMaskPrefixOk

`func (o *GetZoneAuthResponse) GetMaskPrefixOk() (*string, bool)`

GetMaskPrefixOk returns a tuple with the MaskPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaskPrefix

`func (o *GetZoneAuthResponse) SetMaskPrefix(v string)`

SetMaskPrefix sets MaskPrefix field to given value.

### HasMaskPrefix

`func (o *GetZoneAuthResponse) HasMaskPrefix() bool`

HasMaskPrefix returns a boolean if a field has been set.

### GetMemberSoaMnames

`func (o *GetZoneAuthResponse) GetMemberSoaMnames() []ZoneAuthMemberSoaMnames`

GetMemberSoaMnames returns the MemberSoaMnames field if non-nil, zero value otherwise.

### GetMemberSoaMnamesOk

`func (o *GetZoneAuthResponse) GetMemberSoaMnamesOk() (*[]ZoneAuthMemberSoaMnames, bool)`

GetMemberSoaMnamesOk returns a tuple with the MemberSoaMnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSoaMnames

`func (o *GetZoneAuthResponse) SetMemberSoaMnames(v []ZoneAuthMemberSoaMnames)`

SetMemberSoaMnames sets MemberSoaMnames field to given value.

### HasMemberSoaMnames

`func (o *GetZoneAuthResponse) HasMemberSoaMnames() bool`

HasMemberSoaMnames returns a boolean if a field has been set.

### GetMemberSoaSerials

`func (o *GetZoneAuthResponse) GetMemberSoaSerials() []ZoneAuthMemberSoaSerials`

GetMemberSoaSerials returns the MemberSoaSerials field if non-nil, zero value otherwise.

### GetMemberSoaSerialsOk

`func (o *GetZoneAuthResponse) GetMemberSoaSerialsOk() (*[]ZoneAuthMemberSoaSerials, bool)`

GetMemberSoaSerialsOk returns a tuple with the MemberSoaSerials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberSoaSerials

`func (o *GetZoneAuthResponse) SetMemberSoaSerials(v []ZoneAuthMemberSoaSerials)`

SetMemberSoaSerials sets MemberSoaSerials field to given value.

### HasMemberSoaSerials

`func (o *GetZoneAuthResponse) HasMemberSoaSerials() bool`

HasMemberSoaSerials returns a boolean if a field has been set.

### GetMgmPrivate

`func (o *GetZoneAuthResponse) GetMgmPrivate() bool`

GetMgmPrivate returns the MgmPrivate field if non-nil, zero value otherwise.

### GetMgmPrivateOk

`func (o *GetZoneAuthResponse) GetMgmPrivateOk() (*bool, bool)`

GetMgmPrivateOk returns a tuple with the MgmPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivate

`func (o *GetZoneAuthResponse) SetMgmPrivate(v bool)`

SetMgmPrivate sets MgmPrivate field to given value.

### HasMgmPrivate

`func (o *GetZoneAuthResponse) HasMgmPrivate() bool`

HasMgmPrivate returns a boolean if a field has been set.

### GetMgmPrivateOverridable

`func (o *GetZoneAuthResponse) GetMgmPrivateOverridable() bool`

GetMgmPrivateOverridable returns the MgmPrivateOverridable field if non-nil, zero value otherwise.

### GetMgmPrivateOverridableOk

`func (o *GetZoneAuthResponse) GetMgmPrivateOverridableOk() (*bool, bool)`

GetMgmPrivateOverridableOk returns a tuple with the MgmPrivateOverridable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmPrivateOverridable

`func (o *GetZoneAuthResponse) SetMgmPrivateOverridable(v bool)`

SetMgmPrivateOverridable sets MgmPrivateOverridable field to given value.

### HasMgmPrivateOverridable

`func (o *GetZoneAuthResponse) HasMgmPrivateOverridable() bool`

HasMgmPrivateOverridable returns a boolean if a field has been set.

### GetMsAdIntegrated

`func (o *GetZoneAuthResponse) GetMsAdIntegrated() bool`

GetMsAdIntegrated returns the MsAdIntegrated field if non-nil, zero value otherwise.

### GetMsAdIntegratedOk

`func (o *GetZoneAuthResponse) GetMsAdIntegratedOk() (*bool, bool)`

GetMsAdIntegratedOk returns a tuple with the MsAdIntegrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdIntegrated

`func (o *GetZoneAuthResponse) SetMsAdIntegrated(v bool)`

SetMsAdIntegrated sets MsAdIntegrated field to given value.

### HasMsAdIntegrated

`func (o *GetZoneAuthResponse) HasMsAdIntegrated() bool`

HasMsAdIntegrated returns a boolean if a field has been set.

### GetMsAllowTransfer

`func (o *GetZoneAuthResponse) GetMsAllowTransfer() []ZoneAuthMsAllowTransfer`

GetMsAllowTransfer returns the MsAllowTransfer field if non-nil, zero value otherwise.

### GetMsAllowTransferOk

`func (o *GetZoneAuthResponse) GetMsAllowTransferOk() (*[]ZoneAuthMsAllowTransfer, bool)`

GetMsAllowTransferOk returns a tuple with the MsAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAllowTransfer

`func (o *GetZoneAuthResponse) SetMsAllowTransfer(v []ZoneAuthMsAllowTransfer)`

SetMsAllowTransfer sets MsAllowTransfer field to given value.

### HasMsAllowTransfer

`func (o *GetZoneAuthResponse) HasMsAllowTransfer() bool`

HasMsAllowTransfer returns a boolean if a field has been set.

### GetMsAllowTransferMode

`func (o *GetZoneAuthResponse) GetMsAllowTransferMode() string`

GetMsAllowTransferMode returns the MsAllowTransferMode field if non-nil, zero value otherwise.

### GetMsAllowTransferModeOk

`func (o *GetZoneAuthResponse) GetMsAllowTransferModeOk() (*string, bool)`

GetMsAllowTransferModeOk returns a tuple with the MsAllowTransferMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAllowTransferMode

`func (o *GetZoneAuthResponse) SetMsAllowTransferMode(v string)`

SetMsAllowTransferMode sets MsAllowTransferMode field to given value.

### HasMsAllowTransferMode

`func (o *GetZoneAuthResponse) HasMsAllowTransferMode() bool`

HasMsAllowTransferMode returns a boolean if a field has been set.

### GetMsDcNsRecordCreation

`func (o *GetZoneAuthResponse) GetMsDcNsRecordCreation() []ZoneAuthMsDcNsRecordCreation`

GetMsDcNsRecordCreation returns the MsDcNsRecordCreation field if non-nil, zero value otherwise.

### GetMsDcNsRecordCreationOk

`func (o *GetZoneAuthResponse) GetMsDcNsRecordCreationOk() (*[]ZoneAuthMsDcNsRecordCreation, bool)`

GetMsDcNsRecordCreationOk returns a tuple with the MsDcNsRecordCreation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDcNsRecordCreation

`func (o *GetZoneAuthResponse) SetMsDcNsRecordCreation(v []ZoneAuthMsDcNsRecordCreation)`

SetMsDcNsRecordCreation sets MsDcNsRecordCreation field to given value.

### HasMsDcNsRecordCreation

`func (o *GetZoneAuthResponse) HasMsDcNsRecordCreation() bool`

HasMsDcNsRecordCreation returns a boolean if a field has been set.

### GetMsDdnsMode

`func (o *GetZoneAuthResponse) GetMsDdnsMode() string`

GetMsDdnsMode returns the MsDdnsMode field if non-nil, zero value otherwise.

### GetMsDdnsModeOk

`func (o *GetZoneAuthResponse) GetMsDdnsModeOk() (*string, bool)`

GetMsDdnsModeOk returns a tuple with the MsDdnsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsDdnsMode

`func (o *GetZoneAuthResponse) SetMsDdnsMode(v string)`

SetMsDdnsMode sets MsDdnsMode field to given value.

### HasMsDdnsMode

`func (o *GetZoneAuthResponse) HasMsDdnsMode() bool`

HasMsDdnsMode returns a boolean if a field has been set.

### GetMsManaged

`func (o *GetZoneAuthResponse) GetMsManaged() string`

GetMsManaged returns the MsManaged field if non-nil, zero value otherwise.

### GetMsManagedOk

`func (o *GetZoneAuthResponse) GetMsManagedOk() (*string, bool)`

GetMsManagedOk returns a tuple with the MsManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsManaged

`func (o *GetZoneAuthResponse) SetMsManaged(v string)`

SetMsManaged sets MsManaged field to given value.

### HasMsManaged

`func (o *GetZoneAuthResponse) HasMsManaged() bool`

HasMsManaged returns a boolean if a field has been set.

### GetMsPrimaries

`func (o *GetZoneAuthResponse) GetMsPrimaries() []ZoneAuthMsPrimaries`

GetMsPrimaries returns the MsPrimaries field if non-nil, zero value otherwise.

### GetMsPrimariesOk

`func (o *GetZoneAuthResponse) GetMsPrimariesOk() (*[]ZoneAuthMsPrimaries, bool)`

GetMsPrimariesOk returns a tuple with the MsPrimaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsPrimaries

`func (o *GetZoneAuthResponse) SetMsPrimaries(v []ZoneAuthMsPrimaries)`

SetMsPrimaries sets MsPrimaries field to given value.

### HasMsPrimaries

`func (o *GetZoneAuthResponse) HasMsPrimaries() bool`

HasMsPrimaries returns a boolean if a field has been set.

### GetMsReadOnly

`func (o *GetZoneAuthResponse) GetMsReadOnly() bool`

GetMsReadOnly returns the MsReadOnly field if non-nil, zero value otherwise.

### GetMsReadOnlyOk

`func (o *GetZoneAuthResponse) GetMsReadOnlyOk() (*bool, bool)`

GetMsReadOnlyOk returns a tuple with the MsReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsReadOnly

`func (o *GetZoneAuthResponse) SetMsReadOnly(v bool)`

SetMsReadOnly sets MsReadOnly field to given value.

### HasMsReadOnly

`func (o *GetZoneAuthResponse) HasMsReadOnly() bool`

HasMsReadOnly returns a boolean if a field has been set.

### GetMsSecondaries

`func (o *GetZoneAuthResponse) GetMsSecondaries() []ZoneAuthMsSecondaries`

GetMsSecondaries returns the MsSecondaries field if non-nil, zero value otherwise.

### GetMsSecondariesOk

`func (o *GetZoneAuthResponse) GetMsSecondariesOk() (*[]ZoneAuthMsSecondaries, bool)`

GetMsSecondariesOk returns a tuple with the MsSecondaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSecondaries

`func (o *GetZoneAuthResponse) SetMsSecondaries(v []ZoneAuthMsSecondaries)`

SetMsSecondaries sets MsSecondaries field to given value.

### HasMsSecondaries

`func (o *GetZoneAuthResponse) HasMsSecondaries() bool`

HasMsSecondaries returns a boolean if a field has been set.

### GetMsSyncDisabled

`func (o *GetZoneAuthResponse) GetMsSyncDisabled() bool`

GetMsSyncDisabled returns the MsSyncDisabled field if non-nil, zero value otherwise.

### GetMsSyncDisabledOk

`func (o *GetZoneAuthResponse) GetMsSyncDisabledOk() (*bool, bool)`

GetMsSyncDisabledOk returns a tuple with the MsSyncDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSyncDisabled

`func (o *GetZoneAuthResponse) SetMsSyncDisabled(v bool)`

SetMsSyncDisabled sets MsSyncDisabled field to given value.

### HasMsSyncDisabled

`func (o *GetZoneAuthResponse) HasMsSyncDisabled() bool`

HasMsSyncDisabled returns a boolean if a field has been set.

### GetMsSyncMasterName

`func (o *GetZoneAuthResponse) GetMsSyncMasterName() string`

GetMsSyncMasterName returns the MsSyncMasterName field if non-nil, zero value otherwise.

### GetMsSyncMasterNameOk

`func (o *GetZoneAuthResponse) GetMsSyncMasterNameOk() (*string, bool)`

GetMsSyncMasterNameOk returns a tuple with the MsSyncMasterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsSyncMasterName

`func (o *GetZoneAuthResponse) SetMsSyncMasterName(v string)`

SetMsSyncMasterName sets MsSyncMasterName field to given value.

### HasMsSyncMasterName

`func (o *GetZoneAuthResponse) HasMsSyncMasterName() bool`

HasMsSyncMasterName returns a boolean if a field has been set.

### GetNetworkAssociations

`func (o *GetZoneAuthResponse) GetNetworkAssociations() []string`

GetNetworkAssociations returns the NetworkAssociations field if non-nil, zero value otherwise.

### GetNetworkAssociationsOk

`func (o *GetZoneAuthResponse) GetNetworkAssociationsOk() (*[]string, bool)`

GetNetworkAssociationsOk returns a tuple with the NetworkAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAssociations

`func (o *GetZoneAuthResponse) SetNetworkAssociations(v []string)`

SetNetworkAssociations sets NetworkAssociations field to given value.

### HasNetworkAssociations

`func (o *GetZoneAuthResponse) HasNetworkAssociations() bool`

HasNetworkAssociations returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetZoneAuthResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetZoneAuthResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetZoneAuthResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetZoneAuthResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNotifyDelay

`func (o *GetZoneAuthResponse) GetNotifyDelay() int64`

GetNotifyDelay returns the NotifyDelay field if non-nil, zero value otherwise.

### GetNotifyDelayOk

`func (o *GetZoneAuthResponse) GetNotifyDelayOk() (*int64, bool)`

GetNotifyDelayOk returns a tuple with the NotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDelay

`func (o *GetZoneAuthResponse) SetNotifyDelay(v int64)`

SetNotifyDelay sets NotifyDelay field to given value.

### HasNotifyDelay

`func (o *GetZoneAuthResponse) HasNotifyDelay() bool`

HasNotifyDelay returns a boolean if a field has been set.

### GetNsGroup

`func (o *GetZoneAuthResponse) GetNsGroup() string`

GetNsGroup returns the NsGroup field if non-nil, zero value otherwise.

### GetNsGroupOk

`func (o *GetZoneAuthResponse) GetNsGroupOk() (*string, bool)`

GetNsGroupOk returns a tuple with the NsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsGroup

`func (o *GetZoneAuthResponse) SetNsGroup(v string)`

SetNsGroup sets NsGroup field to given value.

### HasNsGroup

`func (o *GetZoneAuthResponse) HasNsGroup() bool`

HasNsGroup returns a boolean if a field has been set.

### GetParent

`func (o *GetZoneAuthResponse) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetZoneAuthResponse) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetZoneAuthResponse) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetZoneAuthResponse) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPrefix

`func (o *GetZoneAuthResponse) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *GetZoneAuthResponse) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *GetZoneAuthResponse) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *GetZoneAuthResponse) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetPrimaryType

`func (o *GetZoneAuthResponse) GetPrimaryType() string`

GetPrimaryType returns the PrimaryType field if non-nil, zero value otherwise.

### GetPrimaryTypeOk

`func (o *GetZoneAuthResponse) GetPrimaryTypeOk() (*string, bool)`

GetPrimaryTypeOk returns a tuple with the PrimaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryType

`func (o *GetZoneAuthResponse) SetPrimaryType(v string)`

SetPrimaryType sets PrimaryType field to given value.

### HasPrimaryType

`func (o *GetZoneAuthResponse) HasPrimaryType() bool`

HasPrimaryType returns a boolean if a field has been set.

### GetRecordNamePolicy

`func (o *GetZoneAuthResponse) GetRecordNamePolicy() string`

GetRecordNamePolicy returns the RecordNamePolicy field if non-nil, zero value otherwise.

### GetRecordNamePolicyOk

`func (o *GetZoneAuthResponse) GetRecordNamePolicyOk() (*string, bool)`

GetRecordNamePolicyOk returns a tuple with the RecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordNamePolicy

`func (o *GetZoneAuthResponse) SetRecordNamePolicy(v string)`

SetRecordNamePolicy sets RecordNamePolicy field to given value.

### HasRecordNamePolicy

`func (o *GetZoneAuthResponse) HasRecordNamePolicy() bool`

HasRecordNamePolicy returns a boolean if a field has been set.

### GetRecordsMonitored

`func (o *GetZoneAuthResponse) GetRecordsMonitored() bool`

GetRecordsMonitored returns the RecordsMonitored field if non-nil, zero value otherwise.

### GetRecordsMonitoredOk

`func (o *GetZoneAuthResponse) GetRecordsMonitoredOk() (*bool, bool)`

GetRecordsMonitoredOk returns a tuple with the RecordsMonitored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsMonitored

`func (o *GetZoneAuthResponse) SetRecordsMonitored(v bool)`

SetRecordsMonitored sets RecordsMonitored field to given value.

### HasRecordsMonitored

`func (o *GetZoneAuthResponse) HasRecordsMonitored() bool`

HasRecordsMonitored returns a boolean if a field has been set.

### GetRemoveSubzones

`func (o *GetZoneAuthResponse) GetRemoveSubzones() bool`

GetRemoveSubzones returns the RemoveSubzones field if non-nil, zero value otherwise.

### GetRemoveSubzonesOk

`func (o *GetZoneAuthResponse) GetRemoveSubzonesOk() (*bool, bool)`

GetRemoveSubzonesOk returns a tuple with the RemoveSubzones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveSubzones

`func (o *GetZoneAuthResponse) SetRemoveSubzones(v bool)`

SetRemoveSubzones sets RemoveSubzones field to given value.

### HasRemoveSubzones

`func (o *GetZoneAuthResponse) HasRemoveSubzones() bool`

HasRemoveSubzones returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *GetZoneAuthResponse) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *GetZoneAuthResponse) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *GetZoneAuthResponse) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *GetZoneAuthResponse) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetRrNotQueriedEnabledTime

`func (o *GetZoneAuthResponse) GetRrNotQueriedEnabledTime() int64`

GetRrNotQueriedEnabledTime returns the RrNotQueriedEnabledTime field if non-nil, zero value otherwise.

### GetRrNotQueriedEnabledTimeOk

`func (o *GetZoneAuthResponse) GetRrNotQueriedEnabledTimeOk() (*int64, bool)`

GetRrNotQueriedEnabledTimeOk returns a tuple with the RrNotQueriedEnabledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrNotQueriedEnabledTime

`func (o *GetZoneAuthResponse) SetRrNotQueriedEnabledTime(v int64)`

SetRrNotQueriedEnabledTime sets RrNotQueriedEnabledTime field to given value.

### HasRrNotQueriedEnabledTime

`func (o *GetZoneAuthResponse) HasRrNotQueriedEnabledTime() bool`

HasRrNotQueriedEnabledTime returns a boolean if a field has been set.

### GetScavengingSettings

`func (o *GetZoneAuthResponse) GetScavengingSettings() ZoneAuthScavengingSettings`

GetScavengingSettings returns the ScavengingSettings field if non-nil, zero value otherwise.

### GetScavengingSettingsOk

`func (o *GetZoneAuthResponse) GetScavengingSettingsOk() (*ZoneAuthScavengingSettings, bool)`

GetScavengingSettingsOk returns a tuple with the ScavengingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScavengingSettings

`func (o *GetZoneAuthResponse) SetScavengingSettings(v ZoneAuthScavengingSettings)`

SetScavengingSettings sets ScavengingSettings field to given value.

### HasScavengingSettings

`func (o *GetZoneAuthResponse) HasScavengingSettings() bool`

HasScavengingSettings returns a boolean if a field has been set.

### GetSetSoaSerialNumber

`func (o *GetZoneAuthResponse) GetSetSoaSerialNumber() bool`

GetSetSoaSerialNumber returns the SetSoaSerialNumber field if non-nil, zero value otherwise.

### GetSetSoaSerialNumberOk

`func (o *GetZoneAuthResponse) GetSetSoaSerialNumberOk() (*bool, bool)`

GetSetSoaSerialNumberOk returns a tuple with the SetSoaSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetSoaSerialNumber

`func (o *GetZoneAuthResponse) SetSetSoaSerialNumber(v bool)`

SetSetSoaSerialNumber sets SetSoaSerialNumber field to given value.

### HasSetSoaSerialNumber

`func (o *GetZoneAuthResponse) HasSetSoaSerialNumber() bool`

HasSetSoaSerialNumber returns a boolean if a field has been set.

### GetSoaDefaultTtl

`func (o *GetZoneAuthResponse) GetSoaDefaultTtl() int64`

GetSoaDefaultTtl returns the SoaDefaultTtl field if non-nil, zero value otherwise.

### GetSoaDefaultTtlOk

`func (o *GetZoneAuthResponse) GetSoaDefaultTtlOk() (*int64, bool)`

GetSoaDefaultTtlOk returns a tuple with the SoaDefaultTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaDefaultTtl

`func (o *GetZoneAuthResponse) SetSoaDefaultTtl(v int64)`

SetSoaDefaultTtl sets SoaDefaultTtl field to given value.

### HasSoaDefaultTtl

`func (o *GetZoneAuthResponse) HasSoaDefaultTtl() bool`

HasSoaDefaultTtl returns a boolean if a field has been set.

### GetSoaEmail

`func (o *GetZoneAuthResponse) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *GetZoneAuthResponse) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *GetZoneAuthResponse) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *GetZoneAuthResponse) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetSoaExpire

`func (o *GetZoneAuthResponse) GetSoaExpire() int64`

GetSoaExpire returns the SoaExpire field if non-nil, zero value otherwise.

### GetSoaExpireOk

`func (o *GetZoneAuthResponse) GetSoaExpireOk() (*int64, bool)`

GetSoaExpireOk returns a tuple with the SoaExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaExpire

`func (o *GetZoneAuthResponse) SetSoaExpire(v int64)`

SetSoaExpire sets SoaExpire field to given value.

### HasSoaExpire

`func (o *GetZoneAuthResponse) HasSoaExpire() bool`

HasSoaExpire returns a boolean if a field has been set.

### GetSoaNegativeTtl

`func (o *GetZoneAuthResponse) GetSoaNegativeTtl() int64`

GetSoaNegativeTtl returns the SoaNegativeTtl field if non-nil, zero value otherwise.

### GetSoaNegativeTtlOk

`func (o *GetZoneAuthResponse) GetSoaNegativeTtlOk() (*int64, bool)`

GetSoaNegativeTtlOk returns a tuple with the SoaNegativeTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaNegativeTtl

`func (o *GetZoneAuthResponse) SetSoaNegativeTtl(v int64)`

SetSoaNegativeTtl sets SoaNegativeTtl field to given value.

### HasSoaNegativeTtl

`func (o *GetZoneAuthResponse) HasSoaNegativeTtl() bool`

HasSoaNegativeTtl returns a boolean if a field has been set.

### GetSoaRefresh

`func (o *GetZoneAuthResponse) GetSoaRefresh() int64`

GetSoaRefresh returns the SoaRefresh field if non-nil, zero value otherwise.

### GetSoaRefreshOk

`func (o *GetZoneAuthResponse) GetSoaRefreshOk() (*int64, bool)`

GetSoaRefreshOk returns a tuple with the SoaRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaRefresh

`func (o *GetZoneAuthResponse) SetSoaRefresh(v int64)`

SetSoaRefresh sets SoaRefresh field to given value.

### HasSoaRefresh

`func (o *GetZoneAuthResponse) HasSoaRefresh() bool`

HasSoaRefresh returns a boolean if a field has been set.

### GetSoaRetry

`func (o *GetZoneAuthResponse) GetSoaRetry() int64`

GetSoaRetry returns the SoaRetry field if non-nil, zero value otherwise.

### GetSoaRetryOk

`func (o *GetZoneAuthResponse) GetSoaRetryOk() (*int64, bool)`

GetSoaRetryOk returns a tuple with the SoaRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaRetry

`func (o *GetZoneAuthResponse) SetSoaRetry(v int64)`

SetSoaRetry sets SoaRetry field to given value.

### HasSoaRetry

`func (o *GetZoneAuthResponse) HasSoaRetry() bool`

HasSoaRetry returns a boolean if a field has been set.

### GetSoaSerial

`func (o *GetZoneAuthResponse) GetSoaSerial() int64`

GetSoaSerial returns the SoaSerial field if non-nil, zero value otherwise.

### GetSoaSerialOk

`func (o *GetZoneAuthResponse) GetSoaSerialOk() (*int64, bool)`

GetSoaSerialOk returns a tuple with the SoaSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaSerial

`func (o *GetZoneAuthResponse) SetSoaSerial(v int64)`

SetSoaSerial sets SoaSerial field to given value.

### HasSoaSerial

`func (o *GetZoneAuthResponse) HasSoaSerial() bool`

HasSoaSerial returns a boolean if a field has been set.

### GetSrgs

`func (o *GetZoneAuthResponse) GetSrgs() []string`

GetSrgs returns the Srgs field if non-nil, zero value otherwise.

### GetSrgsOk

`func (o *GetZoneAuthResponse) GetSrgsOk() (*[]string, bool)`

GetSrgsOk returns a tuple with the Srgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrgs

`func (o *GetZoneAuthResponse) SetSrgs(v []string)`

SetSrgs sets Srgs field to given value.

### HasSrgs

`func (o *GetZoneAuthResponse) HasSrgs() bool`

HasSrgs returns a boolean if a field has been set.

### GetUpdateForwarding

`func (o *GetZoneAuthResponse) GetUpdateForwarding() []ZoneAuthUpdateForwarding`

GetUpdateForwarding returns the UpdateForwarding field if non-nil, zero value otherwise.

### GetUpdateForwardingOk

`func (o *GetZoneAuthResponse) GetUpdateForwardingOk() (*[]ZoneAuthUpdateForwarding, bool)`

GetUpdateForwardingOk returns a tuple with the UpdateForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateForwarding

`func (o *GetZoneAuthResponse) SetUpdateForwarding(v []ZoneAuthUpdateForwarding)`

SetUpdateForwarding sets UpdateForwarding field to given value.

### HasUpdateForwarding

`func (o *GetZoneAuthResponse) HasUpdateForwarding() bool`

HasUpdateForwarding returns a boolean if a field has been set.

### GetUseAllowActiveDir

`func (o *GetZoneAuthResponse) GetUseAllowActiveDir() bool`

GetUseAllowActiveDir returns the UseAllowActiveDir field if non-nil, zero value otherwise.

### GetUseAllowActiveDirOk

`func (o *GetZoneAuthResponse) GetUseAllowActiveDirOk() (*bool, bool)`

GetUseAllowActiveDirOk returns a tuple with the UseAllowActiveDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowActiveDir

`func (o *GetZoneAuthResponse) SetUseAllowActiveDir(v bool)`

SetUseAllowActiveDir sets UseAllowActiveDir field to given value.

### HasUseAllowActiveDir

`func (o *GetZoneAuthResponse) HasUseAllowActiveDir() bool`

HasUseAllowActiveDir returns a boolean if a field has been set.

### GetUseAllowQuery

`func (o *GetZoneAuthResponse) GetUseAllowQuery() bool`

GetUseAllowQuery returns the UseAllowQuery field if non-nil, zero value otherwise.

### GetUseAllowQueryOk

`func (o *GetZoneAuthResponse) GetUseAllowQueryOk() (*bool, bool)`

GetUseAllowQueryOk returns a tuple with the UseAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowQuery

`func (o *GetZoneAuthResponse) SetUseAllowQuery(v bool)`

SetUseAllowQuery sets UseAllowQuery field to given value.

### HasUseAllowQuery

`func (o *GetZoneAuthResponse) HasUseAllowQuery() bool`

HasUseAllowQuery returns a boolean if a field has been set.

### GetUseAllowTransfer

`func (o *GetZoneAuthResponse) GetUseAllowTransfer() bool`

GetUseAllowTransfer returns the UseAllowTransfer field if non-nil, zero value otherwise.

### GetUseAllowTransferOk

`func (o *GetZoneAuthResponse) GetUseAllowTransferOk() (*bool, bool)`

GetUseAllowTransferOk returns a tuple with the UseAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowTransfer

`func (o *GetZoneAuthResponse) SetUseAllowTransfer(v bool)`

SetUseAllowTransfer sets UseAllowTransfer field to given value.

### HasUseAllowTransfer

`func (o *GetZoneAuthResponse) HasUseAllowTransfer() bool`

HasUseAllowTransfer returns a boolean if a field has been set.

### GetUseAllowUpdate

`func (o *GetZoneAuthResponse) GetUseAllowUpdate() bool`

GetUseAllowUpdate returns the UseAllowUpdate field if non-nil, zero value otherwise.

### GetUseAllowUpdateOk

`func (o *GetZoneAuthResponse) GetUseAllowUpdateOk() (*bool, bool)`

GetUseAllowUpdateOk returns a tuple with the UseAllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowUpdate

`func (o *GetZoneAuthResponse) SetUseAllowUpdate(v bool)`

SetUseAllowUpdate sets UseAllowUpdate field to given value.

### HasUseAllowUpdate

`func (o *GetZoneAuthResponse) HasUseAllowUpdate() bool`

HasUseAllowUpdate returns a boolean if a field has been set.

### GetUseAllowUpdateForwarding

`func (o *GetZoneAuthResponse) GetUseAllowUpdateForwarding() bool`

GetUseAllowUpdateForwarding returns the UseAllowUpdateForwarding field if non-nil, zero value otherwise.

### GetUseAllowUpdateForwardingOk

`func (o *GetZoneAuthResponse) GetUseAllowUpdateForwardingOk() (*bool, bool)`

GetUseAllowUpdateForwardingOk returns a tuple with the UseAllowUpdateForwarding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseAllowUpdateForwarding

`func (o *GetZoneAuthResponse) SetUseAllowUpdateForwarding(v bool)`

SetUseAllowUpdateForwarding sets UseAllowUpdateForwarding field to given value.

### HasUseAllowUpdateForwarding

`func (o *GetZoneAuthResponse) HasUseAllowUpdateForwarding() bool`

HasUseAllowUpdateForwarding returns a boolean if a field has been set.

### GetUseCheckNamesPolicy

`func (o *GetZoneAuthResponse) GetUseCheckNamesPolicy() bool`

GetUseCheckNamesPolicy returns the UseCheckNamesPolicy field if non-nil, zero value otherwise.

### GetUseCheckNamesPolicyOk

`func (o *GetZoneAuthResponse) GetUseCheckNamesPolicyOk() (*bool, bool)`

GetUseCheckNamesPolicyOk returns a tuple with the UseCheckNamesPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCheckNamesPolicy

`func (o *GetZoneAuthResponse) SetUseCheckNamesPolicy(v bool)`

SetUseCheckNamesPolicy sets UseCheckNamesPolicy field to given value.

### HasUseCheckNamesPolicy

`func (o *GetZoneAuthResponse) HasUseCheckNamesPolicy() bool`

HasUseCheckNamesPolicy returns a boolean if a field has been set.

### GetUseCopyXferToNotify

`func (o *GetZoneAuthResponse) GetUseCopyXferToNotify() bool`

GetUseCopyXferToNotify returns the UseCopyXferToNotify field if non-nil, zero value otherwise.

### GetUseCopyXferToNotifyOk

`func (o *GetZoneAuthResponse) GetUseCopyXferToNotifyOk() (*bool, bool)`

GetUseCopyXferToNotifyOk returns a tuple with the UseCopyXferToNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCopyXferToNotify

`func (o *GetZoneAuthResponse) SetUseCopyXferToNotify(v bool)`

SetUseCopyXferToNotify sets UseCopyXferToNotify field to given value.

### HasUseCopyXferToNotify

`func (o *GetZoneAuthResponse) HasUseCopyXferToNotify() bool`

HasUseCopyXferToNotify returns a boolean if a field has been set.

### GetUseDdnsForceCreationTimestampUpdate

`func (o *GetZoneAuthResponse) GetUseDdnsForceCreationTimestampUpdate() bool`

GetUseDdnsForceCreationTimestampUpdate returns the UseDdnsForceCreationTimestampUpdate field if non-nil, zero value otherwise.

### GetUseDdnsForceCreationTimestampUpdateOk

`func (o *GetZoneAuthResponse) GetUseDdnsForceCreationTimestampUpdateOk() (*bool, bool)`

GetUseDdnsForceCreationTimestampUpdateOk returns a tuple with the UseDdnsForceCreationTimestampUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsForceCreationTimestampUpdate

`func (o *GetZoneAuthResponse) SetUseDdnsForceCreationTimestampUpdate(v bool)`

SetUseDdnsForceCreationTimestampUpdate sets UseDdnsForceCreationTimestampUpdate field to given value.

### HasUseDdnsForceCreationTimestampUpdate

`func (o *GetZoneAuthResponse) HasUseDdnsForceCreationTimestampUpdate() bool`

HasUseDdnsForceCreationTimestampUpdate returns a boolean if a field has been set.

### GetUseDdnsPatternsRestriction

`func (o *GetZoneAuthResponse) GetUseDdnsPatternsRestriction() bool`

GetUseDdnsPatternsRestriction returns the UseDdnsPatternsRestriction field if non-nil, zero value otherwise.

### GetUseDdnsPatternsRestrictionOk

`func (o *GetZoneAuthResponse) GetUseDdnsPatternsRestrictionOk() (*bool, bool)`

GetUseDdnsPatternsRestrictionOk returns a tuple with the UseDdnsPatternsRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsPatternsRestriction

`func (o *GetZoneAuthResponse) SetUseDdnsPatternsRestriction(v bool)`

SetUseDdnsPatternsRestriction sets UseDdnsPatternsRestriction field to given value.

### HasUseDdnsPatternsRestriction

`func (o *GetZoneAuthResponse) HasUseDdnsPatternsRestriction() bool`

HasUseDdnsPatternsRestriction returns a boolean if a field has been set.

### GetUseDdnsPrincipalSecurity

`func (o *GetZoneAuthResponse) GetUseDdnsPrincipalSecurity() bool`

GetUseDdnsPrincipalSecurity returns the UseDdnsPrincipalSecurity field if non-nil, zero value otherwise.

### GetUseDdnsPrincipalSecurityOk

`func (o *GetZoneAuthResponse) GetUseDdnsPrincipalSecurityOk() (*bool, bool)`

GetUseDdnsPrincipalSecurityOk returns a tuple with the UseDdnsPrincipalSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsPrincipalSecurity

`func (o *GetZoneAuthResponse) SetUseDdnsPrincipalSecurity(v bool)`

SetUseDdnsPrincipalSecurity sets UseDdnsPrincipalSecurity field to given value.

### HasUseDdnsPrincipalSecurity

`func (o *GetZoneAuthResponse) HasUseDdnsPrincipalSecurity() bool`

HasUseDdnsPrincipalSecurity returns a boolean if a field has been set.

### GetUseDdnsRestrictProtected

`func (o *GetZoneAuthResponse) GetUseDdnsRestrictProtected() bool`

GetUseDdnsRestrictProtected returns the UseDdnsRestrictProtected field if non-nil, zero value otherwise.

### GetUseDdnsRestrictProtectedOk

`func (o *GetZoneAuthResponse) GetUseDdnsRestrictProtectedOk() (*bool, bool)`

GetUseDdnsRestrictProtectedOk returns a tuple with the UseDdnsRestrictProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsRestrictProtected

`func (o *GetZoneAuthResponse) SetUseDdnsRestrictProtected(v bool)`

SetUseDdnsRestrictProtected sets UseDdnsRestrictProtected field to given value.

### HasUseDdnsRestrictProtected

`func (o *GetZoneAuthResponse) HasUseDdnsRestrictProtected() bool`

HasUseDdnsRestrictProtected returns a boolean if a field has been set.

### GetUseDdnsRestrictStatic

`func (o *GetZoneAuthResponse) GetUseDdnsRestrictStatic() bool`

GetUseDdnsRestrictStatic returns the UseDdnsRestrictStatic field if non-nil, zero value otherwise.

### GetUseDdnsRestrictStaticOk

`func (o *GetZoneAuthResponse) GetUseDdnsRestrictStaticOk() (*bool, bool)`

GetUseDdnsRestrictStaticOk returns a tuple with the UseDdnsRestrictStatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDdnsRestrictStatic

`func (o *GetZoneAuthResponse) SetUseDdnsRestrictStatic(v bool)`

SetUseDdnsRestrictStatic sets UseDdnsRestrictStatic field to given value.

### HasUseDdnsRestrictStatic

`func (o *GetZoneAuthResponse) HasUseDdnsRestrictStatic() bool`

HasUseDdnsRestrictStatic returns a boolean if a field has been set.

### GetUseDnssecKeyParams

`func (o *GetZoneAuthResponse) GetUseDnssecKeyParams() bool`

GetUseDnssecKeyParams returns the UseDnssecKeyParams field if non-nil, zero value otherwise.

### GetUseDnssecKeyParamsOk

`func (o *GetZoneAuthResponse) GetUseDnssecKeyParamsOk() (*bool, bool)`

GetUseDnssecKeyParamsOk returns a tuple with the UseDnssecKeyParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnssecKeyParams

`func (o *GetZoneAuthResponse) SetUseDnssecKeyParams(v bool)`

SetUseDnssecKeyParams sets UseDnssecKeyParams field to given value.

### HasUseDnssecKeyParams

`func (o *GetZoneAuthResponse) HasUseDnssecKeyParams() bool`

HasUseDnssecKeyParams returns a boolean if a field has been set.

### GetUseExternalPrimary

`func (o *GetZoneAuthResponse) GetUseExternalPrimary() bool`

GetUseExternalPrimary returns the UseExternalPrimary field if non-nil, zero value otherwise.

### GetUseExternalPrimaryOk

`func (o *GetZoneAuthResponse) GetUseExternalPrimaryOk() (*bool, bool)`

GetUseExternalPrimaryOk returns a tuple with the UseExternalPrimary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseExternalPrimary

`func (o *GetZoneAuthResponse) SetUseExternalPrimary(v bool)`

SetUseExternalPrimary sets UseExternalPrimary field to given value.

### HasUseExternalPrimary

`func (o *GetZoneAuthResponse) HasUseExternalPrimary() bool`

HasUseExternalPrimary returns a boolean if a field has been set.

### GetUseGridZoneTimer

`func (o *GetZoneAuthResponse) GetUseGridZoneTimer() bool`

GetUseGridZoneTimer returns the UseGridZoneTimer field if non-nil, zero value otherwise.

### GetUseGridZoneTimerOk

`func (o *GetZoneAuthResponse) GetUseGridZoneTimerOk() (*bool, bool)`

GetUseGridZoneTimerOk returns a tuple with the UseGridZoneTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGridZoneTimer

`func (o *GetZoneAuthResponse) SetUseGridZoneTimer(v bool)`

SetUseGridZoneTimer sets UseGridZoneTimer field to given value.

### HasUseGridZoneTimer

`func (o *GetZoneAuthResponse) HasUseGridZoneTimer() bool`

HasUseGridZoneTimer returns a boolean if a field has been set.

### GetUseImportFrom

`func (o *GetZoneAuthResponse) GetUseImportFrom() bool`

GetUseImportFrom returns the UseImportFrom field if non-nil, zero value otherwise.

### GetUseImportFromOk

`func (o *GetZoneAuthResponse) GetUseImportFromOk() (*bool, bool)`

GetUseImportFromOk returns a tuple with the UseImportFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseImportFrom

`func (o *GetZoneAuthResponse) SetUseImportFrom(v bool)`

SetUseImportFrom sets UseImportFrom field to given value.

### HasUseImportFrom

`func (o *GetZoneAuthResponse) HasUseImportFrom() bool`

HasUseImportFrom returns a boolean if a field has been set.

### GetUseNotifyDelay

`func (o *GetZoneAuthResponse) GetUseNotifyDelay() bool`

GetUseNotifyDelay returns the UseNotifyDelay field if non-nil, zero value otherwise.

### GetUseNotifyDelayOk

`func (o *GetZoneAuthResponse) GetUseNotifyDelayOk() (*bool, bool)`

GetUseNotifyDelayOk returns a tuple with the UseNotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNotifyDelay

`func (o *GetZoneAuthResponse) SetUseNotifyDelay(v bool)`

SetUseNotifyDelay sets UseNotifyDelay field to given value.

### HasUseNotifyDelay

`func (o *GetZoneAuthResponse) HasUseNotifyDelay() bool`

HasUseNotifyDelay returns a boolean if a field has been set.

### GetUseRecordNamePolicy

`func (o *GetZoneAuthResponse) GetUseRecordNamePolicy() bool`

GetUseRecordNamePolicy returns the UseRecordNamePolicy field if non-nil, zero value otherwise.

### GetUseRecordNamePolicyOk

`func (o *GetZoneAuthResponse) GetUseRecordNamePolicyOk() (*bool, bool)`

GetUseRecordNamePolicyOk returns a tuple with the UseRecordNamePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRecordNamePolicy

`func (o *GetZoneAuthResponse) SetUseRecordNamePolicy(v bool)`

SetUseRecordNamePolicy sets UseRecordNamePolicy field to given value.

### HasUseRecordNamePolicy

`func (o *GetZoneAuthResponse) HasUseRecordNamePolicy() bool`

HasUseRecordNamePolicy returns a boolean if a field has been set.

### GetUseScavengingSettings

`func (o *GetZoneAuthResponse) GetUseScavengingSettings() bool`

GetUseScavengingSettings returns the UseScavengingSettings field if non-nil, zero value otherwise.

### GetUseScavengingSettingsOk

`func (o *GetZoneAuthResponse) GetUseScavengingSettingsOk() (*bool, bool)`

GetUseScavengingSettingsOk returns a tuple with the UseScavengingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseScavengingSettings

`func (o *GetZoneAuthResponse) SetUseScavengingSettings(v bool)`

SetUseScavengingSettings sets UseScavengingSettings field to given value.

### HasUseScavengingSettings

`func (o *GetZoneAuthResponse) HasUseScavengingSettings() bool`

HasUseScavengingSettings returns a boolean if a field has been set.

### GetUseSoaEmail

`func (o *GetZoneAuthResponse) GetUseSoaEmail() bool`

GetUseSoaEmail returns the UseSoaEmail field if non-nil, zero value otherwise.

### GetUseSoaEmailOk

`func (o *GetZoneAuthResponse) GetUseSoaEmailOk() (*bool, bool)`

GetUseSoaEmailOk returns a tuple with the UseSoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSoaEmail

`func (o *GetZoneAuthResponse) SetUseSoaEmail(v bool)`

SetUseSoaEmail sets UseSoaEmail field to given value.

### HasUseSoaEmail

`func (o *GetZoneAuthResponse) HasUseSoaEmail() bool`

HasUseSoaEmail returns a boolean if a field has been set.

### GetUsingSrgAssociations

`func (o *GetZoneAuthResponse) GetUsingSrgAssociations() bool`

GetUsingSrgAssociations returns the UsingSrgAssociations field if non-nil, zero value otherwise.

### GetUsingSrgAssociationsOk

`func (o *GetZoneAuthResponse) GetUsingSrgAssociationsOk() (*bool, bool)`

GetUsingSrgAssociationsOk returns a tuple with the UsingSrgAssociations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsingSrgAssociations

`func (o *GetZoneAuthResponse) SetUsingSrgAssociations(v bool)`

SetUsingSrgAssociations sets UsingSrgAssociations field to given value.

### HasUsingSrgAssociations

`func (o *GetZoneAuthResponse) HasUsingSrgAssociations() bool`

HasUsingSrgAssociations returns a boolean if a field has been set.

### GetUuid

`func (o *GetZoneAuthResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetZoneAuthResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetZoneAuthResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetZoneAuthResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetView

`func (o *GetZoneAuthResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetZoneAuthResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetZoneAuthResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetZoneAuthResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZoneFormat

`func (o *GetZoneAuthResponse) GetZoneFormat() string`

GetZoneFormat returns the ZoneFormat field if non-nil, zero value otherwise.

### GetZoneFormatOk

`func (o *GetZoneAuthResponse) GetZoneFormatOk() (*string, bool)`

GetZoneFormatOk returns a tuple with the ZoneFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFormat

`func (o *GetZoneAuthResponse) SetZoneFormat(v string)`

SetZoneFormat sets ZoneFormat field to given value.

### HasZoneFormat

`func (o *GetZoneAuthResponse) HasZoneFormat() bool`

HasZoneFormat returns a boolean if a field has been set.

### GetZoneNotQueriedEnabledTime

`func (o *GetZoneAuthResponse) GetZoneNotQueriedEnabledTime() int64`

GetZoneNotQueriedEnabledTime returns the ZoneNotQueriedEnabledTime field if non-nil, zero value otherwise.

### GetZoneNotQueriedEnabledTimeOk

`func (o *GetZoneAuthResponse) GetZoneNotQueriedEnabledTimeOk() (*int64, bool)`

GetZoneNotQueriedEnabledTimeOk returns a tuple with the ZoneNotQueriedEnabledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNotQueriedEnabledTime

`func (o *GetZoneAuthResponse) SetZoneNotQueriedEnabledTime(v int64)`

SetZoneNotQueriedEnabledTime sets ZoneNotQueriedEnabledTime field to given value.

### HasZoneNotQueriedEnabledTime

`func (o *GetZoneAuthResponse) HasZoneNotQueriedEnabledTime() bool`

HasZoneNotQueriedEnabledTime returns a boolean if a field has been set.

### GetResult

`func (o *GetZoneAuthResponse) GetResult() ZoneAuth`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetZoneAuthResponse) GetResultOk() (*ZoneAuth, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetZoneAuthResponse) SetResult(v ZoneAuth)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetZoneAuthResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


