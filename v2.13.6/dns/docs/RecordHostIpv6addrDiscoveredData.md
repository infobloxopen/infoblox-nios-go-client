# RecordHostIpv6addrDiscoveredData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceModel** | Pointer to **string** | The model name of the end device in the vendor terminology. | [optional] [readonly] 
**DevicePortName** | Pointer to **string** | The system name of the interface associated with the discovered IP address. | [optional] [readonly] 
**DevicePortType** | Pointer to **string** | The hardware type of the interface associated with the discovered IP address. | [optional] [readonly] 
**DeviceType** | Pointer to **string** | The type of end host in vendor terminology. | [optional] [readonly] 
**DeviceVendor** | Pointer to **string** | The vendor name of the end host. | [optional] [readonly] 
**DiscoveredName** | Pointer to **string** | The name of the network device associated with the discovered IP address. | [optional] [readonly] 
**Discoverer** | Pointer to **string** | Specifies whether the IP address was discovered by a NetMRI or NIOS discovery process. | [optional] [readonly] 
**Duid** | Pointer to **string** | For IPv6 address only. The DHCP unique identifier of the discovered host. This is an optional field, and data might not be included. | [optional] [readonly] 
**FirstDiscovered** | Pointer to **int64** | The date and time the IP address was first discovered in Epoch seconds format. | [optional] [readonly] 
**IprgNo** | Pointer to **int64** | The port redundant group number. | [optional] [readonly] 
**IprgState** | Pointer to **string** | The status for the IP address within port redundant group. | [optional] [readonly] 
**IprgType** | Pointer to **string** | The port redundant group type. | [optional] [readonly] 
**LastDiscovered** | Pointer to **int64** | The date and time the IP address was last discovered in Epoch seconds format. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The discovered MAC address for the host. This is the unique identifier of a network device. The discovery acquires the MAC address for hosts that are located on the same network as the Grid member that is running the discovery. This can also be the MAC address of a virtual entity on a specified vSphere server. | [optional] [readonly] 
**MgmtIpAddress** | Pointer to **string** | The management IP address of the end host that has more than one IP. | [optional] [readonly] 
**NetbiosName** | Pointer to **string** | The name returned in the NetBIOS reply or the name you manually register for the discovered host. | [optional] [readonly] 
**NetworkComponentDescription** | Pointer to **string** | A textual description of the switch that is connected to the end device. | [optional] [readonly] 
**NetworkComponentIp** | Pointer to **string** | The IPv4 Address or IPv6 Address of the switch that is connected to the end device. | [optional] [readonly] 
**NetworkComponentModel** | Pointer to **string** | Model name of the switch port connected to the end host in vendor terminology. | [optional] [readonly] 
**NetworkComponentName** | Pointer to **string** | If a reverse lookup was successful for the IP address associated with this switch, the host name is displayed in this field. | [optional] [readonly] 
**NetworkComponentPortDescription** | Pointer to **string** | A textual description of the switch port that is connected to the end device. | [optional] [readonly] 
**NetworkComponentPortName** | Pointer to **string** | The name of the switch port connected to the end device. | [optional] [readonly] 
**NetworkComponentPortNumber** | Pointer to **string** | The number of the switch port connected to the end device. | [optional] [readonly] 
**NetworkComponentType** | Pointer to **string** | Identifies the switch that is connected to the end device. | [optional] [readonly] 
**NetworkComponentVendor** | Pointer to **string** | The vendor name of the switch port connected to the end host. | [optional] [readonly] 
**OpenPorts** | Pointer to **string** | The list of opened ports on the IP address, represented as: \&quot;TCP: 21,22,23 UDP: 137,139\&quot;. Limited to max total 1000 ports. | [optional] [readonly] 
**Os** | Pointer to **string** | The operating system of the detected host or virtual entity. The OS can be one of the following: * Microsoft for all discovered hosts that have a non-null value in the MAC addresses using the NetBIOS discovery method. * A value that a TCP discovery returns. * The OS of a virtual entity on a vSphere server. | [optional] [readonly] 
**PortDuplex** | Pointer to **string** | The negotiated or operational duplex setting of the switch port connected to the end device. | [optional] [readonly] 
**PortLinkStatus** | Pointer to **string** | The link status of the switch port connected to the end device. Indicates whether it is connected. | [optional] [readonly] 
**PortSpeed** | Pointer to **string** | The interface speed, in Mbps, of the switch port. | [optional] [readonly] 
**PortStatus** | Pointer to **string** | The operational status of the switch port. Indicates whether the port is up or down. | [optional] [readonly] 
**PortType** | Pointer to **string** | The type of switch port. | [optional] [readonly] 
**PortVlanDescription** | Pointer to **string** | The description of the VLAN of the switch port that is connected to the end device. | [optional] [readonly] 
**PortVlanName** | Pointer to **string** | The name of the VLAN of the switch port. | [optional] [readonly] 
**PortVlanNumber** | Pointer to **string** | The ID of the VLAN of the switch port. | [optional] [readonly] 
**VAdapter** | Pointer to **string** | The name of the physical network adapter through which the virtual entity is connected to the appliance. | [optional] [readonly] 
**VCluster** | Pointer to **string** | The name of the VMware cluster to which the virtual entity belongs. | [optional] [readonly] 
**VDatacenter** | Pointer to **string** | The name of the vSphere datacenter or container to which the virtual entity belongs. | [optional] [readonly] 
**VEntityName** | Pointer to **string** | The name of the virtual entity. | [optional] [readonly] 
**VEntityType** | Pointer to **string** | The virtual entity type. This can be blank or one of the following: Virtual Machine, Virtual Host, or Virtual Center. Virtual Center represents a VMware vCenter server. | [optional] [readonly] 
**VHost** | Pointer to **string** | The name of the VMware server on which the virtual entity was discovered. | [optional] [readonly] 
**VSwitch** | Pointer to **string** | The name of the switch to which the virtual entity is connected. | [optional] [readonly] 
**VmiName** | Pointer to **string** | Name of the virtual machine. | [optional] [readonly] 
**VmiId** | Pointer to **string** | ID of the virtual machine. | [optional] [readonly] 
**VlanPortGroup** | Pointer to **string** | Port group which the virtual machine belongs to. | [optional] [readonly] 
**VswitchName** | Pointer to **string** | Name of the virtual switch. | [optional] [readonly] 
**VswitchId** | Pointer to **string** | ID of the virtual switch. | [optional] [readonly] 
**VswitchType** | Pointer to **string** | Type of the virtual switch: standard or distributed. | [optional] [readonly] 
**VswitchIpv6Enabled** | Pointer to **bool** | Indicates the virtual switch has IPV6 enabled. | [optional] [readonly] 
**VportName** | Pointer to **string** | Name of the network adapter on the virtual switch connected with the virtual machine. | [optional] [readonly] 
**VportMacAddress** | Pointer to **string** | MAC address of the network adapter on the virtual switch where the virtual machine connected to. | [optional] [readonly] 
**VportLinkStatus** | Pointer to **string** | Link status of the network adapter on the virtual switch where the virtual machine connected to. | [optional] [readonly] 
**VportConfSpeed** | Pointer to **string** | Configured speed of the network adapter on the virtual switch where the virtual machine connected to. Unit is kb. | [optional] [readonly] 
**VportConfMode** | Pointer to **string** | Configured mode of the network adapter on the virtual switch where the virtual machine connected to. | [optional] [readonly] 
**VportSpeed** | Pointer to **string** | Actual speed of the network adapter on the virtual switch where the virtual machine connected to. Unit is kb. | [optional] [readonly] 
**VportMode** | Pointer to **string** | Actual mode of the network adapter on the virtual switch where the virtual machine connected to. | [optional] [readonly] 
**VswitchSegmentType** | Pointer to **string** | Type of the network segment on which the current virtual machine/vport connected to. | [optional] [readonly] 
**VswitchSegmentName** | Pointer to **string** | Name of the network segment on which the current virtual machine/vport connected to. | [optional] [readonly] 
**VswitchSegmentId** | Pointer to **string** | ID of the network segment on which the current virtual machine/vport connected to. | [optional] [readonly] 
**VswitchSegmentPortGroup** | Pointer to **string** | Port group of the network segment on which the current virtual machine/vport connected to. | [optional] [readonly] 
**VswitchAvailablePortsCount** | Pointer to **int64** | Numer of available ports reported by the virtual switch on which the virtual machine/vport connected to. | [optional] [readonly] 
**VswitchTepType** | Pointer to **string** | Type of virtual tunnel endpoint (VTEP) in the virtual switch. | [optional] [readonly] 
**VswitchTepIp** | Pointer to **string** | IP address of the virtual tunnel endpoint (VTEP) in the virtual switch. | [optional] [readonly] 
**VswitchTepPortGroup** | Pointer to **string** | Port group of the virtual tunnel endpoint (VTEP) in the virtual switch. | [optional] [readonly] 
**VswitchTepVlan** | Pointer to **string** | VLAN of the virtual tunnel endpoint (VTEP) in the virtual switch. | [optional] [readonly] 
**VswitchTepDhcpServer** | Pointer to **string** | DHCP server of the virtual tunnel endpoint (VTEP) in the virtual switch. | [optional] [readonly] 
**VswitchTepMulticast** | Pointer to **string** | Muticast address of the virtual tunnel endpoint (VTEP) in the virtual swtich. | [optional] [readonly] 
**VmhostIpAddress** | Pointer to **string** | IP address of the physical node on which the virtual machine is hosted. | [optional] [readonly] 
**VmhostName** | Pointer to **string** | Name of the physical node on which the virtual machine is hosted. | [optional] [readonly] 
**VmhostMacAddress** | Pointer to **string** | MAC address of the physical node on which the virtual machine is hosted. | [optional] [readonly] 
**VmhostSubnetCidr** | Pointer to **int64** | CIDR subnet of the physical node on which the virtual machine is hosted. | [optional] [readonly] 
**VmhostNicNames** | Pointer to **string** | List of all physical port names used by the virtual switch on the physical node on which the virtual machine is hosted. Represented as: \&quot;eth1,eth2,eth3\&quot;. | [optional] [readonly] 
**VmiTenantId** | Pointer to **string** | ID of the tenant which virtual machine belongs to. | [optional] [readonly] 
**CmpType** | Pointer to **string** | If the IP is coming from a Cloud environment, the Cloud Management Platform type. | [optional] [readonly] 
**VmiIpType** | Pointer to **string** | Discovered IP address type. | [optional] [readonly] 
**VmiPrivateAddress** | Pointer to **string** | Private IP address of the virtual machine. | [optional] [readonly] 
**VmiIsPublicAddress** | Pointer to **bool** | Indicates whether the IP address is a public address. | [optional] [readonly] 
**CiscoIseSsid** | Pointer to **string** | The Cisco ISE SSID. | [optional] [readonly] 
**CiscoIseEndpointProfile** | Pointer to **string** | The Endpoint Profile created in Cisco ISE. | [optional] [readonly] 
**CiscoIseSessionState** | Pointer to **string** | The Cisco ISE connection session state. | [optional] [readonly] 
**CiscoIseSecurityGroup** | Pointer to **string** | The Cisco ISE security group name. | [optional] [readonly] 
**TaskName** | Pointer to **string** | The name of the discovery task. | [optional] [readonly] 
**NetworkComponentLocation** | Pointer to **string** | Location of the network component on which the IP address was discovered. | [optional] [readonly] 
**NetworkComponentContact** | Pointer to **string** | Contact information from the network component on which the IP address was discovered. | [optional] [readonly] 
**DeviceLocation** | Pointer to **string** | Location of device on which the IP address was discovered. | [optional] [readonly] 
**DeviceContact** | Pointer to **string** | Contact information from device on which the IP address was discovered. | [optional] [readonly] 
**ApName** | Pointer to **string** | Discovered name of Wireless Access Point. | [optional] [readonly] 
**ApIpAddress** | Pointer to **string** | Discovered IP address of Wireless Access Point. | [optional] [readonly] 
**ApSsid** | Pointer to **string** | Service set identifier (SSID) associated with Wireless Access Point. | [optional] [readonly] 
**BridgeDomain** | Pointer to **string** | Discovered bridge domain. | [optional] [readonly] 
**EndpointGroups** | Pointer to **string** | A comma-separated list of the discovered endpoint groups. | [optional] [readonly] 
**Tenant** | Pointer to **string** | Discovered tenant. | [optional] [readonly] 
**VrfName** | Pointer to **string** | The name of the VRF. | [optional] [readonly] 
**VrfDescription** | Pointer to **string** | Description of the VRF. | [optional] [readonly] 
**VrfRd** | Pointer to **string** | Route distinguisher of the VRF. | [optional] [readonly] 
**BgpAs** | Pointer to **int64** | The BGP autonomous system number. | [optional] [readonly] 

## Methods

### NewRecordHostIpv6addrDiscoveredData

`func NewRecordHostIpv6addrDiscoveredData() *RecordHostIpv6addrDiscoveredData`

NewRecordHostIpv6addrDiscoveredData instantiates a new RecordHostIpv6addrDiscoveredData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordHostIpv6addrDiscoveredDataWithDefaults

`func NewRecordHostIpv6addrDiscoveredDataWithDefaults() *RecordHostIpv6addrDiscoveredData`

NewRecordHostIpv6addrDiscoveredDataWithDefaults instantiates a new RecordHostIpv6addrDiscoveredData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *RecordHostIpv6addrDiscoveredData) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *RecordHostIpv6addrDiscoveredData) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDevicePortName

`func (o *RecordHostIpv6addrDiscoveredData) GetDevicePortName() string`

GetDevicePortName returns the DevicePortName field if non-nil, zero value otherwise.

### GetDevicePortNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDevicePortNameOk() (*string, bool)`

GetDevicePortNameOk returns a tuple with the DevicePortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortName

`func (o *RecordHostIpv6addrDiscoveredData) SetDevicePortName(v string)`

SetDevicePortName sets DevicePortName field to given value.

### HasDevicePortName

`func (o *RecordHostIpv6addrDiscoveredData) HasDevicePortName() bool`

HasDevicePortName returns a boolean if a field has been set.

### GetDevicePortType

`func (o *RecordHostIpv6addrDiscoveredData) GetDevicePortType() string`

GetDevicePortType returns the DevicePortType field if non-nil, zero value otherwise.

### GetDevicePortTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDevicePortTypeOk() (*string, bool)`

GetDevicePortTypeOk returns a tuple with the DevicePortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortType

`func (o *RecordHostIpv6addrDiscoveredData) SetDevicePortType(v string)`

SetDevicePortType sets DevicePortType field to given value.

### HasDevicePortType

`func (o *RecordHostIpv6addrDiscoveredData) HasDevicePortType() bool`

HasDevicePortType returns a boolean if a field has been set.

### GetDeviceType

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *RecordHostIpv6addrDiscoveredData) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *RecordHostIpv6addrDiscoveredData) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *RecordHostIpv6addrDiscoveredData) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *RecordHostIpv6addrDiscoveredData) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDiscoveredName

`func (o *RecordHostIpv6addrDiscoveredData) GetDiscoveredName() string`

GetDiscoveredName returns the DiscoveredName field if non-nil, zero value otherwise.

### GetDiscoveredNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDiscoveredNameOk() (*string, bool)`

GetDiscoveredNameOk returns a tuple with the DiscoveredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredName

`func (o *RecordHostIpv6addrDiscoveredData) SetDiscoveredName(v string)`

SetDiscoveredName sets DiscoveredName field to given value.

### HasDiscoveredName

`func (o *RecordHostIpv6addrDiscoveredData) HasDiscoveredName() bool`

HasDiscoveredName returns a boolean if a field has been set.

### GetDiscoverer

`func (o *RecordHostIpv6addrDiscoveredData) GetDiscoverer() string`

GetDiscoverer returns the Discoverer field if non-nil, zero value otherwise.

### GetDiscovererOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDiscovererOk() (*string, bool)`

GetDiscovererOk returns a tuple with the Discoverer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverer

`func (o *RecordHostIpv6addrDiscoveredData) SetDiscoverer(v string)`

SetDiscoverer sets Discoverer field to given value.

### HasDiscoverer

`func (o *RecordHostIpv6addrDiscoveredData) HasDiscoverer() bool`

HasDiscoverer returns a boolean if a field has been set.

### GetDuid

`func (o *RecordHostIpv6addrDiscoveredData) GetDuid() string`

GetDuid returns the Duid field if non-nil, zero value otherwise.

### GetDuidOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDuidOk() (*string, bool)`

GetDuidOk returns a tuple with the Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuid

`func (o *RecordHostIpv6addrDiscoveredData) SetDuid(v string)`

SetDuid sets Duid field to given value.

### HasDuid

`func (o *RecordHostIpv6addrDiscoveredData) HasDuid() bool`

HasDuid returns a boolean if a field has been set.

### GetFirstDiscovered

`func (o *RecordHostIpv6addrDiscoveredData) GetFirstDiscovered() int64`

GetFirstDiscovered returns the FirstDiscovered field if non-nil, zero value otherwise.

### GetFirstDiscoveredOk

`func (o *RecordHostIpv6addrDiscoveredData) GetFirstDiscoveredOk() (*int64, bool)`

GetFirstDiscoveredOk returns a tuple with the FirstDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDiscovered

`func (o *RecordHostIpv6addrDiscoveredData) SetFirstDiscovered(v int64)`

SetFirstDiscovered sets FirstDiscovered field to given value.

### HasFirstDiscovered

`func (o *RecordHostIpv6addrDiscoveredData) HasFirstDiscovered() bool`

HasFirstDiscovered returns a boolean if a field has been set.

### GetIprgNo

`func (o *RecordHostIpv6addrDiscoveredData) GetIprgNo() int64`

GetIprgNo returns the IprgNo field if non-nil, zero value otherwise.

### GetIprgNoOk

`func (o *RecordHostIpv6addrDiscoveredData) GetIprgNoOk() (*int64, bool)`

GetIprgNoOk returns a tuple with the IprgNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprgNo

`func (o *RecordHostIpv6addrDiscoveredData) SetIprgNo(v int64)`

SetIprgNo sets IprgNo field to given value.

### HasIprgNo

`func (o *RecordHostIpv6addrDiscoveredData) HasIprgNo() bool`

HasIprgNo returns a boolean if a field has been set.

### GetIprgState

`func (o *RecordHostIpv6addrDiscoveredData) GetIprgState() string`

GetIprgState returns the IprgState field if non-nil, zero value otherwise.

### GetIprgStateOk

`func (o *RecordHostIpv6addrDiscoveredData) GetIprgStateOk() (*string, bool)`

GetIprgStateOk returns a tuple with the IprgState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprgState

`func (o *RecordHostIpv6addrDiscoveredData) SetIprgState(v string)`

SetIprgState sets IprgState field to given value.

### HasIprgState

`func (o *RecordHostIpv6addrDiscoveredData) HasIprgState() bool`

HasIprgState returns a boolean if a field has been set.

### GetIprgType

`func (o *RecordHostIpv6addrDiscoveredData) GetIprgType() string`

GetIprgType returns the IprgType field if non-nil, zero value otherwise.

### GetIprgTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetIprgTypeOk() (*string, bool)`

GetIprgTypeOk returns a tuple with the IprgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprgType

`func (o *RecordHostIpv6addrDiscoveredData) SetIprgType(v string)`

SetIprgType sets IprgType field to given value.

### HasIprgType

`func (o *RecordHostIpv6addrDiscoveredData) HasIprgType() bool`

HasIprgType returns a boolean if a field has been set.

### GetLastDiscovered

`func (o *RecordHostIpv6addrDiscoveredData) GetLastDiscovered() int64`

GetLastDiscovered returns the LastDiscovered field if non-nil, zero value otherwise.

### GetLastDiscoveredOk

`func (o *RecordHostIpv6addrDiscoveredData) GetLastDiscoveredOk() (*int64, bool)`

GetLastDiscoveredOk returns a tuple with the LastDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscovered

`func (o *RecordHostIpv6addrDiscoveredData) SetLastDiscovered(v int64)`

SetLastDiscovered sets LastDiscovered field to given value.

### HasLastDiscovered

`func (o *RecordHostIpv6addrDiscoveredData) HasLastDiscovered() bool`

HasLastDiscovered returns a boolean if a field has been set.

### GetMacAddress

`func (o *RecordHostIpv6addrDiscoveredData) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *RecordHostIpv6addrDiscoveredData) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *RecordHostIpv6addrDiscoveredData) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *RecordHostIpv6addrDiscoveredData) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *RecordHostIpv6addrDiscoveredData) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *RecordHostIpv6addrDiscoveredData) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *RecordHostIpv6addrDiscoveredData) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *RecordHostIpv6addrDiscoveredData) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetNetbiosName

`func (o *RecordHostIpv6addrDiscoveredData) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *RecordHostIpv6addrDiscoveredData) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *RecordHostIpv6addrDiscoveredData) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.

### GetNetworkComponentDescription

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentDescription() string`

GetNetworkComponentDescription returns the NetworkComponentDescription field if non-nil, zero value otherwise.

### GetNetworkComponentDescriptionOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentDescriptionOk() (*string, bool)`

GetNetworkComponentDescriptionOk returns a tuple with the NetworkComponentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentDescription

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentDescription(v string)`

SetNetworkComponentDescription sets NetworkComponentDescription field to given value.

### HasNetworkComponentDescription

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentDescription() bool`

HasNetworkComponentDescription returns a boolean if a field has been set.

### GetNetworkComponentIp

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentIp() string`

GetNetworkComponentIp returns the NetworkComponentIp field if non-nil, zero value otherwise.

### GetNetworkComponentIpOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentIpOk() (*string, bool)`

GetNetworkComponentIpOk returns a tuple with the NetworkComponentIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentIp

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentIp(v string)`

SetNetworkComponentIp sets NetworkComponentIp field to given value.

### HasNetworkComponentIp

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentIp() bool`

HasNetworkComponentIp returns a boolean if a field has been set.

### GetNetworkComponentModel

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentModel() string`

GetNetworkComponentModel returns the NetworkComponentModel field if non-nil, zero value otherwise.

### GetNetworkComponentModelOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentModelOk() (*string, bool)`

GetNetworkComponentModelOk returns a tuple with the NetworkComponentModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentModel

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentModel(v string)`

SetNetworkComponentModel sets NetworkComponentModel field to given value.

### HasNetworkComponentModel

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentModel() bool`

HasNetworkComponentModel returns a boolean if a field has been set.

### GetNetworkComponentName

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentName() string`

GetNetworkComponentName returns the NetworkComponentName field if non-nil, zero value otherwise.

### GetNetworkComponentNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentNameOk() (*string, bool)`

GetNetworkComponentNameOk returns a tuple with the NetworkComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentName

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentName(v string)`

SetNetworkComponentName sets NetworkComponentName field to given value.

### HasNetworkComponentName

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentName() bool`

HasNetworkComponentName returns a boolean if a field has been set.

### GetNetworkComponentPortDescription

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentPortDescription() string`

GetNetworkComponentPortDescription returns the NetworkComponentPortDescription field if non-nil, zero value otherwise.

### GetNetworkComponentPortDescriptionOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentPortDescriptionOk() (*string, bool)`

GetNetworkComponentPortDescriptionOk returns a tuple with the NetworkComponentPortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentPortDescription

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentPortDescription(v string)`

SetNetworkComponentPortDescription sets NetworkComponentPortDescription field to given value.

### HasNetworkComponentPortDescription

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentPortDescription() bool`

HasNetworkComponentPortDescription returns a boolean if a field has been set.

### GetNetworkComponentPortName

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentPortName() string`

GetNetworkComponentPortName returns the NetworkComponentPortName field if non-nil, zero value otherwise.

### GetNetworkComponentPortNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentPortNameOk() (*string, bool)`

GetNetworkComponentPortNameOk returns a tuple with the NetworkComponentPortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentPortName

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentPortName(v string)`

SetNetworkComponentPortName sets NetworkComponentPortName field to given value.

### HasNetworkComponentPortName

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentPortName() bool`

HasNetworkComponentPortName returns a boolean if a field has been set.

### GetNetworkComponentPortNumber

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentPortNumber() string`

GetNetworkComponentPortNumber returns the NetworkComponentPortNumber field if non-nil, zero value otherwise.

### GetNetworkComponentPortNumberOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentPortNumberOk() (*string, bool)`

GetNetworkComponentPortNumberOk returns a tuple with the NetworkComponentPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentPortNumber

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentPortNumber(v string)`

SetNetworkComponentPortNumber sets NetworkComponentPortNumber field to given value.

### HasNetworkComponentPortNumber

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentPortNumber() bool`

HasNetworkComponentPortNumber returns a boolean if a field has been set.

### GetNetworkComponentType

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentType() string`

GetNetworkComponentType returns the NetworkComponentType field if non-nil, zero value otherwise.

### GetNetworkComponentTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentTypeOk() (*string, bool)`

GetNetworkComponentTypeOk returns a tuple with the NetworkComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentType

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentType(v string)`

SetNetworkComponentType sets NetworkComponentType field to given value.

### HasNetworkComponentType

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentType() bool`

HasNetworkComponentType returns a boolean if a field has been set.

### GetNetworkComponentVendor

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentVendor() string`

GetNetworkComponentVendor returns the NetworkComponentVendor field if non-nil, zero value otherwise.

### GetNetworkComponentVendorOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentVendorOk() (*string, bool)`

GetNetworkComponentVendorOk returns a tuple with the NetworkComponentVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentVendor

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentVendor(v string)`

SetNetworkComponentVendor sets NetworkComponentVendor field to given value.

### HasNetworkComponentVendor

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentVendor() bool`

HasNetworkComponentVendor returns a boolean if a field has been set.

### GetOpenPorts

`func (o *RecordHostIpv6addrDiscoveredData) GetOpenPorts() string`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *RecordHostIpv6addrDiscoveredData) GetOpenPortsOk() (*string, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *RecordHostIpv6addrDiscoveredData) SetOpenPorts(v string)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *RecordHostIpv6addrDiscoveredData) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetOs

`func (o *RecordHostIpv6addrDiscoveredData) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *RecordHostIpv6addrDiscoveredData) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *RecordHostIpv6addrDiscoveredData) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *RecordHostIpv6addrDiscoveredData) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPortDuplex

`func (o *RecordHostIpv6addrDiscoveredData) GetPortDuplex() string`

GetPortDuplex returns the PortDuplex field if non-nil, zero value otherwise.

### GetPortDuplexOk

`func (o *RecordHostIpv6addrDiscoveredData) GetPortDuplexOk() (*string, bool)`

GetPortDuplexOk returns a tuple with the PortDuplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDuplex

`func (o *RecordHostIpv6addrDiscoveredData) SetPortDuplex(v string)`

SetPortDuplex sets PortDuplex field to given value.

### HasPortDuplex

`func (o *RecordHostIpv6addrDiscoveredData) HasPortDuplex() bool`

HasPortDuplex returns a boolean if a field has been set.

### GetPortLinkStatus

`func (o *RecordHostIpv6addrDiscoveredData) GetPortLinkStatus() string`

GetPortLinkStatus returns the PortLinkStatus field if non-nil, zero value otherwise.

### GetPortLinkStatusOk

`func (o *RecordHostIpv6addrDiscoveredData) GetPortLinkStatusOk() (*string, bool)`

GetPortLinkStatusOk returns a tuple with the PortLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLinkStatus

`func (o *RecordHostIpv6addrDiscoveredData) SetPortLinkStatus(v string)`

SetPortLinkStatus sets PortLinkStatus field to given value.

### HasPortLinkStatus

`func (o *RecordHostIpv6addrDiscoveredData) HasPortLinkStatus() bool`

HasPortLinkStatus returns a boolean if a field has been set.

### GetPortSpeed

`func (o *RecordHostIpv6addrDiscoveredData) GetPortSpeed() string`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *RecordHostIpv6addrDiscoveredData) GetPortSpeedOk() (*string, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *RecordHostIpv6addrDiscoveredData) SetPortSpeed(v string)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *RecordHostIpv6addrDiscoveredData) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### GetPortStatus

`func (o *RecordHostIpv6addrDiscoveredData) GetPortStatus() string`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *RecordHostIpv6addrDiscoveredData) GetPortStatusOk() (*string, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *RecordHostIpv6addrDiscoveredData) SetPortStatus(v string)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *RecordHostIpv6addrDiscoveredData) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetPortType

`func (o *RecordHostIpv6addrDiscoveredData) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *RecordHostIpv6addrDiscoveredData) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *RecordHostIpv6addrDiscoveredData) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetPortVlanDescription

`func (o *RecordHostIpv6addrDiscoveredData) GetPortVlanDescription() string`

GetPortVlanDescription returns the PortVlanDescription field if non-nil, zero value otherwise.

### GetPortVlanDescriptionOk

`func (o *RecordHostIpv6addrDiscoveredData) GetPortVlanDescriptionOk() (*string, bool)`

GetPortVlanDescriptionOk returns a tuple with the PortVlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanDescription

`func (o *RecordHostIpv6addrDiscoveredData) SetPortVlanDescription(v string)`

SetPortVlanDescription sets PortVlanDescription field to given value.

### HasPortVlanDescription

`func (o *RecordHostIpv6addrDiscoveredData) HasPortVlanDescription() bool`

HasPortVlanDescription returns a boolean if a field has been set.

### GetPortVlanName

`func (o *RecordHostIpv6addrDiscoveredData) GetPortVlanName() string`

GetPortVlanName returns the PortVlanName field if non-nil, zero value otherwise.

### GetPortVlanNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetPortVlanNameOk() (*string, bool)`

GetPortVlanNameOk returns a tuple with the PortVlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanName

`func (o *RecordHostIpv6addrDiscoveredData) SetPortVlanName(v string)`

SetPortVlanName sets PortVlanName field to given value.

### HasPortVlanName

`func (o *RecordHostIpv6addrDiscoveredData) HasPortVlanName() bool`

HasPortVlanName returns a boolean if a field has been set.

### GetPortVlanNumber

`func (o *RecordHostIpv6addrDiscoveredData) GetPortVlanNumber() string`

GetPortVlanNumber returns the PortVlanNumber field if non-nil, zero value otherwise.

### GetPortVlanNumberOk

`func (o *RecordHostIpv6addrDiscoveredData) GetPortVlanNumberOk() (*string, bool)`

GetPortVlanNumberOk returns a tuple with the PortVlanNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanNumber

`func (o *RecordHostIpv6addrDiscoveredData) SetPortVlanNumber(v string)`

SetPortVlanNumber sets PortVlanNumber field to given value.

### HasPortVlanNumber

`func (o *RecordHostIpv6addrDiscoveredData) HasPortVlanNumber() bool`

HasPortVlanNumber returns a boolean if a field has been set.

### GetVAdapter

`func (o *RecordHostIpv6addrDiscoveredData) GetVAdapter() string`

GetVAdapter returns the VAdapter field if non-nil, zero value otherwise.

### GetVAdapterOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVAdapterOk() (*string, bool)`

GetVAdapterOk returns a tuple with the VAdapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAdapter

`func (o *RecordHostIpv6addrDiscoveredData) SetVAdapter(v string)`

SetVAdapter sets VAdapter field to given value.

### HasVAdapter

`func (o *RecordHostIpv6addrDiscoveredData) HasVAdapter() bool`

HasVAdapter returns a boolean if a field has been set.

### GetVCluster

`func (o *RecordHostIpv6addrDiscoveredData) GetVCluster() string`

GetVCluster returns the VCluster field if non-nil, zero value otherwise.

### GetVClusterOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVClusterOk() (*string, bool)`

GetVClusterOk returns a tuple with the VCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCluster

`func (o *RecordHostIpv6addrDiscoveredData) SetVCluster(v string)`

SetVCluster sets VCluster field to given value.

### HasVCluster

`func (o *RecordHostIpv6addrDiscoveredData) HasVCluster() bool`

HasVCluster returns a boolean if a field has been set.

### GetVDatacenter

`func (o *RecordHostIpv6addrDiscoveredData) GetVDatacenter() string`

GetVDatacenter returns the VDatacenter field if non-nil, zero value otherwise.

### GetVDatacenterOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVDatacenterOk() (*string, bool)`

GetVDatacenterOk returns a tuple with the VDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDatacenter

`func (o *RecordHostIpv6addrDiscoveredData) SetVDatacenter(v string)`

SetVDatacenter sets VDatacenter field to given value.

### HasVDatacenter

`func (o *RecordHostIpv6addrDiscoveredData) HasVDatacenter() bool`

HasVDatacenter returns a boolean if a field has been set.

### GetVEntityName

`func (o *RecordHostIpv6addrDiscoveredData) GetVEntityName() string`

GetVEntityName returns the VEntityName field if non-nil, zero value otherwise.

### GetVEntityNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVEntityNameOk() (*string, bool)`

GetVEntityNameOk returns a tuple with the VEntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVEntityName

`func (o *RecordHostIpv6addrDiscoveredData) SetVEntityName(v string)`

SetVEntityName sets VEntityName field to given value.

### HasVEntityName

`func (o *RecordHostIpv6addrDiscoveredData) HasVEntityName() bool`

HasVEntityName returns a boolean if a field has been set.

### GetVEntityType

`func (o *RecordHostIpv6addrDiscoveredData) GetVEntityType() string`

GetVEntityType returns the VEntityType field if non-nil, zero value otherwise.

### GetVEntityTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVEntityTypeOk() (*string, bool)`

GetVEntityTypeOk returns a tuple with the VEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVEntityType

`func (o *RecordHostIpv6addrDiscoveredData) SetVEntityType(v string)`

SetVEntityType sets VEntityType field to given value.

### HasVEntityType

`func (o *RecordHostIpv6addrDiscoveredData) HasVEntityType() bool`

HasVEntityType returns a boolean if a field has been set.

### GetVHost

`func (o *RecordHostIpv6addrDiscoveredData) GetVHost() string`

GetVHost returns the VHost field if non-nil, zero value otherwise.

### GetVHostOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVHostOk() (*string, bool)`

GetVHostOk returns a tuple with the VHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVHost

`func (o *RecordHostIpv6addrDiscoveredData) SetVHost(v string)`

SetVHost sets VHost field to given value.

### HasVHost

`func (o *RecordHostIpv6addrDiscoveredData) HasVHost() bool`

HasVHost returns a boolean if a field has been set.

### GetVSwitch

`func (o *RecordHostIpv6addrDiscoveredData) GetVSwitch() string`

GetVSwitch returns the VSwitch field if non-nil, zero value otherwise.

### GetVSwitchOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVSwitchOk() (*string, bool)`

GetVSwitchOk returns a tuple with the VSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSwitch

`func (o *RecordHostIpv6addrDiscoveredData) SetVSwitch(v string)`

SetVSwitch sets VSwitch field to given value.

### HasVSwitch

`func (o *RecordHostIpv6addrDiscoveredData) HasVSwitch() bool`

HasVSwitch returns a boolean if a field has been set.

### GetVmiName

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiName() string`

GetVmiName returns the VmiName field if non-nil, zero value otherwise.

### GetVmiNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiNameOk() (*string, bool)`

GetVmiNameOk returns a tuple with the VmiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiName

`func (o *RecordHostIpv6addrDiscoveredData) SetVmiName(v string)`

SetVmiName sets VmiName field to given value.

### HasVmiName

`func (o *RecordHostIpv6addrDiscoveredData) HasVmiName() bool`

HasVmiName returns a boolean if a field has been set.

### GetVmiId

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiId() string`

GetVmiId returns the VmiId field if non-nil, zero value otherwise.

### GetVmiIdOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiIdOk() (*string, bool)`

GetVmiIdOk returns a tuple with the VmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiId

`func (o *RecordHostIpv6addrDiscoveredData) SetVmiId(v string)`

SetVmiId sets VmiId field to given value.

### HasVmiId

`func (o *RecordHostIpv6addrDiscoveredData) HasVmiId() bool`

HasVmiId returns a boolean if a field has been set.

### GetVlanPortGroup

`func (o *RecordHostIpv6addrDiscoveredData) GetVlanPortGroup() string`

GetVlanPortGroup returns the VlanPortGroup field if non-nil, zero value otherwise.

### GetVlanPortGroupOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVlanPortGroupOk() (*string, bool)`

GetVlanPortGroupOk returns a tuple with the VlanPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPortGroup

`func (o *RecordHostIpv6addrDiscoveredData) SetVlanPortGroup(v string)`

SetVlanPortGroup sets VlanPortGroup field to given value.

### HasVlanPortGroup

`func (o *RecordHostIpv6addrDiscoveredData) HasVlanPortGroup() bool`

HasVlanPortGroup returns a boolean if a field has been set.

### GetVswitchName

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchName() string`

GetVswitchName returns the VswitchName field if non-nil, zero value otherwise.

### GetVswitchNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchNameOk() (*string, bool)`

GetVswitchNameOk returns a tuple with the VswitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchName

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchName(v string)`

SetVswitchName sets VswitchName field to given value.

### HasVswitchName

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchName() bool`

HasVswitchName returns a boolean if a field has been set.

### GetVswitchId

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchId() string`

GetVswitchId returns the VswitchId field if non-nil, zero value otherwise.

### GetVswitchIdOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchIdOk() (*string, bool)`

GetVswitchIdOk returns a tuple with the VswitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchId

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchId(v string)`

SetVswitchId sets VswitchId field to given value.

### HasVswitchId

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchId() bool`

HasVswitchId returns a boolean if a field has been set.

### GetVswitchType

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchType() string`

GetVswitchType returns the VswitchType field if non-nil, zero value otherwise.

### GetVswitchTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTypeOk() (*string, bool)`

GetVswitchTypeOk returns a tuple with the VswitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchType

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchType(v string)`

SetVswitchType sets VswitchType field to given value.

### HasVswitchType

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchType() bool`

HasVswitchType returns a boolean if a field has been set.

### GetVswitchIpv6Enabled

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchIpv6Enabled() bool`

GetVswitchIpv6Enabled returns the VswitchIpv6Enabled field if non-nil, zero value otherwise.

### GetVswitchIpv6EnabledOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchIpv6EnabledOk() (*bool, bool)`

GetVswitchIpv6EnabledOk returns a tuple with the VswitchIpv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchIpv6Enabled

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchIpv6Enabled(v bool)`

SetVswitchIpv6Enabled sets VswitchIpv6Enabled field to given value.

### HasVswitchIpv6Enabled

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchIpv6Enabled() bool`

HasVswitchIpv6Enabled returns a boolean if a field has been set.

### GetVportName

`func (o *RecordHostIpv6addrDiscoveredData) GetVportName() string`

GetVportName returns the VportName field if non-nil, zero value otherwise.

### GetVportNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVportNameOk() (*string, bool)`

GetVportNameOk returns a tuple with the VportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportName

`func (o *RecordHostIpv6addrDiscoveredData) SetVportName(v string)`

SetVportName sets VportName field to given value.

### HasVportName

`func (o *RecordHostIpv6addrDiscoveredData) HasVportName() bool`

HasVportName returns a boolean if a field has been set.

### GetVportMacAddress

`func (o *RecordHostIpv6addrDiscoveredData) GetVportMacAddress() string`

GetVportMacAddress returns the VportMacAddress field if non-nil, zero value otherwise.

### GetVportMacAddressOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVportMacAddressOk() (*string, bool)`

GetVportMacAddressOk returns a tuple with the VportMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportMacAddress

`func (o *RecordHostIpv6addrDiscoveredData) SetVportMacAddress(v string)`

SetVportMacAddress sets VportMacAddress field to given value.

### HasVportMacAddress

`func (o *RecordHostIpv6addrDiscoveredData) HasVportMacAddress() bool`

HasVportMacAddress returns a boolean if a field has been set.

### GetVportLinkStatus

`func (o *RecordHostIpv6addrDiscoveredData) GetVportLinkStatus() string`

GetVportLinkStatus returns the VportLinkStatus field if non-nil, zero value otherwise.

### GetVportLinkStatusOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVportLinkStatusOk() (*string, bool)`

GetVportLinkStatusOk returns a tuple with the VportLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportLinkStatus

`func (o *RecordHostIpv6addrDiscoveredData) SetVportLinkStatus(v string)`

SetVportLinkStatus sets VportLinkStatus field to given value.

### HasVportLinkStatus

`func (o *RecordHostIpv6addrDiscoveredData) HasVportLinkStatus() bool`

HasVportLinkStatus returns a boolean if a field has been set.

### GetVportConfSpeed

`func (o *RecordHostIpv6addrDiscoveredData) GetVportConfSpeed() string`

GetVportConfSpeed returns the VportConfSpeed field if non-nil, zero value otherwise.

### GetVportConfSpeedOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVportConfSpeedOk() (*string, bool)`

GetVportConfSpeedOk returns a tuple with the VportConfSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportConfSpeed

`func (o *RecordHostIpv6addrDiscoveredData) SetVportConfSpeed(v string)`

SetVportConfSpeed sets VportConfSpeed field to given value.

### HasVportConfSpeed

`func (o *RecordHostIpv6addrDiscoveredData) HasVportConfSpeed() bool`

HasVportConfSpeed returns a boolean if a field has been set.

### GetVportConfMode

`func (o *RecordHostIpv6addrDiscoveredData) GetVportConfMode() string`

GetVportConfMode returns the VportConfMode field if non-nil, zero value otherwise.

### GetVportConfModeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVportConfModeOk() (*string, bool)`

GetVportConfModeOk returns a tuple with the VportConfMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportConfMode

`func (o *RecordHostIpv6addrDiscoveredData) SetVportConfMode(v string)`

SetVportConfMode sets VportConfMode field to given value.

### HasVportConfMode

`func (o *RecordHostIpv6addrDiscoveredData) HasVportConfMode() bool`

HasVportConfMode returns a boolean if a field has been set.

### GetVportSpeed

`func (o *RecordHostIpv6addrDiscoveredData) GetVportSpeed() string`

GetVportSpeed returns the VportSpeed field if non-nil, zero value otherwise.

### GetVportSpeedOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVportSpeedOk() (*string, bool)`

GetVportSpeedOk returns a tuple with the VportSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportSpeed

`func (o *RecordHostIpv6addrDiscoveredData) SetVportSpeed(v string)`

SetVportSpeed sets VportSpeed field to given value.

### HasVportSpeed

`func (o *RecordHostIpv6addrDiscoveredData) HasVportSpeed() bool`

HasVportSpeed returns a boolean if a field has been set.

### GetVportMode

`func (o *RecordHostIpv6addrDiscoveredData) GetVportMode() string`

GetVportMode returns the VportMode field if non-nil, zero value otherwise.

### GetVportModeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVportModeOk() (*string, bool)`

GetVportModeOk returns a tuple with the VportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportMode

`func (o *RecordHostIpv6addrDiscoveredData) SetVportMode(v string)`

SetVportMode sets VportMode field to given value.

### HasVportMode

`func (o *RecordHostIpv6addrDiscoveredData) HasVportMode() bool`

HasVportMode returns a boolean if a field has been set.

### GetVswitchSegmentType

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchSegmentType() string`

GetVswitchSegmentType returns the VswitchSegmentType field if non-nil, zero value otherwise.

### GetVswitchSegmentTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchSegmentTypeOk() (*string, bool)`

GetVswitchSegmentTypeOk returns a tuple with the VswitchSegmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentType

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchSegmentType(v string)`

SetVswitchSegmentType sets VswitchSegmentType field to given value.

### HasVswitchSegmentType

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchSegmentType() bool`

HasVswitchSegmentType returns a boolean if a field has been set.

### GetVswitchSegmentName

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchSegmentName() string`

GetVswitchSegmentName returns the VswitchSegmentName field if non-nil, zero value otherwise.

### GetVswitchSegmentNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchSegmentNameOk() (*string, bool)`

GetVswitchSegmentNameOk returns a tuple with the VswitchSegmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentName

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchSegmentName(v string)`

SetVswitchSegmentName sets VswitchSegmentName field to given value.

### HasVswitchSegmentName

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchSegmentName() bool`

HasVswitchSegmentName returns a boolean if a field has been set.

### GetVswitchSegmentId

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchSegmentId() string`

GetVswitchSegmentId returns the VswitchSegmentId field if non-nil, zero value otherwise.

### GetVswitchSegmentIdOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchSegmentIdOk() (*string, bool)`

GetVswitchSegmentIdOk returns a tuple with the VswitchSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentId

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchSegmentId(v string)`

SetVswitchSegmentId sets VswitchSegmentId field to given value.

### HasVswitchSegmentId

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchSegmentId() bool`

HasVswitchSegmentId returns a boolean if a field has been set.

### GetVswitchSegmentPortGroup

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchSegmentPortGroup() string`

GetVswitchSegmentPortGroup returns the VswitchSegmentPortGroup field if non-nil, zero value otherwise.

### GetVswitchSegmentPortGroupOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchSegmentPortGroupOk() (*string, bool)`

GetVswitchSegmentPortGroupOk returns a tuple with the VswitchSegmentPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentPortGroup

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchSegmentPortGroup(v string)`

SetVswitchSegmentPortGroup sets VswitchSegmentPortGroup field to given value.

### HasVswitchSegmentPortGroup

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchSegmentPortGroup() bool`

HasVswitchSegmentPortGroup returns a boolean if a field has been set.

### GetVswitchAvailablePortsCount

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchAvailablePortsCount() int64`

GetVswitchAvailablePortsCount returns the VswitchAvailablePortsCount field if non-nil, zero value otherwise.

### GetVswitchAvailablePortsCountOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchAvailablePortsCountOk() (*int64, bool)`

GetVswitchAvailablePortsCountOk returns a tuple with the VswitchAvailablePortsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchAvailablePortsCount

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchAvailablePortsCount(v int64)`

SetVswitchAvailablePortsCount sets VswitchAvailablePortsCount field to given value.

### HasVswitchAvailablePortsCount

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchAvailablePortsCount() bool`

HasVswitchAvailablePortsCount returns a boolean if a field has been set.

### GetVswitchTepType

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepType() string`

GetVswitchTepType returns the VswitchTepType field if non-nil, zero value otherwise.

### GetVswitchTepTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepTypeOk() (*string, bool)`

GetVswitchTepTypeOk returns a tuple with the VswitchTepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepType

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchTepType(v string)`

SetVswitchTepType sets VswitchTepType field to given value.

### HasVswitchTepType

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchTepType() bool`

HasVswitchTepType returns a boolean if a field has been set.

### GetVswitchTepIp

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepIp() string`

GetVswitchTepIp returns the VswitchTepIp field if non-nil, zero value otherwise.

### GetVswitchTepIpOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepIpOk() (*string, bool)`

GetVswitchTepIpOk returns a tuple with the VswitchTepIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepIp

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchTepIp(v string)`

SetVswitchTepIp sets VswitchTepIp field to given value.

### HasVswitchTepIp

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchTepIp() bool`

HasVswitchTepIp returns a boolean if a field has been set.

### GetVswitchTepPortGroup

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepPortGroup() string`

GetVswitchTepPortGroup returns the VswitchTepPortGroup field if non-nil, zero value otherwise.

### GetVswitchTepPortGroupOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepPortGroupOk() (*string, bool)`

GetVswitchTepPortGroupOk returns a tuple with the VswitchTepPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepPortGroup

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchTepPortGroup(v string)`

SetVswitchTepPortGroup sets VswitchTepPortGroup field to given value.

### HasVswitchTepPortGroup

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchTepPortGroup() bool`

HasVswitchTepPortGroup returns a boolean if a field has been set.

### GetVswitchTepVlan

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepVlan() string`

GetVswitchTepVlan returns the VswitchTepVlan field if non-nil, zero value otherwise.

### GetVswitchTepVlanOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepVlanOk() (*string, bool)`

GetVswitchTepVlanOk returns a tuple with the VswitchTepVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepVlan

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchTepVlan(v string)`

SetVswitchTepVlan sets VswitchTepVlan field to given value.

### HasVswitchTepVlan

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchTepVlan() bool`

HasVswitchTepVlan returns a boolean if a field has been set.

### GetVswitchTepDhcpServer

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepDhcpServer() string`

GetVswitchTepDhcpServer returns the VswitchTepDhcpServer field if non-nil, zero value otherwise.

### GetVswitchTepDhcpServerOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepDhcpServerOk() (*string, bool)`

GetVswitchTepDhcpServerOk returns a tuple with the VswitchTepDhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepDhcpServer

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchTepDhcpServer(v string)`

SetVswitchTepDhcpServer sets VswitchTepDhcpServer field to given value.

### HasVswitchTepDhcpServer

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchTepDhcpServer() bool`

HasVswitchTepDhcpServer returns a boolean if a field has been set.

### GetVswitchTepMulticast

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepMulticast() string`

GetVswitchTepMulticast returns the VswitchTepMulticast field if non-nil, zero value otherwise.

### GetVswitchTepMulticastOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVswitchTepMulticastOk() (*string, bool)`

GetVswitchTepMulticastOk returns a tuple with the VswitchTepMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepMulticast

`func (o *RecordHostIpv6addrDiscoveredData) SetVswitchTepMulticast(v string)`

SetVswitchTepMulticast sets VswitchTepMulticast field to given value.

### HasVswitchTepMulticast

`func (o *RecordHostIpv6addrDiscoveredData) HasVswitchTepMulticast() bool`

HasVswitchTepMulticast returns a boolean if a field has been set.

### GetVmhostIpAddress

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostIpAddress() string`

GetVmhostIpAddress returns the VmhostIpAddress field if non-nil, zero value otherwise.

### GetVmhostIpAddressOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostIpAddressOk() (*string, bool)`

GetVmhostIpAddressOk returns a tuple with the VmhostIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostIpAddress

`func (o *RecordHostIpv6addrDiscoveredData) SetVmhostIpAddress(v string)`

SetVmhostIpAddress sets VmhostIpAddress field to given value.

### HasVmhostIpAddress

`func (o *RecordHostIpv6addrDiscoveredData) HasVmhostIpAddress() bool`

HasVmhostIpAddress returns a boolean if a field has been set.

### GetVmhostName

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostName() string`

GetVmhostName returns the VmhostName field if non-nil, zero value otherwise.

### GetVmhostNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostNameOk() (*string, bool)`

GetVmhostNameOk returns a tuple with the VmhostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostName

`func (o *RecordHostIpv6addrDiscoveredData) SetVmhostName(v string)`

SetVmhostName sets VmhostName field to given value.

### HasVmhostName

`func (o *RecordHostIpv6addrDiscoveredData) HasVmhostName() bool`

HasVmhostName returns a boolean if a field has been set.

### GetVmhostMacAddress

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostMacAddress() string`

GetVmhostMacAddress returns the VmhostMacAddress field if non-nil, zero value otherwise.

### GetVmhostMacAddressOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostMacAddressOk() (*string, bool)`

GetVmhostMacAddressOk returns a tuple with the VmhostMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostMacAddress

`func (o *RecordHostIpv6addrDiscoveredData) SetVmhostMacAddress(v string)`

SetVmhostMacAddress sets VmhostMacAddress field to given value.

### HasVmhostMacAddress

`func (o *RecordHostIpv6addrDiscoveredData) HasVmhostMacAddress() bool`

HasVmhostMacAddress returns a boolean if a field has been set.

### GetVmhostSubnetCidr

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostSubnetCidr() int64`

GetVmhostSubnetCidr returns the VmhostSubnetCidr field if non-nil, zero value otherwise.

### GetVmhostSubnetCidrOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostSubnetCidrOk() (*int64, bool)`

GetVmhostSubnetCidrOk returns a tuple with the VmhostSubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostSubnetCidr

`func (o *RecordHostIpv6addrDiscoveredData) SetVmhostSubnetCidr(v int64)`

SetVmhostSubnetCidr sets VmhostSubnetCidr field to given value.

### HasVmhostSubnetCidr

`func (o *RecordHostIpv6addrDiscoveredData) HasVmhostSubnetCidr() bool`

HasVmhostSubnetCidr returns a boolean if a field has been set.

### GetVmhostNicNames

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostNicNames() string`

GetVmhostNicNames returns the VmhostNicNames field if non-nil, zero value otherwise.

### GetVmhostNicNamesOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmhostNicNamesOk() (*string, bool)`

GetVmhostNicNamesOk returns a tuple with the VmhostNicNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostNicNames

`func (o *RecordHostIpv6addrDiscoveredData) SetVmhostNicNames(v string)`

SetVmhostNicNames sets VmhostNicNames field to given value.

### HasVmhostNicNames

`func (o *RecordHostIpv6addrDiscoveredData) HasVmhostNicNames() bool`

HasVmhostNicNames returns a boolean if a field has been set.

### GetVmiTenantId

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiTenantId() string`

GetVmiTenantId returns the VmiTenantId field if non-nil, zero value otherwise.

### GetVmiTenantIdOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiTenantIdOk() (*string, bool)`

GetVmiTenantIdOk returns a tuple with the VmiTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiTenantId

`func (o *RecordHostIpv6addrDiscoveredData) SetVmiTenantId(v string)`

SetVmiTenantId sets VmiTenantId field to given value.

### HasVmiTenantId

`func (o *RecordHostIpv6addrDiscoveredData) HasVmiTenantId() bool`

HasVmiTenantId returns a boolean if a field has been set.

### GetCmpType

`func (o *RecordHostIpv6addrDiscoveredData) GetCmpType() string`

GetCmpType returns the CmpType field if non-nil, zero value otherwise.

### GetCmpTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetCmpTypeOk() (*string, bool)`

GetCmpTypeOk returns a tuple with the CmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmpType

`func (o *RecordHostIpv6addrDiscoveredData) SetCmpType(v string)`

SetCmpType sets CmpType field to given value.

### HasCmpType

`func (o *RecordHostIpv6addrDiscoveredData) HasCmpType() bool`

HasCmpType returns a boolean if a field has been set.

### GetVmiIpType

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiIpType() string`

GetVmiIpType returns the VmiIpType field if non-nil, zero value otherwise.

### GetVmiIpTypeOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiIpTypeOk() (*string, bool)`

GetVmiIpTypeOk returns a tuple with the VmiIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiIpType

`func (o *RecordHostIpv6addrDiscoveredData) SetVmiIpType(v string)`

SetVmiIpType sets VmiIpType field to given value.

### HasVmiIpType

`func (o *RecordHostIpv6addrDiscoveredData) HasVmiIpType() bool`

HasVmiIpType returns a boolean if a field has been set.

### GetVmiPrivateAddress

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiPrivateAddress() string`

GetVmiPrivateAddress returns the VmiPrivateAddress field if non-nil, zero value otherwise.

### GetVmiPrivateAddressOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiPrivateAddressOk() (*string, bool)`

GetVmiPrivateAddressOk returns a tuple with the VmiPrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiPrivateAddress

`func (o *RecordHostIpv6addrDiscoveredData) SetVmiPrivateAddress(v string)`

SetVmiPrivateAddress sets VmiPrivateAddress field to given value.

### HasVmiPrivateAddress

`func (o *RecordHostIpv6addrDiscoveredData) HasVmiPrivateAddress() bool`

HasVmiPrivateAddress returns a boolean if a field has been set.

### GetVmiIsPublicAddress

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiIsPublicAddress() bool`

GetVmiIsPublicAddress returns the VmiIsPublicAddress field if non-nil, zero value otherwise.

### GetVmiIsPublicAddressOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVmiIsPublicAddressOk() (*bool, bool)`

GetVmiIsPublicAddressOk returns a tuple with the VmiIsPublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiIsPublicAddress

`func (o *RecordHostIpv6addrDiscoveredData) SetVmiIsPublicAddress(v bool)`

SetVmiIsPublicAddress sets VmiIsPublicAddress field to given value.

### HasVmiIsPublicAddress

`func (o *RecordHostIpv6addrDiscoveredData) HasVmiIsPublicAddress() bool`

HasVmiIsPublicAddress returns a boolean if a field has been set.

### GetCiscoIseSsid

`func (o *RecordHostIpv6addrDiscoveredData) GetCiscoIseSsid() string`

GetCiscoIseSsid returns the CiscoIseSsid field if non-nil, zero value otherwise.

### GetCiscoIseSsidOk

`func (o *RecordHostIpv6addrDiscoveredData) GetCiscoIseSsidOk() (*string, bool)`

GetCiscoIseSsidOk returns a tuple with the CiscoIseSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseSsid

`func (o *RecordHostIpv6addrDiscoveredData) SetCiscoIseSsid(v string)`

SetCiscoIseSsid sets CiscoIseSsid field to given value.

### HasCiscoIseSsid

`func (o *RecordHostIpv6addrDiscoveredData) HasCiscoIseSsid() bool`

HasCiscoIseSsid returns a boolean if a field has been set.

### GetCiscoIseEndpointProfile

`func (o *RecordHostIpv6addrDiscoveredData) GetCiscoIseEndpointProfile() string`

GetCiscoIseEndpointProfile returns the CiscoIseEndpointProfile field if non-nil, zero value otherwise.

### GetCiscoIseEndpointProfileOk

`func (o *RecordHostIpv6addrDiscoveredData) GetCiscoIseEndpointProfileOk() (*string, bool)`

GetCiscoIseEndpointProfileOk returns a tuple with the CiscoIseEndpointProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseEndpointProfile

`func (o *RecordHostIpv6addrDiscoveredData) SetCiscoIseEndpointProfile(v string)`

SetCiscoIseEndpointProfile sets CiscoIseEndpointProfile field to given value.

### HasCiscoIseEndpointProfile

`func (o *RecordHostIpv6addrDiscoveredData) HasCiscoIseEndpointProfile() bool`

HasCiscoIseEndpointProfile returns a boolean if a field has been set.

### GetCiscoIseSessionState

`func (o *RecordHostIpv6addrDiscoveredData) GetCiscoIseSessionState() string`

GetCiscoIseSessionState returns the CiscoIseSessionState field if non-nil, zero value otherwise.

### GetCiscoIseSessionStateOk

`func (o *RecordHostIpv6addrDiscoveredData) GetCiscoIseSessionStateOk() (*string, bool)`

GetCiscoIseSessionStateOk returns a tuple with the CiscoIseSessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseSessionState

`func (o *RecordHostIpv6addrDiscoveredData) SetCiscoIseSessionState(v string)`

SetCiscoIseSessionState sets CiscoIseSessionState field to given value.

### HasCiscoIseSessionState

`func (o *RecordHostIpv6addrDiscoveredData) HasCiscoIseSessionState() bool`

HasCiscoIseSessionState returns a boolean if a field has been set.

### GetCiscoIseSecurityGroup

`func (o *RecordHostIpv6addrDiscoveredData) GetCiscoIseSecurityGroup() string`

GetCiscoIseSecurityGroup returns the CiscoIseSecurityGroup field if non-nil, zero value otherwise.

### GetCiscoIseSecurityGroupOk

`func (o *RecordHostIpv6addrDiscoveredData) GetCiscoIseSecurityGroupOk() (*string, bool)`

GetCiscoIseSecurityGroupOk returns a tuple with the CiscoIseSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseSecurityGroup

`func (o *RecordHostIpv6addrDiscoveredData) SetCiscoIseSecurityGroup(v string)`

SetCiscoIseSecurityGroup sets CiscoIseSecurityGroup field to given value.

### HasCiscoIseSecurityGroup

`func (o *RecordHostIpv6addrDiscoveredData) HasCiscoIseSecurityGroup() bool`

HasCiscoIseSecurityGroup returns a boolean if a field has been set.

### GetTaskName

`func (o *RecordHostIpv6addrDiscoveredData) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *RecordHostIpv6addrDiscoveredData) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *RecordHostIpv6addrDiscoveredData) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetNetworkComponentLocation

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentLocation() string`

GetNetworkComponentLocation returns the NetworkComponentLocation field if non-nil, zero value otherwise.

### GetNetworkComponentLocationOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentLocationOk() (*string, bool)`

GetNetworkComponentLocationOk returns a tuple with the NetworkComponentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentLocation

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentLocation(v string)`

SetNetworkComponentLocation sets NetworkComponentLocation field to given value.

### HasNetworkComponentLocation

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentLocation() bool`

HasNetworkComponentLocation returns a boolean if a field has been set.

### GetNetworkComponentContact

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentContact() string`

GetNetworkComponentContact returns the NetworkComponentContact field if non-nil, zero value otherwise.

### GetNetworkComponentContactOk

`func (o *RecordHostIpv6addrDiscoveredData) GetNetworkComponentContactOk() (*string, bool)`

GetNetworkComponentContactOk returns a tuple with the NetworkComponentContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentContact

`func (o *RecordHostIpv6addrDiscoveredData) SetNetworkComponentContact(v string)`

SetNetworkComponentContact sets NetworkComponentContact field to given value.

### HasNetworkComponentContact

`func (o *RecordHostIpv6addrDiscoveredData) HasNetworkComponentContact() bool`

HasNetworkComponentContact returns a boolean if a field has been set.

### GetDeviceLocation

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *RecordHostIpv6addrDiscoveredData) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.

### HasDeviceLocation

`func (o *RecordHostIpv6addrDiscoveredData) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### GetDeviceContact

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceContact() string`

GetDeviceContact returns the DeviceContact field if non-nil, zero value otherwise.

### GetDeviceContactOk

`func (o *RecordHostIpv6addrDiscoveredData) GetDeviceContactOk() (*string, bool)`

GetDeviceContactOk returns a tuple with the DeviceContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceContact

`func (o *RecordHostIpv6addrDiscoveredData) SetDeviceContact(v string)`

SetDeviceContact sets DeviceContact field to given value.

### HasDeviceContact

`func (o *RecordHostIpv6addrDiscoveredData) HasDeviceContact() bool`

HasDeviceContact returns a boolean if a field has been set.

### GetApName

`func (o *RecordHostIpv6addrDiscoveredData) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *RecordHostIpv6addrDiscoveredData) SetApName(v string)`

SetApName sets ApName field to given value.

### HasApName

`func (o *RecordHostIpv6addrDiscoveredData) HasApName() bool`

HasApName returns a boolean if a field has been set.

### GetApIpAddress

`func (o *RecordHostIpv6addrDiscoveredData) GetApIpAddress() string`

GetApIpAddress returns the ApIpAddress field if non-nil, zero value otherwise.

### GetApIpAddressOk

`func (o *RecordHostIpv6addrDiscoveredData) GetApIpAddressOk() (*string, bool)`

GetApIpAddressOk returns a tuple with the ApIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApIpAddress

`func (o *RecordHostIpv6addrDiscoveredData) SetApIpAddress(v string)`

SetApIpAddress sets ApIpAddress field to given value.

### HasApIpAddress

`func (o *RecordHostIpv6addrDiscoveredData) HasApIpAddress() bool`

HasApIpAddress returns a boolean if a field has been set.

### GetApSsid

`func (o *RecordHostIpv6addrDiscoveredData) GetApSsid() string`

GetApSsid returns the ApSsid field if non-nil, zero value otherwise.

### GetApSsidOk

`func (o *RecordHostIpv6addrDiscoveredData) GetApSsidOk() (*string, bool)`

GetApSsidOk returns a tuple with the ApSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApSsid

`func (o *RecordHostIpv6addrDiscoveredData) SetApSsid(v string)`

SetApSsid sets ApSsid field to given value.

### HasApSsid

`func (o *RecordHostIpv6addrDiscoveredData) HasApSsid() bool`

HasApSsid returns a boolean if a field has been set.

### GetBridgeDomain

`func (o *RecordHostIpv6addrDiscoveredData) GetBridgeDomain() string`

GetBridgeDomain returns the BridgeDomain field if non-nil, zero value otherwise.

### GetBridgeDomainOk

`func (o *RecordHostIpv6addrDiscoveredData) GetBridgeDomainOk() (*string, bool)`

GetBridgeDomainOk returns a tuple with the BridgeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeDomain

`func (o *RecordHostIpv6addrDiscoveredData) SetBridgeDomain(v string)`

SetBridgeDomain sets BridgeDomain field to given value.

### HasBridgeDomain

`func (o *RecordHostIpv6addrDiscoveredData) HasBridgeDomain() bool`

HasBridgeDomain returns a boolean if a field has been set.

### GetEndpointGroups

`func (o *RecordHostIpv6addrDiscoveredData) GetEndpointGroups() string`

GetEndpointGroups returns the EndpointGroups field if non-nil, zero value otherwise.

### GetEndpointGroupsOk

`func (o *RecordHostIpv6addrDiscoveredData) GetEndpointGroupsOk() (*string, bool)`

GetEndpointGroupsOk returns a tuple with the EndpointGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroups

`func (o *RecordHostIpv6addrDiscoveredData) SetEndpointGroups(v string)`

SetEndpointGroups sets EndpointGroups field to given value.

### HasEndpointGroups

`func (o *RecordHostIpv6addrDiscoveredData) HasEndpointGroups() bool`

HasEndpointGroups returns a boolean if a field has been set.

### GetTenant

`func (o *RecordHostIpv6addrDiscoveredData) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *RecordHostIpv6addrDiscoveredData) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *RecordHostIpv6addrDiscoveredData) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *RecordHostIpv6addrDiscoveredData) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetVrfName

`func (o *RecordHostIpv6addrDiscoveredData) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *RecordHostIpv6addrDiscoveredData) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *RecordHostIpv6addrDiscoveredData) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.

### GetVrfDescription

`func (o *RecordHostIpv6addrDiscoveredData) GetVrfDescription() string`

GetVrfDescription returns the VrfDescription field if non-nil, zero value otherwise.

### GetVrfDescriptionOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVrfDescriptionOk() (*string, bool)`

GetVrfDescriptionOk returns a tuple with the VrfDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfDescription

`func (o *RecordHostIpv6addrDiscoveredData) SetVrfDescription(v string)`

SetVrfDescription sets VrfDescription field to given value.

### HasVrfDescription

`func (o *RecordHostIpv6addrDiscoveredData) HasVrfDescription() bool`

HasVrfDescription returns a boolean if a field has been set.

### GetVrfRd

`func (o *RecordHostIpv6addrDiscoveredData) GetVrfRd() string`

GetVrfRd returns the VrfRd field if non-nil, zero value otherwise.

### GetVrfRdOk

`func (o *RecordHostIpv6addrDiscoveredData) GetVrfRdOk() (*string, bool)`

GetVrfRdOk returns a tuple with the VrfRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfRd

`func (o *RecordHostIpv6addrDiscoveredData) SetVrfRd(v string)`

SetVrfRd sets VrfRd field to given value.

### HasVrfRd

`func (o *RecordHostIpv6addrDiscoveredData) HasVrfRd() bool`

HasVrfRd returns a boolean if a field has been set.

### GetBgpAs

`func (o *RecordHostIpv6addrDiscoveredData) GetBgpAs() int64`

GetBgpAs returns the BgpAs field if non-nil, zero value otherwise.

### GetBgpAsOk

`func (o *RecordHostIpv6addrDiscoveredData) GetBgpAsOk() (*int64, bool)`

GetBgpAsOk returns a tuple with the BgpAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAs

`func (o *RecordHostIpv6addrDiscoveredData) SetBgpAs(v int64)`

SetBgpAs sets BgpAs field to given value.

### HasBgpAs

`func (o *RecordHostIpv6addrDiscoveredData) HasBgpAs() bool`

HasBgpAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


