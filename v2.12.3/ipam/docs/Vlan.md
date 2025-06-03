# Vlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AssignedTo** | Pointer to **[]string** | List of objects VLAN is assigned to. | [optional] [readonly] 
**Comment** | Pointer to **string** | A descriptive comment for this VLAN. | [optional] 
**Contact** | Pointer to **string** | Contact information for person/team managing or using VLAN. | [optional] 
**Department** | Pointer to **string** | Department where VLAN is used. | [optional] 
**Description** | Pointer to **string** | Description for the VLAN object, may be potentially used for longer VLAN names. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Id** | Pointer to **int64** | VLAN ID value. | [optional] 
**Name** | Pointer to **string** | Name of the VLAN. | [optional] 
**Parent** | Pointer to **string** | The VLAN View or VLAN Range to which this VLAN belongs. | [optional] 
**Reserved** | Pointer to **bool** | When set VLAN can only be assigned to IPAM object manually. | [optional] 
**Status** | Pointer to **string** | Status of VLAN object. Can be Assigned, Unassigned, Reserved. | [optional] [readonly] 

## Methods

### NewVlan

`func NewVlan() *Vlan`

NewVlan instantiates a new Vlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanWithDefaults

`func NewVlanWithDefaults() *Vlan`

NewVlanWithDefaults instantiates a new Vlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Vlan) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Vlan) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Vlan) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Vlan) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAssignedTo

`func (o *Vlan) GetAssignedTo() []string`

GetAssignedTo returns the AssignedTo field if non-nil, zero value otherwise.

### GetAssignedToOk

`func (o *Vlan) GetAssignedToOk() (*[]string, bool)`

GetAssignedToOk returns a tuple with the AssignedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedTo

`func (o *Vlan) SetAssignedTo(v []string)`

SetAssignedTo sets AssignedTo field to given value.

### HasAssignedTo

`func (o *Vlan) HasAssignedTo() bool`

HasAssignedTo returns a boolean if a field has been set.

### GetComment

`func (o *Vlan) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Vlan) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Vlan) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Vlan) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContact

`func (o *Vlan) GetContact() string`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Vlan) GetContactOk() (*string, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Vlan) SetContact(v string)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Vlan) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetDepartment

`func (o *Vlan) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *Vlan) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *Vlan) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *Vlan) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetDescription

`func (o *Vlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Vlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Vlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Vlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExtattrs

`func (o *Vlan) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *Vlan) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *Vlan) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *Vlan) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetId

`func (o *Vlan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Vlan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Vlan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Vlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Vlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Vlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Vlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Vlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *Vlan) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Vlan) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Vlan) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Vlan) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetReserved

`func (o *Vlan) GetReserved() bool`

GetReserved returns the Reserved field if non-nil, zero value otherwise.

### GetReservedOk

`func (o *Vlan) GetReservedOk() (*bool, bool)`

GetReservedOk returns a tuple with the Reserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserved

`func (o *Vlan) SetReserved(v bool)`

SetReserved sets Reserved field to given value.

### HasReserved

`func (o *Vlan) HasReserved() bool`

HasReserved returns a boolean if a field has been set.

### GetStatus

`func (o *Vlan) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Vlan) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Vlan) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Vlan) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


