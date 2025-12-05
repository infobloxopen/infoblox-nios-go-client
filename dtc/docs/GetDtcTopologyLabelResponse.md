# GetDtcTopologyLabelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Field** | Pointer to **string** | The name of the field in the Topology database the label was obtained from. | [optional] [readonly] 
**Label** | Pointer to **string** | The DTC Topology label name. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**DtcTopologyLabel**](DtcTopologyLabel.md) |  | [optional] 

## Methods

### NewGetDtcTopologyLabelResponse

`func NewGetDtcTopologyLabelResponse() *GetDtcTopologyLabelResponse`

NewGetDtcTopologyLabelResponse instantiates a new GetDtcTopologyLabelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcTopologyLabelResponseWithDefaults

`func NewGetDtcTopologyLabelResponseWithDefaults() *GetDtcTopologyLabelResponse`

NewGetDtcTopologyLabelResponseWithDefaults instantiates a new GetDtcTopologyLabelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcTopologyLabelResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcTopologyLabelResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcTopologyLabelResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcTopologyLabelResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetField

`func (o *GetDtcTopologyLabelResponse) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *GetDtcTopologyLabelResponse) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *GetDtcTopologyLabelResponse) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *GetDtcTopologyLabelResponse) HasField() bool`

HasField returns a boolean if a field has been set.

### GetLabel

`func (o *GetDtcTopologyLabelResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetDtcTopologyLabelResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetDtcTopologyLabelResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetDtcTopologyLabelResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUuid

`func (o *GetDtcTopologyLabelResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetDtcTopologyLabelResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetDtcTopologyLabelResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetDtcTopologyLabelResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcTopologyLabelResponse) GetResult() DtcTopologyLabel`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcTopologyLabelResponse) GetResultOk() (*DtcTopologyLabel, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcTopologyLabelResponse) SetResult(v DtcTopologyLabel)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcTopologyLabelResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


