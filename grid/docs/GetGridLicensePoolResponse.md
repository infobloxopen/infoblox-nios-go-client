# GetGridLicensePoolResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**Assigned** | Pointer to **int64** | The number of dynamic licenses allocated to vNIOS appliances. | [optional] [readonly] 
**ExpirationStatus** | Pointer to **string** | The license expiration status. * DELETED - The temporary license has been deleted. * EXPIRED - License/Pool has reached the expiry date. * EXPIRING_SOON - License/Pool expires in 31-90 days. * EXPIRING_VERY_SOON - License/Pool expires in 30 days or earlier. * NOT_EXPIRED - License/Pool has not expired. * PERMANENT - License/Pool does not expire. | [optional] [readonly] 
**ExpiryDate** | Pointer to **int64** | The expiration timestamp of the license. | [optional] [readonly] 
**Installed** | Pointer to **int64** | The total number of dynamic licenses allowed for this license pool. | [optional] [readonly] 
**Key** | Pointer to **string** | The license string for the license pool. | [optional] [readonly] 
**Limit** | Pointer to **string** | The limitation of dynamic license that can be allocated from the license pool. | [optional] [readonly] 
**LimitContext** | Pointer to **string** | The license limit context. | [optional] [readonly] 
**Model** | Pointer to **string** | The supported vNIOS virtual appliance model. | [optional] [readonly] 
**Subpools** | Pointer to [**[]GridLicensePoolSubpools**](GridLicensePoolSubpools.md) | The license pool subpools. | [optional] [readonly] 
**TempAssigned** | Pointer to **int64** | The total number of temporary dynamic licenses allocated to vNIOS appliances. | [optional] [readonly] 
**Type** | Pointer to **string** | The license type. | [optional] [readonly] 
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 
**Result** | Pointer to [**GridLicensePool**](GridLicensePool.md) |  | [optional] 

## Methods

### NewGetGridLicensePoolResponse

`func NewGetGridLicensePoolResponse() *GetGridLicensePoolResponse`

NewGetGridLicensePoolResponse instantiates a new GetGridLicensePoolResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGridLicensePoolResponseWithDefaults

`func NewGetGridLicensePoolResponseWithDefaults() *GetGridLicensePoolResponse`

NewGetGridLicensePoolResponseWithDefaults instantiates a new GetGridLicensePoolResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetGridLicensePoolResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetGridLicensePoolResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetGridLicensePoolResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetGridLicensePoolResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAssigned

`func (o *GetGridLicensePoolResponse) GetAssigned() int64`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *GetGridLicensePoolResponse) GetAssignedOk() (*int64, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *GetGridLicensePoolResponse) SetAssigned(v int64)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *GetGridLicensePoolResponse) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetExpirationStatus

`func (o *GetGridLicensePoolResponse) GetExpirationStatus() string`

GetExpirationStatus returns the ExpirationStatus field if non-nil, zero value otherwise.

### GetExpirationStatusOk

`func (o *GetGridLicensePoolResponse) GetExpirationStatusOk() (*string, bool)`

GetExpirationStatusOk returns a tuple with the ExpirationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationStatus

`func (o *GetGridLicensePoolResponse) SetExpirationStatus(v string)`

SetExpirationStatus sets ExpirationStatus field to given value.

### HasExpirationStatus

`func (o *GetGridLicensePoolResponse) HasExpirationStatus() bool`

HasExpirationStatus returns a boolean if a field has been set.

### GetExpiryDate

`func (o *GetGridLicensePoolResponse) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *GetGridLicensePoolResponse) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *GetGridLicensePoolResponse) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *GetGridLicensePoolResponse) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetInstalled

`func (o *GetGridLicensePoolResponse) GetInstalled() int64`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *GetGridLicensePoolResponse) GetInstalledOk() (*int64, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *GetGridLicensePoolResponse) SetInstalled(v int64)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *GetGridLicensePoolResponse) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetKey

`func (o *GetGridLicensePoolResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetGridLicensePoolResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetGridLicensePoolResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetGridLicensePoolResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLimit

`func (o *GetGridLicensePoolResponse) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetGridLicensePoolResponse) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetGridLicensePoolResponse) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetGridLicensePoolResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitContext

`func (o *GetGridLicensePoolResponse) GetLimitContext() string`

GetLimitContext returns the LimitContext field if non-nil, zero value otherwise.

### GetLimitContextOk

`func (o *GetGridLicensePoolResponse) GetLimitContextOk() (*string, bool)`

GetLimitContextOk returns a tuple with the LimitContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitContext

`func (o *GetGridLicensePoolResponse) SetLimitContext(v string)`

SetLimitContext sets LimitContext field to given value.

### HasLimitContext

`func (o *GetGridLicensePoolResponse) HasLimitContext() bool`

HasLimitContext returns a boolean if a field has been set.

### GetModel

`func (o *GetGridLicensePoolResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetGridLicensePoolResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetGridLicensePoolResponse) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetGridLicensePoolResponse) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSubpools

`func (o *GetGridLicensePoolResponse) GetSubpools() []GridLicensePoolSubpools`

GetSubpools returns the Subpools field if non-nil, zero value otherwise.

### GetSubpoolsOk

`func (o *GetGridLicensePoolResponse) GetSubpoolsOk() (*[]GridLicensePoolSubpools, bool)`

GetSubpoolsOk returns a tuple with the Subpools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubpools

`func (o *GetGridLicensePoolResponse) SetSubpools(v []GridLicensePoolSubpools)`

SetSubpools sets Subpools field to given value.

### HasSubpools

`func (o *GetGridLicensePoolResponse) HasSubpools() bool`

HasSubpools returns a boolean if a field has been set.

### GetTempAssigned

`func (o *GetGridLicensePoolResponse) GetTempAssigned() int64`

GetTempAssigned returns the TempAssigned field if non-nil, zero value otherwise.

### GetTempAssignedOk

`func (o *GetGridLicensePoolResponse) GetTempAssignedOk() (*int64, bool)`

GetTempAssignedOk returns a tuple with the TempAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempAssigned

`func (o *GetGridLicensePoolResponse) SetTempAssigned(v int64)`

SetTempAssigned sets TempAssigned field to given value.

### HasTempAssigned

`func (o *GetGridLicensePoolResponse) HasTempAssigned() bool`

HasTempAssigned returns a boolean if a field has been set.

### GetType

`func (o *GetGridLicensePoolResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGridLicensePoolResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGridLicensePoolResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetGridLicensePoolResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *GetGridLicensePoolResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetGridLicensePoolResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetGridLicensePoolResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetGridLicensePoolResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetResult

`func (o *GetGridLicensePoolResponse) GetResult() GridLicensePool`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetGridLicensePoolResponse) GetResultOk() (*GridLicensePool, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetGridLicensePoolResponse) SetResult(v GridLicensePool)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetGridLicensePoolResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


