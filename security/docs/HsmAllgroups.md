# HsmAllgroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Groups** | Pointer to **[]string** | The list of HSM groups configured on the appliance. | [optional] 

## Methods

### NewHsmAllgroups

`func NewHsmAllgroups() *HsmAllgroups`

NewHsmAllgroups instantiates a new HsmAllgroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHsmAllgroupsWithDefaults

`func NewHsmAllgroupsWithDefaults() *HsmAllgroups`

NewHsmAllgroupsWithDefaults instantiates a new HsmAllgroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *HsmAllgroups) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *HsmAllgroups) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *HsmAllgroups) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *HsmAllgroups) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetGroups

`func (o *HsmAllgroups) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *HsmAllgroups) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *HsmAllgroups) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *HsmAllgroups) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


