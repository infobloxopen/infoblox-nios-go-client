# GetThreatinsightCloudclientResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**BlacklistRpzList** | Pointer to **[]string** | The RPZs to which you apply newly detected domains through the Infoblox Threat Insight Cloud Client. | [optional] 
**Enable** | Pointer to **bool** | Determines whether the Threat Insight in Cloud Client is enabled. | [optional] 
**ForceRefresh** | Pointer to **bool** | Force a refresh if at least one RPZ is configured. | [optional] 
**Interval** | Pointer to **int64** | The time interval (in seconds) for requesting newly detected domains by the Infoblox Threat Insight Cloud Client and applying them to the list of configured RPZs. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**ThreatinsightCloudclient**](ThreatinsightCloudclient.md) |  | [optional] 

## Methods

### NewGetThreatinsightCloudclientResponse

`func NewGetThreatinsightCloudclientResponse() *GetThreatinsightCloudclientResponse`

NewGetThreatinsightCloudclientResponse instantiates a new GetThreatinsightCloudclientResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetThreatinsightCloudclientResponseWithDefaults

`func NewGetThreatinsightCloudclientResponseWithDefaults() *GetThreatinsightCloudclientResponse`

NewGetThreatinsightCloudclientResponseWithDefaults instantiates a new GetThreatinsightCloudclientResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetThreatinsightCloudclientResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetThreatinsightCloudclientResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetThreatinsightCloudclientResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetThreatinsightCloudclientResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetBlacklistRpzList

`func (o *GetThreatinsightCloudclientResponse) GetBlacklistRpzList() []string`

GetBlacklistRpzList returns the BlacklistRpzList field if non-nil, zero value otherwise.

### GetBlacklistRpzListOk

`func (o *GetThreatinsightCloudclientResponse) GetBlacklistRpzListOk() (*[]string, bool)`

GetBlacklistRpzListOk returns a tuple with the BlacklistRpzList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistRpzList

`func (o *GetThreatinsightCloudclientResponse) SetBlacklistRpzList(v []string)`

SetBlacklistRpzList sets BlacklistRpzList field to given value.

### HasBlacklistRpzList

`func (o *GetThreatinsightCloudclientResponse) HasBlacklistRpzList() bool`

HasBlacklistRpzList returns a boolean if a field has been set.

### GetEnable

`func (o *GetThreatinsightCloudclientResponse) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *GetThreatinsightCloudclientResponse) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *GetThreatinsightCloudclientResponse) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *GetThreatinsightCloudclientResponse) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetForceRefresh

`func (o *GetThreatinsightCloudclientResponse) GetForceRefresh() bool`

GetForceRefresh returns the ForceRefresh field if non-nil, zero value otherwise.

### GetForceRefreshOk

`func (o *GetThreatinsightCloudclientResponse) GetForceRefreshOk() (*bool, bool)`

GetForceRefreshOk returns a tuple with the ForceRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceRefresh

`func (o *GetThreatinsightCloudclientResponse) SetForceRefresh(v bool)`

SetForceRefresh sets ForceRefresh field to given value.

### HasForceRefresh

`func (o *GetThreatinsightCloudclientResponse) HasForceRefresh() bool`

HasForceRefresh returns a boolean if a field has been set.

### GetInterval

`func (o *GetThreatinsightCloudclientResponse) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *GetThreatinsightCloudclientResponse) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *GetThreatinsightCloudclientResponse) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *GetThreatinsightCloudclientResponse) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetUuid

`func (o *GetThreatinsightCloudclientResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetThreatinsightCloudclientResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetThreatinsightCloudclientResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetThreatinsightCloudclientResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetThreatinsightCloudclientResponse) GetResult() ThreatinsightCloudclient`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetThreatinsightCloudclientResponse) GetResultOk() (*ThreatinsightCloudclient, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetThreatinsightCloudclientResponse) SetResult(v ThreatinsightCloudclient)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetThreatinsightCloudclientResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


