# DiscoveryStatusReachableInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The overall status of the device. | [optional] [readonly] 
**Message** | Pointer to **string** | The detailed message. | [optional] [readonly] 
**Timestamp** | Pointer to **int64** | The timestamp when the status was generated. | [optional] [readonly] 

## Methods

### NewDiscoveryStatusReachableInfo

`func NewDiscoveryStatusReachableInfo() *DiscoveryStatusReachableInfo`

NewDiscoveryStatusReachableInfo instantiates a new DiscoveryStatusReachableInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryStatusReachableInfoWithDefaults

`func NewDiscoveryStatusReachableInfoWithDefaults() *DiscoveryStatusReachableInfo`

NewDiscoveryStatusReachableInfoWithDefaults instantiates a new DiscoveryStatusReachableInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DiscoveryStatusReachableInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscoveryStatusReachableInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscoveryStatusReachableInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiscoveryStatusReachableInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *DiscoveryStatusReachableInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DiscoveryStatusReachableInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DiscoveryStatusReachableInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DiscoveryStatusReachableInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *DiscoveryStatusReachableInfo) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DiscoveryStatusReachableInfo) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DiscoveryStatusReachableInfo) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DiscoveryStatusReachableInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


