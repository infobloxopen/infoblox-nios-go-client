# ViewScavengingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableScavenging** | Pointer to **bool** | This flag indicates if the resource record scavenging is enabled or not. | [optional] 
**EnableRecurrentScavenging** | Pointer to **bool** | This flag indicates if the recurrent resource record scavenging is enabled or not. | [optional] 
**EnableAutoReclamation** | Pointer to **bool** | This flag indicates if the automatic resource record scavenging is enabled or not. | [optional] 
**EnableRrLastQueried** | Pointer to **bool** | This flag indicates if the resource record last queried monitoring in affected zones is enabled or not. | [optional] 
**EnableZoneLastQueried** | Pointer to **bool** | This flag indicates if the last queried monitoring for affected zones is enabled or not. | [optional] 
**ReclaimAssociatedRecords** | Pointer to **bool** | This flag indicates if the associated resource record scavenging is enabled or not. | [optional] 
**ScavengingSchedule** | Pointer to [**ViewscavengingsettingsScavengingSchedule**](ViewscavengingsettingsScavengingSchedule.md) |  | [optional] 
**ExpressionList** | Pointer to [**ViewscavengingsettingsExpressionList**](ViewscavengingsettingsExpressionList.md) |  | [optional] 
**EaExpressionList** | Pointer to [**ViewscavengingsettingsEaExpressionList**](ViewscavengingsettingsEaExpressionList.md) |  | [optional] 

## Methods

### NewViewScavengingSettings

`func NewViewScavengingSettings() *ViewScavengingSettings`

NewViewScavengingSettings instantiates a new ViewScavengingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewScavengingSettingsWithDefaults

`func NewViewScavengingSettingsWithDefaults() *ViewScavengingSettings`

NewViewScavengingSettingsWithDefaults instantiates a new ViewScavengingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableScavenging

`func (o *ViewScavengingSettings) GetEnableScavenging() bool`

GetEnableScavenging returns the EnableScavenging field if non-nil, zero value otherwise.

### GetEnableScavengingOk

`func (o *ViewScavengingSettings) GetEnableScavengingOk() (*bool, bool)`

GetEnableScavengingOk returns a tuple with the EnableScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScavenging

`func (o *ViewScavengingSettings) SetEnableScavenging(v bool)`

SetEnableScavenging sets EnableScavenging field to given value.

### HasEnableScavenging

`func (o *ViewScavengingSettings) HasEnableScavenging() bool`

HasEnableScavenging returns a boolean if a field has been set.

### GetEnableRecurrentScavenging

`func (o *ViewScavengingSettings) GetEnableRecurrentScavenging() bool`

GetEnableRecurrentScavenging returns the EnableRecurrentScavenging field if non-nil, zero value otherwise.

### GetEnableRecurrentScavengingOk

`func (o *ViewScavengingSettings) GetEnableRecurrentScavengingOk() (*bool, bool)`

GetEnableRecurrentScavengingOk returns a tuple with the EnableRecurrentScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecurrentScavenging

`func (o *ViewScavengingSettings) SetEnableRecurrentScavenging(v bool)`

SetEnableRecurrentScavenging sets EnableRecurrentScavenging field to given value.

### HasEnableRecurrentScavenging

`func (o *ViewScavengingSettings) HasEnableRecurrentScavenging() bool`

HasEnableRecurrentScavenging returns a boolean if a field has been set.

### GetEnableAutoReclamation

`func (o *ViewScavengingSettings) GetEnableAutoReclamation() bool`

GetEnableAutoReclamation returns the EnableAutoReclamation field if non-nil, zero value otherwise.

### GetEnableAutoReclamationOk

`func (o *ViewScavengingSettings) GetEnableAutoReclamationOk() (*bool, bool)`

GetEnableAutoReclamationOk returns a tuple with the EnableAutoReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoReclamation

`func (o *ViewScavengingSettings) SetEnableAutoReclamation(v bool)`

SetEnableAutoReclamation sets EnableAutoReclamation field to given value.

### HasEnableAutoReclamation

`func (o *ViewScavengingSettings) HasEnableAutoReclamation() bool`

HasEnableAutoReclamation returns a boolean if a field has been set.

### GetEnableRrLastQueried

`func (o *ViewScavengingSettings) GetEnableRrLastQueried() bool`

GetEnableRrLastQueried returns the EnableRrLastQueried field if non-nil, zero value otherwise.

### GetEnableRrLastQueriedOk

`func (o *ViewScavengingSettings) GetEnableRrLastQueriedOk() (*bool, bool)`

GetEnableRrLastQueriedOk returns a tuple with the EnableRrLastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRrLastQueried

`func (o *ViewScavengingSettings) SetEnableRrLastQueried(v bool)`

SetEnableRrLastQueried sets EnableRrLastQueried field to given value.

### HasEnableRrLastQueried

`func (o *ViewScavengingSettings) HasEnableRrLastQueried() bool`

HasEnableRrLastQueried returns a boolean if a field has been set.

### GetEnableZoneLastQueried

`func (o *ViewScavengingSettings) GetEnableZoneLastQueried() bool`

GetEnableZoneLastQueried returns the EnableZoneLastQueried field if non-nil, zero value otherwise.

### GetEnableZoneLastQueriedOk

`func (o *ViewScavengingSettings) GetEnableZoneLastQueriedOk() (*bool, bool)`

GetEnableZoneLastQueriedOk returns a tuple with the EnableZoneLastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableZoneLastQueried

`func (o *ViewScavengingSettings) SetEnableZoneLastQueried(v bool)`

SetEnableZoneLastQueried sets EnableZoneLastQueried field to given value.

### HasEnableZoneLastQueried

`func (o *ViewScavengingSettings) HasEnableZoneLastQueried() bool`

HasEnableZoneLastQueried returns a boolean if a field has been set.

### GetReclaimAssociatedRecords

`func (o *ViewScavengingSettings) GetReclaimAssociatedRecords() bool`

GetReclaimAssociatedRecords returns the ReclaimAssociatedRecords field if non-nil, zero value otherwise.

### GetReclaimAssociatedRecordsOk

`func (o *ViewScavengingSettings) GetReclaimAssociatedRecordsOk() (*bool, bool)`

GetReclaimAssociatedRecordsOk returns a tuple with the ReclaimAssociatedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimAssociatedRecords

`func (o *ViewScavengingSettings) SetReclaimAssociatedRecords(v bool)`

SetReclaimAssociatedRecords sets ReclaimAssociatedRecords field to given value.

### HasReclaimAssociatedRecords

`func (o *ViewScavengingSettings) HasReclaimAssociatedRecords() bool`

HasReclaimAssociatedRecords returns a boolean if a field has been set.

### GetScavengingSchedule

`func (o *ViewScavengingSettings) GetScavengingSchedule() ViewscavengingsettingsScavengingSchedule`

GetScavengingSchedule returns the ScavengingSchedule field if non-nil, zero value otherwise.

### GetScavengingScheduleOk

`func (o *ViewScavengingSettings) GetScavengingScheduleOk() (*ViewscavengingsettingsScavengingSchedule, bool)`

GetScavengingScheduleOk returns a tuple with the ScavengingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScavengingSchedule

`func (o *ViewScavengingSettings) SetScavengingSchedule(v ViewscavengingsettingsScavengingSchedule)`

SetScavengingSchedule sets ScavengingSchedule field to given value.

### HasScavengingSchedule

`func (o *ViewScavengingSettings) HasScavengingSchedule() bool`

HasScavengingSchedule returns a boolean if a field has been set.

### GetExpressionList

`func (o *ViewScavengingSettings) GetExpressionList() ViewscavengingsettingsExpressionList`

GetExpressionList returns the ExpressionList field if non-nil, zero value otherwise.

### GetExpressionListOk

`func (o *ViewScavengingSettings) GetExpressionListOk() (*ViewscavengingsettingsExpressionList, bool)`

GetExpressionListOk returns a tuple with the ExpressionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionList

`func (o *ViewScavengingSettings) SetExpressionList(v ViewscavengingsettingsExpressionList)`

SetExpressionList sets ExpressionList field to given value.

### HasExpressionList

`func (o *ViewScavengingSettings) HasExpressionList() bool`

HasExpressionList returns a boolean if a field has been set.

### GetEaExpressionList

`func (o *ViewScavengingSettings) GetEaExpressionList() ViewscavengingsettingsEaExpressionList`

GetEaExpressionList returns the EaExpressionList field if non-nil, zero value otherwise.

### GetEaExpressionListOk

`func (o *ViewScavengingSettings) GetEaExpressionListOk() (*ViewscavengingsettingsEaExpressionList, bool)`

GetEaExpressionListOk returns a tuple with the EaExpressionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaExpressionList

`func (o *ViewScavengingSettings) SetEaExpressionList(v ViewscavengingsettingsEaExpressionList)`

SetEaExpressionList sets EaExpressionList field to given value.

### HasEaExpressionList

`func (o *ViewScavengingSettings) HasEaExpressionList() bool`

HasEaExpressionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


