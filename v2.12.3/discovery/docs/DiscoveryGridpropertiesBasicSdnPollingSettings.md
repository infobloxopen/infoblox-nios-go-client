# DiscoveryGridpropertiesBasicSdnPollingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SdnDiscovery** | Pointer to **bool** | Enable/disable SDN discovery. | [optional] 
**DefaultNetworkView** | Pointer to **string** | Default network view to map collected SDN networks. | [optional] 
**EndHostPolling** | Pointer to **string** | SDN end host polling mode. | [optional] 
**EndHostPollingInterval** | Pointer to **int64** | Valid SDN end host polling interval in seconds. Must be between 1800 and 86400 seconds. | [optional] 
**EndHostPollingSchedule** | Pointer to [**DiscoverygridpropertiesbasicsdnpollingsettingsEndHostPollingSchedule**](DiscoverygridpropertiesbasicsdnpollingsettingsEndHostPollingSchedule.md) |  | [optional] 

## Methods

### NewDiscoveryGridpropertiesBasicSdnPollingSettings

`func NewDiscoveryGridpropertiesBasicSdnPollingSettings() *DiscoveryGridpropertiesBasicSdnPollingSettings`

NewDiscoveryGridpropertiesBasicSdnPollingSettings instantiates a new DiscoveryGridpropertiesBasicSdnPollingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryGridpropertiesBasicSdnPollingSettingsWithDefaults

`func NewDiscoveryGridpropertiesBasicSdnPollingSettingsWithDefaults() *DiscoveryGridpropertiesBasicSdnPollingSettings`

NewDiscoveryGridpropertiesBasicSdnPollingSettingsWithDefaults instantiates a new DiscoveryGridpropertiesBasicSdnPollingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSdnDiscovery

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetSdnDiscovery() bool`

GetSdnDiscovery returns the SdnDiscovery field if non-nil, zero value otherwise.

### GetSdnDiscoveryOk

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetSdnDiscoveryOk() (*bool, bool)`

GetSdnDiscoveryOk returns a tuple with the SdnDiscovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdnDiscovery

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) SetSdnDiscovery(v bool)`

SetSdnDiscovery sets SdnDiscovery field to given value.

### HasSdnDiscovery

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) HasSdnDiscovery() bool`

HasSdnDiscovery returns a boolean if a field has been set.

### GetDefaultNetworkView

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetDefaultNetworkView() string`

GetDefaultNetworkView returns the DefaultNetworkView field if non-nil, zero value otherwise.

### GetDefaultNetworkViewOk

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetDefaultNetworkViewOk() (*string, bool)`

GetDefaultNetworkViewOk returns a tuple with the DefaultNetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkView

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) SetDefaultNetworkView(v string)`

SetDefaultNetworkView sets DefaultNetworkView field to given value.

### HasDefaultNetworkView

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) HasDefaultNetworkView() bool`

HasDefaultNetworkView returns a boolean if a field has been set.

### GetEndHostPolling

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetEndHostPolling() string`

GetEndHostPolling returns the EndHostPolling field if non-nil, zero value otherwise.

### GetEndHostPollingOk

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetEndHostPollingOk() (*string, bool)`

GetEndHostPollingOk returns a tuple with the EndHostPolling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHostPolling

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) SetEndHostPolling(v string)`

SetEndHostPolling sets EndHostPolling field to given value.

### HasEndHostPolling

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) HasEndHostPolling() bool`

HasEndHostPolling returns a boolean if a field has been set.

### GetEndHostPollingInterval

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetEndHostPollingInterval() int64`

GetEndHostPollingInterval returns the EndHostPollingInterval field if non-nil, zero value otherwise.

### GetEndHostPollingIntervalOk

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetEndHostPollingIntervalOk() (*int64, bool)`

GetEndHostPollingIntervalOk returns a tuple with the EndHostPollingInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHostPollingInterval

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) SetEndHostPollingInterval(v int64)`

SetEndHostPollingInterval sets EndHostPollingInterval field to given value.

### HasEndHostPollingInterval

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) HasEndHostPollingInterval() bool`

HasEndHostPollingInterval returns a boolean if a field has been set.

### GetEndHostPollingSchedule

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetEndHostPollingSchedule() DiscoverygridpropertiesbasicsdnpollingsettingsEndHostPollingSchedule`

GetEndHostPollingSchedule returns the EndHostPollingSchedule field if non-nil, zero value otherwise.

### GetEndHostPollingScheduleOk

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) GetEndHostPollingScheduleOk() (*DiscoverygridpropertiesbasicsdnpollingsettingsEndHostPollingSchedule, bool)`

GetEndHostPollingScheduleOk returns a tuple with the EndHostPollingSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndHostPollingSchedule

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) SetEndHostPollingSchedule(v DiscoverygridpropertiesbasicsdnpollingsettingsEndHostPollingSchedule)`

SetEndHostPollingSchedule sets EndHostPollingSchedule field to given value.

### HasEndHostPollingSchedule

`func (o *DiscoveryGridpropertiesBasicSdnPollingSettings) HasEndHostPollingSchedule() bool`

HasEndHostPollingSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


