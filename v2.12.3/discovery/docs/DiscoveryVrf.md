# DiscoveryVrf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Description** | Pointer to **string** | Additional information about the VRF. | [optional] [readonly] 
**Device** | Pointer to **string** | The device to which the VRF belongs. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the VRF. | [optional] [readonly] 
**NetworkView** | Pointer to **string** | The name of the network view in which this VRF resides. | [optional] [readonly] 
**RouteDistinguisher** | Pointer to **string** | The route distinguisher associated with the VRF. | [optional] [readonly] 

## Methods

### NewDiscoveryVrf

`func NewDiscoveryVrf() *DiscoveryVrf`

NewDiscoveryVrf instantiates a new DiscoveryVrf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscoveryVrfWithDefaults

`func NewDiscoveryVrfWithDefaults() *DiscoveryVrf`

NewDiscoveryVrfWithDefaults instantiates a new DiscoveryVrf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DiscoveryVrf) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DiscoveryVrf) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DiscoveryVrf) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DiscoveryVrf) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDescription

`func (o *DiscoveryVrf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscoveryVrf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscoveryVrf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscoveryVrf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDevice

`func (o *DiscoveryVrf) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *DiscoveryVrf) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *DiscoveryVrf) SetDevice(v string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *DiscoveryVrf) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetName

`func (o *DiscoveryVrf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscoveryVrf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscoveryVrf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscoveryVrf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkView

`func (o *DiscoveryVrf) GetNetworkView() string`

GetNetworkView returns the NetworkView field if non-nil, zero value otherwise.

### GetNetworkViewOk

`func (o *DiscoveryVrf) GetNetworkViewOk() (*string, bool)`

GetNetworkViewOk returns a tuple with the NetworkView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkView

`func (o *DiscoveryVrf) SetNetworkView(v string)`

SetNetworkView sets NetworkView field to given value.

### HasNetworkView

`func (o *DiscoveryVrf) HasNetworkView() bool`

HasNetworkView returns a boolean if a field has been set.

### GetRouteDistinguisher

`func (o *DiscoveryVrf) GetRouteDistinguisher() string`

GetRouteDistinguisher returns the RouteDistinguisher field if non-nil, zero value otherwise.

### GetRouteDistinguisherOk

`func (o *DiscoveryVrf) GetRouteDistinguisherOk() (*string, bool)`

GetRouteDistinguisherOk returns a tuple with the RouteDistinguisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteDistinguisher

`func (o *DiscoveryVrf) SetRouteDistinguisher(v string)`

SetRouteDistinguisher sets RouteDistinguisher field to given value.

### HasRouteDistinguisher

`func (o *DiscoveryVrf) HasRouteDistinguisher() bool`

HasRouteDistinguisher returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


