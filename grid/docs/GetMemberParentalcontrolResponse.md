# GetMemberParentalcontrolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**EnableService** | Pointer to **bool** | Determines if the parental control service is enabled. | [optional] 
**Name** | Pointer to **string** | The parental control member hostname. | [optional] [readonly] 
**Result** | Pointer to [**MemberParentalcontrol**](MemberParentalcontrol.md) |  | [optional] 

## Methods

### NewGetMemberParentalcontrolResponse

`func NewGetMemberParentalcontrolResponse() *GetMemberParentalcontrolResponse`

NewGetMemberParentalcontrolResponse instantiates a new GetMemberParentalcontrolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMemberParentalcontrolResponseWithDefaults

`func NewGetMemberParentalcontrolResponseWithDefaults() *GetMemberParentalcontrolResponse`

NewGetMemberParentalcontrolResponseWithDefaults instantiates a new GetMemberParentalcontrolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMemberParentalcontrolResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMemberParentalcontrolResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMemberParentalcontrolResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMemberParentalcontrolResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetEnableService

`func (o *GetMemberParentalcontrolResponse) GetEnableService() bool`

GetEnableService returns the EnableService field if non-nil, zero value otherwise.

### GetEnableServiceOk

`func (o *GetMemberParentalcontrolResponse) GetEnableServiceOk() (*bool, bool)`

GetEnableServiceOk returns a tuple with the EnableService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableService

`func (o *GetMemberParentalcontrolResponse) SetEnableService(v bool)`

SetEnableService sets EnableService field to given value.

### HasEnableService

`func (o *GetMemberParentalcontrolResponse) HasEnableService() bool`

HasEnableService returns a boolean if a field has been set.

### GetName

`func (o *GetMemberParentalcontrolResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMemberParentalcontrolResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMemberParentalcontrolResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetMemberParentalcontrolResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResult

`func (o *GetMemberParentalcontrolResponse) GetResult() MemberParentalcontrol`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMemberParentalcontrolResponse) GetResultOk() (*MemberParentalcontrol, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMemberParentalcontrolResponse) SetResult(v MemberParentalcontrol)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMemberParentalcontrolResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


