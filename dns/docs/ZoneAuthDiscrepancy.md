# ZoneAuthDiscrepancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Description** | Pointer to **string** | Information about the discrepancy. | [optional] [readonly] 
**Severity** | Pointer to **string** | The severity of the discrepancy reported. | [optional] [readonly] 
**Timestamp** | Pointer to **int64** | The time when the DNS integrity check was last run for this zone. | [optional] [readonly] 
**Zone** | Pointer to **string** | The reference of the zone during a search. Otherwise, this is the zone object of the zone to which the discrepancy refers. | [optional] [readonly] 

## Methods

### NewZoneAuthDiscrepancy

`func NewZoneAuthDiscrepancy() *ZoneAuthDiscrepancy`

NewZoneAuthDiscrepancy instantiates a new ZoneAuthDiscrepancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthDiscrepancyWithDefaults

`func NewZoneAuthDiscrepancyWithDefaults() *ZoneAuthDiscrepancy`

NewZoneAuthDiscrepancyWithDefaults instantiates a new ZoneAuthDiscrepancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ZoneAuthDiscrepancy) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ZoneAuthDiscrepancy) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ZoneAuthDiscrepancy) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ZoneAuthDiscrepancy) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDescription

`func (o *ZoneAuthDiscrepancy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ZoneAuthDiscrepancy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ZoneAuthDiscrepancy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ZoneAuthDiscrepancy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *ZoneAuthDiscrepancy) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ZoneAuthDiscrepancy) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ZoneAuthDiscrepancy) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ZoneAuthDiscrepancy) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTimestamp

`func (o *ZoneAuthDiscrepancy) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ZoneAuthDiscrepancy) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ZoneAuthDiscrepancy) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ZoneAuthDiscrepancy) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetZone

`func (o *ZoneAuthDiscrepancy) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ZoneAuthDiscrepancy) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ZoneAuthDiscrepancy) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ZoneAuthDiscrepancy) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


