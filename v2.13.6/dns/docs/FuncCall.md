# FuncCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeName** | **string** | The attribute to be called. | 
**ObjectFunction** | Pointer to **string** | The function to be called. | [optional] 
**Parameters** | Pointer to **map[string]interface{}** | The parameters for the function. | [optional] 
**ResultField** | Pointer to **string** | The result field of the function. | [optional] 
**Object** | Pointer to **string** | The object to be called. | [optional] 
**ObjectParameters** | Pointer to **map[string]interface{}** | The parameters for the object. | [optional] 

## Methods

### NewFuncCall

`func NewFuncCall(attributeName string, ) *FuncCall`

NewFuncCall instantiates a new FuncCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuncCallWithDefaults

`func NewFuncCallWithDefaults() *FuncCall`

NewFuncCallWithDefaults instantiates a new FuncCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeName

`func (o *FuncCall) GetAttributeName() string`

GetAttributeName returns the AttributeName field if non-nil, zero value otherwise.

### GetAttributeNameOk

`func (o *FuncCall) GetAttributeNameOk() (*string, bool)`

GetAttributeNameOk returns a tuple with the AttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeName

`func (o *FuncCall) SetAttributeName(v string)`

SetAttributeName sets AttributeName field to given value.


### GetObjectFunction

`func (o *FuncCall) GetObjectFunction() string`

GetObjectFunction returns the ObjectFunction field if non-nil, zero value otherwise.

### GetObjectFunctionOk

`func (o *FuncCall) GetObjectFunctionOk() (*string, bool)`

GetObjectFunctionOk returns a tuple with the ObjectFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectFunction

`func (o *FuncCall) SetObjectFunction(v string)`

SetObjectFunction sets ObjectFunction field to given value.

### HasObjectFunction

`func (o *FuncCall) HasObjectFunction() bool`

HasObjectFunction returns a boolean if a field has been set.

### GetParameters

`func (o *FuncCall) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *FuncCall) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *FuncCall) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *FuncCall) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetResultField

`func (o *FuncCall) GetResultField() string`

GetResultField returns the ResultField field if non-nil, zero value otherwise.

### GetResultFieldOk

`func (o *FuncCall) GetResultFieldOk() (*string, bool)`

GetResultFieldOk returns a tuple with the ResultField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultField

`func (o *FuncCall) SetResultField(v string)`

SetResultField sets ResultField field to given value.

### HasResultField

`func (o *FuncCall) HasResultField() bool`

HasResultField returns a boolean if a field has been set.

### GetObject

`func (o *FuncCall) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *FuncCall) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *FuncCall) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *FuncCall) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectParameters

`func (o *FuncCall) GetObjectParameters() map[string]interface{}`

GetObjectParameters returns the ObjectParameters field if non-nil, zero value otherwise.

### GetObjectParametersOk

`func (o *FuncCall) GetObjectParametersOk() (*map[string]interface{}, bool)`

GetObjectParametersOk returns a tuple with the ObjectParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectParameters

`func (o *FuncCall) SetObjectParameters(v map[string]interface{})`

SetObjectParameters sets ObjectParameters field to given value.

### HasObjectParameters

`func (o *FuncCall) HasObjectParameters() bool`

HasObjectParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


