# GetGmcgroupResponse

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
**Result** | Pointer to [**Gmcgroup**](Gmcgroup.md) |  | [optional] 

## Methods

### NewGetGmcgroupResponse

`func NewGetGmcgroupResponse() *GetGmcgroupResponse`

NewGetGmcgroupResponse instantiates a new GetGmcgroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGmcgroupResponseWithDefaults

`func NewGetGmcgroupResponseWithDefaults() *GetGmcgroupResponse`

NewGetGmcgroupResponseWithDefaults instantiates a new GetGmcgroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGmcgroupResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGmcgroupResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGmcgroupResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGmcgroupResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetGmcgroupResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetGmcgroupResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetGmcgroupResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetGmcgroupResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetGmcPromotionPolicy

`func (o *GetGmcgroupResponse) GetGmcPromotionPolicy() string`

GetGmcPromotionPolicy returns the GmcPromotionPolicy field if non-nil, zero value otherwise.

### GetGmcPromotionPolicyOk

`func (o *GetGmcgroupResponse) GetGmcPromotionPolicyOk() (*string, bool)`

GetGmcPromotionPolicyOk returns a tuple with the GmcPromotionPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmcPromotionPolicy

`func (o *GetGmcgroupResponse) SetGmcPromotionPolicy(v string)`

SetGmcPromotionPolicy sets GmcPromotionPolicy field to given value.

### HasGmcPromotionPolicy

`func (o *GetGmcgroupResponse) HasGmcPromotionPolicy() bool`

HasGmcPromotionPolicy returns a boolean if a field has been set.

### GetMembers

`func (o *GetGmcgroupResponse) GetMembers() []GmcgroupMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetGmcgroupResponse) GetMembersOk() (*[]GmcgroupMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetGmcgroupResponse) SetMembers(v []GmcgroupMembers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetGmcgroupResponse) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *GetGmcgroupResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGmcgroupResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGmcgroupResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGmcgroupResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetScheduledTime

`func (o *GetGmcgroupResponse) GetScheduledTime() int64`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *GetGmcgroupResponse) GetScheduledTimeOk() (*int64, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *GetGmcgroupResponse) SetScheduledTime(v int64)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *GetGmcgroupResponse) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetTimeZone

`func (o *GetGmcgroupResponse) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetGmcgroupResponse) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetGmcgroupResponse) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *GetGmcgroupResponse) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.

### GetResult

`func (o *GetGmcgroupResponse) GetResult() Gmcgroup`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGmcgroupResponse) GetResultOk() (*Gmcgroup, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGmcgroupResponse) SetResult(v Gmcgroup)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGmcgroupResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


