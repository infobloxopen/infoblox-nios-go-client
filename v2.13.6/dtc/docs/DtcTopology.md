# DtcTopology

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The comment for the DTC TOPOLOGY monitor object; maximum 256 characters. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | Display name of the DTC Topology. | [optional] 
**Rules** | Pointer to **[]map[string]interface{}** | Topology rules. | [optional] 

## Methods

### NewDtcTopology

`func NewDtcTopology() *DtcTopology`

NewDtcTopology instantiates a new DtcTopology object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDtcTopologyWithDefaults

`func NewDtcTopologyWithDefaults() *DtcTopology`

NewDtcTopologyWithDefaults instantiates a new DtcTopology object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *DtcTopology) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *DtcTopology) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *DtcTopology) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *DtcTopology) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *DtcTopology) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DtcTopology) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DtcTopology) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DtcTopology) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *DtcTopology) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *DtcTopology) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *DtcTopology) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *DtcTopology) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetName

`func (o *DtcTopology) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DtcTopology) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DtcTopology) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DtcTopology) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *DtcTopology) GetRules() []map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *DtcTopology) GetRulesOk() (*[]map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *DtcTopology) SetRules(v []map[string]interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *DtcTopology) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


