# OutboundCloudclient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Enable** | Pointer to **bool** | Determines whether the OutBound Cloud Client is enabled. | [optional] 
**GridMember** | Pointer to **string** | The Grid member where our outbound is running. | [optional] 
**Interval** | Pointer to **int64** | The time interval (in seconds) for requesting newly detected domains by the Infoblox Outbound Cloud Client and applying them to the list of configured RPZs. | [optional] 
**OutboundCloudClientEvents** | Pointer to [**[]OutboundCloudclientOutboundCloudClientEvents**](OutboundCloudclientOutboundCloudClientEvents.md) | List of event types to request | [optional] 

## Methods

### NewOutboundCloudclient

`func NewOutboundCloudclient() *OutboundCloudclient`

NewOutboundCloudclient instantiates a new OutboundCloudclient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutboundCloudclientWithDefaults

`func NewOutboundCloudclientWithDefaults() *OutboundCloudclient`

NewOutboundCloudclientWithDefaults instantiates a new OutboundCloudclient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *OutboundCloudclient) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *OutboundCloudclient) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *OutboundCloudclient) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *OutboundCloudclient) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetEnable

`func (o *OutboundCloudclient) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OutboundCloudclient) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OutboundCloudclient) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *OutboundCloudclient) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGridMember

`func (o *OutboundCloudclient) GetGridMember() string`

GetGridMember returns the GridMember field if non-nil, zero value otherwise.

### GetGridMemberOk

`func (o *OutboundCloudclient) GetGridMemberOk() (*string, bool)`

GetGridMemberOk returns a tuple with the GridMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridMember

`func (o *OutboundCloudclient) SetGridMember(v string)`

SetGridMember sets GridMember field to given value.

### HasGridMember

`func (o *OutboundCloudclient) HasGridMember() bool`

HasGridMember returns a boolean if a field has been set.

### GetInterval

`func (o *OutboundCloudclient) GetInterval() int64`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *OutboundCloudclient) GetIntervalOk() (*int64, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *OutboundCloudclient) SetInterval(v int64)`

SetInterval sets Interval field to given value.

### HasInterval

`func (o *OutboundCloudclient) HasInterval() bool`

HasInterval returns a boolean if a field has been set.

### GetOutboundCloudClientEvents

`func (o *OutboundCloudclient) GetOutboundCloudClientEvents() []OutboundCloudclientOutboundCloudClientEvents`

GetOutboundCloudClientEvents returns the OutboundCloudClientEvents field if non-nil, zero value otherwise.

### GetOutboundCloudClientEventsOk

`func (o *OutboundCloudclient) GetOutboundCloudClientEventsOk() (*[]OutboundCloudclientOutboundCloudClientEvents, bool)`

GetOutboundCloudClientEventsOk returns a tuple with the OutboundCloudClientEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundCloudClientEvents

`func (o *OutboundCloudclient) SetOutboundCloudClientEvents(v []OutboundCloudclientOutboundCloudClientEvents)`

SetOutboundCloudClientEvents sets OutboundCloudClientEvents field to given value.

### HasOutboundCloudClientEvents

`func (o *OutboundCloudclient) HasOutboundCloudClientEvents() bool`

HasOutboundCloudClientEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


