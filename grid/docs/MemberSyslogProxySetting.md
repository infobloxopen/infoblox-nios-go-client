# MemberSyslogProxySetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | If set to True, the member receives syslog messages from specified devices, such as syslog servers and routers, and then forwards these messages to an external syslog server. | [optional] 
**TcpEnable** | Pointer to **bool** | If set to True, the appliance can receive messages from other devices via TCP. | [optional] 
**TcpPort** | Pointer to **int64** | The TCP port the appliance must listen on. | [optional] 
**UdpEnable** | Pointer to **bool** | If set to True, the appliance can receive messages from other devices via UDP. | [optional] 
**UdpPort** | Pointer to **int64** | The UDP port the appliance must listen on. | [optional] 
**ClientAcls** | Pointer to [**[]MembersyslogproxysettingClientAcls**](MembersyslogproxysettingClientAcls.md) | This list controls the IP addresses and networks that are allowed to access the syslog proxy. | [optional] 

## Methods

### NewMemberSyslogProxySetting

`func NewMemberSyslogProxySetting() *MemberSyslogProxySetting`

NewMemberSyslogProxySetting instantiates a new MemberSyslogProxySetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSyslogProxySettingWithDefaults

`func NewMemberSyslogProxySettingWithDefaults() *MemberSyslogProxySetting`

NewMemberSyslogProxySettingWithDefaults instantiates a new MemberSyslogProxySetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *MemberSyslogProxySetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MemberSyslogProxySetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MemberSyslogProxySetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *MemberSyslogProxySetting) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetTcpEnable

`func (o *MemberSyslogProxySetting) GetTcpEnable() bool`

GetTcpEnable returns the TcpEnable field if non-nil, zero value otherwise.

### GetTcpEnableOk

`func (o *MemberSyslogProxySetting) GetTcpEnableOk() (*bool, bool)`

GetTcpEnableOk returns a tuple with the TcpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpEnable

`func (o *MemberSyslogProxySetting) SetTcpEnable(v bool)`

SetTcpEnable sets TcpEnable field to given value.

### HasTcpEnable

`func (o *MemberSyslogProxySetting) HasTcpEnable() bool`

HasTcpEnable returns a boolean if a field has been set.

### GetTcpPort

`func (o *MemberSyslogProxySetting) GetTcpPort() int64`

GetTcpPort returns the TcpPort field if non-nil, zero value otherwise.

### GetTcpPortOk

`func (o *MemberSyslogProxySetting) GetTcpPortOk() (*int64, bool)`

GetTcpPortOk returns a tuple with the TcpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpPort

`func (o *MemberSyslogProxySetting) SetTcpPort(v int64)`

SetTcpPort sets TcpPort field to given value.

### HasTcpPort

`func (o *MemberSyslogProxySetting) HasTcpPort() bool`

HasTcpPort returns a boolean if a field has been set.

### GetUdpEnable

`func (o *MemberSyslogProxySetting) GetUdpEnable() bool`

GetUdpEnable returns the UdpEnable field if non-nil, zero value otherwise.

### GetUdpEnableOk

`func (o *MemberSyslogProxySetting) GetUdpEnableOk() (*bool, bool)`

GetUdpEnableOk returns a tuple with the UdpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpEnable

`func (o *MemberSyslogProxySetting) SetUdpEnable(v bool)`

SetUdpEnable sets UdpEnable field to given value.

### HasUdpEnable

`func (o *MemberSyslogProxySetting) HasUdpEnable() bool`

HasUdpEnable returns a boolean if a field has been set.

### GetUdpPort

`func (o *MemberSyslogProxySetting) GetUdpPort() int64`

GetUdpPort returns the UdpPort field if non-nil, zero value otherwise.

### GetUdpPortOk

`func (o *MemberSyslogProxySetting) GetUdpPortOk() (*int64, bool)`

GetUdpPortOk returns a tuple with the UdpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpPort

`func (o *MemberSyslogProxySetting) SetUdpPort(v int64)`

SetUdpPort sets UdpPort field to given value.

### HasUdpPort

`func (o *MemberSyslogProxySetting) HasUdpPort() bool`

HasUdpPort returns a boolean if a field has been set.

### GetClientAcls

`func (o *MemberSyslogProxySetting) GetClientAcls() []MembersyslogproxysettingClientAcls`

GetClientAcls returns the ClientAcls field if non-nil, zero value otherwise.

### GetClientAclsOk

`func (o *MemberSyslogProxySetting) GetClientAclsOk() (*[]MembersyslogproxysettingClientAcls, bool)`

GetClientAclsOk returns a tuple with the ClientAcls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAcls

`func (o *MemberSyslogProxySetting) SetClientAcls(v []MembersyslogproxysettingClientAcls)`

SetClientAcls sets ClientAcls field to given value.

### HasClientAcls

`func (o *MemberSyslogProxySetting) HasClientAcls() bool`

HasClientAcls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


