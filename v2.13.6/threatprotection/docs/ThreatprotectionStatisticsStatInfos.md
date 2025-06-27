# ThreatprotectionStatisticsStatInfos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **int64** | The timestamp when data was collected. | [optional] [readonly] 
**Critical** | Pointer to **map[string]interface{}** | The number of critical events. | [optional] 
**Major** | Pointer to **map[string]interface{}** | The number of major events. | [optional] 
**Warning** | Pointer to **map[string]interface{}** | The number of warning events. | [optional] 
**Informational** | Pointer to **map[string]interface{}** | The number of informational events. | [optional] 
**Total** | Pointer to **map[string]interface{}** | The total number of events. | [optional] 

## Methods

### NewThreatprotectionStatisticsStatInfos

`func NewThreatprotectionStatisticsStatInfos() *ThreatprotectionStatisticsStatInfos`

NewThreatprotectionStatisticsStatInfos instantiates a new ThreatprotectionStatisticsStatInfos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatprotectionStatisticsStatInfosWithDefaults

`func NewThreatprotectionStatisticsStatInfosWithDefaults() *ThreatprotectionStatisticsStatInfos`

NewThreatprotectionStatisticsStatInfosWithDefaults instantiates a new ThreatprotectionStatisticsStatInfos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ThreatprotectionStatisticsStatInfos) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ThreatprotectionStatisticsStatInfos) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ThreatprotectionStatisticsStatInfos) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ThreatprotectionStatisticsStatInfos) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetCritical

`func (o *ThreatprotectionStatisticsStatInfos) GetCritical() map[string]interface{}`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *ThreatprotectionStatisticsStatInfos) GetCriticalOk() (*map[string]interface{}, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *ThreatprotectionStatisticsStatInfos) SetCritical(v map[string]interface{})`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *ThreatprotectionStatisticsStatInfos) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetMajor

`func (o *ThreatprotectionStatisticsStatInfos) GetMajor() map[string]interface{}`

GetMajor returns the Major field if non-nil, zero value otherwise.

### GetMajorOk

`func (o *ThreatprotectionStatisticsStatInfos) GetMajorOk() (*map[string]interface{}, bool)`

GetMajorOk returns a tuple with the Major field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajor

`func (o *ThreatprotectionStatisticsStatInfos) SetMajor(v map[string]interface{})`

SetMajor sets Major field to given value.

### HasMajor

`func (o *ThreatprotectionStatisticsStatInfos) HasMajor() bool`

HasMajor returns a boolean if a field has been set.

### GetWarning

`func (o *ThreatprotectionStatisticsStatInfos) GetWarning() map[string]interface{}`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *ThreatprotectionStatisticsStatInfos) GetWarningOk() (*map[string]interface{}, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *ThreatprotectionStatisticsStatInfos) SetWarning(v map[string]interface{})`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *ThreatprotectionStatisticsStatInfos) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetInformational

`func (o *ThreatprotectionStatisticsStatInfos) GetInformational() map[string]interface{}`

GetInformational returns the Informational field if non-nil, zero value otherwise.

### GetInformationalOk

`func (o *ThreatprotectionStatisticsStatInfos) GetInformationalOk() (*map[string]interface{}, bool)`

GetInformationalOk returns a tuple with the Informational field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInformational

`func (o *ThreatprotectionStatisticsStatInfos) SetInformational(v map[string]interface{})`

SetInformational sets Informational field to given value.

### HasInformational

`func (o *ThreatprotectionStatisticsStatInfos) HasInformational() bool`

HasInformational returns a boolean if a field has been set.

### GetTotal

`func (o *ThreatprotectionStatisticsStatInfos) GetTotal() map[string]interface{}`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ThreatprotectionStatisticsStatInfos) GetTotalOk() (*map[string]interface{}, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ThreatprotectionStatisticsStatInfos) SetTotal(v map[string]interface{})`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ThreatprotectionStatisticsStatInfos) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


