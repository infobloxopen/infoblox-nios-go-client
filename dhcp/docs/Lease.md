# Lease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The IPv4 Address or IPv6 Address of the lease. | [optional] [readonly] 
**BillingClass** | Pointer to **string** | The billing_class value of a DHCP Lease object. This field specifies the class to which this lease is currently billed. This field is for IPv4 leases only. | [optional] [readonly] 
**BindingState** | Pointer to **string** | The binding state for the current lease. Following are some of the values this field can be set to: * ABANDONED: The Infoblox appliance cannot lease this IP address because the appliance received a response when it pinged the address. * ACTIVE: The lease is currently in use by a DHCP client. * EXPIRED: The lease was in use, but the DHCP client never renewed it, so it is no longer valid. * FREE: The lease is available for clients to use. * RELEASED: The DHCP client returned the lease to the appliance. | [optional] [readonly] 
**ClientHostname** | Pointer to **string** | The client_hostname of a DHCP Lease object. This field specifies the host name that the DHCP client sends to the Infoblox appliance using DHCP option 12. | [optional] [readonly] 
**Cltt** | Pointer to **int64** | The CLTT (Client Last Transaction Time) value of a DHCP Lease object. This field specifies the time of the last transaction with the DHCP client for this lease. | [optional] [readonly] 
**DiscoveredData** | Pointer to [**LeaseDiscoveredData**](LeaseDiscoveredData.md) |  | [optional] 
**Ends** | Pointer to **int64** | The end time value of a DHCP Lease object. This field specifies the time when a lease ended. | [optional] [readonly] 
**Fingerprint** | Pointer to **string** | DHCP fingerprint for the lease. | [optional] [readonly] 
**Hardware** | Pointer to **string** | The hardware type of a DHCP Lease object. This field specifies the MAC address of the network interface on which the lease will be used. This field is supported for IPv4 leases, and from NIOS-9.0.6 onwards, also supported for IPv6 leases. | [optional] [readonly] 
**Ipv6Duid** | Pointer to **string** | The DUID value for this lease. This field is only applicable for IPv6 leases. | [optional] [readonly] 
**Ipv6Iaid** | Pointer to **string** | The interface ID of an IPv6 address that the Infoblox appliance leased to the DHCP client. This field is for IPv6 leases only. | [optional] [readonly] 
**Ipv6PreferredLifetime** | Pointer to **int64** | The preferred lifetime value of an IPv6 address that the Infoblox appliance leased to the DHCP client. This field is for IPv6 leases only. | [optional] [readonly] 
**Ipv6PrefixBits** | Pointer to **int64** | Prefix bits for this lease. This field is for IPv6 leases only. | [optional] [readonly] 
**IsInvalidMac** | Pointer to **bool** | This flag reflects whether the MAC address for this lease is invalid. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**LeaseMsAdUserData**](LeaseMsAdUserData.md) |  | [optional] 
**Network** | Pointer to **string** | The network, in \&quot;network/netmask\&quot; format, with which this lease is associated. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this lease resides. | [optional] [readonly] 
**NeverEnds** | Pointer to **bool** | If this field is set to True, the lease does not have an end time. | [optional] [readonly] 
**NeverStarts** | Pointer to **bool** | If this field is set to True, the lease does not have a start time. | [optional] [readonly] 
**NextBindingState** | Pointer to **string** | The subsequent binding state when the current lease expires. This field is for IPv4 leases only. Following are some of the values this field can be set to: * ABANDONED: The Infoblox appliance cannot lease this IP address because the appliance received a response when it pinged the address. * ACTIVE: The lease is currently in use by a DHCP client. * EXPIRED: The lease was in use, but the DHCP client never renewed it, so it is no longer valid. * FREE: The lease is available for clients to use. * RELEASED: The DHCP client returned the lease to the appliance. | [optional] [readonly] 
**OnCommit** | Pointer to **string** | The list of commands to be executed when the lease is granted. | [optional] [readonly] 
**OnExpiry** | Pointer to **string** | The list of commands to be executed when the lease expires. | [optional] [readonly] 
**OnRelease** | Pointer to **string** | The list of commands to be executed when the lease is released. | [optional] [readonly] 
**Option** | Pointer to **string** | The option value of a DHCP Lease object. This field specifies the agent circuit ID and remote ID sent by a DHCP relay agent in DHCP option 82. This field is for IPv4 leases only. | [optional] [readonly] 
**Protocol** | Pointer to **string** | This field determines whether the lease is an IPv4 or IPv6 address. | [optional] [readonly] 
**RemoteId** | Pointer to **string** | This field represents the \&quot;Remote ID\&quot; sub-option of DHCP option 82. Remote ID can be in ASCII form (e.g. &#x60;&#x60;\&quot;abcd\&quot;&#x60;&#x60;) or in colon-separated HEX form (e.g. &#x60;&#x60;1:2:ab:cd&#x60;&#x60;). HEX representation is used only when the sub-option value contains unprintable characters. If a remote ID sub-option value is in ASCII form, it is always enclosed in quotes to prevent ambiguous values (e.g. &#x60;&#x60;\&quot;10:20\&quot;&#x60;&#x60; - ASCII 5-byte string; &#x60;&#x60;10:20&#x60;&#x60; - HEX 2-byte value). * ASCII representation is used if the remote ID sub-option contains only printable ASCII characters (ASCII characters in range &#x60;&#x60;x20-0x7E&#x60;&#x60;). * The backslash symbol (&#x60;&#x60;\\&#x60;&#x60;) is used as an escape symbol to escape the quote symbol (&#x60;&#x60;\&quot;&#x60;&#x60;) in an ASCII string. * Double backslashes (&#x60;&#x60;\\\\&#x60;&#x60;) are used to represent the backslash symbol (&#x60;&#x60;\\&#x60;&#x60;) in an ASCII string. * HEX representation is used only when the remote ID sub-option value contains unprintable characters and is normalized as follows: * starting zero is removed from digits: &#x60;&#x60;1&#x60;&#x60;, &#x60;&#x60;a&#x60;&#x60; - Valid; &#x60;&#x60;01&#x60;&#x60;, &#x60;&#x60;0a&#x60;&#x60; - Invalid; * lowercase characters are used for symbols: &#x60;&#x60;fa&#x60;&#x60; - Valid; &#x60;&#x60;FA&#x60;&#x60; - Invalid. NIOS does not support the convertion between HEX and ASCII formats. Searches are performed using the exact same format and value as the sub-option is represented. Query examples assume the following leases are stored in the database: .. tabularcolumns:: |p{1in}|p{3in}|p{2in}| &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; Number Option field Extracted remote ID field &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; Lease01 agent.remote-id&#x3D; \&quot;00152654358700\&quot; \&quot;00152654358700\&quot; agent.circuit-id&#x3D; \&quot;BX1-PORT-003\&quot; Lease02 agent.remote-id&#x3D;\&quot;Dhcp \&quot;Dhcp Relay 10\&quot; Relay 10\&quot; agent.circuit-id&#x3D;\&quot;Port008\&quot; Lease03 agent.remote-id&#x3D;\&quot;00:01:02\&quot; \&quot;00:01:02\&quot; Lease04 agent.remote-id&#x3D;0:1:2 0:1:2 Lease05 agent.remote-id&#x3D;02:03 2:3 Lease06 agent.remote-id&#x3D;10:20 10:20 Lease07 agent.circuit-id&#x3D; \&quot;no-remote-id\&quot; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; Expected results: .. tabularcolumns:: |p{1.5in}|p{1.5in}|p{3in}| &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; Query Returned leases Comments &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; remote_id&#x3D;01:02 None EXACT query. No results are expected. remote_id&#x3D;\&quot;Dhcp Relay 10\&quot; Lease02 EXACT query for an ASCII value. remote_id&#x3D;0:1:2 Lease04 EXACT query for a HEX value. remote_id&#x3D;00:01:02 None EXACT query for a HEX value. No results are expected as the search value is not normalized to the same format used in the database. remote_id~&#x3D;10 Lease02, Lease06 REGEX query. remote_id~&#x3D;^\&quot;.*1 Lease01, Lease02, REGEX query. Only ASCII Lease03 values are expected due to the starting quote (&#x60;&#x60;\&quot;&#x60;&#x60;) in the search value. remote_id~&#x3D;^[^\&quot;]*2 Lease04, Lease05, REGEX query. Only HEX values Lease06 are expected as the starting quote (&#x60;&#x60;\&quot;&#x60;&#x60;) is excluded from the search value. remote_id&#x3D;\&quot;\&quot; None EXACT query. No results are expected as no leases that contain an empty remote ID value exist in the database. ID value in the database. remote_id~&#x3D;\&quot;\&quot; Lease01, Lease02, REGEX query. This query is Lease03, Lease04, expected to match any Lease05, Lease06 lease that contain remote ID set to any value. &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; &#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D;&#x3D; **NOTE:** Lease07 is not expected to be returned when searching for the remote ID sub-option. | [optional] [readonly] 
**RequestedOptions** | Pointer to **string** | This field contains the option request list received from the client. For DHCPv4, it includes \&quot;Parameter Request List\&quot; data and for DHCPv6, it includes \&quot;Option Request Option\&quot; data. | [optional] [readonly] 
**ServedBy** | Pointer to **string** | The IP address of the server that sends an active lease to a client. | [optional] [readonly] 
**ServerHostName** | Pointer to **string** | The host name of the Grid member or Microsoft DHCP server that issues the lease. | [optional] [readonly] 
**Starts** | Pointer to **int64** | The start time of a DHCP Lease object. This field specifies the time when the lease starts. | [optional] [readonly] 
**Tsfp** | Pointer to **int64** | The TSFP (Time Sent From Partner) value of a DHCP Lease object. This field specifies the time that the current lease state ends, from the point of view of a remote DHCP failover peer. This field is for IPv4 leases only. | [optional] [readonly] 
**Tstp** | Pointer to **int64** | The TSTP (Time Sent To Partner) value of a DHCP Lease object. This field specifies the time that the current lease state ends, from the point of view of a local DHCP failover peer. This field is for IPv4 leases only. | [optional] [readonly] 
**Uid** | Pointer to **string** | The UID (User ID) value of a DHCP Lease object. This field specifies the client identifier that the DHCP client sends the Infoblox appliance (in DHCP option 61) when it acquires the lease. Not all DHCP clients send a UID. This field is for IPv4 leases only. | [optional] [readonly] 
**Username** | Pointer to **string** | The user name that the server has associated with a DHCP Lease object. | [optional] [readonly] 
**Variable** | Pointer to **string** | The variable value of a DHCP Lease object. This field keeps all variables related to the DDNS update of the DHCP lease. The variables related to the DDNS updates of the DHCP lease. The variables can be one of the following: ddns-text: The ddns-text variable is used to record the value of the client&#39;s TXT identification record when the interim DDNS update style has been used to update the DNS service for a particular lease. ddns-fwd-name: When a DDNS update was successfully completed, the ddns-fwd-name variable records the value of the name used when the client&#39;s A record was updated. The server may have used this name when it updated the client&#39;s PTR record. ddns-client-fqdn: If the server is configured to use the interim DDNS update style and is also configured to allow clients to update their own FQDNs, the ddns-client-fqdn variable records the name that the client used when it updated its own FQDN. This is also the name that the server used to update the client&#39;s PTR record. ddns-rev-name: If the server successfully updates the client&#39;s PTR record, this variable will record the name that the DHCP server used for the PTR record. The name to which the PTR record points will be either the ddns-fwd-name or the ddns-client-fqdn. | [optional] [readonly] 

## Methods

### NewLease

`func NewLease() *Lease`

NewLease instantiates a new Lease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaseWithDefaults

`func NewLeaseWithDefaults() *Lease`

NewLeaseWithDefaults instantiates a new Lease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Lease) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Lease) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Lease) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Lease) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *Lease) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Lease) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Lease) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *Lease) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBillingClass

`func (o *Lease) GetBillingClass() string`

GetBillingClass returns the BillingClass field if non-nil, zero value otherwise.

### GetBillingClassOk

`func (o *Lease) GetBillingClassOk() (*string, bool)`

GetBillingClassOk returns a tuple with the BillingClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingClass

`func (o *Lease) SetBillingClass(v string)`

SetBillingClass sets BillingClass field to given value.

### HasBillingClass

`func (o *Lease) HasBillingClass() bool`

HasBillingClass returns a boolean if a field has been set.

### GetBindingState

`func (o *Lease) GetBindingState() string`

GetBindingState returns the BindingState field if non-nil, zero value otherwise.

### GetBindingStateOk

`func (o *Lease) GetBindingStateOk() (*string, bool)`

GetBindingStateOk returns a tuple with the BindingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingState

`func (o *Lease) SetBindingState(v string)`

SetBindingState sets BindingState field to given value.

### HasBindingState

`func (o *Lease) HasBindingState() bool`

HasBindingState returns a boolean if a field has been set.

### GetClientHostname

`func (o *Lease) GetClientHostname() string`

GetClientHostname returns the ClientHostname field if non-nil, zero value otherwise.

### GetClientHostnameOk

`func (o *Lease) GetClientHostnameOk() (*string, bool)`

GetClientHostnameOk returns a tuple with the ClientHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHostname

`func (o *Lease) SetClientHostname(v string)`

SetClientHostname sets ClientHostname field to given value.

### HasClientHostname

`func (o *Lease) HasClientHostname() bool`

HasClientHostname returns a boolean if a field has been set.

### GetCltt

`func (o *Lease) GetCltt() int64`

GetCltt returns the Cltt field if non-nil, zero value otherwise.

### GetClttOk

`func (o *Lease) GetClttOk() (*int64, bool)`

GetClttOk returns a tuple with the Cltt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCltt

`func (o *Lease) SetCltt(v int64)`

SetCltt sets Cltt field to given value.

### HasCltt

`func (o *Lease) HasCltt() bool`

HasCltt returns a boolean if a field has been set.

### GetDiscoveredData

`func (o *Lease) GetDiscoveredData() LeaseDiscoveredData`

GetDiscoveredData returns the DiscoveredData field if non-nil, zero value otherwise.

### GetDiscoveredDataOk

`func (o *Lease) GetDiscoveredDataOk() (*LeaseDiscoveredData, bool)`

GetDiscoveredDataOk returns a tuple with the DiscoveredData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredData

`func (o *Lease) SetDiscoveredData(v LeaseDiscoveredData)`

SetDiscoveredData sets DiscoveredData field to given value.

### HasDiscoveredData

`func (o *Lease) HasDiscoveredData() bool`

HasDiscoveredData returns a boolean if a field has been set.

### GetEnds

`func (o *Lease) GetEnds() int64`

GetEnds returns the Ends field if non-nil, zero value otherwise.

### GetEndsOk

`func (o *Lease) GetEndsOk() (*int64, bool)`

GetEndsOk returns a tuple with the Ends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnds

`func (o *Lease) SetEnds(v int64)`

SetEnds sets Ends field to given value.

### HasEnds

`func (o *Lease) HasEnds() bool`

HasEnds returns a boolean if a field has been set.

### GetFingerprint

`func (o *Lease) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Lease) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Lease) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Lease) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetHardware

`func (o *Lease) GetHardware() string`

GetHardware returns the Hardware field if non-nil, zero value otherwise.

### GetHardwareOk

`func (o *Lease) GetHardwareOk() (*string, bool)`

GetHardwareOk returns a tuple with the Hardware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardware

`func (o *Lease) SetHardware(v string)`

SetHardware sets Hardware field to given value.

### HasHardware

`func (o *Lease) HasHardware() bool`

HasHardware returns a boolean if a field has been set.

### GetIpv6Duid

`func (o *Lease) GetIpv6Duid() string`

GetIpv6Duid returns the Ipv6Duid field if non-nil, zero value otherwise.

### GetIpv6DuidOk

`func (o *Lease) GetIpv6DuidOk() (*string, bool)`

GetIpv6DuidOk returns a tuple with the Ipv6Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Duid

`func (o *Lease) SetIpv6Duid(v string)`

SetIpv6Duid sets Ipv6Duid field to given value.

### HasIpv6Duid

`func (o *Lease) HasIpv6Duid() bool`

HasIpv6Duid returns a boolean if a field has been set.

### GetIpv6Iaid

`func (o *Lease) GetIpv6Iaid() string`

GetIpv6Iaid returns the Ipv6Iaid field if non-nil, zero value otherwise.

### GetIpv6IaidOk

`func (o *Lease) GetIpv6IaidOk() (*string, bool)`

GetIpv6IaidOk returns a tuple with the Ipv6Iaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Iaid

`func (o *Lease) SetIpv6Iaid(v string)`

SetIpv6Iaid sets Ipv6Iaid field to given value.

### HasIpv6Iaid

`func (o *Lease) HasIpv6Iaid() bool`

HasIpv6Iaid returns a boolean if a field has been set.

### GetIpv6PreferredLifetime

`func (o *Lease) GetIpv6PreferredLifetime() int64`

GetIpv6PreferredLifetime returns the Ipv6PreferredLifetime field if non-nil, zero value otherwise.

### GetIpv6PreferredLifetimeOk

`func (o *Lease) GetIpv6PreferredLifetimeOk() (*int64, bool)`

GetIpv6PreferredLifetimeOk returns a tuple with the Ipv6PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PreferredLifetime

`func (o *Lease) SetIpv6PreferredLifetime(v int64)`

SetIpv6PreferredLifetime sets Ipv6PreferredLifetime field to given value.

### HasIpv6PreferredLifetime

`func (o *Lease) HasIpv6PreferredLifetime() bool`

HasIpv6PreferredLifetime returns a boolean if a field has been set.

### GetIpv6PrefixBits

`func (o *Lease) GetIpv6PrefixBits() int64`

GetIpv6PrefixBits returns the Ipv6PrefixBits field if non-nil, zero value otherwise.

### GetIpv6PrefixBitsOk

`func (o *Lease) GetIpv6PrefixBitsOk() (*int64, bool)`

GetIpv6PrefixBitsOk returns a tuple with the Ipv6PrefixBits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6PrefixBits

`func (o *Lease) SetIpv6PrefixBits(v int64)`

SetIpv6PrefixBits sets Ipv6PrefixBits field to given value.

### HasIpv6PrefixBits

`func (o *Lease) HasIpv6PrefixBits() bool`

HasIpv6PrefixBits returns a boolean if a field has been set.

### GetIsInvalidMac

`func (o *Lease) GetIsInvalidMac() bool`

GetIsInvalidMac returns the IsInvalidMac field if non-nil, zero value otherwise.

### GetIsInvalidMacOk

`func (o *Lease) GetIsInvalidMacOk() (*bool, bool)`

GetIsInvalidMacOk returns a tuple with the IsInvalidMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvalidMac

`func (o *Lease) SetIsInvalidMac(v bool)`

SetIsInvalidMac sets IsInvalidMac field to given value.

### HasIsInvalidMac

`func (o *Lease) HasIsInvalidMac() bool`

HasIsInvalidMac returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *Lease) GetMsAdUserData() LeaseMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *Lease) GetMsAdUserDataOk() (*LeaseMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *Lease) SetMsAdUserData(v LeaseMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *Lease) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *Lease) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Lease) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Lease) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Lease) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *Lease) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *Lease) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *Lease) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *Lease) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNeverEnds

`func (o *Lease) GetNeverEnds() bool`

GetNeverEnds returns the NeverEnds field if non-nil, zero value otherwise.

### GetNeverEndsOk

`func (o *Lease) GetNeverEndsOk() (*bool, bool)`

GetNeverEndsOk returns a tuple with the NeverEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverEnds

`func (o *Lease) SetNeverEnds(v bool)`

SetNeverEnds sets NeverEnds field to given value.

### HasNeverEnds

`func (o *Lease) HasNeverEnds() bool`

HasNeverEnds returns a boolean if a field has been set.

### GetNeverStarts

`func (o *Lease) GetNeverStarts() bool`

GetNeverStarts returns the NeverStarts field if non-nil, zero value otherwise.

### GetNeverStartsOk

`func (o *Lease) GetNeverStartsOk() (*bool, bool)`

GetNeverStartsOk returns a tuple with the NeverStarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverStarts

`func (o *Lease) SetNeverStarts(v bool)`

SetNeverStarts sets NeverStarts field to given value.

### HasNeverStarts

`func (o *Lease) HasNeverStarts() bool`

HasNeverStarts returns a boolean if a field has been set.

### GetNextBindingState

`func (o *Lease) GetNextBindingState() string`

GetNextBindingState returns the NextBindingState field if non-nil, zero value otherwise.

### GetNextBindingStateOk

`func (o *Lease) GetNextBindingStateOk() (*string, bool)`

GetNextBindingStateOk returns a tuple with the NextBindingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextBindingState

`func (o *Lease) SetNextBindingState(v string)`

SetNextBindingState sets NextBindingState field to given value.

### HasNextBindingState

`func (o *Lease) HasNextBindingState() bool`

HasNextBindingState returns a boolean if a field has been set.

### GetOnCommit

`func (o *Lease) GetOnCommit() string`

GetOnCommit returns the OnCommit field if non-nil, zero value otherwise.

### GetOnCommitOk

`func (o *Lease) GetOnCommitOk() (*string, bool)`

GetOnCommitOk returns a tuple with the OnCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnCommit

`func (o *Lease) SetOnCommit(v string)`

SetOnCommit sets OnCommit field to given value.

### HasOnCommit

`func (o *Lease) HasOnCommit() bool`

HasOnCommit returns a boolean if a field has been set.

### GetOnExpiry

`func (o *Lease) GetOnExpiry() string`

GetOnExpiry returns the OnExpiry field if non-nil, zero value otherwise.

### GetOnExpiryOk

`func (o *Lease) GetOnExpiryOk() (*string, bool)`

GetOnExpiryOk returns a tuple with the OnExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnExpiry

`func (o *Lease) SetOnExpiry(v string)`

SetOnExpiry sets OnExpiry field to given value.

### HasOnExpiry

`func (o *Lease) HasOnExpiry() bool`

HasOnExpiry returns a boolean if a field has been set.

### GetOnRelease

`func (o *Lease) GetOnRelease() string`

GetOnRelease returns the OnRelease field if non-nil, zero value otherwise.

### GetOnReleaseOk

`func (o *Lease) GetOnReleaseOk() (*string, bool)`

GetOnReleaseOk returns a tuple with the OnRelease field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnRelease

`func (o *Lease) SetOnRelease(v string)`

SetOnRelease sets OnRelease field to given value.

### HasOnRelease

`func (o *Lease) HasOnRelease() bool`

HasOnRelease returns a boolean if a field has been set.

### GetOption

`func (o *Lease) GetOption() string`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *Lease) GetOptionOk() (*string, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *Lease) SetOption(v string)`

SetOption sets Option field to given value.

### HasOption

`func (o *Lease) HasOption() bool`

HasOption returns a boolean if a field has been set.

### GetProtocol

`func (o *Lease) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Lease) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Lease) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Lease) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRemoteId

`func (o *Lease) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *Lease) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *Lease) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *Lease) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.

### GetRequestedOptions

`func (o *Lease) GetRequestedOptions() string`

GetRequestedOptions returns the RequestedOptions field if non-nil, zero value otherwise.

### GetRequestedOptionsOk

`func (o *Lease) GetRequestedOptionsOk() (*string, bool)`

GetRequestedOptionsOk returns a tuple with the RequestedOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedOptions

`func (o *Lease) SetRequestedOptions(v string)`

SetRequestedOptions sets RequestedOptions field to given value.

### HasRequestedOptions

`func (o *Lease) HasRequestedOptions() bool`

HasRequestedOptions returns a boolean if a field has been set.

### GetServedBy

`func (o *Lease) GetServedBy() string`

GetServedBy returns the ServedBy field if non-nil, zero value otherwise.

### GetServedByOk

`func (o *Lease) GetServedByOk() (*string, bool)`

GetServedByOk returns a tuple with the ServedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServedBy

`func (o *Lease) SetServedBy(v string)`

SetServedBy sets ServedBy field to given value.

### HasServedBy

`func (o *Lease) HasServedBy() bool`

HasServedBy returns a boolean if a field has been set.

### GetServerHostName

`func (o *Lease) GetServerHostName() string`

GetServerHostName returns the ServerHostName field if non-nil, zero value otherwise.

### GetServerHostNameOk

`func (o *Lease) GetServerHostNameOk() (*string, bool)`

GetServerHostNameOk returns a tuple with the ServerHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostName

`func (o *Lease) SetServerHostName(v string)`

SetServerHostName sets ServerHostName field to given value.

### HasServerHostName

`func (o *Lease) HasServerHostName() bool`

HasServerHostName returns a boolean if a field has been set.

### GetStarts

`func (o *Lease) GetStarts() int64`

GetStarts returns the Starts field if non-nil, zero value otherwise.

### GetStartsOk

`func (o *Lease) GetStartsOk() (*int64, bool)`

GetStartsOk returns a tuple with the Starts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarts

`func (o *Lease) SetStarts(v int64)`

SetStarts sets Starts field to given value.

### HasStarts

`func (o *Lease) HasStarts() bool`

HasStarts returns a boolean if a field has been set.

### GetTsfp

`func (o *Lease) GetTsfp() int64`

GetTsfp returns the Tsfp field if non-nil, zero value otherwise.

### GetTsfpOk

`func (o *Lease) GetTsfpOk() (*int64, bool)`

GetTsfpOk returns a tuple with the Tsfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsfp

`func (o *Lease) SetTsfp(v int64)`

SetTsfp sets Tsfp field to given value.

### HasTsfp

`func (o *Lease) HasTsfp() bool`

HasTsfp returns a boolean if a field has been set.

### GetTstp

`func (o *Lease) GetTstp() int64`

GetTstp returns the Tstp field if non-nil, zero value otherwise.

### GetTstpOk

`func (o *Lease) GetTstpOk() (*int64, bool)`

GetTstpOk returns a tuple with the Tstp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTstp

`func (o *Lease) SetTstp(v int64)`

SetTstp sets Tstp field to given value.

### HasTstp

`func (o *Lease) HasTstp() bool`

HasTstp returns a boolean if a field has been set.

### GetUid

`func (o *Lease) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *Lease) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *Lease) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *Lease) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *Lease) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Lease) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Lease) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Lease) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetVariable

`func (o *Lease) GetVariable() string`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *Lease) GetVariableOk() (*string, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *Lease) SetVariable(v string)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *Lease) HasVariable() bool`

HasVariable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


