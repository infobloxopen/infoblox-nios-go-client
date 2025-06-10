# GridLicensePoolContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**LastEntitlementUpdate** | Pointer to **int64** | The timestamp when the last pool licenses were updated. | [optional] [readonly] 
**LpcUid** | Pointer to **string** | The world-wide unique ID for the license pool container. | [optional] [readonly] 

## Methods

### NewGridLicensePoolContainer

`func NewGridLicensePoolContainer() *GridLicensePoolContainer`

NewGridLicensePoolContainer instantiates a new GridLicensePoolContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridLicensePoolContainerWithDefaults

`func NewGridLicensePoolContainerWithDefaults() *GridLicensePoolContainer`

NewGridLicensePoolContainerWithDefaults instantiates a new GridLicensePoolContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridLicensePoolContainer) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridLicensePoolContainer) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridLicensePoolContainer) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridLicensePoolContainer) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetLastEntitlementUpdate

`func (o *GridLicensePoolContainer) GetLastEntitlementUpdate() int64`

GetLastEntitlementUpdate returns the LastEntitlementUpdate field if non-nil, zero value otherwise.

### GetLastEntitlementUpdateOk

`func (o *GridLicensePoolContainer) GetLastEntitlementUpdateOk() (*int64, bool)`

GetLastEntitlementUpdateOk returns a tuple with the LastEntitlementUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntitlementUpdate

`func (o *GridLicensePoolContainer) SetLastEntitlementUpdate(v int64)`

SetLastEntitlementUpdate sets LastEntitlementUpdate field to given value.

### HasLastEntitlementUpdate

`func (o *GridLicensePoolContainer) HasLastEntitlementUpdate() bool`

HasLastEntitlementUpdate returns a boolean if a field has been set.

### GetLpcUid

`func (o *GridLicensePoolContainer) GetLpcUid() string`

GetLpcUid returns the LpcUid field if non-nil, zero value otherwise.

### GetLpcUidOk

`func (o *GridLicensePoolContainer) GetLpcUidOk() (*string, bool)`

GetLpcUidOk returns a tuple with the LpcUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpcUid

`func (o *GridLicensePoolContainer) SetLpcUid(v string)`

SetLpcUid sets LpcUid field to given value.

### HasLpcUid

`func (o *GridLicensePoolContainer) HasLpcUid() bool`

HasLpcUid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


