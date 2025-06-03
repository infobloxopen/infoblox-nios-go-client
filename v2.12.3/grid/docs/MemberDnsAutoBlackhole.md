# MemberDnsAutoBlackhole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableFetchesPerServer** | Pointer to **bool** | Enables or disables the configuration of the maximum number of concurrent recursive queries the appliance sends to each upstream DNS server. | [optional] 
**EnableFetchesPerZone** | Pointer to **bool** | Enables or disables the configuration of the maximum number of concurrent recursive queries the appliance sends to each DNS zone. | [optional] 
**EnableHolddown** | Pointer to **bool** | Enables or disables the holddown configuration when the appliance stops sending queries to non-responsive servers. | [optional] 
**FetchesPerServer** | Pointer to **int64** | The maximum number of concurrent recursive queries the appliance sends to a single upstream name server before blocking additional queries to that server. | [optional] 
**FetchesPerZone** | Pointer to **int64** | The maximum number of concurrent recursive queries that a server sends for its domains. | [optional] 
**FpsFreq** | Pointer to **int64** | Determines how often (in number of recursive responses) the appliance recalculates the average timeout ratio for each DNS server. | [optional] 
**Holddown** | Pointer to **int64** | The holddown duration for non-responsive servers. | [optional] 
**HolddownThreshold** | Pointer to **int64** | The number of consecutive timeouts before holding down a non-responsive server. | [optional] 
**HolddownTimeout** | Pointer to **int64** | The minimum time (in seconds) that needs to be passed before a timeout occurs. Note that only these timeouts are counted towards the number of consecutive timeouts. | [optional] 

## Methods

### NewMemberDnsAutoBlackhole

`func NewMemberDnsAutoBlackhole() *MemberDnsAutoBlackhole`

NewMemberDnsAutoBlackhole instantiates a new MemberDnsAutoBlackhole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberDnsAutoBlackholeWithDefaults

`func NewMemberDnsAutoBlackholeWithDefaults() *MemberDnsAutoBlackhole`

NewMemberDnsAutoBlackholeWithDefaults instantiates a new MemberDnsAutoBlackhole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableFetchesPerServer

`func (o *MemberDnsAutoBlackhole) GetEnableFetchesPerServer() bool`

GetEnableFetchesPerServer returns the EnableFetchesPerServer field if non-nil, zero value otherwise.

### GetEnableFetchesPerServerOk

`func (o *MemberDnsAutoBlackhole) GetEnableFetchesPerServerOk() (*bool, bool)`

GetEnableFetchesPerServerOk returns a tuple with the EnableFetchesPerServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFetchesPerServer

`func (o *MemberDnsAutoBlackhole) SetEnableFetchesPerServer(v bool)`

SetEnableFetchesPerServer sets EnableFetchesPerServer field to given value.

### HasEnableFetchesPerServer

`func (o *MemberDnsAutoBlackhole) HasEnableFetchesPerServer() bool`

HasEnableFetchesPerServer returns a boolean if a field has been set.

### GetEnableFetchesPerZone

`func (o *MemberDnsAutoBlackhole) GetEnableFetchesPerZone() bool`

GetEnableFetchesPerZone returns the EnableFetchesPerZone field if non-nil, zero value otherwise.

### GetEnableFetchesPerZoneOk

`func (o *MemberDnsAutoBlackhole) GetEnableFetchesPerZoneOk() (*bool, bool)`

GetEnableFetchesPerZoneOk returns a tuple with the EnableFetchesPerZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableFetchesPerZone

`func (o *MemberDnsAutoBlackhole) SetEnableFetchesPerZone(v bool)`

SetEnableFetchesPerZone sets EnableFetchesPerZone field to given value.

### HasEnableFetchesPerZone

`func (o *MemberDnsAutoBlackhole) HasEnableFetchesPerZone() bool`

HasEnableFetchesPerZone returns a boolean if a field has been set.

### GetEnableHolddown

`func (o *MemberDnsAutoBlackhole) GetEnableHolddown() bool`

GetEnableHolddown returns the EnableHolddown field if non-nil, zero value otherwise.

### GetEnableHolddownOk

`func (o *MemberDnsAutoBlackhole) GetEnableHolddownOk() (*bool, bool)`

GetEnableHolddownOk returns a tuple with the EnableHolddown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableHolddown

`func (o *MemberDnsAutoBlackhole) SetEnableHolddown(v bool)`

SetEnableHolddown sets EnableHolddown field to given value.

### HasEnableHolddown

`func (o *MemberDnsAutoBlackhole) HasEnableHolddown() bool`

HasEnableHolddown returns a boolean if a field has been set.

### GetFetchesPerServer

`func (o *MemberDnsAutoBlackhole) GetFetchesPerServer() int64`

GetFetchesPerServer returns the FetchesPerServer field if non-nil, zero value otherwise.

### GetFetchesPerServerOk

`func (o *MemberDnsAutoBlackhole) GetFetchesPerServerOk() (*int64, bool)`

GetFetchesPerServerOk returns a tuple with the FetchesPerServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchesPerServer

`func (o *MemberDnsAutoBlackhole) SetFetchesPerServer(v int64)`

SetFetchesPerServer sets FetchesPerServer field to given value.

### HasFetchesPerServer

`func (o *MemberDnsAutoBlackhole) HasFetchesPerServer() bool`

HasFetchesPerServer returns a boolean if a field has been set.

### GetFetchesPerZone

`func (o *MemberDnsAutoBlackhole) GetFetchesPerZone() int64`

GetFetchesPerZone returns the FetchesPerZone field if non-nil, zero value otherwise.

### GetFetchesPerZoneOk

`func (o *MemberDnsAutoBlackhole) GetFetchesPerZoneOk() (*int64, bool)`

GetFetchesPerZoneOk returns a tuple with the FetchesPerZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchesPerZone

`func (o *MemberDnsAutoBlackhole) SetFetchesPerZone(v int64)`

SetFetchesPerZone sets FetchesPerZone field to given value.

### HasFetchesPerZone

`func (o *MemberDnsAutoBlackhole) HasFetchesPerZone() bool`

HasFetchesPerZone returns a boolean if a field has been set.

### GetFpsFreq

`func (o *MemberDnsAutoBlackhole) GetFpsFreq() int64`

GetFpsFreq returns the FpsFreq field if non-nil, zero value otherwise.

### GetFpsFreqOk

`func (o *MemberDnsAutoBlackhole) GetFpsFreqOk() (*int64, bool)`

GetFpsFreqOk returns a tuple with the FpsFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFpsFreq

`func (o *MemberDnsAutoBlackhole) SetFpsFreq(v int64)`

SetFpsFreq sets FpsFreq field to given value.

### HasFpsFreq

`func (o *MemberDnsAutoBlackhole) HasFpsFreq() bool`

HasFpsFreq returns a boolean if a field has been set.

### GetHolddown

`func (o *MemberDnsAutoBlackhole) GetHolddown() int64`

GetHolddown returns the Holddown field if non-nil, zero value otherwise.

### GetHolddownOk

`func (o *MemberDnsAutoBlackhole) GetHolddownOk() (*int64, bool)`

GetHolddownOk returns a tuple with the Holddown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolddown

`func (o *MemberDnsAutoBlackhole) SetHolddown(v int64)`

SetHolddown sets Holddown field to given value.

### HasHolddown

`func (o *MemberDnsAutoBlackhole) HasHolddown() bool`

HasHolddown returns a boolean if a field has been set.

### GetHolddownThreshold

`func (o *MemberDnsAutoBlackhole) GetHolddownThreshold() int64`

GetHolddownThreshold returns the HolddownThreshold field if non-nil, zero value otherwise.

### GetHolddownThresholdOk

`func (o *MemberDnsAutoBlackhole) GetHolddownThresholdOk() (*int64, bool)`

GetHolddownThresholdOk returns a tuple with the HolddownThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolddownThreshold

`func (o *MemberDnsAutoBlackhole) SetHolddownThreshold(v int64)`

SetHolddownThreshold sets HolddownThreshold field to given value.

### HasHolddownThreshold

`func (o *MemberDnsAutoBlackhole) HasHolddownThreshold() bool`

HasHolddownThreshold returns a boolean if a field has been set.

### GetHolddownTimeout

`func (o *MemberDnsAutoBlackhole) GetHolddownTimeout() int64`

GetHolddownTimeout returns the HolddownTimeout field if non-nil, zero value otherwise.

### GetHolddownTimeoutOk

`func (o *MemberDnsAutoBlackhole) GetHolddownTimeoutOk() (*int64, bool)`

GetHolddownTimeoutOk returns a tuple with the HolddownTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolddownTimeout

`func (o *MemberDnsAutoBlackhole) SetHolddownTimeout(v int64)`

SetHolddownTimeout sets HolddownTimeout field to given value.

### HasHolddownTimeout

`func (o *MemberDnsAutoBlackhole) HasHolddownTimeout() bool`

HasHolddownTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


