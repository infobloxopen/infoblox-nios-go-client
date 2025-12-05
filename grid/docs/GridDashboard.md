# GridDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AnalyticsTunnelingEventCriticalThreshold** | Pointer to **int64** | The Grid Dashboard critical threshold for Analytics tunneling events. | [optional] 
**AnalyticsTunnelingEventWarningThreshold** | Pointer to **int64** | The Grid Dashboard warning threshold for Analytics tunneling events. | [optional] 
**AtpCriticalEventCriticalThreshold** | Pointer to **int64** | The Grid Dashboard critical threshold for ATP critical events. | [optional] 
**AtpCriticalEventWarningThreshold** | Pointer to **int64** | The Grid Dashboard warning threshold for ATP critical events. | [optional] 
**AtpMajorEventCriticalThreshold** | Pointer to **int64** | The Grid Dashboard critical threshold for ATP major events. | [optional] 
**AtpMajorEventWarningThreshold** | Pointer to **int64** | The Grid Dashboard warning threshold for ATP major events. | [optional] 
**AtpWarningEventCriticalThreshold** | Pointer to **int64** | The Grid Dashboard critical threshold for ATP warning events. | [optional] 
**AtpWarningEventWarningThreshold** | Pointer to **int64** | The Grid Dashboard warning threshold for ATP warning events. | [optional] 
**RpzBlockedHitCriticalThreshold** | Pointer to **int64** | The critical threshold value for blocked RPZ hits in the Grid dashboard. | [optional] 
**RpzBlockedHitWarningThreshold** | Pointer to **int64** | The warning threshold value for blocked RPZ hits in the Grid dashboard. | [optional] 
**RpzPassthruEventCriticalThreshold** | Pointer to **int64** | The Grid Dashboard critical threshold for RPZ passthru events. | [optional] 
**RpzPassthruEventWarningThreshold** | Pointer to **int64** | The Grid Dashboard warning threshold for RPZ passthru events. | [optional] 
**RpzSubstitutedHitCriticalThreshold** | Pointer to **int64** | The critical threshold value for substituted RPZ hits in the Grid dashboard. | [optional] 
**RpzSubstitutedHitWarningThreshold** | Pointer to **int64** | The warning threshold value for substituted RPZ hits in the Grid dashboard. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewGridDashboard

`func NewGridDashboard() *GridDashboard`

NewGridDashboard instantiates a new GridDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDashboardWithDefaults

`func NewGridDashboardWithDefaults() *GridDashboard`

NewGridDashboardWithDefaults instantiates a new GridDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridDashboard) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridDashboard) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridDashboard) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridDashboard) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAnalyticsTunnelingEventCriticalThreshold

`func (o *GridDashboard) GetAnalyticsTunnelingEventCriticalThreshold() int64`

GetAnalyticsTunnelingEventCriticalThreshold returns the AnalyticsTunnelingEventCriticalThreshold field if non-nil, zero value otherwise.

### GetAnalyticsTunnelingEventCriticalThresholdOk

`func (o *GridDashboard) GetAnalyticsTunnelingEventCriticalThresholdOk() (*int64, bool)`

GetAnalyticsTunnelingEventCriticalThresholdOk returns a tuple with the AnalyticsTunnelingEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsTunnelingEventCriticalThreshold

`func (o *GridDashboard) SetAnalyticsTunnelingEventCriticalThreshold(v int64)`

SetAnalyticsTunnelingEventCriticalThreshold sets AnalyticsTunnelingEventCriticalThreshold field to given value.

### HasAnalyticsTunnelingEventCriticalThreshold

`func (o *GridDashboard) HasAnalyticsTunnelingEventCriticalThreshold() bool`

HasAnalyticsTunnelingEventCriticalThreshold returns a boolean if a field has been set.

### GetAnalyticsTunnelingEventWarningThreshold

`func (o *GridDashboard) GetAnalyticsTunnelingEventWarningThreshold() int64`

GetAnalyticsTunnelingEventWarningThreshold returns the AnalyticsTunnelingEventWarningThreshold field if non-nil, zero value otherwise.

### GetAnalyticsTunnelingEventWarningThresholdOk

`func (o *GridDashboard) GetAnalyticsTunnelingEventWarningThresholdOk() (*int64, bool)`

GetAnalyticsTunnelingEventWarningThresholdOk returns a tuple with the AnalyticsTunnelingEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsTunnelingEventWarningThreshold

`func (o *GridDashboard) SetAnalyticsTunnelingEventWarningThreshold(v int64)`

SetAnalyticsTunnelingEventWarningThreshold sets AnalyticsTunnelingEventWarningThreshold field to given value.

### HasAnalyticsTunnelingEventWarningThreshold

`func (o *GridDashboard) HasAnalyticsTunnelingEventWarningThreshold() bool`

HasAnalyticsTunnelingEventWarningThreshold returns a boolean if a field has been set.

### GetAtpCriticalEventCriticalThreshold

`func (o *GridDashboard) GetAtpCriticalEventCriticalThreshold() int64`

GetAtpCriticalEventCriticalThreshold returns the AtpCriticalEventCriticalThreshold field if non-nil, zero value otherwise.

### GetAtpCriticalEventCriticalThresholdOk

`func (o *GridDashboard) GetAtpCriticalEventCriticalThresholdOk() (*int64, bool)`

GetAtpCriticalEventCriticalThresholdOk returns a tuple with the AtpCriticalEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpCriticalEventCriticalThreshold

`func (o *GridDashboard) SetAtpCriticalEventCriticalThreshold(v int64)`

SetAtpCriticalEventCriticalThreshold sets AtpCriticalEventCriticalThreshold field to given value.

### HasAtpCriticalEventCriticalThreshold

`func (o *GridDashboard) HasAtpCriticalEventCriticalThreshold() bool`

HasAtpCriticalEventCriticalThreshold returns a boolean if a field has been set.

### GetAtpCriticalEventWarningThreshold

`func (o *GridDashboard) GetAtpCriticalEventWarningThreshold() int64`

GetAtpCriticalEventWarningThreshold returns the AtpCriticalEventWarningThreshold field if non-nil, zero value otherwise.

### GetAtpCriticalEventWarningThresholdOk

`func (o *GridDashboard) GetAtpCriticalEventWarningThresholdOk() (*int64, bool)`

GetAtpCriticalEventWarningThresholdOk returns a tuple with the AtpCriticalEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpCriticalEventWarningThreshold

`func (o *GridDashboard) SetAtpCriticalEventWarningThreshold(v int64)`

SetAtpCriticalEventWarningThreshold sets AtpCriticalEventWarningThreshold field to given value.

### HasAtpCriticalEventWarningThreshold

`func (o *GridDashboard) HasAtpCriticalEventWarningThreshold() bool`

HasAtpCriticalEventWarningThreshold returns a boolean if a field has been set.

### GetAtpMajorEventCriticalThreshold

`func (o *GridDashboard) GetAtpMajorEventCriticalThreshold() int64`

GetAtpMajorEventCriticalThreshold returns the AtpMajorEventCriticalThreshold field if non-nil, zero value otherwise.

### GetAtpMajorEventCriticalThresholdOk

`func (o *GridDashboard) GetAtpMajorEventCriticalThresholdOk() (*int64, bool)`

GetAtpMajorEventCriticalThresholdOk returns a tuple with the AtpMajorEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpMajorEventCriticalThreshold

`func (o *GridDashboard) SetAtpMajorEventCriticalThreshold(v int64)`

SetAtpMajorEventCriticalThreshold sets AtpMajorEventCriticalThreshold field to given value.

### HasAtpMajorEventCriticalThreshold

`func (o *GridDashboard) HasAtpMajorEventCriticalThreshold() bool`

HasAtpMajorEventCriticalThreshold returns a boolean if a field has been set.

### GetAtpMajorEventWarningThreshold

`func (o *GridDashboard) GetAtpMajorEventWarningThreshold() int64`

GetAtpMajorEventWarningThreshold returns the AtpMajorEventWarningThreshold field if non-nil, zero value otherwise.

### GetAtpMajorEventWarningThresholdOk

`func (o *GridDashboard) GetAtpMajorEventWarningThresholdOk() (*int64, bool)`

GetAtpMajorEventWarningThresholdOk returns a tuple with the AtpMajorEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpMajorEventWarningThreshold

`func (o *GridDashboard) SetAtpMajorEventWarningThreshold(v int64)`

SetAtpMajorEventWarningThreshold sets AtpMajorEventWarningThreshold field to given value.

### HasAtpMajorEventWarningThreshold

`func (o *GridDashboard) HasAtpMajorEventWarningThreshold() bool`

HasAtpMajorEventWarningThreshold returns a boolean if a field has been set.

### GetAtpWarningEventCriticalThreshold

`func (o *GridDashboard) GetAtpWarningEventCriticalThreshold() int64`

GetAtpWarningEventCriticalThreshold returns the AtpWarningEventCriticalThreshold field if non-nil, zero value otherwise.

### GetAtpWarningEventCriticalThresholdOk

`func (o *GridDashboard) GetAtpWarningEventCriticalThresholdOk() (*int64, bool)`

GetAtpWarningEventCriticalThresholdOk returns a tuple with the AtpWarningEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpWarningEventCriticalThreshold

`func (o *GridDashboard) SetAtpWarningEventCriticalThreshold(v int64)`

SetAtpWarningEventCriticalThreshold sets AtpWarningEventCriticalThreshold field to given value.

### HasAtpWarningEventCriticalThreshold

`func (o *GridDashboard) HasAtpWarningEventCriticalThreshold() bool`

HasAtpWarningEventCriticalThreshold returns a boolean if a field has been set.

### GetAtpWarningEventWarningThreshold

`func (o *GridDashboard) GetAtpWarningEventWarningThreshold() int64`

GetAtpWarningEventWarningThreshold returns the AtpWarningEventWarningThreshold field if non-nil, zero value otherwise.

### GetAtpWarningEventWarningThresholdOk

`func (o *GridDashboard) GetAtpWarningEventWarningThresholdOk() (*int64, bool)`

GetAtpWarningEventWarningThresholdOk returns a tuple with the AtpWarningEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpWarningEventWarningThreshold

`func (o *GridDashboard) SetAtpWarningEventWarningThreshold(v int64)`

SetAtpWarningEventWarningThreshold sets AtpWarningEventWarningThreshold field to given value.

### HasAtpWarningEventWarningThreshold

`func (o *GridDashboard) HasAtpWarningEventWarningThreshold() bool`

HasAtpWarningEventWarningThreshold returns a boolean if a field has been set.

### GetRpzBlockedHitCriticalThreshold

`func (o *GridDashboard) GetRpzBlockedHitCriticalThreshold() int64`

GetRpzBlockedHitCriticalThreshold returns the RpzBlockedHitCriticalThreshold field if non-nil, zero value otherwise.

### GetRpzBlockedHitCriticalThresholdOk

`func (o *GridDashboard) GetRpzBlockedHitCriticalThresholdOk() (*int64, bool)`

GetRpzBlockedHitCriticalThresholdOk returns a tuple with the RpzBlockedHitCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzBlockedHitCriticalThreshold

`func (o *GridDashboard) SetRpzBlockedHitCriticalThreshold(v int64)`

SetRpzBlockedHitCriticalThreshold sets RpzBlockedHitCriticalThreshold field to given value.

### HasRpzBlockedHitCriticalThreshold

`func (o *GridDashboard) HasRpzBlockedHitCriticalThreshold() bool`

HasRpzBlockedHitCriticalThreshold returns a boolean if a field has been set.

### GetRpzBlockedHitWarningThreshold

`func (o *GridDashboard) GetRpzBlockedHitWarningThreshold() int64`

GetRpzBlockedHitWarningThreshold returns the RpzBlockedHitWarningThreshold field if non-nil, zero value otherwise.

### GetRpzBlockedHitWarningThresholdOk

`func (o *GridDashboard) GetRpzBlockedHitWarningThresholdOk() (*int64, bool)`

GetRpzBlockedHitWarningThresholdOk returns a tuple with the RpzBlockedHitWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzBlockedHitWarningThreshold

`func (o *GridDashboard) SetRpzBlockedHitWarningThreshold(v int64)`

SetRpzBlockedHitWarningThreshold sets RpzBlockedHitWarningThreshold field to given value.

### HasRpzBlockedHitWarningThreshold

`func (o *GridDashboard) HasRpzBlockedHitWarningThreshold() bool`

HasRpzBlockedHitWarningThreshold returns a boolean if a field has been set.

### GetRpzPassthruEventCriticalThreshold

`func (o *GridDashboard) GetRpzPassthruEventCriticalThreshold() int64`

GetRpzPassthruEventCriticalThreshold returns the RpzPassthruEventCriticalThreshold field if non-nil, zero value otherwise.

### GetRpzPassthruEventCriticalThresholdOk

`func (o *GridDashboard) GetRpzPassthruEventCriticalThresholdOk() (*int64, bool)`

GetRpzPassthruEventCriticalThresholdOk returns a tuple with the RpzPassthruEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPassthruEventCriticalThreshold

`func (o *GridDashboard) SetRpzPassthruEventCriticalThreshold(v int64)`

SetRpzPassthruEventCriticalThreshold sets RpzPassthruEventCriticalThreshold field to given value.

### HasRpzPassthruEventCriticalThreshold

`func (o *GridDashboard) HasRpzPassthruEventCriticalThreshold() bool`

HasRpzPassthruEventCriticalThreshold returns a boolean if a field has been set.

### GetRpzPassthruEventWarningThreshold

`func (o *GridDashboard) GetRpzPassthruEventWarningThreshold() int64`

GetRpzPassthruEventWarningThreshold returns the RpzPassthruEventWarningThreshold field if non-nil, zero value otherwise.

### GetRpzPassthruEventWarningThresholdOk

`func (o *GridDashboard) GetRpzPassthruEventWarningThresholdOk() (*int64, bool)`

GetRpzPassthruEventWarningThresholdOk returns a tuple with the RpzPassthruEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPassthruEventWarningThreshold

`func (o *GridDashboard) SetRpzPassthruEventWarningThreshold(v int64)`

SetRpzPassthruEventWarningThreshold sets RpzPassthruEventWarningThreshold field to given value.

### HasRpzPassthruEventWarningThreshold

`func (o *GridDashboard) HasRpzPassthruEventWarningThreshold() bool`

HasRpzPassthruEventWarningThreshold returns a boolean if a field has been set.

### GetRpzSubstitutedHitCriticalThreshold

`func (o *GridDashboard) GetRpzSubstitutedHitCriticalThreshold() int64`

GetRpzSubstitutedHitCriticalThreshold returns the RpzSubstitutedHitCriticalThreshold field if non-nil, zero value otherwise.

### GetRpzSubstitutedHitCriticalThresholdOk

`func (o *GridDashboard) GetRpzSubstitutedHitCriticalThresholdOk() (*int64, bool)`

GetRpzSubstitutedHitCriticalThresholdOk returns a tuple with the RpzSubstitutedHitCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzSubstitutedHitCriticalThreshold

`func (o *GridDashboard) SetRpzSubstitutedHitCriticalThreshold(v int64)`

SetRpzSubstitutedHitCriticalThreshold sets RpzSubstitutedHitCriticalThreshold field to given value.

### HasRpzSubstitutedHitCriticalThreshold

`func (o *GridDashboard) HasRpzSubstitutedHitCriticalThreshold() bool`

HasRpzSubstitutedHitCriticalThreshold returns a boolean if a field has been set.

### GetRpzSubstitutedHitWarningThreshold

`func (o *GridDashboard) GetRpzSubstitutedHitWarningThreshold() int64`

GetRpzSubstitutedHitWarningThreshold returns the RpzSubstitutedHitWarningThreshold field if non-nil, zero value otherwise.

### GetRpzSubstitutedHitWarningThresholdOk

`func (o *GridDashboard) GetRpzSubstitutedHitWarningThresholdOk() (*int64, bool)`

GetRpzSubstitutedHitWarningThresholdOk returns a tuple with the RpzSubstitutedHitWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzSubstitutedHitWarningThreshold

`func (o *GridDashboard) SetRpzSubstitutedHitWarningThreshold(v int64)`

SetRpzSubstitutedHitWarningThreshold sets RpzSubstitutedHitWarningThreshold field to given value.

### HasRpzSubstitutedHitWarningThreshold

`func (o *GridDashboard) HasRpzSubstitutedHitWarningThreshold() bool`

HasRpzSubstitutedHitWarningThreshold returns a boolean if a field has been set.

### GetUuid

`func (o *GridDashboard) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GridDashboard) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GridDashboard) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GridDashboard) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


