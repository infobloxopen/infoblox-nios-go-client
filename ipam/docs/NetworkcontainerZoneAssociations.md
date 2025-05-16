# NetworkcontainerZoneAssociations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fqdn** | Pointer to **string** | The FQDN of the authoritative forward zone. | [optional] 
**IsDefault** | Pointer to **bool** | True if this is the default zone. | [optional] 
**View** | Pointer to **string** | The view to which the zone belongs. If a view is not specified, the default view is used. | [optional] 

## Methods

### NewNetworkcontainerZoneAssociations

`func NewNetworkcontainerZoneAssociations() *NetworkcontainerZoneAssociations`

NewNetworkcontainerZoneAssociations instantiates a new NetworkcontainerZoneAssociations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkcontainerZoneAssociationsWithDefaults

`func NewNetworkcontainerZoneAssociationsWithDefaults() *NetworkcontainerZoneAssociations`

NewNetworkcontainerZoneAssociationsWithDefaults instantiates a new NetworkcontainerZoneAssociations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFqdn

`func (o *NetworkcontainerZoneAssociations) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *NetworkcontainerZoneAssociations) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *NetworkcontainerZoneAssociations) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *NetworkcontainerZoneAssociations) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetIsDefault

`func (o *NetworkcontainerZoneAssociations) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *NetworkcontainerZoneAssociations) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *NetworkcontainerZoneAssociations) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *NetworkcontainerZoneAssociations) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetView

`func (o *NetworkcontainerZoneAssociations) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *NetworkcontainerZoneAssociations) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *NetworkcontainerZoneAssociations) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *NetworkcontainerZoneAssociations) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


