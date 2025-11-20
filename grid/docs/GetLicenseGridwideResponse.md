# GetLicenseGridwideResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**ExpirationStatus** | Pointer to **string** | The license expiration status. | [optional] [readonly] 
**ExpiryDate** | Pointer to **int64** | The expiration timestamp of the license. | [optional] [readonly] 
**Key** | Pointer to **string** | The license string. | [optional] [readonly] 
**Limit** | Pointer to **string** | The license limit value. | [optional] [readonly] 
**LimitContext** | Pointer to **string** | The license limit context. | [optional] [readonly] 
**Type** | Pointer to **string** | The license type. | [optional] [readonly] 
**Result** | Pointer to [**LicenseGridwide**](LicenseGridwide.md) |  | [optional] 

## Methods

### NewGetLicenseGridwideResponse

`func NewGetLicenseGridwideResponse() *GetLicenseGridwideResponse`

NewGetLicenseGridwideResponse instantiates a new GetLicenseGridwideResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLicenseGridwideResponseWithDefaults

`func NewGetLicenseGridwideResponseWithDefaults() *GetLicenseGridwideResponse`

NewGetLicenseGridwideResponseWithDefaults instantiates a new GetLicenseGridwideResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetLicenseGridwideResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetLicenseGridwideResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetLicenseGridwideResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetLicenseGridwideResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetExpirationStatus

`func (o *GetLicenseGridwideResponse) GetExpirationStatus() string`

GetExpirationStatus returns the ExpirationStatus field if non-nil, zero value otherwise.

### GetExpirationStatusOk

`func (o *GetLicenseGridwideResponse) GetExpirationStatusOk() (*string, bool)`

GetExpirationStatusOk returns a tuple with the ExpirationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationStatus

`func (o *GetLicenseGridwideResponse) SetExpirationStatus(v string)`

SetExpirationStatus sets ExpirationStatus field to given value.

### HasExpirationStatus

`func (o *GetLicenseGridwideResponse) HasExpirationStatus() bool`

HasExpirationStatus returns a boolean if a field has been set.

### GetExpiryDate

`func (o *GetLicenseGridwideResponse) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *GetLicenseGridwideResponse) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *GetLicenseGridwideResponse) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *GetLicenseGridwideResponse) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetKey

`func (o *GetLicenseGridwideResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetLicenseGridwideResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetLicenseGridwideResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetLicenseGridwideResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLimit

`func (o *GetLicenseGridwideResponse) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetLicenseGridwideResponse) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetLicenseGridwideResponse) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetLicenseGridwideResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitContext

`func (o *GetLicenseGridwideResponse) GetLimitContext() string`

GetLimitContext returns the LimitContext field if non-nil, zero value otherwise.

### GetLimitContextOk

`func (o *GetLicenseGridwideResponse) GetLimitContextOk() (*string, bool)`

GetLimitContextOk returns a tuple with the LimitContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitContext

`func (o *GetLicenseGridwideResponse) SetLimitContext(v string)`

SetLimitContext sets LimitContext field to given value.

### HasLimitContext

`func (o *GetLicenseGridwideResponse) HasLimitContext() bool`

HasLimitContext returns a boolean if a field has been set.

### GetType

`func (o *GetLicenseGridwideResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetLicenseGridwideResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetLicenseGridwideResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetLicenseGridwideResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *GetLicenseGridwideResponse) GetResult() LicenseGridwide`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetLicenseGridwideResponse) GetResultOk() (*LicenseGridwide, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetLicenseGridwideResponse) SetResult(v LicenseGridwide)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetLicenseGridwideResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


