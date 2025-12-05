# GetGridDashboardResponse

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
**Result** | Pointer to [**GridDashboard**](GridDashboard.md) |  | [optional] 

## Methods

### NewGetGridDashboardResponse

`func NewGetGridDashboardResponse() *GetGridDashboardResponse`

NewGetGridDashboardResponse instantiates a new GetGridDashboardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridDashboardResponseWithDefaults

`func NewGetGridDashboardResponseWithDefaults() *GetGridDashboardResponse`

NewGetGridDashboardResponseWithDefaults instantiates a new GetGridDashboardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridDashboardResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridDashboardResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridDashboardResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridDashboardResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAnalyticsTunnelingEventCriticalThreshold

`func (o *GetGridDashboardResponse) GetAnalyticsTunnelingEventCriticalThreshold() int64`

GetAnalyticsTunnelingEventCriticalThreshold returns the AnalyticsTunnelingEventCriticalThreshold field if non-nil, zero value otherwise.

### GetAnalyticsTunnelingEventCriticalThresholdOk

`func (o *GetGridDashboardResponse) GetAnalyticsTunnelingEventCriticalThresholdOk() (*int64, bool)`

GetAnalyticsTunnelingEventCriticalThresholdOk returns a tuple with the AnalyticsTunnelingEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsTunnelingEventCriticalThreshold

`func (o *GetGridDashboardResponse) SetAnalyticsTunnelingEventCriticalThreshold(v int64)`

SetAnalyticsTunnelingEventCriticalThreshold sets AnalyticsTunnelingEventCriticalThreshold field to given value.

### HasAnalyticsTunnelingEventCriticalThreshold

`func (o *GetGridDashboardResponse) HasAnalyticsTunnelingEventCriticalThreshold() bool`

HasAnalyticsTunnelingEventCriticalThreshold returns a boolean if a field has been set.

### GetAnalyticsTunnelingEventWarningThreshold

`func (o *GetGridDashboardResponse) GetAnalyticsTunnelingEventWarningThreshold() int64`

GetAnalyticsTunnelingEventWarningThreshold returns the AnalyticsTunnelingEventWarningThreshold field if non-nil, zero value otherwise.

### GetAnalyticsTunnelingEventWarningThresholdOk

`func (o *GetGridDashboardResponse) GetAnalyticsTunnelingEventWarningThresholdOk() (*int64, bool)`

GetAnalyticsTunnelingEventWarningThresholdOk returns a tuple with the AnalyticsTunnelingEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsTunnelingEventWarningThreshold

`func (o *GetGridDashboardResponse) SetAnalyticsTunnelingEventWarningThreshold(v int64)`

SetAnalyticsTunnelingEventWarningThreshold sets AnalyticsTunnelingEventWarningThreshold field to given value.

### HasAnalyticsTunnelingEventWarningThreshold

`func (o *GetGridDashboardResponse) HasAnalyticsTunnelingEventWarningThreshold() bool`

HasAnalyticsTunnelingEventWarningThreshold returns a boolean if a field has been set.

### GetAtpCriticalEventCriticalThreshold

`func (o *GetGridDashboardResponse) GetAtpCriticalEventCriticalThreshold() int64`

GetAtpCriticalEventCriticalThreshold returns the AtpCriticalEventCriticalThreshold field if non-nil, zero value otherwise.

### GetAtpCriticalEventCriticalThresholdOk

`func (o *GetGridDashboardResponse) GetAtpCriticalEventCriticalThresholdOk() (*int64, bool)`

GetAtpCriticalEventCriticalThresholdOk returns a tuple with the AtpCriticalEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpCriticalEventCriticalThreshold

`func (o *GetGridDashboardResponse) SetAtpCriticalEventCriticalThreshold(v int64)`

SetAtpCriticalEventCriticalThreshold sets AtpCriticalEventCriticalThreshold field to given value.

### HasAtpCriticalEventCriticalThreshold

`func (o *GetGridDashboardResponse) HasAtpCriticalEventCriticalThreshold() bool`

HasAtpCriticalEventCriticalThreshold returns a boolean if a field has been set.

### GetAtpCriticalEventWarningThreshold

`func (o *GetGridDashboardResponse) GetAtpCriticalEventWarningThreshold() int64`

GetAtpCriticalEventWarningThreshold returns the AtpCriticalEventWarningThreshold field if non-nil, zero value otherwise.

### GetAtpCriticalEventWarningThresholdOk

`func (o *GetGridDashboardResponse) GetAtpCriticalEventWarningThresholdOk() (*int64, bool)`

GetAtpCriticalEventWarningThresholdOk returns a tuple with the AtpCriticalEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpCriticalEventWarningThreshold

`func (o *GetGridDashboardResponse) SetAtpCriticalEventWarningThreshold(v int64)`

SetAtpCriticalEventWarningThreshold sets AtpCriticalEventWarningThreshold field to given value.

### HasAtpCriticalEventWarningThreshold

`func (o *GetGridDashboardResponse) HasAtpCriticalEventWarningThreshold() bool`

HasAtpCriticalEventWarningThreshold returns a boolean if a field has been set.

### GetAtpMajorEventCriticalThreshold

`func (o *GetGridDashboardResponse) GetAtpMajorEventCriticalThreshold() int64`

GetAtpMajorEventCriticalThreshold returns the AtpMajorEventCriticalThreshold field if non-nil, zero value otherwise.

### GetAtpMajorEventCriticalThresholdOk

`func (o *GetGridDashboardResponse) GetAtpMajorEventCriticalThresholdOk() (*int64, bool)`

GetAtpMajorEventCriticalThresholdOk returns a tuple with the AtpMajorEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpMajorEventCriticalThreshold

`func (o *GetGridDashboardResponse) SetAtpMajorEventCriticalThreshold(v int64)`

SetAtpMajorEventCriticalThreshold sets AtpMajorEventCriticalThreshold field to given value.

### HasAtpMajorEventCriticalThreshold

`func (o *GetGridDashboardResponse) HasAtpMajorEventCriticalThreshold() bool`

HasAtpMajorEventCriticalThreshold returns a boolean if a field has been set.

### GetAtpMajorEventWarningThreshold

`func (o *GetGridDashboardResponse) GetAtpMajorEventWarningThreshold() int64`

GetAtpMajorEventWarningThreshold returns the AtpMajorEventWarningThreshold field if non-nil, zero value otherwise.

### GetAtpMajorEventWarningThresholdOk

`func (o *GetGridDashboardResponse) GetAtpMajorEventWarningThresholdOk() (*int64, bool)`

GetAtpMajorEventWarningThresholdOk returns a tuple with the AtpMajorEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpMajorEventWarningThreshold

`func (o *GetGridDashboardResponse) SetAtpMajorEventWarningThreshold(v int64)`

SetAtpMajorEventWarningThreshold sets AtpMajorEventWarningThreshold field to given value.

### HasAtpMajorEventWarningThreshold

`func (o *GetGridDashboardResponse) HasAtpMajorEventWarningThreshold() bool`

HasAtpMajorEventWarningThreshold returns a boolean if a field has been set.

### GetAtpWarningEventCriticalThreshold

`func (o *GetGridDashboardResponse) GetAtpWarningEventCriticalThreshold() int64`

GetAtpWarningEventCriticalThreshold returns the AtpWarningEventCriticalThreshold field if non-nil, zero value otherwise.

### GetAtpWarningEventCriticalThresholdOk

`func (o *GetGridDashboardResponse) GetAtpWarningEventCriticalThresholdOk() (*int64, bool)`

GetAtpWarningEventCriticalThresholdOk returns a tuple with the AtpWarningEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpWarningEventCriticalThreshold

`func (o *GetGridDashboardResponse) SetAtpWarningEventCriticalThreshold(v int64)`

SetAtpWarningEventCriticalThreshold sets AtpWarningEventCriticalThreshold field to given value.

### HasAtpWarningEventCriticalThreshold

`func (o *GetGridDashboardResponse) HasAtpWarningEventCriticalThreshold() bool`

HasAtpWarningEventCriticalThreshold returns a boolean if a field has been set.

### GetAtpWarningEventWarningThreshold

`func (o *GetGridDashboardResponse) GetAtpWarningEventWarningThreshold() int64`

GetAtpWarningEventWarningThreshold returns the AtpWarningEventWarningThreshold field if non-nil, zero value otherwise.

### GetAtpWarningEventWarningThresholdOk

`func (o *GetGridDashboardResponse) GetAtpWarningEventWarningThresholdOk() (*int64, bool)`

GetAtpWarningEventWarningThresholdOk returns a tuple with the AtpWarningEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtpWarningEventWarningThreshold

`func (o *GetGridDashboardResponse) SetAtpWarningEventWarningThreshold(v int64)`

SetAtpWarningEventWarningThreshold sets AtpWarningEventWarningThreshold field to given value.

### HasAtpWarningEventWarningThreshold

`func (o *GetGridDashboardResponse) HasAtpWarningEventWarningThreshold() bool`

HasAtpWarningEventWarningThreshold returns a boolean if a field has been set.

### GetRpzBlockedHitCriticalThreshold

`func (o *GetGridDashboardResponse) GetRpzBlockedHitCriticalThreshold() int64`

GetRpzBlockedHitCriticalThreshold returns the RpzBlockedHitCriticalThreshold field if non-nil, zero value otherwise.

### GetRpzBlockedHitCriticalThresholdOk

`func (o *GetGridDashboardResponse) GetRpzBlockedHitCriticalThresholdOk() (*int64, bool)`

GetRpzBlockedHitCriticalThresholdOk returns a tuple with the RpzBlockedHitCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzBlockedHitCriticalThreshold

`func (o *GetGridDashboardResponse) SetRpzBlockedHitCriticalThreshold(v int64)`

SetRpzBlockedHitCriticalThreshold sets RpzBlockedHitCriticalThreshold field to given value.

### HasRpzBlockedHitCriticalThreshold

`func (o *GetGridDashboardResponse) HasRpzBlockedHitCriticalThreshold() bool`

HasRpzBlockedHitCriticalThreshold returns a boolean if a field has been set.

### GetRpzBlockedHitWarningThreshold

`func (o *GetGridDashboardResponse) GetRpzBlockedHitWarningThreshold() int64`

GetRpzBlockedHitWarningThreshold returns the RpzBlockedHitWarningThreshold field if non-nil, zero value otherwise.

### GetRpzBlockedHitWarningThresholdOk

`func (o *GetGridDashboardResponse) GetRpzBlockedHitWarningThresholdOk() (*int64, bool)`

GetRpzBlockedHitWarningThresholdOk returns a tuple with the RpzBlockedHitWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzBlockedHitWarningThreshold

`func (o *GetGridDashboardResponse) SetRpzBlockedHitWarningThreshold(v int64)`

SetRpzBlockedHitWarningThreshold sets RpzBlockedHitWarningThreshold field to given value.

### HasRpzBlockedHitWarningThreshold

`func (o *GetGridDashboardResponse) HasRpzBlockedHitWarningThreshold() bool`

HasRpzBlockedHitWarningThreshold returns a boolean if a field has been set.

### GetRpzPassthruEventCriticalThreshold

`func (o *GetGridDashboardResponse) GetRpzPassthruEventCriticalThreshold() int64`

GetRpzPassthruEventCriticalThreshold returns the RpzPassthruEventCriticalThreshold field if non-nil, zero value otherwise.

### GetRpzPassthruEventCriticalThresholdOk

`func (o *GetGridDashboardResponse) GetRpzPassthruEventCriticalThresholdOk() (*int64, bool)`

GetRpzPassthruEventCriticalThresholdOk returns a tuple with the RpzPassthruEventCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPassthruEventCriticalThreshold

`func (o *GetGridDashboardResponse) SetRpzPassthruEventCriticalThreshold(v int64)`

SetRpzPassthruEventCriticalThreshold sets RpzPassthruEventCriticalThreshold field to given value.

### HasRpzPassthruEventCriticalThreshold

`func (o *GetGridDashboardResponse) HasRpzPassthruEventCriticalThreshold() bool`

HasRpzPassthruEventCriticalThreshold returns a boolean if a field has been set.

### GetRpzPassthruEventWarningThreshold

`func (o *GetGridDashboardResponse) GetRpzPassthruEventWarningThreshold() int64`

GetRpzPassthruEventWarningThreshold returns the RpzPassthruEventWarningThreshold field if non-nil, zero value otherwise.

### GetRpzPassthruEventWarningThresholdOk

`func (o *GetGridDashboardResponse) GetRpzPassthruEventWarningThresholdOk() (*int64, bool)`

GetRpzPassthruEventWarningThresholdOk returns a tuple with the RpzPassthruEventWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzPassthruEventWarningThreshold

`func (o *GetGridDashboardResponse) SetRpzPassthruEventWarningThreshold(v int64)`

SetRpzPassthruEventWarningThreshold sets RpzPassthruEventWarningThreshold field to given value.

### HasRpzPassthruEventWarningThreshold

`func (o *GetGridDashboardResponse) HasRpzPassthruEventWarningThreshold() bool`

HasRpzPassthruEventWarningThreshold returns a boolean if a field has been set.

### GetRpzSubstitutedHitCriticalThreshold

`func (o *GetGridDashboardResponse) GetRpzSubstitutedHitCriticalThreshold() int64`

GetRpzSubstitutedHitCriticalThreshold returns the RpzSubstitutedHitCriticalThreshold field if non-nil, zero value otherwise.

### GetRpzSubstitutedHitCriticalThresholdOk

`func (o *GetGridDashboardResponse) GetRpzSubstitutedHitCriticalThresholdOk() (*int64, bool)`

GetRpzSubstitutedHitCriticalThresholdOk returns a tuple with the RpzSubstitutedHitCriticalThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzSubstitutedHitCriticalThreshold

`func (o *GetGridDashboardResponse) SetRpzSubstitutedHitCriticalThreshold(v int64)`

SetRpzSubstitutedHitCriticalThreshold sets RpzSubstitutedHitCriticalThreshold field to given value.

### HasRpzSubstitutedHitCriticalThreshold

`func (o *GetGridDashboardResponse) HasRpzSubstitutedHitCriticalThreshold() bool`

HasRpzSubstitutedHitCriticalThreshold returns a boolean if a field has been set.

### GetRpzSubstitutedHitWarningThreshold

`func (o *GetGridDashboardResponse) GetRpzSubstitutedHitWarningThreshold() int64`

GetRpzSubstitutedHitWarningThreshold returns the RpzSubstitutedHitWarningThreshold field if non-nil, zero value otherwise.

### GetRpzSubstitutedHitWarningThresholdOk

`func (o *GetGridDashboardResponse) GetRpzSubstitutedHitWarningThresholdOk() (*int64, bool)`

GetRpzSubstitutedHitWarningThresholdOk returns a tuple with the RpzSubstitutedHitWarningThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpzSubstitutedHitWarningThreshold

`func (o *GetGridDashboardResponse) SetRpzSubstitutedHitWarningThreshold(v int64)`

SetRpzSubstitutedHitWarningThreshold sets RpzSubstitutedHitWarningThreshold field to given value.

### HasRpzSubstitutedHitWarningThreshold

`func (o *GetGridDashboardResponse) HasRpzSubstitutedHitWarningThreshold() bool`

HasRpzSubstitutedHitWarningThreshold returns a boolean if a field has been set.

### GetUuid

`func (o *GetGridDashboardResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGridDashboardResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGridDashboardResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGridDashboardResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetGridDashboardResponse) GetResult() GridDashboard`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridDashboardResponse) GetResultOk() (*GridDashboard, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridDashboardResponse) SetResult(v GridDashboard)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridDashboardResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


