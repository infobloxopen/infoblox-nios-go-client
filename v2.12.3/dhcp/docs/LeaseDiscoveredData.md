# LeaseDiscoveredData

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

### NewLeaseDiscoveredData

`func NewLeaseDiscoveredData() *LeaseDiscoveredData`

NewLeaseDiscoveredData instantiates a new LeaseDiscoveredData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaseDiscoveredDataWithDefaults

`func NewLeaseDiscoveredDataWithDefaults() *LeaseDiscoveredData`

NewLeaseDiscoveredDataWithDefaults instantiates a new LeaseDiscoveredData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *LeaseDiscoveredData) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *LeaseDiscoveredData) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *LeaseDiscoveredData) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *LeaseDiscoveredData) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDevicePortName

`func (o *LeaseDiscoveredData) GetDevicePortName() string`

GetDevicePortName returns the DevicePortName field if non-nil, zero value otherwise.

### GetDevicePortNameOk

`func (o *LeaseDiscoveredData) GetDevicePortNameOk() (*string, bool)`

GetDevicePortNameOk returns a tuple with the DevicePortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortName

`func (o *LeaseDiscoveredData) SetDevicePortName(v string)`

SetDevicePortName sets DevicePortName field to given value.

### HasDevicePortName

`func (o *LeaseDiscoveredData) HasDevicePortName() bool`

HasDevicePortName returns a boolean if a field has been set.

### GetDevicePortType

`func (o *LeaseDiscoveredData) GetDevicePortType() string`

GetDevicePortType returns the DevicePortType field if non-nil, zero value otherwise.

### GetDevicePortTypeOk

`func (o *LeaseDiscoveredData) GetDevicePortTypeOk() (*string, bool)`

GetDevicePortTypeOk returns a tuple with the DevicePortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortType

`func (o *LeaseDiscoveredData) SetDevicePortType(v string)`

SetDevicePortType sets DevicePortType field to given value.

### HasDevicePortType

`func (o *LeaseDiscoveredData) HasDevicePortType() bool`

HasDevicePortType returns a boolean if a field has been set.

### GetDeviceType

`func (o *LeaseDiscoveredData) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *LeaseDiscoveredData) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *LeaseDiscoveredData) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *LeaseDiscoveredData) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDeviceVendor

`func (o *LeaseDiscoveredData) GetDeviceVendor() string`

GetDeviceVendor returns the DeviceVendor field if non-nil, zero value otherwise.

### GetDeviceVendorOk

`func (o *LeaseDiscoveredData) GetDeviceVendorOk() (*string, bool)`

GetDeviceVendorOk returns a tuple with the DeviceVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceVendor

`func (o *LeaseDiscoveredData) SetDeviceVendor(v string)`

SetDeviceVendor sets DeviceVendor field to given value.

### HasDeviceVendor

`func (o *LeaseDiscoveredData) HasDeviceVendor() bool`

HasDeviceVendor returns a boolean if a field has been set.

### GetDiscoveredName

`func (o *LeaseDiscoveredData) GetDiscoveredName() string`

GetDiscoveredName returns the DiscoveredName field if non-nil, zero value otherwise.

### GetDiscoveredNameOk

`func (o *LeaseDiscoveredData) GetDiscoveredNameOk() (*string, bool)`

GetDiscoveredNameOk returns a tuple with the DiscoveredName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredName

`func (o *LeaseDiscoveredData) SetDiscoveredName(v string)`

SetDiscoveredName sets DiscoveredName field to given value.

### HasDiscoveredName

`func (o *LeaseDiscoveredData) HasDiscoveredName() bool`

HasDiscoveredName returns a boolean if a field has been set.

### GetDiscoverer

`func (o *LeaseDiscoveredData) GetDiscoverer() string`

GetDiscoverer returns the Discoverer field if non-nil, zero value otherwise.

### GetDiscovererOk

`func (o *LeaseDiscoveredData) GetDiscovererOk() (*string, bool)`

GetDiscovererOk returns a tuple with the Discoverer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoverer

`func (o *LeaseDiscoveredData) SetDiscoverer(v string)`

SetDiscoverer sets Discoverer field to given value.

### HasDiscoverer

`func (o *LeaseDiscoveredData) HasDiscoverer() bool`

HasDiscoverer returns a boolean if a field has been set.

### GetDuid

`func (o *LeaseDiscoveredData) GetDuid() string`

GetDuid returns the Duid field if non-nil, zero value otherwise.

### GetDuidOk

`func (o *LeaseDiscoveredData) GetDuidOk() (*string, bool)`

GetDuidOk returns a tuple with the Duid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuid

`func (o *LeaseDiscoveredData) SetDuid(v string)`

SetDuid sets Duid field to given value.

### HasDuid

`func (o *LeaseDiscoveredData) HasDuid() bool`

HasDuid returns a boolean if a field has been set.

### GetFirstDiscovered

`func (o *LeaseDiscoveredData) GetFirstDiscovered() int64`

GetFirstDiscovered returns the FirstDiscovered field if non-nil, zero value otherwise.

### GetFirstDiscoveredOk

`func (o *LeaseDiscoveredData) GetFirstDiscoveredOk() (*int64, bool)`

GetFirstDiscoveredOk returns a tuple with the FirstDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDiscovered

`func (o *LeaseDiscoveredData) SetFirstDiscovered(v int64)`

SetFirstDiscovered sets FirstDiscovered field to given value.

### HasFirstDiscovered

`func (o *LeaseDiscoveredData) HasFirstDiscovered() bool`

HasFirstDiscovered returns a boolean if a field has been set.

### GetIprgNo

`func (o *LeaseDiscoveredData) GetIprgNo() int64`

GetIprgNo returns the IprgNo field if non-nil, zero value otherwise.

### GetIprgNoOk

`func (o *LeaseDiscoveredData) GetIprgNoOk() (*int64, bool)`

GetIprgNoOk returns a tuple with the IprgNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprgNo

`func (o *LeaseDiscoveredData) SetIprgNo(v int64)`

SetIprgNo sets IprgNo field to given value.

### HasIprgNo

`func (o *LeaseDiscoveredData) HasIprgNo() bool`

HasIprgNo returns a boolean if a field has been set.

### GetIprgState

`func (o *LeaseDiscoveredData) GetIprgState() string`

GetIprgState returns the IprgState field if non-nil, zero value otherwise.

### GetIprgStateOk

`func (o *LeaseDiscoveredData) GetIprgStateOk() (*string, bool)`

GetIprgStateOk returns a tuple with the IprgState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprgState

`func (o *LeaseDiscoveredData) SetIprgState(v string)`

SetIprgState sets IprgState field to given value.

### HasIprgState

`func (o *LeaseDiscoveredData) HasIprgState() bool`

HasIprgState returns a boolean if a field has been set.

### GetIprgType

`func (o *LeaseDiscoveredData) GetIprgType() string`

GetIprgType returns the IprgType field if non-nil, zero value otherwise.

### GetIprgTypeOk

`func (o *LeaseDiscoveredData) GetIprgTypeOk() (*string, bool)`

GetIprgTypeOk returns a tuple with the IprgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIprgType

`func (o *LeaseDiscoveredData) SetIprgType(v string)`

SetIprgType sets IprgType field to given value.

### HasIprgType

`func (o *LeaseDiscoveredData) HasIprgType() bool`

HasIprgType returns a boolean if a field has been set.

### GetLastDiscovered

`func (o *LeaseDiscoveredData) GetLastDiscovered() int64`

GetLastDiscovered returns the LastDiscovered field if non-nil, zero value otherwise.

### GetLastDiscoveredOk

`func (o *LeaseDiscoveredData) GetLastDiscoveredOk() (*int64, bool)`

GetLastDiscoveredOk returns a tuple with the LastDiscovered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDiscovered

`func (o *LeaseDiscoveredData) SetLastDiscovered(v int64)`

SetLastDiscovered sets LastDiscovered field to given value.

### HasLastDiscovered

`func (o *LeaseDiscoveredData) HasLastDiscovered() bool`

HasLastDiscovered returns a boolean if a field has been set.

### GetMacAddress

`func (o *LeaseDiscoveredData) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *LeaseDiscoveredData) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *LeaseDiscoveredData) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *LeaseDiscoveredData) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMgmtIpAddress

`func (o *LeaseDiscoveredData) GetMgmtIpAddress() string`

GetMgmtIpAddress returns the MgmtIpAddress field if non-nil, zero value otherwise.

### GetMgmtIpAddressOk

`func (o *LeaseDiscoveredData) GetMgmtIpAddressOk() (*string, bool)`

GetMgmtIpAddressOk returns a tuple with the MgmtIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtIpAddress

`func (o *LeaseDiscoveredData) SetMgmtIpAddress(v string)`

SetMgmtIpAddress sets MgmtIpAddress field to given value.

### HasMgmtIpAddress

`func (o *LeaseDiscoveredData) HasMgmtIpAddress() bool`

HasMgmtIpAddress returns a boolean if a field has been set.

### GetNetbiosName

`func (o *LeaseDiscoveredData) GetNetbiosName() string`

GetNetbiosName returns the NetbiosName field if non-nil, zero value otherwise.

### GetNetbiosNameOk

`func (o *LeaseDiscoveredData) GetNetbiosNameOk() (*string, bool)`

GetNetbiosNameOk returns a tuple with the NetbiosName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetbiosName

`func (o *LeaseDiscoveredData) SetNetbiosName(v string)`

SetNetbiosName sets NetbiosName field to given value.

### HasNetbiosName

`func (o *LeaseDiscoveredData) HasNetbiosName() bool`

HasNetbiosName returns a boolean if a field has been set.

### GetNetworkComponentDescription

`func (o *LeaseDiscoveredData) GetNetworkComponentDescription() string`

GetNetworkComponentDescription returns the NetworkComponentDescription field if non-nil, zero value otherwise.

### GetNetworkComponentDescriptionOk

`func (o *LeaseDiscoveredData) GetNetworkComponentDescriptionOk() (*string, bool)`

GetNetworkComponentDescriptionOk returns a tuple with the NetworkComponentDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentDescription

`func (o *LeaseDiscoveredData) SetNetworkComponentDescription(v string)`

SetNetworkComponentDescription sets NetworkComponentDescription field to given value.

### HasNetworkComponentDescription

`func (o *LeaseDiscoveredData) HasNetworkComponentDescription() bool`

HasNetworkComponentDescription returns a boolean if a field has been set.

### GetNetworkComponentIp

`func (o *LeaseDiscoveredData) GetNetworkComponentIp() string`

GetNetworkComponentIp returns the NetworkComponentIp field if non-nil, zero value otherwise.

### GetNetworkComponentIpOk

`func (o *LeaseDiscoveredData) GetNetworkComponentIpOk() (*string, bool)`

GetNetworkComponentIpOk returns a tuple with the NetworkComponentIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentIp

`func (o *LeaseDiscoveredData) SetNetworkComponentIp(v string)`

SetNetworkComponentIp sets NetworkComponentIp field to given value.

### HasNetworkComponentIp

`func (o *LeaseDiscoveredData) HasNetworkComponentIp() bool`

HasNetworkComponentIp returns a boolean if a field has been set.

### GetNetworkComponentModel

`func (o *LeaseDiscoveredData) GetNetworkComponentModel() string`

GetNetworkComponentModel returns the NetworkComponentModel field if non-nil, zero value otherwise.

### GetNetworkComponentModelOk

`func (o *LeaseDiscoveredData) GetNetworkComponentModelOk() (*string, bool)`

GetNetworkComponentModelOk returns a tuple with the NetworkComponentModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentModel

`func (o *LeaseDiscoveredData) SetNetworkComponentModel(v string)`

SetNetworkComponentModel sets NetworkComponentModel field to given value.

### HasNetworkComponentModel

`func (o *LeaseDiscoveredData) HasNetworkComponentModel() bool`

HasNetworkComponentModel returns a boolean if a field has been set.

### GetNetworkComponentName

`func (o *LeaseDiscoveredData) GetNetworkComponentName() string`

GetNetworkComponentName returns the NetworkComponentName field if non-nil, zero value otherwise.

### GetNetworkComponentNameOk

`func (o *LeaseDiscoveredData) GetNetworkComponentNameOk() (*string, bool)`

GetNetworkComponentNameOk returns a tuple with the NetworkComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentName

`func (o *LeaseDiscoveredData) SetNetworkComponentName(v string)`

SetNetworkComponentName sets NetworkComponentName field to given value.

### HasNetworkComponentName

`func (o *LeaseDiscoveredData) HasNetworkComponentName() bool`

HasNetworkComponentName returns a boolean if a field has been set.

### GetNetworkComponentPortDescription

`func (o *LeaseDiscoveredData) GetNetworkComponentPortDescription() string`

GetNetworkComponentPortDescription returns the NetworkComponentPortDescription field if non-nil, zero value otherwise.

### GetNetworkComponentPortDescriptionOk

`func (o *LeaseDiscoveredData) GetNetworkComponentPortDescriptionOk() (*string, bool)`

GetNetworkComponentPortDescriptionOk returns a tuple with the NetworkComponentPortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentPortDescription

`func (o *LeaseDiscoveredData) SetNetworkComponentPortDescription(v string)`

SetNetworkComponentPortDescription sets NetworkComponentPortDescription field to given value.

### HasNetworkComponentPortDescription

`func (o *LeaseDiscoveredData) HasNetworkComponentPortDescription() bool`

HasNetworkComponentPortDescription returns a boolean if a field has been set.

### GetNetworkComponentPortName

`func (o *LeaseDiscoveredData) GetNetworkComponentPortName() string`

GetNetworkComponentPortName returns the NetworkComponentPortName field if non-nil, zero value otherwise.

### GetNetworkComponentPortNameOk

`func (o *LeaseDiscoveredData) GetNetworkComponentPortNameOk() (*string, bool)`

GetNetworkComponentPortNameOk returns a tuple with the NetworkComponentPortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentPortName

`func (o *LeaseDiscoveredData) SetNetworkComponentPortName(v string)`

SetNetworkComponentPortName sets NetworkComponentPortName field to given value.

### HasNetworkComponentPortName

`func (o *LeaseDiscoveredData) HasNetworkComponentPortName() bool`

HasNetworkComponentPortName returns a boolean if a field has been set.

### GetNetworkComponentPortNumber

`func (o *LeaseDiscoveredData) GetNetworkComponentPortNumber() string`

GetNetworkComponentPortNumber returns the NetworkComponentPortNumber field if non-nil, zero value otherwise.

### GetNetworkComponentPortNumberOk

`func (o *LeaseDiscoveredData) GetNetworkComponentPortNumberOk() (*string, bool)`

GetNetworkComponentPortNumberOk returns a tuple with the NetworkComponentPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentPortNumber

`func (o *LeaseDiscoveredData) SetNetworkComponentPortNumber(v string)`

SetNetworkComponentPortNumber sets NetworkComponentPortNumber field to given value.

### HasNetworkComponentPortNumber

`func (o *LeaseDiscoveredData) HasNetworkComponentPortNumber() bool`

HasNetworkComponentPortNumber returns a boolean if a field has been set.

### GetNetworkComponentType

`func (o *LeaseDiscoveredData) GetNetworkComponentType() string`

GetNetworkComponentType returns the NetworkComponentType field if non-nil, zero value otherwise.

### GetNetworkComponentTypeOk

`func (o *LeaseDiscoveredData) GetNetworkComponentTypeOk() (*string, bool)`

GetNetworkComponentTypeOk returns a tuple with the NetworkComponentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentType

`func (o *LeaseDiscoveredData) SetNetworkComponentType(v string)`

SetNetworkComponentType sets NetworkComponentType field to given value.

### HasNetworkComponentType

`func (o *LeaseDiscoveredData) HasNetworkComponentType() bool`

HasNetworkComponentType returns a boolean if a field has been set.

### GetNetworkComponentVendor

`func (o *LeaseDiscoveredData) GetNetworkComponentVendor() string`

GetNetworkComponentVendor returns the NetworkComponentVendor field if non-nil, zero value otherwise.

### GetNetworkComponentVendorOk

`func (o *LeaseDiscoveredData) GetNetworkComponentVendorOk() (*string, bool)`

GetNetworkComponentVendorOk returns a tuple with the NetworkComponentVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentVendor

`func (o *LeaseDiscoveredData) SetNetworkComponentVendor(v string)`

SetNetworkComponentVendor sets NetworkComponentVendor field to given value.

### HasNetworkComponentVendor

`func (o *LeaseDiscoveredData) HasNetworkComponentVendor() bool`

HasNetworkComponentVendor returns a boolean if a field has been set.

### GetOpenPorts

`func (o *LeaseDiscoveredData) GetOpenPorts() string`

GetOpenPorts returns the OpenPorts field if non-nil, zero value otherwise.

### GetOpenPortsOk

`func (o *LeaseDiscoveredData) GetOpenPortsOk() (*string, bool)`

GetOpenPortsOk returns a tuple with the OpenPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPorts

`func (o *LeaseDiscoveredData) SetOpenPorts(v string)`

SetOpenPorts sets OpenPorts field to given value.

### HasOpenPorts

`func (o *LeaseDiscoveredData) HasOpenPorts() bool`

HasOpenPorts returns a boolean if a field has been set.

### GetOs

`func (o *LeaseDiscoveredData) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *LeaseDiscoveredData) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *LeaseDiscoveredData) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *LeaseDiscoveredData) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPortDuplex

`func (o *LeaseDiscoveredData) GetPortDuplex() string`

GetPortDuplex returns the PortDuplex field if non-nil, zero value otherwise.

### GetPortDuplexOk

`func (o *LeaseDiscoveredData) GetPortDuplexOk() (*string, bool)`

GetPortDuplexOk returns a tuple with the PortDuplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDuplex

`func (o *LeaseDiscoveredData) SetPortDuplex(v string)`

SetPortDuplex sets PortDuplex field to given value.

### HasPortDuplex

`func (o *LeaseDiscoveredData) HasPortDuplex() bool`

HasPortDuplex returns a boolean if a field has been set.

### GetPortLinkStatus

`func (o *LeaseDiscoveredData) GetPortLinkStatus() string`

GetPortLinkStatus returns the PortLinkStatus field if non-nil, zero value otherwise.

### GetPortLinkStatusOk

`func (o *LeaseDiscoveredData) GetPortLinkStatusOk() (*string, bool)`

GetPortLinkStatusOk returns a tuple with the PortLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortLinkStatus

`func (o *LeaseDiscoveredData) SetPortLinkStatus(v string)`

SetPortLinkStatus sets PortLinkStatus field to given value.

### HasPortLinkStatus

`func (o *LeaseDiscoveredData) HasPortLinkStatus() bool`

HasPortLinkStatus returns a boolean if a field has been set.

### GetPortSpeed

`func (o *LeaseDiscoveredData) GetPortSpeed() string`

GetPortSpeed returns the PortSpeed field if non-nil, zero value otherwise.

### GetPortSpeedOk

`func (o *LeaseDiscoveredData) GetPortSpeedOk() (*string, bool)`

GetPortSpeedOk returns a tuple with the PortSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeed

`func (o *LeaseDiscoveredData) SetPortSpeed(v string)`

SetPortSpeed sets PortSpeed field to given value.

### HasPortSpeed

`func (o *LeaseDiscoveredData) HasPortSpeed() bool`

HasPortSpeed returns a boolean if a field has been set.

### GetPortStatus

`func (o *LeaseDiscoveredData) GetPortStatus() string`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *LeaseDiscoveredData) GetPortStatusOk() (*string, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *LeaseDiscoveredData) SetPortStatus(v string)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *LeaseDiscoveredData) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetPortType

`func (o *LeaseDiscoveredData) GetPortType() string`

GetPortType returns the PortType field if non-nil, zero value otherwise.

### GetPortTypeOk

`func (o *LeaseDiscoveredData) GetPortTypeOk() (*string, bool)`

GetPortTypeOk returns a tuple with the PortType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortType

`func (o *LeaseDiscoveredData) SetPortType(v string)`

SetPortType sets PortType field to given value.

### HasPortType

`func (o *LeaseDiscoveredData) HasPortType() bool`

HasPortType returns a boolean if a field has been set.

### GetPortVlanDescription

`func (o *LeaseDiscoveredData) GetPortVlanDescription() string`

GetPortVlanDescription returns the PortVlanDescription field if non-nil, zero value otherwise.

### GetPortVlanDescriptionOk

`func (o *LeaseDiscoveredData) GetPortVlanDescriptionOk() (*string, bool)`

GetPortVlanDescriptionOk returns a tuple with the PortVlanDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanDescription

`func (o *LeaseDiscoveredData) SetPortVlanDescription(v string)`

SetPortVlanDescription sets PortVlanDescription field to given value.

### HasPortVlanDescription

`func (o *LeaseDiscoveredData) HasPortVlanDescription() bool`

HasPortVlanDescription returns a boolean if a field has been set.

### GetPortVlanName

`func (o *LeaseDiscoveredData) GetPortVlanName() string`

GetPortVlanName returns the PortVlanName field if non-nil, zero value otherwise.

### GetPortVlanNameOk

`func (o *LeaseDiscoveredData) GetPortVlanNameOk() (*string, bool)`

GetPortVlanNameOk returns a tuple with the PortVlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanName

`func (o *LeaseDiscoveredData) SetPortVlanName(v string)`

SetPortVlanName sets PortVlanName field to given value.

### HasPortVlanName

`func (o *LeaseDiscoveredData) HasPortVlanName() bool`

HasPortVlanName returns a boolean if a field has been set.

### GetPortVlanNumber

`func (o *LeaseDiscoveredData) GetPortVlanNumber() string`

GetPortVlanNumber returns the PortVlanNumber field if non-nil, zero value otherwise.

### GetPortVlanNumberOk

`func (o *LeaseDiscoveredData) GetPortVlanNumberOk() (*string, bool)`

GetPortVlanNumberOk returns a tuple with the PortVlanNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortVlanNumber

`func (o *LeaseDiscoveredData) SetPortVlanNumber(v string)`

SetPortVlanNumber sets PortVlanNumber field to given value.

### HasPortVlanNumber

`func (o *LeaseDiscoveredData) HasPortVlanNumber() bool`

HasPortVlanNumber returns a boolean if a field has been set.

### GetVAdapter

`func (o *LeaseDiscoveredData) GetVAdapter() string`

GetVAdapter returns the VAdapter field if non-nil, zero value otherwise.

### GetVAdapterOk

`func (o *LeaseDiscoveredData) GetVAdapterOk() (*string, bool)`

GetVAdapterOk returns a tuple with the VAdapter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVAdapter

`func (o *LeaseDiscoveredData) SetVAdapter(v string)`

SetVAdapter sets VAdapter field to given value.

### HasVAdapter

`func (o *LeaseDiscoveredData) HasVAdapter() bool`

HasVAdapter returns a boolean if a field has been set.

### GetVCluster

`func (o *LeaseDiscoveredData) GetVCluster() string`

GetVCluster returns the VCluster field if non-nil, zero value otherwise.

### GetVClusterOk

`func (o *LeaseDiscoveredData) GetVClusterOk() (*string, bool)`

GetVClusterOk returns a tuple with the VCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVCluster

`func (o *LeaseDiscoveredData) SetVCluster(v string)`

SetVCluster sets VCluster field to given value.

### HasVCluster

`func (o *LeaseDiscoveredData) HasVCluster() bool`

HasVCluster returns a boolean if a field has been set.

### GetVDatacenter

`func (o *LeaseDiscoveredData) GetVDatacenter() string`

GetVDatacenter returns the VDatacenter field if non-nil, zero value otherwise.

### GetVDatacenterOk

`func (o *LeaseDiscoveredData) GetVDatacenterOk() (*string, bool)`

GetVDatacenterOk returns a tuple with the VDatacenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVDatacenter

`func (o *LeaseDiscoveredData) SetVDatacenter(v string)`

SetVDatacenter sets VDatacenter field to given value.

### HasVDatacenter

`func (o *LeaseDiscoveredData) HasVDatacenter() bool`

HasVDatacenter returns a boolean if a field has been set.

### GetVEntityName

`func (o *LeaseDiscoveredData) GetVEntityName() string`

GetVEntityName returns the VEntityName field if non-nil, zero value otherwise.

### GetVEntityNameOk

`func (o *LeaseDiscoveredData) GetVEntityNameOk() (*string, bool)`

GetVEntityNameOk returns a tuple with the VEntityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVEntityName

`func (o *LeaseDiscoveredData) SetVEntityName(v string)`

SetVEntityName sets VEntityName field to given value.

### HasVEntityName

`func (o *LeaseDiscoveredData) HasVEntityName() bool`

HasVEntityName returns a boolean if a field has been set.

### GetVEntityType

`func (o *LeaseDiscoveredData) GetVEntityType() string`

GetVEntityType returns the VEntityType field if non-nil, zero value otherwise.

### GetVEntityTypeOk

`func (o *LeaseDiscoveredData) GetVEntityTypeOk() (*string, bool)`

GetVEntityTypeOk returns a tuple with the VEntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVEntityType

`func (o *LeaseDiscoveredData) SetVEntityType(v string)`

SetVEntityType sets VEntityType field to given value.

### HasVEntityType

`func (o *LeaseDiscoveredData) HasVEntityType() bool`

HasVEntityType returns a boolean if a field has been set.

### GetVHost

`func (o *LeaseDiscoveredData) GetVHost() string`

GetVHost returns the VHost field if non-nil, zero value otherwise.

### GetVHostOk

`func (o *LeaseDiscoveredData) GetVHostOk() (*string, bool)`

GetVHostOk returns a tuple with the VHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVHost

`func (o *LeaseDiscoveredData) SetVHost(v string)`

SetVHost sets VHost field to given value.

### HasVHost

`func (o *LeaseDiscoveredData) HasVHost() bool`

HasVHost returns a boolean if a field has been set.

### GetVSwitch

`func (o *LeaseDiscoveredData) GetVSwitch() string`

GetVSwitch returns the VSwitch field if non-nil, zero value otherwise.

### GetVSwitchOk

`func (o *LeaseDiscoveredData) GetVSwitchOk() (*string, bool)`

GetVSwitchOk returns a tuple with the VSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVSwitch

`func (o *LeaseDiscoveredData) SetVSwitch(v string)`

SetVSwitch sets VSwitch field to given value.

### HasVSwitch

`func (o *LeaseDiscoveredData) HasVSwitch() bool`

HasVSwitch returns a boolean if a field has been set.

### GetVmiName

`func (o *LeaseDiscoveredData) GetVmiName() string`

GetVmiName returns the VmiName field if non-nil, zero value otherwise.

### GetVmiNameOk

`func (o *LeaseDiscoveredData) GetVmiNameOk() (*string, bool)`

GetVmiNameOk returns a tuple with the VmiName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiName

`func (o *LeaseDiscoveredData) SetVmiName(v string)`

SetVmiName sets VmiName field to given value.

### HasVmiName

`func (o *LeaseDiscoveredData) HasVmiName() bool`

HasVmiName returns a boolean if a field has been set.

### GetVmiId

`func (o *LeaseDiscoveredData) GetVmiId() string`

GetVmiId returns the VmiId field if non-nil, zero value otherwise.

### GetVmiIdOk

`func (o *LeaseDiscoveredData) GetVmiIdOk() (*string, bool)`

GetVmiIdOk returns a tuple with the VmiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiId

`func (o *LeaseDiscoveredData) SetVmiId(v string)`

SetVmiId sets VmiId field to given value.

### HasVmiId

`func (o *LeaseDiscoveredData) HasVmiId() bool`

HasVmiId returns a boolean if a field has been set.

### GetVlanPortGroup

`func (o *LeaseDiscoveredData) GetVlanPortGroup() string`

GetVlanPortGroup returns the VlanPortGroup field if non-nil, zero value otherwise.

### GetVlanPortGroupOk

`func (o *LeaseDiscoveredData) GetVlanPortGroupOk() (*string, bool)`

GetVlanPortGroupOk returns a tuple with the VlanPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanPortGroup

`func (o *LeaseDiscoveredData) SetVlanPortGroup(v string)`

SetVlanPortGroup sets VlanPortGroup field to given value.

### HasVlanPortGroup

`func (o *LeaseDiscoveredData) HasVlanPortGroup() bool`

HasVlanPortGroup returns a boolean if a field has been set.

### GetVswitchName

`func (o *LeaseDiscoveredData) GetVswitchName() string`

GetVswitchName returns the VswitchName field if non-nil, zero value otherwise.

### GetVswitchNameOk

`func (o *LeaseDiscoveredData) GetVswitchNameOk() (*string, bool)`

GetVswitchNameOk returns a tuple with the VswitchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchName

`func (o *LeaseDiscoveredData) SetVswitchName(v string)`

SetVswitchName sets VswitchName field to given value.

### HasVswitchName

`func (o *LeaseDiscoveredData) HasVswitchName() bool`

HasVswitchName returns a boolean if a field has been set.

### GetVswitchId

`func (o *LeaseDiscoveredData) GetVswitchId() string`

GetVswitchId returns the VswitchId field if non-nil, zero value otherwise.

### GetVswitchIdOk

`func (o *LeaseDiscoveredData) GetVswitchIdOk() (*string, bool)`

GetVswitchIdOk returns a tuple with the VswitchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchId

`func (o *LeaseDiscoveredData) SetVswitchId(v string)`

SetVswitchId sets VswitchId field to given value.

### HasVswitchId

`func (o *LeaseDiscoveredData) HasVswitchId() bool`

HasVswitchId returns a boolean if a field has been set.

### GetVswitchType

`func (o *LeaseDiscoveredData) GetVswitchType() string`

GetVswitchType returns the VswitchType field if non-nil, zero value otherwise.

### GetVswitchTypeOk

`func (o *LeaseDiscoveredData) GetVswitchTypeOk() (*string, bool)`

GetVswitchTypeOk returns a tuple with the VswitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchType

`func (o *LeaseDiscoveredData) SetVswitchType(v string)`

SetVswitchType sets VswitchType field to given value.

### HasVswitchType

`func (o *LeaseDiscoveredData) HasVswitchType() bool`

HasVswitchType returns a boolean if a field has been set.

### GetVswitchIpv6Enabled

`func (o *LeaseDiscoveredData) GetVswitchIpv6Enabled() bool`

GetVswitchIpv6Enabled returns the VswitchIpv6Enabled field if non-nil, zero value otherwise.

### GetVswitchIpv6EnabledOk

`func (o *LeaseDiscoveredData) GetVswitchIpv6EnabledOk() (*bool, bool)`

GetVswitchIpv6EnabledOk returns a tuple with the VswitchIpv6Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchIpv6Enabled

`func (o *LeaseDiscoveredData) SetVswitchIpv6Enabled(v bool)`

SetVswitchIpv6Enabled sets VswitchIpv6Enabled field to given value.

### HasVswitchIpv6Enabled

`func (o *LeaseDiscoveredData) HasVswitchIpv6Enabled() bool`

HasVswitchIpv6Enabled returns a boolean if a field has been set.

### GetVportName

`func (o *LeaseDiscoveredData) GetVportName() string`

GetVportName returns the VportName field if non-nil, zero value otherwise.

### GetVportNameOk

`func (o *LeaseDiscoveredData) GetVportNameOk() (*string, bool)`

GetVportNameOk returns a tuple with the VportName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportName

`func (o *LeaseDiscoveredData) SetVportName(v string)`

SetVportName sets VportName field to given value.

### HasVportName

`func (o *LeaseDiscoveredData) HasVportName() bool`

HasVportName returns a boolean if a field has been set.

### GetVportMacAddress

`func (o *LeaseDiscoveredData) GetVportMacAddress() string`

GetVportMacAddress returns the VportMacAddress field if non-nil, zero value otherwise.

### GetVportMacAddressOk

`func (o *LeaseDiscoveredData) GetVportMacAddressOk() (*string, bool)`

GetVportMacAddressOk returns a tuple with the VportMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportMacAddress

`func (o *LeaseDiscoveredData) SetVportMacAddress(v string)`

SetVportMacAddress sets VportMacAddress field to given value.

### HasVportMacAddress

`func (o *LeaseDiscoveredData) HasVportMacAddress() bool`

HasVportMacAddress returns a boolean if a field has been set.

### GetVportLinkStatus

`func (o *LeaseDiscoveredData) GetVportLinkStatus() string`

GetVportLinkStatus returns the VportLinkStatus field if non-nil, zero value otherwise.

### GetVportLinkStatusOk

`func (o *LeaseDiscoveredData) GetVportLinkStatusOk() (*string, bool)`

GetVportLinkStatusOk returns a tuple with the VportLinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportLinkStatus

`func (o *LeaseDiscoveredData) SetVportLinkStatus(v string)`

SetVportLinkStatus sets VportLinkStatus field to given value.

### HasVportLinkStatus

`func (o *LeaseDiscoveredData) HasVportLinkStatus() bool`

HasVportLinkStatus returns a boolean if a field has been set.

### GetVportConfSpeed

`func (o *LeaseDiscoveredData) GetVportConfSpeed() string`

GetVportConfSpeed returns the VportConfSpeed field if non-nil, zero value otherwise.

### GetVportConfSpeedOk

`func (o *LeaseDiscoveredData) GetVportConfSpeedOk() (*string, bool)`

GetVportConfSpeedOk returns a tuple with the VportConfSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportConfSpeed

`func (o *LeaseDiscoveredData) SetVportConfSpeed(v string)`

SetVportConfSpeed sets VportConfSpeed field to given value.

### HasVportConfSpeed

`func (o *LeaseDiscoveredData) HasVportConfSpeed() bool`

HasVportConfSpeed returns a boolean if a field has been set.

### GetVportConfMode

`func (o *LeaseDiscoveredData) GetVportConfMode() string`

GetVportConfMode returns the VportConfMode field if non-nil, zero value otherwise.

### GetVportConfModeOk

`func (o *LeaseDiscoveredData) GetVportConfModeOk() (*string, bool)`

GetVportConfModeOk returns a tuple with the VportConfMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportConfMode

`func (o *LeaseDiscoveredData) SetVportConfMode(v string)`

SetVportConfMode sets VportConfMode field to given value.

### HasVportConfMode

`func (o *LeaseDiscoveredData) HasVportConfMode() bool`

HasVportConfMode returns a boolean if a field has been set.

### GetVportSpeed

`func (o *LeaseDiscoveredData) GetVportSpeed() string`

GetVportSpeed returns the VportSpeed field if non-nil, zero value otherwise.

### GetVportSpeedOk

`func (o *LeaseDiscoveredData) GetVportSpeedOk() (*string, bool)`

GetVportSpeedOk returns a tuple with the VportSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportSpeed

`func (o *LeaseDiscoveredData) SetVportSpeed(v string)`

SetVportSpeed sets VportSpeed field to given value.

### HasVportSpeed

`func (o *LeaseDiscoveredData) HasVportSpeed() bool`

HasVportSpeed returns a boolean if a field has been set.

### GetVportMode

`func (o *LeaseDiscoveredData) GetVportMode() string`

GetVportMode returns the VportMode field if non-nil, zero value otherwise.

### GetVportModeOk

`func (o *LeaseDiscoveredData) GetVportModeOk() (*string, bool)`

GetVportModeOk returns a tuple with the VportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVportMode

`func (o *LeaseDiscoveredData) SetVportMode(v string)`

SetVportMode sets VportMode field to given value.

### HasVportMode

`func (o *LeaseDiscoveredData) HasVportMode() bool`

HasVportMode returns a boolean if a field has been set.

### GetVswitchSegmentType

`func (o *LeaseDiscoveredData) GetVswitchSegmentType() string`

GetVswitchSegmentType returns the VswitchSegmentType field if non-nil, zero value otherwise.

### GetVswitchSegmentTypeOk

`func (o *LeaseDiscoveredData) GetVswitchSegmentTypeOk() (*string, bool)`

GetVswitchSegmentTypeOk returns a tuple with the VswitchSegmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentType

`func (o *LeaseDiscoveredData) SetVswitchSegmentType(v string)`

SetVswitchSegmentType sets VswitchSegmentType field to given value.

### HasVswitchSegmentType

`func (o *LeaseDiscoveredData) HasVswitchSegmentType() bool`

HasVswitchSegmentType returns a boolean if a field has been set.

### GetVswitchSegmentName

`func (o *LeaseDiscoveredData) GetVswitchSegmentName() string`

GetVswitchSegmentName returns the VswitchSegmentName field if non-nil, zero value otherwise.

### GetVswitchSegmentNameOk

`func (o *LeaseDiscoveredData) GetVswitchSegmentNameOk() (*string, bool)`

GetVswitchSegmentNameOk returns a tuple with the VswitchSegmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentName

`func (o *LeaseDiscoveredData) SetVswitchSegmentName(v string)`

SetVswitchSegmentName sets VswitchSegmentName field to given value.

### HasVswitchSegmentName

`func (o *LeaseDiscoveredData) HasVswitchSegmentName() bool`

HasVswitchSegmentName returns a boolean if a field has been set.

### GetVswitchSegmentId

`func (o *LeaseDiscoveredData) GetVswitchSegmentId() string`

GetVswitchSegmentId returns the VswitchSegmentId field if non-nil, zero value otherwise.

### GetVswitchSegmentIdOk

`func (o *LeaseDiscoveredData) GetVswitchSegmentIdOk() (*string, bool)`

GetVswitchSegmentIdOk returns a tuple with the VswitchSegmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentId

`func (o *LeaseDiscoveredData) SetVswitchSegmentId(v string)`

SetVswitchSegmentId sets VswitchSegmentId field to given value.

### HasVswitchSegmentId

`func (o *LeaseDiscoveredData) HasVswitchSegmentId() bool`

HasVswitchSegmentId returns a boolean if a field has been set.

### GetVswitchSegmentPortGroup

`func (o *LeaseDiscoveredData) GetVswitchSegmentPortGroup() string`

GetVswitchSegmentPortGroup returns the VswitchSegmentPortGroup field if non-nil, zero value otherwise.

### GetVswitchSegmentPortGroupOk

`func (o *LeaseDiscoveredData) GetVswitchSegmentPortGroupOk() (*string, bool)`

GetVswitchSegmentPortGroupOk returns a tuple with the VswitchSegmentPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchSegmentPortGroup

`func (o *LeaseDiscoveredData) SetVswitchSegmentPortGroup(v string)`

SetVswitchSegmentPortGroup sets VswitchSegmentPortGroup field to given value.

### HasVswitchSegmentPortGroup

`func (o *LeaseDiscoveredData) HasVswitchSegmentPortGroup() bool`

HasVswitchSegmentPortGroup returns a boolean if a field has been set.

### GetVswitchAvailablePortsCount

`func (o *LeaseDiscoveredData) GetVswitchAvailablePortsCount() int64`

GetVswitchAvailablePortsCount returns the VswitchAvailablePortsCount field if non-nil, zero value otherwise.

### GetVswitchAvailablePortsCountOk

`func (o *LeaseDiscoveredData) GetVswitchAvailablePortsCountOk() (*int64, bool)`

GetVswitchAvailablePortsCountOk returns a tuple with the VswitchAvailablePortsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchAvailablePortsCount

`func (o *LeaseDiscoveredData) SetVswitchAvailablePortsCount(v int64)`

SetVswitchAvailablePortsCount sets VswitchAvailablePortsCount field to given value.

### HasVswitchAvailablePortsCount

`func (o *LeaseDiscoveredData) HasVswitchAvailablePortsCount() bool`

HasVswitchAvailablePortsCount returns a boolean if a field has been set.

### GetVswitchTepType

`func (o *LeaseDiscoveredData) GetVswitchTepType() string`

GetVswitchTepType returns the VswitchTepType field if non-nil, zero value otherwise.

### GetVswitchTepTypeOk

`func (o *LeaseDiscoveredData) GetVswitchTepTypeOk() (*string, bool)`

GetVswitchTepTypeOk returns a tuple with the VswitchTepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepType

`func (o *LeaseDiscoveredData) SetVswitchTepType(v string)`

SetVswitchTepType sets VswitchTepType field to given value.

### HasVswitchTepType

`func (o *LeaseDiscoveredData) HasVswitchTepType() bool`

HasVswitchTepType returns a boolean if a field has been set.

### GetVswitchTepIp

`func (o *LeaseDiscoveredData) GetVswitchTepIp() string`

GetVswitchTepIp returns the VswitchTepIp field if non-nil, zero value otherwise.

### GetVswitchTepIpOk

`func (o *LeaseDiscoveredData) GetVswitchTepIpOk() (*string, bool)`

GetVswitchTepIpOk returns a tuple with the VswitchTepIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepIp

`func (o *LeaseDiscoveredData) SetVswitchTepIp(v string)`

SetVswitchTepIp sets VswitchTepIp field to given value.

### HasVswitchTepIp

`func (o *LeaseDiscoveredData) HasVswitchTepIp() bool`

HasVswitchTepIp returns a boolean if a field has been set.

### GetVswitchTepPortGroup

`func (o *LeaseDiscoveredData) GetVswitchTepPortGroup() string`

GetVswitchTepPortGroup returns the VswitchTepPortGroup field if non-nil, zero value otherwise.

### GetVswitchTepPortGroupOk

`func (o *LeaseDiscoveredData) GetVswitchTepPortGroupOk() (*string, bool)`

GetVswitchTepPortGroupOk returns a tuple with the VswitchTepPortGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepPortGroup

`func (o *LeaseDiscoveredData) SetVswitchTepPortGroup(v string)`

SetVswitchTepPortGroup sets VswitchTepPortGroup field to given value.

### HasVswitchTepPortGroup

`func (o *LeaseDiscoveredData) HasVswitchTepPortGroup() bool`

HasVswitchTepPortGroup returns a boolean if a field has been set.

### GetVswitchTepVlan

`func (o *LeaseDiscoveredData) GetVswitchTepVlan() string`

GetVswitchTepVlan returns the VswitchTepVlan field if non-nil, zero value otherwise.

### GetVswitchTepVlanOk

`func (o *LeaseDiscoveredData) GetVswitchTepVlanOk() (*string, bool)`

GetVswitchTepVlanOk returns a tuple with the VswitchTepVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepVlan

`func (o *LeaseDiscoveredData) SetVswitchTepVlan(v string)`

SetVswitchTepVlan sets VswitchTepVlan field to given value.

### HasVswitchTepVlan

`func (o *LeaseDiscoveredData) HasVswitchTepVlan() bool`

HasVswitchTepVlan returns a boolean if a field has been set.

### GetVswitchTepDhcpServer

`func (o *LeaseDiscoveredData) GetVswitchTepDhcpServer() string`

GetVswitchTepDhcpServer returns the VswitchTepDhcpServer field if non-nil, zero value otherwise.

### GetVswitchTepDhcpServerOk

`func (o *LeaseDiscoveredData) GetVswitchTepDhcpServerOk() (*string, bool)`

GetVswitchTepDhcpServerOk returns a tuple with the VswitchTepDhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepDhcpServer

`func (o *LeaseDiscoveredData) SetVswitchTepDhcpServer(v string)`

SetVswitchTepDhcpServer sets VswitchTepDhcpServer field to given value.

### HasVswitchTepDhcpServer

`func (o *LeaseDiscoveredData) HasVswitchTepDhcpServer() bool`

HasVswitchTepDhcpServer returns a boolean if a field has been set.

### GetVswitchTepMulticast

`func (o *LeaseDiscoveredData) GetVswitchTepMulticast() string`

GetVswitchTepMulticast returns the VswitchTepMulticast field if non-nil, zero value otherwise.

### GetVswitchTepMulticastOk

`func (o *LeaseDiscoveredData) GetVswitchTepMulticastOk() (*string, bool)`

GetVswitchTepMulticastOk returns a tuple with the VswitchTepMulticast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVswitchTepMulticast

`func (o *LeaseDiscoveredData) SetVswitchTepMulticast(v string)`

SetVswitchTepMulticast sets VswitchTepMulticast field to given value.

### HasVswitchTepMulticast

`func (o *LeaseDiscoveredData) HasVswitchTepMulticast() bool`

HasVswitchTepMulticast returns a boolean if a field has been set.

### GetVmhostIpAddress

`func (o *LeaseDiscoveredData) GetVmhostIpAddress() string`

GetVmhostIpAddress returns the VmhostIpAddress field if non-nil, zero value otherwise.

### GetVmhostIpAddressOk

`func (o *LeaseDiscoveredData) GetVmhostIpAddressOk() (*string, bool)`

GetVmhostIpAddressOk returns a tuple with the VmhostIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostIpAddress

`func (o *LeaseDiscoveredData) SetVmhostIpAddress(v string)`

SetVmhostIpAddress sets VmhostIpAddress field to given value.

### HasVmhostIpAddress

`func (o *LeaseDiscoveredData) HasVmhostIpAddress() bool`

HasVmhostIpAddress returns a boolean if a field has been set.

### GetVmhostName

`func (o *LeaseDiscoveredData) GetVmhostName() string`

GetVmhostName returns the VmhostName field if non-nil, zero value otherwise.

### GetVmhostNameOk

`func (o *LeaseDiscoveredData) GetVmhostNameOk() (*string, bool)`

GetVmhostNameOk returns a tuple with the VmhostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostName

`func (o *LeaseDiscoveredData) SetVmhostName(v string)`

SetVmhostName sets VmhostName field to given value.

### HasVmhostName

`func (o *LeaseDiscoveredData) HasVmhostName() bool`

HasVmhostName returns a boolean if a field has been set.

### GetVmhostMacAddress

`func (o *LeaseDiscoveredData) GetVmhostMacAddress() string`

GetVmhostMacAddress returns the VmhostMacAddress field if non-nil, zero value otherwise.

### GetVmhostMacAddressOk

`func (o *LeaseDiscoveredData) GetVmhostMacAddressOk() (*string, bool)`

GetVmhostMacAddressOk returns a tuple with the VmhostMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostMacAddress

`func (o *LeaseDiscoveredData) SetVmhostMacAddress(v string)`

SetVmhostMacAddress sets VmhostMacAddress field to given value.

### HasVmhostMacAddress

`func (o *LeaseDiscoveredData) HasVmhostMacAddress() bool`

HasVmhostMacAddress returns a boolean if a field has been set.

### GetVmhostSubnetCidr

`func (o *LeaseDiscoveredData) GetVmhostSubnetCidr() int64`

GetVmhostSubnetCidr returns the VmhostSubnetCidr field if non-nil, zero value otherwise.

### GetVmhostSubnetCidrOk

`func (o *LeaseDiscoveredData) GetVmhostSubnetCidrOk() (*int64, bool)`

GetVmhostSubnetCidrOk returns a tuple with the VmhostSubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostSubnetCidr

`func (o *LeaseDiscoveredData) SetVmhostSubnetCidr(v int64)`

SetVmhostSubnetCidr sets VmhostSubnetCidr field to given value.

### HasVmhostSubnetCidr

`func (o *LeaseDiscoveredData) HasVmhostSubnetCidr() bool`

HasVmhostSubnetCidr returns a boolean if a field has been set.

### GetVmhostNicNames

`func (o *LeaseDiscoveredData) GetVmhostNicNames() string`

GetVmhostNicNames returns the VmhostNicNames field if non-nil, zero value otherwise.

### GetVmhostNicNamesOk

`func (o *LeaseDiscoveredData) GetVmhostNicNamesOk() (*string, bool)`

GetVmhostNicNamesOk returns a tuple with the VmhostNicNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmhostNicNames

`func (o *LeaseDiscoveredData) SetVmhostNicNames(v string)`

SetVmhostNicNames sets VmhostNicNames field to given value.

### HasVmhostNicNames

`func (o *LeaseDiscoveredData) HasVmhostNicNames() bool`

HasVmhostNicNames returns a boolean if a field has been set.

### GetVmiTenantId

`func (o *LeaseDiscoveredData) GetVmiTenantId() string`

GetVmiTenantId returns the VmiTenantId field if non-nil, zero value otherwise.

### GetVmiTenantIdOk

`func (o *LeaseDiscoveredData) GetVmiTenantIdOk() (*string, bool)`

GetVmiTenantIdOk returns a tuple with the VmiTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiTenantId

`func (o *LeaseDiscoveredData) SetVmiTenantId(v string)`

SetVmiTenantId sets VmiTenantId field to given value.

### HasVmiTenantId

`func (o *LeaseDiscoveredData) HasVmiTenantId() bool`

HasVmiTenantId returns a boolean if a field has been set.

### GetCmpType

`func (o *LeaseDiscoveredData) GetCmpType() string`

GetCmpType returns the CmpType field if non-nil, zero value otherwise.

### GetCmpTypeOk

`func (o *LeaseDiscoveredData) GetCmpTypeOk() (*string, bool)`

GetCmpTypeOk returns a tuple with the CmpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmpType

`func (o *LeaseDiscoveredData) SetCmpType(v string)`

SetCmpType sets CmpType field to given value.

### HasCmpType

`func (o *LeaseDiscoveredData) HasCmpType() bool`

HasCmpType returns a boolean if a field has been set.

### GetVmiIpType

`func (o *LeaseDiscoveredData) GetVmiIpType() string`

GetVmiIpType returns the VmiIpType field if non-nil, zero value otherwise.

### GetVmiIpTypeOk

`func (o *LeaseDiscoveredData) GetVmiIpTypeOk() (*string, bool)`

GetVmiIpTypeOk returns a tuple with the VmiIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiIpType

`func (o *LeaseDiscoveredData) SetVmiIpType(v string)`

SetVmiIpType sets VmiIpType field to given value.

### HasVmiIpType

`func (o *LeaseDiscoveredData) HasVmiIpType() bool`

HasVmiIpType returns a boolean if a field has been set.

### GetVmiPrivateAddress

`func (o *LeaseDiscoveredData) GetVmiPrivateAddress() string`

GetVmiPrivateAddress returns the VmiPrivateAddress field if non-nil, zero value otherwise.

### GetVmiPrivateAddressOk

`func (o *LeaseDiscoveredData) GetVmiPrivateAddressOk() (*string, bool)`

GetVmiPrivateAddressOk returns a tuple with the VmiPrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiPrivateAddress

`func (o *LeaseDiscoveredData) SetVmiPrivateAddress(v string)`

SetVmiPrivateAddress sets VmiPrivateAddress field to given value.

### HasVmiPrivateAddress

`func (o *LeaseDiscoveredData) HasVmiPrivateAddress() bool`

HasVmiPrivateAddress returns a boolean if a field has been set.

### GetVmiIsPublicAddress

`func (o *LeaseDiscoveredData) GetVmiIsPublicAddress() bool`

GetVmiIsPublicAddress returns the VmiIsPublicAddress field if non-nil, zero value otherwise.

### GetVmiIsPublicAddressOk

`func (o *LeaseDiscoveredData) GetVmiIsPublicAddressOk() (*bool, bool)`

GetVmiIsPublicAddressOk returns a tuple with the VmiIsPublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmiIsPublicAddress

`func (o *LeaseDiscoveredData) SetVmiIsPublicAddress(v bool)`

SetVmiIsPublicAddress sets VmiIsPublicAddress field to given value.

### HasVmiIsPublicAddress

`func (o *LeaseDiscoveredData) HasVmiIsPublicAddress() bool`

HasVmiIsPublicAddress returns a boolean if a field has been set.

### GetCiscoIseSsid

`func (o *LeaseDiscoveredData) GetCiscoIseSsid() string`

GetCiscoIseSsid returns the CiscoIseSsid field if non-nil, zero value otherwise.

### GetCiscoIseSsidOk

`func (o *LeaseDiscoveredData) GetCiscoIseSsidOk() (*string, bool)`

GetCiscoIseSsidOk returns a tuple with the CiscoIseSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseSsid

`func (o *LeaseDiscoveredData) SetCiscoIseSsid(v string)`

SetCiscoIseSsid sets CiscoIseSsid field to given value.

### HasCiscoIseSsid

`func (o *LeaseDiscoveredData) HasCiscoIseSsid() bool`

HasCiscoIseSsid returns a boolean if a field has been set.

### GetCiscoIseEndpointProfile

`func (o *LeaseDiscoveredData) GetCiscoIseEndpointProfile() string`

GetCiscoIseEndpointProfile returns the CiscoIseEndpointProfile field if non-nil, zero value otherwise.

### GetCiscoIseEndpointProfileOk

`func (o *LeaseDiscoveredData) GetCiscoIseEndpointProfileOk() (*string, bool)`

GetCiscoIseEndpointProfileOk returns a tuple with the CiscoIseEndpointProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseEndpointProfile

`func (o *LeaseDiscoveredData) SetCiscoIseEndpointProfile(v string)`

SetCiscoIseEndpointProfile sets CiscoIseEndpointProfile field to given value.

### HasCiscoIseEndpointProfile

`func (o *LeaseDiscoveredData) HasCiscoIseEndpointProfile() bool`

HasCiscoIseEndpointProfile returns a boolean if a field has been set.

### GetCiscoIseSessionState

`func (o *LeaseDiscoveredData) GetCiscoIseSessionState() string`

GetCiscoIseSessionState returns the CiscoIseSessionState field if non-nil, zero value otherwise.

### GetCiscoIseSessionStateOk

`func (o *LeaseDiscoveredData) GetCiscoIseSessionStateOk() (*string, bool)`

GetCiscoIseSessionStateOk returns a tuple with the CiscoIseSessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseSessionState

`func (o *LeaseDiscoveredData) SetCiscoIseSessionState(v string)`

SetCiscoIseSessionState sets CiscoIseSessionState field to given value.

### HasCiscoIseSessionState

`func (o *LeaseDiscoveredData) HasCiscoIseSessionState() bool`

HasCiscoIseSessionState returns a boolean if a field has been set.

### GetCiscoIseSecurityGroup

`func (o *LeaseDiscoveredData) GetCiscoIseSecurityGroup() string`

GetCiscoIseSecurityGroup returns the CiscoIseSecurityGroup field if non-nil, zero value otherwise.

### GetCiscoIseSecurityGroupOk

`func (o *LeaseDiscoveredData) GetCiscoIseSecurityGroupOk() (*string, bool)`

GetCiscoIseSecurityGroupOk returns a tuple with the CiscoIseSecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiscoIseSecurityGroup

`func (o *LeaseDiscoveredData) SetCiscoIseSecurityGroup(v string)`

SetCiscoIseSecurityGroup sets CiscoIseSecurityGroup field to given value.

### HasCiscoIseSecurityGroup

`func (o *LeaseDiscoveredData) HasCiscoIseSecurityGroup() bool`

HasCiscoIseSecurityGroup returns a boolean if a field has been set.

### GetTaskName

`func (o *LeaseDiscoveredData) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *LeaseDiscoveredData) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *LeaseDiscoveredData) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *LeaseDiscoveredData) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetNetworkComponentLocation

`func (o *LeaseDiscoveredData) GetNetworkComponentLocation() string`

GetNetworkComponentLocation returns the NetworkComponentLocation field if non-nil, zero value otherwise.

### GetNetworkComponentLocationOk

`func (o *LeaseDiscoveredData) GetNetworkComponentLocationOk() (*string, bool)`

GetNetworkComponentLocationOk returns a tuple with the NetworkComponentLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentLocation

`func (o *LeaseDiscoveredData) SetNetworkComponentLocation(v string)`

SetNetworkComponentLocation sets NetworkComponentLocation field to given value.

### HasNetworkComponentLocation

`func (o *LeaseDiscoveredData) HasNetworkComponentLocation() bool`

HasNetworkComponentLocation returns a boolean if a field has been set.

### GetNetworkComponentContact

`func (o *LeaseDiscoveredData) GetNetworkComponentContact() string`

GetNetworkComponentContact returns the NetworkComponentContact field if non-nil, zero value otherwise.

### GetNetworkComponentContactOk

`func (o *LeaseDiscoveredData) GetNetworkComponentContactOk() (*string, bool)`

GetNetworkComponentContactOk returns a tuple with the NetworkComponentContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkComponentContact

`func (o *LeaseDiscoveredData) SetNetworkComponentContact(v string)`

SetNetworkComponentContact sets NetworkComponentContact field to given value.

### HasNetworkComponentContact

`func (o *LeaseDiscoveredData) HasNetworkComponentContact() bool`

HasNetworkComponentContact returns a boolean if a field has been set.

### GetDeviceLocation

`func (o *LeaseDiscoveredData) GetDeviceLocation() string`

GetDeviceLocation returns the DeviceLocation field if non-nil, zero value otherwise.

### GetDeviceLocationOk

`func (o *LeaseDiscoveredData) GetDeviceLocationOk() (*string, bool)`

GetDeviceLocationOk returns a tuple with the DeviceLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceLocation

`func (o *LeaseDiscoveredData) SetDeviceLocation(v string)`

SetDeviceLocation sets DeviceLocation field to given value.

### HasDeviceLocation

`func (o *LeaseDiscoveredData) HasDeviceLocation() bool`

HasDeviceLocation returns a boolean if a field has been set.

### GetDeviceContact

`func (o *LeaseDiscoveredData) GetDeviceContact() string`

GetDeviceContact returns the DeviceContact field if non-nil, zero value otherwise.

### GetDeviceContactOk

`func (o *LeaseDiscoveredData) GetDeviceContactOk() (*string, bool)`

GetDeviceContactOk returns a tuple with the DeviceContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceContact

`func (o *LeaseDiscoveredData) SetDeviceContact(v string)`

SetDeviceContact sets DeviceContact field to given value.

### HasDeviceContact

`func (o *LeaseDiscoveredData) HasDeviceContact() bool`

HasDeviceContact returns a boolean if a field has been set.

### GetApName

`func (o *LeaseDiscoveredData) GetApName() string`

GetApName returns the ApName field if non-nil, zero value otherwise.

### GetApNameOk

`func (o *LeaseDiscoveredData) GetApNameOk() (*string, bool)`

GetApNameOk returns a tuple with the ApName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApName

`func (o *LeaseDiscoveredData) SetApName(v string)`

SetApName sets ApName field to given value.

### HasApName

`func (o *LeaseDiscoveredData) HasApName() bool`

HasApName returns a boolean if a field has been set.

### GetApIpAddress

`func (o *LeaseDiscoveredData) GetApIpAddress() string`

GetApIpAddress returns the ApIpAddress field if non-nil, zero value otherwise.

### GetApIpAddressOk

`func (o *LeaseDiscoveredData) GetApIpAddressOk() (*string, bool)`

GetApIpAddressOk returns a tuple with the ApIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApIpAddress

`func (o *LeaseDiscoveredData) SetApIpAddress(v string)`

SetApIpAddress sets ApIpAddress field to given value.

### HasApIpAddress

`func (o *LeaseDiscoveredData) HasApIpAddress() bool`

HasApIpAddress returns a boolean if a field has been set.

### GetApSsid

`func (o *LeaseDiscoveredData) GetApSsid() string`

GetApSsid returns the ApSsid field if non-nil, zero value otherwise.

### GetApSsidOk

`func (o *LeaseDiscoveredData) GetApSsidOk() (*string, bool)`

GetApSsidOk returns a tuple with the ApSsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApSsid

`func (o *LeaseDiscoveredData) SetApSsid(v string)`

SetApSsid sets ApSsid field to given value.

### HasApSsid

`func (o *LeaseDiscoveredData) HasApSsid() bool`

HasApSsid returns a boolean if a field has been set.

### GetBridgeDomain

`func (o *LeaseDiscoveredData) GetBridgeDomain() string`

GetBridgeDomain returns the BridgeDomain field if non-nil, zero value otherwise.

### GetBridgeDomainOk

`func (o *LeaseDiscoveredData) GetBridgeDomainOk() (*string, bool)`

GetBridgeDomainOk returns a tuple with the BridgeDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridgeDomain

`func (o *LeaseDiscoveredData) SetBridgeDomain(v string)`

SetBridgeDomain sets BridgeDomain field to given value.

### HasBridgeDomain

`func (o *LeaseDiscoveredData) HasBridgeDomain() bool`

HasBridgeDomain returns a boolean if a field has been set.

### GetEndpointGroups

`func (o *LeaseDiscoveredData) GetEndpointGroups() string`

GetEndpointGroups returns the EndpointGroups field if non-nil, zero value otherwise.

### GetEndpointGroupsOk

`func (o *LeaseDiscoveredData) GetEndpointGroupsOk() (*string, bool)`

GetEndpointGroupsOk returns a tuple with the EndpointGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointGroups

`func (o *LeaseDiscoveredData) SetEndpointGroups(v string)`

SetEndpointGroups sets EndpointGroups field to given value.

### HasEndpointGroups

`func (o *LeaseDiscoveredData) HasEndpointGroups() bool`

HasEndpointGroups returns a boolean if a field has been set.

### GetTenant

`func (o *LeaseDiscoveredData) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *LeaseDiscoveredData) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *LeaseDiscoveredData) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *LeaseDiscoveredData) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetVrfName

`func (o *LeaseDiscoveredData) GetVrfName() string`

GetVrfName returns the VrfName field if non-nil, zero value otherwise.

### GetVrfNameOk

`func (o *LeaseDiscoveredData) GetVrfNameOk() (*string, bool)`

GetVrfNameOk returns a tuple with the VrfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfName

`func (o *LeaseDiscoveredData) SetVrfName(v string)`

SetVrfName sets VrfName field to given value.

### HasVrfName

`func (o *LeaseDiscoveredData) HasVrfName() bool`

HasVrfName returns a boolean if a field has been set.

### GetVrfDescription

`func (o *LeaseDiscoveredData) GetVrfDescription() string`

GetVrfDescription returns the VrfDescription field if non-nil, zero value otherwise.

### GetVrfDescriptionOk

`func (o *LeaseDiscoveredData) GetVrfDescriptionOk() (*string, bool)`

GetVrfDescriptionOk returns a tuple with the VrfDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfDescription

`func (o *LeaseDiscoveredData) SetVrfDescription(v string)`

SetVrfDescription sets VrfDescription field to given value.

### HasVrfDescription

`func (o *LeaseDiscoveredData) HasVrfDescription() bool`

HasVrfDescription returns a boolean if a field has been set.

### GetVrfRd

`func (o *LeaseDiscoveredData) GetVrfRd() string`

GetVrfRd returns the VrfRd field if non-nil, zero value otherwise.

### GetVrfRdOk

`func (o *LeaseDiscoveredData) GetVrfRdOk() (*string, bool)`

GetVrfRdOk returns a tuple with the VrfRd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfRd

`func (o *LeaseDiscoveredData) SetVrfRd(v string)`

SetVrfRd sets VrfRd field to given value.

### HasVrfRd

`func (o *LeaseDiscoveredData) HasVrfRd() bool`

HasVrfRd returns a boolean if a field has been set.

### GetBgpAs

`func (o *LeaseDiscoveredData) GetBgpAs() int64`

GetBgpAs returns the BgpAs field if non-nil, zero value otherwise.

### GetBgpAsOk

`func (o *LeaseDiscoveredData) GetBgpAsOk() (*int64, bool)`

GetBgpAsOk returns a tuple with the BgpAs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAs

`func (o *LeaseDiscoveredData) SetBgpAs(v int64)`

SetBgpAs sets BgpAs field to given value.

### HasBgpAs

`func (o *LeaseDiscoveredData) HasBgpAs() bool`

HasBgpAs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


