# DbObjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**LastSequenceId** | Pointer to **string** | The last returned sequence ID. | [optional] [readonly] 
**Object** | Pointer to **string** | The record object when supported by WAPI. Otherwise, the value is \&quot;None\&quot;. | [optional] [readonly] 
**ObjectType** | Pointer to **string** | The object type. This is undefined if the object is not supported. | [optional] [readonly] 
**UniqueId** | Pointer to **string** | The unique ID of the requested object. | [optional] [readonly] 

## Methods

### NewDbObjects

`func NewDbObjects() *DbObjects`

NewDbObjects instantiates a new DbObjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbObjectsWithDefaults

`func NewDbObjectsWithDefaults() *DbObjects`

NewDbObjectsWithDefaults instantiates a new DbObjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DbObjects) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DbObjects) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DbObjects) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DbObjects) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetLastSequenceId

`func (o *DbObjects) GetLastSequenceId() string`

GetLastSequenceId returns the LastSequenceId field if non-nil, zero value otherwise.

### GetLastSequenceIdOk

`func (o *DbObjects) GetLastSequenceIdOk() (*string, bool)`

GetLastSequenceIdOk returns a tuple with the LastSequenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSequenceId

`func (o *DbObjects) SetLastSequenceId(v string)`

SetLastSequenceId sets LastSequenceId field to given value.

### HasLastSequenceId

`func (o *DbObjects) HasLastSequenceId() bool`

HasLastSequenceId returns a boolean if a field has been set.

### GetObject

`func (o *DbObjects) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *DbObjects) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *DbObjects) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *DbObjects) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectType

`func (o *DbObjects) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *DbObjects) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *DbObjects) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *DbObjects) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetUniqueId

`func (o *DbObjects) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *DbObjects) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *DbObjects) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *DbObjects) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


