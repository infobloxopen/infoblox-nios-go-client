# GetFederatedrealmsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**Id** | Pointer to **string** | Federated realm id. | [optional] [readonly] 
**Name** | Pointer to **string** | Federated realm name. | [optional] [readonly] 
**Result** | Pointer to [**Federatedrealms**](Federatedrealms.md) |  | [optional] 

## Methods

### NewGetFederatedrealmsResponse

`func NewGetFederatedrealmsResponse() *GetFederatedrealmsResponse`

NewGetFederatedrealmsResponse instantiates a new GetFederatedrealmsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFederatedrealmsResponseWithDefaults

`func NewGetFederatedrealmsResponseWithDefaults() *GetFederatedrealmsResponse`

NewGetFederatedrealmsResponseWithDefaults instantiates a new GetFederatedrealmsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetFederatedrealmsResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetFederatedrealmsResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetFederatedrealmsResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetFederatedrealmsResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetId

`func (o *GetFederatedrealmsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetFederatedrealmsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetFederatedrealmsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetFederatedrealmsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetFederatedrealmsResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetFederatedrealmsResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetFederatedrealmsResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetFederatedrealmsResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetFederatedrealmsResponse) GetResult() Federatedrealms`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetFederatedrealmsResponse) GetResultOk() (*Federatedrealms, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetFederatedrealmsResponse) SetResult(v Federatedrealms)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetFederatedrealmsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


