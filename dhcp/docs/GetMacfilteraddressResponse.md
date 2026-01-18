# GetMacfilteraddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**AuthenticationTime** | Pointer to **int64** | The absolute UNIX time (in seconds) since the address was last authenticated. | [optional] 
**Comment** | Pointer to **string** | Comment for the MAC filter address; maximum 256 characters. | [optional] 
**ExpirationTime** | Pointer to **int64** | The absolute UNIX time (in seconds) until the address expires. | [optional] 
**ExtAttrsPlus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrsMinus** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**ExtAttrs** | Pointer to [**map[string]ExtAttrs**](ExtAttrs.md) | Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}. | [optional] 
**Filter** | Pointer to **string** | Name of the MAC filter to which this address belongs. | [optional] 
**Fingerprint** | Pointer to **string** | DHCP fingerprint for the address. | [optional] [readonly] 
**GuestCustomField1** | Pointer to **string** | Guest custom field 1. | [optional] 
**GuestCustomField2** | Pointer to **string** | Guest custom field 2. | [optional] 
**GuestCustomField3** | Pointer to **string** | Guest custom field 3. | [optional] 
**GuestCustomField4** | Pointer to **string** | Guest custom field 4. | [optional] 
**GuestEmail** | Pointer to **string** | Guest e-mail. | [optional] 
**GuestFirstName** | Pointer to **string** | Guest first name. | [optional] 
**GuestLastName** | Pointer to **string** | Guest last name. | [optional] 
**GuestMiddleName** | Pointer to **string** | Guest middle name. | [optional] 
**GuestPhone** | Pointer to **string** | Guest phone number. | [optional] 
**IsRegisteredUser** | Pointer to **bool** | Determines if the user has been authenticated or not. | [optional] [readonly] 
**Mac** | Pointer to **string** | MAC Address. | [optional] 
**NeverExpires** | Pointer to **bool** | Determines if MAC address expiration is enabled or disabled. | [optional] 
**ReservedForInfoblox** | Pointer to **string** | Reserved for future use. | [optional] 
**Username** | Pointer to **string** | Username for authenticated DHCP purposes. | [optional] 
**Result** | Pointer to [**Macfilteraddress**](Macfilteraddress.md) |  | [optional] 

## Methods

### NewGetMacfilteraddressResponse

`func NewGetMacfilteraddressResponse() *GetMacfilteraddressResponse`

NewGetMacfilteraddressResponse instantiates a new GetMacfilteraddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMacfilteraddressResponseWithDefaults

`func NewGetMacfilteraddressResponseWithDefaults() *GetMacfilteraddressResponse`

NewGetMacfilteraddressResponseWithDefaults instantiates a new GetMacfilteraddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetMacfilteraddressResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetMacfilteraddressResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetMacfilteraddressResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetMacfilteraddressResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthenticationTime

`func (o *GetMacfilteraddressResponse) GetAuthenticationTime() int64`

GetAuthenticationTime returns the AuthenticationTime field if non-nil, zero value otherwise.

### GetAuthenticationTimeOk

`func (o *GetMacfilteraddressResponse) GetAuthenticationTimeOk() (*int64, bool)`

GetAuthenticationTimeOk returns a tuple with the AuthenticationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationTime

`func (o *GetMacfilteraddressResponse) SetAuthenticationTime(v int64)`

SetAuthenticationTime sets AuthenticationTime field to given value.

### HasAuthenticationTime

`func (o *GetMacfilteraddressResponse) HasAuthenticationTime() bool`

HasAuthenticationTime returns a boolean if a field has been set.

### GetComment

`func (o *GetMacfilteraddressResponse) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetMacfilteraddressResponse) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetMacfilteraddressResponse) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetMacfilteraddressResponse) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpirationTime

`func (o *GetMacfilteraddressResponse) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *GetMacfilteraddressResponse) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *GetMacfilteraddressResponse) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *GetMacfilteraddressResponse) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *GetMacfilteraddressResponse) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *GetMacfilteraddressResponse) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *GetMacfilteraddressResponse) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *GetMacfilteraddressResponse) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *GetMacfilteraddressResponse) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *GetMacfilteraddressResponse) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *GetMacfilteraddressResponse) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *GetMacfilteraddressResponse) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *GetMacfilteraddressResponse) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *GetMacfilteraddressResponse) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *GetMacfilteraddressResponse) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *GetMacfilteraddressResponse) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFilter

`func (o *GetMacfilteraddressResponse) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *GetMacfilteraddressResponse) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *GetMacfilteraddressResponse) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *GetMacfilteraddressResponse) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFingerprint

`func (o *GetMacfilteraddressResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *GetMacfilteraddressResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *GetMacfilteraddressResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *GetMacfilteraddressResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetGuestCustomField1

`func (o *GetMacfilteraddressResponse) GetGuestCustomField1() string`

GetGuestCustomField1 returns the GuestCustomField1 field if non-nil, zero value otherwise.

### GetGuestCustomField1Ok

`func (o *GetMacfilteraddressResponse) GetGuestCustomField1Ok() (*string, bool)`

GetGuestCustomField1Ok returns a tuple with the GuestCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField1

`func (o *GetMacfilteraddressResponse) SetGuestCustomField1(v string)`

SetGuestCustomField1 sets GuestCustomField1 field to given value.

### HasGuestCustomField1

`func (o *GetMacfilteraddressResponse) HasGuestCustomField1() bool`

HasGuestCustomField1 returns a boolean if a field has been set.

### GetGuestCustomField2

`func (o *GetMacfilteraddressResponse) GetGuestCustomField2() string`

GetGuestCustomField2 returns the GuestCustomField2 field if non-nil, zero value otherwise.

### GetGuestCustomField2Ok

`func (o *GetMacfilteraddressResponse) GetGuestCustomField2Ok() (*string, bool)`

GetGuestCustomField2Ok returns a tuple with the GuestCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField2

`func (o *GetMacfilteraddressResponse) SetGuestCustomField2(v string)`

SetGuestCustomField2 sets GuestCustomField2 field to given value.

### HasGuestCustomField2

`func (o *GetMacfilteraddressResponse) HasGuestCustomField2() bool`

HasGuestCustomField2 returns a boolean if a field has been set.

### GetGuestCustomField3

`func (o *GetMacfilteraddressResponse) GetGuestCustomField3() string`

GetGuestCustomField3 returns the GuestCustomField3 field if non-nil, zero value otherwise.

### GetGuestCustomField3Ok

`func (o *GetMacfilteraddressResponse) GetGuestCustomField3Ok() (*string, bool)`

GetGuestCustomField3Ok returns a tuple with the GuestCustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField3

`func (o *GetMacfilteraddressResponse) SetGuestCustomField3(v string)`

SetGuestCustomField3 sets GuestCustomField3 field to given value.

### HasGuestCustomField3

`func (o *GetMacfilteraddressResponse) HasGuestCustomField3() bool`

HasGuestCustomField3 returns a boolean if a field has been set.

### GetGuestCustomField4

`func (o *GetMacfilteraddressResponse) GetGuestCustomField4() string`

GetGuestCustomField4 returns the GuestCustomField4 field if non-nil, zero value otherwise.

### GetGuestCustomField4Ok

`func (o *GetMacfilteraddressResponse) GetGuestCustomField4Ok() (*string, bool)`

GetGuestCustomField4Ok returns a tuple with the GuestCustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField4

`func (o *GetMacfilteraddressResponse) SetGuestCustomField4(v string)`

SetGuestCustomField4 sets GuestCustomField4 field to given value.

### HasGuestCustomField4

`func (o *GetMacfilteraddressResponse) HasGuestCustomField4() bool`

HasGuestCustomField4 returns a boolean if a field has been set.

### GetGuestEmail

`func (o *GetMacfilteraddressResponse) GetGuestEmail() string`

GetGuestEmail returns the GuestEmail field if non-nil, zero value otherwise.

### GetGuestEmailOk

`func (o *GetMacfilteraddressResponse) GetGuestEmailOk() (*string, bool)`

GetGuestEmailOk returns a tuple with the GuestEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestEmail

`func (o *GetMacfilteraddressResponse) SetGuestEmail(v string)`

SetGuestEmail sets GuestEmail field to given value.

### HasGuestEmail

`func (o *GetMacfilteraddressResponse) HasGuestEmail() bool`

HasGuestEmail returns a boolean if a field has been set.

### GetGuestFirstName

`func (o *GetMacfilteraddressResponse) GetGuestFirstName() string`

GetGuestFirstName returns the GuestFirstName field if non-nil, zero value otherwise.

### GetGuestFirstNameOk

`func (o *GetMacfilteraddressResponse) GetGuestFirstNameOk() (*string, bool)`

GetGuestFirstNameOk returns a tuple with the GuestFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestFirstName

`func (o *GetMacfilteraddressResponse) SetGuestFirstName(v string)`

SetGuestFirstName sets GuestFirstName field to given value.

### HasGuestFirstName

`func (o *GetMacfilteraddressResponse) HasGuestFirstName() bool`

HasGuestFirstName returns a boolean if a field has been set.

### GetGuestLastName

`func (o *GetMacfilteraddressResponse) GetGuestLastName() string`

GetGuestLastName returns the GuestLastName field if non-nil, zero value otherwise.

### GetGuestLastNameOk

`func (o *GetMacfilteraddressResponse) GetGuestLastNameOk() (*string, bool)`

GetGuestLastNameOk returns a tuple with the GuestLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestLastName

`func (o *GetMacfilteraddressResponse) SetGuestLastName(v string)`

SetGuestLastName sets GuestLastName field to given value.

### HasGuestLastName

`func (o *GetMacfilteraddressResponse) HasGuestLastName() bool`

HasGuestLastName returns a boolean if a field has been set.

### GetGuestMiddleName

`func (o *GetMacfilteraddressResponse) GetGuestMiddleName() string`

GetGuestMiddleName returns the GuestMiddleName field if non-nil, zero value otherwise.

### GetGuestMiddleNameOk

`func (o *GetMacfilteraddressResponse) GetGuestMiddleNameOk() (*string, bool)`

GetGuestMiddleNameOk returns a tuple with the GuestMiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestMiddleName

`func (o *GetMacfilteraddressResponse) SetGuestMiddleName(v string)`

SetGuestMiddleName sets GuestMiddleName field to given value.

### HasGuestMiddleName

`func (o *GetMacfilteraddressResponse) HasGuestMiddleName() bool`

HasGuestMiddleName returns a boolean if a field has been set.

### GetGuestPhone

`func (o *GetMacfilteraddressResponse) GetGuestPhone() string`

GetGuestPhone returns the GuestPhone field if non-nil, zero value otherwise.

### GetGuestPhoneOk

`func (o *GetMacfilteraddressResponse) GetGuestPhoneOk() (*string, bool)`

GetGuestPhoneOk returns a tuple with the GuestPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPhone

`func (o *GetMacfilteraddressResponse) SetGuestPhone(v string)`

SetGuestPhone sets GuestPhone field to given value.

### HasGuestPhone

`func (o *GetMacfilteraddressResponse) HasGuestPhone() bool`

HasGuestPhone returns a boolean if a field has been set.

### GetIsRegisteredUser

`func (o *GetMacfilteraddressResponse) GetIsRegisteredUser() bool`

GetIsRegisteredUser returns the IsRegisteredUser field if non-nil, zero value otherwise.

### GetIsRegisteredUserOk

`func (o *GetMacfilteraddressResponse) GetIsRegisteredUserOk() (*bool, bool)`

GetIsRegisteredUserOk returns a tuple with the IsRegisteredUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegisteredUser

`func (o *GetMacfilteraddressResponse) SetIsRegisteredUser(v bool)`

SetIsRegisteredUser sets IsRegisteredUser field to given value.

### HasIsRegisteredUser

`func (o *GetMacfilteraddressResponse) HasIsRegisteredUser() bool`

HasIsRegisteredUser returns a boolean if a field has been set.

### GetMac

`func (o *GetMacfilteraddressResponse) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GetMacfilteraddressResponse) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GetMacfilteraddressResponse) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GetMacfilteraddressResponse) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNeverExpires

`func (o *GetMacfilteraddressResponse) GetNeverExpires() bool`

GetNeverExpires returns the NeverExpires field if non-nil, zero value otherwise.

### GetNeverExpiresOk

`func (o *GetMacfilteraddressResponse) GetNeverExpiresOk() (*bool, bool)`

GetNeverExpiresOk returns a tuple with the NeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverExpires

`func (o *GetMacfilteraddressResponse) SetNeverExpires(v bool)`

SetNeverExpires sets NeverExpires field to given value.

### HasNeverExpires

`func (o *GetMacfilteraddressResponse) HasNeverExpires() bool`

HasNeverExpires returns a boolean if a field has been set.

### GetReservedForInfoblox

`func (o *GetMacfilteraddressResponse) GetReservedForInfoblox() string`

GetReservedForInfoblox returns the ReservedForInfoblox field if non-nil, zero value otherwise.

### GetReservedForInfobloxOk

`func (o *GetMacfilteraddressResponse) GetReservedForInfobloxOk() (*string, bool)`

GetReservedForInfobloxOk returns a tuple with the ReservedForInfoblox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedForInfoblox

`func (o *GetMacfilteraddressResponse) SetReservedForInfoblox(v string)`

SetReservedForInfoblox sets ReservedForInfoblox field to given value.

### HasReservedForInfoblox

`func (o *GetMacfilteraddressResponse) HasReservedForInfoblox() bool`

HasReservedForInfoblox returns a boolean if a field has been set.

### GetUsername

`func (o *GetMacfilteraddressResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetMacfilteraddressResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetMacfilteraddressResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetMacfilteraddressResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetResult

`func (o *GetMacfilteraddressResponse) GetResult() Macfilteraddress`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMacfilteraddressResponse) GetResultOk() (*Macfilteraddress, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMacfilteraddressResponse) SetResult(v Macfilteraddress)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetMacfilteraddressResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


