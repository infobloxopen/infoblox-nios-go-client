# GetParentalcontrolBlockingpolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the blocking policy. | [optional] 
**Value** | Pointer to **string** | The 32 bit hex value of the blocking policy. | [optional] 
**Result** | Pointer to [**ParentalcontrolBlockingpolicy**](ParentalcontrolBlockingpolicy.md) |  | [optional] 

## Methods

### NewGetParentalcontrolBlockingpolicyResponse

`func NewGetParentalcontrolBlockingpolicyResponse() *GetParentalcontrolBlockingpolicyResponse`

NewGetParentalcontrolBlockingpolicyResponse instantiates a new GetParentalcontrolBlockingpolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetParentalcontrolBlockingpolicyResponseWithDefaults

`func NewGetParentalcontrolBlockingpolicyResponseWithDefaults() *GetParentalcontrolBlockingpolicyResponse`

NewGetParentalcontrolBlockingpolicyResponseWithDefaults instantiates a new GetParentalcontrolBlockingpolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetParentalcontrolBlockingpolicyResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetParentalcontrolBlockingpolicyResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetParentalcontrolBlockingpolicyResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetParentalcontrolBlockingpolicyResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetName

`func (o *GetParentalcontrolBlockingpolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetParentalcontrolBlockingpolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetParentalcontrolBlockingpolicyResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetParentalcontrolBlockingpolicyResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValue

`func (o *GetParentalcontrolBlockingpolicyResponse) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetParentalcontrolBlockingpolicyResponse) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetParentalcontrolBlockingpolicyResponse) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetParentalcontrolBlockingpolicyResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetResult

`func (o *GetParentalcontrolBlockingpolicyResponse) GetResult() ParentalcontrolBlockingpolicy`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetParentalcontrolBlockingpolicyResponse) GetResultOk() (*ParentalcontrolBlockingpolicy, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetParentalcontrolBlockingpolicyResponse) SetResult(v ParentalcontrolBlockingpolicy)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetParentalcontrolBlockingpolicyResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


