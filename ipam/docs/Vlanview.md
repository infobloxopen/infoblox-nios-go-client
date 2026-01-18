# Vlanview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AllowRangeOverlapping** | Pointer to **bool** | When set to true VLAN Ranges under VLAN View can have overlapping ID. | [optional] 
**Comment** | Pointer to **string** | A descriptive comment for this VLAN View. | [optional] 
**EndVlanId** | Pointer to **int64** | End ID for VLAN View. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | Name of the VLAN View. | [optional] 
**PreCreateVlan** | Pointer to **bool** | If set on creation VLAN objects will be created once VLAN View created. | [optional] 
**StartVlanId** | Pointer to **int64** | Start ID for VLAN View. | [optional] 
**VlanNamePrefix** | Pointer to **string** | If set on creation prefix string will be used for VLAN name. | [optional] 

## Methods

### NewVlanview

`func NewVlanview() *Vlanview`

NewVlanview instantiates a new Vlanview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanviewWithDefaults

`func NewVlanviewWithDefaults() *Vlanview`

NewVlanviewWithDefaults instantiates a new Vlanview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Vlanview) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Vlanview) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Vlanview) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Vlanview) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAllowRangeOverlapping

`func (o *Vlanview) GetAllowRangeOverlapping() bool`

GetAllowRangeOverlapping returns the AllowRangeOverlapping field if non-nil, zero value otherwise.

### GetAllowRangeOverlappingOk

`func (o *Vlanview) GetAllowRangeOverlappingOk() (*bool, bool)`

GetAllowRangeOverlappingOk returns a tuple with the AllowRangeOverlapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRangeOverlapping

`func (o *Vlanview) SetAllowRangeOverlapping(v bool)`

SetAllowRangeOverlapping sets AllowRangeOverlapping field to given value.

### HasAllowRangeOverlapping

`func (o *Vlanview) HasAllowRangeOverlapping() bool`

HasAllowRangeOverlapping returns a boolean if a field has been set.

### GetComment

`func (o *Vlanview) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Vlanview) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Vlanview) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Vlanview) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEndVlanId

`func (o *Vlanview) GetEndVlanId() int64`

GetEndVlanId returns the EndVlanId field if non-nil, zero value otherwise.

### GetEndVlanIdOk

`func (o *Vlanview) GetEndVlanIdOk() (*int64, bool)`

GetEndVlanIdOk returns a tuple with the EndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndVlanId

`func (o *Vlanview) SetEndVlanId(v int64)`

SetEndVlanId sets EndVlanId field to given value.

### HasEndVlanId

`func (o *Vlanview) HasEndVlanId() bool`

HasEndVlanId returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Vlanview) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Vlanview) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Vlanview) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Vlanview) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Vlanview) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Vlanview) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Vlanview) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Vlanview) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Vlanview) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Vlanview) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Vlanview) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Vlanview) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *Vlanview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vlanview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vlanview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vlanview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreCreateVlan

`func (o *Vlanview) GetPreCreateVlan() bool`

GetPreCreateVlan returns the PreCreateVlan field if non-nil, zero value otherwise.

### GetPreCreateVlanOk

`func (o *Vlanview) GetPreCreateVlanOk() (*bool, bool)`

GetPreCreateVlanOk returns a tuple with the PreCreateVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCreateVlan

`func (o *Vlanview) SetPreCreateVlan(v bool)`

SetPreCreateVlan sets PreCreateVlan field to given value.

### HasPreCreateVlan

`func (o *Vlanview) HasPreCreateVlan() bool`

HasPreCreateVlan returns a boolean if a field has been set.

### GetStartVlanId

`func (o *Vlanview) GetStartVlanId() int64`

GetStartVlanId returns the StartVlanId field if non-nil, zero value otherwise.

### GetStartVlanIdOk

`func (o *Vlanview) GetStartVlanIdOk() (*int64, bool)`

GetStartVlanIdOk returns a tuple with the StartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartVlanId

`func (o *Vlanview) SetStartVlanId(v int64)`

SetStartVlanId sets StartVlanId field to given value.

### HasStartVlanId

`func (o *Vlanview) HasStartVlanId() bool`

HasStartVlanId returns a boolean if a field has been set.

### GetVlanNamePrefix

`func (o *Vlanview) GetVlanNamePrefix() string`

GetVlanNamePrefix returns the VlanNamePrefix field if non-nil, zero value otherwise.

### GetVlanNamePrefixOk

`func (o *Vlanview) GetVlanNamePrefixOk() (*string, bool)`

GetVlanNamePrefixOk returns a tuple with the VlanNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNamePrefix

`func (o *Vlanview) SetVlanNamePrefix(v string)`

SetVlanNamePrefix sets VlanNamePrefix field to given value.

### HasVlanNamePrefix

`func (o *Vlanview) HasVlanNamePrefix() bool`

HasVlanNamePrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


