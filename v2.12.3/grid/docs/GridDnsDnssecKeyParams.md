# GridDnsDnssecKeyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableKskAutoRollover** | Pointer to **bool** | If set to True, automatic rollovers for the signing key is enabled. | [optional] 
**KskAlgorithm** | Pointer to **string** | Key Signing Key algorithm. Deprecated. | [optional] 
**KskAlgorithms** | Pointer to [**GriddnsdnsseckeyparamsKskAlgorithms**](GriddnsdnsseckeyparamsKskAlgorithms.md) |  | [optional] 
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
**ZskAlgorithms** | Pointer to [**GriddnsdnsseckeyparamsZskAlgorithms**](GriddnsdnsseckeyparamsZskAlgorithms.md) |  | [optional] 
**ZskRollover** | Pointer to **int64** | Zone Signing Key rollover interval, in seconds. | [optional] 
**ZskRolloverMechanism** | Pointer to **string** | Zone Signing Key rollover mechanism. | [optional] 
**ZskSize** | Pointer to **int64** | Zone Signing Key size, in bits. Deprecated. | [optional] 

## Methods

### NewGridDnsDnssecKeyParams

`func NewGridDnsDnssecKeyParams() *GridDnsDnssecKeyParams`

NewGridDnsDnssecKeyParams instantiates a new GridDnsDnssecKeyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGridDnsDnssecKeyParamsWithDefaults

`func NewGridDnsDnssecKeyParamsWithDefaults() *GridDnsDnssecKeyParams`

NewGridDnsDnssecKeyParamsWithDefaults instantiates a new GridDnsDnssecKeyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableKskAutoRollover

`func (o *GridDnsDnssecKeyParams) GetEnableKskAutoRollover() bool`

GetEnableKskAutoRollover returns the EnableKskAutoRollover field if non-nil, zero value otherwise.

### GetEnableKskAutoRolloverOk

`func (o *GridDnsDnssecKeyParams) GetEnableKskAutoRolloverOk() (*bool, bool)`

GetEnableKskAutoRolloverOk returns a tuple with the EnableKskAutoRollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableKskAutoRollover

`func (o *GridDnsDnssecKeyParams) SetEnableKskAutoRollover(v bool)`

SetEnableKskAutoRollover sets EnableKskAutoRollover field to given value.

### HasEnableKskAutoRollover

`func (o *GridDnsDnssecKeyParams) HasEnableKskAutoRollover() bool`

HasEnableKskAutoRollover returns a boolean if a field has been set.

### GetKskAlgorithm

`func (o *GridDnsDnssecKeyParams) GetKskAlgorithm() string`

GetKskAlgorithm returns the KskAlgorithm field if non-nil, zero value otherwise.

### GetKskAlgorithmOk

`func (o *GridDnsDnssecKeyParams) GetKskAlgorithmOk() (*string, bool)`

GetKskAlgorithmOk returns a tuple with the KskAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskAlgorithm

`func (o *GridDnsDnssecKeyParams) SetKskAlgorithm(v string)`

SetKskAlgorithm sets KskAlgorithm field to given value.

### HasKskAlgorithm

`func (o *GridDnsDnssecKeyParams) HasKskAlgorithm() bool`

HasKskAlgorithm returns a boolean if a field has been set.

### GetKskAlgorithms

`func (o *GridDnsDnssecKeyParams) GetKskAlgorithms() GriddnsdnsseckeyparamsKskAlgorithms`

GetKskAlgorithms returns the KskAlgorithms field if non-nil, zero value otherwise.

### GetKskAlgorithmsOk

`func (o *GridDnsDnssecKeyParams) GetKskAlgorithmsOk() (*GriddnsdnsseckeyparamsKskAlgorithms, bool)`

GetKskAlgorithmsOk returns a tuple with the KskAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskAlgorithms

`func (o *GridDnsDnssecKeyParams) SetKskAlgorithms(v GriddnsdnsseckeyparamsKskAlgorithms)`

SetKskAlgorithms sets KskAlgorithms field to given value.

### HasKskAlgorithms

`func (o *GridDnsDnssecKeyParams) HasKskAlgorithms() bool`

HasKskAlgorithms returns a boolean if a field has been set.

### GetKskRollover

`func (o *GridDnsDnssecKeyParams) GetKskRollover() int64`

GetKskRollover returns the KskRollover field if non-nil, zero value otherwise.

### GetKskRolloverOk

`func (o *GridDnsDnssecKeyParams) GetKskRolloverOk() (*int64, bool)`

GetKskRolloverOk returns a tuple with the KskRollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskRollover

`func (o *GridDnsDnssecKeyParams) SetKskRollover(v int64)`

SetKskRollover sets KskRollover field to given value.

### HasKskRollover

`func (o *GridDnsDnssecKeyParams) HasKskRollover() bool`

HasKskRollover returns a boolean if a field has been set.

### GetKskSize

`func (o *GridDnsDnssecKeyParams) GetKskSize() int64`

GetKskSize returns the KskSize field if non-nil, zero value otherwise.

### GetKskSizeOk

`func (o *GridDnsDnssecKeyParams) GetKskSizeOk() (*int64, bool)`

GetKskSizeOk returns a tuple with the KskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskSize

`func (o *GridDnsDnssecKeyParams) SetKskSize(v int64)`

SetKskSize sets KskSize field to given value.

### HasKskSize

`func (o *GridDnsDnssecKeyParams) HasKskSize() bool`

HasKskSize returns a boolean if a field has been set.

### GetNextSecureType

`func (o *GridDnsDnssecKeyParams) GetNextSecureType() string`

GetNextSecureType returns the NextSecureType field if non-nil, zero value otherwise.

### GetNextSecureTypeOk

`func (o *GridDnsDnssecKeyParams) GetNextSecureTypeOk() (*string, bool)`

GetNextSecureTypeOk returns a tuple with the NextSecureType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSecureType

`func (o *GridDnsDnssecKeyParams) SetNextSecureType(v string)`

SetNextSecureType sets NextSecureType field to given value.

### HasNextSecureType

`func (o *GridDnsDnssecKeyParams) HasNextSecureType() bool`

HasNextSecureType returns a boolean if a field has been set.

### GetKskRolloverNotificationConfig

`func (o *GridDnsDnssecKeyParams) GetKskRolloverNotificationConfig() string`

GetKskRolloverNotificationConfig returns the KskRolloverNotificationConfig field if non-nil, zero value otherwise.

### GetKskRolloverNotificationConfigOk

`func (o *GridDnsDnssecKeyParams) GetKskRolloverNotificationConfigOk() (*string, bool)`

GetKskRolloverNotificationConfigOk returns a tuple with the KskRolloverNotificationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskRolloverNotificationConfig

`func (o *GridDnsDnssecKeyParams) SetKskRolloverNotificationConfig(v string)`

SetKskRolloverNotificationConfig sets KskRolloverNotificationConfig field to given value.

### HasKskRolloverNotificationConfig

`func (o *GridDnsDnssecKeyParams) HasKskRolloverNotificationConfig() bool`

HasKskRolloverNotificationConfig returns a boolean if a field has been set.

### GetKskSnmpNotificationEnabled

`func (o *GridDnsDnssecKeyParams) GetKskSnmpNotificationEnabled() bool`

GetKskSnmpNotificationEnabled returns the KskSnmpNotificationEnabled field if non-nil, zero value otherwise.

### GetKskSnmpNotificationEnabledOk

`func (o *GridDnsDnssecKeyParams) GetKskSnmpNotificationEnabledOk() (*bool, bool)`

GetKskSnmpNotificationEnabledOk returns a tuple with the KskSnmpNotificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskSnmpNotificationEnabled

`func (o *GridDnsDnssecKeyParams) SetKskSnmpNotificationEnabled(v bool)`

SetKskSnmpNotificationEnabled sets KskSnmpNotificationEnabled field to given value.

### HasKskSnmpNotificationEnabled

`func (o *GridDnsDnssecKeyParams) HasKskSnmpNotificationEnabled() bool`

HasKskSnmpNotificationEnabled returns a boolean if a field has been set.

### GetKskEmailNotificationEnabled

`func (o *GridDnsDnssecKeyParams) GetKskEmailNotificationEnabled() bool`

GetKskEmailNotificationEnabled returns the KskEmailNotificationEnabled field if non-nil, zero value otherwise.

### GetKskEmailNotificationEnabledOk

`func (o *GridDnsDnssecKeyParams) GetKskEmailNotificationEnabledOk() (*bool, bool)`

GetKskEmailNotificationEnabledOk returns a tuple with the KskEmailNotificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKskEmailNotificationEnabled

`func (o *GridDnsDnssecKeyParams) SetKskEmailNotificationEnabled(v bool)`

SetKskEmailNotificationEnabled sets KskEmailNotificationEnabled field to given value.

### HasKskEmailNotificationEnabled

`func (o *GridDnsDnssecKeyParams) HasKskEmailNotificationEnabled() bool`

HasKskEmailNotificationEnabled returns a boolean if a field has been set.

### GetNsec3SaltMinLength

`func (o *GridDnsDnssecKeyParams) GetNsec3SaltMinLength() int64`

GetNsec3SaltMinLength returns the Nsec3SaltMinLength field if non-nil, zero value otherwise.

### GetNsec3SaltMinLengthOk

`func (o *GridDnsDnssecKeyParams) GetNsec3SaltMinLengthOk() (*int64, bool)`

GetNsec3SaltMinLengthOk returns a tuple with the Nsec3SaltMinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3SaltMinLength

`func (o *GridDnsDnssecKeyParams) SetNsec3SaltMinLength(v int64)`

SetNsec3SaltMinLength sets Nsec3SaltMinLength field to given value.

### HasNsec3SaltMinLength

`func (o *GridDnsDnssecKeyParams) HasNsec3SaltMinLength() bool`

HasNsec3SaltMinLength returns a boolean if a field has been set.

### GetNsec3SaltMaxLength

`func (o *GridDnsDnssecKeyParams) GetNsec3SaltMaxLength() int64`

GetNsec3SaltMaxLength returns the Nsec3SaltMaxLength field if non-nil, zero value otherwise.

### GetNsec3SaltMaxLengthOk

`func (o *GridDnsDnssecKeyParams) GetNsec3SaltMaxLengthOk() (*int64, bool)`

GetNsec3SaltMaxLengthOk returns a tuple with the Nsec3SaltMaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3SaltMaxLength

`func (o *GridDnsDnssecKeyParams) SetNsec3SaltMaxLength(v int64)`

SetNsec3SaltMaxLength sets Nsec3SaltMaxLength field to given value.

### HasNsec3SaltMaxLength

`func (o *GridDnsDnssecKeyParams) HasNsec3SaltMaxLength() bool`

HasNsec3SaltMaxLength returns a boolean if a field has been set.

### GetNsec3Iterations

`func (o *GridDnsDnssecKeyParams) GetNsec3Iterations() int64`

GetNsec3Iterations returns the Nsec3Iterations field if non-nil, zero value otherwise.

### GetNsec3IterationsOk

`func (o *GridDnsDnssecKeyParams) GetNsec3IterationsOk() (*int64, bool)`

GetNsec3IterationsOk returns a tuple with the Nsec3Iterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsec3Iterations

`func (o *GridDnsDnssecKeyParams) SetNsec3Iterations(v int64)`

SetNsec3Iterations sets Nsec3Iterations field to given value.

### HasNsec3Iterations

`func (o *GridDnsDnssecKeyParams) HasNsec3Iterations() bool`

HasNsec3Iterations returns a boolean if a field has been set.

### GetSignatureExpiration

`func (o *GridDnsDnssecKeyParams) GetSignatureExpiration() int64`

GetSignatureExpiration returns the SignatureExpiration field if non-nil, zero value otherwise.

### GetSignatureExpirationOk

`func (o *GridDnsDnssecKeyParams) GetSignatureExpirationOk() (*int64, bool)`

GetSignatureExpirationOk returns a tuple with the SignatureExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureExpiration

`func (o *GridDnsDnssecKeyParams) SetSignatureExpiration(v int64)`

SetSignatureExpiration sets SignatureExpiration field to given value.

### HasSignatureExpiration

`func (o *GridDnsDnssecKeyParams) HasSignatureExpiration() bool`

HasSignatureExpiration returns a boolean if a field has been set.

### GetZskAlgorithm

`func (o *GridDnsDnssecKeyParams) GetZskAlgorithm() string`

GetZskAlgorithm returns the ZskAlgorithm field if non-nil, zero value otherwise.

### GetZskAlgorithmOk

`func (o *GridDnsDnssecKeyParams) GetZskAlgorithmOk() (*string, bool)`

GetZskAlgorithmOk returns a tuple with the ZskAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskAlgorithm

`func (o *GridDnsDnssecKeyParams) SetZskAlgorithm(v string)`

SetZskAlgorithm sets ZskAlgorithm field to given value.

### HasZskAlgorithm

`func (o *GridDnsDnssecKeyParams) HasZskAlgorithm() bool`

HasZskAlgorithm returns a boolean if a field has been set.

### GetZskAlgorithms

`func (o *GridDnsDnssecKeyParams) GetZskAlgorithms() GriddnsdnsseckeyparamsZskAlgorithms`

GetZskAlgorithms returns the ZskAlgorithms field if non-nil, zero value otherwise.

### GetZskAlgorithmsOk

`func (o *GridDnsDnssecKeyParams) GetZskAlgorithmsOk() (*GriddnsdnsseckeyparamsZskAlgorithms, bool)`

GetZskAlgorithmsOk returns a tuple with the ZskAlgorithms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskAlgorithms

`func (o *GridDnsDnssecKeyParams) SetZskAlgorithms(v GriddnsdnsseckeyparamsZskAlgorithms)`

SetZskAlgorithms sets ZskAlgorithms field to given value.

### HasZskAlgorithms

`func (o *GridDnsDnssecKeyParams) HasZskAlgorithms() bool`

HasZskAlgorithms returns a boolean if a field has been set.

### GetZskRollover

`func (o *GridDnsDnssecKeyParams) GetZskRollover() int64`

GetZskRollover returns the ZskRollover field if non-nil, zero value otherwise.

### GetZskRolloverOk

`func (o *GridDnsDnssecKeyParams) GetZskRolloverOk() (*int64, bool)`

GetZskRolloverOk returns a tuple with the ZskRollover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskRollover

`func (o *GridDnsDnssecKeyParams) SetZskRollover(v int64)`

SetZskRollover sets ZskRollover field to given value.

### HasZskRollover

`func (o *GridDnsDnssecKeyParams) HasZskRollover() bool`

HasZskRollover returns a boolean if a field has been set.

### GetZskRolloverMechanism

`func (o *GridDnsDnssecKeyParams) GetZskRolloverMechanism() string`

GetZskRolloverMechanism returns the ZskRolloverMechanism field if non-nil, zero value otherwise.

### GetZskRolloverMechanismOk

`func (o *GridDnsDnssecKeyParams) GetZskRolloverMechanismOk() (*string, bool)`

GetZskRolloverMechanismOk returns a tuple with the ZskRolloverMechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskRolloverMechanism

`func (o *GridDnsDnssecKeyParams) SetZskRolloverMechanism(v string)`

SetZskRolloverMechanism sets ZskRolloverMechanism field to given value.

### HasZskRolloverMechanism

`func (o *GridDnsDnssecKeyParams) HasZskRolloverMechanism() bool`

HasZskRolloverMechanism returns a boolean if a field has been set.

### GetZskSize

`func (o *GridDnsDnssecKeyParams) GetZskSize() int64`

GetZskSize returns the ZskSize field if non-nil, zero value otherwise.

### GetZskSizeOk

`func (o *GridDnsDnssecKeyParams) GetZskSizeOk() (*int64, bool)`

GetZskSizeOk returns a tuple with the ZskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZskSize

`func (o *GridDnsDnssecKeyParams) SetZskSize(v int64)`

SetZskSize sets ZskSize field to given value.

### HasZskSize

`func (o *GridDnsDnssecKeyParams) HasZskSize() bool`

HasZskSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


