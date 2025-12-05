# Natgroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | The NAT group descriptive comment. | [optional] 
**Name** | Pointer to **string** | The name of a NAT group object. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewNatgroup

`func NewNatgroup() *Natgroup`

NewNatgroup instantiates a new Natgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNatgroupWithDefaults

`func NewNatgroupWithDefaults() *Natgroup`

NewNatgroupWithDefaults instantiates a new Natgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Natgroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Natgroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Natgroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Natgroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Natgroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Natgroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Natgroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Natgroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetName

`func (o *Natgroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Natgroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Natgroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Natgroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *Natgroup) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Natgroup) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Natgroup) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Natgroup) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


