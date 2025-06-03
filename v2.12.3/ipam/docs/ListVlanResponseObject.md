# ListVlanResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]Vlan**](Vlan.md) |  | [optional] 

## Methods

### NewListVlanResponseObject

`func NewListVlanResponseObject() *ListVlanResponseObject`

NewListVlanResponseObject instantiates a new ListVlanResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVlanResponseObjectWithDefaults

`func NewListVlanResponseObjectWithDefaults() *ListVlanResponseObject`

NewListVlanResponseObjectWithDefaults instantiates a new ListVlanResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListVlanResponseObject) GetResult() []Vlan`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListVlanResponseObject) GetResultOk() (*[]Vlan, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListVlanResponseObject) SetResult(v []Vlan)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListVlanResponseObject) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


