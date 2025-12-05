# AdAuthService

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

## Methods

### NewAdAuthService

`func NewAdAuthService() *AdAuthService`

NewAdAuthService instantiates a new AdAuthService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdAuthServiceWithDefaults

`func NewAdAuthServiceWithDefaults() *AdAuthService`

NewAdAuthServiceWithDefaults instantiates a new AdAuthService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *AdAuthService) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *AdAuthService) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *AdAuthService) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *AdAuthService) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAdDomain

`func (o *AdAuthService) GetAdDomain() string`

GetAdDomain returns the AdDomain field if non-nil, zero value otherwise.

### GetAdDomainOk

`func (o *AdAuthService) GetAdDomainOk() (*string, bool)`

GetAdDomainOk returns a tuple with the AdDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomain

`func (o *AdAuthService) SetAdDomain(v string)`

SetAdDomain sets AdDomain field to given value.

### HasAdDomain

`func (o *AdAuthService) HasAdDomain() bool`

HasAdDomain returns a boolean if a field has been set.

### GetComment

`func (o *AdAuthService) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *AdAuthService) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *AdAuthService) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *AdAuthService) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDisabled

`func (o *AdAuthService) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AdAuthService) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AdAuthService) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *AdAuthService) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDomainControllers

`func (o *AdAuthService) GetDomainControllers() []AdAuthServiceDomainControllers`

GetDomainControllers returns the DomainControllers field if non-nil, zero value otherwise.

### GetDomainControllersOk

`func (o *AdAuthService) GetDomainControllersOk() (*[]AdAuthServiceDomainControllers, bool)`

GetDomainControllersOk returns a tuple with the DomainControllers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainControllers

`func (o *AdAuthService) SetDomainControllers(v []AdAuthServiceDomainControllers)`

SetDomainControllers sets DomainControllers field to given value.

### HasDomainControllers

`func (o *AdAuthService) HasDomainControllers() bool`

HasDomainControllers returns a boolean if a field has been set.

### GetName

`func (o *AdAuthService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AdAuthService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AdAuthService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AdAuthService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNestedGroupQuerying

`func (o *AdAuthService) GetNestedGroupQuerying() bool`

GetNestedGroupQuerying returns the NestedGroupQuerying field if non-nil, zero value otherwise.

### GetNestedGroupQueryingOk

`func (o *AdAuthService) GetNestedGroupQueryingOk() (*bool, bool)`

GetNestedGroupQueryingOk returns a tuple with the NestedGroupQuerying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedGroupQuerying

`func (o *AdAuthService) SetNestedGroupQuerying(v bool)`

SetNestedGroupQuerying sets NestedGroupQuerying field to given value.

### HasNestedGroupQuerying

`func (o *AdAuthService) HasNestedGroupQuerying() bool`

HasNestedGroupQuerying returns a boolean if a field has been set.

### GetTimeout

`func (o *AdAuthService) GetTimeout() int64`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *AdAuthService) GetTimeoutOk() (*int64, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *AdAuthService) SetTimeout(v int64)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *AdAuthService) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUuid

`func (o *AdAuthService) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AdAuthService) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AdAuthService) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AdAuthService) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


