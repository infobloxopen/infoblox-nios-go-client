# GridObjectsChangesTrackingSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Determines whether the objects changes tracking feature is enabled or not. | [optional] 
**EnableCompletion** | Pointer to **int64** | Determines the percentage of completion for objects changes tracking. | [optional] [readonly] 
**State** | Pointer to **string** | Determines the objects changes tracking enable state. | [optional] [readonly] 
**MaxTimeToTrack** | Pointer to **int64** | Maximum time period in seconds to track the deleted objects changes. You can enter a value from 7200 - 604800 seconds. | [optional] 
**MaxObjsToTrack** | Pointer to **int64** | Maximum number of deleted objects retained for tracking. You can enter a value from 2000 - 20000. | [optional] 

## Methods

### NewGridObjectsChangesTrackingSetting

`func NewGridObjectsChangesTrackingSetting() *GridObjectsChangesTrackingSetting`

NewGridObjectsChangesTrackingSetting instantiates a new GridObjectsChangesTrackingSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridObjectsChangesTrackingSettingWithDefaults

`func NewGridObjectsChangesTrackingSettingWithDefaults() *GridObjectsChangesTrackingSetting`

NewGridObjectsChangesTrackingSettingWithDefaults instantiates a new GridObjectsChangesTrackingSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GridObjectsChangesTrackingSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GridObjectsChangesTrackingSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GridObjectsChangesTrackingSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GridObjectsChangesTrackingSetting) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetEnableCompletion

`func (o *GridObjectsChangesTrackingSetting) GetEnableCompletion() int64`

GetEnableCompletion returns the EnableCompletion field if non-nil, zero value otherwise.

### GetEnableCompletionOk

`func (o *GridObjectsChangesTrackingSetting) GetEnableCompletionOk() (*int64, bool)`

GetEnableCompletionOk returns a tuple with the EnableCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompletion

`func (o *GridObjectsChangesTrackingSetting) SetEnableCompletion(v int64)`

SetEnableCompletion sets EnableCompletion field to given value.

### HasEnableCompletion

`func (o *GridObjectsChangesTrackingSetting) HasEnableCompletion() bool`

HasEnableCompletion returns a boolean if a field has been set.

### GetState

`func (o *GridObjectsChangesTrackingSetting) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GridObjectsChangesTrackingSetting) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GridObjectsChangesTrackingSetting) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GridObjectsChangesTrackingSetting) HasState() bool`

HasState returns a boolean if a field has been set.

### GetMaxTimeToTrack

`func (o *GridObjectsChangesTrackingSetting) GetMaxTimeToTrack() int64`

GetMaxTimeToTrack returns the MaxTimeToTrack field if non-nil, zero value otherwise.

### GetMaxTimeToTrackOk

`func (o *GridObjectsChangesTrackingSetting) GetMaxTimeToTrackOk() (*int64, bool)`

GetMaxTimeToTrackOk returns a tuple with the MaxTimeToTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTimeToTrack

`func (o *GridObjectsChangesTrackingSetting) SetMaxTimeToTrack(v int64)`

SetMaxTimeToTrack sets MaxTimeToTrack field to given value.

### HasMaxTimeToTrack

`func (o *GridObjectsChangesTrackingSetting) HasMaxTimeToTrack() bool`

HasMaxTimeToTrack returns a boolean if a field has been set.

### GetMaxObjsToTrack

`func (o *GridObjectsChangesTrackingSetting) GetMaxObjsToTrack() int64`

GetMaxObjsToTrack returns the MaxObjsToTrack field if non-nil, zero value otherwise.

### GetMaxObjsToTrackOk

`func (o *GridObjectsChangesTrackingSetting) GetMaxObjsToTrackOk() (*int64, bool)`

GetMaxObjsToTrackOk returns a tuple with the MaxObjsToTrack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxObjsToTrack

`func (o *GridObjectsChangesTrackingSetting) SetMaxObjsToTrack(v int64)`

SetMaxObjsToTrack sets MaxObjsToTrack field to given value.

### HasMaxObjsToTrack

`func (o *GridObjectsChangesTrackingSetting) HasMaxObjsToTrack() bool`

HasMaxObjsToTrack returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


