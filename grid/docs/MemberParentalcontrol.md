# MemberParentalcontrol

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**EnableService** | Pointer to **bool** | Determines if the parental control service is enabled. | [optional] 
**Name** | Pointer to **string** | The parental control member hostname. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewMemberParentalcontrol

`func NewMemberParentalcontrol() *MemberParentalcontrol`

NewMemberParentalcontrol instantiates a new MemberParentalcontrol object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberParentalcontrolWithDefaults

`func NewMemberParentalcontrolWithDefaults() *MemberParentalcontrol`

NewMemberParentalcontrolWithDefaults instantiates a new MemberParentalcontrol object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MemberParentalcontrol) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MemberParentalcontrol) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MemberParentalcontrol) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MemberParentalcontrol) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetEnableService

`func (o *MemberParentalcontrol) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *MemberParentalcontrol) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *MemberParentalcontrol) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *MemberParentalcontrol) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetName

`func (o *MemberParentalcontrol) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MemberParentalcontrol) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MemberParentalcontrol) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MemberParentalcontrol) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *MemberParentalcontrol) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MemberParentalcontrol) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MemberParentalcontrol) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MemberParentalcontrol) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


