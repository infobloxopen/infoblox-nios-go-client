# NetworkcontainerSubscribeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledAttributes** | Pointer to **[]string** | The list of Cisco ISE attributes allowed for subscription. | [optional] 
**MappedEaAttributes** | Pointer to [**[]NetworkcontainersubscribesettingsMappedEaAttributes**](NetworkcontainersubscribesettingsMappedEaAttributes.md) | The list of NIOS extensible attributes to Cisco ISE attributes mappings. | [optional] 

## Methods

### NewNetworkcontainerSubscribeSettings

`func NewNetworkcontainerSubscribeSettings() *NetworkcontainerSubscribeSettings`

NewNetworkcontainerSubscribeSettings instantiates a new NetworkcontainerSubscribeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkcontainerSubscribeSettingsWithDefaults

`func NewNetworkcontainerSubscribeSettingsWithDefaults() *NetworkcontainerSubscribeSettings`

NewNetworkcontainerSubscribeSettingsWithDefaults instantiates a new NetworkcontainerSubscribeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledAttributes

`func (o *NetworkcontainerSubscribeSettings) GetEnabledAttributes() []string`

GetEnabledAttributes returns the EnabledAttributes field if non-nil, zero value otherwise.

### GetEnabledAttributesOk

`func (o *NetworkcontainerSubscribeSettings) GetEnabledAttributesOk() (*[]string, bool)`

GetEnabledAttributesOk returns a tuple with the EnabledAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAttributes

`func (o *NetworkcontainerSubscribeSettings) SetEnabledAttributes(v []string)`

SetEnabledAttributes sets EnabledAttributes field to given value.

### HasEnabledAttributes

`func (o *NetworkcontainerSubscribeSettings) HasEnabledAttributes() bool`

HasEnabledAttributes returns a boolean if a field has been set.

### GetMappedEaAttributes

`func (o *NetworkcontainerSubscribeSettings) GetMappedEaAttributes() []NetworkcontainersubscribesettingsMappedEaAttributes`

GetMappedEaAttributes returns the MappedEaAttributes field if non-nil, zero value otherwise.

### GetMappedEaAttributesOk

`func (o *NetworkcontainerSubscribeSettings) GetMappedEaAttributesOk() (*[]NetworkcontainersubscribesettingsMappedEaAttributes, bool)`

GetMappedEaAttributesOk returns a tuple with the MappedEaAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedEaAttributes

`func (o *NetworkcontainerSubscribeSettings) SetMappedEaAttributes(v []NetworkcontainersubscribesettingsMappedEaAttributes)`

SetMappedEaAttributes sets MappedEaAttributes field to given value.

### HasMappedEaAttributes

`func (o *NetworkcontainerSubscribeSettings) HasMappedEaAttributes() bool`

HasMappedEaAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


