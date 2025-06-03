# MemberDnsDnsViewAddressSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewName** | Pointer to **string** | The reference to DNS View | [optional] 
**DnsNotifyTransferSource** | Pointer to **string** | Determines which IP address is used as the source for DDNS notify and transfer operations. | [optional] 
**DnsNotifyTransferSourceAddress** | Pointer to **string** | The source address used if dns_notify_transfer_source type is \&quot;IP\&quot;. | [optional] 
**DnsQuerySourceInterface** | Pointer to **string** | Determines which IP address is used as the source for DDNS query operations. | [optional] 
**DnsQuerySourceAddress** | Pointer to **string** | The source address used if dns_query_source_interface type is \&quot;IP\&quot;. | [optional] 
**EnableNotifySourcePort** | Pointer to **bool** | Determines if the notify source port for a view is enabled or not. | [optional] 
**NotifySourcePort** | Pointer to **int64** | The source port for notify messages. When requesting zone transfers from the primary server, some secondary DNS servers use the source port number (the primary server used to send the notify message) as the destination port number in the zone transfer request. This setting overrides Grid static source port settings. Valid values are between 1 and 63999. The default is selected by BIND. | [optional] 
**EnableQuerySourcePort** | Pointer to **bool** | Determines if the query source port for a view is enabled or not. | [optional] 
**QuerySourcePort** | Pointer to **int64** | The source port for queries. Specifying a source port number for recursive queries ensures that a firewall will allow the response. Valid values are between 1 and 63999. The default is selected by BIND. | [optional] 
**NotifyDelay** | Pointer to **int64** | Specifies the number of seconds of delay the notify messages are sent to secondaries. | [optional] 
**UseSourcePorts** | Pointer to **bool** | Use flag for: enable_notify_source_port , notify_source_port, enable_query_source_port, query_source_port | [optional] 
**UseNotifyDelay** | Pointer to **bool** | Use flag for: notify_delay | [optional] 

## Methods

### NewMemberDnsDnsViewAddressSettings

`func NewMemberDnsDnsViewAddressSettings() *MemberDnsDnsViewAddressSettings`

NewMemberDnsDnsViewAddressSettings instantiates a new MemberDnsDnsViewAddressSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsDnsViewAddressSettingsWithDefaults

`func NewMemberDnsDnsViewAddressSettingsWithDefaults() *MemberDnsDnsViewAddressSettings`

NewMemberDnsDnsViewAddressSettingsWithDefaults instantiates a new MemberDnsDnsViewAddressSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewName

`func (o *MemberDnsDnsViewAddressSettings) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *MemberDnsDnsViewAddressSettings) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *MemberDnsDnsViewAddressSettings) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *MemberDnsDnsViewAddressSettings) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetDnsNotifyTransferSource

`func (o *MemberDnsDnsViewAddressSettings) GetDnsNotifyTransferSource() string`

GetDnsNotifyTransferSource returns the DnsNotifyTransferSource field if non-nil, zero value otherwise.

### GetDnsNotifyTransferSourceOk

`func (o *MemberDnsDnsViewAddressSettings) GetDnsNotifyTransferSourceOk() (*string, bool)`

GetDnsNotifyTransferSourceOk returns a tuple with the DnsNotifyTransferSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNotifyTransferSource

`func (o *MemberDnsDnsViewAddressSettings) SetDnsNotifyTransferSource(v string)`

SetDnsNotifyTransferSource sets DnsNotifyTransferSource field to given value.

### HasDnsNotifyTransferSource

`func (o *MemberDnsDnsViewAddressSettings) HasDnsNotifyTransferSource() bool`

HasDnsNotifyTransferSource returns a boolean if a field has been set.

### GetDnsNotifyTransferSourceAddress

`func (o *MemberDnsDnsViewAddressSettings) GetDnsNotifyTransferSourceAddress() string`

GetDnsNotifyTransferSourceAddress returns the DnsNotifyTransferSourceAddress field if non-nil, zero value otherwise.

### GetDnsNotifyTransferSourceAddressOk

`func (o *MemberDnsDnsViewAddressSettings) GetDnsNotifyTransferSourceAddressOk() (*string, bool)`

GetDnsNotifyTransferSourceAddressOk returns a tuple with the DnsNotifyTransferSourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNotifyTransferSourceAddress

`func (o *MemberDnsDnsViewAddressSettings) SetDnsNotifyTransferSourceAddress(v string)`

SetDnsNotifyTransferSourceAddress sets DnsNotifyTransferSourceAddress field to given value.

### HasDnsNotifyTransferSourceAddress

`func (o *MemberDnsDnsViewAddressSettings) HasDnsNotifyTransferSourceAddress() bool`

HasDnsNotifyTransferSourceAddress returns a boolean if a field has been set.

### GetDnsQuerySourceInterface

`func (o *MemberDnsDnsViewAddressSettings) GetDnsQuerySourceInterface() string`

GetDnsQuerySourceInterface returns the DnsQuerySourceInterface field if non-nil, zero value otherwise.

### GetDnsQuerySourceInterfaceOk

`func (o *MemberDnsDnsViewAddressSettings) GetDnsQuerySourceInterfaceOk() (*string, bool)`

GetDnsQuerySourceInterfaceOk returns a tuple with the DnsQuerySourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQuerySourceInterface

`func (o *MemberDnsDnsViewAddressSettings) SetDnsQuerySourceInterface(v string)`

SetDnsQuerySourceInterface sets DnsQuerySourceInterface field to given value.

### HasDnsQuerySourceInterface

`func (o *MemberDnsDnsViewAddressSettings) HasDnsQuerySourceInterface() bool`

HasDnsQuerySourceInterface returns a boolean if a field has been set.

### GetDnsQuerySourceAddress

`func (o *MemberDnsDnsViewAddressSettings) GetDnsQuerySourceAddress() string`

GetDnsQuerySourceAddress returns the DnsQuerySourceAddress field if non-nil, zero value otherwise.

### GetDnsQuerySourceAddressOk

`func (o *MemberDnsDnsViewAddressSettings) GetDnsQuerySourceAddressOk() (*string, bool)`

GetDnsQuerySourceAddressOk returns a tuple with the DnsQuerySourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsQuerySourceAddress

`func (o *MemberDnsDnsViewAddressSettings) SetDnsQuerySourceAddress(v string)`

SetDnsQuerySourceAddress sets DnsQuerySourceAddress field to given value.

### HasDnsQuerySourceAddress

`func (o *MemberDnsDnsViewAddressSettings) HasDnsQuerySourceAddress() bool`

HasDnsQuerySourceAddress returns a boolean if a field has been set.

### GetEnableNotifySourcePort

`func (o *MemberDnsDnsViewAddressSettings) GetEnableNotifySourcePort() bool`

GetEnableNotifySourcePort returns the EnableNotifySourcePort field if non-nil, zero value otherwise.

### GetEnableNotifySourcePortOk

`func (o *MemberDnsDnsViewAddressSettings) GetEnableNotifySourcePortOk() (*bool, bool)`

GetEnableNotifySourcePortOk returns a tuple with the EnableNotifySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableNotifySourcePort

`func (o *MemberDnsDnsViewAddressSettings) SetEnableNotifySourcePort(v bool)`

SetEnableNotifySourcePort sets EnableNotifySourcePort field to given value.

### HasEnableNotifySourcePort

`func (o *MemberDnsDnsViewAddressSettings) HasEnableNotifySourcePort() bool`

HasEnableNotifySourcePort returns a boolean if a field has been set.

### GetNotifySourcePort

`func (o *MemberDnsDnsViewAddressSettings) GetNotifySourcePort() int64`

GetNotifySourcePort returns the NotifySourcePort field if non-nil, zero value otherwise.

### GetNotifySourcePortOk

`func (o *MemberDnsDnsViewAddressSettings) GetNotifySourcePortOk() (*int64, bool)`

GetNotifySourcePortOk returns a tuple with the NotifySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifySourcePort

`func (o *MemberDnsDnsViewAddressSettings) SetNotifySourcePort(v int64)`

SetNotifySourcePort sets NotifySourcePort field to given value.

### HasNotifySourcePort

`func (o *MemberDnsDnsViewAddressSettings) HasNotifySourcePort() bool`

HasNotifySourcePort returns a boolean if a field has been set.

### GetEnableQuerySourcePort

`func (o *MemberDnsDnsViewAddressSettings) GetEnableQuerySourcePort() bool`

GetEnableQuerySourcePort returns the EnableQuerySourcePort field if non-nil, zero value otherwise.

### GetEnableQuerySourcePortOk

`func (o *MemberDnsDnsViewAddressSettings) GetEnableQuerySourcePortOk() (*bool, bool)`

GetEnableQuerySourcePortOk returns a tuple with the EnableQuerySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableQuerySourcePort

`func (o *MemberDnsDnsViewAddressSettings) SetEnableQuerySourcePort(v bool)`

SetEnableQuerySourcePort sets EnableQuerySourcePort field to given value.

### HasEnableQuerySourcePort

`func (o *MemberDnsDnsViewAddressSettings) HasEnableQuerySourcePort() bool`

HasEnableQuerySourcePort returns a boolean if a field has been set.

### GetQuerySourcePort

`func (o *MemberDnsDnsViewAddressSettings) GetQuerySourcePort() int64`

GetQuerySourcePort returns the QuerySourcePort field if non-nil, zero value otherwise.

### GetQuerySourcePortOk

`func (o *MemberDnsDnsViewAddressSettings) GetQuerySourcePortOk() (*int64, bool)`

GetQuerySourcePortOk returns a tuple with the QuerySourcePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuerySourcePort

`func (o *MemberDnsDnsViewAddressSettings) SetQuerySourcePort(v int64)`

SetQuerySourcePort sets QuerySourcePort field to given value.

### HasQuerySourcePort

`func (o *MemberDnsDnsViewAddressSettings) HasQuerySourcePort() bool`

HasQuerySourcePort returns a boolean if a field has been set.

### GetNotifyDelay

`func (o *MemberDnsDnsViewAddressSettings) GetNotifyDelay() int64`

GetNotifyDelay returns the NotifyDelay field if non-nil, zero value otherwise.

### GetNotifyDelayOk

`func (o *MemberDnsDnsViewAddressSettings) GetNotifyDelayOk() (*int64, bool)`

GetNotifyDelayOk returns a tuple with the NotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyDelay

`func (o *MemberDnsDnsViewAddressSettings) SetNotifyDelay(v int64)`

SetNotifyDelay sets NotifyDelay field to given value.

### HasNotifyDelay

`func (o *MemberDnsDnsViewAddressSettings) HasNotifyDelay() bool`

HasNotifyDelay returns a boolean if a field has been set.

### GetUseSourcePorts

`func (o *MemberDnsDnsViewAddressSettings) GetUseSourcePorts() bool`

GetUseSourcePorts returns the UseSourcePorts field if non-nil, zero value otherwise.

### GetUseSourcePortsOk

`func (o *MemberDnsDnsViewAddressSettings) GetUseSourcePortsOk() (*bool, bool)`

GetUseSourcePortsOk returns a tuple with the UseSourcePorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSourcePorts

`func (o *MemberDnsDnsViewAddressSettings) SetUseSourcePorts(v bool)`

SetUseSourcePorts sets UseSourcePorts field to given value.

### HasUseSourcePorts

`func (o *MemberDnsDnsViewAddressSettings) HasUseSourcePorts() bool`

HasUseSourcePorts returns a boolean if a field has been set.

### GetUseNotifyDelay

`func (o *MemberDnsDnsViewAddressSettings) GetUseNotifyDelay() bool`

GetUseNotifyDelay returns the UseNotifyDelay field if non-nil, zero value otherwise.

### GetUseNotifyDelayOk

`func (o *MemberDnsDnsViewAddressSettings) GetUseNotifyDelayOk() (*bool, bool)`

GetUseNotifyDelayOk returns a tuple with the UseNotifyDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNotifyDelay

`func (o *MemberDnsDnsViewAddressSettings) SetUseNotifyDelay(v bool)`

SetUseNotifyDelay sets UseNotifyDelay field to given value.

### HasUseNotifyDelay

`func (o *MemberDnsDnsViewAddressSettings) HasUseNotifyDelay() bool`

HasUseNotifyDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


