# Orderedranges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Network** | Pointer to **string** | The reference to the network that contains ranges. | [optional] [readonly] 
**Ranges** | Pointer to **[]string** | The ordered list of references to ranges. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewOrderedranges

`func NewOrderedranges() *Orderedranges`

NewOrderedranges instantiates a new Orderedranges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderedrangesWithDefaults

`func NewOrderedrangesWithDefaults() *Orderedranges`

NewOrderedrangesWithDefaults instantiates a new Orderedranges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Orderedranges) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Orderedranges) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Orderedranges) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Orderedranges) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetNetwork

`func (o *Orderedranges) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Orderedranges) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Orderedranges) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Orderedranges) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRanges

`func (o *Orderedranges) GetRanges() []string`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *Orderedranges) GetRangesOk() (*[]string, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *Orderedranges) SetRanges(v []string)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *Orderedranges) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetUuid

`func (o *Orderedranges) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Orderedranges) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Orderedranges) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Orderedranges) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


