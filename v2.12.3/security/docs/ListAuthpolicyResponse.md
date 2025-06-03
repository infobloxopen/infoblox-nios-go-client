# ListAuthpolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]Authpolicy**](Authpolicy.md) |  | [optional] 

## Methods

### NewListAuthpolicyResponse

`func NewListAuthpolicyResponse() *ListAuthpolicyResponse`

NewListAuthpolicyResponse instantiates a new ListAuthpolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthpolicyResponseWithDefaults

`func NewListAuthpolicyResponseWithDefaults() *ListAuthpolicyResponse`

NewListAuthpolicyResponseWithDefaults instantiates a new ListAuthpolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListAuthpolicyResponse) GetResult() []Authpolicy`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListAuthpolicyResponse) GetResultOk() (*[]Authpolicy, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListAuthpolicyResponse) SetResult(v []Authpolicy)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListAuthpolicyResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


