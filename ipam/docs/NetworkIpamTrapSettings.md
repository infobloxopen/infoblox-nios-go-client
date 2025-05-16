# NetworkIpamTrapSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableEmailWarnings** | Pointer to **bool** | Determines whether sending warnings by email is enabled or not. | [optional] 
**EnableSnmpWarnings** | Pointer to **bool** | Determines whether sending warnings by SNMP is enabled or not. | [optional] 

## Methods

### NewNetworkIpamTrapSettings

`func NewNetworkIpamTrapSettings() *NetworkIpamTrapSettings`

NewNetworkIpamTrapSettings instantiates a new NetworkIpamTrapSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkIpamTrapSettingsWithDefaults

`func NewNetworkIpamTrapSettingsWithDefaults() *NetworkIpamTrapSettings`

NewNetworkIpamTrapSettingsWithDefaults instantiates a new NetworkIpamTrapSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableEmailWarnings

`func (o *NetworkIpamTrapSettings) GetEnableEmailWarnings() bool`

GetEnableEmailWarnings returns the EnableEmailWarnings field if non-nil, zero value otherwise.

### GetEnableEmailWarningsOk

`func (o *NetworkIpamTrapSettings) GetEnableEmailWarningsOk() (*bool, bool)`

GetEnableEmailWarningsOk returns a tuple with the EnableEmailWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableEmailWarnings

`func (o *NetworkIpamTrapSettings) SetEnableEmailWarnings(v bool)`

SetEnableEmailWarnings sets EnableEmailWarnings field to given value.

### HasEnableEmailWarnings

`func (o *NetworkIpamTrapSettings) HasEnableEmailWarnings() bool`

HasEnableEmailWarnings returns a boolean if a field has been set.

### GetEnableSnmpWarnings

`func (o *NetworkIpamTrapSettings) GetEnableSnmpWarnings() bool`

GetEnableSnmpWarnings returns the EnableSnmpWarnings field if non-nil, zero value otherwise.

### GetEnableSnmpWarningsOk

`func (o *NetworkIpamTrapSettings) GetEnableSnmpWarningsOk() (*bool, bool)`

GetEnableSnmpWarningsOk returns a tuple with the EnableSnmpWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSnmpWarnings

`func (o *NetworkIpamTrapSettings) SetEnableSnmpWarnings(v bool)`

SetEnableSnmpWarnings sets EnableSnmpWarnings field to given value.

### HasEnableSnmpWarnings

`func (o *NetworkIpamTrapSettings) HasEnableSnmpWarnings() bool`

HasEnableSnmpWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


