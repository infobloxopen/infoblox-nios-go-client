# NsgroupForwardingmemberForwardingServers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this Grid member in FQDN format. | [optional] 
**ForwardersOnly** | Pointer to **bool** | Determines if the appliance sends queries to forwarders only, and not to other internal or Internet root servers. | [optional] 
**ForwardTo** | Pointer to [**NsgroupforwardingmemberforwardingserversForwardTo**](NsgroupforwardingmemberforwardingserversForwardTo.md) |  | [optional] 
**UseOverrideForwarders** | Pointer to **bool** | Use flag for: forward_to | [optional] 

## Methods

### NewNsgroupForwardingmemberForwardingServers

`func NewNsgroupForwardingmemberForwardingServers() *NsgroupForwardingmemberForwardingServers`

NewNsgroupForwardingmemberForwardingServers instantiates a new NsgroupForwardingmemberForwardingServers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNsgroupForwardingmemberForwardingServersWithDefaults

`func NewNsgroupForwardingmemberForwardingServersWithDefaults() *NsgroupForwardingmemberForwardingServers`

NewNsgroupForwardingmemberForwardingServersWithDefaults instantiates a new NsgroupForwardingmemberForwardingServers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NsgroupForwardingmemberForwardingServers) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NsgroupForwardingmemberForwardingServers) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NsgroupForwardingmemberForwardingServers) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NsgroupForwardingmemberForwardingServers) HasName() bool`

HasName returns a boolean if a field has been set.

### GetForwardersOnly

`func (o *NsgroupForwardingmemberForwardingServers) GetForwardersOnly() bool`

GetForwardersOnly returns the ForwardersOnly field if non-nil, zero value otherwise.

### GetForwardersOnlyOk

`func (o *NsgroupForwardingmemberForwardingServers) GetForwardersOnlyOk() (*bool, bool)`

GetForwardersOnlyOk returns a tuple with the ForwardersOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardersOnly

`func (o *NsgroupForwardingmemberForwardingServers) SetForwardersOnly(v bool)`

SetForwardersOnly sets ForwardersOnly field to given value.

### HasForwardersOnly

`func (o *NsgroupForwardingmemberForwardingServers) HasForwardersOnly() bool`

HasForwardersOnly returns a boolean if a field has been set.

### GetForwardTo

`func (o *NsgroupForwardingmemberForwardingServers) GetForwardTo() NsgroupforwardingmemberforwardingserversForwardTo`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *NsgroupForwardingmemberForwardingServers) GetForwardToOk() (*NsgroupforwardingmemberforwardingserversForwardTo, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *NsgroupForwardingmemberForwardingServers) SetForwardTo(v NsgroupforwardingmemberforwardingserversForwardTo)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *NsgroupForwardingmemberForwardingServers) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetUseOverrideForwarders

`func (o *NsgroupForwardingmemberForwardingServers) GetUseOverrideForwarders() bool`

GetUseOverrideForwarders returns the UseOverrideForwarders field if non-nil, zero value otherwise.

### GetUseOverrideForwardersOk

`func (o *NsgroupForwardingmemberForwardingServers) GetUseOverrideForwardersOk() (*bool, bool)`

GetUseOverrideForwardersOk returns a tuple with the UseOverrideForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseOverrideForwarders

`func (o *NsgroupForwardingmemberForwardingServers) SetUseOverrideForwarders(v bool)`

SetUseOverrideForwarders sets UseOverrideForwarders field to given value.

### HasUseOverrideForwarders

`func (o *NsgroupForwardingmemberForwardingServers) HasUseOverrideForwarders() bool`

HasUseOverrideForwarders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


