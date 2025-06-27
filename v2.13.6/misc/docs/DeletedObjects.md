# DeletedObjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ObjectType** | Pointer to **string** | The object type of the deleted object. This is undefined if the object is not supported. | [optional] [readonly] 

## Methods

### NewDeletedObjects

`func NewDeletedObjects() *DeletedObjects`

NewDeletedObjects instantiates a new DeletedObjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletedObjectsWithDefaults

`func NewDeletedObjectsWithDefaults() *DeletedObjects`

NewDeletedObjectsWithDefaults instantiates a new DeletedObjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DeletedObjects) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DeletedObjects) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DeletedObjects) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DeletedObjects) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetObjectType

`func (o *DeletedObjects) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DeletedObjects) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DeletedObjects) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *DeletedObjects) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


