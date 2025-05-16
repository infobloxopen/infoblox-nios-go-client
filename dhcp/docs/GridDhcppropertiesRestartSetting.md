# GridDhcppropertiesRestartSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delay** | Pointer to **int64** | The time duration to delay a restart for a restart group. | [optional] 
**Timeout** | Pointer to **string** | The duration of timeout for a restart group. The value \&quot;-1\&quot; means infinite. | [optional] 
**RestartOffline** | Pointer to **bool** | Determines whether the Grid should try to restart offline member. | [optional] 

## Methods

### NewGridDhcppropertiesRestartSetting

`func NewGridDhcppropertiesRestartSetting() *GridDhcppropertiesRestartSetting`

NewGridDhcppropertiesRestartSetting instantiates a new GridDhcppropertiesRestartSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDhcppropertiesRestartSettingWithDefaults

`func NewGridDhcppropertiesRestartSettingWithDefaults() *GridDhcppropertiesRestartSetting`

NewGridDhcppropertiesRestartSettingWithDefaults instantiates a new GridDhcppropertiesRestartSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelay

`func (o *GridDhcppropertiesRestartSetting) GetDelay() int64`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *GridDhcppropertiesRestartSetting) GetDelayOk() (*int64, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *GridDhcppropertiesRestartSetting) SetDelay(v int64)`

SetDelay sets Delay field to given value.

### HasDelay

`func (o *GridDhcppropertiesRestartSetting) HasDelay() bool`

HasDelay returns a boolean if a field has been set.

### GetTimeout

`func (o *GridDhcppropertiesRestartSetting) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GridDhcppropertiesRestartSetting) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GridDhcppropertiesRestartSetting) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GridDhcppropertiesRestartSetting) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetRestartOffline

`func (o *GridDhcppropertiesRestartSetting) GetRestartOffline() bool`

GetRestartOffline returns the RestartOffline field if non-nil, zero value otherwise.

### GetRestartOfflineOk

`func (o *GridDhcppropertiesRestartSetting) GetRestartOfflineOk() (*bool, bool)`

GetRestartOfflineOk returns a tuple with the RestartOffline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartOffline

`func (o *GridDhcppropertiesRestartSetting) SetRestartOffline(v bool)`

SetRestartOffline sets RestartOffline field to given value.

### HasRestartOffline

`func (o *GridDhcppropertiesRestartSetting) HasRestartOffline() bool`

HasRestartOffline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


