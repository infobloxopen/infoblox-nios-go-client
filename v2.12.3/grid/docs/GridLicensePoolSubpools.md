# GridLicensePoolSubpools

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The license string for the license subpool. | [optional] [readonly] 
**Installed** | Pointer to **int64** | The total number of dynamic licenses allowed for this license subpool. | [optional] [readonly] 
**ExpiryDate** | Pointer to **int64** | License expiration date. | [optional] [readonly] 

## Methods

### NewGridLicensePoolSubpools

`func NewGridLicensePoolSubpools() *GridLicensePoolSubpools`

NewGridLicensePoolSubpools instantiates a new GridLicensePoolSubpools object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridLicensePoolSubpoolsWithDefaults

`func NewGridLicensePoolSubpoolsWithDefaults() *GridLicensePoolSubpools`

NewGridLicensePoolSubpoolsWithDefaults instantiates a new GridLicensePoolSubpools object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GridLicensePoolSubpools) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GridLicensePoolSubpools) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GridLicensePoolSubpools) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GridLicensePoolSubpools) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetInstalled

`func (o *GridLicensePoolSubpools) GetInstalled() int64`

GetInstalled returns the Installed field if non-nil, zero value otherwise.

### GetInstalledOk

`func (o *GridLicensePoolSubpools) GetInstalledOk() (*int64, bool)`

GetInstalledOk returns a tuple with the Installed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalled

`func (o *GridLicensePoolSubpools) SetInstalled(v int64)`

SetInstalled sets Installed field to given value.

### HasInstalled

`func (o *GridLicensePoolSubpools) HasInstalled() bool`

HasInstalled returns a boolean if a field has been set.

### GetExpiryDate

`func (o *GridLicensePoolSubpools) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *GridLicensePoolSubpools) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *GridLicensePoolSubpools) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *GridLicensePoolSubpools) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


