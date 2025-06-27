# NetworkNetwork

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectFunction** | Pointer to **string** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ResultField** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**ObjectParameters** | Pointer to **map[string]interface{}** |  | [optional] 
**ObjectRef** | Pointer to **string** | A WAPI object reference on which the function calls. Either _object or _object_ref must be set. | [optional] 

## Methods

### NewNetworkNetwork

`func NewNetworkNetwork() *NetworkNetwork`

NewNetworkNetwork instantiates a new NetworkNetwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkNetworkWithDefaults

`func NewNetworkNetworkWithDefaults() *NetworkNetwork`

NewNetworkNetworkWithDefaults instantiates a new NetworkNetwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectFunction

`func (o *NetworkNetwork) GetObjectFunction() string`

GetObjectFunction returns the ObjectFunction field if non-nil, zero value otherwise.

### GetObjectFunctionOk

`func (o *NetworkNetwork) GetObjectFunctionOk() (*string, bool)`

GetObjectFunctionOk returns a tuple with the ObjectFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectFunction

`func (o *NetworkNetwork) SetObjectFunction(v string)`

SetObjectFunction sets ObjectFunction field to given value.

### HasObjectFunction

`func (o *NetworkNetwork) HasObjectFunction() bool`

HasObjectFunction returns a boolean if a field has been set.

### GetParameters

`func (o *NetworkNetwork) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *NetworkNetwork) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *NetworkNetwork) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *NetworkNetwork) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetResultField

`func (o *NetworkNetwork) GetResultField() string`

GetResultField returns the ResultField field if non-nil, zero value otherwise.

### GetResultFieldOk

`func (o *NetworkNetwork) GetResultFieldOk() (*string, bool)`

GetResultFieldOk returns a tuple with the ResultField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultField

`func (o *NetworkNetwork) SetResultField(v string)`

SetResultField sets ResultField field to given value.

### HasResultField

`func (o *NetworkNetwork) HasResultField() bool`

HasResultField returns a boolean if a field has been set.

### GetObject

`func (o *NetworkNetwork) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *NetworkNetwork) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *NetworkNetwork) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *NetworkNetwork) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetObjectParameters

`func (o *NetworkNetwork) GetObjectParameters() map[string]interface{}`

GetObjectParameters returns the ObjectParameters field if non-nil, zero value otherwise.

### GetObjectParametersOk

`func (o *NetworkNetwork) GetObjectParametersOk() (*map[string]interface{}, bool)`

GetObjectParametersOk returns a tuple with the ObjectParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectParameters

`func (o *NetworkNetwork) SetObjectParameters(v map[string]interface{})`

SetObjectParameters sets ObjectParameters field to given value.

### HasObjectParameters

`func (o *NetworkNetwork) HasObjectParameters() bool`

HasObjectParameters returns a boolean if a field has been set.

### GetObjectRef

`func (o *NetworkNetwork) GetObjectRef() string`

GetObjectRef returns the ObjectRef field if non-nil, zero value otherwise.

### GetObjectRefOk

`func (o *NetworkNetwork) GetObjectRefOk() (*string, bool)`

GetObjectRefOk returns a tuple with the ObjectRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectRef

`func (o *NetworkNetwork) SetObjectRef(v string)`

SetObjectRef sets ObjectRef field to given value.

### HasObjectRef

`func (o *NetworkNetwork) HasObjectRef() bool`

HasObjectRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


