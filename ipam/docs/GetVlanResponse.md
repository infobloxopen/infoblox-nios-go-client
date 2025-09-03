# GetVlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AssignedTo** | Pointer to **[]string** | List of objects VLAN is assigned to. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment for this VLAN. | [optional] 
**Contact** | Pointer to **string** | Contact information for person/team managing or using VLAN. | [optional] 
**Department** | Pointer to **string** | Department where VLAN is used. | [optional] 
**Description** | Pointer to **string** | Description for the VLAN object, may be potentially used for longer VLAN names. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Id** | Pointer to **int64** | VLAN ID value. | [optional] 
**Name** | Pointer to **string** | Name of the VLAN. | [optional] 
**Parent** | Pointer to **string** | The VLAN View or VLAN Range to which this VLAN belongs. | [optional] 
**Reserved** | Pointer to **bool** | When set VLAN can only be assigned to IPAM object manually. | [optional] 
**Status** | Pointer to **string** | Status of VLAN object. Can be Assigned, Unassigned, Reserved. | [optional] [readonly] 
**Result** | Pointer to [**Vlan**](Vlan.md) |  | [optional] 

## Methods

### NewGetVlanResponse

`func NewGetVlanResponse() *GetVlanResponse`

NewGetVlanResponse instantiates a new GetVlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVlanResponseWithDefaults

`func NewGetVlanResponseWithDefaults() *GetVlanResponse`

NewGetVlanResponseWithDefaults instantiates a new GetVlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetVlanResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetVlanResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetVlanResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetVlanResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAssignedTo

`func (o *GetVlanResponse) GetAssignedTo() []string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *GetVlanResponse) GetAssignedToOk() (*[]string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *GetVlanResponse) SetAssignedTo(v []string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *GetVlanResponse) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetComment

`func (o *GetVlanResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetVlanResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetVlanResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetVlanResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContact

`func (o *GetVlanResponse) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *GetVlanResponse) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *GetVlanResponse) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *GetVlanResponse) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDepartment

`func (o *GetVlanResponse) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *GetVlanResponse) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *GetVlanResponse) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *GetVlanResponse) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDescription

`func (o *GetVlanResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetVlanResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetVlanResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetVlanResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetVlanResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetVlanResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetVlanResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetVlanResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetVlanResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetVlanResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetVlanResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetVlanResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetVlanResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetVlanResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetVlanResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetVlanResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetId

`func (o *GetVlanResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVlanResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVlanResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetVlanResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetVlanResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVlanResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVlanResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVlanResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *GetVlanResponse) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetVlanResponse) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetVlanResponse) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetVlanResponse) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetReserved

`func (o *GetVlanResponse) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *GetVlanResponse) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *GetVlanResponse) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *GetVlanResponse) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetStatus

`func (o *GetVlanResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVlanResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVlanResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVlanResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetResult

`func (o *GetVlanResponse) GetResult() Vlan`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetVlanResponse) GetResultOk() (*Vlan, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetVlanResponse) SetResult(v Vlan)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetVlanResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


