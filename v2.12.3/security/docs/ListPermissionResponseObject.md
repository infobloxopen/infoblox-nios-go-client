# ListPermissionResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]Permission**](Permission.md) |  | [optional] 

## Methods

### NewListPermissionResponseObject

`func NewListPermissionResponseObject() *ListPermissionResponseObject`

NewListPermissionResponseObject instantiates a new ListPermissionResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPermissionResponseObjectWithDefaults

`func NewListPermissionResponseObjectWithDefaults() *ListPermissionResponseObject`

NewListPermissionResponseObjectWithDefaults instantiates a new ListPermissionResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListPermissionResponseObject) GetResult() []Permission`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListPermissionResponseObject) GetResultOk() (*[]Permission, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListPermissionResponseObject) SetResult(v []Permission)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListPermissionResponseObject) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


