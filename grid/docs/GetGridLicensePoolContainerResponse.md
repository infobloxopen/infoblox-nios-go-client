# GetGridLicensePoolContainerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**LastEntitlementUpdate** | Pointer to **int64** | The timestamp when the last pool licenses were updated. | [optional] [readonly] 
**LpcUid** | Pointer to **string** | The world-wide unique ID for the license pool container. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**GridLicensePoolContainer**](GridLicensePoolContainer.md) |  | [optional] 

## Methods

### NewGetGridLicensePoolContainerResponse

`func NewGetGridLicensePoolContainerResponse() *GetGridLicensePoolContainerResponse`

NewGetGridLicensePoolContainerResponse instantiates a new GetGridLicensePoolContainerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridLicensePoolContainerResponseWithDefaults

`func NewGetGridLicensePoolContainerResponseWithDefaults() *GetGridLicensePoolContainerResponse`

NewGetGridLicensePoolContainerResponseWithDefaults instantiates a new GetGridLicensePoolContainerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridLicensePoolContainerResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridLicensePoolContainerResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridLicensePoolContainerResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridLicensePoolContainerResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetLastEntitlementUpdate

`func (o *GetGridLicensePoolContainerResponse) GetLastEntitlementUpdate() int64`

GetLastEntitlementUpdate returns the LastEntitlementUpdate field if non-nil, zero value otherwise.

### GetLastEntitlementUpdateOk

`func (o *GetGridLicensePoolContainerResponse) GetLastEntitlementUpdateOk() (*int64, bool)`

GetLastEntitlementUpdateOk returns a tuple with the LastEntitlementUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEntitlementUpdate

`func (o *GetGridLicensePoolContainerResponse) SetLastEntitlementUpdate(v int64)`

SetLastEntitlementUpdate sets LastEntitlementUpdate field to given value.

### HasLastEntitlementUpdate

`func (o *GetGridLicensePoolContainerResponse) HasLastEntitlementUpdate() bool`

HasLastEntitlementUpdate returns a boolean if a field has been set.

### GetLpcUid

`func (o *GetGridLicensePoolContainerResponse) GetLpcUid() string`

GetLpcUid returns the LpcUid field if non-nil, zero value otherwise.

### GetLpcUidOk

`func (o *GetGridLicensePoolContainerResponse) GetLpcUidOk() (*string, bool)`

GetLpcUidOk returns a tuple with the LpcUid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLpcUid

`func (o *GetGridLicensePoolContainerResponse) SetLpcUid(v string)`

SetLpcUid sets LpcUid field to given value.

### HasLpcUid

`func (o *GetGridLicensePoolContainerResponse) HasLpcUid() bool`

HasLpcUid returns a boolean if a field has been set.

### GetUuid

`func (o *GetGridLicensePoolContainerResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGridLicensePoolContainerResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGridLicensePoolContainerResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGridLicensePoolContainerResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetGridLicensePoolContainerResponse) GetResult() GridLicensePoolContainer`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridLicensePoolContainerResponse) GetResultOk() (*GridLicensePoolContainer, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridLicensePoolContainerResponse) SetResult(v GridLicensePoolContainer)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridLicensePoolContainerResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


