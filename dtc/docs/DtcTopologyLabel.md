# DtcTopologyLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Field** | Pointer to **string** | The name of the field in the Topology database the label was obtained from. | [optional] [readonly] 
**Label** | Pointer to **string** | The DTC Topology label name. | [optional] [readonly] 

## Methods

### NewDtcTopologyLabel

`func NewDtcTopologyLabel() *DtcTopologyLabel`

NewDtcTopologyLabel instantiates a new DtcTopologyLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcTopologyLabelWithDefaults

`func NewDtcTopologyLabelWithDefaults() *DtcTopologyLabel`

NewDtcTopologyLabelWithDefaults instantiates a new DtcTopologyLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcTopologyLabel) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcTopologyLabel) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcTopologyLabel) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcTopologyLabel) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetField

`func (o *DtcTopologyLabel) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *DtcTopologyLabel) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *DtcTopologyLabel) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *DtcTopologyLabel) HasField() bool`

HasField returns a boolean if a field has been set.

### GetLabel

`func (o *DtcTopologyLabel) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *DtcTopologyLabel) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *DtcTopologyLabel) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *DtcTopologyLabel) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


