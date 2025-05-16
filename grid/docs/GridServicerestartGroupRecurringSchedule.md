# GridServicerestartGroupRecurringSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to **[]string** | The list of applicable services for the restart. | [optional] 
**Mode** | Pointer to **string** | The restart method for a Grid restart. | [optional] 
**Schedule** | Pointer to [**GridservicerestartgrouprecurringscheduleSchedule**](GridservicerestartgrouprecurringscheduleSchedule.md) |  | [optional] 
**Force** | Pointer to **bool** | Determines if the Restart Group should have a force restart. | [optional] 

## Methods

### NewGridServicerestartGroupRecurringSchedule

`func NewGridServicerestartGroupRecurringSchedule() *GridServicerestartGroupRecurringSchedule`

NewGridServicerestartGroupRecurringSchedule instantiates a new GridServicerestartGroupRecurringSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridServicerestartGroupRecurringScheduleWithDefaults

`func NewGridServicerestartGroupRecurringScheduleWithDefaults() *GridServicerestartGroupRecurringSchedule`

NewGridServicerestartGroupRecurringScheduleWithDefaults instantiates a new GridServicerestartGroupRecurringSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *GridServicerestartGroupRecurringSchedule) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GridServicerestartGroupRecurringSchedule) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GridServicerestartGroupRecurringSchedule) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *GridServicerestartGroupRecurringSchedule) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetMode

`func (o *GridServicerestartGroupRecurringSchedule) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *GridServicerestartGroupRecurringSchedule) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *GridServicerestartGroupRecurringSchedule) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *GridServicerestartGroupRecurringSchedule) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetSchedule

`func (o *GridServicerestartGroupRecurringSchedule) GetSchedule() GridservicerestartgrouprecurringscheduleSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *GridServicerestartGroupRecurringSchedule) GetScheduleOk() (*GridservicerestartgrouprecurringscheduleSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *GridServicerestartGroupRecurringSchedule) SetSchedule(v GridservicerestartgrouprecurringscheduleSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *GridServicerestartGroupRecurringSchedule) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetForce

`func (o *GridServicerestartGroupRecurringSchedule) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *GridServicerestartGroupRecurringSchedule) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *GridServicerestartGroupRecurringSchedule) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *GridServicerestartGroupRecurringSchedule) HasForce() bool`

HasForce returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


