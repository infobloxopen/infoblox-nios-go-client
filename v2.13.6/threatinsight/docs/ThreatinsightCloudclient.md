# ThreatinsightCloudclient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**BlacklistRpzList** | Pointer to **[]string** | The RPZs to which you apply newly detected domains through the Infoblox Threat Insight Cloud Client. | [optional] 
**Enable** | Pointer to **bool** | Determines whether the Threat Insight in Cloud Client is enabled. | [optional] 
**ForceRefresh** | Pointer to **bool** | Force a refresh if at least one RPZ is configured. | [optional] 
**Interval** | Pointer to **int64** | The time interval (in seconds) for requesting newly detected domains by the Infoblox Threat Insight Cloud Client and applying them to the list of configured RPZs. | [optional] 

## Methods

### NewThreatinsightCloudclient

`func NewThreatinsightCloudclient() *ThreatinsightCloudclient`

NewThreatinsightCloudclient instantiates a new ThreatinsightCloudclient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatinsightCloudclientWithDefaults

`func NewThreatinsightCloudclientWithDefaults() *ThreatinsightCloudclient`

NewThreatinsightCloudclientWithDefaults instantiates a new ThreatinsightCloudclient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *ThreatinsightCloudclient) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *ThreatinsightCloudclient) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *ThreatinsightCloudclient) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *ThreatinsightCloudclient) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBlacklistRpzList

`func (o *ThreatinsightCloudclient) GetBlacklistRpzList() []string`

GetBlacklistRpzList returns the BlacklistRpzList field if non-nil, zero value otherwise.

### GetBlacklistRpzListOk

`func (o *ThreatinsightCloudclient) GetBlacklistRpzListOk() (*[]string, bool)`

GetBlacklistRpzListOk returns a tuple with the BlacklistRpzList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRpzList

`func (o *ThreatinsightCloudclient) SetBlacklistRpzList(v []string)`

SetBlacklistRpzList sets BlacklistRpzList field to given value.

### HasBlacklistRpzList

`func (o *ThreatinsightCloudclient) HasBlacklistRpzList() bool`

HasBlacklistRpzList returns a boolean if a field has been set.

### GetEnable

`func (o *ThreatinsightCloudclient) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ThreatinsightCloudclient) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ThreatinsightCloudclient) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ThreatinsightCloudclient) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetForceRefresh

`func (o *ThreatinsightCloudclient) GetForceRefresh() bool`

GetForceRefresh returns the ForceRefresh field if non-nil, zero value otherwise.

### GetForceRefreshOk

`func (o *ThreatinsightCloudclient) GetForceRefreshOk() (*bool, bool)`

GetForceRefreshOk returns a tuple with the ForceRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRefresh

`func (o *ThreatinsightCloudclient) SetForceRefresh(v bool)`

SetForceRefresh sets ForceRefresh field to given value.

### HasForceRefresh

`func (o *ThreatinsightCloudclient) HasForceRefresh() bool`

HasForceRefresh returns a boolean if a field has been set.

### GetInterval

`func (o *ThreatinsightCloudclient) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ThreatinsightCloudclient) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ThreatinsightCloudclient) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *ThreatinsightCloudclient) HasInterval() bool`

HasInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


