# ListDbObjectsResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]DbObjects**](DbObjects.md) |  | [optional] 

## Methods

### NewListDbObjectsResponseObject

`func NewListDbObjectsResponseObject() *ListDbObjectsResponseObject`

NewListDbObjectsResponseObject instantiates a new ListDbObjectsResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDbObjectsResponseObjectWithDefaults

`func NewListDbObjectsResponseObjectWithDefaults() *ListDbObjectsResponseObject`

NewListDbObjectsResponseObjectWithDefaults instantiates a new ListDbObjectsResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListDbObjectsResponseObject) GetResult() []DbObjects`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListDbObjectsResponseObject) GetResultOk() (*[]DbObjects, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListDbObjectsResponseObject) SetResult(v []DbObjects)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListDbObjectsResponseObject) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


