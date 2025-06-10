# PxgridEndpointSubscribeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnabledAttributes** | Pointer to **[]string** | The list of Cisco ISE attributes allowed for subscription. | [optional] 
**MappedEaAttributes** | Pointer to [**PxgridendpointsubscribesettingsMappedEaAttributes**](PxgridendpointsubscribesettingsMappedEaAttributes.md) |  | [optional] 

## Methods

### NewPxgridEndpointSubscribeSettings

`func NewPxgridEndpointSubscribeSettings() *PxgridEndpointSubscribeSettings`

NewPxgridEndpointSubscribeSettings instantiates a new PxgridEndpointSubscribeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPxgridEndpointSubscribeSettingsWithDefaults

`func NewPxgridEndpointSubscribeSettingsWithDefaults() *PxgridEndpointSubscribeSettings`

NewPxgridEndpointSubscribeSettingsWithDefaults instantiates a new PxgridEndpointSubscribeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabledAttributes

`func (o *PxgridEndpointSubscribeSettings) GetEnabledAttributes() []string`

GetEnabledAttributes returns the EnabledAttributes field if non-nil, zero value otherwise.

### GetEnabledAttributesOk

`func (o *PxgridEndpointSubscribeSettings) GetEnabledAttributesOk() (*[]string, bool)`

GetEnabledAttributesOk returns a tuple with the EnabledAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabledAttributes

`func (o *PxgridEndpointSubscribeSettings) SetEnabledAttributes(v []string)`

SetEnabledAttributes sets EnabledAttributes field to given value.

### HasEnabledAttributes

`func (o *PxgridEndpointSubscribeSettings) HasEnabledAttributes() bool`

HasEnabledAttributes returns a boolean if a field has been set.

### GetMappedEaAttributes

`func (o *PxgridEndpointSubscribeSettings) GetMappedEaAttributes() PxgridendpointsubscribesettingsMappedEaAttributes`

GetMappedEaAttributes returns the MappedEaAttributes field if non-nil, zero value otherwise.

### GetMappedEaAttributesOk

`func (o *PxgridEndpointSubscribeSettings) GetMappedEaAttributesOk() (*PxgridendpointsubscribesettingsMappedEaAttributes, bool)`

GetMappedEaAttributesOk returns a tuple with the MappedEaAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedEaAttributes

`func (o *PxgridEndpointSubscribeSettings) SetMappedEaAttributes(v PxgridendpointsubscribesettingsMappedEaAttributes)`

SetMappedEaAttributes sets MappedEaAttributes field to given value.

### HasMappedEaAttributes

`func (o *PxgridEndpointSubscribeSettings) HasMappedEaAttributes() bool`

HasMappedEaAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


