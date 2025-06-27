# ScheduledtaskChangedObjects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | This is a description of the action that is applied to this object. | [optional] [readonly] 
**Name** | Pointer to **string** | The object name. | [optional] [readonly] 
**Type** | Pointer to **string** | A value of the object type, this may contain objects that are not yet available in WAPI. | [optional] [readonly] 
**ObjectType** | Pointer to **string** | The object type. This is undefined if the object is not yet supported. | [optional] [readonly] 
**Properties** | Pointer to **[]string** | A list of properties that are being changed. | [optional] [readonly] 

## Methods

### NewScheduledtaskChangedObjects

`func NewScheduledtaskChangedObjects() *ScheduledtaskChangedObjects`

NewScheduledtaskChangedObjects instantiates a new ScheduledtaskChangedObjects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduledtaskChangedObjectsWithDefaults

`func NewScheduledtaskChangedObjectsWithDefaults() *ScheduledtaskChangedObjects`

NewScheduledtaskChangedObjectsWithDefaults instantiates a new ScheduledtaskChangedObjects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ScheduledtaskChangedObjects) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ScheduledtaskChangedObjects) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ScheduledtaskChangedObjects) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ScheduledtaskChangedObjects) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetName

`func (o *ScheduledtaskChangedObjects) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScheduledtaskChangedObjects) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScheduledtaskChangedObjects) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScheduledtaskChangedObjects) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ScheduledtaskChangedObjects) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScheduledtaskChangedObjects) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScheduledtaskChangedObjects) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ScheduledtaskChangedObjects) HasType() bool`

HasType returns a boolean if a field has been set.

### GetObjectType

`func (o *ScheduledtaskChangedObjects) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *ScheduledtaskChangedObjects) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *ScheduledtaskChangedObjects) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *ScheduledtaskChangedObjects) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetProperties

`func (o *ScheduledtaskChangedObjects) GetProperties() []string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ScheduledtaskChangedObjects) GetPropertiesOk() (*[]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ScheduledtaskChangedObjects) SetProperties(v []string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ScheduledtaskChangedObjects) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


