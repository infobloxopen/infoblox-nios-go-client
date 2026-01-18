# Macfilteraddress

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

## Methods

### NewMacfilteraddress

`func NewMacfilteraddress() *Macfilteraddress`

NewMacfilteraddress instantiates a new Macfilteraddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacfilteraddressWithDefaults

`func NewMacfilteraddressWithDefaults() *Macfilteraddress`

NewMacfilteraddressWithDefaults instantiates a new Macfilteraddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Macfilteraddress) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Macfilteraddress) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Macfilteraddress) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Macfilteraddress) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetAuthenticationTime

`func (o *Macfilteraddress) GetAuthenticationTime() int64`

GetAuthenticationTime returns the AuthenticationTime field if non-nil, zero value otherwise.

### GetAuthenticationTimeOk

`func (o *Macfilteraddress) GetAuthenticationTimeOk() (*int64, bool)`

GetAuthenticationTimeOk returns a tuple with the AuthenticationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationTime

`func (o *Macfilteraddress) SetAuthenticationTime(v int64)`

SetAuthenticationTime sets AuthenticationTime field to given value.

### HasAuthenticationTime

`func (o *Macfilteraddress) HasAuthenticationTime() bool`

HasAuthenticationTime returns a boolean if a field has been set.

### GetComment

`func (o *Macfilteraddress) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Macfilteraddress) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Macfilteraddress) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Macfilteraddress) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpirationTime

`func (o *Macfilteraddress) GetExpirationTime() int64`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *Macfilteraddress) GetExpirationTimeOk() (*int64, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *Macfilteraddress) SetExpirationTime(v int64)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *Macfilteraddress) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetExtAttrsPlus

`func (o *Macfilteraddress) GetExtAttrsPlus() map[string]ExtAttrs`

GetExtAttrsPlus returns the ExtAttrsPlus field if non-nil, zero value otherwise.

### GetExtAttrsPlusOk

`func (o *Macfilteraddress) GetExtAttrsPlusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsPlusOk returns a tuple with the ExtAttrsPlus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsPlus

`func (o *Macfilteraddress) SetExtAttrsPlus(v map[string]ExtAttrs)`

SetExtAttrsPlus sets ExtAttrsPlus field to given value.

### HasExtAttrsPlus

`func (o *Macfilteraddress) HasExtAttrsPlus() bool`

HasExtAttrsPlus returns a boolean if a field has been set.

### GetExtAttrsMinus

`func (o *Macfilteraddress) GetExtAttrsMinus() map[string]ExtAttrs`

GetExtAttrsMinus returns the ExtAttrsMinus field if non-nil, zero value otherwise.

### GetExtAttrsMinusOk

`func (o *Macfilteraddress) GetExtAttrsMinusOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsMinusOk returns a tuple with the ExtAttrsMinus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrsMinus

`func (o *Macfilteraddress) SetExtAttrsMinus(v map[string]ExtAttrs)`

SetExtAttrsMinus sets ExtAttrsMinus field to given value.

### HasExtAttrsMinus

`func (o *Macfilteraddress) HasExtAttrsMinus() bool`

HasExtAttrsMinus returns a boolean if a field has been set.

### GetExtAttrs

`func (o *Macfilteraddress) GetExtAttrs() map[string]ExtAttrs`

GetExtAttrs returns the ExtAttrs field if non-nil, zero value otherwise.

### GetExtAttrsOk

`func (o *Macfilteraddress) GetExtAttrsOk() (*map[string]ExtAttrs, bool)`

GetExtAttrsOk returns a tuple with the ExtAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtAttrs

`func (o *Macfilteraddress) SetExtAttrs(v map[string]ExtAttrs)`

SetExtAttrs sets ExtAttrs field to given value.

### HasExtAttrs

`func (o *Macfilteraddress) HasExtAttrs() bool`

HasExtAttrs returns a boolean if a field has been set.

### GetFilter

`func (o *Macfilteraddress) GetFilter() string`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Macfilteraddress) GetFilterOk() (*string, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Macfilteraddress) SetFilter(v string)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Macfilteraddress) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetFingerprint

`func (o *Macfilteraddress) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *Macfilteraddress) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *Macfilteraddress) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *Macfilteraddress) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetGuestCustomField1

`func (o *Macfilteraddress) GetGuestCustomField1() string`

GetGuestCustomField1 returns the GuestCustomField1 field if non-nil, zero value otherwise.

### GetGuestCustomField1Ok

`func (o *Macfilteraddress) GetGuestCustomField1Ok() (*string, bool)`

GetGuestCustomField1Ok returns a tuple with the GuestCustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField1

`func (o *Macfilteraddress) SetGuestCustomField1(v string)`

SetGuestCustomField1 sets GuestCustomField1 field to given value.

### HasGuestCustomField1

`func (o *Macfilteraddress) HasGuestCustomField1() bool`

HasGuestCustomField1 returns a boolean if a field has been set.

### GetGuestCustomField2

`func (o *Macfilteraddress) GetGuestCustomField2() string`

GetGuestCustomField2 returns the GuestCustomField2 field if non-nil, zero value otherwise.

### GetGuestCustomField2Ok

`func (o *Macfilteraddress) GetGuestCustomField2Ok() (*string, bool)`

GetGuestCustomField2Ok returns a tuple with the GuestCustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField2

`func (o *Macfilteraddress) SetGuestCustomField2(v string)`

SetGuestCustomField2 sets GuestCustomField2 field to given value.

### HasGuestCustomField2

`func (o *Macfilteraddress) HasGuestCustomField2() bool`

HasGuestCustomField2 returns a boolean if a field has been set.

### GetGuestCustomField3

`func (o *Macfilteraddress) GetGuestCustomField3() string`

GetGuestCustomField3 returns the GuestCustomField3 field if non-nil, zero value otherwise.

### GetGuestCustomField3Ok

`func (o *Macfilteraddress) GetGuestCustomField3Ok() (*string, bool)`

GetGuestCustomField3Ok returns a tuple with the GuestCustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField3

`func (o *Macfilteraddress) SetGuestCustomField3(v string)`

SetGuestCustomField3 sets GuestCustomField3 field to given value.

### HasGuestCustomField3

`func (o *Macfilteraddress) HasGuestCustomField3() bool`

HasGuestCustomField3 returns a boolean if a field has been set.

### GetGuestCustomField4

`func (o *Macfilteraddress) GetGuestCustomField4() string`

GetGuestCustomField4 returns the GuestCustomField4 field if non-nil, zero value otherwise.

### GetGuestCustomField4Ok

`func (o *Macfilteraddress) GetGuestCustomField4Ok() (*string, bool)`

GetGuestCustomField4Ok returns a tuple with the GuestCustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomField4

`func (o *Macfilteraddress) SetGuestCustomField4(v string)`

SetGuestCustomField4 sets GuestCustomField4 field to given value.

### HasGuestCustomField4

`func (o *Macfilteraddress) HasGuestCustomField4() bool`

HasGuestCustomField4 returns a boolean if a field has been set.

### GetGuestEmail

`func (o *Macfilteraddress) GetGuestEmail() string`

GetGuestEmail returns the GuestEmail field if non-nil, zero value otherwise.

### GetGuestEmailOk

`func (o *Macfilteraddress) GetGuestEmailOk() (*string, bool)`

GetGuestEmailOk returns a tuple with the GuestEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestEmail

`func (o *Macfilteraddress) SetGuestEmail(v string)`

SetGuestEmail sets GuestEmail field to given value.

### HasGuestEmail

`func (o *Macfilteraddress) HasGuestEmail() bool`

HasGuestEmail returns a boolean if a field has been set.

### GetGuestFirstName

`func (o *Macfilteraddress) GetGuestFirstName() string`

GetGuestFirstName returns the GuestFirstName field if non-nil, zero value otherwise.

### GetGuestFirstNameOk

`func (o *Macfilteraddress) GetGuestFirstNameOk() (*string, bool)`

GetGuestFirstNameOk returns a tuple with the GuestFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestFirstName

`func (o *Macfilteraddress) SetGuestFirstName(v string)`

SetGuestFirstName sets GuestFirstName field to given value.

### HasGuestFirstName

`func (o *Macfilteraddress) HasGuestFirstName() bool`

HasGuestFirstName returns a boolean if a field has been set.

### GetGuestLastName

`func (o *Macfilteraddress) GetGuestLastName() string`

GetGuestLastName returns the GuestLastName field if non-nil, zero value otherwise.

### GetGuestLastNameOk

`func (o *Macfilteraddress) GetGuestLastNameOk() (*string, bool)`

GetGuestLastNameOk returns a tuple with the GuestLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestLastName

`func (o *Macfilteraddress) SetGuestLastName(v string)`

SetGuestLastName sets GuestLastName field to given value.

### HasGuestLastName

`func (o *Macfilteraddress) HasGuestLastName() bool`

HasGuestLastName returns a boolean if a field has been set.

### GetGuestMiddleName

`func (o *Macfilteraddress) GetGuestMiddleName() string`

GetGuestMiddleName returns the GuestMiddleName field if non-nil, zero value otherwise.

### GetGuestMiddleNameOk

`func (o *Macfilteraddress) GetGuestMiddleNameOk() (*string, bool)`

GetGuestMiddleNameOk returns a tuple with the GuestMiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestMiddleName

`func (o *Macfilteraddress) SetGuestMiddleName(v string)`

SetGuestMiddleName sets GuestMiddleName field to given value.

### HasGuestMiddleName

`func (o *Macfilteraddress) HasGuestMiddleName() bool`

HasGuestMiddleName returns a boolean if a field has been set.

### GetGuestPhone

`func (o *Macfilteraddress) GetGuestPhone() string`

GetGuestPhone returns the GuestPhone field if non-nil, zero value otherwise.

### GetGuestPhoneOk

`func (o *Macfilteraddress) GetGuestPhoneOk() (*string, bool)`

GetGuestPhoneOk returns a tuple with the GuestPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestPhone

`func (o *Macfilteraddress) SetGuestPhone(v string)`

SetGuestPhone sets GuestPhone field to given value.

### HasGuestPhone

`func (o *Macfilteraddress) HasGuestPhone() bool`

HasGuestPhone returns a boolean if a field has been set.

### GetIsRegisteredUser

`func (o *Macfilteraddress) GetIsRegisteredUser() bool`

GetIsRegisteredUser returns the IsRegisteredUser field if non-nil, zero value otherwise.

### GetIsRegisteredUserOk

`func (o *Macfilteraddress) GetIsRegisteredUserOk() (*bool, bool)`

GetIsRegisteredUserOk returns a tuple with the IsRegisteredUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegisteredUser

`func (o *Macfilteraddress) SetIsRegisteredUser(v bool)`

SetIsRegisteredUser sets IsRegisteredUser field to given value.

### HasIsRegisteredUser

`func (o *Macfilteraddress) HasIsRegisteredUser() bool`

HasIsRegisteredUser returns a boolean if a field has been set.

### GetMac

`func (o *Macfilteraddress) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Macfilteraddress) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Macfilteraddress) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Macfilteraddress) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetNeverExpires

`func (o *Macfilteraddress) GetNeverExpires() bool`

GetNeverExpires returns the NeverExpires field if non-nil, zero value otherwise.

### GetNeverExpiresOk

`func (o *Macfilteraddress) GetNeverExpiresOk() (*bool, bool)`

GetNeverExpiresOk returns a tuple with the NeverExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeverExpires

`func (o *Macfilteraddress) SetNeverExpires(v bool)`

SetNeverExpires sets NeverExpires field to given value.

### HasNeverExpires

`func (o *Macfilteraddress) HasNeverExpires() bool`

HasNeverExpires returns a boolean if a field has been set.

### GetReservedForInfoblox

`func (o *Macfilteraddress) GetReservedForInfoblox() string`

GetReservedForInfoblox returns the ReservedForInfoblox field if non-nil, zero value otherwise.

### GetReservedForInfobloxOk

`func (o *Macfilteraddress) GetReservedForInfobloxOk() (*string, bool)`

GetReservedForInfobloxOk returns a tuple with the ReservedForInfoblox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedForInfoblox

`func (o *Macfilteraddress) SetReservedForInfoblox(v string)`

SetReservedForInfoblox sets ReservedForInfoblox field to given value.

### HasReservedForInfoblox

`func (o *Macfilteraddress) HasReservedForInfoblox() bool`

HasReservedForInfoblox returns a boolean if a field has been set.

### GetUsername

`func (o *Macfilteraddress) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Macfilteraddress) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Macfilteraddress) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Macfilteraddress) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


