# DiscoveryStatusSdnCollectionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The overall status of the device. | [optional] [readonly] 
**Message** | Pointer to **string** | The detailed message. | [optional] [readonly] 
**Timestamp** | Pointer to **int64** | The timestamp when the status was generated. | [optional] [readonly] 

## Methods

### NewDiscoveryStatusSdnCollectionInfo

`func NewDiscoveryStatusSdnCollectionInfo() *DiscoveryStatusSdnCollectionInfo`

NewDiscoveryStatusSdnCollectionInfo instantiates a new DiscoveryStatusSdnCollectionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryStatusSdnCollectionInfoWithDefaults

`func NewDiscoveryStatusSdnCollectionInfoWithDefaults() *DiscoveryStatusSdnCollectionInfo`

NewDiscoveryStatusSdnCollectionInfoWithDefaults instantiates a new DiscoveryStatusSdnCollectionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DiscoveryStatusSdnCollectionInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscoveryStatusSdnCollectionInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscoveryStatusSdnCollectionInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiscoveryStatusSdnCollectionInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *DiscoveryStatusSdnCollectionInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DiscoveryStatusSdnCollectionInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DiscoveryStatusSdnCollectionInfo) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DiscoveryStatusSdnCollectionInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *DiscoveryStatusSdnCollectionInfo) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DiscoveryStatusSdnCollectionInfo) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DiscoveryStatusSdnCollectionInfo) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DiscoveryStatusSdnCollectionInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


