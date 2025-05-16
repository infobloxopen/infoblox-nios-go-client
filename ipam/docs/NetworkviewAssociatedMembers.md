# NetworkviewAssociatedMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | Pointer to **string** | The member object associated with a network view. | [optional] [readonly] 
**Failovers** | Pointer to **[]string** | The list of failover objects associated with each member. | [optional] [readonly] 

## Methods

### NewNetworkviewAssociatedMembers

`func NewNetworkviewAssociatedMembers() *NetworkviewAssociatedMembers`

NewNetworkviewAssociatedMembers instantiates a new NetworkviewAssociatedMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkviewAssociatedMembersWithDefaults

`func NewNetworkviewAssociatedMembersWithDefaults() *NetworkviewAssociatedMembers`

NewNetworkviewAssociatedMembersWithDefaults instantiates a new NetworkviewAssociatedMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *NetworkviewAssociatedMembers) GetMember() string`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *NetworkviewAssociatedMembers) GetMemberOk() (*string, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *NetworkviewAssociatedMembers) SetMember(v string)`

SetMember sets Member field to given value.

### HasMember

`func (o *NetworkviewAssociatedMembers) HasMember() bool`

HasMember returns a boolean if a field has been set.

### GetFailovers

`func (o *NetworkviewAssociatedMembers) GetFailovers() []string`

GetFailovers returns the Failovers field if non-nil, zero value otherwise.

### GetFailoversOk

`func (o *NetworkviewAssociatedMembers) GetFailoversOk() (*[]string, bool)`

GetFailoversOk returns a tuple with the Failovers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailovers

`func (o *NetworkviewAssociatedMembers) SetFailovers(v []string)`

SetFailovers sets Failovers field to given value.

### HasFailovers

`func (o *NetworkviewAssociatedMembers) HasFailovers() bool`

HasFailovers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


