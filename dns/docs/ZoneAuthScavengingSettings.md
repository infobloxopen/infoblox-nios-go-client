# ZoneAuthScavengingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableScavenging** | Pointer to **bool** | This flag indicates if the resource record scavenging is enabled or not. | [optional] 
**EnableRecurrentScavenging** | Pointer to **bool** | This flag indicates if the recurrent resource record scavenging is enabled or not. | [optional] 
**EnableAutoReclamation** | Pointer to **bool** | This flag indicates if the automatic resource record scavenging is enabled or not. | [optional] 
**EnableRrLastQueried** | Pointer to **bool** | This flag indicates if the resource record last queried monitoring in affected zones is enabled or not. | [optional] 
**EnableZoneLastQueried** | Pointer to **bool** | This flag indicates if the last queried monitoring for affected zones is enabled or not. | [optional] 
**ReclaimAssociatedRecords** | Pointer to **bool** | This flag indicates if the associated resource record scavenging is enabled or not. | [optional] 
**ScavengingSchedule** | Pointer to [**ZoneauthscavengingsettingsScavengingSchedule**](ZoneauthscavengingsettingsScavengingSchedule.md) |  | [optional] 
**ExpressionList** | Pointer to [**[]ZoneauthscavengingsettingsExpressionList**](ZoneauthscavengingsettingsExpressionList.md) | The expression list. The particular record is treated as reclaimable if expression condition evaluates to &#39;true&#39; for given record if scavenging hasn&#39;t been manually disabled on a given resource record. | [optional] 
**EaExpressionList** | Pointer to [**[]ZoneauthscavengingsettingsEaExpressionList**](ZoneauthscavengingsettingsEaExpressionList.md) | The extensible attributes expression list. The particular record is treated as reclaimable if extensible attributes expression condition evaluates to &#39;true&#39; for given record if scavenging hasn&#39;t been manually disabled on a given resource record. | [optional] 

## Methods

### NewZoneAuthScavengingSettings

`func NewZoneAuthScavengingSettings() *ZoneAuthScavengingSettings`

NewZoneAuthScavengingSettings instantiates a new ZoneAuthScavengingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthScavengingSettingsWithDefaults

`func NewZoneAuthScavengingSettingsWithDefaults() *ZoneAuthScavengingSettings`

NewZoneAuthScavengingSettingsWithDefaults instantiates a new ZoneAuthScavengingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableScavenging

`func (o *ZoneAuthScavengingSettings) GetEnableScavenging() bool`

GetEnableScavenging returns the EnableScavenging field if non-nil, zero value otherwise.

### GetEnableScavengingOk

`func (o *ZoneAuthScavengingSettings) GetEnableScavengingOk() (*bool, bool)`

GetEnableScavengingOk returns a tuple with the EnableScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScavenging

`func (o *ZoneAuthScavengingSettings) SetEnableScavenging(v bool)`

SetEnableScavenging sets EnableScavenging field to given value.

### HasEnableScavenging

`func (o *ZoneAuthScavengingSettings) HasEnableScavenging() bool`

HasEnableScavenging returns a boolean if a field has been set.

### GetEnableRecurrentScavenging

`func (o *ZoneAuthScavengingSettings) GetEnableRecurrentScavenging() bool`

GetEnableRecurrentScavenging returns the EnableRecurrentScavenging field if non-nil, zero value otherwise.

### GetEnableRecurrentScavengingOk

`func (o *ZoneAuthScavengingSettings) GetEnableRecurrentScavengingOk() (*bool, bool)`

GetEnableRecurrentScavengingOk returns a tuple with the EnableRecurrentScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRecurrentScavenging

`func (o *ZoneAuthScavengingSettings) SetEnableRecurrentScavenging(v bool)`

SetEnableRecurrentScavenging sets EnableRecurrentScavenging field to given value.

### HasEnableRecurrentScavenging

`func (o *ZoneAuthScavengingSettings) HasEnableRecurrentScavenging() bool`

HasEnableRecurrentScavenging returns a boolean if a field has been set.

### GetEnableAutoReclamation

`func (o *ZoneAuthScavengingSettings) GetEnableAutoReclamation() bool`

GetEnableAutoReclamation returns the EnableAutoReclamation field if non-nil, zero value otherwise.

### GetEnableAutoReclamationOk

`func (o *ZoneAuthScavengingSettings) GetEnableAutoReclamationOk() (*bool, bool)`

GetEnableAutoReclamationOk returns a tuple with the EnableAutoReclamation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoReclamation

`func (o *ZoneAuthScavengingSettings) SetEnableAutoReclamation(v bool)`

SetEnableAutoReclamation sets EnableAutoReclamation field to given value.

### HasEnableAutoReclamation

`func (o *ZoneAuthScavengingSettings) HasEnableAutoReclamation() bool`

HasEnableAutoReclamation returns a boolean if a field has been set.

### GetEnableRrLastQueried

`func (o *ZoneAuthScavengingSettings) GetEnableRrLastQueried() bool`

GetEnableRrLastQueried returns the EnableRrLastQueried field if non-nil, zero value otherwise.

### GetEnableRrLastQueriedOk

`func (o *ZoneAuthScavengingSettings) GetEnableRrLastQueriedOk() (*bool, bool)`

GetEnableRrLastQueriedOk returns a tuple with the EnableRrLastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRrLastQueried

`func (o *ZoneAuthScavengingSettings) SetEnableRrLastQueried(v bool)`

SetEnableRrLastQueried sets EnableRrLastQueried field to given value.

### HasEnableRrLastQueried

`func (o *ZoneAuthScavengingSettings) HasEnableRrLastQueried() bool`

HasEnableRrLastQueried returns a boolean if a field has been set.

### GetEnableZoneLastQueried

`func (o *ZoneAuthScavengingSettings) GetEnableZoneLastQueried() bool`

GetEnableZoneLastQueried returns the EnableZoneLastQueried field if non-nil, zero value otherwise.

### GetEnableZoneLastQueriedOk

`func (o *ZoneAuthScavengingSettings) GetEnableZoneLastQueriedOk() (*bool, bool)`

GetEnableZoneLastQueriedOk returns a tuple with the EnableZoneLastQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableZoneLastQueried

`func (o *ZoneAuthScavengingSettings) SetEnableZoneLastQueried(v bool)`

SetEnableZoneLastQueried sets EnableZoneLastQueried field to given value.

### HasEnableZoneLastQueried

`func (o *ZoneAuthScavengingSettings) HasEnableZoneLastQueried() bool`

HasEnableZoneLastQueried returns a boolean if a field has been set.

### GetReclaimAssociatedRecords

`func (o *ZoneAuthScavengingSettings) GetReclaimAssociatedRecords() bool`

GetReclaimAssociatedRecords returns the ReclaimAssociatedRecords field if non-nil, zero value otherwise.

### GetReclaimAssociatedRecordsOk

`func (o *ZoneAuthScavengingSettings) GetReclaimAssociatedRecordsOk() (*bool, bool)`

GetReclaimAssociatedRecordsOk returns a tuple with the ReclaimAssociatedRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReclaimAssociatedRecords

`func (o *ZoneAuthScavengingSettings) SetReclaimAssociatedRecords(v bool)`

SetReclaimAssociatedRecords sets ReclaimAssociatedRecords field to given value.

### HasReclaimAssociatedRecords

`func (o *ZoneAuthScavengingSettings) HasReclaimAssociatedRecords() bool`

HasReclaimAssociatedRecords returns a boolean if a field has been set.

### GetScavengingSchedule

`func (o *ZoneAuthScavengingSettings) GetScavengingSchedule() ZoneauthscavengingsettingsScavengingSchedule`

GetScavengingSchedule returns the ScavengingSchedule field if non-nil, zero value otherwise.

### GetScavengingScheduleOk

`func (o *ZoneAuthScavengingSettings) GetScavengingScheduleOk() (*ZoneauthscavengingsettingsScavengingSchedule, bool)`

GetScavengingScheduleOk returns a tuple with the ScavengingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScavengingSchedule

`func (o *ZoneAuthScavengingSettings) SetScavengingSchedule(v ZoneauthscavengingsettingsScavengingSchedule)`

SetScavengingSchedule sets ScavengingSchedule field to given value.

### HasScavengingSchedule

`func (o *ZoneAuthScavengingSettings) HasScavengingSchedule() bool`

HasScavengingSchedule returns a boolean if a field has been set.

### GetExpressionList

`func (o *ZoneAuthScavengingSettings) GetExpressionList() []ZoneauthscavengingsettingsExpressionList`

GetExpressionList returns the ExpressionList field if non-nil, zero value otherwise.

### GetExpressionListOk

`func (o *ZoneAuthScavengingSettings) GetExpressionListOk() (*[]ZoneauthscavengingsettingsExpressionList, bool)`

GetExpressionListOk returns a tuple with the ExpressionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionList

`func (o *ZoneAuthScavengingSettings) SetExpressionList(v []ZoneauthscavengingsettingsExpressionList)`

SetExpressionList sets ExpressionList field to given value.

### HasExpressionList

`func (o *ZoneAuthScavengingSettings) HasExpressionList() bool`

HasExpressionList returns a boolean if a field has been set.

### GetEaExpressionList

`func (o *ZoneAuthScavengingSettings) GetEaExpressionList() []ZoneauthscavengingsettingsEaExpressionList`

GetEaExpressionList returns the EaExpressionList field if non-nil, zero value otherwise.

### GetEaExpressionListOk

`func (o *ZoneAuthScavengingSettings) GetEaExpressionListOk() (*[]ZoneauthscavengingsettingsEaExpressionList, bool)`

GetEaExpressionListOk returns a tuple with the EaExpressionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaExpressionList

`func (o *ZoneAuthScavengingSettings) SetEaExpressionList(v []ZoneauthscavengingsettingsEaExpressionList)`

SetEaExpressionList sets EaExpressionList field to given value.

### HasEaExpressionList

`func (o *ZoneAuthScavengingSettings) HasEaExpressionList() bool`

HasEaExpressionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


