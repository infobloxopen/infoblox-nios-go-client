# GetDtcTopologyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Comment** | Pointer to **string** | The comment for the DTC TOPOLOGY monitor object; maximum 256 characters. | [optional] 
**Extattrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Name** | Pointer to **string** | Display name of the DTC Topology. | [optional] 
**Rules** | Pointer to **[]map[string]interface{}** | Topology rules. | [optional] 
**Result** | Pointer to [**DtcTopology**](DtcTopology.md) |  | [optional] 

## Methods

### NewGetDtcTopologyResponse

`func NewGetDtcTopologyResponse() *GetDtcTopologyResponse`

NewGetDtcTopologyResponse instantiates a new GetDtcTopologyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDtcTopologyResponseWithDefaults

`func NewGetDtcTopologyResponseWithDefaults() *GetDtcTopologyResponse`

NewGetDtcTopologyResponseWithDefaults instantiates a new GetDtcTopologyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDtcTopologyResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDtcTopologyResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDtcTopologyResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDtcTopologyResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetComment

`func (o *GetDtcTopologyResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetDtcTopologyResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetDtcTopologyResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetDtcTopologyResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExtattrs

`func (o *GetDtcTopologyResponse) GetExtattrs() map[string]ExtAttrs`

GetExtattrs returns the Extattrs field if non-nil, zero value otherwise.

### GetExtattrsOk

`func (o *GetDtcTopologyResponse) GetExtattrsOk() (*map[string]ExtAttrs, bool)`

GetExtattrsOk returns a tuple with the Extattrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtattrs

`func (o *GetDtcTopologyResponse) SetExtattrs(v map[string]ExtAttrs)`

SetExtattrs sets Extattrs field to given value.

### HasExtattrs

`func (o *GetDtcTopologyResponse) HasExtattrs() bool`

HasExtattrs returns a boolean if a field has been set.

### GetName

`func (o *GetDtcTopologyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDtcTopologyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDtcTopologyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDtcTopologyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRules

`func (o *GetDtcTopologyResponse) GetRules() []map[string]interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetDtcTopologyResponse) GetRulesOk() (*[]map[string]interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetDtcTopologyResponse) SetRules(v []map[string]interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetDtcTopologyResponse) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetResult

`func (o *GetDtcTopologyResponse) GetResult() DtcTopology`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDtcTopologyResponse) GetResultOk() (*DtcTopology, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDtcTopologyResponse) SetResult(v DtcTopology)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDtcTopologyResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


