# GetAdAuthServiceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AdDomain** | Pointer to **string** | The Active Directory domain to which this server belongs. | [optional] 
**Comment** | Pointer to **string** | The descriptive comment for the AD authentication service. | [optional] 
**Disabled** | Pointer to **bool** | Determines if Active Directory Authentication Service is disabled. | [optional] 
**DomainControllers** | Pointer to [**[]AdAuthServiceDomainControllers**](AdAuthServiceDomainControllers.md) | The AD authentication server list. | [optional] 
**Name** | Pointer to **string** | The AD authentication service name. | [optional] 
**NestedGroupQuerying** | Pointer to **bool** | Determines whether the nested group querying is enabled. | [optional] 
**Timeout** | Pointer to **int64** | The number of seconds that the appliance waits for a response from the AD server. | [optional] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**AdAuthService**](AdAuthService.md) |  | [optional] 

## Methods

### NewGetAdAuthServiceResponse

`func NewGetAdAuthServiceResponse() *GetAdAuthServiceResponse`

NewGetAdAuthServiceResponse instantiates a new GetAdAuthServiceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdAuthServiceResponseWithDefaults

`func NewGetAdAuthServiceResponseWithDefaults() *GetAdAuthServiceResponse`

NewGetAdAuthServiceResponseWithDefaults instantiates a new GetAdAuthServiceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetAdAuthServiceResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetAdAuthServiceResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetAdAuthServiceResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetAdAuthServiceResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAdDomain

`func (o *GetAdAuthServiceResponse) GetAdDomain() string`

GetAdDomain returns the AdDomain field if non-nil, zero value otherwise.

### GetAdDomainOk

`func (o *GetAdAuthServiceResponse) GetAdDomainOk() (*string, bool)`

GetAdDomainOk returns a tuple with the AdDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomain

`func (o *GetAdAuthServiceResponse) SetAdDomain(v string)`

SetAdDomain sets AdDomain field to given value.

### HasAdDomain

`func (o *GetAdAuthServiceResponse) HasAdDomain() bool`

HasAdDomain returns a boolean if a field has been set.

### GetComment

`func (o *GetAdAuthServiceResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetAdAuthServiceResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetAdAuthServiceResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetAdAuthServiceResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *GetAdAuthServiceResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *GetAdAuthServiceResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *GetAdAuthServiceResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *GetAdAuthServiceResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDomainControllers

`func (o *GetAdAuthServiceResponse) GetDomainControllers() []AdAuthServiceDomainControllers`

GetDomainControllers returns the DomainControllers field if non-nil, zero value otherwise.

### GetDomainControllersOk

`func (o *GetAdAuthServiceResponse) GetDomainControllersOk() (*[]AdAuthServiceDomainControllers, bool)`

GetDomainControllersOk returns a tuple with the DomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllers

`func (o *GetAdAuthServiceResponse) SetDomainControllers(v []AdAuthServiceDomainControllers)`

SetDomainControllers sets DomainControllers field to given value.

### HasDomainControllers

`func (o *GetAdAuthServiceResponse) HasDomainControllers() bool`

HasDomainControllers returns a boolean if a field has been set.

### GetName

`func (o *GetAdAuthServiceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAdAuthServiceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAdAuthServiceResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAdAuthServiceResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNestedGroupQuerying

`func (o *GetAdAuthServiceResponse) GetNestedGroupQuerying() bool`

GetNestedGroupQuerying returns the NestedGroupQuerying field if non-nil, zero value otherwise.

### GetNestedGroupQueryingOk

`func (o *GetAdAuthServiceResponse) GetNestedGroupQueryingOk() (*bool, bool)`

GetNestedGroupQueryingOk returns a tuple with the NestedGroupQuerying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedGroupQuerying

`func (o *GetAdAuthServiceResponse) SetNestedGroupQuerying(v bool)`

SetNestedGroupQuerying sets NestedGroupQuerying field to given value.

### HasNestedGroupQuerying

`func (o *GetAdAuthServiceResponse) HasNestedGroupQuerying() bool`

HasNestedGroupQuerying returns a boolean if a field has been set.

### GetTimeout

`func (o *GetAdAuthServiceResponse) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetAdAuthServiceResponse) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetAdAuthServiceResponse) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetAdAuthServiceResponse) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUuid

`func (o *GetAdAuthServiceResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetAdAuthServiceResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetAdAuthServiceResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetAdAuthServiceResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetAdAuthServiceResponse) GetResult() AdAuthService`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetAdAuthServiceResponse) GetResultOk() (*AdAuthService, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetAdAuthServiceResponse) SetResult(v AdAuthService)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetAdAuthServiceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


