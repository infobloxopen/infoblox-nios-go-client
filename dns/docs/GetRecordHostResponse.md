# GetRecordHostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Aliases** | Pointer to **[]string** | This is a list of aliases for the host. The aliases must be in FQDN format. This value can be in unicode format. | [optional] 
**AllowTelnet** | Pointer to **bool** | This field controls whether the credential is used for both the Telnet and SSH credentials. If set to False, the credential is used only for SSH. | [optional] 
**CliCredentials** | Pointer to [**[]RecordHostCliCredentials**](RecordHostCliCredentials.md) | The CLI credentials for the host record. | [optional] 
**CloudInfo** | Pointer to [**RecordHostCloudInfo**](RecordHostCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the record; maximum 256 characters. | [optional] 
**ConfigureForDns** | Pointer to **bool** | When configure_for_dns is false, the host does not have parent zone information. | [optional] 
**CreationTime** | Pointer to **int64** | The time of the record creation in Epoch seconds format. | [optional] [readonly] 
**DdnsProtected** | Pointer to **bool** | Determines if the DDNS updates for this record are allowed or not. | [optional] 
**DeviceDescription** | Pointer to **string** | The description of the device. | [optional] 
**DeviceLocation** | Pointer to **string** | The location of the device. | [optional] 
**DeviceType** | Pointer to **string** | The type of the device. | [optional] 
**DeviceVendor** | Pointer to **string** | The vendor of the device. | [optional] 
**Disable** | Pointer to **bool** | Determines if the record is disabled or not. False means that the record is enabled. | [optional] 
**DisableDiscovery** | Pointer to **bool** | Determines if the discovery for the record is disabled or not. False means that the discovery is enabled. | [optional] 
**DnsAliases** | Pointer to **[]string** | The list of aliases for the host in punycode format. | [optional] 
**DnsName** | Pointer to **string** | The name for a host record in punycode format. | [optional] [readonly] 
**EnableImmediateDiscovery** | Pointer to **bool** | Determines if the discovery for the record should be immediately enabled. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Ipv4addrs** | Pointer to **[]map[string]interface{}** | This is a list of IPv4 Addresses for the host. | [optional] 
**Ipv6addrs** | Pointer to **[]map[string]interface{}** | This is a list of IPv6 Addresses for the host. | [optional] 
**LastQueried** | Pointer to **int64** | The time of the last DNS query in Epoch seconds format. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**RecordHostMsAdUserData**](RecordHostMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | The host name in FQDN format This value can be in unicode format. Regular expression search is not supported for unicode values. | [optional] 
**NetworkView** | Pointer to **string** | The name of the network view in which the host record resides. | [optional] 
**RestartIfNeeded** | Pointer to **bool** | Restarts the member service. | [optional] 
**RrsetOrder** | Pointer to **string** | The value of this field specifies the order in which resource record sets are returned. The possible values are \&quot;cyclic\&quot;, \&quot;random\&quot; and \&quot;fixed\&quot;. | [optional] 
**Snmp3Credential** | Pointer to [**RecordHostSnmp3Credential**](RecordHostSnmp3Credential.md) |  | [optional] 
**SnmpCredential** | Pointer to [**RecordHostSnmpCredential**](RecordHostSnmpCredential.md) |  | [optional] 
**Ttl** | Pointer to **int64** | The Time To Live (TTL) value for record. A 32-bit unsigned integer that represents the duration, in seconds, for which the record is valid (cached). Zero indicates that the record should not be cached. | [optional] 
**UseCliCredentials** | Pointer to **bool** | If set to true, the CLI credential will override member-level settings. | [optional] 
**UseDnsEaInheritance** | Pointer to **bool** | When use_dns_ea_inheritance is True, the EA is inherited from associated zone. | [optional] 
**UseSnmp3Credential** | Pointer to **bool** | Determines if the SNMPv3 credential should be used for the record. | [optional] 
**UseSnmpCredential** | Pointer to **bool** | If set to true, the SNMP credential will override member-level settings. | [optional] 
**UseTtl** | Pointer to **bool** | Use flag for: ttl | [optional] 
**View** | Pointer to **string** | The name of the DNS view in which the record resides. Example: \&quot;external\&quot;. | [optional] 
**Zone** | Pointer to **string** | The name of the zone in which the record resides. Example: \&quot;zone.com\&quot;. If a view is not specified when searching by zone, the default view is used. | [optional] [readonly] 
**Result** | Pointer to [**RecordHost**](RecordHost.md) |  | [optional] 

## Methods

### NewGetRecordHostResponse

`func NewGetRecordHostResponse() *GetRecordHostResponse`

NewGetRecordHostResponse instantiates a new GetRecordHostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRecordHostResponseWithDefaults

`func NewGetRecordHostResponseWithDefaults() *GetRecordHostResponse`

NewGetRecordHostResponseWithDefaults instantiates a new GetRecordHostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetRecordHostResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetRecordHostResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetRecordHostResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetRecordHostResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAliases

`func (o *GetRecordHostResponse) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *GetRecordHostResponse) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *GetRecordHostResponse) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *GetRecordHostResponse) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAllowTelnet

`func (o *GetRecordHostResponse) GetAllowTelnet() bool`

GetAllowTelnet returns the AllowTelnet field if non-nil, zero value otherwise.

### GetAllowTelnetOk

`func (o *GetRecordHostResponse) GetAllowTelnetOk() (*bool, bool)`

GetAllowTelnetOk returns a tuple with the AllowTelnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTelnet

`func (o *GetRecordHostResponse) SetAllowTelnet(v bool)`

SetAllowTelnet sets AllowTelnet field to given value.

### HasAllowTelnet

`func (o *GetRecordHostResponse) HasAllowTelnet() bool`

HasAllowTelnet returns a boolean if a field has been set.

### GetCliCredentials

`func (o *GetRecordHostResponse) GetCliCredentials() []RecordHostCliCredentials`

GetCliCredentials returns the CliCredentials field if non-nil, zero value otherwise.

### GetCliCredentialsOk

`func (o *GetRecordHostResponse) GetCliCredentialsOk() (*[]RecordHostCliCredentials, bool)`

GetCliCredentialsOk returns a tuple with the CliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliCredentials

`func (o *GetRecordHostResponse) SetCliCredentials(v []RecordHostCliCredentials)`

SetCliCredentials sets CliCredentials field to given value.

### HasCliCredentials

`func (o *GetRecordHostResponse) HasCliCredentials() bool`

HasCliCredentials returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetRecordHostResponse) GetCloudInfo() RecordHostCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetRecordHostResponse) GetCloudInfoOk() (*RecordHostCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetRecordHostResponse) SetCloudInfo(v RecordHostCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetRecordHostResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetRecordHostResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetRecordHostResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetRecordHostResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetRecordHostResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetConfigureForDns

`func (o *GetRecordHostResponse) GetConfigureForDns() bool`

GetConfigureForDns returns the ConfigureForDns field if non-nil, zero value otherwise.

### GetConfigureForDnsOk

`func (o *GetRecordHostResponse) GetConfigureForDnsOk() (*bool, bool)`

GetConfigureForDnsOk returns a tuple with the ConfigureForDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigureForDns

`func (o *GetRecordHostResponse) SetConfigureForDns(v bool)`

SetConfigureForDns sets ConfigureForDns field to given value.

### HasConfigureForDns

`func (o *GetRecordHostResponse) HasConfigureForDns() bool`

HasConfigureForDns returns a boolean if a field has been set.

### GetCreationTime

`func (o *GetRecordHostResponse) GetCreationTime() int64`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *GetRecordHostResponse) GetCreationTimeOk() (*int64, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *GetRecordHostResponse) SetCreationTime(v int64)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *GetRecordHostResponse) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDdnsProtected

`func (o *GetRecordHostResponse) GetDdnsProtected() bool`

GetDdnsProtected returns the DdnsProtected field if non-nil, zero value otherwise.

### GetDdnsProtectedOk

`func (o *GetRecordHostResponse) GetDdnsProtectedOk() (*bool, bool)`

GetDdnsProtectedOk returns a tuple with the DdnsProtected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsProtected

`func (o *GetRecordHostResponse) SetDdnsProtected(v bool)`

SetDdnsProtected sets DdnsProtected field to given value.

### HasDdnsProtected

`func (o *GetRecordHostResponse) HasDdnsProtected() bool`

HasDdnsProtected returns a boolean if a field has been set.

### GetDeviceDescription

`func (o *GetRecordHostResponse) GetDeviceDescription() string`

GetDeviceDescription returns the DeviceDescription field if non-nil, zero value otherwise.

### GetDeviceDescriptionOk

`func (o *GetRecordHostResponse) GetDeviceDescriptionOk() (*string, bool)`

GetDeviceDescriptionOk returns a tuple with the DeviceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDescription

`func (o *GetRecordHostResponse) SetDeviceDescription(v string)`

SetDeviceDescription sets DeviceDescription field to given value.

### HasDeviceDescription

`func (o *GetRecordHostResponse) HasDeviceDescription() bool`

HasDeviceDescription returns a boolean if a field has been set.

### GetDeviceLocation

`func (o *GetRecordHostResponse) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *GetRecordHostResponse) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *GetRecordHostResponse) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.

### HasDeviceLocation

`func (o *GetRecordHostResponse) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### GetDeviceType

`func (o *GetRecordHostResponse) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *GetRecordHostResponse) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *GetRecordHostResponse) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *GetRecordHostResponse) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *GetRecordHostResponse) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *GetRecordHostResponse) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *GetRecordHostResponse) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *GetRecordHostResponse) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDisable

`func (o *GetRecordHostResponse) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *GetRecordHostResponse) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *GetRecordHostResponse) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *GetRecordHostResponse) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDisableDiscovery

`func (o *GetRecordHostResponse) GetDisableDiscovery() bool`

GetDisableDiscovery returns the DisableDiscovery field if non-nil, zero value otherwise.

### GetDisableDiscoveryOk

`func (o *GetRecordHostResponse) GetDisableDiscoveryOk() (*bool, bool)`

GetDisableDiscoveryOk returns a tuple with the DisableDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDiscovery

`func (o *GetRecordHostResponse) SetDisableDiscovery(v bool)`

SetDisableDiscovery sets DisableDiscovery field to given value.

### HasDisableDiscovery

`func (o *GetRecordHostResponse) HasDisableDiscovery() bool`

HasDisableDiscovery returns a boolean if a field has been set.

### GetDnsAliases

`func (o *GetRecordHostResponse) GetDnsAliases() []string`

GetDnsAliases returns the DnsAliases field if non-nil, zero value otherwise.

### GetDnsAliasesOk

`func (o *GetRecordHostResponse) GetDnsAliasesOk() (*[]string, bool)`

GetDnsAliasesOk returns a tuple with the DnsAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsAliases

`func (o *GetRecordHostResponse) SetDnsAliases(v []string)`

SetDnsAliases sets DnsAliases field to given value.

### HasDnsAliases

`func (o *GetRecordHostResponse) HasDnsAliases() bool`

HasDnsAliases returns a boolean if a field has been set.

### GetDnsName

`func (o *GetRecordHostResponse) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *GetRecordHostResponse) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *GetRecordHostResponse) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *GetRecordHostResponse) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetEnableImmediateDiscovery

`func (o *GetRecordHostResponse) GetEnableImmediateDiscovery() bool`

GetEnableImmediateDiscovery returns the EnableImmediateDiscovery field if non-nil, zero value otherwise.

### GetEnableImmediateDiscoveryOk

`func (o *GetRecordHostResponse) GetEnableImmediateDiscoveryOk() (*bool, bool)`

GetEnableImmediateDiscoveryOk returns a tuple with the EnableImmediateDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableImmediateDiscovery

`func (o *GetRecordHostResponse) SetEnableImmediateDiscovery(v bool)`

SetEnableImmediateDiscovery sets EnableImmediateDiscovery field to given value.

### HasEnableImmediateDiscovery

`func (o *GetRecordHostResponse) HasEnableImmediateDiscovery() bool`

HasEnableImmediateDiscovery returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetRecordHostResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetRecordHostResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetRecordHostResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetRecordHostResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIpv4addrs

`func (o *GetRecordHostResponse) GetIpv4addrs() []map[string]interface{}`

GetIpv4addrs returns the Ipv4addrs field if non-nil, zero value otherwise.

### GetIpv4addrsOk

`func (o *GetRecordHostResponse) GetIpv4addrsOk() (*[]map[string]interface{}, bool)`

GetIpv4addrsOk returns a tuple with the Ipv4addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4addrs

`func (o *GetRecordHostResponse) SetIpv4addrs(v []map[string]interface{})`

SetIpv4addrs sets Ipv4addrs field to given value.

### HasIpv4addrs

`func (o *GetRecordHostResponse) HasIpv4addrs() bool`

HasIpv4addrs returns a boolean if a field has been set.

### GetIpv6addrs

`func (o *GetRecordHostResponse) GetIpv6addrs() []map[string]interface{}`

GetIpv6addrs returns the Ipv6addrs field if non-nil, zero value otherwise.

### GetIpv6addrsOk

`func (o *GetRecordHostResponse) GetIpv6addrsOk() (*[]map[string]interface{}, bool)`

GetIpv6addrsOk returns a tuple with the Ipv6addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6addrs

`func (o *GetRecordHostResponse) SetIpv6addrs(v []map[string]interface{})`

SetIpv6addrs sets Ipv6addrs field to given value.

### HasIpv6addrs

`func (o *GetRecordHostResponse) HasIpv6addrs() bool`

HasIpv6addrs returns a boolean if a field has been set.

### GetLastQueried

`func (o *GetRecordHostResponse) GetLastQueried() int64`

GetLastQueried returns the LastQueried field if non-nil, zero value otherwise.

### GetLastQueriedOk

`func (o *GetRecordHostResponse) GetLastQueriedOk() (*int64, bool)`

GetLastQueriedOk returns a tuple with the LastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastQueried

`func (o *GetRecordHostResponse) SetLastQueried(v int64)`

SetLastQueried sets LastQueried field to given value.

### HasLastQueried

`func (o *GetRecordHostResponse) HasLastQueried() bool`

HasLastQueried returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetRecordHostResponse) GetMsAdUserData() RecordHostMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetRecordHostResponse) GetMsAdUserDataOk() (*RecordHostMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetRecordHostResponse) SetMsAdUserData(v RecordHostMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetRecordHostResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *GetRecordHostResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRecordHostResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRecordHostResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRecordHostResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetRecordHostResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetRecordHostResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetRecordHostResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetRecordHostResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetRestartIfNeeded

`func (o *GetRecordHostResponse) GetRestartIfNeeded() bool`

GetRestartIfNeeded returns the RestartIfNeeded field if non-nil, zero value otherwise.

### GetRestartIfNeededOk

`func (o *GetRecordHostResponse) GetRestartIfNeededOk() (*bool, bool)`

GetRestartIfNeededOk returns a tuple with the RestartIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartIfNeeded

`func (o *GetRecordHostResponse) SetRestartIfNeeded(v bool)`

SetRestartIfNeeded sets RestartIfNeeded field to given value.

### HasRestartIfNeeded

`func (o *GetRecordHostResponse) HasRestartIfNeeded() bool`

HasRestartIfNeeded returns a boolean if a field has been set.

### GetRrsetOrder

`func (o *GetRecordHostResponse) GetRrsetOrder() string`

GetRrsetOrder returns the RrsetOrder field if non-nil, zero value otherwise.

### GetRrsetOrderOk

`func (o *GetRecordHostResponse) GetRrsetOrderOk() (*string, bool)`

GetRrsetOrderOk returns a tuple with the RrsetOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrsetOrder

`func (o *GetRecordHostResponse) SetRrsetOrder(v string)`

SetRrsetOrder sets RrsetOrder field to given value.

### HasRrsetOrder

`func (o *GetRecordHostResponse) HasRrsetOrder() bool`

HasRrsetOrder returns a boolean if a field has been set.

### GetSnmp3Credential

`func (o *GetRecordHostResponse) GetSnmp3Credential() RecordHostSnmp3Credential`

GetSnmp3Credential returns the Snmp3Credential field if non-nil, zero value otherwise.

### GetSnmp3CredentialOk

`func (o *GetRecordHostResponse) GetSnmp3CredentialOk() (*RecordHostSnmp3Credential, bool)`

GetSnmp3CredentialOk returns a tuple with the Snmp3Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp3Credential

`func (o *GetRecordHostResponse) SetSnmp3Credential(v RecordHostSnmp3Credential)`

SetSnmp3Credential sets Snmp3Credential field to given value.

### HasSnmp3Credential

`func (o *GetRecordHostResponse) HasSnmp3Credential() bool`

HasSnmp3Credential returns a boolean if a field has been set.

### GetSnmpCredential

`func (o *GetRecordHostResponse) GetSnmpCredential() RecordHostSnmpCredential`

GetSnmpCredential returns the SnmpCredential field if non-nil, zero value otherwise.

### GetSnmpCredentialOk

`func (o *GetRecordHostResponse) GetSnmpCredentialOk() (*RecordHostSnmpCredential, bool)`

GetSnmpCredentialOk returns a tuple with the SnmpCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmpCredential

`func (o *GetRecordHostResponse) SetSnmpCredential(v RecordHostSnmpCredential)`

SetSnmpCredential sets SnmpCredential field to given value.

### HasSnmpCredential

`func (o *GetRecordHostResponse) HasSnmpCredential() bool`

HasSnmpCredential returns a boolean if a field has been set.

### GetTtl

`func (o *GetRecordHostResponse) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GetRecordHostResponse) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GetRecordHostResponse) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GetRecordHostResponse) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUseCliCredentials

`func (o *GetRecordHostResponse) GetUseCliCredentials() bool`

GetUseCliCredentials returns the UseCliCredentials field if non-nil, zero value otherwise.

### GetUseCliCredentialsOk

`func (o *GetRecordHostResponse) GetUseCliCredentialsOk() (*bool, bool)`

GetUseCliCredentialsOk returns a tuple with the UseCliCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCliCredentials

`func (o *GetRecordHostResponse) SetUseCliCredentials(v bool)`

SetUseCliCredentials sets UseCliCredentials field to given value.

### HasUseCliCredentials

`func (o *GetRecordHostResponse) HasUseCliCredentials() bool`

HasUseCliCredentials returns a boolean if a field has been set.

### GetUseDnsEaInheritance

`func (o *GetRecordHostResponse) GetUseDnsEaInheritance() bool`

GetUseDnsEaInheritance returns the UseDnsEaInheritance field if non-nil, zero value otherwise.

### GetUseDnsEaInheritanceOk

`func (o *GetRecordHostResponse) GetUseDnsEaInheritanceOk() (*bool, bool)`

GetUseDnsEaInheritanceOk returns a tuple with the UseDnsEaInheritance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseDnsEaInheritance

`func (o *GetRecordHostResponse) SetUseDnsEaInheritance(v bool)`

SetUseDnsEaInheritance sets UseDnsEaInheritance field to given value.

### HasUseDnsEaInheritance

`func (o *GetRecordHostResponse) HasUseDnsEaInheritance() bool`

HasUseDnsEaInheritance returns a boolean if a field has been set.

### GetUseSnmp3Credential

`func (o *GetRecordHostResponse) GetUseSnmp3Credential() bool`

GetUseSnmp3Credential returns the UseSnmp3Credential field if non-nil, zero value otherwise.

### GetUseSnmp3CredentialOk

`func (o *GetRecordHostResponse) GetUseSnmp3CredentialOk() (*bool, bool)`

GetUseSnmp3CredentialOk returns a tuple with the UseSnmp3Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmp3Credential

`func (o *GetRecordHostResponse) SetUseSnmp3Credential(v bool)`

SetUseSnmp3Credential sets UseSnmp3Credential field to given value.

### HasUseSnmp3Credential

`func (o *GetRecordHostResponse) HasUseSnmp3Credential() bool`

HasUseSnmp3Credential returns a boolean if a field has been set.

### GetUseSnmpCredential

`func (o *GetRecordHostResponse) GetUseSnmpCredential() bool`

GetUseSnmpCredential returns the UseSnmpCredential field if non-nil, zero value otherwise.

### GetUseSnmpCredentialOk

`func (o *GetRecordHostResponse) GetUseSnmpCredentialOk() (*bool, bool)`

GetUseSnmpCredentialOk returns a tuple with the UseSnmpCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSnmpCredential

`func (o *GetRecordHostResponse) SetUseSnmpCredential(v bool)`

SetUseSnmpCredential sets UseSnmpCredential field to given value.

### HasUseSnmpCredential

`func (o *GetRecordHostResponse) HasUseSnmpCredential() bool`

HasUseSnmpCredential returns a boolean if a field has been set.

### GetUseTtl

`func (o *GetRecordHostResponse) GetUseTtl() bool`

GetUseTtl returns the UseTtl field if non-nil, zero value otherwise.

### GetUseTtlOk

`func (o *GetRecordHostResponse) GetUseTtlOk() (*bool, bool)`

GetUseTtlOk returns a tuple with the UseTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTtl

`func (o *GetRecordHostResponse) SetUseTtl(v bool)`

SetUseTtl sets UseTtl field to given value.

### HasUseTtl

`func (o *GetRecordHostResponse) HasUseTtl() bool`

HasUseTtl returns a boolean if a field has been set.

### GetView

`func (o *GetRecordHostResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetRecordHostResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetRecordHostResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetRecordHostResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetZone

`func (o *GetRecordHostResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetRecordHostResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetRecordHostResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetRecordHostResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetRecordHostResponse) GetResult() RecordHost`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetRecordHostResponse) GetResultOk() (*RecordHost, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetRecordHostResponse) SetResult(v RecordHost)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetRecordHostResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


