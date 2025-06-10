# Gmcschedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**ActivateGmcGroupSchedule** | Pointer to **bool** | Determines whether the gmc schedule is active. | [optional] 
**GmcGroups** | Pointer to **[]string** | Object array of gmc groups | [optional] [readonly] 

## Methods

### NewGmcschedule

`func NewGmcschedule() *Gmcschedule`

NewGmcschedule instantiates a new Gmcschedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGmcscheduleWithDefaults

`func NewGmcscheduleWithDefaults() *Gmcschedule`

NewGmcscheduleWithDefaults instantiates a new Gmcschedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Gmcschedule) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Gmcschedule) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Gmcschedule) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Gmcschedule) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetActivateGmcGroupSchedule

`func (o *Gmcschedule) GetActivateGmcGroupSchedule() bool`

GetActivateGmcGroupSchedule returns the ActivateGmcGroupSchedule field if non-nil, zero value otherwise.

### GetActivateGmcGroupScheduleOk

`func (o *Gmcschedule) GetActivateGmcGroupScheduleOk() (*bool, bool)`

GetActivateGmcGroupScheduleOk returns a tuple with the ActivateGmcGroupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivateGmcGroupSchedule

`func (o *Gmcschedule) SetActivateGmcGroupSchedule(v bool)`

SetActivateGmcGroupSchedule sets ActivateGmcGroupSchedule field to given value.

### HasActivateGmcGroupSchedule

`func (o *Gmcschedule) HasActivateGmcGroupSchedule() bool`

HasActivateGmcGroupSchedule returns a boolean if a field has been set.

### GetGmcGroups

`func (o *Gmcschedule) GetGmcGroups() []string`

GetGmcGroups returns the GmcGroups field if non-nil, zero value otherwise.

### GetGmcGroupsOk

`func (o *Gmcschedule) GetGmcGroupsOk() (*[]string, bool)`

GetGmcGroupsOk returns a tuple with the GmcGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGmcGroups

`func (o *Gmcschedule) SetGmcGroups(v []string)`

SetGmcGroups sets GmcGroups field to given value.

### HasGmcGroups

`func (o *Gmcschedule) HasGmcGroups() bool`

HasGmcGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


