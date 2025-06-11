# ViewCloudInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DelegatedMember** | Pointer to [**ViewcloudinfoDelegatedMember**](ViewcloudinfoDelegatedMember.md) |  | [optional] 
**DelegatedScope** | Pointer to **string** | Indicates the scope of delegation for the object. This can be one of the following: NONE (outside any delegation), ROOT (the delegation point), SUBTREE (within the scope of a delegation), RECLAIMING (within the scope of a delegation being reclaimed, either as the delegation point or in the subtree). | [optional] [readonly] 
**DelegatedRoot** | Pointer to **string** | Indicates the root of the delegation if delegated_scope is SUBTREE or RECLAIMING. This is not set otherwise. | [optional] [readonly] 
**OwnedByAdaptor** | Pointer to **bool** | Determines whether the object was created by the cloud adapter or not. | [optional] [readonly] 
**Usage** | Pointer to **string** | Indicates the cloud origin of the object. | [optional] [readonly] 
**Tenant** | Pointer to **string** | Reference to the tenant object associated with the object, if any. | [optional] [readonly] 
**MgmtPlatform** | Pointer to **string** | Indicates the specified cloud management platform. | [optional] [readonly] 
**AuthorityType** | Pointer to **string** | Type of authority over the object. | [optional] [readonly] 

## Methods

### NewViewCloudInfo

`func NewViewCloudInfo() *ViewCloudInfo`

NewViewCloudInfo instantiates a new ViewCloudInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewCloudInfoWithDefaults

`func NewViewCloudInfoWithDefaults() *ViewCloudInfo`

NewViewCloudInfoWithDefaults instantiates a new ViewCloudInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelegatedMember

`func (o *ViewCloudInfo) GetDelegatedMember() ViewcloudinfoDelegatedMember`

GetDelegatedMember returns the DelegatedMember field if non-nil, zero value otherwise.

### GetDelegatedMemberOk

`func (o *ViewCloudInfo) GetDelegatedMemberOk() (*ViewcloudinfoDelegatedMember, bool)`

GetDelegatedMemberOk returns a tuple with the DelegatedMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedMember

`func (o *ViewCloudInfo) SetDelegatedMember(v ViewcloudinfoDelegatedMember)`

SetDelegatedMember sets DelegatedMember field to given value.

### HasDelegatedMember

`func (o *ViewCloudInfo) HasDelegatedMember() bool`

HasDelegatedMember returns a boolean if a field has been set.

### GetDelegatedScope

`func (o *ViewCloudInfo) GetDelegatedScope() string`

GetDelegatedScope returns the DelegatedScope field if non-nil, zero value otherwise.

### GetDelegatedScopeOk

`func (o *ViewCloudInfo) GetDelegatedScopeOk() (*string, bool)`

GetDelegatedScopeOk returns a tuple with the DelegatedScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedScope

`func (o *ViewCloudInfo) SetDelegatedScope(v string)`

SetDelegatedScope sets DelegatedScope field to given value.

### HasDelegatedScope

`func (o *ViewCloudInfo) HasDelegatedScope() bool`

HasDelegatedScope returns a boolean if a field has been set.

### GetDelegatedRoot

`func (o *ViewCloudInfo) GetDelegatedRoot() string`

GetDelegatedRoot returns the DelegatedRoot field if non-nil, zero value otherwise.

### GetDelegatedRootOk

`func (o *ViewCloudInfo) GetDelegatedRootOk() (*string, bool)`

GetDelegatedRootOk returns a tuple with the DelegatedRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedRoot

`func (o *ViewCloudInfo) SetDelegatedRoot(v string)`

SetDelegatedRoot sets DelegatedRoot field to given value.

### HasDelegatedRoot

`func (o *ViewCloudInfo) HasDelegatedRoot() bool`

HasDelegatedRoot returns a boolean if a field has been set.

### GetOwnedByAdaptor

`func (o *ViewCloudInfo) GetOwnedByAdaptor() bool`

GetOwnedByAdaptor returns the OwnedByAdaptor field if non-nil, zero value otherwise.

### GetOwnedByAdaptorOk

`func (o *ViewCloudInfo) GetOwnedByAdaptorOk() (*bool, bool)`

GetOwnedByAdaptorOk returns a tuple with the OwnedByAdaptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedByAdaptor

`func (o *ViewCloudInfo) SetOwnedByAdaptor(v bool)`

SetOwnedByAdaptor sets OwnedByAdaptor field to given value.

### HasOwnedByAdaptor

`func (o *ViewCloudInfo) HasOwnedByAdaptor() bool`

HasOwnedByAdaptor returns a boolean if a field has been set.

### GetUsage

`func (o *ViewCloudInfo) GetUsage() string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ViewCloudInfo) GetUsageOk() (*string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ViewCloudInfo) SetUsage(v string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ViewCloudInfo) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetTenant

`func (o *ViewCloudInfo) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *ViewCloudInfo) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *ViewCloudInfo) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *ViewCloudInfo) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetMgmtPlatform

`func (o *ViewCloudInfo) GetMgmtPlatform() string`

GetMgmtPlatform returns the MgmtPlatform field if non-nil, zero value otherwise.

### GetMgmtPlatformOk

`func (o *ViewCloudInfo) GetMgmtPlatformOk() (*string, bool)`

GetMgmtPlatformOk returns a tuple with the MgmtPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMgmtPlatform

`func (o *ViewCloudInfo) SetMgmtPlatform(v string)`

SetMgmtPlatform sets MgmtPlatform field to given value.

### HasMgmtPlatform

`func (o *ViewCloudInfo) HasMgmtPlatform() bool`

HasMgmtPlatform returns a boolean if a field has been set.

### GetAuthorityType

`func (o *ViewCloudInfo) GetAuthorityType() string`

GetAuthorityType returns the AuthorityType field if non-nil, zero value otherwise.

### GetAuthorityTypeOk

`func (o *ViewCloudInfo) GetAuthorityTypeOk() (*string, bool)`

GetAuthorityTypeOk returns a tuple with the AuthorityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorityType

`func (o *ViewCloudInfo) SetAuthorityType(v string)`

SetAuthorityType sets AuthorityType field to given value.

### HasAuthorityType

`func (o *ViewCloudInfo) HasAuthorityType() bool`

HasAuthorityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


