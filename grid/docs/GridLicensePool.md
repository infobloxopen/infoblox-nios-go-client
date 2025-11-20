# GridLicensePool

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

## Methods

### NewGridLicensePool

`func NewGridLicensePool() *GridLicensePool`

NewGridLicensePool instantiates a new GridLicensePool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridLicensePoolWithDefaults

`func NewGridLicensePoolWithDefaults() *GridLicensePool`

NewGridLicensePoolWithDefaults instantiates a new GridLicensePool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GridLicensePool) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GridLicensePool) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GridLicensePool) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GridLicensePool) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAssigned

`func (o *GridLicensePool) GetAssigned() int64`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *GridLicensePool) GetAssignedOk() (*int64, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *GridLicensePool) SetAssigned(v int64)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *GridLicensePool) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetExpirationStatus

`func (o *GridLicensePool) GetExpirationStatus() string`

GetExpirationStatus returns the ExpirationStatus field if non-nil, zero value otherwise.

### GetExpirationStatusOk

`func (o *GridLicensePool) GetExpirationStatusOk() (*string, bool)`

GetExpirationStatusOk returns a tuple with the ExpirationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationStatus

`func (o *GridLicensePool) SetExpirationStatus(v string)`

SetExpirationStatus sets ExpirationStatus field to given value.

### HasExpirationStatus

`func (o *GridLicensePool) HasExpirationStatus() bool`

HasExpirationStatus returns a boolean if a field has been set.

### GetExpiryDate

`func (o *GridLicensePool) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *GridLicensePool) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *GridLicensePool) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *GridLicensePool) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetInstalled

`func (o *GridLicensePool) GetInstalled() int64`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *GridLicensePool) GetInstalledOk() (*int64, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *GridLicensePool) SetInstalled(v int64)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *GridLicensePool) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetKey

`func (o *GridLicensePool) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GridLicensePool) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GridLicensePool) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GridLicensePool) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLimit

`func (o *GridLicensePool) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GridLicensePool) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GridLicensePool) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GridLicensePool) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitContext

`func (o *GridLicensePool) GetLimitContext() string`

GetLimitContext returns the LimitContext field if non-nil, zero value otherwise.

### GetLimitContextOk

`func (o *GridLicensePool) GetLimitContextOk() (*string, bool)`

GetLimitContextOk returns a tuple with the LimitContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitContext

`func (o *GridLicensePool) SetLimitContext(v string)`

SetLimitContext sets LimitContext field to given value.

### HasLimitContext

`func (o *GridLicensePool) HasLimitContext() bool`

HasLimitContext returns a boolean if a field has been set.

### GetModel

`func (o *GridLicensePool) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GridLicensePool) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GridLicensePool) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GridLicensePool) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetSubpools

`func (o *GridLicensePool) GetSubpools() []GridLicensePoolSubpools`

GetSubpools returns the Subpools field if non-nil, zero value otherwise.

### GetSubpoolsOk

`func (o *GridLicensePool) GetSubpoolsOk() (*[]GridLicensePoolSubpools, bool)`

GetSubpoolsOk returns a tuple with the Subpools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubpools

`func (o *GridLicensePool) SetSubpools(v []GridLicensePoolSubpools)`

SetSubpools sets Subpools field to given value.

### HasSubpools

`func (o *GridLicensePool) HasSubpools() bool`

HasSubpools returns a boolean if a field has been set.

### GetTempAssigned

`func (o *GridLicensePool) GetTempAssigned() int64`

GetTempAssigned returns the TempAssigned field if non-nil, zero value otherwise.

### GetTempAssignedOk

`func (o *GridLicensePool) GetTempAssignedOk() (*int64, bool)`

GetTempAssignedOk returns a tuple with the TempAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempAssigned

`func (o *GridLicensePool) SetTempAssigned(v int64)`

SetTempAssigned sets TempAssigned field to given value.

### HasTempAssigned

`func (o *GridLicensePool) HasTempAssigned() bool`

HasTempAssigned returns a boolean if a field has been set.

### GetType

`func (o *GridLicensePool) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GridLicensePool) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GridLicensePool) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GridLicensePool) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


