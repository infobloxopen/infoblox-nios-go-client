# GetGridCloudapiVmResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AvailabilityZone** | Pointer to **string** | Availability zone of the VM. | [optional] [readonly] 
**CloudInfo** | Pointer to [**GridCloudapiVmCloudInfo**](GridCloudapiVmCloudInfo.md) |  | [optional] 
**Comment** | Pointer to **string** | Comment for the vm object; maximum 1024 characters. | [optional] 
**ElasticIpAddress** | Pointer to **string** | Elastic IP address associated with the VM&#39;s primary interface. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
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
**Result** | Pointer to [**GridCloudapiVm**](GridCloudapiVm.md) |  | [optional] 

## Methods

### NewGetGridCloudapiVmResponse

`func NewGetGridCloudapiVmResponse() *GetGridCloudapiVmResponse`

NewGetGridCloudapiVmResponse instantiates a new GetGridCloudapiVmResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridCloudapiVmResponseWithDefaults

`func NewGetGridCloudapiVmResponseWithDefaults() *GetGridCloudapiVmResponse`

NewGetGridCloudapiVmResponseWithDefaults instantiates a new GetGridCloudapiVmResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridCloudapiVmResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridCloudapiVmResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridCloudapiVmResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridCloudapiVmResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *GetGridCloudapiVmResponse) GetAvailabilityZone() string`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *GetGridCloudapiVmResponse) GetAvailabilityZoneOk() (*string, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *GetGridCloudapiVmResponse) SetAvailabilityZone(v string)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *GetGridCloudapiVmResponse) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetCloudInfo

`func (o *GetGridCloudapiVmResponse) GetCloudInfo() GridCloudapiVmCloudInfo`

GetCloudInfo returns the CloudInfo field if non-nil, zero value otherwise.

### GetCloudInfoOk

`func (o *GetGridCloudapiVmResponse) GetCloudInfoOk() (*GridCloudapiVmCloudInfo, bool)`

GetCloudInfoOk returns a tuple with the CloudInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInfo

`func (o *GetGridCloudapiVmResponse) SetCloudInfo(v GridCloudapiVmCloudInfo)`

SetCloudInfo sets CloudInfo field to given value.

### HasCloudInfo

`func (o *GetGridCloudapiVmResponse) HasCloudInfo() bool`

HasCloudInfo returns a boolean if a field has been set.

### GetComment

`func (o *GetGridCloudapiVmResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetGridCloudapiVmResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetGridCloudapiVmResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetGridCloudapiVmResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetElasticIpAddress

`func (o *GetGridCloudapiVmResponse) GetElasticIpAddress() string`

GetElasticIpAddress returns the ElasticIpAddress field if non-nil, zero value otherwise.

### GetElasticIpAddressOk

`func (o *GetGridCloudapiVmResponse) GetElasticIpAddressOk() (*string, bool)`

GetElasticIpAddressOk returns a tuple with the ElasticIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElasticIpAddress

`func (o *GetGridCloudapiVmResponse) SetElasticIpAddress(v string)`

SetElasticIpAddress sets ElasticIpAddress field to given value.

### HasElasticIpAddress

`func (o *GetGridCloudapiVmResponse) HasElasticIpAddress() bool`

HasElasticIpAddress returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetGridCloudapiVmResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetGridCloudapiVmResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetGridCloudapiVmResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetGridCloudapiVmResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetGridCloudapiVmResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetGridCloudapiVmResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetGridCloudapiVmResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetGridCloudapiVmResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetGridCloudapiVmResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetGridCloudapiVmResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetGridCloudapiVmResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetGridCloudapiVmResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFirstSeen

`func (o *GetGridCloudapiVmResponse) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *GetGridCloudapiVmResponse) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *GetGridCloudapiVmResponse) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *GetGridCloudapiVmResponse) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetHostname

`func (o *GetGridCloudapiVmResponse) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetGridCloudapiVmResponse) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetGridCloudapiVmResponse) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetGridCloudapiVmResponse) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *GetGridCloudapiVmResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGridCloudapiVmResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGridCloudapiVmResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetGridCloudapiVmResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKernelId

`func (o *GetGridCloudapiVmResponse) GetKernelId() string`

GetKernelId returns the KernelId field if non-nil, zero value otherwise.

### GetKernelIdOk

`func (o *GetGridCloudapiVmResponse) GetKernelIdOk() (*string, bool)`

GetKernelIdOk returns a tuple with the KernelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernelId

`func (o *GetGridCloudapiVmResponse) SetKernelId(v string)`

SetKernelId sets KernelId field to given value.

### HasKernelId

`func (o *GetGridCloudapiVmResponse) HasKernelId() bool`

HasKernelId returns a boolean if a field has been set.

### GetLastSeen

`func (o *GetGridCloudapiVmResponse) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GetGridCloudapiVmResponse) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GetGridCloudapiVmResponse) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GetGridCloudapiVmResponse) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetName

`func (o *GetGridCloudapiVmResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGridCloudapiVmResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGridCloudapiVmResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGridCloudapiVmResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkCount

`func (o *GetGridCloudapiVmResponse) GetNetworkCount() int64`

GetNetworkCount returns the NetworkCount field if non-nil, zero value otherwise.

### GetNetworkCountOk

`func (o *GetGridCloudapiVmResponse) GetNetworkCountOk() (*int64, bool)`

GetNetworkCountOk returns a tuple with the NetworkCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCount

`func (o *GetGridCloudapiVmResponse) SetNetworkCount(v int64)`

SetNetworkCount sets NetworkCount field to given value.

### HasNetworkCount

`func (o *GetGridCloudapiVmResponse) HasNetworkCount() bool`

HasNetworkCount returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *GetGridCloudapiVmResponse) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *GetGridCloudapiVmResponse) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *GetGridCloudapiVmResponse) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *GetGridCloudapiVmResponse) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetPrimaryMacAddress

`func (o *GetGridCloudapiVmResponse) GetPrimaryMacAddress() string`

GetPrimaryMacAddress returns the PrimaryMacAddress field if non-nil, zero value otherwise.

### GetPrimaryMacAddressOk

`func (o *GetGridCloudapiVmResponse) GetPrimaryMacAddressOk() (*string, bool)`

GetPrimaryMacAddressOk returns a tuple with the PrimaryMacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryMacAddress

`func (o *GetGridCloudapiVmResponse) SetPrimaryMacAddress(v string)`

SetPrimaryMacAddress sets PrimaryMacAddress field to given value.

### HasPrimaryMacAddress

`func (o *GetGridCloudapiVmResponse) HasPrimaryMacAddress() bool`

HasPrimaryMacAddress returns a boolean if a field has been set.

### GetSubnetAddress

`func (o *GetGridCloudapiVmResponse) GetSubnetAddress() string`

GetSubnetAddress returns the SubnetAddress field if non-nil, zero value otherwise.

### GetSubnetAddressOk

`func (o *GetGridCloudapiVmResponse) GetSubnetAddressOk() (*string, bool)`

GetSubnetAddressOk returns a tuple with the SubnetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetAddress

`func (o *GetGridCloudapiVmResponse) SetSubnetAddress(v string)`

SetSubnetAddress sets SubnetAddress field to given value.

### HasSubnetAddress

`func (o *GetGridCloudapiVmResponse) HasSubnetAddress() bool`

HasSubnetAddress returns a boolean if a field has been set.

### GetSubnetCidr

`func (o *GetGridCloudapiVmResponse) GetSubnetCidr() int64`

GetSubnetCidr returns the SubnetCidr field if non-nil, zero value otherwise.

### GetSubnetCidrOk

`func (o *GetGridCloudapiVmResponse) GetSubnetCidrOk() (*int64, bool)`

GetSubnetCidrOk returns a tuple with the SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidr

`func (o *GetGridCloudapiVmResponse) SetSubnetCidr(v int64)`

SetSubnetCidr sets SubnetCidr field to given value.

### HasSubnetCidr

`func (o *GetGridCloudapiVmResponse) HasSubnetCidr() bool`

HasSubnetCidr returns a boolean if a field has been set.

### GetSubnetId

`func (o *GetGridCloudapiVmResponse) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *GetGridCloudapiVmResponse) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *GetGridCloudapiVmResponse) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *GetGridCloudapiVmResponse) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetTenantName

`func (o *GetGridCloudapiVmResponse) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *GetGridCloudapiVmResponse) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *GetGridCloudapiVmResponse) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *GetGridCloudapiVmResponse) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetVmType

`func (o *GetGridCloudapiVmResponse) GetVmType() string`

GetVmType returns the VmType field if non-nil, zero value otherwise.

### GetVmTypeOk

`func (o *GetGridCloudapiVmResponse) GetVmTypeOk() (*string, bool)`

GetVmTypeOk returns a tuple with the VmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmType

`func (o *GetGridCloudapiVmResponse) SetVmType(v string)`

SetVmType sets VmType field to given value.

### HasVmType

`func (o *GetGridCloudapiVmResponse) HasVmType() bool`

HasVmType returns a boolean if a field has been set.

### GetVpcAddress

`func (o *GetGridCloudapiVmResponse) GetVpcAddress() string`

GetVpcAddress returns the VpcAddress field if non-nil, zero value otherwise.

### GetVpcAddressOk

`func (o *GetGridCloudapiVmResponse) GetVpcAddressOk() (*string, bool)`

GetVpcAddressOk returns a tuple with the VpcAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcAddress

`func (o *GetGridCloudapiVmResponse) SetVpcAddress(v string)`

SetVpcAddress sets VpcAddress field to given value.

### HasVpcAddress

`func (o *GetGridCloudapiVmResponse) HasVpcAddress() bool`

HasVpcAddress returns a boolean if a field has been set.

### GetVpcCidr

`func (o *GetGridCloudapiVmResponse) GetVpcCidr() int64`

GetVpcCidr returns the VpcCidr field if non-nil, zero value otherwise.

### GetVpcCidrOk

`func (o *GetGridCloudapiVmResponse) GetVpcCidrOk() (*int64, bool)`

GetVpcCidrOk returns a tuple with the VpcCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcCidr

`func (o *GetGridCloudapiVmResponse) SetVpcCidr(v int64)`

SetVpcCidr sets VpcCidr field to given value.

### HasVpcCidr

`func (o *GetGridCloudapiVmResponse) HasVpcCidr() bool`

HasVpcCidr returns a boolean if a field has been set.

### GetVpcId

`func (o *GetGridCloudapiVmResponse) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *GetGridCloudapiVmResponse) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *GetGridCloudapiVmResponse) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *GetGridCloudapiVmResponse) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### GetVpcName

`func (o *GetGridCloudapiVmResponse) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *GetGridCloudapiVmResponse) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *GetGridCloudapiVmResponse) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *GetGridCloudapiVmResponse) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### GetResult

`func (o *GetGridCloudapiVmResponse) GetResult() GridCloudapiVm`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridCloudapiVmResponse) GetResultOk() (*GridCloudapiVm, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridCloudapiVmResponse) SetResult(v GridCloudapiVm)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridCloudapiVmResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


