# GriddnsattackmitigationDetectUdpDrop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Determines if DNS attack detection is enabled or not. | [optional] 
**High** | Pointer to **int64** | The high threshold value (in percentage) for starting DNS attack detection. | [optional] 
**IntervalMax** | Pointer to **int64** | The maximum number of events that have occurred before processing DNS attack detection. | [optional] 
**IntervalMin** | Pointer to **int64** | The minimum number of events that have occurred before processing DNS attack detection. | [optional] 
**IntervalTime** | Pointer to **int64** | The time interval between detection processing. | [optional] 
**Low** | Pointer to **int64** | The low threshold value (in percentage) for starting DNS attack detection. | [optional] 

## Methods

### NewGriddnsattackmitigationDetectUdpDrop

`func NewGriddnsattackmitigationDetectUdpDrop() *GriddnsattackmitigationDetectUdpDrop`

NewGriddnsattackmitigationDetectUdpDrop instantiates a new GriddnsattackmitigationDetectUdpDrop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGriddnsattackmitigationDetectUdpDropWithDefaults

`func NewGriddnsattackmitigationDetectUdpDropWithDefaults() *GriddnsattackmitigationDetectUdpDrop`

NewGriddnsattackmitigationDetectUdpDropWithDefaults instantiates a new GriddnsattackmitigationDetectUdpDrop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *GriddnsattackmitigationDetectUdpDrop) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GriddnsattackmitigationDetectUdpDrop) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GriddnsattackmitigationDetectUdpDrop) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GriddnsattackmitigationDetectUdpDrop) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetHigh

`func (o *GriddnsattackmitigationDetectUdpDrop) GetHigh() int64`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *GriddnsattackmitigationDetectUdpDrop) GetHighOk() (*int64, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *GriddnsattackmitigationDetectUdpDrop) SetHigh(v int64)`

SetHigh sets High field to given value.

### HasHigh

`func (o *GriddnsattackmitigationDetectUdpDrop) HasHigh() bool`

HasHigh returns a boolean if a field has been set.

### GetIntervalMax

`func (o *GriddnsattackmitigationDetectUdpDrop) GetIntervalMax() int64`

GetIntervalMax returns the IntervalMax field if non-nil, zero value otherwise.

### GetIntervalMaxOk

`func (o *GriddnsattackmitigationDetectUdpDrop) GetIntervalMaxOk() (*int64, bool)`

GetIntervalMaxOk returns a tuple with the IntervalMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMax

`func (o *GriddnsattackmitigationDetectUdpDrop) SetIntervalMax(v int64)`

SetIntervalMax sets IntervalMax field to given value.

### HasIntervalMax

`func (o *GriddnsattackmitigationDetectUdpDrop) HasIntervalMax() bool`

HasIntervalMax returns a boolean if a field has been set.

### GetIntervalMin

`func (o *GriddnsattackmitigationDetectUdpDrop) GetIntervalMin() int64`

GetIntervalMin returns the IntervalMin field if non-nil, zero value otherwise.

### GetIntervalMinOk

`func (o *GriddnsattackmitigationDetectUdpDrop) GetIntervalMinOk() (*int64, bool)`

GetIntervalMinOk returns a tuple with the IntervalMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalMin

`func (o *GriddnsattackmitigationDetectUdpDrop) SetIntervalMin(v int64)`

SetIntervalMin sets IntervalMin field to given value.

### HasIntervalMin

`func (o *GriddnsattackmitigationDetectUdpDrop) HasIntervalMin() bool`

HasIntervalMin returns a boolean if a field has been set.

### GetIntervalTime

`func (o *GriddnsattackmitigationDetectUdpDrop) GetIntervalTime() int64`

GetIntervalTime returns the IntervalTime field if non-nil, zero value otherwise.

### GetIntervalTimeOk

`func (o *GriddnsattackmitigationDetectUdpDrop) GetIntervalTimeOk() (*int64, bool)`

GetIntervalTimeOk returns a tuple with the IntervalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalTime

`func (o *GriddnsattackmitigationDetectUdpDrop) SetIntervalTime(v int64)`

SetIntervalTime sets IntervalTime field to given value.

### HasIntervalTime

`func (o *GriddnsattackmitigationDetectUdpDrop) HasIntervalTime() bool`

HasIntervalTime returns a boolean if a field has been set.

### GetLow

`func (o *GriddnsattackmitigationDetectUdpDrop) GetLow() int64`

GetLow returns the Low field if non-nil, zero value otherwise.

### GetLowOk

`func (o *GriddnsattackmitigationDetectUdpDrop) GetLowOk() (*int64, bool)`

GetLowOk returns a tuple with the Low field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLow

`func (o *GriddnsattackmitigationDetectUdpDrop) SetLow(v int64)`

SetLow sets Low field to given value.

### HasLow

`func (o *GriddnsattackmitigationDetectUdpDrop) HasLow() bool`

HasLow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


