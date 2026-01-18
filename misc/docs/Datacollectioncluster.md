# Datacollectioncluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**EnableRegistration** | Pointer to **bool** | Enable/disable new registration requests | [optional] 
**Name** | Pointer to **string** | Display name for cluster | [optional] [readonly] 

## Methods

### NewDatacollectioncluster

`func NewDatacollectioncluster() *Datacollectioncluster`

NewDatacollectioncluster instantiates a new Datacollectioncluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacollectionclusterWithDefaults

`func NewDatacollectionclusterWithDefaults() *Datacollectioncluster`

NewDatacollectionclusterWithDefaults instantiates a new Datacollectioncluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Datacollectioncluster) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Datacollectioncluster) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Datacollectioncluster) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Datacollectioncluster) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetEnableRegistration

`func (o *Datacollectioncluster) GetEnableRegistration() bool`

GetEnableRegistration returns the EnableRegistration field if non-nil, zero value otherwise.

### GetEnableRegistrationOk

`func (o *Datacollectioncluster) GetEnableRegistrationOk() (*bool, bool)`

GetEnableRegistrationOk returns a tuple with the EnableRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRegistration

`func (o *Datacollectioncluster) SetEnableRegistration(v bool)`

SetEnableRegistration sets EnableRegistration field to given value.

### HasEnableRegistration

`func (o *Datacollectioncluster) HasEnableRegistration() bool`

HasEnableRegistration returns a boolean if a field has been set.

### GetName

`func (o *Datacollectioncluster) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Datacollectioncluster) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Datacollectioncluster) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Datacollectioncluster) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


