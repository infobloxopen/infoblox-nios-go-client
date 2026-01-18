# Orderedresponsepolicyzones

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**RpZones** | Pointer to **[]string** | An ordered list of Response Policy Zone names. | [optional] 
**View** | Pointer to **string** | The DNS View name. | [optional] 

## Methods

### NewOrderedresponsepolicyzones

`func NewOrderedresponsepolicyzones() *Orderedresponsepolicyzones`

NewOrderedresponsepolicyzones instantiates a new Orderedresponsepolicyzones object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderedresponsepolicyzonesWithDefaults

`func NewOrderedresponsepolicyzonesWithDefaults() *Orderedresponsepolicyzones`

NewOrderedresponsepolicyzonesWithDefaults instantiates a new Orderedresponsepolicyzones object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Orderedresponsepolicyzones) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Orderedresponsepolicyzones) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Orderedresponsepolicyzones) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Orderedresponsepolicyzones) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRpZones

`func (o *Orderedresponsepolicyzones) GetRpZones() []string`

GetRpZones returns the RpZones field if non-nil, zero value otherwise.

### GetRpZonesOk

`func (o *Orderedresponsepolicyzones) GetRpZonesOk() (*[]string, bool)`

GetRpZonesOk returns a tuple with the RpZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpZones

`func (o *Orderedresponsepolicyzones) SetRpZones(v []string)`

SetRpZones sets RpZones field to given value.

### HasRpZones

`func (o *Orderedresponsepolicyzones) HasRpZones() bool`

HasRpZones returns a boolean if a field has been set.

### GetView

`func (o *Orderedresponsepolicyzones) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Orderedresponsepolicyzones) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Orderedresponsepolicyzones) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *Orderedresponsepolicyzones) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


