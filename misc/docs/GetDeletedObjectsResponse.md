# GetDeletedObjectsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**ObjectType** | Pointer to **string** | The object type of the deleted object. This is undefined if the object is not supported. | [optional] [readonly] 
**Result** | Pointer to [**DeletedObjects**](DeletedObjects.md) |  | [optional] 

## Methods

### NewGetDeletedObjectsResponse

`func NewGetDeletedObjectsResponse() *GetDeletedObjectsResponse`

NewGetDeletedObjectsResponse instantiates a new GetDeletedObjectsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeletedObjectsResponseWithDefaults

`func NewGetDeletedObjectsResponseWithDefaults() *GetDeletedObjectsResponse`

NewGetDeletedObjectsResponseWithDefaults instantiates a new GetDeletedObjectsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDeletedObjectsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDeletedObjectsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDeletedObjectsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDeletedObjectsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetObjectType

`func (o *GetDeletedObjectsResponse) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *GetDeletedObjectsResponse) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *GetDeletedObjectsResponse) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *GetDeletedObjectsResponse) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetResult

`func (o *GetDeletedObjectsResponse) GetResult() DeletedObjects`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDeletedObjectsResponse) GetResultOk() (*DeletedObjects, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDeletedObjectsResponse) SetResult(v DeletedObjects)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDeletedObjectsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


