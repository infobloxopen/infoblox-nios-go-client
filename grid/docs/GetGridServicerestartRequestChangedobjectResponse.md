# GetGridServicerestartRequestChangedobjectResponse

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
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**GridServicerestartRequestChangedobject**](GridServicerestartRequestChangedobject.md) |  | [optional] 

## Methods

### NewGetGridServicerestartRequestChangedobjectResponse

`func NewGetGridServicerestartRequestChangedobjectResponse() *GetGridServicerestartRequestChangedobjectResponse`

NewGetGridServicerestartRequestChangedobjectResponse instantiates a new GetGridServicerestartRequestChangedobjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridServicerestartRequestChangedobjectResponseWithDefaults

`func NewGetGridServicerestartRequestChangedobjectResponseWithDefaults() *GetGridServicerestartRequestChangedobjectResponse`

NewGetGridServicerestartRequestChangedobjectResponseWithDefaults instantiates a new GetGridServicerestartRequestChangedobjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridServicerestartRequestChangedobjectResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridServicerestartRequestChangedobjectResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAction

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetGridServicerestartRequestChangedobjectResponse) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetGridServicerestartRequestChangedobjectResponse) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetChangedProperties

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetChangedProperties() []string`

GetChangedProperties returns the ChangedProperties field if non-nil, zero value otherwise.

### GetChangedPropertiesOk

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetChangedPropertiesOk() (*[]string, bool)`

GetChangedPropertiesOk returns a tuple with the ChangedProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedProperties

`func (o *GetGridServicerestartRequestChangedobjectResponse) SetChangedProperties(v []string)`

SetChangedProperties sets ChangedProperties field to given value.

### HasChangedProperties

`func (o *GetGridServicerestartRequestChangedobjectResponse) HasChangedProperties() bool`

HasChangedProperties returns a boolean if a field has been set.

### GetChangedTime

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetChangedTime() int64`

GetChangedTime returns the ChangedTime field if non-nil, zero value otherwise.

### GetChangedTimeOk

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetChangedTimeOk() (*int64, bool)`

GetChangedTimeOk returns a tuple with the ChangedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedTime

`func (o *GetGridServicerestartRequestChangedobjectResponse) SetChangedTime(v int64)`

SetChangedTime sets ChangedTime field to given value.

### HasChangedTime

`func (o *GetGridServicerestartRequestChangedobjectResponse) HasChangedTime() bool`

HasChangedTime returns a boolean if a field has been set.

### GetObjectName

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetObjectName() string`

GetObjectName returns the ObjectName field if non-nil, zero value otherwise.

### GetObjectNameOk

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetObjectNameOk() (*string, bool)`

GetObjectNameOk returns a tuple with the ObjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectName

`func (o *GetGridServicerestartRequestChangedobjectResponse) SetObjectName(v string)`

SetObjectName sets ObjectName field to given value.

### HasObjectName

`func (o *GetGridServicerestartRequestChangedobjectResponse) HasObjectName() bool`

HasObjectName returns a boolean if a field has been set.

### GetObjectType

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GetGridServicerestartRequestChangedobjectResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *GetGridServicerestartRequestChangedobjectResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetUserName

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetGridServicerestartRequestChangedobjectResponse) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GetGridServicerestartRequestChangedobjectResponse) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetUuid

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGridServicerestartRequestChangedobjectResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGridServicerestartRequestChangedobjectResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetResult() GridServicerestartRequestChangedobject`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridServicerestartRequestChangedobjectResponse) GetResultOk() (*GridServicerestartRequestChangedobject, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridServicerestartRequestChangedobjectResponse) SetResult(v GridServicerestartRequestChangedobject)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridServicerestartRequestChangedobjectResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


