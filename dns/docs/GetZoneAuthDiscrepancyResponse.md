# GetZoneAuthDiscrepancyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Description** | Pointer to **string** | Information about the discrepancy. | [optional] [readonly] 
**Severity** | Pointer to **string** | The severity of the discrepancy reported. | [optional] [readonly] 
**Timestamp** | Pointer to **int64** | The time when the DNS integrity check was last run for this zone. | [optional] [readonly] 
**Zone** | Pointer to **string** | The reference of the zone during a search. Otherwise, this is the zone object of the zone to which the discrepancy refers. | [optional] [readonly] 
**Result** | Pointer to [**ZoneAuthDiscrepancy**](ZoneAuthDiscrepancy.md) |  | [optional] 

## Methods

### NewGetZoneAuthDiscrepancyResponse

`func NewGetZoneAuthDiscrepancyResponse() *GetZoneAuthDiscrepancyResponse`

NewGetZoneAuthDiscrepancyResponse instantiates a new GetZoneAuthDiscrepancyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetZoneAuthDiscrepancyResponseWithDefaults

`func NewGetZoneAuthDiscrepancyResponseWithDefaults() *GetZoneAuthDiscrepancyResponse`

NewGetZoneAuthDiscrepancyResponseWithDefaults instantiates a new GetZoneAuthDiscrepancyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetZoneAuthDiscrepancyResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetZoneAuthDiscrepancyResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetZoneAuthDiscrepancyResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetZoneAuthDiscrepancyResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDescription

`func (o *GetZoneAuthDiscrepancyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetZoneAuthDiscrepancyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetZoneAuthDiscrepancyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetZoneAuthDiscrepancyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSeverity

`func (o *GetZoneAuthDiscrepancyResponse) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetZoneAuthDiscrepancyResponse) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetZoneAuthDiscrepancyResponse) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetZoneAuthDiscrepancyResponse) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTimestamp

`func (o *GetZoneAuthDiscrepancyResponse) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *GetZoneAuthDiscrepancyResponse) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *GetZoneAuthDiscrepancyResponse) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *GetZoneAuthDiscrepancyResponse) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetZone

`func (o *GetZoneAuthDiscrepancyResponse) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetZoneAuthDiscrepancyResponse) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetZoneAuthDiscrepancyResponse) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetZoneAuthDiscrepancyResponse) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetResult

`func (o *GetZoneAuthDiscrepancyResponse) GetResult() ZoneAuthDiscrepancy`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetZoneAuthDiscrepancyResponse) GetResultOk() (*ZoneAuthDiscrepancy, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetZoneAuthDiscrepancyResponse) SetResult(v ZoneAuthDiscrepancy)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetZoneAuthDiscrepancyResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


