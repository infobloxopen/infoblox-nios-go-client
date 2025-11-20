# GetDatacollectionclusterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**EnableRegistration** | Pointer to **bool** | Enable/disable new registration requests | [optional] 
**Name** | Pointer to **string** | Display name for cluster | [optional] [readonly] 
**Result** | Pointer to [**Datacollectioncluster**](Datacollectioncluster.md) |  | [optional] 

## Methods

### NewGetDatacollectionclusterResponse

`func NewGetDatacollectionclusterResponse() *GetDatacollectionclusterResponse`

NewGetDatacollectionclusterResponse instantiates a new GetDatacollectionclusterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatacollectionclusterResponseWithDefaults

`func NewGetDatacollectionclusterResponseWithDefaults() *GetDatacollectionclusterResponse`

NewGetDatacollectionclusterResponseWithDefaults instantiates a new GetDatacollectionclusterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetDatacollectionclusterResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetDatacollectionclusterResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetDatacollectionclusterResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetDatacollectionclusterResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetEnableRegistration

`func (o *GetDatacollectionclusterResponse) GetEnableRegistration() bool`

GetEnableRegistration returns the EnableRegistration field if non-nil, zero value otherwise.

### GetEnableRegistrationOk

`func (o *GetDatacollectionclusterResponse) GetEnableRegistrationOk() (*bool, bool)`

GetEnableRegistrationOk returns a tuple with the EnableRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRegistration

`func (o *GetDatacollectionclusterResponse) SetEnableRegistration(v bool)`

SetEnableRegistration sets EnableRegistration field to given value.

### HasEnableRegistration

`func (o *GetDatacollectionclusterResponse) HasEnableRegistration() bool`

HasEnableRegistration returns a boolean if a field has been set.

### GetName

`func (o *GetDatacollectionclusterResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDatacollectionclusterResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDatacollectionclusterResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDatacollectionclusterResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetDatacollectionclusterResponse) GetResult() Datacollectioncluster`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetDatacollectionclusterResponse) GetResultOk() (*Datacollectioncluster, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetDatacollectionclusterResponse) SetResult(v Datacollectioncluster)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetDatacollectionclusterResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


