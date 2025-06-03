# DiscoveryDevicePortStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfacesCount** | Pointer to **int64** | The total number of available interfaces on this device. | [optional] [readonly] 
**AdminUpOperUpCount** | Pointer to **int64** | The total number of interfaces which have both administrative and operating states as &#39;UP&#39;. | [optional] [readonly] 
**AdminUpOperDownCount** | Pointer to **int64** | The total number of interfaces which have administrative state &#39;UP&#39; and oper state &#39;DOWN&#39;. | [optional] [readonly] 
**AdminDownOperDownCount** | Pointer to **int64** | The total number of interfaces which have administrative state &#39;DOWN&#39; and operating state &#39;DOWN&#39;. | [optional] [readonly] 

## Methods

### NewDiscoveryDevicePortStats

`func NewDiscoveryDevicePortStats() *DiscoveryDevicePortStats`

NewDiscoveryDevicePortStats instantiates a new DiscoveryDevicePortStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryDevicePortStatsWithDefaults

`func NewDiscoveryDevicePortStatsWithDefaults() *DiscoveryDevicePortStats`

NewDiscoveryDevicePortStatsWithDefaults instantiates a new DiscoveryDevicePortStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfacesCount

`func (o *DiscoveryDevicePortStats) GetInterfacesCount() int64`

GetInterfacesCount returns the InterfacesCount field if non-nil, zero value otherwise.

### GetInterfacesCountOk

`func (o *DiscoveryDevicePortStats) GetInterfacesCountOk() (*int64, bool)`

GetInterfacesCountOk returns a tuple with the InterfacesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfacesCount

`func (o *DiscoveryDevicePortStats) SetInterfacesCount(v int64)`

SetInterfacesCount sets InterfacesCount field to given value.

### HasInterfacesCount

`func (o *DiscoveryDevicePortStats) HasInterfacesCount() bool`

HasInterfacesCount returns a boolean if a field has been set.

### GetAdminUpOperUpCount

`func (o *DiscoveryDevicePortStats) GetAdminUpOperUpCount() int64`

GetAdminUpOperUpCount returns the AdminUpOperUpCount field if non-nil, zero value otherwise.

### GetAdminUpOperUpCountOk

`func (o *DiscoveryDevicePortStats) GetAdminUpOperUpCountOk() (*int64, bool)`

GetAdminUpOperUpCountOk returns a tuple with the AdminUpOperUpCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUpOperUpCount

`func (o *DiscoveryDevicePortStats) SetAdminUpOperUpCount(v int64)`

SetAdminUpOperUpCount sets AdminUpOperUpCount field to given value.

### HasAdminUpOperUpCount

`func (o *DiscoveryDevicePortStats) HasAdminUpOperUpCount() bool`

HasAdminUpOperUpCount returns a boolean if a field has been set.

### GetAdminUpOperDownCount

`func (o *DiscoveryDevicePortStats) GetAdminUpOperDownCount() int64`

GetAdminUpOperDownCount returns the AdminUpOperDownCount field if non-nil, zero value otherwise.

### GetAdminUpOperDownCountOk

`func (o *DiscoveryDevicePortStats) GetAdminUpOperDownCountOk() (*int64, bool)`

GetAdminUpOperDownCountOk returns a tuple with the AdminUpOperDownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUpOperDownCount

`func (o *DiscoveryDevicePortStats) SetAdminUpOperDownCount(v int64)`

SetAdminUpOperDownCount sets AdminUpOperDownCount field to given value.

### HasAdminUpOperDownCount

`func (o *DiscoveryDevicePortStats) HasAdminUpOperDownCount() bool`

HasAdminUpOperDownCount returns a boolean if a field has been set.

### GetAdminDownOperDownCount

`func (o *DiscoveryDevicePortStats) GetAdminDownOperDownCount() int64`

GetAdminDownOperDownCount returns the AdminDownOperDownCount field if non-nil, zero value otherwise.

### GetAdminDownOperDownCountOk

`func (o *DiscoveryDevicePortStats) GetAdminDownOperDownCountOk() (*int64, bool)`

GetAdminDownOperDownCountOk returns a tuple with the AdminDownOperDownCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminDownOperDownCount

`func (o *DiscoveryDevicePortStats) SetAdminDownOperDownCount(v int64)`

SetAdminDownOperDownCount sets AdminDownOperDownCount field to given value.

### HasAdminDownOperDownCount

`func (o *DiscoveryDevicePortStats) HasAdminDownOperDownCount() bool`

HasAdminDownOperDownCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


