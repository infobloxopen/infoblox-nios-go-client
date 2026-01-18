# GridServicerestartRequestChangedobject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Action** | Pointer to **string** | The operation on the changed object. | [optional] [readonly] 
**ChangedProperties** | Pointer to **[]string** | The list of changed properties in the object. | [optional] [readonly] 
**ChangedTime** | Pointer to **int64** | The time when the object was changed. | [optional] [readonly] 
**ObjectName** | Pointer to **string** | The name of the changed object. | [optional] [readonly] 
**ObjectType** | Pointer to **string** | The type of the changed object. This is undefined if the object is not supported. | [optional] [readonly] 
**UserName** | Pointer to **string** | The name of the user who changed the object properties. | [optional] [readonly] 

## Methods

### NewGridServicerestartRequestChangedobject

`func NewGridServicerestartRequestChangedobject() *GridServicerestartRequestChangedobject`

NewGridServicerestartRequestChangedobject instantiates a new GridServicerestartRequestChangedobject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridServicerestartRequestChangedobjectWithDefaults

`func NewGridServicerestartRequestChangedobjectWithDefaults() *GridServicerestartRequestChangedobject`

NewGridServicerestartRequestChangedobjectWithDefaults instantiates a new GridServicerestartRequestChangedobject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridServicerestartRequestChangedobject) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridServicerestartRequestChangedobject) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridServicerestartRequestChangedobject) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridServicerestartRequestChangedobject) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAction

`func (o *GridServicerestartRequestChangedobject) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GridServicerestartRequestChangedobject) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GridServicerestartRequestChangedobject) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GridServicerestartRequestChangedobject) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetChangedProperties

`func (o *GridServicerestartRequestChangedobject) GetChangedProperties() []string`

GetChangedProperties returns the ChangedProperties field if non-nil, zero value otherwise.

### GetChangedPropertiesOk

`func (o *GridServicerestartRequestChangedobject) GetChangedPropertiesOk() (*[]string, bool)`

GetChangedPropertiesOk returns a tuple with the ChangedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedProperties

`func (o *GridServicerestartRequestChangedobject) SetChangedProperties(v []string)`

SetChangedProperties sets ChangedProperties field to given value.

### HasChangedProperties

`func (o *GridServicerestartRequestChangedobject) HasChangedProperties() bool`

HasChangedProperties returns a boolean if a field has been set.

### GetChangedTime

`func (o *GridServicerestartRequestChangedobject) GetChangedTime() int64`

GetChangedTime returns the ChangedTime field if non-nil, zero value otherwise.

### GetChangedTimeOk

`func (o *GridServicerestartRequestChangedobject) GetChangedTimeOk() (*int64, bool)`

GetChangedTimeOk returns a tuple with the ChangedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedTime

`func (o *GridServicerestartRequestChangedobject) SetChangedTime(v int64)`

SetChangedTime sets ChangedTime field to given value.

### HasChangedTime

`func (o *GridServicerestartRequestChangedobject) HasChangedTime() bool`

HasChangedTime returns a boolean if a field has been set.

### GetObjectName

`func (o *GridServicerestartRequestChangedobject) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *GridServicerestartRequestChangedobject) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *GridServicerestartRequestChangedobject) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *GridServicerestartRequestChangedobject) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *GridServicerestartRequestChangedobject) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GridServicerestartRequestChangedobject) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GridServicerestartRequestChangedobject) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *GridServicerestartRequestChangedobject) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetUserName

`func (o *GridServicerestartRequestChangedobject) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GridServicerestartRequestChangedobject) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GridServicerestartRequestChangedobject) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GridServicerestartRequestChangedobject) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


