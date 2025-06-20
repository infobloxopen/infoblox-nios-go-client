# GetDiscoveryDeviceResponse

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
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Interfaces** | Pointer to **[]string** | List of the device interfaces. | [optional] [readonly] 
**Location** | Pointer to **string** | The location of the device. | [optional] [readonly] 
**Model** | Pointer to **string** | The model name of the device. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**DiscoveryDeviceMsAdUserData**](DiscoveryDeviceMsAdUserData.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the device. | [optional] [readonly] 
**Neighbors** | Pointer to **[]string** | List of the device neighbors. | [optional] [readonly] 
**Network** | Pointer to **string** | The ref to the network to which belongs the management IP address belongs. | [optional] [readonly] 
**NetworkInfos** | Pointer to [**[]DiscoveryDeviceNetworkInfos**](DiscoveryDeviceNetworkInfos.md) | The list of networks to which the device interfaces belong. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this device resides. | [optional] [readonly] 
**Networks** | Pointer to **[]string** | The list of networks to which the device interfaces belong. | [optional] [readonly] 
**OsVersion** | Pointer to **string** | The Operating System version running on the device. | [optional] [readonly] 
**PortStats** | Pointer to [**DiscoveryDevicePortStats**](DiscoveryDevicePortStats.md) |  | [optional] 
**PrivilegedPolling** | Pointer to **bool** | A flag indicated that NI should send enable command when interacting with device. | [optional] 
**Type** | Pointer to **string** | The type of the device. | [optional] [readonly] 
**UserDefinedMgmtIp** | Pointer to **string** | User-defined management IP address of the device. | [optional] 
**Vendor** | Pointer to **string** | The vendor name of the device. | [optional] [readonly] 
**VlanInfos** | Pointer to [**[]DiscoveryDeviceVlanInfos**](DiscoveryDeviceVlanInfos.md) | The list of VLAN information associated with the device. | [optional] [readonly] 
**Result** | Pointer to [**DiscoveryDevice**](DiscoveryDevice.md) |  | [optional] 

## Methods

### NewGetDiscoveryDeviceResponse

`func NewGetDiscoveryDeviceResponse() *GetDiscoveryDeviceResponse`

NewGetDiscoveryDeviceResponse instantiates a new GetDiscoveryDeviceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscoveryDeviceResponseWithDefaults

`func NewGetDiscoveryDeviceResponseWithDefaults() *GetDiscoveryDeviceResponse`

NewGetDiscoveryDeviceResponseWithDefaults instantiates a new GetDiscoveryDeviceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDiscoveryDeviceResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDiscoveryDeviceResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDiscoveryDeviceResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDiscoveryDeviceResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetDiscoveryDeviceResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetDiscoveryDeviceResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetDiscoveryDeviceResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetDiscoveryDeviceResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressRef

`func (o *GetDiscoveryDeviceResponse) GetAddressRef() string`

GetAddressRef returns the AddressRef field if non-nil, zero value otherwise.

### GetAddressRefOk

`func (o *GetDiscoveryDeviceResponse) GetAddressRefOk() (*string, bool)`

GetAddressRefOk returns a tuple with the AddressRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressRef

`func (o *GetDiscoveryDeviceResponse) SetAddressRef(v string)`

SetAddressRef sets AddressRef field to given value.

### HasAddressRef

`func (o *GetDiscoveryDeviceResponse) HasAddressRef() bool`

HasAddressRef returns a boolean if a field has been set.

### GetAvailableMgmtIps

`func (o *GetDiscoveryDeviceResponse) GetAvailableMgmtIps() []string`

GetAvailableMgmtIps returns the AvailableMgmtIps field if non-nil, zero value otherwise.

### GetAvailableMgmtIpsOk

`func (o *GetDiscoveryDeviceResponse) GetAvailableMgmtIpsOk() (*[]string, bool)`

GetAvailableMgmtIpsOk returns a tuple with the AvailableMgmtIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMgmtIps

`func (o *GetDiscoveryDeviceResponse) SetAvailableMgmtIps(v []string)`

SetAvailableMgmtIps sets AvailableMgmtIps field to given value.

### HasAvailableMgmtIps

`func (o *GetDiscoveryDeviceResponse) HasAvailableMgmtIps() bool`

HasAvailableMgmtIps returns a boolean if a field has been set.

### GetCapAdminStatusInd

`func (o *GetDiscoveryDeviceResponse) GetCapAdminStatusInd() bool`

GetCapAdminStatusInd returns the CapAdminStatusInd field if non-nil, zero value otherwise.

### GetCapAdminStatusIndOk

`func (o *GetDiscoveryDeviceResponse) GetCapAdminStatusIndOk() (*bool, bool)`

GetCapAdminStatusIndOk returns a tuple with the CapAdminStatusInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapAdminStatusInd

`func (o *GetDiscoveryDeviceResponse) SetCapAdminStatusInd(v bool)`

SetCapAdminStatusInd sets CapAdminStatusInd field to given value.

### HasCapAdminStatusInd

`func (o *GetDiscoveryDeviceResponse) HasCapAdminStatusInd() bool`

HasCapAdminStatusInd returns a boolean if a field has been set.

### GetCapAdminStatusNaReason

`func (o *GetDiscoveryDeviceResponse) GetCapAdminStatusNaReason() string`

GetCapAdminStatusNaReason returns the CapAdminStatusNaReason field if non-nil, zero value otherwise.

### GetCapAdminStatusNaReasonOk

`func (o *GetDiscoveryDeviceResponse) GetCapAdminStatusNaReasonOk() (*string, bool)`

GetCapAdminStatusNaReasonOk returns a tuple with the CapAdminStatusNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapAdminStatusNaReason

`func (o *GetDiscoveryDeviceResponse) SetCapAdminStatusNaReason(v string)`

SetCapAdminStatusNaReason sets CapAdminStatusNaReason field to given value.

### HasCapAdminStatusNaReason

`func (o *GetDiscoveryDeviceResponse) HasCapAdminStatusNaReason() bool`

HasCapAdminStatusNaReason returns a boolean if a field has been set.

### GetCapDescriptionInd

`func (o *GetDiscoveryDeviceResponse) GetCapDescriptionInd() bool`

GetCapDescriptionInd returns the CapDescriptionInd field if non-nil, zero value otherwise.

### GetCapDescriptionIndOk

`func (o *GetDiscoveryDeviceResponse) GetCapDescriptionIndOk() (*bool, bool)`

GetCapDescriptionIndOk returns a tuple with the CapDescriptionInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapDescriptionInd

`func (o *GetDiscoveryDeviceResponse) SetCapDescriptionInd(v bool)`

SetCapDescriptionInd sets CapDescriptionInd field to given value.

### HasCapDescriptionInd

`func (o *GetDiscoveryDeviceResponse) HasCapDescriptionInd() bool`

HasCapDescriptionInd returns a boolean if a field has been set.

### GetCapDescriptionNaReason

`func (o *GetDiscoveryDeviceResponse) GetCapDescriptionNaReason() string`

GetCapDescriptionNaReason returns the CapDescriptionNaReason field if non-nil, zero value otherwise.

### GetCapDescriptionNaReasonOk

`func (o *GetDiscoveryDeviceResponse) GetCapDescriptionNaReasonOk() (*string, bool)`

GetCapDescriptionNaReasonOk returns a tuple with the CapDescriptionNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapDescriptionNaReason

`func (o *GetDiscoveryDeviceResponse) SetCapDescriptionNaReason(v string)`

SetCapDescriptionNaReason sets CapDescriptionNaReason field to given value.

### HasCapDescriptionNaReason

`func (o *GetDiscoveryDeviceResponse) HasCapDescriptionNaReason() bool`

HasCapDescriptionNaReason returns a boolean if a field has been set.

### GetCapNetDeprovisioningInd

`func (o *GetDiscoveryDeviceResponse) GetCapNetDeprovisioningInd() bool`

GetCapNetDeprovisioningInd returns the CapNetDeprovisioningInd field if non-nil, zero value otherwise.

### GetCapNetDeprovisioningIndOk

`func (o *GetDiscoveryDeviceResponse) GetCapNetDeprovisioningIndOk() (*bool, bool)`

GetCapNetDeprovisioningIndOk returns a tuple with the CapNetDeprovisioningInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetDeprovisioningInd

`func (o *GetDiscoveryDeviceResponse) SetCapNetDeprovisioningInd(v bool)`

SetCapNetDeprovisioningInd sets CapNetDeprovisioningInd field to given value.

### HasCapNetDeprovisioningInd

`func (o *GetDiscoveryDeviceResponse) HasCapNetDeprovisioningInd() bool`

HasCapNetDeprovisioningInd returns a boolean if a field has been set.

### GetCapNetDeprovisioningNaReason

`func (o *GetDiscoveryDeviceResponse) GetCapNetDeprovisioningNaReason() string`

GetCapNetDeprovisioningNaReason returns the CapNetDeprovisioningNaReason field if non-nil, zero value otherwise.

### GetCapNetDeprovisioningNaReasonOk

`func (o *GetDiscoveryDeviceResponse) GetCapNetDeprovisioningNaReasonOk() (*string, bool)`

GetCapNetDeprovisioningNaReasonOk returns a tuple with the CapNetDeprovisioningNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetDeprovisioningNaReason

`func (o *GetDiscoveryDeviceResponse) SetCapNetDeprovisioningNaReason(v string)`

SetCapNetDeprovisioningNaReason sets CapNetDeprovisioningNaReason field to given value.

### HasCapNetDeprovisioningNaReason

`func (o *GetDiscoveryDeviceResponse) HasCapNetDeprovisioningNaReason() bool`

HasCapNetDeprovisioningNaReason returns a boolean if a field has been set.

### GetCapNetProvisioningInd

`func (o *GetDiscoveryDeviceResponse) GetCapNetProvisioningInd() bool`

GetCapNetProvisioningInd returns the CapNetProvisioningInd field if non-nil, zero value otherwise.

### GetCapNetProvisioningIndOk

`func (o *GetDiscoveryDeviceResponse) GetCapNetProvisioningIndOk() (*bool, bool)`

GetCapNetProvisioningIndOk returns a tuple with the CapNetProvisioningInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetProvisioningInd

`func (o *GetDiscoveryDeviceResponse) SetCapNetProvisioningInd(v bool)`

SetCapNetProvisioningInd sets CapNetProvisioningInd field to given value.

### HasCapNetProvisioningInd

`func (o *GetDiscoveryDeviceResponse) HasCapNetProvisioningInd() bool`

HasCapNetProvisioningInd returns a boolean if a field has been set.

### GetCapNetProvisioningNaReason

`func (o *GetDiscoveryDeviceResponse) GetCapNetProvisioningNaReason() string`

GetCapNetProvisioningNaReason returns the CapNetProvisioningNaReason field if non-nil, zero value otherwise.

### GetCapNetProvisioningNaReasonOk

`func (o *GetDiscoveryDeviceResponse) GetCapNetProvisioningNaReasonOk() (*string, bool)`

GetCapNetProvisioningNaReasonOk returns a tuple with the CapNetProvisioningNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetProvisioningNaReason

`func (o *GetDiscoveryDeviceResponse) SetCapNetProvisioningNaReason(v string)`

SetCapNetProvisioningNaReason sets CapNetProvisioningNaReason field to given value.

### HasCapNetProvisioningNaReason

`func (o *GetDiscoveryDeviceResponse) HasCapNetProvisioningNaReason() bool`

HasCapNetProvisioningNaReason returns a boolean if a field has been set.

### GetCapNetVlanProvisioningInd

`func (o *GetDiscoveryDeviceResponse) GetCapNetVlanProvisioningInd() bool`

GetCapNetVlanProvisioningInd returns the CapNetVlanProvisioningInd field if non-nil, zero value otherwise.

### GetCapNetVlanProvisioningIndOk

`func (o *GetDiscoveryDeviceResponse) GetCapNetVlanProvisioningIndOk() (*bool, bool)`

GetCapNetVlanProvisioningIndOk returns a tuple with the CapNetVlanProvisioningInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetVlanProvisioningInd

`func (o *GetDiscoveryDeviceResponse) SetCapNetVlanProvisioningInd(v bool)`

SetCapNetVlanProvisioningInd sets CapNetVlanProvisioningInd field to given value.

### HasCapNetVlanProvisioningInd

`func (o *GetDiscoveryDeviceResponse) HasCapNetVlanProvisioningInd() bool`

HasCapNetVlanProvisioningInd returns a boolean if a field has been set.

### GetCapNetVlanProvisioningNaReason

`func (o *GetDiscoveryDeviceResponse) GetCapNetVlanProvisioningNaReason() string`

GetCapNetVlanProvisioningNaReason returns the CapNetVlanProvisioningNaReason field if non-nil, zero value otherwise.

### GetCapNetVlanProvisioningNaReasonOk

`func (o *GetDiscoveryDeviceResponse) GetCapNetVlanProvisioningNaReasonOk() (*string, bool)`

GetCapNetVlanProvisioningNaReasonOk returns a tuple with the CapNetVlanProvisioningNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapNetVlanProvisioningNaReason

`func (o *GetDiscoveryDeviceResponse) SetCapNetVlanProvisioningNaReason(v string)`

SetCapNetVlanProvisioningNaReason sets CapNetVlanProvisioningNaReason field to given value.

### HasCapNetVlanProvisioningNaReason

`func (o *GetDiscoveryDeviceResponse) HasCapNetVlanProvisioningNaReason() bool`

HasCapNetVlanProvisioningNaReason returns a boolean if a field has been set.

### GetCapVlanAssignmentInd

`func (o *GetDiscoveryDeviceResponse) GetCapVlanAssignmentInd() bool`

GetCapVlanAssignmentInd returns the CapVlanAssignmentInd field if non-nil, zero value otherwise.

### GetCapVlanAssignmentIndOk

`func (o *GetDiscoveryDeviceResponse) GetCapVlanAssignmentIndOk() (*bool, bool)`

GetCapVlanAssignmentIndOk returns a tuple with the CapVlanAssignmentInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapVlanAssignmentInd

`func (o *GetDiscoveryDeviceResponse) SetCapVlanAssignmentInd(v bool)`

SetCapVlanAssignmentInd sets CapVlanAssignmentInd field to given value.

### HasCapVlanAssignmentInd

`func (o *GetDiscoveryDeviceResponse) HasCapVlanAssignmentInd() bool`

HasCapVlanAssignmentInd returns a boolean if a field has been set.

### GetCapVlanAssignmentNaReason

`func (o *GetDiscoveryDeviceResponse) GetCapVlanAssignmentNaReason() string`

GetCapVlanAssignmentNaReason returns the CapVlanAssignmentNaReason field if non-nil, zero value otherwise.

### GetCapVlanAssignmentNaReasonOk

`func (o *GetDiscoveryDeviceResponse) GetCapVlanAssignmentNaReasonOk() (*string, bool)`

GetCapVlanAssignmentNaReasonOk returns a tuple with the CapVlanAssignmentNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapVlanAssignmentNaReason

`func (o *GetDiscoveryDeviceResponse) SetCapVlanAssignmentNaReason(v string)`

SetCapVlanAssignmentNaReason sets CapVlanAssignmentNaReason field to given value.

### HasCapVlanAssignmentNaReason

`func (o *GetDiscoveryDeviceResponse) HasCapVlanAssignmentNaReason() bool`

HasCapVlanAssignmentNaReason returns a boolean if a field has been set.

### GetCapVoiceVlanInd

`func (o *GetDiscoveryDeviceResponse) GetCapVoiceVlanInd() bool`

GetCapVoiceVlanInd returns the CapVoiceVlanInd field if non-nil, zero value otherwise.

### GetCapVoiceVlanIndOk

`func (o *GetDiscoveryDeviceResponse) GetCapVoiceVlanIndOk() (*bool, bool)`

GetCapVoiceVlanIndOk returns a tuple with the CapVoiceVlanInd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapVoiceVlanInd

`func (o *GetDiscoveryDeviceResponse) SetCapVoiceVlanInd(v bool)`

SetCapVoiceVlanInd sets CapVoiceVlanInd field to given value.

### HasCapVoiceVlanInd

`func (o *GetDiscoveryDeviceResponse) HasCapVoiceVlanInd() bool`

HasCapVoiceVlanInd returns a boolean if a field has been set.

### GetCapVoiceVlanNaReason

`func (o *GetDiscoveryDeviceResponse) GetCapVoiceVlanNaReason() string`

GetCapVoiceVlanNaReason returns the CapVoiceVlanNaReason field if non-nil, zero value otherwise.

### GetCapVoiceVlanNaReasonOk

`func (o *GetDiscoveryDeviceResponse) GetCapVoiceVlanNaReasonOk() (*string, bool)`

GetCapVoiceVlanNaReasonOk returns a tuple with the CapVoiceVlanNaReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapVoiceVlanNaReason

`func (o *GetDiscoveryDeviceResponse) SetCapVoiceVlanNaReason(v string)`

SetCapVoiceVlanNaReason sets CapVoiceVlanNaReason field to given value.

### HasCapVoiceVlanNaReason

`func (o *GetDiscoveryDeviceResponse) HasCapVoiceVlanNaReason() bool`

HasCapVoiceVlanNaReason returns a boolean if a field has been set.

### GetChassisSerialNumber

`func (o *GetDiscoveryDeviceResponse) GetChassisSerialNumber() string`

GetChassisSerialNumber returns the ChassisSerialNumber field if non-nil, zero value otherwise.

### GetChassisSerialNumberOk

`func (o *GetDiscoveryDeviceResponse) GetChassisSerialNumberOk() (*string, bool)`

GetChassisSerialNumberOk returns a tuple with the ChassisSerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChassisSerialNumber

`func (o *GetDiscoveryDeviceResponse) SetChassisSerialNumber(v string)`

SetChassisSerialNumber sets ChassisSerialNumber field to given value.

### HasChassisSerialNumber

`func (o *GetDiscoveryDeviceResponse) HasChassisSerialNumber() bool`

HasChassisSerialNumber returns a boolean if a field has been set.

### GetDescription

`func (o *GetDiscoveryDeviceResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetDiscoveryDeviceResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetDiscoveryDeviceResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetDiscoveryDeviceResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetDiscoveryDeviceResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetDiscoveryDeviceResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetDiscoveryDeviceResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetDiscoveryDeviceResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetDiscoveryDeviceResponse) GetInterfaces() []string`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetDiscoveryDeviceResponse) GetInterfacesOk() (*[]string, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetDiscoveryDeviceResponse) SetInterfaces(v []string)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetDiscoveryDeviceResponse) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLocation

`func (o *GetDiscoveryDeviceResponse) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetDiscoveryDeviceResponse) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetDiscoveryDeviceResponse) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetDiscoveryDeviceResponse) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetModel

`func (o *GetDiscoveryDeviceResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetDiscoveryDeviceResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetDiscoveryDeviceResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetDiscoveryDeviceResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetDiscoveryDeviceResponse) GetMsAdUserData() DiscoveryDeviceMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetDiscoveryDeviceResponse) GetMsAdUserDataOk() (*DiscoveryDeviceMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetDiscoveryDeviceResponse) SetMsAdUserData(v DiscoveryDeviceMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetDiscoveryDeviceResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetName

`func (o *GetDiscoveryDeviceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDiscoveryDeviceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDiscoveryDeviceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDiscoveryDeviceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeighbors

`func (o *GetDiscoveryDeviceResponse) GetNeighbors() []string`

GetNeighbors returns the Neighbors field if non-nil, zero value otherwise.

### GetNeighborsOk

`func (o *GetDiscoveryDeviceResponse) GetNeighborsOk() (*[]string, bool)`

GetNeighborsOk returns a tuple with the Neighbors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighbors

`func (o *GetDiscoveryDeviceResponse) SetNeighbors(v []string)`

SetNeighbors sets Neighbors field to given value.

### HasNeighbors

`func (o *GetDiscoveryDeviceResponse) HasNeighbors() bool`

HasNeighbors returns a boolean if a field has been set.

### GetNetwork

`func (o *GetDiscoveryDeviceResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetDiscoveryDeviceResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetDiscoveryDeviceResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetDiscoveryDeviceResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkInfos

`func (o *GetDiscoveryDeviceResponse) GetNetworkInfos() []DiscoveryDeviceNetworkInfos`

GetNetworkInfos returns the NetworkInfos field if non-nil, zero value otherwise.

### GetNetworkInfosOk

`func (o *GetDiscoveryDeviceResponse) GetNetworkInfosOk() (*[]DiscoveryDeviceNetworkInfos, bool)`

GetNetworkInfosOk returns a tuple with the NetworkInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfos

`func (o *GetDiscoveryDeviceResponse) SetNetworkInfos(v []DiscoveryDeviceNetworkInfos)`

SetNetworkInfos sets NetworkInfos field to given value.

### HasNetworkInfos

`func (o *GetDiscoveryDeviceResponse) HasNetworkInfos() bool`

HasNetworkInfos returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetDiscoveryDeviceResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetDiscoveryDeviceResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetDiscoveryDeviceResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetDiscoveryDeviceResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetNetworks

`func (o *GetDiscoveryDeviceResponse) GetNetworks() []string`

GetNetworks returns the Networks field if non-nil, zero value otherwise.

### GetNetworksOk

`func (o *GetDiscoveryDeviceResponse) GetNetworksOk() (*[]string, bool)`

GetNetworksOk returns a tuple with the Networks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworks

`func (o *GetDiscoveryDeviceResponse) SetNetworks(v []string)`

SetNetworks sets Networks field to given value.

### HasNetworks

`func (o *GetDiscoveryDeviceResponse) HasNetworks() bool`

HasNetworks returns a boolean if a field has been set.

### GetOsVersion

`func (o *GetDiscoveryDeviceResponse) GetOsVersion() string`

GetOsVersion returns the OsVersion field if non-nil, zero value otherwise.

### GetOsVersionOk

`func (o *GetDiscoveryDeviceResponse) GetOsVersionOk() (*string, bool)`

GetOsVersionOk returns a tuple with the OsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsVersion

`func (o *GetDiscoveryDeviceResponse) SetOsVersion(v string)`

SetOsVersion sets OsVersion field to given value.

### HasOsVersion

`func (o *GetDiscoveryDeviceResponse) HasOsVersion() bool`

HasOsVersion returns a boolean if a field has been set.

### GetPortStats

`func (o *GetDiscoveryDeviceResponse) GetPortStats() DiscoveryDevicePortStats`

GetPortStats returns the PortStats field if non-nil, zero value otherwise.

### GetPortStatsOk

`func (o *GetDiscoveryDeviceResponse) GetPortStatsOk() (*DiscoveryDevicePortStats, bool)`

GetPortStatsOk returns a tuple with the PortStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStats

`func (o *GetDiscoveryDeviceResponse) SetPortStats(v DiscoveryDevicePortStats)`

SetPortStats sets PortStats field to given value.

### HasPortStats

`func (o *GetDiscoveryDeviceResponse) HasPortStats() bool`

HasPortStats returns a boolean if a field has been set.

### GetPrivilegedPolling

`func (o *GetDiscoveryDeviceResponse) GetPrivilegedPolling() bool`

GetPrivilegedPolling returns the PrivilegedPolling field if non-nil, zero value otherwise.

### GetPrivilegedPollingOk

`func (o *GetDiscoveryDeviceResponse) GetPrivilegedPollingOk() (*bool, bool)`

GetPrivilegedPollingOk returns a tuple with the PrivilegedPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivilegedPolling

`func (o *GetDiscoveryDeviceResponse) SetPrivilegedPolling(v bool)`

SetPrivilegedPolling sets PrivilegedPolling field to given value.

### HasPrivilegedPolling

`func (o *GetDiscoveryDeviceResponse) HasPrivilegedPolling() bool`

HasPrivilegedPolling returns a boolean if a field has been set.

### GetType

`func (o *GetDiscoveryDeviceResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDiscoveryDeviceResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDiscoveryDeviceResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDiscoveryDeviceResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserDefinedMgmtIp

`func (o *GetDiscoveryDeviceResponse) GetUserDefinedMgmtIp() string`

GetUserDefinedMgmtIp returns the UserDefinedMgmtIp field if non-nil, zero value otherwise.

### GetUserDefinedMgmtIpOk

`func (o *GetDiscoveryDeviceResponse) GetUserDefinedMgmtIpOk() (*string, bool)`

GetUserDefinedMgmtIpOk returns a tuple with the UserDefinedMgmtIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedMgmtIp

`func (o *GetDiscoveryDeviceResponse) SetUserDefinedMgmtIp(v string)`

SetUserDefinedMgmtIp sets UserDefinedMgmtIp field to given value.

### HasUserDefinedMgmtIp

`func (o *GetDiscoveryDeviceResponse) HasUserDefinedMgmtIp() bool`

HasUserDefinedMgmtIp returns a boolean if a field has been set.

### GetVendor

`func (o *GetDiscoveryDeviceResponse) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GetDiscoveryDeviceResponse) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GetDiscoveryDeviceResponse) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GetDiscoveryDeviceResponse) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVlanInfos

`func (o *GetDiscoveryDeviceResponse) GetVlanInfos() []DiscoveryDeviceVlanInfos`

GetVlanInfos returns the VlanInfos field if non-nil, zero value otherwise.

### GetVlanInfosOk

`func (o *GetDiscoveryDeviceResponse) GetVlanInfosOk() (*[]DiscoveryDeviceVlanInfos, bool)`

GetVlanInfosOk returns a tuple with the VlanInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInfos

`func (o *GetDiscoveryDeviceResponse) SetVlanInfos(v []DiscoveryDeviceVlanInfos)`

SetVlanInfos sets VlanInfos field to given value.

### HasVlanInfos

`func (o *GetDiscoveryDeviceResponse) HasVlanInfos() bool`

HasVlanInfos returns a boolean if a field has been set.

### GetResult

`func (o *GetDiscoveryDeviceResponse) GetResult() DiscoveryDevice`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDiscoveryDeviceResponse) GetResultOk() (*DiscoveryDevice, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDiscoveryDeviceResponse) SetResult(v DiscoveryDevice)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDiscoveryDeviceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


