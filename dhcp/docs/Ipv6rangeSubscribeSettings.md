# Ipv6rangeSubscribeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledAttributes** | Pointer to **[]string** | The list of Cisco ISE attributes allowed for subscription. | [optional] 
**MappedEaAttributes** | Pointer to [**Ipv6rangesubscribesettingsMappedEaAttributes**](Ipv6rangesubscribesettingsMappedEaAttributes.md) |  | [optional] 

## Methods

### NewIpv6rangeSubscribeSettings

`func NewIpv6rangeSubscribeSettings() *Ipv6rangeSubscribeSettings`

NewIpv6rangeSubscribeSettings instantiates a new Ipv6rangeSubscribeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpv6rangeSubscribeSettingsWithDefaults

`func NewIpv6rangeSubscribeSettingsWithDefaults() *Ipv6rangeSubscribeSettings`

NewIpv6rangeSubscribeSettingsWithDefaults instantiates a new Ipv6rangeSubscribeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledAttributes

`func (o *Ipv6rangeSubscribeSettings) GetEnabledAttributes() []string`

GetEnabledAttributes returns the EnabledAttributes field if non-nil, zero value otherwise.

### GetEnabledAttributesOk

`func (o *Ipv6rangeSubscribeSettings) GetEnabledAttributesOk() (*[]string, bool)`

GetEnabledAttributesOk returns a tuple with the EnabledAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAttributes

`func (o *Ipv6rangeSubscribeSettings) SetEnabledAttributes(v []string)`

SetEnabledAttributes sets EnabledAttributes field to given value.

### HasEnabledAttributes

`func (o *Ipv6rangeSubscribeSettings) HasEnabledAttributes() bool`

HasEnabledAttributes returns a boolean if a field has been set.

### GetMappedEaAttributes

`func (o *Ipv6rangeSubscribeSettings) GetMappedEaAttributes() Ipv6rangesubscribesettingsMappedEaAttributes`

GetMappedEaAttributes returns the MappedEaAttributes field if non-nil, zero value otherwise.

### GetMappedEaAttributesOk

`func (o *Ipv6rangeSubscribeSettings) GetMappedEaAttributesOk() (*Ipv6rangesubscribesettingsMappedEaAttributes, bool)`

GetMappedEaAttributesOk returns a tuple with the MappedEaAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedEaAttributes

`func (o *Ipv6rangeSubscribeSettings) SetMappedEaAttributes(v Ipv6rangesubscribesettingsMappedEaAttributes)`

SetMappedEaAttributes sets MappedEaAttributes field to given value.

### HasMappedEaAttributes

`func (o *Ipv6rangeSubscribeSettings) HasMappedEaAttributes() bool`

HasMappedEaAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


