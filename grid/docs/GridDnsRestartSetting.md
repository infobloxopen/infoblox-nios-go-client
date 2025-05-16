# GridDnsRestartSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | Pointer to **int64** | The time duration to delay a restart for a restart group. | [optional] 
**Timeout** | Pointer to **string** | The duration of timeout for a restart group. The value \&quot;-1\&quot; means infinite. | [optional] 
**RestartOffline** | Pointer to **bool** | Determines whether the Grid should try to restart offline member. | [optional] 

## Methods

### NewGridDnsRestartSetting

`func NewGridDnsRestartSetting() *GridDnsRestartSetting`

NewGridDnsRestartSetting instantiates a new GridDnsRestartSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsRestartSettingWithDefaults

`func NewGridDnsRestartSettingWithDefaults() *GridDnsRestartSetting`

NewGridDnsRestartSettingWithDefaults instantiates a new GridDnsRestartSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *GridDnsRestartSetting) GetDelay() int64`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *GridDnsRestartSetting) GetDelayOk() (*int64, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *GridDnsRestartSetting) SetDelay(v int64)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *GridDnsRestartSetting) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetTimeout

`func (o *GridDnsRestartSetting) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GridDnsRestartSetting) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GridDnsRestartSetting) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GridDnsRestartSetting) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRestartOffline

`func (o *GridDnsRestartSetting) GetRestartOffline() bool`

GetRestartOffline returns the RestartOffline field if non-nil, zero value otherwise.

### GetRestartOfflineOk

`func (o *GridDnsRestartSetting) GetRestartOfflineOk() (*bool, bool)`

GetRestartOfflineOk returns a tuple with the RestartOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOffline

`func (o *GridDnsRestartSetting) SetRestartOffline(v bool)`

SetRestartOffline sets RestartOffline field to given value.

### HasRestartOffline

`func (o *GridDnsRestartSetting) HasRestartOffline() bool`

HasRestartOffline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


