# GetGridCloudapiVmaddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Address** | Pointer to **string** | The IP address of the interface. | [optional] [readonly] 
**AddressType** | Pointer to **string** | IP address type (Public, Private, Elastic, Floating, ...). | [optional] [readonly] 
**AssociatedIp** | Pointer to **string** | Reference to associated IPv4 or IPv6 address. | [optional] [readonly] 
**AssociatedObjectTypes** | Pointer to **[]string** | Array of string denoting the types of underlying objects IPv4/IPv6 - \&quot;A\&quot;, \&quot;AAAA\&quot;, \&quot;PTR\&quot;, \&quot;HOST\&quot;, \&quot;FA\&quot;, \&quot;RESERVATION\&quot;, \&quot;UNMANAGED\&quot; + (\&quot;BULKHOST\&quot;, \&quot;DHCP_RANGE\&quot;, \&quot;RESERVED_RANGE\&quot;, \&quot;LEASE\&quot;, \&quot;NETWORK\&quot;, \&quot;BROADCAST\&quot;, \&quot;PENDING\&quot;), | [optional] [readonly] 
**AssociatedObjects** | Pointer to **[]string** | The list of references to the object (Host, Fixed Address, RR, ...) that defines this IP. | [optional] [readonly] 
**CloudInfo** | Pointer to [**GridCloudapiVmaddressCloudInfo**](GridCloudapiVmaddressCloudInfo.md) |  | [optional] 
**DnsNames** | Pointer to **[]string** | The list of all FQDNs associated with the IP address. | [optional] [readonly] 
**ElasticAddress** | Pointer to **string** | Elastic IP address associated with this private address, if this address is a private address; otherwise empty. | [optional] [readonly] 
**InterfaceName** | Pointer to **string** | Name of the interface associated with this IP address. | [optional] [readonly] 
**IsIpv4** | Pointer to **bool** | Indicates whether the address is IPv4 or IPv6. | [optional] [readonly] 
**MacAddress** | Pointer to **string** | The MAC address of the interface. | [optional] [readonly] 
**MsAdUserData** | Pointer to [**GridCloudapiVmaddressMsAdUserData**](GridCloudapiVmaddressMsAdUserData.md) |  | [optional] 
**Network** | Pointer to **string** | The network to which this address belongs, in IPv4 Address/CIDR format. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | Network view name of the delegated object. | [optional] [readonly] 
**PortId** | Pointer to **int64** | Port identifier of the interface. | [optional] [readonly] 
**PrivateAddress** | Pointer to **string** | Private IP address associated with this public (or elastic or floating) address, if this address is a public address; otherwise empty. | [optional] [readonly] 
**PrivateHostname** | Pointer to **string** | Host part of the FQDN of this address if this address is a private address; otherwise empty | [optional] [readonly] 
**PublicAddress** | Pointer to **string** | Public IP address associated with this private address, if this address is a private address; otherwise empty. | [optional] [readonly] 
**PublicHostname** | Pointer to **string** | Host part of the FQDN of this address if this address is a public (or elastic or floating) address; otherwise empty | [optional] [readonly] 
**SubnetAddress** | Pointer to **string** | Network address of the subnet that is the container of this address. | [optional] [readonly] 
**SubnetCidr** | Pointer to **int64** | CIDR of the subnet that is the container of this address. | [optional] [readonly] 
**SubnetId** | Pointer to **string** | Subnet ID that is the container of this address. | [optional] [readonly] 
**Tenant** | Pointer to **string** | The Cloud API Tenant object. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**VmAvailabilityZone** | Pointer to **string** | Availability zone of the VM. | [optional] [readonly] 
**VmComment** | Pointer to **string** | VM comment. | [optional] [readonly] 
**VmCreationTime** | Pointer to **int64** | Date/time the VM was first created as NIOS object. | [optional] [readonly] 
**VmHostname** | Pointer to **string** | Host part of the FQDN of the address attached to the primary interface. | [optional] [readonly] 
**VmId** | Pointer to **string** | The UUID of the Virtual Machine. | [optional] [readonly] 
**VmKernelId** | Pointer to **string** | Kernel ID of the VM that this address is associated with. | [optional] [readonly] 
**VmLastUpdateTime** | Pointer to **int64** | Last time the VM was updated. | [optional] [readonly] 
**VmName** | Pointer to **string** | The name of the Virtual Machine. | [optional] [readonly] 
**VmNetworkCount** | Pointer to **int64** | Count of networks containing all the addresses of the VM. | [optional] [readonly] 
**VmOperatingSystem** | Pointer to **string** | Operating system that the VM is running. | [optional] [readonly] 
**VmType** | Pointer to **string** | Type of the VM this address is associated with. | [optional] [readonly] 
**VmVpcAddress** | Pointer to **string** | Network address of the VPC of the VM that this address is associated with. | [optional] [readonly] 
**VmVpcCidr** | Pointer to **int64** | CIDR of the VPC of the VM that this address is associated with. | [optional] [readonly] 
**VmVpcId** | Pointer to **string** | Identifier of the VPC where the VM is defined. | [optional] [readonly] 
**VmVpcName** | Pointer to **string** | Name of the VPC where the VM is defined. | [optional] [readonly] 
**VmVpcRef** | Pointer to **string** | Reference to the VPC where the VM is defined. | [optional] [readonly] 
**Result** | Pointer to [**GridCloudapiVmaddress**](GridCloudapiVmaddress.md) |  | [optional] 

## Methods

### NewGetGridCloudapiVmaddressResponse

`func NewGetGridCloudapiVmaddressResponse() *GetGridCloudapiVmaddressResponse`

NewGetGridCloudapiVmaddressResponse instantiates a new GetGridCloudapiVmaddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridCloudapiVmaddressResponseWithDefaults

`func NewGetGridCloudapiVmaddressResponseWithDefaults() *GetGridCloudapiVmaddressResponse`

NewGetGridCloudapiVmaddressResponseWithDefaults instantiates a new GetGridCloudapiVmaddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridCloudapiVmaddressResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridCloudapiVmaddressResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridCloudapiVmaddressResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridCloudapiVmaddressResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAddress

`func (o *GetGridCloudapiVmaddressResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetGridCloudapiVmaddressResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetGridCloudapiVmaddressResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetGridCloudapiVmaddressResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAddressType

`func (o *GetGridCloudapiVmaddressResponse) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *GetGridCloudapiVmaddressResponse) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *GetGridCloudapiVmaddressResponse) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *GetGridCloudapiVmaddressResponse) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetAssociatedIp

`func (o *GetGridCloudapiVmaddressResponse) GetAssociatedIp() string`

GetAssociatedIp returns the AssociatedIp field if non-nil, zero value otherwise.

### GetAssociatedIpOk

`func (o *GetGridCloudapiVmaddressResponse) GetAssociatedIpOk() (*string, bool)`

GetAssociatedIpOk returns a tuple with the AssociatedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedIp

`func (o *GetGridCloudapiVmaddressResponse) SetAssociatedIp(v string)`

SetAssociatedIp sets AssociatedIp field to given value.

### HasAssociatedIp

`func (o *GetGridCloudapiVmaddressResponse) HasAssociatedIp() bool`

HasAssociatedIp returns a boolean if a field has been set.

### GetAssociatedObjectTypes

`func (o *GetGridCloudapiVmaddressResponse) GetAssociatedObjectTypes() []string`

GetAssociatedObjectTypes returns the AssociatedObjectTypes field if non-nil, zero value otherwise.

### GetAssociatedObjectTypesOk

`func (o *GetGridCloudapiVmaddressResponse) GetAssociatedObjectTypesOk() (*[]string, bool)`

GetAssociatedObjectTypesOk returns a tuple with the AssociatedObjectTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjectTypes

`func (o *GetGridCloudapiVmaddressResponse) SetAssociatedObjectTypes(v []string)`

SetAssociatedObjectTypes sets AssociatedObjectTypes field to given value.

### HasAssociatedObjectTypes

`func (o *GetGridCloudapiVmaddressResponse) HasAssociatedObjectTypes() bool`

HasAssociatedObjectTypes returns a boolean if a field has been set.

### GetAssociatedObjects

`func (o *GetGridCloudapiVmaddressResponse) GetAssociatedObjects() []string`

GetAssociatedObjects returns the AssociatedObjects field if non-nil, zero value otherwise.

### GetAssociatedObjectsOk

`func (o *GetGridCloudapiVmaddressResponse) GetAssociatedObjectsOk() (*[]string, bool)`

GetAssociatedObjectsOk returns a tuple with the AssociatedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedObjects

`func (o *GetGridCloudapiVmaddressResponse) SetAssociatedObjects(v []string)`

SetAssociatedObjects sets AssociatedObjects field to given value.

### HasAssociatedObjects

`func (o *GetGridCloudapiVmaddressResponse) HasAssociatedObjects() bool`

HasAssociatedObjects returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetGridCloudapiVmaddressResponse) GetCloudInfo() GridCloudapiVmaddressCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetGridCloudapiVmaddressResponse) GetCloudInfoOk() (*GridCloudapiVmaddressCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetGridCloudapiVmaddressResponse) SetCloudInfo(v GridCloudapiVmaddressCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetGridCloudapiVmaddressResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetDnsNames

`func (o *GetGridCloudapiVmaddressResponse) GetDnsNames() []string`

GetDnsNames returns the DnsNames field if non-nil, zero value otherwise.

### GetDnsNamesOk

`func (o *GetGridCloudapiVmaddressResponse) GetDnsNamesOk() (*[]string, bool)`

GetDnsNamesOk returns a tuple with the DnsNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsNames

`func (o *GetGridCloudapiVmaddressResponse) SetDnsNames(v []string)`

SetDnsNames sets DnsNames field to given value.

### HasDnsNames

`func (o *GetGridCloudapiVmaddressResponse) HasDnsNames() bool`

HasDnsNames returns a boolean if a field has been set.

### GetElasticAddress

`func (o *GetGridCloudapiVmaddressResponse) GetElasticAddress() string`

GetElasticAddress returns the ElasticAddress field if non-nil, zero value otherwise.

### GetElasticAddressOk

`func (o *GetGridCloudapiVmaddressResponse) GetElasticAddressOk() (*string, bool)`

GetElasticAddressOk returns a tuple with the ElasticAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticAddress

`func (o *GetGridCloudapiVmaddressResponse) SetElasticAddress(v string)`

SetElasticAddress sets ElasticAddress field to given value.

### HasElasticAddress

`func (o *GetGridCloudapiVmaddressResponse) HasElasticAddress() bool`

HasElasticAddress returns a boolean if a field has been set.

### GetInterfaceName

`func (o *GetGridCloudapiVmaddressResponse) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *GetGridCloudapiVmaddressResponse) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *GetGridCloudapiVmaddressResponse) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *GetGridCloudapiVmaddressResponse) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetIsIpv4

`func (o *GetGridCloudapiVmaddressResponse) GetIsIpv4() bool`

GetIsIpv4 returns the IsIpv4 field if non-nil, zero value otherwise.

### GetIsIpv4Ok

`func (o *GetGridCloudapiVmaddressResponse) GetIsIpv4Ok() (*bool, bool)`

GetIsIpv4Ok returns a tuple with the IsIpv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIpv4

`func (o *GetGridCloudapiVmaddressResponse) SetIsIpv4(v bool)`

SetIsIpv4 sets IsIpv4 field to given value.

### HasIsIpv4

`func (o *GetGridCloudapiVmaddressResponse) HasIsIpv4() bool`

HasIsIpv4 returns a boolean if a field has been set.

### GetMacAddress

`func (o *GetGridCloudapiVmaddressResponse) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *GetGridCloudapiVmaddressResponse) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *GetGridCloudapiVmaddressResponse) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *GetGridCloudapiVmaddressResponse) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetMsAdUserData

`func (o *GetGridCloudapiVmaddressResponse) GetMsAdUserData() GridCloudapiVmaddressMsAdUserData`

GetMsAdUserData returns the MsAdUserData field if non-nil, zero value otherwise.

### GetMsAdUserDataOk

`func (o *GetGridCloudapiVmaddressResponse) GetMsAdUserDataOk() (*GridCloudapiVmaddressMsAdUserData, bool)`

GetMsAdUserDataOk returns a tuple with the MsAdUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsAdUserData

`func (o *GetGridCloudapiVmaddressResponse) SetMsAdUserData(v GridCloudapiVmaddressMsAdUserData)`

SetMsAdUserData sets MsAdUserData field to given value.

### HasMsAdUserData

`func (o *GetGridCloudapiVmaddressResponse) HasMsAdUserData() bool`

HasMsAdUserData returns a boolean if a field has been set.

### GetNetwork

`func (o *GetGridCloudapiVmaddressResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetGridCloudapiVmaddressResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetGridCloudapiVmaddressResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetGridCloudapiVmaddressResponse) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkView

`func (o *GetGridCloudapiVmaddressResponse) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *GetGridCloudapiVmaddressResponse) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *GetGridCloudapiVmaddressResponse) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *GetGridCloudapiVmaddressResponse) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetPortId

`func (o *GetGridCloudapiVmaddressResponse) GetPortId() int64`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *GetGridCloudapiVmaddressResponse) GetPortIdOk() (*int64, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *GetGridCloudapiVmaddressResponse) SetPortId(v int64)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *GetGridCloudapiVmaddressResponse) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPrivateAddress

`func (o *GetGridCloudapiVmaddressResponse) GetPrivateAddress() string`

GetPrivateAddress returns the PrivateAddress field if non-nil, zero value otherwise.

### GetPrivateAddressOk

`func (o *GetGridCloudapiVmaddressResponse) GetPrivateAddressOk() (*string, bool)`

GetPrivateAddressOk returns a tuple with the PrivateAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateAddress

`func (o *GetGridCloudapiVmaddressResponse) SetPrivateAddress(v string)`

SetPrivateAddress sets PrivateAddress field to given value.

### HasPrivateAddress

`func (o *GetGridCloudapiVmaddressResponse) HasPrivateAddress() bool`

HasPrivateAddress returns a boolean if a field has been set.

### GetPrivateHostname

`func (o *GetGridCloudapiVmaddressResponse) GetPrivateHostname() string`

GetPrivateHostname returns the PrivateHostname field if non-nil, zero value otherwise.

### GetPrivateHostnameOk

`func (o *GetGridCloudapiVmaddressResponse) GetPrivateHostnameOk() (*string, bool)`

GetPrivateHostnameOk returns a tuple with the PrivateHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateHostname

`func (o *GetGridCloudapiVmaddressResponse) SetPrivateHostname(v string)`

SetPrivateHostname sets PrivateHostname field to given value.

### HasPrivateHostname

`func (o *GetGridCloudapiVmaddressResponse) HasPrivateHostname() bool`

HasPrivateHostname returns a boolean if a field has been set.

### GetPublicAddress

`func (o *GetGridCloudapiVmaddressResponse) GetPublicAddress() string`

GetPublicAddress returns the PublicAddress field if non-nil, zero value otherwise.

### GetPublicAddressOk

`func (o *GetGridCloudapiVmaddressResponse) GetPublicAddressOk() (*string, bool)`

GetPublicAddressOk returns a tuple with the PublicAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAddress

`func (o *GetGridCloudapiVmaddressResponse) SetPublicAddress(v string)`

SetPublicAddress sets PublicAddress field to given value.

### HasPublicAddress

`func (o *GetGridCloudapiVmaddressResponse) HasPublicAddress() bool`

HasPublicAddress returns a boolean if a field has been set.

### GetPublicHostname

`func (o *GetGridCloudapiVmaddressResponse) GetPublicHostname() string`

GetPublicHostname returns the PublicHostname field if non-nil, zero value otherwise.

### GetPublicHostnameOk

`func (o *GetGridCloudapiVmaddressResponse) GetPublicHostnameOk() (*string, bool)`

GetPublicHostnameOk returns a tuple with the PublicHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicHostname

`func (o *GetGridCloudapiVmaddressResponse) SetPublicHostname(v string)`

SetPublicHostname sets PublicHostname field to given value.

### HasPublicHostname

`func (o *GetGridCloudapiVmaddressResponse) HasPublicHostname() bool`

HasPublicHostname returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *GetGridCloudapiVmaddressResponse) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *GetGridCloudapiVmaddressResponse) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *GetGridCloudapiVmaddressResponse) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *GetGridCloudapiVmaddressResponse) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetSubnetCidr

`func (o *GetGridCloudapiVmaddressResponse) GetSubnetCidr() int64`

GetSubnetCidr returns the SubnetCidr field if non-nil, zero value otherwise.

### GetSubnetCidrOk

`func (o *GetGridCloudapiVmaddressResponse) GetSubnetCidrOk() (*int64, bool)`

GetSubnetCidrOk returns a tuple with the SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidr

`func (o *GetGridCloudapiVmaddressResponse) SetSubnetCidr(v int64)`

SetSubnetCidr sets SubnetCidr field to given value.

### HasSubnetCidr

`func (o *GetGridCloudapiVmaddressResponse) HasSubnetCidr() bool`

HasSubnetCidr returns a boolean if a field has been set.

### GetSubnetId

`func (o *GetGridCloudapiVmaddressResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *GetGridCloudapiVmaddressResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *GetGridCloudapiVmaddressResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *GetGridCloudapiVmaddressResponse) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetTenant

`func (o *GetGridCloudapiVmaddressResponse) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *GetGridCloudapiVmaddressResponse) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *GetGridCloudapiVmaddressResponse) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *GetGridCloudapiVmaddressResponse) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetUuid

`func (o *GetGridCloudapiVmaddressResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGridCloudapiVmaddressResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGridCloudapiVmaddressResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGridCloudapiVmaddressResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVmAvailabilityZone

`func (o *GetGridCloudapiVmaddressResponse) GetVmAvailabilityZone() string`

GetVmAvailabilityZone returns the VmAvailabilityZone field if non-nil, zero value otherwise.

### GetVmAvailabilityZoneOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmAvailabilityZoneOk() (*string, bool)`

GetVmAvailabilityZoneOk returns a tuple with the VmAvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmAvailabilityZone

`func (o *GetGridCloudapiVmaddressResponse) SetVmAvailabilityZone(v string)`

SetVmAvailabilityZone sets VmAvailabilityZone field to given value.

### HasVmAvailabilityZone

`func (o *GetGridCloudapiVmaddressResponse) HasVmAvailabilityZone() bool`

HasVmAvailabilityZone returns a boolean if a field has been set.

### GetVmComment

`func (o *GetGridCloudapiVmaddressResponse) GetVmComment() string`

GetVmComment returns the VmComment field if non-nil, zero value otherwise.

### GetVmCommentOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmCommentOk() (*string, bool)`

GetVmCommentOk returns a tuple with the VmComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmComment

`func (o *GetGridCloudapiVmaddressResponse) SetVmComment(v string)`

SetVmComment sets VmComment field to given value.

### HasVmComment

`func (o *GetGridCloudapiVmaddressResponse) HasVmComment() bool`

HasVmComment returns a boolean if a field has been set.

### GetVmCreationTime

`func (o *GetGridCloudapiVmaddressResponse) GetVmCreationTime() int64`

GetVmCreationTime returns the VmCreationTime field if non-nil, zero value otherwise.

### GetVmCreationTimeOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmCreationTimeOk() (*int64, bool)`

GetVmCreationTimeOk returns a tuple with the VmCreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmCreationTime

`func (o *GetGridCloudapiVmaddressResponse) SetVmCreationTime(v int64)`

SetVmCreationTime sets VmCreationTime field to given value.

### HasVmCreationTime

`func (o *GetGridCloudapiVmaddressResponse) HasVmCreationTime() bool`

HasVmCreationTime returns a boolean if a field has been set.

### GetVmHostname

`func (o *GetGridCloudapiVmaddressResponse) GetVmHostname() string`

GetVmHostname returns the VmHostname field if non-nil, zero value otherwise.

### GetVmHostnameOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmHostnameOk() (*string, bool)`

GetVmHostnameOk returns a tuple with the VmHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHostname

`func (o *GetGridCloudapiVmaddressResponse) SetVmHostname(v string)`

SetVmHostname sets VmHostname field to given value.

### HasVmHostname

`func (o *GetGridCloudapiVmaddressResponse) HasVmHostname() bool`

HasVmHostname returns a boolean if a field has been set.

### GetVmId

`func (o *GetGridCloudapiVmaddressResponse) GetVmId() string`

GetVmId returns the VmId field if non-nil, zero value otherwise.

### GetVmIdOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmIdOk() (*string, bool)`

GetVmIdOk returns a tuple with the VmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmId

`func (o *GetGridCloudapiVmaddressResponse) SetVmId(v string)`

SetVmId sets VmId field to given value.

### HasVmId

`func (o *GetGridCloudapiVmaddressResponse) HasVmId() bool`

HasVmId returns a boolean if a field has been set.

### GetVmKernelId

`func (o *GetGridCloudapiVmaddressResponse) GetVmKernelId() string`

GetVmKernelId returns the VmKernelId field if non-nil, zero value otherwise.

### GetVmKernelIdOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmKernelIdOk() (*string, bool)`

GetVmKernelIdOk returns a tuple with the VmKernelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmKernelId

`func (o *GetGridCloudapiVmaddressResponse) SetVmKernelId(v string)`

SetVmKernelId sets VmKernelId field to given value.

### HasVmKernelId

`func (o *GetGridCloudapiVmaddressResponse) HasVmKernelId() bool`

HasVmKernelId returns a boolean if a field has been set.

### GetVmLastUpdateTime

`func (o *GetGridCloudapiVmaddressResponse) GetVmLastUpdateTime() int64`

GetVmLastUpdateTime returns the VmLastUpdateTime field if non-nil, zero value otherwise.

### GetVmLastUpdateTimeOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmLastUpdateTimeOk() (*int64, bool)`

GetVmLastUpdateTimeOk returns a tuple with the VmLastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmLastUpdateTime

`func (o *GetGridCloudapiVmaddressResponse) SetVmLastUpdateTime(v int64)`

SetVmLastUpdateTime sets VmLastUpdateTime field to given value.

### HasVmLastUpdateTime

`func (o *GetGridCloudapiVmaddressResponse) HasVmLastUpdateTime() bool`

HasVmLastUpdateTime returns a boolean if a field has been set.

### GetVmName

`func (o *GetGridCloudapiVmaddressResponse) GetVmName() string`

GetVmName returns the VmName field if non-nil, zero value otherwise.

### GetVmNameOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmNameOk() (*string, bool)`

GetVmNameOk returns a tuple with the VmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmName

`func (o *GetGridCloudapiVmaddressResponse) SetVmName(v string)`

SetVmName sets VmName field to given value.

### HasVmName

`func (o *GetGridCloudapiVmaddressResponse) HasVmName() bool`

HasVmName returns a boolean if a field has been set.

### GetVmNetworkCount

`func (o *GetGridCloudapiVmaddressResponse) GetVmNetworkCount() int64`

GetVmNetworkCount returns the VmNetworkCount field if non-nil, zero value otherwise.

### GetVmNetworkCountOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmNetworkCountOk() (*int64, bool)`

GetVmNetworkCountOk returns a tuple with the VmNetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmNetworkCount

`func (o *GetGridCloudapiVmaddressResponse) SetVmNetworkCount(v int64)`

SetVmNetworkCount sets VmNetworkCount field to given value.

### HasVmNetworkCount

`func (o *GetGridCloudapiVmaddressResponse) HasVmNetworkCount() bool`

HasVmNetworkCount returns a boolean if a field has been set.

### GetVmOperatingSystem

`func (o *GetGridCloudapiVmaddressResponse) GetVmOperatingSystem() string`

GetVmOperatingSystem returns the VmOperatingSystem field if non-nil, zero value otherwise.

### GetVmOperatingSystemOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmOperatingSystemOk() (*string, bool)`

GetVmOperatingSystemOk returns a tuple with the VmOperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmOperatingSystem

`func (o *GetGridCloudapiVmaddressResponse) SetVmOperatingSystem(v string)`

SetVmOperatingSystem sets VmOperatingSystem field to given value.

### HasVmOperatingSystem

`func (o *GetGridCloudapiVmaddressResponse) HasVmOperatingSystem() bool`

HasVmOperatingSystem returns a boolean if a field has been set.

### GetVmType

`func (o *GetGridCloudapiVmaddressResponse) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmTypeOk() (*string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmType

`func (o *GetGridCloudapiVmaddressResponse) SetVmType(v string)`

SetVmType sets VmType field to given value.

### HasVmType

`func (o *GetGridCloudapiVmaddressResponse) HasVmType() bool`

HasVmType returns a boolean if a field has been set.

### GetVmVpcAddress

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcAddress() string`

GetVmVpcAddress returns the VmVpcAddress field if non-nil, zero value otherwise.

### GetVmVpcAddressOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcAddressOk() (*string, bool)`

GetVmVpcAddressOk returns a tuple with the VmVpcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVpcAddress

`func (o *GetGridCloudapiVmaddressResponse) SetVmVpcAddress(v string)`

SetVmVpcAddress sets VmVpcAddress field to given value.

### HasVmVpcAddress

`func (o *GetGridCloudapiVmaddressResponse) HasVmVpcAddress() bool`

HasVmVpcAddress returns a boolean if a field has been set.

### GetVmVpcCidr

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcCidr() int64`

GetVmVpcCidr returns the VmVpcCidr field if non-nil, zero value otherwise.

### GetVmVpcCidrOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcCidrOk() (*int64, bool)`

GetVmVpcCidrOk returns a tuple with the VmVpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVpcCidr

`func (o *GetGridCloudapiVmaddressResponse) SetVmVpcCidr(v int64)`

SetVmVpcCidr sets VmVpcCidr field to given value.

### HasVmVpcCidr

`func (o *GetGridCloudapiVmaddressResponse) HasVmVpcCidr() bool`

HasVmVpcCidr returns a boolean if a field has been set.

### GetVmVpcId

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcId() string`

GetVmVpcId returns the VmVpcId field if non-nil, zero value otherwise.

### GetVmVpcIdOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcIdOk() (*string, bool)`

GetVmVpcIdOk returns a tuple with the VmVpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVpcId

`func (o *GetGridCloudapiVmaddressResponse) SetVmVpcId(v string)`

SetVmVpcId sets VmVpcId field to given value.

### HasVmVpcId

`func (o *GetGridCloudapiVmaddressResponse) HasVmVpcId() bool`

HasVmVpcId returns a boolean if a field has been set.

### GetVmVpcName

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcName() string`

GetVmVpcName returns the VmVpcName field if non-nil, zero value otherwise.

### GetVmVpcNameOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcNameOk() (*string, bool)`

GetVmVpcNameOk returns a tuple with the VmVpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVpcName

`func (o *GetGridCloudapiVmaddressResponse) SetVmVpcName(v string)`

SetVmVpcName sets VmVpcName field to given value.

### HasVmVpcName

`func (o *GetGridCloudapiVmaddressResponse) HasVmVpcName() bool`

HasVmVpcName returns a boolean if a field has been set.

### GetVmVpcRef

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcRef() string`

GetVmVpcRef returns the VmVpcRef field if non-nil, zero value otherwise.

### GetVmVpcRefOk

`func (o *GetGridCloudapiVmaddressResponse) GetVmVpcRefOk() (*string, bool)`

GetVmVpcRefOk returns a tuple with the VmVpcRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmVpcRef

`func (o *GetGridCloudapiVmaddressResponse) SetVmVpcRef(v string)`

SetVmVpcRef sets VmVpcRef field to given value.

### HasVmVpcRef

`func (o *GetGridCloudapiVmaddressResponse) HasVmVpcRef() bool`

HasVmVpcRef returns a boolean if a field has been set.

### GetResult

`func (o *GetGridCloudapiVmaddressResponse) GetResult() GridCloudapiVmaddress`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridCloudapiVmaddressResponse) GetResultOk() (*GridCloudapiVmaddress, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridCloudapiVmaddressResponse) SetResult(v GridCloudapiVmaddress)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridCloudapiVmaddressResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


