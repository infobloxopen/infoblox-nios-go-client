# MemberLicense

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
**Uuid** | Pointer to **string** | Universally Unique ID assigned for this object | [optional] [readonly] 

## Methods

### NewMemberLicense

`func NewMemberLicense() *MemberLicense`

NewMemberLicense instantiates a new MemberLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberLicenseWithDefaults

`func NewMemberLicenseWithDefaults() *MemberLicense`

NewMemberLicenseWithDefaults instantiates a new MemberLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *MemberLicense) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *MemberLicense) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *MemberLicense) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *MemberLicense) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetExpirationStatus

`func (o *MemberLicense) GetExpirationStatus() string`

GetExpirationStatus returns the ExpirationStatus field if non-nil, zero value otherwise.

### GetExpirationStatusOk

`func (o *MemberLicense) GetExpirationStatusOk() (*string, bool)`

GetExpirationStatusOk returns a tuple with the ExpirationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationStatus

`func (o *MemberLicense) SetExpirationStatus(v string)`

SetExpirationStatus sets ExpirationStatus field to given value.

### HasExpirationStatus

`func (o *MemberLicense) HasExpirationStatus() bool`

HasExpirationStatus returns a boolean if a field has been set.

### GetExpiryDate

`func (o *MemberLicense) GetExpiryDate() int64`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *MemberLicense) GetExpiryDateOk() (*int64, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *MemberLicense) SetExpiryDate(v int64)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *MemberLicense) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetHwid

`func (o *MemberLicense) GetHwid() string`

GetHwid returns the Hwid field if non-nil, zero value otherwise.

### GetHwidOk

`func (o *MemberLicense) GetHwidOk() (*string, bool)`

GetHwidOk returns a tuple with the Hwid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwid

`func (o *MemberLicense) SetHwid(v string)`

SetHwid sets Hwid field to given value.

### HasHwid

`func (o *MemberLicense) HasHwid() bool`

HasHwid returns a boolean if a field has been set.

### GetKey

`func (o *MemberLicense) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MemberLicense) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MemberLicense) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MemberLicense) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetKind

`func (o *MemberLicense) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *MemberLicense) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *MemberLicense) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *MemberLicense) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLimit

`func (o *MemberLicense) GetLimit() string`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MemberLicense) GetLimitOk() (*string, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MemberLicense) SetLimit(v string)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MemberLicense) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetLimitContext

`func (o *MemberLicense) GetLimitContext() string`

GetLimitContext returns the LimitContext field if non-nil, zero value otherwise.

### GetLimitContextOk

`func (o *MemberLicense) GetLimitContextOk() (*string, bool)`

GetLimitContextOk returns a tuple with the LimitContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitContext

`func (o *MemberLicense) SetLimitContext(v string)`

SetLimitContext sets LimitContext field to given value.

### HasLimitContext

`func (o *MemberLicense) HasLimitContext() bool`

HasLimitContext returns a boolean if a field has been set.

### GetType

`func (o *MemberLicense) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MemberLicense) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MemberLicense) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MemberLicense) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUuid

`func (o *MemberLicense) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MemberLicense) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MemberLicense) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MemberLicense) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


