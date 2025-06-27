# GetSmartfolderChildrenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Resource** | Pointer to **string** | The object retuned by the Smart Folder query. | [optional] [readonly] 
**Value** | Pointer to [**SmartfolderChildrenValue**](SmartfolderChildrenValue.md) |  | [optional] 
**ValueType** | Pointer to **string** | The type of the returned value. | [optional] [readonly] 
**Result** | Pointer to [**SmartfolderChildren**](SmartfolderChildren.md) |  | [optional] 

## Methods

### NewGetSmartfolderChildrenResponse

`func NewGetSmartfolderChildrenResponse() *GetSmartfolderChildrenResponse`

NewGetSmartfolderChildrenResponse instantiates a new GetSmartfolderChildrenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSmartfolderChildrenResponseWithDefaults

`func NewGetSmartfolderChildrenResponseWithDefaults() *GetSmartfolderChildrenResponse`

NewGetSmartfolderChildrenResponseWithDefaults instantiates a new GetSmartfolderChildrenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetSmartfolderChildrenResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetSmartfolderChildrenResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetSmartfolderChildrenResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetSmartfolderChildrenResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetResource

`func (o *GetSmartfolderChildrenResponse) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetSmartfolderChildrenResponse) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetSmartfolderChildrenResponse) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *GetSmartfolderChildrenResponse) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetValue

`func (o *GetSmartfolderChildrenResponse) GetValue() SmartfolderChildrenValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetSmartfolderChildrenResponse) GetValueOk() (*SmartfolderChildrenValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetSmartfolderChildrenResponse) SetValue(v SmartfolderChildrenValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetSmartfolderChildrenResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueType

`func (o *GetSmartfolderChildrenResponse) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *GetSmartfolderChildrenResponse) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *GetSmartfolderChildrenResponse) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *GetSmartfolderChildrenResponse) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetResult

`func (o *GetSmartfolderChildrenResponse) GetResult() SmartfolderChildren`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSmartfolderChildrenResponse) GetResultOk() (*SmartfolderChildren, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSmartfolderChildrenResponse) SetResult(v SmartfolderChildren)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSmartfolderChildrenResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


