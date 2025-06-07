# DiscoveryDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Address** | Pointer to **string** | The IPv4 Address or IPv6 Address of the device. | [optional] [readonly] 
**AddressRef** | Pointer to **string** | The ref to management IP address of the device. | [optional] [readonly] 
**AvailableMgmtIps** | Pointer to **[]string** | The list of available management IPs for the device. | [optional] [readonly] 
**CapAdminStatusInd** | Pointer to **bool** | Determines whether to modify the admin status of an interface of the device. | [optional] [readonly] 
**CapAdminStatusNaReason** | Pointer to **string** | The reason that the edit admin status action is not available. | [optional] [readonly] 
**CapDescriptionInd** | Pointer to **bool** | Determines whether to modify the description of an interface on the device. | [optional] [readonly] 
**CapDescriptionNaReason** | Pointer to **string** | The reason that the edit description action is not available. | [optional] [readonly] 
**CapNetDeprovisioningInd** | Pointer to **bool** | Determines whether to deprovision a network from interfaces of the device. | [optional] [readonly] 
**CapNetDeprovisioningNaReason** | Pointer to **string** | The reason that the deprovision a network from interfaces of this device is not available. | [optional] [readonly] 
**CapNetProvisioningInd** | Pointer to **bool** | Determines whether to modify the network associated to an interface of the device. | [optional] [readonly] 
**CapNetProvisioningNaReason** | Pointer to **string** | The reason that network provisioning is not available. | [optional] [readonly] 
**CapNetVlanProvisioningInd** | Pointer to **bool** | Determines whether to create a VLAN and then provision a network to the interface of the device. | [optional] [readonly] 
**CapNetVlanProvisioningNaReason** | Pointer to **string** | The reason that network provisioning on VLAN is not available. | [optional] [readonly] 
**CapVlanAssignmentInd** | Pointer to **bool** | Determines whether to modify the VLAN assignement of an interface of the device. | [optional] [readonly] 
**CapVlanAssignmentNaReason** | Pointer to **string** | The reason that VLAN assignment action is not available. | [optional] [readonly] 
**CapVoiceVlanInd** | Pointer to **bool** | Determines whether to modify the voice VLAN assignment of an interface of the device. | [optional] [readonly] 
**CapVoiceVlanNaReason** | Pointer to **string** | The reason that voice VLAN assignment action is not available. | [optional] [readonly] 
**ChassisSerialNumber** | Pointer to **string** | The device chassis serial number. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the device. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interfaces** | Pointer to **[]map[string]interface{}** | List of the device interfaces. | [optional] [readonly] 
**Location** | Pointer to **string** | The location of the device. | [optional] [readonly] 
**Model** | Pointer to **string** | The model name of the device. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**DiscoveryDeviceMsAdUserData**](DiscoveryDeviceMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the device. | [optional] [readonly] 
**Neighbors** | Pointer to **[]map[string]interface{}** | List of the device neighbors. | [optional] [readonly] 
**Network** | Pointer to **string** | The ref to the network to which belongs the management IP address belongs. | [optional] [readonly] 
**NetworkInfos** | Pointer to [**[]DiscoveryDeviceNetworkInfos**](DiscoveryDeviceNetworkInfos.md) | The list of networks to which the device interfaces belong. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this device resides. | [optional] [readonly] 
**Networks** | Pointer to **[]map[string]interface{}** | The list of networks to which the device interfaces belong. | [optional] [readonly] 
**OsVersion** | Pointer to **string** | The Operating System version running on the device. | [optional] [readonly] 
**PortStats** | Pointer to [**DiscoveryDevicePortStats**](DiscoveryDevicePortStats.md) |  | [optional] 
**PrivilegedPolling** | Pointer to **bool** | A flag indicated that NI should send enable command when interacting with device. | [optional] 
**Type** | Pointer to **string** | The type of the device. | [optional] [readonly] 
**UserDefinedMgmtIp** | Pointer to **string** | User-defined management IP address of the device. | [optional] 
**Vendor** | Pointer to **string** | The vendor name of the device. | [optional] [readonly] 
**VlanInfos** | Pointer to [**[]DiscoveryDeviceVlanInfos**](DiscoveryDeviceVlanInfos.md) | The list of VLAN information associated with the device. | [optional] [readonly] 

## Methods

### NewDiscoveryDevice

`func NewDiscoveryDevice() *DiscoveryDevice`

NewDiscoveryDevice instantiates a new DiscoveryDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryDeviceWithDefaults

`func NewDiscoveryDeviceWithDefaults() *DiscoveryDevice`

NewDiscoveryDeviceWithDefaults instantiates a new DiscoveryDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DiscoveryDevice) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DiscoveryDevice) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DiscoveryDevice) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DiscoveryDevice) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *DiscoveryDevice) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DiscoveryDevice) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DiscoveryDevice) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DiscoveryDevice) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressRef

`func (o *DiscoveryDevice) GetAddressRef() string`

GetAddressRef returns the AddressRef field if non-nil, zero value otherwise.

### GetAddressRefOk

`func (o *DiscoveryDevice) GetAddressRefOk() (*string, bool)`

GetAddressRefOk returns a tuple with the AddressRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRef

`func (o *DiscoveryDevice) SetAddressRef(v string)`

SetAddressRef sets AddressRef field to given value.

### HasAddressRef

`func (o *DiscoveryDevice) HasAddressRef() bool`

HasAddressRef returns a boolean if a field has been set.

### GetAvailableMgmtIps

`func (o *DiscoveryDevice) GetAvailableMgmtIps() []string`

GetAvailableMgmtIps returns the AvailableMgmtIps field if non-nil, zero value otherwise.

### GetAvailableMgmtIpsOk

`func (o *DiscoveryDevice) GetAvailableMgmtIpsOk() (*[]string, bool)`

GetAvailableMgmtIpsOk returns a tuple with the AvailableMgmtIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMgmtIps

`func (o *DiscoveryDevice) SetAvailableMgmtIps(v []string)`

SetAvailableMgmtIps sets AvailableMgmtIps field to given value.

### HasAvailableMgmtIps

`func (o *DiscoveryDevice) HasAvailableMgmtIps() bool`

HasAvailableMgmtIps returns a boolean if a field has been set.

### GetCapAdminStatusInd

`func (o *DiscoveryDevice) GetCapAdminStatusInd() bool`

GetCapAdminStatusInd returns the CapAdminStatusInd field if non-nil, zero value otherwise.

### GetCapAdminStatusIndOk

`func (o *DiscoveryDevice) GetCapAdminStatusIndOk() (*bool, bool)`

GetCapAdminStatusIndOk returns a tuple with the CapAdminStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapAdminStatusInd

`func (o *DiscoveryDevice) SetCapAdminStatusInd(v bool)`

SetCapAdminStatusInd sets CapAdminStatusInd field to given value.

### HasCapAdminStatusInd

`func (o *DiscoveryDevice) HasCapAdminStatusInd() bool`

HasCapAdminStatusInd returns a boolean if a field has been set.

### GetCapAdminStatusNaReason

`func (o *DiscoveryDevice) GetCapAdminStatusNaReason() string`

GetCapAdminStatusNaReason returns the CapAdminStatusNaReason field if non-nil, zero value otherwise.

### GetCapAdminStatusNaReasonOk

`func (o *DiscoveryDevice) GetCapAdminStatusNaReasonOk() (*string, bool)`

GetCapAdminStatusNaReasonOk returns a tuple with the CapAdminStatusNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapAdminStatusNaReason

`func (o *DiscoveryDevice) SetCapAdminStatusNaReason(v string)`

SetCapAdminStatusNaReason sets CapAdminStatusNaReason field to given value.

### HasCapAdminStatusNaReason

`func (o *DiscoveryDevice) HasCapAdminStatusNaReason() bool`

HasCapAdminStatusNaReason returns a boolean if a field has been set.

### GetCapDescriptionInd

`func (o *DiscoveryDevice) GetCapDescriptionInd() bool`

GetCapDescriptionInd returns the CapDescriptionInd field if non-nil, zero value otherwise.

### GetCapDescriptionIndOk

`func (o *DiscoveryDevice) GetCapDescriptionIndOk() (*bool, bool)`

GetCapDescriptionIndOk returns a tuple with the CapDescriptionInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapDescriptionInd

`func (o *DiscoveryDevice) SetCapDescriptionInd(v bool)`

SetCapDescriptionInd sets CapDescriptionInd field to given value.

### HasCapDescriptionInd

`func (o *DiscoveryDevice) HasCapDescriptionInd() bool`

HasCapDescriptionInd returns a boolean if a field has been set.

### GetCapDescriptionNaReason

`func (o *DiscoveryDevice) GetCapDescriptionNaReason() string`

GetCapDescriptionNaReason returns the CapDescriptionNaReason field if non-nil, zero value otherwise.

### GetCapDescriptionNaReasonOk

`func (o *DiscoveryDevice) GetCapDescriptionNaReasonOk() (*string, bool)`

GetCapDescriptionNaReasonOk returns a tuple with the CapDescriptionNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapDescriptionNaReason

`func (o *DiscoveryDevice) SetCapDescriptionNaReason(v string)`

SetCapDescriptionNaReason sets CapDescriptionNaReason field to given value.

### HasCapDescriptionNaReason

`func (o *DiscoveryDevice) HasCapDescriptionNaReason() bool`

HasCapDescriptionNaReason returns a boolean if a field has been set.

### GetCapNetDeprovisioningInd

`func (o *DiscoveryDevice) GetCapNetDeprovisioningInd() bool`

GetCapNetDeprovisioningInd returns the CapNetDeprovisioningInd field if non-nil, zero value otherwise.

### GetCapNetDeprovisioningIndOk

`func (o *DiscoveryDevice) GetCapNetDeprovisioningIndOk() (*bool, bool)`

GetCapNetDeprovisioningIndOk returns a tuple with the CapNetDeprovisioningInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetDeprovisioningInd

`func (o *DiscoveryDevice) SetCapNetDeprovisioningInd(v bool)`

SetCapNetDeprovisioningInd sets CapNetDeprovisioningInd field to given value.

### HasCapNetDeprovisioningInd

`func (o *DiscoveryDevice) HasCapNetDeprovisioningInd() bool`

HasCapNetDeprovisioningInd returns a boolean if a field has been set.

### GetCapNetDeprovisioningNaReason

`func (o *DiscoveryDevice) GetCapNetDeprovisioningNaReason() string`

GetCapNetDeprovisioningNaReason returns the CapNetDeprovisioningNaReason field if non-nil, zero value otherwise.

### GetCapNetDeprovisioningNaReasonOk

`func (o *DiscoveryDevice) GetCapNetDeprovisioningNaReasonOk() (*string, bool)`

GetCapNetDeprovisioningNaReasonOk returns a tuple with the CapNetDeprovisioningNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetDeprovisioningNaReason

`func (o *DiscoveryDevice) SetCapNetDeprovisioningNaReason(v string)`

SetCapNetDeprovisioningNaReason sets CapNetDeprovisioningNaReason field to given value.

### HasCapNetDeprovisioningNaReason

`func (o *DiscoveryDevice) HasCapNetDeprovisioningNaReason() bool`

HasCapNetDeprovisioningNaReason returns a boolean if a field has been set.

### GetCapNetProvisioningInd

`func (o *DiscoveryDevice) GetCapNetProvisioningInd() bool`

GetCapNetProvisioningInd returns the CapNetProvisioningInd field if non-nil, zero value otherwise.

### GetCapNetProvisioningIndOk

`func (o *DiscoveryDevice) GetCapNetProvisioningIndOk() (*bool, bool)`

GetCapNetProvisioningIndOk returns a tuple with the CapNetProvisioningInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetProvisioningInd

`func (o *DiscoveryDevice) SetCapNetProvisioningInd(v bool)`

SetCapNetProvisioningInd sets CapNetProvisioningInd field to given value.

### HasCapNetProvisioningInd

`func (o *DiscoveryDevice) HasCapNetProvisioningInd() bool`

HasCapNetProvisioningInd returns a boolean if a field has been set.

### GetCapNetProvisioningNaReason

`func (o *DiscoveryDevice) GetCapNetProvisioningNaReason() string`

GetCapNetProvisioningNaReason returns the CapNetProvisioningNaReason field if non-nil, zero value otherwise.

### GetCapNetProvisioningNaReasonOk

`func (o *DiscoveryDevice) GetCapNetProvisioningNaReasonOk() (*string, bool)`

GetCapNetProvisioningNaReasonOk returns a tuple with the CapNetProvisioningNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetProvisioningNaReason

`func (o *DiscoveryDevice) SetCapNetProvisioningNaReason(v string)`

SetCapNetProvisioningNaReason sets CapNetProvisioningNaReason field to given value.

### HasCapNetProvisioningNaReason

`func (o *DiscoveryDevice) HasCapNetProvisioningNaReason() bool`

HasCapNetProvisioningNaReason returns a boolean if a field has been set.

### GetCapNetVlanProvisioningInd

`func (o *DiscoveryDevice) GetCapNetVlanProvisioningInd() bool`

GetCapNetVlanProvisioningInd returns the CapNetVlanProvisioningInd field if non-nil, zero value otherwise.

### GetCapNetVlanProvisioningIndOk

`func (o *DiscoveryDevice) GetCapNetVlanProvisioningIndOk() (*bool, bool)`

GetCapNetVlanProvisioningIndOk returns a tuple with the CapNetVlanProvisioningInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetVlanProvisioningInd

`func (o *DiscoveryDevice) SetCapNetVlanProvisioningInd(v bool)`

SetCapNetVlanProvisioningInd sets CapNetVlanProvisioningInd field to given value.

### HasCapNetVlanProvisioningInd

`func (o *DiscoveryDevice) HasCapNetVlanProvisioningInd() bool`

HasCapNetVlanProvisioningInd returns a boolean if a field has been set.

### GetCapNetVlanProvisioningNaReason

`func (o *DiscoveryDevice) GetCapNetVlanProvisioningNaReason() string`

GetCapNetVlanProvisioningNaReason returns the CapNetVlanProvisioningNaReason field if non-nil, zero value otherwise.

### GetCapNetVlanProvisioningNaReasonOk

`func (o *DiscoveryDevice) GetCapNetVlanProvisioningNaReasonOk() (*string, bool)`

GetCapNetVlanProvisioningNaReasonOk returns a tuple with the CapNetVlanProvisioningNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetVlanProvisioningNaReason

`func (o *DiscoveryDevice) SetCapNetVlanProvisioningNaReason(v string)`

SetCapNetVlanProvisioningNaReason sets CapNetVlanProvisioningNaReason field to given value.

### HasCapNetVlanProvisioningNaReason

`func (o *DiscoveryDevice) HasCapNetVlanProvisioningNaReason() bool`

HasCapNetVlanProvisioningNaReason returns a boolean if a field has been set.

### GetCapVlanAssignmentInd

`func (o *DiscoveryDevice) GetCapVlanAssignmentInd() bool`

GetCapVlanAssignmentInd returns the CapVlanAssignmentInd field if non-nil, zero value otherwise.

### GetCapVlanAssignmentIndOk

`func (o *DiscoveryDevice) GetCapVlanAssignmentIndOk() (*bool, bool)`

GetCapVlanAssignmentIndOk returns a tuple with the CapVlanAssignmentInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapVlanAssignmentInd

`func (o *DiscoveryDevice) SetCapVlanAssignmentInd(v bool)`

SetCapVlanAssignmentInd sets CapVlanAssignmentInd field to given value.

### HasCapVlanAssignmentInd

`func (o *DiscoveryDevice) HasCapVlanAssignmentInd() bool`

HasCapVlanAssignmentInd returns a boolean if a field has been set.

### GetCapVlanAssignmentNaReason

`func (o *DiscoveryDevice) GetCapVlanAssignmentNaReason() string`

GetCapVlanAssignmentNaReason returns the CapVlanAssignmentNaReason field if non-nil, zero value otherwise.

### GetCapVlanAssignmentNaReasonOk

`func (o *DiscoveryDevice) GetCapVlanAssignmentNaReasonOk() (*string, bool)`

GetCapVlanAssignmentNaReasonOk returns a tuple with the CapVlanAssignmentNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapVlanAssignmentNaReason

`func (o *DiscoveryDevice) SetCapVlanAssignmentNaReason(v string)`

SetCapVlanAssignmentNaReason sets CapVlanAssignmentNaReason field to given value.

### HasCapVlanAssignmentNaReason

`func (o *DiscoveryDevice) HasCapVlanAssignmentNaReason() bool`

HasCapVlanAssignmentNaReason returns a boolean if a field has been set.

### GetCapVoiceVlanInd

`func (o *DiscoveryDevice) GetCapVoiceVlanInd() bool`

GetCapVoiceVlanInd returns the CapVoiceVlanInd field if non-nil, zero value otherwise.

### GetCapVoiceVlanIndOk

`func (o *DiscoveryDevice) GetCapVoiceVlanIndOk() (*bool, bool)`

GetCapVoiceVlanIndOk returns a tuple with the CapVoiceVlanInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapVoiceVlanInd

`func (o *DiscoveryDevice) SetCapVoiceVlanInd(v bool)`

SetCapVoiceVlanInd sets CapVoiceVlanInd field to given value.

### HasCapVoiceVlanInd

`func (o *DiscoveryDevice) HasCapVoiceVlanInd() bool`

HasCapVoiceVlanInd returns a boolean if a field has been set.

### GetCapVoiceVlanNaReason

`func (o *DiscoveryDevice) GetCapVoiceVlanNaReason() string`

GetCapVoiceVlanNaReason returns the CapVoiceVlanNaReason field if non-nil, zero value otherwise.

### GetCapVoiceVlanNaReasonOk

`func (o *DiscoveryDevice) GetCapVoiceVlanNaReasonOk() (*string, bool)`

GetCapVoiceVlanNaReasonOk returns a tuple with the CapVoiceVlanNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapVoiceVlanNaReason

`func (o *DiscoveryDevice) SetCapVoiceVlanNaReason(v string)`

SetCapVoiceVlanNaReason sets CapVoiceVlanNaReason field to given value.

### HasCapVoiceVlanNaReason

`func (o *DiscoveryDevice) HasCapVoiceVlanNaReason() bool`

HasCapVoiceVlanNaReason returns a boolean if a field has been set.

### GetChassisSerialNumber

`func (o *DiscoveryDevice) GetChassisSerialNumber() string`

GetChassisSerialNumber returns the ChassisSerialNumber field if non-nil, zero value otherwise.

### GetChassisSerialNumberOk

`func (o *DiscoveryDevice) GetChassisSerialNumberOk() (*string, bool)`

GetChassisSerialNumberOk returns a tuple with the ChassisSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerialNumber

`func (o *DiscoveryDevice) SetChassisSerialNumber(v string)`

SetChassisSerialNumber sets ChassisSerialNumber field to given value.

### HasChassisSerialNumber

`func (o *DiscoveryDevice) HasChassisSerialNumber() bool`

HasChassisSerialNumber returns a boolean if a field has been set.

### GetDescription

`func (o *DiscoveryDevice) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveryDevice) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveryDevice) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveryDevice) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtattrs

`func (o *DiscoveryDevice) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *DiscoveryDevice) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *DiscoveryDevice) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *DiscoveryDevice) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetInterfaces

`func (o *DiscoveryDevice) GetInterfaces() []map[string]interface{}`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *DiscoveryDevice) GetInterfacesOk() (*[]map[string]interface{}, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *DiscoveryDevice) SetInterfaces(v []map[string]interface{})`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *DiscoveryDevice) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLocation

`func (o *DiscoveryDevice) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DiscoveryDevice) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DiscoveryDevice) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DiscoveryDevice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetModel

`func (o *DiscoveryDevice) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DiscoveryDevice) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DiscoveryDevice) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DiscoveryDevice) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *DiscoveryDevice) GetMsAdUserData() DiscoveryDeviceMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *DiscoveryDevice) GetMsAdUserDataOk() (*DiscoveryDeviceMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *DiscoveryDevice) SetMsAdUserData(v DiscoveryDeviceMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *DiscoveryDevice) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *DiscoveryDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveryDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveryDevice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscoveryDevice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeighbors

`func (o *DiscoveryDevice) GetNeighbors() []map[string]interface{}`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *DiscoveryDevice) GetNeighborsOk() (*[]map[string]interface{}, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *DiscoveryDevice) SetNeighbors(v []map[string]interface{})`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *DiscoveryDevice) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.

### GetNetwork

`func (o *DiscoveryDevice) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *DiscoveryDevice) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *DiscoveryDevice) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *DiscoveryDevice) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkInfos

`func (o *DiscoveryDevice) GetNetworkInfos() []DiscoveryDeviceNetworkInfos`

GetNetworkInfos returns the NetworkInfos field if non-nil, zero value otherwise.

### GetNetworkInfosOk

`func (o *DiscoveryDevice) GetNetworkInfosOk() (*[]DiscoveryDeviceNetworkInfos, bool)`

GetNetworkInfosOk returns a tuple with the NetworkInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfos

`func (o *DiscoveryDevice) SetNetworkInfos(v []DiscoveryDeviceNetworkInfos)`

SetNetworkInfos sets NetworkInfos field to given value.

### HasNetworkInfos

`func (o *DiscoveryDevice) HasNetworkInfos() bool`

HasNetworkInfos returns a boolean if a field has been set.

### GetNetworkView

`func (o *DiscoveryDevice) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *DiscoveryDevice) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *DiscoveryDevice) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *DiscoveryDevice) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworks

`func (o *DiscoveryDevice) GetNetworks() []map[string]interface{}`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *DiscoveryDevice) GetNetworksOk() (*[]map[string]interface{}, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *DiscoveryDevice) SetNetworks(v []map[string]interface{})`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *DiscoveryDevice) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOsVersion

`func (o *DiscoveryDevice) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *DiscoveryDevice) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *DiscoveryDevice) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *DiscoveryDevice) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPortStats

`func (o *DiscoveryDevice) GetPortStats() DiscoveryDevicePortStats`

GetPortStats returns the PortStats field if non-nil, zero value otherwise.

### GetPortStatsOk

`func (o *DiscoveryDevice) GetPortStatsOk() (*DiscoveryDevicePortStats, bool)`

GetPortStatsOk returns a tuple with the PortStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStats

`func (o *DiscoveryDevice) SetPortStats(v DiscoveryDevicePortStats)`

SetPortStats sets PortStats field to given value.

### HasPortStats

`func (o *DiscoveryDevice) HasPortStats() bool`

HasPortStats returns a boolean if a field has been set.

### GetPrivilegedPolling

`func (o *DiscoveryDevice) GetPrivilegedPolling() bool`

GetPrivilegedPolling returns the PrivilegedPolling field if non-nil, zero value otherwise.

### GetPrivilegedPollingOk

`func (o *DiscoveryDevice) GetPrivilegedPollingOk() (*bool, bool)`

GetPrivilegedPollingOk returns a tuple with the PrivilegedPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedPolling

`func (o *DiscoveryDevice) SetPrivilegedPolling(v bool)`

SetPrivilegedPolling sets PrivilegedPolling field to given value.

### HasPrivilegedPolling

`func (o *DiscoveryDevice) HasPrivilegedPolling() bool`

HasPrivilegedPolling returns a boolean if a field has been set.

### GetType

`func (o *DiscoveryDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DiscoveryDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DiscoveryDevice) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DiscoveryDevice) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserDefinedMgmtIp

`func (o *DiscoveryDevice) GetUserDefinedMgmtIp() string`

GetUserDefinedMgmtIp returns the UserDefinedMgmtIp field if non-nil, zero value otherwise.

### GetUserDefinedMgmtIpOk

`func (o *DiscoveryDevice) GetUserDefinedMgmtIpOk() (*string, bool)`

GetUserDefinedMgmtIpOk returns a tuple with the UserDefinedMgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedMgmtIp

`func (o *DiscoveryDevice) SetUserDefinedMgmtIp(v string)`

SetUserDefinedMgmtIp sets UserDefinedMgmtIp field to given value.

### HasUserDefinedMgmtIp

`func (o *DiscoveryDevice) HasUserDefinedMgmtIp() bool`

HasUserDefinedMgmtIp returns a boolean if a field has been set.

### GetVendor

`func (o *DiscoveryDevice) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *DiscoveryDevice) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *DiscoveryDevice) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *DiscoveryDevice) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVlanInfos

`func (o *DiscoveryDevice) GetVlanInfos() []DiscoveryDeviceVlanInfos`

GetVlanInfos returns the VlanInfos field if non-nil, zero value otherwise.

### GetVlanInfosOk

`func (o *DiscoveryDevice) GetVlanInfosOk() (*[]DiscoveryDeviceVlanInfos, bool)`

GetVlanInfosOk returns a tuple with the VlanInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInfos

`func (o *DiscoveryDevice) SetVlanInfos(v []DiscoveryDeviceVlanInfos)`

SetVlanInfos sets VlanInfos field to given value.

### HasVlanInfos

`func (o *DiscoveryDevice) HasVlanInfos() bool`

HasVlanInfos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


