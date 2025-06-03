# ListDtcServerResponseObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**[]DtcServer**](DtcServer.md) |  | [optional] 

## Methods

### NewListDtcServerResponseObject

`func NewListDtcServerResponseObject() *ListDtcServerResponseObject`

NewListDtcServerResponseObject instantiates a new ListDtcServerResponseObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDtcServerResponseObjectWithDefaults

`func NewListDtcServerResponseObjectWithDefaults() *ListDtcServerResponseObject`

NewListDtcServerResponseObjectWithDefaults instantiates a new ListDtcServerResponseObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ListDtcServerResponseObject) GetResult() []DtcServer`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListDtcServerResponseObject) GetResultOk() (*[]DtcServer, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListDtcServerResponseObject) SetResult(v []DtcServer)`

SetResult sets Result field to given value.

### HasResult

`func (o *ListDtcServerResponseObject) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


