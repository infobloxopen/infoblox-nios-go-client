# Vlanrange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment for this VLAN Range. | [optional] 
**DeleteVlans** | Pointer to **bool** | Vlans delete option. Determines whether all child objects should be removed alongside with the VLAN Range or child objects should be assigned to another parental VLAN Range/View. By default child objects are re-parented. | [optional] 
**EndVlanId** | Pointer to **int64** | End ID for VLAN Range. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | Name of the VLAN Range. | [optional] 
**PreCreateVlan** | Pointer to **bool** | If set on creation VLAN objects will be created once VLAN Range created. | [optional] 
**StartVlanId** | Pointer to **int64** | Start ID for VLAN Range. | [optional] 
**VlanNamePrefix** | Pointer to **string** | If set on creation prefix string will be used for VLAN name. | [optional] 
**VlanView** | Pointer to **string** | The VLAN View to which this VLAN Range belongs. | [optional] 

## Methods

### NewVlanrange

`func NewVlanrange() *Vlanrange`

NewVlanrange instantiates a new Vlanrange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanrangeWithDefaults

`func NewVlanrangeWithDefaults() *Vlanrange`

NewVlanrangeWithDefaults instantiates a new Vlanrange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Vlanrange) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Vlanrange) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Vlanrange) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Vlanrange) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Vlanrange) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Vlanrange) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Vlanrange) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Vlanrange) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeleteVlans

`func (o *Vlanrange) GetDeleteVlans() bool`

GetDeleteVlans returns the DeleteVlans field if non-nil, zero value otherwise.

### GetDeleteVlansOk

`func (o *Vlanrange) GetDeleteVlansOk() (*bool, bool)`

GetDeleteVlansOk returns a tuple with the DeleteVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVlans

`func (o *Vlanrange) SetDeleteVlans(v bool)`

SetDeleteVlans sets DeleteVlans field to given value.

### HasDeleteVlans

`func (o *Vlanrange) HasDeleteVlans() bool`

HasDeleteVlans returns a boolean if a field has been set.

### GetEndVlanId

`func (o *Vlanrange) GetEndVlanId() int64`

GetEndVlanId returns the EndVlanId field if non-nil, zero value otherwise.

### GetEndVlanIdOk

`func (o *Vlanrange) GetEndVlanIdOk() (*int64, bool)`

GetEndVlanIdOk returns a tuple with the EndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndVlanId

`func (o *Vlanrange) SetEndVlanId(v int64)`

SetEndVlanId sets EndVlanId field to given value.

### HasEndVlanId

`func (o *Vlanrange) HasEndVlanId() bool`

HasEndVlanId returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Vlanrange) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Vlanrange) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Vlanrange) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Vlanrange) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Vlanrange) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Vlanrange) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Vlanrange) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Vlanrange) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Vlanrange) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Vlanrange) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Vlanrange) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Vlanrange) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *Vlanrange) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vlanrange) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vlanrange) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vlanrange) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreCreateVlan

`func (o *Vlanrange) GetPreCreateVlan() bool`

GetPreCreateVlan returns the PreCreateVlan field if non-nil, zero value otherwise.

### GetPreCreateVlanOk

`func (o *Vlanrange) GetPreCreateVlanOk() (*bool, bool)`

GetPreCreateVlanOk returns a tuple with the PreCreateVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCreateVlan

`func (o *Vlanrange) SetPreCreateVlan(v bool)`

SetPreCreateVlan sets PreCreateVlan field to given value.

### HasPreCreateVlan

`func (o *Vlanrange) HasPreCreateVlan() bool`

HasPreCreateVlan returns a boolean if a field has been set.

### GetStartVlanId

`func (o *Vlanrange) GetStartVlanId() int64`

GetStartVlanId returns the StartVlanId field if non-nil, zero value otherwise.

### GetStartVlanIdOk

`func (o *Vlanrange) GetStartVlanIdOk() (*int64, bool)`

GetStartVlanIdOk returns a tuple with the StartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartVlanId

`func (o *Vlanrange) SetStartVlanId(v int64)`

SetStartVlanId sets StartVlanId field to given value.

### HasStartVlanId

`func (o *Vlanrange) HasStartVlanId() bool`

HasStartVlanId returns a boolean if a field has been set.

### GetVlanNamePrefix

`func (o *Vlanrange) GetVlanNamePrefix() string`

GetVlanNamePrefix returns the VlanNamePrefix field if non-nil, zero value otherwise.

### GetVlanNamePrefixOk

`func (o *Vlanrange) GetVlanNamePrefixOk() (*string, bool)`

GetVlanNamePrefixOk returns a tuple with the VlanNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNamePrefix

`func (o *Vlanrange) SetVlanNamePrefix(v string)`

SetVlanNamePrefix sets VlanNamePrefix field to given value.

### HasVlanNamePrefix

`func (o *Vlanrange) HasVlanNamePrefix() bool`

HasVlanNamePrefix returns a boolean if a field has been set.

### GetVlanView

`func (o *Vlanrange) GetVlanView() string`

GetVlanView returns the VlanView field if non-nil, zero value otherwise.

### GetVlanViewOk

`func (o *Vlanrange) GetVlanViewOk() (*string, bool)`

GetVlanViewOk returns a tuple with the VlanView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanView

`func (o *Vlanrange) SetVlanView(v string)`

SetVlanView sets VlanView field to given value.

### HasVlanView

`func (o *Vlanrange) HasVlanView() bool`

HasVlanView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


