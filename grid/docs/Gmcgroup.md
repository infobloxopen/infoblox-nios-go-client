# Gmcgroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Comment** | Pointer to **string** | Description of the group | [optional] 
**GmcPromotionPolicy** | Pointer to **string** | This field decides whether the members join back at the same time or sequentially with time gap of 30 seconds. | [optional] 
**Members** | Pointer to [**[]GmcgroupMembers**](GmcgroupMembers.md) | gmcgroup members | [optional] 
**Name** | Pointer to **string** | Group name | [optional] 
**ScheduledTime** | Pointer to **int64** | Absolute time at which the reconnect starts | [optional] 
**TimeZone** | Pointer to **string** | The time zone for scheduling operations. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewGmcgroup

`func NewGmcgroup() *Gmcgroup`

NewGmcgroup instantiates a new Gmcgroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGmcgroupWithDefaults

`func NewGmcgroupWithDefaults() *Gmcgroup`

NewGmcgroupWithDefaults instantiates a new Gmcgroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Gmcgroup) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Gmcgroup) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Gmcgroup) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Gmcgroup) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *Gmcgroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Gmcgroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Gmcgroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Gmcgroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGmcPromotionPolicy

`func (o *Gmcgroup) GetGmcPromotionPolicy() string`

GetGmcPromotionPolicy returns the GmcPromotionPolicy field if non-nil, zero value otherwise.

### GetGmcPromotionPolicyOk

`func (o *Gmcgroup) GetGmcPromotionPolicyOk() (*string, bool)`

GetGmcPromotionPolicyOk returns a tuple with the GmcPromotionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmcPromotionPolicy

`func (o *Gmcgroup) SetGmcPromotionPolicy(v string)`

SetGmcPromotionPolicy sets GmcPromotionPolicy field to given value.

### HasGmcPromotionPolicy

`func (o *Gmcgroup) HasGmcPromotionPolicy() bool`

HasGmcPromotionPolicy returns a boolean if a field has been set.

### GetMembers

`func (o *Gmcgroup) GetMembers() []GmcgroupMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Gmcgroup) GetMembersOk() (*[]GmcgroupMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Gmcgroup) SetMembers(v []GmcgroupMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Gmcgroup) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *Gmcgroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Gmcgroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Gmcgroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Gmcgroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduledTime

`func (o *Gmcgroup) GetScheduledTime() int64`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *Gmcgroup) GetScheduledTimeOk() (*int64, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *Gmcgroup) SetScheduledTime(v int64)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *Gmcgroup) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *Gmcgroup) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *Gmcgroup) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *Gmcgroup) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *Gmcgroup) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetUuid

`func (o *Gmcgroup) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Gmcgroup) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Gmcgroup) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Gmcgroup) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


