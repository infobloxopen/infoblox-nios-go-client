# GetOrderedresponsepolicyzonesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**RpZones** | Pointer to **[]string** | An ordered list of Response Policy Zone names. | [optional] 
**View** | Pointer to **string** | The DNS View name. | [optional] 
**Result** | Pointer to [**Orderedresponsepolicyzones**](Orderedresponsepolicyzones.md) |  | [optional] 

## Methods

### NewGetOrderedresponsepolicyzonesResponse

`func NewGetOrderedresponsepolicyzonesResponse() *GetOrderedresponsepolicyzonesResponse`

NewGetOrderedresponsepolicyzonesResponse instantiates a new GetOrderedresponsepolicyzonesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderedresponsepolicyzonesResponseWithDefaults

`func NewGetOrderedresponsepolicyzonesResponseWithDefaults() *GetOrderedresponsepolicyzonesResponse`

NewGetOrderedresponsepolicyzonesResponseWithDefaults instantiates a new GetOrderedresponsepolicyzonesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetOrderedresponsepolicyzonesResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetOrderedresponsepolicyzonesResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetOrderedresponsepolicyzonesResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetOrderedresponsepolicyzonesResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRpZones

`func (o *GetOrderedresponsepolicyzonesResponse) GetRpZones() []string`

GetRpZones returns the RpZones field if non-nil, zero value otherwise.

### GetRpZonesOk

`func (o *GetOrderedresponsepolicyzonesResponse) GetRpZonesOk() (*[]string, bool)`

GetRpZonesOk returns a tuple with the RpZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRpZones

`func (o *GetOrderedresponsepolicyzonesResponse) SetRpZones(v []string)`

SetRpZones sets RpZones field to given value.

### HasRpZones

`func (o *GetOrderedresponsepolicyzonesResponse) HasRpZones() bool`

HasRpZones returns a boolean if a field has been set.

### GetView

`func (o *GetOrderedresponsepolicyzonesResponse) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *GetOrderedresponsepolicyzonesResponse) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *GetOrderedresponsepolicyzonesResponse) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *GetOrderedresponsepolicyzonesResponse) HasView() bool`

HasView returns a boolean if a field has been set.

### GetResult

`func (o *GetOrderedresponsepolicyzonesResponse) GetResult() Orderedresponsepolicyzones`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetOrderedresponsepolicyzonesResponse) GetResultOk() (*Orderedresponsepolicyzones, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetOrderedresponsepolicyzonesResponse) SetResult(v Orderedresponsepolicyzones)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetOrderedresponsepolicyzonesResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


