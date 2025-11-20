# GetMemberLicenseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**ExpirationStatus** | Pointer to **string** | The license expiration status. | [optional] [readonly] 
**ExpiryDate** | Pointer to **int64** | The expiration timestamp of the license. | [optional] [readonly] 
**Hwid** | Pointer to **string** | The hardware ID of the physical node on which the license is installed. | [optional] [readonly] 
**Key** | Pointer to **string** | License string. | [optional] [readonly] 
**Kind** | Pointer to **string** | The overall type of license: static or dynamic. | [optional] [readonly] 
**Limit** | Pointer to **string** | The license limit value. | [optional] [readonly] 
**LimitContext** | Pointer to **string** | The license limit context. | [optional] [readonly] 
**Type** | Pointer to **string** | The license type. | [optional] [readonly] 
**Result** | Pointer to [**MemberLicense**](MemberLicense.md) |  | [optional] 

## Methods

### NewGetMemberLicenseResponse

`func NewGetMemberLicenseResponse() *GetMemberLicenseResponse`

NewGetMemberLicenseResponse instantiates a new GetMemberLicenseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMemberLicenseResponseWithDefaults

`func NewGetMemberLicenseResponseWithDefaults() *GetMemberLicenseResponse`

NewGetMemberLicenseResponseWithDefaults instantiates a new GetMemberLicenseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMemberLicenseResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMemberLicenseResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMemberLicenseResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMemberLicenseResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetExpirationStatus

`func (o *GetMemberLicenseResponse) GetExpirationStatus() string`

GetExpirationStatus returns the ExpirationStatus field if non-nil, zero value otherwise.

### GetExpirationStatusOk

`func (o *GetMemberLicenseResponse) GetExpirationStatusOk() (*string, bool)`

GetExpirationStatusOk returns a tuple with the ExpirationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationStatus

`func (o *GetMemberLicenseResponse) SetExpirationStatus(v string)`

SetExpirationStatus sets ExpirationStatus field to given value.

### HasExpirationStatus

`func (o *GetMemberLicenseResponse) HasExpirationStatus() bool`

HasExpirationStatus returns a boolean if a field has been set.

### GetExpiryDate

`func (o *GetMemberLicenseResponse) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *GetMemberLicenseResponse) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *GetMemberLicenseResponse) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *GetMemberLicenseResponse) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetHwid

`func (o *GetMemberLicenseResponse) GetHwid() string`

GetHwid returns the Hwid field if non-nil, zero value otherwise.

### GetHwidOk

`func (o *GetMemberLicenseResponse) GetHwidOk() (*string, bool)`

GetHwidOk returns a tuple with the Hwid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwid

`func (o *GetMemberLicenseResponse) SetHwid(v string)`

SetHwid sets Hwid field to given value.

### HasHwid

`func (o *GetMemberLicenseResponse) HasHwid() bool`

HasHwid returns a boolean if a field has been set.

### GetKey

`func (o *GetMemberLicenseResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetMemberLicenseResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetMemberLicenseResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetMemberLicenseResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKind

`func (o *GetMemberLicenseResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetMemberLicenseResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetMemberLicenseResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GetMemberLicenseResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *GetMemberLicenseResponse) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetMemberLicenseResponse) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetMemberLicenseResponse) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetMemberLicenseResponse) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitContext

`func (o *GetMemberLicenseResponse) GetLimitContext() string`

GetLimitContext returns the LimitContext field if non-nil, zero value otherwise.

### GetLimitContextOk

`func (o *GetMemberLicenseResponse) GetLimitContextOk() (*string, bool)`

GetLimitContextOk returns a tuple with the LimitContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitContext

`func (o *GetMemberLicenseResponse) SetLimitContext(v string)`

SetLimitContext sets LimitContext field to given value.

### HasLimitContext

`func (o *GetMemberLicenseResponse) HasLimitContext() bool`

HasLimitContext returns a boolean if a field has been set.

### GetType

`func (o *GetMemberLicenseResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMemberLicenseResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMemberLicenseResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMemberLicenseResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResult

`func (o *GetMemberLicenseResponse) GetResult() MemberLicense`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMemberLicenseResponse) GetResultOk() (*MemberLicense, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMemberLicenseResponse) SetResult(v MemberLicense)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMemberLicenseResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


