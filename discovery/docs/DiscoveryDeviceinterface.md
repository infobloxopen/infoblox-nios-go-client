# DiscoveryDeviceinterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AdminStatus** | Pointer to **string** | Administrative state of the interface. | [optional] [readonly] 
**AdminStatusTaskInfo** | Pointer to [**DiscoveryDeviceinterfaceAdminStatusTaskInfo**](DiscoveryDeviceinterfaceAdminStatusTaskInfo.md) |  | [optional] 
**AggrInterfaceName** | Pointer to **string** | Name of the port channel current interface belongs to. | [optional] [readonly] 
**CapIfAdminStatusInd** | Pointer to **bool** | Determines whether to modify the admin status of the interface. | [optional] [readonly] 
**CapIfAdminStatusNaReason** | Pointer to **string** | The reason that the edit admin status action is not available. | [optional] [readonly] 
**CapIfDescriptionInd** | Pointer to **bool** | Determines whether to modify the description of the interface. | [optional] [readonly] 
**CapIfDescriptionNaReason** | Pointer to **string** | The reason that the edit description action is not available. | [optional] [readonly] 
**CapIfNetDeprovisioningIpv4Ind** | Pointer to **bool** | Determines whether to deprovision a IPv4 network from the interfaces. | [optional] [readonly] 
**CapIfNetDeprovisioningIpv4NaReason** | Pointer to **string** | The reason that the deprovision a IPv4 network from the interface. | [optional] [readonly] 
**CapIfNetDeprovisioningIpv6Ind** | Pointer to **bool** | Determines whether to deprovision a IPv6 network from the interfaces. | [optional] [readonly] 
**CapIfNetDeprovisioningIpv6NaReason** | Pointer to **string** | The reason that the deprovision a IPv6 network from the interface. | [optional] [readonly] 
**CapIfNetProvisioningIpv4Ind** | Pointer to **bool** | Determines whether to modify the IPv4 network associated to the interface. | [optional] [readonly] 
**CapIfNetProvisioningIpv4NaReason** | Pointer to **string** | The reason that IPv4 network provisioning is not available. | [optional] [readonly] 
**CapIfNetProvisioningIpv6Ind** | Pointer to **bool** | Determines whether to modify the IPv6 network associated to the interface. | [optional] [readonly] 
**CapIfNetProvisioningIpv6NaReason** | Pointer to **string** | The reason that IPv6 network provisioning is not available. | [optional] [readonly] 
**CapIfVlanAssignmentInd** | Pointer to **bool** | Determines whether to modify the VLAN assignement of the interface. | [optional] [readonly] 
**CapIfVlanAssignmentNaReason** | Pointer to **string** | The reason that VLAN assignment action is not available. | [optional] [readonly] 
**CapIfVoiceVlanInd** | Pointer to **bool** | Determines whether to modify the voice VLAN assignement of the interface. | [optional] [readonly] 
**CapIfVoiceVlanNaReason** | Pointer to **string** | The reason that voice VLAN assignment action is not available. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the interface. | [optional] [readonly] 
**DescriptionTaskInfo** | Pointer to [**DiscoveryDeviceinterfaceDescriptionTaskInfo**](DiscoveryDeviceinterfaceDescriptionTaskInfo.md) |  | [optional] 
**Device** | Pointer to **string** | The ref to the device to which the interface belongs. | [optional] [readonly] 
**Duplex** | Pointer to **string** | The duplex state of the interface. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**IfaddrInfos** | Pointer to [**[]DiscoveryDeviceinterfaceIfaddrInfos**](DiscoveryDeviceinterfaceIfaddrInfos.md) | List of IFaddr information associated with the interface. | [optional] [readonly] 
**Index** | Pointer to **int64** | The interface index number, as reported by SNMP. | [optional] [readonly] 
**LastChange** | Pointer to **int64** | Timestamp of the last interface property change detected. | [optional] [readonly] 
**LinkAggregation** | Pointer to **bool** | This field indicates if this is a link aggregation interface. | [optional] [readonly] 
**Mac** | Pointer to **string** | The MAC address of the interface. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**DiscoveryDeviceinterfaceMsAdUserData**](DiscoveryDeviceinterfaceMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | The interface system name. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | Th name of the network view. | [optional] [readonly] 
**OperStatus** | Pointer to **string** | Operating state of the interface. | [optional] [readonly] 
**PortFast** | Pointer to **string** | The Port Fast status of the interface. | [optional] [readonly] 
**ReservedObject** | Pointer to **string** | The reference to object(Host/FixedAddress/GridMember) to which this port is reserved. | [optional] [readonly] 
**Speed** | Pointer to **int64** | The interface speed in bps. | [optional] [readonly] 
**TrunkStatus** | Pointer to **string** | Indicates if the interface is tagged as a VLAN trunk or not. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of interface. | [optional] [readonly] 
**VlanInfoTaskInfo** | Pointer to [**DiscoveryDeviceinterfaceVlanInfoTaskInfo**](DiscoveryDeviceinterfaceVlanInfoTaskInfo.md) |  | [optional] 
**VlanInfos** | Pointer to [**[]DiscoveryDeviceinterfaceVlanInfos**](DiscoveryDeviceinterfaceVlanInfos.md) | The list of VLAN information associated with the interface. | [optional] [readonly] 
**VpcPeer** | Pointer to **string** | Aggregated interface name of vPC peer device current port is connected to. | [optional] [readonly] 
**VpcPeerDevice** | Pointer to **string** | The reference to vPC peer device. | [optional] [readonly] 
**VrfDescription** | Pointer to **string** | The description of the Virtual Routing and Forwarding (VRF) associated with the interface. | [optional] [readonly] 
**VrfName** | Pointer to **string** | The name of the Virtual Routing and Forwarding (VRF) associated with the interface. | [optional] [readonly] 
**VrfRd** | Pointer to **string** | The route distinguisher of the Virtual Routing and Forwarding (VRF) associated with the interface. | [optional] [readonly] 

## Methods

### NewDiscoveryDeviceinterface

`func NewDiscoveryDeviceinterface() *DiscoveryDeviceinterface`

NewDiscoveryDeviceinterface instantiates a new DiscoveryDeviceinterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryDeviceinterfaceWithDefaults

`func NewDiscoveryDeviceinterfaceWithDefaults() *DiscoveryDeviceinterface`

NewDiscoveryDeviceinterfaceWithDefaults instantiates a new DiscoveryDeviceinterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DiscoveryDeviceinterface) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DiscoveryDeviceinterface) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DiscoveryDeviceinterface) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DiscoveryDeviceinterface) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAdminStatus

`func (o *DiscoveryDeviceinterface) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *DiscoveryDeviceinterface) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *DiscoveryDeviceinterface) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *DiscoveryDeviceinterface) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetAdminStatusTaskInfo

`func (o *DiscoveryDeviceinterface) GetAdminStatusTaskInfo() DiscoveryDeviceinterfaceAdminStatusTaskInfo`

GetAdminStatusTaskInfo returns the AdminStatusTaskInfo field if non-nil, zero value otherwise.

### GetAdminStatusTaskInfoOk

`func (o *DiscoveryDeviceinterface) GetAdminStatusTaskInfoOk() (*DiscoveryDeviceinterfaceAdminStatusTaskInfo, bool)`

GetAdminStatusTaskInfoOk returns a tuple with the AdminStatusTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatusTaskInfo

`func (o *DiscoveryDeviceinterface) SetAdminStatusTaskInfo(v DiscoveryDeviceinterfaceAdminStatusTaskInfo)`

SetAdminStatusTaskInfo sets AdminStatusTaskInfo field to given value.

### HasAdminStatusTaskInfo

`func (o *DiscoveryDeviceinterface) HasAdminStatusTaskInfo() bool`

HasAdminStatusTaskInfo returns a boolean if a field has been set.

### GetAggrInterfaceName

`func (o *DiscoveryDeviceinterface) GetAggrInterfaceName() string`

GetAggrInterfaceName returns the AggrInterfaceName field if non-nil, zero value otherwise.

### GetAggrInterfaceNameOk

`func (o *DiscoveryDeviceinterface) GetAggrInterfaceNameOk() (*string, bool)`

GetAggrInterfaceNameOk returns a tuple with the AggrInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggrInterfaceName

`func (o *DiscoveryDeviceinterface) SetAggrInterfaceName(v string)`

SetAggrInterfaceName sets AggrInterfaceName field to given value.

### HasAggrInterfaceName

`func (o *DiscoveryDeviceinterface) HasAggrInterfaceName() bool`

HasAggrInterfaceName returns a boolean if a field has been set.

### GetCapIfAdminStatusInd

`func (o *DiscoveryDeviceinterface) GetCapIfAdminStatusInd() bool`

GetCapIfAdminStatusInd returns the CapIfAdminStatusInd field if non-nil, zero value otherwise.

### GetCapIfAdminStatusIndOk

`func (o *DiscoveryDeviceinterface) GetCapIfAdminStatusIndOk() (*bool, bool)`

GetCapIfAdminStatusIndOk returns a tuple with the CapIfAdminStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfAdminStatusInd

`func (o *DiscoveryDeviceinterface) SetCapIfAdminStatusInd(v bool)`

SetCapIfAdminStatusInd sets CapIfAdminStatusInd field to given value.

### HasCapIfAdminStatusInd

`func (o *DiscoveryDeviceinterface) HasCapIfAdminStatusInd() bool`

HasCapIfAdminStatusInd returns a boolean if a field has been set.

### GetCapIfAdminStatusNaReason

`func (o *DiscoveryDeviceinterface) GetCapIfAdminStatusNaReason() string`

GetCapIfAdminStatusNaReason returns the CapIfAdminStatusNaReason field if non-nil, zero value otherwise.

### GetCapIfAdminStatusNaReasonOk

`func (o *DiscoveryDeviceinterface) GetCapIfAdminStatusNaReasonOk() (*string, bool)`

GetCapIfAdminStatusNaReasonOk returns a tuple with the CapIfAdminStatusNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfAdminStatusNaReason

`func (o *DiscoveryDeviceinterface) SetCapIfAdminStatusNaReason(v string)`

SetCapIfAdminStatusNaReason sets CapIfAdminStatusNaReason field to given value.

### HasCapIfAdminStatusNaReason

`func (o *DiscoveryDeviceinterface) HasCapIfAdminStatusNaReason() bool`

HasCapIfAdminStatusNaReason returns a boolean if a field has been set.

### GetCapIfDescriptionInd

`func (o *DiscoveryDeviceinterface) GetCapIfDescriptionInd() bool`

GetCapIfDescriptionInd returns the CapIfDescriptionInd field if non-nil, zero value otherwise.

### GetCapIfDescriptionIndOk

`func (o *DiscoveryDeviceinterface) GetCapIfDescriptionIndOk() (*bool, bool)`

GetCapIfDescriptionIndOk returns a tuple with the CapIfDescriptionInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfDescriptionInd

`func (o *DiscoveryDeviceinterface) SetCapIfDescriptionInd(v bool)`

SetCapIfDescriptionInd sets CapIfDescriptionInd field to given value.

### HasCapIfDescriptionInd

`func (o *DiscoveryDeviceinterface) HasCapIfDescriptionInd() bool`

HasCapIfDescriptionInd returns a boolean if a field has been set.

### GetCapIfDescriptionNaReason

`func (o *DiscoveryDeviceinterface) GetCapIfDescriptionNaReason() string`

GetCapIfDescriptionNaReason returns the CapIfDescriptionNaReason field if non-nil, zero value otherwise.

### GetCapIfDescriptionNaReasonOk

`func (o *DiscoveryDeviceinterface) GetCapIfDescriptionNaReasonOk() (*string, bool)`

GetCapIfDescriptionNaReasonOk returns a tuple with the CapIfDescriptionNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfDescriptionNaReason

`func (o *DiscoveryDeviceinterface) SetCapIfDescriptionNaReason(v string)`

SetCapIfDescriptionNaReason sets CapIfDescriptionNaReason field to given value.

### HasCapIfDescriptionNaReason

`func (o *DiscoveryDeviceinterface) HasCapIfDescriptionNaReason() bool`

HasCapIfDescriptionNaReason returns a boolean if a field has been set.

### GetCapIfNetDeprovisioningIpv4Ind

`func (o *DiscoveryDeviceinterface) GetCapIfNetDeprovisioningIpv4Ind() bool`

GetCapIfNetDeprovisioningIpv4Ind returns the CapIfNetDeprovisioningIpv4Ind field if non-nil, zero value otherwise.

### GetCapIfNetDeprovisioningIpv4IndOk

`func (o *DiscoveryDeviceinterface) GetCapIfNetDeprovisioningIpv4IndOk() (*bool, bool)`

GetCapIfNetDeprovisioningIpv4IndOk returns a tuple with the CapIfNetDeprovisioningIpv4Ind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfNetDeprovisioningIpv4Ind

`func (o *DiscoveryDeviceinterface) SetCapIfNetDeprovisioningIpv4Ind(v bool)`

SetCapIfNetDeprovisioningIpv4Ind sets CapIfNetDeprovisioningIpv4Ind field to given value.

### HasCapIfNetDeprovisioningIpv4Ind

`func (o *DiscoveryDeviceinterface) HasCapIfNetDeprovisioningIpv4Ind() bool`

HasCapIfNetDeprovisioningIpv4Ind returns a boolean if a field has been set.

### GetCapIfNetDeprovisioningIpv4NaReason

`func (o *DiscoveryDeviceinterface) GetCapIfNetDeprovisioningIpv4NaReason() string`

GetCapIfNetDeprovisioningIpv4NaReason returns the CapIfNetDeprovisioningIpv4NaReason field if non-nil, zero value otherwise.

### GetCapIfNetDeprovisioningIpv4NaReasonOk

`func (o *DiscoveryDeviceinterface) GetCapIfNetDeprovisioningIpv4NaReasonOk() (*string, bool)`

GetCapIfNetDeprovisioningIpv4NaReasonOk returns a tuple with the CapIfNetDeprovisioningIpv4NaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfNetDeprovisioningIpv4NaReason

`func (o *DiscoveryDeviceinterface) SetCapIfNetDeprovisioningIpv4NaReason(v string)`

SetCapIfNetDeprovisioningIpv4NaReason sets CapIfNetDeprovisioningIpv4NaReason field to given value.

### HasCapIfNetDeprovisioningIpv4NaReason

`func (o *DiscoveryDeviceinterface) HasCapIfNetDeprovisioningIpv4NaReason() bool`

HasCapIfNetDeprovisioningIpv4NaReason returns a boolean if a field has been set.

### GetCapIfNetDeprovisioningIpv6Ind

`func (o *DiscoveryDeviceinterface) GetCapIfNetDeprovisioningIpv6Ind() bool`

GetCapIfNetDeprovisioningIpv6Ind returns the CapIfNetDeprovisioningIpv6Ind field if non-nil, zero value otherwise.

### GetCapIfNetDeprovisioningIpv6IndOk

`func (o *DiscoveryDeviceinterface) GetCapIfNetDeprovisioningIpv6IndOk() (*bool, bool)`

GetCapIfNetDeprovisioningIpv6IndOk returns a tuple with the CapIfNetDeprovisioningIpv6Ind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfNetDeprovisioningIpv6Ind

`func (o *DiscoveryDeviceinterface) SetCapIfNetDeprovisioningIpv6Ind(v bool)`

SetCapIfNetDeprovisioningIpv6Ind sets CapIfNetDeprovisioningIpv6Ind field to given value.

### HasCapIfNetDeprovisioningIpv6Ind

`func (o *DiscoveryDeviceinterface) HasCapIfNetDeprovisioningIpv6Ind() bool`

HasCapIfNetDeprovisioningIpv6Ind returns a boolean if a field has been set.

### GetCapIfNetDeprovisioningIpv6NaReason

`func (o *DiscoveryDeviceinterface) GetCapIfNetDeprovisioningIpv6NaReason() string`

GetCapIfNetDeprovisioningIpv6NaReason returns the CapIfNetDeprovisioningIpv6NaReason field if non-nil, zero value otherwise.

### GetCapIfNetDeprovisioningIpv6NaReasonOk

`func (o *DiscoveryDeviceinterface) GetCapIfNetDeprovisioningIpv6NaReasonOk() (*string, bool)`

GetCapIfNetDeprovisioningIpv6NaReasonOk returns a tuple with the CapIfNetDeprovisioningIpv6NaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfNetDeprovisioningIpv6NaReason

`func (o *DiscoveryDeviceinterface) SetCapIfNetDeprovisioningIpv6NaReason(v string)`

SetCapIfNetDeprovisioningIpv6NaReason sets CapIfNetDeprovisioningIpv6NaReason field to given value.

### HasCapIfNetDeprovisioningIpv6NaReason

`func (o *DiscoveryDeviceinterface) HasCapIfNetDeprovisioningIpv6NaReason() bool`

HasCapIfNetDeprovisioningIpv6NaReason returns a boolean if a field has been set.

### GetCapIfNetProvisioningIpv4Ind

`func (o *DiscoveryDeviceinterface) GetCapIfNetProvisioningIpv4Ind() bool`

GetCapIfNetProvisioningIpv4Ind returns the CapIfNetProvisioningIpv4Ind field if non-nil, zero value otherwise.

### GetCapIfNetProvisioningIpv4IndOk

`func (o *DiscoveryDeviceinterface) GetCapIfNetProvisioningIpv4IndOk() (*bool, bool)`

GetCapIfNetProvisioningIpv4IndOk returns a tuple with the CapIfNetProvisioningIpv4Ind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfNetProvisioningIpv4Ind

`func (o *DiscoveryDeviceinterface) SetCapIfNetProvisioningIpv4Ind(v bool)`

SetCapIfNetProvisioningIpv4Ind sets CapIfNetProvisioningIpv4Ind field to given value.

### HasCapIfNetProvisioningIpv4Ind

`func (o *DiscoveryDeviceinterface) HasCapIfNetProvisioningIpv4Ind() bool`

HasCapIfNetProvisioningIpv4Ind returns a boolean if a field has been set.

### GetCapIfNetProvisioningIpv4NaReason

`func (o *DiscoveryDeviceinterface) GetCapIfNetProvisioningIpv4NaReason() string`

GetCapIfNetProvisioningIpv4NaReason returns the CapIfNetProvisioningIpv4NaReason field if non-nil, zero value otherwise.

### GetCapIfNetProvisioningIpv4NaReasonOk

`func (o *DiscoveryDeviceinterface) GetCapIfNetProvisioningIpv4NaReasonOk() (*string, bool)`

GetCapIfNetProvisioningIpv4NaReasonOk returns a tuple with the CapIfNetProvisioningIpv4NaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfNetProvisioningIpv4NaReason

`func (o *DiscoveryDeviceinterface) SetCapIfNetProvisioningIpv4NaReason(v string)`

SetCapIfNetProvisioningIpv4NaReason sets CapIfNetProvisioningIpv4NaReason field to given value.

### HasCapIfNetProvisioningIpv4NaReason

`func (o *DiscoveryDeviceinterface) HasCapIfNetProvisioningIpv4NaReason() bool`

HasCapIfNetProvisioningIpv4NaReason returns a boolean if a field has been set.

### GetCapIfNetProvisioningIpv6Ind

`func (o *DiscoveryDeviceinterface) GetCapIfNetProvisioningIpv6Ind() bool`

GetCapIfNetProvisioningIpv6Ind returns the CapIfNetProvisioningIpv6Ind field if non-nil, zero value otherwise.

### GetCapIfNetProvisioningIpv6IndOk

`func (o *DiscoveryDeviceinterface) GetCapIfNetProvisioningIpv6IndOk() (*bool, bool)`

GetCapIfNetProvisioningIpv6IndOk returns a tuple with the CapIfNetProvisioningIpv6Ind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfNetProvisioningIpv6Ind

`func (o *DiscoveryDeviceinterface) SetCapIfNetProvisioningIpv6Ind(v bool)`

SetCapIfNetProvisioningIpv6Ind sets CapIfNetProvisioningIpv6Ind field to given value.

### HasCapIfNetProvisioningIpv6Ind

`func (o *DiscoveryDeviceinterface) HasCapIfNetProvisioningIpv6Ind() bool`

HasCapIfNetProvisioningIpv6Ind returns a boolean if a field has been set.

### GetCapIfNetProvisioningIpv6NaReason

`func (o *DiscoveryDeviceinterface) GetCapIfNetProvisioningIpv6NaReason() string`

GetCapIfNetProvisioningIpv6NaReason returns the CapIfNetProvisioningIpv6NaReason field if non-nil, zero value otherwise.

### GetCapIfNetProvisioningIpv6NaReasonOk

`func (o *DiscoveryDeviceinterface) GetCapIfNetProvisioningIpv6NaReasonOk() (*string, bool)`

GetCapIfNetProvisioningIpv6NaReasonOk returns a tuple with the CapIfNetProvisioningIpv6NaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfNetProvisioningIpv6NaReason

`func (o *DiscoveryDeviceinterface) SetCapIfNetProvisioningIpv6NaReason(v string)`

SetCapIfNetProvisioningIpv6NaReason sets CapIfNetProvisioningIpv6NaReason field to given value.

### HasCapIfNetProvisioningIpv6NaReason

`func (o *DiscoveryDeviceinterface) HasCapIfNetProvisioningIpv6NaReason() bool`

HasCapIfNetProvisioningIpv6NaReason returns a boolean if a field has been set.

### GetCapIfVlanAssignmentInd

`func (o *DiscoveryDeviceinterface) GetCapIfVlanAssignmentInd() bool`

GetCapIfVlanAssignmentInd returns the CapIfVlanAssignmentInd field if non-nil, zero value otherwise.

### GetCapIfVlanAssignmentIndOk

`func (o *DiscoveryDeviceinterface) GetCapIfVlanAssignmentIndOk() (*bool, bool)`

GetCapIfVlanAssignmentIndOk returns a tuple with the CapIfVlanAssignmentInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfVlanAssignmentInd

`func (o *DiscoveryDeviceinterface) SetCapIfVlanAssignmentInd(v bool)`

SetCapIfVlanAssignmentInd sets CapIfVlanAssignmentInd field to given value.

### HasCapIfVlanAssignmentInd

`func (o *DiscoveryDeviceinterface) HasCapIfVlanAssignmentInd() bool`

HasCapIfVlanAssignmentInd returns a boolean if a field has been set.

### GetCapIfVlanAssignmentNaReason

`func (o *DiscoveryDeviceinterface) GetCapIfVlanAssignmentNaReason() string`

GetCapIfVlanAssignmentNaReason returns the CapIfVlanAssignmentNaReason field if non-nil, zero value otherwise.

### GetCapIfVlanAssignmentNaReasonOk

`func (o *DiscoveryDeviceinterface) GetCapIfVlanAssignmentNaReasonOk() (*string, bool)`

GetCapIfVlanAssignmentNaReasonOk returns a tuple with the CapIfVlanAssignmentNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfVlanAssignmentNaReason

`func (o *DiscoveryDeviceinterface) SetCapIfVlanAssignmentNaReason(v string)`

SetCapIfVlanAssignmentNaReason sets CapIfVlanAssignmentNaReason field to given value.

### HasCapIfVlanAssignmentNaReason

`func (o *DiscoveryDeviceinterface) HasCapIfVlanAssignmentNaReason() bool`

HasCapIfVlanAssignmentNaReason returns a boolean if a field has been set.

### GetCapIfVoiceVlanInd

`func (o *DiscoveryDeviceinterface) GetCapIfVoiceVlanInd() bool`

GetCapIfVoiceVlanInd returns the CapIfVoiceVlanInd field if non-nil, zero value otherwise.

### GetCapIfVoiceVlanIndOk

`func (o *DiscoveryDeviceinterface) GetCapIfVoiceVlanIndOk() (*bool, bool)`

GetCapIfVoiceVlanIndOk returns a tuple with the CapIfVoiceVlanInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfVoiceVlanInd

`func (o *DiscoveryDeviceinterface) SetCapIfVoiceVlanInd(v bool)`

SetCapIfVoiceVlanInd sets CapIfVoiceVlanInd field to given value.

### HasCapIfVoiceVlanInd

`func (o *DiscoveryDeviceinterface) HasCapIfVoiceVlanInd() bool`

HasCapIfVoiceVlanInd returns a boolean if a field has been set.

### GetCapIfVoiceVlanNaReason

`func (o *DiscoveryDeviceinterface) GetCapIfVoiceVlanNaReason() string`

GetCapIfVoiceVlanNaReason returns the CapIfVoiceVlanNaReason field if non-nil, zero value otherwise.

### GetCapIfVoiceVlanNaReasonOk

`func (o *DiscoveryDeviceinterface) GetCapIfVoiceVlanNaReasonOk() (*string, bool)`

GetCapIfVoiceVlanNaReasonOk returns a tuple with the CapIfVoiceVlanNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapIfVoiceVlanNaReason

`func (o *DiscoveryDeviceinterface) SetCapIfVoiceVlanNaReason(v string)`

SetCapIfVoiceVlanNaReason sets CapIfVoiceVlanNaReason field to given value.

### HasCapIfVoiceVlanNaReason

`func (o *DiscoveryDeviceinterface) HasCapIfVoiceVlanNaReason() bool`

HasCapIfVoiceVlanNaReason returns a boolean if a field has been set.

### GetDescription

`func (o *DiscoveryDeviceinterface) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveryDeviceinterface) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveryDeviceinterface) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveryDeviceinterface) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDescriptionTaskInfo

`func (o *DiscoveryDeviceinterface) GetDescriptionTaskInfo() DiscoveryDeviceinterfaceDescriptionTaskInfo`

GetDescriptionTaskInfo returns the DescriptionTaskInfo field if non-nil, zero value otherwise.

### GetDescriptionTaskInfoOk

`func (o *DiscoveryDeviceinterface) GetDescriptionTaskInfoOk() (*DiscoveryDeviceinterfaceDescriptionTaskInfo, bool)`

GetDescriptionTaskInfoOk returns a tuple with the DescriptionTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionTaskInfo

`func (o *DiscoveryDeviceinterface) SetDescriptionTaskInfo(v DiscoveryDeviceinterfaceDescriptionTaskInfo)`

SetDescriptionTaskInfo sets DescriptionTaskInfo field to given value.

### HasDescriptionTaskInfo

`func (o *DiscoveryDeviceinterface) HasDescriptionTaskInfo() bool`

HasDescriptionTaskInfo returns a boolean if a field has been set.

### GetDevice

`func (o *DiscoveryDeviceinterface) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiscoveryDeviceinterface) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiscoveryDeviceinterface) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiscoveryDeviceinterface) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDuplex

`func (o *DiscoveryDeviceinterface) GetDuplex() string`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *DiscoveryDeviceinterface) GetDuplexOk() (*string, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *DiscoveryDeviceinterface) SetDuplex(v string)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *DiscoveryDeviceinterface) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetExtattrs

`func (o *DiscoveryDeviceinterface) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *DiscoveryDeviceinterface) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *DiscoveryDeviceinterface) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *DiscoveryDeviceinterface) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetIfaddrInfos

`func (o *DiscoveryDeviceinterface) GetIfaddrInfos() []DiscoveryDeviceinterfaceIfaddrInfos`

GetIfaddrInfos returns the IfaddrInfos field if non-nil, zero value otherwise.

### GetIfaddrInfosOk

`func (o *DiscoveryDeviceinterface) GetIfaddrInfosOk() (*[]DiscoveryDeviceinterfaceIfaddrInfos, bool)`

GetIfaddrInfosOk returns a tuple with the IfaddrInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIfaddrInfos

`func (o *DiscoveryDeviceinterface) SetIfaddrInfos(v []DiscoveryDeviceinterfaceIfaddrInfos)`

SetIfaddrInfos sets IfaddrInfos field to given value.

### HasIfaddrInfos

`func (o *DiscoveryDeviceinterface) HasIfaddrInfos() bool`

HasIfaddrInfos returns a boolean if a field has been set.

### GetIndex

`func (o *DiscoveryDeviceinterface) GetIndex() int64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DiscoveryDeviceinterface) GetIndexOk() (*int64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DiscoveryDeviceinterface) SetIndex(v int64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *DiscoveryDeviceinterface) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetLastChange

`func (o *DiscoveryDeviceinterface) GetLastChange() int64`

GetLastChange returns the LastChange field if non-nil, zero value otherwise.

### GetLastChangeOk

`func (o *DiscoveryDeviceinterface) GetLastChangeOk() (*int64, bool)`

GetLastChangeOk returns a tuple with the LastChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChange

`func (o *DiscoveryDeviceinterface) SetLastChange(v int64)`

SetLastChange sets LastChange field to given value.

### HasLastChange

`func (o *DiscoveryDeviceinterface) HasLastChange() bool`

HasLastChange returns a boolean if a field has been set.

### GetLinkAggregation

`func (o *DiscoveryDeviceinterface) GetLinkAggregation() bool`

GetLinkAggregation returns the LinkAggregation field if non-nil, zero value otherwise.

### GetLinkAggregationOk

`func (o *DiscoveryDeviceinterface) GetLinkAggregationOk() (*bool, bool)`

GetLinkAggregationOk returns a tuple with the LinkAggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAggregation

`func (o *DiscoveryDeviceinterface) SetLinkAggregation(v bool)`

SetLinkAggregation sets LinkAggregation field to given value.

### HasLinkAggregation

`func (o *DiscoveryDeviceinterface) HasLinkAggregation() bool`

HasLinkAggregation returns a boolean if a field has been set.

### GetMac

`func (o *DiscoveryDeviceinterface) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DiscoveryDeviceinterface) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DiscoveryDeviceinterface) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *DiscoveryDeviceinterface) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *DiscoveryDeviceinterface) GetMsAdUserData() DiscoveryDeviceinterfaceMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *DiscoveryDeviceinterface) GetMsAdUserDataOk() (*DiscoveryDeviceinterfaceMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *DiscoveryDeviceinterface) SetMsAdUserData(v DiscoveryDeviceinterfaceMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *DiscoveryDeviceinterface) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *DiscoveryDeviceinterface) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveryDeviceinterface) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveryDeviceinterface) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscoveryDeviceinterface) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *DiscoveryDeviceinterface) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *DiscoveryDeviceinterface) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *DiscoveryDeviceinterface) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *DiscoveryDeviceinterface) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetOperStatus

`func (o *DiscoveryDeviceinterface) GetOperStatus() string`

GetOperStatus returns the OperStatus field if non-nil, zero value otherwise.

### GetOperStatusOk

`func (o *DiscoveryDeviceinterface) GetOperStatusOk() (*string, bool)`

GetOperStatusOk returns a tuple with the OperStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperStatus

`func (o *DiscoveryDeviceinterface) SetOperStatus(v string)`

SetOperStatus sets OperStatus field to given value.

### HasOperStatus

`func (o *DiscoveryDeviceinterface) HasOperStatus() bool`

HasOperStatus returns a boolean if a field has been set.

### GetPortFast

`func (o *DiscoveryDeviceinterface) GetPortFast() string`

GetPortFast returns the PortFast field if non-nil, zero value otherwise.

### GetPortFastOk

`func (o *DiscoveryDeviceinterface) GetPortFastOk() (*string, bool)`

GetPortFastOk returns a tuple with the PortFast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortFast

`func (o *DiscoveryDeviceinterface) SetPortFast(v string)`

SetPortFast sets PortFast field to given value.

### HasPortFast

`func (o *DiscoveryDeviceinterface) HasPortFast() bool`

HasPortFast returns a boolean if a field has been set.

### GetReservedObject

`func (o *DiscoveryDeviceinterface) GetReservedObject() string`

GetReservedObject returns the ReservedObject field if non-nil, zero value otherwise.

### GetReservedObjectOk

`func (o *DiscoveryDeviceinterface) GetReservedObjectOk() (*string, bool)`

GetReservedObjectOk returns a tuple with the ReservedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedObject

`func (o *DiscoveryDeviceinterface) SetReservedObject(v string)`

SetReservedObject sets ReservedObject field to given value.

### HasReservedObject

`func (o *DiscoveryDeviceinterface) HasReservedObject() bool`

HasReservedObject returns a boolean if a field has been set.

### GetSpeed

`func (o *DiscoveryDeviceinterface) GetSpeed() int64`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *DiscoveryDeviceinterface) GetSpeedOk() (*int64, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *DiscoveryDeviceinterface) SetSpeed(v int64)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *DiscoveryDeviceinterface) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetTrunkStatus

`func (o *DiscoveryDeviceinterface) GetTrunkStatus() string`

GetTrunkStatus returns the TrunkStatus field if non-nil, zero value otherwise.

### GetTrunkStatusOk

`func (o *DiscoveryDeviceinterface) GetTrunkStatusOk() (*string, bool)`

GetTrunkStatusOk returns a tuple with the TrunkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunkStatus

`func (o *DiscoveryDeviceinterface) SetTrunkStatus(v string)`

SetTrunkStatus sets TrunkStatus field to given value.

### HasTrunkStatus

`func (o *DiscoveryDeviceinterface) HasTrunkStatus() bool`

HasTrunkStatus returns a boolean if a field has been set.

### GetType

`func (o *DiscoveryDeviceinterface) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveryDeviceinterface) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveryDeviceinterface) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveryDeviceinterface) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVlanInfoTaskInfo

`func (o *DiscoveryDeviceinterface) GetVlanInfoTaskInfo() DiscoveryDeviceinterfaceVlanInfoTaskInfo`

GetVlanInfoTaskInfo returns the VlanInfoTaskInfo field if non-nil, zero value otherwise.

### GetVlanInfoTaskInfoOk

`func (o *DiscoveryDeviceinterface) GetVlanInfoTaskInfoOk() (*DiscoveryDeviceinterfaceVlanInfoTaskInfo, bool)`

GetVlanInfoTaskInfoOk returns a tuple with the VlanInfoTaskInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInfoTaskInfo

`func (o *DiscoveryDeviceinterface) SetVlanInfoTaskInfo(v DiscoveryDeviceinterfaceVlanInfoTaskInfo)`

SetVlanInfoTaskInfo sets VlanInfoTaskInfo field to given value.

### HasVlanInfoTaskInfo

`func (o *DiscoveryDeviceinterface) HasVlanInfoTaskInfo() bool`

HasVlanInfoTaskInfo returns a boolean if a field has been set.

### GetVlanInfos

`func (o *DiscoveryDeviceinterface) GetVlanInfos() []DiscoveryDeviceinterfaceVlanInfos`

GetVlanInfos returns the VlanInfos field if non-nil, zero value otherwise.

### GetVlanInfosOk

`func (o *DiscoveryDeviceinterface) GetVlanInfosOk() (*[]DiscoveryDeviceinterfaceVlanInfos, bool)`

GetVlanInfosOk returns a tuple with the VlanInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInfos

`func (o *DiscoveryDeviceinterface) SetVlanInfos(v []DiscoveryDeviceinterfaceVlanInfos)`

SetVlanInfos sets VlanInfos field to given value.

### HasVlanInfos

`func (o *DiscoveryDeviceinterface) HasVlanInfos() bool`

HasVlanInfos returns a boolean if a field has been set.

### GetVpcPeer

`func (o *DiscoveryDeviceinterface) GetVpcPeer() string`

GetVpcPeer returns the VpcPeer field if non-nil, zero value otherwise.

### GetVpcPeerOk

`func (o *DiscoveryDeviceinterface) GetVpcPeerOk() (*string, bool)`

GetVpcPeerOk returns a tuple with the VpcPeer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeer

`func (o *DiscoveryDeviceinterface) SetVpcPeer(v string)`

SetVpcPeer sets VpcPeer field to given value.

### HasVpcPeer

`func (o *DiscoveryDeviceinterface) HasVpcPeer() bool`

HasVpcPeer returns a boolean if a field has been set.

### GetVpcPeerDevice

`func (o *DiscoveryDeviceinterface) GetVpcPeerDevice() string`

GetVpcPeerDevice returns the VpcPeerDevice field if non-nil, zero value otherwise.

### GetVpcPeerDeviceOk

`func (o *DiscoveryDeviceinterface) GetVpcPeerDeviceOk() (*string, bool)`

GetVpcPeerDeviceOk returns a tuple with the VpcPeerDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcPeerDevice

`func (o *DiscoveryDeviceinterface) SetVpcPeerDevice(v string)`

SetVpcPeerDevice sets VpcPeerDevice field to given value.

### HasVpcPeerDevice

`func (o *DiscoveryDeviceinterface) HasVpcPeerDevice() bool`

HasVpcPeerDevice returns a boolean if a field has been set.

### GetVrfDescription

`func (o *DiscoveryDeviceinterface) GetVrfDescription() string`

GetVrfDescription returns the VrfDescription field if non-nil, zero value otherwise.

### GetVrfDescriptionOk

`func (o *DiscoveryDeviceinterface) GetVrfDescriptionOk() (*string, bool)`

GetVrfDescriptionOk returns a tuple with the VrfDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfDescription

`func (o *DiscoveryDeviceinterface) SetVrfDescription(v string)`

SetVrfDescription sets VrfDescription field to given value.

### HasVrfDescription

`func (o *DiscoveryDeviceinterface) HasVrfDescription() bool`

HasVrfDescription returns a boolean if a field has been set.

### GetVrfName

`func (o *DiscoveryDeviceinterface) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *DiscoveryDeviceinterface) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *DiscoveryDeviceinterface) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *DiscoveryDeviceinterface) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.

### GetVrfRd

`func (o *DiscoveryDeviceinterface) GetVrfRd() string`

GetVrfRd returns the VrfRd field if non-nil, zero value otherwise.

### GetVrfRdOk

`func (o *DiscoveryDeviceinterface) GetVrfRdOk() (*string, bool)`

GetVrfRdOk returns a tuple with the VrfRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfRd

`func (o *DiscoveryDeviceinterface) SetVrfRd(v string)`

SetVrfRd sets VrfRd field to given value.

### HasVrfRd

`func (o *DiscoveryDeviceinterface) HasVrfRd() bool`

HasVrfRd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


