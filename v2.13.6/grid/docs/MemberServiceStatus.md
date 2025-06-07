# MemberServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the current service status. | [optional] [readonly] 
**Status** | Pointer to **string** | The service status. | [optional] [readonly] 
**Service** | Pointer to **string** | The service identifier. | [optional] [readonly] 

## Methods

### NewMemberServiceStatus

`func NewMemberServiceStatus() *MemberServiceStatus`

NewMemberServiceStatus instantiates a new MemberServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberServiceStatusWithDefaults

`func NewMemberServiceStatusWithDefaults() *MemberServiceStatus`

NewMemberServiceStatusWithDefaults instantiates a new MemberServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MemberServiceStatus) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MemberServiceStatus) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MemberServiceStatus) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MemberServiceStatus) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *MemberServiceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MemberServiceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MemberServiceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MemberServiceStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetService

`func (o *MemberServiceStatus) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *MemberServiceStatus) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *MemberServiceStatus) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *MemberServiceStatus) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


