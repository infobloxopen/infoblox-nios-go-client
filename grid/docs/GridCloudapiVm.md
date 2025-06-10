# GridCloudapiVm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AvailabilityZone** | Pointer to **string** | Availability zone of the VM. | [optional] [readonly] 
**CloudInfo** | Pointer to [**GridCloudapiVmCloudInfo**](GridCloudapiVmCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the vm object; maximum 1024 characters. | [optional] 
**ElasticIpAddress** | Pointer to **string** | Elastic IP address associated with the VM&#39;s primary interface. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**FirstSeen** | Pointer to **int64** | The timestamp when the VM was first seen in the system. | [optional] [readonly] 
**Hostname** | Pointer to **string** | Hostname part of the FQDN for the address associated with the VM&#39;s primary interface. | [optional] [readonly] 
**Id** | Pointer to **string** | Unique ID associated with the VM. This is set only when the VM is first created. | [optional] [readonly] 
**KernelId** | Pointer to **string** | Identifier of the kernel that this VM is running; maximum 128 characters. | [optional] 
**LastSeen** | Pointer to **int64** | The timestamp when the last event associated with the VM happened. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the VM. | [optional] 
**NetworkCount** | Pointer to **int64** | Number of Networks containing any address associated with this VM. | [optional] [readonly] 
**OperatingSystem** | Pointer to **string** | Guest Operating system that this VM is running; maximum 128 characters. | [optional] 
**PrimaryMacAddress** | Pointer to **string** | MAC address associated with the VM&#39;s primary interface. | [optional] [readonly] 
**SubnetAddress** | Pointer to **string** | Address of the network that is the container of the address associated with the VM&#39;s primary interface. | [optional] [readonly] 
**SubnetCidr** | Pointer to **int64** | CIDR of the network that is the container of the address associated with the VM&#39;s primary interface. | [optional] [readonly] 
**SubnetId** | Pointer to **string** | Subnet ID of the network that is the container of the address associated with the VM&#39;s primary interface. | [optional] [readonly] 
**TenantName** | Pointer to **string** | Name of the tenant associated with the VM. | [optional] [readonly] 
**VmType** | Pointer to **string** | VM type; maximum 64 characters. | [optional] 
**VpcAddress** | Pointer to **string** | Network address of the parent VPC. | [optional] [readonly] 
**VpcCidr** | Pointer to **int64** | Network CIDR of the parent VPC. | [optional] [readonly] 
**VpcId** | Pointer to **string** | Identifier of the parent VPC. | [optional] [readonly] 
**VpcName** | Pointer to **string** | Name of the parent VPC. | [optional] [readonly] 

## Methods

### NewGridCloudapiVm

`func NewGridCloudapiVm() *GridCloudapiVm`

NewGridCloudapiVm instantiates a new GridCloudapiVm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridCloudapiVmWithDefaults

`func NewGridCloudapiVmWithDefaults() *GridCloudapiVm`

NewGridCloudapiVmWithDefaults instantiates a new GridCloudapiVm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridCloudapiVm) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridCloudapiVm) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridCloudapiVm) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridCloudapiVm) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *GridCloudapiVm) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *GridCloudapiVm) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *GridCloudapiVm) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *GridCloudapiVm) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GridCloudapiVm) GetCloudInfo() GridCloudapiVmCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GridCloudapiVm) GetCloudInfoOk() (*GridCloudapiVmCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GridCloudapiVm) SetCloudInfo(v GridCloudapiVmCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GridCloudapiVm) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GridCloudapiVm) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GridCloudapiVm) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GridCloudapiVm) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GridCloudapiVm) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetElasticIpAddress

`func (o *GridCloudapiVm) GetElasticIpAddress() string`

GetElasticIpAddress returns the ElasticIpAddress field if non-nil, zero value otherwise.

### GetElasticIpAddressOk

`func (o *GridCloudapiVm) GetElasticIpAddressOk() (*string, bool)`

GetElasticIpAddressOk returns a tuple with the ElasticIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticIpAddress

`func (o *GridCloudapiVm) SetElasticIpAddress(v string)`

SetElasticIpAddress sets ElasticIpAddress field to given value.

### HasElasticIpAddress

`func (o *GridCloudapiVm) HasElasticIpAddress() bool`

HasElasticIpAddress returns a boolean if a field has been set.

### GetExtattrs

`func (o *GridCloudapiVm) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GridCloudapiVm) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GridCloudapiVm) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GridCloudapiVm) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetFirstSeen

`func (o *GridCloudapiVm) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *GridCloudapiVm) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *GridCloudapiVm) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *GridCloudapiVm) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetHostname

`func (o *GridCloudapiVm) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GridCloudapiVm) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GridCloudapiVm) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GridCloudapiVm) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *GridCloudapiVm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GridCloudapiVm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GridCloudapiVm) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GridCloudapiVm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKernelId

`func (o *GridCloudapiVm) GetKernelId() string`

GetKernelId returns the KernelId field if non-nil, zero value otherwise.

### GetKernelIdOk

`func (o *GridCloudapiVm) GetKernelIdOk() (*string, bool)`

GetKernelIdOk returns a tuple with the KernelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelId

`func (o *GridCloudapiVm) SetKernelId(v string)`

SetKernelId sets KernelId field to given value.

### HasKernelId

`func (o *GridCloudapiVm) HasKernelId() bool`

HasKernelId returns a boolean if a field has been set.

### GetLastSeen

`func (o *GridCloudapiVm) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GridCloudapiVm) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GridCloudapiVm) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GridCloudapiVm) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetName

`func (o *GridCloudapiVm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GridCloudapiVm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GridCloudapiVm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GridCloudapiVm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkCount

`func (o *GridCloudapiVm) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *GridCloudapiVm) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *GridCloudapiVm) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *GridCloudapiVm) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *GridCloudapiVm) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *GridCloudapiVm) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *GridCloudapiVm) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *GridCloudapiVm) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPrimaryMacAddress

`func (o *GridCloudapiVm) GetPrimaryMacAddress() string`

GetPrimaryMacAddress returns the PrimaryMacAddress field if non-nil, zero value otherwise.

### GetPrimaryMacAddressOk

`func (o *GridCloudapiVm) GetPrimaryMacAddressOk() (*string, bool)`

GetPrimaryMacAddressOk returns a tuple with the PrimaryMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMacAddress

`func (o *GridCloudapiVm) SetPrimaryMacAddress(v string)`

SetPrimaryMacAddress sets PrimaryMacAddress field to given value.

### HasPrimaryMacAddress

`func (o *GridCloudapiVm) HasPrimaryMacAddress() bool`

HasPrimaryMacAddress returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *GridCloudapiVm) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *GridCloudapiVm) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *GridCloudapiVm) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *GridCloudapiVm) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetSubnetCidr

`func (o *GridCloudapiVm) GetSubnetCidr() int64`

GetSubnetCidr returns the SubnetCidr field if non-nil, zero value otherwise.

### GetSubnetCidrOk

`func (o *GridCloudapiVm) GetSubnetCidrOk() (*int64, bool)`

GetSubnetCidrOk returns a tuple with the SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidr

`func (o *GridCloudapiVm) SetSubnetCidr(v int64)`

SetSubnetCidr sets SubnetCidr field to given value.

### HasSubnetCidr

`func (o *GridCloudapiVm) HasSubnetCidr() bool`

HasSubnetCidr returns a boolean if a field has been set.

### GetSubnetId

`func (o *GridCloudapiVm) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *GridCloudapiVm) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *GridCloudapiVm) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *GridCloudapiVm) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetTenantName

`func (o *GridCloudapiVm) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *GridCloudapiVm) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *GridCloudapiVm) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *GridCloudapiVm) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetVmType

`func (o *GridCloudapiVm) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *GridCloudapiVm) GetVmTypeOk() (*string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmType

`func (o *GridCloudapiVm) SetVmType(v string)`

SetVmType sets VmType field to given value.

### HasVmType

`func (o *GridCloudapiVm) HasVmType() bool`

HasVmType returns a boolean if a field has been set.

### GetVpcAddress

`func (o *GridCloudapiVm) GetVpcAddress() string`

GetVpcAddress returns the VpcAddress field if non-nil, zero value otherwise.

### GetVpcAddressOk

`func (o *GridCloudapiVm) GetVpcAddressOk() (*string, bool)`

GetVpcAddressOk returns a tuple with the VpcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcAddress

`func (o *GridCloudapiVm) SetVpcAddress(v string)`

SetVpcAddress sets VpcAddress field to given value.

### HasVpcAddress

`func (o *GridCloudapiVm) HasVpcAddress() bool`

HasVpcAddress returns a boolean if a field has been set.

### GetVpcCidr

`func (o *GridCloudapiVm) GetVpcCidr() int64`

GetVpcCidr returns the VpcCidr field if non-nil, zero value otherwise.

### GetVpcCidrOk

`func (o *GridCloudapiVm) GetVpcCidrOk() (*int64, bool)`

GetVpcCidrOk returns a tuple with the VpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidr

`func (o *GridCloudapiVm) SetVpcCidr(v int64)`

SetVpcCidr sets VpcCidr field to given value.

### HasVpcCidr

`func (o *GridCloudapiVm) HasVpcCidr() bool`

HasVpcCidr returns a boolean if a field has been set.

### GetVpcId

`func (o *GridCloudapiVm) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *GridCloudapiVm) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *GridCloudapiVm) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *GridCloudapiVm) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetVpcName

`func (o *GridCloudapiVm) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *GridCloudapiVm) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *GridCloudapiVm) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *GridCloudapiVm) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


