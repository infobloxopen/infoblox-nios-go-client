# MemberNodeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceStatus** | Pointer to [**MembernodeinfoServiceStatus**](MembernodeinfoServiceStatus.md) |  | [optional] 
**PhysicalOid** | Pointer to **string** | The OID of the physical node. | [optional] [readonly] 
**HaStatus** | Pointer to **string** | Status about the node of an HA pair. | [optional] [readonly] 
**Hwid** | Pointer to **string** | Hardware ID. | [optional] [readonly] 
**Hwmodel** | Pointer to **string** | Hardware model. | [optional] [readonly] 
**Hwtype** | Pointer to **string** | Hardware type. | [optional] [readonly] 
**HostPlatform** | Pointer to **string** | Host Platform | [optional] [readonly] 
**Hypervisor** | Pointer to **string** | Hypervisor | [optional] [readonly] 
**PaidNios** | Pointer to **bool** | True if node is Paid NIOS. | [optional] [readonly] 
**MgmtNetworkSetting** | Pointer to [**MembernodeinfoMgmtNetworkSetting**](MembernodeinfoMgmtNetworkSetting.md) |  | [optional] 
**LanHaPortSetting** | Pointer to [**MembernodeinfoLanHaPortSetting**](MembernodeinfoLanHaPortSetting.md) |  | [optional] 
**MgmtPhysicalSetting** | Pointer to [**MembernodeinfoMgmtPhysicalSetting**](MembernodeinfoMgmtPhysicalSetting.md) |  | [optional] 
**Lan2PhysicalSetting** | Pointer to [**MembernodeinfoLan2PhysicalSetting**](MembernodeinfoLan2PhysicalSetting.md) |  | [optional] 
**NatExternalIp** | Pointer to **string** | The NAT external IP address for the node. | [optional] 
**V6MgmtNetworkSetting** | Pointer to [**MembernodeinfoV6MgmtNetworkSetting**](MembernodeinfoV6MgmtNetworkSetting.md) |  | [optional] 

## Methods

### NewMemberNodeInfo

`func NewMemberNodeInfo() *MemberNodeInfo`

NewMemberNodeInfo instantiates a new MemberNodeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberNodeInfoWithDefaults

`func NewMemberNodeInfoWithDefaults() *MemberNodeInfo`

NewMemberNodeInfoWithDefaults instantiates a new MemberNodeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceStatus

`func (o *MemberNodeInfo) GetServiceStatus() MembernodeinfoServiceStatus`

GetServiceStatus returns the ServiceStatus field if non-nil, zero value otherwise.

### GetServiceStatusOk

`func (o *MemberNodeInfo) GetServiceStatusOk() (*MembernodeinfoServiceStatus, bool)`

GetServiceStatusOk returns a tuple with the ServiceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceStatus

`func (o *MemberNodeInfo) SetServiceStatus(v MembernodeinfoServiceStatus)`

SetServiceStatus sets ServiceStatus field to given value.

### HasServiceStatus

`func (o *MemberNodeInfo) HasServiceStatus() bool`

HasServiceStatus returns a boolean if a field has been set.

### GetPhysicalOid

`func (o *MemberNodeInfo) GetPhysicalOid() string`

GetPhysicalOid returns the PhysicalOid field if non-nil, zero value otherwise.

### GetPhysicalOidOk

`func (o *MemberNodeInfo) GetPhysicalOidOk() (*string, bool)`

GetPhysicalOidOk returns a tuple with the PhysicalOid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalOid

`func (o *MemberNodeInfo) SetPhysicalOid(v string)`

SetPhysicalOid sets PhysicalOid field to given value.

### HasPhysicalOid

`func (o *MemberNodeInfo) HasPhysicalOid() bool`

HasPhysicalOid returns a boolean if a field has been set.

### GetHaStatus

`func (o *MemberNodeInfo) GetHaStatus() string`

GetHaStatus returns the HaStatus field if non-nil, zero value otherwise.

### GetHaStatusOk

`func (o *MemberNodeInfo) GetHaStatusOk() (*string, bool)`

GetHaStatusOk returns a tuple with the HaStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaStatus

`func (o *MemberNodeInfo) SetHaStatus(v string)`

SetHaStatus sets HaStatus field to given value.

### HasHaStatus

`func (o *MemberNodeInfo) HasHaStatus() bool`

HasHaStatus returns a boolean if a field has been set.

### GetHwid

`func (o *MemberNodeInfo) GetHwid() string`

GetHwid returns the Hwid field if non-nil, zero value otherwise.

### GetHwidOk

`func (o *MemberNodeInfo) GetHwidOk() (*string, bool)`

GetHwidOk returns a tuple with the Hwid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwid

`func (o *MemberNodeInfo) SetHwid(v string)`

SetHwid sets Hwid field to given value.

### HasHwid

`func (o *MemberNodeInfo) HasHwid() bool`

HasHwid returns a boolean if a field has been set.

### GetHwmodel

`func (o *MemberNodeInfo) GetHwmodel() string`

GetHwmodel returns the Hwmodel field if non-nil, zero value otherwise.

### GetHwmodelOk

`func (o *MemberNodeInfo) GetHwmodelOk() (*string, bool)`

GetHwmodelOk returns a tuple with the Hwmodel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwmodel

`func (o *MemberNodeInfo) SetHwmodel(v string)`

SetHwmodel sets Hwmodel field to given value.

### HasHwmodel

`func (o *MemberNodeInfo) HasHwmodel() bool`

HasHwmodel returns a boolean if a field has been set.

### GetHwtype

`func (o *MemberNodeInfo) GetHwtype() string`

GetHwtype returns the Hwtype field if non-nil, zero value otherwise.

### GetHwtypeOk

`func (o *MemberNodeInfo) GetHwtypeOk() (*string, bool)`

GetHwtypeOk returns a tuple with the Hwtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwtype

`func (o *MemberNodeInfo) SetHwtype(v string)`

SetHwtype sets Hwtype field to given value.

### HasHwtype

`func (o *MemberNodeInfo) HasHwtype() bool`

HasHwtype returns a boolean if a field has been set.

### GetHostPlatform

`func (o *MemberNodeInfo) GetHostPlatform() string`

GetHostPlatform returns the HostPlatform field if non-nil, zero value otherwise.

### GetHostPlatformOk

`func (o *MemberNodeInfo) GetHostPlatformOk() (*string, bool)`

GetHostPlatformOk returns a tuple with the HostPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostPlatform

`func (o *MemberNodeInfo) SetHostPlatform(v string)`

SetHostPlatform sets HostPlatform field to given value.

### HasHostPlatform

`func (o *MemberNodeInfo) HasHostPlatform() bool`

HasHostPlatform returns a boolean if a field has been set.

### GetHypervisor

`func (o *MemberNodeInfo) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *MemberNodeInfo) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *MemberNodeInfo) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *MemberNodeInfo) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetPaidNios

`func (o *MemberNodeInfo) GetPaidNios() bool`

GetPaidNios returns the PaidNios field if non-nil, zero value otherwise.

### GetPaidNiosOk

`func (o *MemberNodeInfo) GetPaidNiosOk() (*bool, bool)`

GetPaidNiosOk returns a tuple with the PaidNios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidNios

`func (o *MemberNodeInfo) SetPaidNios(v bool)`

SetPaidNios sets PaidNios field to given value.

### HasPaidNios

`func (o *MemberNodeInfo) HasPaidNios() bool`

HasPaidNios returns a boolean if a field has been set.

### GetMgmtNetworkSetting

`func (o *MemberNodeInfo) GetMgmtNetworkSetting() MembernodeinfoMgmtNetworkSetting`

GetMgmtNetworkSetting returns the MgmtNetworkSetting field if non-nil, zero value otherwise.

### GetMgmtNetworkSettingOk

`func (o *MemberNodeInfo) GetMgmtNetworkSettingOk() (*MembernodeinfoMgmtNetworkSetting, bool)`

GetMgmtNetworkSettingOk returns a tuple with the MgmtNetworkSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtNetworkSetting

`func (o *MemberNodeInfo) SetMgmtNetworkSetting(v MembernodeinfoMgmtNetworkSetting)`

SetMgmtNetworkSetting sets MgmtNetworkSetting field to given value.

### HasMgmtNetworkSetting

`func (o *MemberNodeInfo) HasMgmtNetworkSetting() bool`

HasMgmtNetworkSetting returns a boolean if a field has been set.

### GetLanHaPortSetting

`func (o *MemberNodeInfo) GetLanHaPortSetting() MembernodeinfoLanHaPortSetting`

GetLanHaPortSetting returns the LanHaPortSetting field if non-nil, zero value otherwise.

### GetLanHaPortSettingOk

`func (o *MemberNodeInfo) GetLanHaPortSettingOk() (*MembernodeinfoLanHaPortSetting, bool)`

GetLanHaPortSettingOk returns a tuple with the LanHaPortSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanHaPortSetting

`func (o *MemberNodeInfo) SetLanHaPortSetting(v MembernodeinfoLanHaPortSetting)`

SetLanHaPortSetting sets LanHaPortSetting field to given value.

### HasLanHaPortSetting

`func (o *MemberNodeInfo) HasLanHaPortSetting() bool`

HasLanHaPortSetting returns a boolean if a field has been set.

### GetMgmtPhysicalSetting

`func (o *MemberNodeInfo) GetMgmtPhysicalSetting() MembernodeinfoMgmtPhysicalSetting`

GetMgmtPhysicalSetting returns the MgmtPhysicalSetting field if non-nil, zero value otherwise.

### GetMgmtPhysicalSettingOk

`func (o *MemberNodeInfo) GetMgmtPhysicalSettingOk() (*MembernodeinfoMgmtPhysicalSetting, bool)`

GetMgmtPhysicalSettingOk returns a tuple with the MgmtPhysicalSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPhysicalSetting

`func (o *MemberNodeInfo) SetMgmtPhysicalSetting(v MembernodeinfoMgmtPhysicalSetting)`

SetMgmtPhysicalSetting sets MgmtPhysicalSetting field to given value.

### HasMgmtPhysicalSetting

`func (o *MemberNodeInfo) HasMgmtPhysicalSetting() bool`

HasMgmtPhysicalSetting returns a boolean if a field has been set.

### GetLan2PhysicalSetting

`func (o *MemberNodeInfo) GetLan2PhysicalSetting() MembernodeinfoLan2PhysicalSetting`

GetLan2PhysicalSetting returns the Lan2PhysicalSetting field if non-nil, zero value otherwise.

### GetLan2PhysicalSettingOk

`func (o *MemberNodeInfo) GetLan2PhysicalSettingOk() (*MembernodeinfoLan2PhysicalSetting, bool)`

GetLan2PhysicalSettingOk returns a tuple with the Lan2PhysicalSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLan2PhysicalSetting

`func (o *MemberNodeInfo) SetLan2PhysicalSetting(v MembernodeinfoLan2PhysicalSetting)`

SetLan2PhysicalSetting sets Lan2PhysicalSetting field to given value.

### HasLan2PhysicalSetting

`func (o *MemberNodeInfo) HasLan2PhysicalSetting() bool`

HasLan2PhysicalSetting returns a boolean if a field has been set.

### GetNatExternalIp

`func (o *MemberNodeInfo) GetNatExternalIp() string`

GetNatExternalIp returns the NatExternalIp field if non-nil, zero value otherwise.

### GetNatExternalIpOk

`func (o *MemberNodeInfo) GetNatExternalIpOk() (*string, bool)`

GetNatExternalIpOk returns a tuple with the NatExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatExternalIp

`func (o *MemberNodeInfo) SetNatExternalIp(v string)`

SetNatExternalIp sets NatExternalIp field to given value.

### HasNatExternalIp

`func (o *MemberNodeInfo) HasNatExternalIp() bool`

HasNatExternalIp returns a boolean if a field has been set.

### GetV6MgmtNetworkSetting

`func (o *MemberNodeInfo) GetV6MgmtNetworkSetting() MembernodeinfoV6MgmtNetworkSetting`

GetV6MgmtNetworkSetting returns the V6MgmtNetworkSetting field if non-nil, zero value otherwise.

### GetV6MgmtNetworkSettingOk

`func (o *MemberNodeInfo) GetV6MgmtNetworkSettingOk() (*MembernodeinfoV6MgmtNetworkSetting, bool)`

GetV6MgmtNetworkSettingOk returns a tuple with the V6MgmtNetworkSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetV6MgmtNetworkSetting

`func (o *MemberNodeInfo) SetV6MgmtNetworkSetting(v MembernodeinfoV6MgmtNetworkSetting)`

SetV6MgmtNetworkSetting sets V6MgmtNetworkSetting field to given value.

### HasV6MgmtNetworkSetting

`func (o *MemberNodeInfo) HasV6MgmtNetworkSetting() bool`

HasV6MgmtNetworkSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


