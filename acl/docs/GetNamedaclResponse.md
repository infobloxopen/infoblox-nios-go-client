# GetNamedaclResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**AccessList** | Pointer to [**[]NamedaclAccessList**](NamedaclAccessList.md) | The access control list of IPv4/IPv6 addresses, networks, TSIG-based anonymous access controls, and other named ACLs. | [optional] 
**Comment** | Pointer to **string** | Comment for the named ACL; maximum 256 characters. | [optional] 
**ExplodedAccessList** | Pointer to [**[]NamedaclExplodedAccessList**](NamedaclExplodedAccessList.md) | The exploded access list for the named ACL. This list displays all the access control entries in a named ACL and its nested named ACLs, if applicable. | [optional] [readonly] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | The name of the named ACL. | [optional] 
**Result** | Pointer to [**Namedacl**](Namedacl.md) |  | [optional] 

## Methods

### NewGetNamedaclResponse

`func NewGetNamedaclResponse() *GetNamedaclResponse`

NewGetNamedaclResponse instantiates a new GetNamedaclResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNamedaclResponseWithDefaults

`func NewGetNamedaclResponseWithDefaults() *GetNamedaclResponse`

NewGetNamedaclResponseWithDefaults instantiates a new GetNamedaclResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetNamedaclResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetNamedaclResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetNamedaclResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetNamedaclResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAccessList

`func (o *GetNamedaclResponse) GetAccessList() []NamedaclAccessList`

GetAccessList returns the AccessList field if non-nil, zero value otherwise.

### GetAccessListOk

`func (o *GetNamedaclResponse) GetAccessListOk() (*[]NamedaclAccessList, bool)`

GetAccessListOk returns a tuple with the AccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessList

`func (o *GetNamedaclResponse) SetAccessList(v []NamedaclAccessList)`

SetAccessList sets AccessList field to given value.

### HasAccessList

`func (o *GetNamedaclResponse) HasAccessList() bool`

HasAccessList returns a boolean if a field has been set.

### GetComment

`func (o *GetNamedaclResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetNamedaclResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetNamedaclResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetNamedaclResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExplodedAccessList

`func (o *GetNamedaclResponse) GetExplodedAccessList() []NamedaclExplodedAccessList`

GetExplodedAccessList returns the ExplodedAccessList field if non-nil, zero value otherwise.

### GetExplodedAccessListOk

`func (o *GetNamedaclResponse) GetExplodedAccessListOk() (*[]NamedaclExplodedAccessList, bool)`

GetExplodedAccessListOk returns a tuple with the ExplodedAccessList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplodedAccessList

`func (o *GetNamedaclResponse) SetExplodedAccessList(v []NamedaclExplodedAccessList)`

SetExplodedAccessList sets ExplodedAccessList field to given value.

### HasExplodedAccessList

`func (o *GetNamedaclResponse) HasExplodedAccessList() bool`

HasExplodedAccessList returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetNamedaclResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetNamedaclResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetNamedaclResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetNamedaclResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetName

`func (o *GetNamedaclResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNamedaclResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNamedaclResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNamedaclResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetNamedaclResponse) GetResult() Namedacl`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetNamedaclResponse) GetResultOk() (*Namedacl, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetNamedaclResponse) SetResult(v Namedacl)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetNamedaclResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


