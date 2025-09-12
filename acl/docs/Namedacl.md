# Namedacl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AccessList** | Pointer to [**[]NamedaclAccessList**](NamedaclAccessList.md) | The access control list of IPv4/IPv6 addresses, networks, TSIG-based anonymous access controls, and other named ACLs. | [optional] 
**Comment** | Pointer to **string** | Comment for the named ACL; maximum 256 characters. | [optional] 
**ExplodedAccessList** | Pointer to [**[]NamedaclExplodedAccessList**](NamedaclExplodedAccessList.md) | The exploded access list for the named ACL. This list displays all the access control entries in a named ACL and its nested named ACLs, if applicable. | [optional] [readonly] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of the named ACL. | [optional] 

## Methods

### NewNamedacl

`func NewNamedacl() *Namedacl`

NewNamedacl instantiates a new Namedacl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamedaclWithDefaults

`func NewNamedaclWithDefaults() *Namedacl`

NewNamedaclWithDefaults instantiates a new Namedacl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Namedacl) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Namedacl) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Namedacl) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Namedacl) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAccessList

`func (o *Namedacl) GetAccessList() []NamedaclAccessList`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *Namedacl) GetAccessListOk() (*[]NamedaclAccessList, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *Namedacl) SetAccessList(v []NamedaclAccessList)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *Namedacl) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetComment

`func (o *Namedacl) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Namedacl) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Namedacl) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Namedacl) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExplodedAccessList

`func (o *Namedacl) GetExplodedAccessList() []NamedaclExplodedAccessList`

GetExplodedAccessList returns the ExplodedAccessList field if non-nil, zero value otherwise.

### GetExplodedAccessListOk

`func (o *Namedacl) GetExplodedAccessListOk() (*[]NamedaclExplodedAccessList, bool)`

GetExplodedAccessListOk returns a tuple with the ExplodedAccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplodedAccessList

`func (o *Namedacl) SetExplodedAccessList(v []NamedaclExplodedAccessList)`

SetExplodedAccessList sets ExplodedAccessList field to given value.

### HasExplodedAccessList

`func (o *Namedacl) HasExplodedAccessList() bool`

HasExplodedAccessList returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Namedacl) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Namedacl) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Namedacl) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Namedacl) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Namedacl) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Namedacl) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Namedacl) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Namedacl) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Namedacl) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Namedacl) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Namedacl) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Namedacl) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetName

`func (o *Namedacl) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Namedacl) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Namedacl) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Namedacl) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


