# GetCacertificateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ref** | Pointer to **string** | The reference to the object. | [optional] 
**DistinguishedName** | Pointer to **string** | The certificate subject name. | [optional] [readonly] 
**Issuer** | Pointer to **string** | The certificate issuer subject name. | [optional] [readonly] 
**Serial** | Pointer to **string** | The certificate serial number in hex format. | [optional] [readonly] 
**UsedBy** | Pointer to **string** | Information about the CA certificate usage. | [optional] [readonly] 
**ValidNotAfter** | Pointer to **int64** | The date after which the certificate becomes invalid. | [optional] [readonly] 
**ValidNotBefore** | Pointer to **int64** | The date before which the certificate is not valid. | [optional] [readonly] 
**Result** | Pointer to [**Cacertificate**](Cacertificate.md) |  | [optional] 

## Methods

### NewGetCacertificateResponse

`func NewGetCacertificateResponse() *GetCacertificateResponse`

NewGetCacertificateResponse instantiates a new GetCacertificateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCacertificateResponseWithDefaults

`func NewGetCacertificateResponseWithDefaults() *GetCacertificateResponse`

NewGetCacertificateResponseWithDefaults instantiates a new GetCacertificateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRef

`func (o *GetCacertificateResponse) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *GetCacertificateResponse) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *GetCacertificateResponse) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *GetCacertificateResponse) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *GetCacertificateResponse) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *GetCacertificateResponse) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *GetCacertificateResponse) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *GetCacertificateResponse) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetIssuer

`func (o *GetCacertificateResponse) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *GetCacertificateResponse) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *GetCacertificateResponse) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *GetCacertificateResponse) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSerial

`func (o *GetCacertificateResponse) GetSerial() string`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *GetCacertificateResponse) GetSerialOk() (*string, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *GetCacertificateResponse) SetSerial(v string)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *GetCacertificateResponse) HasSerial() bool`

HasSerial returns a boolean if a field has been set.

### GetUsedBy

`func (o *GetCacertificateResponse) GetUsedBy() string`

GetUsedBy returns the UsedBy field if non-nil, zero value otherwise.

### GetUsedByOk

`func (o *GetCacertificateResponse) GetUsedByOk() (*string, bool)`

GetUsedByOk returns a tuple with the UsedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedBy

`func (o *GetCacertificateResponse) SetUsedBy(v string)`

SetUsedBy sets UsedBy field to given value.

### HasUsedBy

`func (o *GetCacertificateResponse) HasUsedBy() bool`

HasUsedBy returns a boolean if a field has been set.

### GetValidNotAfter

`func (o *GetCacertificateResponse) GetValidNotAfter() int64`

GetValidNotAfter returns the ValidNotAfter field if non-nil, zero value otherwise.

### GetValidNotAfterOk

`func (o *GetCacertificateResponse) GetValidNotAfterOk() (*int64, bool)`

GetValidNotAfterOk returns a tuple with the ValidNotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotAfter

`func (o *GetCacertificateResponse) SetValidNotAfter(v int64)`

SetValidNotAfter sets ValidNotAfter field to given value.

### HasValidNotAfter

`func (o *GetCacertificateResponse) HasValidNotAfter() bool`

HasValidNotAfter returns a boolean if a field has been set.

### GetValidNotBefore

`func (o *GetCacertificateResponse) GetValidNotBefore() int64`

GetValidNotBefore returns the ValidNotBefore field if non-nil, zero value otherwise.

### GetValidNotBeforeOk

`func (o *GetCacertificateResponse) GetValidNotBeforeOk() (*int64, bool)`

GetValidNotBeforeOk returns a tuple with the ValidNotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidNotBefore

`func (o *GetCacertificateResponse) SetValidNotBefore(v int64)`

SetValidNotBefore sets ValidNotBefore field to given value.

### HasValidNotBefore

`func (o *GetCacertificateResponse) HasValidNotBefore() bool`

HasValidNotBefore returns a boolean if a field has been set.

### GetResult

`func (o *GetCacertificateResponse) GetResult() Cacertificate`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetCacertificateResponse) GetResultOk() (*Cacertificate, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetCacertificateResponse) SetResult(v Cacertificate)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetCacertificateResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


