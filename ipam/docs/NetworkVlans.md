# NetworkVlans

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vlan** | Pointer to **string** | Reference to the underlying StaticVlan object vlan. | [optional] 
**Id** | Pointer to **int64** | VLAN ID value. | [optional] [readonly] 
**Name** | Pointer to **string** | Name of the VLAN. | [optional] [readonly] 

## Methods

### NewNetworkVlans

`func NewNetworkVlans() *NetworkVlans`

NewNetworkVlans instantiates a new NetworkVlans object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkVlansWithDefaults

`func NewNetworkVlansWithDefaults() *NetworkVlans`

NewNetworkVlansWithDefaults instantiates a new NetworkVlans object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlan

`func (o *NetworkVlans) GetVlan() string`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *NetworkVlans) GetVlanOk() (*string, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *NetworkVlans) SetVlan(v string)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *NetworkVlans) HasVlan() bool`

HasVlan returns a boolean if a field has been set.

### GetId

`func (o *NetworkVlans) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkVlans) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkVlans) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkVlans) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkVlans) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkVlans) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkVlans) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkVlans) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


