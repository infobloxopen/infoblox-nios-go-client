# ZoneAuthDnssecKeyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableKskAutoRollover** | Pointer to **bool** | If set to True, automatic rollovers for the signing key is enabled. | [optional] 
**KskAlgorithm** | Pointer to **string** | Key Signing Key algorithm. Deprecated. | [optional] 
**KskAlgorithms** | Pointer to [**ZoneauthdnsseckeyparamsKskAlgorithms**](ZoneauthdnsseckeyparamsKskAlgorithms.md) |  | [optional] 
**KskRollover** | Pointer to **int64** | Key Signing Key rollover interval, in seconds. | [optional] 
**KskSize** | Pointer to **int64** | Key Signing Key size, in bits. Deprecated. | [optional] 
**NextSecureType** | Pointer to **string** | NSEC (next secure) types. | [optional] 
**KskRolloverNotificationConfig** | Pointer to **string** | This field controls events for which users will be notified. | [optional] 
**KskSnmpNotificationEnabled** | Pointer to **bool** | Enable SNMP notifications for KSK related events. | [optional] 
**KskEmailNotificationEnabled** | Pointer to **bool** | Enable email notifications for KSK related events. | [optional] 
**Nsec3SaltMinLength** | Pointer to **int64** | The minimum length for NSEC3 salts. | [optional] 
**Nsec3SaltMaxLength** | Pointer to **int64** | The maximum length for NSEC3 salts. | [optional] 
**Nsec3Iterations** | Pointer to **int64** | The number of iterations used for hashing NSEC3. | [optional] 
**SignatureExpiration** | Pointer to **int64** | Signature expiration time, in seconds. | [optional] 
**ZskAlgorithm** | Pointer to **string** | Zone Signing Key algorithm. Deprecated. | [optional] 
**ZskAlgorithms** | Pointer to [**ZoneauthdnsseckeyparamsZskAlgorithms**](ZoneauthdnsseckeyparamsZskAlgorithms.md) |  | [optional] 
**ZskRollover** | Pointer to **int64** | Zone Signing Key rollover interval, in seconds. | [optional] 
**ZskRolloverMechanism** | Pointer to **string** | Zone Signing Key rollover mechanism. | [optional] 
**ZskSize** | Pointer to **int64** | Zone Signing Key size, in bits. Deprecated. | [optional] 

## Methods

### NewZoneAuthDnssecKeyParams

`func NewZoneAuthDnssecKeyParams() *ZoneAuthDnssecKeyParams`

NewZoneAuthDnssecKeyParams instantiates a new ZoneAuthDnssecKeyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZoneAuthDnssecKeyParamsWithDefaults

`func NewZoneAuthDnssecKeyParamsWithDefaults() *ZoneAuthDnssecKeyParams`

NewZoneAuthDnssecKeyParamsWithDefaults instantiates a new ZoneAuthDnssecKeyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableKskAutoRollover

`func (o *ZoneAuthDnssecKeyParams) GetEnableKskAutoRollover() bool`

GetEnableKskAutoRollover returns the EnableKskAutoRollover field if non-nil, zero value otherwise.

### GetEnableKskAutoRolloverOk

`func (o *ZoneAuthDnssecKeyParams) GetEnableKskAutoRolloverOk() (*bool, bool)`

GetEnableKskAutoRolloverOk returns a tuple with the EnableKskAutoRollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableKskAutoRollover

`func (o *ZoneAuthDnssecKeyParams) SetEnableKskAutoRollover(v bool)`

SetEnableKskAutoRollover sets EnableKskAutoRollover field to given value.

### HasEnableKskAutoRollover

`func (o *ZoneAuthDnssecKeyParams) HasEnableKskAutoRollover() bool`

HasEnableKskAutoRollover returns a boolean if a field has been set.

### GetKskAlgorithm

`func (o *ZoneAuthDnssecKeyParams) GetKskAlgorithm() string`

GetKskAlgorithm returns the KskAlgorithm field if non-nil, zero value otherwise.

### GetKskAlgorithmOk

`func (o *ZoneAuthDnssecKeyParams) GetKskAlgorithmOk() (*string, bool)`

GetKskAlgorithmOk returns a tuple with the KskAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskAlgorithm

`func (o *ZoneAuthDnssecKeyParams) SetKskAlgorithm(v string)`

SetKskAlgorithm sets KskAlgorithm field to given value.

### HasKskAlgorithm

`func (o *ZoneAuthDnssecKeyParams) HasKskAlgorithm() bool`

HasKskAlgorithm returns a boolean if a field has been set.

### GetKskAlgorithms

`func (o *ZoneAuthDnssecKeyParams) GetKskAlgorithms() ZoneauthdnsseckeyparamsKskAlgorithms`

GetKskAlgorithms returns the KskAlgorithms field if non-nil, zero value otherwise.

### GetKskAlgorithmsOk

`func (o *ZoneAuthDnssecKeyParams) GetKskAlgorithmsOk() (*ZoneauthdnsseckeyparamsKskAlgorithms, bool)`

GetKskAlgorithmsOk returns a tuple with the KskAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskAlgorithms

`func (o *ZoneAuthDnssecKeyParams) SetKskAlgorithms(v ZoneauthdnsseckeyparamsKskAlgorithms)`

SetKskAlgorithms sets KskAlgorithms field to given value.

### HasKskAlgorithms

`func (o *ZoneAuthDnssecKeyParams) HasKskAlgorithms() bool`

HasKskAlgorithms returns a boolean if a field has been set.

### GetKskRollover

`func (o *ZoneAuthDnssecKeyParams) GetKskRollover() int64`

GetKskRollover returns the KskRollover field if non-nil, zero value otherwise.

### GetKskRolloverOk

`func (o *ZoneAuthDnssecKeyParams) GetKskRolloverOk() (*int64, bool)`

GetKskRolloverOk returns a tuple with the KskRollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskRollover

`func (o *ZoneAuthDnssecKeyParams) SetKskRollover(v int64)`

SetKskRollover sets KskRollover field to given value.

### HasKskRollover

`func (o *ZoneAuthDnssecKeyParams) HasKskRollover() bool`

HasKskRollover returns a boolean if a field has been set.

### GetKskSize

`func (o *ZoneAuthDnssecKeyParams) GetKskSize() int64`

GetKskSize returns the KskSize field if non-nil, zero value otherwise.

### GetKskSizeOk

`func (o *ZoneAuthDnssecKeyParams) GetKskSizeOk() (*int64, bool)`

GetKskSizeOk returns a tuple with the KskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskSize

`func (o *ZoneAuthDnssecKeyParams) SetKskSize(v int64)`

SetKskSize sets KskSize field to given value.

### HasKskSize

`func (o *ZoneAuthDnssecKeyParams) HasKskSize() bool`

HasKskSize returns a boolean if a field has been set.

### GetNextSecureType

`func (o *ZoneAuthDnssecKeyParams) GetNextSecureType() string`

GetNextSecureType returns the NextSecureType field if non-nil, zero value otherwise.

### GetNextSecureTypeOk

`func (o *ZoneAuthDnssecKeyParams) GetNextSecureTypeOk() (*string, bool)`

GetNextSecureTypeOk returns a tuple with the NextSecureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSecureType

`func (o *ZoneAuthDnssecKeyParams) SetNextSecureType(v string)`

SetNextSecureType sets NextSecureType field to given value.

### HasNextSecureType

`func (o *ZoneAuthDnssecKeyParams) HasNextSecureType() bool`

HasNextSecureType returns a boolean if a field has been set.

### GetKskRolloverNotificationConfig

`func (o *ZoneAuthDnssecKeyParams) GetKskRolloverNotificationConfig() string`

GetKskRolloverNotificationConfig returns the KskRolloverNotificationConfig field if non-nil, zero value otherwise.

### GetKskRolloverNotificationConfigOk

`func (o *ZoneAuthDnssecKeyParams) GetKskRolloverNotificationConfigOk() (*string, bool)`

GetKskRolloverNotificationConfigOk returns a tuple with the KskRolloverNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskRolloverNotificationConfig

`func (o *ZoneAuthDnssecKeyParams) SetKskRolloverNotificationConfig(v string)`

SetKskRolloverNotificationConfig sets KskRolloverNotificationConfig field to given value.

### HasKskRolloverNotificationConfig

`func (o *ZoneAuthDnssecKeyParams) HasKskRolloverNotificationConfig() bool`

HasKskRolloverNotificationConfig returns a boolean if a field has been set.

### GetKskSnmpNotificationEnabled

`func (o *ZoneAuthDnssecKeyParams) GetKskSnmpNotificationEnabled() bool`

GetKskSnmpNotificationEnabled returns the KskSnmpNotificationEnabled field if non-nil, zero value otherwise.

### GetKskSnmpNotificationEnabledOk

`func (o *ZoneAuthDnssecKeyParams) GetKskSnmpNotificationEnabledOk() (*bool, bool)`

GetKskSnmpNotificationEnabledOk returns a tuple with the KskSnmpNotificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskSnmpNotificationEnabled

`func (o *ZoneAuthDnssecKeyParams) SetKskSnmpNotificationEnabled(v bool)`

SetKskSnmpNotificationEnabled sets KskSnmpNotificationEnabled field to given value.

### HasKskSnmpNotificationEnabled

`func (o *ZoneAuthDnssecKeyParams) HasKskSnmpNotificationEnabled() bool`

HasKskSnmpNotificationEnabled returns a boolean if a field has been set.

### GetKskEmailNotificationEnabled

`func (o *ZoneAuthDnssecKeyParams) GetKskEmailNotificationEnabled() bool`

GetKskEmailNotificationEnabled returns the KskEmailNotificationEnabled field if non-nil, zero value otherwise.

### GetKskEmailNotificationEnabledOk

`func (o *ZoneAuthDnssecKeyParams) GetKskEmailNotificationEnabledOk() (*bool, bool)`

GetKskEmailNotificationEnabledOk returns a tuple with the KskEmailNotificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskEmailNotificationEnabled

`func (o *ZoneAuthDnssecKeyParams) SetKskEmailNotificationEnabled(v bool)`

SetKskEmailNotificationEnabled sets KskEmailNotificationEnabled field to given value.

### HasKskEmailNotificationEnabled

`func (o *ZoneAuthDnssecKeyParams) HasKskEmailNotificationEnabled() bool`

HasKskEmailNotificationEnabled returns a boolean if a field has been set.

### GetNsec3SaltMinLength

`func (o *ZoneAuthDnssecKeyParams) GetNsec3SaltMinLength() int64`

GetNsec3SaltMinLength returns the Nsec3SaltMinLength field if non-nil, zero value otherwise.

### GetNsec3SaltMinLengthOk

`func (o *ZoneAuthDnssecKeyParams) GetNsec3SaltMinLengthOk() (*int64, bool)`

GetNsec3SaltMinLengthOk returns a tuple with the Nsec3SaltMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3SaltMinLength

`func (o *ZoneAuthDnssecKeyParams) SetNsec3SaltMinLength(v int64)`

SetNsec3SaltMinLength sets Nsec3SaltMinLength field to given value.

### HasNsec3SaltMinLength

`func (o *ZoneAuthDnssecKeyParams) HasNsec3SaltMinLength() bool`

HasNsec3SaltMinLength returns a boolean if a field has been set.

### GetNsec3SaltMaxLength

`func (o *ZoneAuthDnssecKeyParams) GetNsec3SaltMaxLength() int64`

GetNsec3SaltMaxLength returns the Nsec3SaltMaxLength field if non-nil, zero value otherwise.

### GetNsec3SaltMaxLengthOk

`func (o *ZoneAuthDnssecKeyParams) GetNsec3SaltMaxLengthOk() (*int64, bool)`

GetNsec3SaltMaxLengthOk returns a tuple with the Nsec3SaltMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3SaltMaxLength

`func (o *ZoneAuthDnssecKeyParams) SetNsec3SaltMaxLength(v int64)`

SetNsec3SaltMaxLength sets Nsec3SaltMaxLength field to given value.

### HasNsec3SaltMaxLength

`func (o *ZoneAuthDnssecKeyParams) HasNsec3SaltMaxLength() bool`

HasNsec3SaltMaxLength returns a boolean if a field has been set.

### GetNsec3Iterations

`func (o *ZoneAuthDnssecKeyParams) GetNsec3Iterations() int64`

GetNsec3Iterations returns the Nsec3Iterations field if non-nil, zero value otherwise.

### GetNsec3IterationsOk

`func (o *ZoneAuthDnssecKeyParams) GetNsec3IterationsOk() (*int64, bool)`

GetNsec3IterationsOk returns a tuple with the Nsec3Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3Iterations

`func (o *ZoneAuthDnssecKeyParams) SetNsec3Iterations(v int64)`

SetNsec3Iterations sets Nsec3Iterations field to given value.

### HasNsec3Iterations

`func (o *ZoneAuthDnssecKeyParams) HasNsec3Iterations() bool`

HasNsec3Iterations returns a boolean if a field has been set.

### GetSignatureExpiration

`func (o *ZoneAuthDnssecKeyParams) GetSignatureExpiration() int64`

GetSignatureExpiration returns the SignatureExpiration field if non-nil, zero value otherwise.

### GetSignatureExpirationOk

`func (o *ZoneAuthDnssecKeyParams) GetSignatureExpirationOk() (*int64, bool)`

GetSignatureExpirationOk returns a tuple with the SignatureExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureExpiration

`func (o *ZoneAuthDnssecKeyParams) SetSignatureExpiration(v int64)`

SetSignatureExpiration sets SignatureExpiration field to given value.

### HasSignatureExpiration

`func (o *ZoneAuthDnssecKeyParams) HasSignatureExpiration() bool`

HasSignatureExpiration returns a boolean if a field has been set.

### GetZskAlgorithm

`func (o *ZoneAuthDnssecKeyParams) GetZskAlgorithm() string`

GetZskAlgorithm returns the ZskAlgorithm field if non-nil, zero value otherwise.

### GetZskAlgorithmOk

`func (o *ZoneAuthDnssecKeyParams) GetZskAlgorithmOk() (*string, bool)`

GetZskAlgorithmOk returns a tuple with the ZskAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskAlgorithm

`func (o *ZoneAuthDnssecKeyParams) SetZskAlgorithm(v string)`

SetZskAlgorithm sets ZskAlgorithm field to given value.

### HasZskAlgorithm

`func (o *ZoneAuthDnssecKeyParams) HasZskAlgorithm() bool`

HasZskAlgorithm returns a boolean if a field has been set.

### GetZskAlgorithms

`func (o *ZoneAuthDnssecKeyParams) GetZskAlgorithms() ZoneauthdnsseckeyparamsZskAlgorithms`

GetZskAlgorithms returns the ZskAlgorithms field if non-nil, zero value otherwise.

### GetZskAlgorithmsOk

`func (o *ZoneAuthDnssecKeyParams) GetZskAlgorithmsOk() (*ZoneauthdnsseckeyparamsZskAlgorithms, bool)`

GetZskAlgorithmsOk returns a tuple with the ZskAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskAlgorithms

`func (o *ZoneAuthDnssecKeyParams) SetZskAlgorithms(v ZoneauthdnsseckeyparamsZskAlgorithms)`

SetZskAlgorithms sets ZskAlgorithms field to given value.

### HasZskAlgorithms

`func (o *ZoneAuthDnssecKeyParams) HasZskAlgorithms() bool`

HasZskAlgorithms returns a boolean if a field has been set.

### GetZskRollover

`func (o *ZoneAuthDnssecKeyParams) GetZskRollover() int64`

GetZskRollover returns the ZskRollover field if non-nil, zero value otherwise.

### GetZskRolloverOk

`func (o *ZoneAuthDnssecKeyParams) GetZskRolloverOk() (*int64, bool)`

GetZskRolloverOk returns a tuple with the ZskRollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskRollover

`func (o *ZoneAuthDnssecKeyParams) SetZskRollover(v int64)`

SetZskRollover sets ZskRollover field to given value.

### HasZskRollover

`func (o *ZoneAuthDnssecKeyParams) HasZskRollover() bool`

HasZskRollover returns a boolean if a field has been set.

### GetZskRolloverMechanism

`func (o *ZoneAuthDnssecKeyParams) GetZskRolloverMechanism() string`

GetZskRolloverMechanism returns the ZskRolloverMechanism field if non-nil, zero value otherwise.

### GetZskRolloverMechanismOk

`func (o *ZoneAuthDnssecKeyParams) GetZskRolloverMechanismOk() (*string, bool)`

GetZskRolloverMechanismOk returns a tuple with the ZskRolloverMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskRolloverMechanism

`func (o *ZoneAuthDnssecKeyParams) SetZskRolloverMechanism(v string)`

SetZskRolloverMechanism sets ZskRolloverMechanism field to given value.

### HasZskRolloverMechanism

`func (o *ZoneAuthDnssecKeyParams) HasZskRolloverMechanism() bool`

HasZskRolloverMechanism returns a boolean if a field has been set.

### GetZskSize

`func (o *ZoneAuthDnssecKeyParams) GetZskSize() int64`

GetZskSize returns the ZskSize field if non-nil, zero value otherwise.

### GetZskSizeOk

`func (o *ZoneAuthDnssecKeyParams) GetZskSizeOk() (*int64, bool)`

GetZskSizeOk returns a tuple with the ZskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskSize

`func (o *ZoneAuthDnssecKeyParams) SetZskSize(v int64)`

SetZskSize sets ZskSize field to given value.

### HasZskSize

`func (o *ZoneAuthDnssecKeyParams) HasZskSize() bool`

HasZskSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


