# ListDiscoveryResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]Discovery**](Discovery.md) |  | [optional] 

## Methods

### NewListDiscoveryResponseObject

`func NewListDiscoveryResponseObject() *ListDiscoveryResponseObject`

NewListDiscoveryResponseObject instantiates a new ListDiscoveryResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDiscoveryResponseObjectWithDefaults

`func NewListDiscoveryResponseObjectWithDefaults() *ListDiscoveryResponseObject`

NewListDiscoveryResponseObjectWithDefaults instantiates a new ListDiscoveryResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListDiscoveryResponseObject) GetResult() []Discovery`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListDiscoveryResponseObject) GetResultOk() (*[]Discovery, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListDiscoveryResponseObject) SetResult(v []Discovery)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListDiscoveryResponseObject) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


