# GetVlanrangeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
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
**Result** | Pointer to [**Vlanrange**](Vlanrange.md) |  | [optional] 

## Methods

### NewGetVlanrangeResponse

`func NewGetVlanrangeResponse() *GetVlanrangeResponse`

NewGetVlanrangeResponse instantiates a new GetVlanrangeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVlanrangeResponseWithDefaults

`func NewGetVlanrangeResponseWithDefaults() *GetVlanrangeResponse`

NewGetVlanrangeResponseWithDefaults instantiates a new GetVlanrangeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetVlanrangeResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetVlanrangeResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetVlanrangeResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetVlanrangeResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetVlanrangeResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetVlanrangeResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetVlanrangeResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetVlanrangeResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDeleteVlans

`func (o *GetVlanrangeResponse) GetDeleteVlans() bool`

GetDeleteVlans returns the DeleteVlans field if non-nil, zero value otherwise.

### GetDeleteVlansOk

`func (o *GetVlanrangeResponse) GetDeleteVlansOk() (*bool, bool)`

GetDeleteVlansOk returns a tuple with the DeleteVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteVlans

`func (o *GetVlanrangeResponse) SetDeleteVlans(v bool)`

SetDeleteVlans sets DeleteVlans field to given value.

### HasDeleteVlans

`func (o *GetVlanrangeResponse) HasDeleteVlans() bool`

HasDeleteVlans returns a boolean if a field has been set.

### GetEndVlanId

`func (o *GetVlanrangeResponse) GetEndVlanId() int64`

GetEndVlanId returns the EndVlanId field if non-nil, zero value otherwise.

### GetEndVlanIdOk

`func (o *GetVlanrangeResponse) GetEndVlanIdOk() (*int64, bool)`

GetEndVlanIdOk returns a tuple with the EndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndVlanId

`func (o *GetVlanrangeResponse) SetEndVlanId(v int64)`

SetEndVlanId sets EndVlanId field to given value.

### HasEndVlanId

`func (o *GetVlanrangeResponse) HasEndVlanId() bool`

HasEndVlanId returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetVlanrangeResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetVlanrangeResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetVlanrangeResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetVlanrangeResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetVlanrangeResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetVlanrangeResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetVlanrangeResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetVlanrangeResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetVlanrangeResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetVlanrangeResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetVlanrangeResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetVlanrangeResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *GetVlanrangeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVlanrangeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVlanrangeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVlanrangeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreCreateVlan

`func (o *GetVlanrangeResponse) GetPreCreateVlan() bool`

GetPreCreateVlan returns the PreCreateVlan field if non-nil, zero value otherwise.

### GetPreCreateVlanOk

`func (o *GetVlanrangeResponse) GetPreCreateVlanOk() (*bool, bool)`

GetPreCreateVlanOk returns a tuple with the PreCreateVlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreCreateVlan

`func (o *GetVlanrangeResponse) SetPreCreateVlan(v bool)`

SetPreCreateVlan sets PreCreateVlan field to given value.

### HasPreCreateVlan

`func (o *GetVlanrangeResponse) HasPreCreateVlan() bool`

HasPreCreateVlan returns a boolean if a field has been set.

### GetStartVlanId

`func (o *GetVlanrangeResponse) GetStartVlanId() int64`

GetStartVlanId returns the StartVlanId field if non-nil, zero value otherwise.

### GetStartVlanIdOk

`func (o *GetVlanrangeResponse) GetStartVlanIdOk() (*int64, bool)`

GetStartVlanIdOk returns a tuple with the StartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartVlanId

`func (o *GetVlanrangeResponse) SetStartVlanId(v int64)`

SetStartVlanId sets StartVlanId field to given value.

### HasStartVlanId

`func (o *GetVlanrangeResponse) HasStartVlanId() bool`

HasStartVlanId returns a boolean if a field has been set.

### GetVlanNamePrefix

`func (o *GetVlanrangeResponse) GetVlanNamePrefix() string`

GetVlanNamePrefix returns the VlanNamePrefix field if non-nil, zero value otherwise.

### GetVlanNamePrefixOk

`func (o *GetVlanrangeResponse) GetVlanNamePrefixOk() (*string, bool)`

GetVlanNamePrefixOk returns a tuple with the VlanNamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanNamePrefix

`func (o *GetVlanrangeResponse) SetVlanNamePrefix(v string)`

SetVlanNamePrefix sets VlanNamePrefix field to given value.

### HasVlanNamePrefix

`func (o *GetVlanrangeResponse) HasVlanNamePrefix() bool`

HasVlanNamePrefix returns a boolean if a field has been set.

### GetVlanView

`func (o *GetVlanrangeResponse) GetVlanView() string`

GetVlanView returns the VlanView field if non-nil, zero value otherwise.

### GetVlanViewOk

`func (o *GetVlanrangeResponse) GetVlanViewOk() (*string, bool)`

GetVlanViewOk returns a tuple with the VlanView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanView

`func (o *GetVlanrangeResponse) SetVlanView(v string)`

SetVlanView sets VlanView field to given value.

### HasVlanView

`func (o *GetVlanrangeResponse) HasVlanView() bool`

HasVlanView returns a boolean if a field has been set.

### GetResult

`func (o *GetVlanrangeResponse) GetResult() Vlanrange`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetVlanrangeResponse) GetResultOk() (*Vlanrange, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetVlanrangeResponse) SetResult(v Vlanrange)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetVlanrangeResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


