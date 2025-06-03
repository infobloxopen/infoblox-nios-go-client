# Ipv4addressDiscoveredData

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

### NewIpv4addressDiscoveredData

`func NewIpv4addressDiscoveredData() *Ipv4addressDiscoveredData`

NewIpv4addressDiscoveredData instantiates a new Ipv4addressDiscoveredData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv4addressDiscoveredDataWithDefaults

`func NewIpv4addressDiscoveredDataWithDefaults() *Ipv4addressDiscoveredData`

NewIpv4addressDiscoveredDataWithDefaults instantiates a new Ipv4addressDiscoveredData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *Ipv4addressDiscoveredData) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *Ipv4addressDiscoveredData) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *Ipv4addressDiscoveredData) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *Ipv4addressDiscoveredData) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDevicePortName

`func (o *Ipv4addressDiscoveredData) GetDevicePortName() string`

GetDevicePortName returns the DevicePortName field if non-nil, zero value otherwise.

### GetDevicePortNameOk

`func (o *Ipv4addressDiscoveredData) GetDevicePortNameOk() (*string, bool)`

GetDevicePortNameOk returns a tuple with the DevicePortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortName

`func (o *Ipv4addressDiscoveredData) SetDevicePortName(v string)`

SetDevicePortName sets DevicePortName field to given value.

### HasDevicePortName

`func (o *Ipv4addressDiscoveredData) HasDevicePortName() bool`

HasDevicePortName returns a boolean if a field has been set.

### GetDevicePortType

`func (o *Ipv4addressDiscoveredData) GetDevicePortType() string`

GetDevicePortType returns the DevicePortType field if non-nil, zero value otherwise.

### GetDevicePortTypeOk

`func (o *Ipv4addressDiscoveredData) GetDevicePortTypeOk() (*string, bool)`

GetDevicePortTypeOk returns a tuple with the DevicePortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortType

`func (o *Ipv4addressDiscoveredData) SetDevicePortType(v string)`

SetDevicePortType sets DevicePortType field to given value.

### HasDevicePortType

`func (o *Ipv4addressDiscoveredData) HasDevicePortType() bool`

HasDevicePortType returns a boolean if a field has been set.

### GetDeviceType

`func (o *Ipv4addressDiscoveredData) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *Ipv4addressDiscoveredData) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *Ipv4addressDiscoveredData) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *Ipv4addressDiscoveredData) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *Ipv4addressDiscoveredData) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *Ipv4addressDiscoveredData) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *Ipv4addressDiscoveredData) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *Ipv4addressDiscoveredData) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDiscoveredName

`func (o *Ipv4addressDiscoveredData) GetDiscoveredName() string`

GetDiscoveredName returns the DiscoveredName field if non-nil, zero value otherwise.

### GetDiscoveredNameOk

`func (o *Ipv4addressDiscoveredData) GetDiscoveredNameOk() (*string, bool)`

GetDiscoveredNameOk returns a tuple with the DiscoveredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredName

`func (o *Ipv4addressDiscoveredData) SetDiscoveredName(v string)`

SetDiscoveredName sets DiscoveredName field to given value.

### HasDiscoveredName

`func (o *Ipv4addressDiscoveredData) HasDiscoveredName() bool`

HasDiscoveredName returns a boolean if a field has been set.

### GetDiscoverer

`func (o *Ipv4addressDiscoveredData) GetDiscoverer() string`

GetDiscoverer returns the Discoverer field if non-nil, zero value otherwise.

### GetDiscovererOk

`func (o *Ipv4addressDiscoveredData) GetDiscovererOk() (*string, bool)`

GetDiscovererOk returns a tuple with the Discoverer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverer

`func (o *Ipv4addressDiscoveredData) SetDiscoverer(v string)`

SetDiscoverer sets Discoverer field to given value.

### HasDiscoverer

`func (o *Ipv4addressDiscoveredData) HasDiscoverer() bool`

HasDiscoverer returns a boolean if a field has been set.

### GetDuid

`func (o *Ipv4addressDiscoveredData) GetDuid() string`

GetDuid returns the Duid field if non-nil, zero value otherwise.

### GetDuidOk

`func (o *Ipv4addressDiscoveredData) GetDuidOk() (*string, bool)`

GetDuidOk returns a tuple with the Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuid

`func (o *Ipv4addressDiscoveredData) SetDuid(v string)`

SetDuid sets Duid field to given value.

### HasDuid

`func (o *Ipv4addressDiscoveredData) HasDuid() bool`

HasDuid returns a boolean if a field has been set.

### GetFirstDiscovered

`func (o *Ipv4addressDiscoveredData) GetFirstDiscovered() int64`

GetFirstDiscovered returns the FirstDiscovered field if non-nil, zero value otherwise.

### GetFirstDiscoveredOk

`func (o *Ipv4addressDiscoveredData) GetFirstDiscoveredOk() (*int64, bool)`

GetFirstDiscoveredOk returns a tuple with the FirstDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDiscovered

`func (o *Ipv4addressDiscoveredData) SetFirstDiscovered(v int64)`

SetFirstDiscovered sets FirstDiscovered field to given value.

### HasFirstDiscovered

`func (o *Ipv4addressDiscoveredData) HasFirstDiscovered() bool`

HasFirstDiscovered returns a boolean if a field has been set.

### GetIprgNo

`func (o *Ipv4addressDiscoveredData) GetIprgNo() int64`

GetIprgNo returns the IprgNo field if non-nil, zero value otherwise.

### GetIprgNoOk

`func (o *Ipv4addressDiscoveredData) GetIprgNoOk() (*int64, bool)`

GetIprgNoOk returns a tuple with the IprgNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprgNo

`func (o *Ipv4addressDiscoveredData) SetIprgNo(v int64)`

SetIprgNo sets IprgNo field to given value.

### HasIprgNo

`func (o *Ipv4addressDiscoveredData) HasIprgNo() bool`

HasIprgNo returns a boolean if a field has been set.

### GetIprgState

`func (o *Ipv4addressDiscoveredData) GetIprgState() string`

GetIprgState returns the IprgState field if non-nil, zero value otherwise.

### GetIprgStateOk

`func (o *Ipv4addressDiscoveredData) GetIprgStateOk() (*string, bool)`

GetIprgStateOk returns a tuple with the IprgState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprgState

`func (o *Ipv4addressDiscoveredData) SetIprgState(v string)`

SetIprgState sets IprgState field to given value.

### HasIprgState

`func (o *Ipv4addressDiscoveredData) HasIprgState() bool`

HasIprgState returns a boolean if a field has been set.

### GetIprgType

`func (o *Ipv4addressDiscoveredData) GetIprgType() string`

GetIprgType returns the IprgType field if non-nil, zero value otherwise.

### GetIprgTypeOk

`func (o *Ipv4addressDiscoveredData) GetIprgTypeOk() (*string, bool)`

GetIprgTypeOk returns a tuple with the IprgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprgType

`func (o *Ipv4addressDiscoveredData) SetIprgType(v string)`

SetIprgType sets IprgType field to given value.

### HasIprgType

`func (o *Ipv4addressDiscoveredData) HasIprgType() bool`

HasIprgType returns a boolean if a field has been set.

### GetLastDiscovered

`func (o *Ipv4addressDiscoveredData) GetLastDiscovered() int64`

GetLastDiscovered returns the LastDiscovered field if non-nil, zero value otherwise.

### GetLastDiscoveredOk

`func (o *Ipv4addressDiscoveredData) GetLastDiscoveredOk() (*int64, bool)`

GetLastDiscoveredOk returns a tuple with the LastDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscovered

`func (o *Ipv4addressDiscoveredData) SetLastDiscovered(v int64)`

SetLastDiscovered sets LastDiscovered field to given value.

### HasLastDiscovered

`func (o *Ipv4addressDiscoveredData) HasLastDiscovered() bool`

HasLastDiscovered returns a boolean if a field has been set.

### GetMacAddress

`func (o *Ipv4addressDiscoveredData) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *Ipv4addressDiscoveredData) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *Ipv4addressDiscoveredData) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *Ipv4addressDiscoveredData) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *Ipv4addressDiscoveredData) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *Ipv4addressDiscoveredData) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *Ipv4addressDiscoveredData) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *Ipv4addressDiscoveredData) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetNetbiosName

`func (o *Ipv4addressDiscoveredData) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *Ipv4addressDiscoveredData) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *Ipv4addressDiscoveredData) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *Ipv4addressDiscoveredData) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.

### GetNetworkComponentDescription

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentDescription() string`

GetNetworkComponentDescription returns the NetworkComponentDescription field if non-nil, zero value otherwise.

### GetNetworkComponentDescriptionOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentDescriptionOk() (*string, bool)`

GetNetworkComponentDescriptionOk returns a tuple with the NetworkComponentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentDescription

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentDescription(v string)`

SetNetworkComponentDescription sets NetworkComponentDescription field to given value.

### HasNetworkComponentDescription

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentDescription() bool`

HasNetworkComponentDescription returns a boolean if a field has been set.

### GetNetworkComponentIp

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentIp() string`

GetNetworkComponentIp returns the NetworkComponentIp field if non-nil, zero value otherwise.

### GetNetworkComponentIpOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentIpOk() (*string, bool)`

GetNetworkComponentIpOk returns a tuple with the NetworkComponentIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentIp

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentIp(v string)`

SetNetworkComponentIp sets NetworkComponentIp field to given value.

### HasNetworkComponentIp

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentIp() bool`

HasNetworkComponentIp returns a boolean if a field has been set.

### GetNetworkComponentModel

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentModel() string`

GetNetworkComponentModel returns the NetworkComponentModel field if non-nil, zero value otherwise.

### GetNetworkComponentModelOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentModelOk() (*string, bool)`

GetNetworkComponentModelOk returns a tuple with the NetworkComponentModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentModel

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentModel(v string)`

SetNetworkComponentModel sets NetworkComponentModel field to given value.

### HasNetworkComponentModel

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentModel() bool`

HasNetworkComponentModel returns a boolean if a field has been set.

### GetNetworkComponentName

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentName() string`

GetNetworkComponentName returns the NetworkComponentName field if non-nil, zero value otherwise.

### GetNetworkComponentNameOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentNameOk() (*string, bool)`

GetNetworkComponentNameOk returns a tuple with the NetworkComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentName

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentName(v string)`

SetNetworkComponentName sets NetworkComponentName field to given value.

### HasNetworkComponentName

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentName() bool`

HasNetworkComponentName returns a boolean if a field has been set.

### GetNetworkComponentPortDescription

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentPortDescription() string`

GetNetworkComponentPortDescription returns the NetworkComponentPortDescription field if non-nil, zero value otherwise.

### GetNetworkComponentPortDescriptionOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentPortDescriptionOk() (*string, bool)`

GetNetworkComponentPortDescriptionOk returns a tuple with the NetworkComponentPortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentPortDescription

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentPortDescription(v string)`

SetNetworkComponentPortDescription sets NetworkComponentPortDescription field to given value.

### HasNetworkComponentPortDescription

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentPortDescription() bool`

HasNetworkComponentPortDescription returns a boolean if a field has been set.

### GetNetworkComponentPortName

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentPortName() string`

GetNetworkComponentPortName returns the NetworkComponentPortName field if non-nil, zero value otherwise.

### GetNetworkComponentPortNameOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentPortNameOk() (*string, bool)`

GetNetworkComponentPortNameOk returns a tuple with the NetworkComponentPortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentPortName

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentPortName(v string)`

SetNetworkComponentPortName sets NetworkComponentPortName field to given value.

### HasNetworkComponentPortName

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentPortName() bool`

HasNetworkComponentPortName returns a boolean if a field has been set.

### GetNetworkComponentPortNumber

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentPortNumber() string`

GetNetworkComponentPortNumber returns the NetworkComponentPortNumber field if non-nil, zero value otherwise.

### GetNetworkComponentPortNumberOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentPortNumberOk() (*string, bool)`

GetNetworkComponentPortNumberOk returns a tuple with the NetworkComponentPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentPortNumber

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentPortNumber(v string)`

SetNetworkComponentPortNumber sets NetworkComponentPortNumber field to given value.

### HasNetworkComponentPortNumber

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentPortNumber() bool`

HasNetworkComponentPortNumber returns a boolean if a field has been set.

### GetNetworkComponentType

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentType() string`

GetNetworkComponentType returns the NetworkComponentType field if non-nil, zero value otherwise.

### GetNetworkComponentTypeOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentTypeOk() (*string, bool)`

GetNetworkComponentTypeOk returns a tuple with the NetworkComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentType

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentType(v string)`

SetNetworkComponentType sets NetworkComponentType field to given value.

### HasNetworkComponentType

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentType() bool`

HasNetworkComponentType returns a boolean if a field has been set.

### GetNetworkComponentVendor

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentVendor() string`

GetNetworkComponentVendor returns the NetworkComponentVendor field if non-nil, zero value otherwise.

### GetNetworkComponentVendorOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentVendorOk() (*string, bool)`

GetNetworkComponentVendorOk returns a tuple with the NetworkComponentVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentVendor

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentVendor(v string)`

SetNetworkComponentVendor sets NetworkComponentVendor field to given value.

### HasNetworkComponentVendor

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentVendor() bool`

HasNetworkComponentVendor returns a boolean if a field has been set.

### GetOpenPorts

`func (o *Ipv4addressDiscoveredData) GetOpenPorts() string`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *Ipv4addressDiscoveredData) GetOpenPortsOk() (*string, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *Ipv4addressDiscoveredData) SetOpenPorts(v string)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *Ipv4addressDiscoveredData) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetOs

`func (o *Ipv4addressDiscoveredData) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *Ipv4addressDiscoveredData) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *Ipv4addressDiscoveredData) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *Ipv4addressDiscoveredData) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPortDuplex

`func (o *Ipv4addressDiscoveredData) GetPortDuplex() string`

GetPortDuplex returns the PortDuplex field if non-nil, zero value otherwise.

### GetPortDuplexOk

`func (o *Ipv4addressDiscoveredData) GetPortDuplexOk() (*string, bool)`

GetPortDuplexOk returns a tuple with the PortDuplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDuplex

`func (o *Ipv4addressDiscoveredData) SetPortDuplex(v string)`

SetPortDuplex sets PortDuplex field to given value.

### HasPortDuplex

`func (o *Ipv4addressDiscoveredData) HasPortDuplex() bool`

HasPortDuplex returns a boolean if a field has been set.

### GetPortLinkStatus

`func (o *Ipv4addressDiscoveredData) GetPortLinkStatus() string`

GetPortLinkStatus returns the PortLinkStatus field if non-nil, zero value otherwise.

### GetPortLinkStatusOk

`func (o *Ipv4addressDiscoveredData) GetPortLinkStatusOk() (*string, bool)`

GetPortLinkStatusOk returns a tuple with the PortLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLinkStatus

`func (o *Ipv4addressDiscoveredData) SetPortLinkStatus(v string)`

SetPortLinkStatus sets PortLinkStatus field to given value.

### HasPortLinkStatus

`func (o *Ipv4addressDiscoveredData) HasPortLinkStatus() bool`

HasPortLinkStatus returns a boolean if a field has been set.

### GetPortSpeed

`func (o *Ipv4addressDiscoveredData) GetPortSpeed() string`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *Ipv4addressDiscoveredData) GetPortSpeedOk() (*string, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *Ipv4addressDiscoveredData) SetPortSpeed(v string)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *Ipv4addressDiscoveredData) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### GetPortStatus

`func (o *Ipv4addressDiscoveredData) GetPortStatus() string`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *Ipv4addressDiscoveredData) GetPortStatusOk() (*string, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *Ipv4addressDiscoveredData) SetPortStatus(v string)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *Ipv4addressDiscoveredData) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetPortType

`func (o *Ipv4addressDiscoveredData) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *Ipv4addressDiscoveredData) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *Ipv4addressDiscoveredData) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *Ipv4addressDiscoveredData) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetPortVlanDescription

`func (o *Ipv4addressDiscoveredData) GetPortVlanDescription() string`

GetPortVlanDescription returns the PortVlanDescription field if non-nil, zero value otherwise.

### GetPortVlanDescriptionOk

`func (o *Ipv4addressDiscoveredData) GetPortVlanDescriptionOk() (*string, bool)`

GetPortVlanDescriptionOk returns a tuple with the PortVlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanDescription

`func (o *Ipv4addressDiscoveredData) SetPortVlanDescription(v string)`

SetPortVlanDescription sets PortVlanDescription field to given value.

### HasPortVlanDescription

`func (o *Ipv4addressDiscoveredData) HasPortVlanDescription() bool`

HasPortVlanDescription returns a boolean if a field has been set.

### GetPortVlanName

`func (o *Ipv4addressDiscoveredData) GetPortVlanName() string`

GetPortVlanName returns the PortVlanName field if non-nil, zero value otherwise.

### GetPortVlanNameOk

`func (o *Ipv4addressDiscoveredData) GetPortVlanNameOk() (*string, bool)`

GetPortVlanNameOk returns a tuple with the PortVlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanName

`func (o *Ipv4addressDiscoveredData) SetPortVlanName(v string)`

SetPortVlanName sets PortVlanName field to given value.

### HasPortVlanName

`func (o *Ipv4addressDiscoveredData) HasPortVlanName() bool`

HasPortVlanName returns a boolean if a field has been set.

### GetPortVlanNumber

`func (o *Ipv4addressDiscoveredData) GetPortVlanNumber() string`

GetPortVlanNumber returns the PortVlanNumber field if non-nil, zero value otherwise.

### GetPortVlanNumberOk

`func (o *Ipv4addressDiscoveredData) GetPortVlanNumberOk() (*string, bool)`

GetPortVlanNumberOk returns a tuple with the PortVlanNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanNumber

`func (o *Ipv4addressDiscoveredData) SetPortVlanNumber(v string)`

SetPortVlanNumber sets PortVlanNumber field to given value.

### HasPortVlanNumber

`func (o *Ipv4addressDiscoveredData) HasPortVlanNumber() bool`

HasPortVlanNumber returns a boolean if a field has been set.

### GetVAdapter

`func (o *Ipv4addressDiscoveredData) GetVAdapter() string`

GetVAdapter returns the VAdapter field if non-nil, zero value otherwise.

### GetVAdapterOk

`func (o *Ipv4addressDiscoveredData) GetVAdapterOk() (*string, bool)`

GetVAdapterOk returns a tuple with the VAdapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAdapter

`func (o *Ipv4addressDiscoveredData) SetVAdapter(v string)`

SetVAdapter sets VAdapter field to given value.

### HasVAdapter

`func (o *Ipv4addressDiscoveredData) HasVAdapter() bool`

HasVAdapter returns a boolean if a field has been set.

### GetVCluster

`func (o *Ipv4addressDiscoveredData) GetVCluster() string`

GetVCluster returns the VCluster field if non-nil, zero value otherwise.

### GetVClusterOk

`func (o *Ipv4addressDiscoveredData) GetVClusterOk() (*string, bool)`

GetVClusterOk returns a tuple with the VCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCluster

`func (o *Ipv4addressDiscoveredData) SetVCluster(v string)`

SetVCluster sets VCluster field to given value.

### HasVCluster

`func (o *Ipv4addressDiscoveredData) HasVCluster() bool`

HasVCluster returns a boolean if a field has been set.

### GetVDatacenter

`func (o *Ipv4addressDiscoveredData) GetVDatacenter() string`

GetVDatacenter returns the VDatacenter field if non-nil, zero value otherwise.

### GetVDatacenterOk

`func (o *Ipv4addressDiscoveredData) GetVDatacenterOk() (*string, bool)`

GetVDatacenterOk returns a tuple with the VDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDatacenter

`func (o *Ipv4addressDiscoveredData) SetVDatacenter(v string)`

SetVDatacenter sets VDatacenter field to given value.

### HasVDatacenter

`func (o *Ipv4addressDiscoveredData) HasVDatacenter() bool`

HasVDatacenter returns a boolean if a field has been set.

### GetVEntityName

`func (o *Ipv4addressDiscoveredData) GetVEntityName() string`

GetVEntityName returns the VEntityName field if non-nil, zero value otherwise.

### GetVEntityNameOk

`func (o *Ipv4addressDiscoveredData) GetVEntityNameOk() (*string, bool)`

GetVEntityNameOk returns a tuple with the VEntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVEntityName

`func (o *Ipv4addressDiscoveredData) SetVEntityName(v string)`

SetVEntityName sets VEntityName field to given value.

### HasVEntityName

`func (o *Ipv4addressDiscoveredData) HasVEntityName() bool`

HasVEntityName returns a boolean if a field has been set.

### GetVEntityType

`func (o *Ipv4addressDiscoveredData) GetVEntityType() string`

GetVEntityType returns the VEntityType field if non-nil, zero value otherwise.

### GetVEntityTypeOk

`func (o *Ipv4addressDiscoveredData) GetVEntityTypeOk() (*string, bool)`

GetVEntityTypeOk returns a tuple with the VEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVEntityType

`func (o *Ipv4addressDiscoveredData) SetVEntityType(v string)`

SetVEntityType sets VEntityType field to given value.

### HasVEntityType

`func (o *Ipv4addressDiscoveredData) HasVEntityType() bool`

HasVEntityType returns a boolean if a field has been set.

### GetVHost

`func (o *Ipv4addressDiscoveredData) GetVHost() string`

GetVHost returns the VHost field if non-nil, zero value otherwise.

### GetVHostOk

`func (o *Ipv4addressDiscoveredData) GetVHostOk() (*string, bool)`

GetVHostOk returns a tuple with the VHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVHost

`func (o *Ipv4addressDiscoveredData) SetVHost(v string)`

SetVHost sets VHost field to given value.

### HasVHost

`func (o *Ipv4addressDiscoveredData) HasVHost() bool`

HasVHost returns a boolean if a field has been set.

### GetVSwitch

`func (o *Ipv4addressDiscoveredData) GetVSwitch() string`

GetVSwitch returns the VSwitch field if non-nil, zero value otherwise.

### GetVSwitchOk

`func (o *Ipv4addressDiscoveredData) GetVSwitchOk() (*string, bool)`

GetVSwitchOk returns a tuple with the VSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSwitch

`func (o *Ipv4addressDiscoveredData) SetVSwitch(v string)`

SetVSwitch sets VSwitch field to given value.

### HasVSwitch

`func (o *Ipv4addressDiscoveredData) HasVSwitch() bool`

HasVSwitch returns a boolean if a field has been set.

### GetVmiName

`func (o *Ipv4addressDiscoveredData) GetVmiName() string`

GetVmiName returns the VmiName field if non-nil, zero value otherwise.

### GetVmiNameOk

`func (o *Ipv4addressDiscoveredData) GetVmiNameOk() (*string, bool)`

GetVmiNameOk returns a tuple with the VmiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiName

`func (o *Ipv4addressDiscoveredData) SetVmiName(v string)`

SetVmiName sets VmiName field to given value.

### HasVmiName

`func (o *Ipv4addressDiscoveredData) HasVmiName() bool`

HasVmiName returns a boolean if a field has been set.

### GetVmiId

`func (o *Ipv4addressDiscoveredData) GetVmiId() string`

GetVmiId returns the VmiId field if non-nil, zero value otherwise.

### GetVmiIdOk

`func (o *Ipv4addressDiscoveredData) GetVmiIdOk() (*string, bool)`

GetVmiIdOk returns a tuple with the VmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiId

`func (o *Ipv4addressDiscoveredData) SetVmiId(v string)`

SetVmiId sets VmiId field to given value.

### HasVmiId

`func (o *Ipv4addressDiscoveredData) HasVmiId() bool`

HasVmiId returns a boolean if a field has been set.

### GetVlanPortGroup

`func (o *Ipv4addressDiscoveredData) GetVlanPortGroup() string`

GetVlanPortGroup returns the VlanPortGroup field if non-nil, zero value otherwise.

### GetVlanPortGroupOk

`func (o *Ipv4addressDiscoveredData) GetVlanPortGroupOk() (*string, bool)`

GetVlanPortGroupOk returns a tuple with the VlanPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPortGroup

`func (o *Ipv4addressDiscoveredData) SetVlanPortGroup(v string)`

SetVlanPortGroup sets VlanPortGroup field to given value.

### HasVlanPortGroup

`func (o *Ipv4addressDiscoveredData) HasVlanPortGroup() bool`

HasVlanPortGroup returns a boolean if a field has been set.

### GetVswitchName

`func (o *Ipv4addressDiscoveredData) GetVswitchName() string`

GetVswitchName returns the VswitchName field if non-nil, zero value otherwise.

### GetVswitchNameOk

`func (o *Ipv4addressDiscoveredData) GetVswitchNameOk() (*string, bool)`

GetVswitchNameOk returns a tuple with the VswitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchName

`func (o *Ipv4addressDiscoveredData) SetVswitchName(v string)`

SetVswitchName sets VswitchName field to given value.

### HasVswitchName

`func (o *Ipv4addressDiscoveredData) HasVswitchName() bool`

HasVswitchName returns a boolean if a field has been set.

### GetVswitchId

`func (o *Ipv4addressDiscoveredData) GetVswitchId() string`

GetVswitchId returns the VswitchId field if non-nil, zero value otherwise.

### GetVswitchIdOk

`func (o *Ipv4addressDiscoveredData) GetVswitchIdOk() (*string, bool)`

GetVswitchIdOk returns a tuple with the VswitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchId

`func (o *Ipv4addressDiscoveredData) SetVswitchId(v string)`

SetVswitchId sets VswitchId field to given value.

### HasVswitchId

`func (o *Ipv4addressDiscoveredData) HasVswitchId() bool`

HasVswitchId returns a boolean if a field has been set.

### GetVswitchType

`func (o *Ipv4addressDiscoveredData) GetVswitchType() string`

GetVswitchType returns the VswitchType field if non-nil, zero value otherwise.

### GetVswitchTypeOk

`func (o *Ipv4addressDiscoveredData) GetVswitchTypeOk() (*string, bool)`

GetVswitchTypeOk returns a tuple with the VswitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchType

`func (o *Ipv4addressDiscoveredData) SetVswitchType(v string)`

SetVswitchType sets VswitchType field to given value.

### HasVswitchType

`func (o *Ipv4addressDiscoveredData) HasVswitchType() bool`

HasVswitchType returns a boolean if a field has been set.

### GetVswitchIpv6Enabled

`func (o *Ipv4addressDiscoveredData) GetVswitchIpv6Enabled() bool`

GetVswitchIpv6Enabled returns the VswitchIpv6Enabled field if non-nil, zero value otherwise.

### GetVswitchIpv6EnabledOk

`func (o *Ipv4addressDiscoveredData) GetVswitchIpv6EnabledOk() (*bool, bool)`

GetVswitchIpv6EnabledOk returns a tuple with the VswitchIpv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchIpv6Enabled

`func (o *Ipv4addressDiscoveredData) SetVswitchIpv6Enabled(v bool)`

SetVswitchIpv6Enabled sets VswitchIpv6Enabled field to given value.

### HasVswitchIpv6Enabled

`func (o *Ipv4addressDiscoveredData) HasVswitchIpv6Enabled() bool`

HasVswitchIpv6Enabled returns a boolean if a field has been set.

### GetVportName

`func (o *Ipv4addressDiscoveredData) GetVportName() string`

GetVportName returns the VportName field if non-nil, zero value otherwise.

### GetVportNameOk

`func (o *Ipv4addressDiscoveredData) GetVportNameOk() (*string, bool)`

GetVportNameOk returns a tuple with the VportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportName

`func (o *Ipv4addressDiscoveredData) SetVportName(v string)`

SetVportName sets VportName field to given value.

### HasVportName

`func (o *Ipv4addressDiscoveredData) HasVportName() bool`

HasVportName returns a boolean if a field has been set.

### GetVportMacAddress

`func (o *Ipv4addressDiscoveredData) GetVportMacAddress() string`

GetVportMacAddress returns the VportMacAddress field if non-nil, zero value otherwise.

### GetVportMacAddressOk

`func (o *Ipv4addressDiscoveredData) GetVportMacAddressOk() (*string, bool)`

GetVportMacAddressOk returns a tuple with the VportMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportMacAddress

`func (o *Ipv4addressDiscoveredData) SetVportMacAddress(v string)`

SetVportMacAddress sets VportMacAddress field to given value.

### HasVportMacAddress

`func (o *Ipv4addressDiscoveredData) HasVportMacAddress() bool`

HasVportMacAddress returns a boolean if a field has been set.

### GetVportLinkStatus

`func (o *Ipv4addressDiscoveredData) GetVportLinkStatus() string`

GetVportLinkStatus returns the VportLinkStatus field if non-nil, zero value otherwise.

### GetVportLinkStatusOk

`func (o *Ipv4addressDiscoveredData) GetVportLinkStatusOk() (*string, bool)`

GetVportLinkStatusOk returns a tuple with the VportLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportLinkStatus

`func (o *Ipv4addressDiscoveredData) SetVportLinkStatus(v string)`

SetVportLinkStatus sets VportLinkStatus field to given value.

### HasVportLinkStatus

`func (o *Ipv4addressDiscoveredData) HasVportLinkStatus() bool`

HasVportLinkStatus returns a boolean if a field has been set.

### GetVportConfSpeed

`func (o *Ipv4addressDiscoveredData) GetVportConfSpeed() string`

GetVportConfSpeed returns the VportConfSpeed field if non-nil, zero value otherwise.

### GetVportConfSpeedOk

`func (o *Ipv4addressDiscoveredData) GetVportConfSpeedOk() (*string, bool)`

GetVportConfSpeedOk returns a tuple with the VportConfSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportConfSpeed

`func (o *Ipv4addressDiscoveredData) SetVportConfSpeed(v string)`

SetVportConfSpeed sets VportConfSpeed field to given value.

### HasVportConfSpeed

`func (o *Ipv4addressDiscoveredData) HasVportConfSpeed() bool`

HasVportConfSpeed returns a boolean if a field has been set.

### GetVportConfMode

`func (o *Ipv4addressDiscoveredData) GetVportConfMode() string`

GetVportConfMode returns the VportConfMode field if non-nil, zero value otherwise.

### GetVportConfModeOk

`func (o *Ipv4addressDiscoveredData) GetVportConfModeOk() (*string, bool)`

GetVportConfModeOk returns a tuple with the VportConfMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportConfMode

`func (o *Ipv4addressDiscoveredData) SetVportConfMode(v string)`

SetVportConfMode sets VportConfMode field to given value.

### HasVportConfMode

`func (o *Ipv4addressDiscoveredData) HasVportConfMode() bool`

HasVportConfMode returns a boolean if a field has been set.

### GetVportSpeed

`func (o *Ipv4addressDiscoveredData) GetVportSpeed() string`

GetVportSpeed returns the VportSpeed field if non-nil, zero value otherwise.

### GetVportSpeedOk

`func (o *Ipv4addressDiscoveredData) GetVportSpeedOk() (*string, bool)`

GetVportSpeedOk returns a tuple with the VportSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportSpeed

`func (o *Ipv4addressDiscoveredData) SetVportSpeed(v string)`

SetVportSpeed sets VportSpeed field to given value.

### HasVportSpeed

`func (o *Ipv4addressDiscoveredData) HasVportSpeed() bool`

HasVportSpeed returns a boolean if a field has been set.

### GetVportMode

`func (o *Ipv4addressDiscoveredData) GetVportMode() string`

GetVportMode returns the VportMode field if non-nil, zero value otherwise.

### GetVportModeOk

`func (o *Ipv4addressDiscoveredData) GetVportModeOk() (*string, bool)`

GetVportModeOk returns a tuple with the VportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportMode

`func (o *Ipv4addressDiscoveredData) SetVportMode(v string)`

SetVportMode sets VportMode field to given value.

### HasVportMode

`func (o *Ipv4addressDiscoveredData) HasVportMode() bool`

HasVportMode returns a boolean if a field has been set.

### GetVswitchSegmentType

`func (o *Ipv4addressDiscoveredData) GetVswitchSegmentType() string`

GetVswitchSegmentType returns the VswitchSegmentType field if non-nil, zero value otherwise.

### GetVswitchSegmentTypeOk

`func (o *Ipv4addressDiscoveredData) GetVswitchSegmentTypeOk() (*string, bool)`

GetVswitchSegmentTypeOk returns a tuple with the VswitchSegmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentType

`func (o *Ipv4addressDiscoveredData) SetVswitchSegmentType(v string)`

SetVswitchSegmentType sets VswitchSegmentType field to given value.

### HasVswitchSegmentType

`func (o *Ipv4addressDiscoveredData) HasVswitchSegmentType() bool`

HasVswitchSegmentType returns a boolean if a field has been set.

### GetVswitchSegmentName

`func (o *Ipv4addressDiscoveredData) GetVswitchSegmentName() string`

GetVswitchSegmentName returns the VswitchSegmentName field if non-nil, zero value otherwise.

### GetVswitchSegmentNameOk

`func (o *Ipv4addressDiscoveredData) GetVswitchSegmentNameOk() (*string, bool)`

GetVswitchSegmentNameOk returns a tuple with the VswitchSegmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentName

`func (o *Ipv4addressDiscoveredData) SetVswitchSegmentName(v string)`

SetVswitchSegmentName sets VswitchSegmentName field to given value.

### HasVswitchSegmentName

`func (o *Ipv4addressDiscoveredData) HasVswitchSegmentName() bool`

HasVswitchSegmentName returns a boolean if a field has been set.

### GetVswitchSegmentId

`func (o *Ipv4addressDiscoveredData) GetVswitchSegmentId() string`

GetVswitchSegmentId returns the VswitchSegmentId field if non-nil, zero value otherwise.

### GetVswitchSegmentIdOk

`func (o *Ipv4addressDiscoveredData) GetVswitchSegmentIdOk() (*string, bool)`

GetVswitchSegmentIdOk returns a tuple with the VswitchSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentId

`func (o *Ipv4addressDiscoveredData) SetVswitchSegmentId(v string)`

SetVswitchSegmentId sets VswitchSegmentId field to given value.

### HasVswitchSegmentId

`func (o *Ipv4addressDiscoveredData) HasVswitchSegmentId() bool`

HasVswitchSegmentId returns a boolean if a field has been set.

### GetVswitchSegmentPortGroup

`func (o *Ipv4addressDiscoveredData) GetVswitchSegmentPortGroup() string`

GetVswitchSegmentPortGroup returns the VswitchSegmentPortGroup field if non-nil, zero value otherwise.

### GetVswitchSegmentPortGroupOk

`func (o *Ipv4addressDiscoveredData) GetVswitchSegmentPortGroupOk() (*string, bool)`

GetVswitchSegmentPortGroupOk returns a tuple with the VswitchSegmentPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentPortGroup

`func (o *Ipv4addressDiscoveredData) SetVswitchSegmentPortGroup(v string)`

SetVswitchSegmentPortGroup sets VswitchSegmentPortGroup field to given value.

### HasVswitchSegmentPortGroup

`func (o *Ipv4addressDiscoveredData) HasVswitchSegmentPortGroup() bool`

HasVswitchSegmentPortGroup returns a boolean if a field has been set.

### GetVswitchAvailablePortsCount

`func (o *Ipv4addressDiscoveredData) GetVswitchAvailablePortsCount() int64`

GetVswitchAvailablePortsCount returns the VswitchAvailablePortsCount field if non-nil, zero value otherwise.

### GetVswitchAvailablePortsCountOk

`func (o *Ipv4addressDiscoveredData) GetVswitchAvailablePortsCountOk() (*int64, bool)`

GetVswitchAvailablePortsCountOk returns a tuple with the VswitchAvailablePortsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchAvailablePortsCount

`func (o *Ipv4addressDiscoveredData) SetVswitchAvailablePortsCount(v int64)`

SetVswitchAvailablePortsCount sets VswitchAvailablePortsCount field to given value.

### HasVswitchAvailablePortsCount

`func (o *Ipv4addressDiscoveredData) HasVswitchAvailablePortsCount() bool`

HasVswitchAvailablePortsCount returns a boolean if a field has been set.

### GetVswitchTepType

`func (o *Ipv4addressDiscoveredData) GetVswitchTepType() string`

GetVswitchTepType returns the VswitchTepType field if non-nil, zero value otherwise.

### GetVswitchTepTypeOk

`func (o *Ipv4addressDiscoveredData) GetVswitchTepTypeOk() (*string, bool)`

GetVswitchTepTypeOk returns a tuple with the VswitchTepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepType

`func (o *Ipv4addressDiscoveredData) SetVswitchTepType(v string)`

SetVswitchTepType sets VswitchTepType field to given value.

### HasVswitchTepType

`func (o *Ipv4addressDiscoveredData) HasVswitchTepType() bool`

HasVswitchTepType returns a boolean if a field has been set.

### GetVswitchTepIp

`func (o *Ipv4addressDiscoveredData) GetVswitchTepIp() string`

GetVswitchTepIp returns the VswitchTepIp field if non-nil, zero value otherwise.

### GetVswitchTepIpOk

`func (o *Ipv4addressDiscoveredData) GetVswitchTepIpOk() (*string, bool)`

GetVswitchTepIpOk returns a tuple with the VswitchTepIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepIp

`func (o *Ipv4addressDiscoveredData) SetVswitchTepIp(v string)`

SetVswitchTepIp sets VswitchTepIp field to given value.

### HasVswitchTepIp

`func (o *Ipv4addressDiscoveredData) HasVswitchTepIp() bool`

HasVswitchTepIp returns a boolean if a field has been set.

### GetVswitchTepPortGroup

`func (o *Ipv4addressDiscoveredData) GetVswitchTepPortGroup() string`

GetVswitchTepPortGroup returns the VswitchTepPortGroup field if non-nil, zero value otherwise.

### GetVswitchTepPortGroupOk

`func (o *Ipv4addressDiscoveredData) GetVswitchTepPortGroupOk() (*string, bool)`

GetVswitchTepPortGroupOk returns a tuple with the VswitchTepPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepPortGroup

`func (o *Ipv4addressDiscoveredData) SetVswitchTepPortGroup(v string)`

SetVswitchTepPortGroup sets VswitchTepPortGroup field to given value.

### HasVswitchTepPortGroup

`func (o *Ipv4addressDiscoveredData) HasVswitchTepPortGroup() bool`

HasVswitchTepPortGroup returns a boolean if a field has been set.

### GetVswitchTepVlan

`func (o *Ipv4addressDiscoveredData) GetVswitchTepVlan() string`

GetVswitchTepVlan returns the VswitchTepVlan field if non-nil, zero value otherwise.

### GetVswitchTepVlanOk

`func (o *Ipv4addressDiscoveredData) GetVswitchTepVlanOk() (*string, bool)`

GetVswitchTepVlanOk returns a tuple with the VswitchTepVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepVlan

`func (o *Ipv4addressDiscoveredData) SetVswitchTepVlan(v string)`

SetVswitchTepVlan sets VswitchTepVlan field to given value.

### HasVswitchTepVlan

`func (o *Ipv4addressDiscoveredData) HasVswitchTepVlan() bool`

HasVswitchTepVlan returns a boolean if a field has been set.

### GetVswitchTepDhcpServer

`func (o *Ipv4addressDiscoveredData) GetVswitchTepDhcpServer() string`

GetVswitchTepDhcpServer returns the VswitchTepDhcpServer field if non-nil, zero value otherwise.

### GetVswitchTepDhcpServerOk

`func (o *Ipv4addressDiscoveredData) GetVswitchTepDhcpServerOk() (*string, bool)`

GetVswitchTepDhcpServerOk returns a tuple with the VswitchTepDhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepDhcpServer

`func (o *Ipv4addressDiscoveredData) SetVswitchTepDhcpServer(v string)`

SetVswitchTepDhcpServer sets VswitchTepDhcpServer field to given value.

### HasVswitchTepDhcpServer

`func (o *Ipv4addressDiscoveredData) HasVswitchTepDhcpServer() bool`

HasVswitchTepDhcpServer returns a boolean if a field has been set.

### GetVswitchTepMulticast

`func (o *Ipv4addressDiscoveredData) GetVswitchTepMulticast() string`

GetVswitchTepMulticast returns the VswitchTepMulticast field if non-nil, zero value otherwise.

### GetVswitchTepMulticastOk

`func (o *Ipv4addressDiscoveredData) GetVswitchTepMulticastOk() (*string, bool)`

GetVswitchTepMulticastOk returns a tuple with the VswitchTepMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepMulticast

`func (o *Ipv4addressDiscoveredData) SetVswitchTepMulticast(v string)`

SetVswitchTepMulticast sets VswitchTepMulticast field to given value.

### HasVswitchTepMulticast

`func (o *Ipv4addressDiscoveredData) HasVswitchTepMulticast() bool`

HasVswitchTepMulticast returns a boolean if a field has been set.

### GetVmhostIpAddress

`func (o *Ipv4addressDiscoveredData) GetVmhostIpAddress() string`

GetVmhostIpAddress returns the VmhostIpAddress field if non-nil, zero value otherwise.

### GetVmhostIpAddressOk

`func (o *Ipv4addressDiscoveredData) GetVmhostIpAddressOk() (*string, bool)`

GetVmhostIpAddressOk returns a tuple with the VmhostIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostIpAddress

`func (o *Ipv4addressDiscoveredData) SetVmhostIpAddress(v string)`

SetVmhostIpAddress sets VmhostIpAddress field to given value.

### HasVmhostIpAddress

`func (o *Ipv4addressDiscoveredData) HasVmhostIpAddress() bool`

HasVmhostIpAddress returns a boolean if a field has been set.

### GetVmhostName

`func (o *Ipv4addressDiscoveredData) GetVmhostName() string`

GetVmhostName returns the VmhostName field if non-nil, zero value otherwise.

### GetVmhostNameOk

`func (o *Ipv4addressDiscoveredData) GetVmhostNameOk() (*string, bool)`

GetVmhostNameOk returns a tuple with the VmhostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostName

`func (o *Ipv4addressDiscoveredData) SetVmhostName(v string)`

SetVmhostName sets VmhostName field to given value.

### HasVmhostName

`func (o *Ipv4addressDiscoveredData) HasVmhostName() bool`

HasVmhostName returns a boolean if a field has been set.

### GetVmhostMacAddress

`func (o *Ipv4addressDiscoveredData) GetVmhostMacAddress() string`

GetVmhostMacAddress returns the VmhostMacAddress field if non-nil, zero value otherwise.

### GetVmhostMacAddressOk

`func (o *Ipv4addressDiscoveredData) GetVmhostMacAddressOk() (*string, bool)`

GetVmhostMacAddressOk returns a tuple with the VmhostMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostMacAddress

`func (o *Ipv4addressDiscoveredData) SetVmhostMacAddress(v string)`

SetVmhostMacAddress sets VmhostMacAddress field to given value.

### HasVmhostMacAddress

`func (o *Ipv4addressDiscoveredData) HasVmhostMacAddress() bool`

HasVmhostMacAddress returns a boolean if a field has been set.

### GetVmhostSubnetCidr

`func (o *Ipv4addressDiscoveredData) GetVmhostSubnetCidr() int64`

GetVmhostSubnetCidr returns the VmhostSubnetCidr field if non-nil, zero value otherwise.

### GetVmhostSubnetCidrOk

`func (o *Ipv4addressDiscoveredData) GetVmhostSubnetCidrOk() (*int64, bool)`

GetVmhostSubnetCidrOk returns a tuple with the VmhostSubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostSubnetCidr

`func (o *Ipv4addressDiscoveredData) SetVmhostSubnetCidr(v int64)`

SetVmhostSubnetCidr sets VmhostSubnetCidr field to given value.

### HasVmhostSubnetCidr

`func (o *Ipv4addressDiscoveredData) HasVmhostSubnetCidr() bool`

HasVmhostSubnetCidr returns a boolean if a field has been set.

### GetVmhostNicNames

`func (o *Ipv4addressDiscoveredData) GetVmhostNicNames() string`

GetVmhostNicNames returns the VmhostNicNames field if non-nil, zero value otherwise.

### GetVmhostNicNamesOk

`func (o *Ipv4addressDiscoveredData) GetVmhostNicNamesOk() (*string, bool)`

GetVmhostNicNamesOk returns a tuple with the VmhostNicNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostNicNames

`func (o *Ipv4addressDiscoveredData) SetVmhostNicNames(v string)`

SetVmhostNicNames sets VmhostNicNames field to given value.

### HasVmhostNicNames

`func (o *Ipv4addressDiscoveredData) HasVmhostNicNames() bool`

HasVmhostNicNames returns a boolean if a field has been set.

### GetVmiTenantId

`func (o *Ipv4addressDiscoveredData) GetVmiTenantId() string`

GetVmiTenantId returns the VmiTenantId field if non-nil, zero value otherwise.

### GetVmiTenantIdOk

`func (o *Ipv4addressDiscoveredData) GetVmiTenantIdOk() (*string, bool)`

GetVmiTenantIdOk returns a tuple with the VmiTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiTenantId

`func (o *Ipv4addressDiscoveredData) SetVmiTenantId(v string)`

SetVmiTenantId sets VmiTenantId field to given value.

### HasVmiTenantId

`func (o *Ipv4addressDiscoveredData) HasVmiTenantId() bool`

HasVmiTenantId returns a boolean if a field has been set.

### GetCmpType

`func (o *Ipv4addressDiscoveredData) GetCmpType() string`

GetCmpType returns the CmpType field if non-nil, zero value otherwise.

### GetCmpTypeOk

`func (o *Ipv4addressDiscoveredData) GetCmpTypeOk() (*string, bool)`

GetCmpTypeOk returns a tuple with the CmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmpType

`func (o *Ipv4addressDiscoveredData) SetCmpType(v string)`

SetCmpType sets CmpType field to given value.

### HasCmpType

`func (o *Ipv4addressDiscoveredData) HasCmpType() bool`

HasCmpType returns a boolean if a field has been set.

### GetVmiIpType

`func (o *Ipv4addressDiscoveredData) GetVmiIpType() string`

GetVmiIpType returns the VmiIpType field if non-nil, zero value otherwise.

### GetVmiIpTypeOk

`func (o *Ipv4addressDiscoveredData) GetVmiIpTypeOk() (*string, bool)`

GetVmiIpTypeOk returns a tuple with the VmiIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiIpType

`func (o *Ipv4addressDiscoveredData) SetVmiIpType(v string)`

SetVmiIpType sets VmiIpType field to given value.

### HasVmiIpType

`func (o *Ipv4addressDiscoveredData) HasVmiIpType() bool`

HasVmiIpType returns a boolean if a field has been set.

### GetVmiPrivateAddress

`func (o *Ipv4addressDiscoveredData) GetVmiPrivateAddress() string`

GetVmiPrivateAddress returns the VmiPrivateAddress field if non-nil, zero value otherwise.

### GetVmiPrivateAddressOk

`func (o *Ipv4addressDiscoveredData) GetVmiPrivateAddressOk() (*string, bool)`

GetVmiPrivateAddressOk returns a tuple with the VmiPrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiPrivateAddress

`func (o *Ipv4addressDiscoveredData) SetVmiPrivateAddress(v string)`

SetVmiPrivateAddress sets VmiPrivateAddress field to given value.

### HasVmiPrivateAddress

`func (o *Ipv4addressDiscoveredData) HasVmiPrivateAddress() bool`

HasVmiPrivateAddress returns a boolean if a field has been set.

### GetVmiIsPublicAddress

`func (o *Ipv4addressDiscoveredData) GetVmiIsPublicAddress() bool`

GetVmiIsPublicAddress returns the VmiIsPublicAddress field if non-nil, zero value otherwise.

### GetVmiIsPublicAddressOk

`func (o *Ipv4addressDiscoveredData) GetVmiIsPublicAddressOk() (*bool, bool)`

GetVmiIsPublicAddressOk returns a tuple with the VmiIsPublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiIsPublicAddress

`func (o *Ipv4addressDiscoveredData) SetVmiIsPublicAddress(v bool)`

SetVmiIsPublicAddress sets VmiIsPublicAddress field to given value.

### HasVmiIsPublicAddress

`func (o *Ipv4addressDiscoveredData) HasVmiIsPublicAddress() bool`

HasVmiIsPublicAddress returns a boolean if a field has been set.

### GetCiscoIseSsid

`func (o *Ipv4addressDiscoveredData) GetCiscoIseSsid() string`

GetCiscoIseSsid returns the CiscoIseSsid field if non-nil, zero value otherwise.

### GetCiscoIseSsidOk

`func (o *Ipv4addressDiscoveredData) GetCiscoIseSsidOk() (*string, bool)`

GetCiscoIseSsidOk returns a tuple with the CiscoIseSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseSsid

`func (o *Ipv4addressDiscoveredData) SetCiscoIseSsid(v string)`

SetCiscoIseSsid sets CiscoIseSsid field to given value.

### HasCiscoIseSsid

`func (o *Ipv4addressDiscoveredData) HasCiscoIseSsid() bool`

HasCiscoIseSsid returns a boolean if a field has been set.

### GetCiscoIseEndpointProfile

`func (o *Ipv4addressDiscoveredData) GetCiscoIseEndpointProfile() string`

GetCiscoIseEndpointProfile returns the CiscoIseEndpointProfile field if non-nil, zero value otherwise.

### GetCiscoIseEndpointProfileOk

`func (o *Ipv4addressDiscoveredData) GetCiscoIseEndpointProfileOk() (*string, bool)`

GetCiscoIseEndpointProfileOk returns a tuple with the CiscoIseEndpointProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseEndpointProfile

`func (o *Ipv4addressDiscoveredData) SetCiscoIseEndpointProfile(v string)`

SetCiscoIseEndpointProfile sets CiscoIseEndpointProfile field to given value.

### HasCiscoIseEndpointProfile

`func (o *Ipv4addressDiscoveredData) HasCiscoIseEndpointProfile() bool`

HasCiscoIseEndpointProfile returns a boolean if a field has been set.

### GetCiscoIseSessionState

`func (o *Ipv4addressDiscoveredData) GetCiscoIseSessionState() string`

GetCiscoIseSessionState returns the CiscoIseSessionState field if non-nil, zero value otherwise.

### GetCiscoIseSessionStateOk

`func (o *Ipv4addressDiscoveredData) GetCiscoIseSessionStateOk() (*string, bool)`

GetCiscoIseSessionStateOk returns a tuple with the CiscoIseSessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseSessionState

`func (o *Ipv4addressDiscoveredData) SetCiscoIseSessionState(v string)`

SetCiscoIseSessionState sets CiscoIseSessionState field to given value.

### HasCiscoIseSessionState

`func (o *Ipv4addressDiscoveredData) HasCiscoIseSessionState() bool`

HasCiscoIseSessionState returns a boolean if a field has been set.

### GetCiscoIseSecurityGroup

`func (o *Ipv4addressDiscoveredData) GetCiscoIseSecurityGroup() string`

GetCiscoIseSecurityGroup returns the CiscoIseSecurityGroup field if non-nil, zero value otherwise.

### GetCiscoIseSecurityGroupOk

`func (o *Ipv4addressDiscoveredData) GetCiscoIseSecurityGroupOk() (*string, bool)`

GetCiscoIseSecurityGroupOk returns a tuple with the CiscoIseSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseSecurityGroup

`func (o *Ipv4addressDiscoveredData) SetCiscoIseSecurityGroup(v string)`

SetCiscoIseSecurityGroup sets CiscoIseSecurityGroup field to given value.

### HasCiscoIseSecurityGroup

`func (o *Ipv4addressDiscoveredData) HasCiscoIseSecurityGroup() bool`

HasCiscoIseSecurityGroup returns a boolean if a field has been set.

### GetTaskName

`func (o *Ipv4addressDiscoveredData) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *Ipv4addressDiscoveredData) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *Ipv4addressDiscoveredData) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *Ipv4addressDiscoveredData) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetNetworkComponentLocation

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentLocation() string`

GetNetworkComponentLocation returns the NetworkComponentLocation field if non-nil, zero value otherwise.

### GetNetworkComponentLocationOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentLocationOk() (*string, bool)`

GetNetworkComponentLocationOk returns a tuple with the NetworkComponentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentLocation

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentLocation(v string)`

SetNetworkComponentLocation sets NetworkComponentLocation field to given value.

### HasNetworkComponentLocation

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentLocation() bool`

HasNetworkComponentLocation returns a boolean if a field has been set.

### GetNetworkComponentContact

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentContact() string`

GetNetworkComponentContact returns the NetworkComponentContact field if non-nil, zero value otherwise.

### GetNetworkComponentContactOk

`func (o *Ipv4addressDiscoveredData) GetNetworkComponentContactOk() (*string, bool)`

GetNetworkComponentContactOk returns a tuple with the NetworkComponentContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentContact

`func (o *Ipv4addressDiscoveredData) SetNetworkComponentContact(v string)`

SetNetworkComponentContact sets NetworkComponentContact field to given value.

### HasNetworkComponentContact

`func (o *Ipv4addressDiscoveredData) HasNetworkComponentContact() bool`

HasNetworkComponentContact returns a boolean if a field has been set.

### GetDeviceLocation

`func (o *Ipv4addressDiscoveredData) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *Ipv4addressDiscoveredData) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *Ipv4addressDiscoveredData) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.

### HasDeviceLocation

`func (o *Ipv4addressDiscoveredData) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### GetDeviceContact

`func (o *Ipv4addressDiscoveredData) GetDeviceContact() string`

GetDeviceContact returns the DeviceContact field if non-nil, zero value otherwise.

### GetDeviceContactOk

`func (o *Ipv4addressDiscoveredData) GetDeviceContactOk() (*string, bool)`

GetDeviceContactOk returns a tuple with the DeviceContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceContact

`func (o *Ipv4addressDiscoveredData) SetDeviceContact(v string)`

SetDeviceContact sets DeviceContact field to given value.

### HasDeviceContact

`func (o *Ipv4addressDiscoveredData) HasDeviceContact() bool`

HasDeviceContact returns a boolean if a field has been set.

### GetApName

`func (o *Ipv4addressDiscoveredData) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *Ipv4addressDiscoveredData) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *Ipv4addressDiscoveredData) SetApName(v string)`

SetApName sets ApName field to given value.

### HasApName

`func (o *Ipv4addressDiscoveredData) HasApName() bool`

HasApName returns a boolean if a field has been set.

### GetApIpAddress

`func (o *Ipv4addressDiscoveredData) GetApIpAddress() string`

GetApIpAddress returns the ApIpAddress field if non-nil, zero value otherwise.

### GetApIpAddressOk

`func (o *Ipv4addressDiscoveredData) GetApIpAddressOk() (*string, bool)`

GetApIpAddressOk returns a tuple with the ApIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApIpAddress

`func (o *Ipv4addressDiscoveredData) SetApIpAddress(v string)`

SetApIpAddress sets ApIpAddress field to given value.

### HasApIpAddress

`func (o *Ipv4addressDiscoveredData) HasApIpAddress() bool`

HasApIpAddress returns a boolean if a field has been set.

### GetApSsid

`func (o *Ipv4addressDiscoveredData) GetApSsid() string`

GetApSsid returns the ApSsid field if non-nil, zero value otherwise.

### GetApSsidOk

`func (o *Ipv4addressDiscoveredData) GetApSsidOk() (*string, bool)`

GetApSsidOk returns a tuple with the ApSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApSsid

`func (o *Ipv4addressDiscoveredData) SetApSsid(v string)`

SetApSsid sets ApSsid field to given value.

### HasApSsid

`func (o *Ipv4addressDiscoveredData) HasApSsid() bool`

HasApSsid returns a boolean if a field has been set.

### GetBridgeDomain

`func (o *Ipv4addressDiscoveredData) GetBridgeDomain() string`

GetBridgeDomain returns the BridgeDomain field if non-nil, zero value otherwise.

### GetBridgeDomainOk

`func (o *Ipv4addressDiscoveredData) GetBridgeDomainOk() (*string, bool)`

GetBridgeDomainOk returns a tuple with the BridgeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeDomain

`func (o *Ipv4addressDiscoveredData) SetBridgeDomain(v string)`

SetBridgeDomain sets BridgeDomain field to given value.

### HasBridgeDomain

`func (o *Ipv4addressDiscoveredData) HasBridgeDomain() bool`

HasBridgeDomain returns a boolean if a field has been set.

### GetEndpointGroups

`func (o *Ipv4addressDiscoveredData) GetEndpointGroups() string`

GetEndpointGroups returns the EndpointGroups field if non-nil, zero value otherwise.

### GetEndpointGroupsOk

`func (o *Ipv4addressDiscoveredData) GetEndpointGroupsOk() (*string, bool)`

GetEndpointGroupsOk returns a tuple with the EndpointGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroups

`func (o *Ipv4addressDiscoveredData) SetEndpointGroups(v string)`

SetEndpointGroups sets EndpointGroups field to given value.

### HasEndpointGroups

`func (o *Ipv4addressDiscoveredData) HasEndpointGroups() bool`

HasEndpointGroups returns a boolean if a field has been set.

### GetTenant

`func (o *Ipv4addressDiscoveredData) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Ipv4addressDiscoveredData) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Ipv4addressDiscoveredData) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Ipv4addressDiscoveredData) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetVrfName

`func (o *Ipv4addressDiscoveredData) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *Ipv4addressDiscoveredData) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *Ipv4addressDiscoveredData) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *Ipv4addressDiscoveredData) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.

### GetVrfDescription

`func (o *Ipv4addressDiscoveredData) GetVrfDescription() string`

GetVrfDescription returns the VrfDescription field if non-nil, zero value otherwise.

### GetVrfDescriptionOk

`func (o *Ipv4addressDiscoveredData) GetVrfDescriptionOk() (*string, bool)`

GetVrfDescriptionOk returns a tuple with the VrfDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfDescription

`func (o *Ipv4addressDiscoveredData) SetVrfDescription(v string)`

SetVrfDescription sets VrfDescription field to given value.

### HasVrfDescription

`func (o *Ipv4addressDiscoveredData) HasVrfDescription() bool`

HasVrfDescription returns a boolean if a field has been set.

### GetVrfRd

`func (o *Ipv4addressDiscoveredData) GetVrfRd() string`

GetVrfRd returns the VrfRd field if non-nil, zero value otherwise.

### GetVrfRdOk

`func (o *Ipv4addressDiscoveredData) GetVrfRdOk() (*string, bool)`

GetVrfRdOk returns a tuple with the VrfRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfRd

`func (o *Ipv4addressDiscoveredData) SetVrfRd(v string)`

SetVrfRd sets VrfRd field to given value.

### HasVrfRd

`func (o *Ipv4addressDiscoveredData) HasVrfRd() bool`

HasVrfRd returns a boolean if a field has been set.

### GetBgpAs

`func (o *Ipv4addressDiscoveredData) GetBgpAs() int64`

GetBgpAs returns the BgpAs field if non-nil, zero value otherwise.

### GetBgpAsOk

`func (o *Ipv4addressDiscoveredData) GetBgpAsOk() (*int64, bool)`

GetBgpAsOk returns a tuple with the BgpAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAs

`func (o *Ipv4addressDiscoveredData) SetBgpAs(v int64)`

SetBgpAs sets BgpAs field to given value.

### HasBgpAs

`func (o *Ipv4addressDiscoveredData) HasBgpAs() bool`

HasBgpAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


