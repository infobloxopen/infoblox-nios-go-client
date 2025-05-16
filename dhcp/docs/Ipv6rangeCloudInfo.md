# Ipv6rangeCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedMember** | Pointer to [**Ipv6rangecloudinfoDelegatedMember**](Ipv6rangecloudinfoDelegatedMember.md) |  | [optional] 
**DelegatedScope** | Pointer to **string** | Indicates the scope of delegation for the object. This can be one of the following: NONE (outside any delegation), ROOT (the delegation point), SUBTREE (within the scope of a delegation), RECLAIMING (within the scope of a delegation being reclaimed, either as the delegation point or in the subtree). | [optional] [readonly] 
**DelegatedRoot** | Pointer to **string** | Indicates the root of the delegation if delegated_scope is SUBTREE or RECLAIMING. This is not set otherwise. | [optional] [readonly] 
**OwnedByAdaptor** | Pointer to **bool** | Determines whether the object was created by the cloud adapter or not. | [optional] [readonly] 
**Usage** | Pointer to **string** | Indicates the cloud origin of the object. | [optional] [readonly] 
**Tenant** | Pointer to **string** | Reference to the tenant object associated with the object, if any. | [optional] [readonly] 
**MgmtPlatform** | Pointer to **string** | Indicates the specified cloud management platform. | [optional] [readonly] 
**AuthorityType** | Pointer to **string** | Type of authority over the object. | [optional] [readonly] 

## Methods

### NewIpv6rangeCloudInfo

`func NewIpv6rangeCloudInfo() *Ipv6rangeCloudInfo`

NewIpv6rangeCloudInfo instantiates a new Ipv6rangeCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6rangeCloudInfoWithDefaults

`func NewIpv6rangeCloudInfoWithDefaults() *Ipv6rangeCloudInfo`

NewIpv6rangeCloudInfoWithDefaults instantiates a new Ipv6rangeCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedMember

`func (o *Ipv6rangeCloudInfo) GetDelegatedMember() Ipv6rangecloudinfoDelegatedMember`

GetDelegatedMember returns the DelegatedMember field if non-nil, zero value otherwise.

### GetDelegatedMemberOk

`func (o *Ipv6rangeCloudInfo) GetDelegatedMemberOk() (*Ipv6rangecloudinfoDelegatedMember, bool)`

GetDelegatedMemberOk returns a tuple with the DelegatedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedMember

`func (o *Ipv6rangeCloudInfo) SetDelegatedMember(v Ipv6rangecloudinfoDelegatedMember)`

SetDelegatedMember sets DelegatedMember field to given value.

### HasDelegatedMember

`func (o *Ipv6rangeCloudInfo) HasDelegatedMember() bool`

HasDelegatedMember returns a boolean if a field has been set.

### GetDelegatedScope

`func (o *Ipv6rangeCloudInfo) GetDelegatedScope() string`

GetDelegatedScope returns the DelegatedScope field if non-nil, zero value otherwise.

### GetDelegatedScopeOk

`func (o *Ipv6rangeCloudInfo) GetDelegatedScopeOk() (*string, bool)`

GetDelegatedScopeOk returns a tuple with the DelegatedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedScope

`func (o *Ipv6rangeCloudInfo) SetDelegatedScope(v string)`

SetDelegatedScope sets DelegatedScope field to given value.

### HasDelegatedScope

`func (o *Ipv6rangeCloudInfo) HasDelegatedScope() bool`

HasDelegatedScope returns a boolean if a field has been set.

### GetDelegatedRoot

`func (o *Ipv6rangeCloudInfo) GetDelegatedRoot() string`

GetDelegatedRoot returns the DelegatedRoot field if non-nil, zero value otherwise.

### GetDelegatedRootOk

`func (o *Ipv6rangeCloudInfo) GetDelegatedRootOk() (*string, bool)`

GetDelegatedRootOk returns a tuple with the DelegatedRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedRoot

`func (o *Ipv6rangeCloudInfo) SetDelegatedRoot(v string)`

SetDelegatedRoot sets DelegatedRoot field to given value.

### HasDelegatedRoot

`func (o *Ipv6rangeCloudInfo) HasDelegatedRoot() bool`

HasDelegatedRoot returns a boolean if a field has been set.

### GetOwnedByAdaptor

`func (o *Ipv6rangeCloudInfo) GetOwnedByAdaptor() bool`

GetOwnedByAdaptor returns the OwnedByAdaptor field if non-nil, zero value otherwise.

### GetOwnedByAdaptorOk

`func (o *Ipv6rangeCloudInfo) GetOwnedByAdaptorOk() (*bool, bool)`

GetOwnedByAdaptorOk returns a tuple with the OwnedByAdaptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedByAdaptor

`func (o *Ipv6rangeCloudInfo) SetOwnedByAdaptor(v bool)`

SetOwnedByAdaptor sets OwnedByAdaptor field to given value.

### HasOwnedByAdaptor

`func (o *Ipv6rangeCloudInfo) HasOwnedByAdaptor() bool`

HasOwnedByAdaptor returns a boolean if a field has been set.

### GetUsage

`func (o *Ipv6rangeCloudInfo) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *Ipv6rangeCloudInfo) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *Ipv6rangeCloudInfo) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *Ipv6rangeCloudInfo) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetTenant

`func (o *Ipv6rangeCloudInfo) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *Ipv6rangeCloudInfo) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *Ipv6rangeCloudInfo) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *Ipv6rangeCloudInfo) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetMgmtPlatform

`func (o *Ipv6rangeCloudInfo) GetMgmtPlatform() string`

GetMgmtPlatform returns the MgmtPlatform field if non-nil, zero value otherwise.

### GetMgmtPlatformOk

`func (o *Ipv6rangeCloudInfo) GetMgmtPlatformOk() (*string, bool)`

GetMgmtPlatformOk returns a tuple with the MgmtPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPlatform

`func (o *Ipv6rangeCloudInfo) SetMgmtPlatform(v string)`

SetMgmtPlatform sets MgmtPlatform field to given value.

### HasMgmtPlatform

`func (o *Ipv6rangeCloudInfo) HasMgmtPlatform() bool`

HasMgmtPlatform returns a boolean if a field has been set.

### GetAuthorityType

`func (o *Ipv6rangeCloudInfo) GetAuthorityType() string`

GetAuthorityType returns the AuthorityType field if non-nil, zero value otherwise.

### GetAuthorityTypeOk

`func (o *Ipv6rangeCloudInfo) GetAuthorityTypeOk() (*string, bool)`

GetAuthorityTypeOk returns a tuple with the AuthorityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityType

`func (o *Ipv6rangeCloudInfo) SetAuthorityType(v string)`

SetAuthorityType sets AuthorityType field to given value.

### HasAuthorityType

`func (o *Ipv6rangeCloudInfo) HasAuthorityType() bool`

HasAuthorityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


