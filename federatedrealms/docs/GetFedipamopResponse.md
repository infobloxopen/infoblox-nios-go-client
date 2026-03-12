# GetFedipamopResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**Fedipamop**](Fedipamop.md) |  | [optional] 

## Methods

### NewGetFedipamopResponse

`func NewGetFedipamopResponse() *GetFedipamopResponse`

NewGetFedipamopResponse instantiates a new GetFedipamopResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFedipamopResponseWithDefaults

`func NewGetFedipamopResponseWithDefaults() *GetFedipamopResponse`

NewGetFedipamopResponseWithDefaults instantiates a new GetFedipamopResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFedipamopResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFedipamopResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFedipamopResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFedipamopResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetUuid

`func (o *GetFedipamopResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetFedipamopResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetFedipamopResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetFedipamopResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetFedipamopResponse) GetResult() Fedipamop`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFedipamopResponse) GetResultOk() (*Fedipamop, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFedipamopResponse) SetResult(v Fedipamop)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFedipamopResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


