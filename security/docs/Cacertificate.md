# Cacertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] [readonly] 
**DistinguishedName** | Pointer to **string** | The certificate subject name. | [optional] [readonly] 
**Issuer** | Pointer to **string** | The certificate issuer subject name. | [optional] [readonly] 
**Serial** | Pointer to **string** | The certificate serial number in hex format. | [optional] [readonly] 
**UsedBy** | Pointer to **string** | Information about the CA certificate usage. | [optional] [readonly] 
**ValidNotAfter** | Pointer to **int64** | The date after which the certificate becomes invalid. | [optional] [readonly] 
**ValidNotBefore** | Pointer to **int64** | The date before which the certificate is not valid. | [optional] [readonly] 

## Methods

### NewCacertificate

`func NewCacertificate() *Cacertificate`

NewCacertificate instantiates a new Cacertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCacertificateWithDefaults

`func NewCacertificateWithDefaults() *Cacertificate`

NewCacertificateWithDefaults instantiates a new Cacertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *Cacertificate) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Cacertificate) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Cacertificate) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Cacertificate) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *Cacertificate) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *Cacertificate) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *Cacertificate) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *Cacertificate) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetIssuer

`func (o *Cacertificate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Cacertificate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Cacertificate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *Cacertificate) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSerial

`func (o *Cacertificate) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *Cacertificate) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *Cacertificate) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *Cacertificate) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUsedBy

`func (o *Cacertificate) GetUsedBy() string`

GetUsedBy returns the UsedBy field if non-nil, zero value otherwise.

### GetUsedByOk

`func (o *Cacertificate) GetUsedByOk() (*string, bool)`

GetUsedByOk returns a tuple with the UsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBy

`func (o *Cacertificate) SetUsedBy(v string)`

SetUsedBy sets UsedBy field to given value.

### HasUsedBy

`func (o *Cacertificate) HasUsedBy() bool`

HasUsedBy returns a boolean if a field has been set.

### GetValidNotAfter

`func (o *Cacertificate) GetValidNotAfter() int64`

GetValidNotAfter returns the ValidNotAfter field if non-nil, zero value otherwise.

### GetValidNotAfterOk

`func (o *Cacertificate) GetValidNotAfterOk() (*int64, bool)`

GetValidNotAfterOk returns a tuple with the ValidNotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotAfter

`func (o *Cacertificate) SetValidNotAfter(v int64)`

SetValidNotAfter sets ValidNotAfter field to given value.

### HasValidNotAfter

`func (o *Cacertificate) HasValidNotAfter() bool`

HasValidNotAfter returns a boolean if a field has been set.

### GetValidNotBefore

`func (o *Cacertificate) GetValidNotBefore() int64`

GetValidNotBefore returns the ValidNotBefore field if non-nil, zero value otherwise.

### GetValidNotBeforeOk

`func (o *Cacertificate) GetValidNotBeforeOk() (*int64, bool)`

GetValidNotBeforeOk returns a tuple with the ValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotBefore

`func (o *Cacertificate) SetValidNotBefore(v int64)`

SetValidNotBefore sets ValidNotBefore field to given value.

### HasValidNotBefore

`func (o *Cacertificate) HasValidNotBefore() bool`

HasValidNotBefore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


